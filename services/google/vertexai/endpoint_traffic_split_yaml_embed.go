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
// gen_go_data -package vertexai -var YAML_endpoint_traffic_split blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/endpoint_traffic_split.yaml

package vertexai

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/endpoint_traffic_split.yaml
var YAML_endpoint_traffic_split = []byte("info:\n  title: VertexAI/EndpointTrafficSplit\n  description: The VertexAI EndpointTrafficSplit resource\n  x-dcl-struct-name: EndpointTrafficSplit\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a EndpointTrafficSplit\n    parameters:\n    - name: EndpointTrafficSplit\n      required: true\n      description: A full instance of a EndpointTrafficSplit\n  apply:\n    description: The function used to apply information about a EndpointTrafficSplit\n    parameters:\n    - name: EndpointTrafficSplit\n      required: true\n      description: A full instance of a EndpointTrafficSplit\ncomponents:\n  schemas:\n    EndpointTrafficSplit:\n      title: EndpointTrafficSplit\n      x-dcl-id: projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: false\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - endpoint\n      - project\n      - location\n      - trafficSplit\n      properties:\n        endpoint:\n          type: string\n          x-dcl-go-name: Endpoint\n          description: The endpoint for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Vertexai/Endpoint\n            field: name\n            parent: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Used to perform consistent read-modify-write updates. If not\n            set, a blind \"overwrite\" update happens.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        trafficSplit:\n          type: array\n          x-dcl-go-name: TrafficSplit\n          description: A map from a DeployedModel's ID to the percentage of this Endpoint's\n            traffic that should be forwarded to that DeployedModel. If a DeployedModel's\n            ID is not listed in this map, then it receives no traffic. The traffic\n            percentage values must add up to 100, or map must be empty if the Endpoint\n            is to not accept any traffic at a moment.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: EndpointTrafficSplitTrafficSplit\n            required:\n            - deployedModelId\n            - trafficPercentage\n            properties:\n              deployedModelId:\n                type: string\n                x-dcl-go-name: DeployedModelId\n                description: A deployed model's id.\n                x-dcl-references:\n                - resource: Vertexai/ModelDeployment\n                  field: id\n              trafficPercentage:\n                type: integer\n                format: int64\n                x-dcl-go-name: TrafficPercentage\n                description: The percentage of this Endpoint's traffic that should\n                  be forwarded to the DeployedModel.\n")

// 3331 bytes
// MD5: c2bfe9fe606dd01818cee385532cf6ca
