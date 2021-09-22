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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/eventarc/alpha/eventarc_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/alpha"
)

// Server implements the gRPC interface for Trigger.
type TriggerServer struct{}

// ProtoToTriggerMatchingCriteria converts a TriggerMatchingCriteria resource from its proto representation.
func ProtoToEventarcAlphaTriggerMatchingCriteria(p *alphapb.EventarcAlphaTriggerMatchingCriteria) *alpha.TriggerMatchingCriteria {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerMatchingCriteria{
		Attribute: dcl.StringOrNil(p.Attribute),
		Value:     dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToTriggerDestination converts a TriggerDestination resource from its proto representation.
func ProtoToEventarcAlphaTriggerDestination(p *alphapb.EventarcAlphaTriggerDestination) *alpha.TriggerDestination {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestination{
		CloudRunService: ProtoToEventarcAlphaTriggerDestinationCloudRunService(p.GetCloudRunService()),
		CloudFunction:   dcl.StringOrNil(p.CloudFunction),
	}
	return obj
}

// ProtoToTriggerDestinationCloudRunService converts a TriggerDestinationCloudRunService resource from its proto representation.
func ProtoToEventarcAlphaTriggerDestinationCloudRunService(p *alphapb.EventarcAlphaTriggerDestinationCloudRunService) *alpha.TriggerDestinationCloudRunService {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestinationCloudRunService{
		Service: dcl.StringOrNil(p.Service),
		Path:    dcl.StringOrNil(p.Path),
		Region:  dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToTriggerTransport converts a TriggerTransport resource from its proto representation.
func ProtoToEventarcAlphaTriggerTransport(p *alphapb.EventarcAlphaTriggerTransport) *alpha.TriggerTransport {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerTransport{
		Pubsub: ProtoToEventarcAlphaTriggerTransportPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToTriggerTransportPubsub converts a TriggerTransportPubsub resource from its proto representation.
func ProtoToEventarcAlphaTriggerTransportPubsub(p *alphapb.EventarcAlphaTriggerTransportPubsub) *alpha.TriggerTransportPubsub {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerTransportPubsub{
		Topic:        dcl.StringOrNil(p.Topic),
		Subscription: dcl.StringOrNil(p.Subscription),
	}
	return obj
}

// ProtoToTrigger converts a Trigger resource from its proto representation.
func ProtoToTrigger(p *alphapb.EventarcAlphaTrigger) *alpha.Trigger {
	obj := &alpha.Trigger{
		Name:           dcl.StringOrNil(p.Name),
		Uid:            dcl.StringOrNil(p.Uid),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		ServiceAccount: dcl.StringOrNil(p.ServiceAccount),
		Destination:    ProtoToEventarcAlphaTriggerDestination(p.GetDestination()),
		Transport:      ProtoToEventarcAlphaTriggerTransport(p.GetTransport()),
		Etag:           dcl.StringOrNil(p.Etag),
		Project:        dcl.StringOrNil(p.Project),
		Location:       dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetMatchingCriteria() {
		obj.MatchingCriteria = append(obj.MatchingCriteria, *ProtoToEventarcAlphaTriggerMatchingCriteria(r))
	}
	return obj
}

// TriggerMatchingCriteriaToProto converts a TriggerMatchingCriteria resource to its proto representation.
func EventarcAlphaTriggerMatchingCriteriaToProto(o *alpha.TriggerMatchingCriteria) *alphapb.EventarcAlphaTriggerMatchingCriteria {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerMatchingCriteria{
		Attribute: dcl.ValueOrEmptyString(o.Attribute),
		Value:     dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// TriggerDestinationToProto converts a TriggerDestination resource to its proto representation.
func EventarcAlphaTriggerDestinationToProto(o *alpha.TriggerDestination) *alphapb.EventarcAlphaTriggerDestination {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestination{
		CloudRunService: EventarcAlphaTriggerDestinationCloudRunServiceToProto(o.CloudRunService),
		CloudFunction:   dcl.ValueOrEmptyString(o.CloudFunction),
	}
	return p
}

// TriggerDestinationCloudRunServiceToProto converts a TriggerDestinationCloudRunService resource to its proto representation.
func EventarcAlphaTriggerDestinationCloudRunServiceToProto(o *alpha.TriggerDestinationCloudRunService) *alphapb.EventarcAlphaTriggerDestinationCloudRunService {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestinationCloudRunService{
		Service: dcl.ValueOrEmptyString(o.Service),
		Path:    dcl.ValueOrEmptyString(o.Path),
		Region:  dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// TriggerTransportToProto converts a TriggerTransport resource to its proto representation.
func EventarcAlphaTriggerTransportToProto(o *alpha.TriggerTransport) *alphapb.EventarcAlphaTriggerTransport {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerTransport{
		Pubsub: EventarcAlphaTriggerTransportPubsubToProto(o.Pubsub),
	}
	return p
}

// TriggerTransportPubsubToProto converts a TriggerTransportPubsub resource to its proto representation.
func EventarcAlphaTriggerTransportPubsubToProto(o *alpha.TriggerTransportPubsub) *alphapb.EventarcAlphaTriggerTransportPubsub {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerTransportPubsub{
		Topic:        dcl.ValueOrEmptyString(o.Topic),
		Subscription: dcl.ValueOrEmptyString(o.Subscription),
	}
	return p
}

// TriggerToProto converts a Trigger resource to its proto representation.
func TriggerToProto(resource *alpha.Trigger) *alphapb.EventarcAlphaTrigger {
	p := &alphapb.EventarcAlphaTrigger{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		Uid:            dcl.ValueOrEmptyString(resource.Uid),
		CreateTime:     dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:     dcl.ValueOrEmptyString(resource.UpdateTime),
		ServiceAccount: dcl.ValueOrEmptyString(resource.ServiceAccount),
		Destination:    EventarcAlphaTriggerDestinationToProto(resource.Destination),
		Transport:      EventarcAlphaTriggerTransportToProto(resource.Transport),
		Etag:           dcl.ValueOrEmptyString(resource.Etag),
		Project:        dcl.ValueOrEmptyString(resource.Project),
		Location:       dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.MatchingCriteria {
		p.MatchingCriteria = append(p.MatchingCriteria, EventarcAlphaTriggerMatchingCriteriaToProto(&r))
	}

	return p
}

// ApplyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) applyTrigger(ctx context.Context, c *alpha.Client, request *alphapb.ApplyEventarcAlphaTriggerRequest) (*alphapb.EventarcAlphaTrigger, error) {
	p := ProtoToTrigger(request.GetResource())
	res, err := c.ApplyTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TriggerToProto(res)
	return r, nil
}

// ApplyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) ApplyEventarcAlphaTrigger(ctx context.Context, request *alphapb.ApplyEventarcAlphaTriggerRequest) (*alphapb.EventarcAlphaTrigger, error) {
	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTrigger(ctx, cl, request)
}

// DeleteTrigger handles the gRPC request by passing it to the underlying Trigger Delete() method.
func (s *TriggerServer) DeleteEventarcAlphaTrigger(ctx context.Context, request *alphapb.DeleteEventarcAlphaTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTrigger(ctx, ProtoToTrigger(request.GetResource()))

}

// ListEventarcAlphaTrigger handles the gRPC request by passing it to the underlying TriggerList() method.
func (s *TriggerServer) ListEventarcAlphaTrigger(ctx context.Context, request *alphapb.ListEventarcAlphaTriggerRequest) (*alphapb.ListEventarcAlphaTriggerResponse, error) {
	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTrigger(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.EventarcAlphaTrigger
	for _, r := range resources.Items {
		rp := TriggerToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListEventarcAlphaTriggerResponse{Items: protos}, nil
}

func createConfigTrigger(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
