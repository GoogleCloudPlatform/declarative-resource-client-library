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

// Server implements the gRPC interface for CertificateTemplate.
type CertificateTemplateServer struct{}

// ProtoToCertificateTemplatePassthroughExtensionsKnownExtensionsEnum converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(e betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum) *beta.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := beta.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateTemplatePredefinedValues converts a CertificateTemplatePredefinedValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValues(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValues) *beta.CertificateTemplatePredefinedValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValues{
		KeyUsage:  ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage) *beta.CertificateTemplatePredefinedValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *beta.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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

// ProtoToCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *beta.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages converts a CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *beta.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesCaOptions converts a CertificateTemplatePredefinedValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesCaOptions(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesCaOptions) *beta.CertificateTemplatePredefinedValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesPolicyIds converts a CertificateTemplatePredefinedValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds) *beta.CertificateTemplatePredefinedValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensions converts a CertificateTemplatePredefinedValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions) *beta.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *beta.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraints converts a CertificateTemplateIdentityConstraints resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplateIdentityConstraints(p *betapb.PrivatecaBetaCertificateTemplateIdentityConstraints) *beta.CertificateTemplateIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplateIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.AllowSubjectAltNamesPassthrough),
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraintsCelExpression converts a CertificateTemplateIdentityConstraintsCelExpression resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression(p *betapb.PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression) *beta.CertificateTemplateIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensions converts a CertificateTemplatePassthroughExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensions(p *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensions) *beta.CertificateTemplatePassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensionsAdditionalExtensions converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions(p *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions) *beta.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplate converts a CertificateTemplate resource from its proto representation.
func ProtoToCertificateTemplate(p *betapb.PrivatecaBetaCertificateTemplate) *beta.CertificateTemplate {
	obj := &beta.CertificateTemplate{
		Name:                  dcl.StringOrNil(p.Name),
		PredefinedValues:      ProtoToPrivatecaBetaCertificateTemplatePredefinedValues(p.GetPredefinedValues()),
		IdentityConstraints:   ProtoToPrivatecaBetaCertificateTemplateIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaBetaCertificateTemplatePassthroughExtensions(p.GetPassthroughExtensions()),
		Description:           dcl.StringOrNil(p.Description),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
	}
	return obj
}

// CertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto(e *beta.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value["CertificateTemplatePassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(v)
	}
	return betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
}

// CertificateTemplatePredefinedValuesToProto converts a CertificateTemplatePredefinedValues resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesToProto(o *beta.CertificateTemplatePredefinedValues) *betapb.PrivatecaBetaCertificateTemplatePredefinedValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValues{
		KeyUsage:  PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateTemplatePredefinedValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsage) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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

// CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplatePredefinedValuesCaOptionsToProto converts a CertificateTemplatePredefinedValuesCaOptions resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesCaOptionsToProto(o *beta.CertificateTemplatePredefinedValuesCaOptions) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateTemplatePredefinedValuesPolicyIdsToProto converts a CertificateTemplatePredefinedValuesPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIdsToProto(o *beta.CertificateTemplatePredefinedValuesPolicyIds) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(o *beta.CertificateTemplatePredefinedValuesAdditionalExtensions) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o *beta.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplateIdentityConstraintsToProto converts a CertificateTemplateIdentityConstraints resource to its proto representation.
func PrivatecaBetaCertificateTemplateIdentityConstraintsToProto(o *beta.CertificateTemplateIdentityConstraints) *betapb.PrivatecaBetaCertificateTemplateIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplateIdentityConstraints{
		CelExpression:                   PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpressionToProto(o.CelExpression),
		AllowSubjectPassthrough:         dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough),
	}
	return p
}

// CertificateTemplateIdentityConstraintsCelExpressionToProto converts a CertificateTemplateIdentityConstraintsCelExpression resource to its proto representation.
func PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpressionToProto(o *beta.CertificateTemplateIdentityConstraintsCelExpression) *betapb.PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// CertificateTemplatePassthroughExtensionsToProto converts a CertificateTemplatePassthroughExtensions resource to its proto representation.
func PrivatecaBetaCertificateTemplatePassthroughExtensionsToProto(o *beta.CertificateTemplatePassthroughExtensions) *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePassthroughExtensions{}
	for _, r := range o.KnownExtensions {
		p.KnownExtensions = append(p.KnownExtensions, betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value[string(r)]))
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(o *beta.CertificateTemplatePassthroughExtensionsAdditionalExtensions) *betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplateToProto converts a CertificateTemplate resource to its proto representation.
func CertificateTemplateToProto(resource *beta.CertificateTemplate) *betapb.PrivatecaBetaCertificateTemplate {
	p := &betapb.PrivatecaBetaCertificateTemplate{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		PredefinedValues:      PrivatecaBetaCertificateTemplatePredefinedValuesToProto(resource.PredefinedValues),
		IdentityConstraints:   PrivatecaBetaCertificateTemplateIdentityConstraintsToProto(resource.IdentityConstraints),
		PassthroughExtensions: PrivatecaBetaCertificateTemplatePassthroughExtensionsToProto(resource.PassthroughExtensions),
		Description:           dcl.ValueOrEmptyString(resource.Description),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) applyCertificateTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCertificateTemplateRequest) (*betapb.PrivatecaBetaCertificateTemplate, error) {
	p := ProtoToCertificateTemplate(request.GetResource())
	res, err := c.ApplyCertificateTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateTemplateToProto(res)
	return r, nil
}

// ApplyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) ApplyPrivatecaBetaCertificateTemplate(ctx context.Context, request *betapb.ApplyPrivatecaBetaCertificateTemplateRequest) (*betapb.PrivatecaBetaCertificateTemplate, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCertificateTemplate(ctx, cl, request)
}

// DeleteCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Delete() method.
func (s *CertificateTemplateServer) DeletePrivatecaBetaCertificateTemplate(ctx context.Context, request *betapb.DeletePrivatecaBetaCertificateTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))

}

// ListPrivatecaBetaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplateList() method.
func (s *CertificateTemplateServer) ListPrivatecaBetaCertificateTemplate(ctx context.Context, request *betapb.ListPrivatecaBetaCertificateTemplateRequest) (*betapb.ListPrivatecaBetaCertificateTemplateResponse, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCertificateTemplate
	for _, r := range resources.Items {
		rp := CertificateTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListPrivatecaBetaCertificateTemplateResponse{Items: protos}, nil
}

func createConfigCertificateTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
