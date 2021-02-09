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
	containeranalysispb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeranalysis/containeranalysis_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis"
)

// Server implements the gRPC interface for Note.
type NoteServer struct{}

// ProtoToNoteVulnerabilitySeverityEnum converts a NoteVulnerabilitySeverityEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilitySeverityEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum) *containeranalysis.NoteVulnerabilitySeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilitySeverityEnum(n[len("NoteVulnerabilitySeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStartKindEnum converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityDetailsAffectedVersionStartKindEnum(n[len("NoteVulnerabilityDetailsAffectedVersionStartKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEndKindEnum converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityDetailsAffectedVersionEndKindEnum(n[len("NoteVulnerabilityDetailsAffectedVersionEndKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsFixedVersionKindEnum converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum) *containeranalysis.NoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityDetailsFixedVersionKindEnum(n[len("NoteVulnerabilityDetailsFixedVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackVectorEnum converts a NoteVulnerabilityCvssV3AttackVectorEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum) *containeranalysis.NoteVulnerabilityCvssV3AttackVectorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3AttackVectorEnum(n[len("NoteVulnerabilityCvssV3AttackVectorEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackComplexityEnum converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum) *containeranalysis.NoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3AttackComplexityEnum(n[len("NoteVulnerabilityCvssV3AttackComplexityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3PrivilegesRequiredEnum converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum) *containeranalysis.NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3PrivilegesRequiredEnum(n[len("NoteVulnerabilityCvssV3PrivilegesRequiredEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3UserInteractionEnum converts a NoteVulnerabilityCvssV3UserInteractionEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum) *containeranalysis.NoteVulnerabilityCvssV3UserInteractionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3UserInteractionEnum(n[len("NoteVulnerabilityCvssV3UserInteractionEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ScopeEnum converts a NoteVulnerabilityCvssV3ScopeEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum) *containeranalysis.NoteVulnerabilityCvssV3ScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3ScopeEnum(n[len("NoteVulnerabilityCvssV3ScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ConfidentialityImpactEnum converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum) *containeranalysis.NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3ConfidentialityImpactEnum(n[len("NoteVulnerabilityCvssV3ConfidentialityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3IntegrityImpactEnum converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum) *containeranalysis.NoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3IntegrityImpactEnum(n[len("NoteVulnerabilityCvssV3IntegrityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AvailabilityImpactEnum converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum) *containeranalysis.NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3AvailabilityImpactEnum(n[len("NoteVulnerabilityCvssV3AvailabilityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionArchitectureEnum converts a NotePackageDistributionArchitectureEnum enum from its proto representation.
func ProtoToContaineranalysisNotePackageDistributionArchitectureEnum(e containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum) *containeranalysis.NotePackageDistributionArchitectureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum_name[int32(e)]; ok {
		e := containeranalysis.NotePackageDistributionArchitectureEnum(n[len("NotePackageDistributionArchitectureEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionLatestVersionKindEnum converts a NotePackageDistributionLatestVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisNotePackageDistributionLatestVersionKindEnum(e containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum) *containeranalysis.NotePackageDistributionLatestVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NotePackageDistributionLatestVersionKindEnum(n[len("NotePackageDistributionLatestVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteDiscoveryAnalysisKindEnum converts a NoteDiscoveryAnalysisKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteDiscoveryAnalysisKindEnum(e containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum) *containeranalysis.NoteDiscoveryAnalysisKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteDiscoveryAnalysisKindEnum(n[len("NoteDiscoveryAnalysisKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl resource from its proto representation.
func ProtoToContaineranalysisNoteRelatedUrl(p *containeranalysispb.ContaineranalysisNoteRelatedUrl) *containeranalysis.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.Url),
		Label: dcl.StringOrNil(p.Label),
	}
	return obj
}

// ProtoToNoteVulnerability converts a NoteVulnerability resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerability(p *containeranalysispb.ContaineranalysisNoteVulnerability) *containeranalysis.NoteVulnerability {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerability{
		CvssScore:        dcl.Float64OrNil(p.CvssScore),
		Severity:         ProtoToContaineranalysisNoteVulnerabilitySeverityEnum(p.GetSeverity()),
		CvssV3:           ProtoToContaineranalysisNoteVulnerabilityCvssV3(p.GetCvssV3()),
		SourceUpdateTime: dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToContaineranalysisNoteVulnerabilityDetails(r))
	}
	for _, r := range p.GetWindowsDetails() {
		obj.WindowsDetails = append(obj.WindowsDetails, *ProtoToContaineranalysisNoteVulnerabilityWindowsDetails(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityDetails converts a NoteVulnerabilityDetails resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetails(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetails) *containeranalysis.NoteVulnerabilityDetails {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetails{
		SeverityName:         dcl.StringOrNil(p.SeverityName),
		Description:          dcl.StringOrNil(p.Description),
		PackageType:          dcl.StringOrNil(p.PackageType),
		AffectedCpeUri:       dcl.StringOrNil(p.AffectedCpeUri),
		AffectedPackage:      dcl.StringOrNil(p.AffectedPackage),
		AffectedVersionStart: ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart(p.GetAffectedVersionStart()),
		AffectedVersionEnd:   ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd(p.GetAffectedVersionEnd()),
		FixedCpeUri:          dcl.StringOrNil(p.FixedCpeUri),
		FixedPackage:         dcl.StringOrNil(p.FixedPackage),
		FixedVersion:         ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersion(p.GetFixedVersion()),
		IsObsolete:           dcl.Bool(p.IsObsolete),
		SourceUpdateTime:     dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStart converts a NoteVulnerabilityDetailsAffectedVersionStart resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStart {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEnd converts a NoteVulnerabilityDetailsAffectedVersionEnd resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEnd {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsFixedVersion converts a NoteVulnerabilityDetailsFixedVersion resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersion(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersion) *containeranalysis.NoteVulnerabilityDetailsFixedVersion {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteVulnerabilityCvssV3 converts a NoteVulnerabilityCvssV3 resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3(p *containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3) *containeranalysis.NoteVulnerabilityCvssV3 {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityCvssV3{
		BaseScore:             dcl.Float64OrNil(p.BaseScore),
		ExploitabilityScore:   dcl.Float64OrNil(p.ExploitabilityScore),
		ImpactScore:           dcl.Float64OrNil(p.ImpactScore),
		AttackVector:          ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(p.GetAttackVector()),
		AttackComplexity:      ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(p.GetAttackComplexity()),
		PrivilegesRequired:    ProtoToContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(p.GetPrivilegesRequired()),
		UserInteraction:       ProtoToContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(p.GetUserInteraction()),
		Scope:                 ProtoToContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(p.GetScope()),
		ConfidentialityImpact: ProtoToContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(p.GetConfidentialityImpact()),
		IntegrityImpact:       ProtoToContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(p.GetIntegrityImpact()),
		AvailabilityImpact:    ProtoToContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(p.GetAvailabilityImpact()),
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetails converts a NoteVulnerabilityWindowsDetails resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityWindowsDetails(p *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails) *containeranalysis.NoteVulnerabilityWindowsDetails {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.StringOrNil(p.CpeUri),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
	}
	for _, r := range p.GetFixingKbs() {
		obj.FixingKbs = append(obj.FixingKbs, *ProtoToContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetailsFixingKbs converts a NoteVulnerabilityWindowsDetailsFixingKbs resource from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs(p *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs) *containeranalysis.NoteVulnerabilityWindowsDetailsFixingKbs {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.StringOrNil(p.Name),
		Url:  dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToNoteBuild converts a NoteBuild resource from its proto representation.
func ProtoToContaineranalysisNoteBuild(p *containeranalysispb.ContaineranalysisNoteBuild) *containeranalysis.NoteBuild {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteBuild{
		BuilderVersion: dcl.StringOrNil(p.BuilderVersion),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage resource from its proto representation.
func ProtoToContaineranalysisNoteImage(p *containeranalysispb.ContaineranalysisNoteImage) *containeranalysis.NoteImage {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.ResourceUrl),
		Fingerprint: ProtoToContaineranalysisNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint resource from its proto representation.
func ProtoToContaineranalysisNoteImageFingerprint(p *containeranalysispb.ContaineranalysisNoteImageFingerprint) *containeranalysis.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.V1Name),
		V2Name: dcl.StringOrNil(p.V2Name),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage resource from its proto representation.
func ProtoToContaineranalysisNotePackage(p *containeranalysispb.ContaineranalysisNotePackage) *containeranalysis.NotePackage {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NotePackage{
		Name: dcl.StringOrNil(p.Name),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution resource from its proto representation.
func ProtoToContaineranalysisNotePackageDistribution(p *containeranalysispb.ContaineranalysisNotePackageDistribution) *containeranalysis.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.CpeUri),
		Architecture:  ProtoToContaineranalysisNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.Maintainer),
		Url:           dcl.StringOrNil(p.Url),
		Description:   dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion resource from its proto representation.
func ProtoToContaineranalysisNotePackageDistributionLatestVersion(p *containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersion) *containeranalysis.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.Epoch),
		Name:     dcl.StringOrNil(p.Name),
		Revision: dcl.StringOrNil(p.Revision),
		Kind:     ProtoToContaineranalysisNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.FullName),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery resource from its proto representation.
func ProtoToContaineranalysisNoteDiscovery(p *containeranalysispb.ContaineranalysisNoteDiscovery) *containeranalysis.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *containeranalysispb.ContaineranalysisNote) *containeranalysis.Note {
	obj := &containeranalysis.Note{
		Name:             dcl.StringOrNil(p.Name),
		ShortDescription: dcl.StringOrNil(p.ShortDescription),
		LongDescription:  dcl.StringOrNil(p.LongDescription),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Vulnerability:    ProtoToContaineranalysisNoteVulnerability(p.GetVulnerability()),
		Build:            ProtoToContaineranalysisNoteBuild(p.GetBuild()),
		Image:            ProtoToContaineranalysisNoteImage(p.GetImage()),
		Package:          ProtoToContaineranalysisNotePackage(p.GetPackage()),
		Discovery:        ProtoToContaineranalysisNoteDiscovery(p.GetDiscovery()),
		Project:          dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetRelatedUrl() {
		obj.RelatedUrl = append(obj.RelatedUrl, *ProtoToContaineranalysisNoteRelatedUrl(r))
	}
	for _, r := range p.GetRelatedNoteNames() {
		obj.RelatedNoteNames = append(obj.RelatedNoteNames, r)
	}
	return obj
}

// NoteVulnerabilitySeverityEnumToProto converts a NoteVulnerabilitySeverityEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilitySeverityEnumToProto(e *containeranalysis.NoteVulnerabilitySeverityEnum) containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum_value["NoteVulnerabilitySeverityEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(e *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStartKindEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum_value["NoteVulnerabilityDetailsAffectedVersionStartKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(e *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEndKindEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum_value["NoteVulnerabilityDetailsAffectedVersionEndKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
}

// NoteVulnerabilityDetailsFixedVersionKindEnumToProto converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnumToProto(e *containeranalysis.NoteVulnerabilityDetailsFixedVersionKindEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum_value["NoteVulnerabilityDetailsFixedVersionKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(0)
}

// NoteVulnerabilityCvssV3AttackVectorEnumToProto converts a NoteVulnerabilityCvssV3AttackVectorEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3AttackVectorEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum_value["NoteVulnerabilityCvssV3AttackVectorEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(0)
}

// NoteVulnerabilityCvssV3AttackComplexityEnumToProto converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3AttackComplexityEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum_value["NoteVulnerabilityCvssV3AttackComplexityEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(0)
}

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3PrivilegesRequiredEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum_value["NoteVulnerabilityCvssV3PrivilegesRequiredEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
}

// NoteVulnerabilityCvssV3UserInteractionEnumToProto converts a NoteVulnerabilityCvssV3UserInteractionEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3UserInteractionEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum_value["NoteVulnerabilityCvssV3UserInteractionEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(0)
}

// NoteVulnerabilityCvssV3ScopeEnumToProto converts a NoteVulnerabilityCvssV3ScopeEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3ScopeEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3ScopeEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum_value["NoteVulnerabilityCvssV3ScopeEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(0)
}

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3ConfidentialityImpactEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum_value["NoteVulnerabilityCvssV3ConfidentialityImpactEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
}

// NoteVulnerabilityCvssV3IntegrityImpactEnumToProto converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3IntegrityImpactEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum_value["NoteVulnerabilityCvssV3IntegrityImpactEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
}

// NoteVulnerabilityCvssV3AvailabilityImpactEnumToProto converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3AvailabilityImpactEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum_value["NoteVulnerabilityCvssV3AvailabilityImpactEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
}

// NotePackageDistributionArchitectureEnumToProto converts a NotePackageDistributionArchitectureEnum enum to its proto representation.
func ContaineranalysisNotePackageDistributionArchitectureEnumToProto(e *containeranalysis.NotePackageDistributionArchitectureEnum) containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum_value["NotePackageDistributionArchitectureEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum(v)
	}
	return containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum(0)
}

// NotePackageDistributionLatestVersionKindEnumToProto converts a NotePackageDistributionLatestVersionKindEnum enum to its proto representation.
func ContaineranalysisNotePackageDistributionLatestVersionKindEnumToProto(e *containeranalysis.NotePackageDistributionLatestVersionKindEnum) containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum_value["NotePackageDistributionLatestVersionKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum(0)
}

// NoteDiscoveryAnalysisKindEnumToProto converts a NoteDiscoveryAnalysisKindEnum enum to its proto representation.
func ContaineranalysisNoteDiscoveryAnalysisKindEnumToProto(e *containeranalysis.NoteDiscoveryAnalysisKindEnum) containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum_value["NoteDiscoveryAnalysisKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum(0)
}

// NoteRelatedUrlToProto converts a NoteRelatedUrl resource to its proto representation.
func ContaineranalysisNoteRelatedUrlToProto(o *containeranalysis.NoteRelatedUrl) *containeranalysispb.ContaineranalysisNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteRelatedUrl{
		Url:   dcl.ValueOrEmptyString(o.Url),
		Label: dcl.ValueOrEmptyString(o.Label),
	}
	return p
}

// NoteVulnerabilityToProto converts a NoteVulnerability resource to its proto representation.
func ContaineranalysisNoteVulnerabilityToProto(o *containeranalysis.NoteVulnerability) *containeranalysispb.ContaineranalysisNoteVulnerability {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerability{
		CvssScore:        dcl.ValueOrEmptyDouble(o.CvssScore),
		Severity:         ContaineranalysisNoteVulnerabilitySeverityEnumToProto(o.Severity),
		CvssV3:           ContaineranalysisNoteVulnerabilityCvssV3ToProto(o.CvssV3),
		SourceUpdateTime: dcl.ValueOrEmptyString(o.SourceUpdateTime),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, ContaineranalysisNoteVulnerabilityDetailsToProto(&r))
	}
	for _, r := range o.WindowsDetails {
		p.WindowsDetails = append(p.WindowsDetails, ContaineranalysisNoteVulnerabilityWindowsDetailsToProto(&r))
	}
	return p
}

// NoteVulnerabilityDetailsToProto converts a NoteVulnerabilityDetails resource to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsToProto(o *containeranalysis.NoteVulnerabilityDetails) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetails {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetails{
		SeverityName:         dcl.ValueOrEmptyString(o.SeverityName),
		Description:          dcl.ValueOrEmptyString(o.Description),
		PackageType:          dcl.ValueOrEmptyString(o.PackageType),
		AffectedCpeUri:       dcl.ValueOrEmptyString(o.AffectedCpeUri),
		AffectedPackage:      dcl.ValueOrEmptyString(o.AffectedPackage),
		AffectedVersionStart: ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartToProto(o.AffectedVersionStart),
		AffectedVersionEnd:   ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndToProto(o.AffectedVersionEnd),
		FixedCpeUri:          dcl.ValueOrEmptyString(o.FixedCpeUri),
		FixedPackage:         dcl.ValueOrEmptyString(o.FixedPackage),
		FixedVersion:         ContaineranalysisNoteVulnerabilityDetailsFixedVersionToProto(o.FixedVersion),
		IsObsolete:           dcl.ValueOrEmptyBool(o.IsObsolete),
		SourceUpdateTime:     dcl.ValueOrEmptyString(o.SourceUpdateTime),
	}
	return p
}

// NoteVulnerabilityDetailsAffectedVersionStartToProto converts a NoteVulnerabilityDetailsAffectedVersionStart resource to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartToProto(o *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStart) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteVulnerabilityDetailsAffectedVersionEndToProto converts a NoteVulnerabilityDetailsAffectedVersionEnd resource to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndToProto(o *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEnd) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteVulnerabilityDetailsFixedVersionToProto converts a NoteVulnerabilityDetailsFixedVersion resource to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsFixedVersionToProto(o *containeranalysis.NoteVulnerabilityDetailsFixedVersion) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersion {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteVulnerabilityCvssV3ToProto converts a NoteVulnerabilityCvssV3 resource to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3ToProto(o *containeranalysis.NoteVulnerabilityCvssV3) *containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3 {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3{
		BaseScore:             dcl.ValueOrEmptyDouble(o.BaseScore),
		ExploitabilityScore:   dcl.ValueOrEmptyDouble(o.ExploitabilityScore),
		ImpactScore:           dcl.ValueOrEmptyDouble(o.ImpactScore),
		AttackVector:          ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnumToProto(o.AttackVector),
		AttackComplexity:      ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnumToProto(o.AttackComplexity),
		PrivilegesRequired:    ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(o.PrivilegesRequired),
		UserInteraction:       ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnumToProto(o.UserInteraction),
		Scope:                 ContaineranalysisNoteVulnerabilityCvssV3ScopeEnumToProto(o.Scope),
		ConfidentialityImpact: ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(o.ConfidentialityImpact),
		IntegrityImpact:       ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(o.IntegrityImpact),
		AvailabilityImpact:    ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(o.AvailabilityImpact),
	}
	return p
}

