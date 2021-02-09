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
	dataprocpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/dataproc_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
)

// Server implements the gRPC interface for WorkflowTemplate.
type WorkflowTemplateServer struct{}

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement resource from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacement(p *dataprocpb.DataprocWorkflowTemplatePlacement) *dataproc.WorkflowTemplatePlacement {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacement{
		ManagedCluster:  ProtoToDataprocWorkflowTemplatePlacementManagedCluster(p.GetManagedCluster()),
		ClusterSelector: ProtoToDataprocWorkflowTemplatePlacementClusterSelector(p.GetClusterSelector()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster resource from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedCluster(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster) *dataproc.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.ClusterName),
		Config:      ProtoToDataprocClusterClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector resource from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementClusterSelector(p *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector) *dataproc.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.Zone),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobs(p *dataprocpb.DataprocWorkflowTemplateJobs) *dataproc.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.StepId),
		HadoopJob:   ProtoToDataprocWorkflowTemplateJobsHadoopJob(p.GetHadoopJob()),
		SparkJob:    ProtoToDataprocWorkflowTemplateJobsSparkJob(p.GetSparkJob()),
		PysparkJob:  ProtoToDataprocWorkflowTemplateJobsPysparkJob(p.GetPysparkJob()),
		HiveJob:     ProtoToDataprocWorkflowTemplateJobsHiveJob(p.GetHiveJob()),
		PigJob:      ProtoToDataprocWorkflowTemplateJobsPigJob(p.GetPigJob()),
		SparkRJob:   ProtoToDataprocWorkflowTemplateJobsSparkRJob(p.GetSparkRJob()),
		SparkSqlJob: ProtoToDataprocWorkflowTemplateJobsSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:   ProtoToDataprocWorkflowTemplateJobsPrestoJob(p.GetPrestoJob()),
		Scheduling:  ProtoToDataprocWorkflowTemplateJobsScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetPrerequisiteStepIds() {
		obj.PrerequisiteStepIds = append(obj.PrerequisiteStepIds, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJob(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJob) *dataproc.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfig converts a WorkflowTemplateJobsHadoopJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJob) *dataproc.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfig converts a WorkflowTemplateJobsSparkJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJob(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJob) *dataproc.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.MainPythonFileUri),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetPythonFileUris() {
		obj.PythonFileUris = append(obj.PythonFileUris, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfig converts a WorkflowTemplateJobsPysparkJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJob(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJob) *dataproc.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList) *dataproc.WorkflowTemplateJobsHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJob(p *dataprocpb.DataprocWorkflowTemplateJobsPigJob) *dataproc.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList) *dataproc.WorkflowTemplateJobsPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig) *dataproc.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJob) *dataproc.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.MainRFileUri),
		LoggingConfig: ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfig converts a WorkflowTemplateJobsSparkRJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob) *dataproc.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.QueryFileUri),
		QueryList:     ProtoToDataprocWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList) *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJob(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJob) *dataproc.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		OutputFormat:      dcl.StringOrNil(p.OutputFormat),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList) *dataproc.WorkflowTemplateJobsPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsScheduling(p *dataprocpb.DataprocWorkflowTemplateJobsScheduling) *dataproc.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.Int64OrNil(p.MaxFailuresTotal),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters resource from its proto representation.
