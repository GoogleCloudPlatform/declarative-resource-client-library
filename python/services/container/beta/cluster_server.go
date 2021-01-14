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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/container/beta/container_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/beta"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterNetworkPolicyProviderEnum converts a ClusterNetworkPolicyProviderEnum enum from its proto representation.
func ProtoToContainerBetaClusterNetworkPolicyProviderEnum(e betapb.ContainerBetaClusterNetworkPolicyProviderEnum) *beta.ClusterNetworkPolicyProviderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNetworkPolicyProviderEnum_name[int32(e)]; ok {
		e := beta.ClusterNetworkPolicyProviderEnum(n[len("ClusterNetworkPolicyProviderEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterDatabaseEncryptionStateEnum converts a ClusterDatabaseEncryptionStateEnum enum from its proto representation.
func ProtoToContainerBetaClusterDatabaseEncryptionStateEnum(e betapb.ContainerBetaClusterDatabaseEncryptionStateEnum) *beta.ClusterDatabaseEncryptionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterDatabaseEncryptionStateEnum_name[int32(e)]; ok {
		e := beta.ClusterDatabaseEncryptionStateEnum(n[len("ClusterDatabaseEncryptionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterMasterAuth converts a ClusterMasterAuth resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuth(p *betapb.ContainerBetaClusterMasterAuth) *beta.ClusterMasterAuth {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuth{
		Username:                dcl.StringOrNil(p.Username),
		Password:                dcl.StringOrNil(p.Password),
		ClientCertificateConfig: ProtoToContainerBetaClusterMasterAuthClientCertificateConfig(p.GetClientCertificateConfig()),
		ClusterCaCertificate:    dcl.StringOrNil(p.ClusterCaCertificate),
		ClientCertificate:       dcl.StringOrNil(p.ClientCertificate),
		ClientKey:               dcl.StringOrNil(p.ClientKey),
	}
	return obj
}

// ProtoToClusterMasterAuthClientCertificateConfig converts a ClusterMasterAuthClientCertificateConfig resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuthClientCertificateConfig(p *betapb.ContainerBetaClusterMasterAuthClientCertificateConfig) *beta.ClusterMasterAuthClientCertificateConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.Bool(p.IssueClientCertificate),
	}
	return obj
}

// ProtoToClusterAddonsConfig converts a ClusterAddonsConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfig(p *betapb.ContainerBetaClusterAddonsConfig) *beta.ClusterAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfig{
		HttpLoadBalancing:        ProtoToContainerBetaClusterAddonsConfigHttpLoadBalancing(p.GetHttpLoadBalancing()),
		HorizontalPodAutoscaling: ProtoToContainerBetaClusterAddonsConfigHorizontalPodAutoscaling(p.GetHorizontalPodAutoscaling()),
		KubernetesDashboard:      ProtoToContainerBetaClusterAddonsConfigKubernetesDashboard(p.GetKubernetesDashboard()),
		NetworkPolicyConfig:      ProtoToContainerBetaClusterAddonsConfigNetworkPolicyConfig(p.GetNetworkPolicyConfig()),
		CloudRunConfig:           ProtoToContainerBetaClusterAddonsConfigCloudRunConfig(p.GetCloudRunConfig()),
	}
	return obj
}

// ProtoToClusterAddonsConfigHttpLoadBalancing converts a ClusterAddonsConfigHttpLoadBalancing resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigHttpLoadBalancing(p *betapb.ContainerBetaClusterAddonsConfigHttpLoadBalancing) *beta.ClusterAddonsConfigHttpLoadBalancing {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigHorizontalPodAutoscaling converts a ClusterAddonsConfigHorizontalPodAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigHorizontalPodAutoscaling(p *betapb.ContainerBetaClusterAddonsConfigHorizontalPodAutoscaling) *beta.ClusterAddonsConfigHorizontalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigKubernetesDashboard converts a ClusterAddonsConfigKubernetesDashboard resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigKubernetesDashboard(p *betapb.ContainerBetaClusterAddonsConfigKubernetesDashboard) *beta.ClusterAddonsConfigKubernetesDashboard {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigNetworkPolicyConfig converts a ClusterAddonsConfigNetworkPolicyConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigNetworkPolicyConfig(p *betapb.ContainerBetaClusterAddonsConfigNetworkPolicyConfig) *beta.ClusterAddonsConfigNetworkPolicyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigCloudRunConfig converts a ClusterAddonsConfigCloudRunConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigCloudRunConfig(p *betapb.ContainerBetaClusterAddonsConfigCloudRunConfig) *beta.ClusterAddonsConfigCloudRunConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigCloudRunConfig{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterNodePools converts a ClusterNodePools resource from its proto representation.
func ProtoToContainerBetaClusterNodePools(p *betapb.ContainerBetaClusterNodePools) *beta.ClusterNodePools {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePools{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToClusterLegacyAbac converts a ClusterLegacyAbac resource from its proto representation.
func ProtoToContainerBetaClusterLegacyAbac(p *betapb.ContainerBetaClusterLegacyAbac) *beta.ClusterLegacyAbac {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterLegacyAbac{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNetworkPolicy converts a ClusterNetworkPolicy resource from its proto representation.
func ProtoToContainerBetaClusterNetworkPolicy(p *betapb.ContainerBetaClusterNetworkPolicy) *beta.ClusterNetworkPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworkPolicy{
		Provider: ProtoToContainerBetaClusterNetworkPolicyProviderEnum(p.GetProvider()),
		Enabled:  dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterIPAllocationPolicy converts a ClusterIPAllocationPolicy resource from its proto representation.
func ProtoToContainerBetaClusterIPAllocationPolicy(p *betapb.ContainerBetaClusterIPAllocationPolicy) *beta.ClusterIPAllocationPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterIPAllocationPolicy{
		UseIPAliases:               dcl.Bool(p.UseIpAliases),
		CreateSubnetwork:           dcl.Bool(p.CreateSubnetwork),
		SubnetworkName:             dcl.StringOrNil(p.SubnetworkName),
		ClusterSecondaryRangeName:  dcl.StringOrNil(p.ClusterSecondaryRangeName),
		ServicesSecondaryRangeName: dcl.StringOrNil(p.ServicesSecondaryRangeName),
		ClusterIPv4CidrBlock:       dcl.StringOrNil(p.ClusterIpv4CidrBlock),
		NodeIPv4CidrBlock:          dcl.StringOrNil(p.NodeIpv4CidrBlock),
		ServicesIPv4CidrBlock:      dcl.StringOrNil(p.ServicesIpv4CidrBlock),
		TPUIPv4CidrBlock:           dcl.StringOrNil(p.TpuIpv4CidrBlock),
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfig converts a ClusterMasterAuthorizedNetworksConfig resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuthorizedNetworksConfig(p *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfig) *beta.ClusterMasterAuthorizedNetworksConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	for _, r := range p.GetCidrBlocks() {
		obj.CidrBlocks = append(obj.CidrBlocks, *ProtoToContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks(r))
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfigCidrBlocks converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks(p *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks) *beta.ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.StringOrNil(p.DisplayName),
		CidrBlock:   dcl.StringOrNil(p.CidrBlock),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization resource from its proto representation.
func ProtoToContainerBetaClusterBinaryAuthorization(p *betapb.ContainerBetaClusterBinaryAuthorization) *beta.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterBinaryAuthorization{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAutoscaling converts a ClusterAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterAutoscaling(p *betapb.ContainerBetaClusterAutoscaling) *beta.ClusterAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.Bool(p.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults(p.GetAutoprovisioningNodePoolDefaults()),
	}
	for _, r := range p.GetResourceLimits() {
		obj.ResourceLimits = append(obj.ResourceLimits, *ProtoToContainerBetaClusterAutoscalingResourceLimits(r))
	}
	return obj
}

// ProtoToClusterAutoscalingResourceLimits converts a ClusterAutoscalingResourceLimits resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingResourceLimits(p *betapb.ContainerBetaClusterAutoscalingResourceLimits) *beta.ClusterAutoscalingResourceLimits {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingResourceLimits{
		ResourceType: dcl.StringOrNil(p.ResourceType),
		Minimum:      dcl.Int64OrNil(p.Minimum),
		Maximum:      dcl.Int64OrNil(p.Maximum),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaults converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:  dcl.StringOrNil(p.ServiceAccount),
		UpgradeSettings: ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p.GetUpgradeSettings()),
		Management:      ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p.GetManagement()),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade: dcl.Bool(p.AutoUpgrade),
		AutoRepair:  dcl.Bool(p.AutoRepair),
	}
	return obj
}

// ProtoToClusterNetworkConfig converts a ClusterNetworkConfig resource from its proto representation.
func ProtoToContainerBetaClusterNetworkConfig(p *betapb.ContainerBetaClusterNetworkConfig) *beta.ClusterNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworkConfig{
		Network:                   dcl.StringOrNil(p.Network),
		Subnetwork:                dcl.StringOrNil(p.Subnetwork),
		EnableIntraNodeVisibility: dcl.Bool(p.EnableIntraNodeVisibility),
	}
	return obj
}

// ProtoToClusterMaintenancePolicy converts a ClusterMaintenancePolicy resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicy(p *betapb.ContainerBetaClusterMaintenancePolicy) *beta.ClusterMaintenancePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicy{
		Window:          ProtoToContainerBetaClusterMaintenancePolicyWindow(p.GetWindow()),
		ResourceVersion: dcl.StringOrNil(p.ResourceVersion),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindow converts a ClusterMaintenancePolicyWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindow) *beta.ClusterMaintenancePolicyWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ProtoToContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow(p.GetDailyMaintenanceWindow()),
		RecurringWindow:        ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindow(p.GetRecurringWindow()),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowDailyMaintenanceWindow converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow) *beta.ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		Duration:  dcl.StringOrNil(p.Duration),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindow converts a ClusterMaintenancePolicyWindowRecurringWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindow) *beta.ClusterMaintenancePolicyWindowRecurringWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow(p.GetWindow()),
		Recurrence: dcl.StringOrNil(p.Recurrence),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindowWindow converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow) *beta.ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		EndTime:   dcl.StringOrNil(p.GetEndTime()),
	}
	return obj
}

