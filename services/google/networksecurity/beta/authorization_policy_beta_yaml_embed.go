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
// gen_go_data -package beta -var YAML_authorization_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networksecurity/beta/authorization_policy.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networksecurity/beta/authorization_policy.yaml
var YAML_authorization_policy = []byte("info:\n  title: NetworkSecurity/AuthorizationPolicy\n  description: The NetworkSecurity AuthorizationPolicy resource\n  x-dcl-struct-name: AuthorizationPolicy\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a AuthorizationPolicy\n    parameters:\n    - name: authorizationPolicy\n      required: true\n      description: A full instance of a AuthorizationPolicy\n  apply:\n    description: The function used to apply information about a AuthorizationPolicy\n    parameters:\n    - name: authorizationPolicy\n      required: true\n      description: A full instance of a AuthorizationPolicy\n  delete:\n    description: The function used to delete a AuthorizationPolicy\n    parameters:\n    - name: authorizationPolicy\n      required: true\n      description: A full instance of a AuthorizationPolicy\n  deleteAll:\n    description: The function used to delete all AuthorizationPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many AuthorizationPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    AuthorizationPolicy:\n      title: AuthorizationPolicy\n      x-dcl-id: projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - action\n      - project\n      - location\n      properties:\n        action:\n          type: string\n          x-dcl-go-name: Action\n          x-dcl-go-type: AuthorizationPolicyActionEnum\n          description: 'Required. The action to take when a rule match is found. Possible\n            values are \"ALLOW\" or \"DENY\". Possible values: ACTION_UNSPECIFIED, ALLOW,\n            DENY'\n          enum:\n          - ACTION_UNSPECIFIED\n          - ALLOW\n          - DENY\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Free-text description of the resource.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the AuthorizationPolicy\n            resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the AuthorizationPolicy resource.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        rules:\n          type: array\n          x-dcl-go-name: Rules\n          description: Optional. List of rules to match. If not set, the action specified\n            in the ‘action’ field will be applied without any additional rule checks.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: AuthorizationPolicyRules\n            properties:\n              destinations:\n                type: array\n                x-dcl-go-name: Destinations\n                description: Optional. List of attributes for the traffic destination.\n                  If not set, the action specified in the ‘action’ field will be applied\n                  without any rule checks for the destination.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: AuthorizationPolicyRulesDestinations\n                  required:\n                  - hosts\n                  - ports\n                  properties:\n                    hosts:\n                      type: array\n                      x-dcl-go-name: Hosts\n                      description: Required. List of host names to match. Matched\n                        against HOST header in http requests. Each host can be an\n                        exact match, or a prefix match (example, “mydomain.*”) or\n                        a suffix match (example, *.myorg.com”) or a presence(any)\n                        match “*”.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    httpHeaderMatch:\n                      type: object\n                      x-dcl-go-name: HttpHeaderMatch\n                      x-dcl-go-type: AuthorizationPolicyRulesDestinationsHttpHeaderMatch\n                      description: Optional. Match against key:value pair in http\n                        header. Provides a flexible match based on HTTP headers, for\n                        potentially advanced use cases.\n                      required:\n                      - headerName\n                      - regexMatch\n                      properties:\n                        headerName:\n                          type: string\n                          x-dcl-go-name: HeaderName\n                          description: Required. The name of the HTTP header to match.\n                            For matching against the HTTP request's authority, use\n                            a headerMatch with the header name \":authority\". For matching\n                            a request's method, use the headerName \":method\".\n                        regexMatch:\n                          type: string\n                          x-dcl-go-name: RegexMatch\n                          description: 'Required. The value of the header must match\n                            the regular expression specified in regexMatch. For regular\n                            expression grammar, please see: en.cppreference.com/w/cpp/regex/ecmascript\n                            For matching against a port specified in the HTTP request,\n                            use a headerMatch with headerName set to Host and a regular\n                            expression that satisfies the RFC2616 Host header''s port\n                            specifier.'\n                    methods:\n                      type: array\n                      x-dcl-go-name: Methods\n                      description: Optional. A list of HTTP methods to match. Should\n                        not be set for gRPC services.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    ports:\n                      type: array\n                      x-dcl-go-name: Ports\n                      description: Required. List of destination ports to match.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: integer\n                        format: int64\n                        x-dcl-go-type: int64\n              sources:\n                type: array\n                x-dcl-go-name: Sources\n                description: Optional. List of attributes for the traffic source.\n                  If not set, the action specified in the ‘action’ field will be applied\n                  without any rule checks for the source.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: AuthorizationPolicyRulesSources\n                  properties:\n                    ipBlocks:\n                      type: array\n                      x-dcl-go-name: IPBlocks\n                      description: Optional. List of CIDR ranges to match based on\n                        source IP address. Single IP (e.g., \"1.2.3.4\") and CIDR (e.g.,\n                        \"1.2.3.0/24\") are supported.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    principals:\n                      type: array\n                      x-dcl-go-name: Principals\n                      description: Optional. List of peer identities to match for\n                        authorization. Each peer can be an exact match, or a prefix\n                        match (example, “namespace/*”) or a suffix match (example,\n                        */service-account”) or a presence match “*”.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 9881 bytes
// MD5: 6069d50715b1ba895ed5e460760e4fc6
