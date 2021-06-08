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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/beta/networksecurity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
)

// Server implements the gRPC interface for AuthorizationPolicy.
type AuthorizationPolicyServer struct{}

// ProtoToAuthorizationPolicyActionEnum converts a AuthorizationPolicyActionEnum enum from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyActionEnum(e betapb.NetworksecurityBetaAuthorizationPolicyActionEnum) *beta.AuthorizationPolicyActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworksecurityBetaAuthorizationPolicyActionEnum_name[int32(e)]; ok {
		e := beta.AuthorizationPolicyActionEnum(n[len("NetworksecurityBetaAuthorizationPolicyActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAuthorizationPolicyRules converts a AuthorizationPolicyRules resource from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRules(p *betapb.NetworksecurityBetaAuthorizationPolicyRules) *beta.AuthorizationPolicyRules {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRules{}
	for _, r := range p.GetSources() {
		obj.Sources = append(obj.Sources, *ProtoToNetworksecurityBetaAuthorizationPolicyRulesSources(r))
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinations(r))
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesSources converts a AuthorizationPolicyRulesSources resource from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRulesSources(p *betapb.NetworksecurityBetaAuthorizationPolicyRulesSources) *beta.AuthorizationPolicyRulesSources {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRulesSources{}
	for _, r := range p.GetPrincipals() {
		obj.Principals = append(obj.Principals, r)
	}
	for _, r := range p.GetIpBlocks() {
		obj.IPBlocks = append(obj.IPBlocks, r)
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesDestinations converts a AuthorizationPolicyRulesDestinations resource from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinations(p *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations) *beta.AuthorizationPolicyRulesDestinations {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRulesDestinations{
		HttpHeaderMatch: ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p.GetHttpHeaderMatch()),
	}
	for _, r := range p.GetHosts() {
		obj.Hosts = append(obj.Hosts, r)
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetPaths() {
		obj.Paths = append(obj.Paths, r)
	}
	for _, r := range p.GetMethods() {
		obj.Methods = append(obj.Methods, r)
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesDestinationsHttpHeaderMatch converts a AuthorizationPolicyRulesDestinationsHttpHeaderMatch resource from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch) *beta.AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRulesDestinationsHttpHeaderMatch{
		HeaderName: dcl.StringOrNil(p.HeaderName),
		RegexMatch: dcl.StringOrNil(p.RegexMatch),
	}
	return obj
}

// ProtoToAuthorizationPolicy converts a AuthorizationPolicy resource from its proto representation.
func ProtoToAuthorizationPolicy(p *betapb.NetworksecurityBetaAuthorizationPolicy) *beta.AuthorizationPolicy {
	obj := &beta.AuthorizationPolicy{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Action:      ProtoToNetworksecurityBetaAuthorizationPolicyActionEnum(p.GetAction()),
		Project:     dcl.StringOrNil(p.Project),
		Location:    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworksecurityBetaAuthorizationPolicyRules(r))
	}
	return obj
}

// AuthorizationPolicyActionEnumToProto converts a AuthorizationPolicyActionEnum enum to its proto representation.
func NetworksecurityBetaAuthorizationPolicyActionEnumToProto(e *beta.AuthorizationPolicyActionEnum) betapb.NetworksecurityBetaAuthorizationPolicyActionEnum {
	if e == nil {
		return betapb.NetworksecurityBetaAuthorizationPolicyActionEnum(0)
	}
	if v, ok := betapb.NetworksecurityBetaAuthorizationPolicyActionEnum_value["AuthorizationPolicyActionEnum"+string(*e)]; ok {
		return betapb.NetworksecurityBetaAuthorizationPolicyActionEnum(v)
	}
	return betapb.NetworksecurityBetaAuthorizationPolicyActionEnum(0)
}

// AuthorizationPolicyRulesToProto converts a AuthorizationPolicyRules resource to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesToProto(o *beta.AuthorizationPolicyRules) *betapb.NetworksecurityBetaAuthorizationPolicyRules {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRules{}
	for _, r := range o.Sources {
		p.Sources = append(p.Sources, NetworksecurityBetaAuthorizationPolicyRulesSourcesToProto(&r))
	}
	for _, r := range o.Destinations {
		p.Destinations = append(p.Destinations, NetworksecurityBetaAuthorizationPolicyRulesDestinationsToProto(&r))
	}
	return p
}

// AuthorizationPolicyRulesSourcesToProto converts a AuthorizationPolicyRulesSources resource to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesSourcesToProto(o *beta.AuthorizationPolicyRulesSources) *betapb.NetworksecurityBetaAuthorizationPolicyRulesSources {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRulesSources{}
	for _, r := range o.Principals {
		p.Principals = append(p.Principals, r)
	}
	for _, r := range o.IPBlocks {
		p.IpBlocks = append(p.IpBlocks, r)
	}
	return p
}

// AuthorizationPolicyRulesDestinationsToProto converts a AuthorizationPolicyRulesDestinations resource to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesDestinationsToProto(o *beta.AuthorizationPolicyRulesDestinations) *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations{
		HttpHeaderMatch: NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o.HttpHeaderMatch),
	}
	for _, r := range o.Hosts {
		p.Hosts = append(p.Hosts, r)
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	for _, r := range o.Paths {
		p.Paths = append(p.Paths, r)
	}
	for _, r := range o.Methods {
		p.Methods = append(p.Methods, r)
	}
	return p
}

// AuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto converts a AuthorizationPolicyRulesDestinationsHttpHeaderMatch resource to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o *beta.AuthorizationPolicyRulesDestinationsHttpHeaderMatch) *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch{
		HeaderName: dcl.ValueOrEmptyString(o.HeaderName),
		RegexMatch: dcl.ValueOrEmptyString(o.RegexMatch),
	}
	return p
}

// AuthorizationPolicyToProto converts a AuthorizationPolicy resource to its proto representation.
func AuthorizationPolicyToProto(resource *beta.AuthorizationPolicy) *betapb.NetworksecurityBetaAuthorizationPolicy {
	p := &betapb.NetworksecurityBetaAuthorizationPolicy{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Action:      NetworksecurityBetaAuthorizationPolicyActionEnumToProto(resource.Action),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Location:    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Rules {
		p.Rules = append(p.Rules, NetworksecurityBetaAuthorizationPolicyRulesToProto(&r))
	}

	return p
}

// ApplyAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) applyAuthorizationPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworksecurityBetaAuthorizationPolicyRequest) (*betapb.NetworksecurityBetaAuthorizationPolicy, error) {
	p := ProtoToAuthorizationPolicy(request.GetResource())
	res, err := c.ApplyAuthorizationPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AuthorizationPolicyToProto(res)
	return r, nil
}

// ApplyAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) ApplyNetworksecurityBetaAuthorizationPolicy(ctx context.Context, request *betapb.ApplyNetworksecurityBetaAuthorizationPolicyRequest) (*betapb.NetworksecurityBetaAuthorizationPolicy, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAuthorizationPolicy(ctx, cl, request)
}

// DeleteAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Delete() method.
func (s *AuthorizationPolicyServer) DeleteNetworksecurityBetaAuthorizationPolicy(ctx context.Context, request *betapb.DeleteNetworksecurityBetaAuthorizationPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAuthorizationPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAuthorizationPolicy(ctx, ProtoToAuthorizationPolicy(request.GetResource()))

}

// ListAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicyList() method.
func (s *AuthorizationPolicyServer) ListNetworksecurityBetaAuthorizationPolicy(ctx context.Context, request *betapb.ListNetworksecurityBetaAuthorizationPolicyRequest) (*betapb.ListNetworksecurityBetaAuthorizationPolicyResponse, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAuthorizationPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworksecurityBetaAuthorizationPolicy
	for _, r := range resources.Items {
		rp := AuthorizationPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListNetworksecurityBetaAuthorizationPolicyResponse{Items: protos}, nil
}

func createConfigAuthorizationPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
