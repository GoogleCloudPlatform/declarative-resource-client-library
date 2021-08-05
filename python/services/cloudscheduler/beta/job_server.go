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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudscheduler/beta/cloudscheduler_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/beta"
)

// Server implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobAppEngineHttpTargetHttpMethodEnum converts a JobAppEngineHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(e betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum) *beta.JobAppEngineHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := beta.JobAppEngineHttpTargetHttpMethodEnum(n[len("CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobHttpTargetHttpMethodEnum converts a JobHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTargetHttpMethodEnum(e betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum) *beta.JobHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := beta.JobHttpTargetHttpMethodEnum(n[len("CloudschedulerBetaJobHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStateEnum converts a JobStateEnum enum from its proto representation.
func ProtoToCloudschedulerBetaJobStateEnum(e betapb.CloudschedulerBetaJobStateEnum) *beta.JobStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudschedulerBetaJobStateEnum_name[int32(e)]; ok {
		e := beta.JobStateEnum(n[len("CloudschedulerBetaJobStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobPubsubTarget converts a JobPubsubTarget resource from its proto representation.
func ProtoToCloudschedulerBetaJobPubsubTarget(p *betapb.CloudschedulerBetaJobPubsubTarget) *beta.JobPubsubTarget {
	if p == nil {
		return nil
	}
	obj := &beta.JobPubsubTarget{
		TopicName: dcl.StringOrNil(p.TopicName),
		Data:      dcl.StringOrNil(p.Data),
	}
	return obj
}

// ProtoToJobAppEngineHttpTarget converts a JobAppEngineHttpTarget resource from its proto representation.
func ProtoToCloudschedulerBetaJobAppEngineHttpTarget(p *betapb.CloudschedulerBetaJobAppEngineHttpTarget) *beta.JobAppEngineHttpTarget {
	if p == nil {
		return nil
	}
	obj := &beta.JobAppEngineHttpTarget{
		HttpMethod:       ProtoToCloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		AppEngineRouting: ProtoToCloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting(p.GetAppEngineRouting()),
		RelativeUri:      dcl.StringOrNil(p.RelativeUri),
		Body:             dcl.StringOrNil(p.Body),
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetAppEngineRouting converts a JobAppEngineHttpTargetAppEngineRouting resource from its proto representation.
func ProtoToCloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting(p *betapb.CloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting) *beta.JobAppEngineHttpTargetAppEngineRouting {
	if p == nil {
		return nil
	}
	obj := &beta.JobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.StringOrNil(p.Service),
		Version:  dcl.StringOrNil(p.Version),
		Instance: dcl.StringOrNil(p.Instance),
		Host:     dcl.StringOrNil(p.Host),
	}
	return obj
}

// ProtoToJobHttpTarget converts a JobHttpTarget resource from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTarget(p *betapb.CloudschedulerBetaJobHttpTarget) *beta.JobHttpTarget {
	if p == nil {
		return nil
	}
	obj := &beta.JobHttpTarget{
		Uri:        dcl.StringOrNil(p.Uri),
		HttpMethod: ProtoToCloudschedulerBetaJobHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		Body:       dcl.StringOrNil(p.Body),
		OAuthToken: ProtoToCloudschedulerBetaJobHttpTargetOAuthToken(p.GetOauthToken()),
		OidcToken:  ProtoToCloudschedulerBetaJobHttpTargetOidcToken(p.GetOidcToken()),
	}
	return obj
}

// ProtoToJobHttpTargetOAuthToken converts a JobHttpTargetOAuthToken resource from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTargetOAuthToken(p *betapb.CloudschedulerBetaJobHttpTargetOAuthToken) *beta.JobHttpTargetOAuthToken {
	if p == nil {
		return nil
	}
	obj := &beta.JobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Scope:               dcl.StringOrNil(p.Scope),
	}
	return obj
}

// ProtoToJobHttpTargetOidcToken converts a JobHttpTargetOidcToken resource from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTargetOidcToken(p *betapb.CloudschedulerBetaJobHttpTargetOidcToken) *beta.JobHttpTargetOidcToken {
	if p == nil {
		return nil
	}
	obj := &beta.JobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Audience:            dcl.StringOrNil(p.Audience),
	}
	return obj
}

// ProtoToJobStatus converts a JobStatus resource from its proto representation.
func ProtoToCloudschedulerBetaJobStatus(p *betapb.CloudschedulerBetaJobStatus) *beta.JobStatus {
	if p == nil {
		return nil
	}
	obj := &beta.JobStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToCloudschedulerBetaJobStatusDetails(r))
	}
	return obj
}

