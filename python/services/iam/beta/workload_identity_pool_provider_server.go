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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/beta/iam_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta"
)

// Server implements the gRPC interface for WorkloadIdentityPoolProvider.
type WorkloadIdentityPoolProviderServer struct{}

// ProtoToWorkloadIdentityPoolProviderStateEnum converts a WorkloadIdentityPoolProviderStateEnum enum from its proto representation.
func ProtoToIamBetaWorkloadIdentityPoolProviderStateEnum(e betapb.IamBetaWorkloadIdentityPoolProviderStateEnum) *beta.WorkloadIdentityPoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IamBetaWorkloadIdentityPoolProviderStateEnum_name[int32(e)]; ok {
		e := beta.WorkloadIdentityPoolProviderStateEnum(n[len("IamBetaWorkloadIdentityPoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPoolProviderAws converts a WorkloadIdentityPoolProviderAws resource from its proto representation.
func ProtoToIamBetaWorkloadIdentityPoolProviderAws(p *betapb.IamBetaWorkloadIdentityPoolProviderAws) *beta.WorkloadIdentityPoolProviderAws {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadIdentityPoolProviderAws{
		AccountId: dcl.StringOrNil(p.AccountId),
	}
	for _, r := range p.GetStsUri() {
		obj.StsUri = append(obj.StsUri, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProviderOidc converts a WorkloadIdentityPoolProviderOidc resource from its proto representation.
func ProtoToIamBetaWorkloadIdentityPoolProviderOidc(p *betapb.IamBetaWorkloadIdentityPoolProviderOidc) *beta.WorkloadIdentityPoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.StringOrNil(p.IssuerUri),
	}
	for _, r := range p.GetAllowedAudiences() {
		obj.AllowedAudiences = append(obj.AllowedAudiences, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProvider converts a WorkloadIdentityPoolProvider resource from its proto representation.
func ProtoToWorkloadIdentityPoolProvider(p *betapb.IamBetaWorkloadIdentityPoolProvider) *beta.WorkloadIdentityPoolProvider {
	obj := &beta.WorkloadIdentityPoolProvider{
		Name:                 dcl.StringOrNil(p.Name),
		DisplayName:          dcl.StringOrNil(p.DisplayName),
		Description:          dcl.StringOrNil(p.Description),
		State:                ProtoToIamBetaWorkloadIdentityPoolProviderStateEnum(p.GetState()),
		Disabled:             dcl.Bool(p.Disabled),
		AttributeCondition:   dcl.StringOrNil(p.AttributeCondition),
		Aws:                  ProtoToIamBetaWorkloadIdentityPoolProviderAws(p.GetAws()),
		Oidc:                 ProtoToIamBetaWorkloadIdentityPoolProviderOidc(p.GetOidc()),
		Project:              dcl.StringOrNil(p.Project),
		Location:             dcl.StringOrNil(p.Location),
		WorkloadIdentityPool: dcl.StringOrNil(p.WorkloadIdentityPool),
	}
	return obj
}

// WorkloadIdentityPoolProviderStateEnumToProto converts a WorkloadIdentityPoolProviderStateEnum enum to its proto representation.
func IamBetaWorkloadIdentityPoolProviderStateEnumToProto(e *beta.WorkloadIdentityPoolProviderStateEnum) betapb.IamBetaWorkloadIdentityPoolProviderStateEnum {
	if e == nil {
		return betapb.IamBetaWorkloadIdentityPoolProviderStateEnum(0)
	}
	if v, ok := betapb.IamBetaWorkloadIdentityPoolProviderStateEnum_value["WorkloadIdentityPoolProviderStateEnum"+string(*e)]; ok {
		return betapb.IamBetaWorkloadIdentityPoolProviderStateEnum(v)
	}
	return betapb.IamBetaWorkloadIdentityPoolProviderStateEnum(0)
}

// WorkloadIdentityPoolProviderAwsToProto converts a WorkloadIdentityPoolProviderAws resource to its proto representation.
func IamBetaWorkloadIdentityPoolProviderAwsToProto(o *beta.WorkloadIdentityPoolProviderAws) *betapb.IamBetaWorkloadIdentityPoolProviderAws {
	if o == nil {
		return nil
	}
	p := &betapb.IamBetaWorkloadIdentityPoolProviderAws{
		AccountId: dcl.ValueOrEmptyString(o.AccountId),
	}
	for _, r := range o.StsUri {
		p.StsUri = append(p.StsUri, r)
	}
	return p
}

// WorkloadIdentityPoolProviderOidcToProto converts a WorkloadIdentityPoolProviderOidc resource to its proto representation.
func IamBetaWorkloadIdentityPoolProviderOidcToProto(o *beta.WorkloadIdentityPoolProviderOidc) *betapb.IamBetaWorkloadIdentityPoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &betapb.IamBetaWorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.ValueOrEmptyString(o.IssuerUri),
	}
	for _, r := range o.AllowedAudiences {
		p.AllowedAudiences = append(p.AllowedAudiences, r)
	}
	return p
}

// WorkloadIdentityPoolProviderToProto converts a WorkloadIdentityPoolProvider resource to its proto representation.
func WorkloadIdentityPoolProviderToProto(resource *beta.WorkloadIdentityPoolProvider) *betapb.IamBetaWorkloadIdentityPoolProvider {
	p := &betapb.IamBetaWorkloadIdentityPoolProvider{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		DisplayName:          dcl.ValueOrEmptyString(resource.DisplayName),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		State:                IamBetaWorkloadIdentityPoolProviderStateEnumToProto(resource.State),
		Disabled:             dcl.ValueOrEmptyBool(resource.Disabled),
		AttributeCondition:   dcl.ValueOrEmptyString(resource.AttributeCondition),
		Aws:                  IamBetaWorkloadIdentityPoolProviderAwsToProto(resource.Aws),
		Oidc:                 IamBetaWorkloadIdentityPoolProviderOidcToProto(resource.Oidc),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Location:             dcl.ValueOrEmptyString(resource.Location),
		WorkloadIdentityPool: dcl.ValueOrEmptyString(resource.WorkloadIdentityPool),
	}

	return p
}

// ApplyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) applyWorkloadIdentityPoolProvider(ctx context.Context, c *beta.Client, request *betapb.ApplyIamBetaWorkloadIdentityPoolProviderRequest) (*betapb.IamBetaWorkloadIdentityPoolProvider, error) {
	p := ProtoToWorkloadIdentityPoolProvider(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolProviderToProto(res)
	return r, nil
}

// ApplyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) ApplyIamBetaWorkloadIdentityPoolProvider(ctx context.Context, request *betapb.ApplyIamBetaWorkloadIdentityPoolProviderRequest) (*betapb.IamBetaWorkloadIdentityPoolProvider, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPoolProvider(ctx, cl, request)
}

// DeleteWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Delete() method.
func (s *WorkloadIdentityPoolProviderServer) DeleteIamBetaWorkloadIdentityPoolProvider(ctx context.Context, request *betapb.DeleteIamBetaWorkloadIdentityPoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))

}

// ListIamBetaWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProviderList() method.
func (s *WorkloadIdentityPoolProviderServer) ListIamBetaWorkloadIdentityPoolProvider(ctx context.Context, request *betapb.ListIamBetaWorkloadIdentityPoolProviderRequest) (*betapb.ListIamBetaWorkloadIdentityPoolProviderResponse, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IamBetaWorkloadIdentityPoolProvider
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolProviderToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListIamBetaWorkloadIdentityPoolProviderResponse{Items: protos}, nil
}

func createConfigWorkloadIdentityPoolProvider(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
