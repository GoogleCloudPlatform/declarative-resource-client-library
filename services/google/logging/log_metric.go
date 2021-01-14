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
package logging

import (
	"context"
	"crypto/sha256"
	"fmt"
	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type LogMetric struct {
	Name             *string                    `json:"name"`
	Description      *string                    `json:"description"`
	Filter           *string                    `json:"filter"`
	Disabled         *bool                      `json:"disabled"`
	MetricDescriptor *LogMetricMetricDescriptor `json:"metricDescriptor"`
	ValueExtractor   *string                    `json:"valueExtractor"`
	LabelExtractors  map[string]string          `json:"labelExtractors"`
	BucketOptions    *LogMetricBucketOptions    `json:"bucketOptions"`
	CreateTime       *string                    `json:"createTime"`
	UpdateTime       *string                    `json:"updateTime"`
	Resolution       *string                    `json:"resolution"`
	Project          *string                    `json:"project"`
}

func (r *LogMetric) String() string {
	return dcl.SprintResource(r)
}

// The enum LogMetricMetricDescriptorLabelsValueTypeEnum.
type LogMetricMetricDescriptorLabelsValueTypeEnum string

// LogMetricMetricDescriptorLabelsValueTypeEnumRef returns a *LogMetricMetricDescriptorLabelsValueTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func LogMetricMetricDescriptorLabelsValueTypeEnumRef(s string) *LogMetricMetricDescriptorLabelsValueTypeEnum {
	if s == "" {
		return nil
	}

	v := LogMetricMetricDescriptorLabelsValueTypeEnum(s)
	return &v
}

