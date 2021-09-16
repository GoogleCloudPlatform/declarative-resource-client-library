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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/alpha/iam_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
)

// Server implements the gRPC interface for WorkloadIdentityPoolProvider.
type WorkloadIdentityPoolProviderServer struct{}

// ProtoToWorkloadIdentityPoolProviderStateEnum converts a WorkloadIdentityPoolProviderStateEnum enum from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolProviderStateEnum(e alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum) *alpha.WorkloadIdentityPoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum_name[int32(e)]; ok {
		e := alpha.WorkloadIdentityPoolProviderStateEnum(n[len("IamAlphaWorkloadIdentityPoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPoolProviderAws converts a WorkloadIdentityPoolProviderAws resource from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolProviderAws(p *alphapb.IamAlphaWorkloadIdentityPoolProviderAws) *alpha.WorkloadIdentityPoolProviderAws {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadIdentityPoolProviderAws{
		AccountId: dcl.StringOrNil(p.AccountId),
	}
	for _, r := range p.GetStsUri() {
		obj.StsUri = append(obj.StsUri, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProviderOidc converts a WorkloadIdentityPoolProviderOidc resource from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolProviderOidc(p *alphapb.IamAlphaWorkloadIdentityPoolProviderOidc) *alpha.WorkloadIdentityPoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.StringOrNil(p.IssuerUri),
	}
	for _, r := range p.GetAllowedAudiences() {
		obj.AllowedAudiences = append(obj.AllowedAudiences, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProvider converts a WorkloadIdentityPoolProvider resource from its proto representation.
func ProtoToWorkloadIdentityPoolProvider(p *alphapb.IamAlphaWorkloadIdentityPoolProvider) *alpha.WorkloadIdentityPoolProvider {
	obj := &alpha.WorkloadIdentityPoolProvider{
		Name:                 dcl.StringOrNil(p.Name),
		DisplayName:          dcl.StringOrNil(p.DisplayName),
		Description:          dcl.StringOrNil(p.Description),
		State:                ProtoToIamAlphaWorkloadIdentityPoolProviderStateEnum(p.GetState()),
		Disabled:             dcl.Bool(p.Disabled),
		AttributeCondition:   dcl.StringOrNil(p.AttributeCondition),
		Aws:                  ProtoToIamAlphaWorkloadIdentityPoolProviderAws(p.GetAws()),
		Oidc:                 ProtoToIamAlphaWorkloadIdentityPoolProviderOidc(p.GetOidc()),
		Project:              dcl.StringOrNil(p.Project),
		Location:             dcl.StringOrNil(p.Location),
		WorkloadIdentityPool: dcl.StringOrNil(p.WorkloadIdentityPool),
	}
	return obj
}

// WorkloadIdentityPoolProviderStateEnumToProto converts a WorkloadIdentityPoolProviderStateEnum enum to its proto representation.
func IamAlphaWorkloadIdentityPoolProviderStateEnumToProto(e *alpha.WorkloadIdentityPoolProviderStateEnum) alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum {
	if e == nil {
		return alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum_value["WorkloadIdentityPoolProviderStateEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum(v)
	}
	return alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum(0)
}

// WorkloadIdentityPoolProviderAwsToProto converts a WorkloadIdentityPoolProviderAws resource to its proto representation.
func IamAlphaWorkloadIdentityPoolProviderAwsToProto(o *alpha.WorkloadIdentityPoolProviderAws) *alphapb.IamAlphaWorkloadIdentityPoolProviderAws {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkloadIdentityPoolProviderAws{
		AccountId: dcl.ValueOrEmptyString(o.AccountId),
	}
	for _, r := range o.StsUri {
		p.StsUri = append(p.StsUri, r)
	}
	return p
}

// WorkloadIdentityPoolProviderOidcToProto converts a WorkloadIdentityPoolProviderOidc resource to its proto representation.
func IamAlphaWorkloadIdentityPoolProviderOidcToProto(o *alpha.WorkloadIdentityPoolProviderOidc) *alphapb.IamAlphaWorkloadIdentityPoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.ValueOrEmptyString(o.IssuerUri),
	}
	for _, r := range o.AllowedAudiences {
		p.AllowedAudiences = append(p.AllowedAudiences, r)
	}
	return p
}

// WorkloadIdentityPoolProviderToProto converts a WorkloadIdentityPoolProvider resource to its proto representation.
func WorkloadIdentityPoolProviderToProto(resource *alpha.WorkloadIdentityPoolProvider) *alphapb.IamAlphaWorkloadIdentityPoolProvider {
	p := &alphapb.IamAlphaWorkloadIdentityPoolProvider{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		DisplayName:          dcl.ValueOrEmptyString(resource.DisplayName),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		State:                IamAlphaWorkloadIdentityPoolProviderStateEnumToProto(resource.State),
		Disabled:             dcl.ValueOrEmptyBool(resource.Disabled),
		AttributeCondition:   dcl.ValueOrEmptyString(resource.AttributeCondition),
		Aws:                  IamAlphaWorkloadIdentityPoolProviderAwsToProto(resource.Aws),
		Oidc:                 IamAlphaWorkloadIdentityPoolProviderOidcToProto(resource.Oidc),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Location:             dcl.ValueOrEmptyString(resource.Location),
		WorkloadIdentityPool: dcl.ValueOrEmptyString(resource.WorkloadIdentityPool),
	}

	return p
}

// ApplyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) applyWorkloadIdentityPoolProvider(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIamAlphaWorkloadIdentityPoolProviderRequest) (*alphapb.IamAlphaWorkloadIdentityPoolProvider, error) {
	p := ProtoToWorkloadIdentityPoolProvider(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolProviderToProto(res)
	return r, nil
}

// ApplyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) ApplyIamAlphaWorkloadIdentityPoolProvider(ctx context.Context, request *alphapb.ApplyIamAlphaWorkloadIdentityPoolProviderRequest) (*alphapb.IamAlphaWorkloadIdentityPoolProvider, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPoolProvider(ctx, cl, request)
}

// DeleteWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Delete() method.
func (s *WorkloadIdentityPoolProviderServer) DeleteIamAlphaWorkloadIdentityPoolProvider(ctx context.Context, request *alphapb.DeleteIamAlphaWorkloadIdentityPoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))

}

// ListIamAlphaWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProviderList() method.
func (s *WorkloadIdentityPoolProviderServer) ListIamAlphaWorkloadIdentityPoolProvider(ctx context.Context, request *alphapb.ListIamAlphaWorkloadIdentityPoolProviderRequest) (*alphapb.ListIamAlphaWorkloadIdentityPoolProviderResponse, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IamAlphaWorkloadIdentityPoolProvider
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolProviderToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListIamAlphaWorkloadIdentityPoolProviderResponse{Items: protos}, nil
}

func createConfigWorkloadIdentityPoolProvider(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
