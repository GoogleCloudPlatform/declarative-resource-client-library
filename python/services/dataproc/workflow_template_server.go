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

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement object from its proto representation.
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

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedCluster(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster) *dataproc.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.GetClusterName()),
		Config:      ProtoToDataprocClusterClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementClusterSelector(p *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector) *dataproc.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobs(p *dataprocpb.DataprocWorkflowTemplateJobs) *dataproc.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.GetStepId()),
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

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJob(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJob) *dataproc.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
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

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfig converts a WorkflowTemplateJobsHadoopJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJob) *dataproc.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
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

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfig converts a WorkflowTemplateJobsSparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJob(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJob) *dataproc.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.GetMainPythonFileUri()),
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

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfig converts a WorkflowTemplateJobsPysparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJob(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJob) *dataproc.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJob(p *dataprocpb.DataprocWorkflowTemplateJobsPigJob) *dataproc.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig) *dataproc.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJob) *dataproc.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.GetMainRFileUri()),
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

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfig converts a WorkflowTemplateJobsSparkRJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob) *dataproc.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:     ProtoToDataprocWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJob(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJob) *dataproc.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		OutputFormat:      dcl.StringOrNil(p.GetOutputFormat()),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList object from its proto representation.
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

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsScheduling(p *dataprocpb.DataprocWorkflowTemplateJobsScheduling) *dataproc.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.GetMaxFailuresPerHour()),
		MaxFailuresTotal:   dcl.Int64OrNil(p.GetMaxFailuresTotal()),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters object from its proto representation.
