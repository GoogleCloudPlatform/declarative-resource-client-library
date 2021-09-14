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
	orgpolicypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/orgpolicy/orgpolicy_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy"
)

// Server implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicySpecRulesAllowAllEnum converts a PolicySpecRulesAllowAllEnum enum from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesAllowAllEnum(e orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum) *orgpolicy.PolicySpecRulesAllowAllEnum {
	if e == 0 {
		return nil
	}
	if n, ok := orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum_name[int32(e)]; ok {
		e := orgpolicy.PolicySpecRulesAllowAllEnum(n[len("OrgpolicyPolicySpecRulesAllowAllEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpecRulesDenyAllEnum converts a PolicySpecRulesDenyAllEnum enum from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesDenyAllEnum(e orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum) *orgpolicy.PolicySpecRulesDenyAllEnum {
	if e == 0 {
		return nil
	}
	if n, ok := orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum_name[int32(e)]; ok {
		e := orgpolicy.PolicySpecRulesDenyAllEnum(n[len("OrgpolicyPolicySpecRulesDenyAllEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpecRulesEnforceEnum converts a PolicySpecRulesEnforceEnum enum from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesEnforceEnum(e orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum) *orgpolicy.PolicySpecRulesEnforceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum_name[int32(e)]; ok {
		e := orgpolicy.PolicySpecRulesEnforceEnum(n[len("OrgpolicyPolicySpecRulesEnforceEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpec converts a PolicySpec resource from its proto representation.
func ProtoToOrgpolicyPolicySpec(p *orgpolicypb.OrgpolicyPolicySpec) *orgpolicy.PolicySpec {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpec{
		Etag:              dcl.StringOrNil(p.Etag),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.InheritFromParent),
		Reset:             dcl.Bool(p.Reset_),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyPolicySpecRules(r))
	}
	return obj
}

// ProtoToPolicySpecRules converts a PolicySpecRules resource from its proto representation.
func ProtoToOrgpolicyPolicySpecRules(p *orgpolicypb.OrgpolicyPolicySpecRules) *orgpolicy.PolicySpecRules {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpecRules{
		Values:    ProtoToOrgpolicyPolicySpecRulesValues(p.GetValues()),
		AllowAll:  ProtoToOrgpolicyPolicySpecRulesAllowAllEnum(p.GetAllowAll()),
		DenyAll:   ProtoToOrgpolicyPolicySpecRulesDenyAllEnum(p.GetDenyAll()),
		Enforce:   ProtoToOrgpolicyPolicySpecRulesEnforceEnum(p.GetEnforce()),
		Condition: ProtoToOrgpolicyPolicySpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicySpecRulesValues converts a PolicySpecRulesValues resource from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesValues(p *orgpolicypb.OrgpolicyPolicySpecRulesValues) *orgpolicy.PolicySpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicySpecRulesCondition converts a PolicySpecRulesCondition resource from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesCondition(p *orgpolicypb.OrgpolicyPolicySpecRulesCondition) *orgpolicy.PolicySpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpecRulesCondition{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *orgpolicypb.OrgpolicyPolicy) *orgpolicy.Policy {
	obj := &orgpolicy.Policy{
		Name:   dcl.StringOrNil(p.Name),
		Spec:   ProtoToOrgpolicyPolicySpec(p.GetSpec()),
		Parent: dcl.StringOrNil(p.Parent),
	}
	return obj
}

// PolicySpecRulesAllowAllEnumToProto converts a PolicySpecRulesAllowAllEnum enum to its proto representation.
func OrgpolicyPolicySpecRulesAllowAllEnumToProto(e *orgpolicy.PolicySpecRulesAllowAllEnum) orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum {
	if e == nil {
		return orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum(0)
	}
	if v, ok := orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum_value["PolicySpecRulesAllowAllEnum"+string(*e)]; ok {
		return orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum(v)
	}
	return orgpolicypb.OrgpolicyPolicySpecRulesAllowAllEnum(0)
}

// PolicySpecRulesDenyAllEnumToProto converts a PolicySpecRulesDenyAllEnum enum to its proto representation.
func OrgpolicyPolicySpecRulesDenyAllEnumToProto(e *orgpolicy.PolicySpecRulesDenyAllEnum) orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum {
	if e == nil {
		return orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum(0)
	}
	if v, ok := orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum_value["PolicySpecRulesDenyAllEnum"+string(*e)]; ok {
		return orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum(v)
	}
	return orgpolicypb.OrgpolicyPolicySpecRulesDenyAllEnum(0)
}

// PolicySpecRulesEnforceEnumToProto converts a PolicySpecRulesEnforceEnum enum to its proto representation.
func OrgpolicyPolicySpecRulesEnforceEnumToProto(e *orgpolicy.PolicySpecRulesEnforceEnum) orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum {
	if e == nil {
		return orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum(0)
	}
	if v, ok := orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum_value["PolicySpecRulesEnforceEnum"+string(*e)]; ok {
		return orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum(v)
	}
	return orgpolicypb.OrgpolicyPolicySpecRulesEnforceEnum(0)
}

// PolicySpecToProto converts a PolicySpec resource to its proto representation.
func OrgpolicyPolicySpecToProto(o *orgpolicy.PolicySpec) *orgpolicypb.OrgpolicyPolicySpec {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpec{
		Etag:              dcl.ValueOrEmptyString(o.Etag),
		UpdateTime:        dcl.ValueOrEmptyString(o.UpdateTime),
		InheritFromParent: dcl.ValueOrEmptyBool(o.InheritFromParent),
		Reset_:            dcl.ValueOrEmptyBool(o.Reset),
	}
	for _, r := range o.Rules {
		p.Rules = append(p.Rules, OrgpolicyPolicySpecRulesToProto(&r))
	}
	return p
}

// PolicySpecRulesToProto converts a PolicySpecRules resource to its proto representation.
func OrgpolicyPolicySpecRulesToProto(o *orgpolicy.PolicySpecRules) *orgpolicypb.OrgpolicyPolicySpecRules {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpecRules{
		Values:    OrgpolicyPolicySpecRulesValuesToProto(o.Values),
		AllowAll:  OrgpolicyPolicySpecRulesAllowAllEnumToProto(o.AllowAll),
		DenyAll:   OrgpolicyPolicySpecRulesDenyAllEnumToProto(o.DenyAll),
		Enforce:   OrgpolicyPolicySpecRulesEnforceEnumToProto(o.Enforce),
		Condition: OrgpolicyPolicySpecRulesConditionToProto(o.Condition),
	}
	return p
}

// PolicySpecRulesValuesToProto converts a PolicySpecRulesValues resource to its proto representation.
func OrgpolicyPolicySpecRulesValuesToProto(o *orgpolicy.PolicySpecRulesValues) *orgpolicypb.OrgpolicyPolicySpecRulesValues {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpecRulesValues{}
	for _, r := range o.AllowedValues {
		p.AllowedValues = append(p.AllowedValues, r)
	}
	for _, r := range o.DeniedValues {
		p.DeniedValues = append(p.DeniedValues, r)
	}
	return p
}

// PolicySpecRulesConditionToProto converts a PolicySpecRulesCondition resource to its proto representation.
func OrgpolicyPolicySpecRulesConditionToProto(o *orgpolicy.PolicySpecRulesCondition) *orgpolicypb.OrgpolicyPolicySpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpecRulesCondition{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *orgpolicy.Policy) *orgpolicypb.OrgpolicyPolicy {
	p := &orgpolicypb.OrgpolicyPolicy{
		Name:   dcl.ValueOrEmptyString(resource.Name),
		Spec:   OrgpolicyPolicySpecToProto(resource.Spec),
		Parent: dcl.ValueOrEmptyString(resource.Parent),
	}

	return p
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *orgpolicy.Client, request *orgpolicypb.ApplyOrgpolicyPolicyRequest) (*orgpolicypb.OrgpolicyPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyOrgpolicyPolicy(ctx context.Context, request *orgpolicypb.ApplyOrgpolicyPolicyRequest) (*orgpolicypb.OrgpolicyPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteOrgpolicyPolicy(ctx context.Context, request *orgpolicypb.DeleteOrgpolicyPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePolicy(ctx, ProtoToPolicy(request.GetResource()))

}

// ListOrgpolicyPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListOrgpolicyPolicy(ctx context.Context, request *orgpolicypb.ListOrgpolicyPolicyRequest) (*orgpolicypb.ListOrgpolicyPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, ProtoToPolicy(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*orgpolicypb.OrgpolicyPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	return &orgpolicypb.ListOrgpolicyPolicyResponse{Items: protos}, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*orgpolicy.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return orgpolicy.NewClient(conf), nil
}
