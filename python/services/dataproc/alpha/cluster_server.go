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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/alpha/dataproc_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *alpha.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := alpha.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *alpha.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := alpha.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterInstanceGroupConfigPreemptibilityEnum converts a ClusterInstanceGroupConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum(e alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum) *alpha.ClusterInstanceGroupConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.ClusterInstanceGroupConfigPreemptibilityEnum(n[len("DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfigSoftwareConfigOptionalComponentsEnum converts a ClusterClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(e alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum) *alpha.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := alpha.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusStateEnum converts a ClusterStatusStateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusStateEnum(e alphapb.DataprocAlphaClusterStatusStateEnum) *alpha.ClusterStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusStateEnum(n[len("DataprocAlphaClusterStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusSubstateEnum converts a ClusterStatusSubstateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusSubstateEnum(e alphapb.DataprocAlphaClusterStatusSubstateEnum) *alpha.ClusterStatusSubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusSubstateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusSubstateEnum(n[len("DataprocAlphaClusterStatusSubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistoryStateEnum converts a ClusterStatusHistoryStateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistoryStateEnum(e alphapb.DataprocAlphaClusterStatusHistoryStateEnum) *alpha.ClusterStatusHistoryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusHistoryStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusHistoryStateEnum(n[len("DataprocAlphaClusterStatusHistoryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistorySubstateEnum converts a ClusterStatusHistorySubstateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistorySubstateEnum(e alphapb.DataprocAlphaClusterStatusHistorySubstateEnum) *alpha.ClusterStatusHistorySubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusHistorySubstateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusHistorySubstateEnum(n[len("DataprocAlphaClusterStatusHistorySubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfig converts a ClusterClusterConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfig(p *alphapb.DataprocAlphaClusterClusterConfig) *alpha.ClusterClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.StagingBucket),
		TempBucket:            dcl.StringOrNil(p.TempBucket),
		GceClusterConfig:      ProtoToDataprocAlphaClusterClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocAlphaClusterInstanceGroupConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocAlphaClusterInstanceGroupConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAlphaClusterInstanceGroupConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocAlphaClusterClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocAlphaClusterClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocAlphaClusterClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocAlphaClusterClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocAlphaClusterClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocAlphaClusterClusterConfigEndpointConfig(p.GetEndpointConfig()),
		GkeClusterConfig:      ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		MetastoreConfig:       ProtoToDataprocAlphaClusterClusterConfigMetastoreConfig(p.GetMetastoreConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocAlphaClusterClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfig converts a ClusterClusterConfigGceClusterConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfig(p *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfig) *alpha.ClusterClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.Zone),
		Network:                 dcl.StringOrNil(p.Network),
		Subnetwork:              dcl.StringOrNil(p.Subnetwork),
		InternalIPOnly:          dcl.Bool(p.InternalIpOnly),
		PrivateIPv6GoogleAccess: ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.ServiceAccount),
		ReservationAffinity:     ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
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
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity(p *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity) *alpha.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity(p *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity) *alpha.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.NodeGroup),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfig converts a ClusterInstanceGroupConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfig(p *alphapb.DataprocAlphaClusterInstanceGroupConfig) *alpha.ClusterInstanceGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfig{
		NumInstances:       dcl.Int64OrNil(p.NumInstances),
		Image:              dcl.StringOrNil(p.Image),
		MachineType:        dcl.StringOrNil(p.MachineType),
		DiskConfig:         ProtoToDataprocAlphaClusterInstanceGroupConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.IsPreemptible),
		Preemptibility:     ProtoToDataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaClusterInstanceGroupConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.MinCpuPlatform),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaClusterInstanceGroupConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigDiskConfig converts a ClusterInstanceGroupConfigDiskConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigDiskConfig(p *alphapb.DataprocAlphaClusterInstanceGroupConfigDiskConfig) *alpha.ClusterInstanceGroupConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.BootDiskType),
		BootDiskSizeGb: dcl.Int64OrNil(p.BootDiskSizeGb),
		NumLocalSsds:   dcl.Int64OrNil(p.NumLocalSsds),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigManagedGroupConfig converts a ClusterInstanceGroupConfigManagedGroupConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigManagedGroupConfig(p *alphapb.DataprocAlphaClusterInstanceGroupConfigManagedGroupConfig) *alpha.ClusterInstanceGroupConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.InstanceTemplateName),
		InstanceGroupManagerName: dcl.StringOrNil(p.InstanceGroupManagerName),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigAccelerators converts a ClusterInstanceGroupConfigAccelerators resource from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigAccelerators(p *alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators) *alpha.ClusterInstanceGroupConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
	}
	return obj
}

// ProtoToClusterClusterConfigSoftwareConfig converts a ClusterClusterConfigSoftwareConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSoftwareConfig(p *alphapb.DataprocAlphaClusterClusterConfigSoftwareConfig) *alpha.ClusterClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.ImageVersion),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterClusterConfigInitializationActions converts a ClusterClusterConfigInitializationActions resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigInitializationActions(p *alphapb.DataprocAlphaClusterClusterConfigInitializationActions) *alpha.ClusterClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.ExecutableFile),
		ExecutionTimeout: dcl.StringOrNil(p.ExecutionTimeout),
	}
	return obj
}

