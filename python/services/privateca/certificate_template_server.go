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
	privatecapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/privateca_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca"
)

// Server implements the gRPC interface for CertificateTemplate.
type CertificateTemplateServer struct{}

// ProtoToCertificateTemplatePassthroughExtensionsKnownExtensionsEnum converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum from its proto representation.
func ProtoToPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(e privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum) *privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_name[int32(e)]; ok {
		e := privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(n[len("PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateTemplatePredefinedValues converts a CertificateTemplatePredefinedValues resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValues(p *privatecapb.PrivatecaCertificateTemplatePredefinedValues) *privateca.CertificateTemplatePredefinedValues {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValues{
		KeyUsage:  ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCertificateTemplatePredefinedValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCertificateTemplatePredefinedValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsage(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsage) *privateca.CertificateTemplatePredefinedValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesCaOptions converts a CertificateTemplatePredefinedValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesCaOptions(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesCaOptions) *privateca.CertificateTemplatePredefinedValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesPolicyIds converts a CertificateTemplatePredefinedValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesPolicyIds(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds) *privateca.CertificateTemplatePredefinedValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensions converts a CertificateTemplatePredefinedValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions) *privateca.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraints converts a CertificateTemplateIdentityConstraints resource from its proto representation.
func ProtoToPrivatecaCertificateTemplateIdentityConstraints(p *privatecapb.PrivatecaCertificateTemplateIdentityConstraints) *privateca.CertificateTemplateIdentityConstraints {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplateIdentityConstraints{
		CelExpression:                   ProtoToPrivatecaCertificateTemplateIdentityConstraintsCelExpression(p.GetCelExpression()),
		AllowSubjectPassthrough:         dcl.Bool(p.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.Bool(p.AllowSubjectAltNamesPassthrough),
	}
	return obj
}

// ProtoToCertificateTemplateIdentityConstraintsCelExpression converts a CertificateTemplateIdentityConstraintsCelExpression resource from its proto representation.
func ProtoToPrivatecaCertificateTemplateIdentityConstraintsCelExpression(p *privatecapb.PrivatecaCertificateTemplateIdentityConstraintsCelExpression) *privateca.CertificateTemplateIdentityConstraintsCelExpression {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.StringOrNil(p.Expression),
		Title:       dcl.StringOrNil(p.Title),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensions converts a CertificateTemplatePassthroughExtensions resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePassthroughExtensions(p *privatecapb.PrivatecaCertificateTemplatePassthroughExtensions) *privateca.CertificateTemplatePassthroughExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePassthroughExtensions{}
	for _, r := range p.GetKnownExtensions() {
		obj.KnownExtensions = append(obj.KnownExtensions, *ProtoToPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(r))
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateTemplatePassthroughExtensionsAdditionalExtensions converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(p *privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions) *privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateTemplate converts a CertificateTemplate resource from its proto representation.
func ProtoToCertificateTemplate(p *privatecapb.PrivatecaCertificateTemplate) *privateca.CertificateTemplate {
	obj := &privateca.CertificateTemplate{
		Name:                  dcl.StringOrNil(p.Name),
		PredefinedValues:      ProtoToPrivatecaCertificateTemplatePredefinedValues(p.GetPredefinedValues()),
		IdentityConstraints:   ProtoToPrivatecaCertificateTemplateIdentityConstraints(p.GetIdentityConstraints()),
		PassthroughExtensions: ProtoToPrivatecaCertificateTemplatePassthroughExtensions(p.GetPassthroughExtensions()),
		Description:           dcl.StringOrNil(p.Description),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
	}
	return obj
}

// CertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto converts a CertificateTemplatePassthroughExtensionsKnownExtensionsEnum enum to its proto representation.
func PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnumToProto(e *privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value["CertificateTemplatePassthroughExtensionsKnownExtensionsEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(v)
	}
	return privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(0)
}

// CertificateTemplatePredefinedValuesToProto converts a CertificateTemplatePredefinedValues resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesToProto(o *privateca.CertificateTemplatePredefinedValues) *privatecapb.PrivatecaCertificateTemplatePredefinedValues {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValues{
		KeyUsage:  PrivatecaCertificateTemplatePredefinedValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaCertificateTemplatePredefinedValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaCertificateTemplatePredefinedValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsage resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsage) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:     PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto converts a CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
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
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
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
func PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplatePredefinedValuesCaOptionsToProto converts a CertificateTemplatePredefinedValuesCaOptions resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesCaOptionsToProto(o *privateca.CertificateTemplatePredefinedValuesCaOptions) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateTemplatePredefinedValuesPolicyIdsToProto converts a CertificateTemplatePredefinedValuesPolicyIds resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesPolicyIdsToProto(o *privateca.CertificateTemplatePredefinedValuesPolicyIds) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensions resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsToProto(o *privateca.CertificateTemplatePredefinedValuesAdditionalExtensions) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto converts a CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdToProto(o *privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) *privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplateIdentityConstraintsToProto converts a CertificateTemplateIdentityConstraints resource to its proto representation.
func PrivatecaCertificateTemplateIdentityConstraintsToProto(o *privateca.CertificateTemplateIdentityConstraints) *privatecapb.PrivatecaCertificateTemplateIdentityConstraints {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplateIdentityConstraints{
		CelExpression:                   PrivatecaCertificateTemplateIdentityConstraintsCelExpressionToProto(o.CelExpression),
		AllowSubjectPassthrough:         dcl.ValueOrEmptyBool(o.AllowSubjectPassthrough),
		AllowSubjectAltNamesPassthrough: dcl.ValueOrEmptyBool(o.AllowSubjectAltNamesPassthrough),
	}
	return p
}

// CertificateTemplateIdentityConstraintsCelExpressionToProto converts a CertificateTemplateIdentityConstraintsCelExpression resource to its proto representation.
func PrivatecaCertificateTemplateIdentityConstraintsCelExpressionToProto(o *privateca.CertificateTemplateIdentityConstraintsCelExpression) *privatecapb.PrivatecaCertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplateIdentityConstraintsCelExpression{
		Expression:  dcl.ValueOrEmptyString(o.Expression),
		Title:       dcl.ValueOrEmptyString(o.Title),
		Description: dcl.ValueOrEmptyString(o.Description),
		Location:    dcl.ValueOrEmptyString(o.Location),
	}
	return p
}

// CertificateTemplatePassthroughExtensionsToProto converts a CertificateTemplatePassthroughExtensions resource to its proto representation.
func PrivatecaCertificateTemplatePassthroughExtensionsToProto(o *privateca.CertificateTemplatePassthroughExtensions) *privatecapb.PrivatecaCertificateTemplatePassthroughExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePassthroughExtensions{}
	for _, r := range o.KnownExtensions {
		p.KnownExtensions = append(p.KnownExtensions, privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum(privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsEnum_value[string(r)]))
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto converts a CertificateTemplatePassthroughExtensionsAdditionalExtensions resource to its proto representation.
func PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsToProto(o *privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions) *privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateTemplateToProto converts a CertificateTemplate resource to its proto representation.
func CertificateTemplateToProto(resource *privateca.CertificateTemplate) *privatecapb.PrivatecaCertificateTemplate {
	p := &privatecapb.PrivatecaCertificateTemplate{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		PredefinedValues:      PrivatecaCertificateTemplatePredefinedValuesToProto(resource.PredefinedValues),
		IdentityConstraints:   PrivatecaCertificateTemplateIdentityConstraintsToProto(resource.IdentityConstraints),
		PassthroughExtensions: PrivatecaCertificateTemplatePassthroughExtensionsToProto(resource.PassthroughExtensions),
		Description:           dcl.ValueOrEmptyString(resource.Description),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) applyCertificateTemplate(ctx context.Context, c *privateca.Client, request *privatecapb.ApplyPrivatecaCertificateTemplateRequest) (*privatecapb.PrivatecaCertificateTemplate, error) {
	p := ProtoToCertificateTemplate(request.GetResource())
	res, err := c.ApplyCertificateTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateTemplateToProto(res)
	return r, nil
}

// ApplyCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Apply() method.
func (s *CertificateTemplateServer) ApplyPrivatecaCertificateTemplate(ctx context.Context, request *privatecapb.ApplyPrivatecaCertificateTemplateRequest) (*privatecapb.PrivatecaCertificateTemplate, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCertificateTemplate(ctx, cl, request)
}

// DeleteCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplate Delete() method.
func (s *CertificateTemplateServer) DeletePrivatecaCertificateTemplate(ctx context.Context, request *privatecapb.DeletePrivatecaCertificateTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))

}

// ListPrivatecaCertificateTemplate handles the gRPC request by passing it to the underlying CertificateTemplateList() method.
func (s *CertificateTemplateServer) ListPrivatecaCertificateTemplate(ctx context.Context, request *privatecapb.ListPrivatecaCertificateTemplateRequest) (*privatecapb.ListPrivatecaCertificateTemplateResponse, error) {
	cl, err := createConfigCertificateTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateTemplate(ctx, ProtoToCertificateTemplate(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*privatecapb.PrivatecaCertificateTemplate
	for _, r := range resources.Items {
		rp := CertificateTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &privatecapb.ListPrivatecaCertificateTemplateResponse{Items: protos}, nil
}

func createConfigCertificateTemplate(ctx context.Context, service_account_file string) (*privateca.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return privateca.NewClient(conf), nil
}
