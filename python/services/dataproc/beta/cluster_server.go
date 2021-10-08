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

// ProtoToClusterClusterConfig converts a ClusterClusterConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfig(p *betapb.DataprocBetaClusterClusterConfig) *beta.ClusterClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
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

// ProtoToClusterClusterConfigGceClusterConfig converts a ClusterClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfig(p *betapb.DataprocBetaClusterClusterConfigGceClusterConfig) *beta.ClusterClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
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

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinity converts a ClusterClusterConfigGceClusterConfigReservationAffinity object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity(p *betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity) *beta.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity(p *betapb.DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity) *beta.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfig converts a ClusterInstanceGroupConfig object from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfig(p *betapb.DataprocBetaClusterInstanceGroupConfig) *beta.ClusterInstanceGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaClusterInstanceGroupConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaClusterInstanceGroupConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaClusterInstanceGroupConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaClusterInstanceGroupConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigDiskConfig converts a ClusterInstanceGroupConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigDiskConfig(p *betapb.DataprocBetaClusterInstanceGroupConfigDiskConfig) *beta.ClusterInstanceGroupConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigManagedGroupConfig converts a ClusterInstanceGroupConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigManagedGroupConfig(p *betapb.DataprocBetaClusterInstanceGroupConfigManagedGroupConfig) *beta.ClusterInstanceGroupConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigAccelerators converts a ClusterInstanceGroupConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaClusterInstanceGroupConfigAccelerators(p *betapb.DataprocBetaClusterInstanceGroupConfigAccelerators) *beta.ClusterInstanceGroupConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterClusterConfigSoftwareConfig converts a ClusterClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSoftwareConfig(p *betapb.DataprocBetaClusterClusterConfigSoftwareConfig) *beta.ClusterClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterClusterConfigInitializationActions converts a ClusterClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigInitializationActions(p *betapb.DataprocBetaClusterClusterConfigInitializationActions) *beta.ClusterClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToClusterClusterConfigEncryptionConfig converts a ClusterClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigEncryptionConfig(p *betapb.DataprocBetaClusterClusterConfigEncryptionConfig) *beta.ClusterClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToClusterClusterConfigAutoscalingConfig converts a ClusterClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigAutoscalingConfig(p *betapb.DataprocBetaClusterClusterConfigAutoscalingConfig) *beta.ClusterClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfig converts a ClusterClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSecurityConfig(p *betapb.DataprocBetaClusterClusterConfigSecurityConfig) *beta.ClusterClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocBetaClusterClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfigKerberosConfig converts a ClusterClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigSecurityConfigKerberosConfig(p *betapb.DataprocBetaClusterClusterConfigSecurityConfigKerberosConfig) *beta.ClusterClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocBetaClusterClusterConfigLifecycleConfig(p *betapb.DataprocBetaClusterClusterConfigLifecycleConfig) *beta.ClusterClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterClusterConfigEndpointConfig converts a ClusterClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigEndpointConfig(p *betapb.DataprocBetaClusterClusterConfigEndpointConfig) *beta.ClusterClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfig converts a ClusterClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGkeClusterConfig(p *betapb.DataprocBetaClusterClusterConfigGkeClusterConfig) *beta.ClusterClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *betapb.DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *beta.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.GetTargetGkeCluster()),
		ClusterNamespace: dcl.StringOrNil(p.GetClusterNamespace()),
	}
	return obj
}

