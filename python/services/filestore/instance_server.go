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
	filestorepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/filestore_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToFilestoreInstanceStateEnum(e filestorepb.FilestoreInstanceStateEnum) *filestore.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filestorepb.FilestoreInstanceStateEnum_name[int32(e)]; ok {
		e := filestore.InstanceStateEnum(n[len("FilestoreInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToFilestoreInstanceTierEnum(e filestorepb.FilestoreInstanceTierEnum) *filestore.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filestorepb.FilestoreInstanceTierEnum_name[int32(e)]; ok {
		e := filestore.InstanceTierEnum(n[len("FilestoreInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsAccessModeEnum converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum from its proto representation.
func ProtoToFilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum(e filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum) *filestore.InstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum_name[int32(e)]; ok {
		e := filestore.InstanceFileSharesNfsExportOptionsAccessModeEnum(n[len("FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsSquashModeEnum converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum from its proto representation.
func ProtoToFilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum(e filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum) *filestore.InstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum_name[int32(e)]; ok {
		e := filestore.InstanceFileSharesNfsExportOptionsSquashModeEnum(n[len("FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworksModesEnum converts a InstanceNetworksModesEnum enum from its proto representation.
func ProtoToFilestoreInstanceNetworksModesEnum(e filestorepb.FilestoreInstanceNetworksModesEnum) *filestore.InstanceNetworksModesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filestorepb.FilestoreInstanceNetworksModesEnum_name[int32(e)]; ok {
		e := filestore.InstanceNetworksModesEnum(n[len("FilestoreInstanceNetworksModesEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileShares converts a InstanceFileShares resource from its proto representation.
func ProtoToFilestoreInstanceFileShares(p *filestorepb.FilestoreInstanceFileShares) *filestore.InstanceFileShares {
	if p == nil {
		return nil
	}
	obj := &filestore.InstanceFileShares{
		Name:         dcl.StringOrNil(p.Name),
		CapacityGb:   dcl.Int64OrNil(p.CapacityGb),
		SourceBackup: dcl.StringOrNil(p.SourceBackup),
	}
	for _, r := range p.GetNfsExportOptions() {
		obj.NfsExportOptions = append(obj.NfsExportOptions, *ProtoToFilestoreInstanceFileSharesNfsExportOptions(r))
	}
	return obj
}

// ProtoToInstanceFileSharesNfsExportOptions converts a InstanceFileSharesNfsExportOptions resource from its proto representation.
func ProtoToFilestoreInstanceFileSharesNfsExportOptions(p *filestorepb.FilestoreInstanceFileSharesNfsExportOptions) *filestore.InstanceFileSharesNfsExportOptions {
	if p == nil {
		return nil
	}
	obj := &filestore.InstanceFileSharesNfsExportOptions{
		AccessMode: ProtoToFilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum(p.GetAccessMode()),
		SquashMode: ProtoToFilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum(p.GetSquashMode()),
		AnonUid:    dcl.Int64OrNil(p.AnonUid),
		AnonGid:    dcl.Int64OrNil(p.AnonGid),
	}
	for _, r := range p.GetIpRanges() {
		obj.IPRanges = append(obj.IPRanges, r)
	}
	return obj
}

// ProtoToInstanceNetworks converts a InstanceNetworks resource from its proto representation.
func ProtoToFilestoreInstanceNetworks(p *filestorepb.FilestoreInstanceNetworks) *filestore.InstanceNetworks {
	if p == nil {
		return nil
	}
	obj := &filestore.InstanceNetworks{
		Network:         dcl.StringOrNil(p.Network),
		ReservedIPRange: dcl.StringOrNil(p.ReservedIpRange),
	}
	for _, r := range p.GetModes() {
		obj.Modes = append(obj.Modes, *ProtoToFilestoreInstanceNetworksModesEnum(r))
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *filestorepb.FilestoreInstance) *filestore.Instance {
	obj := &filestore.Instance{
		Name:          dcl.StringOrNil(p.Name),
		Description:   dcl.StringOrNil(p.Description),
		State:         ProtoToFilestoreInstanceStateEnum(p.GetState()),
		StatusMessage: dcl.StringOrNil(p.StatusMessage),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Tier:          ProtoToFilestoreInstanceTierEnum(p.GetTier()),
		Etag:          dcl.StringOrNil(p.Etag),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetFileShares() {
		obj.FileShares = append(obj.FileShares, *ProtoToFilestoreInstanceFileShares(r))
	}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToFilestoreInstanceNetworks(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func FilestoreInstanceStateEnumToProto(e *filestore.InstanceStateEnum) filestorepb.FilestoreInstanceStateEnum {
	if e == nil {
		return filestorepb.FilestoreInstanceStateEnum(0)
	}
	if v, ok := filestorepb.FilestoreInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return filestorepb.FilestoreInstanceStateEnum(v)
	}
	return filestorepb.FilestoreInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func FilestoreInstanceTierEnumToProto(e *filestore.InstanceTierEnum) filestorepb.FilestoreInstanceTierEnum {
	if e == nil {
		return filestorepb.FilestoreInstanceTierEnum(0)
	}
	if v, ok := filestorepb.FilestoreInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return filestorepb.FilestoreInstanceTierEnum(v)
	}
	return filestorepb.FilestoreInstanceTierEnum(0)
}

// InstanceFileSharesNfsExportOptionsAccessModeEnumToProto converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum to its proto representation.
func FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(e *filestore.InstanceFileSharesNfsExportOptionsAccessModeEnum) filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == nil {
		return filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
	}
	if v, ok := filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum_value["InstanceFileSharesNfsExportOptionsAccessModeEnum"+string(*e)]; ok {
		return filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum(v)
	}
	return filestorepb.FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
}

// InstanceFileSharesNfsExportOptionsSquashModeEnumToProto converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum to its proto representation.
func FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(e *filestore.InstanceFileSharesNfsExportOptionsSquashModeEnum) filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == nil {
		return filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
	}
	if v, ok := filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum_value["InstanceFileSharesNfsExportOptionsSquashModeEnum"+string(*e)]; ok {
		return filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum(v)
	}
	return filestorepb.FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
}

// InstanceNetworksModesEnumToProto converts a InstanceNetworksModesEnum enum to its proto representation.
func FilestoreInstanceNetworksModesEnumToProto(e *filestore.InstanceNetworksModesEnum) filestorepb.FilestoreInstanceNetworksModesEnum {
	if e == nil {
		return filestorepb.FilestoreInstanceNetworksModesEnum(0)
	}
	if v, ok := filestorepb.FilestoreInstanceNetworksModesEnum_value["InstanceNetworksModesEnum"+string(*e)]; ok {
		return filestorepb.FilestoreInstanceNetworksModesEnum(v)
	}
	return filestorepb.FilestoreInstanceNetworksModesEnum(0)
}

// InstanceFileSharesToProto converts a InstanceFileShares resource to its proto representation.
func FilestoreInstanceFileSharesToProto(o *filestore.InstanceFileShares) *filestorepb.FilestoreInstanceFileShares {
	if o == nil {
		return nil
	}
	p := &filestorepb.FilestoreInstanceFileShares{
		Name:         dcl.ValueOrEmptyString(o.Name),
		CapacityGb:   dcl.ValueOrEmptyInt64(o.CapacityGb),
		SourceBackup: dcl.ValueOrEmptyString(o.SourceBackup),
	}
	for _, r := range o.NfsExportOptions {
		p.NfsExportOptions = append(p.NfsExportOptions, FilestoreInstanceFileSharesNfsExportOptionsToProto(&r))
	}
	return p
}

// InstanceFileSharesNfsExportOptionsToProto converts a InstanceFileSharesNfsExportOptions resource to its proto representation.
func FilestoreInstanceFileSharesNfsExportOptionsToProto(o *filestore.InstanceFileSharesNfsExportOptions) *filestorepb.FilestoreInstanceFileSharesNfsExportOptions {
	if o == nil {
		return nil
	}
	p := &filestorepb.FilestoreInstanceFileSharesNfsExportOptions{
		AccessMode: FilestoreInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(o.AccessMode),
		SquashMode: FilestoreInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(o.SquashMode),
		AnonUid:    dcl.ValueOrEmptyInt64(o.AnonUid),
		AnonGid:    dcl.ValueOrEmptyInt64(o.AnonGid),
	}
	for _, r := range o.IPRanges {
		p.IpRanges = append(p.IpRanges, r)
	}
	return p
}

// InstanceNetworksToProto converts a InstanceNetworks resource to its proto representation.
func FilestoreInstanceNetworksToProto(o *filestore.InstanceNetworks) *filestorepb.FilestoreInstanceNetworks {
	if o == nil {
		return nil
	}
	p := &filestorepb.FilestoreInstanceNetworks{
		Network:         dcl.ValueOrEmptyString(o.Network),
		ReservedIpRange: dcl.ValueOrEmptyString(o.ReservedIPRange),
	}
	for _, r := range o.Modes {
		p.Modes = append(p.Modes, filestorepb.FilestoreInstanceNetworksModesEnum(filestorepb.FilestoreInstanceNetworksModesEnum_value[string(r)]))
	}
	for _, r := range o.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *filestore.Instance) *filestorepb.FilestoreInstance {
	p := &filestorepb.FilestoreInstance{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Description:   dcl.ValueOrEmptyString(resource.Description),
		State:         FilestoreInstanceStateEnumToProto(resource.State),
		StatusMessage: dcl.ValueOrEmptyString(resource.StatusMessage),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		Tier:          FilestoreInstanceTierEnumToProto(resource.Tier),
		Etag:          dcl.ValueOrEmptyString(resource.Etag),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.FileShares {
		p.FileShares = append(p.FileShares, FilestoreInstanceFileSharesToProto(&r))
	}
	for _, r := range resource.Networks {
		p.Networks = append(p.Networks, FilestoreInstanceNetworksToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *filestore.Client, request *filestorepb.ApplyFilestoreInstanceRequest) (*filestorepb.FilestoreInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyFilestoreInstance(ctx context.Context, request *filestorepb.ApplyFilestoreInstanceRequest) (*filestorepb.FilestoreInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteFilestoreInstance(ctx context.Context, request *filestorepb.DeleteFilestoreInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListFilestoreInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListFilestoreInstance(ctx context.Context, request *filestorepb.ListFilestoreInstanceRequest) (*filestorepb.ListFilestoreInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*filestorepb.FilestoreInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &filestorepb.ListFilestoreInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*filestore.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return filestore.NewClient(conf), nil
}
