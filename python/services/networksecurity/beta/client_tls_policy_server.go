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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/beta/networksecurity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
)

// Server implements the gRPC interface for ClientTlsPolicy.
type ClientTlsPolicyServer struct{}

// ProtoToClientTlsPolicyClientCertificate converts a ClientTlsPolicyClientCertificate resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificate(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificate) *beta.ClientTlsPolicyClientCertificate {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificate{
		LocalFilepath:               ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepath(p.GetLocalFilepath()),
		GrpcEndpoint:                ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateLocalFilepath converts a ClientTlsPolicyClientCertificateLocalFilepath resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepath(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepath) *beta.ClientTlsPolicyClientCertificateLocalFilepath {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificateLocalFilepath{
		CertificatePath: dcl.StringOrNil(p.CertificatePath),
		PrivateKeyPath:  dcl.StringOrNil(p.PrivateKeyPath),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateGrpcEndpoint converts a ClientTlsPolicyClientCertificateGrpcEndpoint resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint) *beta.ClientTlsPolicyClientCertificateGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificateGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.TargetUri),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateCertificateProviderInstance converts a ClientTlsPolicyClientCertificateCertificateProviderInstance resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance) *beta.ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificateCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.PluginInstance),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCa converts a ClientTlsPolicyServerValidationCa resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCa(p *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa) *beta.ClientTlsPolicyServerValidationCa {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyServerValidationCa{
		CaCertPath:                  dcl.StringOrNil(p.CaCertPath),
		GrpcEndpoint:                ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCaGrpcEndpoint converts a ClientTlsPolicyServerValidationCaGrpcEndpoint resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint(p *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint) *beta.ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyServerValidationCaGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.TargetUri),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCaCertificateProviderInstance converts a ClientTlsPolicyServerValidationCaCertificateProviderInstance resource from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance(p *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance) *beta.ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyServerValidationCaCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.PluginInstance),
	}
	return obj
}

// ProtoToClientTlsPolicy converts a ClientTlsPolicy resource from its proto representation.
func ProtoToClientTlsPolicy(p *betapb.NetworksecurityBetaClientTlsPolicy) *beta.ClientTlsPolicy {
	obj := &beta.ClientTlsPolicy{
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Sni:               dcl.StringOrNil(p.Sni),
		ClientCertificate: ProtoToNetworksecurityBetaClientTlsPolicyClientCertificate(p.GetClientCertificate()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetServerValidationCa() {
		obj.ServerValidationCa = append(obj.ServerValidationCa, *ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCa(r))
	}
	return obj
}

// ClientTlsPolicyClientCertificateToProto converts a ClientTlsPolicyClientCertificate resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateToProto(o *beta.ClientTlsPolicyClientCertificate) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificate {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificate{
		LocalFilepath:               NetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepathToProto(o.LocalFilepath),
		GrpcEndpoint:                NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpointToProto(o.GrpcEndpoint),
		CertificateProviderInstance: NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstanceToProto(o.CertificateProviderInstance),
	}
	return p
}

// ClientTlsPolicyClientCertificateLocalFilepathToProto converts a ClientTlsPolicyClientCertificateLocalFilepath resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepathToProto(o *beta.ClientTlsPolicyClientCertificateLocalFilepath) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepath {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificateLocalFilepath{
		CertificatePath: dcl.ValueOrEmptyString(o.CertificatePath),
		PrivateKeyPath:  dcl.ValueOrEmptyString(o.PrivateKeyPath),
	}
	return p
}

// ClientTlsPolicyClientCertificateGrpcEndpointToProto converts a ClientTlsPolicyClientCertificateGrpcEndpoint resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpointToProto(o *beta.ClientTlsPolicyClientCertificateGrpcEndpoint) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint{
		TargetUri: dcl.ValueOrEmptyString(o.TargetUri),
	}
	return p
}