// ProtoToClusterClusterConfigMetastoreConfig converts a ClusterClusterConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocBetaClusterClusterConfigMetastoreConfig(p *betapb.DataprocBetaClusterClusterConfigMetastoreConfig) *beta.ClusterClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.GetDataprocMetastoreService()),
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus object from its proto representation.
func ProtoToDataprocBetaClusterStatus(p *betapb.DataprocBetaClusterStatus) *beta.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterStatus{
		State:          ProtoToDataprocBetaClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocBetaClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory object from its proto representation.
func ProtoToDataprocBetaClusterStatusHistory(p *betapb.DataprocBetaClusterStatusHistory) *beta.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterStatusHistory{
		State:          ProtoToDataprocBetaClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocBetaClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics object from its proto representation.
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
		Project:     dcl.StringOrNil(p.GetProject()),
		Name:        dcl.StringOrNil(p.GetName()),
		Config:      ProtoToDataprocBetaClusterClusterConfig(p.GetConfig()),
		Status:      ProtoToDataprocBetaClusterStatus(p.GetStatus()),
		ClusterUuid: dcl.StringOrNil(p.GetClusterUuid()),
		Metrics:     ProtoToDataprocBetaClusterMetrics(p.GetMetrics()),
		Location:    dcl.StringOrNil(p.GetLocation()),
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

// ClusterClusterConfigToProto converts a ClusterClusterConfig object to its proto representation.
func DataprocBetaClusterClusterConfigToProto(o *beta.ClusterClusterConfig) *betapb.DataprocBetaClusterClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocBetaClusterClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocBetaClusterInstanceGroupConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocBetaClusterInstanceGroupConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocBetaClusterInstanceGroupConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocBetaClusterClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocBetaClusterClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocBetaClusterClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocBetaClusterClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocBetaClusterClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocBetaClusterClusterConfigEndpointConfigToProto(o.EndpointConfig))
	p.SetGkeClusterConfig(DataprocBetaClusterClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetMetastoreConfig(DataprocBetaClusterClusterConfigMetastoreConfigToProto(o.MetastoreConfig))
	sInitializationActions := make([]*betapb.DataprocBetaClusterClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocBetaClusterClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// ClusterClusterConfigGceClusterConfigToProto converts a ClusterClusterConfigGceClusterConfig object to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigToProto(o *beta.ClusterClusterConfigGceClusterConfig) *betapb.DataprocBetaClusterClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocBetaClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
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
func DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityToProto(o *beta.ClusterClusterConfigGceClusterConfigReservationAffinity) *betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocBetaClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// ClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *beta.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *betapb.DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// ClusterInstanceGroupConfigToProto converts a ClusterInstanceGroupConfig object to its proto representation.
func DataprocBetaClusterInstanceGroupConfigToProto(o *beta.ClusterInstanceGroupConfig) *betapb.DataprocBetaClusterInstanceGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaClusterInstanceGroupConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaClusterInstanceGroupConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaClusterInstanceGroupConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaClusterInstanceGroupConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaClusterInstanceGroupConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterInstanceGroupConfigDiskConfigToProto converts a ClusterInstanceGroupConfigDiskConfig object to its proto representation.
func DataprocBetaClusterInstanceGroupConfigDiskConfigToProto(o *beta.ClusterInstanceGroupConfigDiskConfig) *betapb.DataprocBetaClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterInstanceGroupConfigManagedGroupConfigToProto converts a ClusterInstanceGroupConfigManagedGroupConfig object to its proto representation.
func DataprocBetaClusterInstanceGroupConfigManagedGroupConfigToProto(o *beta.ClusterInstanceGroupConfigManagedGroupConfig) *betapb.DataprocBetaClusterInstanceGroupConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterInstanceGroupConfigAcceleratorsToProto converts a ClusterInstanceGroupConfigAccelerators object to its proto representation.
func DataprocBetaClusterInstanceGroupConfigAcceleratorsToProto(o *beta.ClusterInstanceGroupConfigAccelerators) *betapb.DataprocBetaClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterInstanceGroupConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterClusterConfigSoftwareConfigToProto converts a ClusterClusterConfigSoftwareConfig object to its proto representation.
func DataprocBetaClusterClusterConfigSoftwareConfigToProto(o *beta.ClusterClusterConfigSoftwareConfig) *betapb.DataprocBetaClusterClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum(betapb.DataprocBetaClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// ClusterClusterConfigInitializationActionsToProto converts a ClusterClusterConfigInitializationActions object to its proto representation.
func DataprocBetaClusterClusterConfigInitializationActionsToProto(o *beta.ClusterClusterConfigInitializationActions) *betapb.DataprocBetaClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// ClusterClusterConfigEncryptionConfigToProto converts a ClusterClusterConfigEncryptionConfig object to its proto representation.
func DataprocBetaClusterClusterConfigEncryptionConfigToProto(o *beta.ClusterClusterConfigEncryptionConfig) *betapb.DataprocBetaClusterClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// ClusterClusterConfigAutoscalingConfigToProto converts a ClusterClusterConfigAutoscalingConfig object to its proto representation.
func DataprocBetaClusterClusterConfigAutoscalingConfigToProto(o *beta.ClusterClusterConfigAutoscalingConfig) *betapb.DataprocBetaClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// ClusterClusterConfigSecurityConfigToProto converts a ClusterClusterConfigSecurityConfig object to its proto representation.
func DataprocBetaClusterClusterConfigSecurityConfigToProto(o *beta.ClusterClusterConfigSecurityConfig) *betapb.DataprocBetaClusterClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocBetaClusterClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// ClusterClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocBetaClusterClusterConfigSecurityConfigKerberosConfigToProto(o *beta.ClusterClusterConfigSecurityConfigKerberosConfig) *betapb.DataprocBetaClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigSecurityConfigKerberosConfig{}
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
func DataprocBetaClusterClusterConfigLifecycleConfigToProto(o *beta.ClusterClusterConfigLifecycleConfig) *betapb.DataprocBetaClusterClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// ClusterClusterConfigEndpointConfigToProto converts a ClusterClusterConfigEndpointConfig object to its proto representation.
func DataprocBetaClusterClusterConfigEndpointConfigToProto(o *beta.ClusterClusterConfigEndpointConfig) *betapb.DataprocBetaClusterClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// ClusterClusterConfigGkeClusterConfigToProto converts a ClusterClusterConfigGkeClusterConfig object to its proto representation.
func DataprocBetaClusterClusterConfigGkeClusterConfigToProto(o *beta.ClusterClusterConfigGkeClusterConfig) *betapb.DataprocBetaClusterClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGkeClusterConfig{}
	p.SetNamespacedGkeDeploymentTarget(DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget))
	return p
}

// ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object to its proto representation.
func DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *beta.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *betapb.DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
	p.SetTargetGkeCluster(dcl.ValueOrEmptyString(o.TargetGkeCluster))
	p.SetClusterNamespace(dcl.ValueOrEmptyString(o.ClusterNamespace))
	return p
}

// ClusterClusterConfigMetastoreConfigToProto converts a ClusterClusterConfigMetastoreConfig object to its proto representation.
func DataprocBetaClusterClusterConfigMetastoreConfigToProto(o *beta.ClusterClusterConfigMetastoreConfig) *betapb.DataprocBetaClusterClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterClusterConfigMetastoreConfig{}
	p.SetDataprocMetastoreService(dcl.ValueOrEmptyString(o.DataprocMetastoreService))
	return p
}