// NoteVulnerabilityWindowsDetailsToProto converts a NoteVulnerabilityWindowsDetails resource to its proto representation.
func ContaineranalysisNoteVulnerabilityWindowsDetailsToProto(o *containeranalysis.NoteVulnerabilityWindowsDetails) *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.ValueOrEmptyString(o.CpeUri),
		Name:        dcl.ValueOrEmptyString(o.Name),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	for _, r := range o.FixingKbs {
		p.FixingKbs = append(p.FixingKbs, ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbsToProto(&r))
	}
	return p
}

// NoteVulnerabilityWindowsDetailsFixingKbsToProto converts a NoteVulnerabilityWindowsDetailsFixingKbs resource to its proto representation.
func ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbsToProto(o *containeranalysis.NoteVulnerabilityWindowsDetailsFixingKbs) *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.ValueOrEmptyString(o.Name),
		Url:  dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// NoteBuildToProto converts a NoteBuild resource to its proto representation.
func ContaineranalysisNoteBuildToProto(o *containeranalysis.NoteBuild) *containeranalysispb.ContaineranalysisNoteBuild {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteBuild{
		BuilderVersion: dcl.ValueOrEmptyString(o.BuilderVersion),
	}
	return p
}

// NoteImageToProto converts a NoteImage resource to its proto representation.
func ContaineranalysisNoteImageToProto(o *containeranalysis.NoteImage) *containeranalysispb.ContaineranalysisNoteImage {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteImage{
		ResourceUrl: dcl.ValueOrEmptyString(o.ResourceUrl),
		Fingerprint: ContaineranalysisNoteImageFingerprintToProto(o.Fingerprint),
	}
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint resource to its proto representation.
func ContaineranalysisNoteImageFingerprintToProto(o *containeranalysis.NoteImageFingerprint) *containeranalysispb.ContaineranalysisNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteImageFingerprint{
		V1Name: dcl.ValueOrEmptyString(o.V1Name),
		V2Name: dcl.ValueOrEmptyString(o.V2Name),
	}
	for _, r := range o.V2Blob {
		p.V2Blob = append(p.V2Blob, r)
	}
	return p
}

