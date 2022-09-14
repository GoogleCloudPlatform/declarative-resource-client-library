// Copyright 2022 Google LLC. All Rights Reserved.
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
	clouddeploypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/clouddeploy_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy"
)

// DeliveryPipelineServer implements the gRPC interface for DeliveryPipeline.
type DeliveryPipelineServer struct{}

// ProtoToDeliveryPipelineSerialPipeline converts a DeliveryPipelineSerialPipeline object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipeline(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipeline) *clouddeploy.DeliveryPipelineSerialPipeline {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipeline{}
	for _, r := range p.GetStages() {
		obj.Stages = append(obj.Stages, *ProtoToClouddeployDeliveryPipelineSerialPipelineStages(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStages converts a DeliveryPipelineSerialPipelineStages object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStages(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages) *clouddeploy.DeliveryPipelineSerialPipelineStages {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStages{
		TargetId: dcl.StringOrNil(p.GetTargetId()),
		Strategy: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategy(p.GetStrategy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategy converts a DeliveryPipelineSerialPipelineStagesStrategy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategy{
		Standard: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard(p.GetStandard()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandard converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandard{
		Verify: dcl.Bool(p.GetVerify()),
	}
	return obj
}

// ProtoToDeliveryPipelineCondition converts a DeliveryPipelineCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineCondition(p *clouddeploypb.ClouddeployDeliveryPipelineCondition) *clouddeploy.DeliveryPipelineCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineCondition{
		PipelineReadyCondition:  ProtoToClouddeployDeliveryPipelineConditionPipelineReadyCondition(p.GetPipelineReadyCondition()),
		TargetsPresentCondition: ProtoToClouddeployDeliveryPipelineConditionTargetsPresentCondition(p.GetTargetsPresentCondition()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionPipelineReadyCondition converts a DeliveryPipelineConditionPipelineReadyCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineConditionPipelineReadyCondition(p *clouddeploypb.ClouddeployDeliveryPipelineConditionPipelineReadyCondition) *clouddeploy.DeliveryPipelineConditionPipelineReadyCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineConditionPipelineReadyCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsPresentCondition converts a DeliveryPipelineConditionTargetsPresentCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineConditionTargetsPresentCondition(p *clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsPresentCondition) *clouddeploy.DeliveryPipelineConditionTargetsPresentCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineConditionTargetsPresentCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMissingTargets() {
		obj.MissingTargets = append(obj.MissingTargets, r)
	}
	return obj
}

// ProtoToDeliveryPipeline converts a DeliveryPipeline resource from its proto representation.
func ProtoToDeliveryPipeline(p *clouddeploypb.ClouddeployDeliveryPipeline) *clouddeploy.DeliveryPipeline {
	obj := &clouddeploy.DeliveryPipeline{
		Name:           dcl.StringOrNil(p.GetName()),
		Uid:            dcl.StringOrNil(p.GetUid()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		SerialPipeline: ProtoToClouddeployDeliveryPipelineSerialPipeline(p.GetSerialPipeline()),
		Condition:      ProtoToClouddeployDeliveryPipelineCondition(p.GetCondition()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
		Suspended:      dcl.Bool(p.GetSuspended()),
	}
	return obj
}

// DeliveryPipelineSerialPipelineToProto converts a DeliveryPipelineSerialPipeline object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineToProto(o *clouddeploy.DeliveryPipelineSerialPipeline) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipeline {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipeline{}
	sStages := make([]*clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages, len(o.Stages))
	for i, r := range o.Stages {
		sStages[i] = ClouddeployDeliveryPipelineSerialPipelineStagesToProto(&r)
	}
	p.SetStages(sStages)
	return p
}

// DeliveryPipelineSerialPipelineStagesToProto converts a DeliveryPipelineSerialPipelineStages object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStages) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages{}
	p.SetTargetId(dcl.ValueOrEmptyString(o.TargetId))
	p.SetStrategy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyToProto(o.Strategy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyToProto converts a DeliveryPipelineSerialPipelineStagesStrategy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy{}
	p.SetStandard(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o.Standard))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandard) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	return p
}

// DeliveryPipelineConditionToProto converts a DeliveryPipelineCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionToProto(o *clouddeploy.DeliveryPipelineCondition) *clouddeploypb.ClouddeployDeliveryPipelineCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineCondition{}
	p.SetPipelineReadyCondition(ClouddeployDeliveryPipelineConditionPipelineReadyConditionToProto(o.PipelineReadyCondition))
	p.SetTargetsPresentCondition(ClouddeployDeliveryPipelineConditionTargetsPresentConditionToProto(o.TargetsPresentCondition))
	return p
}

// DeliveryPipelineConditionPipelineReadyConditionToProto converts a DeliveryPipelineConditionPipelineReadyCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionPipelineReadyConditionToProto(o *clouddeploy.DeliveryPipelineConditionPipelineReadyCondition) *clouddeploypb.ClouddeployDeliveryPipelineConditionPipelineReadyCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineConditionPipelineReadyCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// DeliveryPipelineConditionTargetsPresentConditionToProto converts a DeliveryPipelineConditionTargetsPresentCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionTargetsPresentConditionToProto(o *clouddeploy.DeliveryPipelineConditionTargetsPresentCondition) *clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsPresentCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsPresentCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	sMissingTargets := make([]string, len(o.MissingTargets))
	for i, r := range o.MissingTargets {
		sMissingTargets[i] = r
	}
	p.SetMissingTargets(sMissingTargets)
	return p
}

// DeliveryPipelineToProto converts a DeliveryPipeline resource to its proto representation.
func DeliveryPipelineToProto(resource *clouddeploy.DeliveryPipeline) *clouddeploypb.ClouddeployDeliveryPipeline {
	p := &clouddeploypb.ClouddeployDeliveryPipeline{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSerialPipeline(ClouddeployDeliveryPipelineSerialPipelineToProto(resource.SerialPipeline))
	p.SetCondition(ClouddeployDeliveryPipelineConditionToProto(resource.Condition))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSuspended(dcl.ValueOrEmptyBool(resource.Suspended))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) applyDeliveryPipeline(ctx context.Context, c *clouddeploy.Client, request *clouddeploypb.ApplyClouddeployDeliveryPipelineRequest) (*clouddeploypb.ClouddeployDeliveryPipeline, error) {
	p := ProtoToDeliveryPipeline(request.GetResource())
	res, err := c.ApplyDeliveryPipeline(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DeliveryPipelineToProto(res)
	return r, nil
}

// applyClouddeployDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) ApplyClouddeployDeliveryPipeline(ctx context.Context, request *clouddeploypb.ApplyClouddeployDeliveryPipelineRequest) (*clouddeploypb.ClouddeployDeliveryPipeline, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDeliveryPipeline(ctx, cl, request)
}

// DeleteDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Delete() method.
func (s *DeliveryPipelineServer) DeleteClouddeployDeliveryPipeline(ctx context.Context, request *clouddeploypb.DeleteClouddeployDeliveryPipelineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDeliveryPipeline(ctx, ProtoToDeliveryPipeline(request.GetResource()))

}

// ListClouddeployDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipelineList() method.
func (s *DeliveryPipelineServer) ListClouddeployDeliveryPipeline(ctx context.Context, request *clouddeploypb.ListClouddeployDeliveryPipelineRequest) (*clouddeploypb.ListClouddeployDeliveryPipelineResponse, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDeliveryPipeline(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*clouddeploypb.ClouddeployDeliveryPipeline
	for _, r := range resources.Items {
		rp := DeliveryPipelineToProto(r)
		protos = append(protos, rp)
	}
	p := &clouddeploypb.ListClouddeployDeliveryPipelineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDeliveryPipeline(ctx context.Context, service_account_file string) (*clouddeploy.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return clouddeploy.NewClient(conf), nil
}
