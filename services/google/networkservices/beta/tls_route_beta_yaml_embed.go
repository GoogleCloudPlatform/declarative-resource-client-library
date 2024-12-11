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
// gen_go_data -package beta -var YAML_tls_route blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/beta/tls_route.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/beta/tls_route.yaml
var YAML_tls_route = []byte("info:\n  title: NetworkServices/TlsRoute\n  description: The NetworkServices TlsRoute resource\n  x-dcl-struct-name: TlsRoute\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a TlsRoute\n    parameters:\n    - name: tlsRoute\n      required: true\n      description: A full instance of a TlsRoute\n  apply:\n    description: The function used to apply information about a TlsRoute\n    parameters:\n    - name: tlsRoute\n      required: true\n      description: A full instance of a TlsRoute\n  delete:\n    description: The function used to delete a TlsRoute\n    parameters:\n    - name: tlsRoute\n      required: true\n      description: A full instance of a TlsRoute\n  deleteAll:\n    description: The function used to delete all TlsRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many TlsRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    TlsRoute:\n      title: TlsRoute\n      x-dcl-id: projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - rules\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        gateways:\n          type: array\n          x-dcl-go-name: Gateways\n          description: 'Optional. Gateways defines a list of gateways this TlsRoute\n            is attached to, as one of the routing rules to route the requests served\n            by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Gateway\n              field: selfLink\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        meshes:\n          type: array\n          x-dcl-go-name: Meshes\n          description: 'Optional. Meshes defines a list of meshes this TlsRoute is\n            attached to, as one of the routing rules to route the requests served\n            by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`\n            The attached Mesh should be of a type SIDECAR'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Mesh\n              field: selfLink\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the TlsRoute resource. It matches pattern\n            `projects/*/locations/global/tlsRoutes/tls_route_name>`.\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        rules:\n          type: array\n          x-dcl-go-name: Rules\n          description: Required. Rules that define how traffic is routed and handled.\n            At least one RouteRule must be supplied. If there are multiple rules then\n            the action taken will be the first rule to match.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: TlsRouteRules\n            required:\n            - matches\n            - action\n            properties:\n              action:\n                type: object\n                x-dcl-go-name: Action\n                x-dcl-go-type: TlsRouteRulesAction\n                description: Required. The detailed rule defining how to route matched\n                  traffic.\n                required:\n                - destinations\n                properties:\n                  destinations:\n                    type: array\n                    x-dcl-go-name: Destinations\n                    description: Required. The destination services to which traffic\n                      should be forwarded. At least one destination service is required.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: TlsRouteRulesActionDestinations\n                      required:\n                      - serviceName\n                      properties:\n                        serviceName:\n                          type: string\n                          x-dcl-go-name: ServiceName\n                          description: Required. The URL of a BackendService to route\n                            traffic to.\n                          x-dcl-references:\n                          - resource: Compute/BackendService\n                            field: name\n                            format: projects/{{project}}/locations/global/backendServices/{{name}}\n                        weight:\n                          type: integer\n                          format: int64\n                          x-dcl-go-name: Weight\n                          description: 'Optional. Specifies the proportion of requests\n                            forwareded to the backend referenced by the service_name\n                            field. This is computed as: weight/Sum(weights in destinations)\n                            Weights in all destinations does not need to sum up to\n                            100.'\n              matches:\n                type: array\n                x-dcl-go-name: Matches\n                description: Required. RouteMatch defines the predicate used to match\n                  requests to a given action. Multiple match types are \"OR\"ed for\n                  evaluation.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: TlsRouteRulesMatches\n                  properties:\n                    alpn:\n                      type: array\n                      x-dcl-go-name: Alpn\n                      description: 'Optional. ALPN (Application-Layer Protocol Negotiation)\n                        to match against. Examples: \"http/1.1\", \"h2\". At least one\n                        of sni_host and alpn is required. Up to 5 alpns across all\n                        matches can be set.'\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    sniHost:\n                      type: array\n                      x-dcl-go-name: SniHost\n                      description: Optional. SNI (server name indicator) to match\n                        against. SNI will be matched against all wildcard domains,\n                        i.e. www.example.com will be first matched against www.example.com,\n                        then *.example.com, then *.com. Partial wildcards are not\n                        supported, and values like *w.example.com are invalid. At\n                        least one of sni_host and alpn is required. Up to 5 sni hosts\n                        across all matches can be set.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Output only. Server-defined URL of this resource\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 9073 bytes
// MD5: 62362c700d906aeb27d90cf5ee1c8bde
