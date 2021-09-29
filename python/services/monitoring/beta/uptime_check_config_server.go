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

// Server implements the gRPC interface for UptimeCheckConfig.
type UptimeCheckConfigServer struct{}

// ProtoToUptimeCheckConfigResourceGroupResourceTypeEnum converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum(e betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum) *beta.UptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum_name[int32(e)]; ok {
		e := beta.UptimeCheckConfigResourceGroupResourceTypeEnum(n[len("MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckRequestMethodEnum converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum(e betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum) *beta.UptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum_name[int32(e)]; ok {
		e := beta.UptimeCheckConfigHttpCheckRequestMethodEnum(n[len("MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckContentTypeEnum converts a UptimeCheckConfigHttpCheckContentTypeEnum enum from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum(e betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum) *beta.UptimeCheckConfigHttpCheckContentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum_name[int32(e)]; ok {
		e := beta.UptimeCheckConfigHttpCheckContentTypeEnum(n[len("MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigContentMatchersMatcherEnum converts a UptimeCheckConfigContentMatchersMatcherEnum enum from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum(e betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum) *beta.UptimeCheckConfigContentMatchersMatcherEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum_name[int32(e)]; ok {
		e := beta.UptimeCheckConfigContentMatchersMatcherEnum(n[len("MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigMonitoredResource converts a UptimeCheckConfigMonitoredResource resource from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigMonitoredResource(p *betapb.MonitoringBetaUptimeCheckConfigMonitoredResource) *beta.UptimeCheckConfigMonitoredResource {
	if p == nil {
		return nil
	}
	obj := &beta.UptimeCheckConfigMonitoredResource{
		Type: dcl.StringOrNil(p.Type),
	}
	return obj
}

// ProtoToUptimeCheckConfigResourceGroup converts a UptimeCheckConfigResourceGroup resource from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigResourceGroup(p *betapb.MonitoringBetaUptimeCheckConfigResourceGroup) *beta.UptimeCheckConfigResourceGroup {
	if p == nil {
		return nil
	}
	obj := &beta.UptimeCheckConfigResourceGroup{
		GroupId:      dcl.StringOrNil(p.GroupId),
		ResourceType: ProtoToMonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheck converts a UptimeCheckConfigHttpCheck resource from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigHttpCheck(p *betapb.MonitoringBetaUptimeCheckConfigHttpCheck) *beta.UptimeCheckConfigHttpCheck {
	if p == nil {
		return nil
	}
	obj := &beta.UptimeCheckConfigHttpCheck{
		RequestMethod: ProtoToMonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum(p.GetRequestMethod()),
		UseSsl:        dcl.Bool(p.UseSsl),
		Path:          dcl.StringOrNil(p.Path),
		Port:          dcl.Int64OrNil(p.Port),
		AuthInfo:      ProtoToMonitoringBetaUptimeCheckConfigHttpCheckAuthInfo(p.GetAuthInfo()),
		MaskHeaders:   dcl.Bool(p.MaskHeaders),
		ContentType:   ProtoToMonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum(p.GetContentType()),
		ValidateSsl:   dcl.Bool(p.ValidateSsl),
		Body:          dcl.StringOrNil(p.Body),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheckAuthInfo converts a UptimeCheckConfigHttpCheckAuthInfo resource from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigHttpCheckAuthInfo(p *betapb.MonitoringBetaUptimeCheckConfigHttpCheckAuthInfo) *beta.UptimeCheckConfigHttpCheckAuthInfo {
	if p == nil {
		return nil
	}
	obj := &beta.UptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.StringOrNil(p.Username),
		Password: dcl.StringOrNil(p.Password),
	}
	return obj
}

// ProtoToUptimeCheckConfigTcpCheck converts a UptimeCheckConfigTcpCheck resource from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigTcpCheck(p *betapb.MonitoringBetaUptimeCheckConfigTcpCheck) *beta.UptimeCheckConfigTcpCheck {
	if p == nil {
		return nil
	}
	obj := &beta.UptimeCheckConfigTcpCheck{
		Port: dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToUptimeCheckConfigContentMatchers converts a UptimeCheckConfigContentMatchers resource from its proto representation.
func ProtoToMonitoringBetaUptimeCheckConfigContentMatchers(p *betapb.MonitoringBetaUptimeCheckConfigContentMatchers) *beta.UptimeCheckConfigContentMatchers {
	if p == nil {
		return nil
	}
	obj := &beta.UptimeCheckConfigContentMatchers{
		Content: dcl.StringOrNil(p.Content),
		Matcher: ProtoToMonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum(p.GetMatcher()),
	}
	return obj
}

// ProtoToUptimeCheckConfig converts a UptimeCheckConfig resource from its proto representation.
func ProtoToUptimeCheckConfig(p *betapb.MonitoringBetaUptimeCheckConfig) *beta.UptimeCheckConfig {
	obj := &beta.UptimeCheckConfig{
		Name:              dcl.StringOrNil(p.Name),
		DisplayName:       dcl.StringOrNil(p.DisplayName),
		MonitoredResource: ProtoToMonitoringBetaUptimeCheckConfigMonitoredResource(p.GetMonitoredResource()),
		ResourceGroup:     ProtoToMonitoringBetaUptimeCheckConfigResourceGroup(p.GetResourceGroup()),
		HttpCheck:         ProtoToMonitoringBetaUptimeCheckConfigHttpCheck(p.GetHttpCheck()),
		TcpCheck:          ProtoToMonitoringBetaUptimeCheckConfigTcpCheck(p.GetTcpCheck()),
		Period:            dcl.StringOrNil(p.Period),
		Timeout:           dcl.StringOrNil(p.Timeout),
		Project:           dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetContentMatchers() {
		obj.ContentMatchers = append(obj.ContentMatchers, *ProtoToMonitoringBetaUptimeCheckConfigContentMatchers(r))
	}
	for _, r := range p.GetSelectedRegions() {
		obj.SelectedRegions = append(obj.SelectedRegions, r)
	}
	return obj
}

// UptimeCheckConfigResourceGroupResourceTypeEnumToProto converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum to its proto representation.
func MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnumToProto(e *beta.UptimeCheckConfigResourceGroupResourceTypeEnum) betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == nil {
		return betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum_value["UptimeCheckConfigResourceGroupResourceTypeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum(v)
	}
	return betapb.MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnum(0)
}

// UptimeCheckConfigHttpCheckRequestMethodEnumToProto converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum to its proto representation.
func MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnumToProto(e *beta.UptimeCheckConfigHttpCheckRequestMethodEnum) betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == nil {
		return betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum(0)
	}
	if v, ok := betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum_value["UptimeCheckConfigHttpCheckRequestMethodEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum(v)
	}
	return betapb.MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnum(0)
}

// UptimeCheckConfigHttpCheckContentTypeEnumToProto converts a UptimeCheckConfigHttpCheckContentTypeEnum enum to its proto representation.
func MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnumToProto(e *beta.UptimeCheckConfigHttpCheckContentTypeEnum) betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum {
	if e == nil {
		return betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum_value["UptimeCheckConfigHttpCheckContentTypeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum(v)
	}
	return betapb.MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnum(0)
}

// UptimeCheckConfigContentMatchersMatcherEnumToProto converts a UptimeCheckConfigContentMatchersMatcherEnum enum to its proto representation.
func MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnumToProto(e *beta.UptimeCheckConfigContentMatchersMatcherEnum) betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum {
	if e == nil {
		return betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum(0)
	}
	if v, ok := betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum_value["UptimeCheckConfigContentMatchersMatcherEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum(v)
	}
	return betapb.MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnum(0)
}

// UptimeCheckConfigMonitoredResourceToProto converts a UptimeCheckConfigMonitoredResource resource to its proto representation.
func MonitoringBetaUptimeCheckConfigMonitoredResourceToProto(o *beta.UptimeCheckConfigMonitoredResource) *betapb.MonitoringBetaUptimeCheckConfigMonitoredResource {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaUptimeCheckConfigMonitoredResource{
		Type: dcl.ValueOrEmptyString(o.Type),
	}
	p.FilterLabels = make(map[string]string)
	for k, r := range o.FilterLabels {
		p.FilterLabels[k] = r
	}
	return p
}

// UptimeCheckConfigResourceGroupToProto converts a UptimeCheckConfigResourceGroup resource to its proto representation.
func MonitoringBetaUptimeCheckConfigResourceGroupToProto(o *beta.UptimeCheckConfigResourceGroup) *betapb.MonitoringBetaUptimeCheckConfigResourceGroup {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaUptimeCheckConfigResourceGroup{
		GroupId:      dcl.ValueOrEmptyString(o.GroupId),
		ResourceType: MonitoringBetaUptimeCheckConfigResourceGroupResourceTypeEnumToProto(o.ResourceType),
	}
	return p
}

// UptimeCheckConfigHttpCheckToProto converts a UptimeCheckConfigHttpCheck resource to its proto representation.
func MonitoringBetaUptimeCheckConfigHttpCheckToProto(o *beta.UptimeCheckConfigHttpCheck) *betapb.MonitoringBetaUptimeCheckConfigHttpCheck {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaUptimeCheckConfigHttpCheck{
		RequestMethod: MonitoringBetaUptimeCheckConfigHttpCheckRequestMethodEnumToProto(o.RequestMethod),
		UseSsl:        dcl.ValueOrEmptyBool(o.UseSsl),
		Path:          dcl.ValueOrEmptyString(o.Path),
		Port:          dcl.ValueOrEmptyInt64(o.Port),
		AuthInfo:      MonitoringBetaUptimeCheckConfigHttpCheckAuthInfoToProto(o.AuthInfo),
		MaskHeaders:   dcl.ValueOrEmptyBool(o.MaskHeaders),
		ContentType:   MonitoringBetaUptimeCheckConfigHttpCheckContentTypeEnumToProto(o.ContentType),
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
func MonitoringBetaUptimeCheckConfigHttpCheckAuthInfoToProto(o *beta.UptimeCheckConfigHttpCheckAuthInfo) *betapb.MonitoringBetaUptimeCheckConfigHttpCheckAuthInfo {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaUptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.ValueOrEmptyString(o.Username),
		Password: dcl.ValueOrEmptyString(o.Password),
	}
	return p
}

// UptimeCheckConfigTcpCheckToProto converts a UptimeCheckConfigTcpCheck resource to its proto representation.
func MonitoringBetaUptimeCheckConfigTcpCheckToProto(o *beta.UptimeCheckConfigTcpCheck) *betapb.MonitoringBetaUptimeCheckConfigTcpCheck {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaUptimeCheckConfigTcpCheck{
		Port: dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// UptimeCheckConfigContentMatchersToProto converts a UptimeCheckConfigContentMatchers resource to its proto representation.
func MonitoringBetaUptimeCheckConfigContentMatchersToProto(o *beta.UptimeCheckConfigContentMatchers) *betapb.MonitoringBetaUptimeCheckConfigContentMatchers {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaUptimeCheckConfigContentMatchers{
		Content: dcl.ValueOrEmptyString(o.Content),
		Matcher: MonitoringBetaUptimeCheckConfigContentMatchersMatcherEnumToProto(o.Matcher),
	}
	return p
}

// UptimeCheckConfigToProto converts a UptimeCheckConfig resource to its proto representation.
func UptimeCheckConfigToProto(resource *beta.UptimeCheckConfig) *betapb.MonitoringBetaUptimeCheckConfig {
	p := &betapb.MonitoringBetaUptimeCheckConfig{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		DisplayName:       dcl.ValueOrEmptyString(resource.DisplayName),
		MonitoredResource: MonitoringBetaUptimeCheckConfigMonitoredResourceToProto(resource.MonitoredResource),
		ResourceGroup:     MonitoringBetaUptimeCheckConfigResourceGroupToProto(resource.ResourceGroup),
		HttpCheck:         MonitoringBetaUptimeCheckConfigHttpCheckToProto(resource.HttpCheck),
		TcpCheck:          MonitoringBetaUptimeCheckConfigTcpCheckToProto(resource.TcpCheck),
		Period:            dcl.ValueOrEmptyString(resource.Period),
		Timeout:           dcl.ValueOrEmptyString(resource.Timeout),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.ContentMatchers {
		p.ContentMatchers = append(p.ContentMatchers, MonitoringBetaUptimeCheckConfigContentMatchersToProto(&r))
	}
	for _, r := range resource.SelectedRegions {
		p.SelectedRegions = append(p.SelectedRegions, r)
	}

	return p
}

// ApplyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) applyUptimeCheckConfig(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaUptimeCheckConfigRequest) (*betapb.MonitoringBetaUptimeCheckConfig, error) {
	p := ProtoToUptimeCheckConfig(request.GetResource())
	res, err := c.ApplyUptimeCheckConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UptimeCheckConfigToProto(res)
	return r, nil
}

// ApplyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) ApplyMonitoringBetaUptimeCheckConfig(ctx context.Context, request *betapb.ApplyMonitoringBetaUptimeCheckConfigRequest) (*betapb.MonitoringBetaUptimeCheckConfig, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUptimeCheckConfig(ctx, cl, request)
}

// DeleteUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Delete() method.
func (s *UptimeCheckConfigServer) DeleteMonitoringBetaUptimeCheckConfig(ctx context.Context, request *betapb.DeleteMonitoringBetaUptimeCheckConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUptimeCheckConfig(ctx, ProtoToUptimeCheckConfig(request.GetResource()))

}

// ListMonitoringBetaUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfigList() method.
func (s *UptimeCheckConfigServer) ListMonitoringBetaUptimeCheckConfig(ctx context.Context, request *betapb.ListMonitoringBetaUptimeCheckConfigRequest) (*betapb.ListMonitoringBetaUptimeCheckConfigResponse, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUptimeCheckConfig(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaUptimeCheckConfig
	for _, r := range resources.Items {
		rp := UptimeCheckConfigToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListMonitoringBetaUptimeCheckConfigResponse{Items: protos}, nil
}

func createConfigUptimeCheckConfig(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
