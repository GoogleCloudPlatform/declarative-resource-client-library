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
	dataprocpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/dataproc_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
)

// Server implements the gRPC interface for AutoscalingPolicy.
type AutoscalingPolicyServer struct{}

// ProtoToAutoscalingPolicyBasicAlgorithm converts a AutoscalingPolicyBasicAlgorithm resource from its proto representation.
func ProtoToDataprocAutoscalingPolicyBasicAlgorithm(p *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithm) *dataproc.AutoscalingPolicyBasicAlgorithm {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicyBasicAlgorithm{
		YarnConfig:     ProtoToDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(p.GetYarnConfig()),
		CooldownPeriod: dcl.StringOrNil(p.CooldownPeriod),
	}
	return obj
}

// ProtoToAutoscalingPolicyBasicAlgorithmYarnConfig converts a AutoscalingPolicyBasicAlgorithmYarnConfig resource from its proto representation.
func ProtoToDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(p *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithmYarnConfig) *dataproc.AutoscalingPolicyBasicAlgorithmYarnConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.StringOrNil(p.GracefulDecommissionTimeout),
		ScaleUpFactor:               dcl.Float64OrNil(p.ScaleUpFactor),
		ScaleDownFactor:             dcl.Float64OrNil(p.ScaleDownFactor),
		ScaleUpMinWorkerFraction:    dcl.Float64OrNil(p.ScaleUpMinWorkerFraction),
		ScaleDownMinWorkerFraction:  dcl.Float64OrNil(p.ScaleDownMinWorkerFraction),
	}
	return obj
}

// ProtoToAutoscalingPolicyWorkerConfig converts a AutoscalingPolicyWorkerConfig resource from its proto representation.
func ProtoToDataprocAutoscalingPolicyWorkerConfig(p *dataprocpb.DataprocAutoscalingPolicyWorkerConfig) *dataproc.AutoscalingPolicyWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicyWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.MinInstances),
		MaxInstances: dcl.Int64OrNil(p.MaxInstances),
		Weight:       dcl.Int64OrNil(p.Weight),
	}
	return obj
}

// ProtoToAutoscalingPolicySecondaryWorkerConfig converts a AutoscalingPolicySecondaryWorkerConfig resource from its proto representation.
func ProtoToDataprocAutoscalingPolicySecondaryWorkerConfig(p *dataprocpb.DataprocAutoscalingPolicySecondaryWorkerConfig) *dataproc.AutoscalingPolicySecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.MinInstances),
		MaxInstances: dcl.Int64OrNil(p.MaxInstances),
		Weight:       dcl.Int64OrNil(p.Weight),
	}
	return obj
}

// ProtoToAutoscalingPolicy converts a AutoscalingPolicy resource from its proto representation.
func ProtoToAutoscalingPolicy(p *dataprocpb.DataprocAutoscalingPolicy) *dataproc.AutoscalingPolicy {
	obj := &dataproc.AutoscalingPolicy{
		Name:                  dcl.StringOrNil(p.Name),
		BasicAlgorithm:        ProtoToDataprocAutoscalingPolicyBasicAlgorithm(p.GetBasicAlgorithm()),
		WorkerConfig:          ProtoToDataprocAutoscalingPolicyWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAutoscalingPolicySecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
	}
	return obj
}

// AutoscalingPolicyBasicAlgorithmToProto converts a AutoscalingPolicyBasicAlgorithm resource to its proto representation.
func DataprocAutoscalingPolicyBasicAlgorithmToProto(o *dataproc.AutoscalingPolicyBasicAlgorithm) *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithm {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicyBasicAlgorithm{
		YarnConfig:     DataprocAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o.YarnConfig),
		CooldownPeriod: dcl.ValueOrEmptyString(o.CooldownPeriod),
	}
	return p
}

