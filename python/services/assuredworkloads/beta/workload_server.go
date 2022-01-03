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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/beta/assuredworkloads_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
)

// Server implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(e betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum) *beta.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := beta.WorkloadResourcesResourceTypeEnum(n[len("AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadComplianceRegimeEnum converts a WorkloadComplianceRegimeEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadComplianceRegimeEnum(e betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum) *beta.WorkloadComplianceRegimeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum_name[int32(e)]; ok {
		e := beta.WorkloadComplianceRegimeEnum(n[len("AssuredworkloadsBetaWorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResourceSettingsResourceTypeEnum converts a WorkloadResourceSettingsResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(e betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum) *beta.WorkloadResourceSettingsResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum_name[int32(e)]; ok {
		e := beta.WorkloadResourceSettingsResourceTypeEnum(n[len("AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResources(p *betapb.AssuredworkloadsBetaWorkloadResources) *beta.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadKmsSettings(p *betapb.AssuredworkloadsBetaWorkloadKmsSettings) *beta.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.GetRotationPeriod()),
	}
	return obj
}

// ProtoToWorkloadResourceSettings converts a WorkloadResourceSettings object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResourceSettings(p *betapb.AssuredworkloadsBetaWorkloadResourceSettings) *beta.WorkloadResourceSettings {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadResourceSettings{
		ResourceId:   dcl.StringOrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *betapb.AssuredworkloadsBetaWorkload) *beta.Workload {
	obj := &beta.Workload{
		Name:                       dcl.StringOrNil(p.GetName()),
		DisplayName:                dcl.StringOrNil(p.GetDisplayName()),
		ComplianceRegime:           ProtoToAssuredworkloadsBetaWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                 dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:             dcl.StringOrNil(p.GetBillingAccount()),
		ProvisionedResourcesParent: dcl.StringOrNil(p.GetProvisionedResourcesParent()),
		KmsSettings:                ProtoToAssuredworkloadsBetaWorkloadKmsSettings(p.GetKmsSettings()),
		Organization:               dcl.StringOrNil(p.GetOrganization()),
		Location:                   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsBetaWorkloadResources(r))
	}
	for _, r := range p.GetResourceSettings() {
		obj.ResourceSettings = append(obj.ResourceSettings, *ProtoToAssuredworkloadsBetaWorkloadResourceSettings(r))
	}
	return obj
}

// WorkloadResourcesResourceTypeEnumToProto converts a WorkloadResourcesResourceTypeEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadResourcesResourceTypeEnumToProto(e *beta.WorkloadResourcesResourceTypeEnum) betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum_value["WorkloadResourcesResourceTypeEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(0)
}

// WorkloadComplianceRegimeEnumToProto converts a WorkloadComplianceRegimeEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadComplianceRegimeEnumToProto(e *beta.WorkloadComplianceRegimeEnum) betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum_value["WorkloadComplianceRegimeEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum(0)
}

// WorkloadResourceSettingsResourceTypeEnumToProto converts a WorkloadResourceSettingsResourceTypeEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnumToProto(e *beta.WorkloadResourceSettingsResourceTypeEnum) betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum_value["WorkloadResourceSettingsResourceTypeEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(0)
}

// WorkloadResourcesToProto converts a WorkloadResources object to its proto representation.
func AssuredworkloadsBetaWorkloadResourcesToProto(o *beta.WorkloadResources) *betapb.AssuredworkloadsBetaWorkloadResources {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadResources{}
	p.SetResourceId(dcl.ValueOrEmptyInt64(o.ResourceId))
	p.SetResourceType(AssuredworkloadsBetaWorkloadResourcesResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings object to its proto representation.
func AssuredworkloadsBetaWorkloadKmsSettingsToProto(o *beta.WorkloadKmsSettings) *betapb.AssuredworkloadsBetaWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadKmsSettings{}
	p.SetNextRotationTime(dcl.ValueOrEmptyString(o.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(o.RotationPeriod))
	return p
}

// WorkloadResourceSettingsToProto converts a WorkloadResourceSettings object to its proto representation.
func AssuredworkloadsBetaWorkloadResourceSettingsToProto(o *beta.WorkloadResourceSettings) *betapb.AssuredworkloadsBetaWorkloadResourceSettings {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadResourceSettings{}
	p.SetResourceId(dcl.ValueOrEmptyString(o.ResourceId))
	p.SetResourceType(AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *beta.Workload) *betapb.AssuredworkloadsBetaWorkload {
	p := &betapb.AssuredworkloadsBetaWorkload{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetComplianceRegime(AssuredworkloadsBetaWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	p.SetProvisionedResourcesParent(dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent))
	p.SetKmsSettings(AssuredworkloadsBetaWorkloadKmsSettingsToProto(resource.KmsSettings))
	p.SetOrganization(dcl.ValueOrEmptyString(resource.Organization))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sResources := make([]*betapb.AssuredworkloadsBetaWorkloadResources, len(resource.Resources))
	for i, r := range resource.Resources {
		sResources[i] = AssuredworkloadsBetaWorkloadResourcesToProto(&r)
	}
	p.SetResources(sResources)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sResourceSettings := make([]*betapb.AssuredworkloadsBetaWorkloadResourceSettings, len(resource.ResourceSettings))
	for i, r := range resource.ResourceSettings {
		sResourceSettings[i] = AssuredworkloadsBetaWorkloadResourceSettingsToProto(&r)
	}
	p.SetResourceSettings(sResourceSettings)

	return p
}

// applyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *beta.Client, request *betapb.ApplyAssuredworkloadsBetaWorkloadRequest) (*betapb.AssuredworkloadsBetaWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// applyAssuredworkloadsBetaWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsBetaWorkload(ctx context.Context, request *betapb.ApplyAssuredworkloadsBetaWorkloadRequest) (*betapb.AssuredworkloadsBetaWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsBetaWorkload(ctx context.Context, request *betapb.DeleteAssuredworkloadsBetaWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListAssuredworkloadsBetaWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsBetaWorkload(ctx context.Context, request *betapb.ListAssuredworkloadsBetaWorkloadRequest) (*betapb.ListAssuredworkloadsBetaWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, request.GetOrganization(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.AssuredworkloadsBetaWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListAssuredworkloadsBetaWorkloadResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
