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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/pubsub/alpha/pubsub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/alpha"
)

// Server implements the gRPC interface for Topic.
type TopicServer struct{}

// ProtoToTopicMessageStoragePolicy converts a TopicMessageStoragePolicy resource from its proto representation.
func ProtoToPubsubAlphaTopicMessageStoragePolicy(p *alphapb.PubsubAlphaTopicMessageStoragePolicy) *alpha.TopicMessageStoragePolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.TopicMessageStoragePolicy{}
	for _, r := range p.GetAllowedPersistenceRegions() {
		obj.AllowedPersistenceRegions = append(obj.AllowedPersistenceRegions, r)
	}
	return obj
}

// ProtoToTopic converts a Topic resource from its proto representation.
func ProtoToTopic(p *alphapb.PubsubAlphaTopic) *alpha.Topic {
	obj := &alpha.Topic{
		Name:                 dcl.StringOrNil(p.Name),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		MessageStoragePolicy: ProtoToPubsubAlphaTopicMessageStoragePolicy(p.GetMessageStoragePolicy()),
		Project:              dcl.StringOrNil(p.Project),
	}
	return obj
}

// TopicMessageStoragePolicyToProto converts a TopicMessageStoragePolicy resource to its proto representation.
func PubsubAlphaTopicMessageStoragePolicyToProto(o *alpha.TopicMessageStoragePolicy) *alphapb.PubsubAlphaTopicMessageStoragePolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.PubsubAlphaTopicMessageStoragePolicy{}
	for _, r := range o.AllowedPersistenceRegions {
		p.AllowedPersistenceRegions = append(p.AllowedPersistenceRegions, r)
	}
	return p
}

// TopicToProto converts a Topic resource to its proto representation.
func TopicToProto(resource *alpha.Topic) *alphapb.PubsubAlphaTopic {
	p := &alphapb.PubsubAlphaTopic{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		KmsKeyName:           dcl.ValueOrEmptyString(resource.KmsKeyName),
		MessageStoragePolicy: PubsubAlphaTopicMessageStoragePolicyToProto(resource.MessageStoragePolicy),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) applyTopic(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPubsubAlphaTopicRequest) (*alphapb.PubsubAlphaTopic, error) {
	p := ProtoToTopic(request.GetResource())
	res, err := c.ApplyTopic(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TopicToProto(res)
	return r, nil
}

// ApplyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) ApplyPubsubAlphaTopic(ctx context.Context, request *alphapb.ApplyPubsubAlphaTopicRequest) (*alphapb.PubsubAlphaTopic, error) {
	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTopic(ctx, cl, request)
}

// DeleteTopic handles the gRPC request by passing it to the underlying Topic Delete() method.
func (s *TopicServer) DeletePubsubAlphaTopic(ctx context.Context, request *alphapb.DeletePubsubAlphaTopicRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTopic(ctx, ProtoToTopic(request.GetResource()))

}

// ListPubsubAlphaTopic handles the gRPC request by passing it to the underlying TopicList() method.
func (s *TopicServer) ListPubsubAlphaTopic(ctx context.Context, request *alphapb.ListPubsubAlphaTopicRequest) (*alphapb.ListPubsubAlphaTopicResponse, error) {
	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTopic(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PubsubAlphaTopic
	for _, r := range resources.Items {
		rp := TopicToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListPubsubAlphaTopicResponse{Items: protos}, nil
}

func createConfigTopic(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