// ProtoToClusterClusterConfigEncryptionConfig converts a ClusterClusterConfigEncryptionConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigEncryptionConfig(p *alphapb.DataprocAlphaClusterClusterConfigEncryptionConfig) *alpha.ClusterClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GcePdKmsKeyName),
	}
	return obj
}

// ProtoToClusterClusterConfigAutoscalingConfig converts a ClusterClusterConfigAutoscalingConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigAutoscalingConfig(p *alphapb.DataprocAlphaClusterClusterConfigAutoscalingConfig) *alpha.ClusterClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.Policy),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfig converts a ClusterClusterConfigSecurityConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSecurityConfig(p *alphapb.DataprocAlphaClusterClusterConfigSecurityConfig) *alpha.ClusterClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfigKerberosConfig converts a ClusterClusterConfigSecurityConfigKerberosConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig(p *alphapb.DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig) *alpha.ClusterClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocAlphaClusterClusterConfigLifecycleConfig(p *alphapb.DataprocAlphaClusterClusterConfigLifecycleConfig) *alpha.ClusterClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.IdleDeleteTtl),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.AutoDeleteTtl),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterClusterConfigEndpointConfig converts a ClusterClusterConfigEndpointConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigEndpointConfig(p *alphapb.DataprocAlphaClusterClusterConfigEndpointConfig) *alpha.ClusterClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.EnableHttpPortAccess),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfig converts a ClusterClusterConfigGkeClusterConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfig(p *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfig) *alpha.ClusterClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alpha.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.TargetGkeCluster),
		ClusterNamespace: dcl.StringOrNil(p.ClusterNamespace),
	}
	return obj
}

// ProtoToClusterClusterConfigMetastoreConfig converts a ClusterClusterConfigMetastoreConfig resource from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigMetastoreConfig(p *alphapb.DataprocAlphaClusterClusterConfigMetastoreConfig) *alpha.ClusterClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.DataprocMetastoreService),
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus resource from its proto representation.
func ProtoToDataprocAlphaClusterStatus(p *alphapb.DataprocAlphaClusterStatus) *alpha.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterStatus{
		State:          ProtoToDataprocAlphaClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.Detail),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocAlphaClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory resource from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistory(p *alphapb.DataprocAlphaClusterStatusHistory) *alpha.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterStatusHistory{
		State:          ProtoToDataprocAlphaClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.Detail),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocAlphaClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics resource from its proto representation.
func ProtoToDataprocAlphaClusterMetrics(p *alphapb.DataprocAlphaClusterMetrics) *alpha.ClusterMetrics {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterMetrics{}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.DataprocAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Project:     dcl.StringOrNil(p.Project),
		Name:        dcl.StringOrNil(p.Name),
		Config:      ProtoToDataprocAlphaClusterClusterConfig(p.GetConfig()),
		Status:      ProtoToDataprocAlphaClusterStatus(p.GetStatus()),
		ClusterUuid: dcl.StringOrNil(p.ClusterUuid),
		Metrics:     ProtoToDataprocAlphaClusterMetrics(p.GetMetrics()),
		Location:    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetStatusHistory() {
		obj.StatusHistory = append(obj.StatusHistory, *ProtoToDataprocAlphaClusterStatusHistory(r))
	}
	return obj
}

// ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *alpha.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *alpha.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterInstanceGroupConfigPreemptibilityEnumToProto converts a ClusterInstanceGroupConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnumToProto(e *alpha.ClusterInstanceGroupConfigPreemptibilityEnum) alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum_value["ClusterInstanceGroupConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum(0)
}

// ClusterClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a ClusterClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *alpha.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum) alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value["ClusterClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// ClusterStatusStateEnumToProto converts a ClusterStatusStateEnum enum to its proto representation.
func DataprocAlphaClusterStatusStateEnumToProto(e *alpha.ClusterStatusStateEnum) alphapb.DataprocAlphaClusterStatusStateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusStateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusStateEnum_value["ClusterStatusStateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusStateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusStateEnum(0)
}

// ClusterStatusSubstateEnumToProto converts a ClusterStatusSubstateEnum enum to its proto representation.
func DataprocAlphaClusterStatusSubstateEnumToProto(e *alpha.ClusterStatusSubstateEnum) alphapb.DataprocAlphaClusterStatusSubstateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusSubstateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusSubstateEnum_value["ClusterStatusSubstateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusSubstateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusSubstateEnum(0)
}

// ClusterStatusHistoryStateEnumToProto converts a ClusterStatusHistoryStateEnum enum to its proto representation.
func DataprocAlphaClusterStatusHistoryStateEnumToProto(e *alpha.ClusterStatusHistoryStateEnum) alphapb.DataprocAlphaClusterStatusHistoryStateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusHistoryStateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusHistoryStateEnum_value["ClusterStatusHistoryStateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusHistoryStateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusHistoryStateEnum(0)
}

// ClusterStatusHistorySubstateEnumToProto converts a ClusterStatusHistorySubstateEnum enum to its proto representation.
func DataprocAlphaClusterStatusHistorySubstateEnumToProto(e *alpha.ClusterStatusHistorySubstateEnum) alphapb.DataprocAlphaClusterStatusHistorySubstateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusHistorySubstateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusHistorySubstateEnum_value["ClusterStatusHistorySubstateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusHistorySubstateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusHistorySubstateEnum(0)
}

// ClusterClusterConfigToProto converts a ClusterClusterConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigToProto(o *alpha.ClusterClusterConfig) *alphapb.DataprocAlphaClusterClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfig{
		StagingBucket:         dcl.ValueOrEmptyString(o.StagingBucket),
		TempBucket:            dcl.ValueOrEmptyString(o.TempBucket),
		GceClusterConfig:      DataprocAlphaClusterClusterConfigGceClusterConfigToProto(o.GceClusterConfig),
		MasterConfig:          DataprocAlphaClusterInstanceGroupConfigToProto(o.MasterConfig),
		WorkerConfig:          DataprocAlphaClusterInstanceGroupConfigToProto(o.WorkerConfig),
		SecondaryWorkerConfig: DataprocAlphaClusterInstanceGroupConfigToProto(o.SecondaryWorkerConfig),
		SoftwareConfig:        DataprocAlphaClusterClusterConfigSoftwareConfigToProto(o.SoftwareConfig),
		EncryptionConfig:      DataprocAlphaClusterClusterConfigEncryptionConfigToProto(o.EncryptionConfig),
		AutoscalingConfig:     DataprocAlphaClusterClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig),
		SecurityConfig:        DataprocAlphaClusterClusterConfigSecurityConfigToProto(o.SecurityConfig),
		LifecycleConfig:       DataprocAlphaClusterClusterConfigLifecycleConfigToProto(o.LifecycleConfig),
		EndpointConfig:        DataprocAlphaClusterClusterConfigEndpointConfigToProto(o.EndpointConfig),
		GkeClusterConfig:      DataprocAlphaClusterClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig),
		MetastoreConfig:       DataprocAlphaClusterClusterConfigMetastoreConfigToProto(o.MetastoreConfig),
	}
	for _, r := range o.InitializationActions {
		p.InitializationActions = append(p.InitializationActions, DataprocAlphaClusterClusterConfigInitializationActionsToProto(&r))
	}
	return p
}

// ClusterClusterConfigGceClusterConfigToProto converts a ClusterClusterConfigGceClusterConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigToProto(o *alpha.ClusterClusterConfigGceClusterConfig) *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.ValueOrEmptyString(o.Zone),
		Network:                 dcl.ValueOrEmptyString(o.Network),
		Subnetwork:              dcl.ValueOrEmptyString(o.Subnetwork),
		InternalIpOnly:          dcl.ValueOrEmptyBool(o.InternalIPOnly),
		PrivateIpv6GoogleAccess: DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess),
		ServiceAccount:          dcl.ValueOrEmptyString(o.ServiceAccount),
		ReservationAffinity:     DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity),
		NodeGroupAffinity:       DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity),
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
func DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o *alpha.ClusterClusterConfigGceClusterConfigReservationAffinity) *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity resource to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *alpha.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.ValueOrEmptyString(o.NodeGroup),
	}
	return p
}

