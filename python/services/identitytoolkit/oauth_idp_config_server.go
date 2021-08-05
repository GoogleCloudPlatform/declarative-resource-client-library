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
	identitytoolkitpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/identitytoolkit_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
)

// Server implements the gRPC interface for OAuthIdpConfig.
type OAuthIdpConfigServer struct{}

// ProtoToOAuthIdpConfigResponseType converts a OAuthIdpConfigResponseType resource from its proto representation.
func ProtoToIdentitytoolkitOAuthIdpConfigResponseType(p *identitytoolkitpb.IdentitytoolkitOAuthIdpConfigResponseType) *identitytoolkit.OAuthIdpConfigResponseType {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.OAuthIdpConfigResponseType{
		IdToken: dcl.Bool(p.IdToken),
		Code:    dcl.Bool(p.Code),
		Token:   dcl.Bool(p.Token),
	}
	return obj
}

// ProtoToOAuthIdpConfig converts a OAuthIdpConfig resource from its proto representation.
func ProtoToOAuthIdpConfig(p *identitytoolkitpb.IdentitytoolkitOAuthIdpConfig) *identitytoolkit.OAuthIdpConfig {
	obj := &identitytoolkit.OAuthIdpConfig{
		Name:         dcl.StringOrNil(p.Name),
		ClientId:     dcl.StringOrNil(p.ClientId),
		Issuer:       dcl.StringOrNil(p.Issuer),
		DisplayName:  dcl.StringOrNil(p.DisplayName),
		Enabled:      dcl.Bool(p.Enabled),
		ClientSecret: dcl.StringOrNil(p.ClientSecret),
		ResponseType: ProtoToIdentitytoolkitOAuthIdpConfigResponseType(p.GetResponseType()),
		Project:      dcl.StringOrNil(p.Project),
	}
	return obj
}

// OAuthIdpConfigResponseTypeToProto converts a OAuthIdpConfigResponseType resource to its proto representation.
func IdentitytoolkitOAuthIdpConfigResponseTypeToProto(o *identitytoolkit.OAuthIdpConfigResponseType) *identitytoolkitpb.IdentitytoolkitOAuthIdpConfigResponseType {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitOAuthIdpConfigResponseType{
		IdToken: dcl.ValueOrEmptyBool(o.IdToken),
		Code:    dcl.ValueOrEmptyBool(o.Code),
		Token:   dcl.ValueOrEmptyBool(o.Token),
	}
	return p
}

// OAuthIdpConfigToProto converts a OAuthIdpConfig resource to its proto representation.
func OAuthIdpConfigToProto(resource *identitytoolkit.OAuthIdpConfig) *identitytoolkitpb.IdentitytoolkitOAuthIdpConfig {
	p := &identitytoolkitpb.IdentitytoolkitOAuthIdpConfig{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		ClientId:     dcl.ValueOrEmptyString(resource.ClientId),
		Issuer:       dcl.ValueOrEmptyString(resource.Issuer),
		DisplayName:  dcl.ValueOrEmptyString(resource.DisplayName),
		Enabled:      dcl.ValueOrEmptyBool(resource.Enabled),
		ClientSecret: dcl.ValueOrEmptyString(resource.ClientSecret),
		ResponseType: IdentitytoolkitOAuthIdpConfigResponseTypeToProto(resource.ResponseType),
		Project:      dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Apply() method.
func (s *OAuthIdpConfigServer) applyOAuthIdpConfig(ctx context.Context, c *identitytoolkit.Client, request *identitytoolkitpb.ApplyIdentitytoolkitOAuthIdpConfigRequest) (*identitytoolkitpb.IdentitytoolkitOAuthIdpConfig, error) {
	p := ProtoToOAuthIdpConfig(request.GetResource())
	res, err := c.ApplyOAuthIdpConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OAuthIdpConfigToProto(res)
	return r, nil
}

// ApplyOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Apply() method.
func (s *OAuthIdpConfigServer) ApplyIdentitytoolkitOAuthIdpConfig(ctx context.Context, request *identitytoolkitpb.ApplyIdentitytoolkitOAuthIdpConfigRequest) (*identitytoolkitpb.IdentitytoolkitOAuthIdpConfig, error) {
	cl, err := createConfigOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyOAuthIdpConfig(ctx, cl, request)
}

// DeleteOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Delete() method.
func (s *OAuthIdpConfigServer) DeleteIdentitytoolkitOAuthIdpConfig(ctx context.Context, request *identitytoolkitpb.DeleteIdentitytoolkitOAuthIdpConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOAuthIdpConfig(ctx, ProtoToOAuthIdpConfig(request.GetResource()))

}

// ListIdentitytoolkitOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfigList() method.
func (s *OAuthIdpConfigServer) ListIdentitytoolkitOAuthIdpConfig(ctx context.Context, request *identitytoolkitpb.ListIdentitytoolkitOAuthIdpConfigRequest) (*identitytoolkitpb.ListIdentitytoolkitOAuthIdpConfigResponse, error) {
	cl, err := createConfigOAuthIdpConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOAuthIdpConfig(ctx, ProtoToOAuthIdpConfig(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*identitytoolkitpb.IdentitytoolkitOAuthIdpConfig
	for _, r := range resources.Items {
		rp := OAuthIdpConfigToProto(r)
		protos = append(protos, rp)
	}
	return &identitytoolkitpb.ListIdentitytoolkitOAuthIdpConfigResponse{Items: protos}, nil
}

func createConfigOAuthIdpConfig(ctx context.Context, service_account_file string) (*identitytoolkit.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return identitytoolkit.NewClient(conf), nil
}
