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
	assuredworkloadspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/assuredworkloads_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads"
)

// Server implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourcesResourceTypeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum) *assuredworkloads.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadResourcesResourceTypeEnum(n[len("WorkloadResourcesResourceTypeEnum"):])
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
		e := assuredworkloads.WorkloadComplianceRegimeEnum(n[len("WorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources resource from its proto representation.
func ProtoToAssuredworkloadsWorkloadResources(p *assuredworkloadspb.AssuredworkloadsWorkloadResources) *assuredworkloads.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.ResourceId),
		ResourceType: ProtoToAssuredworkloadsWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings resource from its proto representation.
func ProtoToAssuredworkloadsWorkloadKmsSettings(p *assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings) *assuredworkloads.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.RotationPeriod),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *assuredworkloadspb.AssuredworkloadsWorkload) *assuredworkloads.Workload {
	obj := &assuredworkloads.Workload{
		Name:                       dcl.StringOrNil(p.Name),
		DisplayName:                dcl.StringOrNil(p.DisplayName),
		ComplianceRegime:           ProtoToAssuredworkloadsWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                 dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:             dcl.StringOrNil(p.BillingAccount),
		ProvisionedResourcesParent: dcl.StringOrNil(p.ProvisionedResourcesParent),
		KmsSettings:                ProtoToAssuredworkloadsWorkloadKmsSettings(p.GetKmsSettings()),
		Organization:               dcl.StringOrNil(p.Organization),
		Location:                   dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsWorkloadResources(r))
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

// WorkloadResourcesToProto converts a WorkloadResources resource to its proto representation.
func AssuredworkloadsWorkloadResourcesToProto(o *assuredworkloads.WorkloadResources) *assuredworkloadspb.AssuredworkloadsWorkloadResources {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadResources{
		ResourceId:   dcl.ValueOrEmptyInt64(o.ResourceId),
		ResourceType: AssuredworkloadsWorkloadResourcesResourceTypeEnumToProto(o.ResourceType),
	}
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings resource to its proto representation.
func AssuredworkloadsWorkloadKmsSettingsToProto(o *assuredworkloads.WorkloadKmsSettings) *assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings{
		NextRotationTime: dcl.ValueOrEmptyString(o.NextRotationTime),
		RotationPeriod:   dcl.ValueOrEmptyString(o.RotationPeriod),
	}
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *assuredworkloads.Workload) *assuredworkloadspb.AssuredworkloadsWorkload {
	p := &assuredworkloadspb.AssuredworkloadsWorkload{
		Name:                       dcl.ValueOrEmptyString(resource.Name),
		DisplayName:                dcl.ValueOrEmptyString(resource.DisplayName),
		ComplianceRegime:           AssuredworkloadsWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime),
		CreateTime:                 dcl.ValueOrEmptyString(resource.CreateTime),
		BillingAccount:             dcl.ValueOrEmptyString(resource.BillingAccount),
		ProvisionedResourcesParent: dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent),
		KmsSettings:                AssuredworkloadsWorkloadKmsSettingsToProto(resource.KmsSettings),
		Organization:               dcl.ValueOrEmptyString(resource.Organization),
		Location:                   dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Resources {
		p.Resources = append(p.Resources, AssuredworkloadsWorkloadResourcesToProto(&r))
	}

	return p
}

// ApplyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *assuredworkloads.Client, request *assuredworkloadspb.ApplyAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.AssuredworkloadsWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// ApplyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.ApplyAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.AssuredworkloadsWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.DeleteAssuredworkloadsWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.ListAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.ListAssuredworkloadsWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, request.Organization, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*assuredworkloadspb.AssuredworkloadsWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	return &assuredworkloadspb.ListAssuredworkloadsWorkloadResponse{Items: protos}, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*assuredworkloads.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return assuredworkloads.NewClient(conf), nil
}
