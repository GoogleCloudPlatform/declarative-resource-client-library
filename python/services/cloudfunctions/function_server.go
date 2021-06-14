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
	cloudfunctionspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudfunctions/cloudfunctions_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions"
)

// Server implements the gRPC interface for Function.
type FunctionServer struct{}

// ProtoToFunctionStatusEnum converts a FunctionStatusEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionStatusEnum(e cloudfunctionspb.CloudfunctionsFunctionStatusEnum) *cloudfunctions.FunctionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionStatusEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionStatusEnum(n[len("CloudfunctionsFunctionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionVPCConnectorEgressSettingsEnum converts a FunctionVPCConnectorEgressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionVPCConnectorEgressSettingsEnum(e cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum) *cloudfunctions.FunctionVPCConnectorEgressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionVPCConnectorEgressSettingsEnum(n[len("CloudfunctionsFunctionVPCConnectorEgressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionIngressSettingsEnum converts a FunctionIngressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionIngressSettingsEnum(e cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum) *cloudfunctions.FunctionIngressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionIngressSettingsEnum(n[len("CloudfunctionsFunctionIngressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionSourceRepository converts a FunctionSourceRepository resource from its proto representation.
func ProtoToCloudfunctionsFunctionSourceRepository(p *cloudfunctionspb.CloudfunctionsFunctionSourceRepository) *cloudfunctions.FunctionSourceRepository {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.FunctionSourceRepository{
		Url:         dcl.StringOrNil(p.Url),
		DeployedUrl: dcl.StringOrNil(p.DeployedUrl),
	}
	return obj
}

