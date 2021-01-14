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

// Server implements the gRPC interface for NotificationChannel.
type NotificationChannelServer struct{}

// ProtoToNotificationChannelVerificationStatusEnum converts a NotificationChannelVerificationStatusEnum enum from its proto representation.
func ProtoToMonitoringNotificationChannelVerificationStatusEnum(e monitoringpb.MonitoringNotificationChannelVerificationStatusEnum) *monitoring.NotificationChannelVerificationStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringNotificationChannelVerificationStatusEnum_name[int32(e)]; ok {
		e := monitoring.NotificationChannelVerificationStatusEnum(n[len("NotificationChannelVerificationStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotificationChannel converts a NotificationChannel resource from its proto representation.
func ProtoToNotificationChannel(p *monitoringpb.MonitoringNotificationChannel) *monitoring.NotificationChannel {
	obj := &monitoring.NotificationChannel{
		Description:        dcl.StringOrNil(p.Description),
		DisplayName:        dcl.StringOrNil(p.DisplayName),
		Enabled:            dcl.Bool(p.Enabled),
		Name:               dcl.StringOrNil(p.Name),
		Type:               dcl.StringOrNil(p.Type),
		VerificationStatus: ProtoToMonitoringNotificationChannelVerificationStatusEnum(p.GetVerificationStatus()),
		Project:            dcl.StringOrNil(p.Project),
	}
	return obj
}

// NotificationChannelVerificationStatusEnumToProto converts a NotificationChannelVerificationStatusEnum enum to its proto representation.
func MonitoringNotificationChannelVerificationStatusEnumToProto(e *monitoring.NotificationChannelVerificationStatusEnum) monitoringpb.MonitoringNotificationChannelVerificationStatusEnum {
	if e == nil {
		return monitoringpb.MonitoringNotificationChannelVerificationStatusEnum(0)
	}
	if v, ok := monitoringpb.MonitoringNotificationChannelVerificationStatusEnum_value["NotificationChannelVerificationStatusEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringNotificationChannelVerificationStatusEnum(v)
	}
	return monitoringpb.MonitoringNotificationChannelVerificationStatusEnum(0)
}

// NotificationChannelToProto converts a NotificationChannel resource to its proto representation.
func NotificationChannelToProto(resource *monitoring.NotificationChannel) *monitoringpb.MonitoringNotificationChannel {
	p := &monitoringpb.MonitoringNotificationChannel{
		Description:        dcl.ValueOrEmptyString(resource.Description),
		DisplayName:        dcl.ValueOrEmptyString(resource.DisplayName),
		Enabled:            dcl.ValueOrEmptyBool(resource.Enabled),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Type:               dcl.ValueOrEmptyString(resource.Type),
		VerificationStatus: MonitoringNotificationChannelVerificationStatusEnumToProto(resource.VerificationStatus),
		Project:            dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannel Apply() method.
func (s *NotificationChannelServer) applyNotificationChannel(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringNotificationChannelRequest) (*monitoringpb.MonitoringNotificationChannel, error) {
	p := ProtoToNotificationChannel(request.GetResource())
	res, err := c.ApplyNotificationChannel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NotificationChannelToProto(res)
	return r, nil
}

// ApplyNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannel Apply() method.
func (s *NotificationChannelServer) ApplyMonitoringNotificationChannel(ctx context.Context, request *monitoringpb.ApplyMonitoringNotificationChannelRequest) (*monitoringpb.MonitoringNotificationChannel, error) {
	cl, err := createConfigNotificationChannel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNotificationChannel(ctx, cl, request)
}

// DeleteNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannel Delete() method.
func (s *NotificationChannelServer) DeleteMonitoringNotificationChannel(ctx context.Context, request *monitoringpb.DeleteMonitoringNotificationChannelRequest) (*emptypb.Empty, error) {
	cl, err := createConfigNotificationChannel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNotificationChannel(ctx, ProtoToNotificationChannel(request.GetResource()))
}

// ListNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannelList() method.
func (s *NotificationChannelServer) ListMonitoringNotificationChannel(ctx context.Context, request *monitoringpb.ListMonitoringNotificationChannelRequest) (*monitoringpb.ListMonitoringNotificationChannelResponse, error) {
	cl, err := createConfigNotificationChannel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNotificationChannel(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringNotificationChannel
	for _, r := range resources.Items {
		rp := NotificationChannelToProto(r)
		protos = append(protos, rp)
	}
	return &monitoringpb.ListMonitoringNotificationChannelResponse{Items: protos}, nil
}

func createConfigNotificationChannel(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
