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
// gen_go_data -package beta -var YAML_instance blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/datafusion/beta/instance.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/datafusion/beta/instance.yaml
var YAML_instance = []byte("info:\n  title: DataFusion/Instance\n  description: The DataFusion Instance resource\n  x-dcl-struct-name: Instance\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Instance\n    parameters:\n    - name: instance\n      required: true\n      description: A full instance of a Instance\n  apply:\n    description: The function used to apply information about a Instance\n    parameters:\n    - name: instance\n      required: true\n      description: A full instance of a Instance\n  delete:\n    description: The function used to delete a Instance\n    parameters:\n    - name: instance\n      required: true\n      description: A full instance of a Instance\n  deleteAll:\n    description: The function used to delete all Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Instance:\n      title: Instance\n      x-dcl-id: projects/{{project}}/locations/{{location}}/instances/{{name}}\n      x-dcl-locations:\n      - zone\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - type\n      - project\n      - location\n      properties:\n        apiEndpoint:\n          type: string\n          x-dcl-go-name: ApiEndpoint\n          readOnly: true\n          description: Output only. Endpoint on which the REST APIs is accessible.\n          x-kubernetes-immutable: true\n        availableVersion:\n          type: array\n          x-dcl-go-name: AvailableVersion\n          readOnly: true\n          description: Available versions that the instance can be upgraded to.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceAvailableVersion\n            properties:\n              availableFeatures:\n                type: array\n                x-dcl-go-name: AvailableFeatures\n                readOnly: true\n                description: Represents a list of available feature names for a given\n                  version.\n                x-kubernetes-immutable: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n              defaultVersion:\n                type: boolean\n                x-dcl-go-name: DefaultVersion\n                readOnly: true\n                description: Whether this is currently the default version for Cloud\n                  Data Fusion\n                x-kubernetes-immutable: true\n              versionNumber:\n                type: string\n                x-dcl-go-name: VersionNumber\n                readOnly: true\n                description: The version number of the Data Fusion instance, such\n                  as '6.0.1.0'.\n                x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time the instance was created.\n          x-kubernetes-immutable: true\n        dataprocServiceAccount:\n          type: string\n          x-dcl-go-name: DataprocServiceAccount\n          description: User-managed service account to set on Dataproc when Cloud\n            Data Fusion creates Dataproc to run data processing pipelines. This allows\n            users to have fine-grained access control on Dataproc's accesses to cloud\n            resources.\n          x-dcl-references:\n          - resource: Iam/ServiceAccount\n            field: email\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A description of this instance.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Display name for an instance.\n          x-kubernetes-immutable: true\n        enableStackdriverLogging:\n          type: boolean\n          x-dcl-go-name: EnableStackdriverLogging\n          description: Option to enable Stackdriver Logging.\n        enableStackdriverMonitoring:\n          type: boolean\n          x-dcl-go-name: EnableStackdriverMonitoring\n          description: Option to enable Stackdriver Monitoring.\n        gcsBucket:\n          type: string\n          x-dcl-go-name: GcsBucket\n          readOnly: true\n          description: Output only. Cloud Storage bucket generated by Data Fusion\n            in the customer project.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: The resource labels for instance to use to annotate any related\n            underlying resources such as GCE VMs. The character '=' is not allowed\n            to be used within the labels.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.\n          x-kubernetes-immutable: true\n        networkConfig:\n          type: object\n          x-dcl-go-name: NetworkConfig\n          x-dcl-go-type: InstanceNetworkConfig\n          description: Network configuration options. These are required when a private\n            Data Fusion instance is to be created.\n          x-kubernetes-immutable: true\n          properties:\n            ipAllocation:\n              type: string\n              x-dcl-go-name: IPAllocation\n              description: The IP range in CIDR notation to use for the managed Data\n                Fusion instance nodes. This range must not overlap with any other\n                ranges used in the customer network.\n              x-kubernetes-immutable: true\n            network:\n              type: string\n              x-dcl-go-name: Network\n              description: Name of the network in the customer project with which\n                the Tenant Project will be peered for executing pipelines. In case\n                of shared VPC where the network resides in another host project the\n                network should specified in the form of projects/{host-project-id}/global/networks/{network}\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Compute/Network\n                field: name\n        options:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Options\n          description: Map of additional options used to configure the behavior of\n            Data Fusion instance.\n          x-kubernetes-immutable: true\n        p4ServiceAccount:\n          type: string\n          x-dcl-go-name: P4ServiceAccount\n          readOnly: true\n          description: Output only. P4 service account for the customer project.\n          x-kubernetes-immutable: true\n        privateInstance:\n          type: boolean\n          x-dcl-go-name: PrivateInstance\n          description: Specifies whether the Data Fusion instance should be private.\n            If set to true, all Data Fusion nodes will have private IP addresses and\n            will not be able to access the public internet.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        serviceEndpoint:\n          type: string\n          x-dcl-go-name: ServiceEndpoint\n          readOnly: true\n          description: Output only. Endpoint on which the Data Fusion UI is accessible.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: InstanceStateEnum\n          readOnly: true\n          description: 'Output only. The current state of this Data Fusion instance.\n            Possible values: STATE_UNSPECIFIED, ENABLED, DISABLED, UNKNOWN'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ENABLED\n          - DISABLED\n          - UNKNOWN\n        stateMessage:\n          type: string\n          x-dcl-go-name: StateMessage\n          readOnly: true\n          description: Output only. Additional information about the current state\n            of this Data Fusion instance if available.\n          x-kubernetes-immutable: true\n        tenantProjectId:\n          type: string\n          x-dcl-go-name: TenantProjectId\n          readOnly: true\n          description: Output only. The name of the tenant project.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n        type:\n          type: string\n          x-dcl-go-name: Type\n          x-dcl-go-type: InstanceTypeEnum\n          description: 'Required. Instance type. Possible values: TYPE_UNSPECIFIED,\n            BASIC, ENTERPRISE, DEVELOPER'\n          x-kubernetes-immutable: true\n          enum:\n          - TYPE_UNSPECIFIED\n          - BASIC\n          - ENTERPRISE\n          - DEVELOPER\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time the instance was last updated.\n          x-kubernetes-immutable: true\n        version:\n          type: string\n          x-dcl-go-name: Version\n          description: Current version of the Data Fusion.\n        zone:\n          type: string\n          x-dcl-go-name: Zone\n          description: Name of the zone in which the Data Fusion instance will be\n            created. Only DEVELOPER instances use this field.\n          x-kubernetes-immutable: true\n")

// 10492 bytes
// MD5: 79d09d813d545ef516ec9aecb3ec660c
