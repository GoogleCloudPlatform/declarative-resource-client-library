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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/beta/dataproc_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *beta.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := beta.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *beta.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterInstanceGroupConfigPreemptibilityEnum converts a ClusterInstanceGroupConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigPreemptibilityEnum(e betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum) *beta.ClusterInstanceGroupConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.ClusterInstanceGroupConfigPreemptibilityEnum(n[len("DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfigSoftwareConfigOptionalComponentsEnum converts a ClusterClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(e betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum) *beta.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := beta.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusStateEnum converts a ClusterStatusStateEnum enum from its proto representation.
func ProtoToDataprocBetaClusterStatusStateEnum(e betapb.DataprocBetaClusterStatusStateEnum) *beta.ClusterStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterStatusStateEnum_name[int32(e)]; ok {
		e := beta.ClusterStatusStateEnum(n[len("DataprocBetaClusterStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusSubstateEnum converts a ClusterStatusSubstateEnum enum from its proto representation.
func ProtoToDataprocBetaClusterStatusSubstateEnum(e betapb.DataprocBetaClusterStatusSubstateEnum) *beta.ClusterStatusSubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterStatusSubstateEnum_name[int32(e)]; ok {
		e := beta.ClusterStatusSubstateEnum(n[len("DataprocBetaClusterStatusSubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistoryStateEnum converts a ClusterStatusHistoryStateEnum enum from its proto representation.
func ProtoToDataprocBetaClusterStatusHistoryStateEnum(e betapb.DataprocBetaClusterStatusHistoryStateEnum) *beta.ClusterStatusHistoryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterStatusHistoryStateEnum_name[int32(e)]; ok {
		e := beta.ClusterStatusHistoryStateEnum(n[len("DataprocBetaClusterStatusHistoryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistorySubstateEnum converts a ClusterStatusHistorySubstateEnum enum from its proto representation.
func ProtoToDataprocBetaClusterStatusHistorySubstateEnum(e betapb.DataprocBetaClusterStatusHistorySubstateEnum) *beta.ClusterStatusHistorySubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterStatusHistorySubstateEnum_name[int32(e)]; ok {
		e := beta.ClusterStatusHistorySubstateEnum(n[len("DataprocBetaClusterStatusHistorySubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfig converts a ClusterClusterConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfig(p *betapb.DataprocBetaClusterClusterConfig) *beta.ClusterClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.StagingBucket),
		TempBucket:            dcl.StringOrNil(p.TempBucket),
		GceClusterConfig:      ProtoToDataprocBetaClusterClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocBetaClusterInstanceGroupConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocBetaClusterInstanceGroupConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocBetaClusterInstanceGroupConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocBetaClusterClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocBetaClusterClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocBetaClusterClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocBetaClusterClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocBetaClusterClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocBetaClusterClusterConfigEndpointConfig(p.GetEndpointConfig()),
		GkeClusterConfig:      ProtoToDataprocBetaClusterClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		MetastoreConfig:       ProtoToDataprocBetaClusterClusterConfigMetastoreConfig(p.GetMetastoreConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocBetaClusterClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfig converts a ClusterClusterConfigGceClusterConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfig(p *betapb.DataprocBetaClusterClusterConfigGceClusterConfig) *beta.ClusterClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.Zone),
		Network:                 dcl.StringOrNil(p.Network),
		Subnetwork:              dcl.StringOrNil(p.Subnetwork),
		InternalIPOnly:          dcl.Bool(p.InternalIpOnly),
		PrivateIPv6GoogleAccess: ProtoToDataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.ServiceAccount),
		ReservationAffinity:     ProtoToDataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
	}
	for _, r := range p.GetServiceAccountScopes() {
		obj.ServiceAccountScopes = append(obj.ServiceAccountScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinity converts a ClusterClusterConfigGceClusterConfigReservationAffinity resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity(p *betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity) *beta.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity(p *betapb.DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity) *beta.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.NodeGroup),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfig converts a ClusterInstanceGroupConfig resource from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfig(p *betapb.DataprocBetaClusterInstanceGroupConfig) *beta.ClusterInstanceGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfig{
		NumInstances:       dcl.Int64OrNil(p.NumInstances),
		Image:              dcl.StringOrNil(p.Image),
		MachineType:        dcl.StringOrNil(p.MachineType),
		DiskConfig:         ProtoToDataprocBetaClusterInstanceGroupConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.IsPreemptible),
		Preemptibility:     ProtoToDataprocBetaClusterInstanceGroupConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaClusterInstanceGroupConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.MinCpuPlatform),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaClusterInstanceGroupConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigDiskConfig converts a ClusterInstanceGroupConfigDiskConfig resource from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigDiskConfig(p *betapb.DataprocBetaClusterInstanceGroupConfigDiskConfig) *beta.ClusterInstanceGroupConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.BootDiskType),
		BootDiskSizeGb: dcl.Int64OrNil(p.BootDiskSizeGb),
		NumLocalSsds:   dcl.Int64OrNil(p.NumLocalSsds),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigManagedGroupConfig converts a ClusterInstanceGroupConfigManagedGroupConfig resource from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigManagedGroupConfig(p *betapb.DataprocBetaClusterInstanceGroupConfigManagedGroupConfig) *beta.ClusterInstanceGroupConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.InstanceTemplateName),
		InstanceGroupManagerName: dcl.StringOrNil(p.InstanceGroupManagerName),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigAccelerators converts a ClusterInstanceGroupConfigAccelerators resource from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigAccelerators(p *betapb.DataprocBetaClusterInstanceGroupConfigAccelerators) *beta.ClusterInstanceGroupConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
	}
	return obj
}

// ProtoToClusterClusterConfigSoftwareConfig converts a ClusterClusterConfigSoftwareConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSoftwareConfig(p *betapb.DataprocBetaClusterClusterConfigSoftwareConfig) *beta.ClusterClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.ImageVersion),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterClusterConfigInitializationActions converts a ClusterClusterConfigInitializationActions resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigInitializationActions(p *betapb.DataprocBetaClusterClusterConfigInitializationActions) *beta.ClusterClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.ExecutableFile),
		ExecutionTimeout: dcl.StringOrNil(p.ExecutionTimeout),
	}
	return obj
}

// ProtoToClusterClusterConfigEncryptionConfig converts a ClusterClusterConfigEncryptionConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigEncryptionConfig(p *betapb.DataprocBetaClusterClusterConfigEncryptionConfig) *beta.ClusterClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GcePdKmsKeyName),
	}
	return obj
}

// ProtoToClusterClusterConfigAutoscalingConfig converts a ClusterClusterConfigAutoscalingConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigAutoscalingConfig(p *betapb.DataprocBetaClusterClusterConfigAutoscalingConfig) *beta.ClusterClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.Policy),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfig converts a ClusterClusterConfigSecurityConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSecurityConfig(p *betapb.DataprocBetaClusterClusterConfigSecurityConfig) *beta.ClusterClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocBetaClusterClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfigKerberosConfig converts a ClusterClusterConfigSecurityConfigKerberosConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSecurityConfigKerberosConfig(p *betapb.DataprocBetaClusterClusterConfigSecurityConfigKerberosConfig) *beta.ClusterClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigSecurityConfigKerberosConfig{
		EnableKerberos:                dcl.Bool(p.EnableKerberos),
		RootPrincipalPassword:         dcl.StringOrNil(p.RootPrincipalPassword),
		KmsKey:                        dcl.StringOrNil(p.KmsKey),
		Keystore:                      dcl.StringOrNil(p.Keystore),
		Truststore:                    dcl.StringOrNil(p.Truststore),
		KeystorePassword:              dcl.StringOrNil(p.KeystorePassword),
		KeyPassword:                   dcl.StringOrNil(p.KeyPassword),
		TruststorePassword:            dcl.StringOrNil(p.TruststorePassword),
		CrossRealmTrustRealm:          dcl.StringOrNil(p.CrossRealmTrustRealm),
		CrossRealmTrustKdc:            dcl.StringOrNil(p.CrossRealmTrustKdc),
		CrossRealmTrustAdminServer:    dcl.StringOrNil(p.CrossRealmTrustAdminServer),
		CrossRealmTrustSharedPassword: dcl.StringOrNil(p.CrossRealmTrustSharedPassword),
		KdcDbKey:                      dcl.StringOrNil(p.KdcDbKey),
		TgtLifetimeHours:              dcl.Int64OrNil(p.TgtLifetimeHours),
		Realm:                         dcl.StringOrNil(p.Realm),
	}
	return obj
}

// ProtoToClusterClusterConfigLifecycleConfig converts a ClusterClusterConfigLifecycleConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigLifecycleConfig(p *betapb.DataprocBetaClusterClusterConfigLifecycleConfig) *beta.ClusterClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.IdleDeleteTtl),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.AutoDeleteTtl),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterClusterConfigEndpointConfig converts a ClusterClusterConfigEndpointConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigEndpointConfig(p *betapb.DataprocBetaClusterClusterConfigEndpointConfig) *beta.ClusterClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.EnableHttpPortAccess),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfig converts a ClusterClusterConfigGkeClusterConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGkeClusterConfig(p *betapb.DataprocBetaClusterClusterConfigGkeClusterConfig) *beta.ClusterClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *betapb.DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *beta.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.TargetGkeCluster),
		ClusterNamespace: dcl.StringOrNil(p.ClusterNamespace),
	}
	return obj
}

// ProtoToClusterClusterConfigMetastoreConfig converts a ClusterClusterConfigMetastoreConfig resource from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigMetastoreConfig(p *betapb.DataprocBetaClusterClusterConfigMetastoreConfig) *beta.ClusterClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.DataprocMetastoreService),
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus resource from its proto representation.
func ProtoToDataprocBetaClusterStatus(p *betapb.DataprocBetaClusterStatus) *beta.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterStatus{
		State:          ProtoToDataprocBetaClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.Detail),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocBetaClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory resource from its proto representation.
func ProtoToDataprocBetaClusterStatusHistory(p *betapb.DataprocBetaClusterStatusHistory) *beta.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterStatusHistory{
		State:          ProtoToDataprocBetaClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.Detail),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocBetaClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics resource from its proto representation.
func ProtoToDataprocBetaClusterMetrics(p *betapb.DataprocBetaClusterMetrics) *beta.ClusterMetrics {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMetrics{}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *betapb.DataprocBetaCluster) *beta.Cluster {
	obj := &beta.Cluster{
		Project:     dcl.StringOrNil(p.Project),
		Name:        dcl.StringOrNil(p.Name),
		Config:      ProtoToDataprocBetaClusterClusterConfig(p.GetConfig()),
		Status:      ProtoToDataprocBetaClusterStatus(p.GetStatus()),
		ClusterUuid: dcl.StringOrNil(p.ClusterUuid),
		Metrics:     ProtoToDataprocBetaClusterMetrics(p.GetMetrics()),
		Location:    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetStatusHistory() {
		obj.StatusHistory = append(obj.StatusHistory, *ProtoToDataprocBetaClusterStatusHistory(r))
	}
	return obj
}

// ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *beta.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return betapb.DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *beta.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterInstanceGroupConfigPreemptibilityEnumToProto converts a ClusterInstanceGroupConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaClusterInstanceGroupConfigPreemptibilityEnumToProto(e *beta.ClusterInstanceGroupConfigPreemptibilityEnum) betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum_value["ClusterInstanceGroupConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaClusterInstanceGroupConfigPreemptibilityEnum(0)
}

// ClusterClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a ClusterClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *beta.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum) betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value["ClusterClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// ClusterStatusStateEnumToProto converts a ClusterStatusStateEnum enum to its proto representation.
func DataprocBetaClusterStatusStateEnumToProto(e *beta.ClusterStatusStateEnum) betapb.DataprocBetaClusterStatusStateEnum {
	if e == nil {
		return betapb.DataprocBetaClusterStatusStateEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterStatusStateEnum_value["ClusterStatusStateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterStatusStateEnum(v)
	}
	return betapb.DataprocBetaClusterStatusStateEnum(0)
}

// ClusterStatusSubstateEnumToProto converts a ClusterStatusSubstateEnum enum to its proto representation.
func DataprocBetaClusterStatusSubstateEnumToProto(e *beta.ClusterStatusSubstateEnum) betapb.DataprocBetaClusterStatusSubstateEnum {
	if e == nil {
		return betapb.DataprocBetaClusterStatusSubstateEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterStatusSubstateEnum_value["ClusterStatusSubstateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterStatusSubstateEnum(v)
	}
	return betapb.DataprocBetaClusterStatusSubstateEnum(0)
}

// ClusterStatusHistoryStateEnumToProto converts a ClusterStatusHistoryStateEnum enum to its proto representation.
func DataprocBetaClusterStatusHistoryStateEnumToProto(e *beta.ClusterStatusHistoryStateEnum) betapb.DataprocBetaClusterStatusHistoryStateEnum {
	if e == nil {
		return betapb.DataprocBetaClusterStatusHistoryStateEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterStatusHistoryStateEnum_value["ClusterStatusHistoryStateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterStatusHistoryStateEnum(v)
	}
	return betapb.DataprocBetaClusterStatusHistoryStateEnum(0)
}

// ClusterStatusHistorySubstateEnumToProto converts a ClusterStatusHistorySubstateEnum enum to its proto representation.
func DataprocBetaClusterStatusHistorySubstateEnumToProto(e *beta.ClusterStatusHistorySubstateEnum) betapb.DataprocBetaClusterStatusHistorySubstateEnum {
	if e == nil {
		return betapb.DataprocBetaClusterStatusHistorySubstateEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterStatusHistorySubstateEnum_value["ClusterStatusHistorySubstateEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterStatusHistorySubstateEnum(v)
	}
	return betapb.DataprocBetaClusterStatusHistorySubstateEnum(0)
}

// ClusterClusterConfigToProto converts a ClusterClusterConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigToProto(o *beta.ClusterClusterConfig) *betapb.DataprocBetaClusterClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfig{
		StagingBucket:         dcl.ValueOrEmptyString(o.StagingBucket),
		TempBucket:            dcl.ValueOrEmptyString(o.TempBucket),
		GceClusterConfig:      DataprocBetaClusterClusterConfigGceClusterConfigToProto(o.GceClusterConfig),
		MasterConfig:          DataprocBetaClusterInstanceGroupConfigToProto(o.MasterConfig),
		WorkerConfig:          DataprocBetaClusterInstanceGroupConfigToProto(o.WorkerConfig),
		SecondaryWorkerConfig: DataprocBetaClusterInstanceGroupConfigToProto(o.SecondaryWorkerConfig),
		SoftwareConfig:        DataprocBetaClusterClusterConfigSoftwareConfigToProto(o.SoftwareConfig),
		EncryptionConfig:      DataprocBetaClusterClusterConfigEncryptionConfigToProto(o.EncryptionConfig),
		AutoscalingConfig:     DataprocBetaClusterClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig),
		SecurityConfig:        DataprocBetaClusterClusterConfigSecurityConfigToProto(o.SecurityConfig),
		LifecycleConfig:       DataprocBetaClusterClusterConfigLifecycleConfigToProto(o.LifecycleConfig),
		EndpointConfig:        DataprocBetaClusterClusterConfigEndpointConfigToProto(o.EndpointConfig),
		GkeClusterConfig:      DataprocBetaClusterClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig),
		MetastoreConfig:       DataprocBetaClusterClusterConfigMetastoreConfigToProto(o.MetastoreConfig),
	}
	for _, r := range o.InitializationActions {
		p.InitializationActions = append(p.InitializationActions, DataprocBetaClusterClusterConfigInitializationActionsToProto(&r))
	}
	return p
}

// ClusterClusterConfigGceClusterConfigToProto converts a ClusterClusterConfigGceClusterConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigToProto(o *beta.ClusterClusterConfigGceClusterConfig) *betapb.DataprocBetaClusterClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.ValueOrEmptyString(o.Zone),
		Network:                 dcl.ValueOrEmptyString(o.Network),
		Subnetwork:              dcl.ValueOrEmptyString(o.Subnetwork),
		InternalIpOnly:          dcl.ValueOrEmptyBool(o.InternalIPOnly),
		PrivateIpv6GoogleAccess: DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess),
		ServiceAccount:          dcl.ValueOrEmptyString(o.ServiceAccount),
		ReservationAffinity:     DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity),
		NodeGroupAffinity:       DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity),
	}
	for _, r := range o.ServiceAccountScopes {
		p.ServiceAccountScopes = append(p.ServiceAccountScopes, r)
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	p.Metadata = make(map[string]string)
	for k, r := range o.Metadata {
		p.Metadata[k] = r
	}
	return p
}

// ClusterClusterConfigGceClusterConfigReservationAffinityToProto converts a ClusterClusterConfigGceClusterConfigReservationAffinity resource to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o *beta.ClusterClusterConfigGceClusterConfigReservationAffinity) *betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity resource to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *beta.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *betapb.DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.ValueOrEmptyString(o.NodeGroup),
	}
	return p
}

