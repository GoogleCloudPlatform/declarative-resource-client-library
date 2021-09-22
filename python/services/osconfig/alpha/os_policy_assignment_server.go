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

// Server implements the gRPC interface for OSPolicyAssignment.
type OSPolicyAssignmentServer struct{}

// ProtoToOSPolicyAssignmentOSPoliciesModeEnum converts a OSPolicyAssignmentOSPoliciesModeEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum) *alpha.OSPolicyAssignmentOSPoliciesModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesModeEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentExecInterpreterEnum converts a OSPolicyAssignmentExecInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentExecInterpreterEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum) *alpha.OSPolicyAssignmentExecInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentExecInterpreterEnum(n[len("OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentRolloutStateEnum converts a OSPolicyAssignmentRolloutStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentRolloutStateEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum) *alpha.OSPolicyAssignmentRolloutStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentRolloutStateEnum(n[len("OsconfigAlphaOSPolicyAssignmentRolloutStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPolicies converts a OSPolicyAssignmentOSPolicies resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPolicies(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies) *alpha.OSPolicyAssignmentOSPolicies {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPolicies{
		Id:                        dcl.StringOrNil(p.Id),
		Description:               dcl.StringOrNil(p.Description),
		Mode:                      ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(p.GetMode()),
		AllowNoResourceGroupMatch: dcl.Bool(p.AllowNoResourceGroupMatch),
	}
	for _, r := range p.GetResourceGroups() {
		obj.ResourceGroups = append(obj.ResourceGroups, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroups converts a OSPolicyAssignmentOSPoliciesResourceGroups resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups) *alpha.OSPolicyAssignmentOSPoliciesResourceGroups {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroups{
		OSFilter: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter(p.GetOsFilter()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter converts a OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter{
		OSShortName: dcl.StringOrNil(p.OsShortName),
		OSVersion:   dcl.StringOrNil(p.OsVersion),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResources converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResources{
		Id:         dcl.StringOrNil(p.Id),
		Pkg:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p.GetPkg()),
		Repository: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p.GetRepository()),
		Exec:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p.GetExec()),
		File:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p.GetFile()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{
		DesiredState: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(p.GetDesiredState()),
		Apt:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p.GetApt()),
		Deb:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p.GetDeb()),
		Yum:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p.GetYum()),
		Zypper:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p.GetZypper()),
		Rpm:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p.GetRpm()),
		Googet:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p.GetGooget()),
		Msi:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p.GetMsi()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{
		Source:   ProtoToOsconfigAlphaOSPolicyAssignmentFile(p.GetSource()),
		PullDeps: dcl.Bool(p.PullDeps),
	}
	return obj
}

// ProtoToOSPolicyAssignmentFile converts a OSPolicyAssignmentFile resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentFile(p *alphapb.OsconfigAlphaOSPolicyAssignmentFile) *alpha.OSPolicyAssignmentFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentFile{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.LocalPath),
		AllowInsecure: dcl.Bool(p.AllowInsecure),
	}
	return obj
}

// ProtoToOSPolicyAssignmentFileRemote converts a OSPolicyAssignmentFileRemote resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentFileRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentFileRemote) *alpha.OSPolicyAssignmentFileRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentFileRemote{
		Uri:            dcl.StringOrNil(p.Uri),
		Sha256Checksum: dcl.StringOrNil(p.Sha256Checksum),
	}
	return obj
}

