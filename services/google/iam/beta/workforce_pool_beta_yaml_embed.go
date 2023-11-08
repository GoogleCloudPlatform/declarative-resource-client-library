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
// gen_go_data -package beta -var YAML_workforce_pool blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/beta/workforce_pool.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/beta/workforce_pool.yaml
var YAML_workforce_pool = []byte("info:\n  title: Iam/WorkforcePool\n  description: The Iam WorkforcePool resource\n  x-dcl-struct-name: WorkforcePool\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a WorkforcePool\n    parameters:\n    - name: workforcePool\n      required: true\n      description: A full instance of a WorkforcePool\n  apply:\n    description: The function used to apply information about a WorkforcePool\n    parameters:\n    - name: workforcePool\n      required: true\n      description: A full instance of a WorkforcePool\n  delete:\n    description: The function used to delete a WorkforcePool\n    parameters:\n    - name: workforcePool\n      required: true\n      description: A full instance of a WorkforcePool\n  deleteAll:\n    description: The function used to delete all WorkforcePool\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many WorkforcePool\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    WorkforcePool:\n      title: WorkforcePool\n      x-dcl-id: locations/{{location}}/workforcePools/{{name}}\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - parent\n      - location\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A user-specified description of the pool. Cannot exceed 256\n            characters.\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Whether the pool is disabled. You cannot use a disabled pool\n            to exchange tokens, or use existing tokens to access resources. If the\n            pool is re-enabled, existing tokens grant access again.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A user-specified display name of the pool in Google Cloud Console.\n            Cannot exceed 32 characters.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the pool. The ID must be a globally unique string\n            of 6 to 63 lowercase letters, digits, or hyphens. It must start with a\n            letter, and cannot have a trailing hyphen. The prefix `gcp-` is reserved\n            for use by Google, and may not be specified.\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: 'Immutable. The resource name of the parent. Format: `organizations/{org-id}`.'\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: 'Output only. The resource name of the pool. Format: `locations/{location}/workforcePools/{workforce_pool_id}`'\n          x-kubernetes-immutable: true\n        sessionDuration:\n          type: string\n          x-dcl-go-name: SessionDuration\n          description: How long the Google Cloud access tokens, console sign-in sessions,\n            and gcloud sign-in sessions from this pool are valid. Must be greater\n            than 15 minutes (900s) and less than 12 hours (43200s). If `session_duration`\n            is not configured, minted credentials will have a default duration of\n            one hour (3600s).\n          x-dcl-server-default: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: WorkforcePoolStateEnum\n          readOnly: true\n          description: 'Output only. The state of the pool. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETED\n")

// 4456 bytes
// MD5: bcbb6698af72256bc486931e3adfeb01
