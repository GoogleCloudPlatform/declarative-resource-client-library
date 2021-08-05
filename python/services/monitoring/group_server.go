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
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// Server implements the gRPC interface for Group.
type GroupServer struct{}

// ProtoToGroup converts a Group resource from its proto representation.
func ProtoToGroup(p *monitoringpb.MonitoringGroup) *monitoring.Group {
	obj := &monitoring.Group{
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Filter:      dcl.StringOrNil(p.Filter),
		IsCluster:   dcl.Bool(p.IsCluster),
		Name:        dcl.StringOrNil(p.Name),
		ParentName:  dcl.StringOrNil(p.ParentName),
		Project:     dcl.StringOrNil(p.Project),
	}
	return obj
}

// GroupToProto converts a Group resource to its proto representation.
func GroupToProto(resource *monitoring.Group) *monitoringpb.MonitoringGroup {
	p := &monitoringpb.MonitoringGroup{
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		Filter:      dcl.ValueOrEmptyString(resource.Filter),
		IsCluster:   dcl.ValueOrEmptyBool(resource.IsCluster),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		ParentName:  dcl.ValueOrEmptyString(resource.ParentName),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) applyGroup(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringGroupRequest) (*monitoringpb.MonitoringGroup, error) {
	p := ProtoToGroup(request.GetResource())
	res, err := c.ApplyGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GroupToProto(res)
	return r, nil
}

// ApplyGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) ApplyMonitoringGroup(ctx context.Context, request *monitoringpb.ApplyMonitoringGroupRequest) (*monitoringpb.MonitoringGroup, error) {
	cl, err := createConfigGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyGroup(ctx, cl, request)
}

// DeleteGroup handles the gRPC request by passing it to the underlying Group Delete() method.
func (s *GroupServer) DeleteMonitoringGroup(ctx context.Context, request *monitoringpb.DeleteMonitoringGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGroup(ctx, ProtoToGroup(request.GetResource()))

}

// ListMonitoringGroup handles the gRPC request by passing it to the underlying GroupList() method.
func (s *GroupServer) ListMonitoringGroup(ctx context.Context, request *monitoringpb.ListMonitoringGroupRequest) (*monitoringpb.ListMonitoringGroupResponse, error) {
	cl, err := createConfigGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGroup(ctx, ProtoToGroup(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringGroup
	for _, r := range resources.Items {
		rp := GroupToProto(r)
		protos = append(protos, rp)
	}
	return &monitoringpb.ListMonitoringGroupResponse{Items: protos}, nil
}

func createConfigGroup(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
