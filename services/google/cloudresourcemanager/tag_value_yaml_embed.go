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
// gen_go_data -package cloudresourcemanager -var YAML_tag_value blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/tag_value.yaml

package cloudresourcemanager

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/tag_value.yaml
var YAML_tag_value = []byte("info:\n  title: CloudResourceManager/TagValue\n  description: The CloudResourceManager TagValue resource\n  x-dcl-struct-name: TagValue\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a TagValue\n    parameters:\n    - name: tagValue\n      required: true\n      description: A full instance of a TagValue\n  apply:\n    description: The function used to apply information about a TagValue\n    parameters:\n    - name: tagValue\n      required: true\n      description: A full instance of a TagValue\n  delete:\n    description: The function used to delete a TagValue\n    parameters:\n    - name: tagValue\n      required: true\n      description: A full instance of a TagValue\ncomponents:\n  schemas:\n    TagValue:\n      title: TagValue\n      x-dcl-id: tagValues/{{name}}\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - parent\n      - shortName\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Creation time.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. User-assigned description of the TagValue. Must not\n            exceed 256 characters. Read-write.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Optional. Entity tag which users can pass to prevent race conditions.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Immutable. The generated numeric id for the TagValue.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        namespacedName:\n          type: string\n          x-dcl-go-name: NamespacedName\n          readOnly: true\n          description: Output only. Immutable. Namespaced name of the TagValue.\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: Immutable. The resource name of the new TagValue's parent.\n            Must be of the form `tagKeys/{tag_key_id}`.\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n        shortName:\n          type: string\n          x-dcl-go-name: ShortName\n          description: Required. Immutable. The user friendly name for a TagValue.\n            The short name should be unique for TagValuess within the same parent\n            TagKey. The short name must be 1-63 characters, beginning and ending with\n            an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),\n            dots (.), and alphanumerics between.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Update time.\n          x-kubernetes-immutable: true\n")

// 3167 bytes
// MD5: 5ef7f1dc1f9aaa63456914b412a559f8
