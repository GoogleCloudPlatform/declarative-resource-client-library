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

// ProtoToPolicyAdmissionRuleEvaluationModeEnum converts a PolicyAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum(e betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum) *beta.PolicyAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.PolicyAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionRuleEnforcementModeEnum converts a PolicyAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum(e betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum) *beta.PolicyAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := beta.PolicyAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum"):])
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

// ProtoToPolicyAdmissionRule converts a PolicyAdmissionRule resource from its proto representation.
func ProtoToBinaryauthorizationBetaPolicyAdmissionRule(p *betapb.BinaryauthorizationBetaPolicyAdmissionRule) *beta.PolicyAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *betapb.BinaryauthorizationBetaPolicy) *beta.Policy {
	obj := &beta.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationBetaPolicyAdmissionRule(p.GetDefaultAdmissionRule()),
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

// PolicyAdmissionRuleEvaluationModeEnumToProto converts a PolicyAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnumToProto(e *beta.PolicyAdmissionRuleEvaluationModeEnum) betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum_value["PolicyAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnum(0)
}

// PolicyAdmissionRuleEnforcementModeEnumToProto converts a PolicyAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnumToProto(e *beta.PolicyAdmissionRuleEnforcementModeEnum) betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum_value["PolicyAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum(v)
	}
	return betapb.BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnum(0)
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

// PolicyAdmissionRuleToProto converts a PolicyAdmissionRule resource to its proto representation.
func BinaryauthorizationBetaPolicyAdmissionRuleToProto(o *beta.PolicyAdmissionRule) *betapb.BinaryauthorizationBetaPolicyAdmissionRule {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaPolicyAdmissionRule{
		EvaluationMode:  BinaryauthorizationBetaPolicyAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode),
		EnforcementMode: BinaryauthorizationBetaPolicyAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode),
	}
	for _, r := range o.RequireAttestationsBy {
		p.RequireAttestationsBy = append(p.RequireAttestationsBy, r)
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *beta.Policy) *betapb.BinaryauthorizationBetaPolicy {
	p := &betapb.BinaryauthorizationBetaPolicy{
		DefaultAdmissionRule:       BinaryauthorizationBetaPolicyAdmissionRuleToProto(resource.DefaultAdmissionRule),
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

// ListBinaryauthorizationBetaPolicy is a no-op method because Policy has no list method.
func (s *PolicyServer) ListBinaryauthorizationBetaPolicy(_ context.Context, _ *betapb.ListBinaryauthorizationBetaPolicyRequest) (*betapb.ListBinaryauthorizationBetaPolicyResponse, error) {
	return nil, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
