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

// Server implements the gRPC interface for OsPolicyAssignment.
type OsPolicyAssignmentServer struct{}

// ProtoToOsPolicyAssignmentOsPoliciesModeEnum converts a OsPolicyAssignmentOsPoliciesModeEnum enum from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum(e alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum) *alpha.OsPolicyAssignmentOsPoliciesModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum_name[int32(e)]; ok {
		e := alpha.OsPolicyAssignmentOsPoliciesModeEnum(n[len("OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(e alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(n[len("OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(e alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_name[int32(e)]; ok {
		e := alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(n[len("OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOsPolicyAssignmentExecInterpreterEnum converts a OsPolicyAssignmentExecInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentExecInterpreterEnum(e alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum) *alpha.OsPolicyAssignmentExecInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum_name[int32(e)]; ok {
		e := alpha.OsPolicyAssignmentExecInterpreterEnum(n[len("OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(e alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum_name[int32(e)]; ok {
		e := alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(n[len("OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOsPolicyAssignmentRolloutStateEnum converts a OsPolicyAssignmentRolloutStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentRolloutStateEnum(e alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum) *alpha.OsPolicyAssignmentRolloutStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum_name[int32(e)]; ok {
		e := alpha.OsPolicyAssignmentRolloutStateEnum(n[len("OsconfigAlphaOsPolicyAssignmentRolloutStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOsPolicyAssignmentOsPolicies converts a OsPolicyAssignmentOsPolicies resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPolicies(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPolicies) *alpha.OsPolicyAssignmentOsPolicies {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPolicies{
		Id:                        dcl.StringOrNil(p.Id),
		Description:               dcl.StringOrNil(p.Description),
		Mode:                      ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum(p.GetMode()),
		AllowNoResourceGroupMatch: dcl.Bool(p.AllowNoResourceGroupMatch),
	}
	for _, r := range p.GetResourceGroups() {
		obj.ResourceGroups = append(obj.ResourceGroups, *ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroups(r))
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroups converts a OsPolicyAssignmentOsPoliciesResourceGroups resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroups(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroups) *alpha.OsPolicyAssignmentOsPoliciesResourceGroups {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroups{
		OsFilter: ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(p.GetOsFilter()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResources(r))
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter converts a OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{
		OsShortName: dcl.StringOrNil(p.OsShortName),
		OsVersion:   dcl.StringOrNil(p.OsVersion),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResources converts a OsPolicyAssignmentOsPoliciesResourceGroupsResources resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResources(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResources) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResources {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResources{
		Id:         dcl.StringOrNil(p.Id),
		Pkg:        ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(p.GetPkg()),
		Repository: ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(p.GetRepository()),
		Exec:       ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(p.GetExec()),
		File:       ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(p.GetFile()),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{
		DesiredState: ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(p.GetDesiredState()),
		Apt:          ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(p.GetApt()),
		Deb:          ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(p.GetDeb()),
		Yum:          ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(p.GetYum()),
		Zypper:       ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(p.GetZypper()),
		Rpm:          ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(p.GetRpm()),
		Googet:       ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(p.GetGooget()),
		Msi:          ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(p.GetMsi()),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{
		Source:   ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(p.GetSource()),
		PullDeps: dcl.Bool(p.PullDeps),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{
		Remote:        ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.LocalPath),
		AllowInsecure: dcl.Bool(p.AllowInsecure),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{
		Uri:            dcl.StringOrNil(p.Uri),
		Sha256Checksum: dcl.StringOrNil(p.Sha256Checksum),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.Int64OrNil(p.Generation),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{
		Source:   ProtoToOsconfigAlphaOsPolicyAssignmentFile(p.GetSource()),
		PullDeps: dcl.Bool(p.PullDeps),
	}
	return obj
}

// ProtoToOsPolicyAssignmentFile converts a OsPolicyAssignmentFile resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentFile(p *alphapb.OsconfigAlphaOsPolicyAssignmentFile) *alpha.OsPolicyAssignmentFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentFile{
		Remote:        ProtoToOsconfigAlphaOsPolicyAssignmentFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOsPolicyAssignmentFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.LocalPath),
		AllowInsecure: dcl.Bool(p.AllowInsecure),
	}
	return obj
}

// ProtoToOsPolicyAssignmentFileRemote converts a OsPolicyAssignmentFileRemote resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentFileRemote(p *alphapb.OsconfigAlphaOsPolicyAssignmentFileRemote) *alpha.OsPolicyAssignmentFileRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentFileRemote{
		Uri:            dcl.StringOrNil(p.Uri),
		Sha256Checksum: dcl.StringOrNil(p.Sha256Checksum),
	}
	return obj
}

// ProtoToOsPolicyAssignmentFileGcs converts a OsPolicyAssignmentFileGcs resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentFileGcs(p *alphapb.OsconfigAlphaOsPolicyAssignmentFileGcs) *alpha.OsPolicyAssignmentFileGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentFileGcs{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.Int64OrNil(p.Generation),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{
		Source: ProtoToOsconfigAlphaOsPolicyAssignmentFile(p.GetSource()),
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, r)
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{
		Apt:    ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(p.GetApt()),
		Yum:    ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(p.GetYum()),
		Zypper: ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(p.GetArchiveType()),
		Uri:          dcl.StringOrNil(p.Uri),
		Distribution: dcl.StringOrNil(p.Distribution),
		GpgKey:       dcl.StringOrNil(p.GpgKey),
	}
	for _, r := range p.GetComponents() {
		obj.Components = append(obj.Components, r)
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{
		Id:          dcl.StringOrNil(p.Id),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		BaseUrl:     dcl.StringOrNil(p.BaseUrl),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{
		Id:          dcl.StringOrNil(p.Id),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		BaseUrl:     dcl.StringOrNil(p.BaseUrl),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.StringOrNil(p.Name),
		Url:  dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{
		Validate: ProtoToOsconfigAlphaOsPolicyAssignmentExec(p.GetValidate()),
		Enforce:  ProtoToOsconfigAlphaOsPolicyAssignmentExec(p.GetEnforce()),
	}
	return obj
}

// ProtoToOsPolicyAssignmentExec converts a OsPolicyAssignmentExec resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentExec(p *alphapb.OsconfigAlphaOsPolicyAssignmentExec) *alpha.OsPolicyAssignmentExec {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentExec{
		File:        ProtoToOsconfigAlphaOsPolicyAssignmentFile(p.GetFile()),
		Script:      dcl.StringOrNil(p.Script),
		Interpreter: ProtoToOsconfigAlphaOsPolicyAssignmentExecInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(p *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{
		File:        ProtoToOsconfigAlphaOsPolicyAssignmentFile(p.GetFile()),
		Content:     dcl.StringOrNil(p.Content),
		Path:        dcl.StringOrNil(p.Path),
		State:       ProtoToOsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(p.GetState()),
		Permissions: dcl.StringOrNil(p.Permissions),
	}
	return obj
}

// ProtoToOsPolicyAssignmentInstanceFilter converts a OsPolicyAssignmentInstanceFilter resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentInstanceFilter(p *alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilter) *alpha.OsPolicyAssignmentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentInstanceFilter{
		All: dcl.Bool(p.All),
	}
	for _, r := range p.GetOsShortNames() {
		obj.OsShortNames = append(obj.OsShortNames, r)
	}
	for _, r := range p.GetInclusionLabels() {
		obj.InclusionLabels = append(obj.InclusionLabels, *ProtoToOsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabels(r))
	}
	for _, r := range p.GetExclusionLabels() {
		obj.ExclusionLabels = append(obj.ExclusionLabels, *ProtoToOsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabels(r))
	}
	return obj
}

// ProtoToOsPolicyAssignmentInstanceFilterInclusionLabels converts a OsPolicyAssignmentInstanceFilterInclusionLabels resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabels(p *alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabels) *alpha.OsPolicyAssignmentInstanceFilterInclusionLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentInstanceFilterInclusionLabels{}
	return obj
}

// ProtoToOsPolicyAssignmentInstanceFilterExclusionLabels converts a OsPolicyAssignmentInstanceFilterExclusionLabels resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabels(p *alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabels) *alpha.OsPolicyAssignmentInstanceFilterExclusionLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentInstanceFilterExclusionLabels{}
	return obj
}

// ProtoToOsPolicyAssignmentRollout converts a OsPolicyAssignmentRollout resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentRollout(p *alphapb.OsconfigAlphaOsPolicyAssignmentRollout) *alpha.OsPolicyAssignmentRollout {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentRollout{
		DisruptionBudget: ProtoToOsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudget(p.GetDisruptionBudget()),
		MinWaitDuration:  dcl.StringOrNil(p.MinWaitDuration),
	}
	return obj
}

// ProtoToOsPolicyAssignmentRolloutDisruptionBudget converts a OsPolicyAssignmentRolloutDisruptionBudget resource from its proto representation.
func ProtoToOsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudget(p *alphapb.OsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudget) *alpha.OsPolicyAssignmentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &alpha.OsPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.Fixed),
		Percent: dcl.Int64OrNil(p.Percent),
	}
	return obj
}

// ProtoToOsPolicyAssignment converts a OsPolicyAssignment resource from its proto representation.
func ProtoToOsPolicyAssignment(p *alphapb.OsconfigAlphaOsPolicyAssignment) *alpha.OsPolicyAssignment {
	obj := &alpha.OsPolicyAssignment{
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		InstanceFilter:     ProtoToOsconfigAlphaOsPolicyAssignmentInstanceFilter(p.GetInstanceFilter()),
		Rollout:            ProtoToOsconfigAlphaOsPolicyAssignmentRollout(p.GetRollout()),
		RevisionId:         dcl.StringOrNil(p.RevisionId),
		RevisionCreateTime: dcl.StringOrNil(p.GetRevisionCreateTime()),
		RolloutState:       ProtoToOsconfigAlphaOsPolicyAssignmentRolloutStateEnum(p.GetRolloutState()),
		Baseline:           dcl.Bool(p.Baseline),
		Deleted:            dcl.Bool(p.Deleted),
		Reconciling:        dcl.Bool(p.Reconciling),
		Uid:                dcl.StringOrNil(p.Uid),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetOsPolicies() {
		obj.OsPolicies = append(obj.OsPolicies, *ProtoToOsconfigAlphaOsPolicyAssignmentOsPolicies(r))
	}
	return obj
}

// OsPolicyAssignmentOsPoliciesModeEnumToProto converts a OsPolicyAssignmentOsPoliciesModeEnum enum to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnumToProto(e *alpha.OsPolicyAssignmentOsPoliciesModeEnum) alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum_value["OsPolicyAssignmentOsPoliciesModeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum(v)
	}
	return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum(0)
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(e *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum) alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum_value["OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(e *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_value["OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(v)
	}
	return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
}

// OsPolicyAssignmentExecInterpreterEnumToProto converts a OsPolicyAssignmentExecInterpreterEnum enum to its proto representation.
func OsconfigAlphaOsPolicyAssignmentExecInterpreterEnumToProto(e *alpha.OsPolicyAssignmentExecInterpreterEnum) alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum_value["OsPolicyAssignmentExecInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum(0)
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum enum to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumToProto(e *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum) alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum_value["OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(v)
	}
	return alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(0)
}

// OsPolicyAssignmentRolloutStateEnumToProto converts a OsPolicyAssignmentRolloutStateEnum enum to its proto representation.
func OsconfigAlphaOsPolicyAssignmentRolloutStateEnumToProto(e *alpha.OsPolicyAssignmentRolloutStateEnum) alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum_value["OsPolicyAssignmentRolloutStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum(v)
	}
	return alphapb.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum(0)
}

// OsPolicyAssignmentOsPoliciesToProto converts a OsPolicyAssignmentOsPolicies resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesToProto(o *alpha.OsPolicyAssignmentOsPolicies) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPolicies {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPolicies{
		Id:                        dcl.ValueOrEmptyString(o.Id),
		Description:               dcl.ValueOrEmptyString(o.Description),
		Mode:                      OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnumToProto(o.Mode),
		AllowNoResourceGroupMatch: dcl.ValueOrEmptyBool(o.AllowNoResourceGroupMatch),
	}
	for _, r := range o.ResourceGroups {
		p.ResourceGroups = append(p.ResourceGroups, OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsToProto(&r))
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsToProto converts a OsPolicyAssignmentOsPoliciesResourceGroups resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroups) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroups {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroups{
		OsFilter: OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterToProto(o.OsFilter),
	}
	for _, r := range o.Resources {
		p.Resources = append(p.Resources, OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesToProto(&r))
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsOsFilterToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilterToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter{
		OsShortName: dcl.ValueOrEmptyString(o.OsShortName),
		OsVersion:   dcl.ValueOrEmptyString(o.OsVersion),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResources resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResources) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResources {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResources{
		Id:         dcl.ValueOrEmptyString(o.Id),
		Pkg:        OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgToProto(o.Pkg),
		Repository: OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryToProto(o.Repository),
		Exec:       OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecToProto(o.Exec),
		File:       OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileToProto(o.File),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg{
		DesiredState: OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(o.DesiredState),
		Apt:          OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptToProto(o.Apt),
		Deb:          OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebToProto(o.Deb),
		Yum:          OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumToProto(o.Yum),
		Zypper:       OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperToProto(o.Zypper),
		Rpm:          OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmToProto(o.Rpm),
		Googet:       OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetToProto(o.Googet),
		Msi:          OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiToProto(o.Msi),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb{
		Source:   OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceToProto(o.Source),
		PullDeps: dcl.ValueOrEmptyBool(o.PullDeps),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource{
		Remote:        OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o.Remote),
		Gcs:           OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o.Gcs),
		LocalPath:     dcl.ValueOrEmptyString(o.LocalPath),
		AllowInsecure: dcl.ValueOrEmptyBool(o.AllowInsecure),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote{
		Uri:            dcl.ValueOrEmptyString(o.Uri),
		Sha256Checksum: dcl.ValueOrEmptyString(o.Sha256Checksum),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyInt64(o.Generation),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm{
		Source:   OsconfigAlphaOsPolicyAssignmentFileToProto(o.Source),
		PullDeps: dcl.ValueOrEmptyBool(o.PullDeps),
	}
	return p
}

// OsPolicyAssignmentFileToProto converts a OsPolicyAssignmentFile resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentFileToProto(o *alpha.OsPolicyAssignmentFile) *alphapb.OsconfigAlphaOsPolicyAssignmentFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentFile{
		Remote:        OsconfigAlphaOsPolicyAssignmentFileRemoteToProto(o.Remote),
		Gcs:           OsconfigAlphaOsPolicyAssignmentFileGcsToProto(o.Gcs),
		LocalPath:     dcl.ValueOrEmptyString(o.LocalPath),
		AllowInsecure: dcl.ValueOrEmptyBool(o.AllowInsecure),
	}
	return p
}

// OsPolicyAssignmentFileRemoteToProto converts a OsPolicyAssignmentFileRemote resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentFileRemoteToProto(o *alpha.OsPolicyAssignmentFileRemote) *alphapb.OsconfigAlphaOsPolicyAssignmentFileRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentFileRemote{
		Uri:            dcl.ValueOrEmptyString(o.Uri),
		Sha256Checksum: dcl.ValueOrEmptyString(o.Sha256Checksum),
	}
	return p
}

// OsPolicyAssignmentFileGcsToProto converts a OsPolicyAssignmentFileGcs resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentFileGcsToProto(o *alpha.OsPolicyAssignmentFileGcs) *alphapb.OsconfigAlphaOsPolicyAssignmentFileGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentFileGcs{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyInt64(o.Generation),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi{
		Source: OsconfigAlphaOsPolicyAssignmentFileToProto(o.Source),
	}
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, r)
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository{
		Apt:    OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptToProto(o.Apt),
		Yum:    OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumToProto(o.Yum),
		Zypper: OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperToProto(o.Zypper),
		Goo:    OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooToProto(o.Goo),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(o.ArchiveType),
		Uri:          dcl.ValueOrEmptyString(o.Uri),
		Distribution: dcl.ValueOrEmptyString(o.Distribution),
		GpgKey:       dcl.ValueOrEmptyString(o.GpgKey),
	}
	for _, r := range o.Components {
		p.Components = append(p.Components, r)
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum{
		Id:          dcl.ValueOrEmptyString(o.Id),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		BaseUrl:     dcl.ValueOrEmptyString(o.BaseUrl),
	}
	for _, r := range o.GpgKeys {
		p.GpgKeys = append(p.GpgKeys, r)
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper{
		Id:          dcl.ValueOrEmptyString(o.Id),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		BaseUrl:     dcl.ValueOrEmptyString(o.BaseUrl),
	}
	for _, r := range o.GpgKeys {
		p.GpgKeys = append(p.GpgKeys, r)
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.ValueOrEmptyString(o.Name),
		Url:  dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec{
		Validate: OsconfigAlphaOsPolicyAssignmentExecToProto(o.Validate),
		Enforce:  OsconfigAlphaOsPolicyAssignmentExecToProto(o.Enforce),
	}
	return p
}

// OsPolicyAssignmentExecToProto converts a OsPolicyAssignmentExec resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentExecToProto(o *alpha.OsPolicyAssignmentExec) *alphapb.OsconfigAlphaOsPolicyAssignmentExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentExec{
		File:        OsconfigAlphaOsPolicyAssignmentFileToProto(o.File),
		Script:      dcl.ValueOrEmptyString(o.Script),
		Interpreter: OsconfigAlphaOsPolicyAssignmentExecInterpreterEnumToProto(o.Interpreter),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	return p
}

// OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileToProto converts a OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileToProto(o *alpha.OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) *alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile{
		File:        OsconfigAlphaOsPolicyAssignmentFileToProto(o.File),
		Content:     dcl.ValueOrEmptyString(o.Content),
		Path:        dcl.ValueOrEmptyString(o.Path),
		State:       OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnumToProto(o.State),
		Permissions: dcl.ValueOrEmptyString(o.Permissions),
	}
	return p
}

// OsPolicyAssignmentInstanceFilterToProto converts a OsPolicyAssignmentInstanceFilter resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentInstanceFilterToProto(o *alpha.OsPolicyAssignmentInstanceFilter) *alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilter{
		All: dcl.ValueOrEmptyBool(o.All),
	}
	for _, r := range o.OsShortNames {
		p.OsShortNames = append(p.OsShortNames, r)
	}
	for _, r := range o.InclusionLabels {
		p.InclusionLabels = append(p.InclusionLabels, OsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabelsToProto(&r))
	}
	for _, r := range o.ExclusionLabels {
		p.ExclusionLabels = append(p.ExclusionLabels, OsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabelsToProto(&r))
	}
	return p
}

// OsPolicyAssignmentInstanceFilterInclusionLabelsToProto converts a OsPolicyAssignmentInstanceFilterInclusionLabels resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabelsToProto(o *alpha.OsPolicyAssignmentInstanceFilterInclusionLabels) *alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabels{}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// OsPolicyAssignmentInstanceFilterExclusionLabelsToProto converts a OsPolicyAssignmentInstanceFilterExclusionLabels resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabelsToProto(o *alpha.OsPolicyAssignmentInstanceFilterExclusionLabels) *alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabels{}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// OsPolicyAssignmentRolloutToProto converts a OsPolicyAssignmentRollout resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentRolloutToProto(o *alpha.OsPolicyAssignmentRollout) *alphapb.OsconfigAlphaOsPolicyAssignmentRollout {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentRollout{
		DisruptionBudget: OsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudgetToProto(o.DisruptionBudget),
		MinWaitDuration:  dcl.ValueOrEmptyString(o.MinWaitDuration),
	}
	return p
}

// OsPolicyAssignmentRolloutDisruptionBudgetToProto converts a OsPolicyAssignmentRolloutDisruptionBudget resource to its proto representation.
func OsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudgetToProto(o *alpha.OsPolicyAssignmentRolloutDisruptionBudget) *alphapb.OsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.ValueOrEmptyInt64(o.Fixed),
		Percent: dcl.ValueOrEmptyInt64(o.Percent),
	}
	return p
}

