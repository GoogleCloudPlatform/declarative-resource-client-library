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
	iampb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/iam_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
)

// Server implements the gRPC interface for ServiceAccount.
type ServiceAccountServer struct{}

// ProtoToServiceAccountActasResources converts a ServiceAccountActasResources resource from its proto representation.
func ProtoToIamServiceAccountActasResources(p *iampb.IamServiceAccountActasResources) *iam.ServiceAccountActasResources {
	if p == nil {
		return nil
	}
	obj := &iam.ServiceAccountActasResources{}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToIamServiceAccountActasResourcesResources(r))
	}
	return obj
}

// ProtoToServiceAccountActasResourcesResources converts a ServiceAccountActasResourcesResources resource from its proto representation.
func ProtoToIamServiceAccountActasResourcesResources(p *iampb.IamServiceAccountActasResourcesResources) *iam.ServiceAccountActasResourcesResources {
	if p == nil {
		return nil
	}
	obj := &iam.ServiceAccountActasResourcesResources{
		FullResourceName: dcl.StringOrNil(p.FullResourceName),
	}
	return obj
}

// ProtoToServiceAccount converts a ServiceAccount resource from its proto representation.
func ProtoToServiceAccount(p *iampb.IamServiceAccount) *iam.ServiceAccount {
	obj := &iam.ServiceAccount{
		Name:           dcl.StringOrNil(p.Name),
		Project:        dcl.StringOrNil(p.Project),
		UniqueId:       dcl.StringOrNil(p.UniqueId),
		Email:          dcl.StringOrNil(p.Email),
		DisplayName:    dcl.StringOrNil(p.DisplayName),
		Description:    dcl.StringOrNil(p.Description),
		OAuth2ClientId: dcl.StringOrNil(p.Oauth2ClientId),
		ActasResources: ProtoToIamServiceAccountActasResources(p.GetActasResources()),
		Disabled:       dcl.Bool(p.Disabled),
	}
	return obj
}

// ServiceAccountActasResourcesToProto converts a ServiceAccountActasResources resource to its proto representation.
func IamServiceAccountActasResourcesToProto(o *iam.ServiceAccountActasResources) *iampb.IamServiceAccountActasResources {
	if o == nil {
		return nil
	}
	p := &iampb.IamServiceAccountActasResources{}
	for _, r := range o.Resources {
		p.Resources = append(p.Resources, IamServiceAccountActasResourcesResourcesToProto(&r))
	}
	return p
}

// ServiceAccountActasResourcesResourcesToProto converts a ServiceAccountActasResourcesResources resource to its proto representation.
func IamServiceAccountActasResourcesResourcesToProto(o *iam.ServiceAccountActasResourcesResources) *iampb.IamServiceAccountActasResourcesResources {
	if o == nil {
		return nil
	}
	p := &iampb.IamServiceAccountActasResourcesResources{
		FullResourceName: dcl.ValueOrEmptyString(o.FullResourceName),
	}
	return p
}

// ServiceAccountToProto converts a ServiceAccount resource to its proto representation.
func ServiceAccountToProto(resource *iam.ServiceAccount) *iampb.IamServiceAccount {
	p := &iampb.IamServiceAccount{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		Project:        dcl.ValueOrEmptyString(resource.Project),
		UniqueId:       dcl.ValueOrEmptyString(resource.UniqueId),
		Email:          dcl.ValueOrEmptyString(resource.Email),
		DisplayName:    dcl.ValueOrEmptyString(resource.DisplayName),
		Description:    dcl.ValueOrEmptyString(resource.Description),
		Oauth2ClientId: dcl.ValueOrEmptyString(resource.OAuth2ClientId),
		ActasResources: IamServiceAccountActasResourcesToProto(resource.ActasResources),
		Disabled:       dcl.ValueOrEmptyBool(resource.Disabled),
	}

	return p
}

// ApplyServiceAccount handles the gRPC request by passing it to the underlying ServiceAccount Apply() method.
func (s *ServiceAccountServer) applyServiceAccount(ctx context.Context, c *iam.Client, request *iampb.ApplyIamServiceAccountRequest) (*iampb.IamServiceAccount, error) {
	p := ProtoToServiceAccount(request.GetResource())
	res, err := c.ApplyServiceAccount(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceAccountToProto(res)
	return r, nil
}

// ApplyServiceAccount handles the gRPC request by passing it to the underlying ServiceAccount Apply() method.
func (s *ServiceAccountServer) ApplyIamServiceAccount(ctx context.Context, request *iampb.ApplyIamServiceAccountRequest) (*iampb.IamServiceAccount, error) {
	cl, err := createConfigServiceAccount(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServiceAccount(ctx, cl, request)
}

// DeleteServiceAccount handles the gRPC request by passing it to the underlying ServiceAccount Delete() method.
func (s *ServiceAccountServer) DeleteIamServiceAccount(ctx context.Context, request *iampb.DeleteIamServiceAccountRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceAccount(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceAccount(ctx, ProtoToServiceAccount(request.GetResource()))

}

// ListIamServiceAccount handles the gRPC request by passing it to the underlying ServiceAccountList() method.
func (s *ServiceAccountServer) ListIamServiceAccount(ctx context.Context, request *iampb.ListIamServiceAccountRequest) (*iampb.ListIamServiceAccountResponse, error) {
	cl, err := createConfigServiceAccount(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceAccount(ctx, ProtoToServiceAccount(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*iampb.IamServiceAccount
	for _, r := range resources.Items {
		rp := ServiceAccountToProto(r)
		protos = append(protos, rp)
	}
	return &iampb.ListIamServiceAccountResponse{Items: protos}, nil
}

func createConfigServiceAccount(ctx context.Context, service_account_file string) (*iam.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iam.NewClient(conf), nil
}
