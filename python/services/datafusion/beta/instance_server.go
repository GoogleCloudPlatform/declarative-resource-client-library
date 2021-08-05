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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/datafusion/beta/datafusion_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceTypeEnum converts a InstanceTypeEnum enum from its proto representation.
func ProtoToDatafusionBetaInstanceTypeEnum(e betapb.DatafusionBetaInstanceTypeEnum) *beta.InstanceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DatafusionBetaInstanceTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceTypeEnum(n[len("DatafusionBetaInstanceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToDatafusionBetaInstanceStateEnum(e betapb.DatafusionBetaInstanceStateEnum) *beta.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DatafusionBetaInstanceStateEnum_name[int32(e)]; ok {
		e := beta.InstanceStateEnum(n[len("DatafusionBetaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkConfig converts a InstanceNetworkConfig resource from its proto representation.
func ProtoToDatafusionBetaInstanceNetworkConfig(p *betapb.DatafusionBetaInstanceNetworkConfig) *beta.InstanceNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkConfig{
		Network:      dcl.StringOrNil(p.Network),
		IPAllocation: dcl.StringOrNil(p.IpAllocation),
	}
	return obj
}

// ProtoToInstanceAvailableVersion converts a InstanceAvailableVersion resource from its proto representation.
func ProtoToDatafusionBetaInstanceAvailableVersion(p *betapb.DatafusionBetaInstanceAvailableVersion) *beta.InstanceAvailableVersion {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceAvailableVersion{
		VersionNumber:  dcl.StringOrNil(p.VersionNumber),
		DefaultVersion: dcl.Bool(p.DefaultVersion),
	}
	for _, r := range p.GetAvailableFeatures() {
		obj.AvailableFeatures = append(obj.AvailableFeatures, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.DatafusionBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		Name:                        dcl.StringOrNil(p.Name),
		Description:                 dcl.StringOrNil(p.Description),
		Type:                        ProtoToDatafusionBetaInstanceTypeEnum(p.GetType()),
		EnableStackdriverLogging:    dcl.Bool(p.EnableStackdriverLogging),
		EnableStackdriverMonitoring: dcl.Bool(p.EnableStackdriverMonitoring),
		PrivateInstance:             dcl.Bool(p.PrivateInstance),
		NetworkConfig:               ProtoToDatafusionBetaInstanceNetworkConfig(p.GetNetworkConfig()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		State:                       ProtoToDatafusionBetaInstanceStateEnum(p.GetState()),
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
		obj.AvailableVersion = append(obj.AvailableVersion, *ProtoToDatafusionBetaInstanceAvailableVersion(r))
	}
	return obj
}

// InstanceTypeEnumToProto converts a InstanceTypeEnum enum to its proto representation.
func DatafusionBetaInstanceTypeEnumToProto(e *beta.InstanceTypeEnum) betapb.DatafusionBetaInstanceTypeEnum {
	if e == nil {
		return betapb.DatafusionBetaInstanceTypeEnum(0)
	}
	if v, ok := betapb.DatafusionBetaInstanceTypeEnum_value["InstanceTypeEnum"+string(*e)]; ok {
		return betapb.DatafusionBetaInstanceTypeEnum(v)
	}
	return betapb.DatafusionBetaInstanceTypeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func DatafusionBetaInstanceStateEnumToProto(e *beta.InstanceStateEnum) betapb.DatafusionBetaInstanceStateEnum {
	if e == nil {
		return betapb.DatafusionBetaInstanceStateEnum(0)
	}
	if v, ok := betapb.DatafusionBetaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return betapb.DatafusionBetaInstanceStateEnum(v)
	}
	return betapb.DatafusionBetaInstanceStateEnum(0)
}

// InstanceNetworkConfigToProto converts a InstanceNetworkConfig resource to its proto representation.
func DatafusionBetaInstanceNetworkConfigToProto(o *beta.InstanceNetworkConfig) *betapb.DatafusionBetaInstanceNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DatafusionBetaInstanceNetworkConfig{
		Network:      dcl.ValueOrEmptyString(o.Network),
		IpAllocation: dcl.ValueOrEmptyString(o.IPAllocation),
	}
	return p
}

// InstanceAvailableVersionToProto converts a InstanceAvailableVersion resource to its proto representation.
func DatafusionBetaInstanceAvailableVersionToProto(o *beta.InstanceAvailableVersion) *betapb.DatafusionBetaInstanceAvailableVersion {
	if o == nil {
		return nil
	}
	p := &betapb.DatafusionBetaInstanceAvailableVersion{
		VersionNumber:  dcl.ValueOrEmptyString(o.VersionNumber),
		DefaultVersion: dcl.ValueOrEmptyBool(o.DefaultVersion),
	}
	for _, r := range o.AvailableFeatures {
		p.AvailableFeatures = append(p.AvailableFeatures, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.DatafusionBetaInstance {
	p := &betapb.DatafusionBetaInstance{
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		Type:                        DatafusionBetaInstanceTypeEnumToProto(resource.Type),
		EnableStackdriverLogging:    dcl.ValueOrEmptyBool(resource.EnableStackdriverLogging),
		EnableStackdriverMonitoring: dcl.ValueOrEmptyBool(resource.EnableStackdriverMonitoring),
		PrivateInstance:             dcl.ValueOrEmptyBool(resource.PrivateInstance),
		NetworkConfig:               DatafusionBetaInstanceNetworkConfigToProto(resource.NetworkConfig),
		CreateTime:                  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:                  dcl.ValueOrEmptyString(resource.UpdateTime),
		State:                       DatafusionBetaInstanceStateEnumToProto(resource.State),
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
		p.AvailableVersion = append(p.AvailableVersion, DatafusionBetaInstanceAvailableVersionToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyDatafusionBetaInstanceRequest) (*betapb.DatafusionBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyDatafusionBetaInstance(ctx context.Context, request *betapb.ApplyDatafusionBetaInstanceRequest) (*betapb.DatafusionBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteDatafusionBetaInstance(ctx context.Context, request *betapb.DeleteDatafusionBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListDatafusionBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListDatafusionBetaInstance(ctx context.Context, request *betapb.ListDatafusionBetaInstanceRequest) (*betapb.ListDatafusionBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, ProtoToInstance(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DatafusionBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListDatafusionBetaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
