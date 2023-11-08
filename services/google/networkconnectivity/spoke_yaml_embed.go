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
// gen_go_data -package networkconnectivity -var YAML_spoke blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkconnectivity/spoke.yaml

package networkconnectivity

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkconnectivity/spoke.yaml
var YAML_spoke = []byte("info:\n  title: NetworkConnectivity/Spoke\n  description: The NetworkConnectivity Spoke resource\n  x-dcl-struct-name: Spoke\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Spoke\n    parameters:\n    - name: spoke\n      required: true\n      description: A full instance of a Spoke\n  apply:\n    description: The function used to apply information about a Spoke\n    parameters:\n    - name: spoke\n      required: true\n      description: A full instance of a Spoke\n  delete:\n    description: The function used to delete a Spoke\n    parameters:\n    - name: spoke\n      required: true\n      description: A full instance of a Spoke\n  deleteAll:\n    description: The function used to delete all Spoke\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Spoke\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Spoke:\n      title: Spoke\n      x-dcl-id: projects/{{project}}/locations/{{location}}/spokes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - hub\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time the spoke was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of the spoke.\n        hub:\n          type: string\n          x-dcl-go-name: Hub\n          description: Immutable. The URI of the hub that this spoke is attached to.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Networkconnectivity/Hub\n            field: name\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional labels in key:value format. For more information about\n            labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).\n        linkedInterconnectAttachments:\n          type: object\n          x-dcl-go-name: LinkedInterconnectAttachments\n          x-dcl-go-type: SpokeLinkedInterconnectAttachments\n          description: A collection of VLAN attachment resources. These resources\n            should be redundant attachments that all advertise the same prefixes to\n            Google Cloud. Alternatively, in active/passive configurations, all attachments\n            should be capable of advertising the same prefixes.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - linkedVpnTunnels\n          - linkedRouterApplianceInstances\n          - linkedVPCNetwork\n          required:\n          - uris\n          - siteToSiteDataTransfer\n          properties:\n            siteToSiteDataTransfer:\n              type: boolean\n              x-dcl-go-name: SiteToSiteDataTransfer\n              description: A value that controls whether site-to-site data transfer\n                is enabled for these resources. Note that data transfer is available\n                only in supported locations.\n              x-kubernetes-immutable: true\n            uris:\n              type: array\n              x-dcl-go-name: Uris\n              description: The URIs of linked interconnect attachment resources\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n                x-dcl-references:\n                - resource: Compute/InterconnectAttachment\n                  field: selfLink\n        linkedRouterApplianceInstances:\n          type: object\n          x-dcl-go-name: LinkedRouterApplianceInstances\n          x-dcl-go-type: SpokeLinkedRouterApplianceInstances\n          description: The URIs of linked Router appliance resources\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - linkedVpnTunnels\n          - linkedInterconnectAttachments\n          - linkedVPCNetwork\n          required:\n          - instances\n          - siteToSiteDataTransfer\n          properties:\n            instances:\n              type: array\n              x-dcl-go-name: Instances\n              description: The list of router appliance instances\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: SpokeLinkedRouterApplianceInstancesInstances\n                properties:\n                  ipAddress:\n                    type: string\n                    x-dcl-go-name: IPAddress\n                    description: The IP address on the VM to use for peering.\n                    x-kubernetes-immutable: true\n                  virtualMachine:\n                    type: string\n                    x-dcl-go-name: VirtualMachine\n                    description: The URI of the virtual machine resource\n                    x-kubernetes-immutable: true\n                    x-dcl-references:\n                    - resource: Compute/Instance\n                      field: selfLink\n            siteToSiteDataTransfer:\n              type: boolean\n              x-dcl-go-name: SiteToSiteDataTransfer\n              description: A value that controls whether site-to-site data transfer\n                is enabled for these resources. Note that data transfer is available\n                only in supported locations.\n              x-kubernetes-immutable: true\n        linkedVPCNetwork:\n          type: object\n          x-dcl-go-name: LinkedVPCNetwork\n          x-dcl-go-type: SpokeLinkedVPCNetwork\n          description: VPC network that is associated with the spoke.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - linkedVpnTunnels\n          - linkedInterconnectAttachments\n          - linkedRouterApplianceInstances\n          required:\n          - uri\n          properties:\n            excludeExportRanges:\n              type: array\n              x-dcl-go-name: ExcludeExportRanges\n              description: IP ranges encompassing the subnets to be excluded from\n                peering.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            uri:\n              type: string\n              x-dcl-go-name: Uri\n              description: The URI of the VPC network resource.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Compute/Network\n                field: selfLink\n        linkedVpnTunnels:\n          type: object\n          x-dcl-go-name: LinkedVpnTunnels\n          x-dcl-go-type: SpokeLinkedVpnTunnels\n          description: The URIs of linked VPN tunnel resources\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - linkedInterconnectAttachments\n          - linkedRouterApplianceInstances\n          - linkedVPCNetwork\n          required:\n          - uris\n          - siteToSiteDataTransfer\n          properties:\n            siteToSiteDataTransfer:\n              type: boolean\n              x-dcl-go-name: SiteToSiteDataTransfer\n              description: A value that controls whether site-to-site data transfer\n                is enabled for these resources. Note that data transfer is available\n                only in supported locations.\n              x-kubernetes-immutable: true\n            uris:\n              type: array\n              x-dcl-go-name: Uris\n              description: The URIs of linked VPN tunnel resources.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n                x-dcl-references:\n                - resource: Compute/VpnTunnel\n                  field: selfLink\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Immutable. The name of the spoke. Spoke names must be unique.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: SpokeStateEnum\n          readOnly: true\n          description: 'Output only. The current lifecycle state of this spoke. Possible\n            values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - ACTIVE\n          - DELETING\n        uniqueId:\n          type: string\n          x-dcl-go-name: UniqueId\n          readOnly: true\n          description: Output only. The Google-generated UUID for the spoke. This\n            value is unique across all spoke resources. If a spoke is deleted and\n            another with the same name is created, the new spoke is assigned a different\n            unique_id.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time the spoke was last updated.\n          x-kubernetes-immutable: true\n")

// 10368 bytes
// MD5: 60d347ff5891798baec18398f04af5f8
