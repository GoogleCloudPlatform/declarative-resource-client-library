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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for ForwardingRule.
type ForwardingRuleServer struct{}

// ProtoToForwardingRuleIPProtocolEnum converts a ForwardingRuleIPProtocolEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleIPProtocolEnum(e betapb.ComputeBetaForwardingRuleIPProtocolEnum) *beta.ForwardingRuleIPProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleIPProtocolEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleIPProtocolEnum(n[len("ComputeBetaForwardingRuleIPProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleIPVersionEnum converts a ForwardingRuleIPVersionEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleIPVersionEnum(e betapb.ComputeBetaForwardingRuleIPVersionEnum) *beta.ForwardingRuleIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleIPVersionEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleIPVersionEnum(n[len("ComputeBetaForwardingRuleIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleLoadBalancingSchemeEnum converts a ForwardingRuleLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleLoadBalancingSchemeEnum(e betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum) *beta.ForwardingRuleLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleLoadBalancingSchemeEnum(n[len("ComputeBetaForwardingRuleLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilterFilterMatchCriteriaEnum converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(e betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum) *beta.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleNetworkTierEnum converts a ForwardingRuleNetworkTierEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleNetworkTierEnum(e betapb.ComputeBetaForwardingRuleNetworkTierEnum) *beta.ForwardingRuleNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleNetworkTierEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleNetworkTierEnum(n[len("ComputeBetaForwardingRuleNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilter converts a ForwardingRuleMetadataFilter resource from its proto representation.
func ProtoToComputeBetaForwardingRuleMetadataFilter(p *betapb.ComputeBetaForwardingRuleMetadataFilter) *beta.ForwardingRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &beta.ForwardingRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeBetaForwardingRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToForwardingRuleMetadataFilterFilterLabel converts a ForwardingRuleMetadataFilterFilterLabel resource from its proto representation.
func ProtoToComputeBetaForwardingRuleMetadataFilterFilterLabel(p *betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel) *beta.ForwardingRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &beta.ForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToForwardingRule converts a ForwardingRule resource from its proto representation.
func ProtoToForwardingRule(p *betapb.ComputeBetaForwardingRule) *beta.ForwardingRule {
	obj := &beta.ForwardingRule{
		AllPorts:             dcl.Bool(p.AllPorts),
		AllowGlobalAccess:    dcl.Bool(p.AllowGlobalAccess),
		BackendService:       dcl.StringOrNil(p.BackendService),
		CreationTimestamp:    dcl.StringOrNil(p.CreationTimestamp),
		Description:          dcl.StringOrNil(p.Description),
		IPAddress:            dcl.StringOrNil(p.IpAddress),
		IPProtocol:           ProtoToComputeBetaForwardingRuleIPProtocolEnum(p.GetIpProtocol()),
		IPVersion:            ProtoToComputeBetaForwardingRuleIPVersionEnum(p.GetIpVersion()),
		IsMirroringCollector: dcl.Bool(p.IsMirroringCollector),
		LoadBalancingScheme:  ProtoToComputeBetaForwardingRuleLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		Name:                 dcl.StringOrNil(p.Name),
		Network:              dcl.StringOrNil(p.Network),
		NetworkTier:          ProtoToComputeBetaForwardingRuleNetworkTierEnum(p.GetNetworkTier()),
		PortRange:            dcl.StringOrNil(p.PortRange),
		Region:               dcl.StringOrNil(p.Region),
		SelfLink:             dcl.StringOrNil(p.SelfLink),
		ServiceLabel:         dcl.StringOrNil(p.ServiceLabel),
		ServiceName:          dcl.StringOrNil(p.ServiceName),
		Subnetwork:           dcl.StringOrNil(p.Subnetwork),
		Target:               dcl.StringOrNil(p.Target),
		Project:              dcl.StringOrNil(p.Project),
		Location:             dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetMetadataFilter() {
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeBetaForwardingRuleMetadataFilter(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ForwardingRuleIPProtocolEnumToProto converts a ForwardingRuleIPProtocolEnum enum to its proto representation.
func ComputeBetaForwardingRuleIPProtocolEnumToProto(e *beta.ForwardingRuleIPProtocolEnum) betapb.ComputeBetaForwardingRuleIPProtocolEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleIPProtocolEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleIPProtocolEnum_value["ForwardingRuleIPProtocolEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleIPProtocolEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleIPProtocolEnum(0)
}

// ForwardingRuleIPVersionEnumToProto converts a ForwardingRuleIPVersionEnum enum to its proto representation.
func ComputeBetaForwardingRuleIPVersionEnumToProto(e *beta.ForwardingRuleIPVersionEnum) betapb.ComputeBetaForwardingRuleIPVersionEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleIPVersionEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleIPVersionEnum_value["ForwardingRuleIPVersionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleIPVersionEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleIPVersionEnum(0)
}

// ForwardingRuleLoadBalancingSchemeEnumToProto converts a ForwardingRuleLoadBalancingSchemeEnum enum to its proto representation.
func ComputeBetaForwardingRuleLoadBalancingSchemeEnumToProto(e *beta.ForwardingRuleLoadBalancingSchemeEnum) betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum_value["ForwardingRuleLoadBalancingSchemeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum(0)
}

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *beta.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_value["ForwardingRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// ForwardingRuleNetworkTierEnumToProto converts a ForwardingRuleNetworkTierEnum enum to its proto representation.
func ComputeBetaForwardingRuleNetworkTierEnumToProto(e *beta.ForwardingRuleNetworkTierEnum) betapb.ComputeBetaForwardingRuleNetworkTierEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleNetworkTierEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleNetworkTierEnum_value["ForwardingRuleNetworkTierEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleNetworkTierEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleNetworkTierEnum(0)
}

// ForwardingRuleMetadataFilterToProto converts a ForwardingRuleMetadataFilter resource to its proto representation.
func ComputeBetaForwardingRuleMetadataFilterToProto(o *beta.ForwardingRuleMetadataFilter) *betapb.ComputeBetaForwardingRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaForwardingRuleMetadataFilter{
		FilterMatchCriteria: ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria),
	}
	for _, r := range o.FilterLabel {
		p.FilterLabel = append(p.FilterLabel, ComputeBetaForwardingRuleMetadataFilterFilterLabelToProto(&r))
	}
	return p
}

// ForwardingRuleMetadataFilterFilterLabelToProto converts a ForwardingRuleMetadataFilterFilterLabel resource to its proto representation.
func ComputeBetaForwardingRuleMetadataFilterFilterLabelToProto(o *beta.ForwardingRuleMetadataFilterFilterLabel) *betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// ForwardingRuleToProto converts a ForwardingRule resource to its proto representation.
func ForwardingRuleToProto(resource *beta.ForwardingRule) *betapb.ComputeBetaForwardingRule {
	p := &betapb.ComputeBetaForwardingRule{
		AllPorts:             dcl.ValueOrEmptyBool(resource.AllPorts),
		AllowGlobalAccess:    dcl.ValueOrEmptyBool(resource.AllowGlobalAccess),
		BackendService:       dcl.ValueOrEmptyString(resource.BackendService),
		CreationTimestamp:    dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		IpAddress:            dcl.ValueOrEmptyString(resource.IPAddress),
		IpProtocol:           ComputeBetaForwardingRuleIPProtocolEnumToProto(resource.IPProtocol),
		IpVersion:            ComputeBetaForwardingRuleIPVersionEnumToProto(resource.IPVersion),
		IsMirroringCollector: dcl.ValueOrEmptyBool(resource.IsMirroringCollector),
		LoadBalancingScheme:  ComputeBetaForwardingRuleLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme),
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Network:              dcl.ValueOrEmptyString(resource.Network),
		NetworkTier:          ComputeBetaForwardingRuleNetworkTierEnumToProto(resource.NetworkTier),
		PortRange:            dcl.ValueOrEmptyString(resource.PortRange),
		Region:               dcl.ValueOrEmptyString(resource.Region),
		SelfLink:             dcl.ValueOrEmptyString(resource.SelfLink),
		ServiceLabel:         dcl.ValueOrEmptyString(resource.ServiceLabel),
		ServiceName:          dcl.ValueOrEmptyString(resource.ServiceName),
		Subnetwork:           dcl.ValueOrEmptyString(resource.Subnetwork),
		Target:               dcl.ValueOrEmptyString(resource.Target),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Location:             dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.MetadataFilter {
		p.MetadataFilter = append(p.MetadataFilter, ComputeBetaForwardingRuleMetadataFilterToProto(&r))
	}
	for _, r := range resource.Ports {
		p.Ports = append(p.Ports, r)
	}

	return p
}

// ApplyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) applyForwardingRule(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaForwardingRuleRequest) (*betapb.ComputeBetaForwardingRule, error) {
	p := ProtoToForwardingRule(request.GetResource())
	res, err := c.ApplyForwardingRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ForwardingRuleToProto(res)
	return r, nil
}

// ApplyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) ApplyComputeBetaForwardingRule(ctx context.Context, request *betapb.ApplyComputeBetaForwardingRuleRequest) (*betapb.ComputeBetaForwardingRule, error) {
	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyForwardingRule(ctx, cl, request)
}

// DeleteForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Delete() method.
func (s *ForwardingRuleServer) DeleteComputeBetaForwardingRule(ctx context.Context, request *betapb.DeleteComputeBetaForwardingRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))

}

// ListComputeBetaForwardingRule handles the gRPC request by passing it to the underlying ForwardingRuleList() method.
func (s *ForwardingRuleServer) ListComputeBetaForwardingRule(ctx context.Context, request *betapb.ListComputeBetaForwardingRuleRequest) (*betapb.ListComputeBetaForwardingRuleResponse, error) {
	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaForwardingRule
	for _, r := range resources.Items {
		rp := ForwardingRuleToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaForwardingRuleResponse{Items: protos}, nil
}

func createConfigForwardingRule(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
