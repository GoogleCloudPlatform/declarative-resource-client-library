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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/beta/osconfig_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
)

// Server implements the gRPC interface for GuestPolicy.
type GuestPolicyServer struct{}

// ProtoToGuestPolicyPackagesDesiredStateEnum converts a GuestPolicyPackagesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackagesDesiredStateEnum(e betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum) *beta.GuestPolicyPackagesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyPackagesDesiredStateEnum(n[len("GuestPolicyPackagesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackagesManagerEnum converts a GuestPolicyPackagesManagerEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackagesManagerEnum(e betapb.OsconfigBetaGuestPolicyPackagesManagerEnum) *beta.GuestPolicyPackagesManagerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyPackagesManagerEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyPackagesManagerEnum(n[len("GuestPolicyPackagesManagerEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackageRepositoriesAptArchiveTypeEnum converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(e betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum) *beta.GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyPackageRepositoriesAptArchiveTypeEnum(n[len("GuestPolicyPackageRepositoriesAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(e betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) *beta.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(n[len("GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(e betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) *beta.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(n[len("GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(e betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) *beta.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(n[len("GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(e betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) *beta.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(n[len("GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesDesiredStateEnum converts a GuestPolicyRecipesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesDesiredStateEnum(e betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum) *beta.GuestPolicyRecipesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesDesiredStateEnum(n[len("GuestPolicyRecipesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyAssignment converts a GuestPolicyAssignment resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyAssignment(p *betapb.OsconfigBetaGuestPolicyAssignment) *beta.GuestPolicyAssignment {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyAssignment{}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigBetaGuestPolicyAssignmentGroupLabels(r))
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, r)
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, r)
	}
	for _, r := range p.GetInstanceNamePrefixes() {
		obj.InstanceNamePrefixes = append(obj.InstanceNamePrefixes, r)
	}
	for _, r := range p.GetOsTypes() {
		obj.OsTypes = append(obj.OsTypes, *ProtoToOsconfigBetaGuestPolicyAssignmentOsTypes(r))
	}
	return obj
}

// ProtoToGuestPolicyAssignmentGroupLabels converts a GuestPolicyAssignmentGroupLabels resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyAssignmentGroupLabels(p *betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels) *beta.GuestPolicyAssignmentGroupLabels {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyAssignmentGroupLabels{}
	return obj
}

// ProtoToGuestPolicyAssignmentOsTypes converts a GuestPolicyAssignmentOsTypes resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyAssignmentOsTypes(p *betapb.OsconfigBetaGuestPolicyAssignmentOsTypes) *beta.GuestPolicyAssignmentOsTypes {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyAssignmentOsTypes{
		OsShortName:    dcl.StringOrNil(p.OsShortName),
		OsVersion:      dcl.StringOrNil(p.OsVersion),
		OsArchitecture: dcl.StringOrNil(p.OsArchitecture),
	}
	return obj
}

