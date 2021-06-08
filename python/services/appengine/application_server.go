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
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	appenginepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/appengine/appengine_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine"
)

// Server implements the gRPC interface for Application.
type ApplicationServer struct{}

// ProtoToApplicationDatabaseTypeEnum converts a ApplicationDatabaseTypeEnum enum from its proto representation.
func ProtoToAppengineApplicationDatabaseTypeEnum(e appenginepb.AppengineApplicationDatabaseTypeEnum) *appengine.ApplicationDatabaseTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineApplicationDatabaseTypeEnum_name[int32(e)]; ok {
		e := appengine.ApplicationDatabaseTypeEnum(n[len("AppengineApplicationDatabaseTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToApplicationServingStatusEnum converts a ApplicationServingStatusEnum enum from its proto representation.
func ProtoToAppengineApplicationServingStatusEnum(e appenginepb.AppengineApplicationServingStatusEnum) *appengine.ApplicationServingStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineApplicationServingStatusEnum_name[int32(e)]; ok {
		e := appengine.ApplicationServingStatusEnum(n[len("AppengineApplicationServingStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToApplicationDispatchRules converts a ApplicationDispatchRules resource from its proto representation.
func ProtoToAppengineApplicationDispatchRules(p *appenginepb.AppengineApplicationDispatchRules) *appengine.ApplicationDispatchRules {
	if p == nil {
		return nil
	}
	obj := &appengine.ApplicationDispatchRules{
		Domain:  dcl.StringOrNil(p.Domain),
		Path:    dcl.StringOrNil(p.Path),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToApplicationFeatureSettings converts a ApplicationFeatureSettings resource from its proto representation.
func ProtoToAppengineApplicationFeatureSettings(p *appenginepb.AppengineApplicationFeatureSettings) *appengine.ApplicationFeatureSettings {
	if p == nil {
		return nil
	}
	obj := &appengine.ApplicationFeatureSettings{
		SplitHealthChecks:       dcl.Bool(p.SplitHealthChecks),
		UseContainerOptimizedOs: dcl.Bool(p.UseContainerOptimizedOs),
	}
	return obj
}

// ProtoToApplicationIap converts a ApplicationIap resource from its proto representation.
func ProtoToAppengineApplicationIap(p *appenginepb.AppengineApplicationIap) *appengine.ApplicationIap {
	if p == nil {
		return nil
	}
	obj := &appengine.ApplicationIap{
		Enabled:                  dcl.Bool(p.Enabled),
		OAuth2ClientId:           dcl.StringOrNil(p.Oauth2ClientId),
		OAuth2ClientSecret:       dcl.StringOrNil(p.Oauth2ClientSecret),
		OAuth2ClientSecretSha256: dcl.StringOrNil(p.Oauth2ClientSecretSha256),
	}
	return obj
}

// ProtoToApplication converts a Application resource from its proto representation.
func ProtoToApplication(p *appenginepb.AppengineApplication) *appengine.Application {
	obj := &appengine.Application{
		AuthDomain:      dcl.StringOrNil(p.AuthDomain),
		CodeBucket:      dcl.StringOrNil(p.CodeBucket),
		DatabaseType:    ProtoToAppengineApplicationDatabaseTypeEnum(p.GetDatabaseType()),
		DefaultBucket:   dcl.StringOrNil(p.DefaultBucket),
		DefaultHostname: dcl.StringOrNil(p.DefaultHostname),
		FeatureSettings: ProtoToAppengineApplicationFeatureSettings(p.GetFeatureSettings()),
		GcrDomain:       dcl.StringOrNil(p.GcrDomain),
		Iap:             ProtoToAppengineApplicationIap(p.GetIap()),
		Name:            dcl.StringOrNil(p.Name),
		Location:        dcl.StringOrNil(p.Location),
		ServingStatus:   ProtoToAppengineApplicationServingStatusEnum(p.GetServingStatus()),
	}
	for _, r := range p.GetDispatchRules() {
		obj.DispatchRules = append(obj.DispatchRules, *ProtoToAppengineApplicationDispatchRules(r))
	}
	return obj
}

// ApplicationDatabaseTypeEnumToProto converts a ApplicationDatabaseTypeEnum enum to its proto representation.
func AppengineApplicationDatabaseTypeEnumToProto(e *appengine.ApplicationDatabaseTypeEnum) appenginepb.AppengineApplicationDatabaseTypeEnum {
	if e == nil {
		return appenginepb.AppengineApplicationDatabaseTypeEnum(0)
	}
	if v, ok := appenginepb.AppengineApplicationDatabaseTypeEnum_value["ApplicationDatabaseTypeEnum"+string(*e)]; ok {
		return appenginepb.AppengineApplicationDatabaseTypeEnum(v)
	}
	return appenginepb.AppengineApplicationDatabaseTypeEnum(0)
}

// ApplicationServingStatusEnumToProto converts a ApplicationServingStatusEnum enum to its proto representation.
func AppengineApplicationServingStatusEnumToProto(e *appengine.ApplicationServingStatusEnum) appenginepb.AppengineApplicationServingStatusEnum {
	if e == nil {
		return appenginepb.AppengineApplicationServingStatusEnum(0)
	}
	if v, ok := appenginepb.AppengineApplicationServingStatusEnum_value["ApplicationServingStatusEnum"+string(*e)]; ok {
		return appenginepb.AppengineApplicationServingStatusEnum(v)
	}
	return appenginepb.AppengineApplicationServingStatusEnum(0)
}

// ApplicationDispatchRulesToProto converts a ApplicationDispatchRules resource to its proto representation.
func AppengineApplicationDispatchRulesToProto(o *appengine.ApplicationDispatchRules) *appenginepb.AppengineApplicationDispatchRules {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineApplicationDispatchRules{
		Domain:  dcl.ValueOrEmptyString(o.Domain),
		Path:    dcl.ValueOrEmptyString(o.Path),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	return p
}

// ApplicationFeatureSettingsToProto converts a ApplicationFeatureSettings resource to its proto representation.
func AppengineApplicationFeatureSettingsToProto(o *appengine.ApplicationFeatureSettings) *appenginepb.AppengineApplicationFeatureSettings {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineApplicationFeatureSettings{
		SplitHealthChecks:       dcl.ValueOrEmptyBool(o.SplitHealthChecks),
		UseContainerOptimizedOs: dcl.ValueOrEmptyBool(o.UseContainerOptimizedOs),
	}
	return p
}

// ApplicationIapToProto converts a ApplicationIap resource to its proto representation.
func AppengineApplicationIapToProto(o *appengine.ApplicationIap) *appenginepb.AppengineApplicationIap {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineApplicationIap{
		Enabled:                  dcl.ValueOrEmptyBool(o.Enabled),
		Oauth2ClientId:           dcl.ValueOrEmptyString(o.OAuth2ClientId),
		Oauth2ClientSecret:       dcl.ValueOrEmptyString(o.OAuth2ClientSecret),
		Oauth2ClientSecretSha256: dcl.ValueOrEmptyString(o.OAuth2ClientSecretSha256),
	}
	return p
}

// ApplicationToProto converts a Application resource to its proto representation.
func ApplicationToProto(resource *appengine.Application) *appenginepb.AppengineApplication {
	p := &appenginepb.AppengineApplication{
		AuthDomain:      dcl.ValueOrEmptyString(resource.AuthDomain),
		CodeBucket:      dcl.ValueOrEmptyString(resource.CodeBucket),
		DatabaseType:    AppengineApplicationDatabaseTypeEnumToProto(resource.DatabaseType),
		DefaultBucket:   dcl.ValueOrEmptyString(resource.DefaultBucket),
		DefaultHostname: dcl.ValueOrEmptyString(resource.DefaultHostname),
		FeatureSettings: AppengineApplicationFeatureSettingsToProto(resource.FeatureSettings),
		GcrDomain:       dcl.ValueOrEmptyString(resource.GcrDomain),
		Iap:             AppengineApplicationIapToProto(resource.Iap),
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Location:        dcl.ValueOrEmptyString(resource.Location),
		ServingStatus:   AppengineApplicationServingStatusEnumToProto(resource.ServingStatus),
	}
	for _, r := range resource.DispatchRules {
		p.DispatchRules = append(p.DispatchRules, AppengineApplicationDispatchRulesToProto(&r))
	}

	return p
}

// ApplyApplication handles the gRPC request by passing it to the underlying Application Apply() method.
func (s *ApplicationServer) applyApplication(ctx context.Context, c *appengine.Client, request *appenginepb.ApplyAppengineApplicationRequest) (*appenginepb.AppengineApplication, error) {
	p := ProtoToApplication(request.GetResource())
	res, err := c.ApplyApplication(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ApplicationToProto(res)
	return r, nil
}

// ApplyApplication handles the gRPC request by passing it to the underlying Application Apply() method.
func (s *ApplicationServer) ApplyAppengineApplication(ctx context.Context, request *appenginepb.ApplyAppengineApplicationRequest) (*appenginepb.AppengineApplication, error) {
	cl, err := createConfigApplication(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyApplication(ctx, cl, request)
}

// DeleteApplication handles the gRPC request by passing it to the underlying Application Delete() method.
func (s *ApplicationServer) DeleteAppengineApplication(ctx context.Context, request *appenginepb.DeleteAppengineApplicationRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Application")

}

// ListApplication handles the gRPC request by passing it to the underlying ApplicationList() method.
func (s *ApplicationServer) ListAppengineApplication(ctx context.Context, request *appenginepb.ListAppengineApplicationRequest) (*appenginepb.ListAppengineApplicationResponse, error) {
	cl, err := createConfigApplication(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListApplication(ctx)
	if err != nil {
		return nil, err
	}
	var protos []*appenginepb.AppengineApplication
	for _, r := range resources.Items {
		rp := ApplicationToProto(r)
		protos = append(protos, rp)
	}
	return &appenginepb.ListAppengineApplicationResponse{Items: protos}, nil
}

func createConfigApplication(ctx context.Context, service_account_file string) (*appengine.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return appengine.NewClient(conf), nil
}
