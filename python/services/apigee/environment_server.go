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
	apigeepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/apigee_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
)

// Server implements the gRPC interface for Environment.
type EnvironmentServer struct{}

// ProtoToEnvironmentStateEnum converts a EnvironmentStateEnum enum from its proto representation.
func ProtoToApigeeEnvironmentStateEnum(e apigeepb.ApigeeEnvironmentStateEnum) *apigee.EnvironmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeEnvironmentStateEnum_name[int32(e)]; ok {
		e := apigee.EnvironmentStateEnum(n[len("ApigeeEnvironmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironmentProperties converts a EnvironmentProperties resource from its proto representation.
func ProtoToApigeeEnvironmentProperties(p *apigeepb.ApigeeEnvironmentProperties) *apigee.EnvironmentProperties {
	if p == nil {
		return nil
	}
	obj := &apigee.EnvironmentProperties{}
	for _, r := range p.GetProperty() {
		obj.Property = append(obj.Property, *ProtoToApigeeEnvironmentPropertiesProperty(r))
	}
	return obj
}

// ProtoToEnvironmentPropertiesProperty converts a EnvironmentPropertiesProperty resource from its proto representation.
func ProtoToApigeeEnvironmentPropertiesProperty(p *apigeepb.ApigeeEnvironmentPropertiesProperty) *apigee.EnvironmentPropertiesProperty {
	if p == nil {
		return nil
	}
	obj := &apigee.EnvironmentPropertiesProperty{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToEnvironment converts a Environment resource from its proto representation.
func ProtoToEnvironment(p *apigeepb.ApigeeEnvironment) *apigee.Environment {
	obj := &apigee.Environment{
		Name:           dcl.StringOrNil(p.Name),
		Description:    dcl.StringOrNil(p.Description),
		CreatedAt:      dcl.Int64OrNil(p.CreatedAt),
		LastModifiedAt: dcl.Int64OrNil(p.LastModifiedAt),
		Properties:     ProtoToApigeeEnvironmentProperties(p.GetProperties()),
		DisplayName:    dcl.StringOrNil(p.DisplayName),
		State:          ProtoToApigeeEnvironmentStateEnum(p.GetState()),
		Organization:   dcl.StringOrNil(p.Organization),
	}
	return obj
}

// EnvironmentStateEnumToProto converts a EnvironmentStateEnum enum to its proto representation.
func ApigeeEnvironmentStateEnumToProto(e *apigee.EnvironmentStateEnum) apigeepb.ApigeeEnvironmentStateEnum {
	if e == nil {
		return apigeepb.ApigeeEnvironmentStateEnum(0)
	}
	if v, ok := apigeepb.ApigeeEnvironmentStateEnum_value["EnvironmentStateEnum"+string(*e)]; ok {
		return apigeepb.ApigeeEnvironmentStateEnum(v)
	}
	return apigeepb.ApigeeEnvironmentStateEnum(0)
}

// EnvironmentPropertiesToProto converts a EnvironmentProperties resource to its proto representation.
func ApigeeEnvironmentPropertiesToProto(o *apigee.EnvironmentProperties) *apigeepb.ApigeeEnvironmentProperties {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeEnvironmentProperties{}
	for _, r := range o.Property {
		p.Property = append(p.Property, ApigeeEnvironmentPropertiesPropertyToProto(&r))
	}
	return p
}

// EnvironmentPropertiesPropertyToProto converts a EnvironmentPropertiesProperty resource to its proto representation.
func ApigeeEnvironmentPropertiesPropertyToProto(o *apigee.EnvironmentPropertiesProperty) *apigeepb.ApigeeEnvironmentPropertiesProperty {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeEnvironmentPropertiesProperty{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// EnvironmentToProto converts a Environment resource to its proto representation.
func EnvironmentToProto(resource *apigee.Environment) *apigeepb.ApigeeEnvironment {
	p := &apigeepb.ApigeeEnvironment{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		Description:    dcl.ValueOrEmptyString(resource.Description),
		CreatedAt:      dcl.ValueOrEmptyInt64(resource.CreatedAt),
		LastModifiedAt: dcl.ValueOrEmptyInt64(resource.LastModifiedAt),
		Properties:     ApigeeEnvironmentPropertiesToProto(resource.Properties),
		DisplayName:    dcl.ValueOrEmptyString(resource.DisplayName),
		State:          ApigeeEnvironmentStateEnumToProto(resource.State),
		Organization:   dcl.ValueOrEmptyString(resource.Organization),
	}

	return p
}

// ApplyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) applyEnvironment(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeEnvironmentRequest) (*apigeepb.ApigeeEnvironment, error) {
	p := ProtoToEnvironment(request.GetResource())
	res, err := c.ApplyEnvironment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentToProto(res)
	return r, nil
}

// ApplyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) ApplyApigeeEnvironment(ctx context.Context, request *apigeepb.ApplyApigeeEnvironmentRequest) (*apigeepb.ApigeeEnvironment, error) {
	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEnvironment(ctx, cl, request)
}

// DeleteEnvironment handles the gRPC request by passing it to the underlying Environment Delete() method.
func (s *EnvironmentServer) DeleteApigeeEnvironment(ctx context.Context, request *apigeepb.DeleteApigeeEnvironmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironment(ctx, ProtoToEnvironment(request.GetResource()))

}

// ListApigeeEnvironment handles the gRPC request by passing it to the underlying EnvironmentList() method.
func (s *EnvironmentServer) ListApigeeEnvironment(ctx context.Context, request *apigeepb.ListApigeeEnvironmentRequest) (*apigeepb.ListApigeeEnvironmentResponse, error) {
	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironment(ctx, request.Organization)
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeEnvironment
	for _, r := range resources.Items {
		rp := EnvironmentToProto(r)
		protos = append(protos, rp)
	}
	return &apigeepb.ListApigeeEnvironmentResponse{Items: protos}, nil
}

func createConfigEnvironment(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
