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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/beta/logging_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta"
)

// Server implements the gRPC interface for LogExclusion.
type LogExclusionServer struct{}

// ProtoToLogExclusion converts a LogExclusion resource from its proto representation.
func ProtoToLogExclusion(p *betapb.LoggingBetaLogExclusion) *beta.LogExclusion {
	obj := &beta.LogExclusion{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Filter:      dcl.StringOrNil(p.Filter),
		Disabled:    dcl.Bool(p.Disabled),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Parent:      dcl.StringOrNil(p.Parent),
	}
	return obj
}

// LogExclusionToProto converts a LogExclusion resource to its proto representation.
func LogExclusionToProto(resource *beta.LogExclusion) *betapb.LoggingBetaLogExclusion {
	p := &betapb.LoggingBetaLogExclusion{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Filter:      dcl.ValueOrEmptyString(resource.Filter),
		Disabled:    dcl.ValueOrEmptyBool(resource.Disabled),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Parent:      dcl.ValueOrEmptyString(resource.Parent),
	}

	return p
}

// ApplyLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Apply() method.
func (s *LogExclusionServer) applyLogExclusion(ctx context.Context, c *beta.Client, request *betapb.ApplyLoggingBetaLogExclusionRequest) (*betapb.LoggingBetaLogExclusion, error) {
	p := ProtoToLogExclusion(request.GetResource())
	res, err := c.ApplyLogExclusion(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogExclusionToProto(res)
	return r, nil
}

// ApplyLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Apply() method.
func (s *LogExclusionServer) ApplyLoggingBetaLogExclusion(ctx context.Context, request *betapb.ApplyLoggingBetaLogExclusionRequest) (*betapb.LoggingBetaLogExclusion, error) {
	cl, err := createConfigLogExclusion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyLogExclusion(ctx, cl, request)
}

// DeleteLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Delete() method.
func (s *LogExclusionServer) DeleteLoggingBetaLogExclusion(ctx context.Context, request *betapb.DeleteLoggingBetaLogExclusionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogExclusion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogExclusion(ctx, ProtoToLogExclusion(request.GetResource()))

}

// ListLoggingBetaLogExclusion handles the gRPC request by passing it to the underlying LogExclusionList() method.
func (s *LogExclusionServer) ListLoggingBetaLogExclusion(ctx context.Context, request *betapb.ListLoggingBetaLogExclusionRequest) (*betapb.ListLoggingBetaLogExclusionResponse, error) {
	cl, err := createConfigLogExclusion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogExclusion(ctx, ProtoToLogExclusion(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.LoggingBetaLogExclusion
	for _, r := range resources.Items {
		rp := LogExclusionToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListLoggingBetaLogExclusionResponse{Items: protos}, nil
}

func createConfigLogExclusion(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
