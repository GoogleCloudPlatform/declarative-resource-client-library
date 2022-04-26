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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/alpha/clouddeploy_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/alpha"
)

// DeliveryPipelineServer implements the gRPC interface for DeliveryPipeline.
type DeliveryPipelineServer struct{}

// ProtoToDeliveryPipelineSerialPipeline converts a DeliveryPipelineSerialPipeline object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipeline(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipeline) *alpha.DeliveryPipelineSerialPipeline {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipeline{}
	for _, r := range p.GetStages() {
		obj.Stages = append(obj.Stages, *ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStages(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStages converts a DeliveryPipelineSerialPipelineStages object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStages(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages) *alpha.DeliveryPipelineSerialPipelineStages {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStages{
		TargetId: dcl.StringOrNil(p.GetTargetId()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	return obj
}

// ProtoToDeliveryPipelineCondition converts a DeliveryPipelineCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineCondition) *alpha.DeliveryPipelineCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineCondition{
		PipelineReadyCondition:  ProtoToClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition(p.GetPipelineReadyCondition()),
		TargetsPresentCondition: ProtoToClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition(p.GetTargetsPresentCondition()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionPipelineReadyCondition converts a DeliveryPipelineConditionPipelineReadyCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition) *alpha.DeliveryPipelineConditionPipelineReadyCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineConditionPipelineReadyCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsPresentCondition converts a DeliveryPipelineConditionTargetsPresentCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition) *alpha.DeliveryPipelineConditionTargetsPresentCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineConditionTargetsPresentCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMissingTargets() {
		obj.MissingTargets = append(obj.MissingTargets, r)
	}
	return obj
}

// ProtoToDeliveryPipeline converts a DeliveryPipeline resource from its proto representation.
func ProtoToDeliveryPipeline(p *alphapb.ClouddeployAlphaDeliveryPipeline) *alpha.DeliveryPipeline {
	obj := &alpha.DeliveryPipeline{
		Name:           dcl.StringOrNil(p.GetName()),
		Uid:            dcl.StringOrNil(p.GetUid()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		SerialPipeline: ProtoToClouddeployAlphaDeliveryPipelineSerialPipeline(p.GetSerialPipeline()),
		Condition:      ProtoToClouddeployAlphaDeliveryPipelineCondition(p.GetCondition()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// DeliveryPipelineSerialPipelineToProto converts a DeliveryPipelineSerialPipeline object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineToProto(o *alpha.DeliveryPipelineSerialPipeline) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipeline {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipeline{}
	sStages := make([]*alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages, len(o.Stages))
	for i, r := range o.Stages {
		sStages[i] = ClouddeployAlphaDeliveryPipelineSerialPipelineStagesToProto(&r)
	}
	p.SetStages(sStages)
	return p
}

// DeliveryPipelineSerialPipelineStagesToProto converts a DeliveryPipelineSerialPipelineStages object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesToProto(o *alpha.DeliveryPipelineSerialPipelineStages) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages{}
	p.SetTargetId(dcl.ValueOrEmptyString(o.TargetId))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	return p
}

// DeliveryPipelineConditionToProto converts a DeliveryPipelineCondition object to its proto representation.
func ClouddeployAlphaDeliveryPipelineConditionToProto(o *alpha.DeliveryPipelineCondition) *alphapb.ClouddeployAlphaDeliveryPipelineCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineCondition{}
	p.SetPipelineReadyCondition(ClouddeployAlphaDeliveryPipelineConditionPipelineReadyConditionToProto(o.PipelineReadyCondition))
	p.SetTargetsPresentCondition(ClouddeployAlphaDeliveryPipelineConditionTargetsPresentConditionToProto(o.TargetsPresentCondition))
	return p
}

// DeliveryPipelineConditionPipelineReadyConditionToProto converts a DeliveryPipelineConditionPipelineReadyCondition object to its proto representation.
func ClouddeployAlphaDeliveryPipelineConditionPipelineReadyConditionToProto(o *alpha.DeliveryPipelineConditionPipelineReadyCondition) *alphapb.ClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// DeliveryPipelineConditionTargetsPresentConditionToProto converts a DeliveryPipelineConditionTargetsPresentCondition object to its proto representation.
func ClouddeployAlphaDeliveryPipelineConditionTargetsPresentConditionToProto(o *alpha.DeliveryPipelineConditionTargetsPresentCondition) *alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition{}
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
func DeliveryPipelineToProto(resource *alpha.DeliveryPipeline) *alphapb.ClouddeployAlphaDeliveryPipeline {
	p := &alphapb.ClouddeployAlphaDeliveryPipeline{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSerialPipeline(ClouddeployAlphaDeliveryPipelineSerialPipelineToProto(resource.SerialPipeline))
	p.SetCondition(ClouddeployAlphaDeliveryPipelineConditionToProto(resource.Condition))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
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
func (s *DeliveryPipelineServer) applyDeliveryPipeline(ctx context.Context, c *alpha.Client, request *alphapb.ApplyClouddeployAlphaDeliveryPipelineRequest) (*alphapb.ClouddeployAlphaDeliveryPipeline, error) {
	p := ProtoToDeliveryPipeline(request.GetResource())
	res, err := c.ApplyDeliveryPipeline(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DeliveryPipelineToProto(res)
	return r, nil
}

// applyClouddeployAlphaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) ApplyClouddeployAlphaDeliveryPipeline(ctx context.Context, request *alphapb.ApplyClouddeployAlphaDeliveryPipelineRequest) (*alphapb.ClouddeployAlphaDeliveryPipeline, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDeliveryPipeline(ctx, cl, request)
}

// DeleteDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Delete() method.
func (s *DeliveryPipelineServer) DeleteClouddeployAlphaDeliveryPipeline(ctx context.Context, request *alphapb.DeleteClouddeployAlphaDeliveryPipelineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDeliveryPipeline(ctx, ProtoToDeliveryPipeline(request.GetResource()))

}

// ListClouddeployAlphaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipelineList() method.
func (s *DeliveryPipelineServer) ListClouddeployAlphaDeliveryPipeline(ctx context.Context, request *alphapb.ListClouddeployAlphaDeliveryPipelineRequest) (*alphapb.ListClouddeployAlphaDeliveryPipelineResponse, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDeliveryPipeline(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ClouddeployAlphaDeliveryPipeline
	for _, r := range resources.Items {
		rp := DeliveryPipelineToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListClouddeployAlphaDeliveryPipelineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDeliveryPipeline(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
