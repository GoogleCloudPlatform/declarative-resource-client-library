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
package monitoring

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type MetricDescriptor struct {
	SelfLink               *string                          `json:"selfLink"`
	Type                   *string                          `json:"type"`
	Labels                 []MetricDescriptorLabels         `json:"labels"`
	MetricKind             *MetricDescriptorMetricKindEnum  `json:"metricKind"`
	ValueType              *MetricDescriptorValueTypeEnum   `json:"valueType"`
	Unit                   *string                          `json:"unit"`
	Description            *string                          `json:"description"`
	DisplayName            *string                          `json:"displayName"`
	Metadata               *MetricDescriptorMetadata        `json:"metadata"`
	LaunchStage            *MetricDescriptorLaunchStageEnum `json:"launchStage"`
	MonitoredResourceTypes []string                         `json:"monitoredResourceTypes"`
	Project                *string                          `json:"project"`
}

func (r *MetricDescriptor) String() string {
	return dcl.SprintResource(r)
}

// The enum MetricDescriptorLabelsValueTypeEnum.
type MetricDescriptorLabelsValueTypeEnum string

// MetricDescriptorLabelsValueTypeEnumRef returns a *MetricDescriptorLabelsValueTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func MetricDescriptorLabelsValueTypeEnumRef(s string) *MetricDescriptorLabelsValueTypeEnum {
	if s == "" {
		return nil
	}

	v := MetricDescriptorLabelsValueTypeEnum(s)
	return &v
}

