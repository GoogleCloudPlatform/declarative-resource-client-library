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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/alpha/identitytoolkit_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha"
)

// Server implements the gRPC interface for TenantOAuthIdpConfig.
type TenantOAuthIdpConfigServer struct{}

// ProtoToTenantOAuthIdpConfigResponseType converts a TenantOAuthIdpConfigResponseType resource from its proto representation.
func ProtoToIdentitytoolkitAlphaTenantOAuthIdpConfigResponseType(p *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType) *alpha.TenantOAuthIdpConfigResponseType {
	if p == nil {
		return nil
	}
	obj := &alpha.TenantOAuthIdpConfigResponseType{
		IdToken: dcl.Bool(p.IdToken),
		Code:    dcl.Bool(p.Code),
		Token:   dcl.Bool(p.Token),
	}
	return obj
}

// ProtoToTenantOAuthIdpConfig converts a TenantOAuthIdpConfig resource from its proto representation.
func ProtoToTenantOAuthIdpConfig(p *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig) *alpha.TenantOAuthIdpConfig {
	obj := &alpha.TenantOAuthIdpConfig{
		Name:         dcl.StringOrNil(p.Name),
		ClientId:     dcl.StringOrNil(p.ClientId),
		Issuer:       dcl.StringOrNil(p.Issuer),
		DisplayName:  dcl.StringOrNil(p.DisplayName),
		Enabled:      dcl.Bool(p.Enabled),
		ClientSecret: dcl.StringOrNil(p.ClientSecret),
		ResponseType: ProtoToIdentitytoolkitAlphaTenantOAuthIdpConfigResponseType(p.GetResponseType()),
		Project:      dcl.StringOrNil(p.Project),
		Tenant:       dcl.StringOrNil(p.Tenant),
	}
	return obj
}

// TenantOAuthIdpConfigResponseTypeToProto converts a TenantOAuthIdpConfigResponseType resource to its proto representation.
func IdentitytoolkitAlphaTenantOAuthIdpConfigResponseTypeToProto(o *alpha.TenantOAuthIdpConfigResponseType) *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType{
		IdToken: dcl.ValueOrEmptyBool(o.IdToken),
		Code:    dcl.ValueOrEmptyBool(o.Code),
		Token:   dcl.ValueOrEmptyBool(o.Token),
	}
	return p
}

// TenantOAuthIdpConfigToProto converts a TenantOAuthIdpConfig resource to its proto representation.
func TenantOAuthIdpConfigToProto(resource *alpha.TenantOAuthIdpConfig) *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig {
	p := &alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		ClientId:     dcl.ValueOrEmptyString(resource.ClientId),
		Issuer:       dcl.ValueOrEmptyString(resource.Issuer),
		DisplayName:  dcl.ValueOrEmptyString(resource.DisplayName),
		Enabled:      dcl.ValueOrEmptyBool(resource.Enabled),
		ClientSecret: dcl.ValueOrEmptyString(resource.ClientSecret),
		ResponseType: IdentitytoolkitAlphaTenantOAuthIdpConfigResponseTypeToProto(resource.ResponseType),
		Project:      dcl.ValueOrEmptyString(resource.Project),
		Tenant:       dcl.ValueOrEmptyString(resource.Tenant),
	}

	return p
}

// ApplyTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Apply() method.
func (s *TenantOAuthIdpConfigServer) applyTenantOAuthIdpConfig(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig, error) {
	p := ProtoToTenantOAuthIdpConfig(request.GetResource())
	res, err := c.ApplyTenantOAuthIdpConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TenantOAuthIdpConfigToProto(res)
	return r, nil
}

// ApplyTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Apply() method.
func (s *TenantOAuthIdpConfigServer) ApplyIdentitytoolkitAlphaTenantOAuthIdpConfig(ctx context.Context, request *alphapb.ApplyIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig, error) {
	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTenantOAuthIdpConfig(ctx, cl, request)
}

// DeleteTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Delete() method.
func (s *TenantOAuthIdpConfigServer) DeleteIdentitytoolkitAlphaTenantOAuthIdpConfig(ctx context.Context, request *alphapb.DeleteIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTenantOAuthIdpConfig(ctx, ProtoToTenantOAuthIdpConfig(request.GetResource()))

}

// ListIdentitytoolkitAlphaTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfigList() method.
func (s *TenantOAuthIdpConfigServer) ListIdentitytoolkitAlphaTenantOAuthIdpConfig(ctx context.Context, request *alphapb.ListIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*alphapb.ListIdentitytoolkitAlphaTenantOAuthIdpConfigResponse, error) {
	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTenantOAuthIdpConfig(ctx, ProtoToTenantOAuthIdpConfig(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig
	for _, r := range resources.Items {
		rp := TenantOAuthIdpConfigToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListIdentitytoolkitAlphaTenantOAuthIdpConfigResponse{Items: protos}, nil
}

func createConfigTenantOAuthIdpConfig(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
