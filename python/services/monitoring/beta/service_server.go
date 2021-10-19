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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/beta/monitoring_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
)

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceCustom converts a ServiceCustom object from its proto representation.
func ProtoToMonitoringBetaServiceCustom(p *betapb.MonitoringBetaServiceCustom) *beta.ServiceCustom {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceCustom{}
	return obj
}

// ProtoToServiceAppEngine converts a ServiceAppEngine object from its proto representation.
func ProtoToMonitoringBetaServiceAppEngine(p *betapb.MonitoringBetaServiceAppEngine) *beta.ServiceAppEngine {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceAppEngine{
		ModuleId: dcl.StringOrNil(p.GetModuleId()),
	}
	return obj
}

// ProtoToServiceCloudEndpoints converts a ServiceCloudEndpoints object from its proto representation.
func ProtoToMonitoringBetaServiceCloudEndpoints(p *betapb.MonitoringBetaServiceCloudEndpoints) *beta.ServiceCloudEndpoints {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceCloudEndpoints{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToServiceClusterIstio converts a ServiceClusterIstio object from its proto representation.
func ProtoToMonitoringBetaServiceClusterIstio(p *betapb.MonitoringBetaServiceClusterIstio) *beta.ServiceClusterIstio {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceClusterIstio{
		Location:         dcl.StringOrNil(p.GetLocation()),
		ClusterName:      dcl.StringOrNil(p.GetClusterName()),
		ServiceNamespace: dcl.StringOrNil(p.GetServiceNamespace()),
		ServiceName:      dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceMeshIstio converts a ServiceMeshIstio object from its proto representation.
func ProtoToMonitoringBetaServiceMeshIstio(p *betapb.MonitoringBetaServiceMeshIstio) *beta.ServiceMeshIstio {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceMeshIstio{
		MeshUid:          dcl.StringOrNil(p.GetMeshUid()),
		ServiceNamespace: dcl.StringOrNil(p.GetServiceNamespace()),
		ServiceName:      dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceIstioCanonicalService converts a ServiceIstioCanonicalService object from its proto representation.
func ProtoToMonitoringBetaServiceIstioCanonicalService(p *betapb.MonitoringBetaServiceIstioCanonicalService) *beta.ServiceIstioCanonicalService {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceIstioCanonicalService{
		MeshUid:                   dcl.StringOrNil(p.GetMeshUid()),
		CanonicalServiceNamespace: dcl.StringOrNil(p.GetCanonicalServiceNamespace()),
		CanonicalService:          dcl.StringOrNil(p.GetCanonicalService()),
	}
	return obj
}

// ProtoToServiceCloudRun converts a ServiceCloudRun object from its proto representation.
func ProtoToMonitoringBetaServiceCloudRun(p *betapb.MonitoringBetaServiceCloudRun) *beta.ServiceCloudRun {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceCloudRun{
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToServiceGkeNamespace converts a ServiceGkeNamespace object from its proto representation.
func ProtoToMonitoringBetaServiceGkeNamespace(p *betapb.MonitoringBetaServiceGkeNamespace) *beta.ServiceGkeNamespace {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceGkeNamespace{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		ClusterName:   dcl.StringOrNil(p.GetClusterName()),
		NamespaceName: dcl.StringOrNil(p.GetNamespaceName()),
	}
	return obj
}

// ProtoToServiceGkeWorkload converts a ServiceGkeWorkload object from its proto representation.
func ProtoToMonitoringBetaServiceGkeWorkload(p *betapb.MonitoringBetaServiceGkeWorkload) *beta.ServiceGkeWorkload {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceGkeWorkload{
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
func ProtoToMonitoringBetaServiceGkeService(p *betapb.MonitoringBetaServiceGkeService) *beta.ServiceGkeService {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceGkeService{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		ClusterName:   dcl.StringOrNil(p.GetClusterName()),
		NamespaceName: dcl.StringOrNil(p.GetNamespaceName()),
		ServiceName:   dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToServiceTelemetry converts a ServiceTelemetry object from its proto representation.
func ProtoToMonitoringBetaServiceTelemetry(p *betapb.MonitoringBetaServiceTelemetry) *beta.ServiceTelemetry {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTelemetry{
		ResourceName: dcl.StringOrNil(p.GetResourceName()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *betapb.MonitoringBetaService) *beta.Service {
	obj := &beta.Service{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Custom:                ProtoToMonitoringBetaServiceCustom(p.GetCustom()),
		AppEngine:             ProtoToMonitoringBetaServiceAppEngine(p.GetAppEngine()),
		CloudEndpoints:        ProtoToMonitoringBetaServiceCloudEndpoints(p.GetCloudEndpoints()),
		ClusterIstio:          ProtoToMonitoringBetaServiceClusterIstio(p.GetClusterIstio()),
		MeshIstio:             ProtoToMonitoringBetaServiceMeshIstio(p.GetMeshIstio()),
		IstioCanonicalService: ProtoToMonitoringBetaServiceIstioCanonicalService(p.GetIstioCanonicalService()),
		CloudRun:              ProtoToMonitoringBetaServiceCloudRun(p.GetCloudRun()),
		GkeNamespace:          ProtoToMonitoringBetaServiceGkeNamespace(p.GetGkeNamespace()),
		GkeWorkload:           ProtoToMonitoringBetaServiceGkeWorkload(p.GetGkeWorkload()),
		GkeService:            ProtoToMonitoringBetaServiceGkeService(p.GetGkeService()),
		Telemetry:             ProtoToMonitoringBetaServiceTelemetry(p.GetTelemetry()),
		Deleted:               dcl.Bool(p.GetDeleted()),
		Project:               dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ServiceCustomToProto converts a ServiceCustom object to its proto representation.
func MonitoringBetaServiceCustomToProto(o *beta.ServiceCustom) *betapb.MonitoringBetaServiceCustom {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceCustom{}
	return p
}

// ServiceAppEngineToProto converts a ServiceAppEngine object to its proto representation.
func MonitoringBetaServiceAppEngineToProto(o *beta.ServiceAppEngine) *betapb.MonitoringBetaServiceAppEngine {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceAppEngine{}
	p.SetModuleId(dcl.ValueOrEmptyString(o.ModuleId))
	return p
}

// ServiceCloudEndpointsToProto converts a ServiceCloudEndpoints object to its proto representation.
func MonitoringBetaServiceCloudEndpointsToProto(o *beta.ServiceCloudEndpoints) *betapb.MonitoringBetaServiceCloudEndpoints {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceCloudEndpoints{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ServiceClusterIstioToProto converts a ServiceClusterIstio object to its proto representation.
func MonitoringBetaServiceClusterIstioToProto(o *beta.ServiceClusterIstio) *betapb.MonitoringBetaServiceClusterIstio {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceClusterIstio{}
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetServiceNamespace(dcl.ValueOrEmptyString(o.ServiceNamespace))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceMeshIstioToProto converts a ServiceMeshIstio object to its proto representation.
func MonitoringBetaServiceMeshIstioToProto(o *beta.ServiceMeshIstio) *betapb.MonitoringBetaServiceMeshIstio {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceMeshIstio{}
	p.SetMeshUid(dcl.ValueOrEmptyString(o.MeshUid))
	p.SetServiceNamespace(dcl.ValueOrEmptyString(o.ServiceNamespace))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceIstioCanonicalServiceToProto converts a ServiceIstioCanonicalService object to its proto representation.
func MonitoringBetaServiceIstioCanonicalServiceToProto(o *beta.ServiceIstioCanonicalService) *betapb.MonitoringBetaServiceIstioCanonicalService {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceIstioCanonicalService{}
	p.SetMeshUid(dcl.ValueOrEmptyString(o.MeshUid))
	p.SetCanonicalServiceNamespace(dcl.ValueOrEmptyString(o.CanonicalServiceNamespace))
	p.SetCanonicalService(dcl.ValueOrEmptyString(o.CanonicalService))
	return p
}

// ServiceCloudRunToProto converts a ServiceCloudRun object to its proto representation.
func MonitoringBetaServiceCloudRunToProto(o *beta.ServiceCloudRun) *betapb.MonitoringBetaServiceCloudRun {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceCloudRun{}
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// ServiceGkeNamespaceToProto converts a ServiceGkeNamespace object to its proto representation.
func MonitoringBetaServiceGkeNamespaceToProto(o *beta.ServiceGkeNamespace) *betapb.MonitoringBetaServiceGkeNamespace {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceGkeNamespace{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	return p
}

// ServiceGkeWorkloadToProto converts a ServiceGkeWorkload object to its proto representation.
func MonitoringBetaServiceGkeWorkloadToProto(o *beta.ServiceGkeWorkload) *betapb.MonitoringBetaServiceGkeWorkload {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceGkeWorkload{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	p.SetTopLevelControllerType(dcl.ValueOrEmptyString(o.TopLevelControllerType))
	p.SetTopLevelControllerName(dcl.ValueOrEmptyString(o.TopLevelControllerName))
	return p
}

// ServiceGkeServiceToProto converts a ServiceGkeService object to its proto representation.
func MonitoringBetaServiceGkeServiceToProto(o *beta.ServiceGkeService) *betapb.MonitoringBetaServiceGkeService {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceGkeService{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetNamespaceName(dcl.ValueOrEmptyString(o.NamespaceName))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// ServiceTelemetryToProto converts a ServiceTelemetry object to its proto representation.
func MonitoringBetaServiceTelemetryToProto(o *beta.ServiceTelemetry) *betapb.MonitoringBetaServiceTelemetry {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceTelemetry{}
	p.SetResourceName(dcl.ValueOrEmptyString(o.ResourceName))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *beta.Service) *betapb.MonitoringBetaService {
	p := &betapb.MonitoringBetaService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCustom(MonitoringBetaServiceCustomToProto(resource.Custom))
	p.SetAppEngine(MonitoringBetaServiceAppEngineToProto(resource.AppEngine))
	p.SetCloudEndpoints(MonitoringBetaServiceCloudEndpointsToProto(resource.CloudEndpoints))
	p.SetClusterIstio(MonitoringBetaServiceClusterIstioToProto(resource.ClusterIstio))
	p.SetMeshIstio(MonitoringBetaServiceMeshIstioToProto(resource.MeshIstio))
	p.SetIstioCanonicalService(MonitoringBetaServiceIstioCanonicalServiceToProto(resource.IstioCanonicalService))
	p.SetCloudRun(MonitoringBetaServiceCloudRunToProto(resource.CloudRun))
	p.SetGkeNamespace(MonitoringBetaServiceGkeNamespaceToProto(resource.GkeNamespace))
	p.SetGkeWorkload(MonitoringBetaServiceGkeWorkloadToProto(resource.GkeWorkload))
	p.SetGkeService(MonitoringBetaServiceGkeServiceToProto(resource.GkeService))
	p.SetTelemetry(MonitoringBetaServiceTelemetryToProto(resource.Telemetry))
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
func (s *ServiceServer) applyService(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaServiceRequest) (*betapb.MonitoringBetaService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyMonitoringBetaService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyMonitoringBetaService(ctx context.Context, request *betapb.ApplyMonitoringBetaServiceRequest) (*betapb.MonitoringBetaService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteMonitoringBetaService(ctx context.Context, request *betapb.DeleteMonitoringBetaServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListMonitoringBetaService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListMonitoringBetaService(ctx context.Context, request *betapb.ListMonitoringBetaServiceRequest) (*betapb.ListMonitoringBetaServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListMonitoringBetaServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
