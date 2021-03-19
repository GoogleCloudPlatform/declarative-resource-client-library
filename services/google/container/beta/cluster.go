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
}

func (r *Cluster) String() string {
	return dcl.SprintResource(r)
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
	empty                    bool                                         `json:"-"`
	HttpLoadBalancing        *ClusterAddonsConfigHttpLoadBalancing        `json:"httpLoadBalancing"`
	HorizontalPodAutoscaling *ClusterAddonsConfigHorizontalPodAutoscaling `json:"horizontalPodAutoscaling"`
	KubernetesDashboard      *ClusterAddonsConfigKubernetesDashboard      `json:"kubernetesDashboard"`
	NetworkPolicyConfig      *ClusterAddonsConfigNetworkPolicyConfig      `json:"networkPolicyConfig"`
	CloudRunConfig           *ClusterAddonsConfigCloudRunConfig           `json:"cloudRunConfig"`
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
	empty    bool  `json:"-"`
	Disabled *bool `json:"disabled"`
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

type ClusterNodePools struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
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
	empty           bool                                                               `json:"-"`
	OAuthScopes     []string                                                           `json:"oauthScopes"`
	ServiceAccount  *string                                                            `json:"serviceAccount"`
	UpgradeSettings *ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings `json:"upgradeSettings"`
	Management      *ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement      `json:"management"`
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
	empty       bool  `json:"-"`
	AutoUpgrade *bool `json:"autoUpgrade"`
	AutoRepair  *bool `json:"autoRepair"`
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

type ClusterNetworkConfig struct {
	empty                     bool    `json:"-"`
	Network                   *string `json:"network"`
	Subnetwork                *string `json:"subnetwork"`
	EnableIntraNodeVisibility *bool   `json:"enableIntraNodeVisibility"`
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
	empty                 bool    `json:"-"`
	EnablePrivateNodes    *bool   `json:"enablePrivateNodes"`
	EnablePrivateEndpoint *bool   `json:"enablePrivateEndpoint"`
	MasterIPv4CidrBlock   *string `json:"masterIPv4CidrBlock"`
	PrivateEndpoint       *string `json:"privateEndpoint"`
	PublicEndpoint        *string `json:"publicEndpoint"`
	PeeringName           *string `json:"peeringName"`
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
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
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
	empty   bool    `json:"-"`
	Code    *string `json:"code"`
	Message *string `json:"message"`
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
