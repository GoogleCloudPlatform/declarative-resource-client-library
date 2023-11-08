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
// gen_go_data -package bigqueryreservation -var YAML_assignment blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigqueryreservation/assignment.yaml

package bigqueryreservation

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigqueryreservation/assignment.yaml
var YAML_assignment = []byte("info:\n  title: BigqueryReservation/Assignment\n  description: The BigqueryReservation Assignment resource\n  x-dcl-struct-name: Assignment\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Assignment\n    parameters:\n    - name: assignment\n      required: true\n      description: A full instance of a Assignment\n  apply:\n    description: The function used to apply information about a Assignment\n    parameters:\n    - name: assignment\n      required: true\n      description: A full instance of a Assignment\n  delete:\n    description: The function used to delete a Assignment\n    parameters:\n    - name: assignment\n      required: true\n      description: A full instance of a Assignment\n  deleteAll:\n    description: The function used to delete all Assignment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: reservation\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Assignment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: reservation\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Assignment:\n      title: Assignment\n      x-dcl-id: projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - assignee\n      - jobType\n      - reservation\n      properties:\n        assignee:\n          type: string\n          x-dcl-go-name: Assignee\n          description: The resource which will use the reservation. E.g. projects/myproject,\n            folders/123, organizations/456.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n          - resource: Cloudresourcemanager/Folder\n            field: name\n          - resource: Cloudresourcemanager/Organization\n            field: name\n        jobType:\n          type: string\n          x-dcl-go-name: JobType\n          x-dcl-go-type: AssignmentJobTypeEnum\n          description: 'Types of job, which could be specified when using the reservation.\n            Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY'\n          x-kubernetes-immutable: true\n          enum:\n          - JOB_TYPE_UNSPECIFIED\n          - PIPELINE\n          - QUERY\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-extract-if-empty: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the assignment.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-extract-if-empty: true\n        reservation:\n          type: string\n          x-dcl-go-name: Reservation\n          description: The reservation for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Bigqueryreservation/Reservation\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: AssignmentStateEnum\n          readOnly: true\n          description: 'Assignment will remain in PENDING state if no active capacity\n            commitment is present. It will become ACTIVE when some capacity commitment\n            becomes active. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PENDING\n          - ACTIVE\n")

// 4351 bytes
// MD5: 5b2a228770438e48e79d732864ef8957
