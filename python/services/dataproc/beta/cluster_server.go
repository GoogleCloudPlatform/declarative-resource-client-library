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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/beta/dataproc_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *beta.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := beta.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *beta.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigMasterConfigPreemptibilityEnum converts a ClusterConfigMasterConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaClusterConfigMasterConfigPreemptibilityEnum(e betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum) *beta.ClusterConfigMasterConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.ClusterConfigMasterConfigPreemptibilityEnum(n[len("DataprocBetaClusterConfigMasterConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigWorkerConfigPreemptibilityEnum converts a ClusterConfigWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaClusterConfigWorkerConfigPreemptibilityEnum(e betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum) *beta.ClusterConfigWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.ClusterConfigWorkerConfigPreemptibilityEnum(n[len("DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigSecondaryWorkerConfigPreemptibilityEnum converts a ClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(e betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum) *beta.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum(n[len("DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigSoftwareConfigOptionalComponentsEnum converts a ClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum(e betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum) *beta.ClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := beta.ClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum"):])
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

// ProtoToClusterConfig converts a ClusterConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfig(p *betapb.DataprocBetaClusterConfig) *beta.ClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocBetaClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocBetaClusterConfigMasterConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocBetaClusterConfigWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocBetaClusterConfigSecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocBetaClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocBetaClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocBetaClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocBetaClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocBetaClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocBetaClusterConfigEndpointConfig(p.GetEndpointConfig()),
		GkeClusterConfig:      ProtoToDataprocBetaClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		MetastoreConfig:       ProtoToDataprocBetaClusterConfigMetastoreConfig(p.GetMetastoreConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocBetaClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfig converts a ClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigGceClusterConfig(p *betapb.DataprocBetaClusterConfigGceClusterConfig) *beta.ClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:     ProtoToDataprocBetaClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocBetaClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
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
func ProtoToDataprocBetaClusterConfigGceClusterConfigReservationAffinity(p *betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinity) *beta.ClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocBetaClusterConfigGceClusterConfigNodeGroupAffinity(p *betapb.DataprocBetaClusterConfigGceClusterConfigNodeGroupAffinity) *beta.ClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfig converts a ClusterConfigMasterConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigMasterConfig(p *betapb.DataprocBetaClusterConfigMasterConfig) *beta.ClusterConfigMasterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigMasterConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaClusterConfigMasterConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaClusterConfigMasterConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaClusterConfigMasterConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaClusterConfigMasterConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterConfigMasterConfigDiskConfig converts a ClusterConfigMasterConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigMasterConfigDiskConfig(p *betapb.DataprocBetaClusterConfigMasterConfigDiskConfig) *beta.ClusterConfigMasterConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigMasterConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigManagedGroupConfig converts a ClusterConfigMasterConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigMasterConfigManagedGroupConfig(p *betapb.DataprocBetaClusterConfigMasterConfigManagedGroupConfig) *beta.ClusterConfigMasterConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigMasterConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigAccelerators converts a ClusterConfigMasterConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaClusterConfigMasterConfigAccelerators(p *betapb.DataprocBetaClusterConfigMasterConfigAccelerators) *beta.ClusterConfigMasterConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigMasterConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfig converts a ClusterConfigWorkerConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigWorkerConfig(p *betapb.DataprocBetaClusterConfigWorkerConfig) *beta.ClusterConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaClusterConfigWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaClusterConfigWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaClusterConfigWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaClusterConfigWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigDiskConfig converts a ClusterConfigWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigWorkerConfigDiskConfig(p *betapb.DataprocBetaClusterConfigWorkerConfigDiskConfig) *beta.ClusterConfigWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigManagedGroupConfig converts a ClusterConfigWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigWorkerConfigManagedGroupConfig(p *betapb.DataprocBetaClusterConfigWorkerConfigManagedGroupConfig) *beta.ClusterConfigWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigAccelerators converts a ClusterConfigWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaClusterConfigWorkerConfigAccelerators(p *betapb.DataprocBetaClusterConfigWorkerConfigAccelerators) *beta.ClusterConfigWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfig converts a ClusterConfigSecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigSecondaryWorkerConfig(p *betapb.DataprocBetaClusterConfigSecondaryWorkerConfig) *beta.ClusterConfigSecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSecondaryWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigDiskConfig converts a ClusterConfigSecondaryWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigDiskConfig(p *betapb.DataprocBetaClusterConfigSecondaryWorkerConfigDiskConfig) *beta.ClusterConfigSecondaryWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSecondaryWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigManagedGroupConfig converts a ClusterConfigSecondaryWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfig(p *betapb.DataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfig) *beta.ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSecondaryWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigAccelerators converts a ClusterConfigSecondaryWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaClusterConfigSecondaryWorkerConfigAccelerators(p *betapb.DataprocBetaClusterConfigSecondaryWorkerConfigAccelerators) *beta.ClusterConfigSecondaryWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSecondaryWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigSoftwareConfig converts a ClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigSoftwareConfig(p *betapb.DataprocBetaClusterConfigSoftwareConfig) *beta.ClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterConfigInitializationActions converts a ClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocBetaClusterConfigInitializationActions(p *betapb.DataprocBetaClusterConfigInitializationActions) *beta.ClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToClusterConfigEncryptionConfig converts a ClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigEncryptionConfig(p *betapb.DataprocBetaClusterConfigEncryptionConfig) *beta.ClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToClusterConfigAutoscalingConfig converts a ClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigAutoscalingConfig(p *betapb.DataprocBetaClusterConfigAutoscalingConfig) *beta.ClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfig converts a ClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigSecurityConfig(p *betapb.DataprocBetaClusterConfigSecurityConfig) *beta.ClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocBetaClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfigKerberosConfig converts a ClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigSecurityConfigKerberosConfig(p *betapb.DataprocBetaClusterConfigSecurityConfigKerberosConfig) *beta.ClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocBetaClusterConfigLifecycleConfig(p *betapb.DataprocBetaClusterConfigLifecycleConfig) *beta.ClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterConfigEndpointConfig converts a ClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigEndpointConfig(p *betapb.DataprocBetaClusterConfigEndpointConfig) *beta.ClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToClusterConfigGkeClusterConfig converts a ClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigGkeClusterConfig(p *betapb.DataprocBetaClusterConfigGkeClusterConfig) *beta.ClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object from its proto representation.
func ProtoToDataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *betapb.DataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *beta.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.GetTargetGkeCluster()),
		ClusterNamespace: dcl.StringOrNil(p.GetClusterNamespace()),
	}
	return obj
}

// ProtoToClusterConfigMetastoreConfig converts a ClusterConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocBetaClusterConfigMetastoreConfig(p *betapb.DataprocBetaClusterConfigMetastoreConfig) *beta.ClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfigMetastoreConfig{
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
		Config:      ProtoToDataprocBetaClusterConfig(p.GetConfig()),
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

// ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *beta.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return betapb.DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *beta.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterConfigMasterConfigPreemptibilityEnumToProto converts a ClusterConfigMasterConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaClusterConfigMasterConfigPreemptibilityEnumToProto(e *beta.ClusterConfigMasterConfigPreemptibilityEnum) betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum_value["ClusterConfigMasterConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaClusterConfigMasterConfigPreemptibilityEnum(0)
}