// ProtoToClusterDefaultMaxPodsConstraint converts a ClusterDefaultMaxPodsConstraint resource from its proto representation.
func ProtoToContainerBetaClusterDefaultMaxPodsConstraint(p *betapb.ContainerBetaClusterDefaultMaxPodsConstraint) *beta.ClusterDefaultMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.StringOrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfig converts a ClusterResourceUsageExportConfig resource from its proto representation.
func ProtoToContainerBetaClusterResourceUsageExportConfig(p *betapb.ContainerBetaClusterResourceUsageExportConfig) *beta.ClusterResourceUsageExportConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterResourceUsageExportConfig{
		BigqueryDestination:           ProtoToContainerBetaClusterResourceUsageExportConfigBigqueryDestination(p.GetBigqueryDestination()),
		EnableNetworkEgressMonitoring: dcl.Bool(p.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ProtoToContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig(p.GetConsumptionMeteringConfig()),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigBigqueryDestination converts a ClusterResourceUsageExportConfigBigqueryDestination resource from its proto representation.
func ProtoToContainerBetaClusterResourceUsageExportConfigBigqueryDestination(p *betapb.ContainerBetaClusterResourceUsageExportConfigBigqueryDestination) *beta.ClusterResourceUsageExportConfigBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.StringOrNil(p.DatasetId),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigConsumptionMeteringConfig converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource from its proto representation.
func ProtoToContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig(p *betapb.ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig) *beta.ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAuthenticatorGroupsConfig converts a ClusterAuthenticatorGroupsConfig resource from its proto representation.
func ProtoToContainerBetaClusterAuthenticatorGroupsConfig(p *betapb.ContainerBetaClusterAuthenticatorGroupsConfig) *beta.ClusterAuthenticatorGroupsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.Bool(p.Enabled),
		SecurityGroup: dcl.StringOrNil(p.SecurityGroup),
	}
	return obj
}

