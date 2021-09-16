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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/alpha/iam_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
)

// Server implements the gRPC interface for Role.
type RoleServer struct{}

// ProtoToRoleStageEnum converts a RoleStageEnum enum from its proto representation.
func ProtoToIamAlphaRoleStageEnum(e alphapb.IamAlphaRoleStageEnum) *alpha.RoleStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaRoleStageEnum_name[int32(e)]; ok {
		e := alpha.RoleStageEnum(n[len("IamAlphaRoleStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoleLocalizedValues converts a RoleLocalizedValues resource from its proto representation.
func ProtoToIamAlphaRoleLocalizedValues(p *alphapb.IamAlphaRoleLocalizedValues) *alpha.RoleLocalizedValues {
	if p == nil {
		return nil
	}
	obj := &alpha.RoleLocalizedValues{
		LocalizedTitle:       dcl.StringOrNil(p.LocalizedTitle),
		LocalizedDescription: dcl.StringOrNil(p.LocalizedDescription),
	}
	return obj
}

// ProtoToRole converts a Role resource from its proto representation.
func ProtoToRole(p *alphapb.IamAlphaRole) *alpha.Role {
	obj := &alpha.Role{
		Name:            dcl.StringOrNil(p.Name),
		Title:           dcl.StringOrNil(p.Title),
		Description:     dcl.StringOrNil(p.Description),
		LocalizedValues: ProtoToIamAlphaRoleLocalizedValues(p.GetLocalizedValues()),
		LifecyclePhase:  dcl.StringOrNil(p.LifecyclePhase),
		GroupName:       dcl.StringOrNil(p.GroupName),
		GroupTitle:      dcl.StringOrNil(p.GroupTitle),
		Stage:           ProtoToIamAlphaRoleStageEnum(p.GetStage()),
		Etag:            dcl.StringOrNil(p.Etag),
		Deleted:         dcl.Bool(p.Deleted),
		Parent:          dcl.StringOrNil(p.Parent),
	}
	for _, r := range p.GetIncludedPermissions() {
		obj.IncludedPermissions = append(obj.IncludedPermissions, r)
	}
	for _, r := range p.GetIncludedRoles() {
		obj.IncludedRoles = append(obj.IncludedRoles, r)
	}
	return obj
}

// RoleStageEnumToProto converts a RoleStageEnum enum to its proto representation.
func IamAlphaRoleStageEnumToProto(e *alpha.RoleStageEnum) alphapb.IamAlphaRoleStageEnum {
	if e == nil {
		return alphapb.IamAlphaRoleStageEnum(0)
	}
	if v, ok := alphapb.IamAlphaRoleStageEnum_value["RoleStageEnum"+string(*e)]; ok {
		return alphapb.IamAlphaRoleStageEnum(v)
	}
	return alphapb.IamAlphaRoleStageEnum(0)
}

// RoleLocalizedValuesToProto converts a RoleLocalizedValues resource to its proto representation.
func IamAlphaRoleLocalizedValuesToProto(o *alpha.RoleLocalizedValues) *alphapb.IamAlphaRoleLocalizedValues {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaRoleLocalizedValues{
		LocalizedTitle:       dcl.ValueOrEmptyString(o.LocalizedTitle),
		LocalizedDescription: dcl.ValueOrEmptyString(o.LocalizedDescription),
	}
	return p
}

// RoleToProto converts a Role resource to its proto representation.
func RoleToProto(resource *alpha.Role) *alphapb.IamAlphaRole {
	p := &alphapb.IamAlphaRole{
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Title:           dcl.ValueOrEmptyString(resource.Title),
		Description:     dcl.ValueOrEmptyString(resource.Description),
		LocalizedValues: IamAlphaRoleLocalizedValuesToProto(resource.LocalizedValues),
		LifecyclePhase:  dcl.ValueOrEmptyString(resource.LifecyclePhase),
		GroupName:       dcl.ValueOrEmptyString(resource.GroupName),
		GroupTitle:      dcl.ValueOrEmptyString(resource.GroupTitle),
		Stage:           IamAlphaRoleStageEnumToProto(resource.Stage),
		Etag:            dcl.ValueOrEmptyString(resource.Etag),
		Deleted:         dcl.ValueOrEmptyBool(resource.Deleted),
		Parent:          dcl.ValueOrEmptyString(resource.Parent),
	}
	for _, r := range resource.IncludedPermissions {
		p.IncludedPermissions = append(p.IncludedPermissions, r)
	}
	for _, r := range resource.IncludedRoles {
		p.IncludedRoles = append(p.IncludedRoles, r)
	}

	return p
}

// ApplyRole handles the gRPC request by passing it to the underlying Role Apply() method.
func (s *RoleServer) applyRole(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIamAlphaRoleRequest) (*alphapb.IamAlphaRole, error) {
	p := ProtoToRole(request.GetResource())
	res, err := c.ApplyRole(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RoleToProto(res)
	return r, nil
}

// ApplyRole handles the gRPC request by passing it to the underlying Role Apply() method.
func (s *RoleServer) ApplyIamAlphaRole(ctx context.Context, request *alphapb.ApplyIamAlphaRoleRequest) (*alphapb.IamAlphaRole, error) {
	cl, err := createConfigRole(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRole(ctx, cl, request)
}

// DeleteRole handles the gRPC request by passing it to the underlying Role Delete() method.
func (s *RoleServer) DeleteIamAlphaRole(ctx context.Context, request *alphapb.DeleteIamAlphaRoleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRole(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRole(ctx, ProtoToRole(request.GetResource()))

}

// ListIamAlphaRole handles the gRPC request by passing it to the underlying RoleList() method.
func (s *RoleServer) ListIamAlphaRole(ctx context.Context, request *alphapb.ListIamAlphaRoleRequest) (*alphapb.ListIamAlphaRoleResponse, error) {
	cl, err := createConfigRole(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRole(ctx, ProtoToRole(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IamAlphaRole
	for _, r := range resources.Items {
		rp := RoleToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListIamAlphaRoleResponse{Items: protos}, nil
}

func createConfigRole(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
