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
// gen_go_data -package assuredworkloads -var YAML_workload blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/assuredworkloads/workload.yaml

package assuredworkloads

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/assuredworkloads/workload.yaml
var YAML_workload = []byte("info:\n  title: AssuredWorkloads/Workload\n  description: The AssuredWorkloads Workload resource\n  x-dcl-struct-name: Workload\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Workload\n    parameters:\n    - name: workload\n      required: true\n      description: A full instance of a Workload\n  apply:\n    description: The function used to apply information about a Workload\n    parameters:\n    - name: workload\n      required: true\n      description: A full instance of a Workload\n  delete:\n    description: The function used to delete a Workload\n    parameters:\n    - name: workload\n      required: true\n      description: A full instance of a Workload\n  deleteAll:\n    description: The function used to delete all Workload\n    parameters:\n    - name: organization\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Workload\n    parameters:\n    - name: organization\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Workload:\n      title: Workload\n      x-dcl-id: organizations/{{organization}}/locations/{{location}}/workloads/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: organization\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - displayName\n      - complianceRegime\n      - organization\n      - location\n      properties:\n        billingAccount:\n          type: string\n          x-dcl-go-name: BillingAccount\n          description: Optional. Input only. The billing account used for the resources\n            which are direct children of workload. This billing account is initially\n            associated with the resources created as part of Workload creation. After\n            the initial creation of these resources, the customer can change the assigned\n            billing account. The resource name has the form `billingAccounts/{billing_account_id}`.\n            For example, `billingAccounts/012345-567890-ABCDEF`.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/BillingAccount\n            field: name\n          x-dcl-mutable-unreadable: true\n        complianceRegime:\n          type: string\n          x-dcl-go-name: ComplianceRegime\n          x-dcl-go-type: WorkloadComplianceRegimeEnum\n          description: 'Required. Immutable. Compliance Regime associated with this\n            workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH,\n            FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, HITRUST, EU_REGIONS_AND_SUPPORT,\n            CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS,\n            ISR_REGIONS, ISR_REGIONS_AND_SUPPORT, CA_PROTECTED_B, IL5, IL2, JP_REGIONS_AND_SUPPORT,\n            KSA_REGIONS_AND_SUPPORT_WITH_SOVEREIGNTY_CONTROLS, REGIONAL_CONTROLS,\n            HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS, HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS_WITH_US_SUPPORT,\n            IRS_1075'\n          x-kubernetes-immutable: true\n          enum:\n          - COMPLIANCE_REGIME_UNSPECIFIED\n          - IL4\n          - CJIS\n          - FEDRAMP_HIGH\n          - FEDRAMP_MODERATE\n          - US_REGIONAL_ACCESS\n          - HIPAA\n          - HITRUST\n          - EU_REGIONS_AND_SUPPORT\n          - CA_REGIONS_AND_SUPPORT\n          - ITAR\n          - AU_REGIONS_AND_US_SUPPORT\n          - ASSURED_WORKLOADS_FOR_PARTNERS\n          - ISR_REGIONS\n          - ISR_REGIONS_AND_SUPPORT\n          - CA_PROTECTED_B\n          - IL5\n          - IL2\n          - JP_REGIONS_AND_SUPPORT\n          - KSA_REGIONS_AND_SUPPORT_WITH_SOVEREIGNTY_CONTROLS\n          - REGIONAL_CONTROLS\n          - HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS\n          - HEALTHCARE_AND_LIFE_SCIENCES_CONTROLS_WITH_US_SUPPORT\n          - IRS_1075\n        complianceStatus:\n          type: object\n          x-dcl-go-name: ComplianceStatus\n          x-dcl-go-type: WorkloadComplianceStatus\n          readOnly: true\n          description: Output only. Count of active Violations in the Workload.\n          x-kubernetes-immutable: true\n          properties:\n            acknowledgedViolationCount:\n              type: array\n              x-dcl-go-name: AcknowledgedViolationCount\n              description: Number of current orgPolicy violations which are acknowledged.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: integer\n                format: int64\n                x-dcl-go-type: int64\n            activeViolationCount:\n              type: array\n              x-dcl-go-name: ActiveViolationCount\n              description: Number of current orgPolicy violations which are not acknowledged.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: integer\n                format: int64\n                x-dcl-go-type: int64\n        compliantButDisallowedServices:\n          type: array\n          x-dcl-go-name: CompliantButDisallowedServices\n          readOnly: true\n          description: Output only. Urls for services which are compliant for this\n            Assured Workload, but which are currently disallowed by the ResourceUsageRestriction\n            org policy. Invoke workloads.restrictAllowedResources endpoint to allow\n            your project developers to use these services in their environment.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Immutable. The Workload creation timestamp.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: 'Required. The user-assigned display name of the Workload.\n            When present it must be between 4 to 30 characters. Allowed characters\n            are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example:\n            My Workload'\n        ekmProvisioningResponse:\n          type: object\n          x-dcl-go-name: EkmProvisioningResponse\n          x-dcl-go-type: WorkloadEkmProvisioningResponse\n          readOnly: true\n          description: Optional. Represents the Ekm Provisioning State of the given\n            workload.\n          x-kubernetes-immutable: true\n          properties:\n            ekmProvisioningErrorDomain:\n              type: string\n              x-dcl-go-name: EkmProvisioningErrorDomain\n              x-dcl-go-type: WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum\n              description: 'Indicates Ekm provisioning error if any. Possible values:\n                EKM_PROVISIONING_ERROR_DOMAIN_UNSPECIFIED, UNSPECIFIED_ERROR, GOOGLE_SERVER_ERROR,\n                EXTERNAL_USER_ERROR, EXTERNAL_PARTNER_ERROR, TIMEOUT_ERROR'\n              x-kubernetes-immutable: true\n              enum:\n              - EKM_PROVISIONING_ERROR_DOMAIN_UNSPECIFIED\n              - UNSPECIFIED_ERROR\n              - GOOGLE_SERVER_ERROR\n              - EXTERNAL_USER_ERROR\n              - EXTERNAL_PARTNER_ERROR\n              - TIMEOUT_ERROR\n            ekmProvisioningErrorMapping:\n              type: string\n              x-dcl-go-name: EkmProvisioningErrorMapping\n              x-dcl-go-type: WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum\n              description: 'Detailed error message if Ekm provisioning fails Possible\n                values: EKM_PROVISIONING_ERROR_MAPPING_UNSPECIFIED, INVALID_SERVICE_ACCOUNT,\n                MISSING_METRICS_SCOPE_ADMIN_PERMISSION, MISSING_EKM_CONNECTION_ADMIN_PERMISSION'\n              x-kubernetes-immutable: true\n              enum:\n              - EKM_PROVISIONING_ERROR_MAPPING_UNSPECIFIED\n              - INVALID_SERVICE_ACCOUNT\n              - MISSING_METRICS_SCOPE_ADMIN_PERMISSION\n              - MISSING_EKM_CONNECTION_ADMIN_PERMISSION\n            ekmProvisioningState:\n              type: string\n              x-dcl-go-name: EkmProvisioningState\n              x-dcl-go-type: WorkloadEkmProvisioningResponseEkmProvisioningStateEnum\n              description: 'Indicates Ekm enrollment Provisioning of a given workload.\n                Possible values: EKM_PROVISIONING_STATE_UNSPECIFIED, EKM_PROVISIONING_STATE_PENDING,\n                EKM_PROVISIONING_STATE_FAILED, EKM_PROVISIONING_STATE_COMPLETED'\n              x-kubernetes-immutable: true\n              enum:\n              - EKM_PROVISIONING_STATE_UNSPECIFIED\n              - EKM_PROVISIONING_STATE_PENDING\n              - EKM_PROVISIONING_STATE_FAILED\n              - EKM_PROVISIONING_STATE_COMPLETED\n        enableSovereignControls:\n          type: boolean\n          x-dcl-go-name: EnableSovereignControls\n          description: Optional. Indicates the sovereignty status of the given workload.\n            Currently meant to be used by Europe/Canada customers.\n          x-kubernetes-immutable: true\n        kajEnrollmentState:\n          type: string\n          x-dcl-go-name: KajEnrollmentState\n          x-dcl-go-type: WorkloadKajEnrollmentStateEnum\n          readOnly: true\n          description: 'Output only. Represents the KAJ enrollment state of the given\n            workload. Possible values: KAJ_ENROLLMENT_STATE_UNSPECIFIED, KAJ_ENROLLMENT_STATE_PENDING,\n            KAJ_ENROLLMENT_STATE_COMPLETE'\n          x-kubernetes-immutable: true\n          enum:\n          - KAJ_ENROLLMENT_STATE_UNSPECIFIED\n          - KAJ_ENROLLMENT_STATE_PENDING\n          - KAJ_ENROLLMENT_STATE_COMPLETE\n        kmsSettings:\n          type: object\n          x-dcl-go-name: KmsSettings\n          x-dcl-go-type: WorkloadKmsSettings\n          description: '**DEPRECATED** Input only. Settings used to create a CMEK\n            crypto key. When set, a project with a KMS CMEK key is provisioned. This\n            field is deprecated as of Feb 28, 2022. In order to create a Keyring,\n            callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type\n            field.'\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n          required:\n          - nextRotationTime\n          - rotationPeriod\n          properties:\n            nextRotationTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: NextRotationTime\n              description: Required. Input only. Immutable. The time at which the\n                Key Management Service will automatically create a new version of\n                the crypto key and mark it as the primary.\n              x-kubernetes-immutable: true\n            rotationPeriod:\n              type: string\n              x-dcl-go-name: RotationPeriod\n              description: Required. Input only. Immutable. will be advanced by this\n                period when the Key Management Service automatically rotates a key.\n                Must be at least 24 hours and at most 876,000 hours.\n              x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Labels applied to the workload.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the workload.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n          x-dcl-has-long-form: true\n        organization:\n          type: string\n          x-dcl-go-name: Organization\n          description: The organization for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        partner:\n          type: string\n          x-dcl-go-name: Partner\n          x-dcl-go-type: WorkloadPartnerEnum\n          description: 'Optional. Partner regime associated with this workload. Possible\n            values: PARTNER_UNSPECIFIED, LOCAL_CONTROLS_BY_S3NS, SOVEREIGN_CONTROLS_BY_T_SYSTEMS,\n            SOVEREIGN_CONTROLS_BY_SIA_MINSAIT, SOVEREIGN_CONTROLS_BY_PSN, SOVEREIGN_CONTROLS_BY_CNTXT,\n            SOVEREIGN_CONTROLS_BY_CNTXT_NO_EKM'\n          x-kubernetes-immutable: true\n          enum:\n          - PARTNER_UNSPECIFIED\n          - LOCAL_CONTROLS_BY_S3NS\n          - SOVEREIGN_CONTROLS_BY_T_SYSTEMS\n          - SOVEREIGN_CONTROLS_BY_SIA_MINSAIT\n          - SOVEREIGN_CONTROLS_BY_PSN\n          - SOVEREIGN_CONTROLS_BY_CNTXT\n          - SOVEREIGN_CONTROLS_BY_CNTXT_NO_EKM\n        partnerPermissions:\n          type: object\n          x-dcl-go-name: PartnerPermissions\n          x-dcl-go-type: WorkloadPartnerPermissions\n          description: Optional. Permissions granted to the AW Partner SA account\n            for the customer workload\n          x-kubernetes-immutable: true\n          properties:\n            assuredWorkloadsMonitoring:\n              type: boolean\n              x-dcl-go-name: AssuredWorkloadsMonitoring\n              description: Optional. Allow partner to view violation alerts.\n              x-kubernetes-immutable: true\n            dataLogsViewer:\n              type: boolean\n              x-dcl-go-name: DataLogsViewer\n              description: Allow the partner to view inspectability logs and monitoring\n                violations.\n              x-kubernetes-immutable: true\n            serviceAccessApprover:\n              type: boolean\n              x-dcl-go-name: ServiceAccessApprover\n              description: Optional. Allow partner to view access approval logs.\n              x-kubernetes-immutable: true\n        partnerServicesBillingAccount:\n          type: string\n          x-dcl-go-name: PartnerServicesBillingAccount\n          description: Optional. Input only. Billing account necessary for purchasing\n            services from Sovereign Partners. This field is required for creating\n            SIA/PSN/CNTXT partner workloads. The caller should have 'billing.resourceAssociations.create'\n            IAM permission on this billing-account. The format of this string is billingAccounts/AAAAAA-BBBBBB-CCCCCC.\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n        provisionedResourcesParent:\n          type: string\n          x-dcl-go-name: ProvisionedResourcesParent\n          description: 'Input only. The parent resource for the resources managed\n            by this Assured Workload. May be either empty or a folder resource which\n            is a child of the Workload parent. If not specified all resources are\n            created under the parent organization. Format: folders/{folder_id}'\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n        resourceSettings:\n          type: array\n          x-dcl-go-name: ResourceSettings\n          description: Input only. Resource properties that are used to customize\n            workload resources. These properties (such as custom project id) will\n            be used to create workload resources if possible. This field is optional.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: WorkloadResourceSettings\n            properties:\n              displayName:\n                type: string\n                x-dcl-go-name: DisplayName\n                description: User-assigned resource display name. If not empty it\n                  will be used to create a resource with the specified name.\n                x-kubernetes-immutable: true\n              resourceId:\n                type: string\n                x-dcl-go-name: ResourceId\n                description: Resource identifier. For a project this represents projectId.\n                  If the project is already taken, the workload creation will fail.\n                  For KeyRing, this represents the keyring_id. For a folder, don't\n                  set this value as folder_id is assigned by Google.\n                x-kubernetes-immutable: true\n              resourceType:\n                type: string\n                x-dcl-go-name: ResourceType\n                x-dcl-go-type: WorkloadResourceSettingsResourceTypeEnum\n                description: 'Indicates the type of resource. This field should be\n                  specified to correspond the id to the right project type (CONSUMER_PROJECT\n                  or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED,\n                  CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER'\n                x-kubernetes-immutable: true\n                enum:\n                - RESOURCE_TYPE_UNSPECIFIED\n                - CONSUMER_PROJECT\n                - ENCRYPTION_KEYS_PROJECT\n                - KEYRING\n                - CONSUMER_FOLDER\n          x-dcl-mutable-unreadable: true\n        resources:\n          type: array\n          x-dcl-go-name: Resources\n          readOnly: true\n          description: Output only. The resources associated with this workload. These\n            resources will be created when creating the workload. If any of the projects\n            already exist, the workload creation will fail. Always read only.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: WorkloadResources\n            properties:\n              resourceId:\n                type: integer\n                format: int64\n                x-dcl-go-name: ResourceId\n                description: Resource identifier. For a project this represents project_number.\n                x-kubernetes-immutable: true\n              resourceType:\n                type: string\n                x-dcl-go-name: ResourceType\n                x-dcl-go-type: WorkloadResourcesResourceTypeEnum\n                description: 'Indicates the type of resource. Possible values: RESOURCE_TYPE_UNSPECIFIED,\n                  CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER'\n                x-kubernetes-immutable: true\n                enum:\n                - RESOURCE_TYPE_UNSPECIFIED\n                - CONSUMER_PROJECT\n                - ENCRYPTION_KEYS_PROJECT\n                - KEYRING\n                - CONSUMER_FOLDER\n        saaEnrollmentResponse:\n          type: object\n          x-dcl-go-name: SaaEnrollmentResponse\n          x-dcl-go-type: WorkloadSaaEnrollmentResponse\n          readOnly: true\n          description: Output only. Represents the SAA enrollment response of the\n            given workload. SAA enrollment response is queried during workloads.get\n            call. In failure cases, user friendly error message is shown in SAA details\n            page.\n          x-kubernetes-immutable: true\n          properties:\n            setupErrors:\n              type: array\n              x-dcl-go-name: SetupErrors\n              description: Indicates SAA enrollment setup error if any.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: WorkloadSaaEnrollmentResponseSetupErrorsEnum\n                enum:\n                - SETUP_ERROR_UNSPECIFIED\n                - ERROR_INVALID_BASE_SETUP\n                - ERROR_MISSING_EXTERNAL_SIGNING_KEY\n                - ERROR_NOT_ALL_SERVICES_ENROLLED\n                - ERROR_SETUP_CHECK_FAILED\n            setupStatus:\n              type: string\n              x-dcl-go-name: SetupStatus\n              x-dcl-go-type: WorkloadSaaEnrollmentResponseSetupStatusEnum\n              description: 'Indicates SAA enrollment status of a given workload. Possible\n                values: SETUP_STATE_UNSPECIFIED, STATUS_PENDING, STATUS_COMPLETE'\n              x-kubernetes-immutable: true\n              enum:\n              - SETUP_STATE_UNSPECIFIED\n              - STATUS_PENDING\n              - STATUS_COMPLETE\n        violationNotificationsEnabled:\n          type: boolean\n          x-dcl-go-name: ViolationNotificationsEnabled\n          description: Optional. Indicates whether the e-mail notification for a violation\n            is enabled for a workload. This value will be by default True, and if\n            not present will be considered as true. This should only be updated via\n            updateWorkload call. Any Changes to this field during the createWorkload\n            call will not be honored. This will always be true while creating the\n            workload.\n          x-kubernetes-immutable: true\n        workloadOptions:\n          type: object\n          x-dcl-go-name: WorkloadOptions\n          x-dcl-go-type: WorkloadWorkloadOptions\n          description: Optional. Used to specify certain options for a workload during\n            workload creation - currently only supporting KAT Optionality for Regional\n            Controls workloads.\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n          properties:\n            kajEnrollmentType:\n              type: string\n              x-dcl-go-name: KajEnrollmentType\n              x-dcl-go-type: WorkloadWorkloadOptionsKajEnrollmentTypeEnum\n              description: 'Indicates type of KAJ enrollment for the workload. Currently,\n                only specifiying KEY_ACCESS_TRANSPARENCY_OFF is implemented to not\n                enroll in KAT-level KAJ enrollment for Regional Controls workloads.\n                Possible values: KAJ_ENROLLMENT_TYPE_UNSPECIFIED, FULL_KAJ, EKM_ONLY,\n                KEY_ACCESS_TRANSPARENCY_OFF'\n              x-kubernetes-immutable: true\n              enum:\n              - KAJ_ENROLLMENT_TYPE_UNSPECIFIED\n              - FULL_KAJ\n              - EKM_ONLY\n              - KEY_ACCESS_TRANSPARENCY_OFF\n")

// 22584 bytes
// MD5: 441cb7898bdacdf5855bb10eaeac6d9e
