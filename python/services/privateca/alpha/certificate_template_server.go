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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/alpha/privateca_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/alpha"
)

// Server implements the gRPC interface for CertificateTemplate.
type CertificateTemplateServer struct{}

// ProtoToCertificateTemplatePassthroughExtensionsKnownExtensionsEnum converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(e alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum) *alpha.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := alpha.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateTemplatePredefinedValues converts a CertificateTemplatePredefinedValues resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValues(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValues) *alpha.CertificateTemplatePredefinedValues {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValues{
		KeyUsage:  ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage) *alpha.CertificateTemplatePredefinedValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *alpha.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *alpha.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *alpha.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesCaOptions converts a CertificateTemplatePredefinedValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions) *alpha.CertificateTemplatePredefinedValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesPolicyIds converts a CertificateTemplatePredefinedValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds) *alpha.CertificateTemplatePredefinedValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensions converts a CertificateTemplatePredefinedValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions) *alpha.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *alpha.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraints converts a CertificateTemplateIdentityConstraints resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraints(p *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraints) *alpha.CertificateTemplateIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplateIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.AllowSubjectAltNamesPassthrough),
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraintsCelExpression converts a CertificateTemplateIdentityConstraintsCelExpression resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression(p *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression) *alpha.CertificateTemplateIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensions converts a CertificateTemplatePassthroughExtensions resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensions(p *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensions) *alpha.CertificateTemplatePassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensionsAdditionalExtensions converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions) *alpha.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplate converts a CertificateTemplate resource from its proto representation.
func ProtoToCertificateTemplate(p *alphapb.PrivatecaAlphaCertificateTemplate) *alpha.CertificateTemplate {
	obj := &alpha.CertificateTemplate{
		Name:                  dcl.StringOrNil(p.Name),
		PredefinedValues:      ProtoToPrivatecaAlphaCertificateTemplatePredefinedValues(p.GetPredefinedValues()),
		IdentityConstraints:   ProtoToPrivatecaAlphaCertificateTemplateIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaAlphaCertificateTemplatePassthroughExtensions(p.GetPassthroughExtensions()),
		Description:           dcl.StringOrNil(p.Description),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
	}
	return obj
}

// CertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto(e *alpha.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value["CertificateTemplatePassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
}

// CertificateTemplatePredefinedValuesToProto converts a CertificateTemplatePredefinedValues resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesToProto(o *alpha.CertificateTemplatePredefinedValues) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValues {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValues{
		KeyUsage:  PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsage resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsage) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
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
func PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplatePredefinedValuesCaOptionsToProto converts a CertificateTemplatePredefinedValuesCaOptions resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptionsToProto(o *alpha.CertificateTemplatePredefinedValuesCaOptions) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateTemplatePredefinedValuesPolicyIdsToProto converts a CertificateTemplatePredefinedValuesPolicyIds resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIdsToProto(o *alpha.CertificateTemplatePredefinedValuesPolicyIds) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensions resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(o *alpha.CertificateTemplatePredefinedValuesAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o *alpha.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplateIdentityConstraintsToProto converts a CertificateTemplateIdentityConstraints resource to its proto representation.
func PrivatecaAlphaCertificateTemplateIdentityConstraintsToProto(o *alpha.CertificateTemplateIdentityConstraints) *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraints{
		CelExpression:                   PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpressionToProto(o.CelExpression),
		AllowSubjectPassthrough:         dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough),
	}
	return p
}

// CertificateTemplateIdentityConstraintsCelExpressionToProto converts a CertificateTemplateIdentityConstraintsCelExpression resource to its proto representation.
func PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpressionToProto(o *alpha.CertificateTemplateIdentityConstraintsCelExpression) *alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// CertificateTemplatePassthroughExtensionsToProto converts a CertificateTemplatePassthroughExtensions resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePassthroughExtensionsToProto(o *alpha.CertificateTemplatePassthroughExtensions) *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensions{}
	for _, r := range o.KnownExtensions {
		p.KnownExtensions = append(p.KnownExtensions, alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value[string(r)]))
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions resource to its proto representation.
func PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(o *alpha.CertificateTemplatePassthroughExtensionsAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplateToProto converts a CertificateTemplate resource to its proto representation.
func CertificateTemplateToProto(resource *alpha.CertificateTemplate) *alphapb.PrivatecaAlphaCertificateTemplate {
	p := &alphapb.PrivatecaAlphaCertificateTemplate{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		PredefinedValues:      PrivatecaAlphaCertificateTemplatePredefinedValuesToProto(resource.PredefinedValues),
		IdentityConstraints:   PrivatecaAlphaCertificateTemplateIdentityConstraintsToProto(resource.IdentityConstraints),
		PassthroughExtensions: PrivatecaAlphaCertificateTemplatePassthroughExtensionsToProto(resource.PassthroughExtensions),
		Description:           dcl.ValueOrEmptyString(resource.Description),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) applyCertificateTemplate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPrivatecaAlphaCertificateTemplateRequest) (*alphapb.PrivatecaAlphaCertificateTemplate, error) {
	p := ProtoToCertificateTemplate(request.GetResource())
	res, err := c.ApplyCertificateTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateTemplateToProto(res)
	return r, nil
}

// ApplyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) ApplyPrivatecaAlphaCertificateTemplate(ctx context.Context, request *alphapb.ApplyPrivatecaAlphaCertificateTemplateRequest) (*alphapb.PrivatecaAlphaCertificateTemplate, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCertificateTemplate(ctx, cl, request)
}

// DeleteCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Delete() method.
func (s *CertificateTemplateServer) DeletePrivatecaAlphaCertificateTemplate(ctx context.Context, request *alphapb.DeletePrivatecaAlphaCertificateTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))

}

// ListPrivatecaAlphaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplateList() method.
func (s *CertificateTemplateServer) ListPrivatecaAlphaCertificateTemplate(ctx context.Context, request *alphapb.ListPrivatecaAlphaCertificateTemplateRequest) (*alphapb.ListPrivatecaAlphaCertificateTemplateResponse, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateTemplate(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PrivatecaAlphaCertificateTemplate
	for _, r := range resources.Items {
		rp := CertificateTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListPrivatecaAlphaCertificateTemplateResponse{Items: protos}, nil
}

func createConfigCertificateTemplate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
