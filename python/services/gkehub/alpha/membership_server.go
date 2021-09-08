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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// Server implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipStateCodeEnum converts a MembershipStateCodeEnum enum from its proto representation.
func ProtoToGkehubAlphaMembershipStateCodeEnum(e alphapb.GkehubAlphaMembershipStateCodeEnum) *alpha.MembershipStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaMembershipStateCodeEnum_name[int32(e)]; ok {
		e := alpha.MembershipStateCodeEnum(n[len("GkehubAlphaMembershipStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipInfrastructureTypeEnum converts a MembershipInfrastructureTypeEnum enum from its proto representation.
func ProtoToGkehubAlphaMembershipInfrastructureTypeEnum(e alphapb.GkehubAlphaMembershipInfrastructureTypeEnum) *alpha.MembershipInfrastructureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaMembershipInfrastructureTypeEnum_name[int32(e)]; ok {
		e := alpha.MembershipInfrastructureTypeEnum(n[len("GkehubAlphaMembershipInfrastructureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipEndpoint converts a MembershipEndpoint resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpoint(p *alphapb.GkehubAlphaMembershipEndpoint) *alpha.MembershipEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpoint{
		GkeCluster:         ProtoToGkehubAlphaMembershipEndpointGkeCluster(p.GetGkeCluster()),
		KubernetesMetadata: ProtoToGkehubAlphaMembershipEndpointKubernetesMetadata(p.GetKubernetesMetadata()),
		KubernetesResource: ProtoToGkehubAlphaMembershipEndpointKubernetesResource(p.GetKubernetesResource()),
	}
	return obj
}

// ProtoToMembershipEndpointGkeCluster converts a MembershipEndpointGkeCluster resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointGkeCluster(p *alphapb.GkehubAlphaMembershipEndpointGkeCluster) *alpha.MembershipEndpointGkeCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointGkeCluster{
		ResourceLink: dcl.StringOrNil(p.ResourceLink),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesMetadata converts a MembershipEndpointKubernetesMetadata resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesMetadata(p *alphapb.GkehubAlphaMembershipEndpointKubernetesMetadata) *alpha.MembershipEndpointKubernetesMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesMetadata{
		KubernetesApiServerVersion: dcl.StringOrNil(p.KubernetesApiServerVersion),
		NodeProviderId:             dcl.StringOrNil(p.NodeProviderId),
		NodeCount:                  dcl.Int64OrNil(p.NodeCount),
		VcpuCount:                  dcl.Int64OrNil(p.VcpuCount),
		MemoryMb:                   dcl.Int64OrNil(p.MemoryMb),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResource converts a MembershipEndpointKubernetesResource resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResource(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResource) *alpha.MembershipEndpointKubernetesResource {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResource{
		MembershipCrManifest: dcl.StringOrNil(p.MembershipCrManifest),
		ResourceOptions:      ProtoToGkehubAlphaMembershipEndpointKubernetesResourceResourceOptions(p.GetResourceOptions()),
	}
	for _, r := range p.GetMembershipResources() {
		obj.MembershipResources = append(obj.MembershipResources, *ProtoToGkehubAlphaMembershipEndpointKubernetesResourceMembershipResources(r))
	}
	for _, r := range p.GetConnectResources() {
		obj.ConnectResources = append(obj.ConnectResources, *ProtoToGkehubAlphaMembershipEndpointKubernetesResourceConnectResources(r))
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceMembershipResources converts a MembershipEndpointKubernetesResourceMembershipResources resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResourceMembershipResources(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources) *alpha.MembershipEndpointKubernetesResourceMembershipResources {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResourceMembershipResources{
		Manifest:      dcl.StringOrNil(p.Manifest),
		ClusterScoped: dcl.Bool(p.ClusterScoped),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceConnectResources converts a MembershipEndpointKubernetesResourceConnectResources resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResourceConnectResources(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources) *alpha.MembershipEndpointKubernetesResourceConnectResources {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResourceConnectResources{
		Manifest:      dcl.StringOrNil(p.Manifest),
		ClusterScoped: dcl.Bool(p.ClusterScoped),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceResourceOptions converts a MembershipEndpointKubernetesResourceResourceOptions resource from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResourceResourceOptions(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceResourceOptions) *alpha.MembershipEndpointKubernetesResourceResourceOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResourceResourceOptions{
		ConnectVersion: dcl.StringOrNil(p.ConnectVersion),
		V1Beta1Crd:     dcl.Bool(p.V1Beta1Crd),
	}
	return obj
}

// ProtoToMembershipState converts a MembershipState resource from its proto representation.
func ProtoToGkehubAlphaMembershipState(p *alphapb.GkehubAlphaMembershipState) *alpha.MembershipState {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipState{
		Code: ProtoToGkehubAlphaMembershipStateCodeEnum(p.GetCode()),
	}
	return obj
}

// ProtoToMembershipAuthority converts a MembershipAuthority resource from its proto representation.
func ProtoToGkehubAlphaMembershipAuthority(p *alphapb.GkehubAlphaMembershipAuthority) *alpha.MembershipAuthority {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipAuthority{
		Issuer:               dcl.StringOrNil(p.Issuer),
		WorkloadIdentityPool: dcl.StringOrNil(p.WorkloadIdentityPool),
		IdentityProvider:     dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *alphapb.GkehubAlphaMembership) *alpha.Membership {
	obj := &alpha.Membership{
		Endpoint:           ProtoToGkehubAlphaMembershipEndpoint(p.GetEndpoint()),
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		State:              ProtoToGkehubAlphaMembershipState(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:         dcl.StringOrNil(p.GetDeleteTime()),
		ExternalId:         dcl.StringOrNil(p.ExternalId),
		LastConnectionTime: dcl.StringOrNil(p.GetLastConnectionTime()),
		UniqueId:           dcl.StringOrNil(p.UniqueId),
		Authority:          ProtoToGkehubAlphaMembershipAuthority(p.GetAuthority()),
		InfrastructureType: ProtoToGkehubAlphaMembershipInfrastructureTypeEnum(p.GetInfrastructureType()),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	return obj
}

// MembershipStateCodeEnumToProto converts a MembershipStateCodeEnum enum to its proto representation.
func GkehubAlphaMembershipStateCodeEnumToProto(e *alpha.MembershipStateCodeEnum) alphapb.GkehubAlphaMembershipStateCodeEnum {
	if e == nil {
		return alphapb.GkehubAlphaMembershipStateCodeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaMembershipStateCodeEnum_value["MembershipStateCodeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaMembershipStateCodeEnum(v)
	}
	return alphapb.GkehubAlphaMembershipStateCodeEnum(0)
}

// MembershipInfrastructureTypeEnumToProto converts a MembershipInfrastructureTypeEnum enum to its proto representation.
func GkehubAlphaMembershipInfrastructureTypeEnumToProto(e *alpha.MembershipInfrastructureTypeEnum) alphapb.GkehubAlphaMembershipInfrastructureTypeEnum {
	if e == nil {
		return alphapb.GkehubAlphaMembershipInfrastructureTypeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaMembershipInfrastructureTypeEnum_value["MembershipInfrastructureTypeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaMembershipInfrastructureTypeEnum(v)
	}
	return alphapb.GkehubAlphaMembershipInfrastructureTypeEnum(0)
}

// MembershipEndpointToProto converts a MembershipEndpoint resource to its proto representation.
func GkehubAlphaMembershipEndpointToProto(o *alpha.MembershipEndpoint) *alphapb.GkehubAlphaMembershipEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpoint{
		GkeCluster:         GkehubAlphaMembershipEndpointGkeClusterToProto(o.GkeCluster),
		KubernetesMetadata: GkehubAlphaMembershipEndpointKubernetesMetadataToProto(o.KubernetesMetadata),
		KubernetesResource: GkehubAlphaMembershipEndpointKubernetesResourceToProto(o.KubernetesResource),
	}
	return p
}

// MembershipEndpointGkeClusterToProto converts a MembershipEndpointGkeCluster resource to its proto representation.
func GkehubAlphaMembershipEndpointGkeClusterToProto(o *alpha.MembershipEndpointGkeCluster) *alphapb.GkehubAlphaMembershipEndpointGkeCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointGkeCluster{
		ResourceLink: dcl.ValueOrEmptyString(o.ResourceLink),
	}
	return p
}

// MembershipEndpointKubernetesMetadataToProto converts a MembershipEndpointKubernetesMetadata resource to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesMetadataToProto(o *alpha.MembershipEndpointKubernetesMetadata) *alphapb.GkehubAlphaMembershipEndpointKubernetesMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesMetadata{
		KubernetesApiServerVersion: dcl.ValueOrEmptyString(o.KubernetesApiServerVersion),
		NodeProviderId:             dcl.ValueOrEmptyString(o.NodeProviderId),
		NodeCount:                  dcl.ValueOrEmptyInt64(o.NodeCount),
		VcpuCount:                  dcl.ValueOrEmptyInt64(o.VcpuCount),
		MemoryMb:                   dcl.ValueOrEmptyInt64(o.MemoryMb),
		UpdateTime:                 dcl.ValueOrEmptyString(o.UpdateTime),
	}
	return p
}

// MembershipEndpointKubernetesResourceToProto converts a MembershipEndpointKubernetesResource resource to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceToProto(o *alpha.MembershipEndpointKubernetesResource) *alphapb.GkehubAlphaMembershipEndpointKubernetesResource {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResource{
		MembershipCrManifest: dcl.ValueOrEmptyString(o.MembershipCrManifest),
		ResourceOptions:      GkehubAlphaMembershipEndpointKubernetesResourceResourceOptionsToProto(o.ResourceOptions),
	}
	for _, r := range o.MembershipResources {
		p.MembershipResources = append(p.MembershipResources, GkehubAlphaMembershipEndpointKubernetesResourceMembershipResourcesToProto(&r))
	}
	for _, r := range o.ConnectResources {
		p.ConnectResources = append(p.ConnectResources, GkehubAlphaMembershipEndpointKubernetesResourceConnectResourcesToProto(&r))
	}
	return p
}

// MembershipEndpointKubernetesResourceMembershipResourcesToProto converts a MembershipEndpointKubernetesResourceMembershipResources resource to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceMembershipResourcesToProto(o *alpha.MembershipEndpointKubernetesResourceMembershipResources) *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources{
		Manifest:      dcl.ValueOrEmptyString(o.Manifest),
		ClusterScoped: dcl.ValueOrEmptyBool(o.ClusterScoped),
	}
	return p
}

// MembershipEndpointKubernetesResourceConnectResourcesToProto converts a MembershipEndpointKubernetesResourceConnectResources resource to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceConnectResourcesToProto(o *alpha.MembershipEndpointKubernetesResourceConnectResources) *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources{
		Manifest:      dcl.ValueOrEmptyString(o.Manifest),
		ClusterScoped: dcl.ValueOrEmptyBool(o.ClusterScoped),
	}
	return p
}

// MembershipEndpointKubernetesResourceResourceOptionsToProto converts a MembershipEndpointKubernetesResourceResourceOptions resource to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceResourceOptionsToProto(o *alpha.MembershipEndpointKubernetesResourceResourceOptions) *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceResourceOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResourceResourceOptions{
		ConnectVersion: dcl.ValueOrEmptyString(o.ConnectVersion),
		V1Beta1Crd:     dcl.ValueOrEmptyBool(o.V1Beta1Crd),
	}
	return p
}

// MembershipStateToProto converts a MembershipState resource to its proto representation.
func GkehubAlphaMembershipStateToProto(o *alpha.MembershipState) *alphapb.GkehubAlphaMembershipState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipState{
		Code: GkehubAlphaMembershipStateCodeEnumToProto(o.Code),
	}
	return p
}

// MembershipAuthorityToProto converts a MembershipAuthority resource to its proto representation.
func GkehubAlphaMembershipAuthorityToProto(o *alpha.MembershipAuthority) *alphapb.GkehubAlphaMembershipAuthority {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipAuthority{
		Issuer:               dcl.ValueOrEmptyString(o.Issuer),
		WorkloadIdentityPool: dcl.ValueOrEmptyString(o.WorkloadIdentityPool),
		IdentityProvider:     dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *alpha.Membership) *alphapb.GkehubAlphaMembership {
	p := &alphapb.GkehubAlphaMembership{
		Endpoint:           GkehubAlphaMembershipEndpointToProto(resource.Endpoint),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		State:              GkehubAlphaMembershipStateToProto(resource.State),
		CreateTime:         dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:         dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:         dcl.ValueOrEmptyString(resource.DeleteTime),
		ExternalId:         dcl.ValueOrEmptyString(resource.ExternalId),
		LastConnectionTime: dcl.ValueOrEmptyString(resource.LastConnectionTime),
		UniqueId:           dcl.ValueOrEmptyString(resource.UniqueId),
		Authority:          GkehubAlphaMembershipAuthorityToProto(resource.Authority),
		InfrastructureType: GkehubAlphaMembershipInfrastructureTypeEnumToProto(resource.InfrastructureType),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) applyMembership(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaMembershipRequest) (*alphapb.GkehubAlphaMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// ApplyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyGkehubAlphaMembership(ctx context.Context, request *alphapb.ApplyGkehubAlphaMembershipRequest) (*alphapb.GkehubAlphaMembership, error) {
	cl, err := createConfigMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteGkehubAlphaMembership(ctx context.Context, request *alphapb.DeleteGkehubAlphaMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListGkehubAlphaMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListGkehubAlphaMembership(ctx context.Context, request *alphapb.ListGkehubAlphaMembershipRequest) (*alphapb.ListGkehubAlphaMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, ProtoToMembership(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkehubAlphaMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkehubAlphaMembershipResponse{Items: protos}, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