// ProtoToOSPolicyAssignmentFileGcs converts a OSPolicyAssignmentFileGcs resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentFileGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentFileGcs) *alpha.OSPolicyAssignmentFileGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentFileGcs{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.Int64OrNil(p.Generation),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{
		Source:   ProtoToOsconfigAlphaOSPolicyAssignmentFile(p.GetSource()),
		PullDeps: dcl.Bool(p.PullDeps),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{
		Source: ProtoToOsconfigAlphaOSPolicyAssignmentFile(p.GetSource()),
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{
		Apt:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p.GetApt()),
		Yum:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p.GetYum()),
		Zypper: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(p.GetArchiveType()),
		Uri:          dcl.StringOrNil(p.Uri),
		Distribution: dcl.StringOrNil(p.Distribution),
		GpgKey:       dcl.StringOrNil(p.GpgKey),
	}
	for _, r := range p.GetComponents() {
		obj.Components = append(obj.Components, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{
		Id:          dcl.StringOrNil(p.Id),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		BaseUrl:     dcl.StringOrNil(p.BaseUrl),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{
		Id:          dcl.StringOrNil(p.Id),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		BaseUrl:     dcl.StringOrNil(p.BaseUrl),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.StringOrNil(p.Name),
		Url:  dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{
		Validate: ProtoToOsconfigAlphaOSPolicyAssignmentExec(p.GetValidate()),
		Enforce:  ProtoToOsconfigAlphaOSPolicyAssignmentExec(p.GetEnforce()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentExec converts a OSPolicyAssignmentExec resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentExec(p *alphapb.OsconfigAlphaOSPolicyAssignmentExec) *alpha.OSPolicyAssignmentExec {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentExec{
		File:        ProtoToOsconfigAlphaOSPolicyAssignmentFile(p.GetFile()),
		Script:      dcl.StringOrNil(p.Script),
		Interpreter: ProtoToOsconfigAlphaOSPolicyAssignmentExecInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{
		File:        ProtoToOsconfigAlphaOSPolicyAssignmentFile(p.GetFile()),
		Content:     dcl.StringOrNil(p.Content),
		Path:        dcl.StringOrNil(p.Path),
		State:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(p.GetState()),
		Permissions: dcl.StringOrNil(p.Permissions),
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilter converts a OSPolicyAssignmentInstanceFilter resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilter(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilter) *alpha.OSPolicyAssignmentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilter{
		All: dcl.Bool(p.All),
	}
	for _, r := range p.GetOsShortNames() {
		obj.OSShortNames = append(obj.OSShortNames, r)
	}
	for _, r := range p.GetInclusionLabels() {
		obj.InclusionLabels = append(obj.InclusionLabels, *ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels(r))
	}
	for _, r := range p.GetExclusionLabels() {
		obj.ExclusionLabels = append(obj.ExclusionLabels, *ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInclusionLabels converts a OSPolicyAssignmentInstanceFilterInclusionLabels resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels) *alpha.OSPolicyAssignmentInstanceFilterInclusionLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilterInclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterExclusionLabels converts a OSPolicyAssignmentInstanceFilterExclusionLabels resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels) *alpha.OSPolicyAssignmentInstanceFilterExclusionLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilterExclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentRollout converts a OSPolicyAssignmentRollout resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentRollout(p *alphapb.OsconfigAlphaOSPolicyAssignmentRollout) *alpha.OSPolicyAssignmentRollout {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentRollout{
		DisruptionBudget: ProtoToOsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget(p.GetDisruptionBudget()),
		MinWaitDuration:  dcl.StringOrNil(p.MinWaitDuration),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRolloutDisruptionBudget converts a OSPolicyAssignmentRolloutDisruptionBudget resource from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget(p *alphapb.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget) *alpha.OSPolicyAssignmentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.Fixed),
		Percent: dcl.Int64OrNil(p.Percent),
	}
	return obj
}

// ProtoToOSPolicyAssignment converts a OSPolicyAssignment resource from its proto representation.
func ProtoToOSPolicyAssignment(p *alphapb.OsconfigAlphaOSPolicyAssignment) *alpha.OSPolicyAssignment {
	obj := &alpha.OSPolicyAssignment{
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		InstanceFilter:     ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilter(p.GetInstanceFilter()),
		Rollout:            ProtoToOsconfigAlphaOSPolicyAssignmentRollout(p.GetRollout()),
		RevisionId:         dcl.StringOrNil(p.RevisionId),
		RevisionCreateTime: dcl.StringOrNil(p.GetRevisionCreateTime()),
		RolloutState:       ProtoToOsconfigAlphaOSPolicyAssignmentRolloutStateEnum(p.GetRolloutState()),
		Baseline:           dcl.Bool(p.Baseline),
		Deleted:            dcl.Bool(p.Deleted),
		Reconciling:        dcl.Bool(p.Reconciling),
		Uid:                dcl.StringOrNil(p.Uid),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetOsPolicies() {
		obj.OSPolicies = append(obj.OSPolicies, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPolicies(r))
	}
	return obj
}

// OSPolicyAssignmentOSPoliciesModeEnumToProto converts a OSPolicyAssignmentOSPoliciesModeEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesModeEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum_value["OSPolicyAssignmentOSPoliciesModeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
}

// OSPolicyAssignmentExecInterpreterEnumToProto converts a OSPolicyAssignmentExecInterpreterEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentExecInterpreterEnumToProto(e *alpha.OSPolicyAssignmentExecInterpreterEnum) alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum_value["OSPolicyAssignmentExecInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
}

// OSPolicyAssignmentRolloutStateEnumToProto converts a OSPolicyAssignmentRolloutStateEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentRolloutStateEnumToProto(e *alpha.OSPolicyAssignmentRolloutStateEnum) alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum_value["OSPolicyAssignmentRolloutStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesToProto converts a OSPolicyAssignmentOSPolicies resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesToProto(o *alpha.OSPolicyAssignmentOSPolicies) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies{
		Id:                        dcl.ValueOrEmptyString(o.Id),
		Description:               dcl.ValueOrEmptyString(o.Description),
		Mode:                      OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnumToProto(o.Mode),
		AllowNoResourceGroupMatch: dcl.ValueOrEmptyBool(o.AllowNoResourceGroupMatch),
	}
	for _, r := range o.ResourceGroups {
		p.ResourceGroups = append(p.ResourceGroups, OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsToProto(&r))
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroups resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroups) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups{
		OsFilter: OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilterToProto(o.OSFilter),
	}
	for _, r := range o.Resources {
		p.Resources = append(p.Resources, OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(&r))
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsOSFilterToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilterToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter{
		OsShortName: dcl.ValueOrEmptyString(o.OSShortName),
		OsVersion:   dcl.ValueOrEmptyString(o.OSVersion),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResources) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources{
		Id:         dcl.ValueOrEmptyString(o.Id),
		Pkg:        OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o.Pkg),
		Repository: OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o.Repository),
		Exec:       OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o.Exec),
		File:       OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o.File),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{
		DesiredState: OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(o.DesiredState),
		Apt:          OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o.Apt),
		Deb:          OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o.Deb),
		Yum:          OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o.Yum),
		Zypper:       OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o.Zypper),
		Rpm:          OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o.Rpm),
		Googet:       OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o.Googet),
		Msi:          OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o.Msi),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{
		Source:   OsconfigAlphaOSPolicyAssignmentFileToProto(o.Source),
		PullDeps: dcl.ValueOrEmptyBool(o.PullDeps),
	}
	return p
}