// ProtoToClusterPrivateClusterConfig converts a ClusterPrivateClusterConfig resource from its proto representation.
func ProtoToContainerBetaClusterPrivateClusterConfig(p *betapb.ContainerBetaClusterPrivateClusterConfig) *beta.ClusterPrivateClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterPrivateClusterConfig{
		EnablePrivateNodes:    dcl.Bool(p.EnablePrivateNodes),
		EnablePrivateEndpoint: dcl.Bool(p.EnablePrivateEndpoint),
		MasterIPv4CidrBlock:   dcl.StringOrNil(p.MasterIpv4CidrBlock),
		PrivateEndpoint:       dcl.StringOrNil(p.PrivateEndpoint),
		PublicEndpoint:        dcl.StringOrNil(p.PublicEndpoint),
		PeeringName:           dcl.StringOrNil(p.PeeringName),
	}
	return obj
}

// ProtoToClusterDatabaseEncryption converts a ClusterDatabaseEncryption resource from its proto representation.
func ProtoToContainerBetaClusterDatabaseEncryption(p *betapb.ContainerBetaClusterDatabaseEncryption) *beta.ClusterDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterDatabaseEncryption{
		State:   ProtoToContainerBetaClusterDatabaseEncryptionStateEnum(p.GetState()),
		KeyName: dcl.StringOrNil(p.KeyName),
	}
	return obj
}

