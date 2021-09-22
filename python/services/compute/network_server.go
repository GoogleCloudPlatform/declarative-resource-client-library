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

// Server implements the gRPC interface for Network.
type NetworkServer struct{}

// ProtoToNetworkRoutingConfigRoutingModeEnum converts a NetworkRoutingConfigRoutingModeEnum enum from its proto representation.
func ProtoToComputeNetworkRoutingConfigRoutingModeEnum(e computepb.ComputeNetworkRoutingConfigRoutingModeEnum) *compute.NetworkRoutingConfigRoutingModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeNetworkRoutingConfigRoutingModeEnum_name[int32(e)]; ok {
		e := compute.NetworkRoutingConfigRoutingModeEnum(n[len("ComputeNetworkRoutingConfigRoutingModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkRoutingConfig converts a NetworkRoutingConfig resource from its proto representation.
func ProtoToComputeNetworkRoutingConfig(p *computepb.ComputeNetworkRoutingConfig) *compute.NetworkRoutingConfig {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkRoutingConfig{
		RoutingMode: ProtoToComputeNetworkRoutingConfigRoutingModeEnum(p.GetRoutingMode()),
	}
	return obj
}

// ProtoToNetwork converts a Network resource from its proto representation.
func ProtoToNetwork(p *computepb.ComputeNetwork) *compute.Network {
	obj := &compute.Network{
		Description:           dcl.StringOrNil(p.Description),
		GatewayIPv4:           dcl.StringOrNil(p.GatewayIpv4),
		Name:                  dcl.StringOrNil(p.Name),
		AutoCreateSubnetworks: dcl.Bool(p.AutoCreateSubnetworks),
		RoutingConfig:         ProtoToComputeNetworkRoutingConfig(p.GetRoutingConfig()),
		Mtu:                   dcl.Int64OrNil(p.Mtu),
		Project:               dcl.StringOrNil(p.Project),
		SelfLink:              dcl.StringOrNil(p.SelfLink),
	}
	return obj
}

// NetworkRoutingConfigRoutingModeEnumToProto converts a NetworkRoutingConfigRoutingModeEnum enum to its proto representation.
func ComputeNetworkRoutingConfigRoutingModeEnumToProto(e *compute.NetworkRoutingConfigRoutingModeEnum) computepb.ComputeNetworkRoutingConfigRoutingModeEnum {
	if e == nil {
		return computepb.ComputeNetworkRoutingConfigRoutingModeEnum(0)
	}
	if v, ok := computepb.ComputeNetworkRoutingConfigRoutingModeEnum_value["NetworkRoutingConfigRoutingModeEnum"+string(*e)]; ok {
		return computepb.ComputeNetworkRoutingConfigRoutingModeEnum(v)
	}
	return computepb.ComputeNetworkRoutingConfigRoutingModeEnum(0)
}

// NetworkRoutingConfigToProto converts a NetworkRoutingConfig resource to its proto representation.
func ComputeNetworkRoutingConfigToProto(o *compute.NetworkRoutingConfig) *computepb.ComputeNetworkRoutingConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkRoutingConfig{
		RoutingMode: ComputeNetworkRoutingConfigRoutingModeEnumToProto(o.RoutingMode),
	}
	return p
}

// NetworkToProto converts a Network resource to its proto representation.
func NetworkToProto(resource *compute.Network) *computepb.ComputeNetwork {
	p := &computepb.ComputeNetwork{
		Description:           dcl.ValueOrEmptyString(resource.Description),
		GatewayIpv4:           dcl.ValueOrEmptyString(resource.GatewayIPv4),
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		AutoCreateSubnetworks: dcl.ValueOrEmptyBool(resource.AutoCreateSubnetworks),
		RoutingConfig:         ComputeNetworkRoutingConfigToProto(resource.RoutingConfig),
		Mtu:                   dcl.ValueOrEmptyInt64(resource.Mtu),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		SelfLink:              dcl.ValueOrEmptyString(resource.SelfLink),
	}

	return p
}

// ApplyNetwork handles the gRPC request by passing it to the underlying Network Apply() method.
func (s *NetworkServer) applyNetwork(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeNetworkRequest) (*computepb.ComputeNetwork, error) {
	p := ProtoToNetwork(request.GetResource())
	res, err := c.ApplyNetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkToProto(res)
	return r, nil
}

// ApplyNetwork handles the gRPC request by passing it to the underlying Network Apply() method.
func (s *NetworkServer) ApplyComputeNetwork(ctx context.Context, request *computepb.ApplyComputeNetworkRequest) (*computepb.ComputeNetwork, error) {
	cl, err := createConfigNetwork(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNetwork(ctx, cl, request)
}

// DeleteNetwork handles the gRPC request by passing it to the underlying Network Delete() method.
func (s *NetworkServer) DeleteComputeNetwork(ctx context.Context, request *computepb.DeleteComputeNetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetwork(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetwork(ctx, ProtoToNetwork(request.GetResource()))

}

// ListComputeNetwork handles the gRPC request by passing it to the underlying NetworkList() method.
func (s *NetworkServer) ListComputeNetwork(ctx context.Context, request *computepb.ListComputeNetworkRequest) (*computepb.ListComputeNetworkResponse, error) {
	cl, err := createConfigNetwork(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetwork(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeNetwork
	for _, r := range resources.Items {
		rp := NetworkToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeNetworkResponse{Items: protos}, nil
}

func createConfigNetwork(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