// ClusterConfigWorkerConfigPreemptibilityEnumToProto converts a ClusterConfigWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaClusterConfigWorkerConfigPreemptibilityEnumToProto(e *beta.ClusterConfigWorkerConfigPreemptibilityEnum) betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum_value["ClusterConfigWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaClusterConfigWorkerConfigPreemptibilityEnum(0)
}

// ClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto converts a ClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(e *beta.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum) betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum_value["ClusterConfigSecondaryWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
}

// ClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a ClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *beta.ClusterConfigSoftwareConfigOptionalComponentsEnum) betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum_value["ClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum(0)
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

// ClusterConfigToProto converts a ClusterConfig object to its proto representation.
func DataprocBetaClusterConfigToProto(o *beta.ClusterConfig) *betapb.DataprocBetaClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocBetaClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocBetaClusterConfigMasterConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocBetaClusterConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocBetaClusterConfigSecondaryWorkerConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocBetaClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocBetaClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocBetaClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocBetaClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocBetaClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocBetaClusterConfigEndpointConfigToProto(o.EndpointConfig))
	p.SetGkeClusterConfig(DataprocBetaClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetMetastoreConfig(DataprocBetaClusterConfigMetastoreConfigToProto(o.MetastoreConfig))
	sInitializationActions := make([]*betapb.DataprocBetaClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocBetaClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// ClusterConfigGceClusterConfigToProto converts a ClusterConfigGceClusterConfig object to its proto representation.
func DataprocBetaClusterConfigGceClusterConfigToProto(o *beta.ClusterConfigGceClusterConfig) *betapb.DataprocBetaClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocBetaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocBetaClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocBetaClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
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
func DataprocBetaClusterConfigGceClusterConfigReservationAffinityToProto(o *beta.ClusterConfigGceClusterConfigReservationAffinity) *betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocBetaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// ClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocBetaClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *beta.ClusterConfigGceClusterConfigNodeGroupAffinity) *betapb.DataprocBetaClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// ClusterConfigMasterConfigToProto converts a ClusterConfigMasterConfig object to its proto representation.
func DataprocBetaClusterConfigMasterConfigToProto(o *beta.ClusterConfigMasterConfig) *betapb.DataprocBetaClusterConfigMasterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigMasterConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaClusterConfigMasterConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaClusterConfigMasterConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaClusterConfigMasterConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaClusterConfigMasterConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaClusterConfigMasterConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterConfigMasterConfigDiskConfigToProto converts a ClusterConfigMasterConfigDiskConfig object to its proto representation.
func DataprocBetaClusterConfigMasterConfigDiskConfigToProto(o *beta.ClusterConfigMasterConfigDiskConfig) *betapb.DataprocBetaClusterConfigMasterConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigMasterConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterConfigMasterConfigManagedGroupConfigToProto converts a ClusterConfigMasterConfigManagedGroupConfig object to its proto representation.
func DataprocBetaClusterConfigMasterConfigManagedGroupConfigToProto(o *beta.ClusterConfigMasterConfigManagedGroupConfig) *betapb.DataprocBetaClusterConfigMasterConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigMasterConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigMasterConfigAcceleratorsToProto converts a ClusterConfigMasterConfigAccelerators object to its proto representation.
func DataprocBetaClusterConfigMasterConfigAcceleratorsToProto(o *beta.ClusterConfigMasterConfigAccelerators) *betapb.DataprocBetaClusterConfigMasterConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigMasterConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigWorkerConfigToProto converts a ClusterConfigWorkerConfig object to its proto representation.
func DataprocBetaClusterConfigWorkerConfigToProto(o *beta.ClusterConfigWorkerConfig) *betapb.DataprocBetaClusterConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaClusterConfigWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaClusterConfigWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaClusterConfigWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaClusterConfigWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaClusterConfigWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterConfigWorkerConfigDiskConfigToProto converts a ClusterConfigWorkerConfigDiskConfig object to its proto representation.
func DataprocBetaClusterConfigWorkerConfigDiskConfigToProto(o *beta.ClusterConfigWorkerConfigDiskConfig) *betapb.DataprocBetaClusterConfigWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterConfigWorkerConfigManagedGroupConfigToProto converts a ClusterConfigWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocBetaClusterConfigWorkerConfigManagedGroupConfigToProto(o *beta.ClusterConfigWorkerConfigManagedGroupConfig) *betapb.DataprocBetaClusterConfigWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigWorkerConfigAcceleratorsToProto converts a ClusterConfigWorkerConfigAccelerators object to its proto representation.
func DataprocBetaClusterConfigWorkerConfigAcceleratorsToProto(o *beta.ClusterConfigWorkerConfigAccelerators) *betapb.DataprocBetaClusterConfigWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigSecondaryWorkerConfigToProto converts a ClusterConfigSecondaryWorkerConfig object to its proto representation.
func DataprocBetaClusterConfigSecondaryWorkerConfigToProto(o *beta.ClusterConfigSecondaryWorkerConfig) *betapb.DataprocBetaClusterConfigSecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSecondaryWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaClusterConfigSecondaryWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaClusterConfigSecondaryWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaClusterConfigSecondaryWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterConfigSecondaryWorkerConfigDiskConfigToProto converts a ClusterConfigSecondaryWorkerConfigDiskConfig object to its proto representation.
func DataprocBetaClusterConfigSecondaryWorkerConfigDiskConfigToProto(o *beta.ClusterConfigSecondaryWorkerConfigDiskConfig) *betapb.DataprocBetaClusterConfigSecondaryWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSecondaryWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// ClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto converts a ClusterConfigSecondaryWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o *beta.ClusterConfigSecondaryWorkerConfigManagedGroupConfig) *betapb.DataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigSecondaryWorkerConfigAcceleratorsToProto converts a ClusterConfigSecondaryWorkerConfigAccelerators object to its proto representation.
func DataprocBetaClusterConfigSecondaryWorkerConfigAcceleratorsToProto(o *beta.ClusterConfigSecondaryWorkerConfigAccelerators) *betapb.DataprocBetaClusterConfigSecondaryWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSecondaryWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigSoftwareConfigToProto converts a ClusterConfigSoftwareConfig object to its proto representation.
func DataprocBetaClusterConfigSoftwareConfigToProto(o *beta.ClusterConfigSoftwareConfig) *betapb.DataprocBetaClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum(betapb.DataprocBetaClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// ClusterConfigInitializationActionsToProto converts a ClusterConfigInitializationActions object to its proto representation.
func DataprocBetaClusterConfigInitializationActionsToProto(o *beta.ClusterConfigInitializationActions) *betapb.DataprocBetaClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// ClusterConfigEncryptionConfigToProto converts a ClusterConfigEncryptionConfig object to its proto representation.
func DataprocBetaClusterConfigEncryptionConfigToProto(o *beta.ClusterConfigEncryptionConfig) *betapb.DataprocBetaClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// ClusterConfigAutoscalingConfigToProto converts a ClusterConfigAutoscalingConfig object to its proto representation.
func DataprocBetaClusterConfigAutoscalingConfigToProto(o *beta.ClusterConfigAutoscalingConfig) *betapb.DataprocBetaClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// ClusterConfigSecurityConfigToProto converts a ClusterConfigSecurityConfig object to its proto representation.
func DataprocBetaClusterConfigSecurityConfigToProto(o *beta.ClusterConfigSecurityConfig) *betapb.DataprocBetaClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocBetaClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// ClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocBetaClusterConfigSecurityConfigKerberosConfigToProto(o *beta.ClusterConfigSecurityConfigKerberosConfig) *betapb.DataprocBetaClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigSecurityConfigKerberosConfig{}
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
func DataprocBetaClusterConfigLifecycleConfigToProto(o *beta.ClusterConfigLifecycleConfig) *betapb.DataprocBetaClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// ClusterConfigEndpointConfigToProto converts a ClusterConfigEndpointConfig object to its proto representation.
func DataprocBetaClusterConfigEndpointConfigToProto(o *beta.ClusterConfigEndpointConfig) *betapb.DataprocBetaClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// ClusterConfigGkeClusterConfigToProto converts a ClusterConfigGkeClusterConfig object to its proto representation.
func DataprocBetaClusterConfigGkeClusterConfigToProto(o *beta.ClusterConfigGkeClusterConfig) *betapb.DataprocBetaClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigGkeClusterConfig{}
	p.SetNamespacedGkeDeploymentTarget(DataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget))
	return p
}

// ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object to its proto representation.
func DataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *beta.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *betapb.DataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
	p.SetTargetGkeCluster(dcl.ValueOrEmptyString(o.TargetGkeCluster))
	p.SetClusterNamespace(dcl.ValueOrEmptyString(o.ClusterNamespace))
	return p
}

// ClusterConfigMetastoreConfigToProto converts a ClusterConfigMetastoreConfig object to its proto representation.
func DataprocBetaClusterConfigMetastoreConfigToProto(o *beta.ClusterConfigMetastoreConfig) *betapb.DataprocBetaClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaClusterConfigMetastoreConfig{}
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
	p.SetConfig(DataprocBetaClusterConfigToProto(resource.Config))
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
