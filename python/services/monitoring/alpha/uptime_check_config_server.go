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

// Server implements the gRPC interface for UptimeCheckConfig.
type UptimeCheckConfigServer struct{}

// ProtoToUptimeCheckConfigResourceGroupResourceTypeEnum converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(e alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum) *alpha.UptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigResourceGroupResourceTypeEnum(n[len("MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckRequestMethodEnum converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(e alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum) *alpha.UptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigHttpCheckRequestMethodEnum(n[len("MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckContentTypeEnum converts a UptimeCheckConfigHttpCheckContentTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(e alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum) *alpha.UptimeCheckConfigHttpCheckContentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigHttpCheckContentTypeEnum(n[len("MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigContentMatchersMatcherEnum converts a UptimeCheckConfigContentMatchersMatcherEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(e alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum) *alpha.UptimeCheckConfigContentMatchersMatcherEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigContentMatchersMatcherEnum(n[len("MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigMonitoredResource converts a UptimeCheckConfigMonitoredResource resource from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigMonitoredResource(p *alphapb.MonitoringAlphaUptimeCheckConfigMonitoredResource) *alpha.UptimeCheckConfigMonitoredResource {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigMonitoredResource{
		Type: dcl.StringOrNil(p.Type),
	}
	return obj
}

// ProtoToUptimeCheckConfigResourceGroup converts a UptimeCheckConfigResourceGroup resource from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigResourceGroup(p *alphapb.MonitoringAlphaUptimeCheckConfigResourceGroup) *alpha.UptimeCheckConfigResourceGroup {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigResourceGroup{
		GroupId:      dcl.StringOrNil(p.GroupId),
		ResourceType: ProtoToMonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheck converts a UptimeCheckConfigHttpCheck resource from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheck(p *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheck) *alpha.UptimeCheckConfigHttpCheck {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigHttpCheck{
		RequestMethod: ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(p.GetRequestMethod()),
		UseSsl:        dcl.Bool(p.UseSsl),
		Path:          dcl.StringOrNil(p.Path),
		Port:          dcl.Int64OrNil(p.Port),
		AuthInfo:      ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo(p.GetAuthInfo()),
		MaskHeaders:   dcl.Bool(p.MaskHeaders),
		ContentType:   ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(p.GetContentType()),
		ValidateSsl:   dcl.Bool(p.ValidateSsl),
		Body:          dcl.StringOrNil(p.Body),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheckAuthInfo converts a UptimeCheckConfigHttpCheckAuthInfo resource from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo(p *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo) *alpha.UptimeCheckConfigHttpCheckAuthInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.StringOrNil(p.Username),
		Password: dcl.StringOrNil(p.Password),
	}
	return obj
}

// ProtoToUptimeCheckConfigTcpCheck converts a UptimeCheckConfigTcpCheck resource from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigTcpCheck(p *alphapb.MonitoringAlphaUptimeCheckConfigTcpCheck) *alpha.UptimeCheckConfigTcpCheck {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigTcpCheck{
		Port: dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToUptimeCheckConfigContentMatchers converts a UptimeCheckConfigContentMatchers resource from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigContentMatchers(p *alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers) *alpha.UptimeCheckConfigContentMatchers {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigContentMatchers{
		Content: dcl.StringOrNil(p.Content),
		Matcher: ProtoToMonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(p.GetMatcher()),
	}
	return obj
}

