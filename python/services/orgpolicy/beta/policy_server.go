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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/orgpolicy/beta/orgpolicy_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta"
)

// Server implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicySpecRulesAllowAllEnum converts a PolicySpecRulesAllowAllEnum enum from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesAllowAllEnum(e betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum) *beta.PolicySpecRulesAllowAllEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum_name[int32(e)]; ok {
		e := beta.PolicySpecRulesAllowAllEnum(n[len("OrgpolicyBetaPolicySpecRulesAllowAllEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpecRulesDenyAllEnum converts a PolicySpecRulesDenyAllEnum enum from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesDenyAllEnum(e betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum) *beta.PolicySpecRulesDenyAllEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum_name[int32(e)]; ok {
		e := beta.PolicySpecRulesDenyAllEnum(n[len("OrgpolicyBetaPolicySpecRulesDenyAllEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpecRulesEnforceEnum converts a PolicySpecRulesEnforceEnum enum from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesEnforceEnum(e betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum) *beta.PolicySpecRulesEnforceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum_name[int32(e)]; ok {
		e := beta.PolicySpecRulesEnforceEnum(n[len("OrgpolicyBetaPolicySpecRulesEnforceEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpec converts a PolicySpec resource from its proto representation.
func ProtoToOrgpolicyBetaPolicySpec(p *betapb.OrgpolicyBetaPolicySpec) *beta.PolicySpec {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpec{
		Etag:              dcl.StringOrNil(p.Etag),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.InheritFromParent),
		Reset:             dcl.Bool(p.Reset_),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyBetaPolicySpecRules(r))
	}
	return obj
}

// ProtoToPolicySpecRules converts a PolicySpecRules resource from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRules(p *betapb.OrgpolicyBetaPolicySpecRules) *beta.PolicySpecRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpecRules{
		Values:    ProtoToOrgpolicyBetaPolicySpecRulesValues(p.GetValues()),
		AllowAll:  ProtoToOrgpolicyBetaPolicySpecRulesAllowAllEnum(p.GetAllowAll()),
		DenyAll:   ProtoToOrgpolicyBetaPolicySpecRulesDenyAllEnum(p.GetDenyAll()),
		Enforce:   ProtoToOrgpolicyBetaPolicySpecRulesEnforceEnum(p.GetEnforce()),
		Condition: ProtoToOrgpolicyBetaPolicySpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicySpecRulesValues converts a PolicySpecRulesValues resource from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesValues(p *betapb.OrgpolicyBetaPolicySpecRulesValues) *beta.PolicySpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicySpecRulesCondition converts a PolicySpecRulesCondition resource from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesCondition(p *betapb.OrgpolicyBetaPolicySpecRulesCondition) *beta.PolicySpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpecRulesCondition{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *betapb.OrgpolicyBetaPolicy) *beta.Policy {
	obj := &beta.Policy{
		Name:   dcl.StringOrNil(p.Name),
		Spec:   ProtoToOrgpolicyBetaPolicySpec(p.GetSpec()),
		Parent: dcl.StringOrNil(p.Parent),
	}
	return obj
}

// PolicySpecRulesAllowAllEnumToProto converts a PolicySpecRulesAllowAllEnum enum to its proto representation.
func OrgpolicyBetaPolicySpecRulesAllowAllEnumToProto(e *beta.PolicySpecRulesAllowAllEnum) betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum {
	if e == nil {
		return betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum(0)
	}
	if v, ok := betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum_value["PolicySpecRulesAllowAllEnum"+string(*e)]; ok {
		return betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum(v)
	}
	return betapb.OrgpolicyBetaPolicySpecRulesAllowAllEnum(0)
}

// PolicySpecRulesDenyAllEnumToProto converts a PolicySpecRulesDenyAllEnum enum to its proto representation.
func OrgpolicyBetaPolicySpecRulesDenyAllEnumToProto(e *beta.PolicySpecRulesDenyAllEnum) betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum {
	if e == nil {
		return betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum(0)
	}
	if v, ok := betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum_value["PolicySpecRulesDenyAllEnum"+string(*e)]; ok {
		return betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum(v)
	}
	return betapb.OrgpolicyBetaPolicySpecRulesDenyAllEnum(0)
}

// PolicySpecRulesEnforceEnumToProto converts a PolicySpecRulesEnforceEnum enum to its proto representation.
func OrgpolicyBetaPolicySpecRulesEnforceEnumToProto(e *beta.PolicySpecRulesEnforceEnum) betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum {
	if e == nil {
		return betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum(0)
	}
	if v, ok := betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum_value["PolicySpecRulesEnforceEnum"+string(*e)]; ok {
		return betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum(v)
	}
	return betapb.OrgpolicyBetaPolicySpecRulesEnforceEnum(0)
}

// PolicySpecToProto converts a PolicySpec resource to its proto representation.
func OrgpolicyBetaPolicySpecToProto(o *beta.PolicySpec) *betapb.OrgpolicyBetaPolicySpec {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpec{
		Etag:              dcl.ValueOrEmptyString(o.Etag),
		UpdateTime:        dcl.ValueOrEmptyString(o.UpdateTime),
		InheritFromParent: dcl.ValueOrEmptyBool(o.InheritFromParent),
		Reset_:            dcl.ValueOrEmptyBool(o.Reset),
	}
	for _, r := range o.Rules {
		p.Rules = append(p.Rules, OrgpolicyBetaPolicySpecRulesToProto(&r))
	}
	return p
}

// PolicySpecRulesToProto converts a PolicySpecRules resource to its proto representation.
func OrgpolicyBetaPolicySpecRulesToProto(o *beta.PolicySpecRules) *betapb.OrgpolicyBetaPolicySpecRules {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpecRules{
		Values:    OrgpolicyBetaPolicySpecRulesValuesToProto(o.Values),
		AllowAll:  OrgpolicyBetaPolicySpecRulesAllowAllEnumToProto(o.AllowAll),
		DenyAll:   OrgpolicyBetaPolicySpecRulesDenyAllEnumToProto(o.DenyAll),
		Enforce:   OrgpolicyBetaPolicySpecRulesEnforceEnumToProto(o.Enforce),
		Condition: OrgpolicyBetaPolicySpecRulesConditionToProto(o.Condition),
	}
	return p
}

// PolicySpecRulesValuesToProto converts a PolicySpecRulesValues resource to its proto representation.
func OrgpolicyBetaPolicySpecRulesValuesToProto(o *beta.PolicySpecRulesValues) *betapb.OrgpolicyBetaPolicySpecRulesValues {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpecRulesValues{}
	for _, r := range o.AllowedValues {
		p.AllowedValues = append(p.AllowedValues, r)
	}
	for _, r := range o.DeniedValues {
		p.DeniedValues = append(p.DeniedValues, r)
	}
	return p
}

// PolicySpecRulesConditionToProto converts a PolicySpecRulesCondition resource to its proto representation.
func OrgpolicyBetaPolicySpecRulesConditionToProto(o *beta.PolicySpecRulesCondition) *betapb.OrgpolicyBetaPolicySpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpecRulesCondition{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *beta.Policy) *betapb.OrgpolicyBetaPolicy {
	p := &betapb.OrgpolicyBetaPolicy{
		Name:   dcl.ValueOrEmptyString(resource.Name),
		Spec:   OrgpolicyBetaPolicySpecToProto(resource.Spec),
		Parent: dcl.ValueOrEmptyString(resource.Parent),
	}

	return p
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyOrgpolicyBetaPolicyRequest) (*betapb.OrgpolicyBetaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyOrgpolicyBetaPolicy(ctx context.Context, request *betapb.ApplyOrgpolicyBetaPolicyRequest) (*betapb.OrgpolicyBetaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteOrgpolicyBetaPolicy(ctx context.Context, request *betapb.DeleteOrgpolicyBetaPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePolicy(ctx, ProtoToPolicy(request.GetResource()))

}

// ListOrgpolicyBetaPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListOrgpolicyBetaPolicy(ctx context.Context, request *betapb.ListOrgpolicyBetaPolicyRequest) (*betapb.ListOrgpolicyBetaPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, ProtoToPolicy(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.OrgpolicyBetaPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListOrgpolicyBetaPolicyResponse{Items: protos}, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
