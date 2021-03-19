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
	iappb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/iap_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap"
)

// Server implements the gRPC interface for IdentityAwareProxyClient.
type IdentityAwareProxyClientServer struct{}

// ProtoToIdentityAwareProxyClient converts a IdentityAwareProxyClient resource from its proto representation.
func ProtoToIdentityAwareProxyClient(p *iappb.IapIdentityAwareProxyClient) *iap.IdentityAwareProxyClient {
	obj := &iap.IdentityAwareProxyClient{
		Name:        dcl.StringOrNil(p.Name),
		Secret:      dcl.StringOrNil(p.Secret),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Project:     dcl.StringOrNil(p.Project),
		Brand:       dcl.StringOrNil(p.Brand),
	}
	return obj
}

// IdentityAwareProxyClientToProto converts a IdentityAwareProxyClient resource to its proto representation.
func IdentityAwareProxyClientToProto(resource *iap.IdentityAwareProxyClient) *iappb.IapIdentityAwareProxyClient {
	p := &iappb.IapIdentityAwareProxyClient{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Secret:      dcl.ValueOrEmptyString(resource.Secret),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Brand:       dcl.ValueOrEmptyString(resource.Brand),
	}

	return p
}

// ApplyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) applyIdentityAwareProxyClient(ctx context.Context, c *iap.Client, request *iappb.ApplyIapIdentityAwareProxyClientRequest) (*iappb.IapIdentityAwareProxyClient, error) {
	p := ProtoToIdentityAwareProxyClient(request.GetResource())
	res, err := c.ApplyIdentityAwareProxyClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := IdentityAwareProxyClientToProto(res)
	return r, nil
}

// ApplyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) ApplyIapIdentityAwareProxyClient(ctx context.Context, request *iappb.ApplyIapIdentityAwareProxyClientRequest) (*iappb.IapIdentityAwareProxyClient, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyIdentityAwareProxyClient(ctx, cl, request)
}

// DeleteIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Delete() method.
func (s *IdentityAwareProxyClientServer) DeleteIapIdentityAwareProxyClient(ctx context.Context, request *iappb.DeleteIapIdentityAwareProxyClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteIdentityAwareProxyClient(ctx, ProtoToIdentityAwareProxyClient(request.GetResource()))

}

// ListIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClientList() method.
func (s *IdentityAwareProxyClientServer) ListIapIdentityAwareProxyClient(ctx context.Context, request *iappb.ListIapIdentityAwareProxyClientRequest) (*iappb.ListIapIdentityAwareProxyClientResponse, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListIdentityAwareProxyClient(ctx, request.Project, request.Brand)
	if err != nil {
		return nil, err
	}
	var protos []*iappb.IapIdentityAwareProxyClient
	for _, r := range resources.Items {
		rp := IdentityAwareProxyClientToProto(r)
		protos = append(protos, rp)
	}
	return &iappb.ListIapIdentityAwareProxyClientResponse{Items: protos}, nil
}

func createConfigIdentityAwareProxyClient(ctx context.Context, service_account_file string) (*iap.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iap.NewClient(conf), nil
}
