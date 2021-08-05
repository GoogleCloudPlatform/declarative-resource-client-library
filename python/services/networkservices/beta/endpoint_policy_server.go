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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/beta/networkservices_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
)

// Server implements the gRPC interface for EndpointPolicy.
type EndpointPolicyServer struct{}

// ProtoToEndpointPolicyTypeEnum converts a EndpointPolicyTypeEnum enum from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyTypeEnum(e betapb.NetworkservicesBetaEndpointPolicyTypeEnum) *beta.EndpointPolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaEndpointPolicyTypeEnum_name[int32(e)]; ok {
		e := beta.EndpointPolicyTypeEnum(n[len("NetworkservicesBetaEndpointPolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(e betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_name[int32(e)]; ok {
		e := beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(n[len("NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointPolicyEndpointMatcher converts a EndpointPolicyEndpointMatcher resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcher(p *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcher) *beta.EndpointPolicyEndpointMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyEndpointMatcher{
		MetadataLabelMatcher: ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher(p.GetMetadataLabelMatcher()),
	}
	return obj
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcher converts a EndpointPolicyEndpointMatcherMetadataLabelMatcher resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher(p *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher) *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(p.GetMetadataLabelMatchCriteria()),
	}
	for _, r := range p.GetMetadataLabels() {
		obj.MetadataLabels = append(obj.MetadataLabels, *ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(r))
	}
	return obj
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(p *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.StringOrNil(p.LabelName),
		LabelValue: dcl.StringOrNil(p.LabelValue),
	}
	return obj
}

// ProtoToEndpointPolicyTrafficPortSelector converts a EndpointPolicyTrafficPortSelector resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyTrafficPortSelector(p *betapb.NetworkservicesBetaEndpointPolicyTrafficPortSelector) *beta.EndpointPolicyTrafficPortSelector {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyTrafficPortSelector{}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToEndpointPolicy converts a EndpointPolicy resource from its proto representation.
func ProtoToEndpointPolicy(p *betapb.NetworkservicesBetaEndpointPolicy) *beta.EndpointPolicy {
	obj := &beta.EndpointPolicy{
		Name:                dcl.StringOrNil(p.Name),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		Type:                ProtoToNetworkservicesBetaEndpointPolicyTypeEnum(p.GetType()),
		AuthorizationPolicy: dcl.StringOrNil(p.AuthorizationPolicy),
		EndpointMatcher:     ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcher(p.GetEndpointMatcher()),
		TrafficPortSelector: ProtoToNetworkservicesBetaEndpointPolicyTrafficPortSelector(p.GetTrafficPortSelector()),
		Description:         dcl.StringOrNil(p.Description),
		ServerTlsPolicy:     dcl.StringOrNil(p.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.StringOrNil(p.ClientTlsPolicy),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// EndpointPolicyTypeEnumToProto converts a EndpointPolicyTypeEnum enum to its proto representation.
func NetworkservicesBetaEndpointPolicyTypeEnumToProto(e *beta.EndpointPolicyTypeEnum) betapb.NetworkservicesBetaEndpointPolicyTypeEnum {
	if e == nil {
		return betapb.NetworkservicesBetaEndpointPolicyTypeEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaEndpointPolicyTypeEnum_value["EndpointPolicyTypeEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaEndpointPolicyTypeEnum(v)
	}
	return betapb.NetworkservicesBetaEndpointPolicyTypeEnum(0)
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(e *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == nil {
		return betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_value["EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(v)
	}
	return betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
}

// EndpointPolicyEndpointMatcherToProto converts a EndpointPolicyEndpointMatcher resource to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherToProto(o *beta.EndpointPolicyEndpointMatcher) *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyEndpointMatcher{
		MetadataLabelMatcher: NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherToProto(o.MetadataLabelMatcher),
	}
	return p
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcher resource to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherToProto(o *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcher) *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(o.MetadataLabelMatchCriteria),
	}
	for _, r := range o.MetadataLabels {
		p.MetadataLabels = append(p.MetadataLabels, NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(&r))
	}
	return p
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels resource to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(o *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.ValueOrEmptyString(o.LabelName),
		LabelValue: dcl.ValueOrEmptyString(o.LabelValue),
	}
	return p
}

// EndpointPolicyTrafficPortSelectorToProto converts a EndpointPolicyTrafficPortSelector resource to its proto representation.
func NetworkservicesBetaEndpointPolicyTrafficPortSelectorToProto(o *beta.EndpointPolicyTrafficPortSelector) *betapb.NetworkservicesBetaEndpointPolicyTrafficPortSelector {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyTrafficPortSelector{}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// EndpointPolicyToProto converts a EndpointPolicy resource to its proto representation.
func EndpointPolicyToProto(resource *beta.EndpointPolicy) *betapb.NetworkservicesBetaEndpointPolicy {
	p := &betapb.NetworkservicesBetaEndpointPolicy{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		CreateTime:          dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:          dcl.ValueOrEmptyString(resource.UpdateTime),
		Type:                NetworkservicesBetaEndpointPolicyTypeEnumToProto(resource.Type),
		AuthorizationPolicy: dcl.ValueOrEmptyString(resource.AuthorizationPolicy),
		EndpointMatcher:     NetworkservicesBetaEndpointPolicyEndpointMatcherToProto(resource.EndpointMatcher),
		TrafficPortSelector: NetworkservicesBetaEndpointPolicyTrafficPortSelectorToProto(resource.TrafficPortSelector),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		ServerTlsPolicy:     dcl.ValueOrEmptyString(resource.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.ValueOrEmptyString(resource.ClientTlsPolicy),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Apply() method.
func (s *EndpointPolicyServer) applyEndpointPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaEndpointPolicyRequest) (*betapb.NetworkservicesBetaEndpointPolicy, error) {
	p := ProtoToEndpointPolicy(request.GetResource())
	res, err := c.ApplyEndpointPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointPolicyToProto(res)
	return r, nil
}

// ApplyEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Apply() method.
func (s *EndpointPolicyServer) ApplyNetworkservicesBetaEndpointPolicy(ctx context.Context, request *betapb.ApplyNetworkservicesBetaEndpointPolicyRequest) (*betapb.NetworkservicesBetaEndpointPolicy, error) {
	cl, err := createConfigEndpointPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEndpointPolicy(ctx, cl, request)
}

// DeleteEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Delete() method.
func (s *EndpointPolicyServer) DeleteNetworkservicesBetaEndpointPolicy(ctx context.Context, request *betapb.DeleteNetworkservicesBetaEndpointPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpointPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpointPolicy(ctx, ProtoToEndpointPolicy(request.GetResource()))

}

// ListNetworkservicesBetaEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicyList() method.
func (s *EndpointPolicyServer) ListNetworkservicesBetaEndpointPolicy(ctx context.Context, request *betapb.ListNetworkservicesBetaEndpointPolicyRequest) (*betapb.ListNetworkservicesBetaEndpointPolicyResponse, error) {
	cl, err := createConfigEndpointPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpointPolicy(ctx, ProtoToEndpointPolicy(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaEndpointPolicy
	for _, r := range resources.Items {
		rp := EndpointPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListNetworkservicesBetaEndpointPolicyResponse{Items: protos}, nil
}

func createConfigEndpointPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
