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
// gen_go_data -package beta -var YAML_group blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudidentity/beta/group.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudidentity/beta/group.yaml
var YAML_group = []byte("info:\n  title: Cloudidentity/Group\n  description: The Cloudidentity Group resource\n  x-dcl-struct-name: Group\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Group\n    parameters:\n    - name: Group\n      required: true\n      description: A full instance of a Group\n  apply:\n    description: The function used to apply information about a Group\n    parameters:\n    - name: Group\n      required: true\n      description: A full instance of a Group\n  delete:\n    description: The function used to delete a Group\n    parameters:\n    - name: Group\n      required: true\n      description: A full instance of a Group\n  deleteAll:\n    description: The function used to delete all Group\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Group\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Group:\n      title: Group\n      x-dcl-id: groups/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - groupKey\n      - parent\n      - labels\n      properties:\n        additionalGroupKeys:\n          type: array\n          x-dcl-go-name: AdditionalGroupKeys\n          description: Optional. Additional entity key aliases for a Group.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: GroupAdditionalGroupKeys\n            required:\n            - id\n            properties:\n              id:\n                type: string\n                x-dcl-go-name: Id\n                description: The ID of the entity. For Google-managed entities, the\n                  `id` must be the email address of a group or user. For external-identity-mapped\n                  entities, the `id` must be a string conforming to the Identity Source's\n                  requirements. Must be unique within a `namespace`.\n                x-kubernetes-immutable: true\n              namespace:\n                type: string\n                x-dcl-go-name: Namespace\n                description: The namespace in which the entity exists. If not specified,\n                  the `EntityKey` represents a Google-managed entity such as a Google\n                  user or a Google Group. If specified, the `EntityKey` represents\n                  an external-identity-mapped group. The namespace must correspond\n                  to an identity source created in Admin Console and must be in the\n                  form of `identitysources/{identity_source_id}`.\n                x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the `Group` was created.\n          x-kubernetes-immutable: true\n        derivedAliases:\n          type: array\n          x-dcl-go-name: DerivedAliases\n          readOnly: true\n          description: Output only. Aliases for groups derived from domain aliases.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: GroupDerivedAliases\n            required:\n            - id\n            properties:\n              id:\n                type: string\n                x-dcl-go-name: Id\n                description: The ID of the entity. For Google-managed entities, the\n                  `id` must be the email address of a group or user. For external-identity-mapped\n                  entities, the `id` must be a string conforming to the Identity Source's\n                  requirements. Must be unique within a `namespace`.\n                x-kubernetes-immutable: true\n              namespace:\n                type: string\n                x-dcl-go-name: Namespace\n                description: The namespace in which the entity exists. If not specified,\n                  the `EntityKey` represents a Google-managed entity such as a Google\n                  user or a Google Group. If specified, the `EntityKey` represents\n                  an external-identity-mapped group. The namespace must correspond\n                  to an identity source created in Admin Console and must be in the\n                  form of `identitysources/{identity_source_id}`.\n                x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An extended description to help users determine the purpose\n            of a `Group`. Must not be longer than 4,096 characters.\n        directMemberCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: DirectMemberCount\n          readOnly: true\n          description: 'Output only. The number of all direct members. Including groups\n            and users, The special member: all-user-in-domain will be counted as one\n            member. Output only.'\n          x-kubernetes-immutable: true\n        directMemberCountPerType:\n          type: object\n          x-dcl-go-name: DirectMemberCountPerType\n          x-dcl-go-type: GroupDirectMemberCountPerType\n          readOnly: true\n          description: Output only. Direct membership counts grouped by user/group\n            type\n          x-kubernetes-immutable: true\n          properties:\n            groupCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: GroupCount\n              readOnly: true\n              description: Output only. Direct group type membership count\n              x-kubernetes-immutable: true\n            userCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: UserCount\n              readOnly: true\n              description: Output only. Direct user type membership count\n              x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: The display name of the `Group`.\n        dynamicGroupMetadata:\n          type: object\n          x-dcl-go-name: DynamicGroupMetadata\n          x-dcl-go-type: GroupDynamicGroupMetadata\n          description: Optional. Dynamic group metadata like queries and status.\n          properties:\n            queries:\n              type: array\n              x-dcl-go-name: Queries\n              description: Only one entry is supported for now. Memberships will be\n                the union of all queries. Customers can create up to 100 dynamic groups.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: GroupDynamicGroupMetadataQueries\n                properties:\n                  query:\n                    type: string\n                    x-dcl-go-name: Query\n                    description: Query that determines the memberships of the dynamic\n                      group.\n                  resourceType:\n                    type: string\n                    x-dcl-go-name: ResourceType\n                    x-dcl-go-type: GroupDynamicGroupMetadataQueriesResourceTypeEnum\n                    description: ' Possible values: RESOURCE_TYPE_UNSPECIFIED, USER'\n                    enum:\n                    - RESOURCE_TYPE_UNSPECIFIED\n                    - USER\n            status:\n              type: object\n              x-dcl-go-name: Status\n              x-dcl-go-type: GroupDynamicGroupMetadataStatus\n              readOnly: true\n              description: Status of the dynamic group. Output only.\n              properties:\n                status:\n                  type: string\n                  x-dcl-go-name: Status\n                  x-dcl-go-type: GroupDynamicGroupMetadataStatusStatusEnum\n                  description: 'Status of the dynamic group. Possible values: STATUS_UNSPECIFIED,\n                    UP_TO_DATE, UPDATING_MEMBERSHIPS, INVALID_QUERY'\n                  enum:\n                  - STATUS_UNSPECIFIED\n                  - UP_TO_DATE\n                  - UPDATING_MEMBERSHIPS\n                  - INVALID_QUERY\n                statusTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: StatusTime\n                  description: 'The latest time at which the dynamic group is guaranteed\n                    to be in the given status. For example, if status is: UP_TO_DATE\n                    - The latest time at which this dynamic group was confirmed to\n                    be up to date. UPDATING_MEMBERSHIPS - The time at which dynamic\n                    group was created.'\n        groupKey:\n          type: object\n          x-dcl-go-name: GroupKey\n          x-dcl-go-type: GroupGroupKey\n          description: Required. Immutable. The `EntityKey` of the `Group`.\n          x-kubernetes-immutable: true\n          required:\n          - id\n          properties:\n            id:\n              type: string\n              x-dcl-go-name: Id\n              description: The ID of the entity. For Google-managed entities, the\n                `id` must be the email address of a group or user. For external-identity-mapped\n                entities, the `id` must be a string conforming to the Identity Source's\n                requirements. Must be unique within a `namespace`.\n              x-kubernetes-immutable: true\n            namespace:\n              type: string\n              x-dcl-go-name: Namespace\n              description: The namespace in which the entity exists. If not specified,\n                the `EntityKey` represents a Google-managed entity such as a Google\n                user or a Google Group. If specified, the `EntityKey` represents an\n                external-identity-mapped group. The namespace must correspond to an\n                identity source created in Admin Console and must be in the form of\n                `identitysources/{identity_source_id}`.\n              x-kubernetes-immutable: true\n        initialGroupConfig:\n          type: string\n          x-dcl-go-name: InitialGroupConfig\n          x-dcl-go-type: GroupInitialGroupConfigEnum\n          description: 'The initial configuration option for the `Group`. Possible\n            values: INITIAL_GROUP_CONFIG_UNSPECIFIED, WITH_INITIAL_OWNER, EMPTY'\n          x-kubernetes-immutable: true\n          enum:\n          - INITIAL_GROUP_CONFIG_UNSPECIFIED\n          - WITH_INITIAL_OWNER\n          - EMPTY\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Required. One or more label entries that apply to the Group.\n            Currently supported labels contain a key with an empty value. Google Groups\n            are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum`\n            and an empty value. Existing Google Groups can have an additional label\n            with a key of `cloudidentity.googleapis.com/groups.security` and an empty\n            value added to them. **This is an immutable change and the security label\n            cannot be removed once added.** Dynamic groups have a label with a key\n            of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups\n            for Cloud Search have a label with a key of `system/groups/external` and\n            an empty value.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The [resource name](https://cloud.google.com/apis/design/resource_names)\n            of the `Group`. Shall be of the form `groups/{group}`.\n          x-dcl-server-generated-parameter: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: Required. Immutable. The resource name of the entity under\n            which this `Group` resides in the Cloud Identity resource hierarchy. Must\n            be of the form `identitysources/{identity_source}` for external- identity-mapped\n            groups or `customers/{customer}` for Google Groups. The `customer` must\n            begin with \"C\" (for example, 'C046psxkn').\n          x-kubernetes-immutable: true\n        posixGroups:\n          type: array\n          x-dcl-go-name: PosixGroups\n          description: The POSIX groups associated with the `Group`.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: GroupPosixGroups\n            properties:\n              gid:\n                type: string\n                x-dcl-go-name: Gid\n                description: GID of the POSIX group.\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: Name of the POSIX group.\n              systemId:\n                type: string\n                x-dcl-go-name: SystemId\n                description: System identifier for which group name and gid apply\n                  to. If not specified it will default to empty value.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time when the `Group` was last updated.\n          x-kubernetes-immutable: true\n")

// 13524 bytes
// MD5: 0df46f0818020b5622533738c91afd2b
