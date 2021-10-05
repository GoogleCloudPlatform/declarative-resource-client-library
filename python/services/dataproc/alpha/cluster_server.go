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

// ProtoToClusterClusterConfig converts a ClusterClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfig(p *alphapb.DataprocAlphaClusterClusterConfig) *alpha.ClusterClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
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

// ProtoToClusterClusterConfigGceClusterConfig converts a ClusterClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfig(p *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfig) *alpha.ClusterClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
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

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinity converts a ClusterClusterConfigGceClusterConfigReservationAffinity object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity(p *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity) *alpha.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity(p *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity) *alpha.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfig converts a ClusterInstanceGroupConfig object from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfig(p *alphapb.DataprocAlphaClusterInstanceGroupConfig) *alpha.ClusterInstanceGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaClusterInstanceGroupConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaClusterInstanceGroupConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaClusterInstanceGroupConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaClusterInstanceGroupConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigDiskConfig converts a ClusterInstanceGroupConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigDiskConfig(p *alphapb.DataprocAlphaClusterInstanceGroupConfigDiskConfig) *alpha.ClusterInstanceGroupConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigManagedGroupConfig converts a ClusterInstanceGroupConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigManagedGroupConfig(p *alphapb.DataprocAlphaClusterInstanceGroupConfigManagedGroupConfig) *alpha.ClusterInstanceGroupConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigAccelerators converts a ClusterInstanceGroupConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaClusterInstanceGroupConfigAccelerators(p *alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators) *alpha.ClusterInstanceGroupConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterClusterConfigSoftwareConfig converts a ClusterClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSoftwareConfig(p *alphapb.DataprocAlphaClusterClusterConfigSoftwareConfig) *alpha.ClusterClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterClusterConfigInitializationActions converts a ClusterClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigInitializationActions(p *alphapb.DataprocAlphaClusterClusterConfigInitializationActions) *alpha.ClusterClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToClusterClusterConfigEncryptionConfig converts a ClusterClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigEncryptionConfig(p *alphapb.DataprocAlphaClusterClusterConfigEncryptionConfig) *alpha.ClusterClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToClusterClusterConfigAutoscalingConfig converts a ClusterClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigAutoscalingConfig(p *alphapb.DataprocAlphaClusterClusterConfigAutoscalingConfig) *alpha.ClusterClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfig converts a ClusterClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSecurityConfig(p *alphapb.DataprocAlphaClusterClusterConfigSecurityConfig) *alpha.ClusterClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfigKerberosConfig converts a ClusterClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig(p *alphapb.DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig) *alpha.ClusterClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigSecurityConfigKerberosConfig{
		EnableKerberos:                dcl.Bool(p.GetEnableKerberos()),
		RootPrincipalPassword:         dcl.StringOrNil(p.GetRootPrincipalPassword()),
		KmsKey:                        dcl.StringOrNil(p.GetKmsKey()),
		Keystore:                      dcl.StringOrNil(p.GetKeystore()),
		Truststore:                    dcl.StringOrNil(p.GetTruststore()),
		KeystorePassword:              dcl.StringOrNil(p.GetKeystorePassword()),
		KeyPassword:                   dcl.StringOrNil(p.GetKeyPassword()),
		TruststorePassword:            dcl.StringOrNil(p.GetTruststorePassword()),
		CrossRealmTrustRealm:          dcl.StringOrNil(p.GetCrossRealmTrustRealm()),
		CrossRealmTrustKdc:            dcl.StringOrNil(p.GetCrossRealmTrustKdc()),
		CrossRealmTrustAdminServer:    dcl.StringOrNil(p.GetCrossRealmTrustAdminServer()),
		CrossRealmTrustSharedPassword: dcl.StringOrNil(p.GetCrossRealmTrustSharedPassword()),
		KdcDbKey:                      dcl.StringOrNil(p.GetKdcDbKey()),
		TgtLifetimeHours:              dcl.Int64OrNil(p.GetTgtLifetimeHours()),
		Realm:                         dcl.StringOrNil(p.GetRealm()),
	}
	return obj
}

