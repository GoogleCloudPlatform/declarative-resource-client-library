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

// Server implements the gRPC interface for Route.
type RouteServer struct{}

// ProtoToRouteWarningCodeEnum converts a RouteWarningCodeEnum enum from its proto representation.
func ProtoToComputeRouteWarningCodeEnum(e computepb.ComputeRouteWarningCodeEnum) *compute.RouteWarningCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouteWarningCodeEnum_name[int32(e)]; ok {
		e := compute.RouteWarningCodeEnum(n[len("ComputeRouteWarningCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouteWarning converts a RouteWarning resource from its proto representation.
func ProtoToComputeRouteWarning(p *computepb.ComputeRouteWarning) *compute.RouteWarning {
	if p == nil {
		return nil
	}
	obj := &compute.RouteWarning{
		Code:    ProtoToComputeRouteWarningCodeEnum(p.GetCode()),
		Message: dcl.StringOrNil(p.Message),
	}
	return obj
}

// ProtoToRoute converts a Route resource from its proto representation.
func ProtoToRoute(p *computepb.ComputeRoute) *compute.Route {
	obj := &compute.Route{
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
		obj.Warning = append(obj.Warning, *ProtoToComputeRouteWarning(r))
	}
	return obj
}

// RouteWarningCodeEnumToProto converts a RouteWarningCodeEnum enum to its proto representation.
func ComputeRouteWarningCodeEnumToProto(e *compute.RouteWarningCodeEnum) computepb.ComputeRouteWarningCodeEnum {
	if e == nil {
		return computepb.ComputeRouteWarningCodeEnum(0)
	}
	if v, ok := computepb.ComputeRouteWarningCodeEnum_value["RouteWarningCodeEnum"+string(*e)]; ok {
		return computepb.ComputeRouteWarningCodeEnum(v)
	}
	return computepb.ComputeRouteWarningCodeEnum(0)
}

// RouteWarningToProto converts a RouteWarning resource to its proto representation.
func ComputeRouteWarningToProto(o *compute.RouteWarning) *computepb.ComputeRouteWarning {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouteWarning{
		Code:    ComputeRouteWarningCodeEnumToProto(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	p.Data = make(map[string]string)
	for k, r := range o.Data {
		p.Data[k] = r
	}
	return p
}

// RouteToProto converts a Route resource to its proto representation.
func RouteToProto(resource *compute.Route) *computepb.ComputeRoute {
	p := &computepb.ComputeRoute{
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
		p.Warning = append(p.Warning, ComputeRouteWarningToProto(&r))
	}

	return p
}

// ApplyRoute handles the gRPC request by passing it to the underlying Route Apply() method.
func (s *RouteServer) applyRoute(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeRouteRequest) (*computepb.ComputeRoute, error) {
	p := ProtoToRoute(request.GetResource())
	res, err := c.ApplyRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RouteToProto(res)
	return r, nil
}

// ApplyRoute handles the gRPC request by passing it to the underlying Route Apply() method.
func (s *RouteServer) ApplyComputeRoute(ctx context.Context, request *computepb.ApplyComputeRouteRequest) (*computepb.ComputeRoute, error) {
	cl, err := createConfigRoute(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRoute(ctx, cl, request)
}

// DeleteRoute handles the gRPC request by passing it to the underlying Route Delete() method.
func (s *RouteServer) DeleteComputeRoute(ctx context.Context, request *computepb.DeleteComputeRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRoute(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRoute(ctx, ProtoToRoute(request.GetResource()))

}

// ListRoute handles the gRPC request by passing it to the underlying RouteList() method.
func (s *RouteServer) ListComputeRoute(ctx context.Context, request *computepb.ListComputeRouteRequest) (*computepb.ListComputeRouteResponse, error) {
	cl, err := createConfigRoute(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRoute(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeRoute
	for _, r := range resources.Items {
		rp := RouteToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeRouteResponse{Items: protos}, nil
}

func createConfigRoute(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
