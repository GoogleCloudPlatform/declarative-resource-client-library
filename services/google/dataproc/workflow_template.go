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
package dataproc

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type WorkflowTemplate struct {
	Name       *string                      `json:"name"`
	Version    *int64                       `json:"version"`
	CreateTime *string                      `json:"createTime"`
	UpdateTime *string                      `json:"updateTime"`
	Labels     map[string]string            `json:"labels"`
	Placement  *WorkflowTemplatePlacement   `json:"placement"`
	Jobs       []WorkflowTemplateJobs       `json:"jobs"`
	Parameters []WorkflowTemplateParameters `json:"parameters"`
	Project    *string                      `json:"project"`
	Location   *string                      `json:"location"`
}

func (r *WorkflowTemplate) String() string {
	return dcl.SprintResource(r)
}

type WorkflowTemplatePlacement struct {
	empty           bool                                      `json:"-"`
	ManagedCluster  *WorkflowTemplatePlacementManagedCluster  `json:"managedCluster"`
	ClusterSelector *WorkflowTemplatePlacementClusterSelector `json:"clusterSelector"`
}

// This object is used to assert a desired state where this WorkflowTemplatePlacement is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplatePlacement *WorkflowTemplatePlacement = &WorkflowTemplatePlacement{empty: true}

func (r *WorkflowTemplatePlacement) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplatePlacement) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplatePlacementManagedCluster struct {
	empty       bool                  `json:"-"`
	ClusterName *string               `json:"clusterName"`
	Config      *ClusterClusterConfig `json:"config"`
	Labels      map[string]string     `json:"labels"`
}

// This object is used to assert a desired state where this WorkflowTemplatePlacementManagedCluster is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplatePlacementManagedCluster *WorkflowTemplatePlacementManagedCluster = &WorkflowTemplatePlacementManagedCluster{empty: true}

func (r *WorkflowTemplatePlacementManagedCluster) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplatePlacementManagedCluster) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplatePlacementClusterSelector struct {
	empty         bool              `json:"-"`
	Zone          *string           `json:"zone"`
	ClusterLabels map[string]string `json:"clusterLabels"`
}

// This object is used to assert a desired state where this WorkflowTemplatePlacementClusterSelector is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplatePlacementClusterSelector *WorkflowTemplatePlacementClusterSelector = &WorkflowTemplatePlacementClusterSelector{empty: true}