// ClusterStatusToProto converts a ClusterStatus object to its proto representation.
func DataprocBetaClusterStatusToProto(o *beta.ClusterStatus) *betapb.DataprocBetaClusterStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterStatus{}
	p.SetState(DataprocBetaClusterStatusStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocBetaClusterStatusSubstateEnumToProto(o.Substate))
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory object to its proto representation.
func DataprocBetaClusterStatusHistoryToProto(o *beta.ClusterStatusHistory) *betapb.DataprocBetaClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterStatusHistory{}
	p.SetState(DataprocBetaClusterStatusHistoryStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocBetaClusterStatusHistorySubstateEnumToProto(o.Substate))
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics object to its proto representation.
func DataprocBetaClusterMetricsToProto(o *beta.ClusterMetrics) *betapb.DataprocBetaClusterMetrics {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterMetrics{}
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
func ClusterToProto(resource *beta.Cluster) *betapb.DataprocBetaCluster {
	p := &betapb.DataprocBetaCluster{}
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetConfig(DataprocBetaClusterClusterConfigToProto(resource.Config))
	p.SetStatus(DataprocBetaClusterStatusToProto(resource.Status))
	p.SetClusterUuid(dcl.ValueOrEmptyString(resource.ClusterUuid))
	p.SetMetrics(DataprocBetaClusterMetricsToProto(resource.Metrics))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStatusHistory := make([]*betapb.DataprocBetaClusterStatusHistory, len(resource.StatusHistory))
	for i, r := range resource.StatusHistory {
		sStatusHistory[i] = DataprocBetaClusterStatusHistoryToProto(&r)
	}
	p.SetStatusHistory(sStatusHistory)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaClusterRequest) (*betapb.DataprocBetaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyDataprocBetaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocBetaCluster(ctx context.Context, request *betapb.ApplyDataprocBetaClusterRequest) (*betapb.DataprocBetaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocBetaCluster(ctx context.Context, request *betapb.DeleteDataprocBetaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocBetaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocBetaCluster(ctx context.Context, request *betapb.ListDataprocBetaClusterRequest) (*betapb.ListDataprocBetaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataprocBetaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
