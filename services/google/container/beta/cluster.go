// Copyright 2021 Google LLC. All Rights Reserved.
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
package beta

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Cluster struct {
	Name                           *string                                `json:"name"`
	Description                    *string                                `json:"description"`
	InitialNodeCount               *int64                                 `json:"initialNodeCount"`
	MasterAuth                     *ClusterMasterAuth                     `json:"masterAuth"`
	LoggingService                 *string                                `json:"loggingService"`
	MonitoringService              *string                                `json:"monitoringService"`
	Network                        *string                                `json:"network"`
	ClusterIPv4Cidr                *string                                `json:"clusterIPv4Cidr"`
	AddonsConfig                   *ClusterAddonsConfig                   `json:"addonsConfig"`
	Subnetwork                     *string                                `json:"subnetwork"`
	NodePools                      []ClusterNodePools                     `json:"nodePools"`
	Locations                      []string                               `json:"locations"`
	EnableKubernetesAlpha          *bool                                  `json:"enableKubernetesAlpha"`
	ResourceLabels                 map[string]string                      `json:"resourceLabels"`
	LabelFingerprint               *string                                `json:"labelFingerprint"`
	LegacyAbac                     *ClusterLegacyAbac                     `json:"legacyAbac"`
	NetworkPolicy                  *ClusterNetworkPolicy                  `json:"networkPolicy"`
	IPAllocationPolicy             *ClusterIPAllocationPolicy             `json:"ipAllocationPolicy"`
	MasterAuthorizedNetworksConfig *ClusterMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig"`
	BinaryAuthorization            *ClusterBinaryAuthorization            `json:"binaryAuthorization"`
	Autoscaling                    *ClusterAutoscaling                    `json:"autoscaling"`
	NetworkConfig                  *ClusterNetworkConfig                  `json:"networkConfig"`
	MaintenancePolicy              *ClusterMaintenancePolicy              `json:"maintenancePolicy"`
	DefaultMaxPodsConstraint       *ClusterDefaultMaxPodsConstraint       `json:"defaultMaxPodsConstraint"`
	ResourceUsageExportConfig      *ClusterResourceUsageExportConfig      `json:"resourceUsageExportConfig"`
	AuthenticatorGroupsConfig      *ClusterAuthenticatorGroupsConfig      `json:"authenticatorGroupsConfig"`
	PrivateClusterConfig           *ClusterPrivateClusterConfig           `json:"privateClusterConfig"`
	DatabaseEncryption             *ClusterDatabaseEncryption             `json:"databaseEncryption"`
	VerticalPodAutoscaling         *ClusterVerticalPodAutoscaling         `json:"verticalPodAutoscaling"`
	ShieldedNodes                  *ClusterShieldedNodes                  `json:"shieldedNodes"`
	Endpoint                       *string                                `json:"endpoint"`
	MasterVersion                  *string                                `json:"masterVersion"`
	CreateTime                     *string                                `json:"createTime"`
	Status                         *string                                `json:"status"`
	StatusMessage                  *string                                `json:"statusMessage"`
	NodeIPv4CidrSize               *int64                                 `json:"nodeIPv4CidrSize"`
	ServicesIPv4Cidr               *string                                `json:"servicesIPv4Cidr"`
	ExpireTime                     *string                                `json:"expireTime"`
	Location                       *string                                `json:"location"`
	EnableTPU                      *bool                                  `json:"enableTPU"`
	TPUIPv4CidrBlock               *string                                `json:"tpuIPv4CidrBlock"`
	Conditions                     []ClusterConditions                    `json:"conditions"`
	Autopilot                      *ClusterAutopilot                      `json:"autopilot"`
	Project                        *string                                `json:"project"`
	NodeConfig                     *ClusterNodeConfig                     `json:"nodeConfig"`
	ReleaseChannel                 *ClusterReleaseChannel                 `json:"releaseChannel"`
	WorkloadIdentityConfig         *ClusterWorkloadIdentityConfig         `json:"workloadIdentityConfig"`
	NotificationConfig             *ClusterNotificationConfig             `json:"notificationConfig"`
	ConfidentialNodes              *ClusterConfidentialNodes              `json:"confidentialNodes"`
	SelfLink                       *string                                `json:"selfLink"`
	Zone                           *string                                `json:"zone"`
	InitialClusterVersion          *string                                `json:"initialClusterVersion"`
	CurrentMasterVersion           *string                                `json:"currentMasterVersion"`
	CurrentNodeVersion             *string                                `json:"currentNodeVersion"`
	InstanceGroupUrls              []string                               `json:"instanceGroupUrls"`
	CurrentNodeCount               *int64                                 `json:"currentNodeCount"`
	Id                             *string                                `json:"id"`
	PodSecurityPolicyConfig        *ClusterPodSecurityPolicyConfig        `json:"podSecurityPolicyConfig"`
	PrivateCluster                 *bool                                  `json:"privateCluster"`
	MasterIPv4CidrBlock            *string                                `json:"masterIPv4CidrBlock"`
	ClusterTelemetry               *ClusterClusterTelemetry               `json:"clusterTelemetry"`
	TPUConfig                      *ClusterTPUConfig                      `json:"tpuConfig"`
	Master                         *ClusterMaster                         `json:"master"`
}

func (r *Cluster) String() string {
	return dcl.SprintResource(r)
}

// The enum ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum.
type ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum string

// ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumRef returns a *ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumRef(s string) *ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(s)
	return &v
}

