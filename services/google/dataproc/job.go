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
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Job struct {
	Reference               *JobReference         `json:"reference"`
	Placement               *JobPlacement         `json:"placement"`
	HadoopJob               *JobHadoopJob         `json:"hadoopJob"`
	SparkJob                *JobSparkJob          `json:"sparkJob"`
	PysparkJob              *JobPysparkJob        `json:"pysparkJob"`
	HiveJob                 *JobHiveJob           `json:"hiveJob"`
	PigJob                  *JobPigJob            `json:"pigJob"`
	SparkRJob               *JobSparkRJob         `json:"sparkRJob"`
	SparkSqlJob             *JobSparkSqlJob       `json:"sparkSqlJob"`
	PrestoJob               *JobPrestoJob         `json:"prestoJob"`
	Status                  *JobStatus            `json:"status"`
	StatusHistory           []JobStatusHistory    `json:"statusHistory"`
	YarnApplications        []JobYarnApplications `json:"yarnApplications"`
	DriverOutputResourceUri *string               `json:"driverOutputResourceUri"`
	DriverControlFilesUri   *string               `json:"driverControlFilesUri"`
	Labels                  map[string]string     `json:"labels"`
	Scheduling              *JobScheduling        `json:"scheduling"`
	Name                    *string               `json:"name"`
	Done                    *bool                 `json:"done"`
	Region                  *string               `json:"region"`
	Project                 *string               `json:"project"`
}

func (r *Job) String() string {
	return dcl.SprintResource(r)
}

// The enum JobStatusStateEnum.
type JobStatusStateEnum string

// JobStatusStateEnumRef returns a *JobStatusStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobStatusStateEnumRef(s string) *JobStatusStateEnum {
	if s == "" {
		return nil
	}

	v := JobStatusStateEnum(s)
	return &v
}

func (v JobStatusStateEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "CREATING", "RUNNING", "ERROR", "DELETING", "UPDATING", "STOPPING", "STOPPED", "STARTING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobStatusStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobStatusSubstateEnum.
type JobStatusSubstateEnum string

// JobStatusSubstateEnumRef returns a *JobStatusSubstateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobStatusSubstateEnumRef(s string) *JobStatusSubstateEnum {
	if s == "" {
		return nil
	}

	v := JobStatusSubstateEnum(s)
	return &v
}

func (v JobStatusSubstateEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "UNHEALTHY", "STALE_STATUS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobStatusSubstateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobStatusHistoryStateEnum.
type JobStatusHistoryStateEnum string

// JobStatusHistoryStateEnumRef returns a *JobStatusHistoryStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobStatusHistoryStateEnumRef(s string) *JobStatusHistoryStateEnum {
	if s == "" {
		return nil
	}

	v := JobStatusHistoryStateEnum(s)
	return &v
}

func (v JobStatusHistoryStateEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "CREATING", "RUNNING", "ERROR", "DELETING", "UPDATING", "STOPPING", "STOPPED", "STARTING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobStatusHistoryStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobStatusHistorySubstateEnum.
type JobStatusHistorySubstateEnum string

// JobStatusHistorySubstateEnumRef returns a *JobStatusHistorySubstateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobStatusHistorySubstateEnumRef(s string) *JobStatusHistorySubstateEnum {
	if s == "" {
		return nil
	}

	v := JobStatusHistorySubstateEnum(s)
	return &v
}

func (v JobStatusHistorySubstateEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED", "UNHEALTHY", "STALE_STATUS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobStatusHistorySubstateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobYarnApplicationsStateEnum.
type JobYarnApplicationsStateEnum string

// JobYarnApplicationsStateEnumRef returns a *JobYarnApplicationsStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobYarnApplicationsStateEnumRef(s string) *JobYarnApplicationsStateEnum {
	if s == "" {
		return nil
	}

	v := JobYarnApplicationsStateEnum(s)
	return &v
}

func (v JobYarnApplicationsStateEnum) Validate() error {
	for _, s := range []string{"UNKNOWN", "CREATING", "RUNNING", "ERROR", "DELETING", "UPDATING", "STOPPING", "STOPPED", "STARTING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobYarnApplicationsStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type JobReference struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	JobId     *string `json:"jobId"`
}

type jsonJobReference JobReference

func (r *JobReference) UnmarshalJSON(data []byte) error {
	var res jsonJobReference
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobReference
	} else {

		r.ProjectId = res.ProjectId

		r.JobId = res.JobId

	}
	return nil
}