func ProtoToDataprocWorkflowTemplateParameters(p *dataprocpb.DataprocWorkflowTemplateParameters) *dataproc.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Validation:  ProtoToDataprocWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation resource from its proto representation.
func ProtoToDataprocWorkflowTemplateParametersValidation(p *dataprocpb.DataprocWorkflowTemplateParametersValidation) *dataproc.WorkflowTemplateParametersValidation {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParametersValidation{
		Regex:  ProtoToDataprocWorkflowTemplateParametersValidationRegex(p.GetRegex()),
		Values: ProtoToDataprocWorkflowTemplateParametersValidationValues(p.GetValues()),
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex resource from its proto representation.
func ProtoToDataprocWorkflowTemplateParametersValidationRegex(p *dataprocpb.DataprocWorkflowTemplateParametersValidationRegex) *dataproc.WorkflowTemplateParametersValidationRegex {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParametersValidationRegex{}
	for _, r := range p.GetRegexes() {
		obj.Regexes = append(obj.Regexes, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues resource from its proto representation.
func ProtoToDataprocWorkflowTemplateParametersValidationValues(p *dataprocpb.DataprocWorkflowTemplateParametersValidationValues) *dataproc.WorkflowTemplateParametersValidationValues {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParametersValidationValues{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplate converts a WorkflowTemplate resource from its proto representation.
func ProtoToWorkflowTemplate(p *dataprocpb.DataprocWorkflowTemplate) *dataproc.WorkflowTemplate {
	obj := &dataproc.WorkflowTemplate{
		Name:       dcl.StringOrNil(p.Name),
		Version:    dcl.Int64OrNil(p.Version),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocWorkflowTemplatePlacement(p.GetPlacement()),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement resource to its proto representation.
func DataprocWorkflowTemplatePlacementToProto(o *dataproc.WorkflowTemplatePlacement) *dataprocpb.DataprocWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacement{
		ManagedCluster:  DataprocWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster),
		ClusterSelector: DataprocWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector),
	}
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster resource to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterToProto(o *dataproc.WorkflowTemplatePlacementManagedCluster) *dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.ValueOrEmptyString(o.ClusterName),
		Config:      DataprocClusterClusterConfigToProto(o.Config),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector resource to its proto representation.
func DataprocWorkflowTemplatePlacementClusterSelectorToProto(o *dataproc.WorkflowTemplatePlacementClusterSelector) *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector{
		Zone: dcl.ValueOrEmptyString(o.Zone),
	}
	p.ClusterLabels = make(map[string]string)
	for k, r := range o.ClusterLabels {
		p.ClusterLabels[k] = r
	}
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs resource to its proto representation.
func DataprocWorkflowTemplateJobsToProto(o *dataproc.WorkflowTemplateJobs) *dataprocpb.DataprocWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobs{
		StepId:      dcl.ValueOrEmptyString(o.StepId),
		HadoopJob:   DataprocWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob),
		SparkJob:    DataprocWorkflowTemplateJobsSparkJobToProto(o.SparkJob),
		PysparkJob:  DataprocWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob),
		HiveJob:     DataprocWorkflowTemplateJobsHiveJobToProto(o.HiveJob),
		PigJob:      DataprocWorkflowTemplateJobsPigJobToProto(o.PigJob),
		SparkRJob:   DataprocWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob),
		SparkSqlJob: DataprocWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob),
		PrestoJob:   DataprocWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob),
		Scheduling:  DataprocWorkflowTemplateJobsSchedulingToProto(o.Scheduling),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	for _, r := range o.PrerequisiteStepIds {
		p.PrerequisiteStepIds = append(p.PrerequisiteStepIds, r)
	}
	return p
}

// WorkflowTemplateJobsHadoopJobToProto converts a WorkflowTemplateJobsHadoopJob resource to its proto representation.
func DataprocWorkflowTemplateJobsHadoopJobToProto(o *dataproc.WorkflowTemplateJobsHadoopJob) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// WorkflowTemplateJobsHadoopJobLoggingConfigToProto converts a WorkflowTemplateJobsHadoopJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobToProto(o *dataproc.WorkflowTemplateJobsSparkJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob resource to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobToProto(o *dataproc.WorkflowTemplateJobsPysparkJob) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.ValueOrEmptyString(o.MainPythonFileUri),
		LoggingConfig:     DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.PythonFileUris {
		p.PythonFileUris = append(p.PythonFileUris, r)
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// WorkflowTemplateJobsPysparkJobLoggingConfigToProto converts a WorkflowTemplateJobsPysparkJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob resource to its proto representation.
func DataprocWorkflowTemplateJobsHiveJobToProto(o *dataproc.WorkflowTemplateJobsHiveJob) *dataprocpb.DataprocWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
	}
	p.ScriptVariables = make(map[string]string)
	for k, r := range o.ScriptVariables {
		p.ScriptVariables[k] = r
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	return p
}

// WorkflowTemplateJobsHiveJobQueryListToProto converts a WorkflowTemplateJobsHiveJobQueryList resource to its proto representation.
func DataprocWorkflowTemplateJobsHiveJobQueryListToProto(o *dataproc.WorkflowTemplateJobsHiveJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobToProto(o *dataproc.WorkflowTemplateJobsPigJob) *dataprocpb.DataprocWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		LoggingConfig:     DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig),
	}
	p.ScriptVariables = make(map[string]string)
	for k, r := range o.ScriptVariables {
		p.ScriptVariables[k] = r
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobQueryListToProto converts a WorkflowTemplateJobsPigJobQueryList resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobQueryListToProto(o *dataproc.WorkflowTemplateJobsPigJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPigJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobToProto(o *dataproc.WorkflowTemplateJobsSparkRJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.ValueOrEmptyString(o.MainRFileUri),
		LoggingConfig: DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.FileUris {
		p.FileUris = append(p.FileUris, r)
	}
	for _, r := range o.ArchiveUris {
		p.ArchiveUris = append(p.ArchiveUris, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkRJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkRJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:     DataprocWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList),
		LoggingConfig: DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig),
	}
	p.ScriptVariables = make(map[string]string)
	for k, r := range o.ScriptVariables {
		p.ScriptVariables[k] = r
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	for _, r := range o.JarFileUris {
		p.JarFileUris = append(p.JarFileUris, r)
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobQueryListToProto converts a WorkflowTemplateJobsSparkSqlJobQueryList resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob resource to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobToProto(o *dataproc.WorkflowTemplateJobsPrestoJob) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		OutputFormat:      dcl.ValueOrEmptyString(o.OutputFormat),
		LoggingConfig:     DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig),
	}
	for _, r := range o.ClientTags {
		p.ClientTags = append(p.ClientTags, r)
	}
	p.Properties = make(map[string]string)
	for k, r := range o.Properties {
		p.Properties[k] = r
	}
	return p
}

// WorkflowTemplateJobsPrestoJobQueryListToProto converts a WorkflowTemplateJobsPrestoJobQueryList resource to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobQueryListToProto(o *dataproc.WorkflowTemplateJobsPrestoJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling resource to its proto representation.
func DataprocWorkflowTemplateJobsSchedulingToProto(o *dataproc.WorkflowTemplateJobsScheduling) *dataprocpb.DataprocWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.ValueOrEmptyInt64(o.MaxFailuresTotal),
	}
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters resource to its proto representation.
func DataprocWorkflowTemplateParametersToProto(o *dataproc.WorkflowTemplateParameters) *dataprocpb.DataprocWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParameters{
		Name:        dcl.ValueOrEmptyString(o.Name),
		Description: dcl.ValueOrEmptyString(o.Description),
		Validation:  DataprocWorkflowTemplateParametersValidationToProto(o.Validation),
	}
	for _, r := range o.Fields {
		p.Fields = append(p.Fields, r)
	}
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation resource to its proto representation.
func DataprocWorkflowTemplateParametersValidationToProto(o *dataproc.WorkflowTemplateParametersValidation) *dataprocpb.DataprocWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidation{
		Regex:  DataprocWorkflowTemplateParametersValidationRegexToProto(o.Regex),
		Values: DataprocWorkflowTemplateParametersValidationValuesToProto(o.Values),
	}
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex resource to its proto representation.
func DataprocWorkflowTemplateParametersValidationRegexToProto(o *dataproc.WorkflowTemplateParametersValidationRegex) *dataprocpb.DataprocWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidationRegex{}
	for _, r := range o.Regexes {
		p.Regexes = append(p.Regexes, r)
	}
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues resource to its proto representation.
func DataprocWorkflowTemplateParametersValidationValuesToProto(o *dataproc.WorkflowTemplateParametersValidationValues) *dataprocpb.DataprocWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidationValues{}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *dataproc.WorkflowTemplate) *dataprocpb.DataprocWorkflowTemplate {
	p := &dataprocpb.DataprocWorkflowTemplate{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Version:    dcl.ValueOrEmptyInt64(resource.Version),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
		Placement:  DataprocWorkflowTemplatePlacementToProto(resource.Placement),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Jobs {
		p.Jobs = append(p.Jobs, DataprocWorkflowTemplateJobsToProto(&r))
	}
	for _, r := range resource.Parameters {
		p.Parameters = append(p.Parameters, DataprocWorkflowTemplateParametersToProto(&r))
	}

	return p
}

// ApplyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocWorkflowTemplateRequest) (*dataprocpb.DataprocWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// ApplyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.ApplyDataprocWorkflowTemplateRequest) (*dataprocpb.DataprocWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.DeleteDataprocWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.ListDataprocWorkflowTemplateRequest) (*dataprocpb.ListDataprocWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &dataprocpb.ListDataprocWorkflowTemplateResponse{Items: protos}, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
