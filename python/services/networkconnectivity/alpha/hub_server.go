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

// Server implements the gRPC interface for Hub.
type HubServer struct{}

// ProtoToHubStateEnum converts a HubStateEnum enum from its proto representation.
func ProtoToNetworkconnectivityAlphaHubStateEnum(e alphapb.NetworkconnectivityAlphaHubStateEnum) *alpha.HubStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkconnectivityAlphaHubStateEnum_name[int32(e)]; ok {
		e := alpha.HubStateEnum(n[len("NetworkconnectivityAlphaHubStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToHub converts a Hub resource from its proto representation.
func ProtoToHub(p *alphapb.NetworkconnectivityAlphaHub) *alpha.Hub {
	obj := &alpha.Hub{
		Name:        dcl.StringOrNil(p.Name),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.Description),
		UniqueId:    dcl.StringOrNil(p.UniqueId),
		State:       ProtoToNetworkconnectivityAlphaHubStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.Project),
	}
	return obj
}

// HubStateEnumToProto converts a HubStateEnum enum to its proto representation.
func NetworkconnectivityAlphaHubStateEnumToProto(e *alpha.HubStateEnum) alphapb.NetworkconnectivityAlphaHubStateEnum {
	if e == nil {
		return alphapb.NetworkconnectivityAlphaHubStateEnum(0)
	}
	if v, ok := alphapb.NetworkconnectivityAlphaHubStateEnum_value["HubStateEnum"+string(*e)]; ok {
		return alphapb.NetworkconnectivityAlphaHubStateEnum(v)
	}
	return alphapb.NetworkconnectivityAlphaHubStateEnum(0)
}

// HubToProto converts a Hub resource to its proto representation.
func HubToProto(resource *alpha.Hub) *alphapb.NetworkconnectivityAlphaHub {
	p := &alphapb.NetworkconnectivityAlphaHub{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Description: dcl.ValueOrEmptyString(resource.Description),
		UniqueId:    dcl.ValueOrEmptyString(resource.UniqueId),
		State:       NetworkconnectivityAlphaHubStateEnumToProto(resource.State),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyHub handles the gRPC request by passing it to the underlying Hub Apply() method.
func (s *HubServer) applyHub(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkconnectivityAlphaHubRequest) (*alphapb.NetworkconnectivityAlphaHub, error) {
	p := ProtoToHub(request.GetResource())
	res, err := c.ApplyHub(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HubToProto(res)
	return r, nil
}

// ApplyHub handles the gRPC request by passing it to the underlying Hub Apply() method.
func (s *HubServer) ApplyNetworkconnectivityAlphaHub(ctx context.Context, request *alphapb.ApplyNetworkconnectivityAlphaHubRequest) (*alphapb.NetworkconnectivityAlphaHub, error) {
	cl, err := createConfigHub(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyHub(ctx, cl, request)
}

// DeleteHub handles the gRPC request by passing it to the underlying Hub Delete() method.
func (s *HubServer) DeleteNetworkconnectivityAlphaHub(ctx context.Context, request *alphapb.DeleteNetworkconnectivityAlphaHubRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHub(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHub(ctx, ProtoToHub(request.GetResource()))

}

// ListNetworkconnectivityAlphaHub handles the gRPC request by passing it to the underlying HubList() method.
func (s *HubServer) ListNetworkconnectivityAlphaHub(ctx context.Context, request *alphapb.ListNetworkconnectivityAlphaHubRequest) (*alphapb.ListNetworkconnectivityAlphaHubResponse, error) {
	cl, err := createConfigHub(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHub(ctx, ProtoToHub(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkconnectivityAlphaHub
	for _, r := range resources.Items {
		rp := HubToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListNetworkconnectivityAlphaHubResponse{Items: protos}, nil
}

func createConfigHub(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