// ClientTlsPolicyClientCertificateCertificateProviderInstanceToProto converts a ClientTlsPolicyClientCertificateCertificateProviderInstance resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstanceToProto(o *beta.ClientTlsPolicyClientCertificateCertificateProviderInstance) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance{
		PluginInstance: dcl.ValueOrEmptyString(o.PluginInstance),
	}
	return p
}

// ClientTlsPolicyServerValidationCaToProto converts a ClientTlsPolicyServerValidationCa resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyServerValidationCaToProto(o *beta.ClientTlsPolicyServerValidationCa) *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa{
		CaCertPath:                  dcl.ValueOrEmptyString(o.CaCertPath),
		GrpcEndpoint:                NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpointToProto(o.GrpcEndpoint),
		CertificateProviderInstance: NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto(o.CertificateProviderInstance),
	}
	return p
}

// ClientTlsPolicyServerValidationCaGrpcEndpointToProto converts a ClientTlsPolicyServerValidationCaGrpcEndpoint resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpointToProto(o *beta.ClientTlsPolicyServerValidationCaGrpcEndpoint) *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint{
		TargetUri: dcl.ValueOrEmptyString(o.TargetUri),
	}
	return p
}

// ClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto converts a ClientTlsPolicyServerValidationCaCertificateProviderInstance resource to its proto representation.
func NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto(o *beta.ClientTlsPolicyServerValidationCaCertificateProviderInstance) *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance{
		PluginInstance: dcl.ValueOrEmptyString(o.PluginInstance),
	}
	return p
}

// ClientTlsPolicyToProto converts a ClientTlsPolicy resource to its proto representation.
func ClientTlsPolicyToProto(resource *beta.ClientTlsPolicy) *betapb.NetworksecurityBetaClientTlsPolicy {
	p := &betapb.NetworksecurityBetaClientTlsPolicy{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		Sni:               dcl.ValueOrEmptyString(resource.Sni),
		ClientCertificate: NetworksecurityBetaClientTlsPolicyClientCertificateToProto(resource.ClientCertificate),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ServerValidationCa {
		p.ServerValidationCa = append(p.ServerValidationCa, NetworksecurityBetaClientTlsPolicyServerValidationCaToProto(&r))
	}

	return p
}

// ApplyClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Apply() method.
func (s *ClientTlsPolicyServer) applyClientTlsPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworksecurityBetaClientTlsPolicyRequest) (*betapb.NetworksecurityBetaClientTlsPolicy, error) {
	p := ProtoToClientTlsPolicy(request.GetResource())
	res, err := c.ApplyClientTlsPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClientTlsPolicyToProto(res)
	return r, nil
}

// ApplyClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Apply() method.
func (s *ClientTlsPolicyServer) ApplyNetworksecurityBetaClientTlsPolicy(ctx context.Context, request *betapb.ApplyNetworksecurityBetaClientTlsPolicyRequest) (*betapb.NetworksecurityBetaClientTlsPolicy, error) {
	cl, err := createConfigClientTlsPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyClientTlsPolicy(ctx, cl, request)
}

// DeleteClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Delete() method.
func (s *ClientTlsPolicyServer) DeleteNetworksecurityBetaClientTlsPolicy(ctx context.Context, request *betapb.DeleteNetworksecurityBetaClientTlsPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigClientTlsPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteClientTlsPolicy(ctx, ProtoToClientTlsPolicy(request.GetResource()))

}

// ListClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicyList() method.
func (s *ClientTlsPolicyServer) ListNetworksecurityBetaClientTlsPolicy(ctx context.Context, request *betapb.ListNetworksecurityBetaClientTlsPolicyRequest) (*betapb.ListNetworksecurityBetaClientTlsPolicyResponse, error) {
	cl, err := createConfigClientTlsPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListClientTlsPolicy(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworksecurityBetaClientTlsPolicy
	for _, r := range resources.Items {
		rp := ClientTlsPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListNetworksecurityBetaClientTlsPolicyResponse{Items: protos}, nil
}

func createConfigClientTlsPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
