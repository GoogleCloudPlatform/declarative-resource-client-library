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

// ProtoToClusterClusterConfig converts a ClusterClusterConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfig(p *dataprocpb.DataprocClusterClusterConfig) *dataproc.ClusterClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.StagingBucket),
		TempBucket:            dcl.StringOrNil(p.TempBucket),
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

// ProtoToClusterClusterConfigGceClusterConfig converts a ClusterClusterConfigGceClusterConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfig(p *dataprocpb.DataprocClusterClusterConfigGceClusterConfig) *dataproc.ClusterClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.Zone),
		Network:                 dcl.StringOrNil(p.Network),
		Subnetwork:              dcl.StringOrNil(p.Subnetwork),
		InternalIPOnly:          dcl.Bool(p.InternalIpOnly),
		PrivateIPv6GoogleAccess: ProtoToDataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.ServiceAccount),
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

// ProtoToClusterClusterConfigGceClusterConfigReservationAffinity converts a ClusterClusterConfigGceClusterConfigReservationAffinity resource from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfigReservationAffinity(p *dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinity) *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity resource from its proto representation.
func ProtoToDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity(p *dataprocpb.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity) *dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.NodeGroup),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfig converts a ClusterInstanceGroupConfig resource from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfig(p *dataprocpb.DataprocClusterInstanceGroupConfig) *dataproc.ClusterInstanceGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfig{
		NumInstances:       dcl.Int64OrNil(p.NumInstances),
		Image:              dcl.StringOrNil(p.Image),
		MachineType:        dcl.StringOrNil(p.MachineType),
		DiskConfig:         ProtoToDataprocClusterInstanceGroupConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.IsPreemptible),
		Preemptibility:     ProtoToDataprocClusterInstanceGroupConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocClusterInstanceGroupConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.MinCpuPlatform),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocClusterInstanceGroupConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigDiskConfig converts a ClusterInstanceGroupConfigDiskConfig resource from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigDiskConfig(p *dataprocpb.DataprocClusterInstanceGroupConfigDiskConfig) *dataproc.ClusterInstanceGroupConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.BootDiskType),
		BootDiskSizeGb: dcl.Int64OrNil(p.BootDiskSizeGb),
		NumLocalSsds:   dcl.Int64OrNil(p.NumLocalSsds),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigManagedGroupConfig converts a ClusterInstanceGroupConfigManagedGroupConfig resource from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigManagedGroupConfig(p *dataprocpb.DataprocClusterInstanceGroupConfigManagedGroupConfig) *dataproc.ClusterInstanceGroupConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.InstanceTemplateName),
		InstanceGroupManagerName: dcl.StringOrNil(p.InstanceGroupManagerName),
	}
	return obj
}

// ProtoToClusterInstanceGroupConfigAccelerators converts a ClusterInstanceGroupConfigAccelerators resource from its proto representation.
func ProtoToDataprocClusterInstanceGroupConfigAccelerators(p *dataprocpb.DataprocClusterInstanceGroupConfigAccelerators) *dataproc.ClusterInstanceGroupConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
	}
	return obj
}

// ProtoToClusterClusterConfigSoftwareConfig converts a ClusterClusterConfigSoftwareConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigSoftwareConfig(p *dataprocpb.DataprocClusterClusterConfigSoftwareConfig) *dataproc.ClusterClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.ImageVersion),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterClusterConfigInitializationActions converts a ClusterClusterConfigInitializationActions resource from its proto representation.
func ProtoToDataprocClusterClusterConfigInitializationActions(p *dataprocpb.DataprocClusterClusterConfigInitializationActions) *dataproc.ClusterClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.ExecutableFile),
		ExecutionTimeout: dcl.StringOrNil(p.ExecutionTimeout),
	}
	return obj
}

// ProtoToClusterClusterConfigEncryptionConfig converts a ClusterClusterConfigEncryptionConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigEncryptionConfig(p *dataprocpb.DataprocClusterClusterConfigEncryptionConfig) *dataproc.ClusterClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GcePdKmsKeyName),
	}
	return obj
}

