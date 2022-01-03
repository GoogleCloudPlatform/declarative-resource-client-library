// Copyright 2022 Google LLC. All Rights Reserved.
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

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement object from its proto representation.
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

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedCluster(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster) *alpha.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.GetClusterName()),
		Config:      ProtoToDataprocAlphaClusterClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementClusterSelector(p *alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector) *alpha.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobs(p *alphapb.DataprocAlphaWorkflowTemplateJobs) *alpha.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.GetStepId()),
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

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob) *alpha.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
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

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfig converts a WorkflowTemplateJobsHadoopJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig) *alpha.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob) *alpha.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
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

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfig converts a WorkflowTemplateJobsSparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob) *alpha.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.GetMainPythonFileUri()),
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

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfig converts a WorkflowTemplateJobsPysparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig) *alpha.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHiveJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob) *alpha.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJob) *alpha.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig) *alpha.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob) *alpha.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.GetMainRFileUri()),
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

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfig converts a WorkflowTemplateJobsSparkRJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob) *alpha.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:     ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob) *alpha.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		OutputFormat:      dcl.StringOrNil(p.GetOutputFormat()),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig) *alpha.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsScheduling(p *alphapb.DataprocAlphaWorkflowTemplateJobsScheduling) *alpha.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.GetMaxFailuresPerHour()),
		MaxFailuresTotal:   dcl.Int64OrNil(p.GetMaxFailuresTotal()),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParameters(p *alphapb.DataprocAlphaWorkflowTemplateParameters) *alpha.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Validation:  ProtoToDataprocAlphaWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation object from its proto representation.
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

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex object from its proto representation.
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

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues object from its proto representation.
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
		Name:       dcl.StringOrNil(p.GetName()),
		Version:    dcl.Int64OrNil(p.GetVersion()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocAlphaWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.GetDagTimeout()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocAlphaWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocAlphaWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementToProto(o *alpha.WorkflowTemplatePlacement) *alphapb.DataprocAlphaWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacement{}
	p.SetManagedCluster(DataprocAlphaWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster))
	p.SetClusterSelector(DataprocAlphaWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector))
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterToProto(o *alpha.WorkflowTemplatePlacementManagedCluster) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster{}
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetConfig(DataprocAlphaClusterClusterConfigToProto(o.Config))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementClusterSelectorToProto(o *alpha.WorkflowTemplatePlacementClusterSelector) *alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	mClusterLabels := make(map[string]string, len(o.ClusterLabels))
	for k, r := range o.ClusterLabels {
		mClusterLabels[k] = r
	}
	p.SetClusterLabels(mClusterLabels)
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsToProto(o *alpha.WorkflowTemplateJobs) *alphapb.DataprocAlphaWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobs{}
	p.SetStepId(dcl.ValueOrEmptyString(o.StepId))
	p.SetHadoopJob(DataprocAlphaWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob))
	p.SetSparkJob(DataprocAlphaWorkflowTemplateJobsSparkJobToProto(o.SparkJob))
	p.SetPysparkJob(DataprocAlphaWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob))
	p.SetHiveJob(DataprocAlphaWorkflowTemplateJobsHiveJobToProto(o.HiveJob))
	p.SetPigJob(DataprocAlphaWorkflowTemplateJobsPigJobToProto(o.PigJob))
	p.SetSparkRJob(DataprocAlphaWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob))
	p.SetSparkSqlJob(DataprocAlphaWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob))
	p.SetPrestoJob(DataprocAlphaWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob))
	p.SetScheduling(DataprocAlphaWorkflowTemplateJobsSchedulingToProto(o.Scheduling))
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
func DataprocAlphaWorkflowTemplateJobsHadoopJobToProto(o *alpha.WorkflowTemplateJobsHadoopJob) *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsHadoopJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkJobToProto(o *alpha.WorkflowTemplateJobsSparkJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPysparkJobToProto(o *alpha.WorkflowTemplateJobsPysparkJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob{}
	p.SetMainPythonFileUri(dcl.ValueOrEmptyString(o.MainPythonFileUri))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPysparkJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsHiveJobToProto(o *alpha.WorkflowTemplateJobsHiveJob) *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList))
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
func DataprocAlphaWorkflowTemplateJobsHiveJobQueryListToProto(o *alpha.WorkflowTemplateJobsHiveJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobToProto(o *alpha.WorkflowTemplateJobsPigJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsPigJobQueryListToProto(o *alpha.WorkflowTemplateJobsPigJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPigJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkRJobToProto(o *alpha.WorkflowTemplateJobsSparkRJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob{}
	p.SetMainRFileUri(dcl.ValueOrEmptyString(o.MainRFileUri))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkRJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobToProto(o *alpha.WorkflowTemplateJobsSparkSqlJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *alpha.WorkflowTemplateJobsSparkSqlJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobToProto(o *alpha.WorkflowTemplateJobsPrestoJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetOutputFormat(dcl.ValueOrEmptyString(o.OutputFormat))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocAlphaWorkflowTemplateJobsPrestoJobQueryListToProto(o *alpha.WorkflowTemplateJobsPrestoJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPrestoJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSchedulingToProto(o *alpha.WorkflowTemplateJobsScheduling) *alphapb.DataprocAlphaWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsScheduling{}
	p.SetMaxFailuresPerHour(dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour))
	p.SetMaxFailuresTotal(dcl.ValueOrEmptyInt64(o.MaxFailuresTotal))
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersToProto(o *alpha.WorkflowTemplateParameters) *alphapb.DataprocAlphaWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParameters{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetValidation(DataprocAlphaWorkflowTemplateParametersValidationToProto(o.Validation))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationToProto(o *alpha.WorkflowTemplateParametersValidation) *alphapb.DataprocAlphaWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidation{}
	p.SetRegex(DataprocAlphaWorkflowTemplateParametersValidationRegexToProto(o.Regex))
	p.SetValues(DataprocAlphaWorkflowTemplateParametersValidationValuesToProto(o.Values))
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationRegexToProto(o *alpha.WorkflowTemplateParametersValidationRegex) *alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex{}
	sRegexes := make([]string, len(o.Regexes))
	for i, r := range o.Regexes {
		sRegexes[i] = r
	}
	p.SetRegexes(sRegexes)
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationValuesToProto(o *alpha.WorkflowTemplateParametersValidationValues) *alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *alpha.WorkflowTemplate) *alphapb.DataprocAlphaWorkflowTemplate {
	p := &alphapb.DataprocAlphaWorkflowTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyInt64(resource.Version))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetPlacement(DataprocAlphaWorkflowTemplatePlacementToProto(resource.Placement))
	p.SetDagTimeout(dcl.ValueOrEmptyString(resource.DagTimeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sJobs := make([]*alphapb.DataprocAlphaWorkflowTemplateJobs, len(resource.Jobs))
	for i, r := range resource.Jobs {
		sJobs[i] = DataprocAlphaWorkflowTemplateJobsToProto(&r)
	}
	p.SetJobs(sJobs)
	sParameters := make([]*alphapb.DataprocAlphaWorkflowTemplateParameters, len(resource.Parameters))
	for i, r := range resource.Parameters {
		sParameters[i] = DataprocAlphaWorkflowTemplateParametersToProto(&r)
	}
	p.SetParameters(sParameters)

	return p
}

// applyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaWorkflowTemplateRequest) (*alphapb.DataprocAlphaWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// applyDataprocAlphaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.ApplyDataprocAlphaWorkflowTemplateRequest) (*alphapb.DataprocAlphaWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.DeleteDataprocAlphaWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocAlphaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.ListDataprocAlphaWorkflowTemplateRequest) (*alphapb.ListDataprocAlphaWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataprocAlphaWorkflowTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
