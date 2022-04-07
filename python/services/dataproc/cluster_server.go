// Copyright 2022 Google LLC. All Rights Reserved.
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

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *dataproc.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := dataproc.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *dataproc.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := dataproc.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigMasterConfigPreemptibilityEnum converts a ClusterConfigMasterConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocClusterConfigMasterConfigPreemptibilityEnum(e dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum) *dataproc.ClusterConfigMasterConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.ClusterConfigMasterConfigPreemptibilityEnum(n[len("DataprocClusterConfigMasterConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigWorkerConfigPreemptibilityEnum converts a ClusterConfigWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocClusterConfigWorkerConfigPreemptibilityEnum(e dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum) *dataproc.ClusterConfigWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.ClusterConfigWorkerConfigPreemptibilityEnum(n[len("DataprocClusterConfigWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigSecondaryWorkerConfigPreemptibilityEnum converts a ClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum(e dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum) *dataproc.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum(n[len("DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigSoftwareConfigOptionalComponentsEnum converts a ClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocClusterConfigSoftwareConfigOptionalComponentsEnum(e dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum) *dataproc.ClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := dataproc.ClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocClusterConfigSoftwareConfigOptionalComponentsEnum"):])
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

// ProtoToClusterConfig converts a ClusterConfig object from its proto representation.
func ProtoToDataprocClusterConfig(p *dataprocpb.DataprocClusterConfig) *dataproc.ClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocClusterConfigMasterConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocClusterConfigWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocClusterConfigSecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocClusterConfigEndpointConfig(p.GetEndpointConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfig converts a ClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocClusterConfigGceClusterConfig(p *dataprocpb.DataprocClusterConfigGceClusterConfig) *dataproc.ClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:     ProtoToDataprocClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
	}
	for _, r := range p.GetServiceAccountScopes() {
		obj.ServiceAccountScopes = append(obj.ServiceAccountScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigReservationAffinity converts a ClusterConfigGceClusterConfigReservationAffinity object from its proto representation.
func ProtoToDataprocClusterConfigGceClusterConfigReservationAffinity(p *dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinity) *dataproc.ClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocClusterConfigGceClusterConfigNodeGroupAffinity(p *dataprocpb.DataprocClusterConfigGceClusterConfigNodeGroupAffinity) *dataproc.ClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfig converts a ClusterConfigMasterConfig object from its proto representation.
func ProtoToDataprocClusterConfigMasterConfig(p *dataprocpb.DataprocClusterConfigMasterConfig) *dataproc.ClusterConfigMasterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigMasterConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocClusterConfigMasterConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocClusterConfigMasterConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocClusterConfigMasterConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocClusterConfigMasterConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterConfigMasterConfigDiskConfig converts a ClusterConfigMasterConfigDiskConfig object from its proto representation.
func ProtoToDataprocClusterConfigMasterConfigDiskConfig(p *dataprocpb.DataprocClusterConfigMasterConfigDiskConfig) *dataproc.ClusterConfigMasterConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigMasterConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigManagedGroupConfig converts a ClusterConfigMasterConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocClusterConfigMasterConfigManagedGroupConfig(p *dataprocpb.DataprocClusterConfigMasterConfigManagedGroupConfig) *dataproc.ClusterConfigMasterConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigMasterConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigAccelerators converts a ClusterConfigMasterConfigAccelerators object from its proto representation.
func ProtoToDataprocClusterConfigMasterConfigAccelerators(p *dataprocpb.DataprocClusterConfigMasterConfigAccelerators) *dataproc.ClusterConfigMasterConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigMasterConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfig converts a ClusterConfigWorkerConfig object from its proto representation.
func ProtoToDataprocClusterConfigWorkerConfig(p *dataprocpb.DataprocClusterConfigWorkerConfig) *dataproc.ClusterConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocClusterConfigWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocClusterConfigWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocClusterConfigWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocClusterConfigWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigDiskConfig converts a ClusterConfigWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocClusterConfigWorkerConfigDiskConfig(p *dataprocpb.DataprocClusterConfigWorkerConfigDiskConfig) *dataproc.ClusterConfigWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigManagedGroupConfig converts a ClusterConfigWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocClusterConfigWorkerConfigManagedGroupConfig(p *dataprocpb.DataprocClusterConfigWorkerConfigManagedGroupConfig) *dataproc.ClusterConfigWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigAccelerators converts a ClusterConfigWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocClusterConfigWorkerConfigAccelerators(p *dataprocpb.DataprocClusterConfigWorkerConfigAccelerators) *dataproc.ClusterConfigWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfig converts a ClusterConfigSecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocClusterConfigSecondaryWorkerConfig(p *dataprocpb.DataprocClusterConfigSecondaryWorkerConfig) *dataproc.ClusterConfigSecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSecondaryWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocClusterConfigSecondaryWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocClusterConfigSecondaryWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocClusterConfigSecondaryWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigDiskConfig converts a ClusterConfigSecondaryWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocClusterConfigSecondaryWorkerConfigDiskConfig(p *dataprocpb.DataprocClusterConfigSecondaryWorkerConfigDiskConfig) *dataproc.ClusterConfigSecondaryWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSecondaryWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigManagedGroupConfig converts a ClusterConfigSecondaryWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocClusterConfigSecondaryWorkerConfigManagedGroupConfig(p *dataprocpb.DataprocClusterConfigSecondaryWorkerConfigManagedGroupConfig) *dataproc.ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSecondaryWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigAccelerators converts a ClusterConfigSecondaryWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocClusterConfigSecondaryWorkerConfigAccelerators(p *dataprocpb.DataprocClusterConfigSecondaryWorkerConfigAccelerators) *dataproc.ClusterConfigSecondaryWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSecondaryWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigSoftwareConfig converts a ClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocClusterConfigSoftwareConfig(p *dataprocpb.DataprocClusterConfigSoftwareConfig) *dataproc.ClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterConfigInitializationActions converts a ClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocClusterConfigInitializationActions(p *dataprocpb.DataprocClusterConfigInitializationActions) *dataproc.ClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToClusterConfigEncryptionConfig converts a ClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocClusterConfigEncryptionConfig(p *dataprocpb.DataprocClusterConfigEncryptionConfig) *dataproc.ClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToClusterConfigAutoscalingConfig converts a ClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocClusterConfigAutoscalingConfig(p *dataprocpb.DataprocClusterConfigAutoscalingConfig) *dataproc.ClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfig converts a ClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocClusterConfigSecurityConfig(p *dataprocpb.DataprocClusterConfigSecurityConfig) *dataproc.ClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfigKerberosConfig converts a ClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocClusterConfigSecurityConfigKerberosConfig(p *dataprocpb.DataprocClusterConfigSecurityConfigKerberosConfig) *dataproc.ClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigSecurityConfigKerberosConfig{
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

// ProtoToClusterConfigLifecycleConfig converts a ClusterConfigLifecycleConfig object from its proto representation.
func ProtoToDataprocClusterConfigLifecycleConfig(p *dataprocpb.DataprocClusterConfigLifecycleConfig) *dataproc.ClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterConfigEndpointConfig converts a ClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocClusterConfigEndpointConfig(p *dataprocpb.DataprocClusterConfigEndpointConfig) *dataproc.ClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.ClusterConfigEndpointConfig{
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
		Config:      ProtoToDataprocClusterConfig(p.GetConfig()),
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

// ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *dataproc.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return dataprocpb.DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *dataproc.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterConfigMasterConfigPreemptibilityEnumToProto converts a ClusterConfigMasterConfigPreemptibilityEnum enum to its proto representation.
func DataprocClusterConfigMasterConfigPreemptibilityEnumToProto(e *dataproc.ClusterConfigMasterConfigPreemptibilityEnum) dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum_value["ClusterConfigMasterConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocClusterConfigMasterConfigPreemptibilityEnum(0)
}

// ClusterConfigWorkerConfigPreemptibilityEnumToProto converts a ClusterConfigWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocClusterConfigWorkerConfigPreemptibilityEnumToProto(e *dataproc.ClusterConfigWorkerConfigPreemptibilityEnum) dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum_value["ClusterConfigWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocClusterConfigWorkerConfigPreemptibilityEnum(0)
}

// ClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto converts a ClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(e *dataproc.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum) dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum_value["ClusterConfigSecondaryWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
}

// ClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a ClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *dataproc.ClusterConfigSoftwareConfigOptionalComponentsEnum) dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum_value["ClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum(0)
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

// ClusterConfigToProto converts a ClusterConfig object to its proto representation.
func DataprocClusterConfigToProto(o *dataproc.ClusterConfig) *dataprocpb.DataprocClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocClusterConfigMasterConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocClusterConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocClusterConfigSecondaryWorkerConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocClusterConfigEndpointConfigToProto(o.EndpointConfig))
	sInitializationActions := make([]*dataprocpb.DataprocClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// ClusterConfigGceClusterConfigToProto converts a ClusterConfigGceClusterConfig object to its proto representation.
func DataprocClusterConfigGceClusterConfigToProto(o *dataproc.ClusterConfigGceClusterConfig) *dataprocpb.DataprocClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
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

// ClusterConfigGceClusterConfigReservationAffinityToProto converts a ClusterConfigGceClusterConfigReservationAffinity object to its proto representation.
func DataprocClusterConfigGceClusterConfigReservationAffinityToProto(o *dataproc.ClusterConfigGceClusterConfigReservationAffinity) *dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// ClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *dataproc.ClusterConfigGceClusterConfigNodeGroupAffinity) *dataprocpb.DataprocClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// ClusterConfigMasterConfigToProto converts a ClusterConfigMasterConfig object to its proto representation.
func DataprocClusterConfigMasterConfigToProto(o *dataproc.ClusterConfigMasterConfig) *dataprocpb.DataprocClusterConfigMasterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigMasterConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocClusterConfigMasterConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocClusterConfigMasterConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocClusterConfigMasterConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocClusterConfigMasterConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocClusterConfigMasterConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterConfigMasterConfigDiskConfigToProto converts a ClusterConfigMasterConfigDiskConfig object to its proto representation.
func DataprocClusterConfigMasterConfigDiskConfigToProto(o *dataproc.ClusterConfigMasterConfigDiskConfig) *dataprocpb.DataprocClusterConfigMasterConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigMasterConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterConfigMasterConfigManagedGroupConfigToProto converts a ClusterConfigMasterConfigManagedGroupConfig object to its proto representation.
func DataprocClusterConfigMasterConfigManagedGroupConfigToProto(o *dataproc.ClusterConfigMasterConfigManagedGroupConfig) *dataprocpb.DataprocClusterConfigMasterConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigMasterConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigMasterConfigAcceleratorsToProto converts a ClusterConfigMasterConfigAccelerators object to its proto representation.
func DataprocClusterConfigMasterConfigAcceleratorsToProto(o *dataproc.ClusterConfigMasterConfigAccelerators) *dataprocpb.DataprocClusterConfigMasterConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigMasterConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigWorkerConfigToProto converts a ClusterConfigWorkerConfig object to its proto representation.
func DataprocClusterConfigWorkerConfigToProto(o *dataproc.ClusterConfigWorkerConfig) *dataprocpb.DataprocClusterConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocClusterConfigWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocClusterConfigWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocClusterConfigWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocClusterConfigWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocClusterConfigWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterConfigWorkerConfigDiskConfigToProto converts a ClusterConfigWorkerConfigDiskConfig object to its proto representation.
func DataprocClusterConfigWorkerConfigDiskConfigToProto(o *dataproc.ClusterConfigWorkerConfigDiskConfig) *dataprocpb.DataprocClusterConfigWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterConfigWorkerConfigManagedGroupConfigToProto converts a ClusterConfigWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocClusterConfigWorkerConfigManagedGroupConfigToProto(o *dataproc.ClusterConfigWorkerConfigManagedGroupConfig) *dataprocpb.DataprocClusterConfigWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigWorkerConfigAcceleratorsToProto converts a ClusterConfigWorkerConfigAccelerators object to its proto representation.
func DataprocClusterConfigWorkerConfigAcceleratorsToProto(o *dataproc.ClusterConfigWorkerConfigAccelerators) *dataprocpb.DataprocClusterConfigWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigSecondaryWorkerConfigToProto converts a ClusterConfigSecondaryWorkerConfig object to its proto representation.
func DataprocClusterConfigSecondaryWorkerConfigToProto(o *dataproc.ClusterConfigSecondaryWorkerConfig) *dataprocpb.DataprocClusterConfigSecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSecondaryWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocClusterConfigSecondaryWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocClusterConfigSecondaryWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocClusterConfigSecondaryWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterConfigSecondaryWorkerConfigDiskConfigToProto converts a ClusterConfigSecondaryWorkerConfigDiskConfig object to its proto representation.
func DataprocClusterConfigSecondaryWorkerConfigDiskConfigToProto(o *dataproc.ClusterConfigSecondaryWorkerConfigDiskConfig) *dataprocpb.DataprocClusterConfigSecondaryWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSecondaryWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto converts a ClusterConfigSecondaryWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o *dataproc.ClusterConfigSecondaryWorkerConfigManagedGroupConfig) *dataprocpb.DataprocClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigSecondaryWorkerConfigAcceleratorsToProto converts a ClusterConfigSecondaryWorkerConfigAccelerators object to its proto representation.
func DataprocClusterConfigSecondaryWorkerConfigAcceleratorsToProto(o *dataproc.ClusterConfigSecondaryWorkerConfigAccelerators) *dataprocpb.DataprocClusterConfigSecondaryWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSecondaryWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigSoftwareConfigToProto converts a ClusterConfigSoftwareConfig object to its proto representation.
func DataprocClusterConfigSoftwareConfigToProto(o *dataproc.ClusterConfigSoftwareConfig) *dataprocpb.DataprocClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum(dataprocpb.DataprocClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// ClusterConfigInitializationActionsToProto converts a ClusterConfigInitializationActions object to its proto representation.
func DataprocClusterConfigInitializationActionsToProto(o *dataproc.ClusterConfigInitializationActions) *dataprocpb.DataprocClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// ClusterConfigEncryptionConfigToProto converts a ClusterConfigEncryptionConfig object to its proto representation.
func DataprocClusterConfigEncryptionConfigToProto(o *dataproc.ClusterConfigEncryptionConfig) *dataprocpb.DataprocClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// ClusterConfigAutoscalingConfigToProto converts a ClusterConfigAutoscalingConfig object to its proto representation.
func DataprocClusterConfigAutoscalingConfigToProto(o *dataproc.ClusterConfigAutoscalingConfig) *dataprocpb.DataprocClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// ClusterConfigSecurityConfigToProto converts a ClusterConfigSecurityConfig object to its proto representation.
func DataprocClusterConfigSecurityConfigToProto(o *dataproc.ClusterConfigSecurityConfig) *dataprocpb.DataprocClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// ClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocClusterConfigSecurityConfigKerberosConfigToProto(o *dataproc.ClusterConfigSecurityConfigKerberosConfig) *dataprocpb.DataprocClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigSecurityConfigKerberosConfig{}
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

// ClusterConfigLifecycleConfigToProto converts a ClusterConfigLifecycleConfig object to its proto representation.
func DataprocClusterConfigLifecycleConfigToProto(o *dataproc.ClusterConfigLifecycleConfig) *dataprocpb.DataprocClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// ClusterConfigEndpointConfigToProto converts a ClusterConfigEndpointConfig object to its proto representation.
func DataprocClusterConfigEndpointConfigToProto(o *dataproc.ClusterConfigEndpointConfig) *dataprocpb.DataprocClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocClusterConfigEndpointConfig{}
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
	p.SetConfig(DataprocClusterConfigToProto(resource.Config))
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
