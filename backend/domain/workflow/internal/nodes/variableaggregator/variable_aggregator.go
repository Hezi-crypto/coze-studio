/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package variableaggregator

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"runtime/debug"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"

	"github.com/coze-dev/coze-studio/backend/domain/workflow/entity/vo"
	"github.com/coze-dev/coze-studio/backend/domain/workflow/internal/nodes"
	"github.com/coze-dev/coze-studio/backend/pkg/ctxcache"
	"github.com/coze-dev/coze-studio/backend/pkg/lang/ptr"
	"github.com/coze-dev/coze-studio/backend/pkg/safego"
)

type MergeStrategy uint

const (
	FirstNotNullValue MergeStrategy = 1
)

type Config struct {
	MergeStrategy MergeStrategy
	GroupLen      map[string]int
	FullSources   map[string]*nodes.SourceInfo
	NodeKey       vo.NodeKey
	InputSources  []*vo.FieldInfo
	GroupOrder    []string // the order the groups are declared in frontend canvas
}

type VariableAggregator struct {
	config *Config
}

func NewVariableAggregator(_ context.Context, cfg *Config) (*VariableAggregator, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}
	if cfg.MergeStrategy != FirstNotNullValue {
		return nil, fmt.Errorf("merge strategy not supported: %v", cfg.MergeStrategy)
	}
	return &VariableAggregator{config: cfg}, nil
}

func (v *VariableAggregator) Invoke(ctx context.Context, input map[string]any) (_ map[string]any, err error) {
	in, err := inputConverter(input)
	if err != nil {
		return nil, err
	}

	result := make(map[string]any)
	groupToChoice := make(map[string]int)
	for group, length := range v.config.GroupLen {
		for i := 0; i < length; i++ {
			if value, ok := in[group][i]; ok {
				if value != nil {
					result[group] = value
					groupToChoice[group] = i
					break
				}
			}
		}

		if _, ok := result[group]; !ok {
			groupToChoice[group] = -1
		}
	}

	_ = compose.ProcessState(ctx, func(ctx context.Context, state nodes.DynamicStreamContainer) error {
		state.SaveDynamicChoice(v.config.NodeKey, groupToChoice)
		return nil
	})

	ctxcache.Store(ctx, groupChoiceTypeCacheKey, map[string]nodes.FieldStreamType{}) // none of the choices are stream

	groupChoices := make([]any, 0, len(v.config.GroupOrder))
	for _, group := range v.config.GroupOrder {
		choice := groupToChoice[group]
		if choice == -1 {
			groupChoices = append(groupChoices, nil)
		} else {
			groupChoices = append(groupChoices, choice)
		}
	}

	ctxcache.Store(ctx, groupChoiceCacheKey, groupChoices)

	return result, nil
}

const (
	resolvedSourcesCacheKey = "resolved_sources"
	groupChoiceTypeCacheKey = "group_choice_type"
	groupChoiceCacheKey     = "group_choice"
)

