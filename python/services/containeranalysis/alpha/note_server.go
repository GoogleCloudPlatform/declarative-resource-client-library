// Copyright 2022 Google LLC. All Rights Reserved.
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

// NoteServer implements the gRPC interface for Note.
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

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl object from its proto representation.
func ProtoToContaineranalysisAlphaNoteRelatedUrl(p *alphapb.ContaineranalysisAlphaNoteRelatedUrl) *alpha.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.GetUrl()),
		Label: dcl.StringOrNil(p.GetLabel()),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage object from its proto representation.
func ProtoToContaineranalysisAlphaNoteImage(p *alphapb.ContaineranalysisAlphaNoteImage) *alpha.NoteImage {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.GetResourceUrl()),
		Fingerprint: ProtoToContaineranalysisAlphaNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint object from its proto representation.
func ProtoToContaineranalysisAlphaNoteImageFingerprint(p *alphapb.ContaineranalysisAlphaNoteImageFingerprint) *alpha.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.GetV1Name()),
		V2Name: dcl.StringOrNil(p.GetV2Name()),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage object from its proto representation.
func ProtoToContaineranalysisAlphaNotePackage(p *alphapb.ContaineranalysisAlphaNotePackage) *alpha.NotePackage {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackage{
		Name: dcl.StringOrNil(p.GetName()),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisAlphaNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution object from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistribution(p *alphapb.ContaineranalysisAlphaNotePackageDistribution) *alpha.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.GetCpeUri()),
		Architecture:  ProtoToContaineranalysisAlphaNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.GetMaintainer()),
		Url:           dcl.StringOrNil(p.GetUrl()),
		Description:   dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion object from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersion(p *alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion) *alpha.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery object from its proto representation.
func ProtoToContaineranalysisAlphaNoteDiscovery(p *alphapb.ContaineranalysisAlphaNoteDiscovery) *alpha.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNoteDeployment converts a NoteDeployment object from its proto representation.
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

// ProtoToNoteAttestation converts a NoteAttestation object from its proto representation.
func ProtoToContaineranalysisAlphaNoteAttestation(p *alphapb.ContaineranalysisAlphaNoteAttestation) *alpha.NoteAttestation {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteAttestation{
		Hint: ProtoToContaineranalysisAlphaNoteAttestationHint(p.GetHint()),
	}
	return obj
}

