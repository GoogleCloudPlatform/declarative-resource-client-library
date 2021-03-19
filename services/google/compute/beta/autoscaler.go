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
package beta

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Autoscaler struct {
	Id                    *int64                       `json:"id"`
	Name                  *string                      `json:"name"`
	Description           *string                      `json:"description"`
	Target                *string                      `json:"target"`
	AutoscalingPolicy     *AutoscalerAutoscalingPolicy `json:"autoscalingPolicy"`
	Zone                  *string                      `json:"zone"`
	Region                *string                      `json:"region"`
	SelfLink              *string                      `json:"selfLink"`
	Status                *AutoscalerStatusEnum        `json:"status"`
	StatusDetails         []AutoscalerStatusDetails    `json:"statusDetails"`
	RecommendedSize       *int64                       `json:"recommendedSize"`
	SelfLinkWithId        *string                      `json:"selfLinkWithId"`
	ScalingScheduleStatus map[string]string            `json:"scalingScheduleStatus"`
	Project               *string                      `json:"project"`
	CreationTimestamp     *string                      `json:"creationTimestamp"`
	Location              *string                      `json:"location"`
}

func (r *Autoscaler) String() string {
	return dcl.SprintResource(r)
}

// The enum AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum.
type AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum string

// AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumRef returns a *AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumRef(s string) *AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	if s == "" {
		return nil
	}

	v := AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(s)
	return &v
}