// ClusterInstanceGroupConfigToProto converts a ClusterInstanceGroupConfig resource to its proto representation.
func DataprocBetaClusterInstanceGroupConfigToProto(o *beta.ClusterInstanceGroupConfig) *betapb.DataprocBetaClusterInstanceGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfig{
		NumInstances:       dcl.ValueOrEmptyInt64(o.NumInstances),
		Image:              dcl.ValueOrEmptyString(o.Image),
		MachineType:        dcl.ValueOrEmptyString(o.MachineType),
		DiskConfig:         DataprocBetaClusterInstanceGroupConfigDiskConfigToProto(o.DiskConfig),
		IsPreemptible:      dcl.ValueOrEmptyBool(o.IsPreemptible),
		Preemptibility:     DataprocBetaClusterInstanceGroupConfigPreemptibilityEnumToProto(o.Preemptibility),
		ManagedGroupConfig: DataprocBetaClusterInstanceGroupConfigManagedGroupConfigToProto(o.ManagedGroupConfig),
		MinCpuPlatform:     dcl.ValueOrEmptyString(o.MinCpuPlatform),
	}
	for _, r := range o.InstanceNames {
		p.InstanceNames = append(p.InstanceNames, r)
	}
	for _, r := range o.Accelerators {
		p.Accelerators = append(p.Accelerators, DataprocBetaClusterInstanceGroupConfigAcceleratorsToProto(&r))
	}
	return p
}

