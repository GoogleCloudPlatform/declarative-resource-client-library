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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/alpha/networkservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
)

// Server implements the gRPC interface for EndpointPolicy.
type EndpointPolicyServer struct{}

// ProtoToEndpointPolicyTypeEnum converts a EndpointPolicyTypeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaEndpointPolicyTypeEnum(e alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum) *alpha.EndpointPolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum_name[int32(e)]; ok {
		e := alpha.EndpointPolicyTypeEnum(n[len("NetworkservicesAlphaEndpointPolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(e alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) *alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_name[int32(e)]; ok {
		e := alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(n[len("NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointPolicyEndpointMatcher converts a EndpointPolicyEndpointMatcher resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcher(p *alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcher) *alpha.EndpointPolicyEndpointMatcher {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointPolicyEndpointMatcher{
		MetadataLabelMatcher: ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcher(p.GetMetadataLabelMatcher()),
	}
	return obj
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcher converts a EndpointPolicyEndpointMatcherMetadataLabelMatcher resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcher(p *alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcher) *alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(p.GetMetadataLabelMatchCriteria()),
	}
	for _, r := range p.GetMetadataLabels() {
		obj.MetadataLabels = append(obj.MetadataLabels, *ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(r))
	}
	return obj
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(p *alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.StringOrNil(p.LabelName),
		LabelValue: dcl.StringOrNil(p.LabelValue),
	}
	return obj
}

// ProtoToEndpointPolicyTrafficPortSelector converts a EndpointPolicyTrafficPortSelector resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointPolicyTrafficPortSelector(p *alphapb.NetworkservicesAlphaEndpointPolicyTrafficPortSelector) *alpha.EndpointPolicyTrafficPortSelector {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointPolicyTrafficPortSelector{}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToEndpointPolicy converts a EndpointPolicy resource from its proto representation.
func ProtoToEndpointPolicy(p *alphapb.NetworkservicesAlphaEndpointPolicy) *alpha.EndpointPolicy {
	obj := &alpha.EndpointPolicy{
		Name:                dcl.StringOrNil(p.Name),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		Type:                ProtoToNetworkservicesAlphaEndpointPolicyTypeEnum(p.GetType()),
		AuthorizationPolicy: dcl.StringOrNil(p.AuthorizationPolicy),
		EndpointMatcher:     ProtoToNetworkservicesAlphaEndpointPolicyEndpointMatcher(p.GetEndpointMatcher()),
		TrafficPortSelector: ProtoToNetworkservicesAlphaEndpointPolicyTrafficPortSelector(p.GetTrafficPortSelector()),
		Description:         dcl.StringOrNil(p.Description),
		ServerTlsPolicy:     dcl.StringOrNil(p.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.StringOrNil(p.ClientTlsPolicy),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// EndpointPolicyTypeEnumToProto converts a EndpointPolicyTypeEnum enum to its proto representation.
func NetworkservicesAlphaEndpointPolicyTypeEnumToProto(e *alpha.EndpointPolicyTypeEnum) alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum_value["EndpointPolicyTypeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum(v)
	}
	return alphapb.NetworkservicesAlphaEndpointPolicyTypeEnum(0)
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum to its proto representation.
func NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(e *alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_value["EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(v)
	}
	return alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
}

// EndpointPolicyEndpointMatcherToProto converts a EndpointPolicyEndpointMatcher resource to its proto representation.
func NetworkservicesAlphaEndpointPolicyEndpointMatcherToProto(o *alpha.EndpointPolicyEndpointMatcher) *alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcher {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcher{
		MetadataLabelMatcher: NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherToProto(o.MetadataLabelMatcher),
	}
	return p
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcher resource to its proto representation.
func NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherToProto(o *alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcher) *alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(o.MetadataLabelMatchCriteria),
	}
	for _, r := range o.MetadataLabels {
		p.MetadataLabels = append(p.MetadataLabels, NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(&r))
	}
	return p
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels resource to its proto representation.
func NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(o *alpha.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.ValueOrEmptyString(o.LabelName),
		LabelValue: dcl.ValueOrEmptyString(o.LabelValue),
	}
	return p
}

// EndpointPolicyTrafficPortSelectorToProto converts a EndpointPolicyTrafficPortSelector resource to its proto representation.
func NetworkservicesAlphaEndpointPolicyTrafficPortSelectorToProto(o *alpha.EndpointPolicyTrafficPortSelector) *alphapb.NetworkservicesAlphaEndpointPolicyTrafficPortSelector {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointPolicyTrafficPortSelector{}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// EndpointPolicyToProto converts a EndpointPolicy resource to its proto representation.
func EndpointPolicyToProto(resource *alpha.EndpointPolicy) *alphapb.NetworkservicesAlphaEndpointPolicy {
	p := &alphapb.NetworkservicesAlphaEndpointPolicy{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		CreateTime:          dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:          dcl.ValueOrEmptyString(resource.UpdateTime),
		Type:                NetworkservicesAlphaEndpointPolicyTypeEnumToProto(resource.Type),
		AuthorizationPolicy: dcl.ValueOrEmptyString(resource.AuthorizationPolicy),
		EndpointMatcher:     NetworkservicesAlphaEndpointPolicyEndpointMatcherToProto(resource.EndpointMatcher),
		TrafficPortSelector: NetworkservicesAlphaEndpointPolicyTrafficPortSelectorToProto(resource.TrafficPortSelector),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		ServerTlsPolicy:     dcl.ValueOrEmptyString(resource.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.ValueOrEmptyString(resource.ClientTlsPolicy),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Apply() method.
func (s *EndpointPolicyServer) applyEndpointPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaEndpointPolicyRequest) (*alphapb.NetworkservicesAlphaEndpointPolicy, error) {
	p := ProtoToEndpointPolicy(request.GetResource())
	res, err := c.ApplyEndpointPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointPolicyToProto(res)
	return r, nil
}

// ApplyEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Apply() method.
func (s *EndpointPolicyServer) ApplyNetworkservicesAlphaEndpointPolicy(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaEndpointPolicyRequest) (*alphapb.NetworkservicesAlphaEndpointPolicy, error) {
	cl, err := createConfigEndpointPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEndpointPolicy(ctx, cl, request)
}

// DeleteEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Delete() method.
func (s *EndpointPolicyServer) DeleteNetworkservicesAlphaEndpointPolicy(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaEndpointPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpointPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpointPolicy(ctx, ProtoToEndpointPolicy(request.GetResource()))

}

// ListNetworkservicesAlphaEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicyList() method.
func (s *EndpointPolicyServer) ListNetworkservicesAlphaEndpointPolicy(ctx context.Context, request *alphapb.ListNetworkservicesAlphaEndpointPolicyRequest) (*alphapb.ListNetworkservicesAlphaEndpointPolicyResponse, error) {
	cl, err := createConfigEndpointPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpointPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaEndpointPolicy
	for _, r := range resources.Items {
		rp := EndpointPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListNetworkservicesAlphaEndpointPolicyResponse{Items: protos}, nil
}

func createConfigEndpointPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
