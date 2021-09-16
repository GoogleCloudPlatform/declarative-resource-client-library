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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/alpha/filestore_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/alpha"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceStateEnum(e alphapb.FilestoreAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("FilestoreAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceTierEnum(e alphapb.FilestoreAlphaInstanceTierEnum) *alpha.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceTierEnum_name[int32(e)]; ok {
		e := alpha.InstanceTierEnum(n[len("FilestoreAlphaInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsAccessModeEnum converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(e alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum) *alpha.InstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFileSharesNfsExportOptionsAccessModeEnum(n[len("FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsSquashModeEnum converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(e alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum) *alpha.InstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFileSharesNfsExportOptionsSquashModeEnum(n[len("FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworksModesEnum converts a InstanceNetworksModesEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceNetworksModesEnum(e alphapb.FilestoreAlphaInstanceNetworksModesEnum) *alpha.InstanceNetworksModesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceNetworksModesEnum_name[int32(e)]; ok {
		e := alpha.InstanceNetworksModesEnum(n[len("FilestoreAlphaInstanceNetworksModesEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileShares converts a InstanceFileShares resource from its proto representation.
func ProtoToFilestoreAlphaInstanceFileShares(p *alphapb.FilestoreAlphaInstanceFileShares) *alpha.InstanceFileShares {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFileShares{
		Name:         dcl.StringOrNil(p.Name),
		CapacityGb:   dcl.Int64OrNil(p.CapacityGb),
		SourceBackup: dcl.StringOrNil(p.SourceBackup),
	}
	for _, r := range p.GetNfsExportOptions() {
		obj.NfsExportOptions = append(obj.NfsExportOptions, *ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptions(r))
	}
	return obj
}

// ProtoToInstanceFileSharesNfsExportOptions converts a InstanceFileSharesNfsExportOptions resource from its proto representation.
func ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptions(p *alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions) *alpha.InstanceFileSharesNfsExportOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFileSharesNfsExportOptions{
		AccessMode: ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(p.GetAccessMode()),
		SquashMode: ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(p.GetSquashMode()),
		AnonUid:    dcl.Int64OrNil(p.AnonUid),
		AnonGid:    dcl.Int64OrNil(p.AnonGid),
	}
	for _, r := range p.GetIpRanges() {
		obj.IPRanges = append(obj.IPRanges, r)
	}
	return obj
}

// ProtoToInstanceNetworks converts a InstanceNetworks resource from its proto representation.
func ProtoToFilestoreAlphaInstanceNetworks(p *alphapb.FilestoreAlphaInstanceNetworks) *alpha.InstanceNetworks {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworks{
		Network:         dcl.StringOrNil(p.Network),
		ReservedIPRange: dcl.StringOrNil(p.ReservedIpRange),
	}
	for _, r := range p.GetModes() {
		obj.Modes = append(obj.Modes, *ProtoToFilestoreAlphaInstanceNetworksModesEnum(r))
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.FilestoreAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:          dcl.StringOrNil(p.Name),
		Description:   dcl.StringOrNil(p.Description),
		State:         ProtoToFilestoreAlphaInstanceStateEnum(p.GetState()),
		StatusMessage: dcl.StringOrNil(p.StatusMessage),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Tier:          ProtoToFilestoreAlphaInstanceTierEnum(p.GetTier()),
		Etag:          dcl.StringOrNil(p.Etag),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetFileShares() {
		obj.FileShares = append(obj.FileShares, *ProtoToFilestoreAlphaInstanceFileShares(r))
	}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToFilestoreAlphaInstanceNetworks(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func FilestoreAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.FilestoreAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceStateEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func FilestoreAlphaInstanceTierEnumToProto(e *alpha.InstanceTierEnum) alphapb.FilestoreAlphaInstanceTierEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceTierEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceTierEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceTierEnum(0)
}

// InstanceFileSharesNfsExportOptionsAccessModeEnumToProto converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum to its proto representation.
func FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(e *alpha.InstanceFileSharesNfsExportOptionsAccessModeEnum) alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum_value["InstanceFileSharesNfsExportOptionsAccessModeEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
}

// InstanceFileSharesNfsExportOptionsSquashModeEnumToProto converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum to its proto representation.
func FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(e *alpha.InstanceFileSharesNfsExportOptionsSquashModeEnum) alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum_value["InstanceFileSharesNfsExportOptionsSquashModeEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
}

// InstanceNetworksModesEnumToProto converts a InstanceNetworksModesEnum enum to its proto representation.
func FilestoreAlphaInstanceNetworksModesEnumToProto(e *alpha.InstanceNetworksModesEnum) alphapb.FilestoreAlphaInstanceNetworksModesEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceNetworksModesEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceNetworksModesEnum_value["InstanceNetworksModesEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceNetworksModesEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceNetworksModesEnum(0)
}

// InstanceFileSharesToProto converts a InstanceFileShares resource to its proto representation.
func FilestoreAlphaInstanceFileSharesToProto(o *alpha.InstanceFileShares) *alphapb.FilestoreAlphaInstanceFileShares {
	if o == nil {
		return nil
	}
	p := &alphapb.FilestoreAlphaInstanceFileShares{
		Name:         dcl.ValueOrEmptyString(o.Name),
		CapacityGb:   dcl.ValueOrEmptyInt64(o.CapacityGb),
		SourceBackup: dcl.ValueOrEmptyString(o.SourceBackup),
	}
	for _, r := range o.NfsExportOptions {
		p.NfsExportOptions = append(p.NfsExportOptions, FilestoreAlphaInstanceFileSharesNfsExportOptionsToProto(&r))
	}
	return p
}

// InstanceFileSharesNfsExportOptionsToProto converts a InstanceFileSharesNfsExportOptions resource to its proto representation.
func FilestoreAlphaInstanceFileSharesNfsExportOptionsToProto(o *alpha.InstanceFileSharesNfsExportOptions) *alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions{
		AccessMode: FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(o.AccessMode),
		SquashMode: FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(o.SquashMode),
		AnonUid:    dcl.ValueOrEmptyInt64(o.AnonUid),
		AnonGid:    dcl.ValueOrEmptyInt64(o.AnonGid),
	}
	for _, r := range o.IPRanges {
		p.IpRanges = append(p.IpRanges, r)
	}
	return p
}

// InstanceNetworksToProto converts a InstanceNetworks resource to its proto representation.
func FilestoreAlphaInstanceNetworksToProto(o *alpha.InstanceNetworks) *alphapb.FilestoreAlphaInstanceNetworks {
	if o == nil {
		return nil
	}
	p := &alphapb.FilestoreAlphaInstanceNetworks{
		Network:         dcl.ValueOrEmptyString(o.Network),
		ReservedIpRange: dcl.ValueOrEmptyString(o.ReservedIPRange),
	}
	for _, r := range o.Modes {
		p.Modes = append(p.Modes, alphapb.FilestoreAlphaInstanceNetworksModesEnum(alphapb.FilestoreAlphaInstanceNetworksModesEnum_value[string(r)]))
	}
	for _, r := range o.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.FilestoreAlphaInstance {
	p := &alphapb.FilestoreAlphaInstance{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Description:   dcl.ValueOrEmptyString(resource.Description),
		State:         FilestoreAlphaInstanceStateEnumToProto(resource.State),
		StatusMessage: dcl.ValueOrEmptyString(resource.StatusMessage),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		Tier:          FilestoreAlphaInstanceTierEnumToProto(resource.Tier),
		Etag:          dcl.ValueOrEmptyString(resource.Etag),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.FileShares {
		p.FileShares = append(p.FileShares, FilestoreAlphaInstanceFileSharesToProto(&r))
	}
	for _, r := range resource.Networks {
		p.Networks = append(p.Networks, FilestoreAlphaInstanceNetworksToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFilestoreAlphaInstanceRequest) (*alphapb.FilestoreAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyFilestoreAlphaInstance(ctx context.Context, request *alphapb.ApplyFilestoreAlphaInstanceRequest) (*alphapb.FilestoreAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteFilestoreAlphaInstance(ctx context.Context, request *alphapb.DeleteFilestoreAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListFilestoreAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListFilestoreAlphaInstance(ctx context.Context, request *alphapb.ListFilestoreAlphaInstanceRequest) (*alphapb.ListFilestoreAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, ProtoToInstance(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FilestoreAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListFilestoreAlphaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
