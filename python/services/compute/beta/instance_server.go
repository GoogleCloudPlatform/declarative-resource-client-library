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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceDisksInterfaceEnum converts a InstanceDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeBetaInstanceDisksInterfaceEnum(e betapb.ComputeBetaInstanceDisksInterfaceEnum) *beta.InstanceDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceDisksInterfaceEnum_name[int32(e)]; ok {
		e := beta.InstanceDisksInterfaceEnum(n[len("InstanceDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksModeEnum converts a InstanceDisksModeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceDisksModeEnum(e betapb.ComputeBetaInstanceDisksModeEnum) *beta.InstanceDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceDisksModeEnum_name[int32(e)]; ok {
		e := beta.InstanceDisksModeEnum(n[len("InstanceDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksTypeEnum converts a InstanceDisksTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceDisksTypeEnum(e betapb.ComputeBetaInstanceDisksTypeEnum) *beta.InstanceDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceDisksTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceDisksTypeEnum(n[len("InstanceDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsTypeEnum converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(e betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum) *beta.InstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworkInterfacesAccessConfigsTypeEnum(n[len("InstanceNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStatusEnum converts a InstanceStatusEnum enum from its proto representation.
func ProtoToComputeBetaInstanceStatusEnum(e betapb.ComputeBetaInstanceStatusEnum) *beta.InstanceStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceStatusEnum_name[int32(e)]; ok {
		e := beta.InstanceStatusEnum(n[len("InstanceStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisks converts a InstanceDisks resource from its proto representation.
func ProtoToComputeBetaInstanceDisks(p *betapb.ComputeBetaInstanceDisks) *beta.InstanceDisks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisks{
		AutoDelete:        dcl.Bool(p.AutoDelete),
		Boot:              dcl.Bool(p.Boot),
		DeviceName:        dcl.StringOrNil(p.DeviceName),
		DiskEncryptionKey: ProtoToComputeBetaInstanceDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.Index),
		InitializeParams:  ProtoToComputeBetaInstanceDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeBetaInstanceDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeBetaInstanceDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.Source),
		Type:              ProtoToComputeBetaInstanceDisksTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceDisksDiskEncryptionKey converts a InstanceDisksDiskEncryptionKey resource from its proto representation.
func ProtoToComputeBetaInstanceDisksDiskEncryptionKey(p *betapb.ComputeBetaInstanceDisksDiskEncryptionKey) *beta.InstanceDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.RawKey),
		RsaEncryptedKey: dcl.StringOrNil(p.RsaEncryptedKey),
		Sha256:          dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParams converts a InstanceDisksInitializeParams resource from its proto representation.
func ProtoToComputeBetaInstanceDisksInitializeParams(p *betapb.ComputeBetaInstanceDisksInitializeParams) *beta.InstanceDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisksInitializeParams{
		DiskName:                 dcl.StringOrNil(p.DiskName),
		DiskSizeGb:               dcl.Int64OrNil(p.DiskSizeGb),
		DiskType:                 dcl.StringOrNil(p.DiskType),
		SourceImage:              dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey: ProtoToComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParamsSourceImageEncryptionKey converts a InstanceDisksInitializeParamsSourceImageEncryptionKey resource from its proto representation.
func ProtoToComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey(p *betapb.ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey) *beta.InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.StringOrNil(p.RawKey),
		Sha256: dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToInstanceGuestAccelerators converts a InstanceGuestAccelerators resource from its proto representation.
func ProtoToComputeBetaInstanceGuestAccelerators(p *betapb.ComputeBetaInstanceGuestAccelerators) *beta.InstanceGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToInstanceNetworkInterfaces converts a InstanceNetworkInterfaces resource from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfaces(p *betapb.ComputeBetaInstanceNetworkInterfaces) *beta.InstanceNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfaces{
		Name:       dcl.StringOrNil(p.Name),
		Network:    dcl.StringOrNil(p.Network),
		NetworkIP:  dcl.StringOrNil(p.NetworkIp),
		Subnetwork: dcl.StringOrNil(p.Subnetwork),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeBetaInstanceNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAccessConfigs converts a InstanceNetworkInterfacesAccessConfigs resource from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigs(p *betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs) *beta.InstanceNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfacesAccessConfigs{
		Name:  dcl.StringOrNil(p.Name),
		NatIP: dcl.StringOrNil(p.NatIp),
		Type:  ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAliasIPRanges converts a InstanceNetworkInterfacesAliasIPRanges resource from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAliasIPRanges(p *betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges) *beta.InstanceNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.IpCidrRange),
		SubnetworkRangeName: dcl.StringOrNil(p.SubnetworkRangeName),
	}
	return obj
}

// ProtoToInstanceScheduling converts a InstanceScheduling resource from its proto representation.
func ProtoToComputeBetaInstanceScheduling(p *betapb.ComputeBetaInstanceScheduling) *beta.InstanceScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceScheduling{
		AutomaticRestart:  dcl.Bool(p.AutomaticRestart),
		OnHostMaintenance: dcl.StringOrNil(p.OnHostMaintenance),
		Preemptible:       dcl.Bool(p.Preemptible),
	}
	return obj
}

// ProtoToInstanceServiceAccounts converts a InstanceServiceAccounts resource from its proto representation.
func ProtoToComputeBetaInstanceServiceAccounts(p *betapb.ComputeBetaInstanceServiceAccounts) *beta.InstanceServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceServiceAccounts{
		Email: dcl.StringOrNil(p.Email),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceShieldedInstanceConfig converts a InstanceShieldedInstanceConfig resource from its proto representation.
func ProtoToComputeBetaInstanceShieldedInstanceConfig(p *betapb.ComputeBetaInstanceShieldedInstanceConfig) *beta.InstanceShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableVtpm:                dcl.Bool(p.EnableVtpm),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.ComputeBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		CanIPForward:           dcl.Bool(p.CanIpForward),
		CpuPlatform:            dcl.StringOrNil(p.CpuPlatform),
		CreationTimestamp:      dcl.StringOrNil(p.CreationTimestamp),
		DeletionProtection:     dcl.Bool(p.DeletionProtection),
		Description:            dcl.StringOrNil(p.Description),
		Hostname:               dcl.StringOrNil(p.Hostname),
		Id:                     dcl.StringOrNil(p.Id),
		MachineType:            dcl.StringOrNil(p.MachineType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		Name:                   dcl.StringOrNil(p.Name),
		Scheduling:             ProtoToComputeBetaInstanceScheduling(p.GetScheduling()),
		ShieldedInstanceConfig: ProtoToComputeBetaInstanceShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Status:                 ProtoToComputeBetaInstanceStatusEnum(p.GetStatus()),
		StatusMessage:          dcl.StringOrNil(p.StatusMessage),
		Zone:                   dcl.StringOrNil(p.Zone),
		Project:                dcl.StringOrNil(p.Project),
		SelfLink:               dcl.StringOrNil(p.SelfLink),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeBetaInstanceDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeBetaInstanceGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeBetaInstanceNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeBetaInstanceServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// InstanceDisksInterfaceEnumToProto converts a InstanceDisksInterfaceEnum enum to its proto representation.
func ComputeBetaInstanceDisksInterfaceEnumToProto(e *beta.InstanceDisksInterfaceEnum) betapb.ComputeBetaInstanceDisksInterfaceEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceDisksInterfaceEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceDisksInterfaceEnum_value["InstanceDisksInterfaceEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceDisksInterfaceEnum(v)
	}
	return betapb.ComputeBetaInstanceDisksInterfaceEnum(0)
}

// InstanceDisksModeEnumToProto converts a InstanceDisksModeEnum enum to its proto representation.
func ComputeBetaInstanceDisksModeEnumToProto(e *beta.InstanceDisksModeEnum) betapb.ComputeBetaInstanceDisksModeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceDisksModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceDisksModeEnum_value["InstanceDisksModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceDisksModeEnum(v)
	}
	return betapb.ComputeBetaInstanceDisksModeEnum(0)
}

// InstanceDisksTypeEnumToProto converts a InstanceDisksTypeEnum enum to its proto representation.
func ComputeBetaInstanceDisksTypeEnumToProto(e *beta.InstanceDisksTypeEnum) betapb.ComputeBetaInstanceDisksTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceDisksTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceDisksTypeEnum_value["InstanceDisksTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceDisksTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceDisksTypeEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(e *beta.InstanceNetworkInterfacesAccessConfigsTypeEnum) betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum_value["InstanceNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceStatusEnumToProto converts a InstanceStatusEnum enum to its proto representation.
func ComputeBetaInstanceStatusEnumToProto(e *beta.InstanceStatusEnum) betapb.ComputeBetaInstanceStatusEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceStatusEnum_value["InstanceStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceStatusEnum(v)
	}
	return betapb.ComputeBetaInstanceStatusEnum(0)
}

// InstanceDisksToProto converts a InstanceDisks resource to its proto representation.
func ComputeBetaInstanceDisksToProto(o *beta.InstanceDisks) *betapb.ComputeBetaInstanceDisks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisks{
		AutoDelete:        dcl.ValueOrEmptyBool(o.AutoDelete),
		Boot:              dcl.ValueOrEmptyBool(o.Boot),
		DeviceName:        dcl.ValueOrEmptyString(o.DeviceName),
		DiskEncryptionKey: ComputeBetaInstanceDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey),
		Index:             dcl.ValueOrEmptyInt64(o.Index),
		InitializeParams:  ComputeBetaInstanceDisksInitializeParamsToProto(o.InitializeParams),
		Interface:         ComputeBetaInstanceDisksInterfaceEnumToProto(o.Interface),
		Mode:              ComputeBetaInstanceDisksModeEnumToProto(o.Mode),
		Source:            dcl.ValueOrEmptyString(o.Source),
		Type:              ComputeBetaInstanceDisksTypeEnumToProto(o.Type),
	}
	return p
}

// InstanceDisksDiskEncryptionKeyToProto converts a InstanceDisksDiskEncryptionKey resource to its proto representation.
func ComputeBetaInstanceDisksDiskEncryptionKeyToProto(o *beta.InstanceDisksDiskEncryptionKey) *betapb.ComputeBetaInstanceDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisksDiskEncryptionKey{
		RawKey:          dcl.ValueOrEmptyString(o.RawKey),
		RsaEncryptedKey: dcl.ValueOrEmptyString(o.RsaEncryptedKey),
		Sha256:          dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// InstanceDisksInitializeParamsToProto converts a InstanceDisksInitializeParams resource to its proto representation.
func ComputeBetaInstanceDisksInitializeParamsToProto(o *beta.InstanceDisksInitializeParams) *betapb.ComputeBetaInstanceDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisksInitializeParams{
		DiskName:                 dcl.ValueOrEmptyString(o.DiskName),
		DiskSizeGb:               dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		DiskType:                 dcl.ValueOrEmptyString(o.DiskType),
		SourceImage:              dcl.ValueOrEmptyString(o.SourceImage),
		SourceImageEncryptionKey: ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey),
	}
	return p
}

// InstanceDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceDisksInitializeParamsSourceImageEncryptionKey resource to its proto representation.
func ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o *beta.InstanceDisksInitializeParamsSourceImageEncryptionKey) *betapb.ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.ValueOrEmptyString(o.RawKey),
		Sha256: dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// InstanceGuestAcceleratorsToProto converts a InstanceGuestAccelerators resource to its proto representation.
func ComputeBetaInstanceGuestAcceleratorsToProto(o *beta.InstanceGuestAccelerators) *betapb.ComputeBetaInstanceGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGuestAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// InstanceNetworkInterfacesToProto converts a InstanceNetworkInterfaces resource to its proto representation.
func ComputeBetaInstanceNetworkInterfacesToProto(o *beta.InstanceNetworkInterfaces) *betapb.ComputeBetaInstanceNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfaces{
		Name:       dcl.ValueOrEmptyString(o.Name),
		Network:    dcl.ValueOrEmptyString(o.Network),
		NetworkIp:  dcl.ValueOrEmptyString(o.NetworkIP),
		Subnetwork: dcl.ValueOrEmptyString(o.Subnetwork),
	}
	for _, r := range o.AccessConfigs {
		p.AccessConfigs = append(p.AccessConfigs, ComputeBetaInstanceNetworkInterfacesAccessConfigsToProto(&r))
	}
	for _, r := range o.AliasIPRanges {
		p.AliasIpRanges = append(p.AliasIpRanges, ComputeBetaInstanceNetworkInterfacesAliasIPRangesToProto(&r))
	}
	return p
}

// InstanceNetworkInterfacesAccessConfigsToProto converts a InstanceNetworkInterfacesAccessConfigs resource to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAccessConfigsToProto(o *beta.InstanceNetworkInterfacesAccessConfigs) *betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs{
		Name:  dcl.ValueOrEmptyString(o.Name),
		NatIp: dcl.ValueOrEmptyString(o.NatIP),
		Type:  ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type),
	}
	return p
}

// InstanceNetworkInterfacesAliasIPRangesToProto converts a InstanceNetworkInterfacesAliasIPRanges resource to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAliasIPRangesToProto(o *beta.InstanceNetworkInterfacesAliasIPRanges) *betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges{
		IpCidrRange:         dcl.ValueOrEmptyString(o.IPCidrRange),
		SubnetworkRangeName: dcl.ValueOrEmptyString(o.SubnetworkRangeName),
	}
	return p
}

// InstanceSchedulingToProto converts a InstanceScheduling resource to its proto representation.
func ComputeBetaInstanceSchedulingToProto(o *beta.InstanceScheduling) *betapb.ComputeBetaInstanceScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceScheduling{
		AutomaticRestart:  dcl.ValueOrEmptyBool(o.AutomaticRestart),
		OnHostMaintenance: dcl.ValueOrEmptyString(o.OnHostMaintenance),
		Preemptible:       dcl.ValueOrEmptyBool(o.Preemptible),
	}
	return p
}

