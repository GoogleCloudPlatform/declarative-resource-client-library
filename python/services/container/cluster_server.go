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
	containerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/container/container_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterNetworkPolicyProviderEnum converts a ClusterNetworkPolicyProviderEnum enum from its proto representation.
func ProtoToContainerClusterNetworkPolicyProviderEnum(e containerpb.ContainerClusterNetworkPolicyProviderEnum) *container.ClusterNetworkPolicyProviderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNetworkPolicyProviderEnum_name[int32(e)]; ok {
		e := container.ClusterNetworkPolicyProviderEnum(n[len("ClusterNetworkPolicyProviderEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterDatabaseEncryptionStateEnum converts a ClusterDatabaseEncryptionStateEnum enum from its proto representation.
func ProtoToContainerClusterDatabaseEncryptionStateEnum(e containerpb.ContainerClusterDatabaseEncryptionStateEnum) *container.ClusterDatabaseEncryptionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterDatabaseEncryptionStateEnum_name[int32(e)]; ok {
		e := container.ClusterDatabaseEncryptionStateEnum(n[len("ClusterDatabaseEncryptionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterMasterAuth converts a ClusterMasterAuth resource from its proto representation.
func ProtoToContainerClusterMasterAuth(p *containerpb.ContainerClusterMasterAuth) *container.ClusterMasterAuth {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuth{
		Username:                dcl.StringOrNil(p.Username),
		Password:                dcl.StringOrNil(p.Password),
		ClientCertificateConfig: ProtoToContainerClusterMasterAuthClientCertificateConfig(p.GetClientCertificateConfig()),
		ClusterCaCertificate:    dcl.StringOrNil(p.ClusterCaCertificate),
		ClientCertificate:       dcl.StringOrNil(p.ClientCertificate),
		ClientKey:               dcl.StringOrNil(p.ClientKey),
	}
	return obj
}

// ProtoToClusterMasterAuthClientCertificateConfig converts a ClusterMasterAuthClientCertificateConfig resource from its proto representation.
func ProtoToContainerClusterMasterAuthClientCertificateConfig(p *containerpb.ContainerClusterMasterAuthClientCertificateConfig) *container.ClusterMasterAuthClientCertificateConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.Bool(p.IssueClientCertificate),
	}
	return obj
}

// ProtoToClusterAddonsConfig converts a ClusterAddonsConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfig(p *containerpb.ContainerClusterAddonsConfig) *container.ClusterAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfig{
		HttpLoadBalancing:        ProtoToContainerClusterAddonsConfigHttpLoadBalancing(p.GetHttpLoadBalancing()),
		HorizontalPodAutoscaling: ProtoToContainerClusterAddonsConfigHorizontalPodAutoscaling(p.GetHorizontalPodAutoscaling()),
		KubernetesDashboard:      ProtoToContainerClusterAddonsConfigKubernetesDashboard(p.GetKubernetesDashboard()),
		NetworkPolicyConfig:      ProtoToContainerClusterAddonsConfigNetworkPolicyConfig(p.GetNetworkPolicyConfig()),
		CloudRunConfig:           ProtoToContainerClusterAddonsConfigCloudRunConfig(p.GetCloudRunConfig()),
	}
	return obj
}

// ProtoToClusterAddonsConfigHttpLoadBalancing converts a ClusterAddonsConfigHttpLoadBalancing resource from its proto representation.
func ProtoToContainerClusterAddonsConfigHttpLoadBalancing(p *containerpb.ContainerClusterAddonsConfigHttpLoadBalancing) *container.ClusterAddonsConfigHttpLoadBalancing {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigHorizontalPodAutoscaling converts a ClusterAddonsConfigHorizontalPodAutoscaling resource from its proto representation.
func ProtoToContainerClusterAddonsConfigHorizontalPodAutoscaling(p *containerpb.ContainerClusterAddonsConfigHorizontalPodAutoscaling) *container.ClusterAddonsConfigHorizontalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigKubernetesDashboard converts a ClusterAddonsConfigKubernetesDashboard resource from its proto representation.
func ProtoToContainerClusterAddonsConfigKubernetesDashboard(p *containerpb.ContainerClusterAddonsConfigKubernetesDashboard) *container.ClusterAddonsConfigKubernetesDashboard {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigNetworkPolicyConfig converts a ClusterAddonsConfigNetworkPolicyConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigNetworkPolicyConfig(p *containerpb.ContainerClusterAddonsConfigNetworkPolicyConfig) *container.ClusterAddonsConfigNetworkPolicyConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigCloudRunConfig converts a ClusterAddonsConfigCloudRunConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigCloudRunConfig(p *containerpb.ContainerClusterAddonsConfigCloudRunConfig) *container.ClusterAddonsConfigCloudRunConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigCloudRunConfig{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterNodePools converts a ClusterNodePools resource from its proto representation.
func ProtoToContainerClusterNodePools(p *containerpb.ContainerClusterNodePools) *container.ClusterNodePools {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePools{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToClusterLegacyAbac converts a ClusterLegacyAbac resource from its proto representation.
func ProtoToContainerClusterLegacyAbac(p *containerpb.ContainerClusterLegacyAbac) *container.ClusterLegacyAbac {
	if p == nil {
		return nil
	}
	obj := &container.ClusterLegacyAbac{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNetworkPolicy converts a ClusterNetworkPolicy resource from its proto representation.
func ProtoToContainerClusterNetworkPolicy(p *containerpb.ContainerClusterNetworkPolicy) *container.ClusterNetworkPolicy {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNetworkPolicy{
		Provider: ProtoToContainerClusterNetworkPolicyProviderEnum(p.GetProvider()),
		Enabled:  dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterIPAllocationPolicy converts a ClusterIPAllocationPolicy resource from its proto representation.
func ProtoToContainerClusterIPAllocationPolicy(p *containerpb.ContainerClusterIPAllocationPolicy) *container.ClusterIPAllocationPolicy {
	if p == nil {
		return nil
	}
	obj := &container.ClusterIPAllocationPolicy{
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
func ProtoToContainerClusterMasterAuthorizedNetworksConfig(p *containerpb.ContainerClusterMasterAuthorizedNetworksConfig) *container.ClusterMasterAuthorizedNetworksConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	for _, r := range p.GetCidrBlocks() {
		obj.CidrBlocks = append(obj.CidrBlocks, *ProtoToContainerClusterMasterAuthorizedNetworksConfigCidrBlocks(r))
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfigCidrBlocks converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource from its proto representation.
func ProtoToContainerClusterMasterAuthorizedNetworksConfigCidrBlocks(p *containerpb.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks) *container.ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.StringOrNil(p.DisplayName),
		CidrBlock:   dcl.StringOrNil(p.CidrBlock),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization resource from its proto representation.
func ProtoToContainerClusterBinaryAuthorization(p *containerpb.ContainerClusterBinaryAuthorization) *container.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &container.ClusterBinaryAuthorization{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAutoscaling converts a ClusterAutoscaling resource from its proto representation.
func ProtoToContainerClusterAutoscaling(p *containerpb.ContainerClusterAutoscaling) *container.ClusterAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.Bool(p.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaults(p.GetAutoprovisioningNodePoolDefaults()),
	}
	for _, r := range p.GetResourceLimits() {
		obj.ResourceLimits = append(obj.ResourceLimits, *ProtoToContainerClusterAutoscalingResourceLimits(r))
	}
	return obj
}

// ProtoToClusterAutoscalingResourceLimits converts a ClusterAutoscalingResourceLimits resource from its proto representation.
func ProtoToContainerClusterAutoscalingResourceLimits(p *containerpb.ContainerClusterAutoscalingResourceLimits) *container.ClusterAutoscalingResourceLimits {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingResourceLimits{
		ResourceType: dcl.StringOrNil(p.ResourceType),
		Minimum:      dcl.Int64OrNil(p.Minimum),
		Maximum:      dcl.Int64OrNil(p.Maximum),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaults converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaults(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:  dcl.StringOrNil(p.ServiceAccount),
		UpgradeSettings: ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p.GetUpgradeSettings()),
		Management:      ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p.GetManagement()),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade: dcl.Bool(p.AutoUpgrade),
		AutoRepair:  dcl.Bool(p.AutoRepair),
	}
	return obj
}

// ProtoToClusterNetworkConfig converts a ClusterNetworkConfig resource from its proto representation.
func ProtoToContainerClusterNetworkConfig(p *containerpb.ContainerClusterNetworkConfig) *container.ClusterNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNetworkConfig{
		Network:                   dcl.StringOrNil(p.Network),
		Subnetwork:                dcl.StringOrNil(p.Subnetwork),
		EnableIntraNodeVisibility: dcl.Bool(p.EnableIntraNodeVisibility),
	}
	return obj
}

// ProtoToClusterMaintenancePolicy converts a ClusterMaintenancePolicy resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicy(p *containerpb.ContainerClusterMaintenancePolicy) *container.ClusterMaintenancePolicy {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicy{
		Window:          ProtoToContainerClusterMaintenancePolicyWindow(p.GetWindow()),
		ResourceVersion: dcl.StringOrNil(p.ResourceVersion),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindow converts a ClusterMaintenancePolicyWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindow(p *containerpb.ContainerClusterMaintenancePolicyWindow) *container.ClusterMaintenancePolicyWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ProtoToContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow(p.GetDailyMaintenanceWindow()),
		RecurringWindow:        ProtoToContainerClusterMaintenancePolicyWindowRecurringWindow(p.GetRecurringWindow()),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowDailyMaintenanceWindow converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow(p *containerpb.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow) *container.ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		Duration:  dcl.StringOrNil(p.Duration),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindow converts a ClusterMaintenancePolicyWindowRecurringWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindowRecurringWindow(p *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindow) *container.ClusterMaintenancePolicyWindowRecurringWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ProtoToContainerClusterMaintenancePolicyWindowRecurringWindowWindow(p.GetWindow()),
		Recurrence: dcl.StringOrNil(p.Recurrence),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindowWindow converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindowRecurringWindowWindow(p *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow) *container.ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		EndTime:   dcl.StringOrNil(p.GetEndTime()),
	}
	return obj
}

// ProtoToClusterDefaultMaxPodsConstraint converts a ClusterDefaultMaxPodsConstraint resource from its proto representation.
func ProtoToContainerClusterDefaultMaxPodsConstraint(p *containerpb.ContainerClusterDefaultMaxPodsConstraint) *container.ClusterDefaultMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &container.ClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.StringOrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfig converts a ClusterResourceUsageExportConfig resource from its proto representation.
func ProtoToContainerClusterResourceUsageExportConfig(p *containerpb.ContainerClusterResourceUsageExportConfig) *container.ClusterResourceUsageExportConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterResourceUsageExportConfig{
		BigqueryDestination:           ProtoToContainerClusterResourceUsageExportConfigBigqueryDestination(p.GetBigqueryDestination()),
		EnableNetworkEgressMonitoring: dcl.Bool(p.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ProtoToContainerClusterResourceUsageExportConfigConsumptionMeteringConfig(p.GetConsumptionMeteringConfig()),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigBigqueryDestination converts a ClusterResourceUsageExportConfigBigqueryDestination resource from its proto representation.
func ProtoToContainerClusterResourceUsageExportConfigBigqueryDestination(p *containerpb.ContainerClusterResourceUsageExportConfigBigqueryDestination) *container.ClusterResourceUsageExportConfigBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &container.ClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.StringOrNil(p.DatasetId),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigConsumptionMeteringConfig converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource from its proto representation.
func ProtoToContainerClusterResourceUsageExportConfigConsumptionMeteringConfig(p *containerpb.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig) *container.ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAuthenticatorGroupsConfig converts a ClusterAuthenticatorGroupsConfig resource from its proto representation.
func ProtoToContainerClusterAuthenticatorGroupsConfig(p *containerpb.ContainerClusterAuthenticatorGroupsConfig) *container.ClusterAuthenticatorGroupsConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.Bool(p.Enabled),
		SecurityGroup: dcl.StringOrNil(p.SecurityGroup),
	}
	return obj
}

// ProtoToClusterPrivateClusterConfig converts a ClusterPrivateClusterConfig resource from its proto representation.
func ProtoToContainerClusterPrivateClusterConfig(p *containerpb.ContainerClusterPrivateClusterConfig) *container.ClusterPrivateClusterConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterPrivateClusterConfig{
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
func ProtoToContainerClusterDatabaseEncryption(p *containerpb.ContainerClusterDatabaseEncryption) *container.ClusterDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &container.ClusterDatabaseEncryption{
		State:   ProtoToContainerClusterDatabaseEncryptionStateEnum(p.GetState()),
		KeyName: dcl.StringOrNil(p.KeyName),
	}
	return obj
}

// ProtoToClusterVerticalPodAutoscaling converts a ClusterVerticalPodAutoscaling resource from its proto representation.
func ProtoToContainerClusterVerticalPodAutoscaling(p *containerpb.ContainerClusterVerticalPodAutoscaling) *container.ClusterVerticalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterVerticalPodAutoscaling{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterShieldedNodes converts a ClusterShieldedNodes resource from its proto representation.
func ProtoToContainerClusterShieldedNodes(p *containerpb.ContainerClusterShieldedNodes) *container.ClusterShieldedNodes {
	if p == nil {
		return nil
	}
	obj := &container.ClusterShieldedNodes{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterConditions converts a ClusterConditions resource from its proto representation.
func ProtoToContainerClusterConditions(p *containerpb.ContainerClusterConditions) *container.ClusterConditions {
	if p == nil {
		return nil
	}
	obj := &container.ClusterConditions{
		Code:    dcl.StringOrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *containerpb.ContainerCluster) *container.Cluster {
	obj := &container.Cluster{
		Name:                           dcl.StringOrNil(p.Name),
		Description:                    dcl.StringOrNil(p.Description),
		InitialNodeCount:               dcl.Int64OrNil(p.InitialNodeCount),
		MasterAuth:                     ProtoToContainerClusterMasterAuth(p.GetMasterAuth()),
		LoggingService:                 dcl.StringOrNil(p.LoggingService),
		MonitoringService:              dcl.StringOrNil(p.MonitoringService),
		Network:                        dcl.StringOrNil(p.Network),
		ClusterIPv4Cidr:                dcl.StringOrNil(p.ClusterIpv4Cidr),
		AddonsConfig:                   ProtoToContainerClusterAddonsConfig(p.GetAddonsConfig()),
		Subnetwork:                     dcl.StringOrNil(p.Subnetwork),
		EnableKubernetesAlpha:          dcl.Bool(p.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.StringOrNil(p.LabelFingerprint),
		LegacyAbac:                     ProtoToContainerClusterLegacyAbac(p.GetLegacyAbac()),
		NetworkPolicy:                  ProtoToContainerClusterNetworkPolicy(p.GetNetworkPolicy()),
		IPAllocationPolicy:             ProtoToContainerClusterIPAllocationPolicy(p.GetIpAllocationPolicy()),
		MasterAuthorizedNetworksConfig: ProtoToContainerClusterMasterAuthorizedNetworksConfig(p.GetMasterAuthorizedNetworksConfig()),
		BinaryAuthorization:            ProtoToContainerClusterBinaryAuthorization(p.GetBinaryAuthorization()),
		Autoscaling:                    ProtoToContainerClusterAutoscaling(p.GetAutoscaling()),
		NetworkConfig:                  ProtoToContainerClusterNetworkConfig(p.GetNetworkConfig()),
		MaintenancePolicy:              ProtoToContainerClusterMaintenancePolicy(p.GetMaintenancePolicy()),
		DefaultMaxPodsConstraint:       ProtoToContainerClusterDefaultMaxPodsConstraint(p.GetDefaultMaxPodsConstraint()),
		ResourceUsageExportConfig:      ProtoToContainerClusterResourceUsageExportConfig(p.GetResourceUsageExportConfig()),
		AuthenticatorGroupsConfig:      ProtoToContainerClusterAuthenticatorGroupsConfig(p.GetAuthenticatorGroupsConfig()),
		PrivateClusterConfig:           ProtoToContainerClusterPrivateClusterConfig(p.GetPrivateClusterConfig()),
		DatabaseEncryption:             ProtoToContainerClusterDatabaseEncryption(p.GetDatabaseEncryption()),
		VerticalPodAutoscaling:         ProtoToContainerClusterVerticalPodAutoscaling(p.GetVerticalPodAutoscaling()),
		ShieldedNodes:                  ProtoToContainerClusterShieldedNodes(p.GetShieldedNodes()),
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
		obj.NodePools = append(obj.NodePools, *ProtoToContainerClusterNodePools(r))
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerClusterConditions(r))
	}
	return obj
}

// ClusterNetworkPolicyProviderEnumToProto converts a ClusterNetworkPolicyProviderEnum enum to its proto representation.
func ContainerClusterNetworkPolicyProviderEnumToProto(e *container.ClusterNetworkPolicyProviderEnum) containerpb.ContainerClusterNetworkPolicyProviderEnum {
	if e == nil {
		return containerpb.ContainerClusterNetworkPolicyProviderEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNetworkPolicyProviderEnum_value["ClusterNetworkPolicyProviderEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNetworkPolicyProviderEnum(v)
	}
	return containerpb.ContainerClusterNetworkPolicyProviderEnum(0)
}

// ClusterDatabaseEncryptionStateEnumToProto converts a ClusterDatabaseEncryptionStateEnum enum to its proto representation.
func ContainerClusterDatabaseEncryptionStateEnumToProto(e *container.ClusterDatabaseEncryptionStateEnum) containerpb.ContainerClusterDatabaseEncryptionStateEnum {
	if e == nil {
		return containerpb.ContainerClusterDatabaseEncryptionStateEnum(0)
	}
	if v, ok := containerpb.ContainerClusterDatabaseEncryptionStateEnum_value["ClusterDatabaseEncryptionStateEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterDatabaseEncryptionStateEnum(v)
	}
	return containerpb.ContainerClusterDatabaseEncryptionStateEnum(0)
}

// ClusterMasterAuthToProto converts a ClusterMasterAuth resource to its proto representation.
func ContainerClusterMasterAuthToProto(o *container.ClusterMasterAuth) *containerpb.ContainerClusterMasterAuth {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuth{
		Username:                dcl.ValueOrEmptyString(o.Username),
		Password:                dcl.ValueOrEmptyString(o.Password),
		ClientCertificateConfig: ContainerClusterMasterAuthClientCertificateConfigToProto(o.ClientCertificateConfig),
		ClusterCaCertificate:    dcl.ValueOrEmptyString(o.ClusterCaCertificate),
		ClientCertificate:       dcl.ValueOrEmptyString(o.ClientCertificate),
		ClientKey:               dcl.ValueOrEmptyString(o.ClientKey),
	}
	return p
}

// ClusterMasterAuthClientCertificateConfigToProto converts a ClusterMasterAuthClientCertificateConfig resource to its proto representation.
func ContainerClusterMasterAuthClientCertificateConfigToProto(o *container.ClusterMasterAuthClientCertificateConfig) *containerpb.ContainerClusterMasterAuthClientCertificateConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.ValueOrEmptyBool(o.IssueClientCertificate),
	}
	return p
}

// ClusterAddonsConfigToProto converts a ClusterAddonsConfig resource to its proto representation.
func ContainerClusterAddonsConfigToProto(o *container.ClusterAddonsConfig) *containerpb.ContainerClusterAddonsConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfig{
		HttpLoadBalancing:        ContainerClusterAddonsConfigHttpLoadBalancingToProto(o.HttpLoadBalancing),
		HorizontalPodAutoscaling: ContainerClusterAddonsConfigHorizontalPodAutoscalingToProto(o.HorizontalPodAutoscaling),
		KubernetesDashboard:      ContainerClusterAddonsConfigKubernetesDashboardToProto(o.KubernetesDashboard),
		NetworkPolicyConfig:      ContainerClusterAddonsConfigNetworkPolicyConfigToProto(o.NetworkPolicyConfig),
		CloudRunConfig:           ContainerClusterAddonsConfigCloudRunConfigToProto(o.CloudRunConfig),
	}
	return p
}

// ClusterAddonsConfigHttpLoadBalancingToProto converts a ClusterAddonsConfigHttpLoadBalancing resource to its proto representation.
func ContainerClusterAddonsConfigHttpLoadBalancingToProto(o *container.ClusterAddonsConfigHttpLoadBalancing) *containerpb.ContainerClusterAddonsConfigHttpLoadBalancing {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigHorizontalPodAutoscalingToProto converts a ClusterAddonsConfigHorizontalPodAutoscaling resource to its proto representation.
func ContainerClusterAddonsConfigHorizontalPodAutoscalingToProto(o *container.ClusterAddonsConfigHorizontalPodAutoscaling) *containerpb.ContainerClusterAddonsConfigHorizontalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigKubernetesDashboardToProto converts a ClusterAddonsConfigKubernetesDashboard resource to its proto representation.
func ContainerClusterAddonsConfigKubernetesDashboardToProto(o *container.ClusterAddonsConfigKubernetesDashboard) *containerpb.ContainerClusterAddonsConfigKubernetesDashboard {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigNetworkPolicyConfigToProto converts a ClusterAddonsConfigNetworkPolicyConfig resource to its proto representation.
func ContainerClusterAddonsConfigNetworkPolicyConfigToProto(o *container.ClusterAddonsConfigNetworkPolicyConfig) *containerpb.ContainerClusterAddonsConfigNetworkPolicyConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigCloudRunConfigToProto converts a ClusterAddonsConfigCloudRunConfig resource to its proto representation.
func ContainerClusterAddonsConfigCloudRunConfigToProto(o *container.ClusterAddonsConfigCloudRunConfig) *containerpb.ContainerClusterAddonsConfigCloudRunConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigCloudRunConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterNodePoolsToProto converts a ClusterNodePools resource to its proto representation.
func ContainerClusterNodePoolsToProto(o *container.ClusterNodePools) *containerpb.ContainerClusterNodePools {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePools{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// ClusterLegacyAbacToProto converts a ClusterLegacyAbac resource to its proto representation.
func ContainerClusterLegacyAbacToProto(o *container.ClusterLegacyAbac) *containerpb.ContainerClusterLegacyAbac {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterLegacyAbac{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNetworkPolicyToProto converts a ClusterNetworkPolicy resource to its proto representation.
func ContainerClusterNetworkPolicyToProto(o *container.ClusterNetworkPolicy) *containerpb.ContainerClusterNetworkPolicy {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNetworkPolicy{
		Provider: ContainerClusterNetworkPolicyProviderEnumToProto(o.Provider),
		Enabled:  dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterIPAllocationPolicyToProto converts a ClusterIPAllocationPolicy resource to its proto representation.
func ContainerClusterIPAllocationPolicyToProto(o *container.ClusterIPAllocationPolicy) *containerpb.ContainerClusterIPAllocationPolicy {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterIPAllocationPolicy{
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
func ContainerClusterMasterAuthorizedNetworksConfigToProto(o *container.ClusterMasterAuthorizedNetworksConfig) *containerpb.ContainerClusterMasterAuthorizedNetworksConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	for _, r := range o.CidrBlocks {
		p.CidrBlocks = append(p.CidrBlocks, ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(&r))
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigCidrBlocksToProto converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource to its proto representation.
func ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(o *container.ClusterMasterAuthorizedNetworksConfigCidrBlocks) *containerpb.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		CidrBlock:   dcl.ValueOrEmptyString(o.CidrBlock),
	}
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization resource to its proto representation.
func ContainerClusterBinaryAuthorizationToProto(o *container.ClusterBinaryAuthorization) *containerpb.ContainerClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterBinaryAuthorization{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAutoscalingToProto converts a ClusterAutoscaling resource to its proto representation.
func ContainerClusterAutoscalingToProto(o *container.ClusterAutoscaling) *containerpb.ContainerClusterAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.ValueOrEmptyBool(o.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o.AutoprovisioningNodePoolDefaults),
	}
	for _, r := range o.ResourceLimits {
		p.ResourceLimits = append(p.ResourceLimits, ContainerClusterAutoscalingResourceLimitsToProto(&r))
	}
	return p
}

// ClusterAutoscalingResourceLimitsToProto converts a ClusterAutoscalingResourceLimits resource to its proto representation.
func ContainerClusterAutoscalingResourceLimitsToProto(o *container.ClusterAutoscalingResourceLimits) *containerpb.ContainerClusterAutoscalingResourceLimits {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingResourceLimits{
		ResourceType: dcl.ValueOrEmptyString(o.ResourceType),
		Minimum:      dcl.ValueOrEmptyInt64(o.Minimum),
		Maximum:      dcl.ValueOrEmptyInt64(o.Maximum),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaults) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:  dcl.ValueOrEmptyString(o.ServiceAccount),
		UpgradeSettings: ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o.UpgradeSettings),
		Management:      ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o.Management),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade: dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:  dcl.ValueOrEmptyBool(o.AutoRepair),
	}
	return p
}

// ClusterNetworkConfigToProto converts a ClusterNetworkConfig resource to its proto representation.
func ContainerClusterNetworkConfigToProto(o *container.ClusterNetworkConfig) *containerpb.ContainerClusterNetworkConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNetworkConfig{
		Network:                   dcl.ValueOrEmptyString(o.Network),
		Subnetwork:                dcl.ValueOrEmptyString(o.Subnetwork),
		EnableIntraNodeVisibility: dcl.ValueOrEmptyBool(o.EnableIntraNodeVisibility),
	}
	return p
}

// ClusterMaintenancePolicyToProto converts a ClusterMaintenancePolicy resource to its proto representation.
func ContainerClusterMaintenancePolicyToProto(o *container.ClusterMaintenancePolicy) *containerpb.ContainerClusterMaintenancePolicy {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicy{
		Window:          ContainerClusterMaintenancePolicyWindowToProto(o.Window),
		ResourceVersion: dcl.ValueOrEmptyString(o.ResourceVersion),
	}
	return p
}

// ClusterMaintenancePolicyWindowToProto converts a ClusterMaintenancePolicyWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowToProto(o *container.ClusterMaintenancePolicyWindow) *containerpb.ContainerClusterMaintenancePolicyWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o.DailyMaintenanceWindow),
		RecurringWindow:        ContainerClusterMaintenancePolicyWindowRecurringWindowToProto(o.RecurringWindow),
	}
	return p
}

// ClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o *container.ClusterMaintenancePolicyWindowDailyMaintenanceWindow) *containerpb.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		Duration:  dcl.ValueOrEmptyString(o.Duration),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowRecurringWindowToProto(o *container.ClusterMaintenancePolicyWindowRecurringWindow) *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ContainerClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o.Window),
		Recurrence: dcl.ValueOrEmptyString(o.Recurrence),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o *container.ClusterMaintenancePolicyWindowRecurringWindowWindow) *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// ClusterDefaultMaxPodsConstraintToProto converts a ClusterDefaultMaxPodsConstraint resource to its proto representation.
func ContainerClusterDefaultMaxPodsConstraintToProto(o *container.ClusterDefaultMaxPodsConstraint) *containerpb.ContainerClusterDefaultMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyString(o.MaxPodsPerNode),
	}
	return p
}

// ClusterResourceUsageExportConfigToProto converts a ClusterResourceUsageExportConfig resource to its proto representation.
func ContainerClusterResourceUsageExportConfigToProto(o *container.ClusterResourceUsageExportConfig) *containerpb.ContainerClusterResourceUsageExportConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterResourceUsageExportConfig{
		BigqueryDestination:           ContainerClusterResourceUsageExportConfigBigqueryDestinationToProto(o.BigqueryDestination),
		EnableNetworkEgressMonitoring: dcl.ValueOrEmptyBool(o.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ContainerClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o.ConsumptionMeteringConfig),
	}
	return p
}

// ClusterResourceUsageExportConfigBigqueryDestinationToProto converts a ClusterResourceUsageExportConfigBigqueryDestination resource to its proto representation.
func ContainerClusterResourceUsageExportConfigBigqueryDestinationToProto(o *container.ClusterResourceUsageExportConfigBigqueryDestination) *containerpb.ContainerClusterResourceUsageExportConfigBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.ValueOrEmptyString(o.DatasetId),
	}
	return p
}

// ClusterResourceUsageExportConfigConsumptionMeteringConfigToProto converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource to its proto representation.
func ContainerClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o *container.ClusterResourceUsageExportConfigConsumptionMeteringConfig) *containerpb.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAuthenticatorGroupsConfigToProto converts a ClusterAuthenticatorGroupsConfig resource to its proto representation.
func ContainerClusterAuthenticatorGroupsConfigToProto(o *container.ClusterAuthenticatorGroupsConfig) *containerpb.ContainerClusterAuthenticatorGroupsConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.ValueOrEmptyBool(o.Enabled),
		SecurityGroup: dcl.ValueOrEmptyString(o.SecurityGroup),
	}
	return p
}

// ClusterPrivateClusterConfigToProto converts a ClusterPrivateClusterConfig resource to its proto representation.
func ContainerClusterPrivateClusterConfigToProto(o *container.ClusterPrivateClusterConfig) *containerpb.ContainerClusterPrivateClusterConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterPrivateClusterConfig{
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
func ContainerClusterDatabaseEncryptionToProto(o *container.ClusterDatabaseEncryption) *containerpb.ContainerClusterDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterDatabaseEncryption{
		State:   ContainerClusterDatabaseEncryptionStateEnumToProto(o.State),
		KeyName: dcl.ValueOrEmptyString(o.KeyName),
	}
	return p
}

// ClusterVerticalPodAutoscalingToProto converts a ClusterVerticalPodAutoscaling resource to its proto representation.
func ContainerClusterVerticalPodAutoscalingToProto(o *container.ClusterVerticalPodAutoscaling) *containerpb.ContainerClusterVerticalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterVerticalPodAutoscaling{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterShieldedNodesToProto converts a ClusterShieldedNodes resource to its proto representation.
func ContainerClusterShieldedNodesToProto(o *container.ClusterShieldedNodes) *containerpb.ContainerClusterShieldedNodes {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterShieldedNodes{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterConditionsToProto converts a ClusterConditions resource to its proto representation.
func ContainerClusterConditionsToProto(o *container.ClusterConditions) *containerpb.ContainerClusterConditions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterConditions{
		Code:    dcl.ValueOrEmptyString(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *container.Cluster) *containerpb.ContainerCluster {
	p := &containerpb.ContainerCluster{
		Name:                           dcl.ValueOrEmptyString(resource.Name),
		Description:                    dcl.ValueOrEmptyString(resource.Description),
		InitialNodeCount:               dcl.ValueOrEmptyInt64(resource.InitialNodeCount),
		MasterAuth:                     ContainerClusterMasterAuthToProto(resource.MasterAuth),
		LoggingService:                 dcl.ValueOrEmptyString(resource.LoggingService),
		MonitoringService:              dcl.ValueOrEmptyString(resource.MonitoringService),
		Network:                        dcl.ValueOrEmptyString(resource.Network),
		ClusterIpv4Cidr:                dcl.ValueOrEmptyString(resource.ClusterIPv4Cidr),
		AddonsConfig:                   ContainerClusterAddonsConfigToProto(resource.AddonsConfig),
		Subnetwork:                     dcl.ValueOrEmptyString(resource.Subnetwork),
		EnableKubernetesAlpha:          dcl.ValueOrEmptyBool(resource.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.ValueOrEmptyString(resource.LabelFingerprint),
		LegacyAbac:                     ContainerClusterLegacyAbacToProto(resource.LegacyAbac),
		NetworkPolicy:                  ContainerClusterNetworkPolicyToProto(resource.NetworkPolicy),
		IpAllocationPolicy:             ContainerClusterIPAllocationPolicyToProto(resource.IPAllocationPolicy),
		MasterAuthorizedNetworksConfig: ContainerClusterMasterAuthorizedNetworksConfigToProto(resource.MasterAuthorizedNetworksConfig),
		BinaryAuthorization:            ContainerClusterBinaryAuthorizationToProto(resource.BinaryAuthorization),
		Autoscaling:                    ContainerClusterAutoscalingToProto(resource.Autoscaling),
		NetworkConfig:                  ContainerClusterNetworkConfigToProto(resource.NetworkConfig),
		MaintenancePolicy:              ContainerClusterMaintenancePolicyToProto(resource.MaintenancePolicy),
		DefaultMaxPodsConstraint:       ContainerClusterDefaultMaxPodsConstraintToProto(resource.DefaultMaxPodsConstraint),
		ResourceUsageExportConfig:      ContainerClusterResourceUsageExportConfigToProto(resource.ResourceUsageExportConfig),
		AuthenticatorGroupsConfig:      ContainerClusterAuthenticatorGroupsConfigToProto(resource.AuthenticatorGroupsConfig),
		PrivateClusterConfig:           ContainerClusterPrivateClusterConfigToProto(resource.PrivateClusterConfig),
		DatabaseEncryption:             ContainerClusterDatabaseEncryptionToProto(resource.DatabaseEncryption),
		VerticalPodAutoscaling:         ContainerClusterVerticalPodAutoscalingToProto(resource.VerticalPodAutoscaling),
		ShieldedNodes:                  ContainerClusterShieldedNodesToProto(resource.ShieldedNodes),
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
		p.NodePools = append(p.NodePools, ContainerClusterNodePoolsToProto(&r))
	}
	for _, r := range resource.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, ContainerClusterConditionsToProto(&r))
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *container.Client, request *containerpb.ApplyContainerClusterRequest) (*containerpb.ContainerCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerCluster(ctx context.Context, request *containerpb.ApplyContainerClusterRequest) (*containerpb.ContainerCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerCluster(ctx context.Context, request *containerpb.DeleteContainerClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerCluster(ctx context.Context, request *containerpb.ListContainerClusterRequest) (*containerpb.ListContainerClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*containerpb.ContainerCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &containerpb.ListContainerClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*container.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return container.NewClient(conf), nil
}
