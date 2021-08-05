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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/beta/identitytoolkit_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta"
)

// Server implements the gRPC interface for TenantOAuthIdpConfig.
type TenantOAuthIdpConfigServer struct{}

// ProtoToTenantOAuthIdpConfigResponseType converts a TenantOAuthIdpConfigResponseType resource from its proto representation.
func ProtoToIdentitytoolkitBetaTenantOAuthIdpConfigResponseType(p *betapb.IdentitytoolkitBetaTenantOAuthIdpConfigResponseType) *beta.TenantOAuthIdpConfigResponseType {
	if p == nil {
		return nil
	}
	obj := &beta.TenantOAuthIdpConfigResponseType{
		IdToken: dcl.Bool(p.IdToken),
		Code:    dcl.Bool(p.Code),
		Token:   dcl.Bool(p.Token),
	}
	return obj
}

// ProtoToTenantOAuthIdpConfig converts a TenantOAuthIdpConfig resource from its proto representation.
func ProtoToTenantOAuthIdpConfig(p *betapb.IdentitytoolkitBetaTenantOAuthIdpConfig) *beta.TenantOAuthIdpConfig {
	obj := &beta.TenantOAuthIdpConfig{
		Name:         dcl.StringOrNil(p.Name),
		ClientId:     dcl.StringOrNil(p.ClientId),
		Issuer:       dcl.StringOrNil(p.Issuer),
		DisplayName:  dcl.StringOrNil(p.DisplayName),
		Enabled:      dcl.Bool(p.Enabled),
		ClientSecret: dcl.StringOrNil(p.ClientSecret),
		ResponseType: ProtoToIdentitytoolkitBetaTenantOAuthIdpConfigResponseType(p.GetResponseType()),
		Project:      dcl.StringOrNil(p.Project),
		Tenant:       dcl.StringOrNil(p.Tenant),
	}
	return obj
}

// TenantOAuthIdpConfigResponseTypeToProto converts a TenantOAuthIdpConfigResponseType resource to its proto representation.
func IdentitytoolkitBetaTenantOAuthIdpConfigResponseTypeToProto(o *beta.TenantOAuthIdpConfigResponseType) *betapb.IdentitytoolkitBetaTenantOAuthIdpConfigResponseType {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaTenantOAuthIdpConfigResponseType{
		IdToken: dcl.ValueOrEmptyBool(o.IdToken),
		Code:    dcl.ValueOrEmptyBool(o.Code),
		Token:   dcl.ValueOrEmptyBool(o.Token),
	}
	return p
}

// TenantOAuthIdpConfigToProto converts a TenantOAuthIdpConfig resource to its proto representation.
func TenantOAuthIdpConfigToProto(resource *beta.TenantOAuthIdpConfig) *betapb.IdentitytoolkitBetaTenantOAuthIdpConfig {
	p := &betapb.IdentitytoolkitBetaTenantOAuthIdpConfig{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		ClientId:     dcl.ValueOrEmptyString(resource.ClientId),
		Issuer:       dcl.ValueOrEmptyString(resource.Issuer),
		DisplayName:  dcl.ValueOrEmptyString(resource.DisplayName),
		Enabled:      dcl.ValueOrEmptyBool(resource.Enabled),
		ClientSecret: dcl.ValueOrEmptyString(resource.ClientSecret),
		ResponseType: IdentitytoolkitBetaTenantOAuthIdpConfigResponseTypeToProto(resource.ResponseType),
		Project:      dcl.ValueOrEmptyString(resource.Project),
		Tenant:       dcl.ValueOrEmptyString(resource.Tenant),
	}

	return p
}

// ApplyTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Apply() method.
func (s *TenantOAuthIdpConfigServer) applyTenantOAuthIdpConfig(ctx context.Context, c *beta.Client, request *betapb.ApplyIdentitytoolkitBetaTenantOAuthIdpConfigRequest) (*betapb.IdentitytoolkitBetaTenantOAuthIdpConfig, error) {
	p := ProtoToTenantOAuthIdpConfig(request.GetResource())
	res, err := c.ApplyTenantOAuthIdpConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TenantOAuthIdpConfigToProto(res)
	return r, nil
}

// ApplyTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Apply() method.
func (s *TenantOAuthIdpConfigServer) ApplyIdentitytoolkitBetaTenantOAuthIdpConfig(ctx context.Context, request *betapb.ApplyIdentitytoolkitBetaTenantOAuthIdpConfigRequest) (*betapb.IdentitytoolkitBetaTenantOAuthIdpConfig, error) {
	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTenantOAuthIdpConfig(ctx, cl, request)
}

// DeleteTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Delete() method.
func (s *TenantOAuthIdpConfigServer) DeleteIdentitytoolkitBetaTenantOAuthIdpConfig(ctx context.Context, request *betapb.DeleteIdentitytoolkitBetaTenantOAuthIdpConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTenantOAuthIdpConfig(ctx, ProtoToTenantOAuthIdpConfig(request.GetResource()))

}

// ListIdentitytoolkitBetaTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfigList() method.
func (s *TenantOAuthIdpConfigServer) ListIdentitytoolkitBetaTenantOAuthIdpConfig(ctx context.Context, request *betapb.ListIdentitytoolkitBetaTenantOAuthIdpConfigRequest) (*betapb.ListIdentitytoolkitBetaTenantOAuthIdpConfigResponse, error) {
	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTenantOAuthIdpConfig(ctx, request.Project, request.Tenant)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IdentitytoolkitBetaTenantOAuthIdpConfig
	for _, r := range resources.Items {
		rp := TenantOAuthIdpConfigToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListIdentitytoolkitBetaTenantOAuthIdpConfigResponse{Items: protos}, nil
}

func createConfigTenantOAuthIdpConfig(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
