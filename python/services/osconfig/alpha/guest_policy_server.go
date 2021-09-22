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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/alpha/osconfig_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
)

// Server implements the gRPC interface for GuestPolicy.
type GuestPolicyServer struct{}

// ProtoToGuestPolicyPackagesDesiredStateEnum converts a GuestPolicyPackagesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackagesDesiredStateEnum(e alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum) *alpha.GuestPolicyPackagesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyPackagesDesiredStateEnum(n[len("OsconfigAlphaGuestPolicyPackagesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackagesManagerEnum converts a GuestPolicyPackagesManagerEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackagesManagerEnum(e alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum) *alpha.GuestPolicyPackagesManagerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyPackagesManagerEnum(n[len("OsconfigAlphaGuestPolicyPackagesManagerEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackageRepositoriesAptArchiveTypeEnum converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(e alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum) *alpha.GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyPackageRepositoriesAptArchiveTypeEnum(n[len("OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) *alpha.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(n[len("OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) *alpha.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(n[len("OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) *alpha.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(n[len("OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) *alpha.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(n[len("OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesDesiredStateEnum converts a GuestPolicyRecipesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesDesiredStateEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum) *alpha.GuestPolicyRecipesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesDesiredStateEnum(n[len("OsconfigAlphaGuestPolicyRecipesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyAssignment converts a GuestPolicyAssignment resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyAssignment(p *alphapb.OsconfigAlphaGuestPolicyAssignment) *alpha.GuestPolicyAssignment {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyAssignment{}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigAlphaGuestPolicyAssignmentGroupLabels(r))
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
		obj.OSTypes = append(obj.OSTypes, *ProtoToOsconfigAlphaGuestPolicyAssignmentOSTypes(r))
	}
	return obj
}

// ProtoToGuestPolicyAssignmentGroupLabels converts a GuestPolicyAssignmentGroupLabels resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyAssignmentGroupLabels(p *alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels) *alpha.GuestPolicyAssignmentGroupLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyAssignmentGroupLabels{}
	return obj
}

// ProtoToGuestPolicyAssignmentOSTypes converts a GuestPolicyAssignmentOSTypes resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyAssignmentOSTypes(p *alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes) *alpha.GuestPolicyAssignmentOSTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyAssignmentOSTypes{
		OSShortName:    dcl.StringOrNil(p.OsShortName),
		OSVersion:      dcl.StringOrNil(p.OsVersion),
		OSArchitecture: dcl.StringOrNil(p.OsArchitecture),
	}
	return obj
}

