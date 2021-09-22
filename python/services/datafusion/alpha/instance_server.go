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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/datafusion/alpha/datafusion_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/alpha"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceTypeEnum converts a InstanceTypeEnum enum from its proto representation.
func ProtoToDatafusionAlphaInstanceTypeEnum(e alphapb.DatafusionAlphaInstanceTypeEnum) *alpha.InstanceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DatafusionAlphaInstanceTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceTypeEnum(n[len("DatafusionAlphaInstanceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToDatafusionAlphaInstanceStateEnum(e alphapb.DatafusionAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DatafusionAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("DatafusionAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkConfig converts a InstanceNetworkConfig resource from its proto representation.
func ProtoToDatafusionAlphaInstanceNetworkConfig(p *alphapb.DatafusionAlphaInstanceNetworkConfig) *alpha.InstanceNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworkConfig{
		Network:      dcl.StringOrNil(p.Network),
		IPAllocation: dcl.StringOrNil(p.IpAllocation),
	}
	return obj
}

// ProtoToInstanceAvailableVersion converts a InstanceAvailableVersion resource from its proto representation.
func ProtoToDatafusionAlphaInstanceAvailableVersion(p *alphapb.DatafusionAlphaInstanceAvailableVersion) *alpha.InstanceAvailableVersion {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceAvailableVersion{
		VersionNumber:  dcl.StringOrNil(p.VersionNumber),
		DefaultVersion: dcl.Bool(p.DefaultVersion),
	}
	for _, r := range p.GetAvailableFeatures() {
		obj.AvailableFeatures = append(obj.AvailableFeatures, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.DatafusionAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:                        dcl.StringOrNil(p.Name),
		Description:                 dcl.StringOrNil(p.Description),
		Type:                        ProtoToDatafusionAlphaInstanceTypeEnum(p.GetType()),
		EnableStackdriverLogging:    dcl.Bool(p.EnableStackdriverLogging),
		EnableStackdriverMonitoring: dcl.Bool(p.EnableStackdriverMonitoring),
		PrivateInstance:             dcl.Bool(p.PrivateInstance),
		NetworkConfig:               ProtoToDatafusionAlphaInstanceNetworkConfig(p.GetNetworkConfig()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		State:                       ProtoToDatafusionAlphaInstanceStateEnum(p.GetState()),
		StateMessage:                dcl.StringOrNil(p.StateMessage),
		ServiceEndpoint:             dcl.StringOrNil(p.ServiceEndpoint),
		Zone:                        dcl.StringOrNil(p.Zone),
		Version:                     dcl.StringOrNil(p.Version),
		DisplayName:                 dcl.StringOrNil(p.DisplayName),
		ApiEndpoint:                 dcl.StringOrNil(p.ApiEndpoint),
		GcsBucket:                   dcl.StringOrNil(p.GcsBucket),
		P4ServiceAccount:            dcl.StringOrNil(p.P4ServiceAccount),
		TenantProjectId:             dcl.StringOrNil(p.TenantProjectId),
		DataprocServiceAccount:      dcl.StringOrNil(p.DataprocServiceAccount),
		Project:                     dcl.StringOrNil(p.Project),
		Location:                    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetAvailableVersion() {
		obj.AvailableVersion = append(obj.AvailableVersion, *ProtoToDatafusionAlphaInstanceAvailableVersion(r))
	}
	return obj
}

// InstanceTypeEnumToProto converts a InstanceTypeEnum enum to its proto representation.
func DatafusionAlphaInstanceTypeEnumToProto(e *alpha.InstanceTypeEnum) alphapb.DatafusionAlphaInstanceTypeEnum {
	if e == nil {
		return alphapb.DatafusionAlphaInstanceTypeEnum(0)
	}
	if v, ok := alphapb.DatafusionAlphaInstanceTypeEnum_value["InstanceTypeEnum"+string(*e)]; ok {
		return alphapb.DatafusionAlphaInstanceTypeEnum(v)
	}
	return alphapb.DatafusionAlphaInstanceTypeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func DatafusionAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.DatafusionAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.DatafusionAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.DatafusionAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.DatafusionAlphaInstanceStateEnum(v)
	}
	return alphapb.DatafusionAlphaInstanceStateEnum(0)
}

// InstanceNetworkConfigToProto converts a InstanceNetworkConfig resource to its proto representation.
func DatafusionAlphaInstanceNetworkConfigToProto(o *alpha.InstanceNetworkConfig) *alphapb.DatafusionAlphaInstanceNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DatafusionAlphaInstanceNetworkConfig{
		Network:      dcl.ValueOrEmptyString(o.Network),
		IpAllocation: dcl.ValueOrEmptyString(o.IPAllocation),
	}
	return p
}

// InstanceAvailableVersionToProto converts a InstanceAvailableVersion resource to its proto representation.
func DatafusionAlphaInstanceAvailableVersionToProto(o *alpha.InstanceAvailableVersion) *alphapb.DatafusionAlphaInstanceAvailableVersion {
	if o == nil {
		return nil
	}
	p := &alphapb.DatafusionAlphaInstanceAvailableVersion{
		VersionNumber:  dcl.ValueOrEmptyString(o.VersionNumber),
		DefaultVersion: dcl.ValueOrEmptyBool(o.DefaultVersion),
	}
	for _, r := range o.AvailableFeatures {
		p.AvailableFeatures = append(p.AvailableFeatures, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.DatafusionAlphaInstance {
	p := &alphapb.DatafusionAlphaInstance{
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		Type:                        DatafusionAlphaInstanceTypeEnumToProto(resource.Type),
		EnableStackdriverLogging:    dcl.ValueOrEmptyBool(resource.EnableStackdriverLogging),
		EnableStackdriverMonitoring: dcl.ValueOrEmptyBool(resource.EnableStackdriverMonitoring),
		PrivateInstance:             dcl.ValueOrEmptyBool(resource.PrivateInstance),
		NetworkConfig:               DatafusionAlphaInstanceNetworkConfigToProto(resource.NetworkConfig),
		CreateTime:                  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:                  dcl.ValueOrEmptyString(resource.UpdateTime),
		State:                       DatafusionAlphaInstanceStateEnumToProto(resource.State),
		StateMessage:                dcl.ValueOrEmptyString(resource.StateMessage),
		ServiceEndpoint:             dcl.ValueOrEmptyString(resource.ServiceEndpoint),
		Zone:                        dcl.ValueOrEmptyString(resource.Zone),
		Version:                     dcl.ValueOrEmptyString(resource.Version),
		DisplayName:                 dcl.ValueOrEmptyString(resource.DisplayName),
		ApiEndpoint:                 dcl.ValueOrEmptyString(resource.ApiEndpoint),
		GcsBucket:                   dcl.ValueOrEmptyString(resource.GcsBucket),
		P4ServiceAccount:            dcl.ValueOrEmptyString(resource.P4ServiceAccount),
		TenantProjectId:             dcl.ValueOrEmptyString(resource.TenantProjectId),
		DataprocServiceAccount:      dcl.ValueOrEmptyString(resource.DataprocServiceAccount),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
		Location:                    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.AvailableVersion {
		p.AvailableVersion = append(p.AvailableVersion, DatafusionAlphaInstanceAvailableVersionToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDatafusionAlphaInstanceRequest) (*alphapb.DatafusionAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyDatafusionAlphaInstance(ctx context.Context, request *alphapb.ApplyDatafusionAlphaInstanceRequest) (*alphapb.DatafusionAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteDatafusionAlphaInstance(ctx context.Context, request *alphapb.DeleteDatafusionAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListDatafusionAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListDatafusionAlphaInstance(ctx context.Context, request *alphapb.ListDatafusionAlphaInstanceRequest) (*alphapb.ListDatafusionAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DatafusionAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListDatafusionAlphaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
