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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/alpha/cloudkms_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/alpha"
)

// Server implements the gRPC interface for CryptoKey.
type CryptoKeyServer struct{}

// ProtoToCryptoKeyPrimaryStateEnum converts a CryptoKeyPrimaryStateEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryStateEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum) *alpha.CryptoKeyPrimaryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryStateEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryProtectionLevelEnum converts a CryptoKeyPrimaryProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum) *alpha.CryptoKeyPrimaryProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryProtectionLevelEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAlgorithmEnum converts a CryptoKeyPrimaryAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum) *alpha.CryptoKeyPrimaryAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryAlgorithmEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAttestationFormatEnum converts a CryptoKeyPrimaryAttestationFormatEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum) *alpha.CryptoKeyPrimaryAttestationFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryAttestationFormatEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPurposeEnum converts a CryptoKeyPurposeEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPurposeEnum(e alphapb.CloudkmsAlphaCryptoKeyPurposeEnum) *alpha.CryptoKeyPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPurposeEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPurposeEnum(n[len("CloudkmsAlphaCryptoKeyPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateProtectionLevelEnum converts a CryptoKeyVersionTemplateProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(e alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum) *alpha.CryptoKeyVersionTemplateProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyVersionTemplateProtectionLevelEnum(n[len("CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateAlgorithmEnum converts a CryptoKeyVersionTemplateAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(e alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum) *alpha.CryptoKeyVersionTemplateAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyVersionTemplateAlgorithmEnum(n[len("CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimary converts a CryptoKeyPrimary resource from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimary(p *alphapb.CloudkmsAlphaCryptoKeyPrimary) *alpha.CryptoKeyPrimary {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimary{
		Name:                           dcl.StringOrNil(p.Name),
		State:                          ProtoToCloudkmsAlphaCryptoKeyPrimaryStateEnum(p.GetState()),
		ProtectionLevel:                ProtoToCloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:                      ProtoToCloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(p.GetAlgorithm()),
		Attestation:                    ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestation(p.GetAttestation()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		GenerateTime:                   dcl.StringOrNil(p.GetGenerateTime()),
		DestroyTime:                    dcl.StringOrNil(p.GetDestroyTime()),
		DestroyEventTime:               dcl.StringOrNil(p.GetDestroyEventTime()),
		ImportJob:                      dcl.StringOrNil(p.ImportJob),
		ImportTime:                     dcl.StringOrNil(p.GetImportTime()),
		ImportFailureReason:            dcl.StringOrNil(p.ImportFailureReason),
		ExternalProtectionLevelOptions: ProtoToCloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions(p.GetExternalProtectionLevelOptions()),
		ReimportEligible:               dcl.Bool(p.ReimportEligible),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestation converts a CryptoKeyPrimaryAttestation resource from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestation(p *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestation) *alpha.CryptoKeyPrimaryAttestation {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimaryAttestation{
		Format:     ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(p.GetFormat()),
		Content:    dcl.StringOrNil(p.Content),
		CertChains: ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationCertChains(p.GetCertChains()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestationCertChains converts a CryptoKeyPrimaryAttestationCertChains resource from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationCertChains(p *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationCertChains) *alpha.CryptoKeyPrimaryAttestationCertChains {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimaryAttestationCertChains{}
	for _, r := range p.GetCaviumCerts() {
		obj.CaviumCerts = append(obj.CaviumCerts, r)
	}
	for _, r := range p.GetGoogleCardCerts() {
		obj.GoogleCardCerts = append(obj.GoogleCardCerts, r)
	}
	for _, r := range p.GetGooglePartitionCerts() {
		obj.GooglePartitionCerts = append(obj.GooglePartitionCerts, r)
	}
	return obj
}

// ProtoToCryptoKeyPrimaryExternalProtectionLevelOptions converts a CryptoKeyPrimaryExternalProtectionLevelOptions resource from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions(p *alphapb.CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions) *alpha.CryptoKeyPrimaryExternalProtectionLevelOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.StringOrNil(p.ExternalKeyUri),
	}
	return obj
}

// ProtoToCryptoKeyVersionTemplate converts a CryptoKeyVersionTemplate resource from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyVersionTemplate(p *alphapb.CloudkmsAlphaCryptoKeyVersionTemplate) *alpha.CryptoKeyVersionTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyVersionTemplate{
		ProtectionLevel: ProtoToCloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:       ProtoToCloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCryptoKey converts a CryptoKey resource from its proto representation.
func ProtoToCryptoKey(p *alphapb.CloudkmsAlphaCryptoKey) *alpha.CryptoKey {
	obj := &alpha.CryptoKey{
		Name:                     dcl.StringOrNil(p.Name),
		Primary:                  ProtoToCloudkmsAlphaCryptoKeyPrimary(p.GetPrimary()),
		Purpose:                  ProtoToCloudkmsAlphaCryptoKeyPurposeEnum(p.GetPurpose()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		NextRotationTime:         dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:           dcl.StringOrNil(p.RotationPeriod),
		VersionTemplate:          ProtoToCloudkmsAlphaCryptoKeyVersionTemplate(p.GetVersionTemplate()),
		ImportOnly:               dcl.Bool(p.ImportOnly),
		DestroyScheduledDuration: dcl.StringOrNil(p.DestroyScheduledDuration),
		Project:                  dcl.StringOrNil(p.Project),
		Location:                 dcl.StringOrNil(p.Location),
		KeyRing:                  dcl.StringOrNil(p.KeyRing),
	}
	return obj
}

// CryptoKeyPrimaryStateEnumToProto converts a CryptoKeyPrimaryStateEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryStateEnumToProto(e *alpha.CryptoKeyPrimaryStateEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum_value["CryptoKeyPrimaryStateEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum(0)
}

// CryptoKeyPrimaryProtectionLevelEnumToProto converts a CryptoKeyPrimaryProtectionLevelEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnumToProto(e *alpha.CryptoKeyPrimaryProtectionLevelEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum_value["CryptoKeyPrimaryProtectionLevelEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(0)
}

// CryptoKeyPrimaryAlgorithmEnumToProto converts a CryptoKeyPrimaryAlgorithmEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnumToProto(e *alpha.CryptoKeyPrimaryAlgorithmEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum_value["CryptoKeyPrimaryAlgorithmEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(0)
}

