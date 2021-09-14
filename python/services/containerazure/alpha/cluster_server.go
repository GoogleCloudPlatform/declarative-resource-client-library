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

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerazureAlphaClusterStateEnum(e alphapb.ContainerazureAlphaClusterStateEnum) *alpha.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaClusterStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStateEnum(n[len("ContainerazureAlphaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworking converts a ClusterNetworking resource from its proto representation.
func ProtoToContainerazureAlphaClusterNetworking(p *alphapb.ContainerazureAlphaClusterNetworking) *alpha.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterNetworking{
		VirtualNetworkId: dcl.StringOrNil(p.VirtualNetworkId),
	}
	for _, r := range p.GetPodAddressCidrBlocks() {
		obj.PodAddressCidrBlocks = append(obj.PodAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceAddressCidrBlocks() {
		obj.ServiceAddressCidrBlocks = append(obj.ServiceAddressCidrBlocks, r)
	}
	return obj
}

// ProtoToClusterControlPlane converts a ClusterControlPlane resource from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlane(p *alphapb.ContainerazureAlphaClusterControlPlane) *alpha.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlane{
		Version:            dcl.StringOrNil(p.Version),
		SubnetId:           dcl.StringOrNil(p.SubnetId),
		VmSize:             dcl.StringOrNil(p.VmSize),
		SshConfig:          ProtoToContainerazureAlphaClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToContainerazureAlphaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToContainerazureAlphaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToContainerazureAlphaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneSshConfig(p *alphapb.ContainerazureAlphaClusterControlPlaneSshConfig) *alpha.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneRootVolume(p *alphapb.ContainerazureAlphaClusterControlPlaneRootVolume) *alpha.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneMainVolume(p *alphapb.ContainerazureAlphaClusterControlPlaneMainVolume) *alpha.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerazureAlphaClusterControlPlaneDatabaseEncryption) *alpha.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.StringOrNil(p.ResourceGroupId),
		KmsKeyIdentifier: dcl.StringOrNil(p.KmsKeyIdentifier),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization resource from its proto representation.
func ProtoToContainerazureAlphaClusterAuthorization(p *alphapb.ContainerazureAlphaClusterAuthorization) *alpha.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerazureAlphaClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToContainerazureAlphaClusterAuthorizationAdminUsers(p *alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers) *alpha.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToContainerazureAlphaClusterWorkloadIdentityConfig(p *alphapb.ContainerazureAlphaClusterWorkloadIdentityConfig) *alpha.ClusterWorkloadIdentityConfig {
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
func ProtoToCluster(p *alphapb.ContainerazureAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		AzureRegion:            dcl.StringOrNil(p.AzureRegion),
		ResourceGroupId:        dcl.StringOrNil(p.ResourceGroupId),
		Client:                 dcl.StringOrNil(p.Client),
		Networking:             ProtoToContainerazureAlphaClusterNetworking(p.GetNetworking()),
		ControlPlane:           ProtoToContainerazureAlphaClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerazureAlphaClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerazureAlphaClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToContainerazureAlphaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerazureAlphaClusterStateEnumToProto(e *alpha.ClusterStateEnum) alphapb.ContainerazureAlphaClusterStateEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaClusterStateEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaClusterStateEnum(v)
	}
	return alphapb.ContainerazureAlphaClusterStateEnum(0)
}

// ClusterNetworkingToProto converts a ClusterNetworking resource to its proto representation.
func ContainerazureAlphaClusterNetworkingToProto(o *alpha.ClusterNetworking) *alphapb.ContainerazureAlphaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterNetworking{
		VirtualNetworkId: dcl.ValueOrEmptyString(o.VirtualNetworkId),
	}
	for _, r := range o.PodAddressCidrBlocks {
		p.PodAddressCidrBlocks = append(p.PodAddressCidrBlocks, r)
	}
	for _, r := range o.ServiceAddressCidrBlocks {
		p.ServiceAddressCidrBlocks = append(p.ServiceAddressCidrBlocks, r)
	}
	return p
}

// ClusterControlPlaneToProto converts a ClusterControlPlane resource to its proto representation.
func ContainerazureAlphaClusterControlPlaneToProto(o *alpha.ClusterControlPlane) *alphapb.ContainerazureAlphaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlane{
		Version:            dcl.ValueOrEmptyString(o.Version),
		SubnetId:           dcl.ValueOrEmptyString(o.SubnetId),
		VmSize:             dcl.ValueOrEmptyString(o.VmSize),
		SshConfig:          ContainerazureAlphaClusterControlPlaneSshConfigToProto(o.SshConfig),
		RootVolume:         ContainerazureAlphaClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:         ContainerazureAlphaClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption: ContainerazureAlphaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig resource to its proto representation.
func ContainerazureAlphaClusterControlPlaneSshConfigToProto(o *alpha.ClusterControlPlaneSshConfig) *alphapb.ContainerazureAlphaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume resource to its proto representation.
func ContainerazureAlphaClusterControlPlaneRootVolumeToProto(o *alpha.ClusterControlPlaneRootVolume) *alphapb.ContainerazureAlphaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume resource to its proto representation.
func ContainerazureAlphaClusterControlPlaneMainVolumeToProto(o *alpha.ClusterControlPlaneMainVolume) *alphapb.ContainerazureAlphaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneMainVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption resource to its proto representation.
func ContainerazureAlphaClusterControlPlaneDatabaseEncryptionToProto(o *alpha.ClusterControlPlaneDatabaseEncryption) *alphapb.ContainerazureAlphaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.ValueOrEmptyString(o.ResourceGroupId),
		KmsKeyIdentifier: dcl.ValueOrEmptyString(o.KmsKeyIdentifier),
	}
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization resource to its proto representation.
func ContainerazureAlphaClusterAuthorizationToProto(o *alpha.ClusterAuthorization) *alphapb.ContainerazureAlphaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, ContainerazureAlphaClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers resource to its proto representation.
func ContainerazureAlphaClusterAuthorizationAdminUsersToProto(o *alpha.ClusterAuthorizationAdminUsers) *alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig resource to its proto representation.
func ContainerazureAlphaClusterWorkloadIdentityConfigToProto(o *alpha.ClusterWorkloadIdentityConfig) *alphapb.ContainerazureAlphaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.ContainerazureAlphaCluster {
	p := &alphapb.ContainerazureAlphaCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		AzureRegion:            dcl.ValueOrEmptyString(resource.AzureRegion),
		ResourceGroupId:        dcl.ValueOrEmptyString(resource.ResourceGroupId),
		Client:                 dcl.ValueOrEmptyString(resource.Client),
		Networking:             ContainerazureAlphaClusterNetworkingToProto(resource.Networking),
		ControlPlane:           ContainerazureAlphaClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          ContainerazureAlphaClusterAuthorizationToProto(resource.Authorization),
		State:                  ContainerazureAlphaClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: ContainerazureAlphaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaClusterRequest) (*alphapb.ContainerazureAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerazureAlphaCluster(ctx context.Context, request *alphapb.ApplyContainerazureAlphaClusterRequest) (*alphapb.ContainerazureAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerazureAlphaCluster(ctx context.Context, request *alphapb.DeleteContainerazureAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerazureAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerazureAlphaCluster(ctx context.Context, request *alphapb.ListContainerazureAlphaClusterRequest) (*alphapb.ListContainerazureAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, ProtoToCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerazureAlphaClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
