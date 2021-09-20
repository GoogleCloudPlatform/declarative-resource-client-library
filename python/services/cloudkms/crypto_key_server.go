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
	cloudkmspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/cloudkms_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms"
)

// Server implements the gRPC interface for CryptoKey.
type CryptoKeyServer struct{}

// ProtoToCryptoKeyPrimaryStateEnum converts a CryptoKeyPrimaryStateEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryStateEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum) *cloudkms.CryptoKeyPrimaryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryStateEnum(n[len("CloudkmsCryptoKeyPrimaryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryProtectionLevelEnum converts a CryptoKeyPrimaryProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryProtectionLevelEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum) *cloudkms.CryptoKeyPrimaryProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryProtectionLevelEnum(n[len("CloudkmsCryptoKeyPrimaryProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAlgorithmEnum converts a CryptoKeyPrimaryAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAlgorithmEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum) *cloudkms.CryptoKeyPrimaryAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryAlgorithmEnum(n[len("CloudkmsCryptoKeyPrimaryAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAttestationFormatEnum converts a CryptoKeyPrimaryAttestationFormatEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAttestationFormatEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum) *cloudkms.CryptoKeyPrimaryAttestationFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryAttestationFormatEnum(n[len("CloudkmsCryptoKeyPrimaryAttestationFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPurposeEnum converts a CryptoKeyPurposeEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPurposeEnum(e cloudkmspb.CloudkmsCryptoKeyPurposeEnum) *cloudkms.CryptoKeyPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPurposeEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPurposeEnum(n[len("CloudkmsCryptoKeyPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateProtectionLevelEnum converts a CryptoKeyVersionTemplateProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(e cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum) *cloudkms.CryptoKeyVersionTemplateProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyVersionTemplateProtectionLevelEnum(n[len("CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateAlgorithmEnum converts a CryptoKeyVersionTemplateAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyVersionTemplateAlgorithmEnum(e cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum) *cloudkms.CryptoKeyVersionTemplateAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyVersionTemplateAlgorithmEnum(n[len("CloudkmsCryptoKeyVersionTemplateAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimary converts a CryptoKeyPrimary resource from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimary(p *cloudkmspb.CloudkmsCryptoKeyPrimary) *cloudkms.CryptoKeyPrimary {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimary{
		Name:                           dcl.StringOrNil(p.Name),
		State:                          ProtoToCloudkmsCryptoKeyPrimaryStateEnum(p.GetState()),
		ProtectionLevel:                ProtoToCloudkmsCryptoKeyPrimaryProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:                      ProtoToCloudkmsCryptoKeyPrimaryAlgorithmEnum(p.GetAlgorithm()),
		Attestation:                    ProtoToCloudkmsCryptoKeyPrimaryAttestation(p.GetAttestation()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		GenerateTime:                   dcl.StringOrNil(p.GetGenerateTime()),
		DestroyTime:                    dcl.StringOrNil(p.GetDestroyTime()),
		DestroyEventTime:               dcl.StringOrNil(p.GetDestroyEventTime()),
		ImportJob:                      dcl.StringOrNil(p.ImportJob),
		ImportTime:                     dcl.StringOrNil(p.GetImportTime()),
		ImportFailureReason:            dcl.StringOrNil(p.ImportFailureReason),
		ExternalProtectionLevelOptions: ProtoToCloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions(p.GetExternalProtectionLevelOptions()),
		ReimportEligible:               dcl.Bool(p.ReimportEligible),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestation converts a CryptoKeyPrimaryAttestation resource from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAttestation(p *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestation) *cloudkms.CryptoKeyPrimaryAttestation {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimaryAttestation{
		Format:     ProtoToCloudkmsCryptoKeyPrimaryAttestationFormatEnum(p.GetFormat()),
		Content:    dcl.StringOrNil(p.Content),
		CertChains: ProtoToCloudkmsCryptoKeyPrimaryAttestationCertChains(p.GetCertChains()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestationCertChains converts a CryptoKeyPrimaryAttestationCertChains resource from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAttestationCertChains(p *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationCertChains) *cloudkms.CryptoKeyPrimaryAttestationCertChains {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimaryAttestationCertChains{}
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
func ProtoToCloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions(p *cloudkmspb.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions) *cloudkms.CryptoKeyPrimaryExternalProtectionLevelOptions {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.StringOrNil(p.ExternalKeyUri),
	}
	return obj
}

// ProtoToCryptoKeyVersionTemplate converts a CryptoKeyVersionTemplate resource from its proto representation.
func ProtoToCloudkmsCryptoKeyVersionTemplate(p *cloudkmspb.CloudkmsCryptoKeyVersionTemplate) *cloudkms.CryptoKeyVersionTemplate {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyVersionTemplate{
		ProtectionLevel: ProtoToCloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:       ProtoToCloudkmsCryptoKeyVersionTemplateAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCryptoKey converts a CryptoKey resource from its proto representation.
func ProtoToCryptoKey(p *cloudkmspb.CloudkmsCryptoKey) *cloudkms.CryptoKey {
	obj := &cloudkms.CryptoKey{
		Name:                     dcl.StringOrNil(p.Name),
		Primary:                  ProtoToCloudkmsCryptoKeyPrimary(p.GetPrimary()),
		Purpose:                  ProtoToCloudkmsCryptoKeyPurposeEnum(p.GetPurpose()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		NextRotationTime:         dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:           dcl.StringOrNil(p.RotationPeriod),
		VersionTemplate:          ProtoToCloudkmsCryptoKeyVersionTemplate(p.GetVersionTemplate()),
		ImportOnly:               dcl.Bool(p.ImportOnly),
		DestroyScheduledDuration: dcl.StringOrNil(p.DestroyScheduledDuration),
		Project:                  dcl.StringOrNil(p.Project),
		Location:                 dcl.StringOrNil(p.Location),
		KeyRing:                  dcl.StringOrNil(p.KeyRing),
	}
	return obj
}

// CryptoKeyPrimaryStateEnumToProto converts a CryptoKeyPrimaryStateEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryStateEnumToProto(e *cloudkms.CryptoKeyPrimaryStateEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum_value["CryptoKeyPrimaryStateEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum(0)
}

// CryptoKeyPrimaryProtectionLevelEnumToProto converts a CryptoKeyPrimaryProtectionLevelEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryProtectionLevelEnumToProto(e *cloudkms.CryptoKeyPrimaryProtectionLevelEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum_value["CryptoKeyPrimaryProtectionLevelEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum(0)
}

// CryptoKeyPrimaryAlgorithmEnumToProto converts a CryptoKeyPrimaryAlgorithmEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryAlgorithmEnumToProto(e *cloudkms.CryptoKeyPrimaryAlgorithmEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum_value["CryptoKeyPrimaryAlgorithmEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum(0)
}

// CryptoKeyPrimaryAttestationFormatEnumToProto converts a CryptoKeyPrimaryAttestationFormatEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryAttestationFormatEnumToProto(e *cloudkms.CryptoKeyPrimaryAttestationFormatEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum_value["CryptoKeyPrimaryAttestationFormatEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum(0)
}

// CryptoKeyPurposeEnumToProto converts a CryptoKeyPurposeEnum enum to its proto representation.
func CloudkmsCryptoKeyPurposeEnumToProto(e *cloudkms.CryptoKeyPurposeEnum) cloudkmspb.CloudkmsCryptoKeyPurposeEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPurposeEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPurposeEnum_value["CryptoKeyPurposeEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPurposeEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPurposeEnum(0)
}

// CryptoKeyVersionTemplateProtectionLevelEnumToProto converts a CryptoKeyVersionTemplateProtectionLevelEnum enum to its proto representation.
func CloudkmsCryptoKeyVersionTemplateProtectionLevelEnumToProto(e *cloudkms.CryptoKeyVersionTemplateProtectionLevelEnum) cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum_value["CryptoKeyVersionTemplateProtectionLevelEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(0)
}

// CryptoKeyVersionTemplateAlgorithmEnumToProto converts a CryptoKeyVersionTemplateAlgorithmEnum enum to its proto representation.
func CloudkmsCryptoKeyVersionTemplateAlgorithmEnumToProto(e *cloudkms.CryptoKeyVersionTemplateAlgorithmEnum) cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum_value["CryptoKeyVersionTemplateAlgorithmEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum(0)
}

// CryptoKeyPrimaryToProto converts a CryptoKeyPrimary resource to its proto representation.
func CloudkmsCryptoKeyPrimaryToProto(o *cloudkms.CryptoKeyPrimary) *cloudkmspb.CloudkmsCryptoKeyPrimary {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimary{
		Name:                           dcl.ValueOrEmptyString(o.Name),
		State:                          CloudkmsCryptoKeyPrimaryStateEnumToProto(o.State),
		ProtectionLevel:                CloudkmsCryptoKeyPrimaryProtectionLevelEnumToProto(o.ProtectionLevel),
		Algorithm:                      CloudkmsCryptoKeyPrimaryAlgorithmEnumToProto(o.Algorithm),
		Attestation:                    CloudkmsCryptoKeyPrimaryAttestationToProto(o.Attestation),
		CreateTime:                     dcl.ValueOrEmptyString(o.CreateTime),
		GenerateTime:                   dcl.ValueOrEmptyString(o.GenerateTime),
		DestroyTime:                    dcl.ValueOrEmptyString(o.DestroyTime),
		DestroyEventTime:               dcl.ValueOrEmptyString(o.DestroyEventTime),
		ImportJob:                      dcl.ValueOrEmptyString(o.ImportJob),
		ImportTime:                     dcl.ValueOrEmptyString(o.ImportTime),
		ImportFailureReason:            dcl.ValueOrEmptyString(o.ImportFailureReason),
		ExternalProtectionLevelOptions: CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o.ExternalProtectionLevelOptions),
		ReimportEligible:               dcl.ValueOrEmptyBool(o.ReimportEligible),
	}
	return p
}

// CryptoKeyPrimaryAttestationToProto converts a CryptoKeyPrimaryAttestation resource to its proto representation.
func CloudkmsCryptoKeyPrimaryAttestationToProto(o *cloudkms.CryptoKeyPrimaryAttestation) *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestation {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimaryAttestation{
		Format:     CloudkmsCryptoKeyPrimaryAttestationFormatEnumToProto(o.Format),
		Content:    dcl.ValueOrEmptyString(o.Content),
		CertChains: CloudkmsCryptoKeyPrimaryAttestationCertChainsToProto(o.CertChains),
	}
	return p
}

// CryptoKeyPrimaryAttestationCertChainsToProto converts a CryptoKeyPrimaryAttestationCertChains resource to its proto representation.
func CloudkmsCryptoKeyPrimaryAttestationCertChainsToProto(o *cloudkms.CryptoKeyPrimaryAttestationCertChains) *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationCertChains {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationCertChains{}
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
func CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o *cloudkms.CryptoKeyPrimaryExternalProtectionLevelOptions) *cloudkmspb.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.ValueOrEmptyString(o.ExternalKeyUri),
	}
	return p
}

