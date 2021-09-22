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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeraws/alpha/containeraws_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/alpha"
)

// Server implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolConfigRootVolumeVolumeTypeEnum converts a NodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum) *alpha.NodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.NodePoolConfigRootVolumeVolumeTypeEnum(n[len("ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigTaintsEffectEnum converts a NodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigTaintsEffectEnum(e alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum) *alpha.NodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := alpha.NodePoolConfigTaintsEffectEnum(n[len("ContainerawsAlphaNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolStateEnum(e alphapb.ContainerawsAlphaNodePoolStateEnum) *alpha.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.NodePoolStateEnum(n[len("ContainerawsAlphaNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig resource from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfig(p *alphapb.ContainerawsAlphaNodePoolConfig) *alpha.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfig{
		InstanceType:       dcl.StringOrNil(p.InstanceType),
		RootVolume:         ProtoToContainerawsAlphaNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile: dcl.StringOrNil(p.IamInstanceProfile),
		SshConfig:          ProtoToContainerawsAlphaNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerawsAlphaNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume resource from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigRootVolume(p *alphapb.ContainerawsAlphaNodePoolConfigRootVolume) *alpha.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToNodePoolConfigTaints converts a NodePoolConfigTaints resource from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigTaints(p *alphapb.ContainerawsAlphaNodePoolConfigTaints) *alpha.NodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToContainerawsAlphaNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig resource from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigSshConfig(p *alphapb.ContainerawsAlphaNodePoolConfigSshConfig) *alpha.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling resource from its proto representation.
func ProtoToContainerawsAlphaNodePoolAutoscaling(p *alphapb.ContainerawsAlphaNodePoolAutoscaling) *alpha.NodePoolAutoscaling {
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
func ProtoToContainerawsAlphaNodePoolMaxPodsConstraint(p *alphapb.ContainerawsAlphaNodePoolMaxPodsConstraint) *alpha.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *alphapb.ContainerawsAlphaNodePool) *alpha.NodePool {
	obj := &alpha.NodePool{
		Name:              dcl.StringOrNil(p.Name),
		Version:           dcl.StringOrNil(p.Version),
		Config:            ProtoToContainerawsAlphaNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToContainerawsAlphaNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.SubnetId),
		State:             ProtoToContainerawsAlphaNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.Uid),
		Reconciling:       dcl.Bool(p.Reconciling),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.Etag),
		MaxPodsConstraint: ProtoToContainerawsAlphaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
		Cluster:           dcl.StringOrNil(p.Cluster),
	}
	return obj
}

// NodePoolConfigRootVolumeVolumeTypeEnumToProto converts a NodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *alpha.NodePoolConfigRootVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum_value["NodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// NodePoolConfigTaintsEffectEnumToProto converts a NodePoolConfigTaintsEffectEnum enum to its proto representation.
func ContainerawsAlphaNodePoolConfigTaintsEffectEnumToProto(e *alpha.NodePoolConfigTaintsEffectEnum) alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum_value["NodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum(0)
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerawsAlphaNodePoolStateEnumToProto(e *alpha.NodePoolStateEnum) alphapb.ContainerawsAlphaNodePoolStateEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolStateEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolStateEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig resource to its proto representation.
func ContainerawsAlphaNodePoolConfigToProto(o *alpha.NodePoolConfig) *alphapb.ContainerawsAlphaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfig{
		InstanceType:       dcl.ValueOrEmptyString(o.InstanceType),
		RootVolume:         ContainerawsAlphaNodePoolConfigRootVolumeToProto(o.RootVolume),
		IamInstanceProfile: dcl.ValueOrEmptyString(o.IamInstanceProfile),
		SshConfig:          ContainerawsAlphaNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerawsAlphaNodePoolConfigTaintsToProto(&r))
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	for _, r := range o.SecurityGroupIds {
		p.SecurityGroupIds = append(p.SecurityGroupIds, r)
	}
	return p
}

// NodePoolConfigRootVolumeToProto converts a NodePoolConfigRootVolume resource to its proto representation.
func ContainerawsAlphaNodePoolConfigRootVolumeToProto(o *alpha.NodePoolConfigRootVolume) *alphapb.ContainerawsAlphaNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// NodePoolConfigTaintsToProto converts a NodePoolConfigTaints resource to its proto representation.
func ContainerawsAlphaNodePoolConfigTaintsToProto(o *alpha.NodePoolConfigTaints) *alphapb.ContainerawsAlphaNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: ContainerawsAlphaNodePoolConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig resource to its proto representation.
func ContainerawsAlphaNodePoolConfigSshConfigToProto(o *alpha.NodePoolConfigSshConfig) *alphapb.ContainerawsAlphaNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling resource to its proto representation.
func ContainerawsAlphaNodePoolAutoscalingToProto(o *alpha.NodePoolAutoscaling) *alphapb.ContainerawsAlphaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint resource to its proto representation.
func ContainerawsAlphaNodePoolMaxPodsConstraintToProto(o *alpha.NodePoolMaxPodsConstraint) *alphapb.ContainerawsAlphaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *alpha.NodePool) *alphapb.ContainerawsAlphaNodePool {
	p := &alphapb.ContainerawsAlphaNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Config:            ContainerawsAlphaNodePoolConfigToProto(resource.Config),
		Autoscaling:       ContainerawsAlphaNodePoolAutoscalingToProto(resource.Autoscaling),
		SubnetId:          dcl.ValueOrEmptyString(resource.SubnetId),
		State:             ContainerawsAlphaNodePoolStateEnumToProto(resource.State),
		Uid:               dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:       dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:              dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint: ContainerawsAlphaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
		Cluster:           dcl.ValueOrEmptyString(resource.Cluster),
	}

	return p
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaNodePoolRequest) (*alphapb.ContainerawsAlphaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerawsAlphaNodePool(ctx context.Context, request *alphapb.ApplyContainerawsAlphaNodePoolRequest) (*alphapb.ContainerawsAlphaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerawsAlphaNodePool(ctx context.Context, request *alphapb.DeleteContainerawsAlphaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerawsAlphaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerawsAlphaNodePool(ctx context.Context, request *alphapb.ListContainerawsAlphaNodePoolRequest) (*alphapb.ListContainerawsAlphaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.Project, request.Location, request.Cluster)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerawsAlphaNodePoolResponse{Items: protos}, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
