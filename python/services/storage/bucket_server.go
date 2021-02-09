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
	storagepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/storage/storage_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/serialization"
)

// Server implements the gRPC interface for Bucket.
type BucketServer struct{}

// ProtoToBucketLifecycleRuleActionTypeEnum converts a BucketLifecycleRuleActionTypeEnum enum from its proto representation.
func ProtoToStorageBucketLifecycleRuleActionTypeEnum(e storagepb.StorageBucketLifecycleRuleActionTypeEnum) *storage.BucketLifecycleRuleActionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageBucketLifecycleRuleActionTypeEnum_name[int32(e)]; ok {
		e := storage.BucketLifecycleRuleActionTypeEnum(n[len("BucketLifecycleRuleActionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketLifecycleRuleConditionWithStateEnum converts a BucketLifecycleRuleConditionWithStateEnum enum from its proto representation.
func ProtoToStorageBucketLifecycleRuleConditionWithStateEnum(e storagepb.StorageBucketLifecycleRuleConditionWithStateEnum) *storage.BucketLifecycleRuleConditionWithStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageBucketLifecycleRuleConditionWithStateEnum_name[int32(e)]; ok {
		e := storage.BucketLifecycleRuleConditionWithStateEnum(n[len("BucketLifecycleRuleConditionWithStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketStorageClassEnum converts a BucketStorageClassEnum enum from its proto representation.
func ProtoToStorageBucketStorageClassEnum(e storagepb.StorageBucketStorageClassEnum) *storage.BucketStorageClassEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageBucketStorageClassEnum_name[int32(e)]; ok {
		e := storage.BucketStorageClassEnum(n[len("BucketStorageClassEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketCors converts a BucketCors resource from its proto representation.
func ProtoToStorageBucketCors(p *storagepb.StorageBucketCors) *storage.BucketCors {
	if p == nil {
		return nil
	}
	obj := &storage.BucketCors{
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
func ProtoToStorageBucketLifecycle(p *storagepb.StorageBucketLifecycle) *storage.BucketLifecycle {
	if p == nil {
		return nil
	}
	obj := &storage.BucketLifecycle{}
	for _, r := range p.GetRule() {
		obj.Rule = append(obj.Rule, *ProtoToStorageBucketLifecycleRule(r))
	}
	return obj
}

// ProtoToBucketLifecycleRule converts a BucketLifecycleRule resource from its proto representation.
func ProtoToStorageBucketLifecycleRule(p *storagepb.StorageBucketLifecycleRule) *storage.BucketLifecycleRule {
	if p == nil {
		return nil
	}
	obj := &storage.BucketLifecycleRule{
		Action:    ProtoToStorageBucketLifecycleRuleAction(p.GetAction()),
		Condition: ProtoToStorageBucketLifecycleRuleCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToBucketLifecycleRuleAction converts a BucketLifecycleRuleAction resource from its proto representation.
func ProtoToStorageBucketLifecycleRuleAction(p *storagepb.StorageBucketLifecycleRuleAction) *storage.BucketLifecycleRuleAction {
	if p == nil {
		return nil
	}
	obj := &storage.BucketLifecycleRuleAction{
		StorageClass: dcl.StringOrNil(p.StorageClass),
		Type:         ProtoToStorageBucketLifecycleRuleActionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToBucketLifecycleRuleCondition converts a BucketLifecycleRuleCondition resource from its proto representation.
func ProtoToStorageBucketLifecycleRuleCondition(p *storagepb.StorageBucketLifecycleRuleCondition) *storage.BucketLifecycleRuleCondition {
	if p == nil {
		return nil
	}
	obj := &storage.BucketLifecycleRuleCondition{
		Age:              dcl.Int64OrNil(p.Age),
		CreatedBefore:    dcl.StringOrNil(p.GetCreatedBefore()),
		WithState:        ProtoToStorageBucketLifecycleRuleConditionWithStateEnum(p.GetWithState()),
		NumNewerVersions: dcl.Int64OrNil(p.NumNewerVersions),
	}
	for _, r := range p.GetMatchesStorageClass() {
		obj.MatchesStorageClass = append(obj.MatchesStorageClass, r)
	}
	return obj
}

// ProtoToBucketLogging converts a BucketLogging resource from its proto representation.
func ProtoToStorageBucketLogging(p *storagepb.StorageBucketLogging) *storage.BucketLogging {
	if p == nil {
		return nil
	}
	obj := &storage.BucketLogging{
		LogBucket:       dcl.StringOrNil(p.LogBucket),
		LogObjectPrefix: dcl.StringOrNil(p.LogObjectPrefix),
	}
	return obj
}

// ProtoToBucketVersioning converts a BucketVersioning resource from its proto representation.
func ProtoToStorageBucketVersioning(p *storagepb.StorageBucketVersioning) *storage.BucketVersioning {
	if p == nil {
		return nil
	}
	obj := &storage.BucketVersioning{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToBucketWebsite converts a BucketWebsite resource from its proto representation.
func ProtoToStorageBucketWebsite(p *storagepb.StorageBucketWebsite) *storage.BucketWebsite {
	if p == nil {
		return nil
	}
	obj := &storage.BucketWebsite{
		MainPageSuffix: dcl.StringOrNil(p.MainPageSuffix),
		NotFoundPage:   dcl.StringOrNil(p.NotFoundPage),
	}
	return obj
}

// ProtoToBucket converts a Bucket resource from its proto representation.
func ProtoToBucket(p *storagepb.StorageBucket) *storage.Bucket {
	obj := &storage.Bucket{
		Project:      dcl.StringOrNil(p.Project),
		Location:     dcl.StringOrNil(p.Location),
		Name:         dcl.StringOrNil(p.Name),
		Lifecycle:    ProtoToStorageBucketLifecycle(p.GetLifecycle()),
		Logging:      ProtoToStorageBucketLogging(p.GetLogging()),
		StorageClass: ProtoToStorageBucketStorageClassEnum(p.GetStorageClass()),
		Versioning:   ProtoToStorageBucketVersioning(p.GetVersioning()),
		Website:      ProtoToStorageBucketWebsite(p.GetWebsite()),
	}
	for _, r := range p.GetCors() {
		obj.Cors = append(obj.Cors, *ProtoToStorageBucketCors(r))
	}
	return obj
}

// BucketLifecycleRuleActionTypeEnumToProto converts a BucketLifecycleRuleActionTypeEnum enum to its proto representation.
func StorageBucketLifecycleRuleActionTypeEnumToProto(e *storage.BucketLifecycleRuleActionTypeEnum) storagepb.StorageBucketLifecycleRuleActionTypeEnum {
	if e == nil {
		return storagepb.StorageBucketLifecycleRuleActionTypeEnum(0)
	}
	if v, ok := storagepb.StorageBucketLifecycleRuleActionTypeEnum_value["BucketLifecycleRuleActionTypeEnum"+string(*e)]; ok {
		return storagepb.StorageBucketLifecycleRuleActionTypeEnum(v)
	}
	return storagepb.StorageBucketLifecycleRuleActionTypeEnum(0)
}

// BucketLifecycleRuleConditionWithStateEnumToProto converts a BucketLifecycleRuleConditionWithStateEnum enum to its proto representation.
func StorageBucketLifecycleRuleConditionWithStateEnumToProto(e *storage.BucketLifecycleRuleConditionWithStateEnum) storagepb.StorageBucketLifecycleRuleConditionWithStateEnum {
	if e == nil {
		return storagepb.StorageBucketLifecycleRuleConditionWithStateEnum(0)
	}
	if v, ok := storagepb.StorageBucketLifecycleRuleConditionWithStateEnum_value["BucketLifecycleRuleConditionWithStateEnum"+string(*e)]; ok {
		return storagepb.StorageBucketLifecycleRuleConditionWithStateEnum(v)
	}
	return storagepb.StorageBucketLifecycleRuleConditionWithStateEnum(0)
}

// BucketStorageClassEnumToProto converts a BucketStorageClassEnum enum to its proto representation.
func StorageBucketStorageClassEnumToProto(e *storage.BucketStorageClassEnum) storagepb.StorageBucketStorageClassEnum {
	if e == nil {
		return storagepb.StorageBucketStorageClassEnum(0)
	}
	if v, ok := storagepb.StorageBucketStorageClassEnum_value["BucketStorageClassEnum"+string(*e)]; ok {
		return storagepb.StorageBucketStorageClassEnum(v)
	}
	return storagepb.StorageBucketStorageClassEnum(0)
}

// BucketCorsToProto converts a BucketCors resource to its proto representation.
func StorageBucketCorsToProto(o *storage.BucketCors) *storagepb.StorageBucketCors {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketCors{
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
func StorageBucketLifecycleToProto(o *storage.BucketLifecycle) *storagepb.StorageBucketLifecycle {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketLifecycle{}
	for _, r := range o.Rule {
		p.Rule = append(p.Rule, StorageBucketLifecycleRuleToProto(&r))
	}
	return p
}

// BucketLifecycleRuleToProto converts a BucketLifecycleRule resource to its proto representation.
func StorageBucketLifecycleRuleToProto(o *storage.BucketLifecycleRule) *storagepb.StorageBucketLifecycleRule {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketLifecycleRule{
		Action:    StorageBucketLifecycleRuleActionToProto(o.Action),
		Condition: StorageBucketLifecycleRuleConditionToProto(o.Condition),
	}
	return p
}

// BucketLifecycleRuleActionToProto converts a BucketLifecycleRuleAction resource to its proto representation.
func StorageBucketLifecycleRuleActionToProto(o *storage.BucketLifecycleRuleAction) *storagepb.StorageBucketLifecycleRuleAction {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketLifecycleRuleAction{
		StorageClass: dcl.ValueOrEmptyString(o.StorageClass),
		Type:         StorageBucketLifecycleRuleActionTypeEnumToProto(o.Type),
	}
	return p
}

// BucketLifecycleRuleConditionToProto converts a BucketLifecycleRuleCondition resource to its proto representation.
func StorageBucketLifecycleRuleConditionToProto(o *storage.BucketLifecycleRuleCondition) *storagepb.StorageBucketLifecycleRuleCondition {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketLifecycleRuleCondition{
		Age:              dcl.ValueOrEmptyInt64(o.Age),
		CreatedBefore:    dcl.ValueOrEmptyString(o.CreatedBefore),
		WithState:        StorageBucketLifecycleRuleConditionWithStateEnumToProto(o.WithState),
		NumNewerVersions: dcl.ValueOrEmptyInt64(o.NumNewerVersions),
	}
	for _, r := range o.MatchesStorageClass {
		p.MatchesStorageClass = append(p.MatchesStorageClass, r)
	}
	return p
}

// BucketLoggingToProto converts a BucketLogging resource to its proto representation.
func StorageBucketLoggingToProto(o *storage.BucketLogging) *storagepb.StorageBucketLogging {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketLogging{
		LogBucket:       dcl.ValueOrEmptyString(o.LogBucket),
		LogObjectPrefix: dcl.ValueOrEmptyString(o.LogObjectPrefix),
	}
	return p
}

// BucketVersioningToProto converts a BucketVersioning resource to its proto representation.
func StorageBucketVersioningToProto(o *storage.BucketVersioning) *storagepb.StorageBucketVersioning {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketVersioning{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// BucketWebsiteToProto converts a BucketWebsite resource to its proto representation.
func StorageBucketWebsiteToProto(o *storage.BucketWebsite) *storagepb.StorageBucketWebsite {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageBucketWebsite{
		MainPageSuffix: dcl.ValueOrEmptyString(o.MainPageSuffix),
		NotFoundPage:   dcl.ValueOrEmptyString(o.NotFoundPage),
	}
	return p
}

// BucketToProto converts a Bucket resource to its proto representation.
func BucketToProto(resource *storage.Bucket) *storagepb.StorageBucket {
	p := &storagepb.StorageBucket{
		Project:      dcl.ValueOrEmptyString(resource.Project),
		Location:     dcl.ValueOrEmptyString(resource.Location),
		Name:         dcl.ValueOrEmptyString(resource.Name),
		Lifecycle:    StorageBucketLifecycleToProto(resource.Lifecycle),
		Logging:      StorageBucketLoggingToProto(resource.Logging),
		StorageClass: StorageBucketStorageClassEnumToProto(resource.StorageClass),
		Versioning:   StorageBucketVersioningToProto(resource.Versioning),
		Website:      StorageBucketWebsiteToProto(resource.Website),
	}
	for _, r := range resource.Cors {
		p.Cors = append(p.Cors, StorageBucketCorsToProto(&r))
	}

	return p
}

// ApplyBucket handles the gRPC request by passing it to the underlying Bucket Apply() method.
func (s *BucketServer) applyBucket(ctx context.Context, c *storage.Client, request *storagepb.ApplyStorageBucketRequest) (*storagepb.StorageBucket, error) {
	p := ProtoToBucket(request.GetResource())
	res, err := c.ApplyBucket(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BucketToProto(res)
	return r, nil
}

// ApplyBucket handles the gRPC request by passing it to the underlying Bucket Apply() method.
func (s *BucketServer) ApplyStorageBucket(ctx context.Context, request *storagepb.ApplyStorageBucketRequest) (*storagepb.StorageBucket, error) {
	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBucket(ctx, cl, request)
}

func (s *BucketServer) StorageBucketAsHcl(ctx context.Context, request *storagepb.StorageBucketAsHclRequest) (*storagepb.StorageBucketAsHclResponse, error) {
	p := ProtoToBucket(request.GetResource())
	resStr, err := serialization.StorageBucketAsHCL(*p)
	if err != nil {
		return nil, err
	}

	return &storagepb.StorageBucketAsHclResponse{Hcl: resStr}, nil
}

// DeleteBucket handles the gRPC request by passing it to the underlying Bucket Delete() method.
func (s *BucketServer) DeleteStorageBucket(ctx context.Context, request *storagepb.DeleteStorageBucketRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBucket(ctx, ProtoToBucket(request.GetResource()))

}

// ListBucket handles the gRPC request by passing it to the underlying BucketList() method.
func (s *BucketServer) ListStorageBucket(ctx context.Context, request *storagepb.ListStorageBucketRequest) (*storagepb.ListStorageBucketResponse, error) {
	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBucket(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*storagepb.StorageBucket
	for _, r := range resources.Items {
		rp := BucketToProto(r)
		protos = append(protos, rp)
	}
	return &storagepb.ListStorageBucketResponse{Items: protos}, nil
}

func createConfigBucket(ctx context.Context, service_account_file string) (*storage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return storage.NewClient(conf), nil
}
