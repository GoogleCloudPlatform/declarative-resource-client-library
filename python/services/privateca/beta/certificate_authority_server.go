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

// ProtoToCertificateAuthorityConfigPublicKeyFormatEnum converts a CertificateAuthorityConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(e betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum) *beta.CertificateAuthorityConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityConfigPublicKeyFormatEnum(n[len("PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum"):])
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

// ProtoToCertificateAuthorityConfig converts a CertificateAuthorityConfig resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfig(p *betapb.PrivatecaBetaCertificateAuthorityConfig) *beta.CertificateAuthorityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfig{
		SubjectConfig: ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaBetaCertificateAuthorityConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKey(p.GetPublicKey()),
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

// ProtoToCertificateAuthorityConfigX509Config converts a CertificateAuthorityConfigX509Config resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509Config(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509Config) *beta.CertificateAuthorityConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509Config{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage) *beta.CertificateAuthorityConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *beta.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{
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

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigCaOptions converts a CertificateAuthorityConfigX509ConfigCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions) *beta.CertificateAuthorityConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigPolicyIds converts a CertificateAuthorityConfigX509ConfigPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds) *beta.CertificateAuthorityConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensions converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions) *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
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
		Key:    dcl.StringOrNil(p.Key),
		Format: ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(p.GetFormat()),
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
		X509Description:    ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p.GetCertFingerprint()),
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

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509Description converts a CertificateAuthorityCaCertificateDescriptionsX509Description resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description) *beta.CertificateAuthorityCaCertificateDescriptionsX509Description {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509Description{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{
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

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.ServerAuth),
		ClientAuth:      dcl.Bool(p.ClientAuth),
		CodeSigning:     dcl.Bool(p.CodeSigning),
		EmailProtection: dcl.Bool(p.EmailProtection),
		TimeStamping:    dcl.Bool(p.TimeStamping),
		OcspSigning:     dcl.Bool(p.OcspSigning),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.IsCa),
		MaxIssuerPathLength: dcl.Int64OrNil(p.MaxIssuerPathLength),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.Critical),
		Value:    dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
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

// ProtoToCertificateAuthorityAccessUrls converts a CertificateAuthorityAccessUrls resource from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityAccessUrls(p *betapb.PrivatecaBetaCertificateAuthorityAccessUrls) *beta.CertificateAuthorityAccessUrls {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.StringOrNil(p.CaCertificateAccessUrl),
	}
	for _, r := range p.GetCrlAccessUrls() {
		obj.CrlAccessUrls = append(obj.CrlAccessUrls, r)
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
		ExpireTime:        dcl.StringOrNil(p.GetExpireTime()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
		CaPool:            dcl.StringOrNil(p.CaPool),
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

// CertificateAuthorityConfigPublicKeyFormatEnumToProto converts a CertificateAuthorityConfigPublicKeyFormatEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnumToProto(e *beta.CertificateAuthorityConfigPublicKeyFormatEnum) betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum_value["CertificateAuthorityConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(0)
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

// CertificateAuthorityConfigToProto converts a CertificateAuthorityConfig resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigToProto(o *beta.CertificateAuthorityConfig) *betapb.PrivatecaBetaCertificateAuthorityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfig{
		SubjectConfig: PrivatecaBetaCertificateAuthorityConfigSubjectConfigToProto(o.SubjectConfig),
		X509Config:    PrivatecaBetaCertificateAuthorityConfigX509ConfigToProto(o.X509Config),
		PublicKey:     PrivatecaBetaCertificateAuthorityConfigPublicKeyToProto(o.PublicKey),
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

// CertificateAuthorityConfigX509ConfigToProto converts a CertificateAuthorityConfigX509Config resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigToProto(o *beta.CertificateAuthorityConfigX509Config) *betapb.PrivatecaBetaCertificateAuthorityConfigX509Config {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509Config{
		KeyUsage:  PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{
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

// CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityConfigX509ConfigCaOptionsToProto converts a CertificateAuthorityConfigX509ConfigCaOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o *beta.CertificateAuthorityConfigX509ConfigCaOptions) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateAuthorityConfigX509ConfigPolicyIdsToProto converts a CertificateAuthorityConfigX509ConfigPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(o *beta.CertificateAuthorityConfigX509ConfigPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(o *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
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
		Key:    dcl.ValueOrEmptyString(o.Key),
		Format: PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnumToProto(o.Format),
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
		X509Description:    PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o.X509Description),
		PublicKey:          PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o.PublicKey),
		SubjectKeyId:       PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o.SubjectKeyId),
		AuthorityKeyId:     PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o.AuthorityKeyId),
		CertFingerprint:    PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o.CertFingerprint),
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

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto converts a CertificateAuthorityCaCertificateDescriptionsX509Description resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509Description) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description{
		KeyUsage:  PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o.KeyUsage),
		CaOptions: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o.CaOptions),
	}
	for _, r := range o.PolicyIds {
		p.PolicyIds = append(p.PolicyIds, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(&r))
	}
	for _, r := range o.AiaOcspServers {
		p.AiaOcspServers = append(p.AiaOcspServers, r)
	}
	for _, r := range o.AdditionalExtensions {
		p.AdditionalExtensions = append(p.AdditionalExtensions, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(&r))
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{
		BaseKeyUsage:     PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage),
		ExtendedKeyUsage: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage),
	}
	for _, r := range o.UnknownExtendedKeyUsages {
		p.UnknownExtendedKeyUsages = append(p.UnknownExtendedKeyUsages, PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r))
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{
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

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.ValueOrEmptyBool(o.ServerAuth),
		ClientAuth:      dcl.ValueOrEmptyBool(o.ClientAuth),
		CodeSigning:     dcl.ValueOrEmptyBool(o.CodeSigning),
		EmailProtection: dcl.ValueOrEmptyBool(o.EmailProtection),
		TimeStamping:    dcl.ValueOrEmptyBool(o.TimeStamping),
		OcspSigning:     dcl.ValueOrEmptyBool(o.OcspSigning),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{
		IsCa:                dcl.ValueOrEmptyBool(o.IsCa),
		MaxIssuerPathLength: dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	for _, r := range o.ObjectIdPath {
		p.ObjectIdPath = append(p.ObjectIdPath, r)
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{
		ObjectId: PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId),
		Critical: dcl.ValueOrEmptyBool(o.Critical),
		Value:    dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId resource to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
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

// CertificateAuthorityAccessUrlsToProto converts a CertificateAuthorityAccessUrls resource to its proto representation.
func PrivatecaBetaCertificateAuthorityAccessUrlsToProto(o *beta.CertificateAuthorityAccessUrls) *betapb.PrivatecaBetaCertificateAuthorityAccessUrls {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.ValueOrEmptyString(o.CaCertificateAccessUrl),
	}
	for _, r := range o.CrlAccessUrls {
		p.CrlAccessUrls = append(p.CrlAccessUrls, r)
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
		ExpireTime:        dcl.ValueOrEmptyString(resource.ExpireTime),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
		CaPool:            dcl.ValueOrEmptyString(resource.CaPool),
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

	resources, err := cl.ListCertificateAuthority(ctx, request.Project, request.Location, request.CaPool)
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