// ProtoToClusterClusterConfigAutoscalingConfig converts a ClusterClusterConfigAutoscalingConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigAutoscalingConfig(p *dataprocpb.DataprocClusterClusterConfigAutoscalingConfig) *dataproc.ClusterClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.Policy),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfig converts a ClusterClusterConfigSecurityConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigSecurityConfig(p *dataprocpb.DataprocClusterClusterConfigSecurityConfig) *dataproc.ClusterClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocClusterClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterClusterConfigSecurityConfigKerberosConfig converts a ClusterClusterConfigSecurityConfigKerberosConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigSecurityConfigKerberosConfig(p *dataprocpb.DataprocClusterClusterConfigSecurityConfigKerberosConfig) *dataproc.ClusterClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocClusterClusterConfigLifecycleConfig(p *dataprocpb.DataprocClusterClusterConfigLifecycleConfig) *dataproc.ClusterClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.IdleDeleteTtl),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.AutoDeleteTtl),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterClusterConfigEndpointConfig converts a ClusterClusterConfigEndpointConfig resource from its proto representation.
func ProtoToDataprocClusterClusterConfigEndpointConfig(p *dataprocpb.DataprocClusterClusterConfigEndpointConfig) *dataproc.ClusterClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.EnableHttpPortAccess),
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus resource from its proto representation.
func ProtoToDataprocClusterStatus(p *dataprocpb.DataprocClusterStatus) *dataproc.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterStatus{
		State:          ProtoToDataprocClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.Detail),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory resource from its proto representation.
func ProtoToDataprocClusterStatusHistory(p *dataprocpb.DataprocClusterStatusHistory) *dataproc.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterStatusHistory{
		State:          ProtoToDataprocClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.Detail),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics resource from its proto representation.
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
		Project:     dcl.StringOrNil(p.Project),
		Name:        dcl.StringOrNil(p.Name),
		Config:      ProtoToDataprocClusterClusterConfig(p.GetConfig()),
		Status:      ProtoToDataprocClusterStatus(p.GetStatus()),
		ClusterUuid: dcl.StringOrNil(p.ClusterUuid),
		Metrics:     ProtoToDataprocClusterMetrics(p.GetMetrics()),
		Location:    dcl.StringOrNil(p.Location),
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

// ClusterClusterConfigToProto converts a ClusterClusterConfig resource to its proto representation.
func DataprocClusterClusterConfigToProto(o *dataproc.ClusterClusterConfig) *dataprocpb.DataprocClusterClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfig{
		StagingBucket:         dcl.ValueOrEmptyString(o.StagingBucket),
		TempBucket:            dcl.ValueOrEmptyString(o.TempBucket),
		GceClusterConfig:      DataprocClusterClusterConfigGceClusterConfigToProto(o.GceClusterConfig),
		MasterConfig:          DataprocClusterInstanceGroupConfigToProto(o.MasterConfig),
		WorkerConfig:          DataprocClusterInstanceGroupConfigToProto(o.WorkerConfig),
		SecondaryWorkerConfig: DataprocClusterInstanceGroupConfigToProto(o.SecondaryWorkerConfig),
		SoftwareConfig:        DataprocClusterClusterConfigSoftwareConfigToProto(o.SoftwareConfig),
		EncryptionConfig:      DataprocClusterClusterConfigEncryptionConfigToProto(o.EncryptionConfig),
		AutoscalingConfig:     DataprocClusterClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig),
		SecurityConfig:        DataprocClusterClusterConfigSecurityConfigToProto(o.SecurityConfig),
		LifecycleConfig:       DataprocClusterClusterConfigLifecycleConfigToProto(o.LifecycleConfig),
		EndpointConfig:        DataprocClusterClusterConfigEndpointConfigToProto(o.EndpointConfig),
	}
	for _, r := range o.InitializationActions {
		p.InitializationActions = append(p.InitializationActions, DataprocClusterClusterConfigInitializationActionsToProto(&r))
	}
	return p
}

// ClusterClusterConfigGceClusterConfigToProto converts a ClusterClusterConfigGceClusterConfig resource to its proto representation.
func DataprocClusterClusterConfigGceClusterConfigToProto(o *dataproc.ClusterClusterConfigGceClusterConfig) *dataprocpb.DataprocClusterClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigGceClusterConfig{
		Zone:                    dcl.ValueOrEmptyString(o.Zone),
		Network:                 dcl.ValueOrEmptyString(o.Network),
		Subnetwork:              dcl.ValueOrEmptyString(o.Subnetwork),
		InternalIpOnly:          dcl.ValueOrEmptyBool(o.InternalIPOnly),
		PrivateIpv6GoogleAccess: DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess),
		ServiceAccount:          dcl.ValueOrEmptyString(o.ServiceAccount),
		ReservationAffinity:     DataprocClusterClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity),
		NodeGroupAffinity:       DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity),
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
func DataprocClusterClusterConfigGceClusterConfigReservationAffinityToProto(o *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity) *dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterClusterConfigGceClusterConfigNodeGroupAffinity resource to its proto representation.
func DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) *dataprocpb.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.ValueOrEmptyString(o.NodeGroup),
	}
	return p
}