// InstanceServiceAccountsToProto converts a InstanceServiceAccounts resource to its proto representation.
func ComputeBetaInstanceServiceAccountsToProto(o *beta.InstanceServiceAccounts) *betapb.ComputeBetaInstanceServiceAccounts {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceServiceAccounts{
		Email: dcl.ValueOrEmptyString(o.Email),
	}
	for _, r := range o.Scopes {
		p.Scopes = append(p.Scopes, r)
	}
	return p
}

// InstanceShieldedInstanceConfigToProto converts a InstanceShieldedInstanceConfig resource to its proto representation.
func ComputeBetaInstanceShieldedInstanceConfigToProto(o *beta.InstanceShieldedInstanceConfig) *betapb.ComputeBetaInstanceShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableVtpm:                dcl.ValueOrEmptyBool(o.EnableVtpm),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.ComputeBetaInstance {
	p := &betapb.ComputeBetaInstance{
		CanIpForward:           dcl.ValueOrEmptyBool(resource.CanIPForward),
		CpuPlatform:            dcl.ValueOrEmptyString(resource.CpuPlatform),
		CreationTimestamp:      dcl.ValueOrEmptyString(resource.CreationTimestamp),
		DeletionProtection:     dcl.ValueOrEmptyBool(resource.DeletionProtection),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		Hostname:               dcl.ValueOrEmptyString(resource.Hostname),
		Id:                     dcl.ValueOrEmptyString(resource.Id),
		MachineType:            dcl.ValueOrEmptyString(resource.MachineType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(resource.MinCpuPlatform),
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Scheduling:             ComputeBetaInstanceSchedulingToProto(resource.Scheduling),
		ShieldedInstanceConfig: ComputeBetaInstanceShieldedInstanceConfigToProto(resource.ShieldedInstanceConfig),
		Status:                 ComputeBetaInstanceStatusEnumToProto(resource.Status),
		StatusMessage:          dcl.ValueOrEmptyString(resource.StatusMessage),
		Zone:                   dcl.ValueOrEmptyString(resource.Zone),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		SelfLink:               dcl.ValueOrEmptyString(resource.SelfLink),
	}
	for _, r := range resource.Disks {
		p.Disks = append(p.Disks, ComputeBetaInstanceDisksToProto(&r))
	}
	for _, r := range resource.GuestAccelerators {
		p.GuestAccelerators = append(p.GuestAccelerators, ComputeBetaInstanceGuestAcceleratorsToProto(&r))
	}
	for _, r := range resource.NetworkInterfaces {
		p.NetworkInterfaces = append(p.NetworkInterfaces, ComputeBetaInstanceNetworkInterfacesToProto(&r))
	}
	for _, r := range resource.ServiceAccounts {
		p.ServiceAccounts = append(p.ServiceAccounts, ComputeBetaInstanceServiceAccountsToProto(&r))
	}
	for _, r := range resource.Tags {
		p.Tags = append(p.Tags, r)
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInstanceRequest) (*betapb.ComputeBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyComputeBetaInstance(ctx context.Context, request *betapb.ApplyComputeBetaInstanceRequest) (*betapb.ComputeBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteComputeBetaInstance(ctx context.Context, request *betapb.DeleteComputeBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListComputeBetaInstance(ctx context.Context, request *betapb.ListComputeBetaInstanceRequest) (*betapb.ListComputeBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Zone)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