// CryptoKeyPrimaryAttestationFormatEnumToProto converts a CryptoKeyPrimaryAttestationFormatEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnumToProto(e *alpha.CryptoKeyPrimaryAttestationFormatEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum_value["CryptoKeyPrimaryAttestationFormatEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(0)
}

// CryptoKeyPurposeEnumToProto converts a CryptoKeyPurposeEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPurposeEnumToProto(e *alpha.CryptoKeyPurposeEnum) alphapb.CloudkmsAlphaCryptoKeyPurposeEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPurposeEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPurposeEnum_value["CryptoKeyPurposeEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPurposeEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPurposeEnum(0)
}

// CryptoKeyVersionTemplateProtectionLevelEnumToProto converts a CryptoKeyVersionTemplateProtectionLevelEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnumToProto(e *alpha.CryptoKeyVersionTemplateProtectionLevelEnum) alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum_value["CryptoKeyVersionTemplateProtectionLevelEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(0)
}

// CryptoKeyVersionTemplateAlgorithmEnumToProto converts a CryptoKeyVersionTemplateAlgorithmEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnumToProto(e *alpha.CryptoKeyVersionTemplateAlgorithmEnum) alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum_value["CryptoKeyVersionTemplateAlgorithmEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(0)
}

// CryptoKeyPrimaryToProto converts a CryptoKeyPrimary resource to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryToProto(o *alpha.CryptoKeyPrimary) *alphapb.CloudkmsAlphaCryptoKeyPrimary {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimary{
		Name:                           dcl.ValueOrEmptyString(o.Name),
		State:                          CloudkmsAlphaCryptoKeyPrimaryStateEnumToProto(o.State),
		ProtectionLevel:                CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnumToProto(o.ProtectionLevel),
		Algorithm:                      CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnumToProto(o.Algorithm),
		Attestation:                    CloudkmsAlphaCryptoKeyPrimaryAttestationToProto(o.Attestation),
		CreateTime:                     dcl.ValueOrEmptyString(o.CreateTime),
		GenerateTime:                   dcl.ValueOrEmptyString(o.GenerateTime),
		DestroyTime:                    dcl.ValueOrEmptyString(o.DestroyTime),
		DestroyEventTime:               dcl.ValueOrEmptyString(o.DestroyEventTime),
		ImportJob:                      dcl.ValueOrEmptyString(o.ImportJob),
		ImportTime:                     dcl.ValueOrEmptyString(o.ImportTime),
		ImportFailureReason:            dcl.ValueOrEmptyString(o.ImportFailureReason),
		ExternalProtectionLevelOptions: CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o.ExternalProtectionLevelOptions),
		ReimportEligible:               dcl.ValueOrEmptyBool(o.ReimportEligible),
	}
	return p
}

// CryptoKeyPrimaryAttestationToProto converts a CryptoKeyPrimaryAttestation resource to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAttestationToProto(o *alpha.CryptoKeyPrimaryAttestation) *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestation {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestation{
		Format:     CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnumToProto(o.Format),
		Content:    dcl.ValueOrEmptyString(o.Content),
		CertChains: CloudkmsAlphaCryptoKeyPrimaryAttestationCertChainsToProto(o.CertChains),
	}
	return p
}

