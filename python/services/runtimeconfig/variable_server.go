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
	runtimeconfigpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/runtimeconfig/runtimeconfig_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/runtimeconfig"
)

// Server implements the gRPC interface for Variable.
type VariableServer struct{}

// ProtoToVariable converts a Variable resource from its proto representation.
func ProtoToVariable(p *runtimeconfigpb.RuntimeconfigVariable) *runtimeconfig.Variable {
	obj := &runtimeconfig.Variable{
		Name:          dcl.StringOrNil(p.Name),
		RuntimeConfig: dcl.StringOrNil(p.RuntimeConfig),
		Text:          dcl.StringOrNil(p.Text),
		Value:         dcl.StringOrNil(p.Value),
		UpdateTime:    dcl.StringOrNil(p.UpdateTime),
		Project:       dcl.StringOrNil(p.Project),
	}
	return obj
}

// VariableToProto converts a Variable resource to its proto representation.
func VariableToProto(resource *runtimeconfig.Variable) *runtimeconfigpb.RuntimeconfigVariable {
	p := &runtimeconfigpb.RuntimeconfigVariable{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		RuntimeConfig: dcl.ValueOrEmptyString(resource.RuntimeConfig),
		Text:          dcl.ValueOrEmptyString(resource.Text),
		Value:         dcl.ValueOrEmptyString(resource.Value),
		UpdateTime:    dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:       dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyVariable handles the gRPC request by passing it to the underlying Variable Apply() method.
func (s *VariableServer) applyVariable(ctx context.Context, c *runtimeconfig.Client, request *runtimeconfigpb.ApplyRuntimeconfigVariableRequest) (*runtimeconfigpb.RuntimeconfigVariable, error) {
	p := ProtoToVariable(request.GetResource())
	res, err := c.ApplyVariable(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VariableToProto(res)
	return r, nil
}

// ApplyVariable handles the gRPC request by passing it to the underlying Variable Apply() method.
func (s *VariableServer) ApplyRuntimeconfigVariable(ctx context.Context, request *runtimeconfigpb.ApplyRuntimeconfigVariableRequest) (*runtimeconfigpb.RuntimeconfigVariable, error) {
	cl, err := createConfigVariable(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyVariable(ctx, cl, request)
}

// DeleteVariable handles the gRPC request by passing it to the underlying Variable Delete() method.
func (s *VariableServer) DeleteRuntimeconfigVariable(ctx context.Context, request *runtimeconfigpb.DeleteRuntimeconfigVariableRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVariable(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVariable(ctx, ProtoToVariable(request.GetResource()))

}

// ListVariable handles the gRPC request by passing it to the underlying VariableList() method.
func (s *VariableServer) ListRuntimeconfigVariable(ctx context.Context, request *runtimeconfigpb.ListRuntimeconfigVariableRequest) (*runtimeconfigpb.ListRuntimeconfigVariableResponse, error) {
	cl, err := createConfigVariable(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVariable(ctx, request.Project, request.RuntimeConfig)
	if err != nil {
		return nil, err
	}
	var protos []*runtimeconfigpb.RuntimeconfigVariable
	for _, r := range resources.Items {
		rp := VariableToProto(r)
		protos = append(protos, rp)
	}
	return &runtimeconfigpb.ListRuntimeconfigVariableResponse{Items: protos}, nil
}

func createConfigVariable(ctx context.Context, service_account_file string) (*runtimeconfig.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return runtimeconfig.NewClient(conf), nil
}
