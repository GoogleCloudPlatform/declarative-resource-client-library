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

// Server implements the gRPC interface for ServiceAttachment.
type ServiceAttachmentServer struct{}

// ProtoToServiceAttachmentConnectionPreferenceEnum converts a ServiceAttachmentConnectionPreferenceEnum enum from its proto representation.
func ProtoToComputeBetaServiceAttachmentConnectionPreferenceEnum(e betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum) *beta.ServiceAttachmentConnectionPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum_name[int32(e)]; ok {
		e := beta.ServiceAttachmentConnectionPreferenceEnum(n[len("ComputeBetaServiceAttachmentConnectionPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpointsStatusEnum converts a ServiceAttachmentConnectedEndpointsStatusEnum enum from its proto representation.
func ProtoToComputeBetaServiceAttachmentConnectedEndpointsStatusEnum(e betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum) *beta.ServiceAttachmentConnectedEndpointsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum_name[int32(e)]; ok {
		e := beta.ServiceAttachmentConnectedEndpointsStatusEnum(n[len("ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpoints converts a ServiceAttachmentConnectedEndpoints resource from its proto representation.
func ProtoToComputeBetaServiceAttachmentConnectedEndpoints(p *betapb.ComputeBetaServiceAttachmentConnectedEndpoints) *beta.ServiceAttachmentConnectedEndpoints {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceAttachmentConnectedEndpoints{
		Status:          ProtoToComputeBetaServiceAttachmentConnectedEndpointsStatusEnum(p.GetStatus()),
		PscConnectionId: dcl.Int64OrNil(p.PscConnectionId),
		Endpoint:        dcl.StringOrNil(p.Endpoint),
	}
	return obj
}

// ProtoToServiceAttachmentConsumerAcceptLists converts a ServiceAttachmentConsumerAcceptLists resource from its proto representation.
func ProtoToComputeBetaServiceAttachmentConsumerAcceptLists(p *betapb.ComputeBetaServiceAttachmentConsumerAcceptLists) *beta.ServiceAttachmentConsumerAcceptLists {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.StringOrNil(p.ProjectIdOrNum),
		ConnectionLimit: dcl.Int64OrNil(p.ConnectionLimit),
	}
	return obj
}

// ProtoToServiceAttachmentPscServiceAttachmentId converts a ServiceAttachmentPscServiceAttachmentId resource from its proto representation.
func ProtoToComputeBetaServiceAttachmentPscServiceAttachmentId(p *betapb.ComputeBetaServiceAttachmentPscServiceAttachmentId) *beta.ServiceAttachmentPscServiceAttachmentId {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceAttachmentPscServiceAttachmentId{
		High: dcl.Int64OrNil(p.High),
		Low:  dcl.Int64OrNil(p.Low),
	}
	return obj
}

// ProtoToServiceAttachment converts a ServiceAttachment resource from its proto representation.
func ProtoToServiceAttachment(p *betapb.ComputeBetaServiceAttachment) *beta.ServiceAttachment {
	obj := &beta.ServiceAttachment{
		Id:                     dcl.Int64OrNil(p.Id),
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		SelfLink:               dcl.StringOrNil(p.SelfLink),
		Region:                 dcl.StringOrNil(p.Region),
		TargetService:          dcl.StringOrNil(p.TargetService),
		ConnectionPreference:   ProtoToComputeBetaServiceAttachmentConnectionPreferenceEnum(p.GetConnectionPreference()),
		EnableProxyProtocol:    dcl.Bool(p.EnableProxyProtocol),
		PscServiceAttachmentId: ProtoToComputeBetaServiceAttachmentPscServiceAttachmentId(p.GetPscServiceAttachmentId()),
		Fingerprint:            dcl.StringOrNil(p.Fingerprint),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetConnectedEndpoints() {
		obj.ConnectedEndpoints = append(obj.ConnectedEndpoints, *ProtoToComputeBetaServiceAttachmentConnectedEndpoints(r))
	}
	for _, r := range p.GetNatSubnets() {
		obj.NatSubnets = append(obj.NatSubnets, r)
	}
	for _, r := range p.GetConsumerRejectLists() {
		obj.ConsumerRejectLists = append(obj.ConsumerRejectLists, r)
	}
	for _, r := range p.GetConsumerAcceptLists() {
		obj.ConsumerAcceptLists = append(obj.ConsumerAcceptLists, *ProtoToComputeBetaServiceAttachmentConsumerAcceptLists(r))
	}
	return obj
}

// ServiceAttachmentConnectionPreferenceEnumToProto converts a ServiceAttachmentConnectionPreferenceEnum enum to its proto representation.
func ComputeBetaServiceAttachmentConnectionPreferenceEnumToProto(e *beta.ServiceAttachmentConnectionPreferenceEnum) betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum {
	if e == nil {
		return betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum(0)
	}
	if v, ok := betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum_value["ServiceAttachmentConnectionPreferenceEnum"+string(*e)]; ok {
		return betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum(v)
	}
	return betapb.ComputeBetaServiceAttachmentConnectionPreferenceEnum(0)
}

// ServiceAttachmentConnectedEndpointsStatusEnumToProto converts a ServiceAttachmentConnectedEndpointsStatusEnum enum to its proto representation.
func ComputeBetaServiceAttachmentConnectedEndpointsStatusEnumToProto(e *beta.ServiceAttachmentConnectedEndpointsStatusEnum) betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum {
	if e == nil {
		return betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum_value["ServiceAttachmentConnectedEndpointsStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum(v)
	}
	return betapb.ComputeBetaServiceAttachmentConnectedEndpointsStatusEnum(0)
}

// ServiceAttachmentConnectedEndpointsToProto converts a ServiceAttachmentConnectedEndpoints resource to its proto representation.
func ComputeBetaServiceAttachmentConnectedEndpointsToProto(o *beta.ServiceAttachmentConnectedEndpoints) *betapb.ComputeBetaServiceAttachmentConnectedEndpoints {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaServiceAttachmentConnectedEndpoints{
		Status:          ComputeBetaServiceAttachmentConnectedEndpointsStatusEnumToProto(o.Status),
		PscConnectionId: dcl.ValueOrEmptyInt64(o.PscConnectionId),
		Endpoint:        dcl.ValueOrEmptyString(o.Endpoint),
	}
	return p
}

// ServiceAttachmentConsumerAcceptListsToProto converts a ServiceAttachmentConsumerAcceptLists resource to its proto representation.
func ComputeBetaServiceAttachmentConsumerAcceptListsToProto(o *beta.ServiceAttachmentConsumerAcceptLists) *betapb.ComputeBetaServiceAttachmentConsumerAcceptLists {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.ValueOrEmptyString(o.ProjectIdOrNum),
		ConnectionLimit: dcl.ValueOrEmptyInt64(o.ConnectionLimit),
	}
	return p
}

// ServiceAttachmentPscServiceAttachmentIdToProto converts a ServiceAttachmentPscServiceAttachmentId resource to its proto representation.
func ComputeBetaServiceAttachmentPscServiceAttachmentIdToProto(o *beta.ServiceAttachmentPscServiceAttachmentId) *betapb.ComputeBetaServiceAttachmentPscServiceAttachmentId {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaServiceAttachmentPscServiceAttachmentId{
		High: dcl.ValueOrEmptyInt64(o.High),
		Low:  dcl.ValueOrEmptyInt64(o.Low),
	}
	return p
}

// ServiceAttachmentToProto converts a ServiceAttachment resource to its proto representation.
func ServiceAttachmentToProto(resource *beta.ServiceAttachment) *betapb.ComputeBetaServiceAttachment {
	p := &betapb.ComputeBetaServiceAttachment{
		Id:                     dcl.ValueOrEmptyInt64(resource.Id),
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		SelfLink:               dcl.ValueOrEmptyString(resource.SelfLink),
		Region:                 dcl.ValueOrEmptyString(resource.Region),
		TargetService:          dcl.ValueOrEmptyString(resource.TargetService),
		ConnectionPreference:   ComputeBetaServiceAttachmentConnectionPreferenceEnumToProto(resource.ConnectionPreference),
		EnableProxyProtocol:    dcl.ValueOrEmptyBool(resource.EnableProxyProtocol),
		PscServiceAttachmentId: ComputeBetaServiceAttachmentPscServiceAttachmentIdToProto(resource.PscServiceAttachmentId),
		Fingerprint:            dcl.ValueOrEmptyString(resource.Fingerprint),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ConnectedEndpoints {
		p.ConnectedEndpoints = append(p.ConnectedEndpoints, ComputeBetaServiceAttachmentConnectedEndpointsToProto(&r))
	}
	for _, r := range resource.NatSubnets {
		p.NatSubnets = append(p.NatSubnets, r)
	}
	for _, r := range resource.ConsumerRejectLists {
		p.ConsumerRejectLists = append(p.ConsumerRejectLists, r)
	}
	for _, r := range resource.ConsumerAcceptLists {
		p.ConsumerAcceptLists = append(p.ConsumerAcceptLists, ComputeBetaServiceAttachmentConsumerAcceptListsToProto(&r))
	}

	return p
}

// ApplyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) applyServiceAttachment(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaServiceAttachmentRequest) (*betapb.ComputeBetaServiceAttachment, error) {
	p := ProtoToServiceAttachment(request.GetResource())
	res, err := c.ApplyServiceAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceAttachmentToProto(res)
	return r, nil
}

// ApplyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) ApplyComputeBetaServiceAttachment(ctx context.Context, request *betapb.ApplyComputeBetaServiceAttachmentRequest) (*betapb.ComputeBetaServiceAttachment, error) {
	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServiceAttachment(ctx, cl, request)
}

// DeleteServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Delete() method.
func (s *ServiceAttachmentServer) DeleteComputeBetaServiceAttachment(ctx context.Context, request *betapb.DeleteComputeBetaServiceAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceAttachment(ctx, ProtoToServiceAttachment(request.GetResource()))

}

// ListComputeBetaServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachmentList() method.
func (s *ServiceAttachmentServer) ListComputeBetaServiceAttachment(ctx context.Context, request *betapb.ListComputeBetaServiceAttachmentRequest) (*betapb.ListComputeBetaServiceAttachmentResponse, error) {
	cl, err := createConfigServiceAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceAttachment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaServiceAttachment
	for _, r := range resources.Items {
		rp := ServiceAttachmentToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaServiceAttachmentResponse{Items: protos}, nil
}

func createConfigServiceAttachment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