func (v LogMetricMetricDescriptorLabelsValueTypeEnum) Validate() error {
	for _, s := range []string{"STRING", "BOOL", "INT64", "DOUBLE", "DISTRIBUTION", "MONEY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "LogMetricMetricDescriptorLabelsValueTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum LogMetricMetricDescriptorMetricKindEnum.
type LogMetricMetricDescriptorMetricKindEnum string

// LogMetricMetricDescriptorMetricKindEnumRef returns a *LogMetricMetricDescriptorMetricKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func LogMetricMetricDescriptorMetricKindEnumRef(s string) *LogMetricMetricDescriptorMetricKindEnum {
	if s == "" {
		return nil
	}

	v := LogMetricMetricDescriptorMetricKindEnum(s)
	return &v
}

func (v LogMetricMetricDescriptorMetricKindEnum) Validate() error {
	for _, s := range []string{"METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "LogMetricMetricDescriptorMetricKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum LogMetricMetricDescriptorValueTypeEnum.
type LogMetricMetricDescriptorValueTypeEnum string

// LogMetricMetricDescriptorValueTypeEnumRef returns a *LogMetricMetricDescriptorValueTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func LogMetricMetricDescriptorValueTypeEnumRef(s string) *LogMetricMetricDescriptorValueTypeEnum {
	if s == "" {
		return nil
	}

	v := LogMetricMetricDescriptorValueTypeEnum(s)
	return &v
}

func (v LogMetricMetricDescriptorValueTypeEnum) Validate() error {
	for _, s := range []string{"STRING", "BOOL", "INT64", "DOUBLE", "DISTRIBUTION", "MONEY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "LogMetricMetricDescriptorValueTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum LogMetricMetricDescriptorMetadataLaunchStageEnum.
type LogMetricMetricDescriptorMetadataLaunchStageEnum string

// LogMetricMetricDescriptorMetadataLaunchStageEnumRef returns a *LogMetricMetricDescriptorMetadataLaunchStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func LogMetricMetricDescriptorMetadataLaunchStageEnumRef(s string) *LogMetricMetricDescriptorMetadataLaunchStageEnum {
	if s == "" {
		return nil
	}

	v := LogMetricMetricDescriptorMetadataLaunchStageEnum(s)
	return &v
}

func (v LogMetricMetricDescriptorMetadataLaunchStageEnum) Validate() error {
	for _, s := range []string{"LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "LogMetricMetricDescriptorMetadataLaunchStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum LogMetricMetricDescriptorLaunchStageEnum.
type LogMetricMetricDescriptorLaunchStageEnum string

// LogMetricMetricDescriptorLaunchStageEnumRef returns a *LogMetricMetricDescriptorLaunchStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func LogMetricMetricDescriptorLaunchStageEnumRef(s string) *LogMetricMetricDescriptorLaunchStageEnum {
	if s == "" {
		return nil
	}

	v := LogMetricMetricDescriptorLaunchStageEnum(s)
	return &v
}

func (v LogMetricMetricDescriptorLaunchStageEnum) Validate() error {
	for _, s := range []string{"LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "LogMetricMetricDescriptorLaunchStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type LogMetricMetricDescriptor struct {
	empty                  bool                                      `json:"-"`
	Name                   *string                                   `json:"name"`
	Type                   *string                                   `json:"type"`
	Labels                 []LogMetricMetricDescriptorLabels         `json:"labels"`
	MetricKind             *LogMetricMetricDescriptorMetricKindEnum  `json:"metricKind"`
	ValueType              *LogMetricMetricDescriptorValueTypeEnum   `json:"valueType"`
	Unit                   *string                                   `json:"unit"`
	DisplayName            *string                                   `json:"displayName"`
	Metadata               *LogMetricMetricDescriptorMetadata        `json:"metadata"`
	LaunchStage            *LogMetricMetricDescriptorLaunchStageEnum `json:"launchStage"`
	MonitoredResourceTypes []string                                  `json:"monitoredResourceTypes"`
}

// This object is used to assert a desired state where this LogMetricMetricDescriptor is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricMetricDescriptor *LogMetricMetricDescriptor = &LogMetricMetricDescriptor{empty: true}

func (r *LogMetricMetricDescriptor) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricMetricDescriptor) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricMetricDescriptorLabels struct {
	empty       bool                                          `json:"-"`
	Key         *string                                       `json:"key"`
	ValueType   *LogMetricMetricDescriptorLabelsValueTypeEnum `json:"valueType"`
	Description *string                                       `json:"description"`
}

// This object is used to assert a desired state where this LogMetricMetricDescriptorLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricMetricDescriptorLabels *LogMetricMetricDescriptorLabels = &LogMetricMetricDescriptorLabels{empty: true}

func (r *LogMetricMetricDescriptorLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricMetricDescriptorLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricMetricDescriptorMetadata struct {
	empty        bool                                              `json:"-"`
	LaunchStage  *LogMetricMetricDescriptorMetadataLaunchStageEnum `json:"launchStage"`
	SamplePeriod *LogMetricMetricDescriptorMetadataSamplePeriod    `json:"samplePeriod"`
	IngestDelay  *LogMetricMetricDescriptorMetadataIngestDelay     `json:"ingestDelay"`
}

// This object is used to assert a desired state where this LogMetricMetricDescriptorMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricMetricDescriptorMetadata *LogMetricMetricDescriptorMetadata = &LogMetricMetricDescriptorMetadata{empty: true}

func (r *LogMetricMetricDescriptorMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricMetricDescriptorMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricMetricDescriptorMetadataSamplePeriod struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this LogMetricMetricDescriptorMetadataSamplePeriod is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricMetricDescriptorMetadataSamplePeriod *LogMetricMetricDescriptorMetadataSamplePeriod = &LogMetricMetricDescriptorMetadataSamplePeriod{empty: true}

func (r *LogMetricMetricDescriptorMetadataSamplePeriod) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricMetricDescriptorMetadataSamplePeriod) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricMetricDescriptorMetadataIngestDelay struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

// This object is used to assert a desired state where this LogMetricMetricDescriptorMetadataIngestDelay is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricMetricDescriptorMetadataIngestDelay *LogMetricMetricDescriptorMetadataIngestDelay = &LogMetricMetricDescriptorMetadataIngestDelay{empty: true}

func (r *LogMetricMetricDescriptorMetadataIngestDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricMetricDescriptorMetadataIngestDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricBucketOptions struct {
	empty              bool                                      `json:"-"`
	LinearBuckets      *LogMetricBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *LogMetricBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *LogMetricBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

// This object is used to assert a desired state where this LogMetricBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricBucketOptions *LogMetricBucketOptions = &LogMetricBucketOptions{empty: true}

func (r *LogMetricBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

// This object is used to assert a desired state where this LogMetricBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricBucketOptionsLinearBuckets *LogMetricBucketOptionsLinearBuckets = &LogMetricBucketOptionsLinearBuckets{empty: true}

func (r *LogMetricBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

// This object is used to assert a desired state where this LogMetricBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricBucketOptionsExponentialBuckets *LogMetricBucketOptionsExponentialBuckets = &LogMetricBucketOptionsExponentialBuckets{empty: true}

func (r *LogMetricBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type LogMetricBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

// This object is used to assert a desired state where this LogMetricBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyLogMetricBucketOptionsExplicitBuckets *LogMetricBucketOptionsExplicitBuckets = &LogMetricBucketOptionsExplicitBuckets{empty: true}

func (r *LogMetricBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *LogMetricBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *LogMetric) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "logging",
		Type:    "LogMetric",
		Version: "logging",
	}
}

const LogMetricMaxPage = -1

type LogMetricList struct {
	Items []*LogMetric

	nextToken string

	pageSize int32

	project string
}

func (l *LogMetricList) HasNext() bool {
	return l.nextToken != ""
}

func (l *LogMetricList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listLogMetric(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListLogMetric(ctx context.Context, project string) (*LogMetricList, error) {

	return c.ListLogMetricWithMaxResults(ctx, project, LogMetricMaxPage)

}

func (c *Client) ListLogMetricWithMaxResults(ctx context.Context, project string, pageSize int32) (*LogMetricList, error) {
	items, token, err := c.listLogMetric(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &LogMetricList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetLogMetric(ctx context.Context, r *LogMetric) (*LogMetric, error) {
	b, err := c.getLogMetricRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalLogMetric(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeLogMetricNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteLogMetric(ctx context.Context, r *LogMetric) error {
	if r == nil {
		return fmt.Errorf("LogMetric resource is nil")
	}
	c.Config.Logger.Info("Deleting LogMetric...")
	deleteOp := deleteLogMetricOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllLogMetric deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllLogMetric(ctx context.Context, project string, filter func(*LogMetric) bool) error {
	listObj, err := c.ListLogMetric(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllLogMetric(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllLogMetric(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyLogMetric(ctx context.Context, rawDesired *LogMetric, opts ...dcl.ApplyOption) (*LogMetric, error) {
	c.Config.Logger.Info("Beginning ApplyLogMetric...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.logMetricDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []logMetricApiOperation
	if create {
		ops = append(ops, &createLogMetricOperation{})
	} else if recreate {
		ops = append(ops, &deleteLogMetricOperation{})
		ops = append(ops, &createLogMetricOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeLogMetricDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetLogMetric(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeLogMetricNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeLogMetricDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffLogMetric(c, newDesired, newState)
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
