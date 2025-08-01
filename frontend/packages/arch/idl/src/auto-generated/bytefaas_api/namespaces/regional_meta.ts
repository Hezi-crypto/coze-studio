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

import * as common from './common';
import * as admin from './admin';
import * as ticket from './ticket';

export type Int64 = string | number;

export interface GetRegionalMetaRequest {
  /** cluster name */
  cluster: string;
  /** region name */
  region: string;
  /** ID of service */
  service_id: string;
  /** jwt token */
  'X-Jwt-Token'?: string;
}

export interface GetRegionalMetaResponse {
  code?: number;
  data?: RegionalMetaResponseData;
  error?: string;
}

export interface RegionalMeta {
  /** Region name */
  region?: string;
  /** Function ID */
  function_id?: string;
  /** Service ID */
  service_id?: string;
  /** Cluster name */
  cluster?: string;
  /** Function name */
  function_name?: string;
  /** Revision ID */
  revision_id?: string;
  /** Owner of the function */
  owner?: string;
  /** PSM name */
  psm?: string;
  /** Cell name */
  cell?: string;
  /** Disabled status for each zone */
  is_this_zone_disabled?: Record<string, boolean>;
  /** Throttle log bytes per second for each zone */
  zone_throttle_log_bytes_per_sec?: Record<string, number>;
  /** Enable GDPR compliance */
  gdpr_enable?: boolean;
  /** Disable cold start, ture means disable */
  cold_start_disabled?: boolean;
  /** Enable exclusive mode */
  exclusive_mode?: boolean;
  /** Enable async mode */
  async_mode?: boolean;
  /** Enable online mode */
  online_mode?: boolean;
  /** Enable authentication */
  auth_enable?: boolean;
  /** Enable CORS */
  cors_enable?: boolean;
  /** Enable tracing */
  trace_enable?: boolean;
  /** Enable gateway route */
  gateway_route_enable?: boolean;
  /** Enable IPv6 only */
  is_ipv6_only?: boolean;
  /** Enable ZTI */
  zti_enable?: boolean;
  /** Disable HTTP trigger，true means disable */
  http_trigger_disable?: boolean;
  /** Aliases for the function */
  aliases?: Record<string, common.Alias>;
  /** Runtime environment */
  runtime?: string;
  /** Environment name */
  env_name?: string;
  /** List of global KV namespace IDs */
  global_kv_namespace_ids?: Array<string>;
  /** List of local cache namespace IDs */
  local_cache_namespace_ids?: Array<string>;
  /** Protocol type */
  protocol?: string;
  /** Latency in seconds */
  latency_sec?: number;
  /** Network class ID */
  net_class_id?: number;
  /** Environment variables (nested map) */
  envs?: Record<string, Record<string, string>>;
  /** In releasing state */
  in_releasing?: boolean;
  /** Enable reserved DP */
  reserved_dp_enabled?: boolean;
  /** Routing strategy, enums：prefer_reserved, prefer_elastic. */
  routing_strategy?: string;
  /** Disable ByteFaaS error response */
  bytefaas_error_response_disabled?: boolean;
  /** Disable ByteFaaS response header */
  bytefaas_response_header_disabled?: boolean;
  /** Enable colocate scheduling */
  enable_colocate_scheduling?: boolean;
  /** Network mode */
  network_mode?: string;
  /** Enable dynamic load balancing data report */
  dynamic_load_balancing_data_report_enabled?: boolean;
  /** Enable dynamic load balancing weight */
  dynamic_load_balancing_weight_enabled?: boolean;
  /** List of enabled VDCs for dynamic load balancing */
  dynamic_load_balancing_enabled_vdcs?: Array<string>;
  /** Type of dynamic load balancing, enums: round_robin, smooth_weighted_round_robin, weighted_round_robin */
  dynamic_load_balance_type?: string;
  /** Deployment inactive status */
  deployment_inactive?: boolean;
  /** Deployment inactive status for each zone */
  is_this_zone_deployment_inactive?: Record<string, boolean>;
  /** Package name */
  package?: string;
  /** Pod type */
  pod_type?: string;
  /** Plugin name */
  plugin_name?: string;
  /** Allow cold start instance */
  allow_cold_start_instance?: boolean;
  /** Elastic preferred cluster mapping */
  elastic_prefer_cluster?: Record<string, string>;
  /** Reserved preferred cluster mapping */
  reserved_prefer_cluster?: Record<string, string>;
  /** Elastic user preferred cluster mapping */
  elastic_user_preferred_cluster?: Record<string, string>;
  /** Reserved user preferred cluster mapping */
  reserved_user_preferred_cluster?: Record<string, string>;
  /** Elastic auto preferred cluster mapping */
  elastic_auto_preferred_cluster?: Record<string, string>;
  /** Reserved auto preferred cluster mapping */
  reserved_auto_preferred_cluster?: Record<string, string>;
  /** Temporary preferred cluster mapping */
  temp_preferred_cluster?: Record<string, string>;
  /** Formatted elastic preferred cluster list */
  formatted_elastic_prefer_cluster?: Array<common.FormattedPreferCluster>;
  /** Formatted reserved preferred cluster list */
  formatted_reserved_prefer_cluster?: Array<common.FormattedPreferCluster>;
  /** Runtime log writers */
  runtime_log_writers?: string;
  /** System log writers */
  system_log_writers?: string;
  /** Enable runtime JSON log */
  enable_runtime_json_log?: boolean;
  /** Is BytePaaS elastic cluster */
  is_bytepaas_elastic_cluster?: boolean;
  /** Disable service discovery，true means disable */
  disable_service_discovery?: boolean;
  /** Enable resource guarantee */
  resource_guarantee?: boolean;
  /** Disable cgroup v2 */
  disable_cgroup_v2?: boolean;
  /** Enable async result emit event bridge */
  async_result_emit_event_bridge?: boolean;
  /** Runtime stream log bytes per second */
  runtime_stream_log_bytes_per_sec?: number;
  /** System stream log bytes per second */
  system_stream_log_bytes_per_sec?: number;
  /** Throttle log bytes per second */
  throttle_log_bytes_per_sec?: number;
  /** Throttle stdout log bytes per second */
  throttle_stdout_log_bytes_per_sec?: number;
  /** Throttle stderr log bytes per second */
  throttle_stderr_log_bytes_per_sec?: number;
  /** Enable scaling */
  scale_enabled?: boolean;
  /** Scale threshold */
  scale_threshold?: number;
  /** Scale type */
  scale_type?: number;
  /** Scale settings */
  scale_settings?: admin.FuncScaleSettings;
  /** Pod replica limits */
  replica_limit?: Record<string, common.PodReplicaLimit>;
  /** Reserved frozen replicas for each zone */
  zone_reserved_frozen_replicas?: Record<string, number>;
  /** Container runtime */
  container_runtime?: string;
  /** Enable scale optimization */
  enable_scale_optimise?: boolean;
  /** Schedule strategy */
  schedule_strategy?: string;
  /** Dynamic overcommit settings */
  dynamic_overcommit_settings?: Record<
    string,
    ticket.DynamicOvercommitSettings
  >;
  /** Enable overload protection */
  overload_protect_enabled?: boolean;
  /** Frozen CPU milli value */
  frozen_cpu_milli?: number;
  /** Enable federated on-demand resource */
  enable_fed_on_demand_resource?: Record<string, boolean>;
  /** Frozen priority class */
  frozen_priority_class?: string;
  /** Host-unique scheduling config */
  host_uniq?: common.HostUniq;
  /** In cell migration state */
  in_cell_migration?: boolean;
  /** Canary replica limit for each zone */
  zone_canary_replica_limit?: Record<string, number>;
  /** Enable frozen scale */
  frozen_scale_enabled?: boolean;
  /** enable privileged in pod */
  privileged?: boolean;
}

