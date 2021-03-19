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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// Server implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipStateCodeEnum converts a MembershipStateCodeEnum enum from its proto representation.
func ProtoToGkehubBetaMembershipStateCodeEnum(e betapb.GkehubBetaMembershipStateCodeEnum) *beta.MembershipStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaMembershipStateCodeEnum_name[int32(e)]; ok {
		e := beta.MembershipStateCodeEnum(n[len("MembershipStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipInfrastructureTypeEnum converts a MembershipInfrastructureTypeEnum enum from its proto representation.
func ProtoToGkehubBetaMembershipInfrastructureTypeEnum(e betapb.GkehubBetaMembershipInfrastructureTypeEnum) *beta.MembershipInfrastructureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaMembershipInfrastructureTypeEnum_name[int32(e)]; ok {
		e := beta.MembershipInfrastructureTypeEnum(n[len("MembershipInfrastructureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipEndpoint converts a MembershipEndpoint resource from its proto representation.
func ProtoToGkehubBetaMembershipEndpoint(p *betapb.GkehubBetaMembershipEndpoint) *beta.MembershipEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpoint{
		GkeCluster:         ProtoToGkehubBetaMembershipEndpointGkeCluster(p.GetGkeCluster()),
		KubernetesMetadata: ProtoToGkehubBetaMembershipEndpointKubernetesMetadata(p.GetKubernetesMetadata()),
		KubernetesResource: ProtoToGkehubBetaMembershipEndpointKubernetesResource(p.GetKubernetesResource()),
	}
	return obj
}

// ProtoToMembershipEndpointGkeCluster converts a MembershipEndpointGkeCluster resource from its proto representation.
func ProtoToGkehubBetaMembershipEndpointGkeCluster(p *betapb.GkehubBetaMembershipEndpointGkeCluster) *beta.MembershipEndpointGkeCluster {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointGkeCluster{
		ResourceLink: dcl.StringOrNil(p.ResourceLink),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesMetadata converts a MembershipEndpointKubernetesMetadata resource from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesMetadata(p *betapb.GkehubBetaMembershipEndpointKubernetesMetadata) *beta.MembershipEndpointKubernetesMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesMetadata{
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
func ProtoToGkehubBetaMembershipEndpointKubernetesResource(p *betapb.GkehubBetaMembershipEndpointKubernetesResource) *beta.MembershipEndpointKubernetesResource {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResource{
		MembershipCrManifest: dcl.StringOrNil(p.MembershipCrManifest),
		ResourceOptions:      ProtoToGkehubBetaMembershipEndpointKubernetesResourceResourceOptions(p.GetResourceOptions()),
	}
	for _, r := range p.GetMembershipResources() {
		obj.MembershipResources = append(obj.MembershipResources, *ProtoToGkehubBetaMembershipEndpointKubernetesResourceMembershipResources(r))
	}
	for _, r := range p.GetConnectResources() {
		obj.ConnectResources = append(obj.ConnectResources, *ProtoToGkehubBetaMembershipEndpointKubernetesResourceConnectResources(r))
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceMembershipResources converts a MembershipEndpointKubernetesResourceMembershipResources resource from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResourceMembershipResources(p *betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources) *beta.MembershipEndpointKubernetesResourceMembershipResources {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResourceMembershipResources{
		Manifest:      dcl.StringOrNil(p.Manifest),
		ClusterScoped: dcl.Bool(p.ClusterScoped),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceConnectResources converts a MembershipEndpointKubernetesResourceConnectResources resource from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResourceConnectResources(p *betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources) *beta.MembershipEndpointKubernetesResourceConnectResources {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResourceConnectResources{
		Manifest:      dcl.StringOrNil(p.Manifest),
		ClusterScoped: dcl.Bool(p.ClusterScoped),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceResourceOptions converts a MembershipEndpointKubernetesResourceResourceOptions resource from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResourceResourceOptions(p *betapb.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions) *beta.MembershipEndpointKubernetesResourceResourceOptions {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResourceResourceOptions{
		ConnectVersion: dcl.StringOrNil(p.ConnectVersion),
		V1Beta1Crd:     dcl.Bool(p.V1Beta1Crd),
	}
	return obj
}

// ProtoToMembershipState converts a MembershipState resource from its proto representation.
func ProtoToGkehubBetaMembershipState(p *betapb.GkehubBetaMembershipState) *beta.MembershipState {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipState{
		Code: ProtoToGkehubBetaMembershipStateCodeEnum(p.GetCode()),
	}
	return obj
}

// ProtoToMembershipAuthority converts a MembershipAuthority resource from its proto representation.
func ProtoToGkehubBetaMembershipAuthority(p *betapb.GkehubBetaMembershipAuthority) *beta.MembershipAuthority {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipAuthority{
		Issuer:               dcl.StringOrNil(p.Issuer),
		WorkloadIdentityPool: dcl.StringOrNil(p.WorkloadIdentityPool),
		IdentityProvider:     dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *betapb.GkehubBetaMembership) *beta.Membership {
	obj := &beta.Membership{
		Endpoint:           ProtoToGkehubBetaMembershipEndpoint(p.GetEndpoint()),
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		State:              ProtoToGkehubBetaMembershipState(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:         dcl.StringOrNil(p.GetDeleteTime()),
		ExternalId:         dcl.StringOrNil(p.ExternalId),
		LastConnectionTime: dcl.StringOrNil(p.GetLastConnectionTime()),
		UniqueId:           dcl.StringOrNil(p.UniqueId),
		Authority:          ProtoToGkehubBetaMembershipAuthority(p.GetAuthority()),
		InfrastructureType: ProtoToGkehubBetaMembershipInfrastructureTypeEnum(p.GetInfrastructureType()),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	return obj
}

// MembershipStateCodeEnumToProto converts a MembershipStateCodeEnum enum to its proto representation.
func GkehubBetaMembershipStateCodeEnumToProto(e *beta.MembershipStateCodeEnum) betapb.GkehubBetaMembershipStateCodeEnum {
	if e == nil {
		return betapb.GkehubBetaMembershipStateCodeEnum(0)
	}
	if v, ok := betapb.GkehubBetaMembershipStateCodeEnum_value["MembershipStateCodeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaMembershipStateCodeEnum(v)
	}
	return betapb.GkehubBetaMembershipStateCodeEnum(0)
}

// MembershipInfrastructureTypeEnumToProto converts a MembershipInfrastructureTypeEnum enum to its proto representation.
func GkehubBetaMembershipInfrastructureTypeEnumToProto(e *beta.MembershipInfrastructureTypeEnum) betapb.GkehubBetaMembershipInfrastructureTypeEnum {
	if e == nil {
		return betapb.GkehubBetaMembershipInfrastructureTypeEnum(0)
	}
	if v, ok := betapb.GkehubBetaMembershipInfrastructureTypeEnum_value["MembershipInfrastructureTypeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaMembershipInfrastructureTypeEnum(v)
	}
	return betapb.GkehubBetaMembershipInfrastructureTypeEnum(0)
}

// MembershipEndpointToProto converts a MembershipEndpoint resource to its proto representation.
func GkehubBetaMembershipEndpointToProto(o *beta.MembershipEndpoint) *betapb.GkehubBetaMembershipEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpoint{
		GkeCluster:         GkehubBetaMembershipEndpointGkeClusterToProto(o.GkeCluster),
		KubernetesMetadata: GkehubBetaMembershipEndpointKubernetesMetadataToProto(o.KubernetesMetadata),
		KubernetesResource: GkehubBetaMembershipEndpointKubernetesResourceToProto(o.KubernetesResource),
	}
	return p
}

// MembershipEndpointGkeClusterToProto converts a MembershipEndpointGkeCluster resource to its proto representation.
func GkehubBetaMembershipEndpointGkeClusterToProto(o *beta.MembershipEndpointGkeCluster) *betapb.GkehubBetaMembershipEndpointGkeCluster {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointGkeCluster{
		ResourceLink: dcl.ValueOrEmptyString(o.ResourceLink),
	}
	return p
}

// MembershipEndpointKubernetesMetadataToProto converts a MembershipEndpointKubernetesMetadata resource to its proto representation.
func GkehubBetaMembershipEndpointKubernetesMetadataToProto(o *beta.MembershipEndpointKubernetesMetadata) *betapb.GkehubBetaMembershipEndpointKubernetesMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesMetadata{
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
func GkehubBetaMembershipEndpointKubernetesResourceToProto(o *beta.MembershipEndpointKubernetesResource) *betapb.GkehubBetaMembershipEndpointKubernetesResource {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResource{
		MembershipCrManifest: dcl.ValueOrEmptyString(o.MembershipCrManifest),
		ResourceOptions:      GkehubBetaMembershipEndpointKubernetesResourceResourceOptionsToProto(o.ResourceOptions),
	}
	for _, r := range o.MembershipResources {
		p.MembershipResources = append(p.MembershipResources, GkehubBetaMembershipEndpointKubernetesResourceMembershipResourcesToProto(&r))
	}
	for _, r := range o.ConnectResources {
		p.ConnectResources = append(p.ConnectResources, GkehubBetaMembershipEndpointKubernetesResourceConnectResourcesToProto(&r))
	}
	return p
}

// MembershipEndpointKubernetesResourceMembershipResourcesToProto converts a MembershipEndpointKubernetesResourceMembershipResources resource to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceMembershipResourcesToProto(o *beta.MembershipEndpointKubernetesResourceMembershipResources) *betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources{
		Manifest:      dcl.ValueOrEmptyString(o.Manifest),
		ClusterScoped: dcl.ValueOrEmptyBool(o.ClusterScoped),
	}
	return p
}

// MembershipEndpointKubernetesResourceConnectResourcesToProto converts a MembershipEndpointKubernetesResourceConnectResources resource to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceConnectResourcesToProto(o *beta.MembershipEndpointKubernetesResourceConnectResources) *betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources{
		Manifest:      dcl.ValueOrEmptyString(o.Manifest),
		ClusterScoped: dcl.ValueOrEmptyBool(o.ClusterScoped),
	}
	return p
}

// MembershipEndpointKubernetesResourceResourceOptionsToProto converts a MembershipEndpointKubernetesResourceResourceOptions resource to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceResourceOptionsToProto(o *beta.MembershipEndpointKubernetesResourceResourceOptions) *betapb.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions{
		ConnectVersion: dcl.ValueOrEmptyString(o.ConnectVersion),
		V1Beta1Crd:     dcl.ValueOrEmptyBool(o.V1Beta1Crd),
	}
	return p
}

// MembershipStateToProto converts a MembershipState resource to its proto representation.
func GkehubBetaMembershipStateToProto(o *beta.MembershipState) *betapb.GkehubBetaMembershipState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipState{
		Code: GkehubBetaMembershipStateCodeEnumToProto(o.Code),
	}
	return p
}

// MembershipAuthorityToProto converts a MembershipAuthority resource to its proto representation.
func GkehubBetaMembershipAuthorityToProto(o *beta.MembershipAuthority) *betapb.GkehubBetaMembershipAuthority {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipAuthority{
		Issuer:               dcl.ValueOrEmptyString(o.Issuer),
		WorkloadIdentityPool: dcl.ValueOrEmptyString(o.WorkloadIdentityPool),
		IdentityProvider:     dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *beta.Membership) *betapb.GkehubBetaMembership {
	p := &betapb.GkehubBetaMembership{
		Endpoint:           GkehubBetaMembershipEndpointToProto(resource.Endpoint),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		State:              GkehubBetaMembershipStateToProto(resource.State),
		CreateTime:         dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:         dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:         dcl.ValueOrEmptyString(resource.DeleteTime),
		ExternalId:         dcl.ValueOrEmptyString(resource.ExternalId),
		LastConnectionTime: dcl.ValueOrEmptyString(resource.LastConnectionTime),
		UniqueId:           dcl.ValueOrEmptyString(resource.UniqueId),
		Authority:          GkehubBetaMembershipAuthorityToProto(resource.Authority),
		InfrastructureType: GkehubBetaMembershipInfrastructureTypeEnumToProto(resource.InfrastructureType),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) applyMembership(ctx context.Context, c *beta.Client, request *betapb.ApplyGkehubBetaMembershipRequest) (*betapb.GkehubBetaMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// ApplyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyGkehubBetaMembership(ctx context.Context, request *betapb.ApplyGkehubBetaMembershipRequest) (*betapb.GkehubBetaMembership, error) {
	cl, err := createConfigMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteGkehubBetaMembership(ctx context.Context, request *betapb.DeleteGkehubBetaMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListGkehubBetaMembership(ctx context.Context, request *betapb.ListGkehubBetaMembershipRequest) (*betapb.ListGkehubBetaMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkehubBetaMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkehubBetaMembershipResponse{Items: protos}, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
