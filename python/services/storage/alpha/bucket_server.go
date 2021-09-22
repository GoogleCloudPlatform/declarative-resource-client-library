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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/storage/alpha/storage_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/alpha"
)

// Server implements the gRPC interface for Bucket.
type BucketServer struct{}

// ProtoToBucketLifecycleRuleActionTypeEnum converts a BucketLifecycleRuleActionTypeEnum enum from its proto representation.
func ProtoToStorageAlphaBucketLifecycleRuleActionTypeEnum(e alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum) *alpha.BucketLifecycleRuleActionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum_name[int32(e)]; ok {
		e := alpha.BucketLifecycleRuleActionTypeEnum(n[len("StorageAlphaBucketLifecycleRuleActionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketLifecycleRuleConditionWithStateEnum converts a BucketLifecycleRuleConditionWithStateEnum enum from its proto representation.
func ProtoToStorageAlphaBucketLifecycleRuleConditionWithStateEnum(e alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum) *alpha.BucketLifecycleRuleConditionWithStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum_name[int32(e)]; ok {
		e := alpha.BucketLifecycleRuleConditionWithStateEnum(n[len("StorageAlphaBucketLifecycleRuleConditionWithStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketStorageClassEnum converts a BucketStorageClassEnum enum from its proto representation.
func ProtoToStorageAlphaBucketStorageClassEnum(e alphapb.StorageAlphaBucketStorageClassEnum) *alpha.BucketStorageClassEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.StorageAlphaBucketStorageClassEnum_name[int32(e)]; ok {
		e := alpha.BucketStorageClassEnum(n[len("StorageAlphaBucketStorageClassEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketCors converts a BucketCors resource from its proto representation.
func ProtoToStorageAlphaBucketCors(p *alphapb.StorageAlphaBucketCors) *alpha.BucketCors {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketCors{
		MaxAgeSeconds: dcl.Int64OrNil(p.MaxAgeSeconds),
	}
	for _, r := range p.GetMethod() {
		obj.Method = append(obj.Method, r)
	}
	for _, r := range p.GetOrigin() {
		obj.Origin = append(obj.Origin, r)
	}
	for _, r := range p.GetResponseHeader() {
		obj.ResponseHeader = append(obj.ResponseHeader, r)
	}
	return obj
}

// ProtoToBucketLifecycle converts a BucketLifecycle resource from its proto representation.
func ProtoToStorageAlphaBucketLifecycle(p *alphapb.StorageAlphaBucketLifecycle) *alpha.BucketLifecycle {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketLifecycle{}
	for _, r := range p.GetRule() {
		obj.Rule = append(obj.Rule, *ProtoToStorageAlphaBucketLifecycleRule(r))
	}
	return obj
}

// ProtoToBucketLifecycleRule converts a BucketLifecycleRule resource from its proto representation.
func ProtoToStorageAlphaBucketLifecycleRule(p *alphapb.StorageAlphaBucketLifecycleRule) *alpha.BucketLifecycleRule {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketLifecycleRule{
		Action:    ProtoToStorageAlphaBucketLifecycleRuleAction(p.GetAction()),
		Condition: ProtoToStorageAlphaBucketLifecycleRuleCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToBucketLifecycleRuleAction converts a BucketLifecycleRuleAction resource from its proto representation.
func ProtoToStorageAlphaBucketLifecycleRuleAction(p *alphapb.StorageAlphaBucketLifecycleRuleAction) *alpha.BucketLifecycleRuleAction {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketLifecycleRuleAction{
		StorageClass: dcl.StringOrNil(p.StorageClass),
		Type:         ProtoToStorageAlphaBucketLifecycleRuleActionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToBucketLifecycleRuleCondition converts a BucketLifecycleRuleCondition resource from its proto representation.
func ProtoToStorageAlphaBucketLifecycleRuleCondition(p *alphapb.StorageAlphaBucketLifecycleRuleCondition) *alpha.BucketLifecycleRuleCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketLifecycleRuleCondition{
		Age:              dcl.Int64OrNil(p.Age),
		CreatedBefore:    dcl.StringOrNil(p.GetCreatedBefore()),
		WithState:        ProtoToStorageAlphaBucketLifecycleRuleConditionWithStateEnum(p.GetWithState()),
		NumNewerVersions: dcl.Int64OrNil(p.NumNewerVersions),
	}
	for _, r := range p.GetMatchesStorageClass() {
		obj.MatchesStorageClass = append(obj.MatchesStorageClass, r)
	}
	return obj
}

// ProtoToBucketLogging converts a BucketLogging resource from its proto representation.
func ProtoToStorageAlphaBucketLogging(p *alphapb.StorageAlphaBucketLogging) *alpha.BucketLogging {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketLogging{
		LogBucket:       dcl.StringOrNil(p.LogBucket),
		LogObjectPrefix: dcl.StringOrNil(p.LogObjectPrefix),
	}
	return obj
}

// ProtoToBucketVersioning converts a BucketVersioning resource from its proto representation.
func ProtoToStorageAlphaBucketVersioning(p *alphapb.StorageAlphaBucketVersioning) *alpha.BucketVersioning {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketVersioning{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToBucketWebsite converts a BucketWebsite resource from its proto representation.
func ProtoToStorageAlphaBucketWebsite(p *alphapb.StorageAlphaBucketWebsite) *alpha.BucketWebsite {
	if p == nil {
		return nil
	}
	obj := &alpha.BucketWebsite{
		MainPageSuffix: dcl.StringOrNil(p.MainPageSuffix),
		NotFoundPage:   dcl.StringOrNil(p.NotFoundPage),
	}
	return obj
}

// ProtoToBucket converts a Bucket resource from its proto representation.
func ProtoToBucket(p *alphapb.StorageAlphaBucket) *alpha.Bucket {
	obj := &alpha.Bucket{
		Project:      dcl.StringOrNil(p.Project),
		Location:     dcl.StringOrNil(p.Location),
		Name:         dcl.StringOrNil(p.Name),
		Lifecycle:    ProtoToStorageAlphaBucketLifecycle(p.GetLifecycle()),
		Logging:      ProtoToStorageAlphaBucketLogging(p.GetLogging()),
		StorageClass: ProtoToStorageAlphaBucketStorageClassEnum(p.GetStorageClass()),
		Versioning:   ProtoToStorageAlphaBucketVersioning(p.GetVersioning()),
		Website:      ProtoToStorageAlphaBucketWebsite(p.GetWebsite()),
	}
	for _, r := range p.GetCors() {
		obj.Cors = append(obj.Cors, *ProtoToStorageAlphaBucketCors(r))
	}
	return obj
}

// BucketLifecycleRuleActionTypeEnumToProto converts a BucketLifecycleRuleActionTypeEnum enum to its proto representation.
func StorageAlphaBucketLifecycleRuleActionTypeEnumToProto(e *alpha.BucketLifecycleRuleActionTypeEnum) alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum {
	if e == nil {
		return alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum(0)
	}
	if v, ok := alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum_value["BucketLifecycleRuleActionTypeEnum"+string(*e)]; ok {
		return alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum(v)
	}
	return alphapb.StorageAlphaBucketLifecycleRuleActionTypeEnum(0)
}

// BucketLifecycleRuleConditionWithStateEnumToProto converts a BucketLifecycleRuleConditionWithStateEnum enum to its proto representation.
func StorageAlphaBucketLifecycleRuleConditionWithStateEnumToProto(e *alpha.BucketLifecycleRuleConditionWithStateEnum) alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum {
	if e == nil {
		return alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum(0)
	}
	if v, ok := alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum_value["BucketLifecycleRuleConditionWithStateEnum"+string(*e)]; ok {
		return alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum(v)
	}
	return alphapb.StorageAlphaBucketLifecycleRuleConditionWithStateEnum(0)
}

// BucketStorageClassEnumToProto converts a BucketStorageClassEnum enum to its proto representation.
func StorageAlphaBucketStorageClassEnumToProto(e *alpha.BucketStorageClassEnum) alphapb.StorageAlphaBucketStorageClassEnum {
	if e == nil {
		return alphapb.StorageAlphaBucketStorageClassEnum(0)
	}
	if v, ok := alphapb.StorageAlphaBucketStorageClassEnum_value["BucketStorageClassEnum"+string(*e)]; ok {
		return alphapb.StorageAlphaBucketStorageClassEnum(v)
	}
	return alphapb.StorageAlphaBucketStorageClassEnum(0)
}

// BucketCorsToProto converts a BucketCors resource to its proto representation.
func StorageAlphaBucketCorsToProto(o *alpha.BucketCors) *alphapb.StorageAlphaBucketCors {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketCors{
		MaxAgeSeconds: dcl.ValueOrEmptyInt64(o.MaxAgeSeconds),
	}
	for _, r := range o.Method {
		p.Method = append(p.Method, r)
	}
	for _, r := range o.Origin {
		p.Origin = append(p.Origin, r)
	}
	for _, r := range o.ResponseHeader {
		p.ResponseHeader = append(p.ResponseHeader, r)
	}
	return p
}

// BucketLifecycleToProto converts a BucketLifecycle resource to its proto representation.
func StorageAlphaBucketLifecycleToProto(o *alpha.BucketLifecycle) *alphapb.StorageAlphaBucketLifecycle {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketLifecycle{}
	for _, r := range o.Rule {
		p.Rule = append(p.Rule, StorageAlphaBucketLifecycleRuleToProto(&r))
	}
	return p
}

// BucketLifecycleRuleToProto converts a BucketLifecycleRule resource to its proto representation.
func StorageAlphaBucketLifecycleRuleToProto(o *alpha.BucketLifecycleRule) *alphapb.StorageAlphaBucketLifecycleRule {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketLifecycleRule{
		Action:    StorageAlphaBucketLifecycleRuleActionToProto(o.Action),
		Condition: StorageAlphaBucketLifecycleRuleConditionToProto(o.Condition),
	}
	return p
}

// BucketLifecycleRuleActionToProto converts a BucketLifecycleRuleAction resource to its proto representation.
func StorageAlphaBucketLifecycleRuleActionToProto(o *alpha.BucketLifecycleRuleAction) *alphapb.StorageAlphaBucketLifecycleRuleAction {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketLifecycleRuleAction{
		StorageClass: dcl.ValueOrEmptyString(o.StorageClass),
		Type:         StorageAlphaBucketLifecycleRuleActionTypeEnumToProto(o.Type),
	}
	return p
}

// BucketLifecycleRuleConditionToProto converts a BucketLifecycleRuleCondition resource to its proto representation.
func StorageAlphaBucketLifecycleRuleConditionToProto(o *alpha.BucketLifecycleRuleCondition) *alphapb.StorageAlphaBucketLifecycleRuleCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketLifecycleRuleCondition{
		Age:              dcl.ValueOrEmptyInt64(o.Age),
		CreatedBefore:    dcl.ValueOrEmptyString(o.CreatedBefore),
		WithState:        StorageAlphaBucketLifecycleRuleConditionWithStateEnumToProto(o.WithState),
		NumNewerVersions: dcl.ValueOrEmptyInt64(o.NumNewerVersions),
	}
	for _, r := range o.MatchesStorageClass {
		p.MatchesStorageClass = append(p.MatchesStorageClass, r)
	}
	return p
}

// BucketLoggingToProto converts a BucketLogging resource to its proto representation.
func StorageAlphaBucketLoggingToProto(o *alpha.BucketLogging) *alphapb.StorageAlphaBucketLogging {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketLogging{
		LogBucket:       dcl.ValueOrEmptyString(o.LogBucket),
		LogObjectPrefix: dcl.ValueOrEmptyString(o.LogObjectPrefix),
	}
	return p
}

// BucketVersioningToProto converts a BucketVersioning resource to its proto representation.
func StorageAlphaBucketVersioningToProto(o *alpha.BucketVersioning) *alphapb.StorageAlphaBucketVersioning {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketVersioning{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// BucketWebsiteToProto converts a BucketWebsite resource to its proto representation.
func StorageAlphaBucketWebsiteToProto(o *alpha.BucketWebsite) *alphapb.StorageAlphaBucketWebsite {
	if o == nil {
		return nil
	}
	p := &alphapb.StorageAlphaBucketWebsite{
		MainPageSuffix: dcl.ValueOrEmptyString(o.MainPageSuffix),
		NotFoundPage:   dcl.ValueOrEmptyString(o.NotFoundPage),
	}
	return p
}

// BucketToProto converts a Bucket resource to its proto representation.
func BucketToProto(resource *alpha.Bucket) *alphapb.StorageAlphaBucket {
	p := &alphapb.StorageAlphaBucket{
		Project:      dcl.ValueOrEmptyString(resource.Project),
		Location:     dcl.ValueOrEmptyString(resource.Location),
		Name:         dcl.ValueOrEmptyString(resource.Name),
		Lifecycle:    StorageAlphaBucketLifecycleToProto(resource.Lifecycle),
		Logging:      StorageAlphaBucketLoggingToProto(resource.Logging),
		StorageClass: StorageAlphaBucketStorageClassEnumToProto(resource.StorageClass),
		Versioning:   StorageAlphaBucketVersioningToProto(resource.Versioning),
		Website:      StorageAlphaBucketWebsiteToProto(resource.Website),
	}
	for _, r := range resource.Cors {
		p.Cors = append(p.Cors, StorageAlphaBucketCorsToProto(&r))
	}

	return p
}

// ApplyBucket handles the gRPC request by passing it to the underlying Bucket Apply() method.
func (s *BucketServer) applyBucket(ctx context.Context, c *alpha.Client, request *alphapb.ApplyStorageAlphaBucketRequest) (*alphapb.StorageAlphaBucket, error) {
	p := ProtoToBucket(request.GetResource())
	res, err := c.ApplyBucket(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BucketToProto(res)
	return r, nil
}

// ApplyBucket handles the gRPC request by passing it to the underlying Bucket Apply() method.
func (s *BucketServer) ApplyStorageAlphaBucket(ctx context.Context, request *alphapb.ApplyStorageAlphaBucketRequest) (*alphapb.StorageAlphaBucket, error) {
	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBucket(ctx, cl, request)
}

// DeleteBucket handles the gRPC request by passing it to the underlying Bucket Delete() method.
func (s *BucketServer) DeleteStorageAlphaBucket(ctx context.Context, request *alphapb.DeleteStorageAlphaBucketRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBucket(ctx, ProtoToBucket(request.GetResource()))

}

// ListStorageAlphaBucket handles the gRPC request by passing it to the underlying BucketList() method.
func (s *BucketServer) ListStorageAlphaBucket(ctx context.Context, request *alphapb.ListStorageAlphaBucketRequest) (*alphapb.ListStorageAlphaBucketResponse, error) {
	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBucket(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.StorageAlphaBucket
	for _, r := range resources.Items {
		rp := BucketToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListStorageAlphaBucketResponse{Items: protos}, nil
}

func createConfigBucket(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
