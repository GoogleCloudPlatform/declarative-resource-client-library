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

// Server implements the gRPC interface for UptimeCheckConfig.
type UptimeCheckConfigServer struct{}

// ProtoToUptimeCheckConfigResourceGroupResourceTypeEnum converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(e monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum) *monitoring.UptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigResourceGroupResourceTypeEnum(n[len("UptimeCheckConfigResourceGroupResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckRequestMethodEnum converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(e monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum) *monitoring.UptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigHttpCheckRequestMethodEnum(n[len("UptimeCheckConfigHttpCheckRequestMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckContentTypeEnum converts a UptimeCheckConfigHttpCheckContentTypeEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheckContentTypeEnum(e monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum) *monitoring.UptimeCheckConfigHttpCheckContentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigHttpCheckContentTypeEnum(n[len("UptimeCheckConfigHttpCheckContentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigContentMatchersMatcherEnum converts a UptimeCheckConfigContentMatchersMatcherEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigContentMatchersMatcherEnum(e monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum) *monitoring.UptimeCheckConfigContentMatchersMatcherEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigContentMatchersMatcherEnum(n[len("UptimeCheckConfigContentMatchersMatcherEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigMonitoredResource converts a UptimeCheckConfigMonitoredResource resource from its proto representation.
func ProtoToMonitoringUptimeCheckConfigMonitoredResource(p *monitoringpb.MonitoringUptimeCheckConfigMonitoredResource) *monitoring.UptimeCheckConfigMonitoredResource {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigMonitoredResource{
		Type: dcl.StringOrNil(p.Type),
	}
	return obj
}

