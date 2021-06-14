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

// Server implements the gRPC interface for Organization.
type OrganizationServer struct{}

// ProtoToOrganizationRuntimeTypeEnum converts a OrganizationRuntimeTypeEnum enum from its proto representation.
func ProtoToApigeeOrganizationRuntimeTypeEnum(e apigeepb.ApigeeOrganizationRuntimeTypeEnum) *apigee.OrganizationRuntimeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationRuntimeTypeEnum_name[int32(e)]; ok {
		e := apigee.OrganizationRuntimeTypeEnum(n[len("ApigeeOrganizationRuntimeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationSubscriptionTypeEnum converts a OrganizationSubscriptionTypeEnum enum from its proto representation.
func ProtoToApigeeOrganizationSubscriptionTypeEnum(e apigeepb.ApigeeOrganizationSubscriptionTypeEnum) *apigee.OrganizationSubscriptionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationSubscriptionTypeEnum_name[int32(e)]; ok {
		e := apigee.OrganizationSubscriptionTypeEnum(n[len("ApigeeOrganizationSubscriptionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationBillingTypeEnum converts a OrganizationBillingTypeEnum enum from its proto representation.
func ProtoToApigeeOrganizationBillingTypeEnum(e apigeepb.ApigeeOrganizationBillingTypeEnum) *apigee.OrganizationBillingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationBillingTypeEnum_name[int32(e)]; ok {
		e := apigee.OrganizationBillingTypeEnum(n[len("ApigeeOrganizationBillingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationStateEnum converts a OrganizationStateEnum enum from its proto representation.
func ProtoToApigeeOrganizationStateEnum(e apigeepb.ApigeeOrganizationStateEnum) *apigee.OrganizationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationStateEnum_name[int32(e)]; ok {
		e := apigee.OrganizationStateEnum(n[len("ApigeeOrganizationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationProperties converts a OrganizationProperties resource from its proto representation.
func ProtoToApigeeOrganizationProperties(p *apigeepb.ApigeeOrganizationProperties) *apigee.OrganizationProperties {
	if p == nil {
		return nil
	}
	obj := &apigee.OrganizationProperties{}
	for _, r := range p.GetProperty() {
		obj.Property = append(obj.Property, *ProtoToApigeeOrganizationPropertiesProperty(r))
	}
	return obj
}

// ProtoToOrganizationPropertiesProperty converts a OrganizationPropertiesProperty resource from its proto representation.
func ProtoToApigeeOrganizationPropertiesProperty(p *apigeepb.ApigeeOrganizationPropertiesProperty) *apigee.OrganizationPropertiesProperty {
	if p == nil {
		return nil
	}
	obj := &apigee.OrganizationPropertiesProperty{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToOrganization converts a Organization resource from its proto representation.
func ProtoToOrganization(p *apigeepb.ApigeeOrganization) *apigee.Organization {
	obj := &apigee.Organization{
		Name:                             dcl.StringOrNil(p.Name),
		DisplayName:                      dcl.StringOrNil(p.DisplayName),
		Description:                      dcl.StringOrNil(p.Description),
		CreatedAt:                        dcl.Int64OrNil(p.CreatedAt),
		LastModifiedAt:                   dcl.Int64OrNil(p.LastModifiedAt),
		ExpiresAt:                        dcl.Int64OrNil(p.ExpiresAt),
		Properties:                       ProtoToApigeeOrganizationProperties(p.GetProperties()),
		AnalyticsRegion:                  dcl.StringOrNil(p.AnalyticsRegion),
		AuthorizedNetwork:                dcl.StringOrNil(p.AuthorizedNetwork),
		RuntimeType:                      ProtoToApigeeOrganizationRuntimeTypeEnum(p.GetRuntimeType()),
		SubscriptionType:                 ProtoToApigeeOrganizationSubscriptionTypeEnum(p.GetSubscriptionType()),
		BillingType:                      ProtoToApigeeOrganizationBillingTypeEnum(p.GetBillingType()),
		CaCertificate:                    dcl.StringOrNil(p.CaCertificate),
		RuntimeDatabaseEncryptionKeyName: dcl.StringOrNil(p.RuntimeDatabaseEncryptionKeyName),
		ProjectId:                        dcl.StringOrNil(p.ProjectId),
		State:                            ProtoToApigeeOrganizationStateEnum(p.GetState()),
		Parent:                           dcl.StringOrNil(p.Parent),
	}
	for _, r := range p.GetEnvironments() {
		obj.Environments = append(obj.Environments, r)
	}
	return obj
}

// OrganizationRuntimeTypeEnumToProto converts a OrganizationRuntimeTypeEnum enum to its proto representation.
func ApigeeOrganizationRuntimeTypeEnumToProto(e *apigee.OrganizationRuntimeTypeEnum) apigeepb.ApigeeOrganizationRuntimeTypeEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationRuntimeTypeEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationRuntimeTypeEnum_value["OrganizationRuntimeTypeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationRuntimeTypeEnum(v)
	}
	return apigeepb.ApigeeOrganizationRuntimeTypeEnum(0)
}

// OrganizationSubscriptionTypeEnumToProto converts a OrganizationSubscriptionTypeEnum enum to its proto representation.
func ApigeeOrganizationSubscriptionTypeEnumToProto(e *apigee.OrganizationSubscriptionTypeEnum) apigeepb.ApigeeOrganizationSubscriptionTypeEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationSubscriptionTypeEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationSubscriptionTypeEnum_value["OrganizationSubscriptionTypeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationSubscriptionTypeEnum(v)
	}
	return apigeepb.ApigeeOrganizationSubscriptionTypeEnum(0)
}

// OrganizationBillingTypeEnumToProto converts a OrganizationBillingTypeEnum enum to its proto representation.
func ApigeeOrganizationBillingTypeEnumToProto(e *apigee.OrganizationBillingTypeEnum) apigeepb.ApigeeOrganizationBillingTypeEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationBillingTypeEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationBillingTypeEnum_value["OrganizationBillingTypeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationBillingTypeEnum(v)
	}
	return apigeepb.ApigeeOrganizationBillingTypeEnum(0)
}

// OrganizationStateEnumToProto converts a OrganizationStateEnum enum to its proto representation.
func ApigeeOrganizationStateEnumToProto(e *apigee.OrganizationStateEnum) apigeepb.ApigeeOrganizationStateEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationStateEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationStateEnum_value["OrganizationStateEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationStateEnum(v)
	}
	return apigeepb.ApigeeOrganizationStateEnum(0)
}

// OrganizationPropertiesToProto converts a OrganizationProperties resource to its proto representation.
func ApigeeOrganizationPropertiesToProto(o *apigee.OrganizationProperties) *apigeepb.ApigeeOrganizationProperties {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeOrganizationProperties{}
	for _, r := range o.Property {
		p.Property = append(p.Property, ApigeeOrganizationPropertiesPropertyToProto(&r))
	}
	return p
}

// OrganizationPropertiesPropertyToProto converts a OrganizationPropertiesProperty resource to its proto representation.
func ApigeeOrganizationPropertiesPropertyToProto(o *apigee.OrganizationPropertiesProperty) *apigeepb.ApigeeOrganizationPropertiesProperty {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeOrganizationPropertiesProperty{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// OrganizationToProto converts a Organization resource to its proto representation.
func OrganizationToProto(resource *apigee.Organization) *apigeepb.ApigeeOrganization {
	p := &apigeepb.ApigeeOrganization{
		Name:                             dcl.ValueOrEmptyString(resource.Name),
		DisplayName:                      dcl.ValueOrEmptyString(resource.DisplayName),
		Description:                      dcl.ValueOrEmptyString(resource.Description),
		CreatedAt:                        dcl.ValueOrEmptyInt64(resource.CreatedAt),
		LastModifiedAt:                   dcl.ValueOrEmptyInt64(resource.LastModifiedAt),
		ExpiresAt:                        dcl.ValueOrEmptyInt64(resource.ExpiresAt),
		Properties:                       ApigeeOrganizationPropertiesToProto(resource.Properties),
		AnalyticsRegion:                  dcl.ValueOrEmptyString(resource.AnalyticsRegion),
		AuthorizedNetwork:                dcl.ValueOrEmptyString(resource.AuthorizedNetwork),
		RuntimeType:                      ApigeeOrganizationRuntimeTypeEnumToProto(resource.RuntimeType),
		SubscriptionType:                 ApigeeOrganizationSubscriptionTypeEnumToProto(resource.SubscriptionType),
		BillingType:                      ApigeeOrganizationBillingTypeEnumToProto(resource.BillingType),
		CaCertificate:                    dcl.ValueOrEmptyString(resource.CaCertificate),
		RuntimeDatabaseEncryptionKeyName: dcl.ValueOrEmptyString(resource.RuntimeDatabaseEncryptionKeyName),
		ProjectId:                        dcl.ValueOrEmptyString(resource.ProjectId),
		State:                            ApigeeOrganizationStateEnumToProto(resource.State),
		Parent:                           dcl.ValueOrEmptyString(resource.Parent),
	}
	for _, r := range resource.Environments {
		p.Environments = append(p.Environments, r)
	}

	return p
}

// ApplyOrganization handles the gRPC request by passing it to the underlying Organization Apply() method.
func (s *OrganizationServer) applyOrganization(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeOrganizationRequest) (*apigeepb.ApigeeOrganization, error) {
	p := ProtoToOrganization(request.GetResource())
	res, err := c.ApplyOrganization(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OrganizationToProto(res)
	return r, nil
}

// ApplyOrganization handles the gRPC request by passing it to the underlying Organization Apply() method.
func (s *OrganizationServer) ApplyApigeeOrganization(ctx context.Context, request *apigeepb.ApplyApigeeOrganizationRequest) (*apigeepb.ApigeeOrganization, error) {
	cl, err := createConfigOrganization(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyOrganization(ctx, cl, request)
}

// DeleteOrganization handles the gRPC request by passing it to the underlying Organization Delete() method.
func (s *OrganizationServer) DeleteApigeeOrganization(ctx context.Context, request *apigeepb.DeleteApigeeOrganizationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOrganization(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOrganization(ctx, ProtoToOrganization(request.GetResource()))

}

// ListApigeeOrganization handles the gRPC request by passing it to the underlying OrganizationList() method.
func (s *OrganizationServer) ListApigeeOrganization(ctx context.Context, request *apigeepb.ListApigeeOrganizationRequest) (*apigeepb.ListApigeeOrganizationResponse, error) {
	cl, err := createConfigOrganization(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOrganization(ctx)
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeOrganization
	for _, r := range resources.Items {
		rp := OrganizationToProto(r)
		protos = append(protos, rp)
	}
	return &apigeepb.ListApigeeOrganizationResponse{Items: protos}, nil
}

func createConfigOrganization(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
