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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/alpha/containerazure_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
)

// Server implements the gRPC interface for AzureClient.
type AzureClientServer struct{}

// ProtoToAzureClient converts a AzureClient resource from its proto representation.
func ProtoToAzureClient(p *alphapb.ContainerazureAlphaAzureClient) *alpha.AzureClient {
	obj := &alpha.AzureClient{
		Name:          dcl.StringOrNil(p.GetName()),
		TenantId:      dcl.StringOrNil(p.GetTenantId()),
		ApplicationId: dcl.StringOrNil(p.GetApplicationId()),
		Certificate:   dcl.StringOrNil(p.GetCertificate()),
		Uid:           dcl.StringOrNil(p.GetUid()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// AzureClientToProto converts a AzureClient resource to its proto representation.
func AzureClientToProto(resource *alpha.AzureClient) *alphapb.ContainerazureAlphaAzureClient {
	p := &alphapb.ContainerazureAlphaAzureClient{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTenantId(dcl.ValueOrEmptyString(resource.TenantId))
	p.SetApplicationId(dcl.ValueOrEmptyString(resource.ApplicationId))
	p.SetCertificate(dcl.ValueOrEmptyString(resource.Certificate))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyAzureClient handles the gRPC request by passing it to the underlying AzureClient Apply() method.
func (s *AzureClientServer) applyAzureClient(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaAzureClientRequest) (*alphapb.ContainerazureAlphaAzureClient, error) {
	p := ProtoToAzureClient(request.GetResource())
	res, err := c.ApplyAzureClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureClientToProto(res)
	return r, nil
}

// applyContainerazureAlphaAzureClient handles the gRPC request by passing it to the underlying AzureClient Apply() method.
func (s *AzureClientServer) ApplyContainerazureAlphaAzureClient(ctx context.Context, request *alphapb.ApplyContainerazureAlphaAzureClientRequest) (*alphapb.ContainerazureAlphaAzureClient, error) {
	cl, err := createConfigAzureClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAzureClient(ctx, cl, request)
}

// DeleteAzureClient handles the gRPC request by passing it to the underlying AzureClient Delete() method.
func (s *AzureClientServer) DeleteContainerazureAlphaAzureClient(ctx context.Context, request *alphapb.DeleteContainerazureAlphaAzureClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureClient(ctx, ProtoToAzureClient(request.GetResource()))

}

// ListContainerazureAlphaAzureClient handles the gRPC request by passing it to the underlying AzureClientList() method.
func (s *AzureClientServer) ListContainerazureAlphaAzureClient(ctx context.Context, request *alphapb.ListContainerazureAlphaAzureClientRequest) (*alphapb.ListContainerazureAlphaAzureClientResponse, error) {
	cl, err := createConfigAzureClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureClient(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaAzureClient
	for _, r := range resources.Items {
		rp := AzureClientToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContainerazureAlphaAzureClientResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAzureClient(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
