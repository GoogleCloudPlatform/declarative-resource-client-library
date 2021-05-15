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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// Server implements the gRPC interface for FeatureMembership.
type FeatureMembershipServer struct{}

// ProtoToFeatureMembershipConfigmanagement converts a FeatureMembershipConfigmanagement resource from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagement(p *betapb.GkehubBetaFeatureMembershipConfigmanagement) *beta.FeatureMembershipConfigmanagement {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagement{
		ConfigSync:          ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSync(p.GetConfigSync()),
		PolicyController:    ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyController(p.GetPolicyController()),
		Binauthz:            ProtoToGkehubBetaFeatureMembershipConfigmanagementBinauthz(p.GetBinauthz()),
		HierarchyController: ProtoToGkehubBetaFeatureMembershipConfigmanagementHierarchyController(p.GetHierarchyController()),
		Version:             dcl.StringOrNil(p.Version),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSync converts a FeatureMembershipConfigmanagementConfigSync resource from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSync(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync) *beta.FeatureMembershipConfigmanagementConfigSync {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSync{
		Git:          ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncGit(p.GetGit()),
		SourceFormat: dcl.StringOrNil(p.SourceFormat),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncGit converts a FeatureMembershipConfigmanagementConfigSyncGit resource from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementConfigSyncGit(p *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit) *beta.FeatureMembershipConfigmanagementConfigSyncGit {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementConfigSyncGit{
		SyncRepo:     dcl.StringOrNil(p.SyncRepo),
		SyncBranch:   dcl.StringOrNil(p.SyncBranch),
		PolicyDir:    dcl.StringOrNil(p.PolicyDir),
		SyncWaitSecs: dcl.StringOrNil(p.SyncWaitSecs),
		SyncRev:      dcl.StringOrNil(p.SyncRev),
		SecretType:   dcl.StringOrNil(p.SecretType),
		HttpsProxy:   dcl.StringOrNil(p.HttpsProxy),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyController converts a FeatureMembershipConfigmanagementPolicyController resource from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementPolicyController(p *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController) *beta.FeatureMembershipConfigmanagementPolicyController {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementPolicyController{
		Enabled:                  dcl.Bool(p.Enabled),
		ReferentialRulesEnabled:  dcl.Bool(p.ReferentialRulesEnabled),
		LogDeniesEnabled:         dcl.Bool(p.LogDeniesEnabled),
		TemplateLibraryInstalled: dcl.Bool(p.TemplateLibraryInstalled),
		AuditIntervalSeconds:     dcl.StringOrNil(p.AuditIntervalSeconds),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementBinauthz converts a FeatureMembershipConfigmanagementBinauthz resource from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementBinauthz(p *betapb.GkehubBetaFeatureMembershipConfigmanagementBinauthz) *beta.FeatureMembershipConfigmanagementBinauthz {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementHierarchyController converts a FeatureMembershipConfigmanagementHierarchyController resource from its proto representation.
func ProtoToGkehubBetaFeatureMembershipConfigmanagementHierarchyController(p *betapb.GkehubBetaFeatureMembershipConfigmanagementHierarchyController) *beta.FeatureMembershipConfigmanagementHierarchyController {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureMembershipConfigmanagementHierarchyController{
		Enabled:                         dcl.Bool(p.Enabled),
		EnablePodTreeLabels:             dcl.Bool(p.EnablePodTreeLabels),
		EnableHierarchicalResourceQuota: dcl.Bool(p.EnableHierarchicalResourceQuota),
	}
	return obj
}

// ProtoToFeatureMembership converts a FeatureMembership resource from its proto representation.
func ProtoToFeatureMembership(p *betapb.GkehubBetaFeatureMembership) *beta.FeatureMembership {
	obj := &beta.FeatureMembership{
		Membership:       dcl.StringOrNil(p.Membership),
		Feature:          dcl.StringOrNil(p.Feature),
		Location:         dcl.StringOrNil(p.Location),
		Project:          dcl.StringOrNil(p.Project),
		Configmanagement: ProtoToGkehubBetaFeatureMembershipConfigmanagement(p.GetConfigmanagement()),
	}
	return obj
}

// FeatureMembershipConfigmanagementToProto converts a FeatureMembershipConfigmanagement resource to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementToProto(o *beta.FeatureMembershipConfigmanagement) *betapb.GkehubBetaFeatureMembershipConfigmanagement {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagement{
		ConfigSync:          GkehubBetaFeatureMembershipConfigmanagementConfigSyncToProto(o.ConfigSync),
		PolicyController:    GkehubBetaFeatureMembershipConfigmanagementPolicyControllerToProto(o.PolicyController),
		Binauthz:            GkehubBetaFeatureMembershipConfigmanagementBinauthzToProto(o.Binauthz),
		HierarchyController: GkehubBetaFeatureMembershipConfigmanagementHierarchyControllerToProto(o.HierarchyController),
		Version:             dcl.ValueOrEmptyString(o.Version),
	}
	return p
}

// FeatureMembershipConfigmanagementConfigSyncToProto converts a FeatureMembershipConfigmanagementConfigSync resource to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncToProto(o *beta.FeatureMembershipConfigmanagementConfigSync) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSync{
		Git:          GkehubBetaFeatureMembershipConfigmanagementConfigSyncGitToProto(o.Git),
		SourceFormat: dcl.ValueOrEmptyString(o.SourceFormat),
	}
	return p
}

// FeatureMembershipConfigmanagementConfigSyncGitToProto converts a FeatureMembershipConfigmanagementConfigSyncGit resource to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementConfigSyncGitToProto(o *beta.FeatureMembershipConfigmanagementConfigSyncGit) *betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit{
		SyncRepo:     dcl.ValueOrEmptyString(o.SyncRepo),
		SyncBranch:   dcl.ValueOrEmptyString(o.SyncBranch),
		PolicyDir:    dcl.ValueOrEmptyString(o.PolicyDir),
		SyncWaitSecs: dcl.ValueOrEmptyString(o.SyncWaitSecs),
		SyncRev:      dcl.ValueOrEmptyString(o.SyncRev),
		SecretType:   dcl.ValueOrEmptyString(o.SecretType),
		HttpsProxy:   dcl.ValueOrEmptyString(o.HttpsProxy),
	}
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerToProto converts a FeatureMembershipConfigmanagementPolicyController resource to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementPolicyControllerToProto(o *beta.FeatureMembershipConfigmanagementPolicyController) *betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementPolicyController{
		Enabled:                  dcl.ValueOrEmptyBool(o.Enabled),
		ReferentialRulesEnabled:  dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled),
		LogDeniesEnabled:         dcl.ValueOrEmptyBool(o.LogDeniesEnabled),
		TemplateLibraryInstalled: dcl.ValueOrEmptyBool(o.TemplateLibraryInstalled),
		AuditIntervalSeconds:     dcl.ValueOrEmptyString(o.AuditIntervalSeconds),
	}
	for _, r := range o.ExemptableNamespaces {
		p.ExemptableNamespaces = append(p.ExemptableNamespaces, r)
	}
	return p
}

