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

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement object from its proto representation.
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

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedCluster(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster) *beta.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.GetClusterName()),
		Config:      ProtoToDataprocBetaClusterClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementClusterSelector(p *betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector) *beta.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobs(p *betapb.DataprocBetaWorkflowTemplateJobs) *beta.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.GetStepId()),
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

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHadoopJob(p *betapb.DataprocBetaWorkflowTemplateJobsHadoopJob) *beta.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
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

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfig converts a WorkflowTemplateJobsHadoopJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig) *beta.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkJob) *beta.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
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

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfig converts a WorkflowTemplateJobsSparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig) *beta.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPysparkJob(p *betapb.DataprocBetaWorkflowTemplateJobsPysparkJob) *beta.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.GetMainPythonFileUri()),
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

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfig converts a WorkflowTemplateJobsPysparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig) *beta.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHiveJob(p *betapb.DataprocBetaWorkflowTemplateJobsHiveJob) *beta.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJob(p *betapb.DataprocBetaWorkflowTemplateJobsPigJob) *beta.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig) *beta.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkRJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkRJob) *beta.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.GetMainRFileUri()),
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

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfig converts a WorkflowTemplateJobsSparkRJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig) *beta.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob) *beta.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:     ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig) *beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJob(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJob) *beta.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		OutputFormat:      dcl.StringOrNil(p.GetOutputFormat()),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig) *beta.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsScheduling(p *betapb.DataprocBetaWorkflowTemplateJobsScheduling) *beta.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.GetMaxFailuresPerHour()),
		MaxFailuresTotal:   dcl.Int64OrNil(p.GetMaxFailuresTotal()),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParameters(p *betapb.DataprocBetaWorkflowTemplateParameters) *beta.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Validation:  ProtoToDataprocBetaWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation object from its proto representation.
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

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex object from its proto representation.
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

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues object from its proto representation.
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
		Name:       dcl.StringOrNil(p.GetName()),
		Version:    dcl.Int64OrNil(p.GetVersion()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocBetaWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.GetDagTimeout()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocBetaWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocBetaWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementToProto(o *beta.WorkflowTemplatePlacement) *betapb.DataprocBetaWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacement{}
	p.SetManagedCluster(DataprocBetaWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster))
	p.SetClusterSelector(DataprocBetaWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector))
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterToProto(o *beta.WorkflowTemplatePlacementManagedCluster) *betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster{}
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetConfig(DataprocBetaClusterClusterConfigToProto(o.Config))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementClusterSelectorToProto(o *beta.WorkflowTemplatePlacementClusterSelector) *betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	mClusterLabels := make(map[string]string, len(o.ClusterLabels))
	for k, r := range o.ClusterLabels {
		mClusterLabels[k] = r
	}
	p.SetClusterLabels(mClusterLabels)
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs object to its proto representation.
func DataprocBetaWorkflowTemplateJobsToProto(o *beta.WorkflowTemplateJobs) *betapb.DataprocBetaWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobs{}
	p.SetStepId(dcl.ValueOrEmptyString(o.StepId))
	p.SetHadoopJob(DataprocBetaWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob))
	p.SetSparkJob(DataprocBetaWorkflowTemplateJobsSparkJobToProto(o.SparkJob))
	p.SetPysparkJob(DataprocBetaWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob))
	p.SetHiveJob(DataprocBetaWorkflowTemplateJobsHiveJobToProto(o.HiveJob))
	p.SetPigJob(DataprocBetaWorkflowTemplateJobsPigJobToProto(o.PigJob))
	p.SetSparkRJob(DataprocBetaWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob))
	p.SetSparkSqlJob(DataprocBetaWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob))
	p.SetPrestoJob(DataprocBetaWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob))
	p.SetScheduling(DataprocBetaWorkflowTemplateJobsSchedulingToProto(o.Scheduling))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sPrerequisiteStepIds := make([]string, len(o.PrerequisiteStepIds))
	for i, r := range o.PrerequisiteStepIds {
		sPrerequisiteStepIds[i] = r
	}
	p.SetPrerequisiteStepIds(sPrerequisiteStepIds)
	return p
}

// WorkflowTemplateJobsHadoopJobToProto converts a WorkflowTemplateJobsHadoopJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsHadoopJobToProto(o *beta.WorkflowTemplateJobsHadoopJob) *betapb.DataprocBetaWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHadoopJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsHadoopJobLoggingConfigToProto converts a WorkflowTemplateJobsHadoopJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsHadoopJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkJobToProto(o *beta.WorkflowTemplateJobsSparkJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsSparkJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPysparkJobToProto(o *beta.WorkflowTemplateJobsPysparkJob) *betapb.DataprocBetaWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPysparkJob{}
	p.SetMainPythonFileUri(dcl.ValueOrEmptyString(o.MainPythonFileUri))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sPythonFileUris := make([]string, len(o.PythonFileUris))
	for i, r := range o.PythonFileUris {
		sPythonFileUris[i] = r
	}
	p.SetPythonFileUris(sPythonFileUris)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsPysparkJobLoggingConfigToProto converts a WorkflowTemplateJobsPysparkJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPysparkJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsHiveJobToProto(o *beta.WorkflowTemplateJobsHiveJob) *betapb.DataprocBetaWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHiveJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	mScriptVariables := make(map[string]string, len(o.ScriptVariables))
	for k, r := range o.ScriptVariables {
		mScriptVariables[k] = r
	}
	p.SetScriptVariables(mScriptVariables)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	return p
}

