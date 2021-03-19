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
	binaryauthorizationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/binaryauthorization_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
)

// Server implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicyClusterAdmissionRulesEvaluationModeEnum converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum) *binaryauthorization.PolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyClusterAdmissionRulesEvaluationModeEnum(n[len("PolicyClusterAdmissionRulesEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyClusterAdmissionRulesEnforcementModeEnum converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum) *binaryauthorization.PolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyClusterAdmissionRulesEnforcementModeEnum(n[len("PolicyClusterAdmissionRulesEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEvaluationModeEnum converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum) *binaryauthorization.PolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyDefaultAdmissionRuleEvaluationModeEnum(n[len("PolicyDefaultAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyDefaultAdmissionRuleEnforcementModeEnum converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum) *binaryauthorization.PolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyDefaultAdmissionRuleEnforcementModeEnum(n[len("PolicyDefaultAdmissionRuleEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyGlobalPolicyEvaluationModeEnum converts a PolicyGlobalPolicyEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum) *binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum(n[len("PolicyGlobalPolicyEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionWhitelistPatterns converts a PolicyAdmissionWhitelistPatterns resource from its proto representation.
func ProtoToBinaryauthorizationPolicyAdmissionWhitelistPatterns(p *binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns) *binaryauthorization.PolicyAdmissionWhitelistPatterns {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.StringOrNil(p.NamePattern),
	}
	return obj
}

// ProtoToPolicyClusterAdmissionRules converts a PolicyClusterAdmissionRules resource from its proto representation.
func ProtoToBinaryauthorizationPolicyClusterAdmissionRules(p *binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules) *binaryauthorization.PolicyClusterAdmissionRules {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyClusterAdmissionRules{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicyDefaultAdmissionRule converts a PolicyDefaultAdmissionRule resource from its proto representation.
func ProtoToBinaryauthorizationPolicyDefaultAdmissionRule(p *binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRule) *binaryauthorization.PolicyDefaultAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyDefaultAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *binaryauthorizationpb.BinaryauthorizationPolicy) *binaryauthorization.Policy {
	obj := &binaryauthorization.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationPolicyDefaultAdmissionRule(p.GetDefaultAdmissionRule()),
		Description:                dcl.StringOrNil(p.Description),
		GlobalPolicyEvaluationMode: ProtoToBinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(p.GetGlobalPolicyEvaluationMode()),
		SelfLink:                   dcl.StringOrNil(p.SelfLink),
		Project:                    dcl.StringOrNil(p.Project),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetAdmissionWhitelistPatterns() {
		obj.AdmissionWhitelistPatterns = append(obj.AdmissionWhitelistPatterns, *ProtoToBinaryauthorizationPolicyAdmissionWhitelistPatterns(r))
	}
	return obj
}

// PolicyClusterAdmissionRulesEvaluationModeEnumToProto converts a PolicyClusterAdmissionRulesEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnumToProto(e *binaryauthorization.PolicyClusterAdmissionRulesEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum_value["PolicyClusterAdmissionRulesEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnum(0)
}

// PolicyClusterAdmissionRulesEnforcementModeEnumToProto converts a PolicyClusterAdmissionRulesEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnumToProto(e *binaryauthorization.PolicyClusterAdmissionRulesEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum_value["PolicyClusterAdmissionRulesEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnum(0)
}

// PolicyDefaultAdmissionRuleEvaluationModeEnumToProto converts a PolicyDefaultAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(e *binaryauthorization.PolicyDefaultAdmissionRuleEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum_value["PolicyDefaultAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnum(0)
}

// PolicyDefaultAdmissionRuleEnforcementModeEnumToProto converts a PolicyDefaultAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(e *binaryauthorization.PolicyDefaultAdmissionRuleEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum_value["PolicyDefaultAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnum(0)
}

// PolicyGlobalPolicyEvaluationModeEnumToProto converts a PolicyGlobalPolicyEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnumToProto(e *binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum_value["PolicyGlobalPolicyEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum(0)
}

// PolicyAdmissionWhitelistPatternsToProto converts a PolicyAdmissionWhitelistPatterns resource to its proto representation.
func BinaryauthorizationPolicyAdmissionWhitelistPatternsToProto(o *binaryauthorization.PolicyAdmissionWhitelistPatterns) *binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.ValueOrEmptyString(o.NamePattern),
	}
	return p
}

// PolicyClusterAdmissionRulesToProto converts a PolicyClusterAdmissionRules resource to its proto representation.
func BinaryauthorizationPolicyClusterAdmissionRulesToProto(o *binaryauthorization.PolicyClusterAdmissionRules) *binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyClusterAdmissionRules{
		EvaluationMode:  BinaryauthorizationPolicyClusterAdmissionRulesEvaluationModeEnumToProto(o.EvaluationMode),
		EnforcementMode: BinaryauthorizationPolicyClusterAdmissionRulesEnforcementModeEnumToProto(o.EnforcementMode),
	}
	for _, r := range o.RequireAttestationsBy {
		p.RequireAttestationsBy = append(p.RequireAttestationsBy, r)
	}
	return p
}

// PolicyDefaultAdmissionRuleToProto converts a PolicyDefaultAdmissionRule resource to its proto representation.
func BinaryauthorizationPolicyDefaultAdmissionRuleToProto(o *binaryauthorization.PolicyDefaultAdmissionRule) *binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRule {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyDefaultAdmissionRule{
		EvaluationMode:  BinaryauthorizationPolicyDefaultAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode),
		EnforcementMode: BinaryauthorizationPolicyDefaultAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode),
	}
	for _, r := range o.RequireAttestationsBy {
		p.RequireAttestationsBy = append(p.RequireAttestationsBy, r)
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *binaryauthorization.Policy) *binaryauthorizationpb.BinaryauthorizationPolicy {
	p := &binaryauthorizationpb.BinaryauthorizationPolicy{
		DefaultAdmissionRule:       BinaryauthorizationPolicyDefaultAdmissionRuleToProto(resource.DefaultAdmissionRule),
		Description:                dcl.ValueOrEmptyString(resource.Description),
		GlobalPolicyEvaluationMode: BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnumToProto(resource.GlobalPolicyEvaluationMode),
		SelfLink:                   dcl.ValueOrEmptyString(resource.SelfLink),
		Project:                    dcl.ValueOrEmptyString(resource.Project),
		UpdateTime:                 dcl.ValueOrEmptyString(resource.UpdateTime),
	}
	for _, r := range resource.AdmissionWhitelistPatterns {
		p.AdmissionWhitelistPatterns = append(p.AdmissionWhitelistPatterns, BinaryauthorizationPolicyAdmissionWhitelistPatternsToProto(&r))
	}

	return p
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *binaryauthorization.Client, request *binaryauthorizationpb.ApplyBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.BinaryauthorizationPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// ApplyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyBinaryauthorizationPolicy(ctx context.Context, request *binaryauthorizationpb.ApplyBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.BinaryauthorizationPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteBinaryauthorizationPolicy(ctx context.Context, request *binaryauthorizationpb.DeleteBinaryauthorizationPolicyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Policy")

}

// ListPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListBinaryauthorizationPolicy(ctx context.Context, request *binaryauthorizationpb.ListBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.ListBinaryauthorizationPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*binaryauthorizationpb.BinaryauthorizationPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	return &binaryauthorizationpb.ListBinaryauthorizationPolicyResponse{Items: protos}, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*binaryauthorization.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return binaryauthorization.NewClient(conf), nil
}