// ClusterInstanceGroupConfigToProto converts a ClusterInstanceGroupConfig resource to its proto representation.
func DataprocClusterInstanceGroupConfigToProto(o *dataproc.ClusterInstanceGroupConfig) *dataprocpb.DataprocClusterInstanceGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfig{
		NumInstances:       dcl.ValueOrEmptyInt64(o.NumInstances),
		Image:              dcl.ValueOrEmptyString(o.Image),
		MachineType:        dcl.ValueOrEmptyString(o.MachineType),
		DiskConfig:         DataprocClusterInstanceGroupConfigDiskConfigToProto(o.DiskConfig),
		IsPreemptible:      dcl.ValueOrEmptyBool(o.IsPreemptible),
		Preemptibility:     DataprocClusterInstanceGroupConfigPreemptibilityEnumToProto(o.Preemptibility),
		ManagedGroupConfig: DataprocClusterInstanceGroupConfigManagedGroupConfigToProto(o.ManagedGroupConfig),
		MinCpuPlatform:     dcl.ValueOrEmptyString(o.MinCpuPlatform),
	}
	for _, r := range o.InstanceNames {
		p.InstanceNames = append(p.InstanceNames, r)
	}
	for _, r := range o.Accelerators {
		p.Accelerators = append(p.Accelerators, DataprocClusterInstanceGroupConfigAcceleratorsToProto(&r))
	}
	return p
}

// ClusterInstanceGroupConfigDiskConfigToProto converts a ClusterInstanceGroupConfigDiskConfig resource to its proto representation.
func DataprocClusterInstanceGroupConfigDiskConfigToProto(o *dataproc.ClusterInstanceGroupConfigDiskConfig) *dataprocpb.DataprocClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfigDiskConfig{
		BootDiskType:   dcl.ValueOrEmptyString(o.BootDiskType),
		BootDiskSizeGb: dcl.ValueOrEmptyInt64(o.BootDiskSizeGb),
		NumLocalSsds:   dcl.ValueOrEmptyInt64(o.NumLocalSsds),
	}
	return p
}

// ClusterInstanceGroupConfigManagedGroupConfigToProto converts a ClusterInstanceGroupConfigManagedGroupConfig resource to its proto representation.
func DataprocClusterInstanceGroupConfigManagedGroupConfigToProto(o *dataproc.ClusterInstanceGroupConfigManagedGroupConfig) *dataprocpb.DataprocClusterInstanceGroupConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.ValueOrEmptyString(o.InstanceTemplateName),
		InstanceGroupManagerName: dcl.ValueOrEmptyString(o.InstanceGroupManagerName),
	}
	return p
}

// ClusterInstanceGroupConfigAcceleratorsToProto converts a ClusterInstanceGroupConfigAccelerators resource to its proto representation.
func DataprocClusterInstanceGroupConfigAcceleratorsToProto(o *dataproc.ClusterInstanceGroupConfigAccelerators) *dataprocpb.DataprocClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterInstanceGroupConfigAccelerators{
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
	}
	return p
}

// ClusterClusterConfigSoftwareConfigToProto converts a ClusterClusterConfigSoftwareConfig resource to its proto representation.
func DataprocClusterClusterConfigSoftwareConfigToProto(o *dataproc.ClusterClusterConfigSoftwareConfig) *dataprocpb.DataprocClusterClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.ValueOrEmptyString(o.ImageVersion),
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.OptionalComponents {
		p.OptionalComponents = append(p.OptionalComponents, dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum(dataprocpb.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)]))
	}
	return p
}

// ClusterClusterConfigInitializationActionsToProto converts a ClusterClusterConfigInitializationActions resource to its proto representation.
func DataprocClusterClusterConfigInitializationActionsToProto(o *dataproc.ClusterClusterConfigInitializationActions) *dataprocpb.DataprocClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.ValueOrEmptyString(o.ExecutableFile),
		ExecutionTimeout: dcl.ValueOrEmptyString(o.ExecutionTimeout),
	}
	return p
}

// ClusterClusterConfigEncryptionConfigToProto converts a ClusterClusterConfigEncryptionConfig resource to its proto representation.
func DataprocClusterClusterConfigEncryptionConfigToProto(o *dataproc.ClusterClusterConfigEncryptionConfig) *dataprocpb.DataprocClusterClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.ValueOrEmptyString(o.GcePdKmsKeyName),
	}
	return p
}