func (v ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) Validate() error {
	for _, s := range []string{"LOAD_BALANCER_TYPE_UNSPECIFIED", "LOAD_BALANCER_TYPE_EXTERNAL", "LOAD_BALANCER_TYPE_INTERNAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterAddonsConfigIstioConfigAuthEnum.
type ClusterAddonsConfigIstioConfigAuthEnum string

// ClusterAddonsConfigIstioConfigAuthEnumRef returns a *ClusterAddonsConfigIstioConfigAuthEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterAddonsConfigIstioConfigAuthEnumRef(s string) *ClusterAddonsConfigIstioConfigAuthEnum {
	if s == "" {
		return nil
	}

	v := ClusterAddonsConfigIstioConfigAuthEnum(s)
	return &v
}

func (v ClusterAddonsConfigIstioConfigAuthEnum) Validate() error {
	for _, s := range []string{"AUTH_NONE", "AUTH_MUTUAL_TLS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterAddonsConfigIstioConfigAuthEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum.
type ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum string

// ClusterNodePoolsConfigWorkloadMetadataConfigModeEnumRef returns a *ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConfigWorkloadMetadataConfigModeEnumRef(s string) *ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(s)
	return &v
}

func (v ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) Validate() error {
	for _, s := range []string{"MODE_UNSPECIFIED", "GCE_METADATA", "GKE_METADATA"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum.
type ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum string

// ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnumRef returns a *ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnumRef(s string) *ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(s)
	return &v
}

func (v ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "SECURE", "EXPOSE", "GKE_METADATA_SERVER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConfigTaintsEffectEnum.
type ClusterNodePoolsConfigTaintsEffectEnum string

// ClusterNodePoolsConfigTaintsEffectEnumRef returns a *ClusterNodePoolsConfigTaintsEffectEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConfigTaintsEffectEnumRef(s string) *ClusterNodePoolsConfigTaintsEffectEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConfigTaintsEffectEnum(s)
	return &v
}

func (v ClusterNodePoolsConfigTaintsEffectEnum) Validate() error {
	for _, s := range []string{"EFFECT_UNSPECIFIED", "NO_SCHEDULE", "PREFER_NO_SCHEDULE", "NO_EXECUTE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConfigTaintsEffectEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConfigSandboxConfigTypeEnum.
type ClusterNodePoolsConfigSandboxConfigTypeEnum string

// ClusterNodePoolsConfigSandboxConfigTypeEnumRef returns a *ClusterNodePoolsConfigSandboxConfigTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConfigSandboxConfigTypeEnumRef(s string) *ClusterNodePoolsConfigSandboxConfigTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConfigSandboxConfigTypeEnum(s)
	return &v
}

func (v ClusterNodePoolsConfigSandboxConfigTypeEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "GVISOR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConfigSandboxConfigTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum.
type ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum string

// ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumRef returns a *ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumRef(s string) *ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(s)
	return &v
}

func (v ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsStatusEnum.
type ClusterNodePoolsStatusEnum string

// ClusterNodePoolsStatusEnumRef returns a *ClusterNodePoolsStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsStatusEnumRef(s string) *ClusterNodePoolsStatusEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsStatusEnum(s)
	return &v
}

func (v ClusterNodePoolsStatusEnum) Validate() error {
	for _, s := range []string{"STATUS_UNSPECIFIED", "PROVISIONING", "RUNNING", "RECONCILING", "STOPPING", "ERROR", "DEGRADED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConditionsCodeEnum.
type ClusterNodePoolsConditionsCodeEnum string

// ClusterNodePoolsConditionsCodeEnumRef returns a *ClusterNodePoolsConditionsCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConditionsCodeEnumRef(s string) *ClusterNodePoolsConditionsCodeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConditionsCodeEnum(s)
	return &v
}

func (v ClusterNodePoolsConditionsCodeEnum) Validate() error {
	for _, s := range []string{"OK", "CANCELLED", "UNKNOWN", "INVALID_ARGUMENT", "DEADLINE_EXCEEDED", "NOT_FOUND", "ALREADY_EXISTS", "PERMISSION_DENIED", "UNAUTHENTICATED", "RESOURCE_EXHAUSTED", "FAILED_PRECONDITION", "ABORTED", "OUT_OF_RANGE", "UNIMPLEMENTED", "INTERNAL", "UNAVAILABLE", "DATA_LOSS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConditionsCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodePoolsConditionsCanonicalCodeEnum.
type ClusterNodePoolsConditionsCanonicalCodeEnum string

// ClusterNodePoolsConditionsCanonicalCodeEnumRef returns a *ClusterNodePoolsConditionsCanonicalCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodePoolsConditionsCanonicalCodeEnumRef(s string) *ClusterNodePoolsConditionsCanonicalCodeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodePoolsConditionsCanonicalCodeEnum(s)
	return &v
}

func (v ClusterNodePoolsConditionsCanonicalCodeEnum) Validate() error {
	for _, s := range []string{"OK", "CANCELLED", "UNKNOWN", "INVALID_ARGUMENT", "DEADLINE_EXCEEDED", "NOT_FOUND", "ALREADY_EXISTS", "PERMISSION_DENIED", "UNAUTHENTICATED", "RESOURCE_EXHAUSTED", "FAILED_PRECONDITION", "ABORTED", "OUT_OF_RANGE", "UNIMPLEMENTED", "INTERNAL", "UNAVAILABLE", "DATA_LOSS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodePoolsConditionsCanonicalCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNetworkPolicyProviderEnum.
type ClusterNetworkPolicyProviderEnum string

// ClusterNetworkPolicyProviderEnumRef returns a *ClusterNetworkPolicyProviderEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNetworkPolicyProviderEnumRef(s string) *ClusterNetworkPolicyProviderEnum {
	if s == "" {
		return nil
	}

	v := ClusterNetworkPolicyProviderEnum(s)
	return &v
}

func (v ClusterNetworkPolicyProviderEnum) Validate() error {
	for _, s := range []string{"PROVIDER_UNSPECIFIED", "CALICO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNetworkPolicyProviderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterAutoscalingAutoscalingProfileEnum.
type ClusterAutoscalingAutoscalingProfileEnum string

// ClusterAutoscalingAutoscalingProfileEnumRef returns a *ClusterAutoscalingAutoscalingProfileEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterAutoscalingAutoscalingProfileEnumRef(s string) *ClusterAutoscalingAutoscalingProfileEnum {
	if s == "" {
		return nil
	}

	v := ClusterAutoscalingAutoscalingProfileEnum(s)
	return &v
}

func (v ClusterAutoscalingAutoscalingProfileEnum) Validate() error {
	for _, s := range []string{"PROFILE_UNSPECIFIED", "OPTIMIZE_UTILIZATION", "BALANCED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterAutoscalingAutoscalingProfileEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNetworkConfigPrivateIPv6GoogleAccessEnum.
type ClusterNetworkConfigPrivateIPv6GoogleAccessEnum string

// ClusterNetworkConfigPrivateIPv6GoogleAccessEnumRef returns a *ClusterNetworkConfigPrivateIPv6GoogleAccessEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNetworkConfigPrivateIPv6GoogleAccessEnumRef(s string) *ClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	if s == "" {
		return nil
	}

	v := ClusterNetworkConfigPrivateIPv6GoogleAccessEnum(s)
	return &v
}

func (v ClusterNetworkConfigPrivateIPv6GoogleAccessEnum) Validate() error {
	for _, s := range []string{"PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED", "PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED", "PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE", "PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNetworkConfigPrivateIPv6GoogleAccessEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNetworkConfigDatapathProviderEnum.
type ClusterNetworkConfigDatapathProviderEnum string

// ClusterNetworkConfigDatapathProviderEnumRef returns a *ClusterNetworkConfigDatapathProviderEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNetworkConfigDatapathProviderEnumRef(s string) *ClusterNetworkConfigDatapathProviderEnum {
	if s == "" {
		return nil
	}

	v := ClusterNetworkConfigDatapathProviderEnum(s)
	return &v
}

func (v ClusterNetworkConfigDatapathProviderEnum) Validate() error {
	for _, s := range []string{"DATAPATH_PROVIDER_UNSPECIFIED", "LEGACY_DATAPATH", "ADVANCED_DATAPATH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNetworkConfigDatapathProviderEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterDatabaseEncryptionStateEnum.
type ClusterDatabaseEncryptionStateEnum string

// ClusterDatabaseEncryptionStateEnumRef returns a *ClusterDatabaseEncryptionStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterDatabaseEncryptionStateEnumRef(s string) *ClusterDatabaseEncryptionStateEnum {
	if s == "" {
		return nil
	}

	v := ClusterDatabaseEncryptionStateEnum(s)
	return &v
}

func (v ClusterDatabaseEncryptionStateEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "ENCRYPTED", "DECRYPTED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterDatabaseEncryptionStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterConditionsCanonicalCodeEnum.
type ClusterConditionsCanonicalCodeEnum string

// ClusterConditionsCanonicalCodeEnumRef returns a *ClusterConditionsCanonicalCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterConditionsCanonicalCodeEnumRef(s string) *ClusterConditionsCanonicalCodeEnum {
	if s == "" {
		return nil
	}

	v := ClusterConditionsCanonicalCodeEnum(s)
	return &v
}

func (v ClusterConditionsCanonicalCodeEnum) Validate() error {
	for _, s := range []string{"OK", "CANCELLED", "UNKNOWN", "INVALID_ARGUMENT", "DEADLINE_EXCEEDED", "NOT_FOUND", "ALREADY_EXISTS", "PERMISSION_DENIED", "UNAUTHENTICATED", "RESOURCE_EXHAUSTED", "FAILED_PRECONDITION", "ABORTED", "OUT_OF_RANGE", "UNIMPLEMENTED", "INTERNAL", "UNAVAILABLE", "DATA_LOSS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterConditionsCanonicalCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodeConfigWorkloadMetadataConfigModeEnum.
type ClusterNodeConfigWorkloadMetadataConfigModeEnum string

// ClusterNodeConfigWorkloadMetadataConfigModeEnumRef returns a *ClusterNodeConfigWorkloadMetadataConfigModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodeConfigWorkloadMetadataConfigModeEnumRef(s string) *ClusterNodeConfigWorkloadMetadataConfigModeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodeConfigWorkloadMetadataConfigModeEnum(s)
	return &v
}

func (v ClusterNodeConfigWorkloadMetadataConfigModeEnum) Validate() error {
	for _, s := range []string{"MODE_UNSPECIFIED", "GCE_METADATA", "GKE_METADATA"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodeConfigWorkloadMetadataConfigModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum.
type ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum string

// ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnumRef returns a *ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnumRef(s string) *ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(s)
	return &v
}

func (v ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "SECURE", "EXPOSE", "GKE_METADATA_SERVER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodeConfigTaintsEffectEnum.
type ClusterNodeConfigTaintsEffectEnum string

// ClusterNodeConfigTaintsEffectEnumRef returns a *ClusterNodeConfigTaintsEffectEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodeConfigTaintsEffectEnumRef(s string) *ClusterNodeConfigTaintsEffectEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodeConfigTaintsEffectEnum(s)
	return &v
}

func (v ClusterNodeConfigTaintsEffectEnum) Validate() error {
	for _, s := range []string{"EFFECT_UNSPECIFIED", "NO_SCHEDULE", "PREFER_NO_SCHEDULE", "NO_EXECUTE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodeConfigTaintsEffectEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodeConfigSandboxConfigTypeEnum.
type ClusterNodeConfigSandboxConfigTypeEnum string

// ClusterNodeConfigSandboxConfigTypeEnumRef returns a *ClusterNodeConfigSandboxConfigTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodeConfigSandboxConfigTypeEnumRef(s string) *ClusterNodeConfigSandboxConfigTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodeConfigSandboxConfigTypeEnum(s)
	return &v
}

func (v ClusterNodeConfigSandboxConfigTypeEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "GVISOR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodeConfigSandboxConfigTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum.
type ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum string

// ClusterNodeConfigReservationAffinityConsumeReservationTypeEnumRef returns a *ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterNodeConfigReservationAffinityConsumeReservationTypeEnumRef(s string) *ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(s)
	return &v
}

func (v ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterReleaseChannelChannelEnum.
type ClusterReleaseChannelChannelEnum string

// ClusterReleaseChannelChannelEnumRef returns a *ClusterReleaseChannelChannelEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterReleaseChannelChannelEnumRef(s string) *ClusterReleaseChannelChannelEnum {
	if s == "" {
		return nil
	}

	v := ClusterReleaseChannelChannelEnum(s)
	return &v
}

func (v ClusterReleaseChannelChannelEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "RAPID", "REGULAR", "STABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterReleaseChannelChannelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ClusterClusterTelemetryTypeEnum.
type ClusterClusterTelemetryTypeEnum string

// ClusterClusterTelemetryTypeEnumRef returns a *ClusterClusterTelemetryTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ClusterClusterTelemetryTypeEnumRef(s string) *ClusterClusterTelemetryTypeEnum {
	if s == "" {
		return nil
	}

	v := ClusterClusterTelemetryTypeEnum(s)
	return &v
}

func (v ClusterClusterTelemetryTypeEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "GVISOR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ClusterClusterTelemetryTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ClusterMasterAuth struct {
	empty                   bool                                      `json:"-"`
	Username                *string                                   `json:"username"`
	Password                *string                                   `json:"password"`
	ClientCertificateConfig *ClusterMasterAuthClientCertificateConfig `json:"clientCertificateConfig"`
	ClusterCaCertificate    *string                                   `json:"clusterCaCertificate"`
	ClientCertificate       *string                                   `json:"clientCertificate"`
	ClientKey               *string                                   `json:"clientKey"`
}

// This object is used to assert a desired state where this ClusterMasterAuth is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMasterAuth *ClusterMasterAuth = &ClusterMasterAuth{empty: true}

func (r *ClusterMasterAuth) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMasterAuth) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMasterAuthClientCertificateConfig struct {
	empty                  bool  `json:"-"`
	IssueClientCertificate *bool `json:"issueClientCertificate"`
}

// This object is used to assert a desired state where this ClusterMasterAuthClientCertificateConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMasterAuthClientCertificateConfig *ClusterMasterAuthClientCertificateConfig = &ClusterMasterAuthClientCertificateConfig{empty: true}

func (r *ClusterMasterAuthClientCertificateConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMasterAuthClientCertificateConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfig struct {
	empty                            bool                                                 `json:"-"`
	HttpLoadBalancing                *ClusterAddonsConfigHttpLoadBalancing                `json:"httpLoadBalancing"`
	HorizontalPodAutoscaling         *ClusterAddonsConfigHorizontalPodAutoscaling         `json:"horizontalPodAutoscaling"`
	KubernetesDashboard              *ClusterAddonsConfigKubernetesDashboard              `json:"kubernetesDashboard"`
	NetworkPolicyConfig              *ClusterAddonsConfigNetworkPolicyConfig              `json:"networkPolicyConfig"`
	CloudRunConfig                   *ClusterAddonsConfigCloudRunConfig                   `json:"cloudRunConfig"`
	DnsCacheConfig                   *ClusterAddonsConfigDnsCacheConfig                   `json:"dnsCacheConfig"`
	ConfigConnectorConfig            *ClusterAddonsConfigConfigConnectorConfig            `json:"configConnectorConfig"`
	GcePersistentDiskCsiDriverConfig *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfig"`
	IstioConfig                      *ClusterAddonsConfigIstioConfig                      `json:"istioConfig"`
	KalmConfig                       *ClusterAddonsConfigKalmConfig                       `json:"kalmConfig"`
}

// This object is used to assert a desired state where this ClusterAddonsConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfig *ClusterAddonsConfig = &ClusterAddonsConfig{empty: true}

func (r *ClusterAddonsConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigHttpLoadBalancing struct {
	empty    bool  `json:"-"`
	Disabled *bool `json:"disabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigHttpLoadBalancing is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigHttpLoadBalancing *ClusterAddonsConfigHttpLoadBalancing = &ClusterAddonsConfigHttpLoadBalancing{empty: true}

func (r *ClusterAddonsConfigHttpLoadBalancing) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigHttpLoadBalancing) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigHorizontalPodAutoscaling struct {
	empty    bool  `json:"-"`
	Disabled *bool `json:"disabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigHorizontalPodAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigHorizontalPodAutoscaling *ClusterAddonsConfigHorizontalPodAutoscaling = &ClusterAddonsConfigHorizontalPodAutoscaling{empty: true}

func (r *ClusterAddonsConfigHorizontalPodAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigHorizontalPodAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigKubernetesDashboard struct {
	empty    bool  `json:"-"`
	Disabled *bool `json:"disabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigKubernetesDashboard is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigKubernetesDashboard *ClusterAddonsConfigKubernetesDashboard = &ClusterAddonsConfigKubernetesDashboard{empty: true}

func (r *ClusterAddonsConfigKubernetesDashboard) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigKubernetesDashboard) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigNetworkPolicyConfig struct {
	empty    bool  `json:"-"`
	Disabled *bool `json:"disabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigNetworkPolicyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigNetworkPolicyConfig *ClusterAddonsConfigNetworkPolicyConfig = &ClusterAddonsConfigNetworkPolicyConfig{empty: true}

func (r *ClusterAddonsConfigNetworkPolicyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigNetworkPolicyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigCloudRunConfig struct {
	empty            bool                                                   `json:"-"`
	Disabled         *bool                                                  `json:"disabled"`
	LoadBalancerType *ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum `json:"loadBalancerType"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigCloudRunConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigCloudRunConfig *ClusterAddonsConfigCloudRunConfig = &ClusterAddonsConfigCloudRunConfig{empty: true}

func (r *ClusterAddonsConfigCloudRunConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigCloudRunConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigDnsCacheConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigDnsCacheConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigDnsCacheConfig *ClusterAddonsConfigDnsCacheConfig = &ClusterAddonsConfigDnsCacheConfig{empty: true}

func (r *ClusterAddonsConfigDnsCacheConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigDnsCacheConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigConfigConnectorConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigConfigConnectorConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigConfigConnectorConfig *ClusterAddonsConfigConfigConnectorConfig = &ClusterAddonsConfigConfigConnectorConfig{empty: true}

func (r *ClusterAddonsConfigConfigConnectorConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigConfigConnectorConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigGcePersistentDiskCsiDriverConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigGcePersistentDiskCsiDriverConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigGcePersistentDiskCsiDriverConfig *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig = &ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{empty: true}

func (r *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigIstioConfig struct {
	empty    bool                                    `json:"-"`
	Disabled *bool                                   `json:"disabled"`
	Auth     *ClusterAddonsConfigIstioConfigAuthEnum `json:"auth"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigIstioConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigIstioConfig *ClusterAddonsConfigIstioConfig = &ClusterAddonsConfigIstioConfig{empty: true}

func (r *ClusterAddonsConfigIstioConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigIstioConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAddonsConfigKalmConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterAddonsConfigKalmConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAddonsConfigKalmConfig *ClusterAddonsConfigKalmConfig = &ClusterAddonsConfigKalmConfig{empty: true}

func (r *ClusterAddonsConfigKalmConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAddonsConfigKalmConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePools struct {
	empty             bool                               `json:"-"`
	Name              *string                            `json:"name"`
	Config            *ClusterNodePoolsConfig            `json:"config"`
	InitialNodeCount  *int64                             `json:"initialNodeCount"`
	Locations         []string                           `json:"locations"`
	SelfLink          *string                            `json:"selfLink"`
	Version           *string                            `json:"version"`
	InstanceGroupUrls []string                           `json:"instanceGroupUrls"`
	Status            *ClusterNodePoolsStatusEnum        `json:"status"`
	StatusMessage     *string                            `json:"statusMessage"`
	Autoscaling       *ClusterNodePoolsAutoscaling       `json:"autoscaling"`
	Management        *ClusterNodePoolsManagement        `json:"management"`
	MaxPodsConstraint *ClusterNodePoolsMaxPodsConstraint `json:"maxPodsConstraint"`
	Conditions        []ClusterNodePoolsConditions       `json:"conditions"`
	PodIPv4CidrSize   *int64                             `json:"podIPv4CidrSize"`
	UpgradeSettings   *ClusterNodePoolsUpgradeSettings   `json:"upgradeSettings"`
	NetworkConfig     *ClusterNodePoolsNetworkConfig     `json:"networkConfig"`
}

// This object is used to assert a desired state where this ClusterNodePools is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePools *ClusterNodePools = &ClusterNodePools{empty: true}

func (r *ClusterNodePools) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePools) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfig struct {
	empty                  bool                                          `json:"-"`
	MachineType            *string                                       `json:"machineType"`
	DiskSizeGb             *int64                                        `json:"diskSizeGb"`
	OAuthScopes            []string                                      `json:"oauthScopes"`
	ServiceAccount         *string                                       `json:"serviceAccount"`
	Metadata               map[string]string                             `json:"metadata"`
	ImageType              *string                                       `json:"imageType"`
	Labels                 map[string]string                             `json:"labels"`
	LocalSsdCount          *int64                                        `json:"localSsdCount"`
	Tags                   []string                                      `json:"tags"`
	Preemptible            *bool                                         `json:"preemptible"`
	Accelerators           []ClusterNodePoolsConfigAccelerators          `json:"accelerators"`
	DiskType               *string                                       `json:"diskType"`
	MinCpuPlatform         *string                                       `json:"minCpuPlatform"`
	WorkloadMetadataConfig *ClusterNodePoolsConfigWorkloadMetadataConfig `json:"workloadMetadataConfig"`
	Taints                 []ClusterNodePoolsConfigTaints                `json:"taints"`
	SandboxConfig          *ClusterNodePoolsConfigSandboxConfig          `json:"sandboxConfig"`
	NodeGroup              *string                                       `json:"nodeGroup"`
	ReservationAffinity    *ClusterNodePoolsConfigReservationAffinity    `json:"reservationAffinity"`
	ShieldedInstanceConfig *ClusterNodePoolsConfigShieldedInstanceConfig `json:"shieldedInstanceConfig"`
	LinuxNodeConfig        *ClusterNodePoolsConfigLinuxNodeConfig        `json:"linuxNodeConfig"`
	KubeletConfig          *ClusterNodePoolsConfigKubeletConfig          `json:"kubeletConfig"`
	BootDiskKmsKey         *string                                       `json:"bootDiskKmsKey"`
	EphemeralStorageConfig *ClusterNodePoolsConfigEphemeralStorageConfig `json:"ephemeralStorageConfig"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfig *ClusterNodePoolsConfig = &ClusterNodePoolsConfig{empty: true}

func (r *ClusterNodePoolsConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
	AcceleratorType  *string `json:"acceleratorType"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigAccelerators *ClusterNodePoolsConfigAccelerators = &ClusterNodePoolsConfigAccelerators{empty: true}

func (r *ClusterNodePoolsConfigAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigWorkloadMetadataConfig struct {
	empty        bool                                                          `json:"-"`
	Mode         *ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum         `json:"mode"`
	NodeMetadata *ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum `json:"nodeMetadata"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigWorkloadMetadataConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigWorkloadMetadataConfig *ClusterNodePoolsConfigWorkloadMetadataConfig = &ClusterNodePoolsConfigWorkloadMetadataConfig{empty: true}

func (r *ClusterNodePoolsConfigWorkloadMetadataConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigWorkloadMetadataConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigTaints struct {
	empty  bool                                    `json:"-"`
	Key    *string                                 `json:"key"`
	Value  *string                                 `json:"value"`
	Effect *ClusterNodePoolsConfigTaintsEffectEnum `json:"effect"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigTaints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigTaints *ClusterNodePoolsConfigTaints = &ClusterNodePoolsConfigTaints{empty: true}

func (r *ClusterNodePoolsConfigTaints) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigTaints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigSandboxConfig struct {
	empty       bool                                         `json:"-"`
	Type        *ClusterNodePoolsConfigSandboxConfigTypeEnum `json:"type"`
	SandboxType *string                                      `json:"sandboxType"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigSandboxConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigSandboxConfig *ClusterNodePoolsConfigSandboxConfig = &ClusterNodePoolsConfigSandboxConfig{empty: true}

func (r *ClusterNodePoolsConfigSandboxConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigSandboxConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigReservationAffinity struct {
	empty                  bool                                                                 `json:"-"`
	ConsumeReservationType *ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum `json:"consumeReservationType"`
	Key                    *string                                                              `json:"key"`
	Values                 []string                                                             `json:"values"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigReservationAffinity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigReservationAffinity *ClusterNodePoolsConfigReservationAffinity = &ClusterNodePoolsConfigReservationAffinity{empty: true}

func (r *ClusterNodePoolsConfigReservationAffinity) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigReservationAffinity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigShieldedInstanceConfig struct {
	empty                     bool  `json:"-"`
	EnableSecureBoot          *bool `json:"enableSecureBoot"`
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigShieldedInstanceConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigShieldedInstanceConfig *ClusterNodePoolsConfigShieldedInstanceConfig = &ClusterNodePoolsConfigShieldedInstanceConfig{empty: true}

func (r *ClusterNodePoolsConfigShieldedInstanceConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigShieldedInstanceConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigLinuxNodeConfig struct {
	empty   bool              `json:"-"`
	Sysctls map[string]string `json:"sysctls"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigLinuxNodeConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigLinuxNodeConfig *ClusterNodePoolsConfigLinuxNodeConfig = &ClusterNodePoolsConfigLinuxNodeConfig{empty: true}

func (r *ClusterNodePoolsConfigLinuxNodeConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigLinuxNodeConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigKubeletConfig struct {
	empty             bool    `json:"-"`
	CpuManagerPolicy  *string `json:"cpuManagerPolicy"`
	CpuCfsQuota       *bool   `json:"cpuCfsQuota"`
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigKubeletConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigKubeletConfig *ClusterNodePoolsConfigKubeletConfig = &ClusterNodePoolsConfigKubeletConfig{empty: true}

func (r *ClusterNodePoolsConfigKubeletConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigKubeletConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConfigEphemeralStorageConfig struct {
	empty         bool   `json:"-"`
	LocalSsdCount *int64 `json:"localSsdCount"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConfigEphemeralStorageConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConfigEphemeralStorageConfig *ClusterNodePoolsConfigEphemeralStorageConfig = &ClusterNodePoolsConfigEphemeralStorageConfig{empty: true}

func (r *ClusterNodePoolsConfigEphemeralStorageConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConfigEphemeralStorageConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsAutoscaling struct {
	empty           bool   `json:"-"`
	Enabled         *bool  `json:"enabled"`
	MinNodeCount    *int64 `json:"minNodeCount"`
	MaxNodeCount    *int64 `json:"maxNodeCount"`
	Autoprovisioned *bool  `json:"autoprovisioned"`
}

// This object is used to assert a desired state where this ClusterNodePoolsAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsAutoscaling *ClusterNodePoolsAutoscaling = &ClusterNodePoolsAutoscaling{empty: true}

func (r *ClusterNodePoolsAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsManagement struct {
	empty          bool                                      `json:"-"`
	AutoUpgrade    *bool                                     `json:"autoUpgrade"`
	AutoRepair     *bool                                     `json:"autoRepair"`
	UpgradeOptions *ClusterNodePoolsManagementUpgradeOptions `json:"upgradeOptions"`
}

// This object is used to assert a desired state where this ClusterNodePoolsManagement is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsManagement *ClusterNodePoolsManagement = &ClusterNodePoolsManagement{empty: true}

func (r *ClusterNodePoolsManagement) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsManagement) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsManagementUpgradeOptions struct {
	empty                bool    `json:"-"`
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime"`
	Description          *string `json:"description"`
}

// This object is used to assert a desired state where this ClusterNodePoolsManagementUpgradeOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsManagementUpgradeOptions *ClusterNodePoolsManagementUpgradeOptions = &ClusterNodePoolsManagementUpgradeOptions{empty: true}

func (r *ClusterNodePoolsManagementUpgradeOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsManagementUpgradeOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsMaxPodsConstraint struct {
	empty          bool   `json:"-"`
	MaxPodsPerNode *int64 `json:"maxPodsPerNode"`
}

// This object is used to assert a desired state where this ClusterNodePoolsMaxPodsConstraint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsMaxPodsConstraint *ClusterNodePoolsMaxPodsConstraint = &ClusterNodePoolsMaxPodsConstraint{empty: true}

func (r *ClusterNodePoolsMaxPodsConstraint) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsMaxPodsConstraint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsConditions struct {
	empty         bool                                         `json:"-"`
	Code          *ClusterNodePoolsConditionsCodeEnum          `json:"code"`
	Message       *string                                      `json:"message"`
	CanonicalCode *ClusterNodePoolsConditionsCanonicalCodeEnum `json:"canonicalCode"`
}

// This object is used to assert a desired state where this ClusterNodePoolsConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsConditions *ClusterNodePoolsConditions = &ClusterNodePoolsConditions{empty: true}

func (r *ClusterNodePoolsConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsUpgradeSettings struct {
	empty          bool   `json:"-"`
	MaxSurge       *int64 `json:"maxSurge"`
	MaxUnavailable *int64 `json:"maxUnavailable"`
}

// This object is used to assert a desired state where this ClusterNodePoolsUpgradeSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsUpgradeSettings *ClusterNodePoolsUpgradeSettings = &ClusterNodePoolsUpgradeSettings{empty: true}

func (r *ClusterNodePoolsUpgradeSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsUpgradeSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodePoolsNetworkConfig struct {
	empty            bool    `json:"-"`
	CreatePodRange   *bool   `json:"createPodRange"`
	PodRange         *string `json:"podRange"`
	PodIPv4CidrBlock *string `json:"podIPv4CidrBlock"`
}

// This object is used to assert a desired state where this ClusterNodePoolsNetworkConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodePoolsNetworkConfig *ClusterNodePoolsNetworkConfig = &ClusterNodePoolsNetworkConfig{empty: true}

func (r *ClusterNodePoolsNetworkConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodePoolsNetworkConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterLegacyAbac struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterLegacyAbac is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterLegacyAbac *ClusterLegacyAbac = &ClusterLegacyAbac{empty: true}

func (r *ClusterLegacyAbac) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterLegacyAbac) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNetworkPolicy struct {
	empty    bool                              `json:"-"`
	Provider *ClusterNetworkPolicyProviderEnum `json:"provider"`
	Enabled  *bool                             `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterNetworkPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNetworkPolicy *ClusterNetworkPolicy = &ClusterNetworkPolicy{empty: true}

func (r *ClusterNetworkPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNetworkPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterIPAllocationPolicy struct {
	empty                      bool    `json:"-"`
	UseIPAliases               *bool   `json:"useIPAliases"`
	CreateSubnetwork           *bool   `json:"createSubnetwork"`
	SubnetworkName             *string `json:"subnetworkName"`
	ClusterSecondaryRangeName  *string `json:"clusterSecondaryRangeName"`
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName"`
	ClusterIPv4CidrBlock       *string `json:"clusterIPv4CidrBlock"`
	NodeIPv4CidrBlock          *string `json:"nodeIPv4CidrBlock"`
	ServicesIPv4CidrBlock      *string `json:"servicesIPv4CidrBlock"`
	TPUIPv4CidrBlock           *string `json:"tpuIPv4CidrBlock"`
	ClusterIPv4Cidr            *string `json:"clusterIPv4Cidr"`
	NodeIPv4Cidr               *string `json:"nodeIPv4Cidr"`
	ServicesIPv4Cidr           *string `json:"servicesIPv4Cidr"`
	UseRoutes                  *bool   `json:"useRoutes"`
	AllowRouteOverlap          *bool   `json:"allowRouteOverlap"`
}

// This object is used to assert a desired state where this ClusterIPAllocationPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterIPAllocationPolicy *ClusterIPAllocationPolicy = &ClusterIPAllocationPolicy{empty: true}

func (r *ClusterIPAllocationPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterIPAllocationPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMasterAuthorizedNetworksConfig struct {
	empty      bool                                              `json:"-"`
	Enabled    *bool                                             `json:"enabled"`
	CidrBlocks []ClusterMasterAuthorizedNetworksConfigCidrBlocks `json:"cidrBlocks"`
}

// This object is used to assert a desired state where this ClusterMasterAuthorizedNetworksConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMasterAuthorizedNetworksConfig *ClusterMasterAuthorizedNetworksConfig = &ClusterMasterAuthorizedNetworksConfig{empty: true}

func (r *ClusterMasterAuthorizedNetworksConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMasterAuthorizedNetworksConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMasterAuthorizedNetworksConfigCidrBlocks struct {
	empty       bool    `json:"-"`
	DisplayName *string `json:"displayName"`
	CidrBlock   *string `json:"cidrBlock"`
}

// This object is used to assert a desired state where this ClusterMasterAuthorizedNetworksConfigCidrBlocks is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMasterAuthorizedNetworksConfigCidrBlocks *ClusterMasterAuthorizedNetworksConfigCidrBlocks = &ClusterMasterAuthorizedNetworksConfigCidrBlocks{empty: true}

func (r *ClusterMasterAuthorizedNetworksConfigCidrBlocks) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMasterAuthorizedNetworksConfigCidrBlocks) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterBinaryAuthorization struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterBinaryAuthorization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterBinaryAuthorization *ClusterBinaryAuthorization = &ClusterBinaryAuthorization{empty: true}

func (r *ClusterBinaryAuthorization) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterBinaryAuthorization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscaling struct {
	empty                            bool                                                `json:"-"`
	EnableNodeAutoprovisioning       *bool                                               `json:"enableNodeAutoprovisioning"`
	ResourceLimits                   []ClusterAutoscalingResourceLimits                  `json:"resourceLimits"`
	AutoprovisioningNodePoolDefaults *ClusterAutoscalingAutoprovisioningNodePoolDefaults `json:"autoprovisioningNodePoolDefaults"`
	AutoprovisioningLocations        []string                                            `json:"autoprovisioningLocations"`
	AutoscalingProfile               *ClusterAutoscalingAutoscalingProfileEnum           `json:"autoscalingProfile"`
}

// This object is used to assert a desired state where this ClusterAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscaling *ClusterAutoscaling = &ClusterAutoscaling{empty: true}

func (r *ClusterAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscalingResourceLimits struct {
	empty        bool    `json:"-"`
	ResourceType *string `json:"resourceType"`
	Minimum      *int64  `json:"minimum"`
	Maximum      *int64  `json:"maximum"`
}

// This object is used to assert a desired state where this ClusterAutoscalingResourceLimits is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscalingResourceLimits *ClusterAutoscalingResourceLimits = &ClusterAutoscalingResourceLimits{empty: true}

func (r *ClusterAutoscalingResourceLimits) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscalingResourceLimits) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscalingAutoprovisioningNodePoolDefaults struct {
	empty                  bool                                                                      `json:"-"`
	OAuthScopes            []string                                                                  `json:"oauthScopes"`
	ServiceAccount         *string                                                                   `json:"serviceAccount"`
	UpgradeSettings        *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings        `json:"upgradeSettings"`
	Management             *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement             `json:"management"`
	MinCpuPlatform         *string                                                                   `json:"minCpuPlatform"`
	DiskSizeGb             *int64                                                                    `json:"diskSizeGb"`
	DiskType               *string                                                                   `json:"diskType"`
	ShieldedInstanceConfig *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig `json:"shieldedInstanceConfig"`
	BootDiskKmsKey         *string                                                                   `json:"bootDiskKmsKey"`
}

// This object is used to assert a desired state where this ClusterAutoscalingAutoprovisioningNodePoolDefaults is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscalingAutoprovisioningNodePoolDefaults *ClusterAutoscalingAutoprovisioningNodePoolDefaults = &ClusterAutoscalingAutoprovisioningNodePoolDefaults{empty: true}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaults) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaults) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings struct {
	empty          bool   `json:"-"`
	MaxSurge       *int64 `json:"maxSurge"`
	MaxUnavailable *int64 `json:"maxUnavailable"`
}

// This object is used to assert a desired state where this ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings = &ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{empty: true}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement struct {
	empty          bool                                                                        `json:"-"`
	AutoUpgrade    *bool                                                                       `json:"autoUpgrade"`
	AutoRepair     *bool                                                                       `json:"autoRepair"`
	UpgradeOptions *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions `json:"upgradeOptions"`
}

// This object is used to assert a desired state where this ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement = &ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{empty: true}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions struct {
	empty                bool    `json:"-"`
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime"`
	Description          *string `json:"description"`
}

// This object is used to assert a desired state where this ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions = &ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{empty: true}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig struct {
	empty                     bool  `json:"-"`
	EnableSecureBoot          *bool `json:"enableSecureBoot"`
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring"`
}

// This object is used to assert a desired state where this ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig = &ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{empty: true}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNetworkConfig struct {
	empty                     bool                                             `json:"-"`
	Network                   *string                                          `json:"network"`
	Subnetwork                *string                                          `json:"subnetwork"`
	EnableIntraNodeVisibility *bool                                            `json:"enableIntraNodeVisibility"`
	DefaultSnatStatus         *ClusterNetworkConfigDefaultSnatStatus           `json:"defaultSnatStatus"`
	PrivateIPv6GoogleAccess   *ClusterNetworkConfigPrivateIPv6GoogleAccessEnum `json:"privateIPv6GoogleAccess"`
	DatapathProvider          *ClusterNetworkConfigDatapathProviderEnum        `json:"datapathProvider"`
}

// This object is used to assert a desired state where this ClusterNetworkConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNetworkConfig *ClusterNetworkConfig = &ClusterNetworkConfig{empty: true}

func (r *ClusterNetworkConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNetworkConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNetworkConfigDefaultSnatStatus struct {
	empty    bool  `json:"-"`
	Disabled *bool `json:"disabled"`
}

// This object is used to assert a desired state where this ClusterNetworkConfigDefaultSnatStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNetworkConfigDefaultSnatStatus *ClusterNetworkConfigDefaultSnatStatus = &ClusterNetworkConfigDefaultSnatStatus{empty: true}

func (r *ClusterNetworkConfigDefaultSnatStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNetworkConfigDefaultSnatStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMaintenancePolicy struct {
	empty           bool                            `json:"-"`
	Window          *ClusterMaintenancePolicyWindow `json:"window"`
	ResourceVersion *string                         `json:"resourceVersion"`
}

// This object is used to assert a desired state where this ClusterMaintenancePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMaintenancePolicy *ClusterMaintenancePolicy = &ClusterMaintenancePolicy{empty: true}

func (r *ClusterMaintenancePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMaintenancePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMaintenancePolicyWindow struct {
	empty                  bool                                                  `json:"-"`
	DailyMaintenanceWindow *ClusterMaintenancePolicyWindowDailyMaintenanceWindow `json:"dailyMaintenanceWindow"`
	RecurringWindow        *ClusterMaintenancePolicyWindowRecurringWindow        `json:"recurringWindow"`
	MaintenanceExclusions  map[string]string                                     `json:"maintenanceExclusions"`
}

// This object is used to assert a desired state where this ClusterMaintenancePolicyWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMaintenancePolicyWindow *ClusterMaintenancePolicyWindow = &ClusterMaintenancePolicyWindow{empty: true}

func (r *ClusterMaintenancePolicyWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMaintenancePolicyWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMaintenancePolicyWindowDailyMaintenanceWindow struct {
	empty     bool    `json:"-"`
	StartTime *string `json:"startTime"`
	Duration  *string `json:"duration"`
}

// This object is used to assert a desired state where this ClusterMaintenancePolicyWindowDailyMaintenanceWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMaintenancePolicyWindowDailyMaintenanceWindow *ClusterMaintenancePolicyWindowDailyMaintenanceWindow = &ClusterMaintenancePolicyWindowDailyMaintenanceWindow{empty: true}

func (r *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMaintenancePolicyWindowDailyMaintenanceWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMaintenancePolicyWindowRecurringWindow struct {
	empty      bool                                                 `json:"-"`
	Window     *ClusterMaintenancePolicyWindowRecurringWindowWindow `json:"window"`
	Recurrence *string                                              `json:"recurrence"`
}

// This object is used to assert a desired state where this ClusterMaintenancePolicyWindowRecurringWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMaintenancePolicyWindowRecurringWindow *ClusterMaintenancePolicyWindowRecurringWindow = &ClusterMaintenancePolicyWindowRecurringWindow{empty: true}

func (r *ClusterMaintenancePolicyWindowRecurringWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMaintenancePolicyWindowRecurringWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMaintenancePolicyWindowRecurringWindowWindow struct {
	empty     bool    `json:"-"`
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}

// This object is used to assert a desired state where this ClusterMaintenancePolicyWindowRecurringWindowWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMaintenancePolicyWindowRecurringWindowWindow *ClusterMaintenancePolicyWindowRecurringWindowWindow = &ClusterMaintenancePolicyWindowRecurringWindowWindow{empty: true}

func (r *ClusterMaintenancePolicyWindowRecurringWindowWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMaintenancePolicyWindowRecurringWindowWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterDefaultMaxPodsConstraint struct {
	empty          bool    `json:"-"`
	MaxPodsPerNode *string `json:"maxPodsPerNode"`
}

// This object is used to assert a desired state where this ClusterDefaultMaxPodsConstraint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterDefaultMaxPodsConstraint *ClusterDefaultMaxPodsConstraint = &ClusterDefaultMaxPodsConstraint{empty: true}

func (r *ClusterDefaultMaxPodsConstraint) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterDefaultMaxPodsConstraint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterResourceUsageExportConfig struct {
	empty                         bool                                                       `json:"-"`
	BigqueryDestination           *ClusterResourceUsageExportConfigBigqueryDestination       `json:"bigqueryDestination"`
	EnableNetworkEgressMonitoring *bool                                                      `json:"enableNetworkEgressMonitoring"`
	ConsumptionMeteringConfig     *ClusterResourceUsageExportConfigConsumptionMeteringConfig `json:"consumptionMeteringConfig"`
	EnableNetworkEgressMetering   *bool                                                      `json:"enableNetworkEgressMetering"`
}

// This object is used to assert a desired state where this ClusterResourceUsageExportConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterResourceUsageExportConfig *ClusterResourceUsageExportConfig = &ClusterResourceUsageExportConfig{empty: true}

func (r *ClusterResourceUsageExportConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterResourceUsageExportConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterResourceUsageExportConfigBigqueryDestination struct {
	empty     bool    `json:"-"`
	DatasetId *string `json:"datasetId"`
}

// This object is used to assert a desired state where this ClusterResourceUsageExportConfigBigqueryDestination is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterResourceUsageExportConfigBigqueryDestination *ClusterResourceUsageExportConfigBigqueryDestination = &ClusterResourceUsageExportConfigBigqueryDestination{empty: true}

func (r *ClusterResourceUsageExportConfigBigqueryDestination) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterResourceUsageExportConfigBigqueryDestination) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterResourceUsageExportConfigConsumptionMeteringConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterResourceUsageExportConfigConsumptionMeteringConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterResourceUsageExportConfigConsumptionMeteringConfig *ClusterResourceUsageExportConfigConsumptionMeteringConfig = &ClusterResourceUsageExportConfigConsumptionMeteringConfig{empty: true}

func (r *ClusterResourceUsageExportConfigConsumptionMeteringConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterResourceUsageExportConfigConsumptionMeteringConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAuthenticatorGroupsConfig struct {
	empty         bool    `json:"-"`
	Enabled       *bool   `json:"enabled"`
	SecurityGroup *string `json:"securityGroup"`
}

// This object is used to assert a desired state where this ClusterAuthenticatorGroupsConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAuthenticatorGroupsConfig *ClusterAuthenticatorGroupsConfig = &ClusterAuthenticatorGroupsConfig{empty: true}

func (r *ClusterAuthenticatorGroupsConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAuthenticatorGroupsConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterPrivateClusterConfig struct {
	empty                    bool                                                 `json:"-"`
	EnablePrivateNodes       *bool                                                `json:"enablePrivateNodes"`
	EnablePrivateEndpoint    *bool                                                `json:"enablePrivateEndpoint"`
	MasterIPv4CidrBlock      *string                                              `json:"masterIPv4CidrBlock"`
	PrivateEndpoint          *string                                              `json:"privateEndpoint"`
	PublicEndpoint           *string                                              `json:"publicEndpoint"`
	PeeringName              *string                                              `json:"peeringName"`
	MasterGlobalAccessConfig *ClusterPrivateClusterConfigMasterGlobalAccessConfig `json:"masterGlobalAccessConfig"`
}

// This object is used to assert a desired state where this ClusterPrivateClusterConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterPrivateClusterConfig *ClusterPrivateClusterConfig = &ClusterPrivateClusterConfig{empty: true}

func (r *ClusterPrivateClusterConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterPrivateClusterConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterPrivateClusterConfigMasterGlobalAccessConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterPrivateClusterConfigMasterGlobalAccessConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterPrivateClusterConfigMasterGlobalAccessConfig *ClusterPrivateClusterConfigMasterGlobalAccessConfig = &ClusterPrivateClusterConfigMasterGlobalAccessConfig{empty: true}

func (r *ClusterPrivateClusterConfigMasterGlobalAccessConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterPrivateClusterConfigMasterGlobalAccessConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterDatabaseEncryption struct {
	empty   bool                                `json:"-"`
	State   *ClusterDatabaseEncryptionStateEnum `json:"state"`
	KeyName *string                             `json:"keyName"`
}

// This object is used to assert a desired state where this ClusterDatabaseEncryption is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterDatabaseEncryption *ClusterDatabaseEncryption = &ClusterDatabaseEncryption{empty: true}

func (r *ClusterDatabaseEncryption) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterDatabaseEncryption) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterVerticalPodAutoscaling struct {
	empty                      bool  `json:"-"`
	Enabled                    *bool `json:"enabled"`
	EnableExperimentalFeatures *bool `json:"enableExperimentalFeatures"`
}

// This object is used to assert a desired state where this ClusterVerticalPodAutoscaling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterVerticalPodAutoscaling *ClusterVerticalPodAutoscaling = &ClusterVerticalPodAutoscaling{empty: true}

func (r *ClusterVerticalPodAutoscaling) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterVerticalPodAutoscaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterShieldedNodes struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterShieldedNodes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterShieldedNodes *ClusterShieldedNodes = &ClusterShieldedNodes{empty: true}

func (r *ClusterShieldedNodes) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterShieldedNodes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterConditions struct {
	empty         bool                                `json:"-"`
	Code          *string                             `json:"code"`
	Message       *string                             `json:"message"`
	CanonicalCode *ClusterConditionsCanonicalCodeEnum `json:"canonicalCode"`
}

// This object is used to assert a desired state where this ClusterConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterConditions *ClusterConditions = &ClusterConditions{empty: true}

func (r *ClusterConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterAutopilot struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterAutopilot is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterAutopilot *ClusterAutopilot = &ClusterAutopilot{empty: true}

func (r *ClusterAutopilot) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterAutopilot) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfig struct {
	empty                  bool                                     `json:"-"`
	MachineType            *string                                  `json:"machineType"`
	DiskSizeGb             *int64                                   `json:"diskSizeGb"`
	OAuthScopes            []string                                 `json:"oauthScopes"`
	ServiceAccount         *string                                  `json:"serviceAccount"`
	Metadata               map[string]string                        `json:"metadata"`
	ImageType              *string                                  `json:"imageType"`
	Labels                 map[string]string                        `json:"labels"`
	LocalSsdCount          *int64                                   `json:"localSsdCount"`
	Tags                   []string                                 `json:"tags"`
	Preemptible            *bool                                    `json:"preemptible"`
	Accelerators           []ClusterNodeConfigAccelerators          `json:"accelerators"`
	DiskType               *string                                  `json:"diskType"`
	MinCpuPlatform         *string                                  `json:"minCpuPlatform"`
	WorkloadMetadataConfig *ClusterNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfig"`
	Taints                 []ClusterNodeConfigTaints                `json:"taints"`
	SandboxConfig          *ClusterNodeConfigSandboxConfig          `json:"sandboxConfig"`
	NodeGroup              *string                                  `json:"nodeGroup"`
	ReservationAffinity    *ClusterNodeConfigReservationAffinity    `json:"reservationAffinity"`
	ShieldedInstanceConfig *ClusterNodeConfigShieldedInstanceConfig `json:"shieldedInstanceConfig"`
	LinuxNodeConfig        *ClusterNodeConfigLinuxNodeConfig        `json:"linuxNodeConfig"`
	KubeletConfig          *ClusterNodeConfigKubeletConfig          `json:"kubeletConfig"`
	BootDiskKmsKey         *string                                  `json:"bootDiskKmsKey"`
	EphemeralStorageConfig *ClusterNodeConfigEphemeralStorageConfig `json:"ephemeralStorageConfig"`
}

// This object is used to assert a desired state where this ClusterNodeConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfig *ClusterNodeConfig = &ClusterNodeConfig{empty: true}

func (r *ClusterNodeConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigAccelerators struct {
	empty            bool    `json:"-"`
	AcceleratorCount *int64  `json:"acceleratorCount"`
	AcceleratorType  *string `json:"acceleratorType"`
}

// This object is used to assert a desired state where this ClusterNodeConfigAccelerators is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigAccelerators *ClusterNodeConfigAccelerators = &ClusterNodeConfigAccelerators{empty: true}

func (r *ClusterNodeConfigAccelerators) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigAccelerators) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigWorkloadMetadataConfig struct {
	empty        bool                                                     `json:"-"`
	Mode         *ClusterNodeConfigWorkloadMetadataConfigModeEnum         `json:"mode"`
	NodeMetadata *ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum `json:"nodeMetadata"`
}

// This object is used to assert a desired state where this ClusterNodeConfigWorkloadMetadataConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigWorkloadMetadataConfig *ClusterNodeConfigWorkloadMetadataConfig = &ClusterNodeConfigWorkloadMetadataConfig{empty: true}

func (r *ClusterNodeConfigWorkloadMetadataConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigWorkloadMetadataConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigTaints struct {
	empty  bool                               `json:"-"`
	Key    *string                            `json:"key"`
	Value  *string                            `json:"value"`
	Effect *ClusterNodeConfigTaintsEffectEnum `json:"effect"`
}

// This object is used to assert a desired state where this ClusterNodeConfigTaints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigTaints *ClusterNodeConfigTaints = &ClusterNodeConfigTaints{empty: true}

func (r *ClusterNodeConfigTaints) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigTaints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigSandboxConfig struct {
	empty       bool                                    `json:"-"`
	Type        *ClusterNodeConfigSandboxConfigTypeEnum `json:"type"`
	SandboxType *string                                 `json:"sandboxType"`
}

// This object is used to assert a desired state where this ClusterNodeConfigSandboxConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigSandboxConfig *ClusterNodeConfigSandboxConfig = &ClusterNodeConfigSandboxConfig{empty: true}

func (r *ClusterNodeConfigSandboxConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigSandboxConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigReservationAffinity struct {
	empty                  bool                                                            `json:"-"`
	ConsumeReservationType *ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum `json:"consumeReservationType"`
	Key                    *string                                                         `json:"key"`
	Values                 []string                                                        `json:"values"`
}

// This object is used to assert a desired state where this ClusterNodeConfigReservationAffinity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigReservationAffinity *ClusterNodeConfigReservationAffinity = &ClusterNodeConfigReservationAffinity{empty: true}

func (r *ClusterNodeConfigReservationAffinity) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigReservationAffinity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigShieldedInstanceConfig struct {
	empty                     bool  `json:"-"`
	EnableSecureBoot          *bool `json:"enableSecureBoot"`
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring"`
}

// This object is used to assert a desired state where this ClusterNodeConfigShieldedInstanceConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigShieldedInstanceConfig *ClusterNodeConfigShieldedInstanceConfig = &ClusterNodeConfigShieldedInstanceConfig{empty: true}

func (r *ClusterNodeConfigShieldedInstanceConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigShieldedInstanceConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigLinuxNodeConfig struct {
	empty   bool              `json:"-"`
	Sysctls map[string]string `json:"sysctls"`
}

// This object is used to assert a desired state where this ClusterNodeConfigLinuxNodeConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigLinuxNodeConfig *ClusterNodeConfigLinuxNodeConfig = &ClusterNodeConfigLinuxNodeConfig{empty: true}

func (r *ClusterNodeConfigLinuxNodeConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigLinuxNodeConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigKubeletConfig struct {
	empty             bool    `json:"-"`
	CpuManagerPolicy  *string `json:"cpuManagerPolicy"`
	CpuCfsQuota       *bool   `json:"cpuCfsQuota"`
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod"`
}

// This object is used to assert a desired state where this ClusterNodeConfigKubeletConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigKubeletConfig *ClusterNodeConfigKubeletConfig = &ClusterNodeConfigKubeletConfig{empty: true}

func (r *ClusterNodeConfigKubeletConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigKubeletConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNodeConfigEphemeralStorageConfig struct {
	empty         bool   `json:"-"`
	LocalSsdCount *int64 `json:"localSsdCount"`
}

// This object is used to assert a desired state where this ClusterNodeConfigEphemeralStorageConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNodeConfigEphemeralStorageConfig *ClusterNodeConfigEphemeralStorageConfig = &ClusterNodeConfigEphemeralStorageConfig{empty: true}

func (r *ClusterNodeConfigEphemeralStorageConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNodeConfigEphemeralStorageConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterReleaseChannel struct {
	empty   bool                              `json:"-"`
	Channel *ClusterReleaseChannelChannelEnum `json:"channel"`
}

// This object is used to assert a desired state where this ClusterReleaseChannel is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterReleaseChannel *ClusterReleaseChannel = &ClusterReleaseChannel{empty: true}

func (r *ClusterReleaseChannel) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterReleaseChannel) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterWorkloadIdentityConfig struct {
	empty             bool    `json:"-"`
	WorkloadPool      *string `json:"workloadPool"`
	IdentityNamespace *string `json:"identityNamespace"`
	IdentityProvider  *string `json:"identityProvider"`
}

// This object is used to assert a desired state where this ClusterWorkloadIdentityConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterWorkloadIdentityConfig *ClusterWorkloadIdentityConfig = &ClusterWorkloadIdentityConfig{empty: true}

func (r *ClusterWorkloadIdentityConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterWorkloadIdentityConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNotificationConfig struct {
	empty  bool                             `json:"-"`
	Pubsub *ClusterNotificationConfigPubsub `json:"pubsub"`
}

// This object is used to assert a desired state where this ClusterNotificationConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNotificationConfig *ClusterNotificationConfig = &ClusterNotificationConfig{empty: true}

func (r *ClusterNotificationConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNotificationConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterNotificationConfigPubsub struct {
	empty   bool    `json:"-"`
	Enabled *bool   `json:"enabled"`
	Topic   *string `json:"topic"`
}

// This object is used to assert a desired state where this ClusterNotificationConfigPubsub is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterNotificationConfigPubsub *ClusterNotificationConfigPubsub = &ClusterNotificationConfigPubsub{empty: true}

func (r *ClusterNotificationConfigPubsub) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterNotificationConfigPubsub) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterConfidentialNodes struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterConfidentialNodes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterConfidentialNodes *ClusterConfidentialNodes = &ClusterConfidentialNodes{empty: true}

func (r *ClusterConfidentialNodes) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterConfidentialNodes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterPodSecurityPolicyConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

// This object is used to assert a desired state where this ClusterPodSecurityPolicyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterPodSecurityPolicyConfig *ClusterPodSecurityPolicyConfig = &ClusterPodSecurityPolicyConfig{empty: true}

func (r *ClusterPodSecurityPolicyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterPodSecurityPolicyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterClusterTelemetry struct {
	empty bool                             `json:"-"`
	Type  *ClusterClusterTelemetryTypeEnum `json:"type"`
}

// This object is used to assert a desired state where this ClusterClusterTelemetry is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterClusterTelemetry *ClusterClusterTelemetry = &ClusterClusterTelemetry{empty: true}

func (r *ClusterClusterTelemetry) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterClusterTelemetry) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterTPUConfig struct {
	empty                bool    `json:"-"`
	Enabled              *bool   `json:"enabled"`
	UseServiceNetworking *bool   `json:"useServiceNetworking"`
	IPv4CidrBlock        *string `json:"ipv4CidrBlock"`
}

// This object is used to assert a desired state where this ClusterTPUConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterTPUConfig *ClusterTPUConfig = &ClusterTPUConfig{empty: true}

func (r *ClusterTPUConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterTPUConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClusterMaster struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this ClusterMaster is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyClusterMaster *ClusterMaster = &ClusterMaster{empty: true}

func (r *ClusterMaster) String() string {
	return dcl.SprintResource(r)
}

func (r *ClusterMaster) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Cluster) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "container",
		Type:    "Cluster",
		Version: "beta",
	}
}

const ClusterMaxPage = -1

type ClusterList struct {
	Items []*Cluster

	nextToken string

	project string

	location string
}

func (l *ClusterList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ClusterList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listCluster(ctx, l.project, l.location, l.nextToken)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListCluster(ctx context.Context, project, location string) (*ClusterList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listCluster(ctx, project, location, "")
	if err != nil {
		return nil, err
	}
	return &ClusterList{
		Items:     items,
		nextToken: token,

		project: project,

		location: location,
	}, nil

}

func (c *Client) GetCluster(ctx context.Context, r *Cluster) (*Cluster, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getClusterRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalCluster(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeClusterNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteCluster(ctx context.Context, r *Cluster) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Cluster resource is nil")
	}
	c.Config.Logger.Info("Deleting Cluster...")
	deleteOp := deleteClusterOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllCluster deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllCluster(ctx context.Context, project, location string, filter func(*Cluster) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListCluster(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllCluster(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllCluster(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyCluster(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (*Cluster, error) {
	c.Config.Logger.Info("Beginning ApplyCluster...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.clusterDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []clusterApiOperation
	if create {
		ops = append(ops, &createClusterOperation{})
	} else if recreate {

		ops = append(ops, &deleteClusterOperation{})

		ops = append(ops, &createClusterOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeClusterDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetCluster(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createClusterOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapCluster(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeClusterNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeClusterNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeClusterDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffCluster(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