// ProtoToClusterVerticalPodAutoscaling converts a ClusterVerticalPodAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterVerticalPodAutoscaling(p *betapb.ContainerBetaClusterVerticalPodAutoscaling) *beta.ClusterVerticalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterVerticalPodAutoscaling{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterShieldedNodes converts a ClusterShieldedNodes resource from its proto representation.
func ProtoToContainerBetaClusterShieldedNodes(p *betapb.ContainerBetaClusterShieldedNodes) *beta.ClusterShieldedNodes {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterShieldedNodes{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterConditions converts a ClusterConditions resource from its proto representation.
func ProtoToContainerBetaClusterConditions(p *betapb.ContainerBetaClusterConditions) *beta.ClusterConditions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConditions{
		Code:    dcl.StringOrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *betapb.ContainerBetaCluster) *beta.Cluster {
	obj := &beta.Cluster{
		Name:                           dcl.StringOrNil(p.Name),
		Description:                    dcl.StringOrNil(p.Description),
		InitialNodeCount:               dcl.Int64OrNil(p.InitialNodeCount),
		MasterAuth:                     ProtoToContainerBetaClusterMasterAuth(p.GetMasterAuth()),
		LoggingService:                 dcl.StringOrNil(p.LoggingService),
		MonitoringService:              dcl.StringOrNil(p.MonitoringService),
		Network:                        dcl.StringOrNil(p.Network),
		ClusterIPv4Cidr:                dcl.StringOrNil(p.ClusterIpv4Cidr),
		AddonsConfig:                   ProtoToContainerBetaClusterAddonsConfig(p.GetAddonsConfig()),
		Subnetwork:                     dcl.StringOrNil(p.Subnetwork),
		EnableKubernetesAlpha:          dcl.Bool(p.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.StringOrNil(p.LabelFingerprint),
		LegacyAbac:                     ProtoToContainerBetaClusterLegacyAbac(p.GetLegacyAbac()),
		NetworkPolicy:                  ProtoToContainerBetaClusterNetworkPolicy(p.GetNetworkPolicy()),
		IPAllocationPolicy:             ProtoToContainerBetaClusterIPAllocationPolicy(p.GetIpAllocationPolicy()),
		MasterAuthorizedNetworksConfig: ProtoToContainerBetaClusterMasterAuthorizedNetworksConfig(p.GetMasterAuthorizedNetworksConfig()),
		BinaryAuthorization:            ProtoToContainerBetaClusterBinaryAuthorization(p.GetBinaryAuthorization()),
		Autoscaling:                    ProtoToContainerBetaClusterAutoscaling(p.GetAutoscaling()),
		NetworkConfig:                  ProtoToContainerBetaClusterNetworkConfig(p.GetNetworkConfig()),
		MaintenancePolicy:              ProtoToContainerBetaClusterMaintenancePolicy(p.GetMaintenancePolicy()),
		DefaultMaxPodsConstraint:       ProtoToContainerBetaClusterDefaultMaxPodsConstraint(p.GetDefaultMaxPodsConstraint()),
		ResourceUsageExportConfig:      ProtoToContainerBetaClusterResourceUsageExportConfig(p.GetResourceUsageExportConfig()),
		AuthenticatorGroupsConfig:      ProtoToContainerBetaClusterAuthenticatorGroupsConfig(p.GetAuthenticatorGroupsConfig()),
		PrivateClusterConfig:           ProtoToContainerBetaClusterPrivateClusterConfig(p.GetPrivateClusterConfig()),
		DatabaseEncryption:             ProtoToContainerBetaClusterDatabaseEncryption(p.GetDatabaseEncryption()),
		VerticalPodAutoscaling:         ProtoToContainerBetaClusterVerticalPodAutoscaling(p.GetVerticalPodAutoscaling()),
		ShieldedNodes:                  ProtoToContainerBetaClusterShieldedNodes(p.GetShieldedNodes()),
		Endpoint:                       dcl.StringOrNil(p.Endpoint),
		MasterVersion:                  dcl.StringOrNil(p.MasterVersion),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		Status:                         dcl.StringOrNil(p.Status),
		StatusMessage:                  dcl.StringOrNil(p.StatusMessage),
		NodeIPv4CidrSize:               dcl.Int64OrNil(p.NodeIpv4CidrSize),
		ServicesIPv4Cidr:               dcl.StringOrNil(p.ServicesIpv4Cidr),
		ExpireTime:                     dcl.StringOrNil(p.GetExpireTime()),
		Location:                       dcl.StringOrNil(p.Location),
		EnableTPU:                      dcl.Bool(p.EnableTpu),
		TPUIPv4CidrBlock:               dcl.StringOrNil(p.TpuIpv4CidrBlock),
		Project:                        dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetNodePools() {
		obj.NodePools = append(obj.NodePools, *ProtoToContainerBetaClusterNodePools(r))
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerBetaClusterConditions(r))
	}
	return obj
}

// ClusterNetworkPolicyProviderEnumToProto converts a ClusterNetworkPolicyProviderEnum enum to its proto representation.
func ContainerBetaClusterNetworkPolicyProviderEnumToProto(e *beta.ClusterNetworkPolicyProviderEnum) betapb.ContainerBetaClusterNetworkPolicyProviderEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNetworkPolicyProviderEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNetworkPolicyProviderEnum_value["ClusterNetworkPolicyProviderEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNetworkPolicyProviderEnum(v)
	}
	return betapb.ContainerBetaClusterNetworkPolicyProviderEnum(0)
}

// ClusterDatabaseEncryptionStateEnumToProto converts a ClusterDatabaseEncryptionStateEnum enum to its proto representation.
func ContainerBetaClusterDatabaseEncryptionStateEnumToProto(e *beta.ClusterDatabaseEncryptionStateEnum) betapb.ContainerBetaClusterDatabaseEncryptionStateEnum {
	if e == nil {
		return betapb.ContainerBetaClusterDatabaseEncryptionStateEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterDatabaseEncryptionStateEnum_value["ClusterDatabaseEncryptionStateEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterDatabaseEncryptionStateEnum(v)
	}
	return betapb.ContainerBetaClusterDatabaseEncryptionStateEnum(0)
}

// ClusterMasterAuthToProto converts a ClusterMasterAuth resource to its proto representation.
func ContainerBetaClusterMasterAuthToProto(o *beta.ClusterMasterAuth) *betapb.ContainerBetaClusterMasterAuth {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuth{
		Username:                dcl.ValueOrEmptyString(o.Username),
		Password:                dcl.ValueOrEmptyString(o.Password),
		ClientCertificateConfig: ContainerBetaClusterMasterAuthClientCertificateConfigToProto(o.ClientCertificateConfig),
		ClusterCaCertificate:    dcl.ValueOrEmptyString(o.ClusterCaCertificate),
		ClientCertificate:       dcl.ValueOrEmptyString(o.ClientCertificate),
		ClientKey:               dcl.ValueOrEmptyString(o.ClientKey),
	}
	return p
}

// ClusterMasterAuthClientCertificateConfigToProto converts a ClusterMasterAuthClientCertificateConfig resource to its proto representation.
func ContainerBetaClusterMasterAuthClientCertificateConfigToProto(o *beta.ClusterMasterAuthClientCertificateConfig) *betapb.ContainerBetaClusterMasterAuthClientCertificateConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.ValueOrEmptyBool(o.IssueClientCertificate),
	}
	return p
}

// ClusterAddonsConfigToProto converts a ClusterAddonsConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigToProto(o *beta.ClusterAddonsConfig) *betapb.ContainerBetaClusterAddonsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfig{
		HttpLoadBalancing:        ContainerBetaClusterAddonsConfigHttpLoadBalancingToProto(o.HttpLoadBalancing),
		HorizontalPodAutoscaling: ContainerBetaClusterAddonsConfigHorizontalPodAutoscalingToProto(o.HorizontalPodAutoscaling),
		KubernetesDashboard:      ContainerBetaClusterAddonsConfigKubernetesDashboardToProto(o.KubernetesDashboard),
		NetworkPolicyConfig:      ContainerBetaClusterAddonsConfigNetworkPolicyConfigToProto(o.NetworkPolicyConfig),
		CloudRunConfig:           ContainerBetaClusterAddonsConfigCloudRunConfigToProto(o.CloudRunConfig),
	}
	return p
}