// OsPolicyAssignmentToProto converts a OsPolicyAssignment resource to its proto representation.
func OsPolicyAssignmentToProto(resource *alpha.OsPolicyAssignment) *alphapb.OsconfigAlphaOsPolicyAssignment {
	p := &alphapb.OsconfigAlphaOsPolicyAssignment{
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		InstanceFilter:     OsconfigAlphaOsPolicyAssignmentInstanceFilterToProto(resource.InstanceFilter),
		Rollout:            OsconfigAlphaOsPolicyAssignmentRolloutToProto(resource.Rollout),
		RevisionId:         dcl.ValueOrEmptyString(resource.RevisionId),
		RevisionCreateTime: dcl.ValueOrEmptyString(resource.RevisionCreateTime),
		RolloutState:       OsconfigAlphaOsPolicyAssignmentRolloutStateEnumToProto(resource.RolloutState),
		Baseline:           dcl.ValueOrEmptyBool(resource.Baseline),
		Deleted:            dcl.ValueOrEmptyBool(resource.Deleted),
		Reconciling:        dcl.ValueOrEmptyBool(resource.Reconciling),
		Uid:                dcl.ValueOrEmptyString(resource.Uid),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.OsPolicies {
		p.OsPolicies = append(p.OsPolicies, OsconfigAlphaOsPolicyAssignmentOsPoliciesToProto(&r))
	}

	return p
}

// ApplyOsPolicyAssignment handles the gRPC request by passing it to the underlying OsPolicyAssignment Apply() method.
func (s *OsPolicyAssignmentServer) applyOsPolicyAssignment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOsconfigAlphaOsPolicyAssignmentRequest) (*alphapb.OsconfigAlphaOsPolicyAssignment, error) {
	p := ProtoToOsPolicyAssignment(request.GetResource())
	res, err := c.ApplyOsPolicyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OsPolicyAssignmentToProto(res)
	return r, nil
}

