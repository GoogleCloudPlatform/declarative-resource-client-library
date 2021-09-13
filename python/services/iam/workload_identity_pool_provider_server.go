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

// Server implements the gRPC interface for WorkloadIdentityPoolProvider.
type WorkloadIdentityPoolProviderServer struct{}

// ProtoToWorkloadIdentityPoolProviderStateEnum converts a WorkloadIdentityPoolProviderStateEnum enum from its proto representation.
func ProtoToIamWorkloadIdentityPoolProviderStateEnum(e iampb.IamWorkloadIdentityPoolProviderStateEnum) *iam.WorkloadIdentityPoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := iampb.IamWorkloadIdentityPoolProviderStateEnum_name[int32(e)]; ok {
		e := iam.WorkloadIdentityPoolProviderStateEnum(n[len("IamWorkloadIdentityPoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPoolProviderAws converts a WorkloadIdentityPoolProviderAws resource from its proto representation.
func ProtoToIamWorkloadIdentityPoolProviderAws(p *iampb.IamWorkloadIdentityPoolProviderAws) *iam.WorkloadIdentityPoolProviderAws {
	if p == nil {
		return nil
	}
	obj := &iam.WorkloadIdentityPoolProviderAws{
		AccountId: dcl.StringOrNil(p.AccountId),
	}
	for _, r := range p.GetStsUri() {
		obj.StsUri = append(obj.StsUri, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProviderOidc converts a WorkloadIdentityPoolProviderOidc resource from its proto representation.
func ProtoToIamWorkloadIdentityPoolProviderOidc(p *iampb.IamWorkloadIdentityPoolProviderOidc) *iam.WorkloadIdentityPoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &iam.WorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.StringOrNil(p.IssuerUri),
	}
	for _, r := range p.GetAllowedAudiences() {
		obj.AllowedAudiences = append(obj.AllowedAudiences, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProvider converts a WorkloadIdentityPoolProvider resource from its proto representation.
func ProtoToWorkloadIdentityPoolProvider(p *iampb.IamWorkloadIdentityPoolProvider) *iam.WorkloadIdentityPoolProvider {
	obj := &iam.WorkloadIdentityPoolProvider{
		Name:                 dcl.StringOrNil(p.Name),
		DisplayName:          dcl.StringOrNil(p.DisplayName),
		Description:          dcl.StringOrNil(p.Description),
		State:                ProtoToIamWorkloadIdentityPoolProviderStateEnum(p.GetState()),
		Disabled:             dcl.Bool(p.Disabled),
		AttributeCondition:   dcl.StringOrNil(p.AttributeCondition),
		Aws:                  ProtoToIamWorkloadIdentityPoolProviderAws(p.GetAws()),
		Oidc:                 ProtoToIamWorkloadIdentityPoolProviderOidc(p.GetOidc()),
		Project:              dcl.StringOrNil(p.Project),
		Location:             dcl.StringOrNil(p.Location),
		WorkloadIdentityPool: dcl.StringOrNil(p.WorkloadIdentityPool),
	}
	return obj
}

// WorkloadIdentityPoolProviderStateEnumToProto converts a WorkloadIdentityPoolProviderStateEnum enum to its proto representation.
func IamWorkloadIdentityPoolProviderStateEnumToProto(e *iam.WorkloadIdentityPoolProviderStateEnum) iampb.IamWorkloadIdentityPoolProviderStateEnum {
	if e == nil {
		return iampb.IamWorkloadIdentityPoolProviderStateEnum(0)
	}
	if v, ok := iampb.IamWorkloadIdentityPoolProviderStateEnum_value["WorkloadIdentityPoolProviderStateEnum"+string(*e)]; ok {
		return iampb.IamWorkloadIdentityPoolProviderStateEnum(v)
	}
	return iampb.IamWorkloadIdentityPoolProviderStateEnum(0)
}

// WorkloadIdentityPoolProviderAwsToProto converts a WorkloadIdentityPoolProviderAws resource to its proto representation.
func IamWorkloadIdentityPoolProviderAwsToProto(o *iam.WorkloadIdentityPoolProviderAws) *iampb.IamWorkloadIdentityPoolProviderAws {
	if o == nil {
		return nil
	}
	p := &iampb.IamWorkloadIdentityPoolProviderAws{
		AccountId: dcl.ValueOrEmptyString(o.AccountId),
	}
	for _, r := range o.StsUri {
		p.StsUri = append(p.StsUri, r)
	}
	return p
}

// WorkloadIdentityPoolProviderOidcToProto converts a WorkloadIdentityPoolProviderOidc resource to its proto representation.
func IamWorkloadIdentityPoolProviderOidcToProto(o *iam.WorkloadIdentityPoolProviderOidc) *iampb.IamWorkloadIdentityPoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &iampb.IamWorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.ValueOrEmptyString(o.IssuerUri),
	}
	for _, r := range o.AllowedAudiences {
		p.AllowedAudiences = append(p.AllowedAudiences, r)
	}
	return p
}

// WorkloadIdentityPoolProviderToProto converts a WorkloadIdentityPoolProvider resource to its proto representation.
func WorkloadIdentityPoolProviderToProto(resource *iam.WorkloadIdentityPoolProvider) *iampb.IamWorkloadIdentityPoolProvider {
	p := &iampb.IamWorkloadIdentityPoolProvider{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		DisplayName:          dcl.ValueOrEmptyString(resource.DisplayName),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		State:                IamWorkloadIdentityPoolProviderStateEnumToProto(resource.State),
		Disabled:             dcl.ValueOrEmptyBool(resource.Disabled),
		AttributeCondition:   dcl.ValueOrEmptyString(resource.AttributeCondition),
		Aws:                  IamWorkloadIdentityPoolProviderAwsToProto(resource.Aws),
		Oidc:                 IamWorkloadIdentityPoolProviderOidcToProto(resource.Oidc),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Location:             dcl.ValueOrEmptyString(resource.Location),
		WorkloadIdentityPool: dcl.ValueOrEmptyString(resource.WorkloadIdentityPool),
	}

	return p
}

// ApplyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) applyWorkloadIdentityPoolProvider(ctx context.Context, c *iam.Client, request *iampb.ApplyIamWorkloadIdentityPoolProviderRequest) (*iampb.IamWorkloadIdentityPoolProvider, error) {
	p := ProtoToWorkloadIdentityPoolProvider(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolProviderToProto(res)
	return r, nil
}

// ApplyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) ApplyIamWorkloadIdentityPoolProvider(ctx context.Context, request *iampb.ApplyIamWorkloadIdentityPoolProviderRequest) (*iampb.IamWorkloadIdentityPoolProvider, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPoolProvider(ctx, cl, request)
}

// DeleteWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Delete() method.
func (s *WorkloadIdentityPoolProviderServer) DeleteIamWorkloadIdentityPoolProvider(ctx context.Context, request *iampb.DeleteIamWorkloadIdentityPoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))

}

// ListIamWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProviderList() method.
func (s *WorkloadIdentityPoolProviderServer) ListIamWorkloadIdentityPoolProvider(ctx context.Context, request *iampb.ListIamWorkloadIdentityPoolProviderRequest) (*iampb.ListIamWorkloadIdentityPoolProviderResponse, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*iampb.IamWorkloadIdentityPoolProvider
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolProviderToProto(r)
		protos = append(protos, rp)
	}
	return &iampb.ListIamWorkloadIdentityPoolProviderResponse{Items: protos}, nil
}

func createConfigWorkloadIdentityPoolProvider(ctx context.Context, service_account_file string) (*iam.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iam.NewClient(conf), nil
}