// ClusterInstanceGroupConfigDiskConfigToProto converts a ClusterInstanceGroupConfigDiskConfig resource to its proto representation.
func DataprocBetaClusterInstanceGroupConfigDiskConfigToProto(o *beta.ClusterInstanceGroupConfigDiskConfig) *betapb.DataprocBetaClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.ValueOrEmptyString(o.BootDiskType),
		BootDiskSizeGb: dcl.ValueOrEmptyInt64(o.BootDiskSizeGb),
		NumLocalSsds:   dcl.ValueOrEmptyInt64(o.NumLocalSsds),
	}
	return p
}

// ClusterInstanceGroupConfigManagedGroupConfigToProto converts a ClusterInstanceGroupConfigManagedGroupConfig resource to its proto representation.
func DataprocBetaClusterInstanceGroupConfigManagedGroupConfigToProto(o *beta.ClusterInstanceGroupConfigManagedGroupConfig) *betapb.DataprocBetaClusterInstanceGroupConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.ValueOrEmptyString(o.InstanceTemplateName),
		InstanceGroupManagerName: dcl.ValueOrEmptyString(o.InstanceGroupManagerName),
	}
	return p
}

// ClusterInstanceGroupConfigAcceleratorsToProto converts a ClusterInstanceGroupConfigAccelerators resource to its proto representation.
func DataprocBetaClusterInstanceGroupConfigAcceleratorsToProto(o *beta.ClusterInstanceGroupConfigAccelerators) *betapb.DataprocBetaClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
	}
	return p
}

