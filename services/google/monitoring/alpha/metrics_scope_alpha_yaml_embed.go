// Copyright 2023 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_metrics_scope blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/metrics_scope.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/metrics_scope.yaml
var YAML_metrics_scope = []byte("info:\n  title: Monitoring/MetricsScope\n  description: The Monitoring MetricsScope resource\n  x-dcl-struct-name: MetricsScope\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a MetricsScope\n    parameters:\n    - name: metricsScope\n      required: true\n      description: A full instance of a MetricsScope\n  apply:\n    description: The function used to apply information about a MetricsScope\n    parameters:\n    - name: metricsScope\n      required: true\n      description: A full instance of a MetricsScope\ncomponents:\n  schemas:\n    MetricsScope:\n      title: MetricsScope\n      x-dcl-id: locations/global/metricsScopes/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-has-create: false\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when this `Metrics Scope` was created.\n          x-kubernetes-immutable: true\n        monitoredProjects:\n          type: array\n          x-dcl-go-name: MonitoredProjects\n          readOnly: true\n          description: Output only. The list of projects monitored by this `Metrics\n            Scope`.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: MetricsScopeMonitoredProjects\n            properties:\n              createTime:\n                type: string\n                format: date-time\n                x-dcl-go-name: CreateTime\n                readOnly: true\n                description: Output only. The time when this `MonitoredProject` was\n                  created.\n                x-kubernetes-immutable: true\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: 'Immutable. The resource name of the `MonitoredProject`.\n                  On input, the resource name includes the scoping project ID and\n                  monitored project ID. On output, it contains the equivalent project\n                  numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`'\n                x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Immutable. The resource name of the Monitoring Metrics Scope.\n            On input, the resource name can be specified with the scoping project\n            ID or number. On output, the resource name is specified with the scoping\n            project number. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}`'\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time when this `Metrics Scope` record was\n            last updated.\n          x-kubernetes-immutable: true\n")

// 3179 bytes
// MD5: be6bbb3b86b646949a9e2c5d6f16127c