// AutoscalingPolicyBasicAlgorithmYarnConfigToProto converts a AutoscalingPolicyBasicAlgorithmYarnConfig resource to its proto representation.
func DataprocAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o *dataproc.AutoscalingPolicyBasicAlgorithmYarnConfig) *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithmYarnConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.ValueOrEmptyString(o.GracefulDecommissionTimeout),
		ScaleUpFactor:               dcl.ValueOrEmptyDouble(o.ScaleUpFactor),
		ScaleDownFactor:             dcl.ValueOrEmptyDouble(o.ScaleDownFactor),
		ScaleUpMinWorkerFraction:    dcl.ValueOrEmptyDouble(o.ScaleUpMinWorkerFraction),
		ScaleDownMinWorkerFraction:  dcl.ValueOrEmptyDouble(o.ScaleDownMinWorkerFraction),
	}
	return p
}

// AutoscalingPolicyWorkerConfigToProto converts a AutoscalingPolicyWorkerConfig resource to its proto representation.
func DataprocAutoscalingPolicyWorkerConfigToProto(o *dataproc.AutoscalingPolicyWorkerConfig) *dataprocpb.DataprocAutoscalingPolicyWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicyWorkerConfig{
		MinInstances: dcl.ValueOrEmptyInt64(o.MinInstances),
		MaxInstances: dcl.ValueOrEmptyInt64(o.MaxInstances),
		Weight:       dcl.ValueOrEmptyInt64(o.Weight),
	}
	return p
}

// AutoscalingPolicySecondaryWorkerConfigToProto converts a AutoscalingPolicySecondaryWorkerConfig resource to its proto representation.
func DataprocAutoscalingPolicySecondaryWorkerConfigToProto(o *dataproc.AutoscalingPolicySecondaryWorkerConfig) *dataprocpb.DataprocAutoscalingPolicySecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.ValueOrEmptyInt64(o.MinInstances),
		MaxInstances: dcl.ValueOrEmptyInt64(o.MaxInstances),
		Weight:       dcl.ValueOrEmptyInt64(o.Weight),
	}
	return p
}

// AutoscalingPolicyToProto converts a AutoscalingPolicy resource to its proto representation.
func AutoscalingPolicyToProto(resource *dataproc.AutoscalingPolicy) *dataprocpb.DataprocAutoscalingPolicy {
	p := &dataprocpb.DataprocAutoscalingPolicy{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		BasicAlgorithm:        DataprocAutoscalingPolicyBasicAlgorithmToProto(resource.BasicAlgorithm),
		WorkerConfig:          DataprocAutoscalingPolicyWorkerConfigToProto(resource.WorkerConfig),
		SecondaryWorkerConfig: DataprocAutoscalingPolicySecondaryWorkerConfigToProto(resource.SecondaryWorkerConfig),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) applyAutoscalingPolicy(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocAutoscalingPolicyRequest) (*dataprocpb.DataprocAutoscalingPolicy, error) {
	p := ProtoToAutoscalingPolicy(request.GetResource())
	res, err := c.ApplyAutoscalingPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalingPolicyToProto(res)
	return r, nil
}

// ApplyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) ApplyDataprocAutoscalingPolicy(ctx context.Context, request *dataprocpb.ApplyDataprocAutoscalingPolicyRequest) (*dataprocpb.DataprocAutoscalingPolicy, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAutoscalingPolicy(ctx, cl, request)
}

// DeleteAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Delete() method.
func (s *AutoscalingPolicyServer) DeleteDataprocAutoscalingPolicy(ctx context.Context, request *dataprocpb.DeleteDataprocAutoscalingPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscalingPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscalingPolicy(ctx, ProtoToAutoscalingPolicy(request.GetResource()))

}

// ListAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicyList() method.
func (s *AutoscalingPolicyServer) ListDataprocAutoscalingPolicy(ctx context.Context, request *dataprocpb.ListDataprocAutoscalingPolicyRequest) (*dataprocpb.ListDataprocAutoscalingPolicyResponse, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscalingPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocAutoscalingPolicy
	for _, r := range resources.Items {
		rp := AutoscalingPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &dataprocpb.ListDataprocAutoscalingPolicyResponse{Items: protos}, nil
}

func createConfigAutoscalingPolicy(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
