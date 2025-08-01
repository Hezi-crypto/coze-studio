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

export type Int64 = string | number;

export enum GroupStatus {
  Enable = 1,
  Disable = 2,
}

export enum GroupType {
  Normal = 1,
  Continue = 2,
}

export enum PeriodType {
  None = 1,
  Day = 2,
  Week = 3,
  Month = 4,
}

export enum ReceiveType {
  Auto = 1,
  Manual = 2,
}

export enum RewardType {
  Token = 1,
  Coin = 2,
}

export enum TaskStatus {
  Enable = 1,
  Disable = 2,
}

export enum UserControl {
  Disable = 0,
  BlackList = 1,
  WhiteList = 2,
}

export interface Reward {
  type: RewardType;
  amount: Int64;
  receive_type?: ReceiveType;
}

export interface UserTaskButton {
  linked?: string;
  desc?: string;
}
/* eslint-enable */
