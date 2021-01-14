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
	Labels     []WorkflowTemplateLabels     `json:"labels"`
	Placement  *WorkflowTemplatePlacement   `json:"placement"`
	Jobs       []WorkflowTemplateJobs       `json:"jobs"`
	Parameters []WorkflowTemplateParameters `json:"parameters"`
	Project    *string                      `json:"project"`
	Location   *string                      `json:"location"`
}

func (r *WorkflowTemplate) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum.
type WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum string

// WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumRef returns a *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnumRef(s string) *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum {
	if s == "" {
		return nil
	}

	v := WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum(s)
	return &v
}

func (v WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum) Validate() error {
	for _, s := range []string{"NULL_VALUE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type WorkflowTemplateLabels struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateLabels *WorkflowTemplateLabels = &WorkflowTemplateLabels{empty: true}

func (r *WorkflowTemplateLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
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
	empty       bool                                            `json:"-"`
	ClusterName *string                                         `json:"clusterName"`
	Config      *ClusterClusterConfig                           `json:"config"`
	Labels      []WorkflowTemplatePlacementManagedClusterLabels `json:"labels"`
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

type WorkflowTemplatePlacementManagedClusterLabels struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplatePlacementManagedClusterLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplatePlacementManagedClusterLabels *WorkflowTemplatePlacementManagedClusterLabels = &WorkflowTemplatePlacementManagedClusterLabels{empty: true}

func (r *WorkflowTemplatePlacementManagedClusterLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplatePlacementManagedClusterLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplatePlacementClusterSelector struct {
	empty         bool                                                    `json:"-"`
	Zone          *string                                                 `json:"zone"`
	ClusterLabels []WorkflowTemplatePlacementClusterSelectorClusterLabels `json:"clusterLabels"`
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

type WorkflowTemplatePlacementClusterSelectorClusterLabels struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplatePlacementClusterSelectorClusterLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplatePlacementClusterSelectorClusterLabels *WorkflowTemplatePlacementClusterSelectorClusterLabels = &WorkflowTemplatePlacementClusterSelectorClusterLabels{empty: true}

func (r *WorkflowTemplatePlacementClusterSelectorClusterLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplatePlacementClusterSelectorClusterLabels) HashCode() string {
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
	Labels              []WorkflowTemplateJobsLabels     `json:"labels"`
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
	Properties     []WorkflowTemplateJobsHadoopJobProperties   `json:"properties"`
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

type WorkflowTemplateJobsHadoopJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHadoopJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHadoopJobProperties *WorkflowTemplateJobsHadoopJobProperties = &WorkflowTemplateJobsHadoopJobProperties{empty: true}

func (r *WorkflowTemplateJobsHadoopJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHadoopJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHadoopJobLoggingConfig struct {
	empty           bool                                                        `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels struct {
	empty bool                                                                `json:"-"`
	Key   *string                                                             `json:"key"`
	Value *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHadoopJobLoggingConfigDriverLogLevels) HashCode() string {
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
	Properties     []WorkflowTemplateJobsSparkJobProperties   `json:"properties"`
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

type WorkflowTemplateJobsSparkJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkJobProperties *WorkflowTemplateJobsSparkJobProperties = &WorkflowTemplateJobsSparkJobProperties{empty: true}

func (r *WorkflowTemplateJobsSparkJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkJobLoggingConfig struct {
	empty           bool                                                       `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels struct {
	empty bool                                                               `json:"-"`
	Key   *string                                                            `json:"key"`
	Value *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkJobLoggingConfigDriverLogLevels) HashCode() string {
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
	Properties        []WorkflowTemplateJobsPysparkJobProperties   `json:"properties"`
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

type WorkflowTemplateJobsPysparkJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPysparkJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPysparkJobProperties *WorkflowTemplateJobsPysparkJobProperties = &WorkflowTemplateJobsPysparkJobProperties{empty: true}

func (r *WorkflowTemplateJobsPysparkJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPysparkJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPysparkJobLoggingConfig struct {
	empty           bool                                                         `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels struct {
	empty bool                                                                 `json:"-"`
	Key   *string                                                              `json:"key"`
	Value *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPysparkJobLoggingConfigDriverLogLevels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHiveJob struct {
	empty             bool                                         `json:"-"`
	QueryFileUri      *string                                      `json:"queryFileUri"`
	QueryList         *WorkflowTemplateJobsHiveJobQueryList        `json:"queryList"`
	ContinueOnFailure *bool                                        `json:"continueOnFailure"`
	ScriptVariables   []WorkflowTemplateJobsHiveJobScriptVariables `json:"scriptVariables"`
	Properties        []WorkflowTemplateJobsHiveJobProperties      `json:"properties"`
	JarFileUris       []string                                     `json:"jarFileUris"`
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

type WorkflowTemplateJobsHiveJobScriptVariables struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHiveJobScriptVariables is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHiveJobScriptVariables *WorkflowTemplateJobsHiveJobScriptVariables = &WorkflowTemplateJobsHiveJobScriptVariables{empty: true}

func (r *WorkflowTemplateJobsHiveJobScriptVariables) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHiveJobScriptVariables) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsHiveJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsHiveJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsHiveJobProperties *WorkflowTemplateJobsHiveJobProperties = &WorkflowTemplateJobsHiveJobProperties{empty: true}

func (r *WorkflowTemplateJobsHiveJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsHiveJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPigJob struct {
	empty             bool                                        `json:"-"`
	QueryFileUri      *string                                     `json:"queryFileUri"`
	QueryList         *WorkflowTemplateJobsPigJobQueryList        `json:"queryList"`
	ContinueOnFailure *bool                                       `json:"continueOnFailure"`
	ScriptVariables   []WorkflowTemplateJobsPigJobScriptVariables `json:"scriptVariables"`
	Properties        []WorkflowTemplateJobsPigJobProperties      `json:"properties"`
	JarFileUris       []string                                    `json:"jarFileUris"`
	LoggingConfig     *WorkflowTemplateJobsPigJobLoggingConfig    `json:"loggingConfig"`
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

type WorkflowTemplateJobsPigJobScriptVariables struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPigJobScriptVariables is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPigJobScriptVariables *WorkflowTemplateJobsPigJobScriptVariables = &WorkflowTemplateJobsPigJobScriptVariables{empty: true}

func (r *WorkflowTemplateJobsPigJobScriptVariables) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPigJobScriptVariables) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPigJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPigJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPigJobProperties *WorkflowTemplateJobsPigJobProperties = &WorkflowTemplateJobsPigJobProperties{empty: true}

func (r *WorkflowTemplateJobsPigJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPigJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPigJobLoggingConfig struct {
	empty           bool                                                     `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels struct {
	empty bool                                                             `json:"-"`
	Key   *string                                                          `json:"key"`
	Value *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPigJobLoggingConfigDriverLogLevels) HashCode() string {
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
	Properties    []WorkflowTemplateJobsSparkRJobProperties   `json:"properties"`
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

type WorkflowTemplateJobsSparkRJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkRJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkRJobProperties *WorkflowTemplateJobsSparkRJobProperties = &WorkflowTemplateJobsSparkRJobProperties{empty: true}

func (r *WorkflowTemplateJobsSparkRJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkRJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkRJobLoggingConfig struct {
	empty           bool                                                        `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels struct {
	empty bool                                                                `json:"-"`
	Key   *string                                                             `json:"key"`
	Value *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkRJobLoggingConfigDriverLogLevels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkSqlJob struct {
	empty           bool                                             `json:"-"`
	QueryFileUri    *string                                          `json:"queryFileUri"`
	QueryList       *WorkflowTemplateJobsSparkSqlJobQueryList        `json:"queryList"`
	ScriptVariables []WorkflowTemplateJobsSparkSqlJobScriptVariables `json:"scriptVariables"`
	Properties      []WorkflowTemplateJobsSparkSqlJobProperties      `json:"properties"`
	JarFileUris     []string                                         `json:"jarFileUris"`
	LoggingConfig   *WorkflowTemplateJobsSparkSqlJobLoggingConfig    `json:"loggingConfig"`
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

type WorkflowTemplateJobsSparkSqlJobScriptVariables struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkSqlJobScriptVariables is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkSqlJobScriptVariables *WorkflowTemplateJobsSparkSqlJobScriptVariables = &WorkflowTemplateJobsSparkSqlJobScriptVariables{empty: true}

func (r *WorkflowTemplateJobsSparkSqlJobScriptVariables) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkSqlJobScriptVariables) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkSqlJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkSqlJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkSqlJobProperties *WorkflowTemplateJobsSparkSqlJobProperties = &WorkflowTemplateJobsSparkSqlJobProperties{empty: true}

func (r *WorkflowTemplateJobsSparkSqlJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkSqlJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsSparkSqlJobLoggingConfig struct {
	empty           bool                                                          `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels struct {
	empty bool                                                                  `json:"-"`
	Key   *string                                                               `json:"key"`
	Value *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsSparkSqlJobLoggingConfigDriverLogLevels) HashCode() string {
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
	Properties        []WorkflowTemplateJobsPrestoJobProperties   `json:"properties"`
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

type WorkflowTemplateJobsPrestoJobProperties struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPrestoJobProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPrestoJobProperties *WorkflowTemplateJobsPrestoJobProperties = &WorkflowTemplateJobsPrestoJobProperties{empty: true}

func (r *WorkflowTemplateJobsPrestoJobProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPrestoJobProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsPrestoJobLoggingConfig struct {
	empty           bool                                                        `json:"-"`
	DriverLogLevels []WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels `json:"driverLogLevels"`
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

type WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels struct {
	empty bool                                                                `json:"-"`
	Key   *string                                                             `json:"key"`
	Value *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevelsValueEnum `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels = &WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels{empty: true}

func (r *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsPrestoJobLoggingConfigDriverLogLevels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkflowTemplateJobsLabels struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// This object is used to assert a desired state where this WorkflowTemplateJobsLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkflowTemplateJobsLabels *WorkflowTemplateJobsLabels = &WorkflowTemplateJobsLabels{empty: true}

func (r *WorkflowTemplateJobsLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkflowTemplateJobsLabels) HashCode() string {
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