// ProtoToGuestPolicyPackages converts a GuestPolicyPackages resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackages(p *alphapb.OsconfigAlphaGuestPolicyPackages) *alpha.GuestPolicyPackages {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackages{
		Name:         dcl.StringOrNil(p.Name),
		DesiredState: ProtoToOsconfigAlphaGuestPolicyPackagesDesiredStateEnum(p.GetDesiredState()),
		Manager:      ProtoToOsconfigAlphaGuestPolicyPackagesManagerEnum(p.GetManager()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositories converts a GuestPolicyPackageRepositories resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositories(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositories) *alpha.GuestPolicyPackageRepositories {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositories{
		Apt:    ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesApt(p.GetApt()),
		Yum:    ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesYum(p.GetYum()),
		Zypper: ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesApt converts a GuestPolicyPackageRepositoriesApt resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesApt(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesApt) *alpha.GuestPolicyPackageRepositoriesApt {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesApt{
		ArchiveType:  ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(p.GetArchiveType()),
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
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesYum(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesYum) *alpha.GuestPolicyPackageRepositoriesYum {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesYum{
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
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesZypper(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesZypper) *alpha.GuestPolicyPackageRepositoriesZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesZypper{
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
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesGoo(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesGoo) *alpha.GuestPolicyPackageRepositoriesGoo {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesGoo{
		Name: dcl.StringOrNil(p.Name),
		Url:  dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToGuestPolicyRecipes converts a GuestPolicyRecipes resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipes(p *alphapb.OsconfigAlphaGuestPolicyRecipes) *alpha.GuestPolicyRecipes {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipes{
		Name:         dcl.StringOrNil(p.Name),
		Version:      dcl.StringOrNil(p.Version),
		DesiredState: ProtoToOsconfigAlphaGuestPolicyRecipesDesiredStateEnum(p.GetDesiredState()),
	}
	for _, r := range p.GetArtifacts() {
		obj.Artifacts = append(obj.Artifacts, *ProtoToOsconfigAlphaGuestPolicyRecipesArtifacts(r))
	}
	for _, r := range p.GetInstallSteps() {
		obj.InstallSteps = append(obj.InstallSteps, *ProtoToOsconfigAlphaGuestPolicyRecipesInstallSteps(r))
	}
	for _, r := range p.GetUpdateSteps() {
		obj.UpdateSteps = append(obj.UpdateSteps, *ProtoToOsconfigAlphaGuestPolicyRecipesUpdateSteps(r))
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifacts converts a GuestPolicyRecipesArtifacts resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesArtifacts(p *alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts) *alpha.GuestPolicyRecipesArtifacts {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesArtifacts{
		Id:            dcl.StringOrNil(p.Id),
		Remote:        ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsGcs(p.GetGcs()),
		AllowInsecure: dcl.Bool(p.AllowInsecure),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsRemote converts a GuestPolicyRecipesArtifactsRemote resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsRemote(p *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsRemote) *alpha.GuestPolicyRecipesArtifactsRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesArtifactsRemote{
		Uri:      dcl.StringOrNil(p.Uri),
		Checksum: dcl.StringOrNil(p.Checksum),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsGcs converts a GuestPolicyRecipesArtifactsGcs resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsGcs(p *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsGcs) *alpha.GuestPolicyRecipesArtifactsGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesArtifactsGcs{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.Int64OrNil(p.Generation),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallSteps converts a GuestPolicyRecipesInstallSteps resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallSteps(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps) *alpha.GuestPolicyRecipesInstallSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallSteps{
		FileCopy:          ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileCopy converts a GuestPolicyRecipesInstallStepsFileCopy resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy) *alpha.GuestPolicyRecipesInstallStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Overwrite:   dcl.Bool(p.Overwrite),
		Permissions: dcl.StringOrNil(p.Permissions),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtraction converts a GuestPolicyRecipesInstallStepsArchiveExtraction resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction) *alpha.GuestPolicyRecipesInstallStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Type:        ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsMsiInstallation converts a GuestPolicyRecipesInstallStepsMsiInstallation resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation) *alpha.GuestPolicyRecipesInstallStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsMsiInstallation{
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
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation) *alpha.GuestPolicyRecipesInstallStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsRpmInstallation converts a GuestPolicyRecipesInstallStepsRpmInstallation resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation) *alpha.GuestPolicyRecipesInstallStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileExec converts a GuestPolicyRecipesInstallStepsFileExec resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileExec(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileExec) *alpha.GuestPolicyRecipesInstallStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsFileExec{
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
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun) *alpha.GuestPolicyRecipesInstallStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsScriptRun{
		Script:      dcl.StringOrNil(p.Script),
		Interpreter: ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateSteps converts a GuestPolicyRecipesUpdateSteps resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateSteps(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps) *alpha.GuestPolicyRecipesUpdateSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateSteps{
		FileCopy:          ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileCopy converts a GuestPolicyRecipesUpdateStepsFileCopy resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy) *alpha.GuestPolicyRecipesUpdateStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Overwrite:   dcl.Bool(p.Overwrite),
		Permissions: dcl.StringOrNil(p.Permissions),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtraction converts a GuestPolicyRecipesUpdateStepsArchiveExtraction resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction) *alpha.GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.ArtifactId),
		Destination: dcl.StringOrNil(p.Destination),
		Type:        ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsMsiInstallation converts a GuestPolicyRecipesUpdateStepsMsiInstallation resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation) *alpha.GuestPolicyRecipesUpdateStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsMsiInstallation{
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
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation) *alpha.GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsRpmInstallation converts a GuestPolicyRecipesUpdateStepsRpmInstallation resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation) *alpha.GuestPolicyRecipesUpdateStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.ArtifactId),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileExec converts a GuestPolicyRecipesUpdateStepsFileExec resource from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec) *alpha.GuestPolicyRecipesUpdateStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsFileExec{
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
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun) *alpha.GuestPolicyRecipesUpdateStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsScriptRun{
		Script:      dcl.StringOrNil(p.Script),
		Interpreter: ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicy converts a GuestPolicy resource from its proto representation.
func ProtoToGuestPolicy(p *alphapb.OsconfigAlphaGuestPolicy) *alpha.GuestPolicy {
	obj := &alpha.GuestPolicy{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Assignment:  ProtoToOsconfigAlphaGuestPolicyAssignment(p.GetAssignment()),
		Etag:        dcl.StringOrNil(p.Etag),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetPackages() {
		obj.Packages = append(obj.Packages, *ProtoToOsconfigAlphaGuestPolicyPackages(r))
	}
	for _, r := range p.GetPackageRepositories() {
		obj.PackageRepositories = append(obj.PackageRepositories, *ProtoToOsconfigAlphaGuestPolicyPackageRepositories(r))
	}
	for _, r := range p.GetRecipes() {
		obj.Recipes = append(obj.Recipes, *ProtoToOsconfigAlphaGuestPolicyRecipes(r))
	}
	return obj
}

// GuestPolicyPackagesDesiredStateEnumToProto converts a GuestPolicyPackagesDesiredStateEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyPackagesDesiredStateEnumToProto(e *alpha.GuestPolicyPackagesDesiredStateEnum) alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum_value["GuestPolicyPackagesDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum(0)
}

// GuestPolicyPackagesManagerEnumToProto converts a GuestPolicyPackagesManagerEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyPackagesManagerEnumToProto(e *alpha.GuestPolicyPackagesManagerEnum) alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum_value["GuestPolicyPackagesManagerEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum(0)
}

// GuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(e *alpha.GuestPolicyPackageRepositoriesAptArchiveTypeEnum) alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_value["GuestPolicyPackageRepositoriesAptArchiveTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(e *alpha.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(e *alpha.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(e *alpha.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(e *alpha.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesDesiredStateEnumToProto converts a GuestPolicyRecipesDesiredStateEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesDesiredStateEnumToProto(e *alpha.GuestPolicyRecipesDesiredStateEnum) alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum_value["GuestPolicyRecipesDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum(0)
}

// GuestPolicyAssignmentToProto converts a GuestPolicyAssignment resource to its proto representation.
func OsconfigAlphaGuestPolicyAssignmentToProto(o *alpha.GuestPolicyAssignment) *alphapb.OsconfigAlphaGuestPolicyAssignment {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyAssignment{}
	for _, r := range o.GroupLabels {
		p.GroupLabels = append(p.GroupLabels, OsconfigAlphaGuestPolicyAssignmentGroupLabelsToProto(&r))
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
	for _, r := range o.OSTypes {
		p.OsTypes = append(p.OsTypes, OsconfigAlphaGuestPolicyAssignmentOSTypesToProto(&r))
	}
	return p
}

// GuestPolicyAssignmentGroupLabelsToProto converts a GuestPolicyAssignmentGroupLabels resource to its proto representation.
func OsconfigAlphaGuestPolicyAssignmentGroupLabelsToProto(o *alpha.GuestPolicyAssignmentGroupLabels) *alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels{}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// GuestPolicyAssignmentOSTypesToProto converts a GuestPolicyAssignmentOSTypes resource to its proto representation.
func OsconfigAlphaGuestPolicyAssignmentOSTypesToProto(o *alpha.GuestPolicyAssignmentOSTypes) *alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes{
		OsShortName:    dcl.ValueOrEmptyString(o.OSShortName),
		OsVersion:      dcl.ValueOrEmptyString(o.OSVersion),
		OsArchitecture: dcl.ValueOrEmptyString(o.OSArchitecture),
	}
	return p
}

// GuestPolicyPackagesToProto converts a GuestPolicyPackages resource to its proto representation.
func OsconfigAlphaGuestPolicyPackagesToProto(o *alpha.GuestPolicyPackages) *alphapb.OsconfigAlphaGuestPolicyPackages {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackages{
		Name:         dcl.ValueOrEmptyString(o.Name),
		DesiredState: OsconfigAlphaGuestPolicyPackagesDesiredStateEnumToProto(o.DesiredState),
		Manager:      OsconfigAlphaGuestPolicyPackagesManagerEnumToProto(o.Manager),
	}
	return p
}

// GuestPolicyPackageRepositoriesToProto converts a GuestPolicyPackageRepositories resource to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesToProto(o *alpha.GuestPolicyPackageRepositories) *alphapb.OsconfigAlphaGuestPolicyPackageRepositories {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositories{
		Apt:    OsconfigAlphaGuestPolicyPackageRepositoriesAptToProto(o.Apt),
		Yum:    OsconfigAlphaGuestPolicyPackageRepositoriesYumToProto(o.Yum),
		Zypper: OsconfigAlphaGuestPolicyPackageRepositoriesZypperToProto(o.Zypper),
		Goo:    OsconfigAlphaGuestPolicyPackageRepositoriesGooToProto(o.Goo),
	}
	return p
}

// GuestPolicyPackageRepositoriesAptToProto converts a GuestPolicyPackageRepositoriesApt resource to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesAptToProto(o *alpha.GuestPolicyPackageRepositoriesApt) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesApt{
		ArchiveType:  OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(o.ArchiveType),
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
func OsconfigAlphaGuestPolicyPackageRepositoriesYumToProto(o *alpha.GuestPolicyPackageRepositoriesYum) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesYum{
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
func OsconfigAlphaGuestPolicyPackageRepositoriesZypperToProto(o *alpha.GuestPolicyPackageRepositoriesZypper) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesZypper{
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
func OsconfigAlphaGuestPolicyPackageRepositoriesGooToProto(o *alpha.GuestPolicyPackageRepositoriesGoo) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesGoo {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesGoo{
		Name: dcl.ValueOrEmptyString(o.Name),
		Url:  dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// GuestPolicyRecipesToProto converts a GuestPolicyRecipes resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesToProto(o *alpha.GuestPolicyRecipes) *alphapb.OsconfigAlphaGuestPolicyRecipes {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipes{
		Name:         dcl.ValueOrEmptyString(o.Name),
		Version:      dcl.ValueOrEmptyString(o.Version),
		DesiredState: OsconfigAlphaGuestPolicyRecipesDesiredStateEnumToProto(o.DesiredState),
	}
	for _, r := range o.Artifacts {
		p.Artifacts = append(p.Artifacts, OsconfigAlphaGuestPolicyRecipesArtifactsToProto(&r))
	}
	for _, r := range o.InstallSteps {
		p.InstallSteps = append(p.InstallSteps, OsconfigAlphaGuestPolicyRecipesInstallStepsToProto(&r))
	}
	for _, r := range o.UpdateSteps {
		p.UpdateSteps = append(p.UpdateSteps, OsconfigAlphaGuestPolicyRecipesUpdateStepsToProto(&r))
	}
	return p
}

// GuestPolicyRecipesArtifactsToProto converts a GuestPolicyRecipesArtifacts resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesArtifactsToProto(o *alpha.GuestPolicyRecipesArtifacts) *alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts{
		Id:            dcl.ValueOrEmptyString(o.Id),
		Remote:        OsconfigAlphaGuestPolicyRecipesArtifactsRemoteToProto(o.Remote),
		Gcs:           OsconfigAlphaGuestPolicyRecipesArtifactsGcsToProto(o.Gcs),
		AllowInsecure: dcl.ValueOrEmptyBool(o.AllowInsecure),
	}
	return p
}

// GuestPolicyRecipesArtifactsRemoteToProto converts a GuestPolicyRecipesArtifactsRemote resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesArtifactsRemoteToProto(o *alpha.GuestPolicyRecipesArtifactsRemote) *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsRemote{
		Uri:      dcl.ValueOrEmptyString(o.Uri),
		Checksum: dcl.ValueOrEmptyString(o.Checksum),
	}
	return p
}

// GuestPolicyRecipesArtifactsGcsToProto converts a GuestPolicyRecipesArtifactsGcs resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesArtifactsGcsToProto(o *alpha.GuestPolicyRecipesArtifactsGcs) *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsGcs{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyInt64(o.Generation),
	}
	return p
}

// GuestPolicyRecipesInstallStepsToProto converts a GuestPolicyRecipesInstallSteps resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsToProto(o *alpha.GuestPolicyRecipesInstallSteps) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps{
		FileCopy:          OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopyToProto(o.FileCopy),
		ArchiveExtraction: OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o.ArchiveExtraction),
		MsiInstallation:   OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o.MsiInstallation),
		DpkgInstallation:  OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o.DpkgInstallation),
		RpmInstallation:   OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o.RpmInstallation),
		FileExec:          OsconfigAlphaGuestPolicyRecipesInstallStepsFileExecToProto(o.FileExec),
		ScriptRun:         OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunToProto(o.ScriptRun),
	}
	return p
}