// ProtoToClusterClusterConfigLifecycleConfig converts a ClusterClusterConfigLifecycleConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigLifecycleConfig(p *alphapb.DataprocAlphaClusterClusterConfigLifecycleConfig) *alpha.ClusterClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterClusterConfigEndpointConfig converts a ClusterClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigEndpointConfig(p *alphapb.DataprocAlphaClusterClusterConfigEndpointConfig) *alpha.ClusterClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfig converts a ClusterClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfig(p *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfig) *alpha.ClusterClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alpha.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.GetTargetGkeCluster()),
		ClusterNamespace: dcl.StringOrNil(p.GetClusterNamespace()),
	}
	return obj
}

// ProtoToClusterClusterConfigMetastoreConfig converts a ClusterClusterConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocAlphaClusterClusterConfigMetastoreConfig(p *alphapb.DataprocAlphaClusterClusterConfigMetastoreConfig) *alpha.ClusterClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.GetDataprocMetastoreService()),
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus object from its proto representation.
func ProtoToDataprocAlphaClusterStatus(p *alphapb.DataprocAlphaClusterStatus) *alpha.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterStatus{
		State:          ProtoToDataprocAlphaClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocAlphaClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory object from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistory(p *alphapb.DataprocAlphaClusterStatusHistory) *alpha.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterStatusHistory{
		State:          ProtoToDataprocAlphaClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocAlphaClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics object from its proto representation.
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
		Project:     dcl.StringOrNil(p.GetProject()),
		Name:        dcl.StringOrNil(p.GetName()),
		Config:      ProtoToDataprocAlphaClusterClusterConfig(p.GetConfig()),
		Status:      ProtoToDataprocAlphaClusterStatus(p.GetStatus()),
		ClusterUuid: dcl.StringOrNil(p.GetClusterUuid()),
		Metrics:     ProtoToDataprocAlphaClusterMetrics(p.GetMetrics()),
		Location:    dcl.StringOrNil(p.GetLocation()),
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

// ClusterClusterConfigToProto converts a ClusterClusterConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigToProto(o *alpha.ClusterClusterConfig) *alphapb.DataprocAlphaClusterClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocAlphaClusterClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocAlphaClusterInstanceGroupConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocAlphaClusterInstanceGroupConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocAlphaClusterInstanceGroupConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocAlphaClusterClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocAlphaClusterClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocAlphaClusterClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocAlphaClusterClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocAlphaClusterClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocAlphaClusterClusterConfigEndpointConfigToProto(o.EndpointConfig))
	p.SetGkeClusterConfig(DataprocAlphaClusterClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetMetastoreConfig(DataprocAlphaClusterClusterConfigMetastoreConfigToProto(o.MetastoreConfig))
	sInitializationActions := make([]*alphapb.DataprocAlphaClusterClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocAlphaClusterClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// ClusterClusterConfigGceClusterConfigToProto converts a ClusterClusterConfigGceClusterConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigToProto(o *alpha.ClusterClusterConfigGceClusterConfig) *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocAlphaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
	sServiceAccountScopes := make([]string, len(o.ServiceAccountScopes))
	for i, r := range o.ServiceAccountScopes {
		sServiceAccountScopes[i] = r
	}
	p.SetServiceAccountScopes(sServiceAccountScopes)
	sTags := make([]string, len(o.Tags))
	for i, r := range o.Tags {
		sTags[i] = r
	}
	p.SetTags(sTags)
	mMetadata := make(map[string]string, len(o.Metadata))
	for k, r := range o.Metadata {
		mMetadata[k] = r
	}
	p.SetMetadata(mMetadata)
	return p
}

// ClusterClusterConfigGceClusterConfigReservationAffinityToProto converts a ClusterClusterConfigGceClusterConfigReservationAffinity object to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o *alpha.ClusterClusterConfigGceClusterConfigReservationAffinity) *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocAlphaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// ClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *alpha.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// ClusterInstanceGroupConfigToProto converts a ClusterInstanceGroupConfig object to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigToProto(o *alpha.ClusterInstanceGroupConfig) *alphapb.DataprocAlphaClusterInstanceGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaClusterInstanceGroupConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaClusterInstanceGroupConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaClusterInstanceGroupConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaClusterInstanceGroupConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterInstanceGroupConfigDiskConfigToProto converts a ClusterInstanceGroupConfigDiskConfig object to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigDiskConfigToProto(o *alpha.ClusterInstanceGroupConfigDiskConfig) *alphapb.DataprocAlphaClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterInstanceGroupConfigManagedGroupConfigToProto converts a ClusterInstanceGroupConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigManagedGroupConfigToProto(o *alpha.ClusterInstanceGroupConfigManagedGroupConfig) *alphapb.DataprocAlphaClusterInstanceGroupConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterInstanceGroupConfigAcceleratorsToProto converts a ClusterInstanceGroupConfigAccelerators object to its proto representation.
func DataprocAlphaClusterInstanceGroupConfigAcceleratorsToProto(o *alpha.ClusterInstanceGroupConfigAccelerators) *alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterInstanceGroupConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterClusterConfigSoftwareConfigToProto converts a ClusterClusterConfigSoftwareConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigSoftwareConfigToProto(o *alpha.ClusterClusterConfigSoftwareConfig) *alphapb.DataprocAlphaClusterClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(alphapb.DataprocAlphaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// ClusterClusterConfigInitializationActionsToProto converts a ClusterClusterConfigInitializationActions object to its proto representation.
func DataprocAlphaClusterClusterConfigInitializationActionsToProto(o *alpha.ClusterClusterConfigInitializationActions) *alphapb.DataprocAlphaClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// ClusterClusterConfigEncryptionConfigToProto converts a ClusterClusterConfigEncryptionConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigEncryptionConfigToProto(o *alpha.ClusterClusterConfigEncryptionConfig) *alphapb.DataprocAlphaClusterClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// ClusterClusterConfigAutoscalingConfigToProto converts a ClusterClusterConfigAutoscalingConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigAutoscalingConfigToProto(o *alpha.ClusterClusterConfigAutoscalingConfig) *alphapb.DataprocAlphaClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// ClusterClusterConfigSecurityConfigToProto converts a ClusterClusterConfigSecurityConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigSecurityConfigToProto(o *alpha.ClusterClusterConfigSecurityConfig) *alphapb.DataprocAlphaClusterClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// ClusterClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfigToProto(o *alpha.ClusterClusterConfigSecurityConfigKerberosConfig) *alphapb.DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigSecurityConfigKerberosConfig{}
	p.SetEnableKerberos(dcl.ValueOrEmptyBool(o.EnableKerberos))
	p.SetRootPrincipalPassword(dcl.ValueOrEmptyString(o.RootPrincipalPassword))
	p.SetKmsKey(dcl.ValueOrEmptyString(o.KmsKey))
	p.SetKeystore(dcl.ValueOrEmptyString(o.Keystore))
	p.SetTruststore(dcl.ValueOrEmptyString(o.Truststore))
	p.SetKeystorePassword(dcl.ValueOrEmptyString(o.KeystorePassword))
	p.SetKeyPassword(dcl.ValueOrEmptyString(o.KeyPassword))
	p.SetTruststorePassword(dcl.ValueOrEmptyString(o.TruststorePassword))
	p.SetCrossRealmTrustRealm(dcl.ValueOrEmptyString(o.CrossRealmTrustRealm))
	p.SetCrossRealmTrustKdc(dcl.ValueOrEmptyString(o.CrossRealmTrustKdc))
	p.SetCrossRealmTrustAdminServer(dcl.ValueOrEmptyString(o.CrossRealmTrustAdminServer))
	p.SetCrossRealmTrustSharedPassword(dcl.ValueOrEmptyString(o.CrossRealmTrustSharedPassword))
	p.SetKdcDbKey(dcl.ValueOrEmptyString(o.KdcDbKey))
	p.SetTgtLifetimeHours(dcl.ValueOrEmptyInt64(o.TgtLifetimeHours))
	p.SetRealm(dcl.ValueOrEmptyString(o.Realm))
	return p
}

// ClusterClusterConfigLifecycleConfigToProto converts a ClusterClusterConfigLifecycleConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigLifecycleConfigToProto(o *alpha.ClusterClusterConfigLifecycleConfig) *alphapb.DataprocAlphaClusterClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// ClusterClusterConfigEndpointConfigToProto converts a ClusterClusterConfigEndpointConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigEndpointConfigToProto(o *alpha.ClusterClusterConfigEndpointConfig) *alphapb.DataprocAlphaClusterClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// ClusterClusterConfigGkeClusterConfigToProto converts a ClusterClusterConfigGkeClusterConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigGkeClusterConfigToProto(o *alpha.ClusterClusterConfigGkeClusterConfig) *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfig{}
	p.SetNamespacedGkeDeploymentTarget(DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget))
	return p
}

// ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object to its proto representation.
func DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *alpha.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
	p.SetTargetGkeCluster(dcl.ValueOrEmptyString(o.TargetGkeCluster))
	p.SetClusterNamespace(dcl.ValueOrEmptyString(o.ClusterNamespace))
	return p
}