// ProtoToUptimeCheckConfigResourceGroup converts a UptimeCheckConfigResourceGroup resource from its proto representation.
func ProtoToMonitoringUptimeCheckConfigResourceGroup(p *monitoringpb.MonitoringUptimeCheckConfigResourceGroup) *monitoring.UptimeCheckConfigResourceGroup {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigResourceGroup{
		GroupId:      dcl.StringOrNil(p.GroupId),
		ResourceType: ProtoToMonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheck converts a UptimeCheckConfigHttpCheck resource from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheck(p *monitoringpb.MonitoringUptimeCheckConfigHttpCheck) *monitoring.UptimeCheckConfigHttpCheck {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigHttpCheck{
		RequestMethod: ProtoToMonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(p.GetRequestMethod()),
		UseSsl:        dcl.Bool(p.UseSsl),
		Path:          dcl.StringOrNil(p.Path),
		Port:          dcl.Int64OrNil(p.Port),
		AuthInfo:      ProtoToMonitoringUptimeCheckConfigHttpCheckAuthInfo(p.GetAuthInfo()),
		MaskHeaders:   dcl.Bool(p.MaskHeaders),
		ContentType:   ProtoToMonitoringUptimeCheckConfigHttpCheckContentTypeEnum(p.GetContentType()),
		ValidateSsl:   dcl.Bool(p.ValidateSsl),
		Body:          dcl.StringOrNil(p.Body),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheckAuthInfo converts a UptimeCheckConfigHttpCheckAuthInfo resource from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheckAuthInfo(p *monitoringpb.MonitoringUptimeCheckConfigHttpCheckAuthInfo) *monitoring.UptimeCheckConfigHttpCheckAuthInfo {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.StringOrNil(p.Username),
		Password: dcl.StringOrNil(p.Password),
	}
	return obj
}

// ProtoToUptimeCheckConfigTcpCheck converts a UptimeCheckConfigTcpCheck resource from its proto representation.
func ProtoToMonitoringUptimeCheckConfigTcpCheck(p *monitoringpb.MonitoringUptimeCheckConfigTcpCheck) *monitoring.UptimeCheckConfigTcpCheck {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigTcpCheck{
		Port: dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToUptimeCheckConfigContentMatchers converts a UptimeCheckConfigContentMatchers resource from its proto representation.
func ProtoToMonitoringUptimeCheckConfigContentMatchers(p *monitoringpb.MonitoringUptimeCheckConfigContentMatchers) *monitoring.UptimeCheckConfigContentMatchers {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigContentMatchers{
		Content: dcl.StringOrNil(p.Content),
		Matcher: ProtoToMonitoringUptimeCheckConfigContentMatchersMatcherEnum(p.GetMatcher()),
	}
	return obj
}

// ProtoToUptimeCheckConfig converts a UptimeCheckConfig resource from its proto representation.
func ProtoToUptimeCheckConfig(p *monitoringpb.MonitoringUptimeCheckConfig) *monitoring.UptimeCheckConfig {
	obj := &monitoring.UptimeCheckConfig{
		Name:              dcl.StringOrNil(p.Name),
		DisplayName:       dcl.StringOrNil(p.DisplayName),
		MonitoredResource: ProtoToMonitoringUptimeCheckConfigMonitoredResource(p.GetMonitoredResource()),
		ResourceGroup:     ProtoToMonitoringUptimeCheckConfigResourceGroup(p.GetResourceGroup()),
		HttpCheck:         ProtoToMonitoringUptimeCheckConfigHttpCheck(p.GetHttpCheck()),
		TcpCheck:          ProtoToMonitoringUptimeCheckConfigTcpCheck(p.GetTcpCheck()),
		Period:            dcl.StringOrNil(p.Period),
		Timeout:           dcl.StringOrNil(p.Timeout),
		Project:           dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetContentMatchers() {
		obj.ContentMatchers = append(obj.ContentMatchers, *ProtoToMonitoringUptimeCheckConfigContentMatchers(r))
	}
	for _, r := range p.GetPrivateCheckers() {
		obj.PrivateCheckers = append(obj.PrivateCheckers, r)
	}
	for _, r := range p.GetSelectedRegions() {
		obj.SelectedRegions = append(obj.SelectedRegions, r)
	}
	return obj
}

// UptimeCheckConfigResourceGroupResourceTypeEnumToProto converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum to its proto representation.
func MonitoringUptimeCheckConfigResourceGroupResourceTypeEnumToProto(e *monitoring.UptimeCheckConfigResourceGroupResourceTypeEnum) monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum_value["UptimeCheckConfigResourceGroupResourceTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(0)
}

// UptimeCheckConfigHttpCheckRequestMethodEnumToProto converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckRequestMethodEnumToProto(e *monitoring.UptimeCheckConfigHttpCheckRequestMethodEnum) monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum_value["UptimeCheckConfigHttpCheckRequestMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(0)
}

// UptimeCheckConfigHttpCheckContentTypeEnumToProto converts a UptimeCheckConfigHttpCheckContentTypeEnum enum to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckContentTypeEnumToProto(e *monitoring.UptimeCheckConfigHttpCheckContentTypeEnum) monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum_value["UptimeCheckConfigHttpCheckContentTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum(0)
}

// UptimeCheckConfigContentMatchersMatcherEnumToProto converts a UptimeCheckConfigContentMatchersMatcherEnum enum to its proto representation.
func MonitoringUptimeCheckConfigContentMatchersMatcherEnumToProto(e *monitoring.UptimeCheckConfigContentMatchersMatcherEnum) monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum_value["UptimeCheckConfigContentMatchersMatcherEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum(0)
}

// UptimeCheckConfigMonitoredResourceToProto converts a UptimeCheckConfigMonitoredResource resource to its proto representation.
func MonitoringUptimeCheckConfigMonitoredResourceToProto(o *monitoring.UptimeCheckConfigMonitoredResource) *monitoringpb.MonitoringUptimeCheckConfigMonitoredResource {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigMonitoredResource{
		Type: dcl.ValueOrEmptyString(o.Type),
	}
	p.FilterLabels = make(map[string]string)
	for k, r := range o.FilterLabels {
		p.FilterLabels[k] = r
	}
	return p
}

// UptimeCheckConfigResourceGroupToProto converts a UptimeCheckConfigResourceGroup resource to its proto representation.
func MonitoringUptimeCheckConfigResourceGroupToProto(o *monitoring.UptimeCheckConfigResourceGroup) *monitoringpb.MonitoringUptimeCheckConfigResourceGroup {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigResourceGroup{
		GroupId:      dcl.ValueOrEmptyString(o.GroupId),
		ResourceType: MonitoringUptimeCheckConfigResourceGroupResourceTypeEnumToProto(o.ResourceType),
	}
	return p
}

// UptimeCheckConfigHttpCheckToProto converts a UptimeCheckConfigHttpCheck resource to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckToProto(o *monitoring.UptimeCheckConfigHttpCheck) *monitoringpb.MonitoringUptimeCheckConfigHttpCheck {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigHttpCheck{
		RequestMethod: MonitoringUptimeCheckConfigHttpCheckRequestMethodEnumToProto(o.RequestMethod),
		UseSsl:        dcl.ValueOrEmptyBool(o.UseSsl),
		Path:          dcl.ValueOrEmptyString(o.Path),
		Port:          dcl.ValueOrEmptyInt64(o.Port),
		AuthInfo:      MonitoringUptimeCheckConfigHttpCheckAuthInfoToProto(o.AuthInfo),
		MaskHeaders:   dcl.ValueOrEmptyBool(o.MaskHeaders),
		ContentType:   MonitoringUptimeCheckConfigHttpCheckContentTypeEnumToProto(o.ContentType),
		ValidateSsl:   dcl.ValueOrEmptyBool(o.ValidateSsl),
		Body:          dcl.ValueOrEmptyString(o.Body),
	}
	p.Headers = make(map[string]string)
	for k, r := range o.Headers {
		p.Headers[k] = r
	}
	return p
}

// UptimeCheckConfigHttpCheckAuthInfoToProto converts a UptimeCheckConfigHttpCheckAuthInfo resource to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckAuthInfoToProto(o *monitoring.UptimeCheckConfigHttpCheckAuthInfo) *monitoringpb.MonitoringUptimeCheckConfigHttpCheckAuthInfo {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.ValueOrEmptyString(o.Username),
		Password: dcl.ValueOrEmptyString(o.Password),
	}
	return p
}

