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
	loggingpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/logging_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging"
)

// Server implements the gRPC interface for LogExclusion.
type LogExclusionServer struct{}

// ProtoToLogExclusion converts a LogExclusion resource from its proto representation.
func ProtoToLogExclusion(p *loggingpb.LoggingLogExclusion) *logging.LogExclusion {
	obj := &logging.LogExclusion{
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
func LogExclusionToProto(resource *logging.LogExclusion) *loggingpb.LoggingLogExclusion {
	p := &loggingpb.LoggingLogExclusion{
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
func (s *LogExclusionServer) applyLogExclusion(ctx context.Context, c *logging.Client, request *loggingpb.ApplyLoggingLogExclusionRequest) (*loggingpb.LoggingLogExclusion, error) {
	p := ProtoToLogExclusion(request.GetResource())
	res, err := c.ApplyLogExclusion(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogExclusionToProto(res)
	return r, nil
}

// ApplyLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Apply() method.
func (s *LogExclusionServer) ApplyLoggingLogExclusion(ctx context.Context, request *loggingpb.ApplyLoggingLogExclusionRequest) (*loggingpb.LoggingLogExclusion, error) {
	cl, err := createConfigLogExclusion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyLogExclusion(ctx, cl, request)
}

// DeleteLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Delete() method.
func (s *LogExclusionServer) DeleteLoggingLogExclusion(ctx context.Context, request *loggingpb.DeleteLoggingLogExclusionRequest) (*emptypb.Empty, error) {
	cl, err := createConfigLogExclusion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogExclusion(ctx, ProtoToLogExclusion(request.GetResource()))
}

// ListLogExclusion handles the gRPC request by passing it to the underlying LogExclusionList() method.
func (s *LogExclusionServer) ListLoggingLogExclusion(ctx context.Context, request *loggingpb.ListLoggingLogExclusionRequest) (*loggingpb.ListLoggingLogExclusionResponse, error) {
	cl, err := createConfigLogExclusion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogExclusion(ctx, request.Parent)
	if err != nil {
		return nil, err
	}
	var protos []*loggingpb.LoggingLogExclusion
	for _, r := range resources.Items {
		rp := LogExclusionToProto(r)
		protos = append(protos, rp)
	}
	return &loggingpb.ListLoggingLogExclusionResponse{Items: protos}, nil
}

func createConfigLogExclusion(ctx context.Context, service_account_file string) (*logging.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return logging.NewClient(conf), nil
}
