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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/alpha/networksecurity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
)

// Server implements the gRPC interface for ServerTlsPolicy.
type ServerTlsPolicyServer struct{}

// ProtoToServerTlsPolicyServerCertificate converts a ServerTlsPolicyServerCertificate resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificate(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificate) *alpha.ServerTlsPolicyServerCertificate {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificate{
		LocalFilepath:               ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath(p.GetLocalFilepath()),
		GrpcEndpoint:                ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateLocalFilepath converts a ServerTlsPolicyServerCertificateLocalFilepath resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath) *alpha.ServerTlsPolicyServerCertificateLocalFilepath {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificateLocalFilepath{
		CertificatePath: dcl.StringOrNil(p.CertificatePath),
		PrivateKeyPath:  dcl.StringOrNil(p.PrivateKeyPath),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateGrpcEndpoint converts a ServerTlsPolicyServerCertificateGrpcEndpoint resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint) *alpha.ServerTlsPolicyServerCertificateGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificateGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.TargetUri),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateCertificateProviderInstance converts a ServerTlsPolicyServerCertificateCertificateProviderInstance resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance) *alpha.ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificateCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.PluginInstance),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicy converts a ServerTlsPolicyMtlsPolicy resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicy(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicy) *alpha.ServerTlsPolicyMtlsPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicy{}
	for _, r := range p.GetClientValidationCa() {
		obj.ClientValidationCa = append(obj.ClientValidationCa, *ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa(r))
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCa converts a ServerTlsPolicyMtlsPolicyClientValidationCa resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa) *alpha.ServerTlsPolicyMtlsPolicyClientValidationCa {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicyClientValidationCa{
		CaCertPath:                  dcl.StringOrNil(p.CaCertPath),
		GrpcEndpoint:                ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint converts a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.TargetUri),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance converts a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance resource from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.PluginInstance),
	}
	return obj
}

// ProtoToServerTlsPolicy converts a ServerTlsPolicy resource from its proto representation.
func ProtoToServerTlsPolicy(p *alphapb.NetworksecurityAlphaServerTlsPolicy) *alpha.ServerTlsPolicy {
	obj := &alpha.ServerTlsPolicy{
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		AllowOpen:         dcl.Bool(p.AllowOpen),
		ServerCertificate: ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificate(p.GetServerCertificate()),
		MtlsPolicy:        ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicy(p.GetMtlsPolicy()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	return obj
}

// ServerTlsPolicyServerCertificateToProto converts a ServerTlsPolicyServerCertificate resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateToProto(o *alpha.ServerTlsPolicyServerCertificate) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificate {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificate{
		LocalFilepath:               NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepathToProto(o.LocalFilepath),
		GrpcEndpoint:                NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpointToProto(o.GrpcEndpoint),
		CertificateProviderInstance: NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstanceToProto(o.CertificateProviderInstance),
	}
	return p
}

// ServerTlsPolicyServerCertificateLocalFilepathToProto converts a ServerTlsPolicyServerCertificateLocalFilepath resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepathToProto(o *alpha.ServerTlsPolicyServerCertificateLocalFilepath) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath{
		CertificatePath: dcl.ValueOrEmptyString(o.CertificatePath),
		PrivateKeyPath:  dcl.ValueOrEmptyString(o.PrivateKeyPath),
	}
	return p
}

// ServerTlsPolicyServerCertificateGrpcEndpointToProto converts a ServerTlsPolicyServerCertificateGrpcEndpoint resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpointToProto(o *alpha.ServerTlsPolicyServerCertificateGrpcEndpoint) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint{
		TargetUri: dcl.ValueOrEmptyString(o.TargetUri),
	}
	return p
}

// ServerTlsPolicyServerCertificateCertificateProviderInstanceToProto converts a ServerTlsPolicyServerCertificateCertificateProviderInstance resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstanceToProto(o *alpha.ServerTlsPolicyServerCertificateCertificateProviderInstance) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance{
		PluginInstance: dcl.ValueOrEmptyString(o.PluginInstance),
	}
	return p
}

// ServerTlsPolicyMtlsPolicyToProto converts a ServerTlsPolicyMtlsPolicy resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyToProto(o *alpha.ServerTlsPolicyMtlsPolicy) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicy{}
	for _, r := range o.ClientValidationCa {
		p.ClientValidationCa = append(p.ClientValidationCa, NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaToProto(&r))
	}
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCa resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaToProto(o *alpha.ServerTlsPolicyMtlsPolicyClientValidationCa) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa{
		CaCertPath:                  dcl.ValueOrEmptyString(o.CaCertPath),
		GrpcEndpoint:                NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto(o.GrpcEndpoint),
		CertificateProviderInstance: NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto(o.CertificateProviderInstance),
	}
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto(o *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{
		TargetUri: dcl.ValueOrEmptyString(o.TargetUri),
	}
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance resource to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto(o *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{
		PluginInstance: dcl.ValueOrEmptyString(o.PluginInstance),
	}
	return p
}

// ServerTlsPolicyToProto converts a ServerTlsPolicy resource to its proto representation.
func ServerTlsPolicyToProto(resource *alpha.ServerTlsPolicy) *alphapb.NetworksecurityAlphaServerTlsPolicy {
	p := &alphapb.NetworksecurityAlphaServerTlsPolicy{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		AllowOpen:         dcl.ValueOrEmptyBool(resource.AllowOpen),
		ServerCertificate: NetworksecurityAlphaServerTlsPolicyServerCertificateToProto(resource.ServerCertificate),
		MtlsPolicy:        NetworksecurityAlphaServerTlsPolicyMtlsPolicyToProto(resource.MtlsPolicy),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Apply() method.
func (s *ServerTlsPolicyServer) applyServerTlsPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworksecurityAlphaServerTlsPolicyRequest) (*alphapb.NetworksecurityAlphaServerTlsPolicy, error) {
	p := ProtoToServerTlsPolicy(request.GetResource())
	res, err := c.ApplyServerTlsPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServerTlsPolicyToProto(res)
	return r, nil
}

// ApplyServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Apply() method.
func (s *ServerTlsPolicyServer) ApplyNetworksecurityAlphaServerTlsPolicy(ctx context.Context, request *alphapb.ApplyNetworksecurityAlphaServerTlsPolicyRequest) (*alphapb.NetworksecurityAlphaServerTlsPolicy, error) {
	cl, err := createConfigServerTlsPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServerTlsPolicy(ctx, cl, request)
}

// DeleteServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Delete() method.
func (s *ServerTlsPolicyServer) DeleteNetworksecurityAlphaServerTlsPolicy(ctx context.Context, request *alphapb.DeleteNetworksecurityAlphaServerTlsPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServerTlsPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServerTlsPolicy(ctx, ProtoToServerTlsPolicy(request.GetResource()))

}

// ListNetworksecurityAlphaServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicyList() method.
func (s *ServerTlsPolicyServer) ListNetworksecurityAlphaServerTlsPolicy(ctx context.Context, request *alphapb.ListNetworksecurityAlphaServerTlsPolicyRequest) (*alphapb.ListNetworksecurityAlphaServerTlsPolicyResponse, error) {
	cl, err := createConfigServerTlsPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServerTlsPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworksecurityAlphaServerTlsPolicy
	for _, r := range resources.Items {
		rp := ServerTlsPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListNetworksecurityAlphaServerTlsPolicyResponse{Items: protos}, nil
}

func createConfigServerTlsPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
