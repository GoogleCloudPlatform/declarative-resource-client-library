// Copyright 2021 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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

// Server implements the gRPC interface for TargetHttpProxy.
type TargetHttpProxyServer struct{}

// ProtoToTargetHttpProxy converts a TargetHttpProxy resource from its proto representation.
func ProtoToTargetHttpProxy(p *betapb.ComputeBetaTargetHttpProxy) *beta.TargetHttpProxy {
	obj := &beta.TargetHttpProxy{
		Id:          dcl.Int64OrNil(p.Id),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		UrlMap:      dcl.StringOrNil(p.UrlMap),
		Project:     dcl.StringOrNil(p.Project),
	}
	return obj
}

// TargetHttpProxyToProto converts a TargetHttpProxy resource to its proto representation.
func TargetHttpProxyToProto(resource *beta.TargetHttpProxy) *betapb.ComputeBetaTargetHttpProxy {
	p := &betapb.ComputeBetaTargetHttpProxy{
		Id:          dcl.ValueOrEmptyInt64(resource.Id),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		UrlMap:      dcl.ValueOrEmptyString(resource.UrlMap),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyTargetHttpProxy handles the gRPC request by passing it to the underlying TargetHttpProxy Apply() method.
func (s *TargetHttpProxyServer) applyTargetHttpProxy(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaTargetHttpProxyRequest) (*betapb.ComputeBetaTargetHttpProxy, error) {
	p := ProtoToTargetHttpProxy(request.GetResource())
	res, err := c.ApplyTargetHttpProxy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetHttpProxyToProto(res)
	return r, nil
}

// ApplyTargetHttpProxy handles the gRPC request by passing it to the underlying TargetHttpProxy Apply() method.
func (s *TargetHttpProxyServer) ApplyComputeBetaTargetHttpProxy(ctx context.Context, request *betapb.ApplyComputeBetaTargetHttpProxyRequest) (*betapb.ComputeBetaTargetHttpProxy, error) {
	cl, err := createConfigTargetHttpProxy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTargetHttpProxy(ctx, cl, request)
}

// DeleteTargetHttpProxy handles the gRPC request by passing it to the underlying TargetHttpProxy Delete() method.
func (s *TargetHttpProxyServer) DeleteComputeBetaTargetHttpProxy(ctx context.Context, request *betapb.DeleteComputeBetaTargetHttpProxyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTargetHttpProxy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTargetHttpProxy(ctx, ProtoToTargetHttpProxy(request.GetResource()))

}

// ListComputeBetaTargetHttpProxy handles the gRPC request by passing it to the underlying TargetHttpProxyList() method.
func (s *TargetHttpProxyServer) ListComputeBetaTargetHttpProxy(ctx context.Context, request *betapb.ListComputeBetaTargetHttpProxyRequest) (*betapb.ListComputeBetaTargetHttpProxyResponse, error) {
	cl, err := createConfigTargetHttpProxy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTargetHttpProxy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaTargetHttpProxy
	for _, r := range resources.Items {
		rp := TargetHttpProxyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaTargetHttpProxyResponse{Items: protos}, nil
}

func createConfigTargetHttpProxy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
