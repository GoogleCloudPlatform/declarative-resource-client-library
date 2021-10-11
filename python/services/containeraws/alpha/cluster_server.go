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

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerawsAlphaClusterNetworking(p *alphapb.ContainerawsAlphaClusterNetworking) *alpha.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterNetworking{
		VPCId: dcl.StringOrNil(p.GetVpcId()),
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

// ProtoToClusterControlPlane converts a ClusterControlPlane object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlane(p *alphapb.ContainerawsAlphaClusterControlPlane) *alpha.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlane{
		Version:                   dcl.StringOrNil(p.GetVersion()),
		InstanceType:              dcl.StringOrNil(p.GetInstanceType()),
		SshConfig:                 ProtoToContainerawsAlphaClusterControlPlaneSshConfig(p.GetSshConfig()),
		IamInstanceProfile:        dcl.StringOrNil(p.GetIamInstanceProfile()),
		RootVolume:                ProtoToContainerawsAlphaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToContainerawsAlphaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToContainerawsAlphaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToContainerawsAlphaClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
		ProxyConfig:               ProtoToContainerawsAlphaClusterControlPlaneProxyConfig(p.GetProxyConfig()),
	}
	for _, r := range p.GetSubnetIds() {
		obj.SubnetIds = append(obj.SubnetIds, r)
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneSshConfig(p *alphapb.ContainerawsAlphaClusterControlPlaneSshConfig) *alpha.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneRootVolume(p *alphapb.ContainerawsAlphaClusterControlPlaneRootVolume) *alpha.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneMainVolume(p *alphapb.ContainerawsAlphaClusterControlPlaneMainVolume) *alpha.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption) *alpha.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneAwsServicesAuthentication converts a ClusterControlPlaneAwsServicesAuthentication object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneAwsServicesAuthentication(p *alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication) *alpha.ClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.GetRoleArn()),
		RoleSessionName: dcl.StringOrNil(p.GetRoleSessionName()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneProxyConfig(p *alphapb.ContainerawsAlphaClusterControlPlaneProxyConfig) *alpha.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
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

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerawsAlphaClusterAuthorizationAdminUsers(p *alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers) *alpha.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterWorkloadIdentityConfig(p *alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig) *alpha.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.GetIssuerUri()),
		WorkloadPool:     dcl.StringOrNil(p.GetWorkloadPool()),
		IdentityProvider: dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.ContainerawsAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Name:                   dcl.StringOrNil(p.GetName()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Networking:             ProtoToContainerawsAlphaClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.GetAwsRegion()),
		ControlPlane:           ProtoToContainerawsAlphaClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerawsAlphaClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerawsAlphaClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.GetEndpoint()),
		Uid:                    dcl.StringOrNil(p.GetUid()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig: ProtoToContainerawsAlphaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
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

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerawsAlphaClusterNetworkingToProto(o *alpha.ClusterNetworking) *alphapb.ContainerawsAlphaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterNetworking{}
	p.SetVpcId(dcl.ValueOrEmptyString(o.VPCId))
	sPodAddressCidrBlocks := make([]string, len(o.PodAddressCidrBlocks))
	for i, r := range o.PodAddressCidrBlocks {
		sPodAddressCidrBlocks[i] = r
	}
	p.SetPodAddressCidrBlocks(sPodAddressCidrBlocks)
	sServiceAddressCidrBlocks := make([]string, len(o.ServiceAddressCidrBlocks))
	for i, r := range o.ServiceAddressCidrBlocks {
		sServiceAddressCidrBlocks[i] = r
	}
	p.SetServiceAddressCidrBlocks(sServiceAddressCidrBlocks)
	sServiceLoadBalancerSubnetIds := make([]string, len(o.ServiceLoadBalancerSubnetIds))
	for i, r := range o.ServiceLoadBalancerSubnetIds {
		sServiceLoadBalancerSubnetIds[i] = r
	}
	p.SetServiceLoadBalancerSubnetIds(sServiceLoadBalancerSubnetIds)
	return p
}

// ClusterControlPlaneToProto converts a ClusterControlPlane object to its proto representation.
func ContainerawsAlphaClusterControlPlaneToProto(o *alpha.ClusterControlPlane) *alphapb.ContainerawsAlphaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetSshConfig(ContainerawsAlphaClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetRootVolume(ContainerawsAlphaClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerawsAlphaClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerawsAlphaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetAwsServicesAuthentication(ContainerawsAlphaClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication))
	p.SetProxyConfig(ContainerawsAlphaClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
	sSubnetIds := make([]string, len(o.SubnetIds))
	for i, r := range o.SubnetIds {
		sSubnetIds[i] = r
	}
	p.SetSubnetIds(sSubnetIds)
	sSecurityGroupIds := make([]string, len(o.SecurityGroupIds))
	for i, r := range o.SecurityGroupIds {
		sSecurityGroupIds[i] = r
	}
	p.SetSecurityGroupIds(sSecurityGroupIds)
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig object to its proto representation.
func ContainerawsAlphaClusterControlPlaneSshConfigToProto(o *alpha.ClusterControlPlaneSshConfig) *alphapb.ContainerawsAlphaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerawsAlphaClusterControlPlaneRootVolumeToProto(o *alpha.ClusterControlPlaneRootVolume) *alphapb.ContainerawsAlphaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerawsAlphaClusterControlPlaneMainVolumeToProto(o *alpha.ClusterControlPlaneMainVolume) *alphapb.ContainerawsAlphaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerawsAlphaClusterControlPlaneDatabaseEncryptionToProto(o *alpha.ClusterControlPlaneDatabaseEncryption) *alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneAwsServicesAuthenticationToProto converts a ClusterControlPlaneAwsServicesAuthentication object to its proto representation.
func ContainerawsAlphaClusterControlPlaneAwsServicesAuthenticationToProto(o *alpha.ClusterControlPlaneAwsServicesAuthentication) *alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication{}
	p.SetRoleArn(dcl.ValueOrEmptyString(o.RoleArn))
	p.SetRoleSessionName(dcl.ValueOrEmptyString(o.RoleSessionName))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerawsAlphaClusterControlPlaneProxyConfigToProto(o *alpha.ClusterControlPlaneProxyConfig) *alphapb.ContainerawsAlphaClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerawsAlphaClusterAuthorizationToProto(o *alpha.ClusterAuthorization) *alphapb.ContainerawsAlphaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorization{}
	sAdminUsers := make([]*alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerawsAlphaClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerawsAlphaClusterAuthorizationAdminUsersToProto(o *alpha.ClusterAuthorizationAdminUsers) *alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerawsAlphaClusterWorkloadIdentityConfigToProto(o *alpha.ClusterWorkloadIdentityConfig) *alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.ContainerawsAlphaCluster {
	p := &alphapb.ContainerawsAlphaCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetNetworking(ContainerawsAlphaClusterNetworkingToProto(resource.Networking))
	p.SetAwsRegion(dcl.ValueOrEmptyString(resource.AwsRegion))
	p.SetControlPlane(ContainerawsAlphaClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerawsAlphaClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerawsAlphaClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerawsAlphaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaClusterRequest) (*alphapb.ContainerawsAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerawsAlphaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerawsAlphaCluster(ctx context.Context, request *alphapb.ApplyContainerawsAlphaClusterRequest) (*alphapb.ContainerawsAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerawsAlphaCluster(ctx context.Context, request *alphapb.DeleteContainerawsAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerawsAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerawsAlphaCluster(ctx context.Context, request *alphapb.ListContainerawsAlphaClusterRequest) (*alphapb.ListContainerawsAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContainerawsAlphaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
