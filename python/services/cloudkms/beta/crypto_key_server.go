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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/beta/cloudkms_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/beta"
)

// Server implements the gRPC interface for CryptoKey.
type CryptoKeyServer struct{}

// ProtoToCryptoKeyPrimaryStateEnum converts a CryptoKeyPrimaryStateEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimaryStateEnum(e betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum) *beta.CryptoKeyPrimaryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyPrimaryStateEnum(n[len("CloudkmsBetaCryptoKeyPrimaryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryProtectionLevelEnum converts a CryptoKeyPrimaryProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum(e betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum) *beta.CryptoKeyPrimaryProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyPrimaryProtectionLevelEnum(n[len("CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAlgorithmEnum converts a CryptoKeyPrimaryAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimaryAlgorithmEnum(e betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum) *beta.CryptoKeyPrimaryAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyPrimaryAlgorithmEnum(n[len("CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAttestationFormatEnum converts a CryptoKeyPrimaryAttestationFormatEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum(e betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum) *beta.CryptoKeyPrimaryAttestationFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyPrimaryAttestationFormatEnum(n[len("CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPurposeEnum converts a CryptoKeyPurposeEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPurposeEnum(e betapb.CloudkmsBetaCryptoKeyPurposeEnum) *beta.CryptoKeyPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyPurposeEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyPurposeEnum(n[len("CloudkmsBetaCryptoKeyPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateProtectionLevelEnum converts a CryptoKeyVersionTemplateProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum(e betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum) *beta.CryptoKeyVersionTemplateProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyVersionTemplateProtectionLevelEnum(n[len("CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateAlgorithmEnum converts a CryptoKeyVersionTemplateAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum(e betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum) *beta.CryptoKeyVersionTemplateAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum_name[int32(e)]; ok {
		e := beta.CryptoKeyVersionTemplateAlgorithmEnum(n[len("CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimary converts a CryptoKeyPrimary resource from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimary(p *betapb.CloudkmsBetaCryptoKeyPrimary) *beta.CryptoKeyPrimary {
	if p == nil {
		return nil
	}
	obj := &beta.CryptoKeyPrimary{
		Name:                           dcl.StringOrNil(p.Name),
		State:                          ProtoToCloudkmsBetaCryptoKeyPrimaryStateEnum(p.GetState()),
		ProtectionLevel:                ProtoToCloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:                      ProtoToCloudkmsBetaCryptoKeyPrimaryAlgorithmEnum(p.GetAlgorithm()),
		Attestation:                    ProtoToCloudkmsBetaCryptoKeyPrimaryAttestation(p.GetAttestation()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		GenerateTime:                   dcl.StringOrNil(p.GetGenerateTime()),
		DestroyTime:                    dcl.StringOrNil(p.GetDestroyTime()),
		DestroyEventTime:               dcl.StringOrNil(p.GetDestroyEventTime()),
		ImportJob:                      dcl.StringOrNil(p.ImportJob),
		ImportTime:                     dcl.StringOrNil(p.GetImportTime()),
		ImportFailureReason:            dcl.StringOrNil(p.ImportFailureReason),
		ExternalProtectionLevelOptions: ProtoToCloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptions(p.GetExternalProtectionLevelOptions()),
		ReimportEligible:               dcl.Bool(p.ReimportEligible),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestation converts a CryptoKeyPrimaryAttestation resource from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimaryAttestation(p *betapb.CloudkmsBetaCryptoKeyPrimaryAttestation) *beta.CryptoKeyPrimaryAttestation {
	if p == nil {
		return nil
	}
	obj := &beta.CryptoKeyPrimaryAttestation{
		Format:     ProtoToCloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum(p.GetFormat()),
		Content:    dcl.StringOrNil(p.Content),
		CertChains: ProtoToCloudkmsBetaCryptoKeyPrimaryAttestationCertChains(p.GetCertChains()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestationCertChains converts a CryptoKeyPrimaryAttestationCertChains resource from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyPrimaryAttestationCertChains(p *betapb.CloudkmsBetaCryptoKeyPrimaryAttestationCertChains) *beta.CryptoKeyPrimaryAttestationCertChains {
	if p == nil {
		return nil
	}
	obj := &beta.CryptoKeyPrimaryAttestationCertChains{}
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
func ProtoToCloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptions(p *betapb.CloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptions) *beta.CryptoKeyPrimaryExternalProtectionLevelOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.StringOrNil(p.ExternalKeyUri),
	}
	return obj
}

// ProtoToCryptoKeyVersionTemplate converts a CryptoKeyVersionTemplate resource from its proto representation.
func ProtoToCloudkmsBetaCryptoKeyVersionTemplate(p *betapb.CloudkmsBetaCryptoKeyVersionTemplate) *beta.CryptoKeyVersionTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.CryptoKeyVersionTemplate{
		ProtectionLevel: ProtoToCloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:       ProtoToCloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCryptoKey converts a CryptoKey resource from its proto representation.
func ProtoToCryptoKey(p *betapb.CloudkmsBetaCryptoKey) *beta.CryptoKey {
	obj := &beta.CryptoKey{
		Name:                     dcl.StringOrNil(p.Name),
		Primary:                  ProtoToCloudkmsBetaCryptoKeyPrimary(p.GetPrimary()),
		Purpose:                  ProtoToCloudkmsBetaCryptoKeyPurposeEnum(p.GetPurpose()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		NextRotationTime:         dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:           dcl.StringOrNil(p.RotationPeriod),
		VersionTemplate:          ProtoToCloudkmsBetaCryptoKeyVersionTemplate(p.GetVersionTemplate()),
		ImportOnly:               dcl.Bool(p.ImportOnly),
		DestroyScheduledDuration: dcl.StringOrNil(p.DestroyScheduledDuration),
		Project:                  dcl.StringOrNil(p.Project),
		Location:                 dcl.StringOrNil(p.Location),
		KeyRing:                  dcl.StringOrNil(p.KeyRing),
	}
	return obj
}

// CryptoKeyPrimaryStateEnumToProto converts a CryptoKeyPrimaryStateEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryStateEnumToProto(e *beta.CryptoKeyPrimaryStateEnum) betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum_value["CryptoKeyPrimaryStateEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyPrimaryStateEnum(0)
}

// CryptoKeyPrimaryProtectionLevelEnumToProto converts a CryptoKeyPrimaryProtectionLevelEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnumToProto(e *beta.CryptoKeyPrimaryProtectionLevelEnum) betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum_value["CryptoKeyPrimaryProtectionLevelEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnum(0)
}

// CryptoKeyPrimaryAlgorithmEnumToProto converts a CryptoKeyPrimaryAlgorithmEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryAlgorithmEnumToProto(e *beta.CryptoKeyPrimaryAlgorithmEnum) betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum_value["CryptoKeyPrimaryAlgorithmEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyPrimaryAlgorithmEnum(0)
}

// CryptoKeyPrimaryAttestationFormatEnumToProto converts a CryptoKeyPrimaryAttestationFormatEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnumToProto(e *beta.CryptoKeyPrimaryAttestationFormatEnum) betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum_value["CryptoKeyPrimaryAttestationFormatEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnum(0)
}

// CryptoKeyPurposeEnumToProto converts a CryptoKeyPurposeEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyPurposeEnumToProto(e *beta.CryptoKeyPurposeEnum) betapb.CloudkmsBetaCryptoKeyPurposeEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyPurposeEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyPurposeEnum_value["CryptoKeyPurposeEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyPurposeEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyPurposeEnum(0)
}

// CryptoKeyVersionTemplateProtectionLevelEnumToProto converts a CryptoKeyVersionTemplateProtectionLevelEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnumToProto(e *beta.CryptoKeyVersionTemplateProtectionLevelEnum) betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum_value["CryptoKeyVersionTemplateProtectionLevelEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnum(0)
}

// CryptoKeyVersionTemplateAlgorithmEnumToProto converts a CryptoKeyVersionTemplateAlgorithmEnum enum to its proto representation.
func CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnumToProto(e *beta.CryptoKeyVersionTemplateAlgorithmEnum) betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum {
	if e == nil {
		return betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum(0)
	}
	if v, ok := betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum_value["CryptoKeyVersionTemplateAlgorithmEnum"+string(*e)]; ok {
		return betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum(v)
	}
	return betapb.CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnum(0)
}

// CryptoKeyPrimaryToProto converts a CryptoKeyPrimary resource to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryToProto(o *beta.CryptoKeyPrimary) *betapb.CloudkmsBetaCryptoKeyPrimary {
	if o == nil {
		return nil
	}
	p := &betapb.CloudkmsBetaCryptoKeyPrimary{
		Name:                           dcl.ValueOrEmptyString(o.Name),
		State:                          CloudkmsBetaCryptoKeyPrimaryStateEnumToProto(o.State),
		ProtectionLevel:                CloudkmsBetaCryptoKeyPrimaryProtectionLevelEnumToProto(o.ProtectionLevel),
		Algorithm:                      CloudkmsBetaCryptoKeyPrimaryAlgorithmEnumToProto(o.Algorithm),
		Attestation:                    CloudkmsBetaCryptoKeyPrimaryAttestationToProto(o.Attestation),
		CreateTime:                     dcl.ValueOrEmptyString(o.CreateTime),
		GenerateTime:                   dcl.ValueOrEmptyString(o.GenerateTime),
		DestroyTime:                    dcl.ValueOrEmptyString(o.DestroyTime),
		DestroyEventTime:               dcl.ValueOrEmptyString(o.DestroyEventTime),
		ImportJob:                      dcl.ValueOrEmptyString(o.ImportJob),
		ImportTime:                     dcl.ValueOrEmptyString(o.ImportTime),
		ImportFailureReason:            dcl.ValueOrEmptyString(o.ImportFailureReason),
		ExternalProtectionLevelOptions: CloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o.ExternalProtectionLevelOptions),
		ReimportEligible:               dcl.ValueOrEmptyBool(o.ReimportEligible),
	}
	return p
}

// CryptoKeyPrimaryAttestationToProto converts a CryptoKeyPrimaryAttestation resource to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryAttestationToProto(o *beta.CryptoKeyPrimaryAttestation) *betapb.CloudkmsBetaCryptoKeyPrimaryAttestation {
	if o == nil {
		return nil
	}
	p := &betapb.CloudkmsBetaCryptoKeyPrimaryAttestation{
		Format:     CloudkmsBetaCryptoKeyPrimaryAttestationFormatEnumToProto(o.Format),
		Content:    dcl.ValueOrEmptyString(o.Content),
		CertChains: CloudkmsBetaCryptoKeyPrimaryAttestationCertChainsToProto(o.CertChains),
	}
	return p
}

