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

// Server implements the gRPC interface for ServiceAttachment.
type ServiceAttachmentServer struct{}

// ProtoToServiceAttachmentConnectionPreferenceEnum converts a ServiceAttachmentConnectionPreferenceEnum enum from its proto representation.
func ProtoToComputeServiceAttachmentConnectionPreferenceEnum(e computepb.ComputeServiceAttachmentConnectionPreferenceEnum) *compute.ServiceAttachmentConnectionPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeServiceAttachmentConnectionPreferenceEnum_name[int32(e)]; ok {
		e := compute.ServiceAttachmentConnectionPreferenceEnum(n[len("ComputeServiceAttachmentConnectionPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpointsStatusEnum converts a ServiceAttachmentConnectedEndpointsStatusEnum enum from its proto representation.
func ProtoToComputeServiceAttachmentConnectedEndpointsStatusEnum(e computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum) *compute.ServiceAttachmentConnectedEndpointsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum_name[int32(e)]; ok {
		e := compute.ServiceAttachmentConnectedEndpointsStatusEnum(n[len("ComputeServiceAttachmentConnectedEndpointsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpoints converts a ServiceAttachmentConnectedEndpoints resource from its proto representation.
func ProtoToComputeServiceAttachmentConnectedEndpoints(p *computepb.ComputeServiceAttachmentConnectedEndpoints) *compute.ServiceAttachmentConnectedEndpoints {
	if p == nil {
		return nil
	}
	obj := &compute.ServiceAttachmentConnectedEndpoints{
		Status:          ProtoToComputeServiceAttachmentConnectedEndpointsStatusEnum(p.GetStatus()),
		PscConnectionId: dcl.Int64OrNil(p.PscConnectionId),
		Endpoint:        dcl.StringOrNil(p.Endpoint),
	}
	return obj
}

// ProtoToServiceAttachmentConsumerAcceptLists converts a ServiceAttachmentConsumerAcceptLists resource from its proto representation.
func ProtoToComputeServiceAttachmentConsumerAcceptLists(p *computepb.ComputeServiceAttachmentConsumerAcceptLists) *compute.ServiceAttachmentConsumerAcceptLists {
	if p == nil {
		return nil
	}
	obj := &compute.ServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.StringOrNil(p.ProjectIdOrNum),
		ConnectionLimit: dcl.Int64OrNil(p.ConnectionLimit),
	}
	return obj
}

// ProtoToServiceAttachmentPscServiceAttachmentId converts a ServiceAttachmentPscServiceAttachmentId resource from its proto representation.
func ProtoToComputeServiceAttachmentPscServiceAttachmentId(p *computepb.ComputeServiceAttachmentPscServiceAttachmentId) *compute.ServiceAttachmentPscServiceAttachmentId {
	if p == nil {
		return nil
	}
	obj := &compute.ServiceAttachmentPscServiceAttachmentId{
		High: dcl.Int64OrNil(p.High),
		Low:  dcl.Int64OrNil(p.Low),
	}
	return obj
}

// ProtoToServiceAttachment converts a ServiceAttachment resource from its proto representation.
func ProtoToServiceAttachment(p *computepb.ComputeServiceAttachment) *compute.ServiceAttachment {
	obj := &compute.ServiceAttachment{
		Id:                     dcl.Int64OrNil(p.Id),
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		SelfLink:               dcl.StringOrNil(p.SelfLink),
		Region:                 dcl.StringOrNil(p.Region),
		TargetService:          dcl.StringOrNil(p.TargetService),
		ConnectionPreference:   ProtoToComputeServiceAttachmentConnectionPreferenceEnum(p.GetConnectionPreference()),
		EnableProxyProtocol:    dcl.Bool(p.EnableProxyProtocol),
		PscServiceAttachmentId: ProtoToComputeServiceAttachmentPscServiceAttachmentId(p.GetPscServiceAttachmentId()),
		Fingerprint:            dcl.StringOrNil(p.Fingerprint),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetConnectedEndpoints() {
		obj.ConnectedEndpoints = append(obj.ConnectedEndpoints, *ProtoToComputeServiceAttachmentConnectedEndpoints(r))
	}
	for _, r := range p.GetNatSubnets() {
		obj.NatSubnets = append(obj.NatSubnets, r)
	}
	for _, r := range p.GetConsumerRejectLists() {
		obj.ConsumerRejectLists = append(obj.ConsumerRejectLists, r)
	}
	for _, r := range p.GetConsumerAcceptLists() {
		obj.ConsumerAcceptLists = append(obj.ConsumerAcceptLists, *ProtoToComputeServiceAttachmentConsumerAcceptLists(r))
	}
	return obj
}

// ServiceAttachmentConnectionPreferenceEnumToProto converts a ServiceAttachmentConnectionPreferenceEnum enum to its proto representation.
func ComputeServiceAttachmentConnectionPreferenceEnumToProto(e *compute.ServiceAttachmentConnectionPreferenceEnum) computepb.ComputeServiceAttachmentConnectionPreferenceEnum {
	if e == nil {
		return computepb.ComputeServiceAttachmentConnectionPreferenceEnum(0)
	}
	if v, ok := computepb.ComputeServiceAttachmentConnectionPreferenceEnum_value["ServiceAttachmentConnectionPreferenceEnum"+string(*e)]; ok {
		return computepb.ComputeServiceAttachmentConnectionPreferenceEnum(v)
	}
	return computepb.ComputeServiceAttachmentConnectionPreferenceEnum(0)
}

// ServiceAttachmentConnectedEndpointsStatusEnumToProto converts a ServiceAttachmentConnectedEndpointsStatusEnum enum to its proto representation.
func ComputeServiceAttachmentConnectedEndpointsStatusEnumToProto(e *compute.ServiceAttachmentConnectedEndpointsStatusEnum) computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum {
	if e == nil {
		return computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum(0)
	}
	if v, ok := computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum_value["ServiceAttachmentConnectedEndpointsStatusEnum"+string(*e)]; ok {
		return computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum(v)
	}
	return computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum(0)
}

// ServiceAttachmentConnectedEndpointsToProto converts a ServiceAttachmentConnectedEndpoints resource to its proto representation.
func ComputeServiceAttachmentConnectedEndpointsToProto(o *compute.ServiceAttachmentConnectedEndpoints) *computepb.ComputeServiceAttachmentConnectedEndpoints {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeServiceAttachmentConnectedEndpoints{
		Status:          ComputeServiceAttachmentConnectedEndpointsStatusEnumToProto(o.Status),
		PscConnectionId: dcl.ValueOrEmptyInt64(o.PscConnectionId),
		Endpoint:        dcl.ValueOrEmptyString(o.Endpoint),
	}
	return p
}

// ServiceAttachmentConsumerAcceptListsToProto converts a ServiceAttachmentConsumerAcceptLists resource to its proto representation.
func ComputeServiceAttachmentConsumerAcceptListsToProto(o *compute.ServiceAttachmentConsumerAcceptLists) *computepb.ComputeServiceAttachmentConsumerAcceptLists {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.ValueOrEmptyString(o.ProjectIdOrNum),
		ConnectionLimit: dcl.ValueOrEmptyInt64(o.ConnectionLimit),
	}
	return p
}

// ServiceAttachmentPscServiceAttachmentIdToProto converts a ServiceAttachmentPscServiceAttachmentId resource to its proto representation.
func ComputeServiceAttachmentPscServiceAttachmentIdToProto(o *compute.ServiceAttachmentPscServiceAttachmentId) *computepb.ComputeServiceAttachmentPscServiceAttachmentId {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeServiceAttachmentPscServiceAttachmentId{
		High: dcl.ValueOrEmptyInt64(o.High),
		Low:  dcl.ValueOrEmptyInt64(o.Low),
	}
	return p
}

// ServiceAttachmentToProto converts a ServiceAttachment resource to its proto representation.
func ServiceAttachmentToProto(resource *compute.ServiceAttachment) *computepb.ComputeServiceAttachment {
	p := &computepb.ComputeServiceAttachment{
		Id:                     dcl.ValueOrEmptyInt64(resource.Id),
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		SelfLink:               dcl.ValueOrEmptyString(resource.SelfLink),
		Region:                 dcl.ValueOrEmptyString(resource.Region),
		TargetService:          dcl.ValueOrEmptyString(resource.TargetService),
		ConnectionPreference:   ComputeServiceAttachmentConnectionPreferenceEnumToProto(resource.ConnectionPreference),
		EnableProxyProtocol:    dcl.ValueOrEmptyBool(resource.EnableProxyProtocol),
		PscServiceAttachmentId: ComputeServiceAttachmentPscServiceAttachmentIdToProto(resource.PscServiceAttachmentId),
		Fingerprint:            dcl.ValueOrEmptyString(resource.Fingerprint),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ConnectedEndpoints {
		p.ConnectedEndpoints = append(p.ConnectedEndpoints, ComputeServiceAttachmentConnectedEndpointsToProto(&r))
	}
	for _, r := range resource.NatSubnets {
		p.NatSubnets = append(p.NatSubnets, r)
	}
	for _, r := range resource.ConsumerRejectLists {
		p.ConsumerRejectLists = append(p.ConsumerRejectLists, r)
	}
	for _, r := range resource.ConsumerAcceptLists {
		p.ConsumerAcceptLists = append(p.ConsumerAcceptLists, ComputeServiceAttachmentConsumerAcceptListsToProto(&r))
	}

	return p
}

// ApplyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) applyServiceAttachment(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeServiceAttachmentRequest) (*computepb.ComputeServiceAttachment, error) {
	p := ProtoToServiceAttachment(request.GetResource())
	res, err := c.ApplyServiceAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceAttachmentToProto(res)
	return r, nil
}

// ApplyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) ApplyComputeServiceAttachment(ctx context.Context, request *computepb.ApplyComputeServiceAttachmentRequest) (*computepb.ComputeServiceAttachment, error) {
	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServiceAttachment(ctx, cl, request)
}

// DeleteServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Delete() method.
func (s *ServiceAttachmentServer) DeleteComputeServiceAttachment(ctx context.Context, request *computepb.DeleteComputeServiceAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceAttachment(ctx, ProtoToServiceAttachment(request.GetResource()))

}

// ListComputeServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachmentList() method.
func (s *ServiceAttachmentServer) ListComputeServiceAttachment(ctx context.Context, request *computepb.ListComputeServiceAttachmentRequest) (*computepb.ListComputeServiceAttachmentResponse, error) {
	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceAttachment(ctx, ProtoToServiceAttachment(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeServiceAttachment
	for _, r := range resources.Items {
		rp := ServiceAttachmentToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeServiceAttachmentResponse{Items: protos}, nil
}

func createConfigServiceAttachment(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
