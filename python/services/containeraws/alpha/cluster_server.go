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

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterControlPlaneRootVolumeVolumeTypeEnum converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum) *alpha.ClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.ClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterControlPlaneMainVolumeVolumeTypeEnum converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum) *alpha.ClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.ClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterStateEnum(e alphapb.ContainerawsAlphaClusterStateEnum) *alpha.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStateEnum(n[len("ContainerawsAlphaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworking converts a ClusterNetworking resource from its proto representation.
func ProtoToContainerawsAlphaClusterNetworking(p *alphapb.ContainerawsAlphaClusterNetworking) *alpha.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterNetworking{
		VPCId: dcl.StringOrNil(p.VpcId),
	}
	for _, r := range p.GetPodAddressCidrBlocks() {
		obj.PodAddressCidrBlocks = append(obj.PodAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceAddressCidrBlocks() {
		obj.ServiceAddressCidrBlocks = append(obj.ServiceAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceLoadBalancerSubnetIds() {
		obj.ServiceLoadBalancerSubnetIds = append(obj.ServiceLoadBalancerSubnetIds, r)
	}
	return obj
}

// ProtoToClusterControlPlane converts a ClusterControlPlane resource from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlane(p *alphapb.ContainerawsAlphaClusterControlPlane) *alpha.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlane{
		Version:                   dcl.StringOrNil(p.Version),
		InstanceType:              dcl.StringOrNil(p.InstanceType),
		SshConfig:                 ProtoToContainerawsAlphaClusterControlPlaneSshConfig(p.GetSshConfig()),
		IamInstanceProfile:        dcl.StringOrNil(p.IamInstanceProfile),
		RootVolume:                ProtoToContainerawsAlphaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToContainerawsAlphaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToContainerawsAlphaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToContainerawsAlphaClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
	}
	for _, r := range p.GetSubnetIds() {
		obj.SubnetIds = append(obj.SubnetIds, r)
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneSshConfig(p *alphapb.ContainerawsAlphaClusterControlPlaneSshConfig) *alpha.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneRootVolume(p *alphapb.ContainerawsAlphaClusterControlPlaneRootVolume) *alpha.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneMainVolume(p *alphapb.ContainerawsAlphaClusterControlPlaneMainVolume) *alpha.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption) *alpha.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToClusterControlPlaneAwsServicesAuthentication converts a ClusterControlPlaneAwsServicesAuthentication resource from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneAwsServicesAuthentication(p *alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication) *alpha.ClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.RoleArn),
		RoleSessionName: dcl.StringOrNil(p.RoleSessionName),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization resource from its proto representation.
func ProtoToContainerawsAlphaClusterAuthorization(p *alphapb.ContainerawsAlphaClusterAuthorization) *alpha.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerawsAlphaClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToContainerawsAlphaClusterAuthorizationAdminUsers(p *alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers) *alpha.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToContainerawsAlphaClusterWorkloadIdentityConfig(p *alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig) *alpha.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.ContainerawsAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		Networking:             ProtoToContainerawsAlphaClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.AwsRegion),
		ControlPlane:           ProtoToContainerawsAlphaClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerawsAlphaClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerawsAlphaClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToContainerawsAlphaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// ClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *alpha.ClusterControlPlaneRootVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum_value["ClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// ClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *alpha.ClusterControlPlaneMainVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum_value["ClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerawsAlphaClusterStateEnumToProto(e *alpha.ClusterStateEnum) alphapb.ContainerawsAlphaClusterStateEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterStateEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterStateEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterStateEnum(0)
}

// ClusterNetworkingToProto converts a ClusterNetworking resource to its proto representation.
func ContainerawsAlphaClusterNetworkingToProto(o *alpha.ClusterNetworking) *alphapb.ContainerawsAlphaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterNetworking{
		VpcId: dcl.ValueOrEmptyString(o.VPCId),
	}
	for _, r := range o.PodAddressCidrBlocks {
		p.PodAddressCidrBlocks = append(p.PodAddressCidrBlocks, r)
	}
	for _, r := range o.ServiceAddressCidrBlocks {
		p.ServiceAddressCidrBlocks = append(p.ServiceAddressCidrBlocks, r)
	}
	for _, r := range o.ServiceLoadBalancerSubnetIds {
		p.ServiceLoadBalancerSubnetIds = append(p.ServiceLoadBalancerSubnetIds, r)
	}
	return p
}

// ClusterControlPlaneToProto converts a ClusterControlPlane resource to its proto representation.
func ContainerawsAlphaClusterControlPlaneToProto(o *alpha.ClusterControlPlane) *alphapb.ContainerawsAlphaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlane{
		Version:                   dcl.ValueOrEmptyString(o.Version),
		InstanceType:              dcl.ValueOrEmptyString(o.InstanceType),
		SshConfig:                 ContainerawsAlphaClusterControlPlaneSshConfigToProto(o.SshConfig),
		IamInstanceProfile:        dcl.ValueOrEmptyString(o.IamInstanceProfile),
		RootVolume:                ContainerawsAlphaClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:                ContainerawsAlphaClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption:        ContainerawsAlphaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
		AwsServicesAuthentication: ContainerawsAlphaClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication),
	}
	for _, r := range o.SubnetIds {
		p.SubnetIds = append(p.SubnetIds, r)
	}
	for _, r := range o.SecurityGroupIds {
		p.SecurityGroupIds = append(p.SecurityGroupIds, r)
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig resource to its proto representation.
func ContainerawsAlphaClusterControlPlaneSshConfigToProto(o *alpha.ClusterControlPlaneSshConfig) *alphapb.ContainerawsAlphaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume resource to its proto representation.
func ContainerawsAlphaClusterControlPlaneRootVolumeToProto(o *alpha.ClusterControlPlaneRootVolume) *alphapb.ContainerawsAlphaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume resource to its proto representation.
func ContainerawsAlphaClusterControlPlaneMainVolumeToProto(o *alpha.ClusterControlPlaneMainVolume) *alphapb.ContainerawsAlphaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneMainVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption resource to its proto representation.
func ContainerawsAlphaClusterControlPlaneDatabaseEncryptionToProto(o *alpha.ClusterControlPlaneDatabaseEncryption) *alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// ClusterControlPlaneAwsServicesAuthenticationToProto converts a ClusterControlPlaneAwsServicesAuthentication resource to its proto representation.
func ContainerawsAlphaClusterControlPlaneAwsServicesAuthenticationToProto(o *alpha.ClusterControlPlaneAwsServicesAuthentication) *alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.ValueOrEmptyString(o.RoleArn),
		RoleSessionName: dcl.ValueOrEmptyString(o.RoleSessionName),
	}
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization resource to its proto representation.
func ContainerawsAlphaClusterAuthorizationToProto(o *alpha.ClusterAuthorization) *alphapb.ContainerawsAlphaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, ContainerawsAlphaClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers resource to its proto representation.
func ContainerawsAlphaClusterAuthorizationAdminUsersToProto(o *alpha.ClusterAuthorizationAdminUsers) *alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig resource to its proto representation.
func ContainerawsAlphaClusterWorkloadIdentityConfigToProto(o *alpha.ClusterWorkloadIdentityConfig) *alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.ContainerawsAlphaCluster {
	p := &alphapb.ContainerawsAlphaCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		Networking:             ContainerawsAlphaClusterNetworkingToProto(resource.Networking),
		AwsRegion:              dcl.ValueOrEmptyString(resource.AwsRegion),
		ControlPlane:           ContainerawsAlphaClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          ContainerawsAlphaClusterAuthorizationToProto(resource.Authorization),
		State:                  ContainerawsAlphaClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: ContainerawsAlphaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaClusterRequest) (*alphapb.ContainerawsAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerawsAlphaCluster(ctx context.Context, request *alphapb.ApplyContainerawsAlphaClusterRequest) (*alphapb.ContainerawsAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerawsAlphaCluster(ctx context.Context, request *alphapb.DeleteContainerawsAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerawsAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerawsAlphaCluster(ctx context.Context, request *alphapb.ListContainerawsAlphaClusterRequest) (*alphapb.ListContainerawsAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, ProtoToCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerawsAlphaClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