// ClusterClusterConfigSoftwareConfigToProto converts a ClusterClusterConfigSoftwareConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigSoftwareConfigToProto(o *beta.ClusterClusterConfigSoftwareConfig) *betapb.DataprocBetaClusterClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.ValueOrEmptyString(o.ImageVersion),
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.OptionalComponents {
		p.OptionalComponents = append(p.OptionalComponents, betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)]))
	}
	return p
}

// ClusterClusterConfigInitializationActionsToProto converts a ClusterClusterConfigInitializationActions resource to its proto representation.
func DataprocBetaClusterClusterConfigInitializationActionsToProto(o *beta.ClusterClusterConfigInitializationActions) *betapb.DataprocBetaClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.ValueOrEmptyString(o.ExecutableFile),
		ExecutionTimeout: dcl.ValueOrEmptyString(o.ExecutionTimeout),
	}
	return p
}

// ClusterClusterConfigEncryptionConfigToProto converts a ClusterClusterConfigEncryptionConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigEncryptionConfigToProto(o *beta.ClusterClusterConfigEncryptionConfig) *betapb.DataprocBetaClusterClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.ValueOrEmptyString(o.GcePdKmsKeyName),
	}
	return p
}

// ClusterClusterConfigAutoscalingConfigToProto converts a ClusterClusterConfigAutoscalingConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigAutoscalingConfigToProto(o *beta.ClusterClusterConfigAutoscalingConfig) *betapb.DataprocBetaClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigAutoscalingConfig{
		Policy: dcl.ValueOrEmptyString(o.Policy),
	}
	return p
}

