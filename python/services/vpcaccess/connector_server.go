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
	vpcaccesspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vpcaccess/vpcaccess_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess"
)

// Server implements the gRPC interface for Connector.
type ConnectorServer struct{}

// ProtoToConnectorStateEnum converts a ConnectorStateEnum enum from its proto representation.
func ProtoToVpcaccessConnectorStateEnum(e vpcaccesspb.VpcaccessConnectorStateEnum) *vpcaccess.ConnectorStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := vpcaccesspb.VpcaccessConnectorStateEnum_name[int32(e)]; ok {
		e := vpcaccess.ConnectorStateEnum(n[len("ConnectorStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectorSubnet converts a ConnectorSubnet resource from its proto representation.
func ProtoToVpcaccessConnectorSubnet(p *vpcaccesspb.VpcaccessConnectorSubnet) *vpcaccess.ConnectorSubnet {
	if p == nil {
		return nil
	}
	obj := &vpcaccess.ConnectorSubnet{
		Name:      dcl.StringOrNil(p.Name),
		ProjectId: dcl.StringOrNil(p.ProjectId),
	}
	return obj
}

// ProtoToConnector converts a Connector resource from its proto representation.
func ProtoToConnector(p *vpcaccesspb.VpcaccessConnector) *vpcaccess.Connector {
	obj := &vpcaccess.Connector{
		Name:          dcl.StringOrNil(p.Name),
		Network:       dcl.StringOrNil(p.Network),
		IPCidrRange:   dcl.StringOrNil(p.IpCidrRange),
		State:         ProtoToVpcaccessConnectorStateEnum(p.GetState()),
		MinThroughput: dcl.Int64OrNil(p.MinThroughput),
		MaxThroughput: dcl.Int64OrNil(p.MaxThroughput),
		Subnet:        ProtoToVpcaccessConnectorSubnet(p.GetSubnet()),
		MachineType:   dcl.StringOrNil(p.MachineType),
		MinInstances:  dcl.Int64OrNil(p.MinInstances),
		MaxInstances:  dcl.Int64OrNil(p.MaxInstances),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetConnectedProjects() {
		obj.ConnectedProjects = append(obj.ConnectedProjects, r)
	}
	return obj
}

// ConnectorStateEnumToProto converts a ConnectorStateEnum enum to its proto representation.
func VpcaccessConnectorStateEnumToProto(e *vpcaccess.ConnectorStateEnum) vpcaccesspb.VpcaccessConnectorStateEnum {
	if e == nil {
		return vpcaccesspb.VpcaccessConnectorStateEnum(0)
	}
	if v, ok := vpcaccesspb.VpcaccessConnectorStateEnum_value["ConnectorStateEnum"+string(*e)]; ok {
		return vpcaccesspb.VpcaccessConnectorStateEnum(v)
	}
	return vpcaccesspb.VpcaccessConnectorStateEnum(0)
}

// ConnectorSubnetToProto converts a ConnectorSubnet resource to its proto representation.
func VpcaccessConnectorSubnetToProto(o *vpcaccess.ConnectorSubnet) *vpcaccesspb.VpcaccessConnectorSubnet {
	if o == nil {
		return nil
	}
	p := &vpcaccesspb.VpcaccessConnectorSubnet{
		Name:      dcl.ValueOrEmptyString(o.Name),
		ProjectId: dcl.ValueOrEmptyString(o.ProjectId),
	}
	return p
}

// ConnectorToProto converts a Connector resource to its proto representation.
func ConnectorToProto(resource *vpcaccess.Connector) *vpcaccesspb.VpcaccessConnector {
	p := &vpcaccesspb.VpcaccessConnector{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Network:       dcl.ValueOrEmptyString(resource.Network),
		IpCidrRange:   dcl.ValueOrEmptyString(resource.IPCidrRange),
		State:         VpcaccessConnectorStateEnumToProto(resource.State),
		MinThroughput: dcl.ValueOrEmptyInt64(resource.MinThroughput),
		MaxThroughput: dcl.ValueOrEmptyInt64(resource.MaxThroughput),
		Subnet:        VpcaccessConnectorSubnetToProto(resource.Subnet),
		MachineType:   dcl.ValueOrEmptyString(resource.MachineType),
		MinInstances:  dcl.ValueOrEmptyInt64(resource.MinInstances),
		MaxInstances:  dcl.ValueOrEmptyInt64(resource.MaxInstances),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ConnectedProjects {
		p.ConnectedProjects = append(p.ConnectedProjects, r)
	}

	return p
}

// ApplyConnector handles the gRPC request by passing it to the underlying Connector Apply() method.
func (s *ConnectorServer) applyConnector(ctx context.Context, c *vpcaccess.Client, request *vpcaccesspb.ApplyVpcaccessConnectorRequest) (*vpcaccesspb.VpcaccessConnector, error) {
	p := ProtoToConnector(request.GetResource())
	res, err := c.ApplyConnector(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectorToProto(res)
	return r, nil
}

// ApplyConnector handles the gRPC request by passing it to the underlying Connector Apply() method.
func (s *ConnectorServer) ApplyVpcaccessConnector(ctx context.Context, request *vpcaccesspb.ApplyVpcaccessConnectorRequest) (*vpcaccesspb.VpcaccessConnector, error) {
	cl, err := createConfigConnector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyConnector(ctx, cl, request)
}

// DeleteConnector handles the gRPC request by passing it to the underlying Connector Delete() method.
func (s *ConnectorServer) DeleteVpcaccessConnector(ctx context.Context, request *vpcaccesspb.DeleteVpcaccessConnectorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnector(ctx, ProtoToConnector(request.GetResource()))

}

// ListConnector handles the gRPC request by passing it to the underlying ConnectorList() method.
func (s *ConnectorServer) ListVpcaccessConnector(ctx context.Context, request *vpcaccesspb.ListVpcaccessConnectorRequest) (*vpcaccesspb.ListVpcaccessConnectorResponse, error) {
	cl, err := createConfigConnector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnector(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*vpcaccesspb.VpcaccessConnector
	for _, r := range resources.Items {
		rp := ConnectorToProto(r)
		protos = append(protos, rp)
	}
	return &vpcaccesspb.ListVpcaccessConnectorResponse{Items: protos}, nil
}

func createConfigConnector(ctx context.Context, service_account_file string) (*vpcaccess.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return vpcaccess.NewClient(conf), nil
}