// CryptoKeyPrimaryAttestationCertChainsToProto converts a CryptoKeyPrimaryAttestationCertChains resource to its proto representation.
func CloudkmsBetaCryptoKeyPrimaryAttestationCertChainsToProto(o *beta.CryptoKeyPrimaryAttestationCertChains) *betapb.CloudkmsBetaCryptoKeyPrimaryAttestationCertChains {
	if o == nil {
		return nil
	}
	p := &betapb.CloudkmsBetaCryptoKeyPrimaryAttestationCertChains{}
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
func CloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o *beta.CryptoKeyPrimaryExternalProtectionLevelOptions) *betapb.CloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptions {
	if o == nil {
		return nil
	}
	p := &betapb.CloudkmsBetaCryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.ValueOrEmptyString(o.ExternalKeyUri),
	}
	return p
}

// CryptoKeyVersionTemplateToProto converts a CryptoKeyVersionTemplate resource to its proto representation.
func CloudkmsBetaCryptoKeyVersionTemplateToProto(o *beta.CryptoKeyVersionTemplate) *betapb.CloudkmsBetaCryptoKeyVersionTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.CloudkmsBetaCryptoKeyVersionTemplate{
		ProtectionLevel: CloudkmsBetaCryptoKeyVersionTemplateProtectionLevelEnumToProto(o.ProtectionLevel),
		Algorithm:       CloudkmsBetaCryptoKeyVersionTemplateAlgorithmEnumToProto(o.Algorithm),
	}
	return p
}

