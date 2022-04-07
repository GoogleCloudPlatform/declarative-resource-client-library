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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// FeatureMembershipServer implements the gRPC interface for FeatureMembership.
type FeatureMembershipServer struct{}

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
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSync converts a FeatureMembershipConfigmanagementConfigSync object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSync(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync) *beta.FeatureMembershipConfigmanagementConfigSync {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSync{
		Git:          ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncGit(p.GetGit()),
		SourceFormat: dcl.StringOrNil(p.GetSourceFormat()),
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

// ProtoToFeatureMembershipConfigmanagementPolicyController converts a FeatureMembershipConfigmanagementPolicyController object from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyController(p *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController) *beta.FeatureMembershipConfigmanagementPolicyController {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementPolicyController{
		Enabled:                  dcl.Bool(p.GetEnabled()),
		ReferentialRulesEnabled:  dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:         dcl.Bool(p.GetLogDeniesEnabled()),
		TemplateLibraryInstalled: dcl.Bool(p.GetTemplateLibraryInstalled()),
		AuditIntervalSeconds:     dcl.StringOrNil(p.GetAuditIntervalSeconds()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
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

// ProtoToFeatureMembership converts a FeatureMembership resource from its proto representation.
func ProtoToFeatureMembership(p *betapb.GkehubBetaFeatureMembership) *beta.FeatureMembership {
	obj := &beta.FeatureMembership{
		Configmanagement: ProtoToGkehubBetaFeatureMembershipConfigmanagement(p.GetConfigmanagement()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Location:         dcl.StringOrNil(p.GetLocation()),
		Feature:          dcl.StringOrNil(p.GetFeature()),
		Membership:       dcl.StringOrNil(p.GetMembership()),
	}
	return obj
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

// FeatureMembershipConfigmanagementPolicyControllerToProto converts a FeatureMembershipConfigmanagementPolicyController object to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementPolicyControllerToProto(o *beta.FeatureMembershipConfigmanagementPolicyController) *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetTemplateLibraryInstalled(dcl.ValueOrEmptyBool(o.TemplateLibraryInstalled))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyString(o.AuditIntervalSeconds))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
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

// FeatureMembershipToProto converts a FeatureMembership resource to its proto representation.
func FeatureMembershipToProto(resource *beta.FeatureMembership) *betapb.GkehubBetaFeatureMembership {
	p := &betapb.GkehubBetaFeatureMembership{}
	p.SetConfigmanagement(GkehubBetaFeatureMembershipConfigmanagementToProto(resource.Configmanagement))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFeature(dcl.ValueOrEmptyString(resource.Feature))
	p.SetMembership(dcl.ValueOrEmptyString(resource.Membership))

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
