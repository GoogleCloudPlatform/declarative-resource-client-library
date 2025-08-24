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
	cloudbuildpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/cloudbuild_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild"
)

// WorkerPoolServer implements the gRPC interface for WorkerPool.
type WorkerPoolServer struct{}

// ProtoToWorkerPoolStateEnum converts a WorkerPoolStateEnum enum from its proto representation.
func ProtoToCloudbuildWorkerPoolStateEnum(e cloudbuildpb.CloudbuildWorkerPoolStateEnum) *cloudbuild.WorkerPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudbuildpb.CloudbuildWorkerPoolStateEnum_name[int32(e)]; ok {
		e := cloudbuild.WorkerPoolStateEnum(n[len("CloudbuildWorkerPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum converts a WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum enum from its proto representation.
func ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(e cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum) *cloudbuild.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum_name[int32(e)]; ok {
		e := cloudbuild.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(n[len("CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolPrivatePoolV1Config converts a WorkerPoolPrivatePoolV1Config object from its proto representation.
func ProtoToCloudbuildWorkerPoolPrivatePoolV1Config(p *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1Config) *cloudbuild.WorkerPoolPrivatePoolV1Config {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolPrivatePoolV1Config{
		WorkerConfig:          ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig:         ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfig(p.GetNetworkConfig()),
		PrivateServiceConnect: ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnect(p.GetPrivateServiceConnect()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigWorkerConfig converts a WorkerPoolPrivatePoolV1ConfigWorkerConfig object from its proto representation.
func ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfig(p *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfig) *cloudbuild.WorkerPoolPrivatePoolV1ConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolPrivatePoolV1ConfigWorkerConfig{
		MachineType:                dcl.StringOrNil(p.GetMachineType()),
		DiskSizeGb:                 dcl.Int64OrNil(p.GetDiskSizeGb()),
		EnableNestedVirtualization: dcl.Bool(p.GetEnableNestedVirtualization()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigNetworkConfig converts a WorkerPoolPrivatePoolV1ConfigNetworkConfig object from its proto representation.
func ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfig(p *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfig) *cloudbuild.WorkerPoolPrivatePoolV1ConfigNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolPrivatePoolV1ConfigNetworkConfig{
		PeeredNetwork:        dcl.StringOrNil(p.GetPeeredNetwork()),
		PeeredNetworkIPRange: dcl.StringOrNil(p.GetPeeredNetworkIpRange()),
		EgressOption:         ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(p.GetEgressOption()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigPrivateServiceConnect converts a WorkerPoolPrivatePoolV1ConfigPrivateServiceConnect object from its proto representation.
func ProtoToCloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnect(p *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnect) *cloudbuild.WorkerPoolPrivatePoolV1ConfigPrivateServiceConnect {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolPrivatePoolV1ConfigPrivateServiceConnect{
		NetworkAttachment:       dcl.StringOrNil(p.GetNetworkAttachment()),
		PublicIPAddressDisabled: dcl.Bool(p.GetPublicIpAddressDisabled()),
		RouteAllTraffic:         dcl.Bool(p.GetRouteAllTraffic()),
	}
	return obj
}

// ProtoToWorkerPoolWorkerConfig converts a WorkerPoolWorkerConfig object from its proto representation.
func ProtoToCloudbuildWorkerPoolWorkerConfig(p *cloudbuildpb.CloudbuildWorkerPoolWorkerConfig) *cloudbuild.WorkerPoolWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolWorkerConfig{
		MachineType:                dcl.StringOrNil(p.GetMachineType()),
		DiskSizeGb:                 dcl.Int64OrNil(p.GetDiskSizeGb()),
		EnableNestedVirtualization: dcl.Bool(p.GetEnableNestedVirtualization()),
		NoExternalIP:               dcl.Bool(p.GetNoExternalIp()),
	}
	return obj
}

// ProtoToWorkerPoolNetworkConfig converts a WorkerPoolNetworkConfig object from its proto representation.
func ProtoToCloudbuildWorkerPoolNetworkConfig(p *cloudbuildpb.CloudbuildWorkerPoolNetworkConfig) *cloudbuild.WorkerPoolNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolNetworkConfig{
		PeeredNetwork:        dcl.StringOrNil(p.GetPeeredNetwork()),
		PeeredNetworkIPRange: dcl.StringOrNil(p.GetPeeredNetworkIpRange()),
	}
	return obj
}

// ProtoToWorkerPoolPrivateServiceConnect converts a WorkerPoolPrivateServiceConnect object from its proto representation.
func ProtoToCloudbuildWorkerPoolPrivateServiceConnect(p *cloudbuildpb.CloudbuildWorkerPoolPrivateServiceConnect) *cloudbuild.WorkerPoolPrivateServiceConnect {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.WorkerPoolPrivateServiceConnect{
		NetworkAttachment: dcl.StringOrNil(p.GetNetworkAttachment()),
		RouteAllTraffic:   dcl.Bool(p.GetRouteAllTraffic()),
	}
	return obj
}

// ProtoToWorkerPool converts a WorkerPool resource from its proto representation.
func ProtoToWorkerPool(p *cloudbuildpb.CloudbuildWorkerPool) *cloudbuild.WorkerPool {
	obj := &cloudbuild.WorkerPool{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Uid:                   dcl.StringOrNil(p.GetUid()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:            dcl.StringOrNil(p.GetDeleteTime()),
		State:                 ProtoToCloudbuildWorkerPoolStateEnum(p.GetState()),
		PrivatePoolV1Config:   ProtoToCloudbuildWorkerPoolPrivatePoolV1Config(p.GetPrivatePoolV1Config()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		WorkerConfig:          ProtoToCloudbuildWorkerPoolWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig:         ProtoToCloudbuildWorkerPoolNetworkConfig(p.GetNetworkConfig()),
		PrivateServiceConnect: ProtoToCloudbuildWorkerPoolPrivateServiceConnect(p.GetPrivateServiceConnect()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkerPoolStateEnumToProto converts a WorkerPoolStateEnum enum to its proto representation.
func CloudbuildWorkerPoolStateEnumToProto(e *cloudbuild.WorkerPoolStateEnum) cloudbuildpb.CloudbuildWorkerPoolStateEnum {
	if e == nil {
		return cloudbuildpb.CloudbuildWorkerPoolStateEnum(0)
	}
	if v, ok := cloudbuildpb.CloudbuildWorkerPoolStateEnum_value["WorkerPoolStateEnum"+string(*e)]; ok {
		return cloudbuildpb.CloudbuildWorkerPoolStateEnum(v)
	}
	return cloudbuildpb.CloudbuildWorkerPoolStateEnum(0)
}

// WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto converts a WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum enum to its proto representation.
func CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto(e *cloudbuild.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum) cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum {
	if e == nil {
		return cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(0)
	}
	if v, ok := cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum_value["WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"+string(*e)]; ok {
		return cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(v)
	}
	return cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(0)
}

// WorkerPoolPrivatePoolV1ConfigToProto converts a WorkerPoolPrivatePoolV1Config object to its proto representation.
func CloudbuildWorkerPoolPrivatePoolV1ConfigToProto(o *cloudbuild.WorkerPoolPrivatePoolV1Config) *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1Config {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1Config{}
	p.SetWorkerConfig(CloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetNetworkConfig(CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigToProto(o.NetworkConfig))
	p.SetPrivateServiceConnect(CloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnectToProto(o.PrivateServiceConnect))
	return p
}

// WorkerPoolPrivatePoolV1ConfigWorkerConfigToProto converts a WorkerPoolPrivatePoolV1ConfigWorkerConfig object to its proto representation.
func CloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfigToProto(o *cloudbuild.WorkerPoolPrivatePoolV1ConfigWorkerConfig) *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigWorkerConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetEnableNestedVirtualization(dcl.ValueOrEmptyBool(o.EnableNestedVirtualization))
	return p
}

// WorkerPoolPrivatePoolV1ConfigNetworkConfigToProto converts a WorkerPoolPrivatePoolV1ConfigNetworkConfig object to its proto representation.
func CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigToProto(o *cloudbuild.WorkerPoolPrivatePoolV1ConfigNetworkConfig) *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfig{}
	p.SetPeeredNetwork(dcl.ValueOrEmptyString(o.PeeredNetwork))
	p.SetPeeredNetworkIpRange(dcl.ValueOrEmptyString(o.PeeredNetworkIPRange))
	p.SetEgressOption(CloudbuildWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto(o.EgressOption))
	return p
}

// WorkerPoolPrivatePoolV1ConfigPrivateServiceConnectToProto converts a WorkerPoolPrivatePoolV1ConfigPrivateServiceConnect object to its proto representation.
func CloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnectToProto(o *cloudbuild.WorkerPoolPrivatePoolV1ConfigPrivateServiceConnect) *cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnect {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolPrivatePoolV1ConfigPrivateServiceConnect{}
	p.SetNetworkAttachment(dcl.ValueOrEmptyString(o.NetworkAttachment))
	p.SetPublicIpAddressDisabled(dcl.ValueOrEmptyBool(o.PublicIPAddressDisabled))
	p.SetRouteAllTraffic(dcl.ValueOrEmptyBool(o.RouteAllTraffic))
	return p
}

// WorkerPoolWorkerConfigToProto converts a WorkerPoolWorkerConfig object to its proto representation.
func CloudbuildWorkerPoolWorkerConfigToProto(o *cloudbuild.WorkerPoolWorkerConfig) *cloudbuildpb.CloudbuildWorkerPoolWorkerConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolWorkerConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetEnableNestedVirtualization(dcl.ValueOrEmptyBool(o.EnableNestedVirtualization))
	p.SetNoExternalIp(dcl.ValueOrEmptyBool(o.NoExternalIP))
	return p
}

// WorkerPoolNetworkConfigToProto converts a WorkerPoolNetworkConfig object to its proto representation.
func CloudbuildWorkerPoolNetworkConfigToProto(o *cloudbuild.WorkerPoolNetworkConfig) *cloudbuildpb.CloudbuildWorkerPoolNetworkConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolNetworkConfig{}
	p.SetPeeredNetwork(dcl.ValueOrEmptyString(o.PeeredNetwork))
	p.SetPeeredNetworkIpRange(dcl.ValueOrEmptyString(o.PeeredNetworkIPRange))
	return p
}

// WorkerPoolPrivateServiceConnectToProto converts a WorkerPoolPrivateServiceConnect object to its proto representation.
func CloudbuildWorkerPoolPrivateServiceConnectToProto(o *cloudbuild.WorkerPoolPrivateServiceConnect) *cloudbuildpb.CloudbuildWorkerPoolPrivateServiceConnect {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildWorkerPoolPrivateServiceConnect{}
	p.SetNetworkAttachment(dcl.ValueOrEmptyString(o.NetworkAttachment))
	p.SetRouteAllTraffic(dcl.ValueOrEmptyBool(o.RouteAllTraffic))
	return p
}

// WorkerPoolToProto converts a WorkerPool resource to its proto representation.
func WorkerPoolToProto(resource *cloudbuild.WorkerPool) *cloudbuildpb.CloudbuildWorkerPool {
	p := &cloudbuildpb.CloudbuildWorkerPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetState(CloudbuildWorkerPoolStateEnumToProto(resource.State))
	p.SetPrivatePoolV1Config(CloudbuildWorkerPoolPrivatePoolV1ConfigToProto(resource.PrivatePoolV1Config))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkerConfig(CloudbuildWorkerPoolWorkerConfigToProto(resource.WorkerConfig))
	p.SetNetworkConfig(CloudbuildWorkerPoolNetworkConfigToProto(resource.NetworkConfig))
	p.SetPrivateServiceConnect(CloudbuildWorkerPoolPrivateServiceConnectToProto(resource.PrivateServiceConnect))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) applyWorkerPool(ctx context.Context, c *cloudbuild.Client, request *cloudbuildpb.ApplyCloudbuildWorkerPoolRequest) (*cloudbuildpb.CloudbuildWorkerPool, error) {
	p := ProtoToWorkerPool(request.GetResource())
	res, err := c.ApplyWorkerPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkerPoolToProto(res)
	return r, nil
}

// applyCloudbuildWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) ApplyCloudbuildWorkerPool(ctx context.Context, request *cloudbuildpb.ApplyCloudbuildWorkerPoolRequest) (*cloudbuildpb.CloudbuildWorkerPool, error) {
	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkerPool(ctx, cl, request)
}

// DeleteWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Delete() method.
func (s *WorkerPoolServer) DeleteCloudbuildWorkerPool(ctx context.Context, request *cloudbuildpb.DeleteCloudbuildWorkerPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))

}

// ListCloudbuildWorkerPool handles the gRPC request by passing it to the underlying WorkerPoolList() method.
func (s *WorkerPoolServer) ListCloudbuildWorkerPool(ctx context.Context, request *cloudbuildpb.ListCloudbuildWorkerPoolRequest) (*cloudbuildpb.ListCloudbuildWorkerPoolResponse, error) {
	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkerPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*cloudbuildpb.CloudbuildWorkerPool
	for _, r := range resources.Items {
		rp := WorkerPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudbuildpb.ListCloudbuildWorkerPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkerPool(ctx context.Context, service_account_file string) (*cloudbuild.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudbuild.NewClient(conf), nil
}
