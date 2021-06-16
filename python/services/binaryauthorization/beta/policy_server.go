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
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/beta/binaryauthorization_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
)

// Server implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicyClusterAdmissionRulesEvaluationModeEnum converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum) *beta.PolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyClusterAdmissionRulesEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyClusterAdmissionRulesEnforcementModeEnum converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum) *beta.PolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyClusterAdmissionRulesEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEvaluationModeEnum converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum) *beta.PolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyDefaultAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEnforcementModeEnum converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum) *beta.PolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyDefaultAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyGlobalPolicyEvaluationModeEnum converts a PolicyGlobalPolicyEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum) *beta.PolicyGlobalPolicyEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyGlobalPolicyEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionWhitelistPatterns converts a PolicyAdmissionWhitelistPatterns resource from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyAdmissionWhitelistPatterns(p *betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns) *beta.PolicyAdmissionWhitelistPatterns {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.StringOrNil(p.NamePattern),
	}
	return obj
}

// ProtoToPolicyClusterAdmissionRules converts a PolicyClusterAdmissionRules resource from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRules(p *betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules) *beta.PolicyClusterAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyClusterAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyDefaultAdmissionRule converts a PolicyDefaultAdmissionRule resource from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRule(p *betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRule) *beta.PolicyDefaultAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyDefaultAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *betapb.BinaryauthorizationBetaPolicy) *beta.Policy {
	obj := &beta.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationBetaPolicyDefaultAdmissionRule(p.GetDefaultAdmissionRule()),
		Description:                dcl.StringOrNil(p.Description),
		GlobalPolicyEvaluationMode: ProtoToBinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(p.GetGlobalPolicyEvaluationMode()),
		SelfLink:                   dcl.StringOrNil(p.SelfLink),
		Project:                    dcl.StringOrNil(p.Project),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetAdmissionWhitelistPatterns() {
		obj.AdmissionWhitelistPatterns = append(obj.AdmissionWhitelistPatterns, *ProtoToBinaryauthorizationBetaPolicyAdmissionWhitelistPatterns(r))
	}
	return obj
}

// PolicyClusterAdmissionRulesEvaluationModeEnumToProto converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnumToProto(e *beta.PolicyClusterAdmissionRulesEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum_value["PolicyClusterAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum(0)
}

// PolicyClusterAdmissionRulesEnforcementModeEnumToProto converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnumToProto(e *beta.PolicyClusterAdmissionRulesEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum_value["PolicyClusterAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum(0)
}

// PolicyDefaultAdmissionRuleEvaluationModeEnumToProto converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(e *beta.PolicyDefaultAdmissionRuleEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum_value["PolicyDefaultAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
}

// PolicyDefaultAdmissionRuleEnforcementModeEnumToProto converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(e *beta.PolicyDefaultAdmissionRuleEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum_value["PolicyDefaultAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
}

// PolicyGlobalPolicyEvaluationModeEnumToProto converts a PolicyGlobalPolicyEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnumToProto(e *beta.PolicyGlobalPolicyEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum_value["PolicyGlobalPolicyEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum(0)
}

// PolicyAdmissionWhitelistPatternsToProto converts a PolicyAdmissionWhitelistPatterns resource to its proto representation.
func BinaryauthorizationBetaPolicyAdmissionWhitelistPatternsToProto(o *beta.PolicyAdmissionWhitelistPatterns) *betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.ValueOrEmptyString(o.NamePattern),
	}
	return p
}

// PolicyClusterAdmissionRulesToProto converts a PolicyClusterAdmissionRules resource to its proto representation.
func BinaryauthorizationBetaPolicyClusterAdmissionRulesToProto(o *beta.PolicyClusterAdmissionRules) *betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyClusterAdmissionRules{
		EvaluationMode:  BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode),
		EnforcementMode: BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode),
	}
	for _, r := range o.RequireAttestationsBy {
		p.RequireAttestationsBy = append(p.RequireAttestationsBy, r)
	}
	return p
}

// PolicyDefaultAdmissionRuleToProto converts a PolicyDefaultAdmissionRule resource to its proto representation.
func BinaryauthorizationBetaPolicyDefaultAdmissionRuleToProto(o *beta.PolicyDefaultAdmissionRule) *betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRule {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyDefaultAdmissionRule{
		EvaluationMode:  BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode),
		EnforcementMode: BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode),
	}
	for _, r := range o.RequireAttestationsBy {
		p.RequireAttestationsBy = append(p.RequireAttestationsBy, r)
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *beta.Policy) *betapb.BinaryauthorizationBetaPolicy {
	p := &betapb.BinaryauthorizationBetaPolicy{
		DefaultAdmissionRule:       BinaryauthorizationBetaPolicyDefaultAdmissionRuleToProto(resource.DefaultAdmissionRule),
		Description:                dcl.ValueOrEmptyString(resource.Description),
		GlobalPolicyEvaluationMode: BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnumToProto(resource.GlobalPolicyEvaluationMode),
		SelfLink:                   dcl.ValueOrEmptyString(resource.SelfLink),
		Project:                    dcl.ValueOrEmptyString(resource.Project),
		UpdateTime:                 dcl.ValueOrEmptyString(resource.UpdateTime),
	}
	for _, r := range resource.AdmissionWhitelistPatterns {
		p.AdmissionWhitelistPatterns = append(p.AdmissionWhitelistPatterns, BinaryauthorizationBetaPolicyAdmissionWhitelistPatternsToProto(&r))
	}

	return p
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyBinaryauthorizationBetaPolicyRequest) (*betapb.BinaryauthorizationBetaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyBinaryauthorizationBetaPolicy(ctx context.Context, request *betapb.ApplyBinaryauthorizationBetaPolicyRequest) (*betapb.BinaryauthorizationBetaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteBinaryauthorizationBetaPolicy(ctx context.Context, request *betapb.DeleteBinaryauthorizationBetaPolicyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Policy")

}

// ListBinaryauthorizationBetaPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListBinaryauthorizationBetaPolicy(ctx context.Context, request *betapb.ListBinaryauthorizationBetaPolicyRequest) (*betapb.ListBinaryauthorizationBetaPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BinaryauthorizationBetaPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListBinaryauthorizationBetaPolicyResponse{Items: protos}, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
