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
// gen_go_data -package dataproc -var YAML_autoscaling_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataproc/autoscaling_policy.yaml

package dataproc

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataproc/autoscaling_policy.yaml
var YAML_autoscaling_policy = []byte("info:\n  title: Dataproc/AutoscalingPolicy\n  description: The Dataproc AutoscalingPolicy resource\n  x-dcl-struct-name: AutoscalingPolicy\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a AutoscalingPolicy\n    parameters:\n    - name: autoscalingPolicy\n      required: true\n      description: A full instance of a AutoscalingPolicy\n  apply:\n    description: The function used to apply information about a AutoscalingPolicy\n    parameters:\n    - name: autoscalingPolicy\n      required: true\n      description: A full instance of a AutoscalingPolicy\n  delete:\n    description: The function used to delete a AutoscalingPolicy\n    parameters:\n    - name: autoscalingPolicy\n      required: true\n      description: A full instance of a AutoscalingPolicy\n  deleteAll:\n    description: The function used to delete all AutoscalingPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many AutoscalingPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    AutoscalingPolicy:\n      title: AutoscalingPolicy\n      x-dcl-id: projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - basicAlgorithm\n      - workerConfig\n      - project\n      - location\n      properties:\n        basicAlgorithm:\n          type: object\n          x-dcl-go-name: BasicAlgorithm\n          x-dcl-go-type: AutoscalingPolicyBasicAlgorithm\n          required:\n          - yarnConfig\n          properties:\n            cooldownPeriod:\n              type: string\n              x-dcl-go-name: CooldownPeriod\n              description: 'Optional. Duration between scaling events. A scaling period\n                starts after the update operation from the previous event has completed.\n                Bounds: . Default: 2m.'\n              x-dcl-server-default: true\n            yarnConfig:\n              type: object\n              x-dcl-go-name: YarnConfig\n              x-dcl-go-type: AutoscalingPolicyBasicAlgorithmYarnConfig\n              description: Required. YARN autoscaling configuration.\n              required:\n              - gracefulDecommissionTimeout\n              - scaleUpFactor\n              - scaleDownFactor\n              properties:\n                gracefulDecommissionTimeout:\n                  type: string\n                  x-dcl-go-name: GracefulDecommissionTimeout\n                  description: Required. Timeout for YARN graceful decommissioning\n                    of Node Managers. Specifies the duration to wait for jobs to complete\n                    before forcefully removing workers (and potentially interrupting\n                    jobs). Only applicable to downscaling operations.\n                scaleDownFactor:\n                  type: number\n                  format: double\n                  x-dcl-go-name: ScaleDownFactor\n                  description: Required. Fraction of average YARN pending memory in\n                    the last cooldown period for which to remove workers. A scale-down\n                    factor of 1 will result in scaling down so that there is no available\n                    memory remaining after the update (more aggressive scaling). A\n                    scale-down factor of 0 disables removing workers, which can be\n                    beneficial for autoscaling a single job. See .\n                scaleDownMinWorkerFraction:\n                  type: number\n                  format: double\n                  x-dcl-go-name: ScaleDownMinWorkerFraction\n                  description: 'Optional. Minimum scale-down threshold as a fraction\n                    of total cluster size before scaling occurs. For example, in a\n                    20-worker cluster, a threshold of 0.1 means the autoscaler must\n                    recommend at least a 2 worker scale-down for the cluster to scale.\n                    A threshold of 0 means the autoscaler will scale down on any recommended\n                    change. Bounds: . Default: 0.0.'\n                scaleUpFactor:\n                  type: number\n                  format: double\n                  x-dcl-go-name: ScaleUpFactor\n                  description: Required. Fraction of average YARN pending memory in\n                    the last cooldown period for which to add workers. A scale-up\n                    factor of 1.0 will result in scaling up so that there is no pending\n                    memory remaining after the update (more aggressive scaling). A\n                    scale-up factor closer to 0 will result in a smaller magnitude\n                    of scaling up (less aggressive scaling). See .\n                scaleUpMinWorkerFraction:\n                  type: number\n                  format: double\n                  x-dcl-go-name: ScaleUpMinWorkerFraction\n                  description: 'Optional. Minimum scale-up threshold as a fraction\n                    of total cluster size before scaling occurs. For example, in a\n                    20-worker cluster, a threshold of 0.1 means the autoscaler must\n                    recommend at least a 2-worker scale-up for the cluster to scale.\n                    A threshold of 0 means the autoscaler will scale up on any recommended\n                    change. Bounds: . Default: 0.0.'\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The \"resource name\" of the autoscaling policy, as described\n            in https://cloud.google.com/apis/design/resource_names. * For `projects.regions.autoscalingPolicies`,\n            the resource name of the policy has the following format: `projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id}`\n            * For `projects.locations.autoscalingPolicies`, the resource name of the\n            policy has the following format: `projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}`'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        secondaryWorkerConfig:\n          type: object\n          x-dcl-go-name: SecondaryWorkerConfig\n          x-dcl-go-type: AutoscalingPolicySecondaryWorkerConfig\n          description: Optional. Describes how the autoscaler will operate for secondary\n            workers.\n          properties:\n            maxInstances:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxInstances\n              description: 'Optional. Maximum number of instances for this group.\n                Note that by default, clusters will not use secondary workers. Required\n                for secondary workers if the minimum secondary instances is set. Primary\n                workers - Bounds: [min_instances, ). Secondary workers - Bounds: [min_instances,\n                ). Default: 0.'\n            minInstances:\n              type: integer\n              format: int64\n              x-dcl-go-name: MinInstances\n              description: 'Optional. Minimum number of instances for this group.\n                Primary workers - Bounds: . Default: 0.'\n            weight:\n              type: integer\n              format: int64\n              x-dcl-go-name: Weight\n              description: 'Optional. Weight for the instance group, which is used\n                to determine the fraction of total workers in the cluster from this\n                instance group. For example, if primary workers have weight 2, and\n                secondary workers have weight 1, the cluster will have approximately\n                2 primary workers for each secondary worker. The cluster may not reach\n                the specified balance if constrained by min/max bounds or other autoscaling\n                settings. For example, if `max_instances` for secondary workers is\n                0, then only primary workers will be added. The cluster can also be\n                out of balance when created. If weight is not set on any instance\n                group, the cluster will default to equal weight for all groups: the\n                cluster will attempt to maintain an equal number of workers in each\n                group within the configured size bounds for each group. If weight\n                is set for one group only, the cluster will default to zero weight\n                on the unset group. For example if weight is set only on primary workers,\n                the cluster will use primary workers only and no secondary workers.'\n              x-dcl-server-default: true\n        workerConfig:\n          type: object\n          x-dcl-go-name: WorkerConfig\n          x-dcl-go-type: AutoscalingPolicyWorkerConfig\n          description: Required. Describes how the autoscaler will operate for primary\n            workers.\n          required:\n          - maxInstances\n          properties:\n            maxInstances:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxInstances\n              description: 'Required. Maximum number of instances for this group.\n                Required for primary workers. Note that by default, clusters will\n                not use secondary workers. Required for secondary workers if the minimum\n                secondary instances is set. Primary workers - Bounds: [min_instances,\n                ). Secondary workers - Bounds: [min_instances, ). Default: 0.'\n            minInstances:\n              type: integer\n              format: int64\n              x-dcl-go-name: MinInstances\n              description: 'Optional. Minimum number of instances for this group.\n                Primary workers - Bounds: . Default: 0.'\n              x-dcl-server-default: true\n            weight:\n              type: integer\n              format: int64\n              x-dcl-go-name: Weight\n              description: 'Optional. Weight for the instance group, which is used\n                to determine the fraction of total workers in the cluster from this\n                instance group. For example, if primary workers have weight 2, and\n                secondary workers have weight 1, the cluster will have approximately\n                2 primary workers for each secondary worker. The cluster may not reach\n                the specified balance if constrained by min/max bounds or other autoscaling\n                settings. For example, if `max_instances` for secondary workers is\n                0, then only primary workers will be added. The cluster can also be\n                out of balance when created. If weight is not set on any instance\n                group, the cluster will default to equal weight for all groups: the\n                cluster will attempt to maintain an equal number of workers in each\n                group within the configured size bounds for each group. If weight\n                is set for one group only, the cluster will default to zero weight\n                on the unset group. For example if weight is set only on primary workers,\n                the cluster will use primary workers only and no secondary workers.'\n              x-dcl-server-default: true\n")

// 11970 bytes
// MD5: 484c94b8b8380d98a56d1441daf3ec82
