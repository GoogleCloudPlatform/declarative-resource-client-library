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
	eventarcpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/eventarc/eventarc_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
)

// Server implements the gRPC interface for Trigger.
type TriggerServer struct{}

// ProtoToTriggerEventFilters converts a TriggerEventFilters resource from its proto representation.
func ProtoToEventarcTriggerEventFilters(p *eventarcpb.EventarcTriggerEventFilters) *eventarc.TriggerEventFilters {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerEventFilters{
		Attribute: dcl.StringOrNil(p.Attribute),
		Value:     dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToTriggerDestination converts a TriggerDestination resource from its proto representation.
func ProtoToEventarcTriggerDestination(p *eventarcpb.EventarcTriggerDestination) *eventarc.TriggerDestination {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestination{
		CloudRun:      ProtoToEventarcTriggerDestinationCloudRun(p.GetCloudRun()),
		CloudFunction: dcl.StringOrNil(p.CloudFunction),
	}
	return obj
}

// ProtoToTriggerDestinationCloudRun converts a TriggerDestinationCloudRun resource from its proto representation.
func ProtoToEventarcTriggerDestinationCloudRun(p *eventarcpb.EventarcTriggerDestinationCloudRun) *eventarc.TriggerDestinationCloudRun {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestinationCloudRun{
		Service: dcl.StringOrNil(p.Service),
		Path:    dcl.StringOrNil(p.Path),
		Region:  dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToTriggerTransport converts a TriggerTransport resource from its proto representation.
func ProtoToEventarcTriggerTransport(p *eventarcpb.EventarcTriggerTransport) *eventarc.TriggerTransport {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerTransport{
		Pubsub: ProtoToEventarcTriggerTransportPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToTriggerTransportPubsub converts a TriggerTransportPubsub resource from its proto representation.
func ProtoToEventarcTriggerTransportPubsub(p *eventarcpb.EventarcTriggerTransportPubsub) *eventarc.TriggerTransportPubsub {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerTransportPubsub{
		Topic:        dcl.StringOrNil(p.Topic),
		Subscription: dcl.StringOrNil(p.Subscription),
	}
	return obj
}

// ProtoToTrigger converts a Trigger resource from its proto representation.
func ProtoToTrigger(p *eventarcpb.EventarcTrigger) *eventarc.Trigger {
	obj := &eventarc.Trigger{
		Name:           dcl.StringOrNil(p.Name),
		Uid:            dcl.StringOrNil(p.Uid),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		ServiceAccount: dcl.StringOrNil(p.ServiceAccount),
		Destination:    ProtoToEventarcTriggerDestination(p.GetDestination()),
		Transport:      ProtoToEventarcTriggerTransport(p.GetTransport()),
		Etag:           dcl.StringOrNil(p.Etag),
		Project:        dcl.StringOrNil(p.Project),
		Location:       dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetEventFilters() {
		obj.EventFilters = append(obj.EventFilters, *ProtoToEventarcTriggerEventFilters(r))
	}
	return obj
}

// TriggerEventFiltersToProto converts a TriggerEventFilters resource to its proto representation.
func EventarcTriggerEventFiltersToProto(o *eventarc.TriggerEventFilters) *eventarcpb.EventarcTriggerEventFilters {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerEventFilters{
		Attribute: dcl.ValueOrEmptyString(o.Attribute),
		Value:     dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// TriggerDestinationToProto converts a TriggerDestination resource to its proto representation.
func EventarcTriggerDestinationToProto(o *eventarc.TriggerDestination) *eventarcpb.EventarcTriggerDestination {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestination{
		CloudRun:      EventarcTriggerDestinationCloudRunToProto(o.CloudRun),
		CloudFunction: dcl.ValueOrEmptyString(o.CloudFunction),
	}
	return p
}

// TriggerDestinationCloudRunToProto converts a TriggerDestinationCloudRun resource to its proto representation.
func EventarcTriggerDestinationCloudRunToProto(o *eventarc.TriggerDestinationCloudRun) *eventarcpb.EventarcTriggerDestinationCloudRun {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestinationCloudRun{
		Service: dcl.ValueOrEmptyString(o.Service),
		Path:    dcl.ValueOrEmptyString(o.Path),
		Region:  dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// TriggerTransportToProto converts a TriggerTransport resource to its proto representation.
func EventarcTriggerTransportToProto(o *eventarc.TriggerTransport) *eventarcpb.EventarcTriggerTransport {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerTransport{
		Pubsub: EventarcTriggerTransportPubsubToProto(o.Pubsub),
	}
	return p
}

// TriggerTransportPubsubToProto converts a TriggerTransportPubsub resource to its proto representation.
func EventarcTriggerTransportPubsubToProto(o *eventarc.TriggerTransportPubsub) *eventarcpb.EventarcTriggerTransportPubsub {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerTransportPubsub{
		Topic:        dcl.ValueOrEmptyString(o.Topic),
		Subscription: dcl.ValueOrEmptyString(o.Subscription),
	}
	return p
}

// TriggerToProto converts a Trigger resource to its proto representation.
func TriggerToProto(resource *eventarc.Trigger) *eventarcpb.EventarcTrigger {
	p := &eventarcpb.EventarcTrigger{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		Uid:            dcl.ValueOrEmptyString(resource.Uid),
		CreateTime:     dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:     dcl.ValueOrEmptyString(resource.UpdateTime),
		ServiceAccount: dcl.ValueOrEmptyString(resource.ServiceAccount),
		Destination:    EventarcTriggerDestinationToProto(resource.Destination),
		Transport:      EventarcTriggerTransportToProto(resource.Transport),
		Etag:           dcl.ValueOrEmptyString(resource.Etag),
		Project:        dcl.ValueOrEmptyString(resource.Project),
		Location:       dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.EventFilters {
		p.EventFilters = append(p.EventFilters, EventarcTriggerEventFiltersToProto(&r))
	}

	return p
}

// ApplyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) applyTrigger(ctx context.Context, c *eventarc.Client, request *eventarcpb.ApplyEventarcTriggerRequest) (*eventarcpb.EventarcTrigger, error) {
	p := ProtoToTrigger(request.GetResource())
	res, err := c.ApplyTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TriggerToProto(res)
	return r, nil
}

// ApplyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) ApplyEventarcTrigger(ctx context.Context, request *eventarcpb.ApplyEventarcTriggerRequest) (*eventarcpb.EventarcTrigger, error) {
	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTrigger(ctx, cl, request)
}

// DeleteTrigger handles the gRPC request by passing it to the underlying Trigger Delete() method.
func (s *TriggerServer) DeleteEventarcTrigger(ctx context.Context, request *eventarcpb.DeleteEventarcTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTrigger(ctx, ProtoToTrigger(request.GetResource()))

}

// ListTrigger handles the gRPC request by passing it to the underlying TriggerList() method.
func (s *TriggerServer) ListEventarcTrigger(ctx context.Context, request *eventarcpb.ListEventarcTriggerRequest) (*eventarcpb.ListEventarcTriggerResponse, error) {
	cl, err := createConfigTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTrigger(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*eventarcpb.EventarcTrigger
	for _, r := range resources.Items {
		rp := TriggerToProto(r)
		protos = append(protos, rp)
	}
	return &eventarcpb.ListEventarcTriggerResponse{Items: protos}, nil
}

func createConfigTrigger(ctx context.Context, service_account_file string) (*eventarc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return eventarc.NewClient(conf), nil
}
