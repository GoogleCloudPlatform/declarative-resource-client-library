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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// Server implements the gRPC interface for MetricDescriptor.
type MetricDescriptorServer struct{}

// ProtoToMetricDescriptorLabelsValueTypeEnum converts a MetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorLabelsValueTypeEnum(e alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum) *alpha.MetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorLabelsValueTypeEnum(n[len("MonitoringAlphaMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetricKindEnum converts a MetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorMetricKindEnum(e alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum) *alpha.MetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorMetricKindEnum(n[len("MonitoringAlphaMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorValueTypeEnum converts a MetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorValueTypeEnum(e alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum) *alpha.MetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorValueTypeEnum(n[len("MonitoringAlphaMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetadataLaunchStageEnum converts a MetricDescriptorMetadataLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(e alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum) *alpha.MetricDescriptorMetadataLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorMetadataLaunchStageEnum(n[len("MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLaunchStageEnum converts a MetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorLaunchStageEnum(e alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum) *alpha.MetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorLaunchStageEnum(n[len("MonitoringAlphaMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLabels converts a MetricDescriptorLabels resource from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorLabels(p *alphapb.MonitoringAlphaMetricDescriptorLabels) *alpha.MetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.MetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.Key),
		ValueType:   ProtoToMonitoringAlphaMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToMetricDescriptorMetadata converts a MetricDescriptorMetadata resource from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorMetadata(p *alphapb.MonitoringAlphaMetricDescriptorMetadata) *alpha.MetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.MetricDescriptorMetadata{
		LaunchStage:  ProtoToMonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(p.GetLaunchStage()),
		SamplePeriod: dcl.StringOrNil(p.SamplePeriod),
		IngestDelay:  dcl.StringOrNil(p.IngestDelay),
	}
	return obj
}

// ProtoToMetricDescriptor converts a MetricDescriptor resource from its proto representation.
func ProtoToMetricDescriptor(p *alphapb.MonitoringAlphaMetricDescriptor) *alpha.MetricDescriptor {
	obj := &alpha.MetricDescriptor{
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		Type:        dcl.StringOrNil(p.Type),
		MetricKind:  ProtoToMonitoringAlphaMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToMonitoringAlphaMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.Unit),
		Description: dcl.StringOrNil(p.Description),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Metadata:    ProtoToMonitoringAlphaMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToMonitoringAlphaMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToMonitoringAlphaMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// MetricDescriptorLabelsValueTypeEnumToProto converts a MetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorLabelsValueTypeEnumToProto(e *alpha.MetricDescriptorLabelsValueTypeEnum) alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum_value["MetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum(0)
}

// MetricDescriptorMetricKindEnumToProto converts a MetricDescriptorMetricKindEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorMetricKindEnumToProto(e *alpha.MetricDescriptorMetricKindEnum) alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum_value["MetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum(0)
}

// MetricDescriptorValueTypeEnumToProto converts a MetricDescriptorValueTypeEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorValueTypeEnumToProto(e *alpha.MetricDescriptorValueTypeEnum) alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum_value["MetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum(0)
}

// MetricDescriptorMetadataLaunchStageEnumToProto converts a MetricDescriptorMetadataLaunchStageEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorMetadataLaunchStageEnumToProto(e *alpha.MetricDescriptorMetadataLaunchStageEnum) alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum_value["MetricDescriptorMetadataLaunchStageEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(0)
}

// MetricDescriptorLaunchStageEnumToProto converts a MetricDescriptorLaunchStageEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorLaunchStageEnumToProto(e *alpha.MetricDescriptorLaunchStageEnum) alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum_value["MetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum(0)
}

// MetricDescriptorLabelsToProto converts a MetricDescriptorLabels resource to its proto representation.
func MonitoringAlphaMetricDescriptorLabelsToProto(o *alpha.MetricDescriptorLabels) *alphapb.MonitoringAlphaMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaMetricDescriptorLabels{
		Key:         dcl.ValueOrEmptyString(o.Key),
		ValueType:   MonitoringAlphaMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// MetricDescriptorMetadataToProto converts a MetricDescriptorMetadata resource to its proto representation.
func MonitoringAlphaMetricDescriptorMetadataToProto(o *alpha.MetricDescriptorMetadata) *alphapb.MonitoringAlphaMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaMetricDescriptorMetadata{
		LaunchStage:  MonitoringAlphaMetricDescriptorMetadataLaunchStageEnumToProto(o.LaunchStage),
		SamplePeriod: dcl.ValueOrEmptyString(o.SamplePeriod),
		IngestDelay:  dcl.ValueOrEmptyString(o.IngestDelay),
	}
	return p
}

// MetricDescriptorToProto converts a MetricDescriptor resource to its proto representation.
func MetricDescriptorToProto(resource *alpha.MetricDescriptor) *alphapb.MonitoringAlphaMetricDescriptor {
	p := &alphapb.MonitoringAlphaMetricDescriptor{
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		Type:        dcl.ValueOrEmptyString(resource.Type),
		MetricKind:  MonitoringAlphaMetricDescriptorMetricKindEnumToProto(resource.MetricKind),
		ValueType:   MonitoringAlphaMetricDescriptorValueTypeEnumToProto(resource.ValueType),
		Unit:        dcl.ValueOrEmptyString(resource.Unit),
		Description: dcl.ValueOrEmptyString(resource.Description),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		Metadata:    MonitoringAlphaMetricDescriptorMetadataToProto(resource.Metadata),
		LaunchStage: MonitoringAlphaMetricDescriptorLaunchStageEnumToProto(resource.LaunchStage),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Labels {
		p.Labels = append(p.Labels, MonitoringAlphaMetricDescriptorLabelsToProto(&r))
	}
	for _, r := range resource.MonitoredResourceTypes {
		p.MonitoredResourceTypes = append(p.MonitoredResourceTypes, r)
	}

	return p
}

// ApplyMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) applyMetricDescriptor(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaMetricDescriptorRequest) (*alphapb.MonitoringAlphaMetricDescriptor, error) {
	p := ProtoToMetricDescriptor(request.GetResource())
	res, err := c.ApplyMetricDescriptor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricDescriptorToProto(res)
	return r, nil
}

// ApplyMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) ApplyMonitoringAlphaMetricDescriptor(ctx context.Context, request *alphapb.ApplyMonitoringAlphaMetricDescriptorRequest) (*alphapb.MonitoringAlphaMetricDescriptor, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyMetricDescriptor(ctx, cl, request)
}

// DeleteMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Delete() method.
func (s *MetricDescriptorServer) DeleteMonitoringAlphaMetricDescriptor(ctx context.Context, request *alphapb.DeleteMonitoringAlphaMetricDescriptorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetricDescriptor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetricDescriptor(ctx, ProtoToMetricDescriptor(request.GetResource()))

}

// ListMonitoringAlphaMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptorList() method.
func (s *MetricDescriptorServer) ListMonitoringAlphaMetricDescriptor(ctx context.Context, request *alphapb.ListMonitoringAlphaMetricDescriptorRequest) (*alphapb.ListMonitoringAlphaMetricDescriptorResponse, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetricDescriptor(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaMetricDescriptor
	for _, r := range resources.Items {
		rp := MetricDescriptorToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListMonitoringAlphaMetricDescriptorResponse{Items: protos}, nil
}

func createConfigMetricDescriptor(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
