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

// Server implements the gRPC interface for CertificateAuthority.
type CertificateAuthorityServer struct{}

// ProtoToCertificateAuthorityTypeEnum converts a CertificateAuthorityTypeEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityTypeEnum(e betapb.PrivatecaBetaCertificateAuthorityTypeEnum) *beta.CertificateAuthorityTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityTypeEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityTypeEnum(n[len("PrivatecaBetaCertificateAuthorityTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfigPublicKeyTypeEnum converts a CertificateAuthorityConfigPublicKeyTypeEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum(e betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum) *beta.CertificateAuthorityConfigPublicKeyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityConfigPublicKeyTypeEnum(n[len("PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityKeySpecAlgorithmEnum converts a CertificateAuthorityKeySpecAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(e betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum) *beta.CertificateAuthorityKeySpecAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityKeySpecAlgorithmEnum(n[len("PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityTierEnum converts a CertificateAuthorityTierEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityTierEnum(e betapb.PrivatecaBetaCertificateAuthorityTierEnum) *beta.CertificateAuthorityTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityTierEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityTierEnum(n[len("PrivatecaBetaCertificateAuthorityTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityStateEnum converts a CertificateAuthorityStateEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityStateEnum(e betapb.PrivatecaBetaCertificateAuthorityStateEnum) *beta.CertificateAuthorityStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityStateEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityStateEnum(n[len("PrivatecaBetaCertificateAuthorityStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(e betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) *beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(n[len("PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(e betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum) *beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(n[len("PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfig converts a CertificateAuthorityConfig resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfig(p *betapb.PrivatecaBetaCertificateAuthorityConfig) *beta.CertificateAuthorityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfig{
		SubjectConfig:  ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfig(p.GetSubjectConfig()),
		PublicKey:      ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKey(p.GetPublicKey()),
		ReusableConfig: ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfig(p.GetReusableConfig()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfig converts a CertificateAuthorityConfigSubjectConfig resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfig(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfig) *beta.CertificateAuthorityConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfig{
		Subject:        ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject(p.GetSubject()),
		CommonName:     dcl.StringOrNil(p.CommonName),
		SubjectAltName: ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubject converts a CertificateAuthorityConfigSubjectConfigSubject resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject) *beta.CertificateAuthorityConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubject{
		CountryCode:        dcl.StringOrNil(p.CountryCode),
		Organization:       dcl.StringOrNil(p.Organization),
		OrganizationalUnit: dcl.StringOrNil(p.OrganizationalUnit),
		Locality:           dcl.StringOrNil(p.Locality),
		Province:           dcl.StringOrNil(p.Province),
		StreetAddress:      dcl.StringOrNil(p.StreetAddress),
		PostalCode:         dcl.StringOrNil(p.PostalCode),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltName converts a CertificateAuthorityConfigSubjectConfigSubjectAltName resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName) *beta.CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubjectAltName{}
	for _, r := range p.GetDnsNames() {
		obj.DnsNames = append(obj.DnsNames, r)
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	for _, r := range p.GetEmailAddresses() {
		obj.EmailAddresses = append(obj.EmailAddresses, r)
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	for _, r := range p.GetCustomSans() {
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigPublicKey converts a CertificateAuthorityConfigPublicKey resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKey(p *betapb.PrivatecaBetaCertificateAuthorityConfigPublicKey) *beta.CertificateAuthorityConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigPublicKey{
		Key:  dcl.StringOrNil(p.Key),
		Type: ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfig converts a CertificateAuthorityConfigReusableConfig resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfig(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfig) *beta.CertificateAuthorityConfigReusableConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfig{
		ReusableConfig:       dcl.StringOrNil(p.ReusableConfig),
		ReusableConfigValues: ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValues(p.GetReusableConfigValues()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValues converts a CertificateAuthorityConfigReusableConfigReusableConfigValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValues(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValues) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValues{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{
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

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityKeySpec converts a CertificateAuthorityKeySpec resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityKeySpec(p *betapb.PrivatecaBetaCertificateAuthorityKeySpec) *beta.CertificateAuthorityKeySpec {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityKeySpec{
		CloudKmsKeyVersion: dcl.StringOrNil(p.CloudKmsKeyVersion),
		Algorithm:          ProtoToPrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfig converts a CertificateAuthoritySubordinateConfig resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfig(p *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfig) *beta.CertificateAuthoritySubordinateConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthoritySubordinateConfig{
		CertificateAuthority: dcl.StringOrNil(p.CertificateAuthority),
		PemIssuerChain:       ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain(p.GetPemIssuerChain()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfigPemIssuerChain converts a CertificateAuthoritySubordinateConfigPemIssuerChain resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain(p *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain) *beta.CertificateAuthoritySubordinateConfigPemIssuerChain {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthoritySubordinateConfigPemIssuerChain{}
	for _, r := range p.GetPemCertificates() {
		obj.PemCertificates = append(obj.PemCertificates, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptions converts a CertificateAuthorityCaCertificateDescriptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions) *beta.CertificateAuthorityCaCertificateDescriptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptions{
		SubjectDescription: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p.GetSubjectDescription()),
		PublicKey:          ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p.GetCertFingerprint()),
		ConfigValues:       ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValues(p.GetConfigValues()),
	}
	for _, r := range p.GetCrlDistributionPoints() {
		obj.CrlDistributionPoints = append(obj.CrlDistributionPoints, r)
	}
	for _, r := range p.GetAiaIssuingCertificateUrls() {
		obj.AiaIssuingCertificateUrls = append(obj.AiaIssuingCertificateUrls, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescription converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescription resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescription{
		Subject:         ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.HexSerialNumber),
		Lifetime:        dcl.StringOrNil(p.Lifetime),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
		CommonName:      dcl.StringOrNil(p.CommonName),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{
		CommonName:         dcl.StringOrNil(p.CommonName),
		CountryCode:        dcl.StringOrNil(p.CountryCode),
		Organization:       dcl.StringOrNil(p.Organization),
		OrganizationalUnit: dcl.StringOrNil(p.OrganizationalUnit),
		Locality:           dcl.StringOrNil(p.Locality),
		Province:           dcl.StringOrNil(p.Province),
		StreetAddress:      dcl.StringOrNil(p.StreetAddress),
		PostalCode:         dcl.StringOrNil(p.PostalCode),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	for _, r := range p.GetDnsNames() {
		obj.DnsNames = append(obj.DnsNames, r)
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	for _, r := range p.GetEmailAddresses() {
		obj.EmailAddresses = append(obj.EmailAddresses, r)
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	for _, r := range p.GetCustomSans() {
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKey converts a CertificateAuthorityCaCertificateDescriptionsPublicKey resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey) *beta.CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsPublicKey{
		Key:    dcl.StringOrNil(p.Key),
		Format: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(p.GetFormat()),
		Type:   ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectKeyId converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{
		KeyId: dcl.StringOrNil(p.KeyId),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *beta.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.KeyId),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsCertFingerprint converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint) *beta.CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.Sha256Hash),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValues converts a CertificateAuthorityCaCertificateDescriptionsConfigValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValues(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValues) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValues{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{
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

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityAccessUrls converts a CertificateAuthorityAccessUrls resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityAccessUrls(p *betapb.PrivatecaBetaCertificateAuthorityAccessUrls) *beta.CertificateAuthorityAccessUrls {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.StringOrNil(p.CaCertificateAccessUrl),
		CrlAccessUrl:           dcl.StringOrNil(p.CrlAccessUrl),
	}
	for _, r := range p.GetCrlAccessUrls() {
		obj.CrlAccessUrls = append(obj.CrlAccessUrls, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicy converts a CertificateAuthorityCertificatePolicy resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicy(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicy) *beta.CertificateAuthorityCertificatePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicy{
		AllowedConfigList:     ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigList(p.GetAllowedConfigList()),
		OverwriteConfigValues: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValues(p.GetOverwriteConfigValues()),
		AllowedSans:           ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSans(p.GetAllowedSans()),
		MaximumLifetime:       dcl.StringOrNil(p.MaximumLifetime),
		AllowedIssuanceModes:  ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModes(p.GetAllowedIssuanceModes()),
	}
	for _, r := range p.GetAllowedLocationsAndOrganizations() {
		obj.AllowedLocationsAndOrganizations = append(obj.AllowedLocationsAndOrganizations, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(r))
	}
	for _, r := range p.GetAllowedCommonNames() {
		obj.AllowedCommonNames = append(obj.AllowedCommonNames, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigList converts a CertificateAuthorityCertificatePolicyAllowedConfigList resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigList(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigList) *beta.CertificateAuthorityCertificatePolicyAllowedConfigList {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigList{}
	for _, r := range p.GetAllowedConfigValues() {
		obj.AllowedConfigValues = append(obj.AllowedConfigValues, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{
		ReusableConfig:       dcl.StringOrNil(p.ReusableConfig),
		ReusableConfigValues: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(p.GetReusableConfigValues()),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{
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

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValues converts a CertificateAuthorityCertificatePolicyOverwriteConfigValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValues(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValues) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValues{
		ReusableConfig:       dcl.StringOrNil(p.ReusableConfig),
		ReusableConfigValues: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(p.GetReusableConfigValues()),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{
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

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations converts a CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) *beta.CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{
		CountryCode:        dcl.StringOrNil(p.CountryCode),
		Organization:       dcl.StringOrNil(p.Organization),
		OrganizationalUnit: dcl.StringOrNil(p.OrganizationalUnit),
		Locality:           dcl.StringOrNil(p.Locality),
		Province:           dcl.StringOrNil(p.Province),
		StreetAddress:      dcl.StringOrNil(p.StreetAddress),
		PostalCode:         dcl.StringOrNil(p.PostalCode),
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedSans converts a CertificateAuthorityCertificatePolicyAllowedSans resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSans(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSans) *beta.CertificateAuthorityCertificatePolicyAllowedSans {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedSans{
		AllowGlobbingDnsWildcards: dcl.Bool(p.AllowGlobbingDnsWildcards),
		AllowCustomSans:           dcl.Bool(p.AllowCustomSans),
	}
	for _, r := range p.GetAllowedDnsNames() {
		obj.AllowedDnsNames = append(obj.AllowedDnsNames, r)
	}
	for _, r := range p.GetAllowedUris() {
		obj.AllowedUris = append(obj.AllowedUris, r)
	}
	for _, r := range p.GetAllowedEmailAddresses() {
		obj.AllowedEmailAddresses = append(obj.AllowedEmailAddresses, r)
	}
	for _, r := range p.GetAllowedIps() {
		obj.AllowedIps = append(obj.AllowedIps, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCertificatePolicyAllowedIssuanceModes converts a CertificateAuthorityCertificatePolicyAllowedIssuanceModes resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModes(p *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModes) *beta.CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCertificatePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.Bool(p.AllowCsrBasedIssuance),
		AllowConfigBasedIssuance: dcl.Bool(p.AllowConfigBasedIssuance),
	}
	return obj
}

// ProtoToCertificateAuthorityIssuingOptions converts a CertificateAuthorityIssuingOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityIssuingOptions(p *betapb.PrivatecaBetaCertificateAuthorityIssuingOptions) *beta.CertificateAuthorityIssuingOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityIssuingOptions{
		IncludeCaCertUrl:    dcl.Bool(p.IncludeCaCertUrl),
		IncludeCrlAccessUrl: dcl.Bool(p.IncludeCrlAccessUrl),
	}
	return obj
}

// ProtoToCertificateAuthority converts a CertificateAuthority resource from its proto representation.
func ProtoToCertificateAuthority(p *betapb.PrivatecaBetaCertificateAuthority) *beta.CertificateAuthority {
	obj := &beta.CertificateAuthority{
		Name:              dcl.StringOrNil(p.Name),
		Type:              ProtoToPrivatecaBetaCertificateAuthorityTypeEnum(p.GetType()),
		Config:            ProtoToPrivatecaBetaCertificateAuthorityConfig(p.GetConfig()),
		Lifetime:          dcl.StringOrNil(p.Lifetime),
		KeySpec:           ProtoToPrivatecaBetaCertificateAuthorityKeySpec(p.GetKeySpec()),
		SubordinateConfig: ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfig(p.GetSubordinateConfig()),
		Tier:              ProtoToPrivatecaBetaCertificateAuthorityTierEnum(p.GetTier()),
		State:             ProtoToPrivatecaBetaCertificateAuthorityStateEnum(p.GetState()),
		GcsBucket:         dcl.StringOrNil(p.GcsBucket),
		AccessUrls:        ProtoToPrivatecaBetaCertificateAuthorityAccessUrls(p.GetAccessUrls()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:        dcl.StringOrNil(p.GetDeleteTime()),
		CertificatePolicy: ProtoToPrivatecaBetaCertificateAuthorityCertificatePolicy(p.GetCertificatePolicy()),
		IssuingOptions:    ProtoToPrivatecaBetaCertificateAuthorityIssuingOptions(p.GetIssuingOptions()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetPemCaCertificates() {
		obj.PemCaCertificates = append(obj.PemCaCertificates, r)
	}
	for _, r := range p.GetCaCertificateDescriptions() {
		obj.CaCertificateDescriptions = append(obj.CaCertificateDescriptions, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptions(r))
	}
	return obj
}

// CertificateAuthorityTypeEnumToProto converts a CertificateAuthorityTypeEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityTypeEnumToProto(e *beta.CertificateAuthorityTypeEnum) betapb.PrivatecaBetaCertificateAuthorityTypeEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityTypeEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityTypeEnum_value["CertificateAuthorityTypeEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityTypeEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityTypeEnum(0)
}

// CertificateAuthorityConfigPublicKeyTypeEnumToProto converts a CertificateAuthorityConfigPublicKeyTypeEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnumToProto(e *beta.CertificateAuthorityConfigPublicKeyTypeEnum) betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum_value["CertificateAuthorityConfigPublicKeyTypeEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum(0)
}

// CertificateAuthorityKeySpecAlgorithmEnumToProto converts a CertificateAuthorityKeySpecAlgorithmEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnumToProto(e *beta.CertificateAuthorityKeySpecAlgorithmEnum) betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum_value["CertificateAuthorityKeySpecAlgorithmEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(0)
}

// CertificateAuthorityTierEnumToProto converts a CertificateAuthorityTierEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityTierEnumToProto(e *beta.CertificateAuthorityTierEnum) betapb.PrivatecaBetaCertificateAuthorityTierEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityTierEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityTierEnum_value["CertificateAuthorityTierEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityTierEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityTierEnum(0)
}

// CertificateAuthorityStateEnumToProto converts a CertificateAuthorityStateEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityStateEnumToProto(e *beta.CertificateAuthorityStateEnum) betapb.PrivatecaBetaCertificateAuthorityStateEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityStateEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityStateEnum_value["CertificateAuthorityStateEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityStateEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityStateEnum(0)
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(e *beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_value["CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumToProto(e *beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum) betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum_value["CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(0)
}

// CertificateAuthorityConfigToProto converts a CertificateAuthorityConfig resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigToProto(o *beta.CertificateAuthorityConfig) *betapb.PrivatecaBetaCertificateAuthorityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfig{
		SubjectConfig:  PrivatecaBetaCertificateAuthorityConfigSubjectConfigToProto(o.SubjectConfig),
		PublicKey:      PrivatecaBetaCertificateAuthorityConfigPublicKeyToProto(o.PublicKey),
		ReusableConfig: PrivatecaBetaCertificateAuthorityConfigReusableConfigToProto(o.ReusableConfig),
	}
	return p
}

// CertificateAuthorityConfigSubjectConfigToProto converts a CertificateAuthorityConfigSubjectConfig resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigToProto(o *beta.CertificateAuthorityConfigSubjectConfig) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfig{
		Subject:        PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectToProto(o.Subject),
		CommonName:     dcl.ValueOrEmptyString(o.CommonName),
		SubjectAltName: PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName),
	}
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectToProto converts a CertificateAuthorityConfigSubjectConfigSubject resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubject) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject{
		CountryCode:        dcl.ValueOrEmptyString(o.CountryCode),
		Organization:       dcl.ValueOrEmptyString(o.Organization),
		OrganizationalUnit: dcl.ValueOrEmptyString(o.OrganizationalUnit),
		Locality:           dcl.ValueOrEmptyString(o.Locality),
		Province:           dcl.ValueOrEmptyString(o.Province),
		StreetAddress:      dcl.ValueOrEmptyString(o.StreetAddress),
		PostalCode:         dcl.ValueOrEmptyString(o.PostalCode),
	}
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltName resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubjectAltName) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName{}
	for _, r := range o.DnsNames {
		p.DnsNames = append(p.DnsNames, r)
	}
	for _, r := range o.Uris {
		p.Uris = append(p.Uris, r)
	}
	for _, r := range o.EmailAddresses {
		p.EmailAddresses = append(p.EmailAddresses, r)
	}
	for _, r := range o.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, r)
	}
	for _, r := range o.CustomSans {
		p.CustomSans = append(p.CustomSans, PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(&r))
	}
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{
		ObjectId: PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityConfigPublicKeyToProto converts a CertificateAuthorityConfigPublicKey resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigPublicKeyToProto(o *beta.CertificateAuthorityConfigPublicKey) *betapb.PrivatecaBetaCertificateAuthorityConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigPublicKey{
		Key:  dcl.ValueOrEmptyString(o.Key),
		Type: PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnumToProto(o.Type),
	}
	return p
}

// CertificateAuthorityConfigReusableConfigToProto converts a CertificateAuthorityConfigReusableConfig resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigToProto(o *beta.CertificateAuthorityConfigReusableConfig) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfig{
		ReusableConfig:       dcl.ValueOrEmptyString(o.ReusableConfig),
		ReusableConfigValues: PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesToProto(o.ReusableConfigValues),
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValues resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValues) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValues{
		KeyUsage:  PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{
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

// CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityKeySpecToProto converts a CertificateAuthorityKeySpec resource to its proto representation.
func PrivatecaBetaCertificateAuthorityKeySpecToProto(o *beta.CertificateAuthorityKeySpec) *betapb.PrivatecaBetaCertificateAuthorityKeySpec {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityKeySpec{
		CloudKmsKeyVersion: dcl.ValueOrEmptyString(o.CloudKmsKeyVersion),
		Algorithm:          PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnumToProto(o.Algorithm),
	}
	return p
}

// CertificateAuthoritySubordinateConfigToProto converts a CertificateAuthoritySubordinateConfig resource to its proto representation.
func PrivatecaBetaCertificateAuthoritySubordinateConfigToProto(o *beta.CertificateAuthoritySubordinateConfig) *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthoritySubordinateConfig{
		CertificateAuthority: dcl.ValueOrEmptyString(o.CertificateAuthority),
		PemIssuerChain:       PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o.PemIssuerChain),
	}
	return p
}

// CertificateAuthoritySubordinateConfigPemIssuerChainToProto converts a CertificateAuthoritySubordinateConfigPemIssuerChain resource to its proto representation.
func PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o *beta.CertificateAuthoritySubordinateConfigPemIssuerChain) *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain{}
	for _, r := range o.PemCertificates {
		p.PemCertificates = append(p.PemCertificates, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsToProto converts a CertificateAuthorityCaCertificateDescriptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions{
		SubjectDescription: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o.SubjectDescription),
		PublicKey:          PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o.PublicKey),
		SubjectKeyId:       PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o.SubjectKeyId),
		AuthorityKeyId:     PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o.AuthorityKeyId),
		CertFingerprint:    PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o.CertFingerprint),
		ConfigValues:       PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesToProto(o.ConfigValues),
	}
	for _, r := range o.CrlDistributionPoints {
		p.CrlDistributionPoints = append(p.CrlDistributionPoints, r)
	}
	for _, r := range o.AiaIssuingCertificateUrls {
		p.AiaIssuingCertificateUrls = append(p.AiaIssuingCertificateUrls, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescription resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescription) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription{
		Subject:         PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o.Subject),
		SubjectAltName:  PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName),
		HexSerialNumber: dcl.ValueOrEmptyString(o.HexSerialNumber),
		Lifetime:        dcl.ValueOrEmptyString(o.Lifetime),
		NotBeforeTime:   dcl.ValueOrEmptyString(o.NotBeforeTime),
		NotAfterTime:    dcl.ValueOrEmptyString(o.NotAfterTime),
		CommonName:      dcl.ValueOrEmptyString(o.CommonName),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{
		CommonName:         dcl.ValueOrEmptyString(o.CommonName),
		CountryCode:        dcl.ValueOrEmptyString(o.CountryCode),
		Organization:       dcl.ValueOrEmptyString(o.Organization),
		OrganizationalUnit: dcl.ValueOrEmptyString(o.OrganizationalUnit),
		Locality:           dcl.ValueOrEmptyString(o.Locality),
		Province:           dcl.ValueOrEmptyString(o.Province),
		StreetAddress:      dcl.ValueOrEmptyString(o.StreetAddress),
		PostalCode:         dcl.ValueOrEmptyString(o.PostalCode),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	for _, r := range o.DnsNames {
		p.DnsNames = append(p.DnsNames, r)
	}
	for _, r := range o.Uris {
		p.Uris = append(p.Uris, r)
	}
	for _, r := range o.EmailAddresses {
		p.EmailAddresses = append(p.EmailAddresses, r)
	}
	for _, r := range o.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, r)
	}
	for _, r := range o.CustomSans {
		p.CustomSans = append(p.CustomSans, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(&r))
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKey resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsPublicKey) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Format: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(o.Format),
		Type:   PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumToProto(o.Type),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId{
		KeyId: dcl.ValueOrEmptyString(o.KeyId),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{
		KeyId: dcl.ValueOrEmptyString(o.KeyId),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsCertFingerprint) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint{
		Sha256Hash: dcl.ValueOrEmptyString(o.Sha256Hash),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValues resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValues) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValues{
		KeyUsage:  PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{
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

// CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityAccessUrlsToProto converts a CertificateAuthorityAccessUrls resource to its proto representation.
func PrivatecaBetaCertificateAuthorityAccessUrlsToProto(o *beta.CertificateAuthorityAccessUrls) *betapb.PrivatecaBetaCertificateAuthorityAccessUrls {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.ValueOrEmptyString(o.CaCertificateAccessUrl),
		CrlAccessUrl:           dcl.ValueOrEmptyString(o.CrlAccessUrl),
	}
	for _, r := range o.CrlAccessUrls {
		p.CrlAccessUrls = append(p.CrlAccessUrls, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyToProto converts a CertificateAuthorityCertificatePolicy resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyToProto(o *beta.CertificateAuthorityCertificatePolicy) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicy{
		AllowedConfigList:     PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListToProto(o.AllowedConfigList),
		OverwriteConfigValues: PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesToProto(o.OverwriteConfigValues),
		AllowedSans:           PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSansToProto(o.AllowedSans),
		MaximumLifetime:       dcl.ValueOrEmptyString(o.MaximumLifetime),
		AllowedIssuanceModes:  PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModesToProto(o.AllowedIssuanceModes),
	}
	for _, r := range o.AllowedLocationsAndOrganizations {
		p.AllowedLocationsAndOrganizations = append(p.AllowedLocationsAndOrganizations, PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsToProto(&r))
	}
	for _, r := range o.AllowedCommonNames {
		p.AllowedCommonNames = append(p.AllowedCommonNames, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigList resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigList) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigList {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigList{}
	for _, r := range o.AllowedConfigValues {
		p.AllowedConfigValues = append(p.AllowedConfigValues, PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesToProto(&r))
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{
		ReusableConfig:       dcl.ValueOrEmptyString(o.ReusableConfig),
		ReusableConfigValues: PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesToProto(o.ReusableConfigValues),
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{
		KeyUsage:  PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{
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

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValues resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValues) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValues{
		ReusableConfig:       dcl.ValueOrEmptyString(o.ReusableConfig),
		ReusableConfigValues: PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesToProto(o.ReusableConfigValues),
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{
		KeyUsage:  PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{
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

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsToProto converts a CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{
		CountryCode:        dcl.ValueOrEmptyString(o.CountryCode),
		Organization:       dcl.ValueOrEmptyString(o.Organization),
		OrganizationalUnit: dcl.ValueOrEmptyString(o.OrganizationalUnit),
		Locality:           dcl.ValueOrEmptyString(o.Locality),
		Province:           dcl.ValueOrEmptyString(o.Province),
		StreetAddress:      dcl.ValueOrEmptyString(o.StreetAddress),
		PostalCode:         dcl.ValueOrEmptyString(o.PostalCode),
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedSansToProto converts a CertificateAuthorityCertificatePolicyAllowedSans resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSansToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedSans) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSans {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSans{
		AllowGlobbingDnsWildcards: dcl.ValueOrEmptyBool(o.AllowGlobbingDnsWildcards),
		AllowCustomSans:           dcl.ValueOrEmptyBool(o.AllowCustomSans),
	}
	for _, r := range o.AllowedDnsNames {
		p.AllowedDnsNames = append(p.AllowedDnsNames, r)
	}
	for _, r := range o.AllowedUris {
		p.AllowedUris = append(p.AllowedUris, r)
	}
	for _, r := range o.AllowedEmailAddresses {
		p.AllowedEmailAddresses = append(p.AllowedEmailAddresses, r)
	}
	for _, r := range o.AllowedIps {
		p.AllowedIps = append(p.AllowedIps, r)
	}
	return p
}

// CertificateAuthorityCertificatePolicyAllowedIssuanceModesToProto converts a CertificateAuthorityCertificatePolicyAllowedIssuanceModes resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModesToProto(o *beta.CertificateAuthorityCertificatePolicyAllowedIssuanceModes) *betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModes{
		AllowCsrBasedIssuance:    dcl.ValueOrEmptyBool(o.AllowCsrBasedIssuance),
		AllowConfigBasedIssuance: dcl.ValueOrEmptyBool(o.AllowConfigBasedIssuance),
	}
	return p
}

// CertificateAuthorityIssuingOptionsToProto converts a CertificateAuthorityIssuingOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityIssuingOptionsToProto(o *beta.CertificateAuthorityIssuingOptions) *betapb.PrivatecaBetaCertificateAuthorityIssuingOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityIssuingOptions{
		IncludeCaCertUrl:    dcl.ValueOrEmptyBool(o.IncludeCaCertUrl),
		IncludeCrlAccessUrl: dcl.ValueOrEmptyBool(o.IncludeCrlAccessUrl),
	}
	return p
}

// CertificateAuthorityToProto converts a CertificateAuthority resource to its proto representation.
func CertificateAuthorityToProto(resource *beta.CertificateAuthority) *betapb.PrivatecaBetaCertificateAuthority {
	p := &betapb.PrivatecaBetaCertificateAuthority{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Type:              PrivatecaBetaCertificateAuthorityTypeEnumToProto(resource.Type),
		Config:            PrivatecaBetaCertificateAuthorityConfigToProto(resource.Config),
		Lifetime:          dcl.ValueOrEmptyString(resource.Lifetime),
		KeySpec:           PrivatecaBetaCertificateAuthorityKeySpecToProto(resource.KeySpec),
		SubordinateConfig: PrivatecaBetaCertificateAuthoritySubordinateConfigToProto(resource.SubordinateConfig),
		Tier:              PrivatecaBetaCertificateAuthorityTierEnumToProto(resource.Tier),
		State:             PrivatecaBetaCertificateAuthorityStateEnumToProto(resource.State),
		GcsBucket:         dcl.ValueOrEmptyString(resource.GcsBucket),
		AccessUrls:        PrivatecaBetaCertificateAuthorityAccessUrlsToProto(resource.AccessUrls),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:        dcl.ValueOrEmptyString(resource.DeleteTime),
		CertificatePolicy: PrivatecaBetaCertificateAuthorityCertificatePolicyToProto(resource.CertificatePolicy),
		IssuingOptions:    PrivatecaBetaCertificateAuthorityIssuingOptionsToProto(resource.IssuingOptions),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.PemCaCertificates {
		p.PemCaCertificates = append(p.PemCaCertificates, r)
	}
	for _, r := range resource.CaCertificateDescriptions {
		p.CaCertificateDescriptions = append(p.CaCertificateDescriptions, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsToProto(&r))
	}

	return p
}

// ApplyCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Apply() method.
func (s *CertificateAuthorityServer) applyCertificateAuthority(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCertificateAuthorityRequest) (*betapb.PrivatecaBetaCertificateAuthority, error) {
	p := ProtoToCertificateAuthority(request.GetResource())
	res, err := c.ApplyCertificateAuthority(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateAuthorityToProto(res)
	return r, nil
}

// ApplyCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Apply() method.
func (s *CertificateAuthorityServer) ApplyPrivatecaBetaCertificateAuthority(ctx context.Context, request *betapb.ApplyPrivatecaBetaCertificateAuthorityRequest) (*betapb.PrivatecaBetaCertificateAuthority, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCertificateAuthority(ctx, cl, request)
}

// DeleteCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Delete() method.
func (s *CertificateAuthorityServer) DeletePrivatecaBetaCertificateAuthority(ctx context.Context, request *betapb.DeletePrivatecaBetaCertificateAuthorityRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateAuthority(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateAuthority(ctx, ProtoToCertificateAuthority(request.GetResource()))

}

// ListPrivatecaBetaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthorityList() method.
func (s *CertificateAuthorityServer) ListPrivatecaBetaCertificateAuthority(ctx context.Context, request *betapb.ListPrivatecaBetaCertificateAuthorityRequest) (*betapb.ListPrivatecaBetaCertificateAuthorityResponse, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateAuthority(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCertificateAuthority
	for _, r := range resources.Items {
		rp := CertificateAuthorityToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListPrivatecaBetaCertificateAuthorityResponse{Items: protos}, nil
}

func createConfigCertificateAuthority(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