// Transform picks the first non-nil value from each group from a stream of map[group]items.
func (v *VariableAggregator) Transform(ctx context.Context, input *schema.StreamReader[map[string]any]) (
	_ *schema.StreamReader[map[string]any], err error) {
	inStream := streamInputConverter(input)

	resolvedSources, ok := ctxcache.Get[map[string]*nodes.SourceInfo](ctx, resolvedSourcesCacheKey)
	if !ok {
		panic("unable to get resolvesSources from ctx cache.")
	}

	groupToItems := make(map[string][]any)
	groupToChoice := make(map[string]int)
	type skipped struct{}
	type null struct{}
	type stream struct{}

	defer func() {
		if err == nil {
			groupChoiceToStreamType := map[string]nodes.FieldStreamType{}
			for group, choice := range groupToChoice {
				if choice != -1 {
					item := groupToItems[group][choice]
					if _, ok := item.(stream); ok {
						groupChoiceToStreamType[group] = nodes.FieldIsStream
					}
				}
			}

			groupChoices := make([]any, 0, len(v.config.GroupOrder))
			for _, group := range v.config.GroupOrder {
				choice := groupToChoice[group]
				if choice == -1 {
					groupChoices = append(groupChoices, nil)
				} else {
					groupChoices = append(groupChoices, choice)
				}
			}

			// store group -> field type for use in callbacks.OnEnd
			ctxcache.Store(ctx, groupChoiceTypeCacheKey, groupChoiceToStreamType)
			ctxcache.Store(ctx, groupChoiceCacheKey, groupChoices)
		}
	}()

	// goal: find the first non-nil element in each group. 'First' means the smallest index in each group's slice.
	// For a stream element, if the stream source is not skipped, then this stream element is non-nil,
	// even if there's no content in the stream.
	// steps:
	// - for each group, iterate over each element in order, check the element's stream type
	// - if an element is skipped, move on
	// - if an element is stream, pick it
	// - if an element is not stream, actually receive from the stream to check if it's non-nil

	groupToCurrentIndex := make(map[string]int) // the currently known smallest index that is non-nil for each group
	for group, length := range v.config.GroupLen {
		groupToItems[group] = make([]any, length)
		groupToCurrentIndex[group] = math.MaxInt
		for i := 0; i < length; i++ {
			fType := resolvedSources[group].SubSources[strconv.Itoa(i)].FieldType
			if fType == nodes.FieldSkipped {
				groupToItems[group][i] = skipped{}
				continue
			}
			if fType == nodes.FieldIsStream {
				groupToItems[group][i] = stream{}
				if ci, _ := groupToCurrentIndex[group]; i < ci {
					groupToCurrentIndex[group] = i
				}
			}
		}

		hasUndecided := false
		for i := 0; i < length; i++ {
			if groupToItems[group][i] == nil {
				hasUndecided = true
				break
			}

			_, ok := groupToItems[group][i].(stream)
			if ok { // if none of the elements before this one is none-stream, pick this first stream
				groupToChoice[group] = i
				break
			}
		}

		if _, ok := groupToChoice[group]; !ok && !hasUndecided {
			groupToChoice[group] = -1 // all of this group's elements are skipped, won't have any non-nil ones
		}
	}

	allDone := func() bool {
		for group := range v.config.GroupLen {
			_, ok := groupToChoice[group]
			if !ok {
				return false
			}
		}

		return true
	}

	alreadyDone := allDone()
	if alreadyDone { // all groups have made their choices, no need to actually read input streams
		result := make(map[string]any, len(v.config.GroupLen))
		allSkip := true
		for group := range groupToChoice {
			choice := groupToChoice[group]
			if choice == -1 {
				result[group] = nil // all elements of this group are skipped
			} else {
				result[group] = choice
				allSkip = false
			}
		}

		if allSkip { // no need to convert input streams for the output, because all groups are skipped
			_ = compose.ProcessState(ctx, func(ctx context.Context, state nodes.DynamicStreamContainer) error {
				state.SaveDynamicChoice(v.config.NodeKey, groupToChoice)
				return nil
			})
			return schema.StreamReaderFromArray([]map[string]any{result}), nil
		}
	}

	outS := inStream
	if !alreadyDone {
		inCopy := inStream.Copy(2)
		defer inCopy[0].Close()
		outS = inCopy[1]

	recvLoop:
		for {
			chunk, err := inCopy[0].Recv()
			if err != nil {
				if err == io.EOF {
					panic("EOF reached before making choices for all groups")
				}

				return nil, err
			}

			for group, items := range chunk {
				if _, ok := groupToChoice[group]; ok {
					continue // already made the decision for the group.
				}

				for i := range items {
					if i >= groupToCurrentIndex[group] {
						continue
					}

					existing := groupToItems[group][i]
					if existing != nil { // belongs to a stream element
						continue
					}

					// now the item is always a non-stream element
					item := items[i]
					if item == nil {
						groupToItems[group][i] = null{}
					} else {
						groupToItems[group][i] = item
					}

					groupToCurrentIndex[group] = i

					finalized := true
					for j := 0; j < i; j++ {
						indexedItem := groupToItems[group][j]
						if indexedItem == nil { // there exists non-finalized elements in front of the current item
							finalized = false
							break
						}
					}

					if finalized {
						if item == nil { // current item is nil, we need to find the first non-nil element in the group
							foundNonNil := false
							hasUndecided := false
							for j := 0; j < len(groupToItems[group]); j++ {
								indexedItem := groupToItems[group][j]
								if indexedItem != nil {
									_, ok := indexedItem.(skipped)
									if ok {
										continue
									}

									_, ok = indexedItem.(null)
									if ok {
										continue
									}

									groupToChoice[group] = j
									foundNonNil = true
									break
								} else {
									hasUndecided = true
									break
								}
							}
							if !foundNonNil && !hasUndecided {
								groupToChoice[group] = -1 // this group does not have any non-nil value
							}
						} else {
							groupToChoice[group] = i
						}
						if allDone() {
							break recvLoop
						}
					}
				}
			}
		}
	}

	_ = compose.ProcessState(ctx, func(ctx context.Context, state nodes.DynamicStreamContainer) error {
		state.SaveDynamicChoice(v.config.NodeKey, groupToChoice)
		return nil
	})

	actualStream := schema.StreamReaderWithConvert(outS, func(in map[string]map[int]any) (map[string]any, error) {
		out := make(map[string]any)
		for group, items := range in {
			choice, ok := groupToChoice[group]
			if !ok {
				panic(fmt.Sprintf("group %s does not have choice", group))
			}

			if choice < 0 {
				panic(fmt.Sprintf("group %s choice = %d, less than zero, but found actual item in stream", group, choice))
			}

			if _, ok := items[choice]; ok {
				out[group] = items[choice]
			}
		}

		if len(out) == 0 {
			return nil, schema.ErrNoValue
		}

		return out, nil
	})

	nullGroups := make(map[string]any)
	for group, choice := range groupToChoice {
		if choice < 0 {
			nullGroups[group] = nil
		}
	}
	if len(nullGroups) > 0 {
		nullStream := schema.StreamReaderFromArray([]map[string]any{nullGroups})
		return schema.MergeStreamReaders([]*schema.StreamReader[map[string]any]{actualStream, nullStream}), nil
	}

	return actualStream, nil
}

