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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/beta/iap_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta"
)

// Server implements the gRPC interface for IdentityAwareProxyClient.
type IdentityAwareProxyClientServer struct{}

// ProtoToIdentityAwareProxyClient converts a IdentityAwareProxyClient resource from its proto representation.
func ProtoToIdentityAwareProxyClient(p *betapb.IapBetaIdentityAwareProxyClient) *beta.IdentityAwareProxyClient {
	obj := &beta.IdentityAwareProxyClient{
		Name:        dcl.StringOrNil(p.Name),
		Secret:      dcl.StringOrNil(p.Secret),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Project:     dcl.StringOrNil(p.Project),
		Brand:       dcl.StringOrNil(p.Brand),
	}
	return obj
}

// IdentityAwareProxyClientToProto converts a IdentityAwareProxyClient resource to its proto representation.
func IdentityAwareProxyClientToProto(resource *beta.IdentityAwareProxyClient) *betapb.IapBetaIdentityAwareProxyClient {
	p := &betapb.IapBetaIdentityAwareProxyClient{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Secret:      dcl.ValueOrEmptyString(resource.Secret),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Brand:       dcl.ValueOrEmptyString(resource.Brand),
	}

	return p
}

// ApplyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) applyIdentityAwareProxyClient(ctx context.Context, c *beta.Client, request *betapb.ApplyIapBetaIdentityAwareProxyClientRequest) (*betapb.IapBetaIdentityAwareProxyClient, error) {
	p := ProtoToIdentityAwareProxyClient(request.GetResource())
	res, err := c.ApplyIdentityAwareProxyClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := IdentityAwareProxyClientToProto(res)
	return r, nil
}

// ApplyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) ApplyIapBetaIdentityAwareProxyClient(ctx context.Context, request *betapb.ApplyIapBetaIdentityAwareProxyClientRequest) (*betapb.IapBetaIdentityAwareProxyClient, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyIdentityAwareProxyClient(ctx, cl, request)
}

// DeleteIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Delete() method.
func (s *IdentityAwareProxyClientServer) DeleteIapBetaIdentityAwareProxyClient(ctx context.Context, request *betapb.DeleteIapBetaIdentityAwareProxyClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteIdentityAwareProxyClient(ctx, ProtoToIdentityAwareProxyClient(request.GetResource()))

}

// ListIapBetaIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClientList() method.
func (s *IdentityAwareProxyClientServer) ListIapBetaIdentityAwareProxyClient(ctx context.Context, request *betapb.ListIapBetaIdentityAwareProxyClientRequest) (*betapb.ListIapBetaIdentityAwareProxyClientResponse, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListIdentityAwareProxyClient(ctx, request.Project, request.Brand)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IapBetaIdentityAwareProxyClient
	for _, r := range resources.Items {
		rp := IdentityAwareProxyClientToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListIapBetaIdentityAwareProxyClientResponse{Items: protos}, nil
}

func createConfigIdentityAwareProxyClient(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