func (r *WorkflowTemplatePlacementClusterSelector) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplatePlacementClusterSelector) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobs struct {
	empty               bool                             `json:"-"`
	StepId              *string                          `json:"stepId"`
	HadoopJob           *WorkflowTemplateJobsHadoopJob   `json:"hadoopJob"`
	SparkJob            *WorkflowTemplateJobsSparkJob    `json:"sparkJob"`
	PysparkJob          *WorkflowTemplateJobsPysparkJob  `json:"pysparkJob"`
	HiveJob             *WorkflowTemplateJobsHiveJob     `json:"hiveJob"`
	PigJob              *WorkflowTemplateJobsPigJob      `json:"pigJob"`
	SparkRJob           *WorkflowTemplateJobsSparkRJob   `json:"sparkRJob"`
	SparkSqlJob         *WorkflowTemplateJobsSparkSqlJob `json:"sparkSqlJob"`
	PrestoJob           *WorkflowTemplateJobsPrestoJob   `json:"prestoJob"`
	Labels              map[string]string                `json:"labels"`
	Scheduling          *WorkflowTemplateJobsScheduling  `json:"scheduling"`
	PrerequisiteStepIds []string                         `json:"prerequisiteStepIds"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobs *WorkflowTemplateJobs = &WorkflowTemplateJobs{empty: true}

func (r *WorkflowTemplateJobs) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHadoopJob struct {
	empty          bool                                        `json:"-"`
	MainJarFileUri *string                                     `json:"mainJarFileUri"`
	MainClass      *string                                     `json:"mainClass"`
	Args           []string                                    `json:"args"`
	JarFileUris    []string                                    `json:"jarFileUris"`
	FileUris       []string                                    `json:"fileUris"`
	ArchiveUris    []string                                    `json:"archiveUris"`
	Properties     map[string]string                           `json:"properties"`
	LoggingConfig  *WorkflowTemplateJobsHadoopJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHadoopJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHadoopJob *WorkflowTemplateJobsHadoopJob = &WorkflowTemplateJobsHadoopJob{empty: true}

func (r *WorkflowTemplateJobsHadoopJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHadoopJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHadoopJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHadoopJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHadoopJobLoggingConfig *WorkflowTemplateJobsHadoopJobLoggingConfig = &WorkflowTemplateJobsHadoopJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsHadoopJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHadoopJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkJob struct {
	empty          bool                                       `json:"-"`
	MainJarFileUri *string                                    `json:"mainJarFileUri"`
	MainClass      *string                                    `json:"mainClass"`
	Args           []string                                   `json:"args"`
	JarFileUris    []string                                   `json:"jarFileUris"`
	FileUris       []string                                   `json:"fileUris"`
	ArchiveUris    []string                                   `json:"archiveUris"`
	Properties     map[string]string                          `json:"properties"`
	LoggingConfig  *WorkflowTemplateJobsSparkJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkJob *WorkflowTemplateJobsSparkJob = &WorkflowTemplateJobsSparkJob{empty: true}

func (r *WorkflowTemplateJobsSparkJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkJobLoggingConfig *WorkflowTemplateJobsSparkJobLoggingConfig = &WorkflowTemplateJobsSparkJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsSparkJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPysparkJob struct {
	empty             bool                                         `json:"-"`
	MainPythonFileUri *string                                      `json:"mainPythonFileUri"`
	Args              []string                                     `json:"args"`
	PythonFileUris    []string                                     `json:"pythonFileUris"`
	JarFileUris       []string                                     `json:"jarFileUris"`
	FileUris          []string                                     `json:"fileUris"`
	ArchiveUris       []string                                     `json:"archiveUris"`
	Properties        map[string]string                            `json:"properties"`
	LoggingConfig     *WorkflowTemplateJobsPysparkJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPysparkJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPysparkJob *WorkflowTemplateJobsPysparkJob = &WorkflowTemplateJobsPysparkJob{empty: true}

func (r *WorkflowTemplateJobsPysparkJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPysparkJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPysparkJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPysparkJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPysparkJobLoggingConfig *WorkflowTemplateJobsPysparkJobLoggingConfig = &WorkflowTemplateJobsPysparkJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsPysparkJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPysparkJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHiveJob struct {
	empty             bool                                  `json:"-"`
	QueryFileUri      *string                               `json:"queryFileUri"`
	QueryList         *WorkflowTemplateJobsHiveJobQueryList `json:"queryList"`
	ContinueOnFailure *bool                                 `json:"continueOnFailure"`
	ScriptVariables   map[string]string                     `json:"scriptVariables"`
	Properties        map[string]string                     `json:"properties"`
	JarFileUris       []string                              `json:"jarFileUris"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHiveJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHiveJob *WorkflowTemplateJobsHiveJob = &WorkflowTemplateJobsHiveJob{empty: true}

func (r *WorkflowTemplateJobsHiveJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHiveJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHiveJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHiveJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHiveJobQueryList *WorkflowTemplateJobsHiveJobQueryList = &WorkflowTemplateJobsHiveJobQueryList{empty: true}

func (r *WorkflowTemplateJobsHiveJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHiveJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPigJob struct {
	empty             bool                                     `json:"-"`
	QueryFileUri      *string                                  `json:"queryFileUri"`
	QueryList         *WorkflowTemplateJobsPigJobQueryList     `json:"queryList"`
	ContinueOnFailure *bool                                    `json:"continueOnFailure"`
	ScriptVariables   map[string]string                        `json:"scriptVariables"`
	Properties        map[string]string                        `json:"properties"`
	JarFileUris       []string                                 `json:"jarFileUris"`
	LoggingConfig     *WorkflowTemplateJobsPigJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPigJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPigJob *WorkflowTemplateJobsPigJob = &WorkflowTemplateJobsPigJob{empty: true}

func (r *WorkflowTemplateJobsPigJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPigJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPigJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPigJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPigJobQueryList *WorkflowTemplateJobsPigJobQueryList = &WorkflowTemplateJobsPigJobQueryList{empty: true}

func (r *WorkflowTemplateJobsPigJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPigJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPigJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPigJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPigJobLoggingConfig *WorkflowTemplateJobsPigJobLoggingConfig = &WorkflowTemplateJobsPigJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsPigJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPigJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkRJob struct {
	empty         bool                                        `json:"-"`
	MainRFileUri  *string                                     `json:"mainRFileUri"`
	Args          []string                                    `json:"args"`
	FileUris      []string                                    `json:"fileUris"`
	ArchiveUris   []string                                    `json:"archiveUris"`
	Properties    map[string]string                           `json:"properties"`
	LoggingConfig *WorkflowTemplateJobsSparkRJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkRJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkRJob *WorkflowTemplateJobsSparkRJob = &WorkflowTemplateJobsSparkRJob{empty: true}

func (r *WorkflowTemplateJobsSparkRJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkRJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkRJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkRJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkRJobLoggingConfig *WorkflowTemplateJobsSparkRJobLoggingConfig = &WorkflowTemplateJobsSparkRJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsSparkRJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkRJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkSqlJob struct {
	empty           bool                                          `json:"-"`
	QueryFileUri    *string                                       `json:"queryFileUri"`
	QueryList       *WorkflowTemplateJobsSparkSqlJobQueryList     `json:"queryList"`
	ScriptVariables map[string]string                             `json:"scriptVariables"`
	Properties      map[string]string                             `json:"properties"`
	JarFileUris     []string                                      `json:"jarFileUris"`
	LoggingConfig   *WorkflowTemplateJobsSparkSqlJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkSqlJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkSqlJob *WorkflowTemplateJobsSparkSqlJob = &WorkflowTemplateJobsSparkSqlJob{empty: true}

func (r *WorkflowTemplateJobsSparkSqlJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkSqlJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkSqlJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkSqlJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkSqlJobQueryList *WorkflowTemplateJobsSparkSqlJobQueryList = &WorkflowTemplateJobsSparkSqlJobQueryList{empty: true}

func (r *WorkflowTemplateJobsSparkSqlJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkSqlJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkSqlJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkSqlJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkSqlJobLoggingConfig *WorkflowTemplateJobsSparkSqlJobLoggingConfig = &WorkflowTemplateJobsSparkSqlJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPrestoJob struct {
	empty             bool                                        `json:"-"`
	QueryFileUri      *string                                     `json:"queryFileUri"`
	QueryList         *WorkflowTemplateJobsPrestoJobQueryList     `json:"queryList"`
	ContinueOnFailure *bool                                       `json:"continueOnFailure"`
	OutputFormat      *string                                     `json:"outputFormat"`
	ClientTags        []string                                    `json:"clientTags"`
	Properties        map[string]string                           `json:"properties"`
	LoggingConfig     *WorkflowTemplateJobsPrestoJobLoggingConfig `json:"loggingConfig"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPrestoJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPrestoJob *WorkflowTemplateJobsPrestoJob = &WorkflowTemplateJobsPrestoJob{empty: true}

func (r *WorkflowTemplateJobsPrestoJob) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPrestoJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPrestoJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPrestoJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPrestoJobQueryList *WorkflowTemplateJobsPrestoJobQueryList = &WorkflowTemplateJobsPrestoJobQueryList{empty: true}

func (r *WorkflowTemplateJobsPrestoJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPrestoJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPrestoJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPrestoJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPrestoJobLoggingConfig *WorkflowTemplateJobsPrestoJobLoggingConfig = &WorkflowTemplateJobsPrestoJobLoggingConfig{empty: true}

func (r *WorkflowTemplateJobsPrestoJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPrestoJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsScheduling struct {
	empty              bool   `json:"-"`
	MaxFailuresPerHour *int64 `json:"maxFailuresPerHour"`
	MaxFailuresTotal   *int64 `json:"maxFailuresTotal"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsScheduling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsScheduling *WorkflowTemplateJobsScheduling = &WorkflowTemplateJobsScheduling{empty: true}

func (r *WorkflowTemplateJobsScheduling) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsScheduling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateParameters struct {
	empty       bool                                  `json:"-"`
	Name        *string                               `json:"name"`
	Fields      []string                              `json:"fields"`
	Description *string                               `json:"description"`
	Validation  *WorkflowTemplateParametersValidation `json:"validation"`
}

// This object is used to assert a desired state where this WorkflowTemplateParameters is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateParameters *WorkflowTemplateParameters = &WorkflowTemplateParameters{empty: true}

func (r *WorkflowTemplateParameters) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateParameters) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateParametersValidation struct {
	empty  bool                                        `json:"-"`
	Regex  *WorkflowTemplateParametersValidationRegex  `json:"regex"`
	Values *WorkflowTemplateParametersValidationValues `json:"values"`
}

// This object is used to assert a desired state where this WorkflowTemplateParametersValidation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateParametersValidation *WorkflowTemplateParametersValidation = &WorkflowTemplateParametersValidation{empty: true}

func (r *WorkflowTemplateParametersValidation) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateParametersValidation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateParametersValidationRegex struct {
	empty   bool     `json:"-"`
	Regexes []string `json:"regexes"`
}

// This object is used to assert a desired state where this WorkflowTemplateParametersValidationRegex is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateParametersValidationRegex *WorkflowTemplateParametersValidationRegex = &WorkflowTemplateParametersValidationRegex{empty: true}

func (r *WorkflowTemplateParametersValidationRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateParametersValidationRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateParametersValidationValues struct {
	empty  bool     `json:"-"`
	Values []string `json:"values"`
}

// This object is used to assert a desired state where this WorkflowTemplateParametersValidationValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateParametersValidationValues *WorkflowTemplateParametersValidationValues = &WorkflowTemplateParametersValidationValues{empty: true}

func (r *WorkflowTemplateParametersValidationValues) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateParametersValidationValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *WorkflowTemplate) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dataproc",
		Type:    "WorkflowTemplate",
		Version: "dataproc",
	}
}

const WorkflowTemplateMaxPage = -1

type WorkflowTemplateList struct {
	Items []*WorkflowTemplate

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *WorkflowTemplateList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkflowTemplateList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkflowTemplate(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkflowTemplate(ctx context.Context, project, location string) (*WorkflowTemplateList, error) {

	return c.ListWorkflowTemplateWithMaxResults(ctx, project, location, WorkflowTemplateMaxPage)

}

func (c *Client) ListWorkflowTemplateWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*WorkflowTemplateList, error) {
	items, token, err := c.listWorkflowTemplate(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkflowTemplateList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetWorkflowTemplate(ctx context.Context, r *WorkflowTemplate) (*WorkflowTemplate, error) {
	b, err := c.getWorkflowTemplateRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkflowTemplate(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkflowTemplateNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteWorkflowTemplate(ctx context.Context, r *WorkflowTemplate) error {
	if r == nil {
		return fmt.Errorf("WorkflowTemplate resource is nil")
	}
	c.Config.Logger.Info("Deleting WorkflowTemplate...")
	deleteOp := deleteWorkflowTemplateOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkflowTemplate deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkflowTemplate(ctx context.Context, project, location string, filter func(*WorkflowTemplate) bool) error {
	listObj, err := c.ListWorkflowTemplate(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkflowTemplate(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkflowTemplate(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkflowTemplate(ctx context.Context, rawDesired *WorkflowTemplate, opts ...dcl.ApplyOption) (*WorkflowTemplate, error) {
	c.Config.Logger.Info("Beginning ApplyWorkflowTemplate...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.workflowTemplateDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []workflowTemplateApiOperation
	if create {
		ops = append(ops, &createWorkflowTemplateOperation{})
	} else if recreate {

		ops = append(ops, &deleteWorkflowTemplateOperation{})

		ops = append(ops, &createWorkflowTemplateOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeWorkflowTemplateDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetWorkflowTemplate(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkflowTemplateNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkflowTemplateDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkflowTemplate(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