// FeatureMembershipConfigmanagementBinauthzToProto converts a FeatureMembershipConfigmanagementBinauthz resource to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementBinauthzToProto(o *beta.FeatureMembershipConfigmanagementBinauthz) *betapb.GkehubBetaFeatureMembershipConfigmanagementBinauthz {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// FeatureMembershipConfigmanagementHierarchyControllerToProto converts a FeatureMembershipConfigmanagementHierarchyController resource to its proto representation.
func GkehubBetaFeatureMembershipConfigmanagementHierarchyControllerToProto(o *beta.FeatureMembershipConfigmanagementHierarchyController) *betapb.GkehubBetaFeatureMembershipConfigmanagementHierarchyController {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureMembershipConfigmanagementHierarchyController{
		Enabled:                         dcl.ValueOrEmptyBool(o.Enabled),
		EnablePodTreeLabels:             dcl.ValueOrEmptyBool(o.EnablePodTreeLabels),
		EnableHierarchicalResourceQuota: dcl.ValueOrEmptyBool(o.EnableHierarchicalResourceQuota),
	}
	return p
}

// FeatureMembershipToProto converts a FeatureMembership resource to its proto representation.
func FeatureMembershipToProto(resource *beta.FeatureMembership) *betapb.GkehubBetaFeatureMembership {
	p := &betapb.GkehubBetaFeatureMembership{
		Membership:       dcl.ValueOrEmptyString(resource.Membership),
		Feature:          dcl.ValueOrEmptyString(resource.Feature),
		Location:         dcl.ValueOrEmptyString(resource.Location),
		Project:          dcl.ValueOrEmptyString(resource.Project),
		Configmanagement: GkehubBetaFeatureMembershipConfigmanagementToProto(resource.Configmanagement),
	}

	return p
}

// ApplyFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) applyFeatureMembership(ctx context.Context, c *beta.Client, request *betapb.ApplyGkehubBetaFeatureMembershipRequest) (*betapb.GkehubBetaFeatureMembership, error) {
	p := ProtoToFeatureMembership(request.GetResource())
	res, err := c.ApplyFeatureMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureMembershipToProto(res)
	return r, nil
}

// ApplyFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) ApplyGkehubBetaFeatureMembership(ctx context.Context, request *betapb.ApplyGkehubBetaFeatureMembershipRequest) (*betapb.GkehubBetaFeatureMembership, error) {
	cl, err := createConfigFeatureMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFeatureMembership(ctx, cl, request)
}

// DeleteFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Delete() method.
func (s *FeatureMembershipServer) DeleteGkehubBetaFeatureMembership(ctx context.Context, request *betapb.DeleteGkehubBetaFeatureMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeatureMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeatureMembership(ctx, ProtoToFeatureMembership(request.GetResource()))

}

// ListFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembershipList() method.
func (s *FeatureMembershipServer) ListGkehubBetaFeatureMembership(ctx context.Context, request *betapb.ListGkehubBetaFeatureMembershipRequest) (*betapb.ListGkehubBetaFeatureMembershipResponse, error) {
	cl, err := createConfigFeatureMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeatureMembership(ctx, request.Project, request.Location, request.Feature)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkehubBetaFeatureMembership
	for _, r := range resources.Items {
		rp := FeatureMembershipToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkehubBetaFeatureMembershipResponse{Items: protos}, nil
}

func createConfigFeatureMembership(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