func inputConverter(in map[string]any) (converted map[string]map[int]any, err error) {
	converted = make(map[string]map[int]any)

	for k, value := range in {
		m, ok := value.(map[string]any)
		if !ok {
			return nil, errors.New("value is not a map[string]any")
		}
		converted[k] = make(map[int]any, len(m))
		for i, sv := range m {
			index, err := strconv.Atoi(i)
			if err != nil {
				return nil, fmt.Errorf(" converting %s to int failed, err=%v", i, err)
			}
			converted[k][index] = sv
		}
	}

	return converted, nil
}

func streamInputConverter(in *schema.StreamReader[map[string]any]) *schema.StreamReader[map[string]map[int]any] {
	converter := func(input map[string]any) (output map[string]map[int]any, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = safego.NewPanicErr(r, debug.Stack())
			}
		}()
		return inputConverter(input)
	}
	return schema.StreamReaderWithConvert(in, converter)
}

type vaCallbackInput struct {
	Name      string `json:"name"`
	Variables []any  `json:"variables"`
}

func (v *VariableAggregator) Init(ctx context.Context) (context.Context, error) {
	ctx = ctxcache.Init(ctx)

	resolvedSources, err := nodes.ResolveStreamSources(ctx, v.config.FullSources)
	if err != nil {
		return nil, err
	}

	// need this info for callbacks.OnStart, so we put it in cache within Init()
	ctxcache.Store(ctx, resolvedSourcesCacheKey, resolvedSources)

	return ctx, nil
}

type streamMarkerType string

const streamMarker streamMarkerType = "<Stream Data...>"

func (v *VariableAggregator) ToCallbackInput(ctx context.Context, input map[string]any) (map[string]any, error) {
	resolvedSources, ok := ctxcache.Get[map[string]*nodes.SourceInfo](ctx, resolvedSourcesCacheKey)
	if !ok {
		panic("unable to get resolved_sources from ctx cache")
	}

	in, err := inputConverter(input)
	if err != nil {
		return nil, err
	}

	merged := make([]vaCallbackInput, 0, len(in))

	groupLen := v.config.GroupLen

	for groupName, vars := range in {
		orderedVars := make([]any, groupLen[groupName])
		for index := range vars {
			orderedVars[index] = vars[index]
			if len(resolvedSources) > 0 {
				if resolvedSources[groupName].SubSources[strconv.Itoa(index)].FieldType == nodes.FieldIsStream {
					// replace the streams with streamMarker,
					// because we won't read, save to execution history, or display these streams to user
					orderedVars[index] = streamMarker
				}
			}
		}

		merged = append(merged, vaCallbackInput{
			Name:      groupName,
			Variables: orderedVars,
		})
	}

	// Sort merged slice by Name
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].Name < merged[j].Name
	})

	return map[string]any{
		"mergeGroups": merged,
	}, nil
}

