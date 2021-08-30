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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for ForwardingRule.
type ForwardingRuleServer struct{}

// ProtoToForwardingRuleIPProtocolEnum converts a ForwardingRuleIPProtocolEnum enum from its proto representation.
func ProtoToComputeForwardingRuleIPProtocolEnum(e computepb.ComputeForwardingRuleIPProtocolEnum) *compute.ForwardingRuleIPProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleIPProtocolEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleIPProtocolEnum(n[len("ComputeForwardingRuleIPProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleIPVersionEnum converts a ForwardingRuleIPVersionEnum enum from its proto representation.
func ProtoToComputeForwardingRuleIPVersionEnum(e computepb.ComputeForwardingRuleIPVersionEnum) *compute.ForwardingRuleIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleIPVersionEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleIPVersionEnum(n[len("ComputeForwardingRuleIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleLoadBalancingSchemeEnum converts a ForwardingRuleLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeForwardingRuleLoadBalancingSchemeEnum(e computepb.ComputeForwardingRuleLoadBalancingSchemeEnum) *compute.ForwardingRuleLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleLoadBalancingSchemeEnum(n[len("ComputeForwardingRuleLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilterFilterMatchCriteriaEnum converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(e computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum) *compute.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleNetworkTierEnum converts a ForwardingRuleNetworkTierEnum enum from its proto representation.
func ProtoToComputeForwardingRuleNetworkTierEnum(e computepb.ComputeForwardingRuleNetworkTierEnum) *compute.ForwardingRuleNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleNetworkTierEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleNetworkTierEnum(n[len("ComputeForwardingRuleNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilter converts a ForwardingRuleMetadataFilter resource from its proto representation.
func ProtoToComputeForwardingRuleMetadataFilter(p *computepb.ComputeForwardingRuleMetadataFilter) *compute.ForwardingRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &compute.ForwardingRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeForwardingRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToForwardingRuleMetadataFilterFilterLabel converts a ForwardingRuleMetadataFilterFilterLabel resource from its proto representation.
func ProtoToComputeForwardingRuleMetadataFilterFilterLabel(p *computepb.ComputeForwardingRuleMetadataFilterFilterLabel) *compute.ForwardingRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &compute.ForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToForwardingRule converts a ForwardingRule resource from its proto representation.
func ProtoToForwardingRule(p *computepb.ComputeForwardingRule) *compute.ForwardingRule {
	obj := &compute.ForwardingRule{
		AllPorts:             dcl.Bool(p.AllPorts),
		AllowGlobalAccess:    dcl.Bool(p.AllowGlobalAccess),
		LabelFingerprint:     dcl.StringOrNil(p.LabelFingerprint),
		BackendService:       dcl.StringOrNil(p.BackendService),
		CreationTimestamp:    dcl.StringOrNil(p.CreationTimestamp),
		Description:          dcl.StringOrNil(p.Description),
		IPAddress:            dcl.StringOrNil(p.IpAddress),
		IPProtocol:           ProtoToComputeForwardingRuleIPProtocolEnum(p.GetIpProtocol()),
		IPVersion:            ProtoToComputeForwardingRuleIPVersionEnum(p.GetIpVersion()),
		IsMirroringCollector: dcl.Bool(p.IsMirroringCollector),
		LoadBalancingScheme:  ProtoToComputeForwardingRuleLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		Name:                 dcl.StringOrNil(p.Name),
		Network:              dcl.StringOrNil(p.Network),
		NetworkTier:          ProtoToComputeForwardingRuleNetworkTierEnum(p.GetNetworkTier()),
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
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeForwardingRuleMetadataFilter(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ForwardingRuleIPProtocolEnumToProto converts a ForwardingRuleIPProtocolEnum enum to its proto representation.
func ComputeForwardingRuleIPProtocolEnumToProto(e *compute.ForwardingRuleIPProtocolEnum) computepb.ComputeForwardingRuleIPProtocolEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleIPProtocolEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleIPProtocolEnum_value["ForwardingRuleIPProtocolEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleIPProtocolEnum(v)
	}
	return computepb.ComputeForwardingRuleIPProtocolEnum(0)
}

// ForwardingRuleIPVersionEnumToProto converts a ForwardingRuleIPVersionEnum enum to its proto representation.
func ComputeForwardingRuleIPVersionEnumToProto(e *compute.ForwardingRuleIPVersionEnum) computepb.ComputeForwardingRuleIPVersionEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleIPVersionEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleIPVersionEnum_value["ForwardingRuleIPVersionEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleIPVersionEnum(v)
	}
	return computepb.ComputeForwardingRuleIPVersionEnum(0)
}

// ForwardingRuleLoadBalancingSchemeEnumToProto converts a ForwardingRuleLoadBalancingSchemeEnum enum to its proto representation.
func ComputeForwardingRuleLoadBalancingSchemeEnumToProto(e *compute.ForwardingRuleLoadBalancingSchemeEnum) computepb.ComputeForwardingRuleLoadBalancingSchemeEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleLoadBalancingSchemeEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleLoadBalancingSchemeEnum_value["ForwardingRuleLoadBalancingSchemeEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleLoadBalancingSchemeEnum(v)
	}
	return computepb.ComputeForwardingRuleLoadBalancingSchemeEnum(0)
}

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *compute.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum_value["ForwardingRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// ForwardingRuleNetworkTierEnumToProto converts a ForwardingRuleNetworkTierEnum enum to its proto representation.
func ComputeForwardingRuleNetworkTierEnumToProto(e *compute.ForwardingRuleNetworkTierEnum) computepb.ComputeForwardingRuleNetworkTierEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleNetworkTierEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleNetworkTierEnum_value["ForwardingRuleNetworkTierEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleNetworkTierEnum(v)
	}
	return computepb.ComputeForwardingRuleNetworkTierEnum(0)
}

// ForwardingRuleMetadataFilterToProto converts a ForwardingRuleMetadataFilter resource to its proto representation.
func ComputeForwardingRuleMetadataFilterToProto(o *compute.ForwardingRuleMetadataFilter) *computepb.ComputeForwardingRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeForwardingRuleMetadataFilter{
		FilterMatchCriteria: ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria),
	}
	for _, r := range o.FilterLabel {
		p.FilterLabel = append(p.FilterLabel, ComputeForwardingRuleMetadataFilterFilterLabelToProto(&r))
	}
	return p
}

// ForwardingRuleMetadataFilterFilterLabelToProto converts a ForwardingRuleMetadataFilterFilterLabel resource to its proto representation.
func ComputeForwardingRuleMetadataFilterFilterLabelToProto(o *compute.ForwardingRuleMetadataFilterFilterLabel) *computepb.ComputeForwardingRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// ForwardingRuleToProto converts a ForwardingRule resource to its proto representation.
func ForwardingRuleToProto(resource *compute.ForwardingRule) *computepb.ComputeForwardingRule {
	p := &computepb.ComputeForwardingRule{
		AllPorts:             dcl.ValueOrEmptyBool(resource.AllPorts),
		AllowGlobalAccess:    dcl.ValueOrEmptyBool(resource.AllowGlobalAccess),
		LabelFingerprint:     dcl.ValueOrEmptyString(resource.LabelFingerprint),
		BackendService:       dcl.ValueOrEmptyString(resource.BackendService),
		CreationTimestamp:    dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		IpAddress:            dcl.ValueOrEmptyString(resource.IPAddress),
		IpProtocol:           ComputeForwardingRuleIPProtocolEnumToProto(resource.IPProtocol),
		IpVersion:            ComputeForwardingRuleIPVersionEnumToProto(resource.IPVersion),
		IsMirroringCollector: dcl.ValueOrEmptyBool(resource.IsMirroringCollector),
		LoadBalancingScheme:  ComputeForwardingRuleLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme),
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Network:              dcl.ValueOrEmptyString(resource.Network),
		NetworkTier:          ComputeForwardingRuleNetworkTierEnumToProto(resource.NetworkTier),
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
		p.MetadataFilter = append(p.MetadataFilter, ComputeForwardingRuleMetadataFilterToProto(&r))
	}
	for _, r := range resource.Ports {
		p.Ports = append(p.Ports, r)
	}

	return p
}

// ApplyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) applyForwardingRule(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeForwardingRuleRequest) (*computepb.ComputeForwardingRule, error) {
	p := ProtoToForwardingRule(request.GetResource())
	res, err := c.ApplyForwardingRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ForwardingRuleToProto(res)
	return r, nil
}

// ApplyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) ApplyComputeForwardingRule(ctx context.Context, request *computepb.ApplyComputeForwardingRuleRequest) (*computepb.ComputeForwardingRule, error) {
	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyForwardingRule(ctx, cl, request)
}

// DeleteForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Delete() method.
func (s *ForwardingRuleServer) DeleteComputeForwardingRule(ctx context.Context, request *computepb.DeleteComputeForwardingRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))

}

// ListComputeForwardingRule handles the gRPC request by passing it to the underlying ForwardingRuleList() method.
func (s *ForwardingRuleServer) ListComputeForwardingRule(ctx context.Context, request *computepb.ListComputeForwardingRuleRequest) (*computepb.ListComputeForwardingRuleResponse, error) {
	cl, err := createConfigForwardingRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeForwardingRule
	for _, r := range resources.Items {
		rp := ForwardingRuleToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeForwardingRuleResponse{Items: protos}, nil
}

func createConfigForwardingRule(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
