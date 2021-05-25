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

// Server implements the gRPC interface for LogMetric.
type LogMetricServer struct{}

// ProtoToLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum converts a LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum(e loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum) *logging.LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum(n[len("LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorMetricKindEnum converts a LogMetricMetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorMetricKindEnum(e loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum) *logging.LogMetricMetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorMetricKindEnum(n[len("LogMetricMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorValueTypeEnum converts a LogMetricMetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorValueTypeEnum(e loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum) *logging.LogMetricMetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorValueTypeEnum(n[len("LogMetricMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorMetadataLaunchStageEnum converts a LogMetricMetricDescriptorMetadataLaunchStageEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorMetadataLaunchStageEnum(e loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum) *logging.LogMetricMetricDescriptorMetadataLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorMetadataLaunchStageEnum(n[len("LogMetricMetricDescriptorMetadataLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorLaunchStageEnum converts a LogMetricMetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorLaunchStageEnum(e loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum) *logging.LogMetricMetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorLaunchStageEnum(n[len("LogMetricMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptor converts a LogMetricMetricDescriptor resource from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptor(p *loggingpb.LoggingLogMetricMetricDescriptor) *logging.LogMetricMetricDescriptor {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricMetricDescriptor{
		Name:        dcl.StringOrNil(p.Name),
		Type:        dcl.StringOrNil(p.Type),
		MetricKind:  ProtoToLoggingLogMetricMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToLoggingLogMetricMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.Unit),
		Description: dcl.StringOrNil(p.Description),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Metadata:    ProtoToLoggingLogMetricMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToLoggingLogMetricMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
	}
	for _, r := range p.GetDescriptorLabels() {
		obj.DescriptorLabels = append(obj.DescriptorLabels, *ProtoToLoggingLogMetricMetricDescriptorDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorDescriptorLabels converts a LogMetricMetricDescriptorDescriptorLabels resource from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorDescriptorLabels(p *loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabels) *logging.LogMetricMetricDescriptorDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricMetricDescriptorDescriptorLabels{
		Key:         dcl.StringOrNil(p.Key),
		ValueType:   ProtoToLoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorMetadata converts a LogMetricMetricDescriptorMetadata resource from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorMetadata(p *loggingpb.LoggingLogMetricMetricDescriptorMetadata) *logging.LogMetricMetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricMetricDescriptorMetadata{
		LaunchStage:  ProtoToLoggingLogMetricMetricDescriptorMetadataLaunchStageEnum(p.GetLaunchStage()),
		SamplePeriod: dcl.StringOrNil(p.SamplePeriod),
		IngestDelay:  dcl.StringOrNil(p.IngestDelay),
	}
	return obj
}

// ProtoToLogMetricBucketOptions converts a LogMetricBucketOptions resource from its proto representation.
func ProtoToLoggingLogMetricBucketOptions(p *loggingpb.LoggingLogMetricBucketOptions) *logging.LogMetricBucketOptions {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptions{
		LinearBuckets:      ProtoToLoggingLogMetricBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToLoggingLogMetricBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToLoggingLogMetricBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsLinearBuckets converts a LogMetricBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToLoggingLogMetricBucketOptionsLinearBuckets(p *loggingpb.LoggingLogMetricBucketOptionsLinearBuckets) *logging.LogMetricBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExponentialBuckets converts a LogMetricBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToLoggingLogMetricBucketOptionsExponentialBuckets(p *loggingpb.LoggingLogMetricBucketOptionsExponentialBuckets) *logging.LogMetricBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExplicitBuckets converts a LogMetricBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToLoggingLogMetricBucketOptionsExplicitBuckets(p *loggingpb.LoggingLogMetricBucketOptionsExplicitBuckets) *logging.LogMetricBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToLogMetric converts a LogMetric resource from its proto representation.
func ProtoToLogMetric(p *loggingpb.LoggingLogMetric) *logging.LogMetric {
	obj := &logging.LogMetric{
		Name:             dcl.StringOrNil(p.Name),
		Description:      dcl.StringOrNil(p.Description),
		Filter:           dcl.StringOrNil(p.Filter),
		Disabled:         dcl.Bool(p.Disabled),
		MetricDescriptor: ProtoToLoggingLogMetricMetricDescriptor(p.GetMetricDescriptor()),
		ValueExtractor:   dcl.StringOrNil(p.ValueExtractor),
		BucketOptions:    ProtoToLoggingLogMetricBucketOptions(p.GetBucketOptions()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Project:          dcl.StringOrNil(p.Project),
	}
	return obj
}

// LogMetricMetricDescriptorDescriptorLabelsValueTypeEnumToProto converts a LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnumToProto(e *logging.LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum) loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum_value["LogMetricMetricDescriptorDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnum(0)
}

// LogMetricMetricDescriptorMetricKindEnumToProto converts a LogMetricMetricDescriptorMetricKindEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorMetricKindEnumToProto(e *logging.LogMetricMetricDescriptorMetricKindEnum) loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum_value["LogMetricMetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum(0)
}

// LogMetricMetricDescriptorValueTypeEnumToProto converts a LogMetricMetricDescriptorValueTypeEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorValueTypeEnumToProto(e *logging.LogMetricMetricDescriptorValueTypeEnum) loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum_value["LogMetricMetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum(0)
}

// LogMetricMetricDescriptorMetadataLaunchStageEnumToProto converts a LogMetricMetricDescriptorMetadataLaunchStageEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorMetadataLaunchStageEnumToProto(e *logging.LogMetricMetricDescriptorMetadataLaunchStageEnum) loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum_value["LogMetricMetricDescriptorMetadataLaunchStageEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorMetadataLaunchStageEnum(0)
}

// LogMetricMetricDescriptorLaunchStageEnumToProto converts a LogMetricMetricDescriptorLaunchStageEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorLaunchStageEnumToProto(e *logging.LogMetricMetricDescriptorLaunchStageEnum) loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum_value["LogMetricMetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum(0)
}