// ProtoToUptimeCheckConfig converts a UptimeCheckConfig resource from its proto representation.
func ProtoToUptimeCheckConfig(p *alphapb.MonitoringAlphaUptimeCheckConfig) *alpha.UptimeCheckConfig {
	obj := &alpha.UptimeCheckConfig{
		Name:              dcl.StringOrNil(p.Name),
		DisplayName:       dcl.StringOrNil(p.DisplayName),
		MonitoredResource: ProtoToMonitoringAlphaUptimeCheckConfigMonitoredResource(p.GetMonitoredResource()),
		ResourceGroup:     ProtoToMonitoringAlphaUptimeCheckConfigResourceGroup(p.GetResourceGroup()),
		HttpCheck:         ProtoToMonitoringAlphaUptimeCheckConfigHttpCheck(p.GetHttpCheck()),
		TcpCheck:          ProtoToMonitoringAlphaUptimeCheckConfigTcpCheck(p.GetTcpCheck()),
		Period:            dcl.StringOrNil(p.Period),
		Timeout:           dcl.StringOrNil(p.Timeout),
		Project:           dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetContentMatchers() {
		obj.ContentMatchers = append(obj.ContentMatchers, *ProtoToMonitoringAlphaUptimeCheckConfigContentMatchers(r))
	}
	for _, r := range p.GetSelectedRegions() {
		obj.SelectedRegions = append(obj.SelectedRegions, r)
	}
	return obj
}

// UptimeCheckConfigResourceGroupResourceTypeEnumToProto converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnumToProto(e *alpha.UptimeCheckConfigResourceGroupResourceTypeEnum) alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum_value["UptimeCheckConfigResourceGroupResourceTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(0)
}

// UptimeCheckConfigHttpCheckRequestMethodEnumToProto converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnumToProto(e *alpha.UptimeCheckConfigHttpCheckRequestMethodEnum) alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum_value["UptimeCheckConfigHttpCheckRequestMethodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(0)
}

// UptimeCheckConfigHttpCheckContentTypeEnumToProto converts a UptimeCheckConfigHttpCheckContentTypeEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnumToProto(e *alpha.UptimeCheckConfigHttpCheckContentTypeEnum) alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum_value["UptimeCheckConfigHttpCheckContentTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(0)
}

// UptimeCheckConfigContentMatchersMatcherEnumToProto converts a UptimeCheckConfigContentMatchersMatcherEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnumToProto(e *alpha.UptimeCheckConfigContentMatchersMatcherEnum) alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum_value["UptimeCheckConfigContentMatchersMatcherEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(0)
}

// UptimeCheckConfigMonitoredResourceToProto converts a UptimeCheckConfigMonitoredResource resource to its proto representation.
func MonitoringAlphaUptimeCheckConfigMonitoredResourceToProto(o *alpha.UptimeCheckConfigMonitoredResource) *alphapb.MonitoringAlphaUptimeCheckConfigMonitoredResource {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigMonitoredResource{
		Type: dcl.ValueOrEmptyString(o.Type),
	}
	p.FilterLabels = make(map[string]string)
	for k, r := range o.FilterLabels {
		p.FilterLabels[k] = r
	}
	return p
}

// UptimeCheckConfigResourceGroupToProto converts a UptimeCheckConfigResourceGroup resource to its proto representation.
func MonitoringAlphaUptimeCheckConfigResourceGroupToProto(o *alpha.UptimeCheckConfigResourceGroup) *alphapb.MonitoringAlphaUptimeCheckConfigResourceGroup {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigResourceGroup{
		GroupId:      dcl.ValueOrEmptyString(o.GroupId),
		ResourceType: MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnumToProto(o.ResourceType),
	}
	return p
}

// UptimeCheckConfigHttpCheckToProto converts a UptimeCheckConfigHttpCheck resource to its proto representation.
func MonitoringAlphaUptimeCheckConfigHttpCheckToProto(o *alpha.UptimeCheckConfigHttpCheck) *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheck {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigHttpCheck{
		RequestMethod: MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnumToProto(o.RequestMethod),
		UseSsl:        dcl.ValueOrEmptyBool(o.UseSsl),
		Path:          dcl.ValueOrEmptyString(o.Path),
		Port:          dcl.ValueOrEmptyInt64(o.Port),
		AuthInfo:      MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfoToProto(o.AuthInfo),
		MaskHeaders:   dcl.ValueOrEmptyBool(o.MaskHeaders),
		ContentType:   MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnumToProto(o.ContentType),
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
func MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfoToProto(o *alpha.UptimeCheckConfigHttpCheckAuthInfo) *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.ValueOrEmptyString(o.Username),
		Password: dcl.ValueOrEmptyString(o.Password),
	}
	return p
}

