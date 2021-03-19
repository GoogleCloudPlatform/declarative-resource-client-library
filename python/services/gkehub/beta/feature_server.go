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
		e := beta.FeatureResourceStateStateEnum(n[len("FeatureResourceStateStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureSpecMulticlusteringressBillingEnum converts a FeatureSpecMulticlusteringressBillingEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureSpecMulticlusteringressBillingEnum(e betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum) *beta.FeatureSpecMulticlusteringressBillingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum_name[int32(e)]; ok {
		e := beta.FeatureSpecMulticlusteringressBillingEnum(n[len("FeatureSpecMulticlusteringressBillingEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureSpecHelloworldFeatureTestThirdEnum converts a FeatureSpecHelloworldFeatureTestThirdEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum(e betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum) *beta.FeatureSpecHelloworldFeatureTestThirdEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum_name[int32(e)]; ok {
		e := beta.FeatureSpecHelloworldFeatureTestThirdEnum(n[len("FeatureSpecHelloworldFeatureTestThirdEnum"):])
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
		e := beta.FeatureStateStateCodeEnum(n[len("FeatureStateStateCodeEnum"):])
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
		State: ProtoToGkehubBetaFeatureResourceStateStateEnum(p.GetState()),
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
		Helloworld:          ProtoToGkehubBetaFeatureSpecHelloworld(p.GetHelloworld()),
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
		Billing:          ProtoToGkehubBetaFeatureSpecMulticlusteringressBillingEnum(p.GetBilling()),
	}
	return obj
}

// ProtoToFeatureSpecHelloworld converts a FeatureSpecHelloworld resource from its proto representation.
func ProtoToGkehubBetaFeatureSpecHelloworld(p *betapb.GkehubBetaFeatureSpecHelloworld) *beta.FeatureSpecHelloworld {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecHelloworld{
		FeatureTest:  ProtoToGkehubBetaFeatureSpecHelloworldFeatureTest(p.GetFeatureTest()),
		CustomConfig: dcl.StringOrNil(p.CustomConfig),
	}
	return obj
}

// ProtoToFeatureSpecHelloworldFeatureTest converts a FeatureSpecHelloworldFeatureTest resource from its proto representation.
func ProtoToGkehubBetaFeatureSpecHelloworldFeatureTest(p *betapb.GkehubBetaFeatureSpecHelloworldFeatureTest) *beta.FeatureSpecHelloworldFeatureTest {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecHelloworldFeatureTest{
		First:   dcl.StringOrNil(p.First),
		Second:  dcl.Int64OrNil(p.Second),
		Third:   ProtoToGkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum(p.GetThird()),
		Fourth:  dcl.StringOrNil(p.Fourth),
		Fifth:   ProtoToGkehubBetaFeatureSpecHelloworldFeatureTestFifth(p.GetFifth()),
		Sixth:   dcl.Int64OrNil(p.Sixth),
		Seventh: dcl.StringOrNil(p.Seventh),
	}
	for _, r := range p.GetEighth() {
		obj.Eighth = append(obj.Eighth, *ProtoToGkehubBetaFeatureSpecHelloworldFeatureTestEighth(r))
	}
	return obj
}

// ProtoToFeatureSpecHelloworldFeatureTestFifth converts a FeatureSpecHelloworldFeatureTestFifth resource from its proto representation.
func ProtoToGkehubBetaFeatureSpecHelloworldFeatureTestFifth(p *betapb.GkehubBetaFeatureSpecHelloworldFeatureTestFifth) *beta.FeatureSpecHelloworldFeatureTestFifth {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecHelloworldFeatureTestFifth{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToFeatureSpecHelloworldFeatureTestEighth converts a FeatureSpecHelloworldFeatureTestEighth resource from its proto representation.
func ProtoToGkehubBetaFeatureSpecHelloworldFeatureTestEighth(p *betapb.GkehubBetaFeatureSpecHelloworldFeatureTestEighth) *beta.FeatureSpecHelloworldFeatureTestEighth {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecHelloworldFeatureTestEighth{
		First:  dcl.StringOrNil(p.First),
		Second: dcl.Int64OrNil(p.Second),
	}
	return obj
}

// ProtoToFeatureState converts a FeatureState resource from its proto representation.
func ProtoToGkehubBetaFeatureState(p *betapb.GkehubBetaFeatureState) *beta.FeatureState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureState{
		State:      ProtoToGkehubBetaFeatureStateState(p.GetState()),
		Helloworld: ProtoToGkehubBetaFeatureStateHelloworld(p.GetHelloworld()),
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
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToFeatureStateHelloworld converts a FeatureStateHelloworld resource from its proto representation.
func ProtoToGkehubBetaFeatureStateHelloworld(p *betapb.GkehubBetaFeatureStateHelloworld) *beta.FeatureStateHelloworld {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureStateHelloworld{}
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

// FeatureSpecMulticlusteringressBillingEnumToProto converts a FeatureSpecMulticlusteringressBillingEnum enum to its proto representation.
func GkehubBetaFeatureSpecMulticlusteringressBillingEnumToProto(e *beta.FeatureSpecMulticlusteringressBillingEnum) betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum_value["FeatureSpecMulticlusteringressBillingEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum(v)
	}
	return betapb.GkehubBetaFeatureSpecMulticlusteringressBillingEnum(0)
}

// FeatureSpecHelloworldFeatureTestThirdEnumToProto converts a FeatureSpecHelloworldFeatureTestThirdEnum enum to its proto representation.
func GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnumToProto(e *beta.FeatureSpecHelloworldFeatureTestThirdEnum) betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum_value["FeatureSpecHelloworldFeatureTestThirdEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum(v)
	}
	return betapb.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum(0)
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
		State: GkehubBetaFeatureResourceStateStateEnumToProto(o.State),
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
		Helloworld:          GkehubBetaFeatureSpecHelloworldToProto(o.Helloworld),
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
		Billing:          GkehubBetaFeatureSpecMulticlusteringressBillingEnumToProto(o.Billing),
	}
	return p
}

// FeatureSpecHelloworldToProto converts a FeatureSpecHelloworld resource to its proto representation.
func GkehubBetaFeatureSpecHelloworldToProto(o *beta.FeatureSpecHelloworld) *betapb.GkehubBetaFeatureSpecHelloworld {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecHelloworld{
		FeatureTest:  GkehubBetaFeatureSpecHelloworldFeatureTestToProto(o.FeatureTest),
		CustomConfig: dcl.ValueOrEmptyString(o.CustomConfig),
	}
	return p
}

// FeatureSpecHelloworldFeatureTestToProto converts a FeatureSpecHelloworldFeatureTest resource to its proto representation.
func GkehubBetaFeatureSpecHelloworldFeatureTestToProto(o *beta.FeatureSpecHelloworldFeatureTest) *betapb.GkehubBetaFeatureSpecHelloworldFeatureTest {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecHelloworldFeatureTest{
		First:   dcl.ValueOrEmptyString(o.First),
		Second:  dcl.ValueOrEmptyInt64(o.Second),
		Third:   GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnumToProto(o.Third),
		Fourth:  dcl.ValueOrEmptyString(o.Fourth),
		Fifth:   GkehubBetaFeatureSpecHelloworldFeatureTestFifthToProto(o.Fifth),
		Sixth:   dcl.ValueOrEmptyInt64(o.Sixth),
		Seventh: dcl.ValueOrEmptyString(o.Seventh),
	}
	for _, r := range o.Eighth {
		p.Eighth = append(p.Eighth, GkehubBetaFeatureSpecHelloworldFeatureTestEighthToProto(&r))
	}
	p.Ninth = make(map[string]string)
	for k, r := range o.Ninth {
		p.Ninth[k] = r
	}
	return p
}

// FeatureSpecHelloworldFeatureTestFifthToProto converts a FeatureSpecHelloworldFeatureTestFifth resource to its proto representation.
func GkehubBetaFeatureSpecHelloworldFeatureTestFifthToProto(o *beta.FeatureSpecHelloworldFeatureTestFifth) *betapb.GkehubBetaFeatureSpecHelloworldFeatureTestFifth {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecHelloworldFeatureTestFifth{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// FeatureSpecHelloworldFeatureTestEighthToProto converts a FeatureSpecHelloworldFeatureTestEighth resource to its proto representation.
func GkehubBetaFeatureSpecHelloworldFeatureTestEighthToProto(o *beta.FeatureSpecHelloworldFeatureTestEighth) *betapb.GkehubBetaFeatureSpecHelloworldFeatureTestEighth {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecHelloworldFeatureTestEighth{
		First:  dcl.ValueOrEmptyString(o.First),
		Second: dcl.ValueOrEmptyInt64(o.Second),
	}
	return p
}

// FeatureStateToProto converts a FeatureState resource to its proto representation.
func GkehubBetaFeatureStateToProto(o *beta.FeatureState) *betapb.GkehubBetaFeatureState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureState{
		State:      GkehubBetaFeatureStateStateToProto(o.State),
		Helloworld: GkehubBetaFeatureStateHelloworldToProto(o.Helloworld),
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

// FeatureStateHelloworldToProto converts a FeatureStateHelloworld resource to its proto representation.
func GkehubBetaFeatureStateHelloworldToProto(o *beta.FeatureStateHelloworld) *betapb.GkehubBetaFeatureStateHelloworld {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureStateHelloworld{}
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
