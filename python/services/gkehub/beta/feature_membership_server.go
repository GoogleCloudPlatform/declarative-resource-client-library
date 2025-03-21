// Copyright 2025 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// FeatureMembershipServer implements the gRPC interface for FeatureMembership.
type FeatureMembershipServer struct{}

// ProtoToFeatureMembershipMeshManagementEnum converts a FeatureMembershipMeshManagementEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipMeshManagementEnum(e betapb.GkehubBetaFeatureMembershipMeshManagementEnum) *beta.FeatureMembershipMeshManagementEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipMeshManagementEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipMeshManagementEnum(n[len("GkehubBetaFeatureMembershipMeshManagementEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMeshControlPlaneEnum converts a FeatureMembershipMeshControlPlaneEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipMeshControlPlaneEnum(e betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum) *beta.FeatureMembershipMeshControlPlaneEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipMeshControlPlaneEnum(n[len("GkehubBetaFeatureMembershipMeshControlPlaneEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum converts a FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(e betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) *beta.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(n[len("GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipConfigmanagementManagementEnum converts a FeatureMembershipConfigmanagementManagementEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementManagementEnum(e betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum) *beta.FeatureMembershipConfigmanagementManagementEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipConfigmanagementManagementEnum(n[len("GkehubBetaFeatureMembershipConfigmanagementManagementEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(e betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(n[len("GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(e betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(n[len("GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(e betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(n[len("GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum(e betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum_name[int32(e)]; ok {
		e := beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum(n[len("GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMesh converts a FeatureMembershipMesh object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipMesh(p *betapb.GkehubBetaFeatureMembershipMesh) *beta.FeatureMembershipMesh {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipMesh{
		Management:   ProtoToGkehubBetaFeatureMembershipMeshManagementEnum(p.GetManagement()),
		ControlPlane: ProtoToGkehubBetaFeatureMembershipMeshControlPlaneEnum(p.GetControlPlane()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagement converts a FeatureMembershipConfigmanagement object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagement(p *betapb.GkehubBetaFeatureMembershipConfigmanagement) *beta.FeatureMembershipConfigmanagement {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagement{
		ConfigSync:          ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSync(p.GetConfigSync()),
		PolicyController:    ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyController(p.GetPolicyController()),
		Binauthz:            ProtoToGkehubBetaFeatureMembershipConfigmanagementBinauthz(p.GetBinauthz()),
		HierarchyController: ProtoToGkehubBetaFeatureMembershipConfigmanagementHierarchyController(p.GetHierarchyController()),
		Version:             dcl.StringOrNil(p.GetVersion()),
		Management:          ProtoToGkehubBetaFeatureMembershipConfigmanagementManagementEnum(p.GetManagement()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSync converts a FeatureMembershipConfigmanagementConfigSync object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSync(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync) *beta.FeatureMembershipConfigmanagementConfigSync {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSync{
		Git:                           ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncGit(p.GetGit()),
		SourceFormat:                  dcl.StringOrNil(p.GetSourceFormat()),
		Enabled:                       dcl.Bool(p.GetEnabled()),
		StopSyncing:                   dcl.Bool(p.GetStopSyncing()),
		PreventDrift:                  dcl.Bool(p.GetPreventDrift()),
		MetricsGcpServiceAccountEmail: dcl.StringOrNil(p.GetMetricsGcpServiceAccountEmail()),
		Oci:                           ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncOci(p.GetOci()),
	}
	for _, r := range p.GetDeploymentOverrides() {
		obj.DeploymentOverrides = append(obj.DeploymentOverrides, *ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides(r))
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides converts a FeatureMembershipConfigmanagementConfigSyncDeploymentOverrides object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides) *beta.FeatureMembershipConfigmanagementConfigSyncDeploymentOverrides {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSyncDeploymentOverrides{
		DeploymentName:      dcl.StringOrNil(p.GetDeploymentName()),
		DeploymentNamespace: dcl.StringOrNil(p.GetDeploymentNamespace()),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers(r))
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers converts a FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers) *beta.FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers{
		ContainerName: dcl.StringOrNil(p.GetContainerName()),
		CpuRequest:    dcl.StringOrNil(p.GetCpuRequest()),
		MemoryRequest: dcl.StringOrNil(p.GetMemoryRequest()),
		CpuLimit:      dcl.StringOrNil(p.GetCpuLimit()),
		MemoryLimit:   dcl.StringOrNil(p.GetMemoryLimit()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncGit converts a FeatureMembershipConfigmanagementConfigSyncGit object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncGit(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit) *beta.FeatureMembershipConfigmanagementConfigSyncGit {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSyncGit{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		SyncBranch:             dcl.StringOrNil(p.GetSyncBranch()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SyncRev:                dcl.StringOrNil(p.GetSyncRev()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		HttpsProxy:             dcl.StringOrNil(p.GetHttpsProxy()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncOci converts a FeatureMembershipConfigmanagementConfigSyncOci object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncOci(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncOci) *beta.FeatureMembershipConfigmanagementConfigSyncOci {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSyncOci{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyController converts a FeatureMembershipConfigmanagementPolicyController object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyController(p *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController) *beta.FeatureMembershipConfigmanagementPolicyController {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementPolicyController{
		Enabled:                  dcl.Bool(p.GetEnabled()),
		ReferentialRulesEnabled:  dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:         dcl.Bool(p.GetLogDeniesEnabled()),
		MutationEnabled:          dcl.Bool(p.GetMutationEnabled()),
		Monitoring:               ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoring(p.GetMonitoring()),
		TemplateLibraryInstalled: dcl.Bool(p.GetTemplateLibraryInstalled()),
		AuditIntervalSeconds:     dcl.StringOrNil(p.GetAuditIntervalSeconds()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerMonitoring converts a FeatureMembershipConfigmanagementPolicyControllerMonitoring object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoring(p *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoring) *beta.FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(r))
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementBinauthz converts a FeatureMembershipConfigmanagementBinauthz object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementBinauthz(p *betapb.GkehubBetaFeatureMembershipConfigmanagementBinauthz) *beta.FeatureMembershipConfigmanagementBinauthz {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementHierarchyController converts a FeatureMembershipConfigmanagementHierarchyController object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementHierarchyController(p *betapb.GkehubBetaFeatureMembershipConfigmanagementHierarchyController) *beta.FeatureMembershipConfigmanagementHierarchyController {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementHierarchyController{
		Enabled:                         dcl.Bool(p.GetEnabled()),
		EnablePodTreeLabels:             dcl.Bool(p.GetEnablePodTreeLabels()),
		EnableHierarchicalResourceQuota: dcl.Bool(p.GetEnableHierarchicalResourceQuota()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontroller converts a FeatureMembershipPolicycontroller object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontroller(p *betapb.GkehubBetaFeatureMembershipPolicycontroller) *beta.FeatureMembershipPolicycontroller {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontroller{
		Version:                   dcl.StringOrNil(p.GetVersion()),
		PolicyControllerHubConfig: ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfig(p.GetPolicyControllerHubConfig()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfig converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfig object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfig(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfig) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfig{
		InstallSpec:              ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(p.GetInstallSpec()),
		ReferentialRulesEnabled:  dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:         dcl.Bool(p.GetLogDeniesEnabled()),
		MutationEnabled:          dcl.Bool(p.GetMutationEnabled()),
		Monitoring:               ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(p.GetMonitoring()),
		AuditIntervalSeconds:     dcl.Int64OrNil(p.GetAuditIntervalSeconds()),
		ConstraintViolationLimit: dcl.Int64OrNil(p.GetConstraintViolationLimit()),
		PolicyContent:            ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(p.GetPolicyContent()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(r))
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{
		TemplateLibrary: ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(p.GetTemplateLibrary()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{
		Installation: ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(p.GetInstallation()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles{}
	for _, r := range p.GetExemptedNamespaces() {
		obj.ExemptedNamespaces = append(obj.ExemptedNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs{
		ReplicaCount:       dcl.Int64OrNil(p.GetReplicaCount()),
		ContainerResources: ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources(p.GetContainerResources()),
		PodAffinity:        ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum(p.GetPodAffinity()),
	}
	for _, r := range p.GetPodTolerations() {
		obj.PodTolerations = append(obj.PodTolerations, *ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations(r))
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources{
		Limits:   ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits(p.GetLimits()),
		Requests: ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests(p.GetRequests()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits{
		Memory: dcl.StringOrNil(p.GetMemory()),
		Cpu:    dcl.StringOrNil(p.GetCpu()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests{
		Memory: dcl.StringOrNil(p.GetMemory()),
		Cpu:    dcl.StringOrNil(p.GetCpu()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations(p *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations) *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations{
		Key:      dcl.StringOrNil(p.GetKey()),
		Operator: dcl.StringOrNil(p.GetOperator()),
		Value:    dcl.StringOrNil(p.GetValue()),
		Effect:   dcl.StringOrNil(p.GetEffect()),
	}
	return obj
}

// ProtoToFeatureMembership converts a FeatureMembership resource from its proto representation.
func ProtoToFeatureMembership(p *betapb.GkehubBetaFeatureMembership) *beta.FeatureMembership {
	obj := &beta.FeatureMembership{
		Mesh:               ProtoToGkehubBetaFeatureMembershipMesh(p.GetMesh()),
		Configmanagement:   ProtoToGkehubBetaFeatureMembershipConfigmanagement(p.GetConfigmanagement()),
		Policycontroller:   ProtoToGkehubBetaFeatureMembershipPolicycontroller(p.GetPolicycontroller()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		Feature:            dcl.StringOrNil(p.GetFeature()),
		Membership:         dcl.StringOrNil(p.GetMembership()),
		MembershipLocation: dcl.StringOrNil(p.GetMembershipLocation()),
	}
	return obj
}

// FeatureMembershipMeshManagementEnumToProto converts a FeatureMembershipMeshManagementEnum enum to its proto representation.
func GkehubBetaFeatureMembershipMeshManagementEnumToProto(e *beta.FeatureMembershipMeshManagementEnum) betapb.GkehubBetaFeatureMembershipMeshManagementEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipMeshManagementEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipMeshManagementEnum_value["FeatureMembershipMeshManagementEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipMeshManagementEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipMeshManagementEnum(0)
}

// FeatureMembershipMeshControlPlaneEnumToProto converts a FeatureMembershipMeshControlPlaneEnum enum to its proto representation.
func GkehubBetaFeatureMembershipMeshControlPlaneEnumToProto(e *beta.FeatureMembershipMeshControlPlaneEnum) betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum_value["FeatureMembershipMeshControlPlaneEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipMeshControlPlaneEnum(0)
}

// FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumToProto converts a FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum enum to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumToProto(e *beta.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_value["FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(0)
}

// FeatureMembershipConfigmanagementManagementEnumToProto converts a FeatureMembershipConfigmanagementManagementEnum enum to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementManagementEnumToProto(e *beta.FeatureMembershipConfigmanagementManagementEnum) betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum_value["FeatureMembershipConfigmanagementManagementEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipConfigmanagementManagementEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum enum to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumToProto(e *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum) betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum enum to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumToProto(e *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum) betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum enum to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumToProto(e *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum) betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum enum to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnumToProto(e *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum) betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum(v)
	}
	return betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnum(0)
}

// FeatureMembershipMeshToProto converts a FeatureMembershipMesh object to its proto representation.
func GkehubBetaFeatureMembershipMeshToProto(o *beta.FeatureMembershipMesh) *betapb.GkehubBetaFeatureMembershipMesh {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipMesh{}
	p.SetManagement(GkehubBetaFeatureMembershipMeshManagementEnumToProto(o.Management))
	p.SetControlPlane(GkehubBetaFeatureMembershipMeshControlPlaneEnumToProto(o.ControlPlane))
	return p
}

// FeatureMembershipConfigmanagementToProto converts a FeatureMembershipConfigmanagement object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementToProto(o *beta.FeatureMembershipConfigmanagement) *betapb.GkehubBetaFeatureMembershipConfigmanagement {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagement{}
	p.SetConfigSync(GkehubBetaFeatureMembershipConfigmanagementConfigSyncToProto(o.ConfigSync))
	p.SetPolicyController(GkehubBetaFeatureMembershipConfigmanagementPolicyControllerToProto(o.PolicyController))
	p.SetBinauthz(GkehubBetaFeatureMembershipConfigmanagementBinauthzToProto(o.Binauthz))
	p.SetHierarchyController(GkehubBetaFeatureMembershipConfigmanagementHierarchyControllerToProto(o.HierarchyController))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetManagement(GkehubBetaFeatureMembershipConfigmanagementManagementEnumToProto(o.Management))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncToProto converts a FeatureMembershipConfigmanagementConfigSync object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncToProto(o *beta.FeatureMembershipConfigmanagementConfigSync) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync{}
	p.SetGit(GkehubBetaFeatureMembershipConfigmanagementConfigSyncGitToProto(o.Git))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetStopSyncing(dcl.ValueOrEmptyBool(o.StopSyncing))
	p.SetPreventDrift(dcl.ValueOrEmptyBool(o.PreventDrift))
	p.SetMetricsGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.MetricsGcpServiceAccountEmail))
	p.SetOci(GkehubBetaFeatureMembershipConfigmanagementConfigSyncOciToProto(o.Oci))
	sDeploymentOverrides := make([]*betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides, len(o.DeploymentOverrides))
	for i, r := range o.DeploymentOverrides {
		sDeploymentOverrides[i] = GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesToProto(&r)
	}
	p.SetDeploymentOverrides(sDeploymentOverrides)
	return p
}

// FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesToProto converts a FeatureMembershipConfigmanagementConfigSyncDeploymentOverrides object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesToProto(o *beta.FeatureMembershipConfigmanagementConfigSyncDeploymentOverrides) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverrides{}
	p.SetDeploymentName(dcl.ValueOrEmptyString(o.DeploymentName))
	p.SetDeploymentNamespace(dcl.ValueOrEmptyString(o.DeploymentNamespace))
	sContainers := make([]*betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers, len(o.Containers))
	for i, r := range o.Containers {
		sContainers[i] = GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersToProto(&r)
	}
	p.SetContainers(sContainers)
	return p
}

// FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersToProto converts a FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainersToProto(o *beta.FeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesContainers{}
	p.SetContainerName(dcl.ValueOrEmptyString(o.ContainerName))
	p.SetCpuRequest(dcl.ValueOrEmptyString(o.CpuRequest))
	p.SetMemoryRequest(dcl.ValueOrEmptyString(o.MemoryRequest))
	p.SetCpuLimit(dcl.ValueOrEmptyString(o.CpuLimit))
	p.SetMemoryLimit(dcl.ValueOrEmptyString(o.MemoryLimit))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncGitToProto converts a FeatureMembershipConfigmanagementConfigSyncGit object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncGitToProto(o *beta.FeatureMembershipConfigmanagementConfigSyncGit) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetSyncBranch(dcl.ValueOrEmptyString(o.SyncBranch))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSyncRev(dcl.ValueOrEmptyString(o.SyncRev))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetHttpsProxy(dcl.ValueOrEmptyString(o.HttpsProxy))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncOciToProto converts a FeatureMembershipConfigmanagementConfigSyncOci object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncOciToProto(o *beta.FeatureMembershipConfigmanagementConfigSyncOci) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncOci {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncOci{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerToProto converts a FeatureMembershipConfigmanagementPolicyController object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementPolicyControllerToProto(o *beta.FeatureMembershipConfigmanagementPolicyController) *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetMutationEnabled(dcl.ValueOrEmptyBool(o.MutationEnabled))
	p.SetMonitoring(GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringToProto(o.Monitoring))
	p.SetTemplateLibraryInstalled(dcl.ValueOrEmptyBool(o.TemplateLibraryInstalled))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyString(o.AuditIntervalSeconds))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerMonitoringToProto converts a FeatureMembershipConfigmanagementPolicyControllerMonitoring object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringToProto(o *beta.FeatureMembershipConfigmanagementPolicyControllerMonitoring) *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	sBackends := make([]betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum, len(o.Backends))
	for i, r := range o.Backends {
		sBackends[i] = betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_value[string(r)])
	}
	p.SetBackends(sBackends)
	return p
}

// FeatureMembershipConfigmanagementBinauthzToProto converts a FeatureMembershipConfigmanagementBinauthz object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementBinauthzToProto(o *beta.FeatureMembershipConfigmanagementBinauthz) *betapb.GkehubBetaFeatureMembershipConfigmanagementBinauthz {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementBinauthz{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// FeatureMembershipConfigmanagementHierarchyControllerToProto converts a FeatureMembershipConfigmanagementHierarchyController object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementHierarchyControllerToProto(o *beta.FeatureMembershipConfigmanagementHierarchyController) *betapb.GkehubBetaFeatureMembershipConfigmanagementHierarchyController {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementHierarchyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetEnablePodTreeLabels(dcl.ValueOrEmptyBool(o.EnablePodTreeLabels))
	p.SetEnableHierarchicalResourceQuota(dcl.ValueOrEmptyBool(o.EnableHierarchicalResourceQuota))
	return p
}

// FeatureMembershipPolicycontrollerToProto converts a FeatureMembershipPolicycontroller object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerToProto(o *beta.FeatureMembershipPolicycontroller) *betapb.GkehubBetaFeatureMembershipPolicycontroller {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontroller{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetPolicyControllerHubConfig(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigToProto(o.PolicyControllerHubConfig))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfig object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfig) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	p.SetInstallSpec(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumToProto(o.InstallSpec))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetMutationEnabled(dcl.ValueOrEmptyBool(o.MutationEnabled))
	p.SetMonitoring(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringToProto(o.Monitoring))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyInt64(o.AuditIntervalSeconds))
	p.SetConstraintViolationLimit(dcl.ValueOrEmptyInt64(o.ConstraintViolationLimit))
	p.SetPolicyContent(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentToProto(o.PolicyContent))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
	mDeploymentConfigs := make(map[string]*betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs, len(o.DeploymentConfigs))
	for k, r := range o.DeploymentConfigs {
		mDeploymentConfigs[k] = GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsToProto(&r)
	}
	p.SetDeploymentConfigs(mDeploymentConfigs)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	sBackends := make([]betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum, len(o.Backends))
	for i, r := range o.Backends {
		sBackends[i] = betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum_value[string(r)])
	}
	p.SetBackends(sBackends)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	p.SetTemplateLibrary(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryToProto(o.TemplateLibrary))
	mBundles := make(map[string]*betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles, len(o.Bundles))
	for k, r := range o.Bundles {
		mBundles[k] = GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundlesToProto(&r)
	}
	p.SetBundles(mBundles)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	p.SetInstallation(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumToProto(o.Installation))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundlesToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundlesToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles{}
	sExemptedNamespaces := make([]string, len(o.ExemptedNamespaces))
	for i, r := range o.ExemptedNamespaces {
		sExemptedNamespaces[i] = r
	}
	p.SetExemptedNamespaces(sExemptedNamespaces)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigs{}
	p.SetReplicaCount(dcl.ValueOrEmptyInt64(o.ReplicaCount))
	p.SetContainerResources(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesToProto(o.ContainerResources))
	p.SetPodAffinity(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodAffinityEnumToProto(o.PodAffinity))
	sPodTolerations := make([]*betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations, len(o.PodTolerations))
	for i, r := range o.PodTolerations {
		sPodTolerations[i] = GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerationsToProto(&r)
	}
	p.SetPodTolerations(sPodTolerations)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources{}
	p.SetLimits(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimitsToProto(o.Limits))
	p.SetRequests(GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequestsToProto(o.Requests))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimitsToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimitsToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits{}
	p.SetMemory(dcl.ValueOrEmptyString(o.Memory))
	p.SetCpu(dcl.ValueOrEmptyString(o.Cpu))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequestsToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequestsToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests{}
	p.SetMemory(dcl.ValueOrEmptyString(o.Memory))
	p.SetCpu(dcl.ValueOrEmptyString(o.Cpu))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerationsToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations object to its proto representation.
func GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerationsToProto(o *beta.FeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations) *betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsPodTolerations{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetOperator(dcl.ValueOrEmptyString(o.Operator))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetEffect(dcl.ValueOrEmptyString(o.Effect))
	return p
}

// FeatureMembershipToProto converts a FeatureMembership resource to its proto representation.
func FeatureMembershipToProto(resource *beta.FeatureMembership) *betapb.GkehubBetaFeatureMembership {
	p := &betapb.GkehubBetaFeatureMembership{}
	p.SetMesh(GkehubBetaFeatureMembershipMeshToProto(resource.Mesh))
	p.SetConfigmanagement(GkehubBetaFeatureMembershipConfigmanagementToProto(resource.Configmanagement))
	p.SetPolicycontroller(GkehubBetaFeatureMembershipPolicycontrollerToProto(resource.Policycontroller))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFeature(dcl.ValueOrEmptyString(resource.Feature))
	p.SetMembership(dcl.ValueOrEmptyString(resource.Membership))
	p.SetMembershipLocation(dcl.ValueOrEmptyString(resource.MembershipLocation))

	return p
}

// applyFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) applyFeatureMembership(ctx context.Context, c *beta.Client, request *betapb.ApplyGkehubBetaFeatureMembershipRequest) (*betapb.GkehubBetaFeatureMembership, error) {
	p := ProtoToFeatureMembership(request.GetResource())
	res, err := c.ApplyFeatureMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureMembershipToProto(res)
	return r, nil
}

// applyGkehubBetaFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) ApplyGkehubBetaFeatureMembership(ctx context.Context, request *betapb.ApplyGkehubBetaFeatureMembershipRequest) (*betapb.GkehubBetaFeatureMembership, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeatureMembership(ctx, cl, request)
}

// DeleteFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Delete() method.
func (s *FeatureMembershipServer) DeleteGkehubBetaFeatureMembership(ctx context.Context, request *betapb.DeleteGkehubBetaFeatureMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeatureMembership(ctx, ProtoToFeatureMembership(request.GetResource()))

}

// ListGkehubBetaFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembershipList() method.
func (s *FeatureMembershipServer) ListGkehubBetaFeatureMembership(ctx context.Context, request *betapb.ListGkehubBetaFeatureMembershipRequest) (*betapb.ListGkehubBetaFeatureMembershipResponse, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeatureMembership(ctx, request.GetProject(), request.GetLocation(), request.GetFeature())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkehubBetaFeatureMembership
	for _, r := range resources.Items {
		rp := FeatureMembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListGkehubBetaFeatureMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeatureMembership(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
