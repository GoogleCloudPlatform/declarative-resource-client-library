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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/beta/privateca_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
)

// Server implements the gRPC interface for CaPool.
type CaPoolServer struct{}

// ProtoToCaPoolTierEnum converts a CaPoolTierEnum enum from its proto representation.
func ProtoToPrivatecaBetaCaPoolTierEnum(e betapb.PrivatecaBetaCaPoolTierEnum) *beta.CaPoolTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCaPoolTierEnum_name[int32(e)]; ok {
		e := beta.CaPoolTierEnum(n[len("PrivatecaBetaCaPoolTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(e betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_name[int32(e)]; ok {
		e := beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(n[len("PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(e betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) *beta.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := beta.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCaPoolIssuancePolicy converts a CaPoolIssuancePolicy resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicy(p *betapb.PrivatecaBetaCaPoolIssuancePolicy) *beta.CaPoolIssuancePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicy{
		MaximumLifetime:       dcl.StringOrNil(p.MaximumLifetime),
		AllowedIssuanceModes:  ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes(p.GetAllowedIssuanceModes()),
		BaselineValues:        ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValues(p.GetBaselineValues()),
		IdentityConstraints:   ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions(p.GetPassthroughExtensions()),
	}
	for _, r := range p.GetAllowedKeyTypes() {
		obj.AllowedKeyTypes = append(obj.AllowedKeyTypes, *ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypes converts a CaPoolIssuancePolicyAllowedKeyTypes resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes) *beta.CaPoolIssuancePolicyAllowedKeyTypes {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedKeyTypes{
		Rsa:           ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa(p.GetRsa()),
		EllipticCurve: ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p.GetEllipticCurve()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesRsa converts a CaPoolIssuancePolicyAllowedKeyTypesRsa resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa) *beta.CaPoolIssuancePolicyAllowedKeyTypesRsa {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedKeyTypesRsa{
		MinModulusSize: dcl.Int64OrNil(p.MinModulusSize),
		MaxModulusSize: dcl.Int64OrNil(p.MaxModulusSize),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{
		SignatureAlgorithm: ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyAllowedIssuanceModes converts a CaPoolIssuancePolicyAllowedIssuanceModes resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes(p *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes) *beta.CaPoolIssuancePolicyAllowedIssuanceModes {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.Bool(p.AllowCsrBasedIssuance),
		AllowConfigBasedIssuance: dcl.Bool(p.AllowConfigBasedIssuance),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValues converts a CaPoolIssuancePolicyBaselineValues resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValues(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValues) *beta.CaPoolIssuancePolicyBaselineValues {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValues{
		KeyUsage:  ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{
		DigitalSignature:  dcl.Bool(p.DigitalSignature),
		ContentCommitment: dcl.Bool(p.ContentCommitment),
		KeyEncipherment:   dcl.Bool(p.KeyEncipherment),
		DataEncipherment:  dcl.Bool(p.DataEncipherment),
		KeyAgreement:      dcl.Bool(p.KeyAgreement),
		CertSign:          dcl.Bool(p.CertSign),
		CrlSign:           dcl.Bool(p.CrlSign),
		EncipherOnly:      dcl.Bool(p.EncipherOnly),
		DecipherOnly:      dcl.Bool(p.DecipherOnly),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage converts a CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages converts a CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesCaOptions converts a CaPoolIssuancePolicyBaselineValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions) *beta.CaPoolIssuancePolicyBaselineValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesPolicyIds converts a CaPoolIssuancePolicyBaselineValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds) *beta.CaPoolIssuancePolicyBaselineValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensions converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraints converts a CaPoolIssuancePolicyIdentityConstraints resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraints(p *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraints) *beta.CaPoolIssuancePolicyIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.AllowSubjectAltNamesPassthrough),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyIdentityConstraintsCelExpression converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression(p *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression) *beta.CaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensions converts a CaPoolIssuancePolicyPassthroughExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions) *beta.CaPoolIssuancePolicyPassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyPassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions(p *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *beta.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCaPoolPublishingOptions converts a CaPoolPublishingOptions resource from its proto representation.
func ProtoToPrivatecaBetaCaPoolPublishingOptions(p *betapb.PrivatecaBetaCaPoolPublishingOptions) *beta.CaPoolPublishingOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CaPoolPublishingOptions{
		PublishCaCert: dcl.Bool(p.PublishCaCert),
		PublishCrl:    dcl.Bool(p.PublishCrl),
	}
	return obj
}

// ProtoToCaPool converts a CaPool resource from its proto representation.
func ProtoToCaPool(p *betapb.PrivatecaBetaCaPool) *beta.CaPool {
	obj := &beta.CaPool{
		Name:              dcl.StringOrNil(p.Name),
		Tier:              ProtoToPrivatecaBetaCaPoolTierEnum(p.GetTier()),
		IssuancePolicy:    ProtoToPrivatecaBetaCaPoolIssuancePolicy(p.GetIssuancePolicy()),
		PublishingOptions: ProtoToPrivatecaBetaCaPoolPublishingOptions(p.GetPublishingOptions()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	return obj
}

// CaPoolTierEnumToProto converts a CaPoolTierEnum enum to its proto representation.
func PrivatecaBetaCaPoolTierEnumToProto(e *beta.CaPoolTierEnum) betapb.PrivatecaBetaCaPoolTierEnum {
	if e == nil {
		return betapb.PrivatecaBetaCaPoolTierEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCaPoolTierEnum_value["CaPoolTierEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCaPoolTierEnum(v)
	}
	return betapb.PrivatecaBetaCaPoolTierEnum(0)
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum enum to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(e *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum) betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum {
	if e == nil {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum_value["CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(v)
	}
	return betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnum(0)
}

// CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto converts a CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnumToProto(e *beta.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum) betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value["CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(v)
	}
	return betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(0)
}

// CaPoolIssuancePolicyToProto converts a CaPoolIssuancePolicy resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyToProto(o *beta.CaPoolIssuancePolicy) *betapb.PrivatecaBetaCaPoolIssuancePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicy{
		MaximumLifetime:       dcl.ValueOrEmptyString(o.MaximumLifetime),
		AllowedIssuanceModes:  PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o.AllowedIssuanceModes),
		BaselineValues:        PrivatecaBetaCaPoolIssuancePolicyBaselineValuesToProto(o.BaselineValues),
		IdentityConstraints:   PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsToProto(o.IdentityConstraints),
		PassthroughExtensions: PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsToProto(o.PassthroughExtensions),
	}
	for _, r := range o.AllowedKeyTypes {
		p.AllowedKeyTypes = append(p.AllowedKeyTypes, PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesToProto(&r))
	}
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesToProto converts a CaPoolIssuancePolicyAllowedKeyTypes resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesToProto(o *beta.CaPoolIssuancePolicyAllowedKeyTypes) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypes{
		Rsa:           PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o.Rsa),
		EllipticCurve: PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o.EllipticCurve),
	}
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesRsaToProto converts a CaPoolIssuancePolicyAllowedKeyTypesRsa resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsaToProto(o *beta.CaPoolIssuancePolicyAllowedKeyTypesRsa) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesRsa{
		MinModulusSize: dcl.ValueOrEmptyInt64(o.MinModulusSize),
		MaxModulusSize: dcl.ValueOrEmptyInt64(o.MaxModulusSize),
	}
	return p
}

// CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto converts a CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveToProto(o *beta.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{
		SignatureAlgorithm: PrivatecaBetaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumToProto(o.SignatureAlgorithm),
	}
	return p
}

// CaPoolIssuancePolicyAllowedIssuanceModesToProto converts a CaPoolIssuancePolicyAllowedIssuanceModes resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModesToProto(o *beta.CaPoolIssuancePolicyAllowedIssuanceModes) *betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.ValueOrEmptyBool(o.AllowCsrBasedIssuance),
		AllowConfigBasedIssuance: dcl.ValueOrEmptyBool(o.AllowConfigBasedIssuance),
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesToProto converts a CaPoolIssuancePolicyBaselineValues resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesToProto(o *beta.CaPoolIssuancePolicyBaselineValues) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValues{
		KeyUsage:  PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsage resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsage) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsageToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{
		DigitalSignature:  dcl.ValueOrEmptyBool(o.DigitalSignature),
		ContentCommitment: dcl.ValueOrEmptyBool(o.ContentCommitment),
		KeyEncipherment:   dcl.ValueOrEmptyBool(o.KeyEncipherment),
		DataEncipherment:  dcl.ValueOrEmptyBool(o.DataEncipherment),
		KeyAgreement:      dcl.ValueOrEmptyBool(o.KeyAgreement),
		CertSign:          dcl.ValueOrEmptyBool(o.CertSign),
		CrlSign:           dcl.ValueOrEmptyBool(o.CrlSign),
		EncipherOnly:      dcl.ValueOrEmptyBool(o.EncipherOnly),
		DecipherOnly:      dcl.ValueOrEmptyBool(o.DecipherOnly),
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesCaOptionsToProto converts a CaPoolIssuancePolicyBaselineValuesCaOptions resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptionsToProto(o *beta.CaPoolIssuancePolicyBaselineValuesCaOptions) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesPolicyIdsToProto converts a CaPoolIssuancePolicyBaselineValuesPolicyIds resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIdsToProto(o *beta.CaPoolIssuancePolicyBaselineValuesPolicyIds) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsToProto(o *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensions{
		ObjectId: PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto converts a CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectIdToProto(o *beta.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsToProto converts a CaPoolIssuancePolicyIdentityConstraints resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsToProto(o *beta.CaPoolIssuancePolicyIdentityConstraints) *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraints{
		CelExpression:                   PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o.CelExpression),
		AllowSubjectPassthrough:         dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough),
	}
	return p
}

// CaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto converts a CaPoolIssuancePolicyIdentityConstraintsCelExpression resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpressionToProto(o *beta.CaPoolIssuancePolicyIdentityConstraintsCelExpression) *betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyIdentityConstraintsCelExpression{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensions resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsToProto(o *beta.CaPoolIssuancePolicyPassthroughExtensions) *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensions{}
	for _, r := range o.KnownExtensions {
		p.KnownExtensions = append(p.KnownExtensions, betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum_value[string(r)]))
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(&r))
	}
	return p
}

// CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto converts a CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensionsToProto(o *beta.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions) *betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CaPoolPublishingOptionsToProto converts a CaPoolPublishingOptions resource to its proto representation.
func PrivatecaBetaCaPoolPublishingOptionsToProto(o *beta.CaPoolPublishingOptions) *betapb.PrivatecaBetaCaPoolPublishingOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCaPoolPublishingOptions{
		PublishCaCert: dcl.ValueOrEmptyBool(o.PublishCaCert),
		PublishCrl:    dcl.ValueOrEmptyBool(o.PublishCrl),
	}
	return p
}

// CaPoolToProto converts a CaPool resource to its proto representation.
func CaPoolToProto(resource *beta.CaPool) *betapb.PrivatecaBetaCaPool {
	p := &betapb.PrivatecaBetaCaPool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Tier:              PrivatecaBetaCaPoolTierEnumToProto(resource.Tier),
		IssuancePolicy:    PrivatecaBetaCaPoolIssuancePolicyToProto(resource.IssuancePolicy),
		PublishingOptions: PrivatecaBetaCaPoolPublishingOptionsToProto(resource.PublishingOptions),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCaPool handles the gRPC request by passing it to the underlying CaPool Apply() method.
func (s *CaPoolServer) applyCaPool(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCaPoolRequest) (*betapb.PrivatecaBetaCaPool, error) {
	p := ProtoToCaPool(request.GetResource())
	res, err := c.ApplyCaPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CaPoolToProto(res)
	return r, nil
}

// ApplyCaPool handles the gRPC request by passing it to the underlying CaPool Apply() method.
func (s *CaPoolServer) ApplyPrivatecaBetaCaPool(ctx context.Context, request *betapb.ApplyPrivatecaBetaCaPoolRequest) (*betapb.PrivatecaBetaCaPool, error) {
	cl, err := createConfigCaPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCaPool(ctx, cl, request)
}

// DeleteCaPool handles the gRPC request by passing it to the underlying CaPool Delete() method.
func (s *CaPoolServer) DeletePrivatecaBetaCaPool(ctx context.Context, request *betapb.DeletePrivatecaBetaCaPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCaPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCaPool(ctx, ProtoToCaPool(request.GetResource()))

}

// ListPrivatecaBetaCaPool handles the gRPC request by passing it to the underlying CaPoolList() method.
func (s *CaPoolServer) ListPrivatecaBetaCaPool(ctx context.Context, request *betapb.ListPrivatecaBetaCaPoolRequest) (*betapb.ListPrivatecaBetaCaPoolResponse, error) {
	cl, err := createConfigCaPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCaPool(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCaPool
	for _, r := range resources.Items {
		rp := CaPoolToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListPrivatecaBetaCaPoolResponse{Items: protos}, nil
}

func createConfigCaPool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
