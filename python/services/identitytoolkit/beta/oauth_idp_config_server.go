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

// Server implements the gRPC interface for OAuthIdpConfig.
type OAuthIdpConfigServer struct{}

// ProtoToOAuthIdpConfigResponseType converts a OAuthIdpConfigResponseType resource from its proto representation.
func ProtoToIdentitytoolkitBetaOAuthIdpConfigResponseType(p *betapb.IdentitytoolkitBetaOAuthIdpConfigResponseType) *beta.OAuthIdpConfigResponseType {
	if p == nil {
		return nil
	}
	obj := &beta.OAuthIdpConfigResponseType{
		IdToken: dcl.Bool(p.IdToken),
		Code:    dcl.Bool(p.Code),
		Token:   dcl.Bool(p.Token),
	}
	return obj
}

// ProtoToOAuthIdpConfig converts a OAuthIdpConfig resource from its proto representation.
func ProtoToOAuthIdpConfig(p *betapb.IdentitytoolkitBetaOAuthIdpConfig) *beta.OAuthIdpConfig {
	obj := &beta.OAuthIdpConfig{
		Name:         dcl.StringOrNil(p.Name),
		ClientId:     dcl.StringOrNil(p.ClientId),
		Issuer:       dcl.StringOrNil(p.Issuer),
		DisplayName:  dcl.StringOrNil(p.DisplayName),
		Enabled:      dcl.Bool(p.Enabled),
		ClientSecret: dcl.StringOrNil(p.ClientSecret),
		ResponseType: ProtoToIdentitytoolkitBetaOAuthIdpConfigResponseType(p.GetResponseType()),
		Project:      dcl.StringOrNil(p.Project),
	}
	return obj
}

// OAuthIdpConfigResponseTypeToProto converts a OAuthIdpConfigResponseType resource to its proto representation.
func IdentitytoolkitBetaOAuthIdpConfigResponseTypeToProto(o *beta.OAuthIdpConfigResponseType) *betapb.IdentitytoolkitBetaOAuthIdpConfigResponseType {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaOAuthIdpConfigResponseType{
		IdToken: dcl.ValueOrEmptyBool(o.IdToken),
		Code:    dcl.ValueOrEmptyBool(o.Code),
		Token:   dcl.ValueOrEmptyBool(o.Token),
	}
	return p
}

// OAuthIdpConfigToProto converts a OAuthIdpConfig resource to its proto representation.
func OAuthIdpConfigToProto(resource *beta.OAuthIdpConfig) *betapb.IdentitytoolkitBetaOAuthIdpConfig {
	p := &betapb.IdentitytoolkitBetaOAuthIdpConfig{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		ClientId:     dcl.ValueOrEmptyString(resource.ClientId),
		Issuer:       dcl.ValueOrEmptyString(resource.Issuer),
		DisplayName:  dcl.ValueOrEmptyString(resource.DisplayName),
		Enabled:      dcl.ValueOrEmptyBool(resource.Enabled),
		ClientSecret: dcl.ValueOrEmptyString(resource.ClientSecret),
		ResponseType: IdentitytoolkitBetaOAuthIdpConfigResponseTypeToProto(resource.ResponseType),
		Project:      dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Apply() method.
func (s *OAuthIdpConfigServer) applyOAuthIdpConfig(ctx context.Context, c *beta.Client, request *betapb.ApplyIdentitytoolkitBetaOAuthIdpConfigRequest) (*betapb.IdentitytoolkitBetaOAuthIdpConfig, error) {
	p := ProtoToOAuthIdpConfig(request.GetResource())
	res, err := c.ApplyOAuthIdpConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OAuthIdpConfigToProto(res)
	return r, nil
}

// ApplyOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Apply() method.
func (s *OAuthIdpConfigServer) ApplyIdentitytoolkitBetaOAuthIdpConfig(ctx context.Context, request *betapb.ApplyIdentitytoolkitBetaOAuthIdpConfigRequest) (*betapb.IdentitytoolkitBetaOAuthIdpConfig, error) {
	cl, err := createConfigOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyOAuthIdpConfig(ctx, cl, request)
}

// DeleteOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Delete() method.
func (s *OAuthIdpConfigServer) DeleteIdentitytoolkitBetaOAuthIdpConfig(ctx context.Context, request *betapb.DeleteIdentitytoolkitBetaOAuthIdpConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOAuthIdpConfig(ctx, ProtoToOAuthIdpConfig(request.GetResource()))

}

// ListIdentitytoolkitBetaOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfigList() method.
func (s *OAuthIdpConfigServer) ListIdentitytoolkitBetaOAuthIdpConfig(ctx context.Context, request *betapb.ListIdentitytoolkitBetaOAuthIdpConfigRequest) (*betapb.ListIdentitytoolkitBetaOAuthIdpConfigResponse, error) {
	cl, err := createConfigOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOAuthIdpConfig(ctx, ProtoToOAuthIdpConfig(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IdentitytoolkitBetaOAuthIdpConfig
	for _, r := range resources.Items {
		rp := OAuthIdpConfigToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListIdentitytoolkitBetaOAuthIdpConfigResponse{Items: protos}, nil
}

func createConfigOAuthIdpConfig(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
