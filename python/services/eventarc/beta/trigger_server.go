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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/eventarc/beta/eventarc_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
)

// Server implements the gRPC interface for Trigger.
type TriggerServer struct{}

// ProtoToTriggerMatchingCriteria converts a TriggerMatchingCriteria resource from its proto representation.
func ProtoToEventarcBetaTriggerMatchingCriteria(p *betapb.EventarcBetaTriggerMatchingCriteria) *beta.TriggerMatchingCriteria {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerMatchingCriteria{
		Attribute: dcl.StringOrNil(p.Attribute),
		Value:     dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToTriggerDestination converts a TriggerDestination resource from its proto representation.
func ProtoToEventarcBetaTriggerDestination(p *betapb.EventarcBetaTriggerDestination) *beta.TriggerDestination {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestination{
		CloudRunService: ProtoToEventarcBetaTriggerDestinationCloudRunService(p.GetCloudRunService()),
	}
	return obj
}

// ProtoToTriggerDestinationCloudRunService converts a TriggerDestinationCloudRunService resource from its proto representation.
func ProtoToEventarcBetaTriggerDestinationCloudRunService(p *betapb.EventarcBetaTriggerDestinationCloudRunService) *beta.TriggerDestinationCloudRunService {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestinationCloudRunService{
		Service: dcl.StringOrNil(p.Service),
		Path:    dcl.StringOrNil(p.Path),
		Region:  dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToTriggerTransport converts a TriggerTransport resource from its proto representation.
func ProtoToEventarcBetaTriggerTransport(p *betapb.EventarcBetaTriggerTransport) *beta.TriggerTransport {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerTransport{
		Pubsub: ProtoToEventarcBetaTriggerTransportPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToTriggerTransportPubsub converts a TriggerTransportPubsub resource from its proto representation.
func ProtoToEventarcBetaTriggerTransportPubsub(p *betapb.EventarcBetaTriggerTransportPubsub) *beta.TriggerTransportPubsub {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerTransportPubsub{
		Topic:        dcl.StringOrNil(p.Topic),
		Subscription: dcl.StringOrNil(p.Subscription),
	}
	return obj
}

// ProtoToTrigger converts a Trigger resource from its proto representation.
func ProtoToTrigger(p *betapb.EventarcBetaTrigger) *beta.Trigger {
	obj := &beta.Trigger{
		Name:           dcl.StringOrNil(p.Name),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		ServiceAccount: dcl.StringOrNil(p.ServiceAccount),
		Destination:    ProtoToEventarcBetaTriggerDestination(p.GetDestination()),
		Transport:      ProtoToEventarcBetaTriggerTransport(p.GetTransport()),
		Etag:           dcl.StringOrNil(p.Etag),
		Project:        dcl.StringOrNil(p.Project),
		Location:       dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetMatchingCriteria() {
		obj.MatchingCriteria = append(obj.MatchingCriteria, *ProtoToEventarcBetaTriggerMatchingCriteria(r))
	}
	return obj
}

// TriggerMatchingCriteriaToProto converts a TriggerMatchingCriteria resource to its proto representation.
func EventarcBetaTriggerMatchingCriteriaToProto(o *beta.TriggerMatchingCriteria) *betapb.EventarcBetaTriggerMatchingCriteria {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerMatchingCriteria{
		Attribute: dcl.ValueOrEmptyString(o.Attribute),
		Value:     dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// TriggerDestinationToProto converts a TriggerDestination resource to its proto representation.
func EventarcBetaTriggerDestinationToProto(o *beta.TriggerDestination) *betapb.EventarcBetaTriggerDestination {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestination{
		CloudRunService: EventarcBetaTriggerDestinationCloudRunServiceToProto(o.CloudRunService),
	}
	return p
}

// TriggerDestinationCloudRunServiceToProto converts a TriggerDestinationCloudRunService resource to its proto representation.
func EventarcBetaTriggerDestinationCloudRunServiceToProto(o *beta.TriggerDestinationCloudRunService) *betapb.EventarcBetaTriggerDestinationCloudRunService {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestinationCloudRunService{
		Service: dcl.ValueOrEmptyString(o.Service),
		Path:    dcl.ValueOrEmptyString(o.Path),
		Region:  dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// TriggerTransportToProto converts a TriggerTransport resource to its proto representation.
func EventarcBetaTriggerTransportToProto(o *beta.TriggerTransport) *betapb.EventarcBetaTriggerTransport {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerTransport{
		Pubsub: EventarcBetaTriggerTransportPubsubToProto(o.Pubsub),
	}
	return p
}

// TriggerTransportPubsubToProto converts a TriggerTransportPubsub resource to its proto representation.
func EventarcBetaTriggerTransportPubsubToProto(o *beta.TriggerTransportPubsub) *betapb.EventarcBetaTriggerTransportPubsub {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerTransportPubsub{
		Topic:        dcl.ValueOrEmptyString(o.Topic),
		Subscription: dcl.ValueOrEmptyString(o.Subscription),
	}
	return p
}

// TriggerToProto converts a Trigger resource to its proto representation.
func TriggerToProto(resource *beta.Trigger) *betapb.EventarcBetaTrigger {
	p := &betapb.EventarcBetaTrigger{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		CreateTime:     dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:     dcl.ValueOrEmptyString(resource.UpdateTime),
		ServiceAccount: dcl.ValueOrEmptyString(resource.ServiceAccount),
		Destination:    EventarcBetaTriggerDestinationToProto(resource.Destination),
		Transport:      EventarcBetaTriggerTransportToProto(resource.Transport),
		Etag:           dcl.ValueOrEmptyString(resource.Etag),
		Project:        dcl.ValueOrEmptyString(resource.Project),
		Location:       dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.MatchingCriteria {
		p.MatchingCriteria = append(p.MatchingCriteria, EventarcBetaTriggerMatchingCriteriaToProto(&r))
	}

	return p
}

// ApplyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) applyTrigger(ctx context.Context, c *beta.Client, request *betapb.ApplyEventarcBetaTriggerRequest) (*betapb.EventarcBetaTrigger, error) {
	p := ProtoToTrigger(request.GetResource())
	res, err := c.ApplyTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TriggerToProto(res)
	return r, nil
}

// ApplyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) ApplyEventarcBetaTrigger(ctx context.Context, request *betapb.ApplyEventarcBetaTriggerRequest) (*betapb.EventarcBetaTrigger, error) {
	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTrigger(ctx, cl, request)
}

// DeleteTrigger handles the gRPC request by passing it to the underlying Trigger Delete() method.
func (s *TriggerServer) DeleteEventarcBetaTrigger(ctx context.Context, request *betapb.DeleteEventarcBetaTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTrigger(ctx, ProtoToTrigger(request.GetResource()))

}

// ListEventarcBetaTrigger handles the gRPC request by passing it to the underlying TriggerList() method.
func (s *TriggerServer) ListEventarcBetaTrigger(ctx context.Context, request *betapb.ListEventarcBetaTriggerRequest) (*betapb.ListEventarcBetaTriggerResponse, error) {
	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTrigger(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.EventarcBetaTrigger
	for _, r := range resources.Items {
		rp := TriggerToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListEventarcBetaTriggerResponse{Items: protos}, nil
}

func createConfigTrigger(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
