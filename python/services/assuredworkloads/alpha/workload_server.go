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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/alpha/assuredworkloads_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/alpha"
)

// Server implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(e alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum) *alpha.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.WorkloadResourcesResourceTypeEnum(n[len("AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadComplianceRegimeEnum converts a WorkloadComplianceRegimeEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadComplianceRegimeEnum(e alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum) *alpha.WorkloadComplianceRegimeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum_name[int32(e)]; ok {
		e := alpha.WorkloadComplianceRegimeEnum(n[len("AssuredworkloadsAlphaWorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResourceSettingsResourceTypeEnum converts a WorkloadResourceSettingsResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(e alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum) *alpha.WorkloadResourceSettingsResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.WorkloadResourceSettingsResourceTypeEnum(n[len("AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources resource from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResources(p *alphapb.AssuredworkloadsAlphaWorkloadResources) *alpha.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.ResourceId),
		ResourceType: ProtoToAssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings resource from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadKmsSettings(p *alphapb.AssuredworkloadsAlphaWorkloadKmsSettings) *alpha.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.RotationPeriod),
	}
	return obj
}

// ProtoToWorkloadResourceSettings converts a WorkloadResourceSettings resource from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResourceSettings(p *alphapb.AssuredworkloadsAlphaWorkloadResourceSettings) *alpha.WorkloadResourceSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadResourceSettings{
		ResourceId:   dcl.StringOrNil(p.ResourceId),
		ResourceType: ProtoToAssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *alphapb.AssuredworkloadsAlphaWorkload) *alpha.Workload {
	obj := &alpha.Workload{
		Name:                       dcl.StringOrNil(p.Name),
		DisplayName:                dcl.StringOrNil(p.DisplayName),
		ComplianceRegime:           ProtoToAssuredworkloadsAlphaWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                 dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:             dcl.StringOrNil(p.BillingAccount),
		ProvisionedResourcesParent: dcl.StringOrNil(p.ProvisionedResourcesParent),
		KmsSettings:                ProtoToAssuredworkloadsAlphaWorkloadKmsSettings(p.GetKmsSettings()),
		Organization:               dcl.StringOrNil(p.Organization),
		Location:                   dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsAlphaWorkloadResources(r))
	}
	for _, r := range p.GetResourceSettings() {
		obj.ResourceSettings = append(obj.ResourceSettings, *ProtoToAssuredworkloadsAlphaWorkloadResourceSettings(r))
	}
	return obj
}

// WorkloadResourcesResourceTypeEnumToProto converts a WorkloadResourcesResourceTypeEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnumToProto(e *alpha.WorkloadResourcesResourceTypeEnum) alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum_value["WorkloadResourcesResourceTypeEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(0)
}

// WorkloadComplianceRegimeEnumToProto converts a WorkloadComplianceRegimeEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadComplianceRegimeEnumToProto(e *alpha.WorkloadComplianceRegimeEnum) alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum_value["WorkloadComplianceRegimeEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum(0)
}

// WorkloadResourceSettingsResourceTypeEnumToProto converts a WorkloadResourceSettingsResourceTypeEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnumToProto(e *alpha.WorkloadResourceSettingsResourceTypeEnum) alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum_value["WorkloadResourceSettingsResourceTypeEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(0)
}

// WorkloadResourcesToProto converts a WorkloadResources resource to its proto representation.
func AssuredworkloadsAlphaWorkloadResourcesToProto(o *alpha.WorkloadResources) *alphapb.AssuredworkloadsAlphaWorkloadResources {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadResources{
		ResourceId:   dcl.ValueOrEmptyInt64(o.ResourceId),
		ResourceType: AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnumToProto(o.ResourceType),
	}
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings resource to its proto representation.
func AssuredworkloadsAlphaWorkloadKmsSettingsToProto(o *alpha.WorkloadKmsSettings) *alphapb.AssuredworkloadsAlphaWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadKmsSettings{
		NextRotationTime: dcl.ValueOrEmptyString(o.NextRotationTime),
		RotationPeriod:   dcl.ValueOrEmptyString(o.RotationPeriod),
	}
	return p
}

// WorkloadResourceSettingsToProto converts a WorkloadResourceSettings resource to its proto representation.
func AssuredworkloadsAlphaWorkloadResourceSettingsToProto(o *alpha.WorkloadResourceSettings) *alphapb.AssuredworkloadsAlphaWorkloadResourceSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadResourceSettings{
		ResourceId:   dcl.ValueOrEmptyString(o.ResourceId),
		ResourceType: AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnumToProto(o.ResourceType),
	}
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *alpha.Workload) *alphapb.AssuredworkloadsAlphaWorkload {
	p := &alphapb.AssuredworkloadsAlphaWorkload{
		Name:                       dcl.ValueOrEmptyString(resource.Name),
		DisplayName:                dcl.ValueOrEmptyString(resource.DisplayName),
		ComplianceRegime:           AssuredworkloadsAlphaWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime),
		CreateTime:                 dcl.ValueOrEmptyString(resource.CreateTime),
		BillingAccount:             dcl.ValueOrEmptyString(resource.BillingAccount),
		ProvisionedResourcesParent: dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent),
		KmsSettings:                AssuredworkloadsAlphaWorkloadKmsSettingsToProto(resource.KmsSettings),
		Organization:               dcl.ValueOrEmptyString(resource.Organization),
		Location:                   dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Resources {
		p.Resources = append(p.Resources, AssuredworkloadsAlphaWorkloadResourcesToProto(&r))
	}
	for _, r := range resource.ResourceSettings {
		p.ResourceSettings = append(p.ResourceSettings, AssuredworkloadsAlphaWorkloadResourceSettingsToProto(&r))
	}

	return p
}

// ApplyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *alpha.Client, request *alphapb.ApplyAssuredworkloadsAlphaWorkloadRequest) (*alphapb.AssuredworkloadsAlphaWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// ApplyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsAlphaWorkload(ctx context.Context, request *alphapb.ApplyAssuredworkloadsAlphaWorkloadRequest) (*alphapb.AssuredworkloadsAlphaWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsAlphaWorkload(ctx context.Context, request *alphapb.DeleteAssuredworkloadsAlphaWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListAssuredworkloadsAlphaWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsAlphaWorkload(ctx context.Context, request *alphapb.ListAssuredworkloadsAlphaWorkloadRequest) (*alphapb.ListAssuredworkloadsAlphaWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, ProtoToWorkload(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.AssuredworkloadsAlphaWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListAssuredworkloadsAlphaWorkloadResponse{Items: protos}, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
