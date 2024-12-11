// Copyright 2024 Google LLC. All Rights Reserved.
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
// gen_go_data -package beta -var YAML_service_attachment blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/service_attachment.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/service_attachment.yaml
var YAML_service_attachment = []byte("info:\n  title: Compute/ServiceAttachment\n  description: Represents a ServiceAttachment resource.\n  x-dcl-struct-name: ServiceAttachment\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: API documentation\n    url: https://cloud.google.com/compute/docs/reference/rest/beta/serviceAttachments\n  x-dcl-guides:\n  - text: Configuring Private Service Connect to access services\n    url: https://cloud.google.com/vpc/docs/configure-private-service-connect-services\npaths:\n  get:\n    description: The function used to get information about a ServiceAttachment\n    parameters:\n    - name: serviceAttachment\n      required: true\n      description: A full instance of a ServiceAttachment\n  apply:\n    description: The function used to apply information about a ServiceAttachment\n    parameters:\n    - name: serviceAttachment\n      required: true\n      description: A full instance of a ServiceAttachment\n  delete:\n    description: The function used to delete a ServiceAttachment\n    parameters:\n    - name: serviceAttachment\n      required: true\n      description: A full instance of a ServiceAttachment\n  deleteAll:\n    description: The function used to delete all ServiceAttachment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ServiceAttachment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ServiceAttachment:\n      title: ServiceAttachment\n      x-dcl-id: projects/{{project}}/regions/{{location}}/serviceAttachments/{{name}}\n      x-dcl-locations:\n      - region\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - targetService\n      - connectionPreference\n      - natSubnets\n      - project\n      - location\n      properties:\n        connectedEndpoints:\n          type: array\n          x-dcl-go-name: ConnectedEndpoints\n          readOnly: true\n          description: An array of connections for all the consumers connected to\n            this service attachment.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ServiceAttachmentConnectedEndpoints\n            properties:\n              endpoint:\n                type: string\n                x-dcl-go-name: Endpoint\n                description: The url of a connected endpoint.\n              pscConnectionId:\n                type: integer\n                format: int64\n                x-dcl-go-name: PscConnectionId\n                description: The PSC connection id of the connected endpoint.\n              status:\n                type: string\n                x-dcl-go-name: Status\n                x-dcl-go-type: ServiceAttachmentConnectedEndpointsStatusEnum\n                description: 'The status of a connected endpoint to this service attachment.\n                  Possible values: PENDING, RUNNING, DONE'\n                enum:\n                - PENDING\n                - RUNNING\n                - DONE\n        connectionPreference:\n          type: string\n          x-dcl-go-name: ConnectionPreference\n          x-dcl-go-type: ServiceAttachmentConnectionPreferenceEnum\n          description: 'The connection preference of service attachment. The value\n            can be set to `ACCEPT_AUTOMATIC`. An `ACCEPT_AUTOMATIC` service attachment\n            is one that always accepts the connection from consumer forwarding rules.\n            Possible values: CONNECTION_PREFERENCE_UNSPECIFIED, ACCEPT_AUTOMATIC,\n            ACCEPT_MANUAL'\n          enum:\n          - CONNECTION_PREFERENCE_UNSPECIFIED\n          - ACCEPT_AUTOMATIC\n          - ACCEPT_MANUAL\n        consumerAcceptLists:\n          type: array\n          x-dcl-go-name: ConsumerAcceptLists\n          description: Projects that are allowed to connect to this service attachment.\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: object\n            x-dcl-go-type: ServiceAttachmentConsumerAcceptLists\n            required:\n            - projectIdOrNum\n            properties:\n              connectionLimit:\n                type: integer\n                format: int64\n                x-dcl-go-name: ConnectionLimit\n                description: The value of the limit to set.\n              projectIdOrNum:\n                type: string\n                x-dcl-go-name: ProjectIdOrNum\n                description: The project id or number for the project to set the limit\n                  for.\n                x-dcl-references:\n                - resource: Cloudresourcemanager/Project\n                  field: name\n        consumerRejectLists:\n          type: array\n          x-dcl-go-name: ConsumerRejectLists\n          description: Projects that are not allowed to connect to this service attachment.\n            The project can be specified using its id or number.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Cloudresourcemanager/Project\n              field: name\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n        enableProxyProtocol:\n          type: boolean\n          x-dcl-go-name: EnableProxyProtocol\n          description: If true, enable the proxy protocol which is for supplying client\n            TCP/IP address data in TCP connections that traverse proxies on their\n            way to destination servers.\n          x-kubernetes-immutable: true\n        fingerprint:\n          type: string\n          x-dcl-go-name: Fingerprint\n          readOnly: true\n          description: Fingerprint of this resource. This field is used internally\n            during updates of this resource.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource type. The server generates\n            this identifier.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `)?` which means the first character must be a lowercase letter,\n            and all following characters must be a dash, lowercase letter, or digit,\n            except the last character, which cannot be a dash.\n          x-kubernetes-immutable: true\n        natSubnets:\n          type: array\n          x-dcl-go-name: NatSubnets\n          description: An array of URLs where each entry is the URL of a subnet provided\n            by the service producer to use for NAT in this service attachment.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Compute/Subnetwork\n              field: selfLink\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        pscServiceAttachmentId:\n          type: object\n          x-dcl-go-name: PscServiceAttachmentId\n          x-dcl-go-type: ServiceAttachmentPscServiceAttachmentId\n          readOnly: true\n          description: An 128-bit global unique ID of the PSC service attachment.\n          x-kubernetes-immutable: true\n          properties:\n            high:\n              type: integer\n              format: int64\n              x-dcl-go-name: High\n              readOnly: true\n              x-kubernetes-immutable: true\n            low:\n              type: integer\n              format: int64\n              x-dcl-go-name: Low\n              readOnly: true\n              x-kubernetes-immutable: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          readOnly: true\n          description: URL of the region where the service attachment resides. This\n            field applies only to the region resource. You must specify this field\n            as part of the HTTP request URL. It is not settable as a field in the\n            request body.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        targetService:\n          type: string\n          x-dcl-go-name: TargetService\n          description: The URL of a service serving the endpoint identified by this\n            service attachment.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/ForwardingRule\n            field: selfLink\n")

// 9886 bytes
// MD5: 4e3bee73d93bec95f4ec0b50172091da