// ProtoToGuestPolicyPackages converts a GuestPolicyPackages resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackages(p *betapb.OsconfigBetaGuestPolicyPackages) *beta.GuestPolicyPackages {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackages{
		Name:         dcl.StringOrNil(p.Name),
		DesiredState: ProtoToOsconfigBetaGuestPolicyPackagesDesiredStateEnum(p.GetDesiredState()),
		Manager:      ProtoToOsconfigBetaGuestPolicyPackagesManagerEnum(p.GetManager()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositories converts a GuestPolicyPackageRepositories resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositories(p *betapb.OsconfigBetaGuestPolicyPackageRepositories) *beta.GuestPolicyPackageRepositories {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositories{
		Apt:    ProtoToOsconfigBetaGuestPolicyPackageRepositoriesApt(p.GetApt()),
		Yum:    ProtoToOsconfigBetaGuestPolicyPackageRepositoriesYum(p.GetYum()),
		Zypper: ProtoToOsconfigBetaGuestPolicyPackageRepositoriesZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigBetaGuestPolicyPackageRepositoriesGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesApt converts a GuestPolicyPackageRepositoriesApt resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesApt(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesApt) *beta.GuestPolicyPackageRepositoriesApt {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesApt{
		ArchiveType:  ProtoToOsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(p.GetArchiveType()),
		Uri:          dcl.StringOrNil(p.Uri),
		Distribution: dcl.StringOrNil(p.Distribution),
		GpgKey:       dcl.StringOrNil(p.GpgKey),
	}
	for _, r := range p.GetComponents() {
		obj.Components = append(obj.Components, r)
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesYum converts a GuestPolicyPackageRepositoriesYum resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesYum(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesYum) *beta.GuestPolicyPackageRepositoriesYum {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesYum{
		Id:          dcl.StringOrNil(p.Id),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		BaseUrl:     dcl.StringOrNil(p.BaseUrl),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesZypper converts a GuestPolicyPackageRepositoriesZypper resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesZypper(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesZypper) *beta.GuestPolicyPackageRepositoriesZypper {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesZypper{
		Id:          dcl.StringOrNil(p.Id),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		BaseUrl:     dcl.StringOrNil(p.BaseUrl),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesGoo converts a GuestPolicyPackageRepositoriesGoo resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesGoo(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesGoo) *beta.GuestPolicyPackageRepositoriesGoo {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesGoo{
		Name: dcl.StringOrNil(p.Name),
		Url:  dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToGuestPolicyRecipes converts a GuestPolicyRecipes resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipes(p *betapb.OsconfigBetaGuestPolicyRecipes) *beta.GuestPolicyRecipes {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipes{
		Name:         dcl.StringOrNil(p.Name),
		Version:      dcl.StringOrNil(p.Version),
		DesiredState: ProtoToOsconfigBetaGuestPolicyRecipesDesiredStateEnum(p.GetDesiredState()),
	}
	for _, r := range p.GetArtifacts() {
		obj.Artifacts = append(obj.Artifacts, *ProtoToOsconfigBetaGuestPolicyRecipesArtifacts(r))
	}
	for _, r := range p.GetInstallSteps() {
		obj.InstallSteps = append(obj.InstallSteps, *ProtoToOsconfigBetaGuestPolicyRecipesInstallSteps(r))
	}
	for _, r := range p.GetUpdateSteps() {
		obj.UpdateSteps = append(obj.UpdateSteps, *ProtoToOsconfigBetaGuestPolicyRecipesUpdateSteps(r))
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifacts converts a GuestPolicyRecipesArtifacts resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesArtifacts(p *betapb.OsconfigBetaGuestPolicyRecipesArtifacts) *beta.GuestPolicyRecipesArtifacts {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesArtifacts{
		Id:            dcl.StringOrNil(p.Id),
		Remote:        ProtoToOsconfigBetaGuestPolicyRecipesArtifactsRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaGuestPolicyRecipesArtifactsGcs(p.GetGcs()),
		AllowInsecure: dcl.Bool(p.AllowInsecure),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsRemote converts a GuestPolicyRecipesArtifactsRemote resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesArtifactsRemote(p *betapb.OsconfigBetaGuestPolicyRecipesArtifactsRemote) *beta.GuestPolicyRecipesArtifactsRemote {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesArtifactsRemote{
		Uri:      dcl.StringOrNil(p.Uri),
		Checksum: dcl.StringOrNil(p.Checksum),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsGcs converts a GuestPolicyRecipesArtifactsGcs resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesArtifactsGcs(p *betapb.OsconfigBetaGuestPolicyRecipesArtifactsGcs) *beta.GuestPolicyRecipesArtifactsGcs {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesArtifactsGcs{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.Int64OrNil(p.Generation),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallSteps converts a GuestPolicyRecipesInstallSteps resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallSteps(p *betapb.OsconfigBetaGuestPolicyRecipesInstallSteps) *beta.GuestPolicyRecipesInstallSteps {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallSteps{
		FileCopy:          ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileCopy converts a GuestPolicyRecipesInstallStepsFileCopy resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileCopy(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy) *beta.GuestPolicyRecipesInstallStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Overwrite:   dcl.Bool(p.Overwrite),
		Permissions: dcl.StringOrNil(p.Permissions),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtraction converts a GuestPolicyRecipesInstallStepsArchiveExtraction resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction) *beta.GuestPolicyRecipesInstallStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Type:        ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsMsiInstallation converts a GuestPolicyRecipesInstallStepsMsiInstallation resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation) *beta.GuestPolicyRecipesInstallStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsMsiInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	for _, r := range p.GetFlags() {
		obj.Flags = append(obj.Flags, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsDpkgInstallation converts a GuestPolicyRecipesInstallStepsDpkgInstallation resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation) *beta.GuestPolicyRecipesInstallStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsRpmInstallation converts a GuestPolicyRecipesInstallStepsRpmInstallation resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation) *beta.GuestPolicyRecipesInstallStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileExec converts a GuestPolicyRecipesInstallStepsFileExec resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileExec(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec) *beta.GuestPolicyRecipesInstallStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsFileExec{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
		LocalPath:  dcl.StringOrNil(p.LocalPath),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsScriptRun converts a GuestPolicyRecipesInstallStepsScriptRun resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRun(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun) *beta.GuestPolicyRecipesInstallStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsScriptRun{
		Script:      dcl.StringOrNil(p.Script),
		Interpreter: ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateSteps converts a GuestPolicyRecipesUpdateSteps resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateSteps(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps) *beta.GuestPolicyRecipesUpdateSteps {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateSteps{
		FileCopy:          ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileCopy converts a GuestPolicyRecipesUpdateStepsFileCopy resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy) *beta.GuestPolicyRecipesUpdateStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Overwrite:   dcl.Bool(p.Overwrite),
		Permissions: dcl.StringOrNil(p.Permissions),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtraction converts a GuestPolicyRecipesUpdateStepsArchiveExtraction resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction) *beta.GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Type:        ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsMsiInstallation converts a GuestPolicyRecipesUpdateStepsMsiInstallation resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation) *beta.GuestPolicyRecipesUpdateStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsMsiInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	for _, r := range p.GetFlags() {
		obj.Flags = append(obj.Flags, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsDpkgInstallation converts a GuestPolicyRecipesUpdateStepsDpkgInstallation resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation) *beta.GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsRpmInstallation converts a GuestPolicyRecipesUpdateStepsRpmInstallation resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation) *beta.GuestPolicyRecipesUpdateStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileExec converts a GuestPolicyRecipesUpdateStepsFileExec resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileExec(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec) *beta.GuestPolicyRecipesUpdateStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsFileExec{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
		LocalPath:  dcl.StringOrNil(p.LocalPath),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsScriptRun converts a GuestPolicyRecipesUpdateStepsScriptRun resource from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun) *beta.GuestPolicyRecipesUpdateStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsScriptRun{
		Script:      dcl.StringOrNil(p.Script),
		Interpreter: ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicy converts a GuestPolicy resource from its proto representation.
func ProtoToGuestPolicy(p *betapb.OsconfigBetaGuestPolicy) *beta.GuestPolicy {
	obj := &beta.GuestPolicy{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Assignment:  ProtoToOsconfigBetaGuestPolicyAssignment(p.GetAssignment()),
		Etag:        dcl.StringOrNil(p.Etag),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetPackages() {
		obj.Packages = append(obj.Packages, *ProtoToOsconfigBetaGuestPolicyPackages(r))
	}
	for _, r := range p.GetPackageRepositories() {
		obj.PackageRepositories = append(obj.PackageRepositories, *ProtoToOsconfigBetaGuestPolicyPackageRepositories(r))
	}
	for _, r := range p.GetRecipes() {
		obj.Recipes = append(obj.Recipes, *ProtoToOsconfigBetaGuestPolicyRecipes(r))
	}
	return obj
}

// GuestPolicyPackagesDesiredStateEnumToProto converts a GuestPolicyPackagesDesiredStateEnum enum to its proto representation.
func OsconfigBetaGuestPolicyPackagesDesiredStateEnumToProto(e *beta.GuestPolicyPackagesDesiredStateEnum) betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum_value["GuestPolicyPackagesDesiredStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum(0)
}

// GuestPolicyPackagesManagerEnumToProto converts a GuestPolicyPackagesManagerEnum enum to its proto representation.
func OsconfigBetaGuestPolicyPackagesManagerEnumToProto(e *beta.GuestPolicyPackagesManagerEnum) betapb.OsconfigBetaGuestPolicyPackagesManagerEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyPackagesManagerEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyPackagesManagerEnum_value["GuestPolicyPackagesManagerEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyPackagesManagerEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyPackagesManagerEnum(0)
}

// GuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(e *beta.GuestPolicyPackageRepositoriesAptArchiveTypeEnum) betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_value["GuestPolicyPackageRepositoriesAptArchiveTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(e *beta.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(e *beta.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(e *beta.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(e *beta.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesDesiredStateEnumToProto converts a GuestPolicyRecipesDesiredStateEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesDesiredStateEnumToProto(e *beta.GuestPolicyRecipesDesiredStateEnum) betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum_value["GuestPolicyRecipesDesiredStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum(0)
}

// GuestPolicyAssignmentToProto converts a GuestPolicyAssignment resource to its proto representation.
func OsconfigBetaGuestPolicyAssignmentToProto(o *beta.GuestPolicyAssignment) *betapb.OsconfigBetaGuestPolicyAssignment {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyAssignment{}
	for _, r := range o.GroupLabels {
		p.GroupLabels = append(p.GroupLabels, OsconfigBetaGuestPolicyAssignmentGroupLabelsToProto(&r))
	}
	for _, r := range o.Zones {
		p.Zones = append(p.Zones, r)
	}
	for _, r := range o.Instances {
		p.Instances = append(p.Instances, r)
	}
	for _, r := range o.InstanceNamePrefixes {
		p.InstanceNamePrefixes = append(p.InstanceNamePrefixes, r)
	}
	for _, r := range o.OsTypes {
		p.OsTypes = append(p.OsTypes, OsconfigBetaGuestPolicyAssignmentOsTypesToProto(&r))
	}
	return p
}

// GuestPolicyAssignmentGroupLabelsToProto converts a GuestPolicyAssignmentGroupLabels resource to its proto representation.
func OsconfigBetaGuestPolicyAssignmentGroupLabelsToProto(o *beta.GuestPolicyAssignmentGroupLabels) *betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels{}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// GuestPolicyAssignmentOsTypesToProto converts a GuestPolicyAssignmentOsTypes resource to its proto representation.
func OsconfigBetaGuestPolicyAssignmentOsTypesToProto(o *beta.GuestPolicyAssignmentOsTypes) *betapb.OsconfigBetaGuestPolicyAssignmentOsTypes {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyAssignmentOsTypes{
		OsShortName:    dcl.ValueOrEmptyString(o.OsShortName),
		OsVersion:      dcl.ValueOrEmptyString(o.OsVersion),
		OsArchitecture: dcl.ValueOrEmptyString(o.OsArchitecture),
	}
	return p
}

// GuestPolicyPackagesToProto converts a GuestPolicyPackages resource to its proto representation.
func OsconfigBetaGuestPolicyPackagesToProto(o *beta.GuestPolicyPackages) *betapb.OsconfigBetaGuestPolicyPackages {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackages{
		Name:         dcl.ValueOrEmptyString(o.Name),
		DesiredState: OsconfigBetaGuestPolicyPackagesDesiredStateEnumToProto(o.DesiredState),
		Manager:      OsconfigBetaGuestPolicyPackagesManagerEnumToProto(o.Manager),
	}
	return p
}

// GuestPolicyPackageRepositoriesToProto converts a GuestPolicyPackageRepositories resource to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesToProto(o *beta.GuestPolicyPackageRepositories) *betapb.OsconfigBetaGuestPolicyPackageRepositories {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositories{
		Apt:    OsconfigBetaGuestPolicyPackageRepositoriesAptToProto(o.Apt),
		Yum:    OsconfigBetaGuestPolicyPackageRepositoriesYumToProto(o.Yum),
		Zypper: OsconfigBetaGuestPolicyPackageRepositoriesZypperToProto(o.Zypper),
		Goo:    OsconfigBetaGuestPolicyPackageRepositoriesGooToProto(o.Goo),
	}
	return p
}

// GuestPolicyPackageRepositoriesAptToProto converts a GuestPolicyPackageRepositoriesApt resource to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesAptToProto(o *beta.GuestPolicyPackageRepositoriesApt) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesApt {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesApt{
		ArchiveType:  OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(o.ArchiveType),
		Uri:          dcl.ValueOrEmptyString(o.Uri),
		Distribution: dcl.ValueOrEmptyString(o.Distribution),
		GpgKey:       dcl.ValueOrEmptyString(o.GpgKey),
	}
	for _, r := range o.Components {
		p.Components = append(p.Components, r)
	}
	return p
}

// GuestPolicyPackageRepositoriesYumToProto converts a GuestPolicyPackageRepositoriesYum resource to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesYumToProto(o *beta.GuestPolicyPackageRepositoriesYum) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesYum {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesYum{
		Id:          dcl.ValueOrEmptyString(o.Id),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		BaseUrl:     dcl.ValueOrEmptyString(o.BaseUrl),
	}
	for _, r := range o.GpgKeys {
		p.GpgKeys = append(p.GpgKeys, r)
	}
	return p
}

// GuestPolicyPackageRepositoriesZypperToProto converts a GuestPolicyPackageRepositoriesZypper resource to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesZypperToProto(o *beta.GuestPolicyPackageRepositoriesZypper) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesZypper {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesZypper{
		Id:          dcl.ValueOrEmptyString(o.Id),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		BaseUrl:     dcl.ValueOrEmptyString(o.BaseUrl),
	}
	for _, r := range o.GpgKeys {
		p.GpgKeys = append(p.GpgKeys, r)
	}
	return p
}

// GuestPolicyPackageRepositoriesGooToProto converts a GuestPolicyPackageRepositoriesGoo resource to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesGooToProto(o *beta.GuestPolicyPackageRepositoriesGoo) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesGoo {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesGoo{
		Name: dcl.ValueOrEmptyString(o.Name),
		Url:  dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// GuestPolicyRecipesToProto converts a GuestPolicyRecipes resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesToProto(o *beta.GuestPolicyRecipes) *betapb.OsconfigBetaGuestPolicyRecipes {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipes{
		Name:         dcl.ValueOrEmptyString(o.Name),
		Version:      dcl.ValueOrEmptyString(o.Version),
		DesiredState: OsconfigBetaGuestPolicyRecipesDesiredStateEnumToProto(o.DesiredState),
	}
	for _, r := range o.Artifacts {
		p.Artifacts = append(p.Artifacts, OsconfigBetaGuestPolicyRecipesArtifactsToProto(&r))
	}
	for _, r := range o.InstallSteps {
		p.InstallSteps = append(p.InstallSteps, OsconfigBetaGuestPolicyRecipesInstallStepsToProto(&r))
	}
	for _, r := range o.UpdateSteps {
		p.UpdateSteps = append(p.UpdateSteps, OsconfigBetaGuestPolicyRecipesUpdateStepsToProto(&r))
	}
	return p
}

// GuestPolicyRecipesArtifactsToProto converts a GuestPolicyRecipesArtifacts resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesArtifactsToProto(o *beta.GuestPolicyRecipesArtifacts) *betapb.OsconfigBetaGuestPolicyRecipesArtifacts {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesArtifacts{
		Id:            dcl.ValueOrEmptyString(o.Id),
		Remote:        OsconfigBetaGuestPolicyRecipesArtifactsRemoteToProto(o.Remote),
		Gcs:           OsconfigBetaGuestPolicyRecipesArtifactsGcsToProto(o.Gcs),
		AllowInsecure: dcl.ValueOrEmptyBool(o.AllowInsecure),
	}
	return p
}

// GuestPolicyRecipesArtifactsRemoteToProto converts a GuestPolicyRecipesArtifactsRemote resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesArtifactsRemoteToProto(o *beta.GuestPolicyRecipesArtifactsRemote) *betapb.OsconfigBetaGuestPolicyRecipesArtifactsRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesArtifactsRemote{
		Uri:      dcl.ValueOrEmptyString(o.Uri),
		Checksum: dcl.ValueOrEmptyString(o.Checksum),
	}
	return p
}

// GuestPolicyRecipesArtifactsGcsToProto converts a GuestPolicyRecipesArtifactsGcs resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesArtifactsGcsToProto(o *beta.GuestPolicyRecipesArtifactsGcs) *betapb.OsconfigBetaGuestPolicyRecipesArtifactsGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesArtifactsGcs{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyInt64(o.Generation),
	}
	return p
}

// GuestPolicyRecipesInstallStepsToProto converts a GuestPolicyRecipesInstallSteps resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsToProto(o *beta.GuestPolicyRecipesInstallSteps) *betapb.OsconfigBetaGuestPolicyRecipesInstallSteps {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallSteps{
		FileCopy:          OsconfigBetaGuestPolicyRecipesInstallStepsFileCopyToProto(o.FileCopy),
		ArchiveExtraction: OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o.ArchiveExtraction),
		MsiInstallation:   OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o.MsiInstallation),
		DpkgInstallation:  OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o.DpkgInstallation),
		RpmInstallation:   OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o.RpmInstallation),
		FileExec:          OsconfigBetaGuestPolicyRecipesInstallStepsFileExecToProto(o.FileExec),
		ScriptRun:         OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunToProto(o.ScriptRun),
	}
	return p
}