// ClusterClusterConfigSecurityConfigToProto converts a ClusterClusterConfigSecurityConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigSecurityConfigToProto(o *beta.ClusterClusterConfigSecurityConfig) *betapb.DataprocBetaClusterClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigSecurityConfig{
		KerberosConfig: DataprocBetaClusterClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig),
	}
	return p
}

// ClusterClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterClusterConfigSecurityConfigKerberosConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigSecurityConfigKerberosConfigToProto(o *beta.ClusterClusterConfigSecurityConfigKerberosConfig) *betapb.DataprocBetaClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigSecurityConfigKerberosConfig{
		EnableKerberos:                dcl.ValueOrEmptyBool(o.EnableKerberos),
		RootPrincipalPassword:         dcl.ValueOrEmptyString(o.RootPrincipalPassword),
		KmsKey:                        dcl.ValueOrEmptyString(o.KmsKey),
		Keystore:                      dcl.ValueOrEmptyString(o.Keystore),
		Truststore:                    dcl.ValueOrEmptyString(o.Truststore),
		KeystorePassword:              dcl.ValueOrEmptyString(o.KeystorePassword),
		KeyPassword:                   dcl.ValueOrEmptyString(o.KeyPassword),
		TruststorePassword:            dcl.ValueOrEmptyString(o.TruststorePassword),
		CrossRealmTrustRealm:          dcl.ValueOrEmptyString(o.CrossRealmTrustRealm),
		CrossRealmTrustKdc:            dcl.ValueOrEmptyString(o.CrossRealmTrustKdc),
		CrossRealmTrustAdminServer:    dcl.ValueOrEmptyString(o.CrossRealmTrustAdminServer),
		CrossRealmTrustSharedPassword: dcl.ValueOrEmptyString(o.CrossRealmTrustSharedPassword),
		KdcDbKey:                      dcl.ValueOrEmptyString(o.KdcDbKey),
		TgtLifetimeHours:              dcl.ValueOrEmptyInt64(o.TgtLifetimeHours),
		Realm:                         dcl.ValueOrEmptyString(o.Realm),
	}
	return p
}