// This object is used to assert a desired state where this JobReference is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobReference *JobReference = &JobReference{empty: true}

func (r *JobReference) Empty() bool {
	return r.empty
}

func (r *JobReference) String() string {
	return dcl.SprintResource(r)
}

func (r *JobReference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPlacement struct {
	empty         bool              `json:"-"`
	ClusterName   *string           `json:"clusterName"`
	ClusterUuid   *string           `json:"clusterUuid"`
	ClusterLabels map[string]string `json:"clusterLabels"`
}

type jsonJobPlacement JobPlacement

func (r *JobPlacement) UnmarshalJSON(data []byte) error {
	var res jsonJobPlacement
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPlacement
	} else {

		r.ClusterName = res.ClusterName

		r.ClusterUuid = res.ClusterUuid

		r.ClusterLabels = res.ClusterLabels

	}
	return nil
}

// This object is used to assert a desired state where this JobPlacement is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPlacement *JobPlacement = &JobPlacement{empty: true}

func (r *JobPlacement) Empty() bool {
	return r.empty
}

func (r *JobPlacement) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPlacement) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHadoopJob struct {
	empty          bool                       `json:"-"`
	MainJarFileUri *string                    `json:"mainJarFileUri"`
	MainClass      *string                    `json:"mainClass"`
	Args           []string                   `json:"args"`
	JarFileUris    []string                   `json:"jarFileUris"`
	FileUris       []string                   `json:"fileUris"`
	ArchiveUris    []string                   `json:"archiveUris"`
	Properties     map[string]string          `json:"properties"`
	LoggingConfig  *JobHadoopJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobHadoopJob JobHadoopJob

func (r *JobHadoopJob) UnmarshalJSON(data []byte) error {
	var res jsonJobHadoopJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHadoopJob
	} else {

		r.MainJarFileUri = res.MainJarFileUri

		r.MainClass = res.MainClass

		r.Args = res.Args

		r.JarFileUris = res.JarFileUris

		r.FileUris = res.FileUris

		r.ArchiveUris = res.ArchiveUris

		r.Properties = res.Properties

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobHadoopJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobHadoopJob *JobHadoopJob = &JobHadoopJob{empty: true}

func (r *JobHadoopJob) Empty() bool {
	return r.empty
}

func (r *JobHadoopJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHadoopJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHadoopJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobHadoopJobLoggingConfig JobHadoopJobLoggingConfig

func (r *JobHadoopJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobHadoopJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHadoopJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobHadoopJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobHadoopJobLoggingConfig *JobHadoopJobLoggingConfig = &JobHadoopJobLoggingConfig{empty: true}

func (r *JobHadoopJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobHadoopJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHadoopJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkJob struct {
	empty          bool                      `json:"-"`
	MainJarFileUri *string                   `json:"mainJarFileUri"`
	MainClass      *string                   `json:"mainClass"`
	Args           []string                  `json:"args"`
	JarFileUris    []string                  `json:"jarFileUris"`
	FileUris       []string                  `json:"fileUris"`
	ArchiveUris    []string                  `json:"archiveUris"`
	Properties     map[string]string         `json:"properties"`
	LoggingConfig  *JobSparkJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobSparkJob JobSparkJob

func (r *JobSparkJob) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkJob
	} else {

		r.MainJarFileUri = res.MainJarFileUri

		r.MainClass = res.MainClass

		r.Args = res.Args

		r.JarFileUris = res.JarFileUris

		r.FileUris = res.FileUris

		r.ArchiveUris = res.ArchiveUris

		r.Properties = res.Properties

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkJob *JobSparkJob = &JobSparkJob{empty: true}

func (r *JobSparkJob) Empty() bool {
	return r.empty
}

func (r *JobSparkJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobSparkJobLoggingConfig JobSparkJobLoggingConfig

func (r *JobSparkJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkJobLoggingConfig *JobSparkJobLoggingConfig = &JobSparkJobLoggingConfig{empty: true}

func (r *JobSparkJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobSparkJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPysparkJob struct {
	empty             bool                        `json:"-"`
	MainPythonFileUri *string                     `json:"mainPythonFileUri"`
	Args              []string                    `json:"args"`
	PythonFileUris    []string                    `json:"pythonFileUris"`
	JarFileUris       []string                    `json:"jarFileUris"`
	FileUris          []string                    `json:"fileUris"`
	ArchiveUris       []string                    `json:"archiveUris"`
	Properties        map[string]string           `json:"properties"`
	LoggingConfig     *JobPysparkJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobPysparkJob JobPysparkJob

func (r *JobPysparkJob) UnmarshalJSON(data []byte) error {
	var res jsonJobPysparkJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPysparkJob
	} else {

		r.MainPythonFileUri = res.MainPythonFileUri

		r.Args = res.Args

		r.PythonFileUris = res.PythonFileUris

		r.JarFileUris = res.JarFileUris

		r.FileUris = res.FileUris

		r.ArchiveUris = res.ArchiveUris

		r.Properties = res.Properties

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobPysparkJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPysparkJob *JobPysparkJob = &JobPysparkJob{empty: true}

func (r *JobPysparkJob) Empty() bool {
	return r.empty
}

func (r *JobPysparkJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPysparkJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPysparkJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobPysparkJobLoggingConfig JobPysparkJobLoggingConfig

func (r *JobPysparkJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobPysparkJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPysparkJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobPysparkJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPysparkJobLoggingConfig *JobPysparkJobLoggingConfig = &JobPysparkJobLoggingConfig{empty: true}

func (r *JobPysparkJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobPysparkJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPysparkJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHiveJob struct {
	empty             bool                 `json:"-"`
	QueryFileUri      *string              `json:"queryFileUri"`
	QueryList         *JobHiveJobQueryList `json:"queryList"`
	ContinueOnFailure *bool                `json:"continueOnFailure"`
	ScriptVariables   map[string]string    `json:"scriptVariables"`
	Properties        map[string]string    `json:"properties"`
	JarFileUris       []string             `json:"jarFileUris"`
}

type jsonJobHiveJob JobHiveJob

func (r *JobHiveJob) UnmarshalJSON(data []byte) error {
	var res jsonJobHiveJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHiveJob
	} else {

		r.QueryFileUri = res.QueryFileUri

		r.QueryList = res.QueryList

		r.ContinueOnFailure = res.ContinueOnFailure

		r.ScriptVariables = res.ScriptVariables

		r.Properties = res.Properties

		r.JarFileUris = res.JarFileUris

	}
	return nil
}

// This object is used to assert a desired state where this JobHiveJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobHiveJob *JobHiveJob = &JobHiveJob{empty: true}

func (r *JobHiveJob) Empty() bool {
	return r.empty
}

func (r *JobHiveJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHiveJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHiveJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

type jsonJobHiveJobQueryList JobHiveJobQueryList

func (r *JobHiveJobQueryList) UnmarshalJSON(data []byte) error {
	var res jsonJobHiveJobQueryList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHiveJobQueryList
	} else {

		r.Queries = res.Queries

	}
	return nil
}

// This object is used to assert a desired state where this JobHiveJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobHiveJobQueryList *JobHiveJobQueryList = &JobHiveJobQueryList{empty: true}

func (r *JobHiveJobQueryList) Empty() bool {
	return r.empty
}

func (r *JobHiveJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHiveJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPigJob struct {
	empty             bool                    `json:"-"`
	QueryFileUri      *string                 `json:"queryFileUri"`
	QueryList         *JobPigJobQueryList     `json:"queryList"`
	ContinueOnFailure *bool                   `json:"continueOnFailure"`
	ScriptVariables   map[string]string       `json:"scriptVariables"`
	Properties        map[string]string       `json:"properties"`
	JarFileUris       []string                `json:"jarFileUris"`
	LoggingConfig     *JobPigJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobPigJob JobPigJob

func (r *JobPigJob) UnmarshalJSON(data []byte) error {
	var res jsonJobPigJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPigJob
	} else {

		r.QueryFileUri = res.QueryFileUri

		r.QueryList = res.QueryList

		r.ContinueOnFailure = res.ContinueOnFailure

		r.ScriptVariables = res.ScriptVariables

		r.Properties = res.Properties

		r.JarFileUris = res.JarFileUris

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobPigJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPigJob *JobPigJob = &JobPigJob{empty: true}

func (r *JobPigJob) Empty() bool {
	return r.empty
}

func (r *JobPigJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPigJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPigJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

type jsonJobPigJobQueryList JobPigJobQueryList

func (r *JobPigJobQueryList) UnmarshalJSON(data []byte) error {
	var res jsonJobPigJobQueryList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPigJobQueryList
	} else {

		r.Queries = res.Queries

	}
	return nil
}

// This object is used to assert a desired state where this JobPigJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPigJobQueryList *JobPigJobQueryList = &JobPigJobQueryList{empty: true}

func (r *JobPigJobQueryList) Empty() bool {
	return r.empty
}

func (r *JobPigJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPigJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPigJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobPigJobLoggingConfig JobPigJobLoggingConfig

func (r *JobPigJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobPigJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPigJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobPigJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPigJobLoggingConfig *JobPigJobLoggingConfig = &JobPigJobLoggingConfig{empty: true}

func (r *JobPigJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobPigJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPigJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkRJob struct {
	empty         bool                       `json:"-"`
	MainRFileUri  *string                    `json:"mainRFileUri"`
	Args          []string                   `json:"args"`
	FileUris      []string                   `json:"fileUris"`
	ArchiveUris   []string                   `json:"archiveUris"`
	Properties    map[string]string          `json:"properties"`
	LoggingConfig *JobSparkRJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobSparkRJob JobSparkRJob

func (r *JobSparkRJob) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkRJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkRJob
	} else {

		r.MainRFileUri = res.MainRFileUri

		r.Args = res.Args

		r.FileUris = res.FileUris

		r.ArchiveUris = res.ArchiveUris

		r.Properties = res.Properties

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkRJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkRJob *JobSparkRJob = &JobSparkRJob{empty: true}

func (r *JobSparkRJob) Empty() bool {
	return r.empty
}

func (r *JobSparkRJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkRJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkRJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobSparkRJobLoggingConfig JobSparkRJobLoggingConfig

func (r *JobSparkRJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkRJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkRJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkRJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkRJobLoggingConfig *JobSparkRJobLoggingConfig = &JobSparkRJobLoggingConfig{empty: true}

func (r *JobSparkRJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobSparkRJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkRJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkSqlJob struct {
	empty           bool                         `json:"-"`
	QueryFileUri    *string                      `json:"queryFileUri"`
	QueryList       *JobSparkSqlJobQueryList     `json:"queryList"`
	ScriptVariables map[string]string            `json:"scriptVariables"`
	Properties      map[string]string            `json:"properties"`
	JarFileUris     []string                     `json:"jarFileUris"`
	LoggingConfig   *JobSparkSqlJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobSparkSqlJob JobSparkSqlJob

func (r *JobSparkSqlJob) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkSqlJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkSqlJob
	} else {

		r.QueryFileUri = res.QueryFileUri

		r.QueryList = res.QueryList

		r.ScriptVariables = res.ScriptVariables

		r.Properties = res.Properties

		r.JarFileUris = res.JarFileUris

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkSqlJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkSqlJob *JobSparkSqlJob = &JobSparkSqlJob{empty: true}

func (r *JobSparkSqlJob) Empty() bool {
	return r.empty
}

func (r *JobSparkSqlJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkSqlJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkSqlJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

type jsonJobSparkSqlJobQueryList JobSparkSqlJobQueryList

func (r *JobSparkSqlJobQueryList) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkSqlJobQueryList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkSqlJobQueryList
	} else {

		r.Queries = res.Queries

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkSqlJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkSqlJobQueryList *JobSparkSqlJobQueryList = &JobSparkSqlJobQueryList{empty: true}

func (r *JobSparkSqlJobQueryList) Empty() bool {
	return r.empty
}

func (r *JobSparkSqlJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkSqlJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobSparkSqlJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobSparkSqlJobLoggingConfig JobSparkSqlJobLoggingConfig

func (r *JobSparkSqlJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobSparkSqlJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobSparkSqlJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobSparkSqlJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobSparkSqlJobLoggingConfig *JobSparkSqlJobLoggingConfig = &JobSparkSqlJobLoggingConfig{empty: true}

func (r *JobSparkSqlJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobSparkSqlJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobSparkSqlJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPrestoJob struct {
	empty             bool                       `json:"-"`
	QueryFileUri      *string                    `json:"queryFileUri"`
	QueryList         *JobPrestoJobQueryList     `json:"queryList"`
	ContinueOnFailure *bool                      `json:"continueOnFailure"`
	OutputFormat      *string                    `json:"outputFormat"`
	ClientTags        []string                   `json:"clientTags"`
	Properties        map[string]string          `json:"properties"`
	LoggingConfig     *JobPrestoJobLoggingConfig `json:"loggingConfig"`
}

type jsonJobPrestoJob JobPrestoJob

func (r *JobPrestoJob) UnmarshalJSON(data []byte) error {
	var res jsonJobPrestoJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPrestoJob
	} else {

		r.QueryFileUri = res.QueryFileUri

		r.QueryList = res.QueryList

		r.ContinueOnFailure = res.ContinueOnFailure

		r.OutputFormat = res.OutputFormat

		r.ClientTags = res.ClientTags

		r.Properties = res.Properties

		r.LoggingConfig = res.LoggingConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobPrestoJob is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPrestoJob *JobPrestoJob = &JobPrestoJob{empty: true}

func (r *JobPrestoJob) Empty() bool {
	return r.empty
}

func (r *JobPrestoJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPrestoJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPrestoJobQueryList struct {
	empty   bool     `json:"-"`
	Queries []string `json:"queries"`
}

type jsonJobPrestoJobQueryList JobPrestoJobQueryList

func (r *JobPrestoJobQueryList) UnmarshalJSON(data []byte) error {
	var res jsonJobPrestoJobQueryList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPrestoJobQueryList
	} else {

		r.Queries = res.Queries

	}
	return nil
}

// This object is used to assert a desired state where this JobPrestoJobQueryList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPrestoJobQueryList *JobPrestoJobQueryList = &JobPrestoJobQueryList{empty: true}

func (r *JobPrestoJobQueryList) Empty() bool {
	return r.empty
}

func (r *JobPrestoJobQueryList) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPrestoJobQueryList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobPrestoJobLoggingConfig struct {
	empty           bool              `json:"-"`
	DriverLogLevels map[string]string `json:"driverLogLevels"`
}

type jsonJobPrestoJobLoggingConfig JobPrestoJobLoggingConfig

func (r *JobPrestoJobLoggingConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobPrestoJobLoggingConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPrestoJobLoggingConfig
	} else {

		r.DriverLogLevels = res.DriverLogLevels

	}
	return nil
}

// This object is used to assert a desired state where this JobPrestoJobLoggingConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobPrestoJobLoggingConfig *JobPrestoJobLoggingConfig = &JobPrestoJobLoggingConfig{empty: true}

func (r *JobPrestoJobLoggingConfig) Empty() bool {
	return r.empty
}

func (r *JobPrestoJobLoggingConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPrestoJobLoggingConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobStatus struct {
	empty          bool                   `json:"-"`
	State          *JobStatusStateEnum    `json:"state"`
	Details        *string                `json:"details"`
	StateStartTime *string                `json:"stateStartTime"`
	Substate       *JobStatusSubstateEnum `json:"substate"`
}

type jsonJobStatus JobStatus

func (r *JobStatus) UnmarshalJSON(data []byte) error {
	var res jsonJobStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobStatus
	} else {

		r.State = res.State

		r.Details = res.Details

		r.StateStartTime = res.StateStartTime

		r.Substate = res.Substate

	}
	return nil
}

// This object is used to assert a desired state where this JobStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobStatus *JobStatus = &JobStatus{empty: true}

func (r *JobStatus) Empty() bool {
	return r.empty
}

func (r *JobStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *JobStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobStatusHistory struct {
	empty          bool                          `json:"-"`
	State          *JobStatusHistoryStateEnum    `json:"state"`
	Details        *string                       `json:"details"`
	StateStartTime *string                       `json:"stateStartTime"`
	Substate       *JobStatusHistorySubstateEnum `json:"substate"`
}

type jsonJobStatusHistory JobStatusHistory

func (r *JobStatusHistory) UnmarshalJSON(data []byte) error {
	var res jsonJobStatusHistory
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobStatusHistory
	} else {

		r.State = res.State

		r.Details = res.Details

		r.StateStartTime = res.StateStartTime

		r.Substate = res.Substate

	}
	return nil
}

// This object is used to assert a desired state where this JobStatusHistory is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobStatusHistory *JobStatusHistory = &JobStatusHistory{empty: true}

func (r *JobStatusHistory) Empty() bool {
	return r.empty
}

func (r *JobStatusHistory) String() string {
	return dcl.SprintResource(r)
}

func (r *JobStatusHistory) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobYarnApplications struct {
	empty       bool                          `json:"-"`
	Name        *string                       `json:"name"`
	State       *JobYarnApplicationsStateEnum `json:"state"`
	Progress    *float64                      `json:"progress"`
	TrackingUrl *string                       `json:"trackingUrl"`
}

type jsonJobYarnApplications JobYarnApplications

func (r *JobYarnApplications) UnmarshalJSON(data []byte) error {
	var res jsonJobYarnApplications
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobYarnApplications
	} else {

		r.Name = res.Name

		r.State = res.State

		r.Progress = res.Progress

		r.TrackingUrl = res.TrackingUrl

	}
	return nil
}

// This object is used to assert a desired state where this JobYarnApplications is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobYarnApplications *JobYarnApplications = &JobYarnApplications{empty: true}

func (r *JobYarnApplications) Empty() bool {
	return r.empty
}

func (r *JobYarnApplications) String() string {
	return dcl.SprintResource(r)
}

func (r *JobYarnApplications) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobScheduling struct {
	empty              bool   `json:"-"`
	MaxFailuresPerHour *int64 `json:"maxFailuresPerHour"`
	MaxFailuresTotal   *int64 `json:"maxFailuresTotal"`
}

type jsonJobScheduling JobScheduling

func (r *JobScheduling) UnmarshalJSON(data []byte) error {
	var res jsonJobScheduling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobScheduling
	} else {

		r.MaxFailuresPerHour = res.MaxFailuresPerHour

		r.MaxFailuresTotal = res.MaxFailuresTotal

	}
	return nil
}

// This object is used to assert a desired state where this JobScheduling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyJobScheduling *JobScheduling = &JobScheduling{empty: true}

func (r *JobScheduling) Empty() bool {
	return r.empty
}

func (r *JobScheduling) String() string {
	return dcl.SprintResource(r)
}

func (r *JobScheduling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Job) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dataproc",
		Type:    "Job",
		Version: "dataproc",
	}
}

const JobMaxPage = -1

type JobList struct {
	Items []*Job

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *JobList) HasNext() bool {
	return l.nextToken != ""
}

func (l *JobList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listJob(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListJob(ctx context.Context, project, region string) (*JobList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListJobWithMaxResults(ctx, project, region, JobMaxPage)

}

func (c *Client) ListJobWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*JobList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listJob(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &JobList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

func (c *Client) GetJob(ctx context.Context, r *Job) (*Job, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getJobRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalJob(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeJobNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteJob(ctx context.Context, r *Job) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Job resource is nil")
	}
	c.Config.Logger.Info("Deleting Job...")
	deleteOp := deleteJobOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllJob deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllJob(ctx context.Context, project, region string, filter func(*Job) bool) error {
	listObj, err := c.ListJob(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllJob(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllJob(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyJob(ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (*Job, error) {

	var resultNewState *Job
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyJobHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyJobHelper(c *Client, ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (*Job, error) {
	c.Config.Logger.Info("Beginning ApplyJob...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.jobDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToJobOp(opStrings, fieldDiffs, opts)
	if err != nil {
		return nil, err
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []jobApiOperation
	if create {
		ops = append(ops, &createJobOperation{})
	} else if recreate {

		ops = append(ops, &deleteJobOperation{})

		ops = append(ops, &createJobOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeJobDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetJob(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createJobOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapJob(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeJobNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeJobNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeJobDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffJob(c, newDesired, newState)
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
