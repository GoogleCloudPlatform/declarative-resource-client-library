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
	iampb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/iam_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
)

// WorkforcePoolProviderServer implements the gRPC interface for WorkforcePoolProvider.
type WorkforcePoolProviderServer struct{}

// ProtoToWorkforcePoolProviderStateEnum converts a WorkforcePoolProviderStateEnum enum from its proto representation.
func ProtoToIamWorkforcePoolProviderStateEnum(e iampb.IamWorkforcePoolProviderStateEnum) *iam.WorkforcePoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := iampb.IamWorkforcePoolProviderStateEnum_name[int32(e)]; ok {
		e := iam.WorkforcePoolProviderStateEnum(n[len("IamWorkforcePoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePoolProviderSaml converts a WorkforcePoolProviderSaml object from its proto representation.
func ProtoToIamWorkforcePoolProviderSaml(p *iampb.IamWorkforcePoolProviderSaml) *iam.WorkforcePoolProviderSaml {
	if p == nil {
		return nil
	}
	obj := &iam.WorkforcePoolProviderSaml{
		IdpMetadataXml: dcl.StringOrNil(p.GetIdpMetadataXml()),
	}
	return obj
}

// ProtoToWorkforcePoolProviderOidc converts a WorkforcePoolProviderOidc object from its proto representation.
func ProtoToIamWorkforcePoolProviderOidc(p *iampb.IamWorkforcePoolProviderOidc) *iam.WorkforcePoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &iam.WorkforcePoolProviderOidc{
		IssuerUri: dcl.StringOrNil(p.GetIssuerUri()),
		ClientId:  dcl.StringOrNil(p.GetClientId()),
	}
	return obj
}

// ProtoToWorkforcePoolProvider converts a WorkforcePoolProvider resource from its proto representation.
func ProtoToWorkforcePoolProvider(p *iampb.IamWorkforcePoolProvider) *iam.WorkforcePoolProvider {
	obj := &iam.WorkforcePoolProvider{
		Name:               dcl.StringOrNil(p.GetName()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToIamWorkforcePoolProviderStateEnum(p.GetState()),
		Disabled:           dcl.Bool(p.GetDisabled()),
		AttributeCondition: dcl.StringOrNil(p.GetAttributeCondition()),
		Saml:               ProtoToIamWorkforcePoolProviderSaml(p.GetSaml()),
		Oidc:               ProtoToIamWorkforcePoolProviderOidc(p.GetOidc()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		WorkforcePool:      dcl.StringOrNil(p.GetWorkforcePool()),
	}
	return obj
}

// WorkforcePoolProviderStateEnumToProto converts a WorkforcePoolProviderStateEnum enum to its proto representation.
func IamWorkforcePoolProviderStateEnumToProto(e *iam.WorkforcePoolProviderStateEnum) iampb.IamWorkforcePoolProviderStateEnum {
	if e == nil {
		return iampb.IamWorkforcePoolProviderStateEnum(0)
	}
	if v, ok := iampb.IamWorkforcePoolProviderStateEnum_value["WorkforcePoolProviderStateEnum"+string(*e)]; ok {
		return iampb.IamWorkforcePoolProviderStateEnum(v)
	}
	return iampb.IamWorkforcePoolProviderStateEnum(0)
}

// WorkforcePoolProviderSamlToProto converts a WorkforcePoolProviderSaml object to its proto representation.
func IamWorkforcePoolProviderSamlToProto(o *iam.WorkforcePoolProviderSaml) *iampb.IamWorkforcePoolProviderSaml {
	if o == nil {
		return nil
	}
	p := &iampb.IamWorkforcePoolProviderSaml{}
	p.SetIdpMetadataXml(dcl.ValueOrEmptyString(o.IdpMetadataXml))
	return p
}

// WorkforcePoolProviderOidcToProto converts a WorkforcePoolProviderOidc object to its proto representation.
func IamWorkforcePoolProviderOidcToProto(o *iam.WorkforcePoolProviderOidc) *iampb.IamWorkforcePoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &iampb.IamWorkforcePoolProviderOidc{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetClientId(dcl.ValueOrEmptyString(o.ClientId))
	return p
}

// WorkforcePoolProviderToProto converts a WorkforcePoolProvider resource to its proto representation.
func WorkforcePoolProviderToProto(resource *iam.WorkforcePoolProvider) *iampb.IamWorkforcePoolProvider {
	p := &iampb.IamWorkforcePoolProvider{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamWorkforcePoolProviderStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetAttributeCondition(dcl.ValueOrEmptyString(resource.AttributeCondition))
	p.SetSaml(IamWorkforcePoolProviderSamlToProto(resource.Saml))
	p.SetOidc(IamWorkforcePoolProviderOidcToProto(resource.Oidc))
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
func (s *WorkforcePoolProviderServer) applyWorkforcePoolProvider(ctx context.Context, c *iam.Client, request *iampb.ApplyIamWorkforcePoolProviderRequest) (*iampb.IamWorkforcePoolProvider, error) {
	p := ProtoToWorkforcePoolProvider(request.GetResource())
	res, err := c.ApplyWorkforcePoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkforcePoolProviderToProto(res)
	return r, nil
}

// applyIamWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Apply() method.
func (s *WorkforcePoolProviderServer) ApplyIamWorkforcePoolProvider(ctx context.Context, request *iampb.ApplyIamWorkforcePoolProviderRequest) (*iampb.IamWorkforcePoolProvider, error) {
	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkforcePoolProvider(ctx, cl, request)
}

// DeleteWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Delete() method.
func (s *WorkforcePoolProviderServer) DeleteIamWorkforcePoolProvider(ctx context.Context, request *iampb.DeleteIamWorkforcePoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkforcePoolProvider(ctx, ProtoToWorkforcePoolProvider(request.GetResource()))

}

// ListIamWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProviderList() method.
func (s *WorkforcePoolProviderServer) ListIamWorkforcePoolProvider(ctx context.Context, request *iampb.ListIamWorkforcePoolProviderRequest) (*iampb.ListIamWorkforcePoolProviderResponse, error) {
	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkforcePoolProvider(ctx, request.GetLocation(), request.GetWorkforcePool())
	if err != nil {
		return nil, err
	}
	var protos []*iampb.IamWorkforcePoolProvider
	for _, r := range resources.Items {
		rp := WorkforcePoolProviderToProto(r)
		protos = append(protos, rp)
	}
	p := &iampb.ListIamWorkforcePoolProviderResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkforcePoolProvider(ctx context.Context, service_account_file string) (*iam.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iam.NewClient(conf), nil
}