// ClusterClusterConfigLifecycleConfigToProto converts a ClusterClusterConfigLifecycleConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigLifecycleConfigToProto(o *beta.ClusterClusterConfigLifecycleConfig) *betapb.DataprocBetaClusterClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.ValueOrEmptyString(o.IdleDeleteTtl),
		AutoDeleteTime: dcl.ValueOrEmptyString(o.AutoDeleteTime),
		AutoDeleteTtl:  dcl.ValueOrEmptyString(o.AutoDeleteTtl),
		IdleStartTime:  dcl.ValueOrEmptyString(o.IdleStartTime),
	}
	return p
}

// ClusterClusterConfigEndpointConfigToProto converts a ClusterClusterConfigEndpointConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigEndpointConfigToProto(o *beta.ClusterClusterConfigEndpointConfig) *betapb.DataprocBetaClusterClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.ValueOrEmptyBool(o.EnableHttpPortAccess),
	}
	p.HttpPorts = make(map[string]string)
	for k, r := range o.HttpPorts {
		p.HttpPorts[k] = r
	}
	return p
}

// ClusterClusterConfigGkeClusterConfigToProto converts a ClusterClusterConfigGkeClusterConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigGkeClusterConfigToProto(o *beta.ClusterClusterConfigGkeClusterConfig) *betapb.DataprocBetaClusterClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget),
	}
	return p
}

// ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget resource to its proto representation.
func DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *beta.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *betapb.DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.ValueOrEmptyString(o.TargetGkeCluster),
		ClusterNamespace: dcl.ValueOrEmptyString(o.ClusterNamespace),
	}
	return p
}

// ClusterClusterConfigMetastoreConfigToProto converts a ClusterClusterConfigMetastoreConfig resource to its proto representation.
func DataprocBetaClusterClusterConfigMetastoreConfigToProto(o *beta.ClusterClusterConfigMetastoreConfig) *betapb.DataprocBetaClusterClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.ValueOrEmptyString(o.DataprocMetastoreService),
	}
	return p
}

// ClusterStatusToProto converts a ClusterStatus resource to its proto representation.
func DataprocBetaClusterStatusToProto(o *beta.ClusterStatus) *betapb.DataprocBetaClusterStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterStatus{
		State:          DataprocBetaClusterStatusStateEnumToProto(o.State),
		Detail:         dcl.ValueOrEmptyString(o.Detail),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocBetaClusterStatusSubstateEnumToProto(o.Substate),
	}
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory resource to its proto representation.
func DataprocBetaClusterStatusHistoryToProto(o *beta.ClusterStatusHistory) *betapb.DataprocBetaClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterStatusHistory{
		State:          DataprocBetaClusterStatusHistoryStateEnumToProto(o.State),
		Detail:         dcl.ValueOrEmptyString(o.Detail),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocBetaClusterStatusHistorySubstateEnumToProto(o.Substate),
	}
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics resource to its proto representation.
func DataprocBetaClusterMetricsToProto(o *beta.ClusterMetrics) *betapb.DataprocBetaClusterMetrics {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterMetrics{}
	p.HdfsMetrics = make(map[string]string)
	for k, r := range o.HdfsMetrics {
		p.HdfsMetrics[k] = r
	}
	p.YarnMetrics = make(map[string]string)
	for k, r := range o.YarnMetrics {
		p.YarnMetrics[k] = r
	}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *beta.Cluster) *betapb.DataprocBetaCluster {
	p := &betapb.DataprocBetaCluster{
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Config:      DataprocBetaClusterClusterConfigToProto(resource.Config),
		Status:      DataprocBetaClusterStatusToProto(resource.Status),
		ClusterUuid: dcl.ValueOrEmptyString(resource.ClusterUuid),
		Metrics:     DataprocBetaClusterMetricsToProto(resource.Metrics),
		Location:    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.StatusHistory {
		p.StatusHistory = append(p.StatusHistory, DataprocBetaClusterStatusHistoryToProto(&r))
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaClusterRequest) (*betapb.DataprocBetaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocBetaCluster(ctx context.Context, request *betapb.ApplyDataprocBetaClusterRequest) (*betapb.DataprocBetaCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocBetaCluster(ctx context.Context, request *betapb.DeleteDataprocBetaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocBetaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocBetaCluster(ctx context.Context, request *betapb.ListDataprocBetaClusterRequest) (*betapb.ListDataprocBetaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, ProtoToCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListDataprocBetaClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
