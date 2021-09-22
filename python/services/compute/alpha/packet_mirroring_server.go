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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// Server implements the gRPC interface for PacketMirroring.
type PacketMirroringServer struct{}

// ProtoToPacketMirroringFilterDirectionEnum converts a PacketMirroringFilterDirectionEnum enum from its proto representation.
func ProtoToComputeAlphaPacketMirroringFilterDirectionEnum(e alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum) *alpha.PacketMirroringFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum_name[int32(e)]; ok {
		e := alpha.PacketMirroringFilterDirectionEnum(n[len("ComputeAlphaPacketMirroringFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringEnableEnum converts a PacketMirroringEnableEnum enum from its proto representation.
func ProtoToComputeAlphaPacketMirroringEnableEnum(e alphapb.ComputeAlphaPacketMirroringEnableEnum) *alpha.PacketMirroringEnableEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaPacketMirroringEnableEnum_name[int32(e)]; ok {
		e := alpha.PacketMirroringEnableEnum(n[len("ComputeAlphaPacketMirroringEnableEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringNetwork converts a PacketMirroringNetwork resource from its proto representation.
func ProtoToComputeAlphaPacketMirroringNetwork(p *alphapb.ComputeAlphaPacketMirroringNetwork) *alpha.PacketMirroringNetwork {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringNetwork{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringCollectorIlb converts a PacketMirroringCollectorIlb resource from its proto representation.
func ProtoToComputeAlphaPacketMirroringCollectorIlb(p *alphapb.ComputeAlphaPacketMirroringCollectorIlb) *alpha.PacketMirroringCollectorIlb {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringCollectorIlb{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResources converts a PacketMirroringMirroredResources resource from its proto representation.
func ProtoToComputeAlphaPacketMirroringMirroredResources(p *alphapb.ComputeAlphaPacketMirroringMirroredResources) *alpha.PacketMirroringMirroredResources {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringMirroredResources{}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputeAlphaPacketMirroringMirroredResourcesSubnetworks(r))
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToComputeAlphaPacketMirroringMirroredResourcesInstances(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesSubnetworks converts a PacketMirroringMirroredResourcesSubnetworks resource from its proto representation.
func ProtoToComputeAlphaPacketMirroringMirroredResourcesSubnetworks(p *alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks) *alpha.PacketMirroringMirroredResourcesSubnetworks {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesInstances converts a PacketMirroringMirroredResourcesInstances resource from its proto representation.
func ProtoToComputeAlphaPacketMirroringMirroredResourcesInstances(p *alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances) *alpha.PacketMirroringMirroredResourcesInstances {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringMirroredResourcesInstances{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringFilter converts a PacketMirroringFilter resource from its proto representation.
func ProtoToComputeAlphaPacketMirroringFilter(p *alphapb.ComputeAlphaPacketMirroringFilter) *alpha.PacketMirroringFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringFilter{
		Direction: ProtoToComputeAlphaPacketMirroringFilterDirectionEnum(p.GetDirection()),
	}
	for _, r := range p.GetCidrRanges() {
		obj.CidrRanges = append(obj.CidrRanges, r)
	}
	for _, r := range p.GetIpProtocols() {
		obj.IPProtocols = append(obj.IPProtocols, r)
	}
	return obj
}

// ProtoToPacketMirroring converts a PacketMirroring resource from its proto representation.
func ProtoToPacketMirroring(p *alphapb.ComputeAlphaPacketMirroring) *alpha.PacketMirroring {
	obj := &alpha.PacketMirroring{
		Id:                dcl.Int64OrNil(p.Id),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Region:            dcl.StringOrNil(p.Region),
		Network:           ProtoToComputeAlphaPacketMirroringNetwork(p.GetNetwork()),
		Priority:          dcl.Int64OrNil(p.Priority),
		CollectorIlb:      ProtoToComputeAlphaPacketMirroringCollectorIlb(p.GetCollectorIlb()),
		MirroredResources: ProtoToComputeAlphaPacketMirroringMirroredResources(p.GetMirroredResources()),
		Filter:            ProtoToComputeAlphaPacketMirroringFilter(p.GetFilter()),
		Enable:            ProtoToComputeAlphaPacketMirroringEnableEnum(p.GetEnable()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	return obj
}

// PacketMirroringFilterDirectionEnumToProto converts a PacketMirroringFilterDirectionEnum enum to its proto representation.
func ComputeAlphaPacketMirroringFilterDirectionEnumToProto(e *alpha.PacketMirroringFilterDirectionEnum) alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum {
	if e == nil {
		return alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum_value["PacketMirroringFilterDirectionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum(v)
	}
	return alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum(0)
}

// PacketMirroringEnableEnumToProto converts a PacketMirroringEnableEnum enum to its proto representation.
func ComputeAlphaPacketMirroringEnableEnumToProto(e *alpha.PacketMirroringEnableEnum) alphapb.ComputeAlphaPacketMirroringEnableEnum {
	if e == nil {
		return alphapb.ComputeAlphaPacketMirroringEnableEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaPacketMirroringEnableEnum_value["PacketMirroringEnableEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaPacketMirroringEnableEnum(v)
	}
	return alphapb.ComputeAlphaPacketMirroringEnableEnum(0)
}

// PacketMirroringNetworkToProto converts a PacketMirroringNetwork resource to its proto representation.
func ComputeAlphaPacketMirroringNetworkToProto(o *alpha.PacketMirroringNetwork) *alphapb.ComputeAlphaPacketMirroringNetwork {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringNetwork{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringCollectorIlbToProto converts a PacketMirroringCollectorIlb resource to its proto representation.
func ComputeAlphaPacketMirroringCollectorIlbToProto(o *alpha.PacketMirroringCollectorIlb) *alphapb.ComputeAlphaPacketMirroringCollectorIlb {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringCollectorIlb{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringMirroredResourcesToProto converts a PacketMirroringMirroredResources resource to its proto representation.
func ComputeAlphaPacketMirroringMirroredResourcesToProto(o *alpha.PacketMirroringMirroredResources) *alphapb.ComputeAlphaPacketMirroringMirroredResources {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringMirroredResources{}
	for _, r := range o.Subnetworks {
		p.Subnetworks = append(p.Subnetworks, ComputeAlphaPacketMirroringMirroredResourcesSubnetworksToProto(&r))
	}
	for _, r := range o.Instances {
		p.Instances = append(p.Instances, ComputeAlphaPacketMirroringMirroredResourcesInstancesToProto(&r))
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	return p
}

// PacketMirroringMirroredResourcesSubnetworksToProto converts a PacketMirroringMirroredResourcesSubnetworks resource to its proto representation.
func ComputeAlphaPacketMirroringMirroredResourcesSubnetworksToProto(o *alpha.PacketMirroringMirroredResourcesSubnetworks) *alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringMirroredResourcesInstancesToProto converts a PacketMirroringMirroredResourcesInstances resource to its proto representation.
func ComputeAlphaPacketMirroringMirroredResourcesInstancesToProto(o *alpha.PacketMirroringMirroredResourcesInstances) *alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringFilterToProto converts a PacketMirroringFilter resource to its proto representation.
func ComputeAlphaPacketMirroringFilterToProto(o *alpha.PacketMirroringFilter) *alphapb.ComputeAlphaPacketMirroringFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringFilter{
		Direction: ComputeAlphaPacketMirroringFilterDirectionEnumToProto(o.Direction),
	}
	for _, r := range o.CidrRanges {
		p.CidrRanges = append(p.CidrRanges, r)
	}
	for _, r := range o.IPProtocols {
		p.IpProtocols = append(p.IpProtocols, r)
	}
	return p
}

// PacketMirroringToProto converts a PacketMirroring resource to its proto representation.
func PacketMirroringToProto(resource *alpha.PacketMirroring) *alphapb.ComputeAlphaPacketMirroring {
	p := &alphapb.ComputeAlphaPacketMirroring{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		Network:           ComputeAlphaPacketMirroringNetworkToProto(resource.Network),
		Priority:          dcl.ValueOrEmptyInt64(resource.Priority),
		CollectorIlb:      ComputeAlphaPacketMirroringCollectorIlbToProto(resource.CollectorIlb),
		MirroredResources: ComputeAlphaPacketMirroringMirroredResourcesToProto(resource.MirroredResources),
		Filter:            ComputeAlphaPacketMirroringFilterToProto(resource.Filter),
		Enable:            ComputeAlphaPacketMirroringEnableEnumToProto(resource.Enable),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) applyPacketMirroring(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaPacketMirroringRequest) (*alphapb.ComputeAlphaPacketMirroring, error) {
	p := ProtoToPacketMirroring(request.GetResource())
	res, err := c.ApplyPacketMirroring(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PacketMirroringToProto(res)
	return r, nil
}

// ApplyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) ApplyComputeAlphaPacketMirroring(ctx context.Context, request *alphapb.ApplyComputeAlphaPacketMirroringRequest) (*alphapb.ComputeAlphaPacketMirroring, error) {
	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPacketMirroring(ctx, cl, request)
}

// DeletePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Delete() method.
func (s *PacketMirroringServer) DeleteComputeAlphaPacketMirroring(ctx context.Context, request *alphapb.DeleteComputeAlphaPacketMirroringRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePacketMirroring(ctx, ProtoToPacketMirroring(request.GetResource()))

}

// ListComputeAlphaPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroringList() method.
func (s *PacketMirroringServer) ListComputeAlphaPacketMirroring(ctx context.Context, request *alphapb.ListComputeAlphaPacketMirroringRequest) (*alphapb.ListComputeAlphaPacketMirroringResponse, error) {
	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPacketMirroring(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaPacketMirroring
	for _, r := range resources.Items {
		rp := PacketMirroringToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListComputeAlphaPacketMirroringResponse{Items: protos}, nil
}

func createConfigPacketMirroring(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
