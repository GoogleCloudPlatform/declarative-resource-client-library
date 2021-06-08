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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceDisksInterfaceEnum converts a InstanceDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeInstanceDisksInterfaceEnum(e computepb.ComputeInstanceDisksInterfaceEnum) *compute.InstanceDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceDisksInterfaceEnum_name[int32(e)]; ok {
		e := compute.InstanceDisksInterfaceEnum(n[len("ComputeInstanceDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksModeEnum converts a InstanceDisksModeEnum enum from its proto representation.
func ProtoToComputeInstanceDisksModeEnum(e computepb.ComputeInstanceDisksModeEnum) *compute.InstanceDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceDisksModeEnum_name[int32(e)]; ok {
		e := compute.InstanceDisksModeEnum(n[len("ComputeInstanceDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksTypeEnum converts a InstanceDisksTypeEnum enum from its proto representation.
func ProtoToComputeInstanceDisksTypeEnum(e computepb.ComputeInstanceDisksTypeEnum) *compute.InstanceDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceDisksTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceDisksTypeEnum(n[len("ComputeInstanceDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsTypeEnum converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(e computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum) *compute.InstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceNetworkInterfacesAccessConfigsTypeEnum(n[len("ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStatusEnum converts a InstanceStatusEnum enum from its proto representation.
func ProtoToComputeInstanceStatusEnum(e computepb.ComputeInstanceStatusEnum) *compute.InstanceStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceStatusEnum_name[int32(e)]; ok {
		e := compute.InstanceStatusEnum(n[len("ComputeInstanceStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisks converts a InstanceDisks resource from its proto representation.
func ProtoToComputeInstanceDisks(p *computepb.ComputeInstanceDisks) *compute.InstanceDisks {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisks{
		AutoDelete:        dcl.Bool(p.AutoDelete),
		Boot:              dcl.Bool(p.Boot),
		DeviceName:        dcl.StringOrNil(p.DeviceName),
		DiskEncryptionKey: ProtoToComputeInstanceDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.Index),
		InitializeParams:  ProtoToComputeInstanceDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeInstanceDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeInstanceDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.Source),
		Type:              ProtoToComputeInstanceDisksTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceDisksDiskEncryptionKey converts a InstanceDisksDiskEncryptionKey resource from its proto representation.
func ProtoToComputeInstanceDisksDiskEncryptionKey(p *computepb.ComputeInstanceDisksDiskEncryptionKey) *compute.InstanceDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.RawKey),
		RsaEncryptedKey: dcl.StringOrNil(p.RsaEncryptedKey),
		Sha256:          dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParams converts a InstanceDisksInitializeParams resource from its proto representation.
func ProtoToComputeInstanceDisksInitializeParams(p *computepb.ComputeInstanceDisksInitializeParams) *compute.InstanceDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisksInitializeParams{
		DiskName:                 dcl.StringOrNil(p.DiskName),
		DiskSizeGb:               dcl.Int64OrNil(p.DiskSizeGb),
		DiskType:                 dcl.StringOrNil(p.DiskType),
		SourceImage:              dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey: ProtoToComputeInstanceDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParamsSourceImageEncryptionKey converts a InstanceDisksInitializeParamsSourceImageEncryptionKey resource from its proto representation.
func ProtoToComputeInstanceDisksInitializeParamsSourceImageEncryptionKey(p *computepb.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey) *compute.InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.StringOrNil(p.RawKey),
		Sha256: dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToInstanceGuestAccelerators converts a InstanceGuestAccelerators resource from its proto representation.
func ProtoToComputeInstanceGuestAccelerators(p *computepb.ComputeInstanceGuestAccelerators) *compute.InstanceGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToInstanceNetworkInterfaces converts a InstanceNetworkInterfaces resource from its proto representation.
func ProtoToComputeInstanceNetworkInterfaces(p *computepb.ComputeInstanceNetworkInterfaces) *compute.InstanceNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfaces{
		Name:       dcl.StringOrNil(p.Name),
		Network:    dcl.StringOrNil(p.Network),
		NetworkIP:  dcl.StringOrNil(p.NetworkIp),
		Subnetwork: dcl.StringOrNil(p.Subnetwork),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeInstanceNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeInstanceNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAccessConfigs converts a InstanceNetworkInterfacesAccessConfigs resource from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAccessConfigs(p *computepb.ComputeInstanceNetworkInterfacesAccessConfigs) *compute.InstanceNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfacesAccessConfigs{
		Name:  dcl.StringOrNil(p.Name),
		NatIP: dcl.StringOrNil(p.NatIp),
		Type:  ProtoToComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAliasIPRanges converts a InstanceNetworkInterfacesAliasIPRanges resource from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAliasIPRanges(p *computepb.ComputeInstanceNetworkInterfacesAliasIPRanges) *compute.InstanceNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.IpCidrRange),
		SubnetworkRangeName: dcl.StringOrNil(p.SubnetworkRangeName),
	}
	return obj
}

// ProtoToInstanceScheduling converts a InstanceScheduling resource from its proto representation.
func ProtoToComputeInstanceScheduling(p *computepb.ComputeInstanceScheduling) *compute.InstanceScheduling {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceScheduling{
		AutomaticRestart:  dcl.Bool(p.AutomaticRestart),
		OnHostMaintenance: dcl.StringOrNil(p.OnHostMaintenance),
		Preemptible:       dcl.Bool(p.Preemptible),
	}
	return obj
}

// ProtoToInstanceServiceAccounts converts a InstanceServiceAccounts resource from its proto representation.
func ProtoToComputeInstanceServiceAccounts(p *computepb.ComputeInstanceServiceAccounts) *compute.InstanceServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceServiceAccounts{
		Email: dcl.StringOrNil(p.Email),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceShieldedInstanceConfig converts a InstanceShieldedInstanceConfig resource from its proto representation.
func ProtoToComputeInstanceShieldedInstanceConfig(p *computepb.ComputeInstanceShieldedInstanceConfig) *compute.InstanceShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableVtpm:                dcl.Bool(p.EnableVtpm),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *computepb.ComputeInstance) *compute.Instance {
	obj := &compute.Instance{
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
		Scheduling:             ProtoToComputeInstanceScheduling(p.GetScheduling()),
		ShieldedInstanceConfig: ProtoToComputeInstanceShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Status:                 ProtoToComputeInstanceStatusEnum(p.GetStatus()),
		StatusMessage:          dcl.StringOrNil(p.StatusMessage),
		Zone:                   dcl.StringOrNil(p.Zone),
		Project:                dcl.StringOrNil(p.Project),
		SelfLink:               dcl.StringOrNil(p.SelfLink),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeInstanceDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeInstanceGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeInstanceNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeInstanceServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// InstanceDisksInterfaceEnumToProto converts a InstanceDisksInterfaceEnum enum to its proto representation.
func ComputeInstanceDisksInterfaceEnumToProto(e *compute.InstanceDisksInterfaceEnum) computepb.ComputeInstanceDisksInterfaceEnum {
	if e == nil {
		return computepb.ComputeInstanceDisksInterfaceEnum(0)
	}
	if v, ok := computepb.ComputeInstanceDisksInterfaceEnum_value["InstanceDisksInterfaceEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceDisksInterfaceEnum(v)
	}
	return computepb.ComputeInstanceDisksInterfaceEnum(0)
}

// InstanceDisksModeEnumToProto converts a InstanceDisksModeEnum enum to its proto representation.
func ComputeInstanceDisksModeEnumToProto(e *compute.InstanceDisksModeEnum) computepb.ComputeInstanceDisksModeEnum {
	if e == nil {
		return computepb.ComputeInstanceDisksModeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceDisksModeEnum_value["InstanceDisksModeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceDisksModeEnum(v)
	}
	return computepb.ComputeInstanceDisksModeEnum(0)
}

// InstanceDisksTypeEnumToProto converts a InstanceDisksTypeEnum enum to its proto representation.
func ComputeInstanceDisksTypeEnumToProto(e *compute.InstanceDisksTypeEnum) computepb.ComputeInstanceDisksTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceDisksTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceDisksTypeEnum_value["InstanceDisksTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceDisksTypeEnum(v)
	}
	return computepb.ComputeInstanceDisksTypeEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(e *compute.InstanceNetworkInterfacesAccessConfigsTypeEnum) computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum_value["InstanceNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceStatusEnumToProto converts a InstanceStatusEnum enum to its proto representation.
func ComputeInstanceStatusEnumToProto(e *compute.InstanceStatusEnum) computepb.ComputeInstanceStatusEnum {
	if e == nil {
		return computepb.ComputeInstanceStatusEnum(0)
	}
	if v, ok := computepb.ComputeInstanceStatusEnum_value["InstanceStatusEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceStatusEnum(v)
	}
	return computepb.ComputeInstanceStatusEnum(0)
}

// InstanceDisksToProto converts a InstanceDisks resource to its proto representation.
func ComputeInstanceDisksToProto(o *compute.InstanceDisks) *computepb.ComputeInstanceDisks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisks{
		AutoDelete:        dcl.ValueOrEmptyBool(o.AutoDelete),
		Boot:              dcl.ValueOrEmptyBool(o.Boot),
		DeviceName:        dcl.ValueOrEmptyString(o.DeviceName),
		DiskEncryptionKey: ComputeInstanceDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey),
		Index:             dcl.ValueOrEmptyInt64(o.Index),
		InitializeParams:  ComputeInstanceDisksInitializeParamsToProto(o.InitializeParams),
		Interface:         ComputeInstanceDisksInterfaceEnumToProto(o.Interface),
		Mode:              ComputeInstanceDisksModeEnumToProto(o.Mode),
		Source:            dcl.ValueOrEmptyString(o.Source),
		Type:              ComputeInstanceDisksTypeEnumToProto(o.Type),
	}
	return p
}

// InstanceDisksDiskEncryptionKeyToProto converts a InstanceDisksDiskEncryptionKey resource to its proto representation.
func ComputeInstanceDisksDiskEncryptionKeyToProto(o *compute.InstanceDisksDiskEncryptionKey) *computepb.ComputeInstanceDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisksDiskEncryptionKey{
		RawKey:          dcl.ValueOrEmptyString(o.RawKey),
		RsaEncryptedKey: dcl.ValueOrEmptyString(o.RsaEncryptedKey),
		Sha256:          dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// InstanceDisksInitializeParamsToProto converts a InstanceDisksInitializeParams resource to its proto representation.
func ComputeInstanceDisksInitializeParamsToProto(o *compute.InstanceDisksInitializeParams) *computepb.ComputeInstanceDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisksInitializeParams{
		DiskName:                 dcl.ValueOrEmptyString(o.DiskName),
		DiskSizeGb:               dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		DiskType:                 dcl.ValueOrEmptyString(o.DiskType),
		SourceImage:              dcl.ValueOrEmptyString(o.SourceImage),
		SourceImageEncryptionKey: ComputeInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey),
	}
	return p
}

// InstanceDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceDisksInitializeParamsSourceImageEncryptionKey resource to its proto representation.
func ComputeInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o *compute.InstanceDisksInitializeParamsSourceImageEncryptionKey) *computepb.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.ValueOrEmptyString(o.RawKey),
		Sha256: dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// InstanceGuestAcceleratorsToProto converts a InstanceGuestAccelerators resource to its proto representation.
func ComputeInstanceGuestAcceleratorsToProto(o *compute.InstanceGuestAccelerators) *computepb.ComputeInstanceGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGuestAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// InstanceNetworkInterfacesToProto converts a InstanceNetworkInterfaces resource to its proto representation.
func ComputeInstanceNetworkInterfacesToProto(o *compute.InstanceNetworkInterfaces) *computepb.ComputeInstanceNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfaces{
		Name:       dcl.ValueOrEmptyString(o.Name),
		Network:    dcl.ValueOrEmptyString(o.Network),
		NetworkIp:  dcl.ValueOrEmptyString(o.NetworkIP),
		Subnetwork: dcl.ValueOrEmptyString(o.Subnetwork),
	}
	for _, r := range o.AccessConfigs {
		p.AccessConfigs = append(p.AccessConfigs, ComputeInstanceNetworkInterfacesAccessConfigsToProto(&r))
	}
	for _, r := range o.AliasIPRanges {
		p.AliasIpRanges = append(p.AliasIpRanges, ComputeInstanceNetworkInterfacesAliasIPRangesToProto(&r))
	}
	return p
}

// InstanceNetworkInterfacesAccessConfigsToProto converts a InstanceNetworkInterfacesAccessConfigs resource to its proto representation.
func ComputeInstanceNetworkInterfacesAccessConfigsToProto(o *compute.InstanceNetworkInterfacesAccessConfigs) *computepb.ComputeInstanceNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfacesAccessConfigs{
		Name:  dcl.ValueOrEmptyString(o.Name),
		NatIp: dcl.ValueOrEmptyString(o.NatIP),
		Type:  ComputeInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type),
	}
	return p
}

// InstanceNetworkInterfacesAliasIPRangesToProto converts a InstanceNetworkInterfacesAliasIPRanges resource to its proto representation.
func ComputeInstanceNetworkInterfacesAliasIPRangesToProto(o *compute.InstanceNetworkInterfacesAliasIPRanges) *computepb.ComputeInstanceNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfacesAliasIPRanges{
		IpCidrRange:         dcl.ValueOrEmptyString(o.IPCidrRange),
		SubnetworkRangeName: dcl.ValueOrEmptyString(o.SubnetworkRangeName),
	}
	return p
}

// InstanceSchedulingToProto converts a InstanceScheduling resource to its proto representation.
func ComputeInstanceSchedulingToProto(o *compute.InstanceScheduling) *computepb.ComputeInstanceScheduling {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceScheduling{
		AutomaticRestart:  dcl.ValueOrEmptyBool(o.AutomaticRestart),
		OnHostMaintenance: dcl.ValueOrEmptyString(o.OnHostMaintenance),
		Preemptible:       dcl.ValueOrEmptyBool(o.Preemptible),
	}
	return p
}

// InstanceServiceAccountsToProto converts a InstanceServiceAccounts resource to its proto representation.
func ComputeInstanceServiceAccountsToProto(o *compute.InstanceServiceAccounts) *computepb.ComputeInstanceServiceAccounts {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceServiceAccounts{
		Email: dcl.ValueOrEmptyString(o.Email),
	}
	for _, r := range o.Scopes {
		p.Scopes = append(p.Scopes, r)
	}
	return p
}

// InstanceShieldedInstanceConfigToProto converts a InstanceShieldedInstanceConfig resource to its proto representation.
func ComputeInstanceShieldedInstanceConfigToProto(o *compute.InstanceShieldedInstanceConfig) *computepb.ComputeInstanceShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableVtpm:                dcl.ValueOrEmptyBool(o.EnableVtpm),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *compute.Instance) *computepb.ComputeInstance {
	p := &computepb.ComputeInstance{
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
		Scheduling:             ComputeInstanceSchedulingToProto(resource.Scheduling),
		ShieldedInstanceConfig: ComputeInstanceShieldedInstanceConfigToProto(resource.ShieldedInstanceConfig),
		Status:                 ComputeInstanceStatusEnumToProto(resource.Status),
		StatusMessage:          dcl.ValueOrEmptyString(resource.StatusMessage),
		Zone:                   dcl.ValueOrEmptyString(resource.Zone),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		SelfLink:               dcl.ValueOrEmptyString(resource.SelfLink),
	}
	for _, r := range resource.Disks {
		p.Disks = append(p.Disks, ComputeInstanceDisksToProto(&r))
	}
	for _, r := range resource.GuestAccelerators {
		p.GuestAccelerators = append(p.GuestAccelerators, ComputeInstanceGuestAcceleratorsToProto(&r))
	}
	for _, r := range resource.NetworkInterfaces {
		p.NetworkInterfaces = append(p.NetworkInterfaces, ComputeInstanceNetworkInterfacesToProto(&r))
	}
	for _, r := range resource.ServiceAccounts {
		p.ServiceAccounts = append(p.ServiceAccounts, ComputeInstanceServiceAccountsToProto(&r))
	}
	for _, r := range resource.Tags {
		p.Tags = append(p.Tags, r)
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInstanceRequest) (*computepb.ComputeInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyComputeInstance(ctx context.Context, request *computepb.ApplyComputeInstanceRequest) (*computepb.ComputeInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteComputeInstance(ctx context.Context, request *computepb.DeleteComputeInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListComputeInstance(ctx context.Context, request *computepb.ListComputeInstanceRequest) (*computepb.ListComputeInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Zone)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