// OSPolicyAssignmentFileToProto converts a OSPolicyAssignmentFile resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentFileToProto(o *alpha.OSPolicyAssignmentFile) *alphapb.OsconfigAlphaOSPolicyAssignmentFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentFile{
		Remote:        OsconfigAlphaOSPolicyAssignmentFileRemoteToProto(o.Remote),
		Gcs:           OsconfigAlphaOSPolicyAssignmentFileGcsToProto(o.Gcs),
		LocalPath:     dcl.ValueOrEmptyString(o.LocalPath),
		AllowInsecure: dcl.ValueOrEmptyBool(o.AllowInsecure),
	}
	return p
}

// OSPolicyAssignmentFileRemoteToProto converts a OSPolicyAssignmentFileRemote resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentFileRemoteToProto(o *alpha.OSPolicyAssignmentFileRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentFileRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentFileRemote{
		Uri:            dcl.ValueOrEmptyString(o.Uri),
		Sha256Checksum: dcl.ValueOrEmptyString(o.Sha256Checksum),
	}
	return p
}

// OSPolicyAssignmentFileGcsToProto converts a OSPolicyAssignmentFileGcs resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentFileGcsToProto(o *alpha.OSPolicyAssignmentFileGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentFileGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentFileGcs{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyInt64(o.Generation),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{
		Source:   OsconfigAlphaOSPolicyAssignmentFileToProto(o.Source),
		PullDeps: dcl.ValueOrEmptyBool(o.PullDeps),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{
		Source: OsconfigAlphaOSPolicyAssignmentFileToProto(o.Source),
	}
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, r)
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{
		Apt:    OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o.Apt),
		Yum:    OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o.Yum),
		Zypper: OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o.Zypper),
		Goo:    OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o.Goo),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(o.ArchiveType),
		Uri:          dcl.ValueOrEmptyString(o.Uri),
		Distribution: dcl.ValueOrEmptyString(o.Distribution),
		GpgKey:       dcl.ValueOrEmptyString(o.GpgKey),
	}
	for _, r := range o.Components {
		p.Components = append(p.Components, r)
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{
		Id:          dcl.ValueOrEmptyString(o.Id),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		BaseUrl:     dcl.ValueOrEmptyString(o.BaseUrl),
	}
	for _, r := range o.GpgKeys {
		p.GpgKeys = append(p.GpgKeys, r)
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{
		Id:          dcl.ValueOrEmptyString(o.Id),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		BaseUrl:     dcl.ValueOrEmptyString(o.BaseUrl),
	}
	for _, r := range o.GpgKeys {
		p.GpgKeys = append(p.GpgKeys, r)
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.ValueOrEmptyString(o.Name),
		Url:  dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{
		Validate: OsconfigAlphaOSPolicyAssignmentExecToProto(o.Validate),
		Enforce:  OsconfigAlphaOSPolicyAssignmentExecToProto(o.Enforce),
	}
	return p
}

// OSPolicyAssignmentExecToProto converts a OSPolicyAssignmentExec resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentExecToProto(o *alpha.OSPolicyAssignmentExec) *alphapb.OsconfigAlphaOSPolicyAssignmentExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentExec{
		File:        OsconfigAlphaOSPolicyAssignmentFileToProto(o.File),
		Script:      dcl.ValueOrEmptyString(o.Script),
		Interpreter: OsconfigAlphaOSPolicyAssignmentExecInterpreterEnumToProto(o.Interpreter),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{
		File:        OsconfigAlphaOSPolicyAssignmentFileToProto(o.File),
		Content:     dcl.ValueOrEmptyString(o.Content),
		Path:        dcl.ValueOrEmptyString(o.Path),
		State:       OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(o.State),
		Permissions: dcl.ValueOrEmptyString(o.Permissions),
	}
	return p
}

// OSPolicyAssignmentInstanceFilterToProto converts a OSPolicyAssignmentInstanceFilter resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterToProto(o *alpha.OSPolicyAssignmentInstanceFilter) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilter{
		All: dcl.ValueOrEmptyBool(o.All),
	}
	for _, r := range o.OSShortNames {
		p.OsShortNames = append(p.OsShortNames, r)
	}
	for _, r := range o.InclusionLabels {
		p.InclusionLabels = append(p.InclusionLabels, OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(&r))
	}
	for _, r := range o.ExclusionLabels {
		p.ExclusionLabels = append(p.ExclusionLabels, OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(&r))
	}
	return p
}

// OSPolicyAssignmentInstanceFilterInclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterInclusionLabels resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(o *alpha.OSPolicyAssignmentInstanceFilterInclusionLabels) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels{}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// OSPolicyAssignmentInstanceFilterExclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterExclusionLabels resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(o *alpha.OSPolicyAssignmentInstanceFilterExclusionLabels) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels{}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// OSPolicyAssignmentRolloutToProto converts a OSPolicyAssignmentRollout resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentRolloutToProto(o *alpha.OSPolicyAssignmentRollout) *alphapb.OsconfigAlphaOSPolicyAssignmentRollout {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentRollout{
		DisruptionBudget: OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudgetToProto(o.DisruptionBudget),
		MinWaitDuration:  dcl.ValueOrEmptyString(o.MinWaitDuration),
	}
	return p
}

// OSPolicyAssignmentRolloutDisruptionBudgetToProto converts a OSPolicyAssignmentRolloutDisruptionBudget resource to its proto representation.
func OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudgetToProto(o *alpha.OSPolicyAssignmentRolloutDisruptionBudget) *alphapb.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.ValueOrEmptyInt64(o.Fixed),
		Percent: dcl.ValueOrEmptyInt64(o.Percent),
	}
	return p
}

// OSPolicyAssignmentToProto converts a OSPolicyAssignment resource to its proto representation.
func OSPolicyAssignmentToProto(resource *alpha.OSPolicyAssignment) *alphapb.OsconfigAlphaOSPolicyAssignment {
	p := &alphapb.OsconfigAlphaOSPolicyAssignment{
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		InstanceFilter:     OsconfigAlphaOSPolicyAssignmentInstanceFilterToProto(resource.InstanceFilter),
		Rollout:            OsconfigAlphaOSPolicyAssignmentRolloutToProto(resource.Rollout),
		RevisionId:         dcl.ValueOrEmptyString(resource.RevisionId),
		RevisionCreateTime: dcl.ValueOrEmptyString(resource.RevisionCreateTime),
		RolloutState:       OsconfigAlphaOSPolicyAssignmentRolloutStateEnumToProto(resource.RolloutState),
		Baseline:           dcl.ValueOrEmptyBool(resource.Baseline),
		Deleted:            dcl.ValueOrEmptyBool(resource.Deleted),
		Reconciling:        dcl.ValueOrEmptyBool(resource.Reconciling),
		Uid:                dcl.ValueOrEmptyString(resource.Uid),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.OSPolicies {
		p.OsPolicies = append(p.OsPolicies, OsconfigAlphaOSPolicyAssignmentOSPoliciesToProto(&r))
	}

	return p
}

// ApplyOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) applyOSPolicyAssignment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOsconfigAlphaOSPolicyAssignmentRequest) (*alphapb.OsconfigAlphaOSPolicyAssignment, error) {
	p := ProtoToOSPolicyAssignment(request.GetResource())
	res, err := c.ApplyOSPolicyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OSPolicyAssignmentToProto(res)
	return r, nil
}

// ApplyOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) ApplyOsconfigAlphaOSPolicyAssignment(ctx context.Context, request *alphapb.ApplyOsconfigAlphaOSPolicyAssignmentRequest) (*alphapb.OsconfigAlphaOSPolicyAssignment, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyOSPolicyAssignment(ctx, cl, request)
}

// DeleteOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Delete() method.
func (s *OSPolicyAssignmentServer) DeleteOsconfigAlphaOSPolicyAssignment(ctx context.Context, request *alphapb.DeleteOsconfigAlphaOSPolicyAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOSPolicyAssignment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOSPolicyAssignment(ctx, ProtoToOSPolicyAssignment(request.GetResource()))

}

// ListOsconfigAlphaOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignmentList() method.
func (s *OSPolicyAssignmentServer) ListOsconfigAlphaOSPolicyAssignment(ctx context.Context, request *alphapb.ListOsconfigAlphaOSPolicyAssignmentRequest) (*alphapb.ListOsconfigAlphaOSPolicyAssignmentResponse, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOSPolicyAssignment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OsconfigAlphaOSPolicyAssignment
	for _, r := range resources.Items {
		rp := OSPolicyAssignmentToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListOsconfigAlphaOSPolicyAssignmentResponse{Items: protos}, nil
}

func createConfigOSPolicyAssignment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