// ClusterAddonsConfigHttpLoadBalancingToProto converts a ClusterAddonsConfigHttpLoadBalancing resource to its proto representation.
func ContainerBetaClusterAddonsConfigHttpLoadBalancingToProto(o *beta.ClusterAddonsConfigHttpLoadBalancing) *betapb.ContainerBetaClusterAddonsConfigHttpLoadBalancing {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigHorizontalPodAutoscalingToProto converts a ClusterAddonsConfigHorizontalPodAutoscaling resource to its proto representation.
func ContainerBetaClusterAddonsConfigHorizontalPodAutoscalingToProto(o *beta.ClusterAddonsConfigHorizontalPodAutoscaling) *betapb.ContainerBetaClusterAddonsConfigHorizontalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigKubernetesDashboardToProto converts a ClusterAddonsConfigKubernetesDashboard resource to its proto representation.
func ContainerBetaClusterAddonsConfigKubernetesDashboardToProto(o *beta.ClusterAddonsConfigKubernetesDashboard) *betapb.ContainerBetaClusterAddonsConfigKubernetesDashboard {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigNetworkPolicyConfigToProto converts a ClusterAddonsConfigNetworkPolicyConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigNetworkPolicyConfigToProto(o *beta.ClusterAddonsConfigNetworkPolicyConfig) *betapb.ContainerBetaClusterAddonsConfigNetworkPolicyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigCloudRunConfigToProto converts a ClusterAddonsConfigCloudRunConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigCloudRunConfigToProto(o *beta.ClusterAddonsConfigCloudRunConfig) *betapb.ContainerBetaClusterAddonsConfigCloudRunConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigCloudRunConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterNodePoolsToProto converts a ClusterNodePools resource to its proto representation.
func ContainerBetaClusterNodePoolsToProto(o *beta.ClusterNodePools) *betapb.ContainerBetaClusterNodePools {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePools{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ClusterLegacyAbacToProto converts a ClusterLegacyAbac resource to its proto representation.
func ContainerBetaClusterLegacyAbacToProto(o *beta.ClusterLegacyAbac) *betapb.ContainerBetaClusterLegacyAbac {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterLegacyAbac{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNetworkPolicyToProto converts a ClusterNetworkPolicy resource to its proto representation.
func ContainerBetaClusterNetworkPolicyToProto(o *beta.ClusterNetworkPolicy) *betapb.ContainerBetaClusterNetworkPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNetworkPolicy{
		Provider: ContainerBetaClusterNetworkPolicyProviderEnumToProto(o.Provider),
		Enabled:  dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterIPAllocationPolicyToProto converts a ClusterIPAllocationPolicy resource to its proto representation.
func ContainerBetaClusterIPAllocationPolicyToProto(o *beta.ClusterIPAllocationPolicy) *betapb.ContainerBetaClusterIPAllocationPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterIPAllocationPolicy{
		UseIpAliases:               dcl.ValueOrEmptyBool(o.UseIPAliases),
		CreateSubnetwork:           dcl.ValueOrEmptyBool(o.CreateSubnetwork),
		SubnetworkName:             dcl.ValueOrEmptyString(o.SubnetworkName),
		ClusterSecondaryRangeName:  dcl.ValueOrEmptyString(o.ClusterSecondaryRangeName),
		ServicesSecondaryRangeName: dcl.ValueOrEmptyString(o.ServicesSecondaryRangeName),
		ClusterIpv4CidrBlock:       dcl.ValueOrEmptyString(o.ClusterIPv4CidrBlock),
		NodeIpv4CidrBlock:          dcl.ValueOrEmptyString(o.NodeIPv4CidrBlock),
		ServicesIpv4CidrBlock:      dcl.ValueOrEmptyString(o.ServicesIPv4CidrBlock),
		TpuIpv4CidrBlock:           dcl.ValueOrEmptyString(o.TPUIPv4CidrBlock),
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigToProto converts a ClusterMasterAuthorizedNetworksConfig resource to its proto representation.
func ContainerBetaClusterMasterAuthorizedNetworksConfigToProto(o *beta.ClusterMasterAuthorizedNetworksConfig) *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	for _, r := range o.CidrBlocks {
		p.CidrBlocks = append(p.CidrBlocks, ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(&r))
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigCidrBlocksToProto converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource to its proto representation.
func ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(o *beta.ClusterMasterAuthorizedNetworksConfigCidrBlocks) *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		CidrBlock:   dcl.ValueOrEmptyString(o.CidrBlock),
	}
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization resource to its proto representation.
func ContainerBetaClusterBinaryAuthorizationToProto(o *beta.ClusterBinaryAuthorization) *betapb.ContainerBetaClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterBinaryAuthorization{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAutoscalingToProto converts a ClusterAutoscaling resource to its proto representation.
func ContainerBetaClusterAutoscalingToProto(o *beta.ClusterAutoscaling) *betapb.ContainerBetaClusterAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.ValueOrEmptyBool(o.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o.AutoprovisioningNodePoolDefaults),
	}
	for _, r := range o.ResourceLimits {
		p.ResourceLimits = append(p.ResourceLimits, ContainerBetaClusterAutoscalingResourceLimitsToProto(&r))
	}
	return p
}

// ClusterAutoscalingResourceLimitsToProto converts a ClusterAutoscalingResourceLimits resource to its proto representation.
func ContainerBetaClusterAutoscalingResourceLimitsToProto(o *beta.ClusterAutoscalingResourceLimits) *betapb.ContainerBetaClusterAutoscalingResourceLimits {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingResourceLimits{
		ResourceType: dcl.ValueOrEmptyString(o.ResourceType),
		Minimum:      dcl.ValueOrEmptyInt64(o.Minimum),
		Maximum:      dcl.ValueOrEmptyInt64(o.Maximum),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaults) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:  dcl.ValueOrEmptyString(o.ServiceAccount),
		UpgradeSettings: ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o.UpgradeSettings),
		Management:      ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o.Management),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade: dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:  dcl.ValueOrEmptyBool(o.AutoRepair),
	}
	return p
}

// ClusterNetworkConfigToProto converts a ClusterNetworkConfig resource to its proto representation.
func ContainerBetaClusterNetworkConfigToProto(o *beta.ClusterNetworkConfig) *betapb.ContainerBetaClusterNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNetworkConfig{
		Network:                   dcl.ValueOrEmptyString(o.Network),
		Subnetwork:                dcl.ValueOrEmptyString(o.Subnetwork),
		EnableIntraNodeVisibility: dcl.ValueOrEmptyBool(o.EnableIntraNodeVisibility),
	}
	return p
}

// ClusterMaintenancePolicyToProto converts a ClusterMaintenancePolicy resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyToProto(o *beta.ClusterMaintenancePolicy) *betapb.ContainerBetaClusterMaintenancePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicy{
		Window:          ContainerBetaClusterMaintenancePolicyWindowToProto(o.Window),
		ResourceVersion: dcl.ValueOrEmptyString(o.ResourceVersion),
	}
	return p
}

// ClusterMaintenancePolicyWindowToProto converts a ClusterMaintenancePolicyWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowToProto(o *beta.ClusterMaintenancePolicyWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o.DailyMaintenanceWindow),
		RecurringWindow:        ContainerBetaClusterMaintenancePolicyWindowRecurringWindowToProto(o.RecurringWindow),
	}
	return p
}

// ClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o *beta.ClusterMaintenancePolicyWindowDailyMaintenanceWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		Duration:  dcl.ValueOrEmptyString(o.Duration),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowRecurringWindowToProto(o *beta.ClusterMaintenancePolicyWindowRecurringWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o.Window),
		Recurrence: dcl.ValueOrEmptyString(o.Recurrence),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o *beta.ClusterMaintenancePolicyWindowRecurringWindowWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// ClusterDefaultMaxPodsConstraintToProto converts a ClusterDefaultMaxPodsConstraint resource to its proto representation.
func ContainerBetaClusterDefaultMaxPodsConstraintToProto(o *beta.ClusterDefaultMaxPodsConstraint) *betapb.ContainerBetaClusterDefaultMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyString(o.MaxPodsPerNode),
	}
	return p
}

// ClusterResourceUsageExportConfigToProto converts a ClusterResourceUsageExportConfig resource to its proto representation.
func ContainerBetaClusterResourceUsageExportConfigToProto(o *beta.ClusterResourceUsageExportConfig) *betapb.ContainerBetaClusterResourceUsageExportConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterResourceUsageExportConfig{
		BigqueryDestination:           ContainerBetaClusterResourceUsageExportConfigBigqueryDestinationToProto(o.BigqueryDestination),
		EnableNetworkEgressMonitoring: dcl.ValueOrEmptyBool(o.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o.ConsumptionMeteringConfig),
	}
	return p
}

