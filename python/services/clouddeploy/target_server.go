// Copyright 2025 Google LLC. All Rights Reserved.
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

// TargetServer implements the gRPC interface for Target.
type TargetServer struct{}

// ProtoToTargetExecutionConfigsUsagesEnum converts a TargetExecutionConfigsUsagesEnum enum from its proto representation.
func ProtoToClouddeployTargetExecutionConfigsUsagesEnum(e clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum) *clouddeploy.TargetExecutionConfigsUsagesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum_name[int32(e)]; ok {
		e := clouddeploy.TargetExecutionConfigsUsagesEnum(n[len("ClouddeployTargetExecutionConfigsUsagesEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetGke converts a TargetGke object from its proto representation.
func ProtoToClouddeployTargetGke(p *clouddeploypb.ClouddeployTargetGke) *clouddeploy.TargetGke {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetGke{
		Cluster:     dcl.StringOrNil(p.GetCluster()),
		InternalIP:  dcl.Bool(p.GetInternalIp()),
		ProxyUrl:    dcl.StringOrNil(p.GetProxyUrl()),
		DnsEndpoint: dcl.Bool(p.GetDnsEndpoint()),
	}
	return obj
}

// ProtoToTargetAnthosCluster converts a TargetAnthosCluster object from its proto representation.
func ProtoToClouddeployTargetAnthosCluster(p *clouddeploypb.ClouddeployTargetAnthosCluster) *clouddeploy.TargetAnthosCluster {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetAnthosCluster{
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToTargetExecutionConfigs converts a TargetExecutionConfigs object from its proto representation.
func ProtoToClouddeployTargetExecutionConfigs(p *clouddeploypb.ClouddeployTargetExecutionConfigs) *clouddeploy.TargetExecutionConfigs {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetExecutionConfigs{
		WorkerPool:       dcl.StringOrNil(p.GetWorkerPool()),
		ServiceAccount:   dcl.StringOrNil(p.GetServiceAccount()),
		ArtifactStorage:  dcl.StringOrNil(p.GetArtifactStorage()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
		Verbose:          dcl.Bool(p.GetVerbose()),
	}
	for _, r := range p.GetUsages() {
		obj.Usages = append(obj.Usages, *ProtoToClouddeployTargetExecutionConfigsUsagesEnum(r))
	}
	return obj
}

// ProtoToTargetRun converts a TargetRun object from its proto representation.
func ProtoToClouddeployTargetRun(p *clouddeploypb.ClouddeployTargetRun) *clouddeploy.TargetRun {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetRun{
		Location: dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToTargetMultiTarget converts a TargetMultiTarget object from its proto representation.
func ProtoToClouddeployTargetMultiTarget(p *clouddeploypb.ClouddeployTargetMultiTarget) *clouddeploy.TargetMultiTarget {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetMultiTarget{}
	for _, r := range p.GetTargetIds() {
		obj.TargetIds = append(obj.TargetIds, r)
	}
	return obj
}

// ProtoToTargetCustomTarget converts a TargetCustomTarget object from its proto representation.
func ProtoToClouddeployTargetCustomTarget(p *clouddeploypb.ClouddeployTargetCustomTarget) *clouddeploy.TargetCustomTarget {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetCustomTarget{
		CustomTargetType: dcl.StringOrNil(p.GetCustomTargetType()),
	}
	return obj
}

// ProtoToTargetAssociatedEntities converts a TargetAssociatedEntities object from its proto representation.
func ProtoToClouddeployTargetAssociatedEntities(p *clouddeploypb.ClouddeployTargetAssociatedEntities) *clouddeploy.TargetAssociatedEntities {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetAssociatedEntities{}
	for _, r := range p.GetGkeClusters() {
		obj.GkeClusters = append(obj.GkeClusters, *ProtoToClouddeployTargetAssociatedEntitiesGkeClusters(r))
	}
	for _, r := range p.GetAnthosClusters() {
		obj.AnthosClusters = append(obj.AnthosClusters, *ProtoToClouddeployTargetAssociatedEntitiesAnthosClusters(r))
	}
	return obj
}

// ProtoToTargetAssociatedEntitiesGkeClusters converts a TargetAssociatedEntitiesGkeClusters object from its proto representation.
func ProtoToClouddeployTargetAssociatedEntitiesGkeClusters(p *clouddeploypb.ClouddeployTargetAssociatedEntitiesGkeClusters) *clouddeploy.TargetAssociatedEntitiesGkeClusters {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetAssociatedEntitiesGkeClusters{
		Cluster:    dcl.StringOrNil(p.GetCluster()),
		InternalIP: dcl.Bool(p.GetInternalIp()),
		ProxyUrl:   dcl.StringOrNil(p.GetProxyUrl()),
	}
	return obj
}

// ProtoToTargetAssociatedEntitiesAnthosClusters converts a TargetAssociatedEntitiesAnthosClusters object from its proto representation.
func ProtoToClouddeployTargetAssociatedEntitiesAnthosClusters(p *clouddeploypb.ClouddeployTargetAssociatedEntitiesAnthosClusters) *clouddeploy.TargetAssociatedEntitiesAnthosClusters {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.TargetAssociatedEntitiesAnthosClusters{
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToTarget converts a Target resource from its proto representation.
func ProtoToTarget(p *clouddeploypb.ClouddeployTarget) *clouddeploy.Target {
	obj := &clouddeploy.Target{
		Name:            dcl.StringOrNil(p.GetName()),
		TargetId:        dcl.StringOrNil(p.GetTargetId()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		RequireApproval: dcl.Bool(p.GetRequireApproval()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Gke:             ProtoToClouddeployTargetGke(p.GetGke()),
		AnthosCluster:   ProtoToClouddeployTargetAnthosCluster(p.GetAnthosCluster()),
		Etag:            dcl.StringOrNil(p.GetEtag()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		Run:             ProtoToClouddeployTargetRun(p.GetRun()),
		MultiTarget:     ProtoToClouddeployTargetMultiTarget(p.GetMultiTarget()),
		CustomTarget:    ProtoToClouddeployTargetCustomTarget(p.GetCustomTarget()),
	}
	for _, r := range p.GetExecutionConfigs() {
		obj.ExecutionConfigs = append(obj.ExecutionConfigs, *ProtoToClouddeployTargetExecutionConfigs(r))
	}
	return obj
}

// TargetExecutionConfigsUsagesEnumToProto converts a TargetExecutionConfigsUsagesEnum enum to its proto representation.
func ClouddeployTargetExecutionConfigsUsagesEnumToProto(e *clouddeploy.TargetExecutionConfigsUsagesEnum) clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum {
	if e == nil {
		return clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum(0)
	}
	if v, ok := clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum_value["TargetExecutionConfigsUsagesEnum"+string(*e)]; ok {
		return clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum(v)
	}
	return clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum(0)
}

// TargetGkeToProto converts a TargetGke object to its proto representation.
func ClouddeployTargetGkeToProto(o *clouddeploy.TargetGke) *clouddeploypb.ClouddeployTargetGke {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetGke{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetInternalIp(dcl.ValueOrEmptyBool(o.InternalIP))
	p.SetProxyUrl(dcl.ValueOrEmptyString(o.ProxyUrl))
	p.SetDnsEndpoint(dcl.ValueOrEmptyBool(o.DnsEndpoint))
	return p
}

// TargetAnthosClusterToProto converts a TargetAnthosCluster object to its proto representation.
func ClouddeployTargetAnthosClusterToProto(o *clouddeploy.TargetAnthosCluster) *clouddeploypb.ClouddeployTargetAnthosCluster {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetAnthosCluster{}
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// TargetExecutionConfigsToProto converts a TargetExecutionConfigs object to its proto representation.
func ClouddeployTargetExecutionConfigsToProto(o *clouddeploy.TargetExecutionConfigs) *clouddeploypb.ClouddeployTargetExecutionConfigs {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetExecutionConfigs{}
	p.SetWorkerPool(dcl.ValueOrEmptyString(o.WorkerPool))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetArtifactStorage(dcl.ValueOrEmptyString(o.ArtifactStorage))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	p.SetVerbose(dcl.ValueOrEmptyBool(o.Verbose))
	sUsages := make([]clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum, len(o.Usages))
	for i, r := range o.Usages {
		sUsages[i] = clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum(clouddeploypb.ClouddeployTargetExecutionConfigsUsagesEnum_value[string(r)])
	}
	p.SetUsages(sUsages)
	return p
}

// TargetRunToProto converts a TargetRun object to its proto representation.
func ClouddeployTargetRunToProto(o *clouddeploy.TargetRun) *clouddeploypb.ClouddeployTargetRun {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetRun{}
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// TargetMultiTargetToProto converts a TargetMultiTarget object to its proto representation.
func ClouddeployTargetMultiTargetToProto(o *clouddeploy.TargetMultiTarget) *clouddeploypb.ClouddeployTargetMultiTarget {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetMultiTarget{}
	sTargetIds := make([]string, len(o.TargetIds))
	for i, r := range o.TargetIds {
		sTargetIds[i] = r
	}
	p.SetTargetIds(sTargetIds)
	return p
}

// TargetCustomTargetToProto converts a TargetCustomTarget object to its proto representation.
func ClouddeployTargetCustomTargetToProto(o *clouddeploy.TargetCustomTarget) *clouddeploypb.ClouddeployTargetCustomTarget {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetCustomTarget{}
	p.SetCustomTargetType(dcl.ValueOrEmptyString(o.CustomTargetType))
	return p
}

// TargetAssociatedEntitiesToProto converts a TargetAssociatedEntities object to its proto representation.
func ClouddeployTargetAssociatedEntitiesToProto(o *clouddeploy.TargetAssociatedEntities) *clouddeploypb.ClouddeployTargetAssociatedEntities {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetAssociatedEntities{}
	sGkeClusters := make([]*clouddeploypb.ClouddeployTargetAssociatedEntitiesGkeClusters, len(o.GkeClusters))
	for i, r := range o.GkeClusters {
		sGkeClusters[i] = ClouddeployTargetAssociatedEntitiesGkeClustersToProto(&r)
	}
	p.SetGkeClusters(sGkeClusters)
	sAnthosClusters := make([]*clouddeploypb.ClouddeployTargetAssociatedEntitiesAnthosClusters, len(o.AnthosClusters))
	for i, r := range o.AnthosClusters {
		sAnthosClusters[i] = ClouddeployTargetAssociatedEntitiesAnthosClustersToProto(&r)
	}
	p.SetAnthosClusters(sAnthosClusters)
	return p
}

// TargetAssociatedEntitiesGkeClustersToProto converts a TargetAssociatedEntitiesGkeClusters object to its proto representation.
func ClouddeployTargetAssociatedEntitiesGkeClustersToProto(o *clouddeploy.TargetAssociatedEntitiesGkeClusters) *clouddeploypb.ClouddeployTargetAssociatedEntitiesGkeClusters {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetAssociatedEntitiesGkeClusters{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetInternalIp(dcl.ValueOrEmptyBool(o.InternalIP))
	p.SetProxyUrl(dcl.ValueOrEmptyString(o.ProxyUrl))
	return p
}

// TargetAssociatedEntitiesAnthosClustersToProto converts a TargetAssociatedEntitiesAnthosClusters object to its proto representation.
func ClouddeployTargetAssociatedEntitiesAnthosClustersToProto(o *clouddeploy.TargetAssociatedEntitiesAnthosClusters) *clouddeploypb.ClouddeployTargetAssociatedEntitiesAnthosClusters {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployTargetAssociatedEntitiesAnthosClusters{}
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// TargetToProto converts a Target resource to its proto representation.
func TargetToProto(resource *clouddeploy.Target) *clouddeploypb.ClouddeployTarget {
	p := &clouddeploypb.ClouddeployTarget{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTargetId(dcl.ValueOrEmptyString(resource.TargetId))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetRequireApproval(dcl.ValueOrEmptyBool(resource.RequireApproval))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGke(ClouddeployTargetGkeToProto(resource.Gke))
	p.SetAnthosCluster(ClouddeployTargetAnthosClusterToProto(resource.AnthosCluster))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetRun(ClouddeployTargetRunToProto(resource.Run))
	p.SetMultiTarget(ClouddeployTargetMultiTargetToProto(resource.MultiTarget))
	p.SetCustomTarget(ClouddeployTargetCustomTargetToProto(resource.CustomTarget))
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
	sExecutionConfigs := make([]*clouddeploypb.ClouddeployTargetExecutionConfigs, len(resource.ExecutionConfigs))
	for i, r := range resource.ExecutionConfigs {
		sExecutionConfigs[i] = ClouddeployTargetExecutionConfigsToProto(&r)
	}
	p.SetExecutionConfigs(sExecutionConfigs)
	mDeployParameters := make(map[string]string, len(resource.DeployParameters))
	for k, r := range resource.DeployParameters {
		mDeployParameters[k] = r
	}
	p.SetDeployParameters(mDeployParameters)
	mAssociatedEntities := make(map[string]*clouddeploypb.ClouddeployTargetAssociatedEntities, len(resource.AssociatedEntities))
	for k, r := range resource.AssociatedEntities {
		mAssociatedEntities[k] = ClouddeployTargetAssociatedEntitiesToProto(&r)
	}
	p.SetAssociatedEntities(mAssociatedEntities)

	return p
}

// applyTarget handles the gRPC request by passing it to the underlying Target Apply() method.
func (s *TargetServer) applyTarget(ctx context.Context, c *clouddeploy.Client, request *clouddeploypb.ApplyClouddeployTargetRequest) (*clouddeploypb.ClouddeployTarget, error) {
	p := ProtoToTarget(request.GetResource())
	res, err := c.ApplyTarget(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetToProto(res)
	return r, nil
}

// applyClouddeployTarget handles the gRPC request by passing it to the underlying Target Apply() method.
func (s *TargetServer) ApplyClouddeployTarget(ctx context.Context, request *clouddeploypb.ApplyClouddeployTargetRequest) (*clouddeploypb.ClouddeployTarget, error) {
	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTarget(ctx, cl, request)
}

// DeleteTarget handles the gRPC request by passing it to the underlying Target Delete() method.
func (s *TargetServer) DeleteClouddeployTarget(ctx context.Context, request *clouddeploypb.DeleteClouddeployTargetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTarget(ctx, ProtoToTarget(request.GetResource()))

}

// ListClouddeployTarget handles the gRPC request by passing it to the underlying TargetList() method.
func (s *TargetServer) ListClouddeployTarget(ctx context.Context, request *clouddeploypb.ListClouddeployTargetRequest) (*clouddeploypb.ListClouddeployTargetResponse, error) {
	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTarget(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*clouddeploypb.ClouddeployTarget
	for _, r := range resources.Items {
		rp := TargetToProto(r)
		protos = append(protos, rp)
	}
	p := &clouddeploypb.ListClouddeployTargetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTarget(ctx context.Context, service_account_file string) (*clouddeploy.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return clouddeploy.NewClient(conf), nil
}