// ApplyOsPolicyAssignment handles the gRPC request by passing it to the underlying OsPolicyAssignment Apply() method.
func (s *OsPolicyAssignmentServer) ApplyOsconfigAlphaOsPolicyAssignment(ctx context.Context, request *alphapb.ApplyOsconfigAlphaOsPolicyAssignmentRequest) (*alphapb.OsconfigAlphaOsPolicyAssignment, error) {
	cl, err := createConfigOsPolicyAssignment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyOsPolicyAssignment(ctx, cl, request)
}

// DeleteOsPolicyAssignment handles the gRPC request by passing it to the underlying OsPolicyAssignment Delete() method.
func (s *OsPolicyAssignmentServer) DeleteOsconfigAlphaOsPolicyAssignment(ctx context.Context, request *alphapb.DeleteOsconfigAlphaOsPolicyAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOsPolicyAssignment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOsPolicyAssignment(ctx, ProtoToOsPolicyAssignment(request.GetResource()))

}

// ListOsconfigAlphaOsPolicyAssignment handles the gRPC request by passing it to the underlying OsPolicyAssignmentList() method.
func (s *OsPolicyAssignmentServer) ListOsconfigAlphaOsPolicyAssignment(ctx context.Context, request *alphapb.ListOsconfigAlphaOsPolicyAssignmentRequest) (*alphapb.ListOsconfigAlphaOsPolicyAssignmentResponse, error) {
	cl, err := createConfigOsPolicyAssignment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOsPolicyAssignment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OsconfigAlphaOsPolicyAssignment
	for _, r := range resources.Items {
		rp := OsPolicyAssignmentToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListOsconfigAlphaOsPolicyAssignmentResponse{Items: protos}, nil
}

func createConfigOsPolicyAssignment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
