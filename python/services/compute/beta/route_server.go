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

// Server implements the gRPC interface for Route.
type RouteServer struct{}

// ProtoToRouteWarningCodeEnum converts a RouteWarningCodeEnum enum from its proto representation.
func ProtoToComputeBetaRouteWarningCodeEnum(e betapb.ComputeBetaRouteWarningCodeEnum) *beta.RouteWarningCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouteWarningCodeEnum_name[int32(e)]; ok {
		e := beta.RouteWarningCodeEnum(n[len("ComputeBetaRouteWarningCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouteWarning converts a RouteWarning resource from its proto representation.
func ProtoToComputeBetaRouteWarning(p *betapb.ComputeBetaRouteWarning) *beta.RouteWarning {
	if p == nil {
		return nil
	}
	obj := &beta.RouteWarning{
		Code:    ProtoToComputeBetaRouteWarningCodeEnum(p.GetCode()),
		Message: dcl.StringOrNil(p.Message),
	}
	return obj
}

// ProtoToRoute converts a Route resource from its proto representation.
func ProtoToRoute(p *betapb.ComputeBetaRoute) *beta.Route {
	obj := &beta.Route{
		Id:               dcl.Int64OrNil(p.Id),
		Name:             dcl.StringOrNil(p.Name),
		Description:      dcl.StringOrNil(p.Description),
		Network:          dcl.StringOrNil(p.Network),
		DestRange:        dcl.StringOrNil(p.DestRange),
		Priority:         dcl.Int64OrNil(p.Priority),
		NextHopInstance:  dcl.StringOrNil(p.NextHopInstance),
		NextHopIP:        dcl.StringOrNil(p.NextHopIp),
		NextHopNetwork:   dcl.StringOrNil(p.NextHopNetwork),
		NextHopGateway:   dcl.StringOrNil(p.NextHopGateway),
		NextHopPeering:   dcl.StringOrNil(p.NextHopPeering),
		NextHopIlb:       dcl.StringOrNil(p.NextHopIlb),
		NextHopVpnTunnel: dcl.StringOrNil(p.NextHopVpnTunnel),
		SelfLink:         dcl.StringOrNil(p.SelfLink),
		Project:          dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetTag() {
		obj.Tag = append(obj.Tag, r)
	}
	for _, r := range p.GetWarning() {
		obj.Warning = append(obj.Warning, *ProtoToComputeBetaRouteWarning(r))
	}
	return obj
}

// RouteWarningCodeEnumToProto converts a RouteWarningCodeEnum enum to its proto representation.
func ComputeBetaRouteWarningCodeEnumToProto(e *beta.RouteWarningCodeEnum) betapb.ComputeBetaRouteWarningCodeEnum {
	if e == nil {
		return betapb.ComputeBetaRouteWarningCodeEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouteWarningCodeEnum_value["RouteWarningCodeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouteWarningCodeEnum(v)
	}
	return betapb.ComputeBetaRouteWarningCodeEnum(0)
}

// RouteWarningToProto converts a RouteWarning resource to its proto representation.
func ComputeBetaRouteWarningToProto(o *beta.RouteWarning) *betapb.ComputeBetaRouteWarning {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouteWarning{
		Code:    ComputeBetaRouteWarningCodeEnumToProto(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	p.Data = make(map[string]string)
	for k, r := range o.Data {
		p.Data[k] = r
	}
	return p
}

// RouteToProto converts a Route resource to its proto representation.
func RouteToProto(resource *beta.Route) *betapb.ComputeBetaRoute {
	p := &betapb.ComputeBetaRoute{
		Id:               dcl.ValueOrEmptyInt64(resource.Id),
		Name:             dcl.ValueOrEmptyString(resource.Name),
		Description:      dcl.ValueOrEmptyString(resource.Description),
		Network:          dcl.ValueOrEmptyString(resource.Network),
		DestRange:        dcl.ValueOrEmptyString(resource.DestRange),
		Priority:         dcl.ValueOrEmptyInt64(resource.Priority),
		NextHopInstance:  dcl.ValueOrEmptyString(resource.NextHopInstance),
		NextHopIp:        dcl.ValueOrEmptyString(resource.NextHopIP),
		NextHopNetwork:   dcl.ValueOrEmptyString(resource.NextHopNetwork),
		NextHopGateway:   dcl.ValueOrEmptyString(resource.NextHopGateway),
		NextHopPeering:   dcl.ValueOrEmptyString(resource.NextHopPeering),
		NextHopIlb:       dcl.ValueOrEmptyString(resource.NextHopIlb),
		NextHopVpnTunnel: dcl.ValueOrEmptyString(resource.NextHopVpnTunnel),
		SelfLink:         dcl.ValueOrEmptyString(resource.SelfLink),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Tag {
		p.Tag = append(p.Tag, r)
	}
	for _, r := range resource.Warning {
		p.Warning = append(p.Warning, ComputeBetaRouteWarningToProto(&r))
	}

	return p
}

// ApplyRoute handles the gRPC request by passing it to the underlying Route Apply() method.
func (s *RouteServer) applyRoute(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaRouteRequest) (*betapb.ComputeBetaRoute, error) {
	p := ProtoToRoute(request.GetResource())
	res, err := c.ApplyRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RouteToProto(res)
	return r, nil
}

// ApplyRoute handles the gRPC request by passing it to the underlying Route Apply() method.
func (s *RouteServer) ApplyComputeBetaRoute(ctx context.Context, request *betapb.ApplyComputeBetaRouteRequest) (*betapb.ComputeBetaRoute, error) {
	cl, err := createConfigRoute(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRoute(ctx, cl, request)
}

// DeleteRoute handles the gRPC request by passing it to the underlying Route Delete() method.
func (s *RouteServer) DeleteComputeBetaRoute(ctx context.Context, request *betapb.DeleteComputeBetaRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRoute(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRoute(ctx, ProtoToRoute(request.GetResource()))

}

// ListComputeBetaRoute handles the gRPC request by passing it to the underlying RouteList() method.
func (s *RouteServer) ListComputeBetaRoute(ctx context.Context, request *betapb.ListComputeBetaRouteRequest) (*betapb.ListComputeBetaRouteResponse, error) {
	cl, err := createConfigRoute(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRoute(ctx, ProtoToRoute(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaRoute
	for _, r := range resources.Items {
		rp := RouteToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaRouteResponse{Items: protos}, nil
}

func createConfigRoute(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