func (v *VariableAggregator) ToCallbackOutput(ctx context.Context, output map[string]any) (*nodes.StructuredCallbackOutput, error) {
	dynamicStreamType, ok := ctxcache.Get[map[string]nodes.FieldStreamType](ctx, groupChoiceTypeCacheKey)
	if !ok {
		panic("unable to get dynamic stream types from ctx cache")
	}

	groupChoices, ok := ctxcache.Get[[]any](ctx, groupChoiceCacheKey)
	if !ok {
		panic("unable to get group choices from ctx cache")
	}

	if len(dynamicStreamType) == 0 {
		return &nodes.StructuredCallbackOutput{
			Output:    output,
			RawOutput: output,
			Extra: map[string]any{
				"variable_select": groupChoices,
			},
		}, nil
	}

	newOut := maps.Clone(output)
	for k := range output {
		if t, ok := dynamicStreamType[k]; ok && t == nodes.FieldIsStream {
			newOut[k] = streamMarker
		}
	}

	return &nodes.StructuredCallbackOutput{
		Output:    newOut,
		RawOutput: newOut,
		Extra: map[string]any{
			"variable_select": groupChoices,
		},
	}, nil
}

func concatVACallbackInputs(vs [][]vaCallbackInput) ([]vaCallbackInput, error) {
	if len(vs) == 0 {
		return nil, nil
	}

	init := slices.Clone(vs[0])
	for i := 1; i < len(vs); i++ {
		next := vs[i]
		for j := 0; j < len(next); j++ {
			oneGroup := next[j]
			groupName := oneGroup.Name
			var (
				existingGroup *vaCallbackInput
				nextIndex     = len(init)
				currentIndex  int
			)
			for k := 0; k < len(init); k++ {
				if init[k].Name == groupName {
					existingGroup = ptr.Of(init[k])
					currentIndex = k
				} else if init[k].Name > groupName && k < nextIndex {
					nextIndex = k
				}
			}

			if existingGroup == nil {
				after := slices.Clone(init[nextIndex:])
				init = append(init[:nextIndex], oneGroup)
				init = append(init, after...)
			} else {
				for vi := 0; vi < len(oneGroup.Variables); vi++ {
					newV := oneGroup.Variables[vi]
					if newV == nil {
						if vi >= len(existingGroup.Variables) {
							for i := len(existingGroup.Variables); i <= vi; i++ {
								existingGroup.Variables = append(existingGroup.Variables, nil)
							}
						}
						continue
					}
					if newStr, ok := newV.(string); ok {
						if strings.HasSuffix(newStr, nodes.KeyIsFinished) {
							newStr = strings.TrimSuffix(newStr, nodes.KeyIsFinished)
						}
						newV = newStr
					}
					for ei := len(existingGroup.Variables); ei <= vi; ei++ {
						existingGroup.Variables = append(existingGroup.Variables, nil)
					}
					ev := existingGroup.Variables[vi]
					if ev == nil {
						existingGroup.Variables[vi] = oneGroup.Variables[vi]
					} else {
						if evStr, ok := ev.(streamMarkerType); !ok {
							return nil, fmt.Errorf("multiple stream chunk when concating VACallbackInputs, variable %s is not string", ev)
						} else {
							if evStr != streamMarker || newV.(streamMarkerType) != streamMarker {
								return nil, fmt.Errorf("multiple stream chunk when concating VACallbackInputs, variable %s is not streamMarker", ev)
							}
							existingGroup.Variables[vi] = evStr
						}
					}
				}
				init[currentIndex] = *existingGroup
			}
		}
	}

	return init, nil
}

func concatStreamMarkers(_ []streamMarkerType) (streamMarkerType, error) {
	return streamMarker, nil
}

func init() {
	nodes.RegisterStreamChunkConcatFunc(concatVACallbackInputs)
	nodes.RegisterStreamChunkConcatFunc(concatStreamMarkers)
}
