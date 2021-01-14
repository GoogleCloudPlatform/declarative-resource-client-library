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

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum converts a WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(e dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(n[len("WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplateLabels converts a WorkflowTemplateLabels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateLabels(p *dataprocpb.DataprocWorkflowTemplateLabels) *dataproc.WorkflowTemplateLabels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateLabels{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

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
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToDataprocWorkflowTemplatePlacementManagedClusterLabels(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterLabels converts a WorkflowTemplatePlacementManagedClusterLabels resource from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterLabels(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterLabels) *dataproc.WorkflowTemplatePlacementManagedClusterLabels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterLabels{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
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
	for _, r := range p.GetClusterLabels() {
		obj.ClusterLabels = append(obj.ClusterLabels, *ProtoToDataprocWorkflowTemplatePlacementClusterSelectorClusterLabels(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelectorClusterLabels converts a WorkflowTemplatePlacementClusterSelectorClusterLabels resource from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementClusterSelectorClusterLabels(p *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelectorClusterLabels) *dataproc.WorkflowTemplatePlacementClusterSelectorClusterLabels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementClusterSelectorClusterLabels{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
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
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToDataprocWorkflowTemplateJobsLabels(r))
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
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsHadoopJobProperties(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJobProperties converts a WorkflowTemplateJobsHadoopJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobProperties) *dataproc.WorkflowTemplateJobsHadoopJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfig converts a WorkflowTemplateJobsHadoopJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
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
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsSparkJobProperties(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJobProperties converts a WorkflowTemplateJobsSparkJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJobProperties) *dataproc.WorkflowTemplateJobsSparkJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfig converts a WorkflowTemplateJobsSparkJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
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
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsPysparkJobProperties(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJobProperties converts a WorkflowTemplateJobsPysparkJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobProperties) *dataproc.WorkflowTemplateJobsPysparkJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfig converts a WorkflowTemplateJobsPysparkJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
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
	for _, r := range p.GetScriptVariables() {
		obj.ScriptVariables = append(obj.ScriptVariables, *ProtoToDataprocWorkflowTemplateJobsHiveJobScriptVariables(r))
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsHiveJobProperties(r))
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

// ProtoToWorkflowTemplateJobsHiveJobScriptVariables converts a WorkflowTemplateJobsHiveJobScriptVariables resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJobScriptVariables(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJobScriptVariables) *dataproc.WorkflowTemplateJobsHiveJobScriptVariables {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJobScriptVariables{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobProperties converts a WorkflowTemplateJobsHiveJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJobProperties) *dataproc.WorkflowTemplateJobsHiveJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
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
	for _, r := range p.GetScriptVariables() {
		obj.ScriptVariables = append(obj.ScriptVariables, *ProtoToDataprocWorkflowTemplateJobsPigJobScriptVariables(r))
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsPigJobProperties(r))
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

// ProtoToWorkflowTemplateJobsPigJobScriptVariables converts a WorkflowTemplateJobsPigJobScriptVariables resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobScriptVariables(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobScriptVariables) *dataproc.WorkflowTemplateJobsPigJobScriptVariables {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobScriptVariables{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobProperties converts a WorkflowTemplateJobsPigJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobProperties) *dataproc.WorkflowTemplateJobsPigJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig) *dataproc.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
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
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsSparkRJobProperties(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJobProperties converts a WorkflowTemplateJobsSparkRJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobProperties) *dataproc.WorkflowTemplateJobsSparkRJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfig converts a WorkflowTemplateJobsSparkRJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
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
	for _, r := range p.GetScriptVariables() {
		obj.ScriptVariables = append(obj.ScriptVariables, *ProtoToDataprocWorkflowTemplateJobsSparkSqlJobScriptVariables(r))
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsSparkSqlJobProperties(r))
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

// ProtoToWorkflowTemplateJobsSparkSqlJobScriptVariables converts a WorkflowTemplateJobsSparkSqlJobScriptVariables resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobScriptVariables(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobScriptVariables) *dataproc.WorkflowTemplateJobsSparkSqlJobScriptVariables {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobScriptVariables{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobProperties converts a WorkflowTemplateJobsSparkSqlJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobProperties) *dataproc.WorkflowTemplateJobsSparkSqlJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
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
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDataprocWorkflowTemplateJobsPrestoJobProperties(r))
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

// ProtoToWorkflowTemplateJobsPrestoJobProperties converts a WorkflowTemplateJobsPrestoJobProperties resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobProperties(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobProperties) *dataproc.WorkflowTemplateJobsPrestoJobProperties {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobProperties{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	for _, r := range p.GetDriverLogLevels() {
		obj.DriverLogLevels = append(obj.DriverLogLevels, *ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(r))
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels converts a WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{
		Key:   dcl.StringOrNil(p.Key),
		Value: ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(p.GetValue()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobsLabels converts a WorkflowTemplateJobsLabels resource from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsLabels(p *dataprocpb.DataprocWorkflowTemplateJobsLabels) *dataproc.WorkflowTemplateJobsLabels {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsLabels{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
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
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToDataprocWorkflowTemplateLabels(r))
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum enum to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumToProto(e *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum) dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum_value["WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(0)
}

// WorkflowTemplateLabelsToProto converts a WorkflowTemplateLabels resource to its proto representation.
func DataprocWorkflowTemplateLabelsToProto(o *dataproc.WorkflowTemplateLabels) *dataprocpb.DataprocWorkflowTemplateLabels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateLabels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
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
	for _, r := range o.Labels {
		p.Labels = append(p.Labels, DataprocWorkflowTemplatePlacementManagedClusterLabelsToProto(&r))
	}
	return p
}

// WorkflowTemplatePlacementManagedClusterLabelsToProto converts a WorkflowTemplatePlacementManagedClusterLabels resource to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterLabelsToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterLabels) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterLabels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterLabels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
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
	for _, r := range o.ClusterLabels {
		p.ClusterLabels = append(p.ClusterLabels, DataprocWorkflowTemplatePlacementClusterSelectorClusterLabelsToProto(&r))
	}
	return p
}

// WorkflowTemplatePlacementClusterSelectorClusterLabelsToProto converts a WorkflowTemplatePlacementClusterSelectorClusterLabels resource to its proto representation.
func DataprocWorkflowTemplatePlacementClusterSelectorClusterLabelsToProto(o *dataproc.WorkflowTemplatePlacementClusterSelectorClusterLabels) *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelectorClusterLabels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementClusterSelectorClusterLabels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
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
	for _, r := range o.Labels {
		p.Labels = append(p.Labels, DataprocWorkflowTemplateJobsLabelsToProto(&r))
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
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsHadoopJobPropertiesToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsHadoopJobPropertiesToProto converts a WorkflowTemplateJobsHadoopJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsHadoopJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsHadoopJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsHadoopJobLoggingConfigToProto converts a WorkflowTemplateJobsHadoopJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
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
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsSparkJobPropertiesToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsSparkJobPropertiesToProto converts a WorkflowTemplateJobsSparkJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsSparkJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsSparkJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsSparkJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
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
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsPysparkJobPropertiesToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsPysparkJobPropertiesToProto converts a WorkflowTemplateJobsPysparkJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsPysparkJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsPysparkJobLoggingConfigToProto converts a WorkflowTemplateJobsPysparkJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
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
	for _, r := range o.ScriptVariables {
		p.ScriptVariables = append(p.ScriptVariables, DataprocWorkflowTemplateJobsHiveJobScriptVariablesToProto(&r))
	}
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsHiveJobPropertiesToProto(&r))
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

// WorkflowTemplateJobsHiveJobScriptVariablesToProto converts a WorkflowTemplateJobsHiveJobScriptVariables resource to its proto representation.
func DataprocWorkflowTemplateJobsHiveJobScriptVariablesToProto(o *dataproc.WorkflowTemplateJobsHiveJobScriptVariables) *dataprocpb.DataprocWorkflowTemplateJobsHiveJobScriptVariables {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJobScriptVariables{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsHiveJobPropertiesToProto converts a WorkflowTemplateJobsHiveJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsHiveJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsHiveJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsHiveJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
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
	for _, r := range o.ScriptVariables {
		p.ScriptVariables = append(p.ScriptVariables, DataprocWorkflowTemplateJobsPigJobScriptVariablesToProto(&r))
	}
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsPigJobPropertiesToProto(&r))
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

// WorkflowTemplateJobsPigJobScriptVariablesToProto converts a WorkflowTemplateJobsPigJobScriptVariables resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobScriptVariablesToProto(o *dataproc.WorkflowTemplateJobsPigJobScriptVariables) *dataprocpb.DataprocWorkflowTemplateJobsPigJobScriptVariables {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobScriptVariables{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsPigJobPropertiesToProto converts a WorkflowTemplateJobsPigJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsPigJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsPigJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPigJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
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
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsSparkRJobPropertiesToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsSparkRJobPropertiesToProto converts a WorkflowTemplateJobsSparkRJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsSparkRJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsSparkRJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkRJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
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
	for _, r := range o.ScriptVariables {
		p.ScriptVariables = append(p.ScriptVariables, DataprocWorkflowTemplateJobsSparkSqlJobScriptVariablesToProto(&r))
	}
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsSparkSqlJobPropertiesToProto(&r))
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

// WorkflowTemplateJobsSparkSqlJobScriptVariablesToProto converts a WorkflowTemplateJobsSparkSqlJobScriptVariables resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobScriptVariablesToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobScriptVariables) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobScriptVariables {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobScriptVariables{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobPropertiesToProto converts a WorkflowTemplateJobsSparkSqlJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
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
	for _, r := range o.Properties {
		p.Properties = append(p.Properties, DataprocWorkflowTemplateJobsPrestoJobPropertiesToProto(&r))
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

// WorkflowTemplateJobsPrestoJobPropertiesToProto converts a WorkflowTemplateJobsPrestoJobProperties resource to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobPropertiesToProto(o *dataproc.WorkflowTemplateJobsPrestoJobProperties) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobProperties {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobProperties{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig resource to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig{}
	for _, r := range o.DriverLogLevels {
		p.DriverLogLevels = append(p.DriverLogLevels, DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsToProto(&r))
	}
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels resource to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsToProto(o *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: DataprocWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumToProto(o.Value),
	}
	return p
}

// WorkflowTemplateJobsLabelsToProto converts a WorkflowTemplateJobsLabels resource to its proto representation.
func DataprocWorkflowTemplateJobsLabelsToProto(o *dataproc.WorkflowTemplateJobsLabels) *dataprocpb.DataprocWorkflowTemplateJobsLabels {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsLabels{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
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
	for _, r := range resource.Labels {
		p.Labels = append(p.Labels, DataprocWorkflowTemplateLabelsToProto(&r))
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
