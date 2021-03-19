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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/beta/dataproc_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

// Server implements the gRPC interface for WorkflowTemplate.
type WorkflowTemplateServer struct{}

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacement(p *betapb.DataprocBetaWorkflowTemplatePlacement) *beta.WorkflowTemplatePlacement {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacement{
		ManagedCluster:  ProtoToDataprocBetaWorkflowTemplatePlacementManagedCluster(p.GetManagedCluster()),
		ClusterSelector: ProtoToDataprocBetaWorkflowTemplatePlacementClusterSelector(p.GetClusterSelector()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedCluster(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster) *beta.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.ClusterName),
		Config:      ProtoToDataprocBetaClusterClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementClusterSelector(p *betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector) *beta.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.Zone),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobs(p *betapb.DataprocBetaWorkflowTemplateJobs) *beta.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.StepId),
		HadoopJob:   ProtoToDataprocBetaWorkflowTemplateJobsHadoopJob(p.GetHadoopJob()),
		SparkJob:    ProtoToDataprocBetaWorkflowTemplateJobsSparkJob(p.GetSparkJob()),
		PysparkJob:  ProtoToDataprocBetaWorkflowTemplateJobsPysparkJob(p.GetPysparkJob()),
		HiveJob:     ProtoToDataprocBetaWorkflowTemplateJobsHiveJob(p.GetHiveJob()),
		PigJob:      ProtoToDataprocBetaWorkflowTemplateJobsPigJob(p.GetPigJob()),
		SparkRJob:   ProtoToDataprocBetaWorkflowTemplateJobsSparkRJob(p.GetSparkRJob()),
		SparkSqlJob: ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:   ProtoToDataprocBetaWorkflowTemplateJobsPrestoJob(p.GetPrestoJob()),
		Scheduling:  ProtoToDataprocBetaWorkflowTemplateJobsScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetPrerequisiteStepIds() {
		obj.PrerequisiteStepIds = append(obj.PrerequisiteStepIds, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHadoopJob(p *betapb.DataprocBetaWorkflowTemplateJobsHadoopJob) *beta.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig) *beta.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkJob) *beta.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.MainJarFileUri),
		MainClass:      dcl.StringOrNil(p.MainClass),
		LoggingConfig:  ProtoToDataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig) *beta.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPysparkJob(p *betapb.DataprocBetaWorkflowTemplateJobsPysparkJob) *beta.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.MainPythonFileUri),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig) *beta.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHiveJob(p *betapb.DataprocBetaWorkflowTemplateJobsHiveJob) *beta.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHiveJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList) *beta.WorkflowTemplateJobsHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJob(p *betapb.DataprocBetaWorkflowTemplateJobsPigJob) *beta.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList) *beta.WorkflowTemplateJobsPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig) *beta.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkRJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkRJob) *beta.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.MainRFileUri),
		LoggingConfig: ProtoToDataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig) *beta.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob) *beta.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.QueryFileUri),
		QueryList:     ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList) *beta.WorkflowTemplateJobsSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig) *beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJob(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJob) *beta.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.QueryFileUri),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.ContinueOnFailure),
		OutputFormat:      dcl.StringOrNil(p.OutputFormat),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList) *beta.WorkflowTemplateJobsPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig) *beta.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsScheduling(p *betapb.DataprocBetaWorkflowTemplateJobsScheduling) *beta.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.Int64OrNil(p.MaxFailuresTotal),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParameters(p *betapb.DataprocBetaWorkflowTemplateParameters) *beta.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Validation:  ProtoToDataprocBetaWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParametersValidation(p *betapb.DataprocBetaWorkflowTemplateParametersValidation) *beta.WorkflowTemplateParametersValidation {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParametersValidation{
		Regex:  ProtoToDataprocBetaWorkflowTemplateParametersValidationRegex(p.GetRegex()),
		Values: ProtoToDataprocBetaWorkflowTemplateParametersValidationValues(p.GetValues()),
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParametersValidationRegex(p *betapb.DataprocBetaWorkflowTemplateParametersValidationRegex) *beta.WorkflowTemplateParametersValidationRegex {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParametersValidationRegex{}
	for _, r := range p.GetRegexes() {
		obj.Regexes = append(obj.Regexes, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues resource from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParametersValidationValues(p *betapb.DataprocBetaWorkflowTemplateParametersValidationValues) *beta.WorkflowTemplateParametersValidationValues {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParametersValidationValues{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplate converts a WorkflowTemplate resource from its proto representation.
func ProtoToWorkflowTemplate(p *betapb.DataprocBetaWorkflowTemplate) *beta.WorkflowTemplate {
	obj := &beta.WorkflowTemplate{
		Name:       dcl.StringOrNil(p.Name),
		Version:    dcl.Int64OrNil(p.Version),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocBetaWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.DagTimeout),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocBetaWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocBetaWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement resource to its proto representation.
func DataprocBetaWorkflowTemplatePlacementToProto(o *beta.WorkflowTemplatePlacement) *betapb.DataprocBetaWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacement{
		ManagedCluster:  DataprocBetaWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster),
		ClusterSelector: DataprocBetaWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector),
	}
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster resource to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterToProto(o *beta.WorkflowTemplatePlacementManagedCluster) *betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.ValueOrEmptyString(o.ClusterName),
		Config:      DataprocBetaClusterClusterConfigToProto(o.Config),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector resource to its proto representation.
func DataprocBetaWorkflowTemplatePlacementClusterSelectorToProto(o *beta.WorkflowTemplatePlacementClusterSelector) *betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector{
		Zone: dcl.ValueOrEmptyString(o.Zone),
	}
	p.ClusterLabels = make(map[string]string)
	for k, r := range o.ClusterLabels {
		p.ClusterLabels[k] = r
	}
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsToProto(o *beta.WorkflowTemplateJobs) *betapb.DataprocBetaWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobs{
		StepId:      dcl.ValueOrEmptyString(o.StepId),
		HadoopJob:   DataprocBetaWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob),
		SparkJob:    DataprocBetaWorkflowTemplateJobsSparkJobToProto(o.SparkJob),
		PysparkJob:  DataprocBetaWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob),
		HiveJob:     DataprocBetaWorkflowTemplateJobsHiveJobToProto(o.HiveJob),
		PigJob:      DataprocBetaWorkflowTemplateJobsPigJobToProto(o.PigJob),
		SparkRJob:   DataprocBetaWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob),
		SparkSqlJob: DataprocBetaWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob),
		PrestoJob:   DataprocBetaWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob),
		Scheduling:  DataprocBetaWorkflowTemplateJobsSchedulingToProto(o.Scheduling),
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
func DataprocBetaWorkflowTemplateJobsHadoopJobToProto(o *beta.WorkflowTemplateJobsHadoopJob) *betapb.DataprocBetaWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsHadoopJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkJobToProto(o *beta.WorkflowTemplateJobsSparkJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.ValueOrEmptyString(o.MainJarFileUri),
		MainClass:      dcl.ValueOrEmptyString(o.MainClass),
		LoggingConfig:  DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsPysparkJobToProto(o *beta.WorkflowTemplateJobsPysparkJob) *betapb.DataprocBetaWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.ValueOrEmptyString(o.MainPythonFileUri),
		LoggingConfig:     DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPysparkJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsHiveJobToProto(o *beta.WorkflowTemplateJobsHiveJob) *betapb.DataprocBetaWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocBetaWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList),
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
func DataprocBetaWorkflowTemplateJobsHiveJobQueryListToProto(o *beta.WorkflowTemplateJobsHiveJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobToProto(o *beta.WorkflowTemplateJobsPigJob) *betapb.DataprocBetaWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocBetaWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		LoggingConfig:     DataprocBetaWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsPigJobQueryListToProto(o *beta.WorkflowTemplateJobsPigJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPigJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkRJobToProto(o *beta.WorkflowTemplateJobsSparkRJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.ValueOrEmptyString(o.MainRFileUri),
		LoggingConfig: DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkRJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobToProto(o *beta.WorkflowTemplateJobsSparkSqlJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:     DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList),
		LoggingConfig: DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *beta.WorkflowTemplateJobsSparkSqlJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobToProto(o *beta.WorkflowTemplateJobsPrestoJob) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.ValueOrEmptyString(o.QueryFileUri),
		QueryList:         DataprocBetaWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList),
		ContinueOnFailure: dcl.ValueOrEmptyBool(o.ContinueOnFailure),
		OutputFormat:      dcl.ValueOrEmptyString(o.OutputFormat),
		LoggingConfig:     DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig),
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
func DataprocBetaWorkflowTemplateJobsPrestoJobQueryListToProto(o *beta.WorkflowTemplateJobsPrestoJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range o.Queries {
		p.Queries = append(p.Queries, r)
	}
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPrestoJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig{}
	p.DriverLogLevels = make(map[string]string)
	for k, r := range o.DriverLogLevels {
		p.DriverLogLevels[k] = r
	}
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling resource to its proto representation.
func DataprocBetaWorkflowTemplateJobsSchedulingToProto(o *beta.WorkflowTemplateJobsScheduling) *betapb.DataprocBetaWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour),
		MaxFailuresTotal:   dcl.ValueOrEmptyInt64(o.MaxFailuresTotal),
	}
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters resource to its proto representation.
func DataprocBetaWorkflowTemplateParametersToProto(o *beta.WorkflowTemplateParameters) *betapb.DataprocBetaWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParameters{
		Name:        dcl.ValueOrEmptyString(o.Name),
		Description: dcl.ValueOrEmptyString(o.Description),
		Validation:  DataprocBetaWorkflowTemplateParametersValidationToProto(o.Validation),
	}
	for _, r := range o.Fields {
		p.Fields = append(p.Fields, r)
	}
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation resource to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationToProto(o *beta.WorkflowTemplateParametersValidation) *betapb.DataprocBetaWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidation{
		Regex:  DataprocBetaWorkflowTemplateParametersValidationRegexToProto(o.Regex),
		Values: DataprocBetaWorkflowTemplateParametersValidationValuesToProto(o.Values),
	}
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex resource to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationRegexToProto(o *beta.WorkflowTemplateParametersValidationRegex) *betapb.DataprocBetaWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidationRegex{}
	for _, r := range o.Regexes {
		p.Regexes = append(p.Regexes, r)
	}
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues resource to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationValuesToProto(o *beta.WorkflowTemplateParametersValidationValues) *betapb.DataprocBetaWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidationValues{}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *beta.WorkflowTemplate) *betapb.DataprocBetaWorkflowTemplate {
	p := &betapb.DataprocBetaWorkflowTemplate{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Version:    dcl.ValueOrEmptyInt64(resource.Version),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
		Placement:  DataprocBetaWorkflowTemplatePlacementToProto(resource.Placement),
		DagTimeout: dcl.ValueOrEmptyString(resource.DagTimeout),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Jobs {
		p.Jobs = append(p.Jobs, DataprocBetaWorkflowTemplateJobsToProto(&r))
	}
	for _, r := range resource.Parameters {
		p.Parameters = append(p.Parameters, DataprocBetaWorkflowTemplateParametersToProto(&r))
	}

	return p
}

// ApplyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaWorkflowTemplateRequest) (*betapb.DataprocBetaWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// ApplyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.ApplyDataprocBetaWorkflowTemplateRequest) (*betapb.DataprocBetaWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.DeleteDataprocBetaWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.ListDataprocBetaWorkflowTemplateRequest) (*betapb.ListDataprocBetaWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListDataprocBetaWorkflowTemplateResponse{Items: protos}, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
