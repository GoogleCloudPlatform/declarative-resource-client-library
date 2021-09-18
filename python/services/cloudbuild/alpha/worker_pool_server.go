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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/alpha/cloudbuild_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/alpha"
)

// Server implements the gRPC interface for WorkerPool.
type WorkerPoolServer struct{}

// ProtoToWorkerPoolStateEnum converts a WorkerPoolStateEnum enum from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolStateEnum(e alphapb.CloudbuildAlphaWorkerPoolStateEnum) *alpha.WorkerPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudbuildAlphaWorkerPoolStateEnum_name[int32(e)]; ok {
		e := alpha.WorkerPoolStateEnum(n[len("CloudbuildAlphaWorkerPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolWorkerConfig converts a WorkerPoolWorkerConfig resource from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolWorkerConfig(p *alphapb.CloudbuildAlphaWorkerPoolWorkerConfig) *alpha.WorkerPoolWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolWorkerConfig{
		MachineType:  dcl.StringOrNil(p.MachineType),
		DiskSizeGb:   dcl.Int64OrNil(p.DiskSizeGb),
		NoExternalIP: dcl.Bool(p.NoExternalIp),
	}
	return obj
}

// ProtoToWorkerPoolNetworkConfig converts a WorkerPoolNetworkConfig resource from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolNetworkConfig(p *alphapb.CloudbuildAlphaWorkerPoolNetworkConfig) *alpha.WorkerPoolNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolNetworkConfig{
		PeeredNetwork: dcl.StringOrNil(p.PeeredNetwork),
	}
	return obj
}

// ProtoToWorkerPool converts a WorkerPool resource from its proto representation.
func ProtoToWorkerPool(p *alphapb.CloudbuildAlphaWorkerPool) *alpha.WorkerPool {
	obj := &alpha.WorkerPool{
		Name:          dcl.StringOrNil(p.Name),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		State:         ProtoToCloudbuildAlphaWorkerPoolStateEnum(p.GetState()),
		WorkerConfig:  ProtoToCloudbuildAlphaWorkerPoolWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig: ProtoToCloudbuildAlphaWorkerPoolNetworkConfig(p.GetNetworkConfig()),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	return obj
}

// WorkerPoolStateEnumToProto converts a WorkerPoolStateEnum enum to its proto representation.
func CloudbuildAlphaWorkerPoolStateEnumToProto(e *alpha.WorkerPoolStateEnum) alphapb.CloudbuildAlphaWorkerPoolStateEnum {
	if e == nil {
		return alphapb.CloudbuildAlphaWorkerPoolStateEnum(0)
	}
	if v, ok := alphapb.CloudbuildAlphaWorkerPoolStateEnum_value["WorkerPoolStateEnum"+string(*e)]; ok {
		return alphapb.CloudbuildAlphaWorkerPoolStateEnum(v)
	}
	return alphapb.CloudbuildAlphaWorkerPoolStateEnum(0)
}

// WorkerPoolWorkerConfigToProto converts a WorkerPoolWorkerConfig resource to its proto representation.
func CloudbuildAlphaWorkerPoolWorkerConfigToProto(o *alpha.WorkerPoolWorkerConfig) *alphapb.CloudbuildAlphaWorkerPoolWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolWorkerConfig{
		MachineType:  dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:   dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		NoExternalIp: dcl.ValueOrEmptyBool(o.NoExternalIP),
	}
	return p
}

// WorkerPoolNetworkConfigToProto converts a WorkerPoolNetworkConfig resource to its proto representation.
func CloudbuildAlphaWorkerPoolNetworkConfigToProto(o *alpha.WorkerPoolNetworkConfig) *alphapb.CloudbuildAlphaWorkerPoolNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolNetworkConfig{
		PeeredNetwork: dcl.ValueOrEmptyString(o.PeeredNetwork),
	}
	return p
}

// WorkerPoolToProto converts a WorkerPool resource to its proto representation.
func WorkerPoolToProto(resource *alpha.WorkerPool) *alphapb.CloudbuildAlphaWorkerPool {
	p := &alphapb.CloudbuildAlphaWorkerPool{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:    dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:    dcl.ValueOrEmptyString(resource.DeleteTime),
		State:         CloudbuildAlphaWorkerPoolStateEnumToProto(resource.State),
		WorkerConfig:  CloudbuildAlphaWorkerPoolWorkerConfigToProto(resource.WorkerConfig),
		NetworkConfig: CloudbuildAlphaWorkerPoolNetworkConfigToProto(resource.NetworkConfig),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) applyWorkerPool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudbuildAlphaWorkerPoolRequest) (*alphapb.CloudbuildAlphaWorkerPool, error) {
	p := ProtoToWorkerPool(request.GetResource())
	res, err := c.ApplyWorkerPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkerPoolToProto(res)
	return r, nil
}

// ApplyWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) ApplyCloudbuildAlphaWorkerPool(ctx context.Context, request *alphapb.ApplyCloudbuildAlphaWorkerPoolRequest) (*alphapb.CloudbuildAlphaWorkerPool, error) {
	cl, err := createConfigWorkerPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkerPool(ctx, cl, request)
}

// DeleteWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Delete() method.
func (s *WorkerPoolServer) DeleteCloudbuildAlphaWorkerPool(ctx context.Context, request *alphapb.DeleteCloudbuildAlphaWorkerPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkerPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))

}

// ListCloudbuildAlphaWorkerPool handles the gRPC request by passing it to the underlying WorkerPoolList() method.
func (s *WorkerPoolServer) ListCloudbuildAlphaWorkerPool(ctx context.Context, request *alphapb.ListCloudbuildAlphaWorkerPoolRequest) (*alphapb.ListCloudbuildAlphaWorkerPoolResponse, error) {
	cl, err := createConfigWorkerPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudbuildAlphaWorkerPool
	for _, r := range resources.Items {
		rp := WorkerPoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListCloudbuildAlphaWorkerPoolResponse{Items: protos}, nil
}

func createConfigWorkerPool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
