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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeranalysis/beta/containeranalysis_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta"
)

// Server implements the gRPC interface for Note.
type NoteServer struct{}

// ProtoToNoteVulnerabilitySeverityEnum converts a NoteVulnerabilitySeverityEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilitySeverityEnum(e betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum) *beta.NoteVulnerabilitySeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilitySeverityEnum(n[len("NoteVulnerabilitySeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStartKindEnum converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum) *beta.NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityDetailsAffectedVersionStartKindEnum(n[len("NoteVulnerabilityDetailsAffectedVersionStartKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEndKindEnum converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum) *beta.NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityDetailsAffectedVersionEndKindEnum(n[len("NoteVulnerabilityDetailsAffectedVersionEndKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsFixedVersionKindEnum converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum) *beta.NoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityDetailsFixedVersionKindEnum(n[len("NoteVulnerabilityDetailsFixedVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackVectorEnum converts a NoteVulnerabilityCvssV3AttackVectorEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum) *beta.NoteVulnerabilityCvssV3AttackVectorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3AttackVectorEnum(n[len("NoteVulnerabilityCvssV3AttackVectorEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackComplexityEnum converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum) *beta.NoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3AttackComplexityEnum(n[len("NoteVulnerabilityCvssV3AttackComplexityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3PrivilegesRequiredEnum converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum) *beta.NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3PrivilegesRequiredEnum(n[len("NoteVulnerabilityCvssV3PrivilegesRequiredEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3UserInteractionEnum converts a NoteVulnerabilityCvssV3UserInteractionEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum) *beta.NoteVulnerabilityCvssV3UserInteractionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3UserInteractionEnum(n[len("NoteVulnerabilityCvssV3UserInteractionEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ScopeEnum converts a NoteVulnerabilityCvssV3ScopeEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum) *beta.NoteVulnerabilityCvssV3ScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3ScopeEnum(n[len("NoteVulnerabilityCvssV3ScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ConfidentialityImpactEnum converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum) *beta.NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3ConfidentialityImpactEnum(n[len("NoteVulnerabilityCvssV3ConfidentialityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3IntegrityImpactEnum converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum) *beta.NoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3IntegrityImpactEnum(n[len("NoteVulnerabilityCvssV3IntegrityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AvailabilityImpactEnum converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum) *beta.NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3AvailabilityImpactEnum(n[len("NoteVulnerabilityCvssV3AvailabilityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteBuildSignatureKeyTypeEnum converts a NoteBuildSignatureKeyTypeEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(e betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum) *beta.NoteBuildSignatureKeyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum_name[int32(e)]; ok {
		e := beta.NoteBuildSignatureKeyTypeEnum(n[len("NoteBuildSignatureKeyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionArchitectureEnum converts a NotePackageDistributionArchitectureEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistributionArchitectureEnum(e betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum) *beta.NotePackageDistributionArchitectureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum_name[int32(e)]; ok {
		e := beta.NotePackageDistributionArchitectureEnum(n[len("NotePackageDistributionArchitectureEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionLatestVersionKindEnum converts a NotePackageDistributionLatestVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(e betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum) *beta.NotePackageDistributionLatestVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum_name[int32(e)]; ok {
		e := beta.NotePackageDistributionLatestVersionKindEnum(n[len("NotePackageDistributionLatestVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteDiscoveryAnalysisKindEnum converts a NoteDiscoveryAnalysisKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(e betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum) *beta.NoteDiscoveryAnalysisKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum_name[int32(e)]; ok {
		e := beta.NoteDiscoveryAnalysisKindEnum(n[len("NoteDiscoveryAnalysisKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl resource from its proto representation.
func ProtoToContaineranalysisBetaNoteRelatedUrl(p *betapb.ContaineranalysisBetaNoteRelatedUrl) *beta.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &beta.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.Url),
		Label: dcl.StringOrNil(p.Label),
	}
	return obj
}

// ProtoToNoteVulnerability converts a NoteVulnerability resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerability(p *betapb.ContaineranalysisBetaNoteVulnerability) *beta.NoteVulnerability {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerability{
		CvssScore:        dcl.Float64OrNil(p.CvssScore),
		Severity:         ProtoToContaineranalysisBetaNoteVulnerabilitySeverityEnum(p.GetSeverity()),
		CvssV3:           ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3(p.GetCvssV3()),
		SourceUpdateTime: dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToContaineranalysisBetaNoteVulnerabilityDetails(r))
	}
	for _, r := range p.GetWindowsDetails() {
		obj.WindowsDetails = append(obj.WindowsDetails, *ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetails(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityDetails converts a NoteVulnerabilityDetails resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetails(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetails) *beta.NoteVulnerabilityDetails {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetails{
		SeverityName:         dcl.StringOrNil(p.SeverityName),
		Description:          dcl.StringOrNil(p.Description),
		PackageType:          dcl.StringOrNil(p.PackageType),
		AffectedCpeUri:       dcl.StringOrNil(p.AffectedCpeUri),
		AffectedPackage:      dcl.StringOrNil(p.AffectedPackage),
		AffectedVersionStart: ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart(p.GetAffectedVersionStart()),
		AffectedVersionEnd:   ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd(p.GetAffectedVersionEnd()),
		FixedCpeUri:          dcl.StringOrNil(p.FixedCpeUri),
		FixedPackage:         dcl.StringOrNil(p.FixedPackage),
		FixedVersion:         ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion(p.GetFixedVersion()),
		IsObsolete:           dcl.Bool(p.IsObsolete),
		SourceUpdateTime:     dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStart converts a NoteVulnerabilityDetailsAffectedVersionStart resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart) *beta.NoteVulnerabilityDetailsAffectedVersionStart {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEnd converts a NoteVulnerabilityDetailsAffectedVersionEnd resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd) *beta.NoteVulnerabilityDetailsAffectedVersionEnd {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsFixedVersion converts a NoteVulnerabilityDetailsFixedVersion resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion) *beta.NoteVulnerabilityDetailsFixedVersion {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteVulnerabilityCvssV3 converts a NoteVulnerabilityCvssV3 resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3(p *betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3) *beta.NoteVulnerabilityCvssV3 {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityCvssV3{
		BaseScore:             dcl.Float64OrNil(p.BaseScore),
		ExploitabilityScore:   dcl.Float64OrNil(p.ExploitabilityScore),
		ImpactScore:           dcl.Float64OrNil(p.ImpactScore),
		AttackVector:          ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(p.GetAttackVector()),
		AttackComplexity:      ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(p.GetAttackComplexity()),
		PrivilegesRequired:    ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(p.GetPrivilegesRequired()),
		UserInteraction:       ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(p.GetUserInteraction()),
		Scope:                 ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(p.GetScope()),
		ConfidentialityImpact: ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(p.GetConfidentialityImpact()),
		IntegrityImpact:       ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(p.GetIntegrityImpact()),
		AvailabilityImpact:    ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(p.GetAvailabilityImpact()),
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetails converts a NoteVulnerabilityWindowsDetails resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetails(p *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails) *beta.NoteVulnerabilityWindowsDetails {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.StringOrNil(p.CpeUri),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
	}
	for _, r := range p.GetFixingKbs() {
		obj.FixingKbs = append(obj.FixingKbs, *ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetailsFixingKbs converts a NoteVulnerabilityWindowsDetailsFixingKbs resource from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs(p *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs) *beta.NoteVulnerabilityWindowsDetailsFixingKbs {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.StringOrNil(p.Name),
		Url:  dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToNoteBuild converts a NoteBuild resource from its proto representation.
func ProtoToContaineranalysisBetaNoteBuild(p *betapb.ContaineranalysisBetaNoteBuild) *beta.NoteBuild {
	if p == nil {
		return nil
	}
	obj := &beta.NoteBuild{
		BuilderVersion: dcl.StringOrNil(p.BuilderVersion),
		Signature:      ProtoToContaineranalysisBetaNoteBuildSignature(p.GetSignature()),
	}
	return obj
}

// ProtoToNoteBuildSignature converts a NoteBuildSignature resource from its proto representation.
func ProtoToContaineranalysisBetaNoteBuildSignature(p *betapb.ContaineranalysisBetaNoteBuildSignature) *beta.NoteBuildSignature {
	if p == nil {
		return nil
	}
	obj := &beta.NoteBuildSignature{
		PublicKey: dcl.StringOrNil(p.PublicKey),
		Signature: dcl.StringOrNil(p.Signature),
		KeyId:     dcl.StringOrNil(p.KeyId),
		KeyType:   ProtoToContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(p.GetKeyType()),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage resource from its proto representation.
func ProtoToContaineranalysisBetaNoteImage(p *betapb.ContaineranalysisBetaNoteImage) *beta.NoteImage {
	if p == nil {
		return nil
	}
	obj := &beta.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.ResourceUrl),
		Fingerprint: ProtoToContaineranalysisBetaNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint resource from its proto representation.
func ProtoToContaineranalysisBetaNoteImageFingerprint(p *betapb.ContaineranalysisBetaNoteImageFingerprint) *beta.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &beta.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.V1Name),
		V2Name: dcl.StringOrNil(p.V2Name),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage resource from its proto representation.
func ProtoToContaineranalysisBetaNotePackage(p *betapb.ContaineranalysisBetaNotePackage) *beta.NotePackage {
	if p == nil {
		return nil
	}
	obj := &beta.NotePackage{
		Name: dcl.StringOrNil(p.Name),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisBetaNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution resource from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistribution(p *betapb.ContaineranalysisBetaNotePackageDistribution) *beta.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &beta.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.CpeUri),
		Architecture:  ProtoToContaineranalysisBetaNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisBetaNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.Maintainer),
		Url:           dcl.StringOrNil(p.Url),
		Description:   dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion resource from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistributionLatestVersion(p *betapb.ContaineranalysisBetaNotePackageDistributionLatestVersion) *beta.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &beta.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery resource from its proto representation.
func ProtoToContaineranalysisBetaNoteDiscovery(p *betapb.ContaineranalysisBetaNoteDiscovery) *beta.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &beta.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNoteBaseImage converts a NoteBaseImage resource from its proto representation.
func ProtoToContaineranalysisBetaNoteBaseImage(p *betapb.ContaineranalysisBetaNoteBaseImage) *beta.NoteBaseImage {
	if p == nil {
		return nil
	}
	obj := &beta.NoteBaseImage{
		ResourceUrl: dcl.StringOrNil(p.ResourceUrl),
		Fingerprint: ProtoToContaineranalysisBetaNoteBaseImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteBaseImageFingerprint converts a NoteBaseImageFingerprint resource from its proto representation.
func ProtoToContaineranalysisBetaNoteBaseImageFingerprint(p *betapb.ContaineranalysisBetaNoteBaseImageFingerprint) *beta.NoteBaseImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &beta.NoteBaseImageFingerprint{
		V1Name: dcl.StringOrNil(p.V1Name),
		V2Name: dcl.StringOrNil(p.V2Name),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNoteDeployable converts a NoteDeployable resource from its proto representation.
func ProtoToContaineranalysisBetaNoteDeployable(p *betapb.ContaineranalysisBetaNoteDeployable) *beta.NoteDeployable {
	if p == nil {
		return nil
	}
	obj := &beta.NoteDeployable{}
	for _, r := range p.GetResourceUri() {
		obj.ResourceUri = append(obj.ResourceUri, r)
	}
	return obj
}

// ProtoToNoteAttestationAuthority converts a NoteAttestationAuthority resource from its proto representation.
func ProtoToContaineranalysisBetaNoteAttestationAuthority(p *betapb.ContaineranalysisBetaNoteAttestationAuthority) *beta.NoteAttestationAuthority {
	if p == nil {
		return nil
	}
	obj := &beta.NoteAttestationAuthority{
		Hint: ProtoToContaineranalysisBetaNoteAttestationAuthorityHint(p.GetHint()),
	}
	return obj
}

// ProtoToNoteAttestationAuthorityHint converts a NoteAttestationAuthorityHint resource from its proto representation.
func ProtoToContaineranalysisBetaNoteAttestationAuthorityHint(p *betapb.ContaineranalysisBetaNoteAttestationAuthorityHint) *beta.NoteAttestationAuthorityHint {
	if p == nil {
		return nil
	}
	obj := &beta.NoteAttestationAuthorityHint{
		HumanReadableName: dcl.StringOrNil(p.HumanReadableName),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *betapb.ContaineranalysisBetaNote) *beta.Note {
	obj := &beta.Note{
		Name:                 dcl.StringOrNil(p.Name),
		ShortDescription:     dcl.StringOrNil(p.ShortDescription),
		LongDescription:      dcl.StringOrNil(p.LongDescription),
		ExpirationTime:       dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:           dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Vulnerability:        ProtoToContaineranalysisBetaNoteVulnerability(p.GetVulnerability()),
		Build:                ProtoToContaineranalysisBetaNoteBuild(p.GetBuild()),
		Image:                ProtoToContaineranalysisBetaNoteImage(p.GetImage()),
		Package:              ProtoToContaineranalysisBetaNotePackage(p.GetPackage()),
		Discovery:            ProtoToContaineranalysisBetaNoteDiscovery(p.GetDiscovery()),
		BaseImage:            ProtoToContaineranalysisBetaNoteBaseImage(p.GetBaseImage()),
		Deployable:           ProtoToContaineranalysisBetaNoteDeployable(p.GetDeployable()),
		AttestationAuthority: ProtoToContaineranalysisBetaNoteAttestationAuthority(p.GetAttestationAuthority()),
		Project:              dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetRelatedUrl() {
		obj.RelatedUrl = append(obj.RelatedUrl, *ProtoToContaineranalysisBetaNoteRelatedUrl(r))
	}
	for _, r := range p.GetRelatedNoteNames() {
		obj.RelatedNoteNames = append(obj.RelatedNoteNames, r)
	}
	return obj
}

// NoteVulnerabilitySeverityEnumToProto converts a NoteVulnerabilitySeverityEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilitySeverityEnumToProto(e *beta.NoteVulnerabilitySeverityEnum) betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum_value["NoteVulnerabilitySeverityEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(e *beta.NoteVulnerabilityDetailsAffectedVersionStartKindEnum) betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum_value["NoteVulnerabilityDetailsAffectedVersionStartKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(e *beta.NoteVulnerabilityDetailsAffectedVersionEndKindEnum) betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum_value["NoteVulnerabilityDetailsAffectedVersionEndKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
}

// NoteVulnerabilityDetailsFixedVersionKindEnumToProto converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnumToProto(e *beta.NoteVulnerabilityDetailsFixedVersionKindEnum) betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum_value["NoteVulnerabilityDetailsFixedVersionKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(0)
}

// NoteVulnerabilityCvssV3AttackVectorEnumToProto converts a NoteVulnerabilityCvssV3AttackVectorEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnumToProto(e *beta.NoteVulnerabilityCvssV3AttackVectorEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum_value["NoteVulnerabilityCvssV3AttackVectorEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(0)
}

// NoteVulnerabilityCvssV3AttackComplexityEnumToProto converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnumToProto(e *beta.NoteVulnerabilityCvssV3AttackComplexityEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum_value["NoteVulnerabilityCvssV3AttackComplexityEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(0)
}

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(e *beta.NoteVulnerabilityCvssV3PrivilegesRequiredEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum_value["NoteVulnerabilityCvssV3PrivilegesRequiredEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
}

// NoteVulnerabilityCvssV3UserInteractionEnumToProto converts a NoteVulnerabilityCvssV3UserInteractionEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnumToProto(e *beta.NoteVulnerabilityCvssV3UserInteractionEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum_value["NoteVulnerabilityCvssV3UserInteractionEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(0)
}

// NoteVulnerabilityCvssV3ScopeEnumToProto converts a NoteVulnerabilityCvssV3ScopeEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnumToProto(e *beta.NoteVulnerabilityCvssV3ScopeEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum_value["NoteVulnerabilityCvssV3ScopeEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(0)
}

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(e *beta.NoteVulnerabilityCvssV3ConfidentialityImpactEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum_value["NoteVulnerabilityCvssV3ConfidentialityImpactEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
}

// NoteVulnerabilityCvssV3IntegrityImpactEnumToProto converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(e *beta.NoteVulnerabilityCvssV3IntegrityImpactEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum_value["NoteVulnerabilityCvssV3IntegrityImpactEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
}

// NoteVulnerabilityCvssV3AvailabilityImpactEnumToProto converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(e *beta.NoteVulnerabilityCvssV3AvailabilityImpactEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum_value["NoteVulnerabilityCvssV3AvailabilityImpactEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
}

// NoteBuildSignatureKeyTypeEnumToProto converts a NoteBuildSignatureKeyTypeEnum enum to its proto representation.
func ContaineranalysisBetaNoteBuildSignatureKeyTypeEnumToProto(e *beta.NoteBuildSignatureKeyTypeEnum) betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum_value["NoteBuildSignatureKeyTypeEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(0)
}

// NotePackageDistributionArchitectureEnumToProto converts a NotePackageDistributionArchitectureEnum enum to its proto representation.
func ContaineranalysisBetaNotePackageDistributionArchitectureEnumToProto(e *beta.NotePackageDistributionArchitectureEnum) betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum_value["NotePackageDistributionArchitectureEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum(v)
	}
	return betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum(0)
}

// NotePackageDistributionLatestVersionKindEnumToProto converts a NotePackageDistributionLatestVersionKindEnum enum to its proto representation.
func ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnumToProto(e *beta.NotePackageDistributionLatestVersionKindEnum) betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum_value["NotePackageDistributionLatestVersionKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(0)
}

// NoteDiscoveryAnalysisKindEnumToProto converts a NoteDiscoveryAnalysisKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteDiscoveryAnalysisKindEnumToProto(e *beta.NoteDiscoveryAnalysisKindEnum) betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum_value["NoteDiscoveryAnalysisKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(0)
}

// NoteRelatedUrlToProto converts a NoteRelatedUrl resource to its proto representation.
func ContaineranalysisBetaNoteRelatedUrlToProto(o *beta.NoteRelatedUrl) *betapb.ContaineranalysisBetaNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteRelatedUrl{
		Url:   dcl.ValueOrEmptyString(o.Url),
		Label: dcl.ValueOrEmptyString(o.Label),
	}
	return p
}

// NoteVulnerabilityToProto converts a NoteVulnerability resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityToProto(o *beta.NoteVulnerability) *betapb.ContaineranalysisBetaNoteVulnerability {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerability{
		CvssScore:        dcl.ValueOrEmptyDouble(o.CvssScore),
		Severity:         ContaineranalysisBetaNoteVulnerabilitySeverityEnumToProto(o.Severity),
		CvssV3:           ContaineranalysisBetaNoteVulnerabilityCvssV3ToProto(o.CvssV3),
		SourceUpdateTime: dcl.ValueOrEmptyString(o.SourceUpdateTime),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, ContaineranalysisBetaNoteVulnerabilityDetailsToProto(&r))
	}
	for _, r := range o.WindowsDetails {
		p.WindowsDetails = append(p.WindowsDetails, ContaineranalysisBetaNoteVulnerabilityWindowsDetailsToProto(&r))
	}
	return p
}

// NoteVulnerabilityDetailsToProto converts a NoteVulnerabilityDetails resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsToProto(o *beta.NoteVulnerabilityDetails) *betapb.ContaineranalysisBetaNoteVulnerabilityDetails {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetails{
		SeverityName:         dcl.ValueOrEmptyString(o.SeverityName),
		Description:          dcl.ValueOrEmptyString(o.Description),
		PackageType:          dcl.ValueOrEmptyString(o.PackageType),
		AffectedCpeUri:       dcl.ValueOrEmptyString(o.AffectedCpeUri),
		AffectedPackage:      dcl.ValueOrEmptyString(o.AffectedPackage),
		AffectedVersionStart: ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartToProto(o.AffectedVersionStart),
		AffectedVersionEnd:   ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndToProto(o.AffectedVersionEnd),
		FixedCpeUri:          dcl.ValueOrEmptyString(o.FixedCpeUri),
		FixedPackage:         dcl.ValueOrEmptyString(o.FixedPackage),
		FixedVersion:         ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionToProto(o.FixedVersion),
		IsObsolete:           dcl.ValueOrEmptyBool(o.IsObsolete),
		SourceUpdateTime:     dcl.ValueOrEmptyString(o.SourceUpdateTime),
	}
	return p
}

// NoteVulnerabilityDetailsAffectedVersionStartToProto converts a NoteVulnerabilityDetailsAffectedVersionStart resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartToProto(o *beta.NoteVulnerabilityDetailsAffectedVersionStart) *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteVulnerabilityDetailsAffectedVersionEndToProto converts a NoteVulnerabilityDetailsAffectedVersionEnd resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndToProto(o *beta.NoteVulnerabilityDetailsAffectedVersionEnd) *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteVulnerabilityDetailsFixedVersionToProto converts a NoteVulnerabilityDetailsFixedVersion resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionToProto(o *beta.NoteVulnerabilityDetailsFixedVersion) *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteVulnerabilityCvssV3ToProto converts a NoteVulnerabilityCvssV3 resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3ToProto(o *beta.NoteVulnerabilityCvssV3) *betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3 {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3{
		BaseScore:             dcl.ValueOrEmptyDouble(o.BaseScore),
		ExploitabilityScore:   dcl.ValueOrEmptyDouble(o.ExploitabilityScore),
		ImpactScore:           dcl.ValueOrEmptyDouble(o.ImpactScore),
		AttackVector:          ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnumToProto(o.AttackVector),
		AttackComplexity:      ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnumToProto(o.AttackComplexity),
		PrivilegesRequired:    ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(o.PrivilegesRequired),
		UserInteraction:       ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnumToProto(o.UserInteraction),
		Scope:                 ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnumToProto(o.Scope),
		ConfidentialityImpact: ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(o.ConfidentialityImpact),
		IntegrityImpact:       ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(o.IntegrityImpact),
		AvailabilityImpact:    ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(o.AvailabilityImpact),
	}
	return p
}

// NoteVulnerabilityWindowsDetailsToProto converts a NoteVulnerabilityWindowsDetails resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityWindowsDetailsToProto(o *beta.NoteVulnerabilityWindowsDetails) *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.ValueOrEmptyString(o.CpeUri),
		Name:        dcl.ValueOrEmptyString(o.Name),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	for _, r := range o.FixingKbs {
		p.FixingKbs = append(p.FixingKbs, ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbsToProto(&r))
	}
	return p
}

// NoteVulnerabilityWindowsDetailsFixingKbsToProto converts a NoteVulnerabilityWindowsDetailsFixingKbs resource to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbsToProto(o *beta.NoteVulnerabilityWindowsDetailsFixingKbs) *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.ValueOrEmptyString(o.Name),
		Url:  dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// NoteBuildToProto converts a NoteBuild resource to its proto representation.
func ContaineranalysisBetaNoteBuildToProto(o *beta.NoteBuild) *betapb.ContaineranalysisBetaNoteBuild {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteBuild{
		BuilderVersion: dcl.ValueOrEmptyString(o.BuilderVersion),
		Signature:      ContaineranalysisBetaNoteBuildSignatureToProto(o.Signature),
	}
	return p
}

// NoteBuildSignatureToProto converts a NoteBuildSignature resource to its proto representation.
func ContaineranalysisBetaNoteBuildSignatureToProto(o *beta.NoteBuildSignature) *betapb.ContaineranalysisBetaNoteBuildSignature {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteBuildSignature{
		PublicKey: dcl.ValueOrEmptyString(o.PublicKey),
		Signature: dcl.ValueOrEmptyString(o.Signature),
		KeyId:     dcl.ValueOrEmptyString(o.KeyId),
		KeyType:   ContaineranalysisBetaNoteBuildSignatureKeyTypeEnumToProto(o.KeyType),
	}
	return p
}

// NoteImageToProto converts a NoteImage resource to its proto representation.
func ContaineranalysisBetaNoteImageToProto(o *beta.NoteImage) *betapb.ContaineranalysisBetaNoteImage {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteImage{
		ResourceUrl: dcl.ValueOrEmptyString(o.ResourceUrl),
		Fingerprint: ContaineranalysisBetaNoteImageFingerprintToProto(o.Fingerprint),
	}
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint resource to its proto representation.
func ContaineranalysisBetaNoteImageFingerprintToProto(o *beta.NoteImageFingerprint) *betapb.ContaineranalysisBetaNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteImageFingerprint{
		V1Name: dcl.ValueOrEmptyString(o.V1Name),
		V2Name: dcl.ValueOrEmptyString(o.V2Name),
	}
	for _, r := range o.V2Blob {
		p.V2Blob = append(p.V2Blob, r)
	}
	return p
}

// NotePackageToProto converts a NotePackage resource to its proto representation.
func ContaineranalysisBetaNotePackageToProto(o *beta.NotePackage) *betapb.ContaineranalysisBetaNotePackage {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNotePackage{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	for _, r := range o.Distribution {
		p.Distribution = append(p.Distribution, ContaineranalysisBetaNotePackageDistributionToProto(&r))
	}
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution resource to its proto representation.
func ContaineranalysisBetaNotePackageDistributionToProto(o *beta.NotePackageDistribution) *betapb.ContaineranalysisBetaNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNotePackageDistribution{
		CpeUri:        dcl.ValueOrEmptyString(o.CpeUri),
		Architecture:  ContaineranalysisBetaNotePackageDistributionArchitectureEnumToProto(o.Architecture),
		LatestVersion: ContaineranalysisBetaNotePackageDistributionLatestVersionToProto(o.LatestVersion),
		Maintainer:    dcl.ValueOrEmptyString(o.Maintainer),
		Url:           dcl.ValueOrEmptyString(o.Url),
		Description:   dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion resource to its proto representation.
func ContaineranalysisBetaNotePackageDistributionLatestVersionToProto(o *beta.NotePackageDistributionLatestVersion) *betapb.ContaineranalysisBetaNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNotePackageDistributionLatestVersion{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery resource to its proto representation.
func ContaineranalysisBetaNoteDiscoveryToProto(o *beta.NoteDiscovery) *betapb.ContaineranalysisBetaNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteDiscovery{
		AnalysisKind: ContaineranalysisBetaNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind),
	}
	return p
}

// NoteBaseImageToProto converts a NoteBaseImage resource to its proto representation.
func ContaineranalysisBetaNoteBaseImageToProto(o *beta.NoteBaseImage) *betapb.ContaineranalysisBetaNoteBaseImage {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteBaseImage{
		ResourceUrl: dcl.ValueOrEmptyString(o.ResourceUrl),
		Fingerprint: ContaineranalysisBetaNoteBaseImageFingerprintToProto(o.Fingerprint),
	}
	return p
}

// NoteBaseImageFingerprintToProto converts a NoteBaseImageFingerprint resource to its proto representation.
func ContaineranalysisBetaNoteBaseImageFingerprintToProto(o *beta.NoteBaseImageFingerprint) *betapb.ContaineranalysisBetaNoteBaseImageFingerprint {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteBaseImageFingerprint{
		V1Name: dcl.ValueOrEmptyString(o.V1Name),
		V2Name: dcl.ValueOrEmptyString(o.V2Name),
	}
	for _, r := range o.V2Blob {
		p.V2Blob = append(p.V2Blob, r)
	}
	return p
}

// NoteDeployableToProto converts a NoteDeployable resource to its proto representation.
func ContaineranalysisBetaNoteDeployableToProto(o *beta.NoteDeployable) *betapb.ContaineranalysisBetaNoteDeployable {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteDeployable{}
	for _, r := range o.ResourceUri {
		p.ResourceUri = append(p.ResourceUri, r)
	}
	return p
}

// NoteAttestationAuthorityToProto converts a NoteAttestationAuthority resource to its proto representation.
func ContaineranalysisBetaNoteAttestationAuthorityToProto(o *beta.NoteAttestationAuthority) *betapb.ContaineranalysisBetaNoteAttestationAuthority {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteAttestationAuthority{
		Hint: ContaineranalysisBetaNoteAttestationAuthorityHintToProto(o.Hint),
	}
	return p
}

// NoteAttestationAuthorityHintToProto converts a NoteAttestationAuthorityHint resource to its proto representation.
func ContaineranalysisBetaNoteAttestationAuthorityHintToProto(o *beta.NoteAttestationAuthorityHint) *betapb.ContaineranalysisBetaNoteAttestationAuthorityHint {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteAttestationAuthorityHint{
		HumanReadableName: dcl.ValueOrEmptyString(o.HumanReadableName),
	}
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *beta.Note) *betapb.ContaineranalysisBetaNote {
	p := &betapb.ContaineranalysisBetaNote{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		ShortDescription:     dcl.ValueOrEmptyString(resource.ShortDescription),
		LongDescription:      dcl.ValueOrEmptyString(resource.LongDescription),
		ExpirationTime:       dcl.ValueOrEmptyString(resource.ExpirationTime),
		CreateTime:           dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:           dcl.ValueOrEmptyString(resource.UpdateTime),
		Vulnerability:        ContaineranalysisBetaNoteVulnerabilityToProto(resource.Vulnerability),
		Build:                ContaineranalysisBetaNoteBuildToProto(resource.Build),
		Image:                ContaineranalysisBetaNoteImageToProto(resource.Image),
		Package:              ContaineranalysisBetaNotePackageToProto(resource.Package),
		Discovery:            ContaineranalysisBetaNoteDiscoveryToProto(resource.Discovery),
		BaseImage:            ContaineranalysisBetaNoteBaseImageToProto(resource.BaseImage),
		Deployable:           ContaineranalysisBetaNoteDeployableToProto(resource.Deployable),
		AttestationAuthority: ContaineranalysisBetaNoteAttestationAuthorityToProto(resource.AttestationAuthority),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.RelatedUrl {
		p.RelatedUrl = append(p.RelatedUrl, ContaineranalysisBetaNoteRelatedUrlToProto(&r))
	}
	for _, r := range resource.RelatedNoteNames {
		p.RelatedNoteNames = append(p.RelatedNoteNames, r)
	}

	return p
}

// ApplyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) applyNote(ctx context.Context, c *beta.Client, request *betapb.ApplyContaineranalysisBetaNoteRequest) (*betapb.ContaineranalysisBetaNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// ApplyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisBetaNote(ctx context.Context, request *betapb.ApplyContaineranalysisBetaNoteRequest) (*betapb.ContaineranalysisBetaNote, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisBetaNote(ctx context.Context, request *betapb.DeleteContaineranalysisBetaNoteRequest) (*emptypb.Empty, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))
}

// ListNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisBetaNote(ctx context.Context, request *betapb.ListContaineranalysisBetaNoteRequest) (*betapb.ListContaineranalysisBetaNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContaineranalysisBetaNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListContaineranalysisBetaNoteResponse{Items: protos}, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