// CryptoKeyPrimaryAttestationCertChainsToProto converts a CryptoKeyPrimaryAttestationCertChains resource to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAttestationCertChainsToProto(o *alpha.CryptoKeyPrimaryAttestationCertChains) *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationCertChains {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationCertChains{}
	for _, r := range o.CaviumCerts {
		p.CaviumCerts = append(p.CaviumCerts, r)
	}
	for _, r := range o.GoogleCardCerts {
		p.GoogleCardCerts = append(p.GoogleCardCerts, r)
	}
	for _, r := range o.GooglePartitionCerts {
		p.GooglePartitionCerts = append(p.GooglePartitionCerts, r)
	}
	return p
}

// CryptoKeyPrimaryExternalProtectionLevelOptionsToProto converts a CryptoKeyPrimaryExternalProtectionLevelOptions resource to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o *alpha.CryptoKeyPrimaryExternalProtectionLevelOptions) *alphapb.CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.ValueOrEmptyString(o.ExternalKeyUri),
	}
	return p
}

// CryptoKeyVersionTemplateToProto converts a CryptoKeyVersionTemplate resource to its proto representation.
func CloudkmsAlphaCryptoKeyVersionTemplateToProto(o *alpha.CryptoKeyVersionTemplate) *alphapb.CloudkmsAlphaCryptoKeyVersionTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyVersionTemplate{
		ProtectionLevel: CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnumToProto(o.ProtectionLevel),
		Algorithm:       CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnumToProto(o.Algorithm),
	}
	return p
}

// CryptoKeyToProto converts a CryptoKey resource to its proto representation.
func CryptoKeyToProto(resource *alpha.CryptoKey) *alphapb.CloudkmsAlphaCryptoKey {
	p := &alphapb.CloudkmsAlphaCryptoKey{
		Name:                     dcl.ValueOrEmptyString(resource.Name),
		Primary:                  CloudkmsAlphaCryptoKeyPrimaryToProto(resource.Primary),
		Purpose:                  CloudkmsAlphaCryptoKeyPurposeEnumToProto(resource.Purpose),
		CreateTime:               dcl.ValueOrEmptyString(resource.CreateTime),
		NextRotationTime:         dcl.ValueOrEmptyString(resource.NextRotationTime),
		RotationPeriod:           dcl.ValueOrEmptyString(resource.RotationPeriod),
		VersionTemplate:          CloudkmsAlphaCryptoKeyVersionTemplateToProto(resource.VersionTemplate),
		ImportOnly:               dcl.ValueOrEmptyBool(resource.ImportOnly),
		DestroyScheduledDuration: dcl.ValueOrEmptyString(resource.DestroyScheduledDuration),
		Project:                  dcl.ValueOrEmptyString(resource.Project),
		Location:                 dcl.ValueOrEmptyString(resource.Location),
		KeyRing:                  dcl.ValueOrEmptyString(resource.KeyRing),
	}

	return p
}

// ApplyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) applyCryptoKey(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudkmsAlphaCryptoKeyRequest) (*alphapb.CloudkmsAlphaCryptoKey, error) {
	p := ProtoToCryptoKey(request.GetResource())
	res, err := c.ApplyCryptoKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CryptoKeyToProto(res)
	return r, nil
}

// ApplyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) ApplyCloudkmsAlphaCryptoKey(ctx context.Context, request *alphapb.ApplyCloudkmsAlphaCryptoKeyRequest) (*alphapb.CloudkmsAlphaCryptoKey, error) {
	cl, err := createConfigCryptoKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCryptoKey(ctx, cl, request)
}

// DeleteCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Delete() method.
func (s *CryptoKeyServer) DeleteCloudkmsAlphaCryptoKey(ctx context.Context, request *alphapb.DeleteCloudkmsAlphaCryptoKeyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for CryptoKey")

}

// ListCloudkmsAlphaCryptoKey handles the gRPC request by passing it to the underlying CryptoKeyList() method.
func (s *CryptoKeyServer) ListCloudkmsAlphaCryptoKey(ctx context.Context, request *alphapb.ListCloudkmsAlphaCryptoKeyRequest) (*alphapb.ListCloudkmsAlphaCryptoKeyResponse, error) {
	cl, err := createConfigCryptoKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCryptoKey(ctx, ProtoToCryptoKey(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudkmsAlphaCryptoKey
	for _, r := range resources.Items {
		rp := CryptoKeyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListCloudkmsAlphaCryptoKeyResponse{Items: protos}, nil
}

func createConfigCryptoKey(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