// ProtoToJobStatusDetails converts a JobStatusDetails resource from its proto representation.
func ProtoToCloudschedulerBetaJobStatusDetails(p *betapb.CloudschedulerBetaJobStatusDetails) *beta.JobStatusDetails {
	if p == nil {
		return nil
	}
	obj := &beta.JobStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobRetryConfig converts a JobRetryConfig resource from its proto representation.
func ProtoToCloudschedulerBetaJobRetryConfig(p *betapb.CloudschedulerBetaJobRetryConfig) *beta.JobRetryConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobRetryConfig{
		RetryCount:         dcl.Int64OrNil(p.RetryCount),
		MaxRetryDuration:   dcl.StringOrNil(p.MaxRetryDuration),
		MinBackoffDuration: dcl.StringOrNil(p.MinBackoffDuration),
		MaxBackoffDuration: dcl.StringOrNil(p.MaxBackoffDuration),
		MaxDoublings:       dcl.Int64OrNil(p.MaxDoublings),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *betapb.CloudschedulerBetaJob) *beta.Job {
	obj := &beta.Job{
		Name:                dcl.StringOrNil(p.Name),
		Description:         dcl.StringOrNil(p.Description),
		PubsubTarget:        ProtoToCloudschedulerBetaJobPubsubTarget(p.GetPubsubTarget()),
		AppEngineHttpTarget: ProtoToCloudschedulerBetaJobAppEngineHttpTarget(p.GetAppEngineHttpTarget()),
		HttpTarget:          ProtoToCloudschedulerBetaJobHttpTarget(p.GetHttpTarget()),
		Schedule:            dcl.StringOrNil(p.Schedule),
		TimeZone:            dcl.StringOrNil(p.TimeZone),
		UserUpdateTime:      dcl.StringOrNil(p.GetUserUpdateTime()),
		State:               ProtoToCloudschedulerBetaJobStateEnum(p.GetState()),
		Status:              ProtoToCloudschedulerBetaJobStatus(p.GetStatus()),
		ScheduleTime:        dcl.StringOrNil(p.GetScheduleTime()),
		LastAttemptTime:     dcl.StringOrNil(p.GetLastAttemptTime()),
		RetryConfig:         ProtoToCloudschedulerBetaJobRetryConfig(p.GetRetryConfig()),
		AttemptDeadline:     dcl.StringOrNil(p.AttemptDeadline),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// JobAppEngineHttpTargetHttpMethodEnumToProto converts a JobAppEngineHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnumToProto(e *beta.JobAppEngineHttpTargetHttpMethodEnum) betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum {
	if e == nil {
		return betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(0)
	}
	if v, ok := betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum_value["JobAppEngineHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(v)
	}
	return betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(0)
}

// JobHttpTargetHttpMethodEnumToProto converts a JobHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerBetaJobHttpTargetHttpMethodEnumToProto(e *beta.JobHttpTargetHttpMethodEnum) betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum {
	if e == nil {
		return betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum(0)
	}
	if v, ok := betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum_value["JobHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum(v)
	}
	return betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum(0)
}

// JobStateEnumToProto converts a JobStateEnum enum to its proto representation.
func CloudschedulerBetaJobStateEnumToProto(e *beta.JobStateEnum) betapb.CloudschedulerBetaJobStateEnum {
	if e == nil {
		return betapb.CloudschedulerBetaJobStateEnum(0)
	}
	if v, ok := betapb.CloudschedulerBetaJobStateEnum_value["JobStateEnum"+string(*e)]; ok {
		return betapb.CloudschedulerBetaJobStateEnum(v)
	}
	return betapb.CloudschedulerBetaJobStateEnum(0)
}

// JobPubsubTargetToProto converts a JobPubsubTarget resource to its proto representation.
func CloudschedulerBetaJobPubsubTargetToProto(o *beta.JobPubsubTarget) *betapb.CloudschedulerBetaJobPubsubTarget {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobPubsubTarget{
		TopicName: dcl.ValueOrEmptyString(o.TopicName),
		Data:      dcl.ValueOrEmptyString(o.Data),
	}
	p.Attributes = make(map[string]string)
	for k, r := range o.Attributes {
		p.Attributes[k] = r
	}
	return p
}

// JobAppEngineHttpTargetToProto converts a JobAppEngineHttpTarget resource to its proto representation.
func CloudschedulerBetaJobAppEngineHttpTargetToProto(o *beta.JobAppEngineHttpTarget) *betapb.CloudschedulerBetaJobAppEngineHttpTarget {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobAppEngineHttpTarget{
		HttpMethod:       CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnumToProto(o.HttpMethod),
		AppEngineRouting: CloudschedulerBetaJobAppEngineHttpTargetAppEngineRoutingToProto(o.AppEngineRouting),
		RelativeUri:      dcl.ValueOrEmptyString(o.RelativeUri),
		Body:             dcl.ValueOrEmptyString(o.Body),
	}
	p.Headers = make(map[string]string)
	for k, r := range o.Headers {
		p.Headers[k] = r
	}
	return p
}

// JobAppEngineHttpTargetAppEngineRoutingToProto converts a JobAppEngineHttpTargetAppEngineRouting resource to its proto representation.
func CloudschedulerBetaJobAppEngineHttpTargetAppEngineRoutingToProto(o *beta.JobAppEngineHttpTargetAppEngineRouting) *betapb.CloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.ValueOrEmptyString(o.Service),
		Version:  dcl.ValueOrEmptyString(o.Version),
		Instance: dcl.ValueOrEmptyString(o.Instance),
		Host:     dcl.ValueOrEmptyString(o.Host),
	}
	return p
}

