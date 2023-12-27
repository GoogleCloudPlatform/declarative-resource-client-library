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
// gen_go_data -package alpha -var YAML_tenant blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/alpha/tenant.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/alpha/tenant.yaml
var YAML_tenant = []byte("info:\n  title: IdentityToolkit/Tenant\n  description: The IdentityToolkit Tenant resource\n  x-dcl-struct-name: Tenant\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Tenant\n    parameters:\n    - name: tenant\n      required: true\n      description: A full instance of a Tenant\n  apply:\n    description: The function used to apply information about a Tenant\n    parameters:\n    - name: tenant\n      required: true\n      description: A full instance of a Tenant\n  delete:\n    description: The function used to delete a Tenant\n    parameters:\n    - name: tenant\n      required: true\n      description: A full instance of a Tenant\n  deleteAll:\n    description: The function used to delete all Tenant\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Tenant\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Tenant:\n      title: Tenant\n      x-dcl-id: projects/{{project}}/tenants/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      properties:\n        allowPasswordSignup:\n          type: boolean\n          x-dcl-go-name: AllowPasswordSignup\n          description: Whether to allow email/password user authentication.\n        disableAuth:\n          type: boolean\n          x-dcl-go-name: DisableAuth\n          description: Whether authentication is disabled for the tenant. If true,\n            the users under the disabled tenant are not allowed to sign-in. Admins\n            of the disabled tenant are not able to manage its users.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Display name of the tenant.\n        enableAnonymousUser:\n          type: boolean\n          x-dcl-go-name: EnableAnonymousUser\n          description: Whether to enable anonymous user authentication.\n        enableEmailLinkSignin:\n          type: boolean\n          x-dcl-go-name: EnableEmailLinkSignin\n          description: Whether to enable email link user authentication.\n        mfaConfig:\n          type: object\n          x-dcl-go-name: MfaConfig\n          x-dcl-go-type: TenantMfaConfig\n          description: The tenant-level configuration of MFA options.\n          properties:\n            enabledProviders:\n              type: array\n              x-dcl-go-name: EnabledProviders\n              description: A list of usable second factors for this project.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: TenantMfaConfigEnabledProvidersEnum\n                enum:\n                - PROVIDER_UNSPECIFIED\n                - PHONE_SMS\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: TenantMfaConfigStateEnum\n              description: 'Whether MultiFactor Authentication has been enabled for\n                this project. Possible values: STATE_UNSPECIFIED, DISABLED, ENABLED,\n                MANDATORY'\n              enum:\n              - STATE_UNSPECIFIED\n              - DISABLED\n              - ENABLED\n              - MANDATORY\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. Resource name of a tenant. For example: \"projects/{project-id}/tenants/{tenant-id}\"'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        testPhoneNumbers:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: TestPhoneNumbers\n          description: A map of <test phone number, fake code> pairs that can be used\n            for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/)\n            and a maximum of 10 pairs can be added (error will be thrown once exceeded).\n")

// 4460 bytes
// MD5: aa6af3a4219d535093bb7b5e7d189eec
