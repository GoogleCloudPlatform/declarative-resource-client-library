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
// gen_go_data -package alpha -var YAML_google_channel_config blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/eventarc/alpha/google_channel_config.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/eventarc/alpha/google_channel_config.yaml
var YAML_google_channel_config = []byte("info:\n  title: Eventarc/GoogleChannelConfig\n  description: The Eventarc GoogleChannelConfig resource\n  x-dcl-struct-name: GoogleChannelConfig\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a GoogleChannelConfig\n    parameters:\n    - name: googleChannelConfig\n      required: true\n      description: A full instance of a GoogleChannelConfig\n  apply:\n    description: The function used to apply information about a GoogleChannelConfig\n    parameters:\n    - name: googleChannelConfig\n      required: true\n      description: A full instance of a GoogleChannelConfig\n  delete:\n    description: The function used to delete a GoogleChannelConfig\n    parameters:\n    - name: googleChannelConfig\n      required: true\n      description: A full instance of a GoogleChannelConfig\ncomponents:\n  schemas:\n    GoogleChannelConfig:\n      title: GoogleChannelConfig\n      x-dcl-id: projects/{{project}}/locations/{{location}}/googleChannelConfig\n      x-dcl-parent-container: project\n      x-dcl-has-create: false\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        cryptoKeyName:\n          type: string\n          x-dcl-go-name: CryptoKeyName\n          description: Optional. Resource name of a KMS crypto key (managed by the\n            user) used to encrypt/decrypt their event data. It must match the pattern\n            `projects/*/locations/*/keyRings/*/cryptoKeys/*`.\n          x-dcl-references:\n          - resource: Cloudkms/CryptoKey\n            field: name\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. The resource name of the config. Must be in the format\n            of, `projects/{project}/locations/{location}/googleChannelConfig`.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last-modified time.\n          x-kubernetes-immutable: true\n")

// 2632 bytes
// MD5: 0c3f04539396904e0f6a7fd5eba9d3d2