// JobHttpTargetToProto converts a JobHttpTarget resource to its proto representation.
func CloudschedulerBetaJobHttpTargetToProto(o *beta.JobHttpTarget) *betapb.CloudschedulerBetaJobHttpTarget {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobHttpTarget{
		Uri:        dcl.ValueOrEmptyString(o.Uri),
		HttpMethod: CloudschedulerBetaJobHttpTargetHttpMethodEnumToProto(o.HttpMethod),
		Body:       dcl.ValueOrEmptyString(o.Body),
		OauthToken: CloudschedulerBetaJobHttpTargetOAuthTokenToProto(o.OAuthToken),
		OidcToken:  CloudschedulerBetaJobHttpTargetOidcTokenToProto(o.OidcToken),
	}
	p.Headers = make(map[string]string)
	for k, r := range o.Headers {
		p.Headers[k] = r
	}
	return p
}

// JobHttpTargetOAuthTokenToProto converts a JobHttpTargetOAuthToken resource to its proto representation.
func CloudschedulerBetaJobHttpTargetOAuthTokenToProto(o *beta.JobHttpTargetOAuthToken) *betapb.CloudschedulerBetaJobHttpTargetOAuthToken {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Scope:               dcl.ValueOrEmptyString(o.Scope),
	}
	return p
}

// JobHttpTargetOidcTokenToProto converts a JobHttpTargetOidcToken resource to its proto representation.
func CloudschedulerBetaJobHttpTargetOidcTokenToProto(o *beta.JobHttpTargetOidcToken) *betapb.CloudschedulerBetaJobHttpTargetOidcToken {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Audience:            dcl.ValueOrEmptyString(o.Audience),
	}
	return p
}

// JobStatusToProto converts a JobStatus resource to its proto representation.
func CloudschedulerBetaJobStatusToProto(o *beta.JobStatus) *betapb.CloudschedulerBetaJobStatus {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, CloudschedulerBetaJobStatusDetailsToProto(&r))
	}
	return p
}

// JobStatusDetailsToProto converts a JobStatusDetails resource to its proto representation.
func CloudschedulerBetaJobStatusDetailsToProto(o *beta.JobStatusDetails) *betapb.CloudschedulerBetaJobStatusDetails {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobRetryConfigToProto converts a JobRetryConfig resource to its proto representation.
func CloudschedulerBetaJobRetryConfigToProto(o *beta.JobRetryConfig) *betapb.CloudschedulerBetaJobRetryConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobRetryConfig{
		RetryCount:         dcl.ValueOrEmptyInt64(o.RetryCount),
		MaxRetryDuration:   dcl.ValueOrEmptyString(o.MaxRetryDuration),
		MinBackoffDuration: dcl.ValueOrEmptyString(o.MinBackoffDuration),
		MaxBackoffDuration: dcl.ValueOrEmptyString(o.MaxBackoffDuration),
		MaxDoublings:       dcl.ValueOrEmptyInt64(o.MaxDoublings),
	}
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *beta.Job) *betapb.CloudschedulerBetaJob {
	p := &betapb.CloudschedulerBetaJob{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		PubsubTarget:        CloudschedulerBetaJobPubsubTargetToProto(resource.PubsubTarget),
		AppEngineHttpTarget: CloudschedulerBetaJobAppEngineHttpTargetToProto(resource.AppEngineHttpTarget),
		HttpTarget:          CloudschedulerBetaJobHttpTargetToProto(resource.HttpTarget),
		Schedule:            dcl.ValueOrEmptyString(resource.Schedule),
		TimeZone:            dcl.ValueOrEmptyString(resource.TimeZone),
		UserUpdateTime:      dcl.ValueOrEmptyString(resource.UserUpdateTime),
		State:               CloudschedulerBetaJobStateEnumToProto(resource.State),
		Status:              CloudschedulerBetaJobStatusToProto(resource.Status),
		ScheduleTime:        dcl.ValueOrEmptyString(resource.ScheduleTime),
		LastAttemptTime:     dcl.ValueOrEmptyString(resource.LastAttemptTime),
		RetryConfig:         CloudschedulerBetaJobRetryConfigToProto(resource.RetryConfig),
		AttemptDeadline:     dcl.ValueOrEmptyString(resource.AttemptDeadline),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudschedulerBetaJobRequest) (*betapb.CloudschedulerBetaJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyCloudschedulerBetaJob(ctx context.Context, request *betapb.ApplyCloudschedulerBetaJobRequest) (*betapb.CloudschedulerBetaJob, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteCloudschedulerBetaJob(ctx context.Context, request *betapb.DeleteCloudschedulerBetaJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListCloudschedulerBetaJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListCloudschedulerBetaJob(ctx context.Context, request *betapb.ListCloudschedulerBetaJobRequest) (*betapb.ListCloudschedulerBetaJobResponse, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, ProtoToJob(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudschedulerBetaJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListCloudschedulerBetaJobResponse{Items: protos}, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
