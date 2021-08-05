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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for FirewallPolicyRule.
type FirewallPolicyRuleServer struct{}

// ProtoToFirewallPolicyRuleDirectionEnum converts a FirewallPolicyRuleDirectionEnum enum from its proto representation.
func ProtoToComputeFirewallPolicyRuleDirectionEnum(e computepb.ComputeFirewallPolicyRuleDirectionEnum) *compute.FirewallPolicyRuleDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeFirewallPolicyRuleDirectionEnum_name[int32(e)]; ok {
		e := compute.FirewallPolicyRuleDirectionEnum(n[len("ComputeFirewallPolicyRuleDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFirewallPolicyRuleMatch converts a FirewallPolicyRuleMatch resource from its proto representation.
func ProtoToComputeFirewallPolicyRuleMatch(p *computepb.ComputeFirewallPolicyRuleMatch) *compute.FirewallPolicyRuleMatch {
	if p == nil {
		return nil
	}
	obj := &compute.FirewallPolicyRuleMatch{}
	for _, r := range p.GetSrcIpRanges() {
		obj.SrcIPRanges = append(obj.SrcIPRanges, r)
	}
	for _, r := range p.GetDestIpRanges() {
		obj.DestIPRanges = append(obj.DestIPRanges, r)
	}
	for _, r := range p.GetLayer4Configs() {
		obj.Layer4Configs = append(obj.Layer4Configs, *ProtoToComputeFirewallPolicyRuleMatchLayer4Configs(r))
	}
	return obj
}

// ProtoToFirewallPolicyRuleMatchLayer4Configs converts a FirewallPolicyRuleMatchLayer4Configs resource from its proto representation.
func ProtoToComputeFirewallPolicyRuleMatchLayer4Configs(p *computepb.ComputeFirewallPolicyRuleMatchLayer4Configs) *compute.FirewallPolicyRuleMatchLayer4Configs {
	if p == nil {
		return nil
	}
	obj := &compute.FirewallPolicyRuleMatchLayer4Configs{
		IPProtocol: dcl.StringOrNil(p.IpProtocol),
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToFirewallPolicyRule converts a FirewallPolicyRule resource from its proto representation.
func ProtoToFirewallPolicyRule(p *computepb.ComputeFirewallPolicyRule) *compute.FirewallPolicyRule {
	obj := &compute.FirewallPolicyRule{
		Description:    dcl.StringOrNil(p.Description),
		Priority:       dcl.Int64OrNil(p.Priority),
		Match:          ProtoToComputeFirewallPolicyRuleMatch(p.GetMatch()),
		Action:         dcl.StringOrNil(p.Action),
		Direction:      ProtoToComputeFirewallPolicyRuleDirectionEnum(p.GetDirection()),
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
func ComputeFirewallPolicyRuleDirectionEnumToProto(e *compute.FirewallPolicyRuleDirectionEnum) computepb.ComputeFirewallPolicyRuleDirectionEnum {
	if e == nil {
		return computepb.ComputeFirewallPolicyRuleDirectionEnum(0)
	}
	if v, ok := computepb.ComputeFirewallPolicyRuleDirectionEnum_value["FirewallPolicyRuleDirectionEnum"+string(*e)]; ok {
		return computepb.ComputeFirewallPolicyRuleDirectionEnum(v)
	}
	return computepb.ComputeFirewallPolicyRuleDirectionEnum(0)
}

// FirewallPolicyRuleMatchToProto converts a FirewallPolicyRuleMatch resource to its proto representation.
func ComputeFirewallPolicyRuleMatchToProto(o *compute.FirewallPolicyRuleMatch) *computepb.ComputeFirewallPolicyRuleMatch {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeFirewallPolicyRuleMatch{}
	for _, r := range o.SrcIPRanges {
		p.SrcIpRanges = append(p.SrcIpRanges, r)
	}
	for _, r := range o.DestIPRanges {
		p.DestIpRanges = append(p.DestIpRanges, r)
	}
	for _, r := range o.Layer4Configs {
		p.Layer4Configs = append(p.Layer4Configs, ComputeFirewallPolicyRuleMatchLayer4ConfigsToProto(&r))
	}
	return p
}

// FirewallPolicyRuleMatchLayer4ConfigsToProto converts a FirewallPolicyRuleMatchLayer4Configs resource to its proto representation.
func ComputeFirewallPolicyRuleMatchLayer4ConfigsToProto(o *compute.FirewallPolicyRuleMatchLayer4Configs) *computepb.ComputeFirewallPolicyRuleMatchLayer4Configs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeFirewallPolicyRuleMatchLayer4Configs{
		IpProtocol: dcl.ValueOrEmptyString(o.IPProtocol),
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// FirewallPolicyRuleToProto converts a FirewallPolicyRule resource to its proto representation.
func FirewallPolicyRuleToProto(resource *compute.FirewallPolicyRule) *computepb.ComputeFirewallPolicyRule {
	p := &computepb.ComputeFirewallPolicyRule{
		Description:    dcl.ValueOrEmptyString(resource.Description),
		Priority:       dcl.ValueOrEmptyInt64(resource.Priority),
		Match:          ComputeFirewallPolicyRuleMatchToProto(resource.Match),
		Action:         dcl.ValueOrEmptyString(resource.Action),
		Direction:      ComputeFirewallPolicyRuleDirectionEnumToProto(resource.Direction),
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
func (s *FirewallPolicyRuleServer) applyFirewallPolicyRule(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeFirewallPolicyRuleRequest) (*computepb.ComputeFirewallPolicyRule, error) {
	p := ProtoToFirewallPolicyRule(request.GetResource())
	res, err := c.ApplyFirewallPolicyRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyRuleToProto(res)
	return r, nil
}

// ApplyFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Apply() method.
func (s *FirewallPolicyRuleServer) ApplyComputeFirewallPolicyRule(ctx context.Context, request *computepb.ApplyComputeFirewallPolicyRuleRequest) (*computepb.ComputeFirewallPolicyRule, error) {
	cl, err := createConfigFirewallPolicyRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicyRule(ctx, cl, request)
}

// DeleteFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Delete() method.
func (s *FirewallPolicyRuleServer) DeleteComputeFirewallPolicyRule(ctx context.Context, request *computepb.DeleteComputeFirewallPolicyRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicyRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicyRule(ctx, ProtoToFirewallPolicyRule(request.GetResource()))

}

// ListComputeFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRuleList() method.
func (s *FirewallPolicyRuleServer) ListComputeFirewallPolicyRule(ctx context.Context, request *computepb.ListComputeFirewallPolicyRuleRequest) (*computepb.ListComputeFirewallPolicyRuleResponse, error) {
	cl, err := createConfigFirewallPolicyRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicyRule(ctx, ProtoToFirewallPolicyRule(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeFirewallPolicyRule
	for _, r := range resources.Items {
		rp := FirewallPolicyRuleToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeFirewallPolicyRuleResponse{Items: protos}, nil
}

func createConfigFirewallPolicyRule(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
