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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/alpha/containerazure_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
)

// Server implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerazureAlphaNodePoolStateEnum(e alphapb.ContainerazureAlphaNodePoolStateEnum) *alpha.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.NodePoolStateEnum(n[len("ContainerazureAlphaNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig resource from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfig(p *alphapb.ContainerazureAlphaNodePoolConfig) *alpha.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfig{
		VmSize:     dcl.StringOrNil(p.VmSize),
		RootVolume: ProtoToContainerazureAlphaNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:  ProtoToContainerazureAlphaNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume resource from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfigRootVolume(p *alphapb.ContainerazureAlphaNodePoolConfigRootVolume) *alpha.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig resource from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfigSshConfig(p *alphapb.ContainerazureAlphaNodePoolConfigSshConfig) *alpha.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling resource from its proto representation.
func ProtoToContainerazureAlphaNodePoolAutoscaling(p *alphapb.ContainerazureAlphaNodePoolAutoscaling) *alpha.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToContainerazureAlphaNodePoolMaxPodsConstraint(p *alphapb.ContainerazureAlphaNodePoolMaxPodsConstraint) *alpha.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *alphapb.ContainerazureAlphaNodePool) *alpha.NodePool {
	obj := &alpha.NodePool{
		Name:                  dcl.StringOrNil(p.Name),
		Version:               dcl.StringOrNil(p.Version),
		Config:                ProtoToContainerazureAlphaNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.SubnetId),
		Autoscaling:           ProtoToContainerazureAlphaNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToContainerazureAlphaNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.Uid),
		Reconciling:           dcl.Bool(p.Reconciling),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.Etag),
		MaxPodsConstraint:     ProtoToContainerazureAlphaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		AzureAvailabilityZone: dcl.StringOrNil(p.AzureAvailabilityZone),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
		Cluster:               dcl.StringOrNil(p.Cluster),
	}
	return obj
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerazureAlphaNodePoolStateEnumToProto(e *alpha.NodePoolStateEnum) alphapb.ContainerazureAlphaNodePoolStateEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaNodePoolStateEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaNodePoolStateEnum(v)
	}
	return alphapb.ContainerazureAlphaNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig resource to its proto representation.
func ContainerazureAlphaNodePoolConfigToProto(o *alpha.NodePoolConfig) *alphapb.ContainerazureAlphaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfig{
		VmSize:     dcl.ValueOrEmptyString(o.VmSize),
		RootVolume: ContainerazureAlphaNodePoolConfigRootVolumeToProto(o.RootVolume),
		SshConfig:  ContainerazureAlphaNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// NodePoolConfigRootVolumeToProto converts a NodePoolConfigRootVolume resource to its proto representation.
func ContainerazureAlphaNodePoolConfigRootVolumeToProto(o *alpha.NodePoolConfigRootVolume) *alphapb.ContainerazureAlphaNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfigRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig resource to its proto representation.
func ContainerazureAlphaNodePoolConfigSshConfigToProto(o *alpha.NodePoolConfigSshConfig) *alphapb.ContainerazureAlphaNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfigSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling resource to its proto representation.
func ContainerazureAlphaNodePoolAutoscalingToProto(o *alpha.NodePoolAutoscaling) *alphapb.ContainerazureAlphaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint resource to its proto representation.
func ContainerazureAlphaNodePoolMaxPodsConstraintToProto(o *alpha.NodePoolMaxPodsConstraint) *alphapb.ContainerazureAlphaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *alpha.NodePool) *alphapb.ContainerazureAlphaNodePool {
	p := &alphapb.ContainerazureAlphaNodePool{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		Version:               dcl.ValueOrEmptyString(resource.Version),
		Config:                ContainerazureAlphaNodePoolConfigToProto(resource.Config),
		SubnetId:              dcl.ValueOrEmptyString(resource.SubnetId),
		Autoscaling:           ContainerazureAlphaNodePoolAutoscalingToProto(resource.Autoscaling),
		State:                 ContainerazureAlphaNodePoolStateEnumToProto(resource.State),
		Uid:                   dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:           dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                  dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint:     ContainerazureAlphaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		AzureAvailabilityZone: dcl.ValueOrEmptyString(resource.AzureAvailabilityZone),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
		Cluster:               dcl.ValueOrEmptyString(resource.Cluster),
	}

	return p
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaNodePoolRequest) (*alphapb.ContainerazureAlphaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerazureAlphaNodePool(ctx context.Context, request *alphapb.ApplyContainerazureAlphaNodePoolRequest) (*alphapb.ContainerazureAlphaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerazureAlphaNodePool(ctx context.Context, request *alphapb.DeleteContainerazureAlphaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerazureAlphaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerazureAlphaNodePool(ctx context.Context, request *alphapb.ListContainerazureAlphaNodePoolRequest) (*alphapb.ListContainerazureAlphaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, ProtoToNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerazureAlphaNodePoolResponse{Items: protos}, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
