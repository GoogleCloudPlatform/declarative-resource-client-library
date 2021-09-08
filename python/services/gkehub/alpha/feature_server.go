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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// Server implements the gRPC interface for Feature.
type FeatureServer struct{}

// ProtoToFeatureResourceStateStateEnum converts a FeatureResourceStateStateEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureResourceStateStateEnum(e alphapb.GkehubAlphaFeatureResourceStateStateEnum) *alpha.FeatureResourceStateStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureResourceStateStateEnum_name[int32(e)]; ok {
		e := alpha.FeatureResourceStateStateEnum(n[len("GkehubAlphaFeatureResourceStateStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateStateCodeEnum converts a FeatureStateStateCodeEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureStateStateCodeEnum(e alphapb.GkehubAlphaFeatureStateStateCodeEnum) *alpha.FeatureStateStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureStateStateCodeEnum_name[int32(e)]; ok {
		e := alpha.FeatureStateStateCodeEnum(n[len("GkehubAlphaFeatureStateStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum converts a FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(e alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum) *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum_name[int32(e)]; ok {
		e := alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(n[len("GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureResourceState converts a FeatureResourceState resource from its proto representation.
func ProtoToGkehubAlphaFeatureResourceState(p *alphapb.GkehubAlphaFeatureResourceState) *alpha.FeatureResourceState {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureResourceState{
		State:        ProtoToGkehubAlphaFeatureResourceStateStateEnum(p.GetState()),
		HasResources: dcl.Bool(p.HasResources),
	}
	return obj
}

// ProtoToFeatureSpec converts a FeatureSpec resource from its proto representation.
func ProtoToGkehubAlphaFeatureSpec(p *alphapb.GkehubAlphaFeatureSpec) *alpha.FeatureSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpec{
		Multiclusteringress: ProtoToGkehubAlphaFeatureSpecMulticlusteringress(p.GetMulticlusteringress()),
		Cloudauditlogging:   ProtoToGkehubAlphaFeatureSpecCloudauditlogging(p.GetCloudauditlogging()),
	}
	return obj
}

// ProtoToFeatureSpecMulticlusteringress converts a FeatureSpecMulticlusteringress resource from its proto representation.
func ProtoToGkehubAlphaFeatureSpecMulticlusteringress(p *alphapb.GkehubAlphaFeatureSpecMulticlusteringress) *alpha.FeatureSpecMulticlusteringress {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.StringOrNil(p.ConfigMembership),
	}
	return obj
}

// ProtoToFeatureSpecCloudauditlogging converts a FeatureSpecCloudauditlogging resource from its proto representation.
func ProtoToGkehubAlphaFeatureSpecCloudauditlogging(p *alphapb.GkehubAlphaFeatureSpecCloudauditlogging) *alpha.FeatureSpecCloudauditlogging {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecCloudauditlogging{}
	for _, r := range p.GetAllowlistedServiceAccounts() {
		obj.AllowlistedServiceAccounts = append(obj.AllowlistedServiceAccounts, r)
	}
	return obj
}

// ProtoToFeatureState converts a FeatureState resource from its proto representation.
func ProtoToGkehubAlphaFeatureState(p *alphapb.GkehubAlphaFeatureState) *alpha.FeatureState {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureState{
		State:       ProtoToGkehubAlphaFeatureStateState(p.GetState()),
		Servicemesh: ProtoToGkehubAlphaFeatureStateServicemesh(p.GetServicemesh()),
	}
	return obj
}

// ProtoToFeatureStateState converts a FeatureStateState resource from its proto representation.
func ProtoToGkehubAlphaFeatureStateState(p *alphapb.GkehubAlphaFeatureStateState) *alpha.FeatureStateState {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateState{
		Code:        ProtoToGkehubAlphaFeatureStateStateCodeEnum(p.GetCode()),
		Description: dcl.StringOrNil(p.Description),
		UpdateTime:  dcl.StringOrNil(p.UpdateTime),
	}
	return obj
}

// ProtoToFeatureStateServicemesh converts a FeatureStateServicemesh resource from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemesh(p *alphapb.GkehubAlphaFeatureStateServicemesh) *alpha.FeatureStateServicemesh {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemesh{}
	for _, r := range p.GetAnalysisMessages() {
		obj.AnalysisMessages = append(obj.AnalysisMessages, *ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessages(r))
	}
	return obj
}

// ProtoToFeatureStateServicemeshAnalysisMessages converts a FeatureStateServicemeshAnalysisMessages resource from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessages(p *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages) *alpha.FeatureStateServicemeshAnalysisMessages {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemeshAnalysisMessages{
		MessageBase: ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase(p.GetMessageBase()),
		Description: dcl.StringOrNil(p.Description),
	}
	for _, r := range p.GetResourcePaths() {
		obj.ResourcePaths = append(obj.ResourcePaths, r)
	}
	return obj
}

// ProtoToFeatureStateServicemeshAnalysisMessagesMessageBase converts a FeatureStateServicemeshAnalysisMessagesMessageBase resource from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase(p *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase) *alpha.FeatureStateServicemeshAnalysisMessagesMessageBase {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemeshAnalysisMessagesMessageBase{
		Type:             ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType(p.GetType()),
		Level:            ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(p.GetLevel()),
		DocumentationUrl: dcl.StringOrNil(p.DocumentationUrl),
	}
	return obj
}

// ProtoToFeatureStateServicemeshAnalysisMessagesMessageBaseType converts a FeatureStateServicemeshAnalysisMessagesMessageBaseType resource from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType(p *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType) *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseType{
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Code:        dcl.StringOrNil(p.Code),
	}
	return obj
}

