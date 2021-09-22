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

// Server implements the gRPC interface for PacketMirroring.
type PacketMirroringServer struct{}

// ProtoToPacketMirroringFilterDirectionEnum converts a PacketMirroringFilterDirectionEnum enum from its proto representation.
func ProtoToComputePacketMirroringFilterDirectionEnum(e computepb.ComputePacketMirroringFilterDirectionEnum) *compute.PacketMirroringFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputePacketMirroringFilterDirectionEnum_name[int32(e)]; ok {
		e := compute.PacketMirroringFilterDirectionEnum(n[len("ComputePacketMirroringFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringEnableEnum converts a PacketMirroringEnableEnum enum from its proto representation.
func ProtoToComputePacketMirroringEnableEnum(e computepb.ComputePacketMirroringEnableEnum) *compute.PacketMirroringEnableEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputePacketMirroringEnableEnum_name[int32(e)]; ok {
		e := compute.PacketMirroringEnableEnum(n[len("ComputePacketMirroringEnableEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringNetwork converts a PacketMirroringNetwork resource from its proto representation.
func ProtoToComputePacketMirroringNetwork(p *computepb.ComputePacketMirroringNetwork) *compute.PacketMirroringNetwork {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringNetwork{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringCollectorIlb converts a PacketMirroringCollectorIlb resource from its proto representation.
func ProtoToComputePacketMirroringCollectorIlb(p *computepb.ComputePacketMirroringCollectorIlb) *compute.PacketMirroringCollectorIlb {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringCollectorIlb{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResources converts a PacketMirroringMirroredResources resource from its proto representation.
func ProtoToComputePacketMirroringMirroredResources(p *computepb.ComputePacketMirroringMirroredResources) *compute.PacketMirroringMirroredResources {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringMirroredResources{}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputePacketMirroringMirroredResourcesSubnetworks(r))
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToComputePacketMirroringMirroredResourcesInstances(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesSubnetworks converts a PacketMirroringMirroredResourcesSubnetworks resource from its proto representation.
func ProtoToComputePacketMirroringMirroredResourcesSubnetworks(p *computepb.ComputePacketMirroringMirroredResourcesSubnetworks) *compute.PacketMirroringMirroredResourcesSubnetworks {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesInstances converts a PacketMirroringMirroredResourcesInstances resource from its proto representation.
func ProtoToComputePacketMirroringMirroredResourcesInstances(p *computepb.ComputePacketMirroringMirroredResourcesInstances) *compute.PacketMirroringMirroredResourcesInstances {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringMirroredResourcesInstances{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringFilter converts a PacketMirroringFilter resource from its proto representation.
func ProtoToComputePacketMirroringFilter(p *computepb.ComputePacketMirroringFilter) *compute.PacketMirroringFilter {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringFilter{
		Direction: ProtoToComputePacketMirroringFilterDirectionEnum(p.GetDirection()),
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
func ProtoToPacketMirroring(p *computepb.ComputePacketMirroring) *compute.PacketMirroring {
	obj := &compute.PacketMirroring{
		Id:                dcl.Int64OrNil(p.Id),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Region:            dcl.StringOrNil(p.Region),
		Network:           ProtoToComputePacketMirroringNetwork(p.GetNetwork()),
		Priority:          dcl.Int64OrNil(p.Priority),
		CollectorIlb:      ProtoToComputePacketMirroringCollectorIlb(p.GetCollectorIlb()),
		MirroredResources: ProtoToComputePacketMirroringMirroredResources(p.GetMirroredResources()),
		Filter:            ProtoToComputePacketMirroringFilter(p.GetFilter()),
		Enable:            ProtoToComputePacketMirroringEnableEnum(p.GetEnable()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	return obj
}

// PacketMirroringFilterDirectionEnumToProto converts a PacketMirroringFilterDirectionEnum enum to its proto representation.
func ComputePacketMirroringFilterDirectionEnumToProto(e *compute.PacketMirroringFilterDirectionEnum) computepb.ComputePacketMirroringFilterDirectionEnum {
	if e == nil {
		return computepb.ComputePacketMirroringFilterDirectionEnum(0)
	}
	if v, ok := computepb.ComputePacketMirroringFilterDirectionEnum_value["PacketMirroringFilterDirectionEnum"+string(*e)]; ok {
		return computepb.ComputePacketMirroringFilterDirectionEnum(v)
	}
	return computepb.ComputePacketMirroringFilterDirectionEnum(0)
}

// PacketMirroringEnableEnumToProto converts a PacketMirroringEnableEnum enum to its proto representation.
func ComputePacketMirroringEnableEnumToProto(e *compute.PacketMirroringEnableEnum) computepb.ComputePacketMirroringEnableEnum {
	if e == nil {
		return computepb.ComputePacketMirroringEnableEnum(0)
	}
	if v, ok := computepb.ComputePacketMirroringEnableEnum_value["PacketMirroringEnableEnum"+string(*e)]; ok {
		return computepb.ComputePacketMirroringEnableEnum(v)
	}
	return computepb.ComputePacketMirroringEnableEnum(0)
}

// PacketMirroringNetworkToProto converts a PacketMirroringNetwork resource to its proto representation.
func ComputePacketMirroringNetworkToProto(o *compute.PacketMirroringNetwork) *computepb.ComputePacketMirroringNetwork {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringNetwork{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringCollectorIlbToProto converts a PacketMirroringCollectorIlb resource to its proto representation.
func ComputePacketMirroringCollectorIlbToProto(o *compute.PacketMirroringCollectorIlb) *computepb.ComputePacketMirroringCollectorIlb {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringCollectorIlb{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringMirroredResourcesToProto converts a PacketMirroringMirroredResources resource to its proto representation.
func ComputePacketMirroringMirroredResourcesToProto(o *compute.PacketMirroringMirroredResources) *computepb.ComputePacketMirroringMirroredResources {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringMirroredResources{}
	for _, r := range o.Subnetworks {
		p.Subnetworks = append(p.Subnetworks, ComputePacketMirroringMirroredResourcesSubnetworksToProto(&r))
	}
	for _, r := range o.Instances {
		p.Instances = append(p.Instances, ComputePacketMirroringMirroredResourcesInstancesToProto(&r))
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	return p
}

// PacketMirroringMirroredResourcesSubnetworksToProto converts a PacketMirroringMirroredResourcesSubnetworks resource to its proto representation.
func ComputePacketMirroringMirroredResourcesSubnetworksToProto(o *compute.PacketMirroringMirroredResourcesSubnetworks) *computepb.ComputePacketMirroringMirroredResourcesSubnetworks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringMirroredResourcesInstancesToProto converts a PacketMirroringMirroredResourcesInstances resource to its proto representation.
func ComputePacketMirroringMirroredResourcesInstancesToProto(o *compute.PacketMirroringMirroredResourcesInstances) *computepb.ComputePacketMirroringMirroredResourcesInstances {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringMirroredResourcesInstances{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringFilterToProto converts a PacketMirroringFilter resource to its proto representation.
func ComputePacketMirroringFilterToProto(o *compute.PacketMirroringFilter) *computepb.ComputePacketMirroringFilter {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringFilter{
		Direction: ComputePacketMirroringFilterDirectionEnumToProto(o.Direction),
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
func PacketMirroringToProto(resource *compute.PacketMirroring) *computepb.ComputePacketMirroring {
	p := &computepb.ComputePacketMirroring{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		Network:           ComputePacketMirroringNetworkToProto(resource.Network),
		Priority:          dcl.ValueOrEmptyInt64(resource.Priority),
		CollectorIlb:      ComputePacketMirroringCollectorIlbToProto(resource.CollectorIlb),
		MirroredResources: ComputePacketMirroringMirroredResourcesToProto(resource.MirroredResources),
		Filter:            ComputePacketMirroringFilterToProto(resource.Filter),
		Enable:            ComputePacketMirroringEnableEnumToProto(resource.Enable),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) applyPacketMirroring(ctx context.Context, c *compute.Client, request *computepb.ApplyComputePacketMirroringRequest) (*computepb.ComputePacketMirroring, error) {
	p := ProtoToPacketMirroring(request.GetResource())
	res, err := c.ApplyPacketMirroring(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PacketMirroringToProto(res)
	return r, nil
}

// ApplyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) ApplyComputePacketMirroring(ctx context.Context, request *computepb.ApplyComputePacketMirroringRequest) (*computepb.ComputePacketMirroring, error) {
	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPacketMirroring(ctx, cl, request)
}

// DeletePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Delete() method.
func (s *PacketMirroringServer) DeleteComputePacketMirroring(ctx context.Context, request *computepb.DeleteComputePacketMirroringRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePacketMirroring(ctx, ProtoToPacketMirroring(request.GetResource()))

}

// ListComputePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroringList() method.
func (s *PacketMirroringServer) ListComputePacketMirroring(ctx context.Context, request *computepb.ListComputePacketMirroringRequest) (*computepb.ListComputePacketMirroringResponse, error) {
	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPacketMirroring(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputePacketMirroring
	for _, r := range resources.Items {
		rp := PacketMirroringToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputePacketMirroringResponse{Items: protos}, nil
}

func createConfigPacketMirroring(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