// UptimeCheckConfigTcpCheckToProto converts a UptimeCheckConfigTcpCheck resource to its proto representation.
func MonitoringUptimeCheckConfigTcpCheckToProto(o *monitoring.UptimeCheckConfigTcpCheck) *monitoringpb.MonitoringUptimeCheckConfigTcpCheck {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigTcpCheck{
		Port: dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// UptimeCheckConfigContentMatchersToProto converts a UptimeCheckConfigContentMatchers resource to its proto representation.
func MonitoringUptimeCheckConfigContentMatchersToProto(o *monitoring.UptimeCheckConfigContentMatchers) *monitoringpb.MonitoringUptimeCheckConfigContentMatchers {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigContentMatchers{
		Content: dcl.ValueOrEmptyString(o.Content),
		Matcher: MonitoringUptimeCheckConfigContentMatchersMatcherEnumToProto(o.Matcher),
	}
	return p
}

// UptimeCheckConfigToProto converts a UptimeCheckConfig resource to its proto representation.
func UptimeCheckConfigToProto(resource *monitoring.UptimeCheckConfig) *monitoringpb.MonitoringUptimeCheckConfig {
	p := &monitoringpb.MonitoringUptimeCheckConfig{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		DisplayName:       dcl.ValueOrEmptyString(resource.DisplayName),
		MonitoredResource: MonitoringUptimeCheckConfigMonitoredResourceToProto(resource.MonitoredResource),
		ResourceGroup:     MonitoringUptimeCheckConfigResourceGroupToProto(resource.ResourceGroup),
		HttpCheck:         MonitoringUptimeCheckConfigHttpCheckToProto(resource.HttpCheck),
		TcpCheck:          MonitoringUptimeCheckConfigTcpCheckToProto(resource.TcpCheck),
		Period:            dcl.ValueOrEmptyString(resource.Period),
		Timeout:           dcl.ValueOrEmptyString(resource.Timeout),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.ContentMatchers {
		p.ContentMatchers = append(p.ContentMatchers, MonitoringUptimeCheckConfigContentMatchersToProto(&r))
	}
	for _, r := range resource.PrivateCheckers {
		p.PrivateCheckers = append(p.PrivateCheckers, r)
	}
	for _, r := range resource.SelectedRegions {
		p.SelectedRegions = append(p.SelectedRegions, r)
	}

	return p
}

// ApplyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) applyUptimeCheckConfig(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringUptimeCheckConfigRequest) (*monitoringpb.MonitoringUptimeCheckConfig, error) {
	p := ProtoToUptimeCheckConfig(request.GetResource())
	res, err := c.ApplyUptimeCheckConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UptimeCheckConfigToProto(res)
	return r, nil
}

// ApplyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) ApplyMonitoringUptimeCheckConfig(ctx context.Context, request *monitoringpb.ApplyMonitoringUptimeCheckConfigRequest) (*monitoringpb.MonitoringUptimeCheckConfig, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUptimeCheckConfig(ctx, cl, request)
}

// DeleteUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Delete() method.
func (s *UptimeCheckConfigServer) DeleteMonitoringUptimeCheckConfig(ctx context.Context, request *monitoringpb.DeleteMonitoringUptimeCheckConfigRequest) (*emptypb.Empty, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUptimeCheckConfig(ctx, ProtoToUptimeCheckConfig(request.GetResource()))
}

// ListUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfigList() method.
func (s *UptimeCheckConfigServer) ListMonitoringUptimeCheckConfig(ctx context.Context, request *monitoringpb.ListMonitoringUptimeCheckConfigRequest) (*monitoringpb.ListMonitoringUptimeCheckConfigResponse, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUptimeCheckConfig(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringUptimeCheckConfig
	for _, r := range resources.Items {
		rp := UptimeCheckConfigToProto(r)
		protos = append(protos, rp)
	}
	return &monitoringpb.ListMonitoringUptimeCheckConfigResponse{Items: protos}, nil
}

func createConfigUptimeCheckConfig(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
