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

// Server implements the gRPC interface for VpnTunnel.
type VpnTunnelServer struct{}

// ProtoToVpnTunnelStatusEnum converts a VpnTunnelStatusEnum enum from its proto representation.
func ProtoToComputeBetaVpnTunnelStatusEnum(e betapb.ComputeBetaVpnTunnelStatusEnum) *beta.VpnTunnelStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaVpnTunnelStatusEnum_name[int32(e)]; ok {
		e := beta.VpnTunnelStatusEnum(n[len("ComputeBetaVpnTunnelStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToVpnTunnel converts a VpnTunnel resource from its proto representation.
func ProtoToVpnTunnel(p *betapb.ComputeBetaVpnTunnel) *beta.VpnTunnel {
	obj := &beta.VpnTunnel{
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
		Status:                       ProtoToComputeBetaVpnTunnelStatusEnum(p.GetStatus()),
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
func ComputeBetaVpnTunnelStatusEnumToProto(e *beta.VpnTunnelStatusEnum) betapb.ComputeBetaVpnTunnelStatusEnum {
	if e == nil {
		return betapb.ComputeBetaVpnTunnelStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaVpnTunnelStatusEnum_value["VpnTunnelStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaVpnTunnelStatusEnum(v)
	}
	return betapb.ComputeBetaVpnTunnelStatusEnum(0)
}

// VpnTunnelToProto converts a VpnTunnel resource to its proto representation.
func VpnTunnelToProto(resource *beta.VpnTunnel) *betapb.ComputeBetaVpnTunnel {
	p := &betapb.ComputeBetaVpnTunnel{
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
		Status:                       ComputeBetaVpnTunnelStatusEnumToProto(resource.Status),
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
func (s *VpnTunnelServer) applyVpnTunnel(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaVpnTunnelRequest) (*betapb.ComputeBetaVpnTunnel, error) {
	p := ProtoToVpnTunnel(request.GetResource())
	res, err := c.ApplyVpnTunnel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VpnTunnelToProto(res)
	return r, nil
}

// ApplyVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Apply() method.
func (s *VpnTunnelServer) ApplyComputeBetaVpnTunnel(ctx context.Context, request *betapb.ApplyComputeBetaVpnTunnelRequest) (*betapb.ComputeBetaVpnTunnel, error) {
	cl, err := createConfigVpnTunnel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyVpnTunnel(ctx, cl, request)
}

// DeleteVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Delete() method.
func (s *VpnTunnelServer) DeleteComputeBetaVpnTunnel(ctx context.Context, request *betapb.DeleteComputeBetaVpnTunnelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVpnTunnel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVpnTunnel(ctx, ProtoToVpnTunnel(request.GetResource()))

}

// ListComputeBetaVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnelList() method.
func (s *VpnTunnelServer) ListComputeBetaVpnTunnel(ctx context.Context, request *betapb.ListComputeBetaVpnTunnelRequest) (*betapb.ListComputeBetaVpnTunnelResponse, error) {
	cl, err := createConfigVpnTunnel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVpnTunnel(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaVpnTunnel
	for _, r := range resources.Items {
		rp := VpnTunnelToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaVpnTunnelResponse{Items: protos}, nil
}

func createConfigVpnTunnel(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