// ClusterResourceUsageExportConfigBigqueryDestinationToProto converts a ClusterResourceUsageExportConfigBigqueryDestination resource to its proto representation.
func ContainerBetaClusterResourceUsageExportConfigBigqueryDestinationToProto(o *beta.ClusterResourceUsageExportConfigBigqueryDestination) *betapb.ContainerBetaClusterResourceUsageExportConfigBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.ValueOrEmptyString(o.DatasetId),
	}
	return p
}

// ClusterResourceUsageExportConfigConsumptionMeteringConfigToProto converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource to its proto representation.
func ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o *beta.ClusterResourceUsageExportConfigConsumptionMeteringConfig) *betapb.ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAuthenticatorGroupsConfigToProto converts a ClusterAuthenticatorGroupsConfig resource to its proto representation.
func ContainerBetaClusterAuthenticatorGroupsConfigToProto(o *beta.ClusterAuthenticatorGroupsConfig) *betapb.ContainerBetaClusterAuthenticatorGroupsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.ValueOrEmptyBool(o.Enabled),
		SecurityGroup: dcl.ValueOrEmptyString(o.SecurityGroup),
	}
	return p
}

// ClusterPrivateClusterConfigToProto converts a ClusterPrivateClusterConfig resource to its proto representation.
func ContainerBetaClusterPrivateClusterConfigToProto(o *beta.ClusterPrivateClusterConfig) *betapb.ContainerBetaClusterPrivateClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterPrivateClusterConfig{
		EnablePrivateNodes:    dcl.ValueOrEmptyBool(o.EnablePrivateNodes),
		EnablePrivateEndpoint: dcl.ValueOrEmptyBool(o.EnablePrivateEndpoint),
		MasterIpv4CidrBlock:   dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock),
		PrivateEndpoint:       dcl.ValueOrEmptyString(o.PrivateEndpoint),
		PublicEndpoint:        dcl.ValueOrEmptyString(o.PublicEndpoint),
		PeeringName:           dcl.ValueOrEmptyString(o.PeeringName),
	}
	return p
}

// ClusterDatabaseEncryptionToProto converts a ClusterDatabaseEncryption resource to its proto representation.
func ContainerBetaClusterDatabaseEncryptionToProto(o *beta.ClusterDatabaseEncryption) *betapb.ContainerBetaClusterDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterDatabaseEncryption{
		State:   ContainerBetaClusterDatabaseEncryptionStateEnumToProto(o.State),
		KeyName: dcl.ValueOrEmptyString(o.KeyName),
	}
	return p
}

