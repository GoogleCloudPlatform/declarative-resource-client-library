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

// Server implements the gRPC interface for PacketMirroring.
type PacketMirroringServer struct{}

// ProtoToPacketMirroringFilterDirectionEnum converts a PacketMirroringFilterDirectionEnum enum from its proto representation.
func ProtoToComputeBetaPacketMirroringFilterDirectionEnum(e betapb.ComputeBetaPacketMirroringFilterDirectionEnum) *beta.PacketMirroringFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaPacketMirroringFilterDirectionEnum_name[int32(e)]; ok {
		e := beta.PacketMirroringFilterDirectionEnum(n[len("ComputeBetaPacketMirroringFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringEnableEnum converts a PacketMirroringEnableEnum enum from its proto representation.
func ProtoToComputeBetaPacketMirroringEnableEnum(e betapb.ComputeBetaPacketMirroringEnableEnum) *beta.PacketMirroringEnableEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaPacketMirroringEnableEnum_name[int32(e)]; ok {
		e := beta.PacketMirroringEnableEnum(n[len("ComputeBetaPacketMirroringEnableEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringNetwork converts a PacketMirroringNetwork resource from its proto representation.
func ProtoToComputeBetaPacketMirroringNetwork(p *betapb.ComputeBetaPacketMirroringNetwork) *beta.PacketMirroringNetwork {
	if p == nil {
		return nil
	}
	obj := &beta.PacketMirroringNetwork{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringCollectorIlb converts a PacketMirroringCollectorIlb resource from its proto representation.
func ProtoToComputeBetaPacketMirroringCollectorIlb(p *betapb.ComputeBetaPacketMirroringCollectorIlb) *beta.PacketMirroringCollectorIlb {
	if p == nil {
		return nil
	}
	obj := &beta.PacketMirroringCollectorIlb{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResources converts a PacketMirroringMirroredResources resource from its proto representation.
func ProtoToComputeBetaPacketMirroringMirroredResources(p *betapb.ComputeBetaPacketMirroringMirroredResources) *beta.PacketMirroringMirroredResources {
	if p == nil {
		return nil
	}
	obj := &beta.PacketMirroringMirroredResources{}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputeBetaPacketMirroringMirroredResourcesSubnetworks(r))
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToComputeBetaPacketMirroringMirroredResourcesInstances(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesSubnetworks converts a PacketMirroringMirroredResourcesSubnetworks resource from its proto representation.
func ProtoToComputeBetaPacketMirroringMirroredResourcesSubnetworks(p *betapb.ComputeBetaPacketMirroringMirroredResourcesSubnetworks) *beta.PacketMirroringMirroredResourcesSubnetworks {
	if p == nil {
		return nil
	}
	obj := &beta.PacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesInstances converts a PacketMirroringMirroredResourcesInstances resource from its proto representation.
func ProtoToComputeBetaPacketMirroringMirroredResourcesInstances(p *betapb.ComputeBetaPacketMirroringMirroredResourcesInstances) *beta.PacketMirroringMirroredResourcesInstances {
	if p == nil {
		return nil
	}
	obj := &beta.PacketMirroringMirroredResourcesInstances{
		Url:          dcl.StringOrNil(p.Url),
		CanonicalUrl: dcl.StringOrNil(p.CanonicalUrl),
	}
	return obj
}

// ProtoToPacketMirroringFilter converts a PacketMirroringFilter resource from its proto representation.
func ProtoToComputeBetaPacketMirroringFilter(p *betapb.ComputeBetaPacketMirroringFilter) *beta.PacketMirroringFilter {
	if p == nil {
		return nil
	}
	obj := &beta.PacketMirroringFilter{
		Direction: ProtoToComputeBetaPacketMirroringFilterDirectionEnum(p.GetDirection()),
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
func ProtoToPacketMirroring(p *betapb.ComputeBetaPacketMirroring) *beta.PacketMirroring {
	obj := &beta.PacketMirroring{
		Id:                dcl.Int64OrNil(p.Id),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Region:            dcl.StringOrNil(p.Region),
		Network:           ProtoToComputeBetaPacketMirroringNetwork(p.GetNetwork()),
		Priority:          dcl.Int64OrNil(p.Priority),
		CollectorIlb:      ProtoToComputeBetaPacketMirroringCollectorIlb(p.GetCollectorIlb()),
		MirroredResources: ProtoToComputeBetaPacketMirroringMirroredResources(p.GetMirroredResources()),
		Filter:            ProtoToComputeBetaPacketMirroringFilter(p.GetFilter()),
		Enable:            ProtoToComputeBetaPacketMirroringEnableEnum(p.GetEnable()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	return obj
}

// PacketMirroringFilterDirectionEnumToProto converts a PacketMirroringFilterDirectionEnum enum to its proto representation.
func ComputeBetaPacketMirroringFilterDirectionEnumToProto(e *beta.PacketMirroringFilterDirectionEnum) betapb.ComputeBetaPacketMirroringFilterDirectionEnum {
	if e == nil {
		return betapb.ComputeBetaPacketMirroringFilterDirectionEnum(0)
	}
	if v, ok := betapb.ComputeBetaPacketMirroringFilterDirectionEnum_value["PacketMirroringFilterDirectionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaPacketMirroringFilterDirectionEnum(v)
	}
	return betapb.ComputeBetaPacketMirroringFilterDirectionEnum(0)
}

// PacketMirroringEnableEnumToProto converts a PacketMirroringEnableEnum enum to its proto representation.
func ComputeBetaPacketMirroringEnableEnumToProto(e *beta.PacketMirroringEnableEnum) betapb.ComputeBetaPacketMirroringEnableEnum {
	if e == nil {
		return betapb.ComputeBetaPacketMirroringEnableEnum(0)
	}
	if v, ok := betapb.ComputeBetaPacketMirroringEnableEnum_value["PacketMirroringEnableEnum"+string(*e)]; ok {
		return betapb.ComputeBetaPacketMirroringEnableEnum(v)
	}
	return betapb.ComputeBetaPacketMirroringEnableEnum(0)
}

// PacketMirroringNetworkToProto converts a PacketMirroringNetwork resource to its proto representation.
func ComputeBetaPacketMirroringNetworkToProto(o *beta.PacketMirroringNetwork) *betapb.ComputeBetaPacketMirroringNetwork {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaPacketMirroringNetwork{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringCollectorIlbToProto converts a PacketMirroringCollectorIlb resource to its proto representation.
func ComputeBetaPacketMirroringCollectorIlbToProto(o *beta.PacketMirroringCollectorIlb) *betapb.ComputeBetaPacketMirroringCollectorIlb {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaPacketMirroringCollectorIlb{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringMirroredResourcesToProto converts a PacketMirroringMirroredResources resource to its proto representation.
func ComputeBetaPacketMirroringMirroredResourcesToProto(o *beta.PacketMirroringMirroredResources) *betapb.ComputeBetaPacketMirroringMirroredResources {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaPacketMirroringMirroredResources{}
	for _, r := range o.Subnetworks {
		p.Subnetworks = append(p.Subnetworks, ComputeBetaPacketMirroringMirroredResourcesSubnetworksToProto(&r))
	}
	for _, r := range o.Instances {
		p.Instances = append(p.Instances, ComputeBetaPacketMirroringMirroredResourcesInstancesToProto(&r))
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	return p
}

// PacketMirroringMirroredResourcesSubnetworksToProto converts a PacketMirroringMirroredResourcesSubnetworks resource to its proto representation.
func ComputeBetaPacketMirroringMirroredResourcesSubnetworksToProto(o *beta.PacketMirroringMirroredResourcesSubnetworks) *betapb.ComputeBetaPacketMirroringMirroredResourcesSubnetworks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaPacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringMirroredResourcesInstancesToProto converts a PacketMirroringMirroredResourcesInstances resource to its proto representation.
func ComputeBetaPacketMirroringMirroredResourcesInstancesToProto(o *beta.PacketMirroringMirroredResourcesInstances) *betapb.ComputeBetaPacketMirroringMirroredResourcesInstances {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaPacketMirroringMirroredResourcesInstances{
		Url:          dcl.ValueOrEmptyString(o.Url),
		CanonicalUrl: dcl.ValueOrEmptyString(o.CanonicalUrl),
	}
	return p
}

// PacketMirroringFilterToProto converts a PacketMirroringFilter resource to its proto representation.
func ComputeBetaPacketMirroringFilterToProto(o *beta.PacketMirroringFilter) *betapb.ComputeBetaPacketMirroringFilter {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaPacketMirroringFilter{
		Direction: ComputeBetaPacketMirroringFilterDirectionEnumToProto(o.Direction),
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
func PacketMirroringToProto(resource *beta.PacketMirroring) *betapb.ComputeBetaPacketMirroring {
	p := &betapb.ComputeBetaPacketMirroring{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		Network:           ComputeBetaPacketMirroringNetworkToProto(resource.Network),
		Priority:          dcl.ValueOrEmptyInt64(resource.Priority),
		CollectorIlb:      ComputeBetaPacketMirroringCollectorIlbToProto(resource.CollectorIlb),
		MirroredResources: ComputeBetaPacketMirroringMirroredResourcesToProto(resource.MirroredResources),
		Filter:            ComputeBetaPacketMirroringFilterToProto(resource.Filter),
		Enable:            ComputeBetaPacketMirroringEnableEnumToProto(resource.Enable),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) applyPacketMirroring(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaPacketMirroringRequest) (*betapb.ComputeBetaPacketMirroring, error) {
	p := ProtoToPacketMirroring(request.GetResource())
	res, err := c.ApplyPacketMirroring(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PacketMirroringToProto(res)
	return r, nil
}

// ApplyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) ApplyComputeBetaPacketMirroring(ctx context.Context, request *betapb.ApplyComputeBetaPacketMirroringRequest) (*betapb.ComputeBetaPacketMirroring, error) {
	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPacketMirroring(ctx, cl, request)
}

// DeletePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Delete() method.
func (s *PacketMirroringServer) DeleteComputeBetaPacketMirroring(ctx context.Context, request *betapb.DeleteComputeBetaPacketMirroringRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePacketMirroring(ctx, ProtoToPacketMirroring(request.GetResource()))

}

// ListComputeBetaPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroringList() method.
func (s *PacketMirroringServer) ListComputeBetaPacketMirroring(ctx context.Context, request *betapb.ListComputeBetaPacketMirroringRequest) (*betapb.ListComputeBetaPacketMirroringResponse, error) {
	cl, err := createConfigPacketMirroring(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPacketMirroring(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaPacketMirroring
	for _, r := range resources.Items {
		rp := PacketMirroringToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaPacketMirroringResponse{Items: protos}, nil
}

func createConfigPacketMirroring(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
