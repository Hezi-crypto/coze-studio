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
 
// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
/* eslint-disable */
/* tslint:disable */
// @ts-nocheck

import * as api from './namespaces/api';
import * as common from './namespaces/common';
import * as rpc from './namespaces/rpc';

export { api, common, rpc };
export * from './namespaces/api';
export * from './namespaces/common';
export * from './namespaces/rpc';

export type Int64 = string | number;

export default class IncentiveOperationService<T> {
  private request: any = () => {
    throw new Error('IncentiveOperationService.request is undefined');
  };
  private baseURL: string | ((path: string) => string) = '';

  constructor(options?: {
    baseURL?: string | ((path: string) => string);
    request?<R>(
      params: {
        url: string;
        method: 'GET' | 'DELETE' | 'POST' | 'PUT' | 'PATCH';
        data?: any;
        params?: any;
        headers?: any;
      },
      options?: T,
    ): Promise<R>;
  }) {
    this.request = options?.request || this.request;
    this.baseURL = options?.baseURL || '';
  }

  private genBaseURL(path: string) {
    return typeof this.baseURL === 'string'
      ? this.baseURL + path
      : this.baseURL(path);
  }

  /**
   * GET /api/marketplace/incentive/user_task/list
   *
   * 任务中心展示用户任务
   */
  ListUserTask(
    req: api.ListUserTaskRequest,
    options?: T,
  ): Promise<api.ListUserTaskResponse> {
    const _req = req;
    const url = this.genBaseURL('/api/marketplace/incentive/user_task/list');
    const method = 'GET';
    const params = {
      page_size: _req['page_size'],
      page_token: _req['page_token'],
    };
    return this.request({ url, method, params }, options);
  }

  /**
   * GET /api/marketplace/admin/incentive/task/list
   *
   * 任务管理, 任务查询
   */
  ListTask(
    req: api.ListTaskRequest,
    options?: T,
  ): Promise<api.ListTaskResponse> {
    const _req = req;
    const url = this.genBaseURL('/api/marketplace/admin/incentive/task/list');
    const method = 'GET';
    const params = {
      page: _req['page'],
      size: _req['size'],
      task_id: _req['task_id'],
      task_status: _req['task_status'],
      task_name: _req['task_name'],
    };
    return this.request({ url, method, params }, options);
  }

  /**
   * POST /api/marketplace/admin/incentive/task/update
   *
   * 任务管理, 任务修改
   */
  UpdateTask(
    req: api.UpdateTaskRequest,
    options?: T,
  ): Promise<api.UpdateTaskResponse> {
    const _req = req;
    const url = this.genBaseURL('/api/marketplace/admin/incentive/task/update');
    const method = 'POST';
    const data = {
      task_id: _req['task_id'],
      bill_desc_starling_key: _req['bill_desc_starling_key'],
      user_task_desc_starling_key: _req['user_task_desc_starling_key'],
      button: _req['button'],
      reward: _req['reward'],
      times_limit: _req['times_limit'],
      limit_restriction: _req['limit_restriction'],
      peroid_type: _req['peroid_type'],
      peroid_value: _req['peroid_value'],
      hide: _req['hide'],
      drop_when_finish: _req['drop_when_finish'],
      task_desc: _req['task_desc'],
      event_trigger_key: _req['event_trigger_key'],
      user_control: _req['user_control'],
      black_list: _req['black_list'],
      white_list: _req['white_list'],
    };
    return this.request({ url, method, data }, options);
  }

  /**
   * POST /api/marketplace/admin/incentive/task/status/update
   *
   * 任务管理, 任务上架/下架
   */
  UpdateTaskStatus(
    req: api.UpdateTaskStatusRequest,
    options?: T,
  ): Promise<api.UpdateTaskStatusResponse> {
    const _req = req;
    const url = this.genBaseURL(
      '/api/marketplace/admin/incentive/task/status/update',
    );
    const method = 'POST';
    const data = { task_id: _req['task_id'], task_status: _req['task_status'] };
    return this.request({ url, method, data }, options);
  }
}
/* eslint-enable */