// ProtoToNoteAttestationHint converts a NoteAttestationHint object from its proto representation.
func ProtoToContaineranalysisAlphaNoteAttestationHint(p *alphapb.ContaineranalysisAlphaNoteAttestationHint) *alpha.NoteAttestationHint {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteAttestationHint{
		HumanReadableName: dcl.StringOrNil(p.GetHumanReadableName()),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *alphapb.ContaineranalysisAlphaNote) *alpha.Note {
	obj := &alpha.Note{
		Name:             dcl.StringOrNil(p.GetName()),
		ShortDescription: dcl.StringOrNil(p.GetShortDescription()),
		LongDescription:  dcl.StringOrNil(p.GetLongDescription()),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Image:            ProtoToContaineranalysisAlphaNoteImage(p.GetImage()),
		Package:          ProtoToContaineranalysisAlphaNotePackage(p.GetPackage()),
		Discovery:        ProtoToContaineranalysisAlphaNoteDiscovery(p.GetDiscovery()),
		Deployment:       ProtoToContaineranalysisAlphaNoteDeployment(p.GetDeployment()),
		Attestation:      ProtoToContaineranalysisAlphaNoteAttestation(p.GetAttestation()),
		Project:          dcl.StringOrNil(p.GetProject()),
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

// NoteRelatedUrlToProto converts a NoteRelatedUrl object to its proto representation.
func ContaineranalysisAlphaNoteRelatedUrlToProto(o *alpha.NoteRelatedUrl) *alphapb.ContaineranalysisAlphaNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteRelatedUrl{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	return p
}

// NoteImageToProto converts a NoteImage object to its proto representation.
func ContaineranalysisAlphaNoteImageToProto(o *alpha.NoteImage) *alphapb.ContaineranalysisAlphaNoteImage {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteImage{}
	p.SetResourceUrl(dcl.ValueOrEmptyString(o.ResourceUrl))
	p.SetFingerprint(ContaineranalysisAlphaNoteImageFingerprintToProto(o.Fingerprint))
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint object to its proto representation.
func ContaineranalysisAlphaNoteImageFingerprintToProto(o *alpha.NoteImageFingerprint) *alphapb.ContaineranalysisAlphaNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteImageFingerprint{}
	p.SetV1Name(dcl.ValueOrEmptyString(o.V1Name))
	p.SetV2Name(dcl.ValueOrEmptyString(o.V2Name))
	sV2Blob := make([]string, len(o.V2Blob))
	for i, r := range o.V2Blob {
		sV2Blob[i] = r
	}
	p.SetV2Blob(sV2Blob)
	return p
}

// NotePackageToProto converts a NotePackage object to its proto representation.
func ContaineranalysisAlphaNotePackageToProto(o *alpha.NotePackage) *alphapb.ContaineranalysisAlphaNotePackage {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackage{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	sDistribution := make([]*alphapb.ContaineranalysisAlphaNotePackageDistribution, len(o.Distribution))
	for i, r := range o.Distribution {
		sDistribution[i] = ContaineranalysisAlphaNotePackageDistributionToProto(&r)
	}
	p.SetDistribution(sDistribution)
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution object to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionToProto(o *alpha.NotePackageDistribution) *alphapb.ContaineranalysisAlphaNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackageDistribution{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetArchitecture(ContaineranalysisAlphaNotePackageDistributionArchitectureEnumToProto(o.Architecture))
	p.SetLatestVersion(ContaineranalysisAlphaNotePackageDistributionLatestVersionToProto(o.LatestVersion))
	p.SetMaintainer(dcl.ValueOrEmptyString(o.Maintainer))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion object to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionLatestVersionToProto(o *alpha.NotePackageDistributionLatestVersion) *alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery object to its proto representation.
func ContaineranalysisAlphaNoteDiscoveryToProto(o *alpha.NoteDiscovery) *alphapb.ContaineranalysisAlphaNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteDiscovery{}
	p.SetAnalysisKind(ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind))
	return p
}

// NoteDeploymentToProto converts a NoteDeployment object to its proto representation.
func ContaineranalysisAlphaNoteDeploymentToProto(o *alpha.NoteDeployment) *alphapb.ContaineranalysisAlphaNoteDeployment {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteDeployment{}
	sResourceUri := make([]string, len(o.ResourceUri))
	for i, r := range o.ResourceUri {
		sResourceUri[i] = r
	}
	p.SetResourceUri(sResourceUri)
	return p
}

// NoteAttestationToProto converts a NoteAttestation object to its proto representation.
func ContaineranalysisAlphaNoteAttestationToProto(o *alpha.NoteAttestation) *alphapb.ContaineranalysisAlphaNoteAttestation {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteAttestation{}
	p.SetHint(ContaineranalysisAlphaNoteAttestationHintToProto(o.Hint))
	return p
}

// NoteAttestationHintToProto converts a NoteAttestationHint object to its proto representation.
func ContaineranalysisAlphaNoteAttestationHintToProto(o *alpha.NoteAttestationHint) *alphapb.ContaineranalysisAlphaNoteAttestationHint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteAttestationHint{}
	p.SetHumanReadableName(dcl.ValueOrEmptyString(o.HumanReadableName))
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *alpha.Note) *alphapb.ContaineranalysisAlphaNote {
	p := &alphapb.ContaineranalysisAlphaNote{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetShortDescription(dcl.ValueOrEmptyString(resource.ShortDescription))
	p.SetLongDescription(dcl.ValueOrEmptyString(resource.LongDescription))
	p.SetExpirationTime(dcl.ValueOrEmptyString(resource.ExpirationTime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetImage(ContaineranalysisAlphaNoteImageToProto(resource.Image))
	p.SetPackage(ContaineranalysisAlphaNotePackageToProto(resource.Package))
	p.SetDiscovery(ContaineranalysisAlphaNoteDiscoveryToProto(resource.Discovery))
	p.SetDeployment(ContaineranalysisAlphaNoteDeploymentToProto(resource.Deployment))
	p.SetAttestation(ContaineranalysisAlphaNoteAttestationToProto(resource.Attestation))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sRelatedUrl := make([]*alphapb.ContaineranalysisAlphaNoteRelatedUrl, len(resource.RelatedUrl))
	for i, r := range resource.RelatedUrl {
		sRelatedUrl[i] = ContaineranalysisAlphaNoteRelatedUrlToProto(&r)
	}
	p.SetRelatedUrl(sRelatedUrl)

	return p
}

// applyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) applyNote(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContaineranalysisAlphaNoteRequest) (*alphapb.ContaineranalysisAlphaNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// applyContaineranalysisAlphaNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisAlphaNote(ctx context.Context, request *alphapb.ApplyContaineranalysisAlphaNoteRequest) (*alphapb.ContaineranalysisAlphaNote, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisAlphaNote(ctx context.Context, request *alphapb.DeleteContaineranalysisAlphaNoteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))

}

// ListContaineranalysisAlphaNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisAlphaNote(ctx context.Context, request *alphapb.ListContaineranalysisAlphaNoteRequest) (*alphapb.ListContaineranalysisAlphaNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContaineranalysisAlphaNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContaineranalysisAlphaNoteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