// LogMetricMetricDescriptorToProto converts a LogMetricMetricDescriptor resource to its proto representation.
func LoggingLogMetricMetricDescriptorToProto(o *logging.LogMetricMetricDescriptor) *loggingpb.LoggingLogMetricMetricDescriptor {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricMetricDescriptor{
		Name:        dcl.ValueOrEmptyString(o.Name),
		Type:        dcl.ValueOrEmptyString(o.Type),
		MetricKind:  LoggingLogMetricMetricDescriptorMetricKindEnumToProto(o.MetricKind),
		ValueType:   LoggingLogMetricMetricDescriptorValueTypeEnumToProto(o.ValueType),
		Unit:        dcl.ValueOrEmptyString(o.Unit),
		Description: dcl.ValueOrEmptyString(o.Description),
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		Metadata:    LoggingLogMetricMetricDescriptorMetadataToProto(o.Metadata),
		LaunchStage: LoggingLogMetricMetricDescriptorLaunchStageEnumToProto(o.LaunchStage),
	}
	for _, r := range o.DescriptorLabels {
		p.DescriptorLabels = append(p.DescriptorLabels, LoggingLogMetricMetricDescriptorDescriptorLabelsToProto(&r))
	}
	for _, r := range o.MonitoredResourceTypes {
		p.MonitoredResourceTypes = append(p.MonitoredResourceTypes, r)
	}
	return p
}

// LogMetricMetricDescriptorDescriptorLabelsToProto converts a LogMetricMetricDescriptorDescriptorLabels resource to its proto representation.
func LoggingLogMetricMetricDescriptorDescriptorLabelsToProto(o *logging.LogMetricMetricDescriptorDescriptorLabels) *loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricMetricDescriptorDescriptorLabels{
		Key:         dcl.ValueOrEmptyString(o.Key),
		ValueType:   LoggingLogMetricMetricDescriptorDescriptorLabelsValueTypeEnumToProto(o.ValueType),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// LogMetricMetricDescriptorMetadataToProto converts a LogMetricMetricDescriptorMetadata resource to its proto representation.
func LoggingLogMetricMetricDescriptorMetadataToProto(o *logging.LogMetricMetricDescriptorMetadata) *loggingpb.LoggingLogMetricMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricMetricDescriptorMetadata{
		LaunchStage:  LoggingLogMetricMetricDescriptorMetadataLaunchStageEnumToProto(o.LaunchStage),
		SamplePeriod: dcl.ValueOrEmptyString(o.SamplePeriod),
		IngestDelay:  dcl.ValueOrEmptyString(o.IngestDelay),
	}
	return p
}