func ProtoToDataprocWorkflowTemplateParameters(p *dataprocpb.DataprocWorkflowTemplateParameters) *dataproc.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Validation:  ProtoToDataprocWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation object from its proto representation.
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

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex object from its proto representation.
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

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues object from its proto representation.
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
		Name:       dcl.StringOrNil(p.GetName()),
		Version:    dcl.Int64OrNil(p.GetVersion()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.GetDagTimeout()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement object to its proto representation.
func DataprocWorkflowTemplatePlacementToProto(o *dataproc.WorkflowTemplatePlacement) *dataprocpb.DataprocWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacement{}
	p.SetManagedCluster(DataprocWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster))
	p.SetClusterSelector(DataprocWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector))
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterToProto(o *dataproc.WorkflowTemplatePlacementManagedCluster) *dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster{}
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetConfig(DataprocClusterClusterConfigToProto(o.Config))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector object to its proto representation.
func DataprocWorkflowTemplatePlacementClusterSelectorToProto(o *dataproc.WorkflowTemplatePlacementClusterSelector) *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	mClusterLabels := make(map[string]string, len(o.ClusterLabels))
	for k, r := range o.ClusterLabels {
		mClusterLabels[k] = r
	}
	p.SetClusterLabels(mClusterLabels)
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs object to its proto representation.
func DataprocWorkflowTemplateJobsToProto(o *dataproc.WorkflowTemplateJobs) *dataprocpb.DataprocWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobs{}
	p.SetStepId(dcl.ValueOrEmptyString(o.StepId))
	p.SetHadoopJob(DataprocWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob))
	p.SetSparkJob(DataprocWorkflowTemplateJobsSparkJobToProto(o.SparkJob))
	p.SetPysparkJob(DataprocWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob))
	p.SetHiveJob(DataprocWorkflowTemplateJobsHiveJobToProto(o.HiveJob))
	p.SetPigJob(DataprocWorkflowTemplateJobsPigJobToProto(o.PigJob))
	p.SetSparkRJob(DataprocWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob))
	p.SetSparkSqlJob(DataprocWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob))
	p.SetPrestoJob(DataprocWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob))
	p.SetScheduling(DataprocWorkflowTemplateJobsSchedulingToProto(o.Scheduling))
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
func DataprocWorkflowTemplateJobsHadoopJobToProto(o *dataproc.WorkflowTemplateJobsHadoopJob) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob object to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobToProto(o *dataproc.WorkflowTemplateJobsSparkJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob object to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobToProto(o *dataproc.WorkflowTemplateJobsPysparkJob) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJob{}
	p.SetMainPythonFileUri(dcl.ValueOrEmptyString(o.MainPythonFileUri))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob object to its proto representation.
func DataprocWorkflowTemplateJobsHiveJobToProto(o *dataproc.WorkflowTemplateJobsHiveJob) *dataprocpb.DataprocWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList))
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
func DataprocWorkflowTemplateJobsHiveJobQueryListToProto(o *dataproc.WorkflowTemplateJobsHiveJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob object to its proto representation.
func DataprocWorkflowTemplateJobsPigJobToProto(o *dataproc.WorkflowTemplateJobsPigJob) *dataprocpb.DataprocWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsPigJobQueryListToProto(o *dataproc.WorkflowTemplateJobsPigJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig object to its proto representation.
func DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPigJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob object to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobToProto(o *dataproc.WorkflowTemplateJobsSparkRJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJob{}
	p.SetMainRFileUri(dcl.ValueOrEmptyString(o.MainRFileUri))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob object to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob object to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobToProto(o *dataproc.WorkflowTemplateJobsPrestoJob) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetOutputFormat(dcl.ValueOrEmptyString(o.OutputFormat))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsPrestoJobQueryListToProto(o *dataproc.WorkflowTemplateJobsPrestoJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig object to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling object to its proto representation.
func DataprocWorkflowTemplateJobsSchedulingToProto(o *dataproc.WorkflowTemplateJobsScheduling) *dataprocpb.DataprocWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsScheduling{}
	p.SetMaxFailuresPerHour(dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour))
	p.SetMaxFailuresTotal(dcl.ValueOrEmptyInt64(o.MaxFailuresTotal))
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters object to its proto representation.
func DataprocWorkflowTemplateParametersToProto(o *dataproc.WorkflowTemplateParameters) *dataprocpb.DataprocWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParameters{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetValidation(DataprocWorkflowTemplateParametersValidationToProto(o.Validation))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation object to its proto representation.
func DataprocWorkflowTemplateParametersValidationToProto(o *dataproc.WorkflowTemplateParametersValidation) *dataprocpb.DataprocWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidation{}
	p.SetRegex(DataprocWorkflowTemplateParametersValidationRegexToProto(o.Regex))
	p.SetValues(DataprocWorkflowTemplateParametersValidationValuesToProto(o.Values))
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex object to its proto representation.
func DataprocWorkflowTemplateParametersValidationRegexToProto(o *dataproc.WorkflowTemplateParametersValidationRegex) *dataprocpb.DataprocWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidationRegex{}
	sRegexes := make([]string, len(o.Regexes))
	for i, r := range o.Regexes {
		sRegexes[i] = r
	}
	p.SetRegexes(sRegexes)
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues object to its proto representation.
func DataprocWorkflowTemplateParametersValidationValuesToProto(o *dataproc.WorkflowTemplateParametersValidationValues) *dataprocpb.DataprocWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidationValues{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *dataproc.WorkflowTemplate) *dataprocpb.DataprocWorkflowTemplate {
	p := &dataprocpb.DataprocWorkflowTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyInt64(resource.Version))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetPlacement(DataprocWorkflowTemplatePlacementToProto(resource.Placement))
	p.SetDagTimeout(dcl.ValueOrEmptyString(resource.DagTimeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sJobs := make([]*dataprocpb.DataprocWorkflowTemplateJobs, len(resource.Jobs))
	for i, r := range resource.Jobs {
		sJobs[i] = DataprocWorkflowTemplateJobsToProto(&r)
	}
	p.SetJobs(sJobs)
	sParameters := make([]*dataprocpb.DataprocWorkflowTemplateParameters, len(resource.Parameters))
	for i, r := range resource.Parameters {
		sParameters[i] = DataprocWorkflowTemplateParametersToProto(&r)
	}
	p.SetParameters(sParameters)

	return p
}

// applyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocWorkflowTemplateRequest) (*dataprocpb.DataprocWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// applyDataprocWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.ApplyDataprocWorkflowTemplateRequest) (*dataprocpb.DataprocWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.DeleteDataprocWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.ListDataprocWorkflowTemplateRequest) (*dataprocpb.ListDataprocWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &dataprocpb.ListDataprocWorkflowTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
