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
// gen_go_data -package orgpolicy -var YAML_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/orgpolicy/policy.yaml

package orgpolicy

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/orgpolicy/policy.yaml
var YAML_policy = []byte("info:\n  title: OrgPolicy/Policy\n  description: An organization policy gives you programmatic control over your organization's\n    cloud resources.  Using Organization Policies, you will be able to configure constraints\n    across your entire resource hierarchy.\n  x-dcl-struct-name: Policy\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: REST API\n    url: https://cloud.google.com/resource-manager/docs/reference/orgpolicy/rest/v2/organizations.policies\n  x-dcl-guides:\n  - text: Understanding Org Policy concepts\n    url: https://cloud.google.com/resource-manager/docs/organization-policy/overview\n  - text: The resource hierarchy\n    url: https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy\n  - text: All valid constraints\n    url: https://cloud.google.com/resource-manager/docs/organization-policy/org-policy-constraints\npaths:\n  get:\n    description: The function used to get information about a Policy\n    parameters:\n    - name: Policy\n      required: true\n      description: A full instance of a Policy\n  apply:\n    description: The function used to apply information about a Policy\n    parameters:\n    - name: Policy\n      required: true\n      description: A full instance of a Policy\n  delete:\n    description: The function used to delete a Policy\n    parameters:\n    - name: Policy\n      required: true\n      description: A full instance of a Policy\n  deleteAll:\n    description: The function used to delete all Policy\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Policy\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Policy:\n      title: Policy\n      x-dcl-id: '{{parent}}/policies/{{name}}'\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - parent\n      properties:\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Immutable. The resource name of the Policy. Must be one of\n            the following forms, where constraint_name is the name of the constraint\n            which this Policy configures: * `projects/{project_number}/policies/{constraint_name}`\n            * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}`\n            For example, \"projects/123/policies/compute.disableSerialPortAccess\".\n            Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable\n            name for API requests, but responses will return the name using the equivalent\n            project number.'\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: The parent of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Folder\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        spec:\n          type: object\n          x-dcl-go-name: Spec\n          x-dcl-go-type: PolicySpec\n          description: Basic information about the Organization Policy.\n          properties:\n            etag:\n              type: string\n              x-dcl-go-name: Etag\n              readOnly: true\n              description: An opaque tag indicating the current version of the `Policy`,\n                used for concurrency control. This field is ignored if used in a `CreatePolicy`\n                request. When the `Policy` is returned from either a `GetPolicy` or\n                a `ListPolicies` request, this `etag` indicates the version of the\n                current `Policy` to use when executing a read-modify-write loop. When\n                the `Policy` is returned from a `GetEffectivePolicy` request, the\n                `etag` will be unset.\n            inheritFromParent:\n              type: boolean\n              x-dcl-go-name: InheritFromParent\n              description: Determines the inheritance behavior for this `Policy`.\n                If `inherit_from_parent` is true, PolicyRules set higher up in the\n                hierarchy (up to the closest root) are inherited and present in the\n                effective policy. If it is false, then no rules are inherited, and\n                this Policy becomes the new root for evaluation. This field can be\n                set only for Policies which configure list constraints.\n            reset:\n              type: boolean\n              x-dcl-go-name: Reset\n              description: Ignores policies set above this resource and restores the\n                `constraint_default` enforcement behavior of the specific `Constraint`\n                at this resource. This field can be set in policies for either list\n                or boolean constraints. If set, `rules` must be empty and `inherit_from_parent`\n                must be set to false.\n            rules:\n              type: array\n              x-dcl-go-name: Rules\n              description: 'Up to 10 PolicyRules are allowed. In Policies for boolean\n                constraints, the following requirements apply: - There must be one\n                and only one PolicyRule where condition is unset. - BooleanPolicyRules\n                with conditions must set `enforced` to the opposite of the PolicyRule\n                without a condition. - During policy evaluation, PolicyRules with\n                conditions that are true for a target resource take precedence.'\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: PolicySpecRules\n                properties:\n                  allowAll:\n                    type: boolean\n                    x-dcl-go-name: AllowAll\n                    description: Setting this to true means that all values are allowed.\n                      This field can be set only in Policies for list constraints.\n                    x-dcl-conflicts:\n                    - values\n                    - denyAll\n                    - enforce\n                  condition:\n                    type: object\n                    x-dcl-go-name: Condition\n                    x-dcl-go-type: PolicySpecRulesCondition\n                    description: 'A condition which determines whether this rule is\n                      used in the evaluation of the policy. When set, the `expression`\n                      field in the `Expr'' must include from 1 to 10 subexpressions,\n                      joined by the \"||\" or \"&&\" operators. Each subexpression must\n                      be of the form \"resource.matchTag(''/tag_key_short_name, ''tag_value_short_name'')\".\n                      or \"resource.matchTagId(''tagKeys/key_id'', ''tagValues/value_id'')\".\n                      where key_name and value_name are the resource names for Label\n                      Keys and Values. These names are available from the Tag Manager\n                      Service. An example expression is: \"resource.matchTag(''123456789/environment,\n                      ''prod'')\". or \"resource.matchTagId(''tagKeys/123'', ''tagValues/456'')\".'\n                    properties:\n                      description:\n                        type: string\n                        x-dcl-go-name: Description\n                        description: Optional. Description of the expression. This\n                          is a longer text which describes the expression, e.g. when\n                          hovered over it in a UI.\n                      expression:\n                        type: string\n                        x-dcl-go-name: Expression\n                        description: Textual representation of an expression in Common\n                          Expression Language syntax.\n                      location:\n                        type: string\n                        x-dcl-go-name: Location\n                        description: Optional. String indicating the location of the\n                          expression for error reporting, e.g. a file name and a position\n                          in the file.\n                      title:\n                        type: string\n                        x-dcl-go-name: Title\n                        description: Optional. Title for the expression, i.e. a short\n                          string describing its purpose. This can be used e.g. in\n                          UIs which allow to enter the expression.\n                  denyAll:\n                    type: boolean\n                    x-dcl-go-name: DenyAll\n                    description: Setting this to true means that all values are denied.\n                      This field can be set only in Policies for list constraints.\n                    x-dcl-conflicts:\n                    - values\n                    - allowAll\n                    - enforce\n                  enforce:\n                    type: boolean\n                    x-dcl-go-name: Enforce\n                    description: If `true`, then the `Policy` is enforced. If `false`,\n                      then any configuration is acceptable. This field can be set\n                      only in Policies for boolean constraints.\n                    x-dcl-conflicts:\n                    - values\n                    - allowAll\n                    - denyAll\n                  values:\n                    type: object\n                    x-dcl-go-name: Values\n                    x-dcl-go-type: PolicySpecRulesValues\n                    description: List of values to be used for this PolicyRule. This\n                      field can be set only in Policies for list constraints.\n                    x-dcl-conflicts:\n                    - allowAll\n                    - denyAll\n                    - enforce\n                    properties:\n                      allowedValues:\n                        type: array\n                        x-dcl-go-name: AllowedValues\n                        description: List of values allowed at this resource.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      deniedValues:\n                        type: array\n                        x-dcl-go-name: DeniedValues\n                        description: List of values denied at this resource.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              readOnly: true\n              description: Output only. The time stamp this was previously updated.\n                This represents the last time a call to `CreatePolicy` or `UpdatePolicy`\n                was made for that `Policy`.\n")

// 11333 bytes
// MD5: ff5acdcb9b6c77da038e193849d10e32
