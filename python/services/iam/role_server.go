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
	iampb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/iam_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
)

// Server implements the gRPC interface for Role.
type RoleServer struct{}

// ProtoToRoleStageEnum converts a RoleStageEnum enum from its proto representation.
func ProtoToIamRoleStageEnum(e iampb.IamRoleStageEnum) *iam.RoleStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := iampb.IamRoleStageEnum_name[int32(e)]; ok {
		e := iam.RoleStageEnum(n[len("IamRoleStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoleLocalizedValues converts a RoleLocalizedValues resource from its proto representation.
func ProtoToIamRoleLocalizedValues(p *iampb.IamRoleLocalizedValues) *iam.RoleLocalizedValues {
	if p == nil {
		return nil
	}
	obj := &iam.RoleLocalizedValues{
		LocalizedTitle:       dcl.StringOrNil(p.LocalizedTitle),
		LocalizedDescription: dcl.StringOrNil(p.LocalizedDescription),
	}
	return obj
}

// ProtoToRole converts a Role resource from its proto representation.
func ProtoToRole(p *iampb.IamRole) *iam.Role {
	obj := &iam.Role{
		Name:            dcl.StringOrNil(p.Name),
		Title:           dcl.StringOrNil(p.Title),
		Description:     dcl.StringOrNil(p.Description),
		LocalizedValues: ProtoToIamRoleLocalizedValues(p.GetLocalizedValues()),
		LifecyclePhase:  dcl.StringOrNil(p.LifecyclePhase),
		GroupName:       dcl.StringOrNil(p.GroupName),
		GroupTitle:      dcl.StringOrNil(p.GroupTitle),
		Stage:           ProtoToIamRoleStageEnum(p.GetStage()),
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
func IamRoleStageEnumToProto(e *iam.RoleStageEnum) iampb.IamRoleStageEnum {
	if e == nil {
		return iampb.IamRoleStageEnum(0)
	}
	if v, ok := iampb.IamRoleStageEnum_value["RoleStageEnum"+string(*e)]; ok {
		return iampb.IamRoleStageEnum(v)
	}
	return iampb.IamRoleStageEnum(0)
}

// RoleLocalizedValuesToProto converts a RoleLocalizedValues resource to its proto representation.
func IamRoleLocalizedValuesToProto(o *iam.RoleLocalizedValues) *iampb.IamRoleLocalizedValues {
	if o == nil {
		return nil
	}
	p := &iampb.IamRoleLocalizedValues{
		LocalizedTitle:       dcl.ValueOrEmptyString(o.LocalizedTitle),
		LocalizedDescription: dcl.ValueOrEmptyString(o.LocalizedDescription),
	}
	return p
}

// RoleToProto converts a Role resource to its proto representation.
func RoleToProto(resource *iam.Role) *iampb.IamRole {
	p := &iampb.IamRole{
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Title:           dcl.ValueOrEmptyString(resource.Title),
		Description:     dcl.ValueOrEmptyString(resource.Description),
		LocalizedValues: IamRoleLocalizedValuesToProto(resource.LocalizedValues),
		LifecyclePhase:  dcl.ValueOrEmptyString(resource.LifecyclePhase),
		GroupName:       dcl.ValueOrEmptyString(resource.GroupName),
		GroupTitle:      dcl.ValueOrEmptyString(resource.GroupTitle),
		Stage:           IamRoleStageEnumToProto(resource.Stage),
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
func (s *RoleServer) applyRole(ctx context.Context, c *iam.Client, request *iampb.ApplyIamRoleRequest) (*iampb.IamRole, error) {
	p := ProtoToRole(request.GetResource())
	res, err := c.ApplyRole(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RoleToProto(res)
	return r, nil
}

// ApplyRole handles the gRPC request by passing it to the underlying Role Apply() method.
func (s *RoleServer) ApplyIamRole(ctx context.Context, request *iampb.ApplyIamRoleRequest) (*iampb.IamRole, error) {
	cl, err := createConfigRole(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRole(ctx, cl, request)
}

// DeleteRole handles the gRPC request by passing it to the underlying Role Delete() method.
func (s *RoleServer) DeleteIamRole(ctx context.Context, request *iampb.DeleteIamRoleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRole(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRole(ctx, ProtoToRole(request.GetResource()))

}

// ListRole handles the gRPC request by passing it to the underlying RoleList() method.
func (s *RoleServer) ListIamRole(ctx context.Context, request *iampb.ListIamRoleRequest) (*iampb.ListIamRoleResponse, error) {
	cl, err := createConfigRole(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRole(ctx, request.Parent)
	if err != nil {
		return nil, err
	}
	var protos []*iampb.IamRole
	for _, r := range resources.Items {
		rp := RoleToProto(r)
		protos = append(protos, rp)
	}
	return &iampb.ListIamRoleResponse{Items: protos}, nil
}

func createConfigRole(ctx context.Context, service_account_file string) (*iam.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iam.NewClient(conf), nil
}
