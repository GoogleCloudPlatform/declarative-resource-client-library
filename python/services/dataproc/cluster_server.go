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
	dataprocpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/dataproc_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *dataproc.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := dataproc.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := dataproc.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterInstanceGroupConfigPreemptibilityEnum converts a ClusterInstanceGroupConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigPreemptibilityEnum(e dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum) *dataproc.ClusterInstanceGroupConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.ClusterInstanceGroupConfigPreemptibilityEnum(n[len("DataprocClusterInstanceGroupConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfigSoftwareConfigOptionalComponentsEnum converts a ClusterClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(e dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum) *dataproc.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := dataproc.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusStateEnum converts a ClusterStatusStateEnum enum from its proto representation.
func ProtoToDataprocClusterStatusStateEnum(e dataprocpb.DataprocClusterStatusStateEnum) *dataproc.ClusterStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterStatusStateEnum_name[int32(e)]; ok {
		e := dataproc.ClusterStatusStateEnum(n[len("DataprocClusterStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusSubstateEnum converts a ClusterStatusSubstateEnum enum from its proto representation.
func ProtoToDataprocClusterStatusSubstateEnum(e dataprocpb.DataprocClusterStatusSubstateEnum) *dataproc.ClusterStatusSubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterStatusSubstateEnum_name[int32(e)]; ok {
		e := dataproc.ClusterStatusSubstateEnum(n[len("DataprocClusterStatusSubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistoryStateEnum converts a ClusterStatusHistoryStateEnum enum from its proto representation.
func ProtoToDataprocClusterStatusHistoryStateEnum(e dataprocpb.DataprocClusterStatusHistoryStateEnum) *dataproc.ClusterStatusHistoryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterStatusHistoryStateEnum_name[int32(e)]; ok {
		e := dataproc.ClusterStatusHistoryStateEnum(n[len("DataprocClusterStatusHistoryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistorySubstateEnum converts a ClusterStatusHistorySubstateEnum enum from its proto representation.
func ProtoToDataprocClusterStatusHistorySubstateEnum(e dataprocpb.DataprocClusterStatusHistorySubstateEnum) *dataproc.ClusterStatusHistorySubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterStatusHistorySubstateEnum_name[int32(e)]; ok {
		e := dataproc.ClusterStatusHistorySubstateEnum(n[len("DataprocClusterStatusHistorySubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterConfig converts a ClusterClusterConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfig(p *dataprocpb.DataprocClusterClusterConfig) *dataproc.ClusterClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocClusterClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocClusterInstanceGroupConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocClusterInstanceGroupConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocClusterInstanceGroupConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocClusterClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocClusterClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocClusterClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocClusterClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocClusterClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocClusterClusterConfigEndpointConfig(p.GetEndpointConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocClusterClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfig converts a ClusterClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfig(p *dataprocpb.DataprocClusterClusterConfigGceClusterConfig) *dataproc.ClusterClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:     ProtoToDataprocClusterClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
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
func ProtoToDataprocClusterClusterConfigGceClusterConfigReservationAffinity(p *dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinity) *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity(p *dataprocpb.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity) *dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfig converts a ClusterInstanceGroupConfig object from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfig(p *dataprocpb.DataprocClusterInstanceGroupConfig) *dataproc.ClusterInstanceGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocClusterInstanceGroupConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocClusterInstanceGroupConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocClusterInstanceGroupConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocClusterInstanceGroupConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigDiskConfig converts a ClusterInstanceGroupConfigDiskConfig object from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigDiskConfig(p *dataprocpb.DataprocClusterInstanceGroupConfigDiskConfig) *dataproc.ClusterInstanceGroupConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigManagedGroupConfig converts a ClusterInstanceGroupConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigManagedGroupConfig(p *dataprocpb.DataprocClusterInstanceGroupConfigManagedGroupConfig) *dataproc.ClusterInstanceGroupConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigAccelerators converts a ClusterInstanceGroupConfigAccelerators object from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigAccelerators(p *dataprocpb.DataprocClusterInstanceGroupConfigAccelerators) *dataproc.ClusterInstanceGroupConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterClusterConfigSoftwareConfig converts a ClusterClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigSoftwareConfig(p *dataprocpb.DataprocClusterClusterConfigSoftwareConfig) *dataproc.ClusterClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterClusterConfigInitializationActions converts a ClusterClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocClusterClusterConfigInitializationActions(p *dataprocpb.DataprocClusterClusterConfigInitializationActions) *dataproc.ClusterClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToClusterClusterConfigEncryptionConfig converts a ClusterClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigEncryptionConfig(p *dataprocpb.DataprocClusterClusterConfigEncryptionConfig) *dataproc.ClusterClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToClusterClusterConfigAutoscalingConfig converts a ClusterClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigAutoscalingConfig(p *dataprocpb.DataprocClusterClusterConfigAutoscalingConfig) *dataproc.ClusterClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfig converts a ClusterClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigSecurityConfig(p *dataprocpb.DataprocClusterClusterConfigSecurityConfig) *dataproc.ClusterClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocClusterClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfigKerberosConfig converts a ClusterClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigSecurityConfigKerberosConfig(p *dataprocpb.DataprocClusterClusterConfigSecurityConfigKerberosConfig) *dataproc.ClusterClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocClusterClusterConfigLifecycleConfig(p *dataprocpb.DataprocClusterClusterConfigLifecycleConfig) *dataproc.ClusterClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterClusterConfigEndpointConfig converts a ClusterClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocClusterClusterConfigEndpointConfig(p *dataprocpb.DataprocClusterClusterConfigEndpointConfig) *dataproc.ClusterClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus object from its proto representation.
func ProtoToDataprocClusterStatus(p *dataprocpb.DataprocClusterStatus) *dataproc.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterStatus{
		State:          ProtoToDataprocClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory object from its proto representation.
func ProtoToDataprocClusterStatusHistory(p *dataprocpb.DataprocClusterStatusHistory) *dataproc.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterStatusHistory{
		State:          ProtoToDataprocClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics object from its proto representation.
func ProtoToDataprocClusterMetrics(p *dataprocpb.DataprocClusterMetrics) *dataproc.ClusterMetrics {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterMetrics{}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *dataprocpb.DataprocCluster) *dataproc.Cluster {
	obj := &dataproc.Cluster{
		Project:     dcl.StringOrNil(p.GetProject()),
		Name:        dcl.StringOrNil(p.GetName()),
		Config:      ProtoToDataprocClusterClusterConfig(p.GetConfig()),
		Status:      ProtoToDataprocClusterStatus(p.GetStatus()),
		ClusterUuid: dcl.StringOrNil(p.GetClusterUuid()),
		Metrics:     ProtoToDataprocClusterMetrics(p.GetMetrics()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetStatusHistory() {
		obj.StatusHistory = append(obj.StatusHistory, *ProtoToDataprocClusterStatusHistory(r))
	}
	return obj
}

// ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *dataproc.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return dataprocpb.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterInstanceGroupConfigPreemptibilityEnumToProto converts a ClusterInstanceGroupConfigPreemptibilityEnum enum to its proto representation.
func DataprocClusterInstanceGroupConfigPreemptibilityEnumToProto(e *dataproc.ClusterInstanceGroupConfigPreemptibilityEnum) dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum_value["ClusterInstanceGroupConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocClusterInstanceGroupConfigPreemptibilityEnum(0)
}

// ClusterClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a ClusterClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *dataproc.ClusterClusterConfigSoftwareConfigOptionalComponentsEnum) dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value["ClusterClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// ClusterStatusStateEnumToProto converts a ClusterStatusStateEnum enum to its proto representation.
func DataprocClusterStatusStateEnumToProto(e *dataproc.ClusterStatusStateEnum) dataprocpb.DataprocClusterStatusStateEnum {
	if e == nil {
		return dataprocpb.DataprocClusterStatusStateEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterStatusStateEnum_value["ClusterStatusStateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterStatusStateEnum(v)
	}
	return dataprocpb.DataprocClusterStatusStateEnum(0)
}

// ClusterStatusSubstateEnumToProto converts a ClusterStatusSubstateEnum enum to its proto representation.
func DataprocClusterStatusSubstateEnumToProto(e *dataproc.ClusterStatusSubstateEnum) dataprocpb.DataprocClusterStatusSubstateEnum {
	if e == nil {
		return dataprocpb.DataprocClusterStatusSubstateEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterStatusSubstateEnum_value["ClusterStatusSubstateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterStatusSubstateEnum(v)
	}
	return dataprocpb.DataprocClusterStatusSubstateEnum(0)
}

// ClusterStatusHistoryStateEnumToProto converts a ClusterStatusHistoryStateEnum enum to its proto representation.
func DataprocClusterStatusHistoryStateEnumToProto(e *dataproc.ClusterStatusHistoryStateEnum) dataprocpb.DataprocClusterStatusHistoryStateEnum {
	if e == nil {
		return dataprocpb.DataprocClusterStatusHistoryStateEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterStatusHistoryStateEnum_value["ClusterStatusHistoryStateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterStatusHistoryStateEnum(v)
	}
	return dataprocpb.DataprocClusterStatusHistoryStateEnum(0)
}

// ClusterStatusHistorySubstateEnumToProto converts a ClusterStatusHistorySubstateEnum enum to its proto representation.
func DataprocClusterStatusHistorySubstateEnumToProto(e *dataproc.ClusterStatusHistorySubstateEnum) dataprocpb.DataprocClusterStatusHistorySubstateEnum {
	if e == nil {
		return dataprocpb.DataprocClusterStatusHistorySubstateEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterStatusHistorySubstateEnum_value["ClusterStatusHistorySubstateEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterStatusHistorySubstateEnum(v)
	}
	return dataprocpb.DataprocClusterStatusHistorySubstateEnum(0)
}

// ClusterClusterConfigToProto converts a ClusterClusterConfig object to its proto representation.
func DataprocClusterClusterConfigToProto(o *dataproc.ClusterClusterConfig) *dataprocpb.DataprocClusterClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocClusterClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocClusterInstanceGroupConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocClusterInstanceGroupConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocClusterInstanceGroupConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocClusterClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocClusterClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocClusterClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocClusterClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocClusterClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocClusterClusterConfigEndpointConfigToProto(o.EndpointConfig))
	sInitializationActions := make([]*dataprocpb.DataprocClusterClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocClusterClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// ClusterClusterConfigGceClusterConfigToProto converts a ClusterClusterConfigGceClusterConfig object to its proto representation.
func DataprocClusterClusterConfigGceClusterConfigToProto(o *dataproc.ClusterClusterConfigGceClusterConfig) *dataprocpb.DataprocClusterClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocClusterClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
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
func DataprocClusterClusterConfigGceClusterConfigReservationAffinityToProto(o *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity) *dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// ClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *dataprocpb.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// ClusterInstanceGroupConfigToProto converts a ClusterInstanceGroupConfig object to its proto representation.
func DataprocClusterInstanceGroupConfigToProto(o *dataproc.ClusterInstanceGroupConfig) *dataprocpb.DataprocClusterInstanceGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocClusterInstanceGroupConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocClusterInstanceGroupConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocClusterInstanceGroupConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocClusterInstanceGroupConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocClusterInstanceGroupConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterInstanceGroupConfigDiskConfigToProto converts a ClusterInstanceGroupConfigDiskConfig object to its proto representation.
func DataprocClusterInstanceGroupConfigDiskConfigToProto(o *dataproc.ClusterInstanceGroupConfigDiskConfig) *dataprocpb.DataprocClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterInstanceGroupConfigManagedGroupConfigToProto converts a ClusterInstanceGroupConfigManagedGroupConfig object to its proto representation.
func DataprocClusterInstanceGroupConfigManagedGroupConfigToProto(o *dataproc.ClusterInstanceGroupConfigManagedGroupConfig) *dataprocpb.DataprocClusterInstanceGroupConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterInstanceGroupConfigAcceleratorsToProto converts a ClusterInstanceGroupConfigAccelerators object to its proto representation.
func DataprocClusterInstanceGroupConfigAcceleratorsToProto(o *dataproc.ClusterInstanceGroupConfigAccelerators) *dataprocpb.DataprocClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterClusterConfigSoftwareConfigToProto converts a ClusterClusterConfigSoftwareConfig object to its proto representation.
func DataprocClusterClusterConfigSoftwareConfigToProto(o *dataproc.ClusterClusterConfigSoftwareConfig) *dataprocpb.DataprocClusterClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// ClusterClusterConfigInitializationActionsToProto converts a ClusterClusterConfigInitializationActions object to its proto representation.
func DataprocClusterClusterConfigInitializationActionsToProto(o *dataproc.ClusterClusterConfigInitializationActions) *dataprocpb.DataprocClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// ClusterClusterConfigEncryptionConfigToProto converts a ClusterClusterConfigEncryptionConfig object to its proto representation.
func DataprocClusterClusterConfigEncryptionConfigToProto(o *dataproc.ClusterClusterConfigEncryptionConfig) *dataprocpb.DataprocClusterClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// ClusterClusterConfigAutoscalingConfigToProto converts a ClusterClusterConfigAutoscalingConfig object to its proto representation.
func DataprocClusterClusterConfigAutoscalingConfigToProto(o *dataproc.ClusterClusterConfigAutoscalingConfig) *dataprocpb.DataprocClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// ClusterClusterConfigSecurityConfigToProto converts a ClusterClusterConfigSecurityConfig object to its proto representation.
func DataprocClusterClusterConfigSecurityConfigToProto(o *dataproc.ClusterClusterConfigSecurityConfig) *dataprocpb.DataprocClusterClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocClusterClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// ClusterClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocClusterClusterConfigSecurityConfigKerberosConfigToProto(o *dataproc.ClusterClusterConfigSecurityConfigKerberosConfig) *dataprocpb.DataprocClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigSecurityConfigKerberosConfig{}
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
func DataprocClusterClusterConfigLifecycleConfigToProto(o *dataproc.ClusterClusterConfigLifecycleConfig) *dataprocpb.DataprocClusterClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// ClusterClusterConfigEndpointConfigToProto converts a ClusterClusterConfigEndpointConfig object to its proto representation.
func DataprocClusterClusterConfigEndpointConfigToProto(o *dataproc.ClusterClusterConfigEndpointConfig) *dataprocpb.DataprocClusterClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// ClusterStatusToProto converts a ClusterStatus object to its proto representation.
func DataprocClusterStatusToProto(o *dataproc.ClusterStatus) *dataprocpb.DataprocClusterStatus {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterStatus{}
	p.SetState(DataprocClusterStatusStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocClusterStatusSubstateEnumToProto(o.Substate))
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory object to its proto representation.
func DataprocClusterStatusHistoryToProto(o *dataproc.ClusterStatusHistory) *dataprocpb.DataprocClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterStatusHistory{}
	p.SetState(DataprocClusterStatusHistoryStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocClusterStatusHistorySubstateEnumToProto(o.Substate))
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics object to its proto representation.
func DataprocClusterMetricsToProto(o *dataproc.ClusterMetrics) *dataprocpb.DataprocClusterMetrics {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterMetrics{}
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
func ClusterToProto(resource *dataproc.Cluster) *dataprocpb.DataprocCluster {
	p := &dataprocpb.DataprocCluster{}
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetConfig(DataprocClusterClusterConfigToProto(resource.Config))
	p.SetStatus(DataprocClusterStatusToProto(resource.Status))
	p.SetClusterUuid(dcl.ValueOrEmptyString(resource.ClusterUuid))
	p.SetMetrics(DataprocClusterMetricsToProto(resource.Metrics))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStatusHistory := make([]*dataprocpb.DataprocClusterStatusHistory, len(resource.StatusHistory))
	for i, r := range resource.StatusHistory {
		sStatusHistory[i] = DataprocClusterStatusHistoryToProto(&r)
	}
	p.SetStatusHistory(sStatusHistory)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocClusterRequest) (*dataprocpb.DataprocCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyDataprocCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocCluster(ctx context.Context, request *dataprocpb.ApplyDataprocClusterRequest) (*dataprocpb.DataprocCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocCluster(ctx context.Context, request *dataprocpb.DeleteDataprocClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocCluster(ctx context.Context, request *dataprocpb.ListDataprocClusterRequest) (*dataprocpb.ListDataprocClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &dataprocpb.ListDataprocClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
