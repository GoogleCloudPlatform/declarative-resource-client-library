// Copyright 2022 Google LLC. All Rights Reserved.
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

// WorkforcePoolProviderServer implements the gRPC interface for WorkforcePoolProvider.
type WorkforcePoolProviderServer struct{}

// ProtoToWorkforcePoolProviderStateEnum converts a WorkforcePoolProviderStateEnum enum from its proto representation.
func ProtoToIamBetaWorkforcePoolProviderStateEnum(e betapb.IamBetaWorkforcePoolProviderStateEnum) *beta.WorkforcePoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IamBetaWorkforcePoolProviderStateEnum_name[int32(e)]; ok {
		e := beta.WorkforcePoolProviderStateEnum(n[len("IamBetaWorkforcePoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePoolProviderSaml converts a WorkforcePoolProviderSaml object from its proto representation.
func ProtoToIamBetaWorkforcePoolProviderSaml(p *betapb.IamBetaWorkforcePoolProviderSaml) *beta.WorkforcePoolProviderSaml {
	if p == nil {
		return nil
	}
	obj := &beta.WorkforcePoolProviderSaml{
		IdpMetadataXml: dcl.StringOrNil(p.GetIdpMetadataXml()),
	}
	return obj
}

// ProtoToWorkforcePoolProviderOidc converts a WorkforcePoolProviderOidc object from its proto representation.
func ProtoToIamBetaWorkforcePoolProviderOidc(p *betapb.IamBetaWorkforcePoolProviderOidc) *beta.WorkforcePoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &beta.WorkforcePoolProviderOidc{
		IssuerUri: dcl.StringOrNil(p.GetIssuerUri()),
		ClientId:  dcl.StringOrNil(p.GetClientId()),
	}
	return obj
}

// ProtoToWorkforcePoolProvider converts a WorkforcePoolProvider resource from its proto representation.
func ProtoToWorkforcePoolProvider(p *betapb.IamBetaWorkforcePoolProvider) *beta.WorkforcePoolProvider {
	obj := &beta.WorkforcePoolProvider{
		Name:               dcl.StringOrNil(p.GetName()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToIamBetaWorkforcePoolProviderStateEnum(p.GetState()),
		Disabled:           dcl.Bool(p.GetDisabled()),
		AttributeCondition: dcl.StringOrNil(p.GetAttributeCondition()),
		Saml:               ProtoToIamBetaWorkforcePoolProviderSaml(p.GetSaml()),
		Oidc:               ProtoToIamBetaWorkforcePoolProviderOidc(p.GetOidc()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		WorkforcePool:      dcl.StringOrNil(p.GetWorkforcePool()),
	}
	return obj
}

// WorkforcePoolProviderStateEnumToProto converts a WorkforcePoolProviderStateEnum enum to its proto representation.
func IamBetaWorkforcePoolProviderStateEnumToProto(e *beta.WorkforcePoolProviderStateEnum) betapb.IamBetaWorkforcePoolProviderStateEnum {
	if e == nil {
		return betapb.IamBetaWorkforcePoolProviderStateEnum(0)
	}
	if v, ok := betapb.IamBetaWorkforcePoolProviderStateEnum_value["WorkforcePoolProviderStateEnum"+string(*e)]; ok {
		return betapb.IamBetaWorkforcePoolProviderStateEnum(v)
	}
	return betapb.IamBetaWorkforcePoolProviderStateEnum(0)
}

// WorkforcePoolProviderSamlToProto converts a WorkforcePoolProviderSaml object to its proto representation.
func IamBetaWorkforcePoolProviderSamlToProto(o *beta.WorkforcePoolProviderSaml) *betapb.IamBetaWorkforcePoolProviderSaml {
	if o == nil {
		return nil
	}
	p := &betapb.IamBetaWorkforcePoolProviderSaml{}
	p.SetIdpMetadataXml(dcl.ValueOrEmptyString(o.IdpMetadataXml))
	return p
}

// WorkforcePoolProviderOidcToProto converts a WorkforcePoolProviderOidc object to its proto representation.
func IamBetaWorkforcePoolProviderOidcToProto(o *beta.WorkforcePoolProviderOidc) *betapb.IamBetaWorkforcePoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &betapb.IamBetaWorkforcePoolProviderOidc{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetClientId(dcl.ValueOrEmptyString(o.ClientId))
	return p
}

// WorkforcePoolProviderToProto converts a WorkforcePoolProvider resource to its proto representation.
func WorkforcePoolProviderToProto(resource *beta.WorkforcePoolProvider) *betapb.IamBetaWorkforcePoolProvider {
	p := &betapb.IamBetaWorkforcePoolProvider{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamBetaWorkforcePoolProviderStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetAttributeCondition(dcl.ValueOrEmptyString(resource.AttributeCondition))
	p.SetSaml(IamBetaWorkforcePoolProviderSamlToProto(resource.Saml))
	p.SetOidc(IamBetaWorkforcePoolProviderOidcToProto(resource.Oidc))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetWorkforcePool(dcl.ValueOrEmptyString(resource.WorkforcePool))
	mAttributeMapping := make(map[string]string, len(resource.AttributeMapping))
	for k, r := range resource.AttributeMapping {
		mAttributeMapping[k] = r
	}
	p.SetAttributeMapping(mAttributeMapping)

	return p
}

// applyWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Apply() method.
func (s *WorkforcePoolProviderServer) applyWorkforcePoolProvider(ctx context.Context, c *beta.Client, request *betapb.ApplyIamBetaWorkforcePoolProviderRequest) (*betapb.IamBetaWorkforcePoolProvider, error) {
	p := ProtoToWorkforcePoolProvider(request.GetResource())
	res, err := c.ApplyWorkforcePoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkforcePoolProviderToProto(res)
	return r, nil
}

// applyIamBetaWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Apply() method.
func (s *WorkforcePoolProviderServer) ApplyIamBetaWorkforcePoolProvider(ctx context.Context, request *betapb.ApplyIamBetaWorkforcePoolProviderRequest) (*betapb.IamBetaWorkforcePoolProvider, error) {
	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkforcePoolProvider(ctx, cl, request)
}

// DeleteWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Delete() method.
func (s *WorkforcePoolProviderServer) DeleteIamBetaWorkforcePoolProvider(ctx context.Context, request *betapb.DeleteIamBetaWorkforcePoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkforcePoolProvider(ctx, ProtoToWorkforcePoolProvider(request.GetResource()))

}

// ListIamBetaWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProviderList() method.
func (s *WorkforcePoolProviderServer) ListIamBetaWorkforcePoolProvider(ctx context.Context, request *betapb.ListIamBetaWorkforcePoolProviderRequest) (*betapb.ListIamBetaWorkforcePoolProviderResponse, error) {
	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkforcePoolProvider(ctx, request.GetLocation(), request.GetWorkforcePool())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IamBetaWorkforcePoolProvider
	for _, r := range resources.Items {
		rp := WorkforcePoolProviderToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListIamBetaWorkforcePoolProviderResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkforcePoolProvider(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
