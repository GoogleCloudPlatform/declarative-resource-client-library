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
	assuredworkloadspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/assuredworkloads_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads"
)

// WorkloadServer implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourcesResourceTypeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum) *assuredworkloads.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadResourcesResourceTypeEnum(n[len("AssuredworkloadsWorkloadResourcesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadComplianceRegimeEnum converts a WorkloadComplianceRegimeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadComplianceRegimeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum) *assuredworkloads.WorkloadComplianceRegimeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadComplianceRegimeEnum(n[len("AssuredworkloadsWorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResourceSettingsResourceTypeEnum converts a WorkloadResourceSettingsResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum) *assuredworkloads.WorkloadResourceSettingsResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadResourceSettingsResourceTypeEnum(n[len("AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources object from its proto representation.
func ProtoToAssuredworkloadsWorkloadResources(p *assuredworkloadspb.AssuredworkloadsWorkloadResources) *assuredworkloads.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings object from its proto representation.
func ProtoToAssuredworkloadsWorkloadKmsSettings(p *assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings) *assuredworkloads.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.GetRotationPeriod()),
	}
	return obj
}

// ProtoToWorkloadResourceSettings converts a WorkloadResourceSettings object from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourceSettings(p *assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings) *assuredworkloads.WorkloadResourceSettings {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadResourceSettings{
		ResourceId:   dcl.StringOrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *assuredworkloadspb.AssuredworkloadsWorkload) *assuredworkloads.Workload {
	obj := &assuredworkloads.Workload{
		Name:                       dcl.StringOrNil(p.GetName()),
		DisplayName:                dcl.StringOrNil(p.GetDisplayName()),
		ComplianceRegime:           ProtoToAssuredworkloadsWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                 dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:             dcl.StringOrNil(p.GetBillingAccount()),
		ProvisionedResourcesParent: dcl.StringOrNil(p.GetProvisionedResourcesParent()),
		KmsSettings:                ProtoToAssuredworkloadsWorkloadKmsSettings(p.GetKmsSettings()),
		Organization:               dcl.StringOrNil(p.GetOrganization()),
		Location:                   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsWorkloadResources(r))
	}
	for _, r := range p.GetResourceSettings() {
		obj.ResourceSettings = append(obj.ResourceSettings, *ProtoToAssuredworkloadsWorkloadResourceSettings(r))
	}
	return obj
}

// WorkloadResourcesResourceTypeEnumToProto converts a WorkloadResourcesResourceTypeEnum enum to its proto representation.
func AssuredworkloadsWorkloadResourcesResourceTypeEnumToProto(e *assuredworkloads.WorkloadResourcesResourceTypeEnum) assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum_value["WorkloadResourcesResourceTypeEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum(0)
}

// WorkloadComplianceRegimeEnumToProto converts a WorkloadComplianceRegimeEnum enum to its proto representation.
func AssuredworkloadsWorkloadComplianceRegimeEnumToProto(e *assuredworkloads.WorkloadComplianceRegimeEnum) assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum_value["WorkloadComplianceRegimeEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum(0)
}

// WorkloadResourceSettingsResourceTypeEnumToProto converts a WorkloadResourceSettingsResourceTypeEnum enum to its proto representation.
func AssuredworkloadsWorkloadResourceSettingsResourceTypeEnumToProto(e *assuredworkloads.WorkloadResourceSettingsResourceTypeEnum) assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum_value["WorkloadResourceSettingsResourceTypeEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(0)
}

// WorkloadResourcesToProto converts a WorkloadResources object to its proto representation.
func AssuredworkloadsWorkloadResourcesToProto(o *assuredworkloads.WorkloadResources) *assuredworkloadspb.AssuredworkloadsWorkloadResources {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadResources{}
	p.SetResourceId(dcl.ValueOrEmptyInt64(o.ResourceId))
	p.SetResourceType(AssuredworkloadsWorkloadResourcesResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings object to its proto representation.
func AssuredworkloadsWorkloadKmsSettingsToProto(o *assuredworkloads.WorkloadKmsSettings) *assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings{}
	p.SetNextRotationTime(dcl.ValueOrEmptyString(o.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(o.RotationPeriod))
	return p
}

// WorkloadResourceSettingsToProto converts a WorkloadResourceSettings object to its proto representation.
func AssuredworkloadsWorkloadResourceSettingsToProto(o *assuredworkloads.WorkloadResourceSettings) *assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings{}
	p.SetResourceId(dcl.ValueOrEmptyString(o.ResourceId))
	p.SetResourceType(AssuredworkloadsWorkloadResourceSettingsResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *assuredworkloads.Workload) *assuredworkloadspb.AssuredworkloadsWorkload {
	p := &assuredworkloadspb.AssuredworkloadsWorkload{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetComplianceRegime(AssuredworkloadsWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	p.SetProvisionedResourcesParent(dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent))
	p.SetKmsSettings(AssuredworkloadsWorkloadKmsSettingsToProto(resource.KmsSettings))
	p.SetOrganization(dcl.ValueOrEmptyString(resource.Organization))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sResources := make([]*assuredworkloadspb.AssuredworkloadsWorkloadResources, len(resource.Resources))
	for i, r := range resource.Resources {
		sResources[i] = AssuredworkloadsWorkloadResourcesToProto(&r)
	}
	p.SetResources(sResources)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sResourceSettings := make([]*assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings, len(resource.ResourceSettings))
	for i, r := range resource.ResourceSettings {
		sResourceSettings[i] = AssuredworkloadsWorkloadResourceSettingsToProto(&r)
	}
	p.SetResourceSettings(sResourceSettings)

	return p
}

// applyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *assuredworkloads.Client, request *assuredworkloadspb.ApplyAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.AssuredworkloadsWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// applyAssuredworkloadsWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.ApplyAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.AssuredworkloadsWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.DeleteAssuredworkloadsWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListAssuredworkloadsWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.ListAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.ListAssuredworkloadsWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, request.GetOrganization(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*assuredworkloadspb.AssuredworkloadsWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	p := &assuredworkloadspb.ListAssuredworkloadsWorkloadResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*assuredworkloads.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return assuredworkloads.NewClient(conf), nil
}
