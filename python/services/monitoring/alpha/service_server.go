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

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceCustom converts a ServiceCustom object from its proto representation.
func ProtoToMonitoringAlphaServiceCustom(p *alphapb.MonitoringAlphaServiceCustom) *alpha.ServiceCustom {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceCustom{}
	return obj
}

// ProtoToServiceAppEngine converts a ServiceAppEngine object from its proto representation.
func ProtoToMonitoringAlphaServiceAppEngine(p *alphapb.MonitoringAlphaServiceAppEngine) *alpha.ServiceAppEngine {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAppEngine{
		ModuleId: dcl.StringOrNil(p.GetModuleId()),
	}
	return obj
}

// ProtoToServiceCloudEndpoints converts a ServiceCloudEndpoints object from its proto representation.
func ProtoToMonitoringAlphaServiceCloudEndpoints(p *alphapb.MonitoringAlphaServiceCloudEndpoints) *alpha.ServiceCloudEndpoints {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceCloudEndpoints{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToServiceClusterIstio converts a ServiceClusterIstio object from its proto representation.
func ProtoToMonitoringAlphaServiceClusterIstio(p *alphapb.MonitoringAlphaServiceClusterIstio) *alpha.ServiceClusterIstio {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceClusterIstio{
		Location:         dcl.StringOrNil(p.GetLocation()),
		ClusterName:      dcl.StringOrNil(p.GetClusterName()),
		ServiceNamespace: dcl.StringOrNil(p.GetServiceNamespace()),
		ServiceName:      dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceMeshIstio converts a ServiceMeshIstio object from its proto representation.
func ProtoToMonitoringAlphaServiceMeshIstio(p *alphapb.MonitoringAlphaServiceMeshIstio) *alpha.ServiceMeshIstio {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceMeshIstio{
		MeshUid:          dcl.StringOrNil(p.GetMeshUid()),
		ServiceNamespace: dcl.StringOrNil(p.GetServiceNamespace()),
		ServiceName:      dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceIstioCanonicalService converts a ServiceIstioCanonicalService object from its proto representation.
func ProtoToMonitoringAlphaServiceIstioCanonicalService(p *alphapb.MonitoringAlphaServiceIstioCanonicalService) *alpha.ServiceIstioCanonicalService {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceIstioCanonicalService{
		MeshUid:                   dcl.StringOrNil(p.GetMeshUid()),
		CanonicalServiceNamespace: dcl.StringOrNil(p.GetCanonicalServiceNamespace()),
		CanonicalService:          dcl.StringOrNil(p.GetCanonicalService()),
	}
	return obj
}

// ProtoToServiceCloudRun converts a ServiceCloudRun object from its proto representation.
func ProtoToMonitoringAlphaServiceCloudRun(p *alphapb.MonitoringAlphaServiceCloudRun) *alpha.ServiceCloudRun {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceCloudRun{
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToServiceGkeNamespace converts a ServiceGkeNamespace object from its proto representation.
func ProtoToMonitoringAlphaServiceGkeNamespace(p *alphapb.MonitoringAlphaServiceGkeNamespace) *alpha.ServiceGkeNamespace {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceGkeNamespace{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		ClusterName:   dcl.StringOrNil(p.GetClusterName()),
		NamespaceName: dcl.StringOrNil(p.GetNamespaceName()),
	}
	return obj
}

// ProtoToServiceGkeWorkload converts a ServiceGkeWorkload object from its proto representation.
func ProtoToMonitoringAlphaServiceGkeWorkload(p *alphapb.MonitoringAlphaServiceGkeWorkload) *alpha.ServiceGkeWorkload {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceGkeWorkload{
		ProjectId:              dcl.StringOrNil(p.GetProjectId()),
		Location:               dcl.StringOrNil(p.GetLocation()),
		ClusterName:            dcl.StringOrNil(p.GetClusterName()),
		NamespaceName:          dcl.StringOrNil(p.GetNamespaceName()),
		TopLevelControllerType: dcl.StringOrNil(p.GetTopLevelControllerType()),
		TopLevelControllerName: dcl.StringOrNil(p.GetTopLevelControllerName()),
	}
	return obj
}

// ProtoToServiceGkeService converts a ServiceGkeService object from its proto representation.
func ProtoToMonitoringAlphaServiceGkeService(p *alphapb.MonitoringAlphaServiceGkeService) *alpha.ServiceGkeService {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceGkeService{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		ClusterName:   dcl.StringOrNil(p.GetClusterName()),
		NamespaceName: dcl.StringOrNil(p.GetNamespaceName()),
		ServiceName:   dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceTelemetry converts a ServiceTelemetry object from its proto representation.
func ProtoToMonitoringAlphaServiceTelemetry(p *alphapb.MonitoringAlphaServiceTelemetry) *alpha.ServiceTelemetry {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTelemetry{
		ResourceName: dcl.StringOrNil(p.GetResourceName()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *alphapb.MonitoringAlphaService) *alpha.Service {
	obj := &alpha.Service{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Custom:                ProtoToMonitoringAlphaServiceCustom(p.GetCustom()),
		AppEngine:             ProtoToMonitoringAlphaServiceAppEngine(p.GetAppEngine()),
		CloudEndpoints:        ProtoToMonitoringAlphaServiceCloudEndpoints(p.GetCloudEndpoints()),
		ClusterIstio:          ProtoToMonitoringAlphaServiceClusterIstio(p.GetClusterIstio()),
		MeshIstio:             ProtoToMonitoringAlphaServiceMeshIstio(p.GetMeshIstio()),
		IstioCanonicalService: ProtoToMonitoringAlphaServiceIstioCanonicalService(p.GetIstioCanonicalService()),
		CloudRun:              ProtoToMonitoringAlphaServiceCloudRun(p.GetCloudRun()),
		GkeNamespace:          ProtoToMonitoringAlphaServiceGkeNamespace(p.GetGkeNamespace()),
		GkeWorkload:           ProtoToMonitoringAlphaServiceGkeWorkload(p.GetGkeWorkload()),
		GkeService:            ProtoToMonitoringAlphaServiceGkeService(p.GetGkeService()),
		Telemetry:             ProtoToMonitoringAlphaServiceTelemetry(p.GetTelemetry()),
		Deleted:               dcl.Bool(p.GetDeleted()),
		Project:               dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ServiceCustomToProto converts a ServiceCustom object to its proto representation.
func MonitoringAlphaServiceCustomToProto(o *alpha.ServiceCustom) *alphapb.MonitoringAlphaServiceCustom {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceCustom{}
	return p
}

// ServiceAppEngineToProto converts a ServiceAppEngine object to its proto representation.
func MonitoringAlphaServiceAppEngineToProto(o *alpha.ServiceAppEngine) *alphapb.MonitoringAlphaServiceAppEngine {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceAppEngine{}
	p.SetModuleId(dcl.ValueOrEmptyString(o.ModuleId))
	return p
}

// ServiceCloudEndpointsToProto converts a ServiceCloudEndpoints object to its proto representation.
func MonitoringAlphaServiceCloudEndpointsToProto(o *alpha.ServiceCloudEndpoints) *alphapb.MonitoringAlphaServiceCloudEndpoints {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceCloudEndpoints{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ServiceClusterIstioToProto converts a ServiceClusterIstio object to its proto representation.
func MonitoringAlphaServiceClusterIstioToProto(o *alpha.ServiceClusterIstio) *alphapb.MonitoringAlphaServiceClusterIstio {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceClusterIstio{}
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetServiceNamespace(dcl.ValueOrEmptyString(o.ServiceNamespace))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceMeshIstioToProto converts a ServiceMeshIstio object to its proto representation.
func MonitoringAlphaServiceMeshIstioToProto(o *alpha.ServiceMeshIstio) *alphapb.MonitoringAlphaServiceMeshIstio {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceMeshIstio{}
	p.SetMeshUid(dcl.ValueOrEmptyString(o.MeshUid))
	p.SetServiceNamespace(dcl.ValueOrEmptyString(o.ServiceNamespace))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceIstioCanonicalServiceToProto converts a ServiceIstioCanonicalService object to its proto representation.
func MonitoringAlphaServiceIstioCanonicalServiceToProto(o *alpha.ServiceIstioCanonicalService) *alphapb.MonitoringAlphaServiceIstioCanonicalService {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceIstioCanonicalService{}
	p.SetMeshUid(dcl.ValueOrEmptyString(o.MeshUid))
	p.SetCanonicalServiceNamespace(dcl.ValueOrEmptyString(o.CanonicalServiceNamespace))
	p.SetCanonicalService(dcl.ValueOrEmptyString(o.CanonicalService))
	return p
}

// ServiceCloudRunToProto converts a ServiceCloudRun object to its proto representation.
func MonitoringAlphaServiceCloudRunToProto(o *alpha.ServiceCloudRun) *alphapb.MonitoringAlphaServiceCloudRun {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceCloudRun{}
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// ServiceGkeNamespaceToProto converts a ServiceGkeNamespace object to its proto representation.
func MonitoringAlphaServiceGkeNamespaceToProto(o *alpha.ServiceGkeNamespace) *alphapb.MonitoringAlphaServiceGkeNamespace {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceGkeNamespace{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	return p
}

// ServiceGkeWorkloadToProto converts a ServiceGkeWorkload object to its proto representation.
func MonitoringAlphaServiceGkeWorkloadToProto(o *alpha.ServiceGkeWorkload) *alphapb.MonitoringAlphaServiceGkeWorkload {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceGkeWorkload{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	p.SetTopLevelControllerType(dcl.ValueOrEmptyString(o.TopLevelControllerType))
	p.SetTopLevelControllerName(dcl.ValueOrEmptyString(o.TopLevelControllerName))
	return p
}

// ServiceGkeServiceToProto converts a ServiceGkeService object to its proto representation.
func MonitoringAlphaServiceGkeServiceToProto(o *alpha.ServiceGkeService) *alphapb.MonitoringAlphaServiceGkeService {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceGkeService{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceTelemetryToProto converts a ServiceTelemetry object to its proto representation.
func MonitoringAlphaServiceTelemetryToProto(o *alpha.ServiceTelemetry) *alphapb.MonitoringAlphaServiceTelemetry {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceTelemetry{}
	p.SetResourceName(dcl.ValueOrEmptyString(o.ResourceName))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *alpha.Service) *alphapb.MonitoringAlphaService {
	p := &alphapb.MonitoringAlphaService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCustom(MonitoringAlphaServiceCustomToProto(resource.Custom))
	p.SetAppEngine(MonitoringAlphaServiceAppEngineToProto(resource.AppEngine))
	p.SetCloudEndpoints(MonitoringAlphaServiceCloudEndpointsToProto(resource.CloudEndpoints))
	p.SetClusterIstio(MonitoringAlphaServiceClusterIstioToProto(resource.ClusterIstio))
	p.SetMeshIstio(MonitoringAlphaServiceMeshIstioToProto(resource.MeshIstio))
	p.SetIstioCanonicalService(MonitoringAlphaServiceIstioCanonicalServiceToProto(resource.IstioCanonicalService))
	p.SetCloudRun(MonitoringAlphaServiceCloudRunToProto(resource.CloudRun))
	p.SetGkeNamespace(MonitoringAlphaServiceGkeNamespaceToProto(resource.GkeNamespace))
	p.SetGkeWorkload(MonitoringAlphaServiceGkeWorkloadToProto(resource.GkeWorkload))
	p.SetGkeService(MonitoringAlphaServiceGkeServiceToProto(resource.GkeService))
	p.SetTelemetry(MonitoringAlphaServiceTelemetryToProto(resource.Telemetry))
	p.SetDeleted(dcl.ValueOrEmptyBool(resource.Deleted))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mUserLabels := make(map[string]string, len(resource.UserLabels))
	for k, r := range resource.UserLabels {
		mUserLabels[k] = r
	}
	p.SetUserLabels(mUserLabels)

	return p
}

// applyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaServiceRequest) (*alphapb.MonitoringAlphaService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyMonitoringAlphaService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyMonitoringAlphaService(ctx context.Context, request *alphapb.ApplyMonitoringAlphaServiceRequest) (*alphapb.MonitoringAlphaService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteMonitoringAlphaService(ctx context.Context, request *alphapb.DeleteMonitoringAlphaServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListMonitoringAlphaService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListMonitoringAlphaService(ctx context.Context, request *alphapb.ListMonitoringAlphaServiceRequest) (*alphapb.ListMonitoringAlphaServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