// ClusterInstanceGroupConfigToProto converts a ClusterInstanceGroupConfig resource to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigToProto(o *alpha.ClusterInstanceGroupConfig) *alphapb.DataprocAlphaClusterInstanceGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfig{
		NumInstances:       dcl.ValueOrEmptyInt64(o.NumInstances),
		Image:              dcl.ValueOrEmptyString(o.Image),
		MachineType:        dcl.ValueOrEmptyString(o.MachineType),
		DiskConfig:         DataprocAlphaClusterInstanceGroupConfigDiskConfigToProto(o.DiskConfig),
		IsPreemptible:      dcl.ValueOrEmptyBool(o.IsPreemptible),
		Preemptibility:     DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnumToProto(o.Preemptibility),
		ManagedGroupConfig: DataprocAlphaClusterInstanceGroupConfigManagedGroupConfigToProto(o.ManagedGroupConfig),
		MinCpuPlatform:     dcl.ValueOrEmptyString(o.MinCpuPlatform),
	}
	for _, r := range o.InstanceNames {
		p.InstanceNames = append(p.InstanceNames, r)
	}
	for _, r := range o.Accelerators {
		p.Accelerators = append(p.Accelerators, DataprocAlphaClusterInstanceGroupConfigAcceleratorsToProto(&r))
	}
	return p
}

// ClusterInstanceGroupConfigDiskConfigToProto converts a ClusterInstanceGroupConfigDiskConfig resource to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigDiskConfigToProto(o *alpha.ClusterInstanceGroupConfigDiskConfig) *alphapb.DataprocAlphaClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.ValueOrEmptyString(o.BootDiskType),
		BootDiskSizeGb: dcl.ValueOrEmptyInt64(o.BootDiskSizeGb),
		NumLocalSsds:   dcl.ValueOrEmptyInt64(o.NumLocalSsds),
	}
	return p
}

// ClusterInstanceGroupConfigManagedGroupConfigToProto converts a ClusterInstanceGroupConfigManagedGroupConfig resource to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigManagedGroupConfigToProto(o *alpha.ClusterInstanceGroupConfigManagedGroupConfig) *alphapb.DataprocAlphaClusterInstanceGroupConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.ValueOrEmptyString(o.InstanceTemplateName),
		InstanceGroupManagerName: dcl.ValueOrEmptyString(o.InstanceGroupManagerName),
	}
	return p
}

// ClusterInstanceGroupConfigAcceleratorsToProto converts a ClusterInstanceGroupConfigAccelerators resource to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigAcceleratorsToProto(o *alpha.ClusterInstanceGroupConfigAccelerators) *alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
	}
	return p
}

// ClusterClusterConfigSoftwareConfigToProto converts a ClusterClusterConfigSoftwareConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigSoftwareConfigToProto(o *alpha.ClusterClusterConfigSoftwareConfig) *alphapb.DataprocAlphaClusterClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.ValueOrEmptyString(o.ImageVersion),
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.OptionalComponents {
		p.OptionalComponents = append(p.OptionalComponents, alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)]))
	}
	return p
}

// ClusterClusterConfigInitializationActionsToProto converts a ClusterClusterConfigInitializationActions resource to its proto representation.
func DataprocAlphaClusterClusterConfigInitializationActionsToProto(o *alpha.ClusterClusterConfigInitializationActions) *alphapb.DataprocAlphaClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.ValueOrEmptyString(o.ExecutableFile),
		ExecutionTimeout: dcl.ValueOrEmptyString(o.ExecutionTimeout),
	}
	return p
}

// ClusterClusterConfigEncryptionConfigToProto converts a ClusterClusterConfigEncryptionConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigEncryptionConfigToProto(o *alpha.ClusterClusterConfigEncryptionConfig) *alphapb.DataprocAlphaClusterClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.ValueOrEmptyString(o.GcePdKmsKeyName),
	}
	return p
}

// ClusterClusterConfigAutoscalingConfigToProto converts a ClusterClusterConfigAutoscalingConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigAutoscalingConfigToProto(o *alpha.ClusterClusterConfigAutoscalingConfig) *alphapb.DataprocAlphaClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigAutoscalingConfig{
		Policy: dcl.ValueOrEmptyString(o.Policy),
	}
	return p
}

// ClusterClusterConfigSecurityConfigToProto converts a ClusterClusterConfigSecurityConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigSecurityConfigToProto(o *alpha.ClusterClusterConfigSecurityConfig) *alphapb.DataprocAlphaClusterClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigSecurityConfig{
		KerberosConfig: DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig),
	}
	return p
}

// ClusterClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterClusterConfigSecurityConfigKerberosConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfigToProto(o *alpha.ClusterClusterConfigSecurityConfigKerberosConfig) *alphapb.DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig{
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
func DataprocAlphaClusterClusterConfigLifecycleConfigToProto(o *alpha.ClusterClusterConfigLifecycleConfig) *alphapb.DataprocAlphaClusterClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.ValueOrEmptyString(o.IdleDeleteTtl),
		AutoDeleteTime: dcl.ValueOrEmptyString(o.AutoDeleteTime),
		AutoDeleteTtl:  dcl.ValueOrEmptyString(o.AutoDeleteTtl),
		IdleStartTime:  dcl.ValueOrEmptyString(o.IdleStartTime),
	}
	return p
}

// ClusterClusterConfigEndpointConfigToProto converts a ClusterClusterConfigEndpointConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigEndpointConfigToProto(o *alpha.ClusterClusterConfigEndpointConfig) *alphapb.DataprocAlphaClusterClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.ValueOrEmptyBool(o.EnableHttpPortAccess),
	}
	p.HttpPorts = make(map[string]string)
	for k, r := range o.HttpPorts {
		p.HttpPorts[k] = r
	}
	return p
}

// ClusterClusterConfigGkeClusterConfigToProto converts a ClusterClusterConfigGkeClusterConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigGkeClusterConfigToProto(o *alpha.ClusterClusterConfigGkeClusterConfig) *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget),
	}
	return p
}

// ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget resource to its proto representation.
func DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *alpha.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.ValueOrEmptyString(o.TargetGkeCluster),
		ClusterNamespace: dcl.ValueOrEmptyString(o.ClusterNamespace),
	}
	return p
}

// ClusterClusterConfigMetastoreConfigToProto converts a ClusterClusterConfigMetastoreConfig resource to its proto representation.
func DataprocAlphaClusterClusterConfigMetastoreConfigToProto(o *alpha.ClusterClusterConfigMetastoreConfig) *alphapb.DataprocAlphaClusterClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.ValueOrEmptyString(o.DataprocMetastoreService),
	}
	return p
}

// ClusterStatusToProto converts a ClusterStatus resource to its proto representation.
func DataprocAlphaClusterStatusToProto(o *alpha.ClusterStatus) *alphapb.DataprocAlphaClusterStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterStatus{
		State:          DataprocAlphaClusterStatusStateEnumToProto(o.State),
		Detail:         dcl.ValueOrEmptyString(o.Detail),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocAlphaClusterStatusSubstateEnumToProto(o.Substate),
	}
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory resource to its proto representation.
func DataprocAlphaClusterStatusHistoryToProto(o *alpha.ClusterStatusHistory) *alphapb.DataprocAlphaClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterStatusHistory{
		State:          DataprocAlphaClusterStatusHistoryStateEnumToProto(o.State),
		Detail:         dcl.ValueOrEmptyString(o.Detail),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocAlphaClusterStatusHistorySubstateEnumToProto(o.Substate),
	}
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics resource to its proto representation.
func DataprocAlphaClusterMetricsToProto(o *alpha.ClusterMetrics) *alphapb.DataprocAlphaClusterMetrics {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterMetrics{}
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
func ClusterToProto(resource *alpha.Cluster) *alphapb.DataprocAlphaCluster {
	p := &alphapb.DataprocAlphaCluster{
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Config:      DataprocAlphaClusterClusterConfigToProto(resource.Config),
		Status:      DataprocAlphaClusterStatusToProto(resource.Status),
		ClusterUuid: dcl.ValueOrEmptyString(resource.ClusterUuid),
		Metrics:     DataprocAlphaClusterMetricsToProto(resource.Metrics),
		Location:    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.StatusHistory {
		p.StatusHistory = append(p.StatusHistory, DataprocAlphaClusterStatusHistoryToProto(&r))
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaClusterRequest) (*alphapb.DataprocAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocAlphaCluster(ctx context.Context, request *alphapb.ApplyDataprocAlphaClusterRequest) (*alphapb.DataprocAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocAlphaCluster(ctx context.Context, request *alphapb.DeleteDataprocAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocAlphaCluster(ctx context.Context, request *alphapb.ListDataprocAlphaClusterRequest) (*alphapb.ListDataprocAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListDataprocAlphaClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