// LogMetricBucketOptionsToProto converts a LogMetricBucketOptions resource to its proto representation.
func LoggingLogMetricBucketOptionsToProto(o *logging.LogMetricBucketOptions) *loggingpb.LoggingLogMetricBucketOptions {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptions{
		LinearBuckets:      LoggingLogMetricBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: LoggingLogMetricBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    LoggingLogMetricBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// LogMetricBucketOptionsLinearBucketsToProto converts a LogMetricBucketOptionsLinearBuckets resource to its proto representation.
func LoggingLogMetricBucketOptionsLinearBucketsToProto(o *logging.LogMetricBucketOptionsLinearBuckets) *loggingpb.LoggingLogMetricBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// LogMetricBucketOptionsExponentialBucketsToProto converts a LogMetricBucketOptionsExponentialBuckets resource to its proto representation.
func LoggingLogMetricBucketOptionsExponentialBucketsToProto(o *logging.LogMetricBucketOptionsExponentialBuckets) *loggingpb.LoggingLogMetricBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// LogMetricBucketOptionsExplicitBucketsToProto converts a LogMetricBucketOptionsExplicitBuckets resource to its proto representation.
func LoggingLogMetricBucketOptionsExplicitBucketsToProto(o *logging.LogMetricBucketOptionsExplicitBuckets) *loggingpb.LoggingLogMetricBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// LogMetricToProto converts a LogMetric resource to its proto representation.
func LogMetricToProto(resource *logging.LogMetric) *loggingpb.LoggingLogMetric {
	p := &loggingpb.LoggingLogMetric{
		Name:             dcl.ValueOrEmptyString(resource.Name),
		Description:      dcl.ValueOrEmptyString(resource.Description),
		Filter:           dcl.ValueOrEmptyString(resource.Filter),
		Disabled:         dcl.ValueOrEmptyBool(resource.Disabled),
		MetricDescriptor: LoggingLogMetricMetricDescriptorToProto(resource.MetricDescriptor),
		ValueExtractor:   dcl.ValueOrEmptyString(resource.ValueExtractor),
		BucketOptions:    LoggingLogMetricBucketOptionsToProto(resource.BucketOptions),
		CreateTime:       dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:       dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyLogMetric handles the gRPC request by passing it to the underlying LogMetric Apply() method.
func (s *LogMetricServer) applyLogMetric(ctx context.Context, c *logging.Client, request *loggingpb.ApplyLoggingLogMetricRequest) (*loggingpb.LoggingLogMetric, error) {
	p := ProtoToLogMetric(request.GetResource())
	res, err := c.ApplyLogMetric(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogMetricToProto(res)
	return r, nil
}

// ApplyLogMetric handles the gRPC request by passing it to the underlying LogMetric Apply() method.
func (s *LogMetricServer) ApplyLoggingLogMetric(ctx context.Context, request *loggingpb.ApplyLoggingLogMetricRequest) (*loggingpb.LoggingLogMetric, error) {
	cl, err := createConfigLogMetric(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyLogMetric(ctx, cl, request)
}

// DeleteLogMetric handles the gRPC request by passing it to the underlying LogMetric Delete() method.
func (s *LogMetricServer) DeleteLoggingLogMetric(ctx context.Context, request *loggingpb.DeleteLoggingLogMetricRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogMetric(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogMetric(ctx, ProtoToLogMetric(request.GetResource()))

}

// ListLogMetric handles the gRPC request by passing it to the underlying LogMetricList() method.
func (s *LogMetricServer) ListLoggingLogMetric(ctx context.Context, request *loggingpb.ListLoggingLogMetricRequest) (*loggingpb.ListLoggingLogMetricResponse, error) {
	cl, err := createConfigLogMetric(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogMetric(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*loggingpb.LoggingLogMetric
	for _, r := range resources.Items {
		rp := LogMetricToProto(r)
		protos = append(protos, rp)
	}
	return &loggingpb.ListLoggingLogMetricResponse{Items: protos}, nil
}

func createConfigLogMetric(ctx context.Context, service_account_file string) (*logging.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return logging.NewClient(conf), nil
}