// NotePackageToProto converts a NotePackage resource to its proto representation.
func ContaineranalysisNotePackageToProto(o *containeranalysis.NotePackage) *containeranalysispb.ContaineranalysisNotePackage {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNotePackage{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	for _, r := range o.Distribution {
		p.Distribution = append(p.Distribution, ContaineranalysisNotePackageDistributionToProto(&r))
	}
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution resource to its proto representation.
func ContaineranalysisNotePackageDistributionToProto(o *containeranalysis.NotePackageDistribution) *containeranalysispb.ContaineranalysisNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNotePackageDistribution{
		CpeUri:        dcl.ValueOrEmptyString(o.CpeUri),
		Architecture:  ContaineranalysisNotePackageDistributionArchitectureEnumToProto(o.Architecture),
		LatestVersion: ContaineranalysisNotePackageDistributionLatestVersionToProto(o.LatestVersion),
		Maintainer:    dcl.ValueOrEmptyString(o.Maintainer),
		Url:           dcl.ValueOrEmptyString(o.Url),
		Description:   dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion resource to its proto representation.
func ContaineranalysisNotePackageDistributionLatestVersionToProto(o *containeranalysis.NotePackageDistributionLatestVersion) *containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersion{
		Epoch:    dcl.ValueOrEmptyInt64(o.Epoch),
		Name:     dcl.ValueOrEmptyString(o.Name),
		Revision: dcl.ValueOrEmptyString(o.Revision),
		Kind:     ContaineranalysisNotePackageDistributionLatestVersionKindEnumToProto(o.Kind),
		FullName: dcl.ValueOrEmptyString(o.FullName),
	}
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery resource to its proto representation.
func ContaineranalysisNoteDiscoveryToProto(o *containeranalysis.NoteDiscovery) *containeranalysispb.ContaineranalysisNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteDiscovery{
		AnalysisKind: ContaineranalysisNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind),
	}
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *containeranalysis.Note) *containeranalysispb.ContaineranalysisNote {
	p := &containeranalysispb.ContaineranalysisNote{
		Name:             dcl.ValueOrEmptyString(resource.Name),
		ShortDescription: dcl.ValueOrEmptyString(resource.ShortDescription),
		LongDescription:  dcl.ValueOrEmptyString(resource.LongDescription),
		ExpirationTime:   dcl.ValueOrEmptyString(resource.ExpirationTime),
		CreateTime:       dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:       dcl.ValueOrEmptyString(resource.UpdateTime),
		Vulnerability:    ContaineranalysisNoteVulnerabilityToProto(resource.Vulnerability),
		Build:            ContaineranalysisNoteBuildToProto(resource.Build),
		Image:            ContaineranalysisNoteImageToProto(resource.Image),
		Package:          ContaineranalysisNotePackageToProto(resource.Package),
		Discovery:        ContaineranalysisNoteDiscoveryToProto(resource.Discovery),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.RelatedUrl {
		p.RelatedUrl = append(p.RelatedUrl, ContaineranalysisNoteRelatedUrlToProto(&r))
	}
	for _, r := range resource.RelatedNoteNames {
		p.RelatedNoteNames = append(p.RelatedNoteNames, r)
	}

	return p
}

// ApplyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) applyNote(ctx context.Context, c *containeranalysis.Client, request *containeranalysispb.ApplyContaineranalysisNoteRequest) (*containeranalysispb.ContaineranalysisNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// ApplyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisNote(ctx context.Context, request *containeranalysispb.ApplyContaineranalysisNoteRequest) (*containeranalysispb.ContaineranalysisNote, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisNote(ctx context.Context, request *containeranalysispb.DeleteContaineranalysisNoteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))

}

// ListNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisNote(ctx context.Context, request *containeranalysispb.ListContaineranalysisNoteRequest) (*containeranalysispb.ListContaineranalysisNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*containeranalysispb.ContaineranalysisNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	return &containeranalysispb.ListContaineranalysisNoteResponse{Items: protos}, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*containeranalysis.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containeranalysis.NewClient(conf), nil
}