/** Regional metadata response data */
export interface RegionalMetaResponseData {
  /** Region name */
  region?: string;
  /** Function ID */
  function_id?: string;
  /** Service ID */
  service_id?: string;
  /** Cluster name */
  cluster?: string;
  /** Function name */
  function_name?: string;
  /** Revision ID */
  revision_id?: string;
  /** Owner of the function */
  owner?: string;
  /** PSM name */
  psm?: string;
  /** Cell name */
  cell?: string;
  /** Disabled status for each zone */
  is_this_zone_disabled?: Record<string, boolean>;
  /** Throttle log bytes per second for each zone */
  zone_throttle_log_bytes_per_sec?: Record<string, number>;
  /** Enable GDPR compliance */
  gdpr_enable?: boolean;
  /** Disable cold start, ture means disable */
  cold_start_disabled?: boolean;
  /** Enable exclusive mode */
  exclusive_mode?: boolean;
  /** Enable async mode */
  async_mode?: boolean;
  /** Enable online mode */
  online_mode?: boolean;
  /** Enable authentication */
  auth_enable?: boolean;
  /** Enable CORS */
  cors_enable?: boolean;
  /** Enable tracing */
  trace_enable?: boolean;
  /** Enable gateway route */
  gateway_route_enable?: boolean;
  /** Enable IPv6 only */
  is_ipv6_only?: boolean;
  /** Enable ZTI */
  zti_enable?: boolean;
  /** Disable HTTP trigger，true means disable */
  http_trigger_disable?: boolean;
  /** Aliases for the function */
  aliases?: Record<string, common.Alias>;
  /** Runtime environment */
  runtime?: string;
  /** Environment name */
  env_name?: string;
  /** List of global KV namespace IDs */
  global_kv_namespace_ids?: Array<string>;
  /** List of local cache namespace IDs */
  local_cache_namespace_ids?: Array<string>;
  /** Protocol type */
  protocol?: string;
  /** Latency in seconds */
  latency_sec?: number;
  /** Network class ID */
  net_class_id?: number;
  /** Environment variables (nested map) */
  envs?: Record<string, Record<string, string>>;
  /** In releasing state */
  in_releasing?: boolean;
  /** Enable reserved DP */
  reserved_dp_enabled?: boolean;
  /** Routing strategy, enums：prefer_reserved, prefer_elastic. */
  routing_strategy?: string;
  /** Disable ByteFaaS error response */
  bytefaas_error_response_disabled?: boolean;
  /** Disable ByteFaaS response header */
  bytefaas_response_header_disabled?: boolean;
  /** Enable colocate scheduling */
  enable_colocate_scheduling?: boolean;
  /** Network mode */
  network_mode?: string;
  /** Enable dynamic load balancing data report */
  dynamic_load_balancing_data_report_enabled?: boolean;
  /** Enable dynamic load balancing weight */
  dynamic_load_balancing_weight_enabled?: boolean;
  /** List of enabled VDCs for dynamic load balancing */
  dynamic_load_balancing_enabled_vdcs?: Array<string>;
  /** Type of dynamic load balancing, enums: round_robin, smooth_weighted_round_robin, weighted_round_robin */
  dynamic_load_balance_type?: string;
  /** Deployment inactive status */
  deployment_inactive?: boolean;
  /** Deployment inactive status for each zone */
  is_this_zone_deployment_inactive?: Record<string, boolean>;
  /** Package name */
  package?: string;
  /** Pod type */
  pod_type?: string;
  /** Plugin name */
  plugin_name?: string;
  /** Allow cold start instance */
  allow_cold_start_instance?: boolean;
  /** Elastic preferred cluster mapping */
  elastic_prefer_cluster?: Record<string, string>;
  /** Reserved preferred cluster mapping */
  reserved_prefer_cluster?: Record<string, string>;
  /** Elastic user preferred cluster mapping */
  elastic_user_preferred_cluster?: Record<string, string>;
  /** Reserved user preferred cluster mapping */
  reserved_user_preferred_cluster?: Record<string, string>;
  /** Elastic auto preferred cluster mapping */
  elastic_auto_preferred_cluster?: Record<string, string>;
  /** Reserved auto preferred cluster mapping */
  reserved_auto_preferred_cluster?: Record<string, string>;
  /** Temporary preferred cluster mapping */
  temp_preferred_cluster?: Record<string, string>;
  /** Formatted elastic preferred cluster list */
  formatted_elastic_prefer_cluster?: Array<common.FormattedPreferCluster>;
  /** Formatted reserved preferred cluster list */
  formatted_reserved_prefer_cluster?: Array<common.FormattedPreferCluster>;
  /** Runtime log writers */
  runtime_log_writers?: string;
  /** System log writers */
  system_log_writers?: string;
  /** Enable runtime JSON log */
  enable_runtime_json_log?: boolean;
  /** Is BytePaaS elastic cluster */
  is_bytepaas_elastic_cluster?: boolean;
  /** Disable service discovery，true means disable */
  disable_service_discovery?: boolean;
  /** Enable resource guarantee */
  resource_guarantee?: boolean;
  /** Disable cgroup v2 */
  disable_cgroup_v2?: boolean;
  /** Enable async result emit event bridge */
  async_result_emit_event_bridge?: boolean;
  /** Runtime stream log bytes per second */
  runtime_stream_log_bytes_per_sec?: number;
  /** System stream log bytes per second */
  system_stream_log_bytes_per_sec?: number;
  /** Throttle log bytes per second */
  throttle_log_bytes_per_sec?: number;
  /** Throttle stdout log bytes per second */
  throttle_stdout_log_bytes_per_sec?: number;
  /** Throttle stderr log bytes per second */
  throttle_stderr_log_bytes_per_sec?: number;
  /** Enable scaling */
  scale_enabled?: boolean;
  /** Scale threshold */
  scale_threshold?: number;
  /** Scale type */
  scale_type?: number;
  /** Scale settings */
  scale_settings?: admin.FuncScaleSettings;
  /** Pod replica limits */
  replica_limit?: Record<string, common.PodReplicaLimit>;
  /** Reserved frozen replicas for each zone */
  zone_reserved_frozen_replicas?: Record<string, number>;
  /** Container runtime */
  container_runtime?: string;
  /** Enable scale optimization */
  enable_scale_optimise?: boolean;
  /** Schedule strategy */
  schedule_strategy?: string;
  /** Dynamic overcommit settings */
  dynamic_overcommit_settings?: Record<
    string,
    ticket.DynamicOvercommitSettings
  >;
  /** Enable overload protection */
  overload_protect_enabled?: boolean;
  /** Frozen CPU milli value */
  frozen_cpu_milli?: number;
  /** Enable federated on-demand resource */
  enable_fed_on_demand_resource?: Record<string, boolean>;
  /** Frozen priority class */
  frozen_priority_class?: string;
  /** Enable frozen scale */
  frozen_scale_enabled?: boolean;
  /** Log link */
  log_link?: string;
  /** Stream log link */
  stream_log_link?: string;
  /** Argos link */
  argos_link?: string;
  /** Grafana link */
  grafana_link?: string;
  /** List of metrics links */
  metrics_links?: Array<string>;
  /** Host unique identifier */
  host_uniq?: common.HostUniq;
  /** In cell migration state */
  in_cell_migration?: boolean;
  /** Canary replica limit for each zone */
  zone_canary_replica_limit?: Record<string, number>;
  /** enable privileged in pod */
  privileged?: boolean;
}

