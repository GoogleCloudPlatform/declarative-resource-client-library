// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package alpha -var YAML_metadata_store blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/alpha/metadata_store.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/alpha/metadata_store.yaml
var YAML_metadata_store = []byte("info:\n  title: VertexAI/MetadataStore\n  description: The VertexAI MetadataStore resource\n  x-dcl-struct-name: MetadataStore\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a MetadataStore\n    parameters:\n    - name: MetadataStore\n      required: true\n      description: A full instance of a MetadataStore\n  apply:\n    description: The function used to apply information about a MetadataStore\n    parameters:\n    - name: MetadataStore\n      required: true\n      description: A full instance of a MetadataStore\n  delete:\n    description: The function used to delete a MetadataStore\n    parameters:\n    - name: MetadataStore\n      required: true\n      description: A full instance of a MetadataStore\n  deleteAll:\n    description: The function used to delete all MetadataStore\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many MetadataStore\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    MetadataStore:\n      title: MetadataStore\n      x-dcl-id: projects/{{project}}/locations/{{location}}/metadataStores/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Timestamp when this MetadataStore was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Description of the MetadataStore.\n          x-kubernetes-immutable: true\n        encryptionSpec:\n          type: object\n          x-dcl-go-name: EncryptionSpec\n          x-dcl-go-type: MetadataStoreEncryptionSpec\n          description: Customer-managed encryption key spec for a Metadata Store.\n            If set, this Metadata Store and all sub-resources of this Metadata Store\n            are secured using this key.\n          x-kubernetes-immutable: true\n          required:\n          - kmsKeyName\n          properties:\n            kmsKeyName:\n              type: string\n              x-dcl-go-name: KmsKeyName\n              description: 'Required. The Cloud KMS resource identifier of the customer\n                managed encryption key used to protect a resource. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.\n                The key needs to be in the same region as where the compute resource\n                is created.'\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Cloudkms/CryptoKey\n                field: name\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the MetadataStore instance.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: object\n          x-dcl-go-name: State\n          x-dcl-go-type: MetadataStoreState\n          readOnly: true\n          description: Output only. State information of the MetadataStore.\n          x-kubernetes-immutable: true\n          properties:\n            diskUtilizationBytes:\n              type: integer\n              format: int64\n              x-dcl-go-name: DiskUtilizationBytes\n              description: The disk utilization of the MetadataStore in bytes.\n              x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Timestamp when this MetadataStore was last updated.\n          x-kubernetes-immutable: true\n")

// 4507 bytes
// MD5: 56798b141ce09836853ce61eaf62ce8e