// UptimeCheckConfigTcpCheckToProto converts a UptimeCheckConfigTcpCheck resource to its proto representation.
func MonitoringAlphaUptimeCheckConfigTcpCheckToProto(o *alpha.UptimeCheckConfigTcpCheck) *alphapb.MonitoringAlphaUptimeCheckConfigTcpCheck {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigTcpCheck{
		Port: dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// UptimeCheckConfigContentMatchersToProto converts a UptimeCheckConfigContentMatchers resource to its proto representation.
func MonitoringAlphaUptimeCheckConfigContentMatchersToProto(o *alpha.UptimeCheckConfigContentMatchers) *alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers{
		Content: dcl.ValueOrEmptyString(o.Content),
		Matcher: MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnumToProto(o.Matcher),
	}
	return p
}

// UptimeCheckConfigToProto converts a UptimeCheckConfig resource to its proto representation.
func UptimeCheckConfigToProto(resource *alpha.UptimeCheckConfig) *alphapb.MonitoringAlphaUptimeCheckConfig {
	p := &alphapb.MonitoringAlphaUptimeCheckConfig{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		DisplayName:       dcl.ValueOrEmptyString(resource.DisplayName),
		MonitoredResource: MonitoringAlphaUptimeCheckConfigMonitoredResourceToProto(resource.MonitoredResource),
		ResourceGroup:     MonitoringAlphaUptimeCheckConfigResourceGroupToProto(resource.ResourceGroup),
		HttpCheck:         MonitoringAlphaUptimeCheckConfigHttpCheckToProto(resource.HttpCheck),
		TcpCheck:          MonitoringAlphaUptimeCheckConfigTcpCheckToProto(resource.TcpCheck),
		Period:            dcl.ValueOrEmptyString(resource.Period),
		Timeout:           dcl.ValueOrEmptyString(resource.Timeout),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.ContentMatchers {
		p.ContentMatchers = append(p.ContentMatchers, MonitoringAlphaUptimeCheckConfigContentMatchersToProto(&r))
	}
	for _, r := range resource.SelectedRegions {
		p.SelectedRegions = append(p.SelectedRegions, r)
	}

	return p
}

// ApplyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) applyUptimeCheckConfig(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaUptimeCheckConfigRequest) (*alphapb.MonitoringAlphaUptimeCheckConfig, error) {
	p := ProtoToUptimeCheckConfig(request.GetResource())
	res, err := c.ApplyUptimeCheckConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UptimeCheckConfigToProto(res)
	return r, nil
}

// ApplyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) ApplyMonitoringAlphaUptimeCheckConfig(ctx context.Context, request *alphapb.ApplyMonitoringAlphaUptimeCheckConfigRequest) (*alphapb.MonitoringAlphaUptimeCheckConfig, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUptimeCheckConfig(ctx, cl, request)
}

// DeleteUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Delete() method.
func (s *UptimeCheckConfigServer) DeleteMonitoringAlphaUptimeCheckConfig(ctx context.Context, request *alphapb.DeleteMonitoringAlphaUptimeCheckConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUptimeCheckConfig(ctx, ProtoToUptimeCheckConfig(request.GetResource()))

}

// ListMonitoringAlphaUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfigList() method.
func (s *UptimeCheckConfigServer) ListMonitoringAlphaUptimeCheckConfig(ctx context.Context, request *alphapb.ListMonitoringAlphaUptimeCheckConfigRequest) (*alphapb.ListMonitoringAlphaUptimeCheckConfigResponse, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUptimeCheckConfig(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaUptimeCheckConfig
	for _, r := range resources.Items {
		rp := UptimeCheckConfigToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListMonitoringAlphaUptimeCheckConfigResponse{Items: protos}, nil
}

func createConfigUptimeCheckConfig(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
