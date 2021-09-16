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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/orgpolicy/alpha/orgpolicy_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/alpha"
)

// Server implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicySpecRulesAllowAllEnum converts a PolicySpecRulesAllowAllEnum enum from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesAllowAllEnum(e alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum) *alpha.PolicySpecRulesAllowAllEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum_name[int32(e)]; ok {
		e := alpha.PolicySpecRulesAllowAllEnum(n[len("OrgpolicyAlphaPolicySpecRulesAllowAllEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpecRulesDenyAllEnum converts a PolicySpecRulesDenyAllEnum enum from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesDenyAllEnum(e alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum) *alpha.PolicySpecRulesDenyAllEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum_name[int32(e)]; ok {
		e := alpha.PolicySpecRulesDenyAllEnum(n[len("OrgpolicyAlphaPolicySpecRulesDenyAllEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpecRulesEnforceEnum converts a PolicySpecRulesEnforceEnum enum from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesEnforceEnum(e alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum) *alpha.PolicySpecRulesEnforceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum_name[int32(e)]; ok {
		e := alpha.PolicySpecRulesEnforceEnum(n[len("OrgpolicyAlphaPolicySpecRulesEnforceEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicySpec converts a PolicySpec resource from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpec(p *alphapb.OrgpolicyAlphaPolicySpec) *alpha.PolicySpec {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpec{
		Etag:              dcl.StringOrNil(p.Etag),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.InheritFromParent),
		Reset:             dcl.Bool(p.Reset_),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyAlphaPolicySpecRules(r))
	}
	return obj
}

// ProtoToPolicySpecRules converts a PolicySpecRules resource from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRules(p *alphapb.OrgpolicyAlphaPolicySpecRules) *alpha.PolicySpecRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpecRules{
		Values:    ProtoToOrgpolicyAlphaPolicySpecRulesValues(p.GetValues()),
		AllowAll:  ProtoToOrgpolicyAlphaPolicySpecRulesAllowAllEnum(p.GetAllowAll()),
		DenyAll:   ProtoToOrgpolicyAlphaPolicySpecRulesDenyAllEnum(p.GetDenyAll()),
		Enforce:   ProtoToOrgpolicyAlphaPolicySpecRulesEnforceEnum(p.GetEnforce()),
		Condition: ProtoToOrgpolicyAlphaPolicySpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicySpecRulesValues converts a PolicySpecRulesValues resource from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesValues(p *alphapb.OrgpolicyAlphaPolicySpecRulesValues) *alpha.PolicySpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicySpecRulesCondition converts a PolicySpecRulesCondition resource from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesCondition(p *alphapb.OrgpolicyAlphaPolicySpecRulesCondition) *alpha.PolicySpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpecRulesCondition{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *alphapb.OrgpolicyAlphaPolicy) *alpha.Policy {
	obj := &alpha.Policy{
		Name:   dcl.StringOrNil(p.Name),
		Spec:   ProtoToOrgpolicyAlphaPolicySpec(p.GetSpec()),
		Parent: dcl.StringOrNil(p.Parent),
	}
	return obj
}

// PolicySpecRulesAllowAllEnumToProto converts a PolicySpecRulesAllowAllEnum enum to its proto representation.
func OrgpolicyAlphaPolicySpecRulesAllowAllEnumToProto(e *alpha.PolicySpecRulesAllowAllEnum) alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum {
	if e == nil {
		return alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum(0)
	}
	if v, ok := alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum_value["PolicySpecRulesAllowAllEnum"+string(*e)]; ok {
		return alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum(v)
	}
	return alphapb.OrgpolicyAlphaPolicySpecRulesAllowAllEnum(0)
}

// PolicySpecRulesDenyAllEnumToProto converts a PolicySpecRulesDenyAllEnum enum to its proto representation.
func OrgpolicyAlphaPolicySpecRulesDenyAllEnumToProto(e *alpha.PolicySpecRulesDenyAllEnum) alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum {
	if e == nil {
		return alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum(0)
	}
	if v, ok := alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum_value["PolicySpecRulesDenyAllEnum"+string(*e)]; ok {
		return alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum(v)
	}
	return alphapb.OrgpolicyAlphaPolicySpecRulesDenyAllEnum(0)
}

// PolicySpecRulesEnforceEnumToProto converts a PolicySpecRulesEnforceEnum enum to its proto representation.
func OrgpolicyAlphaPolicySpecRulesEnforceEnumToProto(e *alpha.PolicySpecRulesEnforceEnum) alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum {
	if e == nil {
		return alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum(0)
	}
	if v, ok := alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum_value["PolicySpecRulesEnforceEnum"+string(*e)]; ok {
		return alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum(v)
	}
	return alphapb.OrgpolicyAlphaPolicySpecRulesEnforceEnum(0)
}

// PolicySpecToProto converts a PolicySpec resource to its proto representation.
func OrgpolicyAlphaPolicySpecToProto(o *alpha.PolicySpec) *alphapb.OrgpolicyAlphaPolicySpec {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpec{
		Etag:              dcl.ValueOrEmptyString(o.Etag),
		UpdateTime:        dcl.ValueOrEmptyString(o.UpdateTime),
		InheritFromParent: dcl.ValueOrEmptyBool(o.InheritFromParent),
		Reset_:            dcl.ValueOrEmptyBool(o.Reset),
	}
	for _, r := range o.Rules {
		p.Rules = append(p.Rules, OrgpolicyAlphaPolicySpecRulesToProto(&r))
	}
	return p
}

// PolicySpecRulesToProto converts a PolicySpecRules resource to its proto representation.
func OrgpolicyAlphaPolicySpecRulesToProto(o *alpha.PolicySpecRules) *alphapb.OrgpolicyAlphaPolicySpecRules {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpecRules{
		Values:    OrgpolicyAlphaPolicySpecRulesValuesToProto(o.Values),
		AllowAll:  OrgpolicyAlphaPolicySpecRulesAllowAllEnumToProto(o.AllowAll),
		DenyAll:   OrgpolicyAlphaPolicySpecRulesDenyAllEnumToProto(o.DenyAll),
		Enforce:   OrgpolicyAlphaPolicySpecRulesEnforceEnumToProto(o.Enforce),
		Condition: OrgpolicyAlphaPolicySpecRulesConditionToProto(o.Condition),
	}
	return p
}

// PolicySpecRulesValuesToProto converts a PolicySpecRulesValues resource to its proto representation.
func OrgpolicyAlphaPolicySpecRulesValuesToProto(o *alpha.PolicySpecRulesValues) *alphapb.OrgpolicyAlphaPolicySpecRulesValues {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpecRulesValues{}
	for _, r := range o.AllowedValues {
		p.AllowedValues = append(p.AllowedValues, r)
	}
	for _, r := range o.DeniedValues {
		p.DeniedValues = append(p.DeniedValues, r)
	}
	return p
}

// PolicySpecRulesConditionToProto converts a PolicySpecRulesCondition resource to its proto representation.
func OrgpolicyAlphaPolicySpecRulesConditionToProto(o *alpha.PolicySpecRulesCondition) *alphapb.OrgpolicyAlphaPolicySpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpecRulesCondition{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *alpha.Policy) *alphapb.OrgpolicyAlphaPolicy {
	p := &alphapb.OrgpolicyAlphaPolicy{
		Name:   dcl.ValueOrEmptyString(resource.Name),
		Spec:   OrgpolicyAlphaPolicySpecToProto(resource.Spec),
		Parent: dcl.ValueOrEmptyString(resource.Parent),
	}

	return p
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOrgpolicyAlphaPolicyRequest) (*alphapb.OrgpolicyAlphaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyOrgpolicyAlphaPolicy(ctx context.Context, request *alphapb.ApplyOrgpolicyAlphaPolicyRequest) (*alphapb.OrgpolicyAlphaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteOrgpolicyAlphaPolicy(ctx context.Context, request *alphapb.DeleteOrgpolicyAlphaPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePolicy(ctx, ProtoToPolicy(request.GetResource()))

}

// ListOrgpolicyAlphaPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListOrgpolicyAlphaPolicy(ctx context.Context, request *alphapb.ListOrgpolicyAlphaPolicyRequest) (*alphapb.ListOrgpolicyAlphaPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, ProtoToPolicy(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OrgpolicyAlphaPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListOrgpolicyAlphaPolicyResponse{Items: protos}, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
