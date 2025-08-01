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
 
// FIXME: Unable to resolve path to module 'vitest/config'
import { defaultExclude } from 'vitest/config';
import { defineConfig } from '@coze-arch/vitest-config';

export default defineConfig({
  preset: 'node',
  dirname: __dirname,
  test: {
    testTimeout: 30 * 1000,
    globals: true,
    mockReset: false,
    coverage: {
      provider: 'v8',
      exclude: ['.eslintrc.js', 'lib', ...defaultExclude],
    },
  },
});
