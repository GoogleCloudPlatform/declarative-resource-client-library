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

// Server implements the gRPC interface for Tenant.
type TenantServer struct{}

// ProtoToTenantMfaConfigStateEnum converts a TenantMfaConfigStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaTenantMfaConfigStateEnum(e betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum) *beta.TenantMfaConfigStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum_name[int32(e)]; ok {
		e := beta.TenantMfaConfigStateEnum(n[len("IdentitytoolkitBetaTenantMfaConfigStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToTenantMfaConfigEnabledProvidersEnum converts a TenantMfaConfigEnabledProvidersEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum(e betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum) *beta.TenantMfaConfigEnabledProvidersEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum_name[int32(e)]; ok {
		e := beta.TenantMfaConfigEnabledProvidersEnum(n[len("IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum"):])
		return &e
	}
	return nil
}

// ProtoToTenantMfaConfig converts a TenantMfaConfig resource from its proto representation.
func ProtoToIdentitytoolkitBetaTenantMfaConfig(p *betapb.IdentitytoolkitBetaTenantMfaConfig) *beta.TenantMfaConfig {
	if p == nil {
		return nil
	}
	obj := &beta.TenantMfaConfig{
		State: ProtoToIdentitytoolkitBetaTenantMfaConfigStateEnum(p.GetState()),
	}
	for _, r := range p.GetEnabledProviders() {
		obj.EnabledProviders = append(obj.EnabledProviders, *ProtoToIdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum(r))
	}
	return obj
}

// ProtoToTenant converts a Tenant resource from its proto representation.
func ProtoToTenant(p *betapb.IdentitytoolkitBetaTenant) *beta.Tenant {
	obj := &beta.Tenant{
		Name:                  dcl.StringOrNil(p.Name),
		DisplayName:           dcl.StringOrNil(p.DisplayName),
		AllowPasswordSignup:   dcl.Bool(p.AllowPasswordSignup),
		EnableEmailLinkSignin: dcl.Bool(p.EnableEmailLinkSignin),
		DisableAuth:           dcl.Bool(p.DisableAuth),
		EnableAnonymousUser:   dcl.Bool(p.EnableAnonymousUser),
		MfaConfig:             ProtoToIdentitytoolkitBetaTenantMfaConfig(p.GetMfaConfig()),
		Project:               dcl.StringOrNil(p.Project),
	}
	return obj
}

// TenantMfaConfigStateEnumToProto converts a TenantMfaConfigStateEnum enum to its proto representation.
func IdentitytoolkitBetaTenantMfaConfigStateEnumToProto(e *beta.TenantMfaConfigStateEnum) betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum_value["TenantMfaConfigStateEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum(v)
	}
	return betapb.IdentitytoolkitBetaTenantMfaConfigStateEnum(0)
}

// TenantMfaConfigEnabledProvidersEnumToProto converts a TenantMfaConfigEnabledProvidersEnum enum to its proto representation.
func IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnumToProto(e *beta.TenantMfaConfigEnabledProvidersEnum) betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum_value["TenantMfaConfigEnabledProvidersEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum(v)
	}
	return betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum(0)
}

// TenantMfaConfigToProto converts a TenantMfaConfig resource to its proto representation.
func IdentitytoolkitBetaTenantMfaConfigToProto(o *beta.TenantMfaConfig) *betapb.IdentitytoolkitBetaTenantMfaConfig {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaTenantMfaConfig{
		State: IdentitytoolkitBetaTenantMfaConfigStateEnumToProto(o.State),
	}
	for _, r := range o.EnabledProviders {
		p.EnabledProviders = append(p.EnabledProviders, betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum(betapb.IdentitytoolkitBetaTenantMfaConfigEnabledProvidersEnum_value[string(r)]))
	}
	return p
}

// TenantToProto converts a Tenant resource to its proto representation.
func TenantToProto(resource *beta.Tenant) *betapb.IdentitytoolkitBetaTenant {
	p := &betapb.IdentitytoolkitBetaTenant{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		DisplayName:           dcl.ValueOrEmptyString(resource.DisplayName),
		AllowPasswordSignup:   dcl.ValueOrEmptyBool(resource.AllowPasswordSignup),
		EnableEmailLinkSignin: dcl.ValueOrEmptyBool(resource.EnableEmailLinkSignin),
		DisableAuth:           dcl.ValueOrEmptyBool(resource.DisableAuth),
		EnableAnonymousUser:   dcl.ValueOrEmptyBool(resource.EnableAnonymousUser),
		MfaConfig:             IdentitytoolkitBetaTenantMfaConfigToProto(resource.MfaConfig),
		Project:               dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyTenant handles the gRPC request by passing it to the underlying Tenant Apply() method.
func (s *TenantServer) applyTenant(ctx context.Context, c *beta.Client, request *betapb.ApplyIdentitytoolkitBetaTenantRequest) (*betapb.IdentitytoolkitBetaTenant, error) {
	p := ProtoToTenant(request.GetResource())
	res, err := c.ApplyTenant(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TenantToProto(res)
	return r, nil
}

// ApplyTenant handles the gRPC request by passing it to the underlying Tenant Apply() method.
func (s *TenantServer) ApplyIdentitytoolkitBetaTenant(ctx context.Context, request *betapb.ApplyIdentitytoolkitBetaTenantRequest) (*betapb.IdentitytoolkitBetaTenant, error) {
	cl, err := createConfigTenant(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTenant(ctx, cl, request)
}

// DeleteTenant handles the gRPC request by passing it to the underlying Tenant Delete() method.
func (s *TenantServer) DeleteIdentitytoolkitBetaTenant(ctx context.Context, request *betapb.DeleteIdentitytoolkitBetaTenantRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTenant(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTenant(ctx, ProtoToTenant(request.GetResource()))

}

// ListIdentitytoolkitBetaTenant handles the gRPC request by passing it to the underlying TenantList() method.
func (s *TenantServer) ListIdentitytoolkitBetaTenant(ctx context.Context, request *betapb.ListIdentitytoolkitBetaTenantRequest) (*betapb.ListIdentitytoolkitBetaTenantResponse, error) {
	cl, err := createConfigTenant(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTenant(ctx, ProtoToTenant(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IdentitytoolkitBetaTenant
	for _, r := range resources.Items {
		rp := TenantToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListIdentitytoolkitBetaTenantResponse{Items: protos}, nil
}

func createConfigTenant(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
