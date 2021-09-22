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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudscheduler/alpha/cloudscheduler_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/alpha"
)

// Server implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobAppEngineHttpTargetHttpMethodEnum converts a JobAppEngineHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(e alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum) *alpha.JobAppEngineHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := alpha.JobAppEngineHttpTargetHttpMethodEnum(n[len("CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobHttpTargetHttpMethodEnum converts a JobHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTargetHttpMethodEnum(e alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum) *alpha.JobHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := alpha.JobHttpTargetHttpMethodEnum(n[len("CloudschedulerAlphaJobHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStateEnum converts a JobStateEnum enum from its proto representation.
func ProtoToCloudschedulerAlphaJobStateEnum(e alphapb.CloudschedulerAlphaJobStateEnum) *alpha.JobStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudschedulerAlphaJobStateEnum_name[int32(e)]; ok {
		e := alpha.JobStateEnum(n[len("CloudschedulerAlphaJobStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobPubsubTarget converts a JobPubsubTarget resource from its proto representation.
func ProtoToCloudschedulerAlphaJobPubsubTarget(p *alphapb.CloudschedulerAlphaJobPubsubTarget) *alpha.JobPubsubTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.JobPubsubTarget{
		TopicName: dcl.StringOrNil(p.TopicName),
		Data:      dcl.StringOrNil(p.Data),
	}
	return obj
}

// ProtoToJobAppEngineHttpTarget converts a JobAppEngineHttpTarget resource from its proto representation.
func ProtoToCloudschedulerAlphaJobAppEngineHttpTarget(p *alphapb.CloudschedulerAlphaJobAppEngineHttpTarget) *alpha.JobAppEngineHttpTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.JobAppEngineHttpTarget{
		HttpMethod:       ProtoToCloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		AppEngineRouting: ProtoToCloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting(p.GetAppEngineRouting()),
		RelativeUri:      dcl.StringOrNil(p.RelativeUri),
		Body:             dcl.StringOrNil(p.Body),
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetAppEngineRouting converts a JobAppEngineHttpTargetAppEngineRouting resource from its proto representation.
func ProtoToCloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting(p *alphapb.CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting) *alpha.JobAppEngineHttpTargetAppEngineRouting {
	if p == nil {
		return nil
	}
	obj := &alpha.JobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.StringOrNil(p.Service),
		Version:  dcl.StringOrNil(p.Version),
		Instance: dcl.StringOrNil(p.Instance),
		Host:     dcl.StringOrNil(p.Host),
	}
	return obj
}

// ProtoToJobHttpTarget converts a JobHttpTarget resource from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTarget(p *alphapb.CloudschedulerAlphaJobHttpTarget) *alpha.JobHttpTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.JobHttpTarget{
		Uri:        dcl.StringOrNil(p.Uri),
		HttpMethod: ProtoToCloudschedulerAlphaJobHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		Body:       dcl.StringOrNil(p.Body),
		OAuthToken: ProtoToCloudschedulerAlphaJobHttpTargetOAuthToken(p.GetOauthToken()),
		OidcToken:  ProtoToCloudschedulerAlphaJobHttpTargetOidcToken(p.GetOidcToken()),
	}
	return obj
}

// ProtoToJobHttpTargetOAuthToken converts a JobHttpTargetOAuthToken resource from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTargetOAuthToken(p *alphapb.CloudschedulerAlphaJobHttpTargetOAuthToken) *alpha.JobHttpTargetOAuthToken {
	if p == nil {
		return nil
	}
	obj := &alpha.JobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Scope:               dcl.StringOrNil(p.Scope),
	}
	return obj
}

// ProtoToJobHttpTargetOidcToken converts a JobHttpTargetOidcToken resource from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTargetOidcToken(p *alphapb.CloudschedulerAlphaJobHttpTargetOidcToken) *alpha.JobHttpTargetOidcToken {
	if p == nil {
		return nil
	}
	obj := &alpha.JobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Audience:            dcl.StringOrNil(p.Audience),
	}
	return obj
}

// ProtoToJobStatus converts a JobStatus resource from its proto representation.
func ProtoToCloudschedulerAlphaJobStatus(p *alphapb.CloudschedulerAlphaJobStatus) *alpha.JobStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.JobStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToCloudschedulerAlphaJobStatusDetails(r))
	}
	return obj
}

