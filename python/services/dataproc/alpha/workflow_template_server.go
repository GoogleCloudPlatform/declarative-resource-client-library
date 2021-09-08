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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/alpha/dataproc_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
)

// Server implements the gRPC interface for WorkflowTemplate.
type WorkflowTemplateServer struct{}

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacement(p *alphapb.DataprocAlphaWorkflowTemplatePlacement) *alpha.WorkflowTemplatePlacement {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacement{
		ManagedCluster:  ProtoToDataprocAlphaWorkflowTemplatePlacementManagedCluster(p.GetManagedCluster()),
		ClusterSelector: ProtoToDataprocAlphaWorkflowTemplatePlacementClusterSelector(p.GetClusterSelector()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedCluster(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster) *alpha.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.ClusterName),
		Config:      ProtoToDataprocAlphaClusterClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementClusterSelector(p *alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector) *alpha.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.Zone),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobs(p *alphapb.DataprocAlphaWorkflowTemplateJobs) *alpha.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.StepId),
		HadoopJob:   ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJob(p.GetHadoopJob()),
		SparkJob:    ProtoToDataprocAlphaWorkflowTemplateJobsSparkJob(p.GetSparkJob()),
		PysparkJob:  ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJob(p.GetPysparkJob()),
		HiveJob:     ProtoToDataprocAlphaWorkflowTemplateJobsHiveJob(p.GetHiveJob()),
		PigJob:      ProtoToDataprocAlphaWorkflowTemplateJobsPigJob(p.GetPigJob()),
		SparkRJob:   ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJob(p.GetSparkRJob()),
		SparkSqlJob: ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:   ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJob(p.GetPrestoJob()),
		Scheduling:  ProtoToDataprocAlphaWorkflowTemplateJobsScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetPrerequisiteStepIds() {
		obj.PrerequisiteStepIds = append(obj.PrerequisiteStepIds, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob) *alpha.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig) *alpha.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob) *alpha.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob) *alpha.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.MainPythonFileUri),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig) *alpha.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHiveJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob) *alpha.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHiveJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList) *alpha.WorkflowTemplateJobsHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJob) *alpha.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList) *alpha.WorkflowTemplateJobsPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig) *alpha.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob) *alpha.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.MainRFileUri),
		LoggingConfig: ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob) *alpha.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.QueryFileUri),
		QueryList:     ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList) *alpha.WorkflowTemplateJobsSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob) *alpha.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		OutputFormat:      dcl.StringOrNil(p.OutputFormat),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList) *alpha.WorkflowTemplateJobsPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig) *alpha.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsScheduling(p *alphapb.DataprocAlphaWorkflowTemplateJobsScheduling) *alpha.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.Int64OrNil(p.MaxFailuresTotal),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParameters(p *alphapb.DataprocAlphaWorkflowTemplateParameters) *alpha.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Validation:  ProtoToDataprocAlphaWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParametersValidation(p *alphapb.DataprocAlphaWorkflowTemplateParametersValidation) *alpha.WorkflowTemplateParametersValidation {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParametersValidation{
		Regex:  ProtoToDataprocAlphaWorkflowTemplateParametersValidationRegex(p.GetRegex()),
		Values: ProtoToDataprocAlphaWorkflowTemplateParametersValidationValues(p.GetValues()),
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParametersValidationRegex(p *alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex) *alpha.WorkflowTemplateParametersValidationRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParametersValidationRegex{}
	for _, r := range p.GetRegexes() {
		obj.Regexes = append(obj.Regexes, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues resource from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParametersValidationValues(p *alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues) *alpha.WorkflowTemplateParametersValidationValues {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParametersValidationValues{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplate converts a WorkflowTemplate resource from its proto representation.
func ProtoToWorkflowTemplate(p *alphapb.DataprocAlphaWorkflowTemplate) *alpha.WorkflowTemplate {
	obj := &alpha.WorkflowTemplate{
		Name:       dcl.StringOrNil(p.Name),
		Version:    dcl.Int64OrNil(p.Version),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocAlphaWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.DagTimeout),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocAlphaWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocAlphaWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement resource to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementToProto(o *alpha.WorkflowTemplatePlacement) *alphapb.DataprocAlphaWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacement{
		ManagedCluster:  DataprocAlphaWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster),
		ClusterSelector: DataprocAlphaWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector),
	}
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster resource to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterToProto(o *alpha.WorkflowTemplatePlacementManagedCluster) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.ValueOrEmptyString(o.ClusterName),
		Config:      DataprocAlphaClusterClusterConfigToProto(o.Config),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector resource to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementClusterSelectorToProto(o *alpha.WorkflowTemplatePlacementClusterSelector) *alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector{
		Zone: dcl.ValueOrEmptyString(o.Zone),
	}
	p.ClusterLabels = make(map[string]string)
	for k, r := range o.ClusterLabels {
		p.ClusterLabels[k] = r
	}
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsToProto(o *alpha.WorkflowTemplateJobs) *alphapb.DataprocAlphaWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobs{
		StepId:      dcl.ValueOrEmptyString(o.StepId),
		HadoopJob:   DataprocAlphaWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob),
		SparkJob:    DataprocAlphaWorkflowTemplateJobsSparkJobToProto(o.SparkJob),
		PysparkJob:  DataprocAlphaWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob),
		HiveJob:     DataprocAlphaWorkflowTemplateJobsHiveJobToProto(o.HiveJob),
		PigJob:      DataprocAlphaWorkflowTemplateJobsPigJobToProto(o.PigJob),
		SparkRJob:   DataprocAlphaWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob),
		SparkSqlJob: DataprocAlphaWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob),
		PrestoJob:   DataprocAlphaWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob),
		Scheduling:  DataprocAlphaWorkflowTemplateJobsSchedulingToProto(o.Scheduling),
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
func DataprocAlphaWorkflowTemplateJobsHadoopJobToProto(o *alpha.WorkflowTemplateJobsHadoopJob) *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsHadoopJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkJobToProto(o *alpha.WorkflowTemplateJobsSparkJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPysparkJobToProto(o *alpha.WorkflowTemplateJobsPysparkJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.ValueOrEmptyString(o.MainPythonFileUri),
		LoggingConfig:     DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPysparkJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsHiveJobToProto(o *alpha.WorkflowTemplateJobsHiveJob) *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocAlphaWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList),
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
func DataprocAlphaWorkflowTemplateJobsHiveJobQueryListToProto(o *alpha.WorkflowTemplateJobsHiveJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobToProto(o *alpha.WorkflowTemplateJobsPigJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocAlphaWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		LoggingConfig:     DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsPigJobQueryListToProto(o *alpha.WorkflowTemplateJobsPigJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPigJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkRJobToProto(o *alpha.WorkflowTemplateJobsSparkRJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.ValueOrEmptyString(o.MainRFileUri),
		LoggingConfig: DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkRJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobToProto(o *alpha.WorkflowTemplateJobsSparkSqlJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:     DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList),
		LoggingConfig: DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *alpha.WorkflowTemplateJobsSparkSqlJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobToProto(o *alpha.WorkflowTemplateJobsPrestoJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocAlphaWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		OutputFormat:      dcl.ValueOrEmptyString(o.OutputFormat),
		LoggingConfig:     DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocAlphaWorkflowTemplateJobsPrestoJobQueryListToProto(o *alpha.WorkflowTemplateJobsPrestoJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPrestoJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling resource to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSchedulingToProto(o *alpha.WorkflowTemplateJobsScheduling) *alphapb.DataprocAlphaWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.ValueOrEmptyInt64(o.MaxFailuresTotal),
	}
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters resource to its proto representation.
func DataprocAlphaWorkflowTemplateParametersToProto(o *alpha.WorkflowTemplateParameters) *alphapb.DataprocAlphaWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParameters{
		Name:        dcl.ValueOrEmptyString(o.Name),
		Description: dcl.ValueOrEmptyString(o.Description),
		Validation:  DataprocAlphaWorkflowTemplateParametersValidationToProto(o.Validation),
	}
	for _, r := range o.Fields {
		p.Fields = append(p.Fields, r)
	}
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation resource to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationToProto(o *alpha.WorkflowTemplateParametersValidation) *alphapb.DataprocAlphaWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidation{
		Regex:  DataprocAlphaWorkflowTemplateParametersValidationRegexToProto(o.Regex),
		Values: DataprocAlphaWorkflowTemplateParametersValidationValuesToProto(o.Values),
	}
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex resource to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationRegexToProto(o *alpha.WorkflowTemplateParametersValidationRegex) *alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex{}
	for _, r := range o.Regexes {
		p.Regexes = append(p.Regexes, r)
	}
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues resource to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationValuesToProto(o *alpha.WorkflowTemplateParametersValidationValues) *alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues{}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *alpha.WorkflowTemplate) *alphapb.DataprocAlphaWorkflowTemplate {
	p := &alphapb.DataprocAlphaWorkflowTemplate{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Version:    dcl.ValueOrEmptyInt64(resource.Version),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
		Placement:  DataprocAlphaWorkflowTemplatePlacementToProto(resource.Placement),
		DagTimeout: dcl.ValueOrEmptyString(resource.DagTimeout),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Jobs {
		p.Jobs = append(p.Jobs, DataprocAlphaWorkflowTemplateJobsToProto(&r))
	}
	for _, r := range resource.Parameters {
		p.Parameters = append(p.Parameters, DataprocAlphaWorkflowTemplateParametersToProto(&r))
	}

	return p
}

// ApplyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaWorkflowTemplateRequest) (*alphapb.DataprocAlphaWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// ApplyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.ApplyDataprocAlphaWorkflowTemplateRequest) (*alphapb.DataprocAlphaWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.DeleteDataprocAlphaWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocAlphaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.ListDataprocAlphaWorkflowTemplateRequest) (*alphapb.ListDataprocAlphaWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListDataprocAlphaWorkflowTemplateResponse{Items: protos}, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
