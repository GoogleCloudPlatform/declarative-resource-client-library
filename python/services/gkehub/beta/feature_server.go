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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// Server implements the gRPC interface for Feature.
type FeatureServer struct{}

// ProtoToFeatureSpec converts a FeatureSpec resource from its proto representation.
func ProtoToGkehubBetaFeatureSpec(p *betapb.GkehubBetaFeatureSpec) *beta.FeatureSpec {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpec{
		Multiclusteringress: ProtoToGkehubBetaFeatureSpecMulticlusteringress(p.GetMulticlusteringress()),
	}
	return obj
}

// ProtoToFeatureSpecMulticlusteringress converts a FeatureSpecMulticlusteringress resource from its proto representation.
func ProtoToGkehubBetaFeatureSpecMulticlusteringress(p *betapb.GkehubBetaFeatureSpecMulticlusteringress) *beta.FeatureSpecMulticlusteringress {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.StringOrNil(p.ConfigMembership),
	}
	return obj
}

// ProtoToFeature converts a Feature resource from its proto representation.
func ProtoToFeature(p *betapb.GkehubBetaFeature) *beta.Feature {
	obj := &beta.Feature{
		Name:       dcl.StringOrNil(p.Name),
		Spec:       ProtoToGkehubBetaFeatureSpec(p.GetSpec()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime: dcl.StringOrNil(p.GetDeleteTime()),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	return obj
}

// FeatureSpecToProto converts a FeatureSpec resource to its proto representation.
func GkehubBetaFeatureSpecToProto(o *beta.FeatureSpec) *betapb.GkehubBetaFeatureSpec {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpec{
		Multiclusteringress: GkehubBetaFeatureSpecMulticlusteringressToProto(o.Multiclusteringress),
	}
	return p
}

// FeatureSpecMulticlusteringressToProto converts a FeatureSpecMulticlusteringress resource to its proto representation.
func GkehubBetaFeatureSpecMulticlusteringressToProto(o *beta.FeatureSpecMulticlusteringress) *betapb.GkehubBetaFeatureSpecMulticlusteringress {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecMulticlusteringress{
		ConfigMembership: dcl.ValueOrEmptyString(o.ConfigMembership),
	}
	return p
}

// FeatureToProto converts a Feature resource to its proto representation.
func FeatureToProto(resource *beta.Feature) *betapb.GkehubBetaFeature {
	p := &betapb.GkehubBetaFeature{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Spec:       GkehubBetaFeatureSpecToProto(resource.Spec),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime: dcl.ValueOrEmptyString(resource.DeleteTime),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) applyFeature(ctx context.Context, c *beta.Client, request *betapb.ApplyGkehubBetaFeatureRequest) (*betapb.GkehubBetaFeature, error) {
	p := ProtoToFeature(request.GetResource())
	res, err := c.ApplyFeature(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureToProto(res)
	return r, nil
}

// ApplyFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) ApplyGkehubBetaFeature(ctx context.Context, request *betapb.ApplyGkehubBetaFeatureRequest) (*betapb.GkehubBetaFeature, error) {
	cl, err := createConfigFeature(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFeature(ctx, cl, request)
}

// DeleteFeature handles the gRPC request by passing it to the underlying Feature Delete() method.
func (s *FeatureServer) DeleteGkehubBetaFeature(ctx context.Context, request *betapb.DeleteGkehubBetaFeatureRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeature(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeature(ctx, ProtoToFeature(request.GetResource()))

}

// ListFeature handles the gRPC request by passing it to the underlying FeatureList() method.
func (s *FeatureServer) ListGkehubBetaFeature(ctx context.Context, request *betapb.ListGkehubBetaFeatureRequest) (*betapb.ListGkehubBetaFeatureResponse, error) {
	cl, err := createConfigFeature(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeature(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkehubBetaFeature
	for _, r := range resources.Items {
		rp := FeatureToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkehubBetaFeatureResponse{Items: protos}, nil
}

func createConfigFeature(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