func (v AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) Validate() error {
	for _, s := range []string{"GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AutoscalerAutoscalingPolicyModeEnum.
type AutoscalerAutoscalingPolicyModeEnum string

// AutoscalerAutoscalingPolicyModeEnumRef returns a *AutoscalerAutoscalingPolicyModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AutoscalerAutoscalingPolicyModeEnumRef(s string) *AutoscalerAutoscalingPolicyModeEnum {
	if s == "" {
		return nil
	}

	v := AutoscalerAutoscalingPolicyModeEnum(s)
	return &v
}

func (v AutoscalerAutoscalingPolicyModeEnum) Validate() error {
	for _, s := range []string{"OFF", "ON", "ONLY_SCALE_OUT", "ONLY_UP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AutoscalerAutoscalingPolicyModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AutoscalerStatusEnum.
type AutoscalerStatusEnum string

// AutoscalerStatusEnumRef returns a *AutoscalerStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func AutoscalerStatusEnumRef(s string) *AutoscalerStatusEnum {
	if s == "" {
		return nil
	}

	v := AutoscalerStatusEnum(s)
	return &v
}

func (v AutoscalerStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "RUNNING", "DONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AutoscalerStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AutoscalerStatusDetailsTypeEnum.
type AutoscalerStatusDetailsTypeEnum string

// AutoscalerStatusDetailsTypeEnumRef returns a *AutoscalerStatusDetailsTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AutoscalerStatusDetailsTypeEnumRef(s string) *AutoscalerStatusDetailsTypeEnum {
	if s == "" {
		return nil
	}

	v := AutoscalerStatusDetailsTypeEnum(s)
	return &v
}

func (v AutoscalerStatusDetailsTypeEnum) Validate() error {
	for _, s := range []string{"PATH", "OTHER", "PARAMETER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AutoscalerStatusDetailsTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AutoscalerAutoscalingPolicy struct {
	empty                    bool                                                  `json:"-"`
	MinNumReplicas           *int64                                                `json:"minNumReplicas"`
	MaxNumReplicas           *int64                                                `json:"maxNumReplicas"`
	ScaleInControl           *AutoscalerAutoscalingPolicyScaleInControl            `json:"scaleInControl"`
	CoolDownPeriodSec        *int64                                                `json:"coolDownPeriodSec"`
	CpuUtilization           *AutoscalerAutoscalingPolicyCpuUtilization            `json:"cpuUtilization"`
	CustomMetricUtilizations []AutoscalerAutoscalingPolicyCustomMetricUtilizations `json:"customMetricUtilizations"`
	LoadBalancingUtilization *AutoscalerAutoscalingPolicyLoadBalancingUtilization  `json:"loadBalancingUtilization"`
	Mode                     *AutoscalerAutoscalingPolicyModeEnum                  `json:"mode"`
}

// This object is used to assert a desired state where this AutoscalerAutoscalingPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerAutoscalingPolicy *AutoscalerAutoscalingPolicy = &AutoscalerAutoscalingPolicy{empty: true}

func (r *AutoscalerAutoscalingPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerAutoscalingPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalerAutoscalingPolicyScaleInControl struct {
	empty               bool                                                          `json:"-"`
	MaxScaledInReplicas *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas `json:"maxScaledInReplicas"`
	TimeWindowSec       *int64                                                        `json:"timeWindowSec"`
}

// This object is used to assert a desired state where this AutoscalerAutoscalingPolicyScaleInControl is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerAutoscalingPolicyScaleInControl *AutoscalerAutoscalingPolicyScaleInControl = &AutoscalerAutoscalingPolicyScaleInControl{empty: true}

func (r *AutoscalerAutoscalingPolicyScaleInControl) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerAutoscalingPolicyScaleInControl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas struct {
	empty      bool   `json:"-"`
	Fixed      *int64 `json:"fixed"`
	Percent    *int64 `json:"percent"`
	Calculated *int64 `json:"calculated"`
}

// This object is used to assert a desired state where this AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas = &AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{empty: true}

func (r *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalerAutoscalingPolicyCpuUtilization struct {
	empty             bool     `json:"-"`
	UtilizationTarget *float64 `json:"utilizationTarget"`
}

// This object is used to assert a desired state where this AutoscalerAutoscalingPolicyCpuUtilization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerAutoscalingPolicyCpuUtilization *AutoscalerAutoscalingPolicyCpuUtilization = &AutoscalerAutoscalingPolicyCpuUtilization{empty: true}

func (r *AutoscalerAutoscalingPolicyCpuUtilization) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerAutoscalingPolicyCpuUtilization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalerAutoscalingPolicyCustomMetricUtilizations struct {
	empty                    bool                                                                          `json:"-"`
	Metric                   *string                                                                       `json:"metric"`
	UtilizationTarget        *float64                                                                      `json:"utilizationTarget"`
	UtilizationTargetType    *AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum `json:"utilizationTargetType"`
	Filter                   *string                                                                       `json:"filter"`
	SingleInstanceAssignment *float64                                                                      `json:"singleInstanceAssignment"`
}

// This object is used to assert a desired state where this AutoscalerAutoscalingPolicyCustomMetricUtilizations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerAutoscalingPolicyCustomMetricUtilizations *AutoscalerAutoscalingPolicyCustomMetricUtilizations = &AutoscalerAutoscalingPolicyCustomMetricUtilizations{empty: true}

func (r *AutoscalerAutoscalingPolicyCustomMetricUtilizations) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerAutoscalingPolicyCustomMetricUtilizations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalerAutoscalingPolicyLoadBalancingUtilization struct {
	empty             bool     `json:"-"`
	UtilizationTarget *float64 `json:"utilizationTarget"`
}

// This object is used to assert a desired state where this AutoscalerAutoscalingPolicyLoadBalancingUtilization is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerAutoscalingPolicyLoadBalancingUtilization *AutoscalerAutoscalingPolicyLoadBalancingUtilization = &AutoscalerAutoscalingPolicyLoadBalancingUtilization{empty: true}

func (r *AutoscalerAutoscalingPolicyLoadBalancingUtilization) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerAutoscalingPolicyLoadBalancingUtilization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalerStatusDetails struct {
	empty   bool                             `json:"-"`
	Message *string                          `json:"message"`
	Type    *AutoscalerStatusDetailsTypeEnum `json:"type"`
}

// This object is used to assert a desired state where this AutoscalerStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalerStatusDetails *AutoscalerStatusDetails = &AutoscalerStatusDetails{empty: true}

func (r *AutoscalerStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalerStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Autoscaler) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Autoscaler",
		Version: "beta",
	}
}

const AutoscalerMaxPage = -1

type AutoscalerList struct {
	Items []*Autoscaler

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *AutoscalerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AutoscalerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAutoscaler(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAutoscaler(ctx context.Context, project, location string) (*AutoscalerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListAutoscalerWithMaxResults(ctx, project, location, AutoscalerMaxPage)

}

func (c *Client) ListAutoscalerWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*AutoscalerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listAutoscaler(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AutoscalerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetAutoscaler(ctx context.Context, r *Autoscaler) (*Autoscaler, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getAutoscalerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAutoscaler(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAutoscalerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAutoscaler(ctx context.Context, r *Autoscaler) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Autoscaler resource is nil")
	}
	c.Config.Logger.Info("Deleting Autoscaler...")
	deleteOp := deleteAutoscalerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAutoscaler deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAutoscaler(ctx context.Context, project, location string, filter func(*Autoscaler) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListAutoscaler(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllAutoscaler(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAutoscaler(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAutoscaler(ctx context.Context, rawDesired *Autoscaler, opts ...dcl.ApplyOption) (*Autoscaler, error) {
	c.Config.Logger.Info("Beginning ApplyAutoscaler...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.autoscalerDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []autoscalerApiOperation
	if create {
		ops = append(ops, &createAutoscalerOperation{})
	} else if recreate {

		ops = append(ops, &deleteAutoscalerOperation{})

		ops = append(ops, &createAutoscalerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAutoscalerDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAutoscaler(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAutoscalerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAutoscaler(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAutoscalerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAutoscalerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAutoscalerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAutoscaler(c, newDesired, newState)
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