func (v MetricDescriptorLabelsValueTypeEnum) Validate() error {
	for _, s := range []string{"STRING", "BOOL", "INT64"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MetricDescriptorLabelsValueTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MetricDescriptorMetricKindEnum.
type MetricDescriptorMetricKindEnum string

// MetricDescriptorMetricKindEnumRef returns a *MetricDescriptorMetricKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func MetricDescriptorMetricKindEnumRef(s string) *MetricDescriptorMetricKindEnum {
	if s == "" {
		return nil
	}

	v := MetricDescriptorMetricKindEnum(s)
	return &v
}

func (v MetricDescriptorMetricKindEnum) Validate() error {
	for _, s := range []string{"METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MetricDescriptorMetricKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MetricDescriptorValueTypeEnum.
type MetricDescriptorValueTypeEnum string

// MetricDescriptorValueTypeEnumRef returns a *MetricDescriptorValueTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func MetricDescriptorValueTypeEnumRef(s string) *MetricDescriptorValueTypeEnum {
	if s == "" {
		return nil
	}

	v := MetricDescriptorValueTypeEnum(s)
	return &v
}

func (v MetricDescriptorValueTypeEnum) Validate() error {
	for _, s := range []string{"STRING", "BOOL", "INT64"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MetricDescriptorValueTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MetricDescriptorMetadataLaunchStageEnum.
type MetricDescriptorMetadataLaunchStageEnum string

// MetricDescriptorMetadataLaunchStageEnumRef returns a *MetricDescriptorMetadataLaunchStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func MetricDescriptorMetadataLaunchStageEnumRef(s string) *MetricDescriptorMetadataLaunchStageEnum {
	if s == "" {
		return nil
	}

	v := MetricDescriptorMetadataLaunchStageEnum(s)
	return &v
}

func (v MetricDescriptorMetadataLaunchStageEnum) Validate() error {
	for _, s := range []string{"LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MetricDescriptorMetadataLaunchStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MetricDescriptorLaunchStageEnum.
type MetricDescriptorLaunchStageEnum string

// MetricDescriptorLaunchStageEnumRef returns a *MetricDescriptorLaunchStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func MetricDescriptorLaunchStageEnumRef(s string) *MetricDescriptorLaunchStageEnum {
	if s == "" {
		return nil
	}

	v := MetricDescriptorLaunchStageEnum(s)
	return &v
}

func (v MetricDescriptorLaunchStageEnum) Validate() error {
	for _, s := range []string{"LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MetricDescriptorLaunchStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type MetricDescriptorLabels struct {
	empty       bool                                 `json:"-"`
	Key         *string                              `json:"key"`
	ValueType   *MetricDescriptorLabelsValueTypeEnum `json:"valueType"`
	Description *string                              `json:"description"`
}

type jsonMetricDescriptorLabels MetricDescriptorLabels

func (r *MetricDescriptorLabels) UnmarshalJSON(data []byte) error {
	var res jsonMetricDescriptorLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMetricDescriptorLabels
	} else {

		r.Key = res.Key

		r.ValueType = res.ValueType

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this MetricDescriptorLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyMetricDescriptorLabels *MetricDescriptorLabels = &MetricDescriptorLabels{empty: true}

func (r *MetricDescriptorLabels) Empty() bool {
	return r.empty
}

func (r *MetricDescriptorLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *MetricDescriptorLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MetricDescriptorMetadata struct {
	empty        bool                                     `json:"-"`
	LaunchStage  *MetricDescriptorMetadataLaunchStageEnum `json:"launchStage"`
	SamplePeriod *string                                  `json:"samplePeriod"`
	IngestDelay  *string                                  `json:"ingestDelay"`
}

type jsonMetricDescriptorMetadata MetricDescriptorMetadata

func (r *MetricDescriptorMetadata) UnmarshalJSON(data []byte) error {
	var res jsonMetricDescriptorMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMetricDescriptorMetadata
	} else {

		r.LaunchStage = res.LaunchStage

		r.SamplePeriod = res.SamplePeriod

		r.IngestDelay = res.IngestDelay

	}
	return nil
}

// This object is used to assert a desired state where this MetricDescriptorMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyMetricDescriptorMetadata *MetricDescriptorMetadata = &MetricDescriptorMetadata{empty: true}

func (r *MetricDescriptorMetadata) Empty() bool {
	return r.empty
}

func (r *MetricDescriptorMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *MetricDescriptorMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *MetricDescriptor) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "monitoring",
		Type:    "MetricDescriptor",
		Version: "monitoring",
	}
}

const MetricDescriptorMaxPage = -1

type MetricDescriptorList struct {
	Items []*MetricDescriptor

	nextToken string

	pageSize int32

	project string
}

func (l *MetricDescriptorList) HasNext() bool {
	return l.nextToken != ""
}

func (l *MetricDescriptorList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listMetricDescriptor(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListMetricDescriptor(ctx context.Context, project string) (*MetricDescriptorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListMetricDescriptorWithMaxResults(ctx, project, MetricDescriptorMaxPage)

}

func (c *Client) ListMetricDescriptorWithMaxResults(ctx context.Context, project string, pageSize int32) (*MetricDescriptorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listMetricDescriptor(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &MetricDescriptorList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *MetricDescriptor) URLNormalized() *MetricDescriptor {
	normalized := dcl.Copy(*r).(MetricDescriptor)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Type = r.Type
	normalized.Unit = dcl.SelfLinkToName(r.Unit)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetMetricDescriptor(ctx context.Context, r *MetricDescriptor) (*MetricDescriptor, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getMetricDescriptorRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalMetricDescriptor(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Type = r.Type

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeMetricDescriptorNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteMetricDescriptor(ctx context.Context, r *MetricDescriptor) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("MetricDescriptor resource is nil")
	}
	c.Config.Logger.Info("Deleting MetricDescriptor...")
	deleteOp := deleteMetricDescriptorOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllMetricDescriptor deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllMetricDescriptor(ctx context.Context, project string, filter func(*MetricDescriptor) bool) error {
	listObj, err := c.ListMetricDescriptor(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllMetricDescriptor(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllMetricDescriptor(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyMetricDescriptor(ctx context.Context, rawDesired *MetricDescriptor, opts ...dcl.ApplyOption) (*MetricDescriptor, error) {
	var resultNewState *MetricDescriptor
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyMetricDescriptorHelper(c, ctx, rawDesired, opts...)
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

func applyMetricDescriptorHelper(c *Client, ctx context.Context, rawDesired *MetricDescriptor, opts ...dcl.ApplyOption) (*MetricDescriptor, error) {
	c.Config.Logger.Info("Beginning ApplyMetricDescriptor...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.metricDescriptorDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToMetricDescriptorDiffs(c.Config, fieldDiffs, opts)
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
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []metricDescriptorApiOperation
	if create {
		ops = append(ops, &createMetricDescriptorOperation{})
	} else if recreate {
		ops = append(ops, &deleteMetricDescriptorOperation{})
		ops = append(ops, &createMetricDescriptorOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeMetricDescriptorDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetMetricDescriptor(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createMetricDescriptorOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapMetricDescriptor(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeMetricDescriptorNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeMetricDescriptorNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeMetricDescriptorDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffMetricDescriptor(c, newDesired, newState)
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
