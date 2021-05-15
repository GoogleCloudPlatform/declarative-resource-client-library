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
	apigeepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/apigee_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
)

// Server implements the gRPC interface for Attachment.
type AttachmentServer struct{}

// ProtoToAttachment converts a Attachment resource from its proto representation.
func ProtoToAttachment(p *apigeepb.ApigeeAttachment) *apigee.Attachment {
	obj := &apigee.Attachment{
		Name:         dcl.StringOrNil(p.Name),
		Environment:  dcl.StringOrNil(p.Environment),
		CreatedAt:    dcl.Int64OrNil(p.CreatedAt),
		Organization: dcl.StringOrNil(p.Organization),
		Envgroup:     dcl.StringOrNil(p.Envgroup),
	}
	return obj
}

// AttachmentToProto converts a Attachment resource to its proto representation.
func AttachmentToProto(resource *apigee.Attachment) *apigeepb.ApigeeAttachment {
	p := &apigeepb.ApigeeAttachment{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		Environment:  dcl.ValueOrEmptyString(resource.Environment),
		CreatedAt:    dcl.ValueOrEmptyInt64(resource.CreatedAt),
		Organization: dcl.ValueOrEmptyString(resource.Organization),
		Envgroup:     dcl.ValueOrEmptyString(resource.Envgroup),
	}

	return p
}

// ApplyAttachment handles the gRPC request by passing it to the underlying Attachment Apply() method.
func (s *AttachmentServer) applyAttachment(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeAttachmentRequest) (*apigeepb.ApigeeAttachment, error) {
	p := ProtoToAttachment(request.GetResource())
	res, err := c.ApplyAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AttachmentToProto(res)
	return r, nil
}

// ApplyAttachment handles the gRPC request by passing it to the underlying Attachment Apply() method.
func (s *AttachmentServer) ApplyApigeeAttachment(ctx context.Context, request *apigeepb.ApplyApigeeAttachmentRequest) (*apigeepb.ApigeeAttachment, error) {
	cl, err := createConfigAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAttachment(ctx, cl, request)
}

// DeleteAttachment handles the gRPC request by passing it to the underlying Attachment Delete() method.
func (s *AttachmentServer) DeleteApigeeAttachment(ctx context.Context, request *apigeepb.DeleteApigeeAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAttachment(ctx, ProtoToAttachment(request.GetResource()))

}

// ListAttachment handles the gRPC request by passing it to the underlying AttachmentList() method.
func (s *AttachmentServer) ListApigeeAttachment(ctx context.Context, request *apigeepb.ListApigeeAttachmentRequest) (*apigeepb.ListApigeeAttachmentResponse, error) {
	cl, err := createConfigAttachment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAttachment(ctx, request.Organization, request.Envgroup)
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeAttachment
	for _, r := range resources.Items {
		rp := AttachmentToProto(r)
		protos = append(protos, rp)
	}
	return &apigeepb.ListApigeeAttachmentResponse{Items: protos}, nil
}

func createConfigAttachment(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