// ClusterClusterConfigMetastoreConfigToProto converts a ClusterClusterConfigMetastoreConfig object to its proto representation.
func DataprocAlphaClusterClusterConfigMetastoreConfigToProto(o *alpha.ClusterClusterConfigMetastoreConfig) *alphapb.DataprocAlphaClusterClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterClusterConfigMetastoreConfig{}
	p.SetDataprocMetastoreService(dcl.ValueOrEmptyString(o.DataprocMetastoreService))
	return p
}

// ClusterStatusToProto converts a ClusterStatus object to its proto representation.
func DataprocAlphaClusterStatusToProto(o *alpha.ClusterStatus) *alphapb.DataprocAlphaClusterStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterStatus{}
	p.SetState(DataprocAlphaClusterStatusStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocAlphaClusterStatusSubstateEnumToProto(o.Substate))
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory object to its proto representation.
func DataprocAlphaClusterStatusHistoryToProto(o *alpha.ClusterStatusHistory) *alphapb.DataprocAlphaClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterStatusHistory{}
	p.SetState(DataprocAlphaClusterStatusHistoryStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocAlphaClusterStatusHistorySubstateEnumToProto(o.Substate))
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics object to its proto representation.
func DataprocAlphaClusterMetricsToProto(o *alpha.ClusterMetrics) *alphapb.DataprocAlphaClusterMetrics {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterMetrics{}
	mHdfsMetrics := make(map[string]string, len(o.HdfsMetrics))
	for k, r := range o.HdfsMetrics {
		mHdfsMetrics[k] = r
	}
	p.SetHdfsMetrics(mHdfsMetrics)
	mYarnMetrics := make(map[string]string, len(o.YarnMetrics))
	for k, r := range o.YarnMetrics {
		mYarnMetrics[k] = r
	}
	p.SetYarnMetrics(mYarnMetrics)
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.DataprocAlphaCluster {
	p := &alphapb.DataprocAlphaCluster{}
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetConfig(DataprocAlphaClusterClusterConfigToProto(resource.Config))
	p.SetStatus(DataprocAlphaClusterStatusToProto(resource.Status))
	p.SetClusterUuid(dcl.ValueOrEmptyString(resource.ClusterUuid))
	p.SetMetrics(DataprocAlphaClusterMetricsToProto(resource.Metrics))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStatusHistory := make([]*alphapb.DataprocAlphaClusterStatusHistory, len(resource.StatusHistory))
	for i, r := range resource.StatusHistory {
		sStatusHistory[i] = DataprocAlphaClusterStatusHistoryToProto(&r)
	}
	p.SetStatusHistory(sStatusHistory)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaClusterRequest) (*alphapb.DataprocAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyDataprocAlphaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocAlphaCluster(ctx context.Context, request *alphapb.ApplyDataprocAlphaClusterRequest) (*alphapb.DataprocAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocAlphaCluster(ctx context.Context, request *alphapb.DeleteDataprocAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocAlphaCluster(ctx context.Context, request *alphapb.ListDataprocAlphaClusterRequest) (*alphapb.ListDataprocAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataprocAlphaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