// CryptoKeyVersionTemplateToProto converts a CryptoKeyVersionTemplate resource to its proto representation.
func CloudkmsCryptoKeyVersionTemplateToProto(o *cloudkms.CryptoKeyVersionTemplate) *cloudkmspb.CloudkmsCryptoKeyVersionTemplate {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyVersionTemplate{
		ProtectionLevel: CloudkmsCryptoKeyVersionTemplateProtectionLevelEnumToProto(o.ProtectionLevel),
		Algorithm:       CloudkmsCryptoKeyVersionTemplateAlgorithmEnumToProto(o.Algorithm),
	}
	return p
}

// CryptoKeyToProto converts a CryptoKey resource to its proto representation.
func CryptoKeyToProto(resource *cloudkms.CryptoKey) *cloudkmspb.CloudkmsCryptoKey {
	p := &cloudkmspb.CloudkmsCryptoKey{
		Name:                     dcl.ValueOrEmptyString(resource.Name),
		Primary:                  CloudkmsCryptoKeyPrimaryToProto(resource.Primary),
		Purpose:                  CloudkmsCryptoKeyPurposeEnumToProto(resource.Purpose),
		CreateTime:               dcl.ValueOrEmptyString(resource.CreateTime),
		NextRotationTime:         dcl.ValueOrEmptyString(resource.NextRotationTime),
		RotationPeriod:           dcl.ValueOrEmptyString(resource.RotationPeriod),
		VersionTemplate:          CloudkmsCryptoKeyVersionTemplateToProto(resource.VersionTemplate),
		ImportOnly:               dcl.ValueOrEmptyBool(resource.ImportOnly),
		DestroyScheduledDuration: dcl.ValueOrEmptyString(resource.DestroyScheduledDuration),
		Project:                  dcl.ValueOrEmptyString(resource.Project),
		Location:                 dcl.ValueOrEmptyString(resource.Location),
		KeyRing:                  dcl.ValueOrEmptyString(resource.KeyRing),
	}

	return p
}

// ApplyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) applyCryptoKey(ctx context.Context, c *cloudkms.Client, request *cloudkmspb.ApplyCloudkmsCryptoKeyRequest) (*cloudkmspb.CloudkmsCryptoKey, error) {
	p := ProtoToCryptoKey(request.GetResource())
	res, err := c.ApplyCryptoKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CryptoKeyToProto(res)
	return r, nil
}

// ApplyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) ApplyCloudkmsCryptoKey(ctx context.Context, request *cloudkmspb.ApplyCloudkmsCryptoKeyRequest) (*cloudkmspb.CloudkmsCryptoKey, error) {
	cl, err := createConfigCryptoKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCryptoKey(ctx, cl, request)
}

// DeleteCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Delete() method.
func (s *CryptoKeyServer) DeleteCloudkmsCryptoKey(ctx context.Context, request *cloudkmspb.DeleteCloudkmsCryptoKeyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for CryptoKey")

}

// ListCloudkmsCryptoKey handles the gRPC request by passing it to the underlying CryptoKeyList() method.
func (s *CryptoKeyServer) ListCloudkmsCryptoKey(ctx context.Context, request *cloudkmspb.ListCloudkmsCryptoKeyRequest) (*cloudkmspb.ListCloudkmsCryptoKeyResponse, error) {
	cl, err := createConfigCryptoKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCryptoKey(ctx, ProtoToCryptoKey(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*cloudkmspb.CloudkmsCryptoKey
	for _, r := range resources.Items {
		rp := CryptoKeyToProto(r)
		protos = append(protos, rp)
	}
	return &cloudkmspb.ListCloudkmsCryptoKeyResponse{Items: protos}, nil
}

func createConfigCryptoKey(ctx context.Context, service_account_file string) (*cloudkms.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudkms.NewClient(conf), nil
}
