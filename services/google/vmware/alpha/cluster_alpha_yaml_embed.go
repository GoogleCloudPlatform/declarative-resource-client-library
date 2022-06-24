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
// gen_go_data -package alpha -var YAML_cluster blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/cluster.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/cluster.yaml
var YAML_cluster = []byte("info:\n  title: Vmware/Cluster\n  description: The Vmware Cluster resource\n  x-dcl-struct-name: Cluster\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Cluster\n    parameters:\n    - name: Cluster\n      required: true\n      description: A full instance of a Cluster\n  apply:\n    description: The function used to apply information about a Cluster\n    parameters:\n    - name: Cluster\n      required: true\n      description: A full instance of a Cluster\n  delete:\n    description: The function used to delete a Cluster\n    parameters:\n    - name: Cluster\n      required: true\n      description: A full instance of a Cluster\n  deleteAll:\n    description: The function used to delete all Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: privatecloud\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: privatecloud\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Cluster:\n      title: Cluster\n      x-dcl-id: projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - nodeTypeId\n      - nodeCount\n      - project\n      - location\n      - privateCloud\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Creation time of this resource in RFC3339 text\n            format.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        management:\n          type: boolean\n          x-dcl-go-name: Management\n          readOnly: true\n          description: Output only. True if the cluster is a management cluster; false\n            otherwise. There can only be one management cluster in a private cloud\n            and it has to be the first one.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of this cluster. Resource names are schemeless\n            URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\n            For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud/clusters/my-cluster`'\n          x-kubernetes-immutable: true\n        nodeCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: NodeCount\n          description: Required. Number of nodes in this cluster.\n        nodeTypeId:\n          type: string\n          x-dcl-go-name: NodeTypeId\n          description: 'Required. The canonical identifier of node types (`NodeType`)\n            in this cluster. For example: standard-72.'\n        privateCloud:\n          type: string\n          x-dcl-go-name: PrivateCloud\n          description: The privateCloud for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Vmware/PrivateCloud\n            field: name\n            parent: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ClusterStateEnum\n          readOnly: true\n          description: 'Output only. State of the resource. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, CREATING, UPDATING, FAILED, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - CREATING\n          - UPDATING\n          - FAILED\n          - DELETED\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Last update time of this resource in RFC3339 text\n            format.\n          x-kubernetes-immutable: true\n")

// 4677 bytes
// MD5: 0a2844705b8d6cb1ebd5937dba59795b
