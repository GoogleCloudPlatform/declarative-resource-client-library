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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/alpha/iap_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/alpha"
)

// Server implements the gRPC interface for IdentityAwareProxyClient.
type IdentityAwareProxyClientServer struct{}

// ProtoToIdentityAwareProxyClient converts a IdentityAwareProxyClient resource from its proto representation.
func ProtoToIdentityAwareProxyClient(p *alphapb.IapAlphaIdentityAwareProxyClient) *alpha.IdentityAwareProxyClient {
	obj := &alpha.IdentityAwareProxyClient{
		Name:        dcl.StringOrNil(p.Name),
		Secret:      dcl.StringOrNil(p.Secret),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Project:     dcl.StringOrNil(p.Project),
		Brand:       dcl.StringOrNil(p.Brand),
	}
	return obj
}

// IdentityAwareProxyClientToProto converts a IdentityAwareProxyClient resource to its proto representation.
func IdentityAwareProxyClientToProto(resource *alpha.IdentityAwareProxyClient) *alphapb.IapAlphaIdentityAwareProxyClient {
	p := &alphapb.IapAlphaIdentityAwareProxyClient{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Secret:      dcl.ValueOrEmptyString(resource.Secret),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Brand:       dcl.ValueOrEmptyString(resource.Brand),
	}

	return p
}

// ApplyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) applyIdentityAwareProxyClient(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIapAlphaIdentityAwareProxyClientRequest) (*alphapb.IapAlphaIdentityAwareProxyClient, error) {
	p := ProtoToIdentityAwareProxyClient(request.GetResource())
	res, err := c.ApplyIdentityAwareProxyClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := IdentityAwareProxyClientToProto(res)
	return r, nil
}

// ApplyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) ApplyIapAlphaIdentityAwareProxyClient(ctx context.Context, request *alphapb.ApplyIapAlphaIdentityAwareProxyClientRequest) (*alphapb.IapAlphaIdentityAwareProxyClient, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyIdentityAwareProxyClient(ctx, cl, request)
}

// DeleteIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Delete() method.
func (s *IdentityAwareProxyClientServer) DeleteIapAlphaIdentityAwareProxyClient(ctx context.Context, request *alphapb.DeleteIapAlphaIdentityAwareProxyClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteIdentityAwareProxyClient(ctx, ProtoToIdentityAwareProxyClient(request.GetResource()))

}

// ListIapAlphaIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClientList() method.
func (s *IdentityAwareProxyClientServer) ListIapAlphaIdentityAwareProxyClient(ctx context.Context, request *alphapb.ListIapAlphaIdentityAwareProxyClientRequest) (*alphapb.ListIapAlphaIdentityAwareProxyClientResponse, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListIdentityAwareProxyClient(ctx, ProtoToIdentityAwareProxyClient(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IapAlphaIdentityAwareProxyClient
	for _, r := range resources.Items {
		rp := IdentityAwareProxyClientToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListIapAlphaIdentityAwareProxyClientResponse{Items: protos}, nil
}

func createConfigIdentityAwareProxyClient(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