// ProtoToFunctionHttpsTrigger converts a FunctionHttpsTrigger resource from its proto representation.
func ProtoToCloudfunctionsFunctionHttpsTrigger(p *cloudfunctionspb.CloudfunctionsFunctionHttpsTrigger) *cloudfunctions.FunctionHttpsTrigger {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.FunctionHttpsTrigger{
		Url: dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToFunctionEventTrigger converts a FunctionEventTrigger resource from its proto representation.
func ProtoToCloudfunctionsFunctionEventTrigger(p *cloudfunctionspb.CloudfunctionsFunctionEventTrigger) *cloudfunctions.FunctionEventTrigger {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.FunctionEventTrigger{
		EventType:     dcl.StringOrNil(p.EventType),
		Resource:      dcl.StringOrNil(p.Resource),
		Service:       dcl.StringOrNil(p.Service),
		FailurePolicy: dcl.Bool(p.FailurePolicy),
	}
	return obj
}

// ProtoToFunction converts a Function resource from its proto representation.
func ProtoToFunction(p *cloudfunctionspb.CloudfunctionsFunction) *cloudfunctions.Function {
	obj := &cloudfunctions.Function{
		Name:                       dcl.StringOrNil(p.Name),
		Description:                dcl.StringOrNil(p.Description),
		SourceArchiveUrl:           dcl.StringOrNil(p.SourceArchiveUrl),
		SourceRepository:           ProtoToCloudfunctionsFunctionSourceRepository(p.GetSourceRepository()),
		HttpsTrigger:               ProtoToCloudfunctionsFunctionHttpsTrigger(p.GetHttpsTrigger()),
		EventTrigger:               ProtoToCloudfunctionsFunctionEventTrigger(p.GetEventTrigger()),
		Status:                     ProtoToCloudfunctionsFunctionStatusEnum(p.GetStatus()),
		EntryPoint:                 dcl.StringOrNil(p.EntryPoint),
		Runtime:                    dcl.StringOrNil(p.Runtime),
		Timeout:                    dcl.StringOrNil(p.Timeout),
		AvailableMemoryMb:          dcl.Int64OrNil(p.AvailableMemoryMb),
		ServiceAccountEmail:        dcl.StringOrNil(p.ServiceAccountEmail),
		UpdateTime:                 dcl.StringOrNil(p.UpdateTime),
		VersionId:                  dcl.Int64OrNil(p.VersionId),
		Network:                    dcl.StringOrNil(p.Network),
		MaxInstances:               dcl.Int64OrNil(p.MaxInstances),
		VPCConnector:               dcl.StringOrNil(p.VpcConnector),
		VPCConnectorEgressSettings: ProtoToCloudfunctionsFunctionVPCConnectorEgressSettingsEnum(p.GetVpcConnectorEgressSettings()),
		IngressSettings:            ProtoToCloudfunctionsFunctionIngressSettingsEnum(p.GetIngressSettings()),
		Region:                     dcl.StringOrNil(p.Region),
		Project:                    dcl.StringOrNil(p.Project),
	}
	return obj
}

// FunctionStatusEnumToProto converts a FunctionStatusEnum enum to its proto representation.
func CloudfunctionsFunctionStatusEnumToProto(e *cloudfunctions.FunctionStatusEnum) cloudfunctionspb.CloudfunctionsFunctionStatusEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionStatusEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionStatusEnum_value["FunctionStatusEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionStatusEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionStatusEnum(0)
}

// FunctionVPCConnectorEgressSettingsEnumToProto converts a FunctionVPCConnectorEgressSettingsEnum enum to its proto representation.
func CloudfunctionsFunctionVPCConnectorEgressSettingsEnumToProto(e *cloudfunctions.FunctionVPCConnectorEgressSettingsEnum) cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum_value["FunctionVPCConnectorEgressSettingsEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum(0)
}

// FunctionIngressSettingsEnumToProto converts a FunctionIngressSettingsEnum enum to its proto representation.
func CloudfunctionsFunctionIngressSettingsEnumToProto(e *cloudfunctions.FunctionIngressSettingsEnum) cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum_value["FunctionIngressSettingsEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum(0)
}

// FunctionSourceRepositoryToProto converts a FunctionSourceRepository resource to its proto representation.
func CloudfunctionsFunctionSourceRepositoryToProto(o *cloudfunctions.FunctionSourceRepository) *cloudfunctionspb.CloudfunctionsFunctionSourceRepository {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsFunctionSourceRepository{
		Url:         dcl.ValueOrEmptyString(o.Url),
		DeployedUrl: dcl.ValueOrEmptyString(o.DeployedUrl),
	}
	return p
}

// FunctionHttpsTriggerToProto converts a FunctionHttpsTrigger resource to its proto representation.
func CloudfunctionsFunctionHttpsTriggerToProto(o *cloudfunctions.FunctionHttpsTrigger) *cloudfunctionspb.CloudfunctionsFunctionHttpsTrigger {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsFunctionHttpsTrigger{
		Url: dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// FunctionEventTriggerToProto converts a FunctionEventTrigger resource to its proto representation.
func CloudfunctionsFunctionEventTriggerToProto(o *cloudfunctions.FunctionEventTrigger) *cloudfunctionspb.CloudfunctionsFunctionEventTrigger {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsFunctionEventTrigger{
		EventType:     dcl.ValueOrEmptyString(o.EventType),
		Resource:      dcl.ValueOrEmptyString(o.Resource),
		Service:       dcl.ValueOrEmptyString(o.Service),
		FailurePolicy: dcl.ValueOrEmptyBool(o.FailurePolicy),
	}
	return p
}

// FunctionToProto converts a Function resource to its proto representation.
func FunctionToProto(resource *cloudfunctions.Function) *cloudfunctionspb.CloudfunctionsFunction {
	p := &cloudfunctionspb.CloudfunctionsFunction{
		Name:                       dcl.ValueOrEmptyString(resource.Name),
		Description:                dcl.ValueOrEmptyString(resource.Description),
		SourceArchiveUrl:           dcl.ValueOrEmptyString(resource.SourceArchiveUrl),
		SourceRepository:           CloudfunctionsFunctionSourceRepositoryToProto(resource.SourceRepository),
		HttpsTrigger:               CloudfunctionsFunctionHttpsTriggerToProto(resource.HttpsTrigger),
		EventTrigger:               CloudfunctionsFunctionEventTriggerToProto(resource.EventTrigger),
		Status:                     CloudfunctionsFunctionStatusEnumToProto(resource.Status),
		EntryPoint:                 dcl.ValueOrEmptyString(resource.EntryPoint),
		Runtime:                    dcl.ValueOrEmptyString(resource.Runtime),
		Timeout:                    dcl.ValueOrEmptyString(resource.Timeout),
		AvailableMemoryMb:          dcl.ValueOrEmptyInt64(resource.AvailableMemoryMb),
		ServiceAccountEmail:        dcl.ValueOrEmptyString(resource.ServiceAccountEmail),
		UpdateTime:                 dcl.ValueOrEmptyString(resource.UpdateTime),
		VersionId:                  dcl.ValueOrEmptyInt64(resource.VersionId),
		Network:                    dcl.ValueOrEmptyString(resource.Network),
		MaxInstances:               dcl.ValueOrEmptyInt64(resource.MaxInstances),
		VpcConnector:               dcl.ValueOrEmptyString(resource.VPCConnector),
		VpcConnectorEgressSettings: CloudfunctionsFunctionVPCConnectorEgressSettingsEnumToProto(resource.VPCConnectorEgressSettings),
		IngressSettings:            CloudfunctionsFunctionIngressSettingsEnumToProto(resource.IngressSettings),
		Region:                     dcl.ValueOrEmptyString(resource.Region),
		Project:                    dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyFunction handles the gRPC request by passing it to the underlying Function Apply() method.
func (s *FunctionServer) applyFunction(ctx context.Context, c *cloudfunctions.Client, request *cloudfunctionspb.ApplyCloudfunctionsFunctionRequest) (*cloudfunctionspb.CloudfunctionsFunction, error) {
	p := ProtoToFunction(request.GetResource())
	res, err := c.ApplyFunction(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FunctionToProto(res)
	return r, nil
}

// ApplyFunction handles the gRPC request by passing it to the underlying Function Apply() method.
func (s *FunctionServer) ApplyCloudfunctionsFunction(ctx context.Context, request *cloudfunctionspb.ApplyCloudfunctionsFunctionRequest) (*cloudfunctionspb.CloudfunctionsFunction, error) {
	cl, err := createConfigFunction(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFunction(ctx, cl, request)
}

// DeleteFunction handles the gRPC request by passing it to the underlying Function Delete() method.
func (s *FunctionServer) DeleteCloudfunctionsFunction(ctx context.Context, request *cloudfunctionspb.DeleteCloudfunctionsFunctionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFunction(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFunction(ctx, ProtoToFunction(request.GetResource()))

}

// ListCloudfunctionsFunction handles the gRPC request by passing it to the underlying FunctionList() method.
func (s *FunctionServer) ListCloudfunctionsFunction(ctx context.Context, request *cloudfunctionspb.ListCloudfunctionsFunctionRequest) (*cloudfunctionspb.ListCloudfunctionsFunctionResponse, error) {
	cl, err := createConfigFunction(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFunction(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*cloudfunctionspb.CloudfunctionsFunction
	for _, r := range resources.Items {
		rp := FunctionToProto(r)
		protos = append(protos, rp)
	}
	return &cloudfunctionspb.ListCloudfunctionsFunctionResponse{Items: protos}, nil
}

func createConfigFunction(ctx context.Context, service_account_file string) (*cloudfunctions.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudfunctions.NewClient(conf), nil
}
