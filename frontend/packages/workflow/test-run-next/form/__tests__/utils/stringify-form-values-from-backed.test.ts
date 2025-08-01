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
 
import { describe, it, expect } from 'vitest';

import { stringifyFormValuesFromBacked } from '../../src/utils/stringify-form-values-from-backed';

describe('stringifyFormValuesFromBacked', () => {
  // 测试输入为空的情况
  it('should return undefined when input is null or undefined', () => {
    expect(stringifyFormValuesFromBacked(null as any)).toBeUndefined();
    expect(stringifyFormValuesFromBacked(undefined as any)).toBeUndefined();
  });

  // 测试输入包含字符串和布尔值的情况
  it('should return the same string and boolean values', () => {
    const input = {
      str: 'hello',
      bool: true,
    };
    const result = stringifyFormValuesFromBacked(input);
    expect(result).toEqual({
      str: 'hello',
      bool: true,
    });
  });

  // 测试输入包含对象和数组的情况
  it('should stringify objects and arrays', () => {
    const input = {
      obj: { key: 'value' },
      arr: [1, 2, 3],
    };
    const result = stringifyFormValuesFromBacked(input);
    expect(result).toEqual({
      obj: '{"key":"value"}',
      arr: '[1,2,3]',
    });
  });

  // 测试输入包含 null 和 undefined 的情况
  it('should set null and undefined values to undefined in the result', () => {
    const input = {
      nullValue: null,
      undefinedValue: undefined,
    };
    const result = stringifyFormValuesFromBacked(input);
    expect(result).toEqual({
      nullValue: undefined,
      undefinedValue: undefined,
    });
  });
});