// GuestPolicyRecipesInstallStepsFileCopyToProto converts a GuestPolicyRecipesInstallStepsFileCopy resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsFileCopyToProto(o *beta.GuestPolicyRecipesInstallStepsFileCopy) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Overwrite:   dcl.ValueOrEmptyBool(o.Overwrite),
		Permissions: dcl.ValueOrEmptyString(o.Permissions),
	}
	return p
}

// GuestPolicyRecipesInstallStepsArchiveExtractionToProto converts a GuestPolicyRecipesInstallStepsArchiveExtraction resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o *beta.GuestPolicyRecipesInstallStepsArchiveExtraction) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Type:        OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(o.Type),
	}
	return p
}

// GuestPolicyRecipesInstallStepsMsiInstallationToProto converts a GuestPolicyRecipesInstallStepsMsiInstallation resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o *beta.GuestPolicyRecipesInstallStepsMsiInstallation) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	for _, r := range o.Flags {
		p.Flags = append(p.Flags, r)
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyRecipesInstallStepsDpkgInstallationToProto converts a GuestPolicyRecipesInstallStepsDpkgInstallation resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o *beta.GuestPolicyRecipesInstallStepsDpkgInstallation) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesInstallStepsRpmInstallationToProto converts a GuestPolicyRecipesInstallStepsRpmInstallation resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o *beta.GuestPolicyRecipesInstallStepsRpmInstallation) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesInstallStepsFileExecToProto converts a GuestPolicyRecipesInstallStepsFileExec resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsFileExecToProto(o *beta.GuestPolicyRecipesInstallStepsFileExec) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
		LocalPath:  dcl.ValueOrEmptyString(o.LocalPath),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyRecipesInstallStepsScriptRunToProto converts a GuestPolicyRecipesInstallStepsScriptRun resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunToProto(o *beta.GuestPolicyRecipesInstallStepsScriptRun) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun{
		Script:      dcl.ValueOrEmptyString(o.Script),
		Interpreter: OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(o.Interpreter),
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyRecipesUpdateStepsToProto converts a GuestPolicyRecipesUpdateSteps resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsToProto(o *beta.GuestPolicyRecipesUpdateSteps) *betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps{
		FileCopy:          OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopyToProto(o.FileCopy),
		ArchiveExtraction: OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o.ArchiveExtraction),
		MsiInstallation:   OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o.MsiInstallation),
		DpkgInstallation:  OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o.DpkgInstallation),
		RpmInstallation:   OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o.RpmInstallation),
		FileExec:          OsconfigBetaGuestPolicyRecipesUpdateStepsFileExecToProto(o.FileExec),
		ScriptRun:         OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunToProto(o.ScriptRun),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsFileCopyToProto converts a GuestPolicyRecipesUpdateStepsFileCopy resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopyToProto(o *beta.GuestPolicyRecipesUpdateStepsFileCopy) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Overwrite:   dcl.ValueOrEmptyBool(o.Overwrite),
		Permissions: dcl.ValueOrEmptyString(o.Permissions),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtraction resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o *beta.GuestPolicyRecipesUpdateStepsArchiveExtraction) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Type:        OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(o.Type),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsMsiInstallationToProto converts a GuestPolicyRecipesUpdateStepsMsiInstallation resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o *beta.GuestPolicyRecipesUpdateStepsMsiInstallation) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	for _, r := range o.Flags {
		p.Flags = append(p.Flags, r)
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyRecipesUpdateStepsDpkgInstallationToProto converts a GuestPolicyRecipesUpdateStepsDpkgInstallation resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o *beta.GuestPolicyRecipesUpdateStepsDpkgInstallation) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsRpmInstallationToProto converts a GuestPolicyRecipesUpdateStepsRpmInstallation resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o *beta.GuestPolicyRecipesUpdateStepsRpmInstallation) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsFileExecToProto converts a GuestPolicyRecipesUpdateStepsFileExec resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsFileExecToProto(o *beta.GuestPolicyRecipesUpdateStepsFileExec) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
		LocalPath:  dcl.ValueOrEmptyString(o.LocalPath),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyRecipesUpdateStepsScriptRunToProto converts a GuestPolicyRecipesUpdateStepsScriptRun resource to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunToProto(o *beta.GuestPolicyRecipesUpdateStepsScriptRun) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun{
		Script:      dcl.ValueOrEmptyString(o.Script),
		Interpreter: OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(o.Interpreter),
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyToProto converts a GuestPolicy resource to its proto representation.
func GuestPolicyToProto(resource *beta.GuestPolicy) *betapb.OsconfigBetaGuestPolicy {
	p := &betapb.OsconfigBetaGuestPolicy{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Assignment:  OsconfigBetaGuestPolicyAssignmentToProto(resource.Assignment),
		Etag:        dcl.ValueOrEmptyString(resource.Etag),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Packages {
		p.Packages = append(p.Packages, OsconfigBetaGuestPolicyPackagesToProto(&r))
	}
	for _, r := range resource.PackageRepositories {
		p.PackageRepositories = append(p.PackageRepositories, OsconfigBetaGuestPolicyPackageRepositoriesToProto(&r))
	}
	for _, r := range resource.Recipes {
		p.Recipes = append(p.Recipes, OsconfigBetaGuestPolicyRecipesToProto(&r))
	}

	return p
}

// ApplyGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) applyGuestPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyOsconfigBetaGuestPolicyRequest) (*betapb.OsconfigBetaGuestPolicy, error) {
	p := ProtoToGuestPolicy(request.GetResource())
	res, err := c.ApplyGuestPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GuestPolicyToProto(res)
	return r, nil
}

// ApplyGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) ApplyOsconfigBetaGuestPolicy(ctx context.Context, request *betapb.ApplyOsconfigBetaGuestPolicyRequest) (*betapb.OsconfigBetaGuestPolicy, error) {
	cl, err := createConfigGuestPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyGuestPolicy(ctx, cl, request)
}

// DeleteGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Delete() method.
func (s *GuestPolicyServer) DeleteOsconfigBetaGuestPolicy(ctx context.Context, request *betapb.DeleteOsconfigBetaGuestPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGuestPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGuestPolicy(ctx, ProtoToGuestPolicy(request.GetResource()))

}

// ListGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicyList() method.
func (s *GuestPolicyServer) ListOsconfigBetaGuestPolicy(ctx context.Context, request *betapb.ListOsconfigBetaGuestPolicyRequest) (*betapb.ListOsconfigBetaGuestPolicyResponse, error) {
	cl, err := createConfigGuestPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGuestPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.OsconfigBetaGuestPolicy
	for _, r := range resources.Items {
		rp := GuestPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListOsconfigBetaGuestPolicyResponse{Items: protos}, nil
}

func createConfigGuestPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