// ProtoToFeature converts a Feature resource from its proto representation.
func ProtoToFeature(p *alphapb.GkehubAlphaFeature) *alpha.Feature {
	obj := &alpha.Feature{
		Name:          dcl.StringOrNil(p.Name),
		ResourceState: ProtoToGkehubAlphaFeatureResourceState(p.GetResourceState()),
		Spec:          ProtoToGkehubAlphaFeatureSpec(p.GetSpec()),
		State:         ProtoToGkehubAlphaFeatureState(p.GetState()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	return obj
}

// FeatureResourceStateStateEnumToProto converts a FeatureResourceStateStateEnum enum to its proto representation.
func GkehubAlphaFeatureResourceStateStateEnumToProto(e *alpha.FeatureResourceStateStateEnum) alphapb.GkehubAlphaFeatureResourceStateStateEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureResourceStateStateEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureResourceStateStateEnum_value["FeatureResourceStateStateEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureResourceStateStateEnum(v)
	}
	return alphapb.GkehubAlphaFeatureResourceStateStateEnum(0)
}

// FeatureStateStateCodeEnumToProto converts a FeatureStateStateCodeEnum enum to its proto representation.
func GkehubAlphaFeatureStateStateCodeEnumToProto(e *alpha.FeatureStateStateCodeEnum) alphapb.GkehubAlphaFeatureStateStateCodeEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureStateStateCodeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureStateStateCodeEnum_value["FeatureStateStateCodeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureStateStateCodeEnum(v)
	}
	return alphapb.GkehubAlphaFeatureStateStateCodeEnum(0)
}

// FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumToProto converts a FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum enum to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumToProto(e *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum) alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum_value["FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(v)
	}
	return alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(0)
}

// FeatureResourceStateToProto converts a FeatureResourceState resource to its proto representation.
func GkehubAlphaFeatureResourceStateToProto(o *alpha.FeatureResourceState) *alphapb.GkehubAlphaFeatureResourceState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureResourceState{
		State:        GkehubAlphaFeatureResourceStateStateEnumToProto(o.State),
		HasResources: dcl.ValueOrEmptyBool(o.HasResources),
	}
	return p
}

// FeatureSpecToProto converts a FeatureSpec resource to its proto representation.
func GkehubAlphaFeatureSpecToProto(o *alpha.FeatureSpec) *alphapb.GkehubAlphaFeatureSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpec{
		Multiclusteringress: GkehubAlphaFeatureSpecMulticlusteringressToProto(o.Multiclusteringress),
		Cloudauditlogging:   GkehubAlphaFeatureSpecCloudauditloggingToProto(o.Cloudauditlogging),
	}
	return p
}

// FeatureSpecMulticlusteringressToProto converts a FeatureSpecMulticlusteringress resource to its proto representation.
func GkehubAlphaFeatureSpecMulticlusteringressToProto(o *alpha.FeatureSpecMulticlusteringress) *alphapb.GkehubAlphaFeatureSpecMulticlusteringress {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecMulticlusteringress{
		ConfigMembership: dcl.ValueOrEmptyString(o.ConfigMembership),
	}
	return p
}

// FeatureSpecCloudauditloggingToProto converts a FeatureSpecCloudauditlogging resource to its proto representation.
func GkehubAlphaFeatureSpecCloudauditloggingToProto(o *alpha.FeatureSpecCloudauditlogging) *alphapb.GkehubAlphaFeatureSpecCloudauditlogging {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecCloudauditlogging{}
	for _, r := range o.AllowlistedServiceAccounts {
		p.AllowlistedServiceAccounts = append(p.AllowlistedServiceAccounts, r)
	}
	return p
}

