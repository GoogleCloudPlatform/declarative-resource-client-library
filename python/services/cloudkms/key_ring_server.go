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
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudkmspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/cloudkms_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms"
)

// Server implements the gRPC interface for KeyRing.
type KeyRingServer struct{}

// ProtoToKeyRing converts a KeyRing resource from its proto representation.
func ProtoToKeyRing(p *cloudkmspb.CloudkmsKeyRing) *cloudkms.KeyRing {
	obj := &cloudkms.KeyRing{
		Name:       dcl.StringOrNil(p.Name),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	return obj
}

// KeyRingToProto converts a KeyRing resource to its proto representation.
func KeyRingToProto(resource *cloudkms.KeyRing) *cloudkmspb.CloudkmsKeyRing {
	p := &cloudkmspb.CloudkmsKeyRing{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyKeyRing handles the gRPC request by passing it to the underlying KeyRing Apply() method.
func (s *KeyRingServer) applyKeyRing(ctx context.Context, c *cloudkms.Client, request *cloudkmspb.ApplyCloudkmsKeyRingRequest) (*cloudkmspb.CloudkmsKeyRing, error) {
	p := ProtoToKeyRing(request.GetResource())
	res, err := c.ApplyKeyRing(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyRingToProto(res)
	return r, nil
}

// ApplyKeyRing handles the gRPC request by passing it to the underlying KeyRing Apply() method.
func (s *KeyRingServer) ApplyCloudkmsKeyRing(ctx context.Context, request *cloudkmspb.ApplyCloudkmsKeyRingRequest) (*cloudkmspb.CloudkmsKeyRing, error) {
	cl, err := createConfigKeyRing(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyKeyRing(ctx, cl, request)
}

// DeleteKeyRing handles the gRPC request by passing it to the underlying KeyRing Delete() method.
func (s *KeyRingServer) DeleteCloudkmsKeyRing(ctx context.Context, request *cloudkmspb.DeleteCloudkmsKeyRingRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for KeyRing")

}

// ListCloudkmsKeyRing handles the gRPC request by passing it to the underlying KeyRingList() method.
func (s *KeyRingServer) ListCloudkmsKeyRing(ctx context.Context, request *cloudkmspb.ListCloudkmsKeyRingRequest) (*cloudkmspb.ListCloudkmsKeyRingResponse, error) {
	cl, err := createConfigKeyRing(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKeyRing(ctx, ProtoToKeyRing(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*cloudkmspb.CloudkmsKeyRing
	for _, r := range resources.Items {
		rp := KeyRingToProto(r)
		protos = append(protos, rp)
	}
	return &cloudkmspb.ListCloudkmsKeyRingResponse{Items: protos}, nil
}

func createConfigKeyRing(ctx context.Context, service_account_file string) (*cloudkms.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudkms.NewClient(conf), nil
}
