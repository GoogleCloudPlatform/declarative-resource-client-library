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

// ProtoToPolicyAdmissionRuleEvaluationModeEnum converts a PolicyAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum) *binaryauthorization.PolicyAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionRuleEnforcementModeEnum converts a PolicyAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum(e binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum) *binaryauthorization.PolicyAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := binaryauthorization.PolicyAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum"):])
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
		e := binaryauthorization.PolicyGlobalPolicyEvaluationModeEnum(n[len("BinaryauthorizationPolicyGlobalPolicyEvaluationModeEnum"):])
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

// ProtoToPolicyAdmissionRule converts a PolicyAdmissionRule resource from its proto representation.
func ProtoToBinaryauthorizationPolicyAdmissionRule(p *binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRule) *binaryauthorization.PolicyAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.PolicyAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *binaryauthorizationpb.BinaryauthorizationPolicy) *binaryauthorization.Policy {
	obj := &binaryauthorization.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationPolicyAdmissionRule(p.GetDefaultAdmissionRule()),
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

// PolicyAdmissionRuleEvaluationModeEnumToProto converts a PolicyAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnumToProto(e *binaryauthorization.PolicyAdmissionRuleEvaluationModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum_value["PolicyAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnum(0)
}

// PolicyAdmissionRuleEnforcementModeEnumToProto converts a PolicyAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnumToProto(e *binaryauthorization.PolicyAdmissionRuleEnforcementModeEnum) binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum_value["PolicyAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnum(0)
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

// PolicyAdmissionRuleToProto converts a PolicyAdmissionRule resource to its proto representation.
func BinaryauthorizationPolicyAdmissionRuleToProto(o *binaryauthorization.PolicyAdmissionRule) *binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRule {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationPolicyAdmissionRule{
		EvaluationMode:  BinaryauthorizationPolicyAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode),
		EnforcementMode: BinaryauthorizationPolicyAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode),
	}
	for _, r := range o.RequireAttestationsBy {
		p.RequireAttestationsBy = append(p.RequireAttestationsBy, r)
	}
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *binaryauthorization.Policy) *binaryauthorizationpb.BinaryauthorizationPolicy {
	p := &binaryauthorizationpb.BinaryauthorizationPolicy{
		DefaultAdmissionRule:       BinaryauthorizationPolicyAdmissionRuleToProto(resource.DefaultAdmissionRule),
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

// ListBinaryauthorizationPolicy is a no-op method because Policy has no list method.
func (s *PolicyServer) ListBinaryauthorizationPolicy(_ context.Context, _ *binaryauthorizationpb.ListBinaryauthorizationPolicyRequest) (*binaryauthorizationpb.ListBinaryauthorizationPolicyResponse, error) {
	return nil, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*binaryauthorization.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return binaryauthorization.NewClient(conf), nil
}