// ProtoToJobStatusDetails converts a JobStatusDetails resource from its proto representation.
func ProtoToCloudschedulerAlphaJobStatusDetails(p *alphapb.CloudschedulerAlphaJobStatusDetails) *alpha.JobStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.JobStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToJobRetryConfig converts a JobRetryConfig resource from its proto representation.
func ProtoToCloudschedulerAlphaJobRetryConfig(p *alphapb.CloudschedulerAlphaJobRetryConfig) *alpha.JobRetryConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.JobRetryConfig{
		RetryCount:         dcl.Int64OrNil(p.RetryCount),
		MaxRetryDuration:   dcl.StringOrNil(p.MaxRetryDuration),
		MinBackoffDuration: dcl.StringOrNil(p.MinBackoffDuration),
		MaxBackoffDuration: dcl.StringOrNil(p.MaxBackoffDuration),
		MaxDoublings:       dcl.Int64OrNil(p.MaxDoublings),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *alphapb.CloudschedulerAlphaJob) *alpha.Job {
	obj := &alpha.Job{
		Name:                dcl.StringOrNil(p.Name),
		Description:         dcl.StringOrNil(p.Description),
		PubsubTarget:        ProtoToCloudschedulerAlphaJobPubsubTarget(p.GetPubsubTarget()),
		AppEngineHttpTarget: ProtoToCloudschedulerAlphaJobAppEngineHttpTarget(p.GetAppEngineHttpTarget()),
		HttpTarget:          ProtoToCloudschedulerAlphaJobHttpTarget(p.GetHttpTarget()),
		Schedule:            dcl.StringOrNil(p.Schedule),
		TimeZone:            dcl.StringOrNil(p.TimeZone),
		UserUpdateTime:      dcl.StringOrNil(p.GetUserUpdateTime()),
		State:               ProtoToCloudschedulerAlphaJobStateEnum(p.GetState()),
		Status:              ProtoToCloudschedulerAlphaJobStatus(p.GetStatus()),
		ScheduleTime:        dcl.StringOrNil(p.GetScheduleTime()),
		LastAttemptTime:     dcl.StringOrNil(p.GetLastAttemptTime()),
		RetryConfig:         ProtoToCloudschedulerAlphaJobRetryConfig(p.GetRetryConfig()),
		AttemptDeadline:     dcl.StringOrNil(p.AttemptDeadline),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// JobAppEngineHttpTargetHttpMethodEnumToProto converts a JobAppEngineHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnumToProto(e *alpha.JobAppEngineHttpTargetHttpMethodEnum) alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum {
	if e == nil {
		return alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(0)
	}
	if v, ok := alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum_value["JobAppEngineHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(v)
	}
	return alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(0)
}

// JobHttpTargetHttpMethodEnumToProto converts a JobHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerAlphaJobHttpTargetHttpMethodEnumToProto(e *alpha.JobHttpTargetHttpMethodEnum) alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum {
	if e == nil {
		return alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum(0)
	}
	if v, ok := alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum_value["JobHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum(v)
	}
	return alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum(0)
}

// JobStateEnumToProto converts a JobStateEnum enum to its proto representation.
func CloudschedulerAlphaJobStateEnumToProto(e *alpha.JobStateEnum) alphapb.CloudschedulerAlphaJobStateEnum {
	if e == nil {
		return alphapb.CloudschedulerAlphaJobStateEnum(0)
	}
	if v, ok := alphapb.CloudschedulerAlphaJobStateEnum_value["JobStateEnum"+string(*e)]; ok {
		return alphapb.CloudschedulerAlphaJobStateEnum(v)
	}
	return alphapb.CloudschedulerAlphaJobStateEnum(0)
}

// JobPubsubTargetToProto converts a JobPubsubTarget resource to its proto representation.
func CloudschedulerAlphaJobPubsubTargetToProto(o *alpha.JobPubsubTarget) *alphapb.CloudschedulerAlphaJobPubsubTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobPubsubTarget{
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
func CloudschedulerAlphaJobAppEngineHttpTargetToProto(o *alpha.JobAppEngineHttpTarget) *alphapb.CloudschedulerAlphaJobAppEngineHttpTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobAppEngineHttpTarget{
		HttpMethod:       CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnumToProto(o.HttpMethod),
		AppEngineRouting: CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRoutingToProto(o.AppEngineRouting),
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
func CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRoutingToProto(o *alpha.JobAppEngineHttpTargetAppEngineRouting) *alphapb.CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.ValueOrEmptyString(o.Service),
		Version:  dcl.ValueOrEmptyString(o.Version),
		Instance: dcl.ValueOrEmptyString(o.Instance),
		Host:     dcl.ValueOrEmptyString(o.Host),
	}
	return p
}