// FeatureStateToProto converts a FeatureState resource to its proto representation.
func GkehubAlphaFeatureStateToProto(o *alpha.FeatureState) *alphapb.GkehubAlphaFeatureState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureState{
		State:       GkehubAlphaFeatureStateStateToProto(o.State),
		Servicemesh: GkehubAlphaFeatureStateServicemeshToProto(o.Servicemesh),
	}
	return p
}

// FeatureStateStateToProto converts a FeatureStateState resource to its proto representation.
func GkehubAlphaFeatureStateStateToProto(o *alpha.FeatureStateState) *alphapb.GkehubAlphaFeatureStateState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateState{
		Code:        GkehubAlphaFeatureStateStateCodeEnumToProto(o.Code),
		Description: dcl.ValueOrEmptyString(o.Description),
		UpdateTime:  dcl.ValueOrEmptyString(o.UpdateTime),
	}
	return p
}

// FeatureStateServicemeshToProto converts a FeatureStateServicemesh resource to its proto representation.
func GkehubAlphaFeatureStateServicemeshToProto(o *alpha.FeatureStateServicemesh) *alphapb.GkehubAlphaFeatureStateServicemesh {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemesh{}
	for _, r := range o.AnalysisMessages {
		p.AnalysisMessages = append(p.AnalysisMessages, GkehubAlphaFeatureStateServicemeshAnalysisMessagesToProto(&r))
	}
	return p
}

// FeatureStateServicemeshAnalysisMessagesToProto converts a FeatureStateServicemeshAnalysisMessages resource to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesToProto(o *alpha.FeatureStateServicemeshAnalysisMessages) *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages{
		MessageBase: GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseToProto(o.MessageBase),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	for _, r := range o.ResourcePaths {
		p.ResourcePaths = append(p.ResourcePaths, r)
	}
	p.Args = make(map[string]string)
	for k, r := range o.Args {
		p.Args[k] = r
	}
	return p
}

// FeatureStateServicemeshAnalysisMessagesMessageBaseToProto converts a FeatureStateServicemeshAnalysisMessagesMessageBase resource to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseToProto(o *alpha.FeatureStateServicemeshAnalysisMessagesMessageBase) *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase{
		Type:             GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseTypeToProto(o.Type),
		Level:            GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumToProto(o.Level),
		DocumentationUrl: dcl.ValueOrEmptyString(o.DocumentationUrl),
	}
	return p
}

// FeatureStateServicemeshAnalysisMessagesMessageBaseTypeToProto converts a FeatureStateServicemeshAnalysisMessagesMessageBaseType resource to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseTypeToProto(o *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseType) *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType{
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		Code:        dcl.ValueOrEmptyString(o.Code),
	}
	return p
}

// FeatureToProto converts a Feature resource to its proto representation.
func FeatureToProto(resource *alpha.Feature) *alphapb.GkehubAlphaFeature {
	p := &alphapb.GkehubAlphaFeature{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		ResourceState: GkehubAlphaFeatureResourceStateToProto(resource.ResourceState),
		Spec:          GkehubAlphaFeatureSpecToProto(resource.Spec),
		State:         GkehubAlphaFeatureStateToProto(resource.State),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:    dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:    dcl.ValueOrEmptyString(resource.DeleteTime),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) applyFeature(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaFeatureRequest) (*alphapb.GkehubAlphaFeature, error) {
	p := ProtoToFeature(request.GetResource())
	res, err := c.ApplyFeature(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureToProto(res)
	return r, nil
}

// ApplyFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) ApplyGkehubAlphaFeature(ctx context.Context, request *alphapb.ApplyGkehubAlphaFeatureRequest) (*alphapb.GkehubAlphaFeature, error) {
	cl, err := createConfigFeature(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFeature(ctx, cl, request)
}

// DeleteFeature handles the gRPC request by passing it to the underlying Feature Delete() method.
func (s *FeatureServer) DeleteGkehubAlphaFeature(ctx context.Context, request *alphapb.DeleteGkehubAlphaFeatureRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeature(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeature(ctx, ProtoToFeature(request.GetResource()))

}

// ListGkehubAlphaFeature handles the gRPC request by passing it to the underlying FeatureList() method.
func (s *FeatureServer) ListGkehubAlphaFeature(ctx context.Context, request *alphapb.ListGkehubAlphaFeatureRequest) (*alphapb.ListGkehubAlphaFeatureResponse, error) {
	cl, err := createConfigFeature(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeature(ctx, ProtoToFeature(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkehubAlphaFeature
	for _, r := range resources.Items {
		rp := FeatureToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkehubAlphaFeatureResponse{Items: protos}, nil
}

func createConfigFeature(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
