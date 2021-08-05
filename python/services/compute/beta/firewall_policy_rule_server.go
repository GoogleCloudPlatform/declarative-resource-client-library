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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for FirewallPolicyRule.
type FirewallPolicyRuleServer struct{}

// ProtoToFirewallPolicyRuleDirectionEnum converts a FirewallPolicyRuleDirectionEnum enum from its proto representation.
func ProtoToComputeBetaFirewallPolicyRuleDirectionEnum(e betapb.ComputeBetaFirewallPolicyRuleDirectionEnum) *beta.FirewallPolicyRuleDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaFirewallPolicyRuleDirectionEnum_name[int32(e)]; ok {
		e := beta.FirewallPolicyRuleDirectionEnum(n[len("ComputeBetaFirewallPolicyRuleDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFirewallPolicyRuleMatch converts a FirewallPolicyRuleMatch resource from its proto representation.
func ProtoToComputeBetaFirewallPolicyRuleMatch(p *betapb.ComputeBetaFirewallPolicyRuleMatch) *beta.FirewallPolicyRuleMatch {
	if p == nil {
		return nil
	}
	obj := &beta.FirewallPolicyRuleMatch{}
	for _, r := range p.GetSrcIpRanges() {
		obj.SrcIPRanges = append(obj.SrcIPRanges, r)
	}
	for _, r := range p.GetDestIpRanges() {
		obj.DestIPRanges = append(obj.DestIPRanges, r)
	}
	for _, r := range p.GetLayer4Configs() {
		obj.Layer4Configs = append(obj.Layer4Configs, *ProtoToComputeBetaFirewallPolicyRuleMatchLayer4Configs(r))
	}
	return obj
}

// ProtoToFirewallPolicyRuleMatchLayer4Configs converts a FirewallPolicyRuleMatchLayer4Configs resource from its proto representation.
func ProtoToComputeBetaFirewallPolicyRuleMatchLayer4Configs(p *betapb.ComputeBetaFirewallPolicyRuleMatchLayer4Configs) *beta.FirewallPolicyRuleMatchLayer4Configs {
	if p == nil {
		return nil
	}
	obj := &beta.FirewallPolicyRuleMatchLayer4Configs{
		IPProtocol: dcl.StringOrNil(p.IpProtocol),
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToFirewallPolicyRule converts a FirewallPolicyRule resource from its proto representation.
func ProtoToFirewallPolicyRule(p *betapb.ComputeBetaFirewallPolicyRule) *beta.FirewallPolicyRule {
	obj := &beta.FirewallPolicyRule{
		Description:    dcl.StringOrNil(p.Description),
		Priority:       dcl.Int64OrNil(p.Priority),
		Match:          ProtoToComputeBetaFirewallPolicyRuleMatch(p.GetMatch()),
		Action:         dcl.StringOrNil(p.Action),
		Direction:      ProtoToComputeBetaFirewallPolicyRuleDirectionEnum(p.GetDirection()),
		EnableLogging:  dcl.Bool(p.EnableLogging),
		RuleTupleCount: dcl.Int64OrNil(p.RuleTupleCount),
		Disabled:       dcl.Bool(p.Disabled),
		Kind:           dcl.StringOrNil(p.Kind),
		FirewallPolicy: dcl.StringOrNil(p.FirewallPolicy),
	}
	for _, r := range p.GetTargetResources() {
		obj.TargetResources = append(obj.TargetResources, r)
	}
	for _, r := range p.GetTargetServiceAccounts() {
		obj.TargetServiceAccounts = append(obj.TargetServiceAccounts, r)
	}
	return obj
}

// FirewallPolicyRuleDirectionEnumToProto converts a FirewallPolicyRuleDirectionEnum enum to its proto representation.
func ComputeBetaFirewallPolicyRuleDirectionEnumToProto(e *beta.FirewallPolicyRuleDirectionEnum) betapb.ComputeBetaFirewallPolicyRuleDirectionEnum {
	if e == nil {
		return betapb.ComputeBetaFirewallPolicyRuleDirectionEnum(0)
	}
	if v, ok := betapb.ComputeBetaFirewallPolicyRuleDirectionEnum_value["FirewallPolicyRuleDirectionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaFirewallPolicyRuleDirectionEnum(v)
	}
	return betapb.ComputeBetaFirewallPolicyRuleDirectionEnum(0)
}

// FirewallPolicyRuleMatchToProto converts a FirewallPolicyRuleMatch resource to its proto representation.
func ComputeBetaFirewallPolicyRuleMatchToProto(o *beta.FirewallPolicyRuleMatch) *betapb.ComputeBetaFirewallPolicyRuleMatch {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaFirewallPolicyRuleMatch{}
	for _, r := range o.SrcIPRanges {
		p.SrcIpRanges = append(p.SrcIpRanges, r)
	}
	for _, r := range o.DestIPRanges {
		p.DestIpRanges = append(p.DestIpRanges, r)
	}
	for _, r := range o.Layer4Configs {
		p.Layer4Configs = append(p.Layer4Configs, ComputeBetaFirewallPolicyRuleMatchLayer4ConfigsToProto(&r))
	}
	return p
}

// FirewallPolicyRuleMatchLayer4ConfigsToProto converts a FirewallPolicyRuleMatchLayer4Configs resource to its proto representation.
func ComputeBetaFirewallPolicyRuleMatchLayer4ConfigsToProto(o *beta.FirewallPolicyRuleMatchLayer4Configs) *betapb.ComputeBetaFirewallPolicyRuleMatchLayer4Configs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaFirewallPolicyRuleMatchLayer4Configs{
		IpProtocol: dcl.ValueOrEmptyString(o.IPProtocol),
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// FirewallPolicyRuleToProto converts a FirewallPolicyRule resource to its proto representation.
func FirewallPolicyRuleToProto(resource *beta.FirewallPolicyRule) *betapb.ComputeBetaFirewallPolicyRule {
	p := &betapb.ComputeBetaFirewallPolicyRule{
		Description:    dcl.ValueOrEmptyString(resource.Description),
		Priority:       dcl.ValueOrEmptyInt64(resource.Priority),
		Match:          ComputeBetaFirewallPolicyRuleMatchToProto(resource.Match),
		Action:         dcl.ValueOrEmptyString(resource.Action),
		Direction:      ComputeBetaFirewallPolicyRuleDirectionEnumToProto(resource.Direction),
		EnableLogging:  dcl.ValueOrEmptyBool(resource.EnableLogging),
		RuleTupleCount: dcl.ValueOrEmptyInt64(resource.RuleTupleCount),
		Disabled:       dcl.ValueOrEmptyBool(resource.Disabled),
		Kind:           dcl.ValueOrEmptyString(resource.Kind),
		FirewallPolicy: dcl.ValueOrEmptyString(resource.FirewallPolicy),
	}
	for _, r := range resource.TargetResources {
		p.TargetResources = append(p.TargetResources, r)
	}
	for _, r := range resource.TargetServiceAccounts {
		p.TargetServiceAccounts = append(p.TargetServiceAccounts, r)
	}

	return p
}

// ApplyFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Apply() method.
func (s *FirewallPolicyRuleServer) applyFirewallPolicyRule(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaFirewallPolicyRuleRequest) (*betapb.ComputeBetaFirewallPolicyRule, error) {
	p := ProtoToFirewallPolicyRule(request.GetResource())
	res, err := c.ApplyFirewallPolicyRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyRuleToProto(res)
	return r, nil
}

// ApplyFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Apply() method.
func (s *FirewallPolicyRuleServer) ApplyComputeBetaFirewallPolicyRule(ctx context.Context, request *betapb.ApplyComputeBetaFirewallPolicyRuleRequest) (*betapb.ComputeBetaFirewallPolicyRule, error) {
	cl, err := createConfigFirewallPolicyRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicyRule(ctx, cl, request)
}

// DeleteFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Delete() method.
func (s *FirewallPolicyRuleServer) DeleteComputeBetaFirewallPolicyRule(ctx context.Context, request *betapb.DeleteComputeBetaFirewallPolicyRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicyRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicyRule(ctx, ProtoToFirewallPolicyRule(request.GetResource()))

}

// ListComputeBetaFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRuleList() method.
func (s *FirewallPolicyRuleServer) ListComputeBetaFirewallPolicyRule(ctx context.Context, request *betapb.ListComputeBetaFirewallPolicyRuleRequest) (*betapb.ListComputeBetaFirewallPolicyRuleResponse, error) {
	cl, err := createConfigFirewallPolicyRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicyRule(ctx, ProtoToFirewallPolicyRule(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaFirewallPolicyRule
	for _, r := range resources.Items {
		rp := FirewallPolicyRuleToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaFirewallPolicyRuleResponse{Items: protos}, nil
}

func createConfigFirewallPolicyRule(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
