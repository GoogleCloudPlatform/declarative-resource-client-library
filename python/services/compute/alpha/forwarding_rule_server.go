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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// Server implements the gRPC interface for ForwardingRule.
type ForwardingRuleServer struct{}

// ProtoToForwardingRuleIPProtocolEnum converts a ForwardingRuleIPProtocolEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleIPProtocolEnum(e alphapb.ComputeAlphaForwardingRuleIPProtocolEnum) *alpha.ForwardingRuleIPProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleIPProtocolEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleIPProtocolEnum(n[len("ComputeAlphaForwardingRuleIPProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleIPVersionEnum converts a ForwardingRuleIPVersionEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleIPVersionEnum(e alphapb.ComputeAlphaForwardingRuleIPVersionEnum) *alpha.ForwardingRuleIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleIPVersionEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleIPVersionEnum(n[len("ComputeAlphaForwardingRuleIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleLoadBalancingSchemeEnum converts a ForwardingRuleLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleLoadBalancingSchemeEnum(e alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum) *alpha.ForwardingRuleLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleLoadBalancingSchemeEnum(n[len("ComputeAlphaForwardingRuleLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilterFilterMatchCriteriaEnum converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(e alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum) *alpha.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleNetworkTierEnum converts a ForwardingRuleNetworkTierEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleNetworkTierEnum(e alphapb.ComputeAlphaForwardingRuleNetworkTierEnum) *alpha.ForwardingRuleNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleNetworkTierEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleNetworkTierEnum(n[len("ComputeAlphaForwardingRuleNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilter converts a ForwardingRuleMetadataFilter resource from its proto representation.
func ProtoToComputeAlphaForwardingRuleMetadataFilter(p *alphapb.ComputeAlphaForwardingRuleMetadataFilter) *alpha.ForwardingRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.ForwardingRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeAlphaForwardingRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToForwardingRuleMetadataFilterFilterLabel converts a ForwardingRuleMetadataFilterFilterLabel resource from its proto representation.
func ProtoToComputeAlphaForwardingRuleMetadataFilterFilterLabel(p *alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel) *alpha.ForwardingRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &alpha.ForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToForwardingRule converts a ForwardingRule resource from its proto representation.
func ProtoToForwardingRule(p *alphapb.ComputeAlphaForwardingRule) *alpha.ForwardingRule {
	obj := &alpha.ForwardingRule{
		AllPorts:             dcl.Bool(p.AllPorts),
		AllowGlobalAccess:    dcl.Bool(p.AllowGlobalAccess),
		LabelFingerprint:     dcl.StringOrNil(p.LabelFingerprint),
		BackendService:       dcl.StringOrNil(p.BackendService),
		CreationTimestamp:    dcl.StringOrNil(p.CreationTimestamp),
		Description:          dcl.StringOrNil(p.Description),
		IPAddress:            dcl.StringOrNil(p.IpAddress),
		IPProtocol:           ProtoToComputeAlphaForwardingRuleIPProtocolEnum(p.GetIpProtocol()),
		IPVersion:            ProtoToComputeAlphaForwardingRuleIPVersionEnum(p.GetIpVersion()),
		IsMirroringCollector: dcl.Bool(p.IsMirroringCollector),
		LoadBalancingScheme:  ProtoToComputeAlphaForwardingRuleLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		Name:                 dcl.StringOrNil(p.Name),
		Network:              dcl.StringOrNil(p.Network),
		NetworkTier:          ProtoToComputeAlphaForwardingRuleNetworkTierEnum(p.GetNetworkTier()),
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
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeAlphaForwardingRuleMetadataFilter(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ForwardingRuleIPProtocolEnumToProto converts a ForwardingRuleIPProtocolEnum enum to its proto representation.
func ComputeAlphaForwardingRuleIPProtocolEnumToProto(e *alpha.ForwardingRuleIPProtocolEnum) alphapb.ComputeAlphaForwardingRuleIPProtocolEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleIPProtocolEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleIPProtocolEnum_value["ForwardingRuleIPProtocolEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleIPProtocolEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleIPProtocolEnum(0)
}

// ForwardingRuleIPVersionEnumToProto converts a ForwardingRuleIPVersionEnum enum to its proto representation.
func ComputeAlphaForwardingRuleIPVersionEnumToProto(e *alpha.ForwardingRuleIPVersionEnum) alphapb.ComputeAlphaForwardingRuleIPVersionEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleIPVersionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleIPVersionEnum_value["ForwardingRuleIPVersionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleIPVersionEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleIPVersionEnum(0)
}

// ForwardingRuleLoadBalancingSchemeEnumToProto converts a ForwardingRuleLoadBalancingSchemeEnum enum to its proto representation.
func ComputeAlphaForwardingRuleLoadBalancingSchemeEnumToProto(e *alpha.ForwardingRuleLoadBalancingSchemeEnum) alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum_value["ForwardingRuleLoadBalancingSchemeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum(0)
}

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *alpha.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_value["ForwardingRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// ForwardingRuleNetworkTierEnumToProto converts a ForwardingRuleNetworkTierEnum enum to its proto representation.
func ComputeAlphaForwardingRuleNetworkTierEnumToProto(e *alpha.ForwardingRuleNetworkTierEnum) alphapb.ComputeAlphaForwardingRuleNetworkTierEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleNetworkTierEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleNetworkTierEnum_value["ForwardingRuleNetworkTierEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleNetworkTierEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleNetworkTierEnum(0)
}

// ForwardingRuleMetadataFilterToProto converts a ForwardingRuleMetadataFilter resource to its proto representation.
func ComputeAlphaForwardingRuleMetadataFilterToProto(o *alpha.ForwardingRuleMetadataFilter) *alphapb.ComputeAlphaForwardingRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaForwardingRuleMetadataFilter{
		FilterMatchCriteria: ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria),
	}
	for _, r := range o.FilterLabel {
		p.FilterLabel = append(p.FilterLabel, ComputeAlphaForwardingRuleMetadataFilterFilterLabelToProto(&r))
	}
	return p
}

// ForwardingRuleMetadataFilterFilterLabelToProto converts a ForwardingRuleMetadataFilterFilterLabel resource to its proto representation.
func ComputeAlphaForwardingRuleMetadataFilterFilterLabelToProto(o *alpha.ForwardingRuleMetadataFilterFilterLabel) *alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// ForwardingRuleToProto converts a ForwardingRule resource to its proto representation.
func ForwardingRuleToProto(resource *alpha.ForwardingRule) *alphapb.ComputeAlphaForwardingRule {
	p := &alphapb.ComputeAlphaForwardingRule{
		AllPorts:             dcl.ValueOrEmptyBool(resource.AllPorts),
		AllowGlobalAccess:    dcl.ValueOrEmptyBool(resource.AllowGlobalAccess),
		LabelFingerprint:     dcl.ValueOrEmptyString(resource.LabelFingerprint),
		BackendService:       dcl.ValueOrEmptyString(resource.BackendService),
		CreationTimestamp:    dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		IpAddress:            dcl.ValueOrEmptyString(resource.IPAddress),
		IpProtocol:           ComputeAlphaForwardingRuleIPProtocolEnumToProto(resource.IPProtocol),
		IpVersion:            ComputeAlphaForwardingRuleIPVersionEnumToProto(resource.IPVersion),
		IsMirroringCollector: dcl.ValueOrEmptyBool(resource.IsMirroringCollector),
		LoadBalancingScheme:  ComputeAlphaForwardingRuleLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme),
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Network:              dcl.ValueOrEmptyString(resource.Network),
		NetworkTier:          ComputeAlphaForwardingRuleNetworkTierEnumToProto(resource.NetworkTier),
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
		p.MetadataFilter = append(p.MetadataFilter, ComputeAlphaForwardingRuleMetadataFilterToProto(&r))
	}
	for _, r := range resource.Ports {
		p.Ports = append(p.Ports, r)
	}

	return p
}

// ApplyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) applyForwardingRule(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaForwardingRuleRequest) (*alphapb.ComputeAlphaForwardingRule, error) {
	p := ProtoToForwardingRule(request.GetResource())
	res, err := c.ApplyForwardingRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ForwardingRuleToProto(res)
	return r, nil
}

// ApplyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) ApplyComputeAlphaForwardingRule(ctx context.Context, request *alphapb.ApplyComputeAlphaForwardingRuleRequest) (*alphapb.ComputeAlphaForwardingRule, error) {
	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyForwardingRule(ctx, cl, request)
}

// DeleteForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Delete() method.
func (s *ForwardingRuleServer) DeleteComputeAlphaForwardingRule(ctx context.Context, request *alphapb.DeleteComputeAlphaForwardingRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))

}

// ListComputeAlphaForwardingRule handles the gRPC request by passing it to the underlying ForwardingRuleList() method.
func (s *ForwardingRuleServer) ListComputeAlphaForwardingRule(ctx context.Context, request *alphapb.ListComputeAlphaForwardingRuleRequest) (*alphapb.ListComputeAlphaForwardingRuleResponse, error) {
	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaForwardingRule
	for _, r := range resources.Items {
		rp := ForwardingRuleToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListComputeAlphaForwardingRuleResponse{Items: protos}, nil
}

func createConfigForwardingRule(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
