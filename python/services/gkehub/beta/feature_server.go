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

// ProtoToFeatureResourceStateStateEnum converts a FeatureResourceStateStateEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureResourceStateStateEnum(e betapb.GkehubBetaFeatureResourceStateStateEnum) *beta.FeatureResourceStateStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureResourceStateStateEnum_name[int32(e)]; ok {
		e := beta.FeatureResourceStateStateEnum(n[len("GkehubBetaFeatureResourceStateStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateStateCodeEnum converts a FeatureStateStateCodeEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureStateStateCodeEnum(e betapb.GkehubBetaFeatureStateStateCodeEnum) *beta.FeatureStateStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureStateStateCodeEnum_name[int32(e)]; ok {
		e := beta.FeatureStateStateCodeEnum(n[len("GkehubBetaFeatureStateStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureResourceState converts a FeatureResourceState resource from its proto representation.
func ProtoToGkehubBetaFeatureResourceState(p *betapb.GkehubBetaFeatureResourceState) *beta.FeatureResourceState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureResourceState{
		State:        ProtoToGkehubBetaFeatureResourceStateStateEnum(p.GetState()),
		HasResources: dcl.Bool(p.HasResources),
	}
	return obj
}

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

// ProtoToFeatureState converts a FeatureState resource from its proto representation.
func ProtoToGkehubBetaFeatureState(p *betapb.GkehubBetaFeatureState) *beta.FeatureState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureState{
		State: ProtoToGkehubBetaFeatureStateState(p.GetState()),
	}
	return obj
}

// ProtoToFeatureStateState converts a FeatureStateState resource from its proto representation.
func ProtoToGkehubBetaFeatureStateState(p *betapb.GkehubBetaFeatureStateState) *beta.FeatureStateState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureStateState{
		Code:        ProtoToGkehubBetaFeatureStateStateCodeEnum(p.GetCode()),
		Description: dcl.StringOrNil(p.Description),
		UpdateTime:  dcl.StringOrNil(p.UpdateTime),
	}
	return obj
}

// ProtoToFeature converts a Feature resource from its proto representation.
func ProtoToFeature(p *betapb.GkehubBetaFeature) *beta.Feature {
	obj := &beta.Feature{
		Name:          dcl.StringOrNil(p.Name),
		ResourceState: ProtoToGkehubBetaFeatureResourceState(p.GetResourceState()),
		Spec:          ProtoToGkehubBetaFeatureSpec(p.GetSpec()),
		State:         ProtoToGkehubBetaFeatureState(p.GetState()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	return obj
}

// FeatureResourceStateStateEnumToProto converts a FeatureResourceStateStateEnum enum to its proto representation.
func GkehubBetaFeatureResourceStateStateEnumToProto(e *beta.FeatureResourceStateStateEnum) betapb.GkehubBetaFeatureResourceStateStateEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureResourceStateStateEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureResourceStateStateEnum_value["FeatureResourceStateStateEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureResourceStateStateEnum(v)
	}
	return betapb.GkehubBetaFeatureResourceStateStateEnum(0)
}

// FeatureStateStateCodeEnumToProto converts a FeatureStateStateCodeEnum enum to its proto representation.
func GkehubBetaFeatureStateStateCodeEnumToProto(e *beta.FeatureStateStateCodeEnum) betapb.GkehubBetaFeatureStateStateCodeEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureStateStateCodeEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureStateStateCodeEnum_value["FeatureStateStateCodeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureStateStateCodeEnum(v)
	}
	return betapb.GkehubBetaFeatureStateStateCodeEnum(0)
}

// FeatureResourceStateToProto converts a FeatureResourceState resource to its proto representation.
func GkehubBetaFeatureResourceStateToProto(o *beta.FeatureResourceState) *betapb.GkehubBetaFeatureResourceState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureResourceState{
		State:        GkehubBetaFeatureResourceStateStateEnumToProto(o.State),
		HasResources: dcl.ValueOrEmptyBool(o.HasResources),
	}
	return p
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

// FeatureStateToProto converts a FeatureState resource to its proto representation.
func GkehubBetaFeatureStateToProto(o *beta.FeatureState) *betapb.GkehubBetaFeatureState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureState{
		State: GkehubBetaFeatureStateStateToProto(o.State),
	}
	return p
}

// FeatureStateStateToProto converts a FeatureStateState resource to its proto representation.
func GkehubBetaFeatureStateStateToProto(o *beta.FeatureStateState) *betapb.GkehubBetaFeatureStateState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureStateState{
		Code:        GkehubBetaFeatureStateStateCodeEnumToProto(o.Code),
		Description: dcl.ValueOrEmptyString(o.Description),
		UpdateTime:  dcl.ValueOrEmptyString(o.UpdateTime),
	}
	return p
}

// FeatureToProto converts a Feature resource to its proto representation.
func FeatureToProto(resource *beta.Feature) *betapb.GkehubBetaFeature {
	p := &betapb.GkehubBetaFeature{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		ResourceState: GkehubBetaFeatureResourceStateToProto(resource.ResourceState),
		Spec:          GkehubBetaFeatureSpecToProto(resource.Spec),
		State:         GkehubBetaFeatureStateToProto(resource.State),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:    dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:    dcl.ValueOrEmptyString(resource.DeleteTime),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
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
