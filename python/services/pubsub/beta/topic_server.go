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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/pubsub/beta/pubsub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/beta"
)

// Server implements the gRPC interface for Topic.
type TopicServer struct{}

// ProtoToTopicMessageStoragePolicy converts a TopicMessageStoragePolicy resource from its proto representation.
func ProtoToPubsubBetaTopicMessageStoragePolicy(p *betapb.PubsubBetaTopicMessageStoragePolicy) *beta.TopicMessageStoragePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.TopicMessageStoragePolicy{}
	for _, r := range p.GetAllowedPersistenceRegions() {
		obj.AllowedPersistenceRegions = append(obj.AllowedPersistenceRegions, r)
	}
	return obj
}

// ProtoToTopic converts a Topic resource from its proto representation.
func ProtoToTopic(p *betapb.PubsubBetaTopic) *beta.Topic {
	obj := &beta.Topic{
		Name:                 dcl.StringOrNil(p.Name),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		MessageStoragePolicy: ProtoToPubsubBetaTopicMessageStoragePolicy(p.GetMessageStoragePolicy()),
		Project:              dcl.StringOrNil(p.Project),
	}
	return obj
}

// TopicMessageStoragePolicyToProto converts a TopicMessageStoragePolicy resource to its proto representation.
func PubsubBetaTopicMessageStoragePolicyToProto(o *beta.TopicMessageStoragePolicy) *betapb.PubsubBetaTopicMessageStoragePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.PubsubBetaTopicMessageStoragePolicy{}
	for _, r := range o.AllowedPersistenceRegions {
		p.AllowedPersistenceRegions = append(p.AllowedPersistenceRegions, r)
	}
	return p
}

// TopicToProto converts a Topic resource to its proto representation.
func TopicToProto(resource *beta.Topic) *betapb.PubsubBetaTopic {
	p := &betapb.PubsubBetaTopic{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		KmsKeyName:           dcl.ValueOrEmptyString(resource.KmsKeyName),
		MessageStoragePolicy: PubsubBetaTopicMessageStoragePolicyToProto(resource.MessageStoragePolicy),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) applyTopic(ctx context.Context, c *beta.Client, request *betapb.ApplyPubsubBetaTopicRequest) (*betapb.PubsubBetaTopic, error) {
	p := ProtoToTopic(request.GetResource())
	res, err := c.ApplyTopic(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TopicToProto(res)
	return r, nil
}

// ApplyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) ApplyPubsubBetaTopic(ctx context.Context, request *betapb.ApplyPubsubBetaTopicRequest) (*betapb.PubsubBetaTopic, error) {
	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTopic(ctx, cl, request)
}

// DeleteTopic handles the gRPC request by passing it to the underlying Topic Delete() method.
func (s *TopicServer) DeletePubsubBetaTopic(ctx context.Context, request *betapb.DeletePubsubBetaTopicRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTopic(ctx, ProtoToTopic(request.GetResource()))

}

// ListPubsubBetaTopic handles the gRPC request by passing it to the underlying TopicList() method.
func (s *TopicServer) ListPubsubBetaTopic(ctx context.Context, request *betapb.ListPubsubBetaTopicRequest) (*betapb.ListPubsubBetaTopicResponse, error) {
	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTopic(ctx, ProtoToTopic(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PubsubBetaTopic
	for _, r := range resources.Items {
		rp := TopicToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListPubsubBetaTopicResponse{Items: protos}, nil
}

func createConfigTopic(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