export interface UpdateRegionalMetaRequest {
  /** Function ID */
  function_id?: string;
  /** Function name */
  function_name?: string;
  /** Revision ID */
  revision_id?: string;
  /** Owner of the function */
  owner?: string;
  /** PSM name */
  psm?: string;
  /** Cell name */
  cell?: string;
  /** Disabled status for each zone */
  is_this_zone_disabled?: Record<string, boolean>;
  /** Throttle log bytes per second for each zone */
  zone_throttle_log_bytes_per_sec?: Record<string, number>;
  /** Enable GDPR compliance */
  gdpr_enable?: boolean;
  /** Disable cold start */
  cold_start_disabled?: boolean;
  /** Enable exclusive mode */
  exclusive_mode?: boolean;
  /** Enable async mode */
  async_mode?: boolean;
  /** Enable online mode */
  online_mode?: boolean;
  /** Enable authentication */
  auth_enable?: boolean;
  /** Enable CORS */
  cors_enable?: boolean;
  /** Enable tracing */
  trace_enable?: boolean;
  /** Enable gateway route */
  gateway_route_enable?: boolean;
  /** Enable IPv6 only */
  is_ipv6_only?: boolean;
  /** Enable ZTI */
  zti_enable?: boolean;
  /** Disable HTTP trigger，true means disable */
  http_trigger_disable?: boolean;
  /** Aliases for the function */
  aliases?: Record<string, common.Alias>;
  /** Runtime environment */
  runtime?: string;
  /** Environment name */
  env_name?: string;
  /** List of global KV namespace IDs */
  global_kv_namespace_ids?: Array<string>;
  /** List of local cache namespace IDs */
  local_cache_namespace_ids?: Array<string>;
  /** Protocol type */
  protocol?: string;
  /** Latency in seconds */
  latency_sec?: number;
  /** Network class ID */
  net_class_id?: number;
  /** Environment variables (nested map) */
  envs?: Record<string, Record<string, string>>;
  /** In releasing state */
  in_releasing?: boolean;
  /** Enable reserved DP */
  reserved_dp_enabled?: boolean;
  /** Routing strategy, enums：prefer_reserved, prefer_elastic. */
  routing_strategy?: string;
  /** Disable ByteFaaS error response */
  bytefaas_error_response_disabled?: boolean;
  /** Disable ByteFaaS response header */
  bytefaas_response_header_disabled?: boolean;
  /** Enable colocate scheduling */
  enable_colocate_scheduling?: boolean;
  /** Network mode */
  network_mode?: string;
  /** Enable dynamic load balancing data report */
  dynamic_load_balancing_data_report_enabled?: boolean;
  /** Enable dynamic load balancing weight */
  dynamic_load_balancing_weight_enabled?: boolean;
  /** List of enabled VDCs for dynamic load balancing */
  dynamic_load_balancing_enabled_vdcs?: Array<string>;
  /** Type of dynamic load balancing, enums: round_robin, smooth_weighted_round_robin, weighted_round_robin */
  dynamic_load_balance_type?: string;
  /** Deployment inactive status */
  deployment_inactive?: boolean;
  /** Deployment inactive status for each zone */
  is_this_zone_deployment_inactive?: Record<string, boolean>;
  /** Package name */
  package?: string;
  /** Pod type */
  pod_type?: string;
  /** Plugin name */
  plugin_name?: string;
  /** Allow cold start instance */
  allow_cold_start_instance?: boolean;
  /** Elastic preferred cluster mapping */
  elastic_prefer_cluster?: Record<string, string>;
  /** Reserved preferred cluster mapping */
  reserved_prefer_cluster?: Record<string, string>;
  /** Elastic user preferred cluster mapping */
  elastic_user_preferred_cluster?: Record<string, string>;
  /** Reserved user preferred cluster mapping */
  reserved_user_preferred_cluster?: Record<string, string>;
  /** Elastic auto preferred cluster mapping */
  elastic_auto_preferred_cluster?: Record<string, string>;
  /** Reserved auto preferred cluster mapping */
  reserved_auto_preferred_cluster?: Record<string, string>;
  /** Temporary preferred cluster mapping */
  temp_preferred_cluster?: Record<string, string>;
  /** Formatted elastic preferred cluster list */
  formatted_elastic_prefer_cluster?: Array<common.FormattedPreferCluster>;
  /** Formatted reserved preferred cluster list */
  formatted_reserved_prefer_cluster?: Array<common.FormattedPreferCluster>;
  /** Runtime log writers */
  runtime_log_writers?: string;
  /** System log writers */
  system_log_writers?: string;
  /** Enable runtime JSON log */
  is_bytepaas_elastic_cluster?: boolean;
  /** Disable service discovery，true means disable */
  disable_service_discovery?: boolean;
  /** Enable resource guarantee */
  resource_guarantee?: boolean;
  /** Disable cgroup v2 */
  disable_cgroup_v2?: boolean;
  /** Enable async result emit event bridge */
  async_result_emit_event_bridge?: boolean;
  /** Runtime stream log bytes per second */
  runtime_stream_log_bytes_per_sec?: number;
  /** System stream log bytes per second */
  system_stream_log_bytes_per_sec?: number;
  /** Throttle log bytes per second */
  throttle_log_bytes_per_sec?: number;
  /** Throttle stdout log bytes per second */
  throttle_stdout_log_bytes_per_sec?: number;
  /** Throttle stderr log bytes per second */
  throttle_stderr_log_bytes_per_sec?: number;
  /** Enable scaling */
  scale_enabled?: boolean;
  /** Scale threshold */
  scale_threshold?: number;
  /** Scale type */
  scale_type?: number;
  /** Scale settings */
  scale_settings?: admin.FuncScaleSettings;
  /** Pod replica limits */
  replica_limit?: Record<string, common.PodReplicaLimit>;
  /** Reserved frozen replicas for each zone */
  zone_reserved_frozen_replicas?: Record<string, number>;
  /** Container runtime */
  container_runtime?: string;
  /** Enable scale optimization */
  enable_scale_optimise?: boolean;
  /** Schedule strategy */
  schedule_strategy?: string;
  /** Dynamic overcommit settings */
  dynamic_overcommit_settings?: Record<
    string,
    ticket.DynamicOvercommitSettings
  >;
  /** Enable overload protection */
  overload_protect_enabled?: boolean;
  /** Frozen CPU milli value */
  frozen_cpu_milli?: number;
  /** Enable federated on-demand resource */
  enable_fed_on_demand_resource?: Record<string, boolean>;
  /** Frozen priority class */
  frozen_priority_class?: string;
  /** Enable frozen scale */
  host_uniq?: common.HostUniq;
  /** Canary replica limit for each zone */
  zone_canary_replica_limit?: Record<string, number>;
  /** Enable frozen scale */
  frozen_scale_enabled?: boolean;
  /** enable privileged in pod */
  privileged?: boolean;
  /** In cell migration state */
  in_cell_migration?: boolean;
  region?: string;
  service_id?: string;
  cluster?: string;
  'X-Jwt-Token'?: string;
}

export interface UpdateRegionalMetaResponse {
  code?: number;
  data?: RegionalMeta;
  error?: string;
}
/* eslint-enable */
