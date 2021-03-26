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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/alpha/networksecurity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
)

// Server implements the gRPC interface for AuthorizationPolicy.
type AuthorizationPolicyServer struct{}

// ProtoToAuthorizationPolicyActionEnum converts a AuthorizationPolicyActionEnum enum from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyActionEnum(e alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum) *alpha.AuthorizationPolicyActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum_name[int32(e)]; ok {
		e := alpha.AuthorizationPolicyActionEnum(n[len("AuthorizationPolicyActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAuthorizationPolicyRules converts a AuthorizationPolicyRules resource from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyRules(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRules) *alpha.AuthorizationPolicyRules {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRules{}
	for _, r := range p.GetSources() {
		obj.Sources = append(obj.Sources, *ProtoToNetworksecurityAlphaAuthorizationPolicyRulesSources(r))
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinations(r))
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesSources converts a AuthorizationPolicyRulesSources resource from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyRulesSources(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources) *alpha.AuthorizationPolicyRulesSources {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRulesSources{}
	for _, r := range p.GetPrincipals() {
		obj.Principals = append(obj.Principals, r)
	}
	for _, r := range p.GetIpBlocks() {
		obj.IPBlocks = append(obj.IPBlocks, r)
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesDestinations converts a AuthorizationPolicyRulesDestinations resource from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinations(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations) *alpha.AuthorizationPolicyRulesDestinations {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRulesDestinations{
		HttpHeaderMatch: ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p.GetHttpHeaderMatch()),
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
func ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch) *alpha.AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRulesDestinationsHttpHeaderMatch{
		HeaderName: dcl.StringOrNil(p.HeaderName),
		RegexMatch: dcl.StringOrNil(p.RegexMatch),
	}
	return obj
}

// ProtoToAuthorizationPolicy converts a AuthorizationPolicy resource from its proto representation.
func ProtoToAuthorizationPolicy(p *alphapb.NetworksecurityAlphaAuthorizationPolicy) *alpha.AuthorizationPolicy {
	obj := &alpha.AuthorizationPolicy{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Action:      ProtoToNetworksecurityAlphaAuthorizationPolicyActionEnum(p.GetAction()),
		Project:     dcl.StringOrNil(p.Project),
		Location:    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworksecurityAlphaAuthorizationPolicyRules(r))
	}
	return obj
}

// AuthorizationPolicyActionEnumToProto converts a AuthorizationPolicyActionEnum enum to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyActionEnumToProto(e *alpha.AuthorizationPolicyActionEnum) alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum {
	if e == nil {
		return alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum(0)
	}
	if v, ok := alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum_value["AuthorizationPolicyActionEnum"+string(*e)]; ok {
		return alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum(v)
	}
	return alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum(0)
}

// AuthorizationPolicyRulesToProto converts a AuthorizationPolicyRules resource to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyRulesToProto(o *alpha.AuthorizationPolicyRules) *alphapb.NetworksecurityAlphaAuthorizationPolicyRules {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRules{}
	for _, r := range o.Sources {
		p.Sources = append(p.Sources, NetworksecurityAlphaAuthorizationPolicyRulesSourcesToProto(&r))
	}
	for _, r := range o.Destinations {
		p.Destinations = append(p.Destinations, NetworksecurityAlphaAuthorizationPolicyRulesDestinationsToProto(&r))
	}
	return p
}

// AuthorizationPolicyRulesSourcesToProto converts a AuthorizationPolicyRulesSources resource to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyRulesSourcesToProto(o *alpha.AuthorizationPolicyRulesSources) *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources{}
	for _, r := range o.Principals {
		p.Principals = append(p.Principals, r)
	}
	for _, r := range o.IPBlocks {
		p.IpBlocks = append(p.IpBlocks, r)
	}
	return p
}

// AuthorizationPolicyRulesDestinationsToProto converts a AuthorizationPolicyRulesDestinations resource to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyRulesDestinationsToProto(o *alpha.AuthorizationPolicyRulesDestinations) *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations{
		HttpHeaderMatch: NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o.HttpHeaderMatch),
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
func NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o *alpha.AuthorizationPolicyRulesDestinationsHttpHeaderMatch) *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch{
		HeaderName: dcl.ValueOrEmptyString(o.HeaderName),
		RegexMatch: dcl.ValueOrEmptyString(o.RegexMatch),
	}
	return p
}

// AuthorizationPolicyToProto converts a AuthorizationPolicy resource to its proto representation.
func AuthorizationPolicyToProto(resource *alpha.AuthorizationPolicy) *alphapb.NetworksecurityAlphaAuthorizationPolicy {
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicy{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Action:      NetworksecurityAlphaAuthorizationPolicyActionEnumToProto(resource.Action),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Location:    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Rules {
		p.Rules = append(p.Rules, NetworksecurityAlphaAuthorizationPolicyRulesToProto(&r))
	}

	return p
}

// ApplyAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) applyAuthorizationPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworksecurityAlphaAuthorizationPolicyRequest) (*alphapb.NetworksecurityAlphaAuthorizationPolicy, error) {
	p := ProtoToAuthorizationPolicy(request.GetResource())
	res, err := c.ApplyAuthorizationPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AuthorizationPolicyToProto(res)
	return r, nil
}

// ApplyAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) ApplyNetworksecurityAlphaAuthorizationPolicy(ctx context.Context, request *alphapb.ApplyNetworksecurityAlphaAuthorizationPolicyRequest) (*alphapb.NetworksecurityAlphaAuthorizationPolicy, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAuthorizationPolicy(ctx, cl, request)
}

// DeleteAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Delete() method.
func (s *AuthorizationPolicyServer) DeleteNetworksecurityAlphaAuthorizationPolicy(ctx context.Context, request *alphapb.DeleteNetworksecurityAlphaAuthorizationPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAuthorizationPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAuthorizationPolicy(ctx, ProtoToAuthorizationPolicy(request.GetResource()))

}

// ListAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicyList() method.
func (s *AuthorizationPolicyServer) ListNetworksecurityAlphaAuthorizationPolicy(ctx context.Context, request *alphapb.ListNetworksecurityAlphaAuthorizationPolicyRequest) (*alphapb.ListNetworksecurityAlphaAuthorizationPolicyResponse, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAuthorizationPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworksecurityAlphaAuthorizationPolicy
	for _, r := range resources.Items {
		rp := AuthorizationPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListNetworksecurityAlphaAuthorizationPolicyResponse{Items: protos}, nil
}

func createConfigAuthorizationPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
