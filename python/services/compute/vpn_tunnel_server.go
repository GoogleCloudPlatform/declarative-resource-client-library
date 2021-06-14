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

// Server implements the gRPC interface for VpnTunnel.
type VpnTunnelServer struct{}

// ProtoToVpnTunnelStatusEnum converts a VpnTunnelStatusEnum enum from its proto representation.
func ProtoToComputeVpnTunnelStatusEnum(e computepb.ComputeVpnTunnelStatusEnum) *compute.VpnTunnelStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeVpnTunnelStatusEnum_name[int32(e)]; ok {
		e := compute.VpnTunnelStatusEnum(n[len("ComputeVpnTunnelStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToVpnTunnel converts a VpnTunnel resource from its proto representation.
func ProtoToVpnTunnel(p *computepb.ComputeVpnTunnel) *compute.VpnTunnel {
	obj := &compute.VpnTunnel{
		Id:                           dcl.Int64OrNil(p.Id),
		Name:                         dcl.StringOrNil(p.Name),
		Description:                  dcl.StringOrNil(p.Description),
		Region:                       dcl.StringOrNil(p.Region),
		TargetVpnGateway:             dcl.StringOrNil(p.TargetVpnGateway),
		VpnGateway:                   dcl.StringOrNil(p.VpnGateway),
		VpnGatewayInterface:          dcl.Int64OrNil(p.VpnGatewayInterface),
		PeerExternalGateway:          dcl.StringOrNil(p.PeerExternalGateway),
		PeerExternalGatewayInterface: dcl.Int64OrNil(p.PeerExternalGatewayInterface),
		PeerGcpGateway:               dcl.StringOrNil(p.PeerGcpGateway),
		Router:                       dcl.StringOrNil(p.Router),
		PeerIP:                       dcl.StringOrNil(p.PeerIp),
		SharedSecret:                 dcl.StringOrNil(p.SharedSecret),
		SharedSecretHash:             dcl.StringOrNil(p.SharedSecretHash),
		Status:                       ProtoToComputeVpnTunnelStatusEnum(p.GetStatus()),
		SelfLink:                     dcl.StringOrNil(p.SelfLink),
		IkeVersion:                   dcl.Int64OrNil(p.IkeVersion),
		DetailedStatus:               dcl.StringOrNil(p.DetailedStatus),
		Project:                      dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetLocalTrafficSelector() {
		obj.LocalTrafficSelector = append(obj.LocalTrafficSelector, r)
	}
	for _, r := range p.GetRemoteTrafficSelector() {
		obj.RemoteTrafficSelector = append(obj.RemoteTrafficSelector, r)
	}
	return obj
}

// VpnTunnelStatusEnumToProto converts a VpnTunnelStatusEnum enum to its proto representation.
func ComputeVpnTunnelStatusEnumToProto(e *compute.VpnTunnelStatusEnum) computepb.ComputeVpnTunnelStatusEnum {
	if e == nil {
		return computepb.ComputeVpnTunnelStatusEnum(0)
	}
	if v, ok := computepb.ComputeVpnTunnelStatusEnum_value["VpnTunnelStatusEnum"+string(*e)]; ok {
		return computepb.ComputeVpnTunnelStatusEnum(v)
	}
	return computepb.ComputeVpnTunnelStatusEnum(0)
}

// VpnTunnelToProto converts a VpnTunnel resource to its proto representation.
func VpnTunnelToProto(resource *compute.VpnTunnel) *computepb.ComputeVpnTunnel {
	p := &computepb.ComputeVpnTunnel{
		Id:                           dcl.ValueOrEmptyInt64(resource.Id),
		Name:                         dcl.ValueOrEmptyString(resource.Name),
		Description:                  dcl.ValueOrEmptyString(resource.Description),
		Region:                       dcl.ValueOrEmptyString(resource.Region),
		TargetVpnGateway:             dcl.ValueOrEmptyString(resource.TargetVpnGateway),
		VpnGateway:                   dcl.ValueOrEmptyString(resource.VpnGateway),
		VpnGatewayInterface:          dcl.ValueOrEmptyInt64(resource.VpnGatewayInterface),
		PeerExternalGateway:          dcl.ValueOrEmptyString(resource.PeerExternalGateway),
		PeerExternalGatewayInterface: dcl.ValueOrEmptyInt64(resource.PeerExternalGatewayInterface),
		PeerGcpGateway:               dcl.ValueOrEmptyString(resource.PeerGcpGateway),
		Router:                       dcl.ValueOrEmptyString(resource.Router),
		PeerIp:                       dcl.ValueOrEmptyString(resource.PeerIP),
		SharedSecret:                 dcl.ValueOrEmptyString(resource.SharedSecret),
		SharedSecretHash:             dcl.ValueOrEmptyString(resource.SharedSecretHash),
		Status:                       ComputeVpnTunnelStatusEnumToProto(resource.Status),
		SelfLink:                     dcl.ValueOrEmptyString(resource.SelfLink),
		IkeVersion:                   dcl.ValueOrEmptyInt64(resource.IkeVersion),
		DetailedStatus:               dcl.ValueOrEmptyString(resource.DetailedStatus),
		Project:                      dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.LocalTrafficSelector {
		p.LocalTrafficSelector = append(p.LocalTrafficSelector, r)
	}
	for _, r := range resource.RemoteTrafficSelector {
		p.RemoteTrafficSelector = append(p.RemoteTrafficSelector, r)
	}

	return p
}

// ApplyVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Apply() method.
func (s *VpnTunnelServer) applyVpnTunnel(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeVpnTunnelRequest) (*computepb.ComputeVpnTunnel, error) {
	p := ProtoToVpnTunnel(request.GetResource())
	res, err := c.ApplyVpnTunnel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VpnTunnelToProto(res)
	return r, nil
}

// ApplyVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Apply() method.
func (s *VpnTunnelServer) ApplyComputeVpnTunnel(ctx context.Context, request *computepb.ApplyComputeVpnTunnelRequest) (*computepb.ComputeVpnTunnel, error) {
	cl, err := createConfigVpnTunnel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyVpnTunnel(ctx, cl, request)
}

// DeleteVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Delete() method.
func (s *VpnTunnelServer) DeleteComputeVpnTunnel(ctx context.Context, request *computepb.DeleteComputeVpnTunnelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVpnTunnel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVpnTunnel(ctx, ProtoToVpnTunnel(request.GetResource()))

}

// ListComputeVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnelList() method.
func (s *VpnTunnelServer) ListComputeVpnTunnel(ctx context.Context, request *computepb.ListComputeVpnTunnelRequest) (*computepb.ListComputeVpnTunnelResponse, error) {
	cl, err := createConfigVpnTunnel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVpnTunnel(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeVpnTunnel
	for _, r := range resources.Items {
		rp := VpnTunnelToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeVpnTunnelResponse{Items: protos}, nil
}

func createConfigVpnTunnel(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
