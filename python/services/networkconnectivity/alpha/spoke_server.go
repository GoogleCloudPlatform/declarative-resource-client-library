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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkconnectivity/alpha/networkconnectivity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/alpha"
)

// Server implements the gRPC interface for Spoke.
type SpokeServer struct{}

// ProtoToSpokeStateEnum converts a SpokeStateEnum enum from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeStateEnum(e alphapb.NetworkconnectivityAlphaSpokeStateEnum) *alpha.SpokeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkconnectivityAlphaSpokeStateEnum_name[int32(e)]; ok {
		e := alpha.SpokeStateEnum(n[len("NetworkconnectivityAlphaSpokeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToSpokeLinkedRouterApplianceInstances converts a SpokeLinkedRouterApplianceInstances resource from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances(p *alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances) *alpha.SpokeLinkedRouterApplianceInstances {
	if p == nil {
		return nil
	}
	obj := &alpha.SpokeLinkedRouterApplianceInstances{
		VirtualMachine: dcl.StringOrNil(p.VirtualMachine),
		IPAddress:      dcl.StringOrNil(p.IpAddress),
	}
	return obj
}

// ProtoToSpoke converts a Spoke resource from its proto representation.
func ProtoToSpoke(p *alphapb.NetworkconnectivityAlphaSpoke) *alpha.Spoke {
	obj := &alpha.Spoke{
		Name:        dcl.StringOrNil(p.Name),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.Description),
		Hub:         dcl.StringOrNil(p.Hub),
		UniqueId:    dcl.StringOrNil(p.UniqueId),
		State:       ProtoToNetworkconnectivityAlphaSpokeStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.Project),
		Location:    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetLinkedVpnTunnels() {
		obj.LinkedVpnTunnels = append(obj.LinkedVpnTunnels, r)
	}
	for _, r := range p.GetLinkedInterconnectAttachments() {
		obj.LinkedInterconnectAttachments = append(obj.LinkedInterconnectAttachments, r)
	}
	for _, r := range p.GetLinkedRouterApplianceInstances() {
		obj.LinkedRouterApplianceInstances = append(obj.LinkedRouterApplianceInstances, *ProtoToNetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances(r))
	}
	return obj
}

// SpokeStateEnumToProto converts a SpokeStateEnum enum to its proto representation.
func NetworkconnectivityAlphaSpokeStateEnumToProto(e *alpha.SpokeStateEnum) alphapb.NetworkconnectivityAlphaSpokeStateEnum {
	if e == nil {
		return alphapb.NetworkconnectivityAlphaSpokeStateEnum(0)
	}
	if v, ok := alphapb.NetworkconnectivityAlphaSpokeStateEnum_value["SpokeStateEnum"+string(*e)]; ok {
		return alphapb.NetworkconnectivityAlphaSpokeStateEnum(v)
	}
	return alphapb.NetworkconnectivityAlphaSpokeStateEnum(0)
}

// SpokeLinkedRouterApplianceInstancesToProto converts a SpokeLinkedRouterApplianceInstances resource to its proto representation.
func NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesToProto(o *alpha.SpokeLinkedRouterApplianceInstances) *alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances{
		VirtualMachine: dcl.ValueOrEmptyString(o.VirtualMachine),
		IpAddress:      dcl.ValueOrEmptyString(o.IPAddress),
	}
	return p
}

// SpokeToProto converts a Spoke resource to its proto representation.
func SpokeToProto(resource *alpha.Spoke) *alphapb.NetworkconnectivityAlphaSpoke {
	p := &alphapb.NetworkconnectivityAlphaSpoke{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Hub:         dcl.ValueOrEmptyString(resource.Hub),
		UniqueId:    dcl.ValueOrEmptyString(resource.UniqueId),
		State:       NetworkconnectivityAlphaSpokeStateEnumToProto(resource.State),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Location:    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.LinkedVpnTunnels {
		p.LinkedVpnTunnels = append(p.LinkedVpnTunnels, r)
	}
	for _, r := range resource.LinkedInterconnectAttachments {
		p.LinkedInterconnectAttachments = append(p.LinkedInterconnectAttachments, r)
	}
	for _, r := range resource.LinkedRouterApplianceInstances {
		p.LinkedRouterApplianceInstances = append(p.LinkedRouterApplianceInstances, NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesToProto(&r))
	}

	return p
}

// ApplySpoke handles the gRPC request by passing it to the underlying Spoke Apply() method.
func (s *SpokeServer) applySpoke(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkconnectivityAlphaSpokeRequest) (*alphapb.NetworkconnectivityAlphaSpoke, error) {
	p := ProtoToSpoke(request.GetResource())
	res, err := c.ApplySpoke(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SpokeToProto(res)
	return r, nil
}

// ApplySpoke handles the gRPC request by passing it to the underlying Spoke Apply() method.
func (s *SpokeServer) ApplyNetworkconnectivityAlphaSpoke(ctx context.Context, request *alphapb.ApplyNetworkconnectivityAlphaSpokeRequest) (*alphapb.NetworkconnectivityAlphaSpoke, error) {
	cl, err := createConfigSpoke(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySpoke(ctx, cl, request)
}

// DeleteSpoke handles the gRPC request by passing it to the underlying Spoke Delete() method.
func (s *SpokeServer) DeleteNetworkconnectivityAlphaSpoke(ctx context.Context, request *alphapb.DeleteNetworkconnectivityAlphaSpokeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSpoke(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSpoke(ctx, ProtoToSpoke(request.GetResource()))

}

// ListNetworkconnectivityAlphaSpoke handles the gRPC request by passing it to the underlying SpokeList() method.
func (s *SpokeServer) ListNetworkconnectivityAlphaSpoke(ctx context.Context, request *alphapb.ListNetworkconnectivityAlphaSpokeRequest) (*alphapb.ListNetworkconnectivityAlphaSpokeResponse, error) {
	cl, err := createConfigSpoke(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSpoke(ctx, ProtoToSpoke(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkconnectivityAlphaSpoke
	for _, r := range resources.Items {
		rp := SpokeToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListNetworkconnectivityAlphaSpokeResponse{Items: protos}, nil
}

func createConfigSpoke(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