// ClusterVerticalPodAutoscalingToProto converts a ClusterVerticalPodAutoscaling resource to its proto representation.
func ContainerBetaClusterVerticalPodAutoscalingToProto(o *beta.ClusterVerticalPodAutoscaling) *betapb.ContainerBetaClusterVerticalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterVerticalPodAutoscaling{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterShieldedNodesToProto converts a ClusterShieldedNodes resource to its proto representation.
func ContainerBetaClusterShieldedNodesToProto(o *beta.ClusterShieldedNodes) *betapb.ContainerBetaClusterShieldedNodes {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterShieldedNodes{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterConditionsToProto converts a ClusterConditions resource to its proto representation.
func ContainerBetaClusterConditionsToProto(o *beta.ClusterConditions) *betapb.ContainerBetaClusterConditions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterConditions{
		Code:    dcl.ValueOrEmptyString(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *beta.Cluster) *betapb.ContainerBetaCluster {
	p := &betapb.ContainerBetaCluster{
		Name:                           dcl.ValueOrEmptyString(resource.Name),
		Description:                    dcl.ValueOrEmptyString(resource.Description),
		InitialNodeCount:               dcl.ValueOrEmptyInt64(resource.InitialNodeCount),
		MasterAuth:                     ContainerBetaClusterMasterAuthToProto(resource.MasterAuth),
		LoggingService:                 dcl.ValueOrEmptyString(resource.LoggingService),
		MonitoringService:              dcl.ValueOrEmptyString(resource.MonitoringService),
		Network:                        dcl.ValueOrEmptyString(resource.Network),
		ClusterIpv4Cidr:                dcl.ValueOrEmptyString(resource.ClusterIPv4Cidr),
		AddonsConfig:                   ContainerBetaClusterAddonsConfigToProto(resource.AddonsConfig),
		Subnetwork:                     dcl.ValueOrEmptyString(resource.Subnetwork),
		EnableKubernetesAlpha:          dcl.ValueOrEmptyBool(resource.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.ValueOrEmptyString(resource.LabelFingerprint),
		LegacyAbac:                     ContainerBetaClusterLegacyAbacToProto(resource.LegacyAbac),
		NetworkPolicy:                  ContainerBetaClusterNetworkPolicyToProto(resource.NetworkPolicy),
		IpAllocationPolicy:             ContainerBetaClusterIPAllocationPolicyToProto(resource.IPAllocationPolicy),
		MasterAuthorizedNetworksConfig: ContainerBetaClusterMasterAuthorizedNetworksConfigToProto(resource.MasterAuthorizedNetworksConfig),
		BinaryAuthorization:            ContainerBetaClusterBinaryAuthorizationToProto(resource.BinaryAuthorization),
		Autoscaling:                    ContainerBetaClusterAutoscalingToProto(resource.Autoscaling),
		NetworkConfig:                  ContainerBetaClusterNetworkConfigToProto(resource.NetworkConfig),
		MaintenancePolicy:              ContainerBetaClusterMaintenancePolicyToProto(resource.MaintenancePolicy),
		DefaultMaxPodsConstraint:       ContainerBetaClusterDefaultMaxPodsConstraintToProto(resource.DefaultMaxPodsConstraint),
		ResourceUsageExportConfig:      ContainerBetaClusterResourceUsageExportConfigToProto(resource.ResourceUsageExportConfig),
		AuthenticatorGroupsConfig:      ContainerBetaClusterAuthenticatorGroupsConfigToProto(resource.AuthenticatorGroupsConfig),
		PrivateClusterConfig:           ContainerBetaClusterPrivateClusterConfigToProto(resource.PrivateClusterConfig),
		DatabaseEncryption:             ContainerBetaClusterDatabaseEncryptionToProto(resource.DatabaseEncryption),
		VerticalPodAutoscaling:         ContainerBetaClusterVerticalPodAutoscalingToProto(resource.VerticalPodAutoscaling),
		ShieldedNodes:                  ContainerBetaClusterShieldedNodesToProto(resource.ShieldedNodes),
		Endpoint:                       dcl.ValueOrEmptyString(resource.Endpoint),
		MasterVersion:                  dcl.ValueOrEmptyString(resource.MasterVersion),
		CreateTime:                     dcl.ValueOrEmptyString(resource.CreateTime),
		Status:                         dcl.ValueOrEmptyString(resource.Status),
		StatusMessage:                  dcl.ValueOrEmptyString(resource.StatusMessage),
		NodeIpv4CidrSize:               dcl.ValueOrEmptyInt64(resource.NodeIPv4CidrSize),
		ServicesIpv4Cidr:               dcl.ValueOrEmptyString(resource.ServicesIPv4Cidr),
		ExpireTime:                     dcl.ValueOrEmptyString(resource.ExpireTime),
		Location:                       dcl.ValueOrEmptyString(resource.Location),
		EnableTpu:                      dcl.ValueOrEmptyBool(resource.EnableTPU),
		TpuIpv4CidrBlock:               dcl.ValueOrEmptyString(resource.TPUIPv4CidrBlock),
		Project:                        dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.NodePools {
		p.NodePools = append(p.NodePools, ContainerBetaClusterNodePoolsToProto(&r))
	}
	for _, r := range resource.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, ContainerBetaClusterConditionsToProto(&r))
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerBetaClusterRequest) (*betapb.ContainerBetaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerBetaCluster(ctx context.Context, request *betapb.ApplyContainerBetaClusterRequest) (*betapb.ContainerBetaCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerBetaCluster(ctx context.Context, request *betapb.DeleteContainerBetaClusterRequest) (*emptypb.Empty, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))
}

// ListCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerBetaCluster(ctx context.Context, request *betapb.ListContainerBetaClusterRequest) (*betapb.ListContainerBetaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerBetaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListContainerBetaClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