// WorkflowTemplateJobsHiveJobQueryListToProto converts a WorkflowTemplateJobsHiveJobQueryList object to its proto representation.
func DataprocBetaWorkflowTemplateJobsHiveJobQueryListToProto(o *beta.WorkflowTemplateJobsHiveJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobToProto(o *beta.WorkflowTemplateJobsPigJob) *betapb.DataprocBetaWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig))
	mScriptVariables := make(map[string]string, len(o.ScriptVariables))
	for k, r := range o.ScriptVariables {
		mScriptVariables[k] = r
	}
	p.SetScriptVariables(mScriptVariables)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	return p
}

// WorkflowTemplateJobsPigJobQueryListToProto converts a WorkflowTemplateJobsPigJobQueryList object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobQueryListToProto(o *beta.WorkflowTemplateJobsPigJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPigJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkRJobToProto(o *beta.WorkflowTemplateJobsSparkRJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkRJob{}
	p.SetMainRFileUri(dcl.ValueOrEmptyString(o.MainRFileUri))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsSparkRJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkRJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkRJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobToProto(o *beta.WorkflowTemplateJobsSparkSqlJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig))
	mScriptVariables := make(map[string]string, len(o.ScriptVariables))
	for k, r := range o.ScriptVariables {
		mScriptVariables[k] = r
	}
	p.SetScriptVariables(mScriptVariables)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	return p
}

// WorkflowTemplateJobsSparkSqlJobQueryListToProto converts a WorkflowTemplateJobsSparkSqlJobQueryList object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *beta.WorkflowTemplateJobsSparkSqlJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobToProto(o *beta.WorkflowTemplateJobsPrestoJob) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetOutputFormat(dcl.ValueOrEmptyString(o.OutputFormat))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig))
	sClientTags := make([]string, len(o.ClientTags))
	for i, r := range o.ClientTags {
		sClientTags[i] = r
	}
	p.SetClientTags(sClientTags)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsPrestoJobQueryListToProto converts a WorkflowTemplateJobsPrestoJobQueryList object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobQueryListToProto(o *beta.WorkflowTemplateJobsPrestoJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPrestoJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSchedulingToProto(o *beta.WorkflowTemplateJobsScheduling) *betapb.DataprocBetaWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsScheduling{}
	p.SetMaxFailuresPerHour(dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour))
	p.SetMaxFailuresTotal(dcl.ValueOrEmptyInt64(o.MaxFailuresTotal))
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters object to its proto representation.
func DataprocBetaWorkflowTemplateParametersToProto(o *beta.WorkflowTemplateParameters) *betapb.DataprocBetaWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParameters{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetValidation(DataprocBetaWorkflowTemplateParametersValidationToProto(o.Validation))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation object to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationToProto(o *beta.WorkflowTemplateParametersValidation) *betapb.DataprocBetaWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidation{}
	p.SetRegex(DataprocBetaWorkflowTemplateParametersValidationRegexToProto(o.Regex))
	p.SetValues(DataprocBetaWorkflowTemplateParametersValidationValuesToProto(o.Values))
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex object to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationRegexToProto(o *beta.WorkflowTemplateParametersValidationRegex) *betapb.DataprocBetaWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidationRegex{}
	sRegexes := make([]string, len(o.Regexes))
	for i, r := range o.Regexes {
		sRegexes[i] = r
	}
	p.SetRegexes(sRegexes)
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues object to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationValuesToProto(o *beta.WorkflowTemplateParametersValidationValues) *betapb.DataprocBetaWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidationValues{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *beta.WorkflowTemplate) *betapb.DataprocBetaWorkflowTemplate {
	p := &betapb.DataprocBetaWorkflowTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyInt64(resource.Version))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetPlacement(DataprocBetaWorkflowTemplatePlacementToProto(resource.Placement))
	p.SetDagTimeout(dcl.ValueOrEmptyString(resource.DagTimeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sJobs := make([]*betapb.DataprocBetaWorkflowTemplateJobs, len(resource.Jobs))
	for i, r := range resource.Jobs {
		sJobs[i] = DataprocBetaWorkflowTemplateJobsToProto(&r)
	}
	p.SetJobs(sJobs)
	sParameters := make([]*betapb.DataprocBetaWorkflowTemplateParameters, len(resource.Parameters))
	for i, r := range resource.Parameters {
		sParameters[i] = DataprocBetaWorkflowTemplateParametersToProto(&r)
	}
	p.SetParameters(sParameters)

	return p
}

// applyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaWorkflowTemplateRequest) (*betapb.DataprocBetaWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// applyDataprocBetaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.ApplyDataprocBetaWorkflowTemplateRequest) (*betapb.DataprocBetaWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.DeleteDataprocBetaWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocBetaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.ListDataprocBetaWorkflowTemplateRequest) (*betapb.ListDataprocBetaWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataprocBetaWorkflowTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
