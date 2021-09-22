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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/storage/beta/storage_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/beta"
)

// Server implements the gRPC interface for Bucket.
type BucketServer struct{}

// ProtoToBucketLifecycleRuleActionTypeEnum converts a BucketLifecycleRuleActionTypeEnum enum from its proto representation.
func ProtoToStorageBetaBucketLifecycleRuleActionTypeEnum(e betapb.StorageBetaBucketLifecycleRuleActionTypeEnum) *beta.BucketLifecycleRuleActionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.StorageBetaBucketLifecycleRuleActionTypeEnum_name[int32(e)]; ok {
		e := beta.BucketLifecycleRuleActionTypeEnum(n[len("StorageBetaBucketLifecycleRuleActionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketLifecycleRuleConditionWithStateEnum converts a BucketLifecycleRuleConditionWithStateEnum enum from its proto representation.
func ProtoToStorageBetaBucketLifecycleRuleConditionWithStateEnum(e betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum) *beta.BucketLifecycleRuleConditionWithStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum_name[int32(e)]; ok {
		e := beta.BucketLifecycleRuleConditionWithStateEnum(n[len("StorageBetaBucketLifecycleRuleConditionWithStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketStorageClassEnum converts a BucketStorageClassEnum enum from its proto representation.
func ProtoToStorageBetaBucketStorageClassEnum(e betapb.StorageBetaBucketStorageClassEnum) *beta.BucketStorageClassEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.StorageBetaBucketStorageClassEnum_name[int32(e)]; ok {
		e := beta.BucketStorageClassEnum(n[len("StorageBetaBucketStorageClassEnum"):])
		return &e
	}
	return nil
}

// ProtoToBucketCors converts a BucketCors resource from its proto representation.
func ProtoToStorageBetaBucketCors(p *betapb.StorageBetaBucketCors) *beta.BucketCors {
	if p == nil {
		return nil
	}
	obj := &beta.BucketCors{
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
func ProtoToStorageBetaBucketLifecycle(p *betapb.StorageBetaBucketLifecycle) *beta.BucketLifecycle {
	if p == nil {
		return nil
	}
	obj := &beta.BucketLifecycle{}
	for _, r := range p.GetRule() {
		obj.Rule = append(obj.Rule, *ProtoToStorageBetaBucketLifecycleRule(r))
	}
	return obj
}

// ProtoToBucketLifecycleRule converts a BucketLifecycleRule resource from its proto representation.
func ProtoToStorageBetaBucketLifecycleRule(p *betapb.StorageBetaBucketLifecycleRule) *beta.BucketLifecycleRule {
	if p == nil {
		return nil
	}
	obj := &beta.BucketLifecycleRule{
		Action:    ProtoToStorageBetaBucketLifecycleRuleAction(p.GetAction()),
		Condition: ProtoToStorageBetaBucketLifecycleRuleCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToBucketLifecycleRuleAction converts a BucketLifecycleRuleAction resource from its proto representation.
func ProtoToStorageBetaBucketLifecycleRuleAction(p *betapb.StorageBetaBucketLifecycleRuleAction) *beta.BucketLifecycleRuleAction {
	if p == nil {
		return nil
	}
	obj := &beta.BucketLifecycleRuleAction{
		StorageClass: dcl.StringOrNil(p.StorageClass),
		Type:         ProtoToStorageBetaBucketLifecycleRuleActionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToBucketLifecycleRuleCondition converts a BucketLifecycleRuleCondition resource from its proto representation.
func ProtoToStorageBetaBucketLifecycleRuleCondition(p *betapb.StorageBetaBucketLifecycleRuleCondition) *beta.BucketLifecycleRuleCondition {
	if p == nil {
		return nil
	}
	obj := &beta.BucketLifecycleRuleCondition{
		Age:              dcl.Int64OrNil(p.Age),
		CreatedBefore:    dcl.StringOrNil(p.GetCreatedBefore()),
		WithState:        ProtoToStorageBetaBucketLifecycleRuleConditionWithStateEnum(p.GetWithState()),
		NumNewerVersions: dcl.Int64OrNil(p.NumNewerVersions),
	}
	for _, r := range p.GetMatchesStorageClass() {
		obj.MatchesStorageClass = append(obj.MatchesStorageClass, r)
	}
	return obj
}

// ProtoToBucketLogging converts a BucketLogging resource from its proto representation.
func ProtoToStorageBetaBucketLogging(p *betapb.StorageBetaBucketLogging) *beta.BucketLogging {
	if p == nil {
		return nil
	}
	obj := &beta.BucketLogging{
		LogBucket:       dcl.StringOrNil(p.LogBucket),
		LogObjectPrefix: dcl.StringOrNil(p.LogObjectPrefix),
	}
	return obj
}

// ProtoToBucketVersioning converts a BucketVersioning resource from its proto representation.
func ProtoToStorageBetaBucketVersioning(p *betapb.StorageBetaBucketVersioning) *beta.BucketVersioning {
	if p == nil {
		return nil
	}
	obj := &beta.BucketVersioning{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToBucketWebsite converts a BucketWebsite resource from its proto representation.
func ProtoToStorageBetaBucketWebsite(p *betapb.StorageBetaBucketWebsite) *beta.BucketWebsite {
	if p == nil {
		return nil
	}
	obj := &beta.BucketWebsite{
		MainPageSuffix: dcl.StringOrNil(p.MainPageSuffix),
		NotFoundPage:   dcl.StringOrNil(p.NotFoundPage),
	}
	return obj
}

// ProtoToBucket converts a Bucket resource from its proto representation.
func ProtoToBucket(p *betapb.StorageBetaBucket) *beta.Bucket {
	obj := &beta.Bucket{
		Project:      dcl.StringOrNil(p.Project),
		Location:     dcl.StringOrNil(p.Location),
		Name:         dcl.StringOrNil(p.Name),
		Lifecycle:    ProtoToStorageBetaBucketLifecycle(p.GetLifecycle()),
		Logging:      ProtoToStorageBetaBucketLogging(p.GetLogging()),
		StorageClass: ProtoToStorageBetaBucketStorageClassEnum(p.GetStorageClass()),
		Versioning:   ProtoToStorageBetaBucketVersioning(p.GetVersioning()),
		Website:      ProtoToStorageBetaBucketWebsite(p.GetWebsite()),
	}
	for _, r := range p.GetCors() {
		obj.Cors = append(obj.Cors, *ProtoToStorageBetaBucketCors(r))
	}
	return obj
}

// BucketLifecycleRuleActionTypeEnumToProto converts a BucketLifecycleRuleActionTypeEnum enum to its proto representation.
func StorageBetaBucketLifecycleRuleActionTypeEnumToProto(e *beta.BucketLifecycleRuleActionTypeEnum) betapb.StorageBetaBucketLifecycleRuleActionTypeEnum {
	if e == nil {
		return betapb.StorageBetaBucketLifecycleRuleActionTypeEnum(0)
	}
	if v, ok := betapb.StorageBetaBucketLifecycleRuleActionTypeEnum_value["BucketLifecycleRuleActionTypeEnum"+string(*e)]; ok {
		return betapb.StorageBetaBucketLifecycleRuleActionTypeEnum(v)
	}
	return betapb.StorageBetaBucketLifecycleRuleActionTypeEnum(0)
}

// BucketLifecycleRuleConditionWithStateEnumToProto converts a BucketLifecycleRuleConditionWithStateEnum enum to its proto representation.
func StorageBetaBucketLifecycleRuleConditionWithStateEnumToProto(e *beta.BucketLifecycleRuleConditionWithStateEnum) betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum {
	if e == nil {
		return betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum(0)
	}
	if v, ok := betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum_value["BucketLifecycleRuleConditionWithStateEnum"+string(*e)]; ok {
		return betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum(v)
	}
	return betapb.StorageBetaBucketLifecycleRuleConditionWithStateEnum(0)
}

// BucketStorageClassEnumToProto converts a BucketStorageClassEnum enum to its proto representation.
func StorageBetaBucketStorageClassEnumToProto(e *beta.BucketStorageClassEnum) betapb.StorageBetaBucketStorageClassEnum {
	if e == nil {
		return betapb.StorageBetaBucketStorageClassEnum(0)
	}
	if v, ok := betapb.StorageBetaBucketStorageClassEnum_value["BucketStorageClassEnum"+string(*e)]; ok {
		return betapb.StorageBetaBucketStorageClassEnum(v)
	}
	return betapb.StorageBetaBucketStorageClassEnum(0)
}

// BucketCorsToProto converts a BucketCors resource to its proto representation.
func StorageBetaBucketCorsToProto(o *beta.BucketCors) *betapb.StorageBetaBucketCors {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketCors{
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
func StorageBetaBucketLifecycleToProto(o *beta.BucketLifecycle) *betapb.StorageBetaBucketLifecycle {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketLifecycle{}
	for _, r := range o.Rule {
		p.Rule = append(p.Rule, StorageBetaBucketLifecycleRuleToProto(&r))
	}
	return p
}

// BucketLifecycleRuleToProto converts a BucketLifecycleRule resource to its proto representation.
func StorageBetaBucketLifecycleRuleToProto(o *beta.BucketLifecycleRule) *betapb.StorageBetaBucketLifecycleRule {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketLifecycleRule{
		Action:    StorageBetaBucketLifecycleRuleActionToProto(o.Action),
		Condition: StorageBetaBucketLifecycleRuleConditionToProto(o.Condition),
	}
	return p
}

// BucketLifecycleRuleActionToProto converts a BucketLifecycleRuleAction resource to its proto representation.
func StorageBetaBucketLifecycleRuleActionToProto(o *beta.BucketLifecycleRuleAction) *betapb.StorageBetaBucketLifecycleRuleAction {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketLifecycleRuleAction{
		StorageClass: dcl.ValueOrEmptyString(o.StorageClass),
		Type:         StorageBetaBucketLifecycleRuleActionTypeEnumToProto(o.Type),
	}
	return p
}

// BucketLifecycleRuleConditionToProto converts a BucketLifecycleRuleCondition resource to its proto representation.
func StorageBetaBucketLifecycleRuleConditionToProto(o *beta.BucketLifecycleRuleCondition) *betapb.StorageBetaBucketLifecycleRuleCondition {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketLifecycleRuleCondition{
		Age:              dcl.ValueOrEmptyInt64(o.Age),
		CreatedBefore:    dcl.ValueOrEmptyString(o.CreatedBefore),
		WithState:        StorageBetaBucketLifecycleRuleConditionWithStateEnumToProto(o.WithState),
		NumNewerVersions: dcl.ValueOrEmptyInt64(o.NumNewerVersions),
	}
	for _, r := range o.MatchesStorageClass {
		p.MatchesStorageClass = append(p.MatchesStorageClass, r)
	}
	return p
}

// BucketLoggingToProto converts a BucketLogging resource to its proto representation.
func StorageBetaBucketLoggingToProto(o *beta.BucketLogging) *betapb.StorageBetaBucketLogging {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketLogging{
		LogBucket:       dcl.ValueOrEmptyString(o.LogBucket),
		LogObjectPrefix: dcl.ValueOrEmptyString(o.LogObjectPrefix),
	}
	return p
}

// BucketVersioningToProto converts a BucketVersioning resource to its proto representation.
func StorageBetaBucketVersioningToProto(o *beta.BucketVersioning) *betapb.StorageBetaBucketVersioning {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketVersioning{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// BucketWebsiteToProto converts a BucketWebsite resource to its proto representation.
func StorageBetaBucketWebsiteToProto(o *beta.BucketWebsite) *betapb.StorageBetaBucketWebsite {
	if o == nil {
		return nil
	}
	p := &betapb.StorageBetaBucketWebsite{
		MainPageSuffix: dcl.ValueOrEmptyString(o.MainPageSuffix),
		NotFoundPage:   dcl.ValueOrEmptyString(o.NotFoundPage),
	}
	return p
}

// BucketToProto converts a Bucket resource to its proto representation.
func BucketToProto(resource *beta.Bucket) *betapb.StorageBetaBucket {
	p := &betapb.StorageBetaBucket{
		Project:      dcl.ValueOrEmptyString(resource.Project),
		Location:     dcl.ValueOrEmptyString(resource.Location),
		Name:         dcl.ValueOrEmptyString(resource.Name),
		Lifecycle:    StorageBetaBucketLifecycleToProto(resource.Lifecycle),
		Logging:      StorageBetaBucketLoggingToProto(resource.Logging),
		StorageClass: StorageBetaBucketStorageClassEnumToProto(resource.StorageClass),
		Versioning:   StorageBetaBucketVersioningToProto(resource.Versioning),
		Website:      StorageBetaBucketWebsiteToProto(resource.Website),
	}
	for _, r := range resource.Cors {
		p.Cors = append(p.Cors, StorageBetaBucketCorsToProto(&r))
	}

	return p
}

// ApplyBucket handles the gRPC request by passing it to the underlying Bucket Apply() method.
func (s *BucketServer) applyBucket(ctx context.Context, c *beta.Client, request *betapb.ApplyStorageBetaBucketRequest) (*betapb.StorageBetaBucket, error) {
	p := ProtoToBucket(request.GetResource())
	res, err := c.ApplyBucket(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BucketToProto(res)
	return r, nil
}

// ApplyBucket handles the gRPC request by passing it to the underlying Bucket Apply() method.
func (s *BucketServer) ApplyStorageBetaBucket(ctx context.Context, request *betapb.ApplyStorageBetaBucketRequest) (*betapb.StorageBetaBucket, error) {
	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBucket(ctx, cl, request)
}

// DeleteBucket handles the gRPC request by passing it to the underlying Bucket Delete() method.
func (s *BucketServer) DeleteStorageBetaBucket(ctx context.Context, request *betapb.DeleteStorageBetaBucketRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBucket(ctx, ProtoToBucket(request.GetResource()))

}

// ListStorageBetaBucket handles the gRPC request by passing it to the underlying BucketList() method.
func (s *BucketServer) ListStorageBetaBucket(ctx context.Context, request *betapb.ListStorageBetaBucketRequest) (*betapb.ListStorageBetaBucketResponse, error) {
	cl, err := createConfigBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBucket(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.StorageBetaBucket
	for _, r := range resources.Items {
		rp := BucketToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListStorageBetaBucketResponse{Items: protos}, nil
}

func createConfigBucket(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
