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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/alpha/dataproc_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
)

// Server implements the gRPC interface for AutoscalingPolicy.
type AutoscalingPolicyServer struct{}

// ProtoToAutoscalingPolicyBasicAlgorithm converts a AutoscalingPolicyBasicAlgorithm resource from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithm(p *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithm) *alpha.AutoscalingPolicyBasicAlgorithm {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicyBasicAlgorithm{
		YarnConfig:     ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig(p.GetYarnConfig()),
		CooldownPeriod: dcl.StringOrNil(p.CooldownPeriod),
	}
	return obj
}

// ProtoToAutoscalingPolicyBasicAlgorithmYarnConfig converts a AutoscalingPolicyBasicAlgorithmYarnConfig resource from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig(p *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig) *alpha.AutoscalingPolicyBasicAlgorithmYarnConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.StringOrNil(p.GracefulDecommissionTimeout),
		ScaleUpFactor:               dcl.Float64OrNil(p.ScaleUpFactor),
		ScaleDownFactor:             dcl.Float64OrNil(p.ScaleDownFactor),
		ScaleUpMinWorkerFraction:    dcl.Float64OrNil(p.ScaleUpMinWorkerFraction),
		ScaleDownMinWorkerFraction:  dcl.Float64OrNil(p.ScaleDownMinWorkerFraction),
	}
	return obj
}

// ProtoToAutoscalingPolicyWorkerConfig converts a AutoscalingPolicyWorkerConfig resource from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicyWorkerConfig(p *alphapb.DataprocAlphaAutoscalingPolicyWorkerConfig) *alpha.AutoscalingPolicyWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicyWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.MinInstances),
		MaxInstances: dcl.Int64OrNil(p.MaxInstances),
		Weight:       dcl.Int64OrNil(p.Weight),
	}
	return obj
}

// ProtoToAutoscalingPolicySecondaryWorkerConfig converts a AutoscalingPolicySecondaryWorkerConfig resource from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicySecondaryWorkerConfig(p *alphapb.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig) *alpha.AutoscalingPolicySecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.MinInstances),
		MaxInstances: dcl.Int64OrNil(p.MaxInstances),
		Weight:       dcl.Int64OrNil(p.Weight),
	}
	return obj
}

// ProtoToAutoscalingPolicy converts a AutoscalingPolicy resource from its proto representation.
func ProtoToAutoscalingPolicy(p *alphapb.DataprocAlphaAutoscalingPolicy) *alpha.AutoscalingPolicy {
	obj := &alpha.AutoscalingPolicy{
		Name:                  dcl.StringOrNil(p.Name),
		BasicAlgorithm:        ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithm(p.GetBasicAlgorithm()),
		WorkerConfig:          ProtoToDataprocAlphaAutoscalingPolicyWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAlphaAutoscalingPolicySecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
	}
	return obj
}

// AutoscalingPolicyBasicAlgorithmToProto converts a AutoscalingPolicyBasicAlgorithm resource to its proto representation.
func DataprocAlphaAutoscalingPolicyBasicAlgorithmToProto(o *alpha.AutoscalingPolicyBasicAlgorithm) *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithm {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithm{
		YarnConfig:     DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o.YarnConfig),
		CooldownPeriod: dcl.ValueOrEmptyString(o.CooldownPeriod),
	}
	return p
}

// AutoscalingPolicyBasicAlgorithmYarnConfigToProto converts a AutoscalingPolicyBasicAlgorithmYarnConfig resource to its proto representation.
func DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o *alpha.AutoscalingPolicyBasicAlgorithmYarnConfig) *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.ValueOrEmptyString(o.GracefulDecommissionTimeout),
		ScaleUpFactor:               dcl.ValueOrEmptyDouble(o.ScaleUpFactor),
		ScaleDownFactor:             dcl.ValueOrEmptyDouble(o.ScaleDownFactor),
		ScaleUpMinWorkerFraction:    dcl.ValueOrEmptyDouble(o.ScaleUpMinWorkerFraction),
		ScaleDownMinWorkerFraction:  dcl.ValueOrEmptyDouble(o.ScaleDownMinWorkerFraction),
	}
	return p
}

// AutoscalingPolicyWorkerConfigToProto converts a AutoscalingPolicyWorkerConfig resource to its proto representation.
func DataprocAlphaAutoscalingPolicyWorkerConfigToProto(o *alpha.AutoscalingPolicyWorkerConfig) *alphapb.DataprocAlphaAutoscalingPolicyWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicyWorkerConfig{
		MinInstances: dcl.ValueOrEmptyInt64(o.MinInstances),
		MaxInstances: dcl.ValueOrEmptyInt64(o.MaxInstances),
		Weight:       dcl.ValueOrEmptyInt64(o.Weight),
	}
	return p
}

// AutoscalingPolicySecondaryWorkerConfigToProto converts a AutoscalingPolicySecondaryWorkerConfig resource to its proto representation.
func DataprocAlphaAutoscalingPolicySecondaryWorkerConfigToProto(o *alpha.AutoscalingPolicySecondaryWorkerConfig) *alphapb.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.ValueOrEmptyInt64(o.MinInstances),
		MaxInstances: dcl.ValueOrEmptyInt64(o.MaxInstances),
		Weight:       dcl.ValueOrEmptyInt64(o.Weight),
	}
	return p
}

// AutoscalingPolicyToProto converts a AutoscalingPolicy resource to its proto representation.
func AutoscalingPolicyToProto(resource *alpha.AutoscalingPolicy) *alphapb.DataprocAlphaAutoscalingPolicy {
	p := &alphapb.DataprocAlphaAutoscalingPolicy{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		BasicAlgorithm:        DataprocAlphaAutoscalingPolicyBasicAlgorithmToProto(resource.BasicAlgorithm),
		WorkerConfig:          DataprocAlphaAutoscalingPolicyWorkerConfigToProto(resource.WorkerConfig),
		SecondaryWorkerConfig: DataprocAlphaAutoscalingPolicySecondaryWorkerConfigToProto(resource.SecondaryWorkerConfig),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) applyAutoscalingPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaAutoscalingPolicyRequest) (*alphapb.DataprocAlphaAutoscalingPolicy, error) {
	p := ProtoToAutoscalingPolicy(request.GetResource())
	res, err := c.ApplyAutoscalingPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalingPolicyToProto(res)
	return r, nil
}

// ApplyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) ApplyDataprocAlphaAutoscalingPolicy(ctx context.Context, request *alphapb.ApplyDataprocAlphaAutoscalingPolicyRequest) (*alphapb.DataprocAlphaAutoscalingPolicy, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAutoscalingPolicy(ctx, cl, request)
}

// DeleteAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Delete() method.
func (s *AutoscalingPolicyServer) DeleteDataprocAlphaAutoscalingPolicy(ctx context.Context, request *alphapb.DeleteDataprocAlphaAutoscalingPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscalingPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscalingPolicy(ctx, ProtoToAutoscalingPolicy(request.GetResource()))

}

// ListDataprocAlphaAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicyList() method.
func (s *AutoscalingPolicyServer) ListDataprocAlphaAutoscalingPolicy(ctx context.Context, request *alphapb.ListDataprocAlphaAutoscalingPolicyRequest) (*alphapb.ListDataprocAlphaAutoscalingPolicyResponse, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscalingPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaAutoscalingPolicy
	for _, r := range resources.Items {
		rp := AutoscalingPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListDataprocAlphaAutoscalingPolicyResponse{Items: protos}, nil
}

func createConfigAutoscalingPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
