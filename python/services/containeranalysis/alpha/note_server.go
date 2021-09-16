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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeranalysis/alpha/containeranalysis_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/alpha"
)

// Server implements the gRPC interface for Note.
type NoteServer struct{}

// ProtoToNotePackageDistributionArchitectureEnum converts a NotePackageDistributionArchitectureEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionArchitectureEnum(e alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum) *alpha.NotePackageDistributionArchitectureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum_name[int32(e)]; ok {
		e := alpha.NotePackageDistributionArchitectureEnum(n[len("ContaineranalysisAlphaNotePackageDistributionArchitectureEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionLatestVersionKindEnum converts a NotePackageDistributionLatestVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(e alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum) *alpha.NotePackageDistributionLatestVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum_name[int32(e)]; ok {
		e := alpha.NotePackageDistributionLatestVersionKindEnum(n[len("ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteDiscoveryAnalysisKindEnum converts a NoteDiscoveryAnalysisKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(e alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum) *alpha.NoteDiscoveryAnalysisKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum_name[int32(e)]; ok {
		e := alpha.NoteDiscoveryAnalysisKindEnum(n[len("ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteRelatedUrl(p *alphapb.ContaineranalysisAlphaNoteRelatedUrl) *alpha.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.Url),
		Label: dcl.StringOrNil(p.Label),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteImage(p *alphapb.ContaineranalysisAlphaNoteImage) *alpha.NoteImage {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.ResourceUrl),
		Fingerprint: ProtoToContaineranalysisAlphaNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteImageFingerprint(p *alphapb.ContaineranalysisAlphaNoteImageFingerprint) *alpha.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.V1Name),
		V2Name: dcl.StringOrNil(p.V2Name),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage resource from its proto representation.
func ProtoToContaineranalysisAlphaNotePackage(p *alphapb.ContaineranalysisAlphaNotePackage) *alpha.NotePackage {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackage{
		Name: dcl.StringOrNil(p.Name),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisAlphaNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution resource from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistribution(p *alphapb.ContaineranalysisAlphaNotePackageDistribution) *alpha.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.CpeUri),
		Architecture:  ProtoToContaineranalysisAlphaNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.Maintainer),
		Url:           dcl.StringOrNil(p.Url),
		Description:   dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion resource from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersion(p *alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion) *alpha.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteDiscovery(p *alphapb.ContaineranalysisAlphaNoteDiscovery) *alpha.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNoteDeployment converts a NoteDeployment resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteDeployment(p *alphapb.ContaineranalysisAlphaNoteDeployment) *alpha.NoteDeployment {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteDeployment{}
	for _, r := range p.GetResourceUri() {
		obj.ResourceUri = append(obj.ResourceUri, r)
	}
	return obj
}

// ProtoToNoteAttestation converts a NoteAttestation resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteAttestation(p *alphapb.ContaineranalysisAlphaNoteAttestation) *alpha.NoteAttestation {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteAttestation{
		Hint: ProtoToContaineranalysisAlphaNoteAttestationHint(p.GetHint()),
	}
	return obj
}

// ProtoToNoteAttestationHint converts a NoteAttestationHint resource from its proto representation.
func ProtoToContaineranalysisAlphaNoteAttestationHint(p *alphapb.ContaineranalysisAlphaNoteAttestationHint) *alpha.NoteAttestationHint {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteAttestationHint{
		HumanReadableName: dcl.StringOrNil(p.HumanReadableName),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *alphapb.ContaineranalysisAlphaNote) *alpha.Note {
	obj := &alpha.Note{
		Name:             dcl.StringOrNil(p.Name),
		ShortDescription: dcl.StringOrNil(p.ShortDescription),
		LongDescription:  dcl.StringOrNil(p.LongDescription),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Image:            ProtoToContaineranalysisAlphaNoteImage(p.GetImage()),
		Package:          ProtoToContaineranalysisAlphaNotePackage(p.GetPackage()),
		Discovery:        ProtoToContaineranalysisAlphaNoteDiscovery(p.GetDiscovery()),
		Deployment:       ProtoToContaineranalysisAlphaNoteDeployment(p.GetDeployment()),
		Attestation:      ProtoToContaineranalysisAlphaNoteAttestation(p.GetAttestation()),
		Project:          dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetRelatedUrl() {
		obj.RelatedUrl = append(obj.RelatedUrl, *ProtoToContaineranalysisAlphaNoteRelatedUrl(r))
	}
	return obj
}

// NotePackageDistributionArchitectureEnumToProto converts a NotePackageDistributionArchitectureEnum enum to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionArchitectureEnumToProto(e *alpha.NotePackageDistributionArchitectureEnum) alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum_value["NotePackageDistributionArchitectureEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum(0)
}

// NotePackageDistributionLatestVersionKindEnumToProto converts a NotePackageDistributionLatestVersionKindEnum enum to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnumToProto(e *alpha.NotePackageDistributionLatestVersionKindEnum) alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum_value["NotePackageDistributionLatestVersionKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(0)
}

// NoteDiscoveryAnalysisKindEnumToProto converts a NoteDiscoveryAnalysisKindEnum enum to its proto representation.
func ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnumToProto(e *alpha.NoteDiscoveryAnalysisKindEnum) alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum_value["NoteDiscoveryAnalysisKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(0)
}

// NoteRelatedUrlToProto converts a NoteRelatedUrl resource to its proto representation.
func ContaineranalysisAlphaNoteRelatedUrlToProto(o *alpha.NoteRelatedUrl) *alphapb.ContaineranalysisAlphaNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteRelatedUrl{
		Url:   dcl.ValueOrEmptyString(o.Url),
		Label: dcl.ValueOrEmptyString(o.Label),
	}
	return p
}

// NoteImageToProto converts a NoteImage resource to its proto representation.
func ContaineranalysisAlphaNoteImageToProto(o *alpha.NoteImage) *alphapb.ContaineranalysisAlphaNoteImage {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteImage{
		ResourceUrl: dcl.ValueOrEmptyString(o.ResourceUrl),
		Fingerprint: ContaineranalysisAlphaNoteImageFingerprintToProto(o.Fingerprint),
	}
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint resource to its proto representation.
func ContaineranalysisAlphaNoteImageFingerprintToProto(o *alpha.NoteImageFingerprint) *alphapb.ContaineranalysisAlphaNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteImageFingerprint{
		V1Name: dcl.ValueOrEmptyString(o.V1Name),
		V2Name: dcl.ValueOrEmptyString(o.V2Name),
	}
	for _, r := range o.V2Blob {
		p.V2Blob = append(p.V2Blob, r)
	}
	return p
}

// NotePackageToProto converts a NotePackage resource to its proto representation.
func ContaineranalysisAlphaNotePackageToProto(o *alpha.NotePackage) *alphapb.ContaineranalysisAlphaNotePackage {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackage{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	for _, r := range o.Distribution {
		p.Distribution = append(p.Distribution, ContaineranalysisAlphaNotePackageDistributionToProto(&r))
	}
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution resource to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionToProto(o *alpha.NotePackageDistribution) *alphapb.ContaineranalysisAlphaNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackageDistribution{
		CpeUri:        dcl.ValueOrEmptyString(o.CpeUri),
		Architecture:  ContaineranalysisAlphaNotePackageDistributionArchitectureEnumToProto(o.Architecture),
		LatestVersion: ContaineranalysisAlphaNotePackageDistributionLatestVersionToProto(o.LatestVersion),
		Maintainer:    dcl.ValueOrEmptyString(o.Maintainer),
		Url:           dcl.ValueOrEmptyString(o.Url),
		Description:   dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion resource to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionLatestVersionToProto(o *alpha.NotePackageDistributionLatestVersion) *alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery resource to its proto representation.
func ContaineranalysisAlphaNoteDiscoveryToProto(o *alpha.NoteDiscovery) *alphapb.ContaineranalysisAlphaNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteDiscovery{
		AnalysisKind: ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind),
	}
	return p
}

// NoteDeploymentToProto converts a NoteDeployment resource to its proto representation.
func ContaineranalysisAlphaNoteDeploymentToProto(o *alpha.NoteDeployment) *alphapb.ContaineranalysisAlphaNoteDeployment {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteDeployment{}
	for _, r := range o.ResourceUri {
		p.ResourceUri = append(p.ResourceUri, r)
	}
	return p
}

// NoteAttestationToProto converts a NoteAttestation resource to its proto representation.
func ContaineranalysisAlphaNoteAttestationToProto(o *alpha.NoteAttestation) *alphapb.ContaineranalysisAlphaNoteAttestation {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteAttestation{
		Hint: ContaineranalysisAlphaNoteAttestationHintToProto(o.Hint),
	}
	return p
}

// NoteAttestationHintToProto converts a NoteAttestationHint resource to its proto representation.
func ContaineranalysisAlphaNoteAttestationHintToProto(o *alpha.NoteAttestationHint) *alphapb.ContaineranalysisAlphaNoteAttestationHint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteAttestationHint{
		HumanReadableName: dcl.ValueOrEmptyString(o.HumanReadableName),
	}
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *alpha.Note) *alphapb.ContaineranalysisAlphaNote {
	p := &alphapb.ContaineranalysisAlphaNote{
		Name:             dcl.ValueOrEmptyString(resource.Name),
		ShortDescription: dcl.ValueOrEmptyString(resource.ShortDescription),
		LongDescription:  dcl.ValueOrEmptyString(resource.LongDescription),
		ExpirationTime:   dcl.ValueOrEmptyString(resource.ExpirationTime),
		CreateTime:       dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:       dcl.ValueOrEmptyString(resource.UpdateTime),
		Image:            ContaineranalysisAlphaNoteImageToProto(resource.Image),
		Package:          ContaineranalysisAlphaNotePackageToProto(resource.Package),
		Discovery:        ContaineranalysisAlphaNoteDiscoveryToProto(resource.Discovery),
		Deployment:       ContaineranalysisAlphaNoteDeploymentToProto(resource.Deployment),
		Attestation:      ContaineranalysisAlphaNoteAttestationToProto(resource.Attestation),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.RelatedUrl {
		p.RelatedUrl = append(p.RelatedUrl, ContaineranalysisAlphaNoteRelatedUrlToProto(&r))
	}

	return p
}

// ApplyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) applyNote(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContaineranalysisAlphaNoteRequest) (*alphapb.ContaineranalysisAlphaNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// ApplyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisAlphaNote(ctx context.Context, request *alphapb.ApplyContaineranalysisAlphaNoteRequest) (*alphapb.ContaineranalysisAlphaNote, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisAlphaNote(ctx context.Context, request *alphapb.DeleteContaineranalysisAlphaNoteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))

}

// ListContaineranalysisAlphaNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisAlphaNote(ctx context.Context, request *alphapb.ListContaineranalysisAlphaNoteRequest) (*alphapb.ListContaineranalysisAlphaNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, ProtoToNote(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContaineranalysisAlphaNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContaineranalysisAlphaNoteResponse{Items: protos}, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