// CryptoKeyToProto converts a CryptoKey resource to its proto representation.
func CryptoKeyToProto(resource *beta.CryptoKey) *betapb.CloudkmsBetaCryptoKey {
	p := &betapb.CloudkmsBetaCryptoKey{
		Name:                     dcl.ValueOrEmptyString(resource.Name),
		Primary:                  CloudkmsBetaCryptoKeyPrimaryToProto(resource.Primary),
		Purpose:                  CloudkmsBetaCryptoKeyPurposeEnumToProto(resource.Purpose),
		CreateTime:               dcl.ValueOrEmptyString(resource.CreateTime),
		NextRotationTime:         dcl.ValueOrEmptyString(resource.NextRotationTime),
		RotationPeriod:           dcl.ValueOrEmptyString(resource.RotationPeriod),
		VersionTemplate:          CloudkmsBetaCryptoKeyVersionTemplateToProto(resource.VersionTemplate),
		ImportOnly:               dcl.ValueOrEmptyBool(resource.ImportOnly),
		DestroyScheduledDuration: dcl.ValueOrEmptyString(resource.DestroyScheduledDuration),
		Project:                  dcl.ValueOrEmptyString(resource.Project),
		Location:                 dcl.ValueOrEmptyString(resource.Location),
		KeyRing:                  dcl.ValueOrEmptyString(resource.KeyRing),
	}

	return p
}

// ApplyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) applyCryptoKey(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudkmsBetaCryptoKeyRequest) (*betapb.CloudkmsBetaCryptoKey, error) {
	p := ProtoToCryptoKey(request.GetResource())
	res, err := c.ApplyCryptoKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CryptoKeyToProto(res)
	return r, nil
}

// ApplyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) ApplyCloudkmsBetaCryptoKey(ctx context.Context, request *betapb.ApplyCloudkmsBetaCryptoKeyRequest) (*betapb.CloudkmsBetaCryptoKey, error) {
	cl, err := createConfigCryptoKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCryptoKey(ctx, cl, request)
}

// DeleteCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Delete() method.
func (s *CryptoKeyServer) DeleteCloudkmsBetaCryptoKey(ctx context.Context, request *betapb.DeleteCloudkmsBetaCryptoKeyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for CryptoKey")

}

// ListCloudkmsBetaCryptoKey handles the gRPC request by passing it to the underlying CryptoKeyList() method.
func (s *CryptoKeyServer) ListCloudkmsBetaCryptoKey(ctx context.Context, request *betapb.ListCloudkmsBetaCryptoKeyRequest) (*betapb.ListCloudkmsBetaCryptoKeyResponse, error) {
	cl, err := createConfigCryptoKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCryptoKey(ctx, request.Project, request.Location, request.KeyRing)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudkmsBetaCryptoKey
	for _, r := range resources.Items {
		rp := CryptoKeyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListCloudkmsBetaCryptoKeyResponse{Items: protos}, nil
}

func createConfigCryptoKey(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