// GuestPolicyRecipesInstallStepsFileCopyToProto converts a GuestPolicyRecipesInstallStepsFileCopy resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopyToProto(o *alpha.GuestPolicyRecipesInstallStepsFileCopy) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Overwrite:   dcl.ValueOrEmptyBool(o.Overwrite),
		Permissions: dcl.ValueOrEmptyString(o.Permissions),
	}
	return p
}

// GuestPolicyRecipesInstallStepsArchiveExtractionToProto converts a GuestPolicyRecipesInstallStepsArchiveExtraction resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o *alpha.GuestPolicyRecipesInstallStepsArchiveExtraction) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Type:        OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(o.Type),
	}
	return p
}

// GuestPolicyRecipesInstallStepsMsiInstallationToProto converts a GuestPolicyRecipesInstallStepsMsiInstallation resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o *alpha.GuestPolicyRecipesInstallStepsMsiInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation{
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
func OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o *alpha.GuestPolicyRecipesInstallStepsDpkgInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesInstallStepsRpmInstallationToProto converts a GuestPolicyRecipesInstallStepsRpmInstallation resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o *alpha.GuestPolicyRecipesInstallStepsRpmInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesInstallStepsFileExecToProto converts a GuestPolicyRecipesInstallStepsFileExec resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsFileExecToProto(o *alpha.GuestPolicyRecipesInstallStepsFileExec) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileExec{
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
func OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunToProto(o *alpha.GuestPolicyRecipesInstallStepsScriptRun) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun{
		Script:      dcl.ValueOrEmptyString(o.Script),
		Interpreter: OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(o.Interpreter),
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyRecipesUpdateStepsToProto converts a GuestPolicyRecipesUpdateSteps resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsToProto(o *alpha.GuestPolicyRecipesUpdateSteps) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps{
		FileCopy:          OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopyToProto(o.FileCopy),
		ArchiveExtraction: OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o.ArchiveExtraction),
		MsiInstallation:   OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o.MsiInstallation),
		DpkgInstallation:  OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o.DpkgInstallation),
		RpmInstallation:   OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o.RpmInstallation),
		FileExec:          OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExecToProto(o.FileExec),
		ScriptRun:         OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunToProto(o.ScriptRun),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsFileCopyToProto converts a GuestPolicyRecipesUpdateStepsFileCopy resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopyToProto(o *alpha.GuestPolicyRecipesUpdateStepsFileCopy) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Overwrite:   dcl.ValueOrEmptyBool(o.Overwrite),
		Permissions: dcl.ValueOrEmptyString(o.Permissions),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtraction resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o *alpha.GuestPolicyRecipesUpdateStepsArchiveExtraction) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction{
		ArtifactId:  dcl.ValueOrEmptyString(o.ArtifactId),
		Destination: dcl.ValueOrEmptyString(o.Destination),
		Type:        OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(o.Type),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsMsiInstallationToProto converts a GuestPolicyRecipesUpdateStepsMsiInstallation resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o *alpha.GuestPolicyRecipesUpdateStepsMsiInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation{
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
func OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o *alpha.GuestPolicyRecipesUpdateStepsDpkgInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsRpmInstallationToProto converts a GuestPolicyRecipesUpdateStepsRpmInstallation resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o *alpha.GuestPolicyRecipesUpdateStepsRpmInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation{
		ArtifactId: dcl.ValueOrEmptyString(o.ArtifactId),
	}
	return p
}

// GuestPolicyRecipesUpdateStepsFileExecToProto converts a GuestPolicyRecipesUpdateStepsFileExec resource to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExecToProto(o *alpha.GuestPolicyRecipesUpdateStepsFileExec) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec{
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
func OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunToProto(o *alpha.GuestPolicyRecipesUpdateStepsScriptRun) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun{
		Script:      dcl.ValueOrEmptyString(o.Script),
		Interpreter: OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(o.Interpreter),
	}
	for _, r := range o.AllowedExitCodes {
		p.AllowedExitCodes = append(p.AllowedExitCodes, r)
	}
	return p
}