// JobHttpTargetToProto converts a JobHttpTarget resource to its proto representation.
func CloudschedulerAlphaJobHttpTargetToProto(o *alpha.JobHttpTarget) *alphapb.CloudschedulerAlphaJobHttpTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobHttpTarget{
		Uri:        dcl.ValueOrEmptyString(o.Uri),
		HttpMethod: CloudschedulerAlphaJobHttpTargetHttpMethodEnumToProto(o.HttpMethod),
		Body:       dcl.ValueOrEmptyString(o.Body),
		OauthToken: CloudschedulerAlphaJobHttpTargetOAuthTokenToProto(o.OAuthToken),
		OidcToken:  CloudschedulerAlphaJobHttpTargetOidcTokenToProto(o.OidcToken),
	}
	p.Headers = make(map[string]string)
	for k, r := range o.Headers {
		p.Headers[k] = r
	}
	return p
}

// JobHttpTargetOAuthTokenToProto converts a JobHttpTargetOAuthToken resource to its proto representation.
func CloudschedulerAlphaJobHttpTargetOAuthTokenToProto(o *alpha.JobHttpTargetOAuthToken) *alphapb.CloudschedulerAlphaJobHttpTargetOAuthToken {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Scope:               dcl.ValueOrEmptyString(o.Scope),
	}
	return p
}

// JobHttpTargetOidcTokenToProto converts a JobHttpTargetOidcToken resource to its proto representation.
func CloudschedulerAlphaJobHttpTargetOidcTokenToProto(o *alpha.JobHttpTargetOidcToken) *alphapb.CloudschedulerAlphaJobHttpTargetOidcToken {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Audience:            dcl.ValueOrEmptyString(o.Audience),
	}
	return p
}

// JobStatusToProto converts a JobStatus resource to its proto representation.
func CloudschedulerAlphaJobStatusToProto(o *alpha.JobStatus) *alphapb.CloudschedulerAlphaJobStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, CloudschedulerAlphaJobStatusDetailsToProto(&r))
	}
	return p
}

// JobStatusDetailsToProto converts a JobStatusDetails resource to its proto representation.
func CloudschedulerAlphaJobStatusDetailsToProto(o *alpha.JobStatusDetails) *alphapb.CloudschedulerAlphaJobStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// JobRetryConfigToProto converts a JobRetryConfig resource to its proto representation.
func CloudschedulerAlphaJobRetryConfigToProto(o *alpha.JobRetryConfig) *alphapb.CloudschedulerAlphaJobRetryConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobRetryConfig{
		RetryCount:         dcl.ValueOrEmptyInt64(o.RetryCount),
		MaxRetryDuration:   dcl.ValueOrEmptyString(o.MaxRetryDuration),
		MinBackoffDuration: dcl.ValueOrEmptyString(o.MinBackoffDuration),
		MaxBackoffDuration: dcl.ValueOrEmptyString(o.MaxBackoffDuration),
		MaxDoublings:       dcl.ValueOrEmptyInt64(o.MaxDoublings),
	}
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *alpha.Job) *alphapb.CloudschedulerAlphaJob {
	p := &alphapb.CloudschedulerAlphaJob{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		PubsubTarget:        CloudschedulerAlphaJobPubsubTargetToProto(resource.PubsubTarget),
		AppEngineHttpTarget: CloudschedulerAlphaJobAppEngineHttpTargetToProto(resource.AppEngineHttpTarget),
		HttpTarget:          CloudschedulerAlphaJobHttpTargetToProto(resource.HttpTarget),
		Schedule:            dcl.ValueOrEmptyString(resource.Schedule),
		TimeZone:            dcl.ValueOrEmptyString(resource.TimeZone),
		UserUpdateTime:      dcl.ValueOrEmptyString(resource.UserUpdateTime),
		State:               CloudschedulerAlphaJobStateEnumToProto(resource.State),
		Status:              CloudschedulerAlphaJobStatusToProto(resource.Status),
		ScheduleTime:        dcl.ValueOrEmptyString(resource.ScheduleTime),
		LastAttemptTime:     dcl.ValueOrEmptyString(resource.LastAttemptTime),
		RetryConfig:         CloudschedulerAlphaJobRetryConfigToProto(resource.RetryConfig),
		AttemptDeadline:     dcl.ValueOrEmptyString(resource.AttemptDeadline),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudschedulerAlphaJobRequest) (*alphapb.CloudschedulerAlphaJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// ApplyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyCloudschedulerAlphaJob(ctx context.Context, request *alphapb.ApplyCloudschedulerAlphaJobRequest) (*alphapb.CloudschedulerAlphaJob, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteCloudschedulerAlphaJob(ctx context.Context, request *alphapb.DeleteCloudschedulerAlphaJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListCloudschedulerAlphaJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListCloudschedulerAlphaJob(ctx context.Context, request *alphapb.ListCloudschedulerAlphaJobRequest) (*alphapb.ListCloudschedulerAlphaJobResponse, error) {
	cl, err := createConfigJob(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudschedulerAlphaJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListCloudschedulerAlphaJobResponse{Items: protos}, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
