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

// Server implements the gRPC interface for ServiceAttachment.
type ServiceAttachmentServer struct{}

// ProtoToServiceAttachmentConnectionPreferenceEnum converts a ServiceAttachmentConnectionPreferenceEnum enum from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConnectionPreferenceEnum(e alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum) *alpha.ServiceAttachmentConnectionPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum_name[int32(e)]; ok {
		e := alpha.ServiceAttachmentConnectionPreferenceEnum(n[len("ComputeAlphaServiceAttachmentConnectionPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpointsStatusEnum converts a ServiceAttachmentConnectedEndpointsStatusEnum enum from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(e alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum) *alpha.ServiceAttachmentConnectedEndpointsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum_name[int32(e)]; ok {
		e := alpha.ServiceAttachmentConnectedEndpointsStatusEnum(n[len("ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpoints converts a ServiceAttachmentConnectedEndpoints resource from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConnectedEndpoints(p *alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints) *alpha.ServiceAttachmentConnectedEndpoints {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAttachmentConnectedEndpoints{
		Status:          ProtoToComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(p.GetStatus()),
		PscConnectionId: dcl.Int64OrNil(p.PscConnectionId),
		Endpoint:        dcl.StringOrNil(p.Endpoint),
	}
	return obj
}

// ProtoToServiceAttachmentConsumerAcceptLists converts a ServiceAttachmentConsumerAcceptLists resource from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConsumerAcceptLists(p *alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists) *alpha.ServiceAttachmentConsumerAcceptLists {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.StringOrNil(p.ProjectIdOrNum),
		ConnectionLimit: dcl.Int64OrNil(p.ConnectionLimit),
	}
	return obj
}

// ProtoToServiceAttachmentPscServiceAttachmentId converts a ServiceAttachmentPscServiceAttachmentId resource from its proto representation.
func ProtoToComputeAlphaServiceAttachmentPscServiceAttachmentId(p *alphapb.ComputeAlphaServiceAttachmentPscServiceAttachmentId) *alpha.ServiceAttachmentPscServiceAttachmentId {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAttachmentPscServiceAttachmentId{
		High: dcl.Int64OrNil(p.High),
		Low:  dcl.Int64OrNil(p.Low),
	}
	return obj
}

// ProtoToServiceAttachment converts a ServiceAttachment resource from its proto representation.
func ProtoToServiceAttachment(p *alphapb.ComputeAlphaServiceAttachment) *alpha.ServiceAttachment {
	obj := &alpha.ServiceAttachment{
		Id:                     dcl.Int64OrNil(p.Id),
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		SelfLink:               dcl.StringOrNil(p.SelfLink),
		Region:                 dcl.StringOrNil(p.Region),
		TargetService:          dcl.StringOrNil(p.TargetService),
		ConnectionPreference:   ProtoToComputeAlphaServiceAttachmentConnectionPreferenceEnum(p.GetConnectionPreference()),
		EnableProxyProtocol:    dcl.Bool(p.EnableProxyProtocol),
		PscServiceAttachmentId: ProtoToComputeAlphaServiceAttachmentPscServiceAttachmentId(p.GetPscServiceAttachmentId()),
		Fingerprint:            dcl.StringOrNil(p.Fingerprint),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetConnectedEndpoints() {
		obj.ConnectedEndpoints = append(obj.ConnectedEndpoints, *ProtoToComputeAlphaServiceAttachmentConnectedEndpoints(r))
	}
	for _, r := range p.GetNatSubnets() {
		obj.NatSubnets = append(obj.NatSubnets, r)
	}
	for _, r := range p.GetConsumerRejectLists() {
		obj.ConsumerRejectLists = append(obj.ConsumerRejectLists, r)
	}
	for _, r := range p.GetConsumerAcceptLists() {
		obj.ConsumerAcceptLists = append(obj.ConsumerAcceptLists, *ProtoToComputeAlphaServiceAttachmentConsumerAcceptLists(r))
	}
	return obj
}

// ServiceAttachmentConnectionPreferenceEnumToProto converts a ServiceAttachmentConnectionPreferenceEnum enum to its proto representation.
func ComputeAlphaServiceAttachmentConnectionPreferenceEnumToProto(e *alpha.ServiceAttachmentConnectionPreferenceEnum) alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum {
	if e == nil {
		return alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum_value["ServiceAttachmentConnectionPreferenceEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum(v)
	}
	return alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum(0)
}

// ServiceAttachmentConnectedEndpointsStatusEnumToProto converts a ServiceAttachmentConnectedEndpointsStatusEnum enum to its proto representation.
func ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnumToProto(e *alpha.ServiceAttachmentConnectedEndpointsStatusEnum) alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum {
	if e == nil {
		return alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum_value["ServiceAttachmentConnectedEndpointsStatusEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(v)
	}
	return alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(0)
}

// ServiceAttachmentConnectedEndpointsToProto converts a ServiceAttachmentConnectedEndpoints resource to its proto representation.
func ComputeAlphaServiceAttachmentConnectedEndpointsToProto(o *alpha.ServiceAttachmentConnectedEndpoints) *alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints{
		Status:          ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnumToProto(o.Status),
		PscConnectionId: dcl.ValueOrEmptyInt64(o.PscConnectionId),
		Endpoint:        dcl.ValueOrEmptyString(o.Endpoint),
	}
	return p
}

// ServiceAttachmentConsumerAcceptListsToProto converts a ServiceAttachmentConsumerAcceptLists resource to its proto representation.
func ComputeAlphaServiceAttachmentConsumerAcceptListsToProto(o *alpha.ServiceAttachmentConsumerAcceptLists) *alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.ValueOrEmptyString(o.ProjectIdOrNum),
		ConnectionLimit: dcl.ValueOrEmptyInt64(o.ConnectionLimit),
	}
	return p
}

// ServiceAttachmentPscServiceAttachmentIdToProto converts a ServiceAttachmentPscServiceAttachmentId resource to its proto representation.
func ComputeAlphaServiceAttachmentPscServiceAttachmentIdToProto(o *alpha.ServiceAttachmentPscServiceAttachmentId) *alphapb.ComputeAlphaServiceAttachmentPscServiceAttachmentId {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaServiceAttachmentPscServiceAttachmentId{
		High: dcl.ValueOrEmptyInt64(o.High),
		Low:  dcl.ValueOrEmptyInt64(o.Low),
	}
	return p
}

// ServiceAttachmentToProto converts a ServiceAttachment resource to its proto representation.
func ServiceAttachmentToProto(resource *alpha.ServiceAttachment) *alphapb.ComputeAlphaServiceAttachment {
	p := &alphapb.ComputeAlphaServiceAttachment{
		Id:                     dcl.ValueOrEmptyInt64(resource.Id),
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		SelfLink:               dcl.ValueOrEmptyString(resource.SelfLink),
		Region:                 dcl.ValueOrEmptyString(resource.Region),
		TargetService:          dcl.ValueOrEmptyString(resource.TargetService),
		ConnectionPreference:   ComputeAlphaServiceAttachmentConnectionPreferenceEnumToProto(resource.ConnectionPreference),
		EnableProxyProtocol:    dcl.ValueOrEmptyBool(resource.EnableProxyProtocol),
		PscServiceAttachmentId: ComputeAlphaServiceAttachmentPscServiceAttachmentIdToProto(resource.PscServiceAttachmentId),
		Fingerprint:            dcl.ValueOrEmptyString(resource.Fingerprint),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ConnectedEndpoints {
		p.ConnectedEndpoints = append(p.ConnectedEndpoints, ComputeAlphaServiceAttachmentConnectedEndpointsToProto(&r))
	}
	for _, r := range resource.NatSubnets {
		p.NatSubnets = append(p.NatSubnets, r)
	}
	for _, r := range resource.ConsumerRejectLists {
		p.ConsumerRejectLists = append(p.ConsumerRejectLists, r)
	}
	for _, r := range resource.ConsumerAcceptLists {
		p.ConsumerAcceptLists = append(p.ConsumerAcceptLists, ComputeAlphaServiceAttachmentConsumerAcceptListsToProto(&r))
	}

	return p
}

// ApplyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) applyServiceAttachment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaServiceAttachmentRequest) (*alphapb.ComputeAlphaServiceAttachment, error) {
	p := ProtoToServiceAttachment(request.GetResource())
	res, err := c.ApplyServiceAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceAttachmentToProto(res)
	return r, nil
}

// ApplyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) ApplyComputeAlphaServiceAttachment(ctx context.Context, request *alphapb.ApplyComputeAlphaServiceAttachmentRequest) (*alphapb.ComputeAlphaServiceAttachment, error) {
	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServiceAttachment(ctx, cl, request)
}

// DeleteServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Delete() method.
func (s *ServiceAttachmentServer) DeleteComputeAlphaServiceAttachment(ctx context.Context, request *alphapb.DeleteComputeAlphaServiceAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceAttachment(ctx, ProtoToServiceAttachment(request.GetResource()))

}

// ListComputeAlphaServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachmentList() method.
func (s *ServiceAttachmentServer) ListComputeAlphaServiceAttachment(ctx context.Context, request *alphapb.ListComputeAlphaServiceAttachmentRequest) (*alphapb.ListComputeAlphaServiceAttachmentResponse, error) {
	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceAttachment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaServiceAttachment
	for _, r := range resources.Items {
		rp := ServiceAttachmentToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListComputeAlphaServiceAttachmentResponse{Items: protos}, nil
}

func createConfigServiceAttachment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