// GuestPolicyToProto converts a GuestPolicy resource to its proto representation.
func GuestPolicyToProto(resource *alpha.GuestPolicy) *alphapb.OsconfigAlphaGuestPolicy {
	p := &alphapb.OsconfigAlphaGuestPolicy{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Assignment:  OsconfigAlphaGuestPolicyAssignmentToProto(resource.Assignment),
		Etag:        dcl.ValueOrEmptyString(resource.Etag),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Packages {
		p.Packages = append(p.Packages, OsconfigAlphaGuestPolicyPackagesToProto(&r))
	}
	for _, r := range resource.PackageRepositories {
		p.PackageRepositories = append(p.PackageRepositories, OsconfigAlphaGuestPolicyPackageRepositoriesToProto(&r))
	}
	for _, r := range resource.Recipes {
		p.Recipes = append(p.Recipes, OsconfigAlphaGuestPolicyRecipesToProto(&r))
	}

	return p
}

// ApplyGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) applyGuestPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOsconfigAlphaGuestPolicyRequest) (*alphapb.OsconfigAlphaGuestPolicy, error) {
	p := ProtoToGuestPolicy(request.GetResource())
	res, err := c.ApplyGuestPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GuestPolicyToProto(res)
	return r, nil
}

// ApplyGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) ApplyOsconfigAlphaGuestPolicy(ctx context.Context, request *alphapb.ApplyOsconfigAlphaGuestPolicyRequest) (*alphapb.OsconfigAlphaGuestPolicy, error) {
	cl, err := createConfigGuestPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyGuestPolicy(ctx, cl, request)
}

// DeleteGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Delete() method.
func (s *GuestPolicyServer) DeleteOsconfigAlphaGuestPolicy(ctx context.Context, request *alphapb.DeleteOsconfigAlphaGuestPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGuestPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGuestPolicy(ctx, ProtoToGuestPolicy(request.GetResource()))

}

// ListOsconfigAlphaGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicyList() method.
func (s *GuestPolicyServer) ListOsconfigAlphaGuestPolicy(ctx context.Context, request *alphapb.ListOsconfigAlphaGuestPolicyRequest) (*alphapb.ListOsconfigAlphaGuestPolicyResponse, error) {
	cl, err := createConfigGuestPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGuestPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OsconfigAlphaGuestPolicy
	for _, r := range resources.Items {
		rp := GuestPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListOsconfigAlphaGuestPolicyResponse{Items: protos}, nil
}

func createConfigGuestPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