// ClusterClusterConfigAutoscalingConfigToProto converts a ClusterClusterConfigAutoscalingConfig resource to its proto representation.
func DataprocClusterClusterConfigAutoscalingConfigToProto(o *dataproc.ClusterClusterConfigAutoscalingConfig) *dataprocpb.DataprocClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigAutoscalingConfig{
		Policy: dcl.ValueOrEmptyString(o.Policy),
	}
	return p
}

// ClusterClusterConfigSecurityConfigToProto converts a ClusterClusterConfigSecurityConfig resource to its proto representation.
func DataprocClusterClusterConfigSecurityConfigToProto(o *dataproc.ClusterClusterConfigSecurityConfig) *dataprocpb.DataprocClusterClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigSecurityConfig{
		KerberosConfig: DataprocClusterClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig),
	}
	return p
}

// ClusterClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterClusterConfigSecurityConfigKerberosConfig resource to its proto representation.
func DataprocClusterClusterConfigSecurityConfigKerberosConfigToProto(o *dataproc.ClusterClusterConfigSecurityConfigKerberosConfig) *dataprocpb.DataprocClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigSecurityConfigKerberosConfig{
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
func DataprocClusterClusterConfigLifecycleConfigToProto(o *dataproc.ClusterClusterConfigLifecycleConfig) *dataprocpb.DataprocClusterClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.ValueOrEmptyString(o.IdleDeleteTtl),
		AutoDeleteTime: dcl.ValueOrEmptyString(o.AutoDeleteTime),
		AutoDeleteTtl:  dcl.ValueOrEmptyString(o.AutoDeleteTtl),
		IdleStartTime:  dcl.ValueOrEmptyString(o.IdleStartTime),
	}
	return p
}

// ClusterClusterConfigEndpointConfigToProto converts a ClusterClusterConfigEndpointConfig resource to its proto representation.
func DataprocClusterClusterConfigEndpointConfigToProto(o *dataproc.ClusterClusterConfigEndpointConfig) *dataprocpb.DataprocClusterClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.ValueOrEmptyBool(o.EnableHttpPortAccess),
	}
	p.HttpPorts = make(map[string]string)
	for k, r := range o.HttpPorts {
		p.HttpPorts[k] = r
	}
	return p
}

// ClusterStatusToProto converts a ClusterStatus resource to its proto representation.
func DataprocClusterStatusToProto(o *dataproc.ClusterStatus) *dataprocpb.DataprocClusterStatus {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterStatus{
		State:          DataprocClusterStatusStateEnumToProto(o.State),
		Detail:         dcl.ValueOrEmptyString(o.Detail),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocClusterStatusSubstateEnumToProto(o.Substate),
	}
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory resource to its proto representation.
func DataprocClusterStatusHistoryToProto(o *dataproc.ClusterStatusHistory) *dataprocpb.DataprocClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterStatusHistory{
		State:          DataprocClusterStatusHistoryStateEnumToProto(o.State),
		Detail:         dcl.ValueOrEmptyString(o.Detail),
		StateStartTime: dcl.ValueOrEmptyString(o.StateStartTime),
		Substate:       DataprocClusterStatusHistorySubstateEnumToProto(o.Substate),
	}
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics resource to its proto representation.
func DataprocClusterMetricsToProto(o *dataproc.ClusterMetrics) *dataprocpb.DataprocClusterMetrics {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterMetrics{}
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
func ClusterToProto(resource *dataproc.Cluster) *dataprocpb.DataprocCluster {
	p := &dataprocpb.DataprocCluster{
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Config:      DataprocClusterClusterConfigToProto(resource.Config),
		Status:      DataprocClusterStatusToProto(resource.Status),
		ClusterUuid: dcl.ValueOrEmptyString(resource.ClusterUuid),
		Metrics:     DataprocClusterMetricsToProto(resource.Metrics),
		Location:    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.StatusHistory {
		p.StatusHistory = append(p.StatusHistory, DataprocClusterStatusHistoryToProto(&r))
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocClusterRequest) (*dataprocpb.DataprocCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocCluster(ctx context.Context, request *dataprocpb.ApplyDataprocClusterRequest) (*dataprocpb.DataprocCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocCluster(ctx context.Context, request *dataprocpb.DeleteDataprocClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocCluster(ctx context.Context, request *dataprocpb.ListDataprocClusterRequest) (*dataprocpb.ListDataprocClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, ProtoToCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &dataprocpb.ListDataprocClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
