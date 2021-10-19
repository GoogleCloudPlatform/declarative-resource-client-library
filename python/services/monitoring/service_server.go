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

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceCustom converts a ServiceCustom object from its proto representation.
func ProtoToMonitoringServiceCustom(p *monitoringpb.MonitoringServiceCustom) *monitoring.ServiceCustom {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceCustom{}
	return obj
}

// ProtoToServiceAppEngine converts a ServiceAppEngine object from its proto representation.
func ProtoToMonitoringServiceAppEngine(p *monitoringpb.MonitoringServiceAppEngine) *monitoring.ServiceAppEngine {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceAppEngine{
		ModuleId: dcl.StringOrNil(p.GetModuleId()),
	}
	return obj
}

// ProtoToServiceCloudEndpoints converts a ServiceCloudEndpoints object from its proto representation.
func ProtoToMonitoringServiceCloudEndpoints(p *monitoringpb.MonitoringServiceCloudEndpoints) *monitoring.ServiceCloudEndpoints {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceCloudEndpoints{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToServiceClusterIstio converts a ServiceClusterIstio object from its proto representation.
func ProtoToMonitoringServiceClusterIstio(p *monitoringpb.MonitoringServiceClusterIstio) *monitoring.ServiceClusterIstio {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceClusterIstio{
		Location:         dcl.StringOrNil(p.GetLocation()),
		ClusterName:      dcl.StringOrNil(p.GetClusterName()),
		ServiceNamespace: dcl.StringOrNil(p.GetServiceNamespace()),
		ServiceName:      dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceMeshIstio converts a ServiceMeshIstio object from its proto representation.
func ProtoToMonitoringServiceMeshIstio(p *monitoringpb.MonitoringServiceMeshIstio) *monitoring.ServiceMeshIstio {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceMeshIstio{
		MeshUid:          dcl.StringOrNil(p.GetMeshUid()),
		ServiceNamespace: dcl.StringOrNil(p.GetServiceNamespace()),
		ServiceName:      dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceIstioCanonicalService converts a ServiceIstioCanonicalService object from its proto representation.
func ProtoToMonitoringServiceIstioCanonicalService(p *monitoringpb.MonitoringServiceIstioCanonicalService) *monitoring.ServiceIstioCanonicalService {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceIstioCanonicalService{
		MeshUid:                   dcl.StringOrNil(p.GetMeshUid()),
		CanonicalServiceNamespace: dcl.StringOrNil(p.GetCanonicalServiceNamespace()),
		CanonicalService:          dcl.StringOrNil(p.GetCanonicalService()),
	}
	return obj
}

// ProtoToServiceCloudRun converts a ServiceCloudRun object from its proto representation.
func ProtoToMonitoringServiceCloudRun(p *monitoringpb.MonitoringServiceCloudRun) *monitoring.ServiceCloudRun {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceCloudRun{
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToServiceGkeNamespace converts a ServiceGkeNamespace object from its proto representation.
func ProtoToMonitoringServiceGkeNamespace(p *monitoringpb.MonitoringServiceGkeNamespace) *monitoring.ServiceGkeNamespace {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceGkeNamespace{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		ClusterName:   dcl.StringOrNil(p.GetClusterName()),
		NamespaceName: dcl.StringOrNil(p.GetNamespaceName()),
	}
	return obj
}

// ProtoToServiceGkeWorkload converts a ServiceGkeWorkload object from its proto representation.
func ProtoToMonitoringServiceGkeWorkload(p *monitoringpb.MonitoringServiceGkeWorkload) *monitoring.ServiceGkeWorkload {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceGkeWorkload{
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
func ProtoToMonitoringServiceGkeService(p *monitoringpb.MonitoringServiceGkeService) *monitoring.ServiceGkeService {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceGkeService{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		ClusterName:   dcl.StringOrNil(p.GetClusterName()),
		NamespaceName: dcl.StringOrNil(p.GetNamespaceName()),
		ServiceName:   dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceTelemetry converts a ServiceTelemetry object from its proto representation.
func ProtoToMonitoringServiceTelemetry(p *monitoringpb.MonitoringServiceTelemetry) *monitoring.ServiceTelemetry {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceTelemetry{
		ResourceName: dcl.StringOrNil(p.GetResourceName()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *monitoringpb.MonitoringService) *monitoring.Service {
	obj := &monitoring.Service{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Custom:                ProtoToMonitoringServiceCustom(p.GetCustom()),
		AppEngine:             ProtoToMonitoringServiceAppEngine(p.GetAppEngine()),
		CloudEndpoints:        ProtoToMonitoringServiceCloudEndpoints(p.GetCloudEndpoints()),
		ClusterIstio:          ProtoToMonitoringServiceClusterIstio(p.GetClusterIstio()),
		MeshIstio:             ProtoToMonitoringServiceMeshIstio(p.GetMeshIstio()),
		IstioCanonicalService: ProtoToMonitoringServiceIstioCanonicalService(p.GetIstioCanonicalService()),
		CloudRun:              ProtoToMonitoringServiceCloudRun(p.GetCloudRun()),
		GkeNamespace:          ProtoToMonitoringServiceGkeNamespace(p.GetGkeNamespace()),
		GkeWorkload:           ProtoToMonitoringServiceGkeWorkload(p.GetGkeWorkload()),
		GkeService:            ProtoToMonitoringServiceGkeService(p.GetGkeService()),
		Telemetry:             ProtoToMonitoringServiceTelemetry(p.GetTelemetry()),
		Deleted:               dcl.Bool(p.GetDeleted()),
		Project:               dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ServiceCustomToProto converts a ServiceCustom object to its proto representation.
func MonitoringServiceCustomToProto(o *monitoring.ServiceCustom) *monitoringpb.MonitoringServiceCustom {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceCustom{}
	return p
}

// ServiceAppEngineToProto converts a ServiceAppEngine object to its proto representation.
func MonitoringServiceAppEngineToProto(o *monitoring.ServiceAppEngine) *monitoringpb.MonitoringServiceAppEngine {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceAppEngine{}
	p.SetModuleId(dcl.ValueOrEmptyString(o.ModuleId))
	return p
}

// ServiceCloudEndpointsToProto converts a ServiceCloudEndpoints object to its proto representation.
func MonitoringServiceCloudEndpointsToProto(o *monitoring.ServiceCloudEndpoints) *monitoringpb.MonitoringServiceCloudEndpoints {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceCloudEndpoints{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ServiceClusterIstioToProto converts a ServiceClusterIstio object to its proto representation.
func MonitoringServiceClusterIstioToProto(o *monitoring.ServiceClusterIstio) *monitoringpb.MonitoringServiceClusterIstio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceClusterIstio{}
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetServiceNamespace(dcl.ValueOrEmptyString(o.ServiceNamespace))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceMeshIstioToProto converts a ServiceMeshIstio object to its proto representation.
func MonitoringServiceMeshIstioToProto(o *monitoring.ServiceMeshIstio) *monitoringpb.MonitoringServiceMeshIstio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceMeshIstio{}
	p.SetMeshUid(dcl.ValueOrEmptyString(o.MeshUid))
	p.SetServiceNamespace(dcl.ValueOrEmptyString(o.ServiceNamespace))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceIstioCanonicalServiceToProto converts a ServiceIstioCanonicalService object to its proto representation.
func MonitoringServiceIstioCanonicalServiceToProto(o *monitoring.ServiceIstioCanonicalService) *monitoringpb.MonitoringServiceIstioCanonicalService {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceIstioCanonicalService{}
	p.SetMeshUid(dcl.ValueOrEmptyString(o.MeshUid))
	p.SetCanonicalServiceNamespace(dcl.ValueOrEmptyString(o.CanonicalServiceNamespace))
	p.SetCanonicalService(dcl.ValueOrEmptyString(o.CanonicalService))
	return p
}

// ServiceCloudRunToProto converts a ServiceCloudRun object to its proto representation.
func MonitoringServiceCloudRunToProto(o *monitoring.ServiceCloudRun) *monitoringpb.MonitoringServiceCloudRun {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceCloudRun{}
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// ServiceGkeNamespaceToProto converts a ServiceGkeNamespace object to its proto representation.
func MonitoringServiceGkeNamespaceToProto(o *monitoring.ServiceGkeNamespace) *monitoringpb.MonitoringServiceGkeNamespace {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceGkeNamespace{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	return p
}

// ServiceGkeWorkloadToProto converts a ServiceGkeWorkload object to its proto representation.
func MonitoringServiceGkeWorkloadToProto(o *monitoring.ServiceGkeWorkload) *monitoringpb.MonitoringServiceGkeWorkload {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceGkeWorkload{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	p.SetTopLevelControllerType(dcl.ValueOrEmptyString(o.TopLevelControllerType))
	p.SetTopLevelControllerName(dcl.ValueOrEmptyString(o.TopLevelControllerName))
	return p
}

// ServiceGkeServiceToProto converts a ServiceGkeService object to its proto representation.
func MonitoringServiceGkeServiceToProto(o *monitoring.ServiceGkeService) *monitoringpb.MonitoringServiceGkeService {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceGkeService{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceTelemetryToProto converts a ServiceTelemetry object to its proto representation.
func MonitoringServiceTelemetryToProto(o *monitoring.ServiceTelemetry) *monitoringpb.MonitoringServiceTelemetry {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceTelemetry{}
	p.SetResourceName(dcl.ValueOrEmptyString(o.ResourceName))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *monitoring.Service) *monitoringpb.MonitoringService {
	p := &monitoringpb.MonitoringService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCustom(MonitoringServiceCustomToProto(resource.Custom))
	p.SetAppEngine(MonitoringServiceAppEngineToProto(resource.AppEngine))
	p.SetCloudEndpoints(MonitoringServiceCloudEndpointsToProto(resource.CloudEndpoints))
	p.SetClusterIstio(MonitoringServiceClusterIstioToProto(resource.ClusterIstio))
	p.SetMeshIstio(MonitoringServiceMeshIstioToProto(resource.MeshIstio))
	p.SetIstioCanonicalService(MonitoringServiceIstioCanonicalServiceToProto(resource.IstioCanonicalService))
	p.SetCloudRun(MonitoringServiceCloudRunToProto(resource.CloudRun))
	p.SetGkeNamespace(MonitoringServiceGkeNamespaceToProto(resource.GkeNamespace))
	p.SetGkeWorkload(MonitoringServiceGkeWorkloadToProto(resource.GkeWorkload))
	p.SetGkeService(MonitoringServiceGkeServiceToProto(resource.GkeService))
	p.SetTelemetry(MonitoringServiceTelemetryToProto(resource.Telemetry))
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
func (s *ServiceServer) applyService(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringServiceRequest) (*monitoringpb.MonitoringService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyMonitoringService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyMonitoringService(ctx context.Context, request *monitoringpb.ApplyMonitoringServiceRequest) (*monitoringpb.MonitoringService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteMonitoringService(ctx context.Context, request *monitoringpb.DeleteMonitoringServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListMonitoringService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListMonitoringService(ctx context.Context, request *monitoringpb.ListMonitoringServiceRequest) (*monitoringpb.ListMonitoringServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &monitoringpb.ListMonitoringServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
