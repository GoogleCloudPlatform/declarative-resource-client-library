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

type AutoscalingPolicy struct {
	Name                  *string                                 `json:"name"`
	BasicAlgorithm        *AutoscalingPolicyBasicAlgorithm        `json:"basicAlgorithm"`
	WorkerConfig          *AutoscalingPolicyWorkerConfig          `json:"workerConfig"`
	SecondaryWorkerConfig *AutoscalingPolicySecondaryWorkerConfig `json:"secondaryWorkerConfig"`
	Project               *string                                 `json:"project"`
	Location              *string                                 `json:"location"`
}

func (r *AutoscalingPolicy) String() string {
	return dcl.SprintResource(r)
}

type AutoscalingPolicyBasicAlgorithm struct {
	empty          bool                                       `json:"-"`
	YarnConfig     *AutoscalingPolicyBasicAlgorithmYarnConfig `json:"yarnConfig"`
	CooldownPeriod *string                                    `json:"cooldownPeriod"`
}

type jsonAutoscalingPolicyBasicAlgorithm AutoscalingPolicyBasicAlgorithm

func (r *AutoscalingPolicyBasicAlgorithm) UnmarshalJSON(data []byte) error {
	var res jsonAutoscalingPolicyBasicAlgorithm
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAutoscalingPolicyBasicAlgorithm
	} else {

		r.YarnConfig = res.YarnConfig

		r.CooldownPeriod = res.CooldownPeriod

	}
	return nil
}

// This object is used to assert a desired state where this AutoscalingPolicyBasicAlgorithm is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalingPolicyBasicAlgorithm *AutoscalingPolicyBasicAlgorithm = &AutoscalingPolicyBasicAlgorithm{empty: true}

func (r *AutoscalingPolicyBasicAlgorithm) Empty() bool {
	return r.empty
}

func (r *AutoscalingPolicyBasicAlgorithm) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalingPolicyBasicAlgorithm) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalingPolicyBasicAlgorithmYarnConfig struct {
	empty                       bool     `json:"-"`
	GracefulDecommissionTimeout *string  `json:"gracefulDecommissionTimeout"`
	ScaleUpFactor               *float64 `json:"scaleUpFactor"`
	ScaleDownFactor             *float64 `json:"scaleDownFactor"`
	ScaleUpMinWorkerFraction    *float64 `json:"scaleUpMinWorkerFraction"`
	ScaleDownMinWorkerFraction  *float64 `json:"scaleDownMinWorkerFraction"`
}

type jsonAutoscalingPolicyBasicAlgorithmYarnConfig AutoscalingPolicyBasicAlgorithmYarnConfig

func (r *AutoscalingPolicyBasicAlgorithmYarnConfig) UnmarshalJSON(data []byte) error {
	var res jsonAutoscalingPolicyBasicAlgorithmYarnConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAutoscalingPolicyBasicAlgorithmYarnConfig
	} else {

		r.GracefulDecommissionTimeout = res.GracefulDecommissionTimeout

		r.ScaleUpFactor = res.ScaleUpFactor

		r.ScaleDownFactor = res.ScaleDownFactor

		r.ScaleUpMinWorkerFraction = res.ScaleUpMinWorkerFraction

		r.ScaleDownMinWorkerFraction = res.ScaleDownMinWorkerFraction

	}
	return nil
}

// This object is used to assert a desired state where this AutoscalingPolicyBasicAlgorithmYarnConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalingPolicyBasicAlgorithmYarnConfig *AutoscalingPolicyBasicAlgorithmYarnConfig = &AutoscalingPolicyBasicAlgorithmYarnConfig{empty: true}

func (r *AutoscalingPolicyBasicAlgorithmYarnConfig) Empty() bool {
	return r.empty
}

func (r *AutoscalingPolicyBasicAlgorithmYarnConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalingPolicyBasicAlgorithmYarnConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalingPolicyWorkerConfig struct {
	empty        bool   `json:"-"`
	MinInstances *int64 `json:"minInstances"`
	MaxInstances *int64 `json:"maxInstances"`
	Weight       *int64 `json:"weight"`
}

type jsonAutoscalingPolicyWorkerConfig AutoscalingPolicyWorkerConfig

func (r *AutoscalingPolicyWorkerConfig) UnmarshalJSON(data []byte) error {
	var res jsonAutoscalingPolicyWorkerConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAutoscalingPolicyWorkerConfig
	} else {

		r.MinInstances = res.MinInstances

		r.MaxInstances = res.MaxInstances

		r.Weight = res.Weight

	}
	return nil
}

// This object is used to assert a desired state where this AutoscalingPolicyWorkerConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalingPolicyWorkerConfig *AutoscalingPolicyWorkerConfig = &AutoscalingPolicyWorkerConfig{empty: true}

func (r *AutoscalingPolicyWorkerConfig) Empty() bool {
	return r.empty
}

func (r *AutoscalingPolicyWorkerConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalingPolicyWorkerConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AutoscalingPolicySecondaryWorkerConfig struct {
	empty        bool   `json:"-"`
	MinInstances *int64 `json:"minInstances"`
	MaxInstances *int64 `json:"maxInstances"`
	Weight       *int64 `json:"weight"`
}

type jsonAutoscalingPolicySecondaryWorkerConfig AutoscalingPolicySecondaryWorkerConfig

func (r *AutoscalingPolicySecondaryWorkerConfig) UnmarshalJSON(data []byte) error {
	var res jsonAutoscalingPolicySecondaryWorkerConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAutoscalingPolicySecondaryWorkerConfig
	} else {

		r.MinInstances = res.MinInstances

		r.MaxInstances = res.MaxInstances

		r.Weight = res.Weight

	}
	return nil
}

// This object is used to assert a desired state where this AutoscalingPolicySecondaryWorkerConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAutoscalingPolicySecondaryWorkerConfig *AutoscalingPolicySecondaryWorkerConfig = &AutoscalingPolicySecondaryWorkerConfig{empty: true}

func (r *AutoscalingPolicySecondaryWorkerConfig) Empty() bool {
	return r.empty
}

func (r *AutoscalingPolicySecondaryWorkerConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *AutoscalingPolicySecondaryWorkerConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AutoscalingPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dataproc",
		Type:    "AutoscalingPolicy",
		Version: "dataproc",
	}
}

const AutoscalingPolicyMaxPage = -1

type AutoscalingPolicyList struct {
	Items []*AutoscalingPolicy

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *AutoscalingPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AutoscalingPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAutoscalingPolicy(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAutoscalingPolicy(ctx context.Context, project, location string) (*AutoscalingPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAutoscalingPolicyWithMaxResults(ctx, project, location, AutoscalingPolicyMaxPage)

}

func (c *Client) ListAutoscalingPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*AutoscalingPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAutoscalingPolicy(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AutoscalingPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetAutoscalingPolicy(ctx context.Context, r *AutoscalingPolicy) (*AutoscalingPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAutoscalingPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAutoscalingPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAutoscalingPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAutoscalingPolicy(ctx context.Context, r *AutoscalingPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AutoscalingPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting AutoscalingPolicy...")
	deleteOp := deleteAutoscalingPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAutoscalingPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAutoscalingPolicy(ctx context.Context, project, location string, filter func(*AutoscalingPolicy) bool) error {
	listObj, err := c.ListAutoscalingPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllAutoscalingPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAutoscalingPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAutoscalingPolicy(ctx context.Context, rawDesired *AutoscalingPolicy, opts ...dcl.ApplyOption) (*AutoscalingPolicy, error) {

	var resultNewState *AutoscalingPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAutoscalingPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyAutoscalingPolicyHelper(c *Client, ctx context.Context, rawDesired *AutoscalingPolicy, opts ...dcl.ApplyOption) (*AutoscalingPolicy, error) {
	c.Config.Logger.Info("Beginning ApplyAutoscalingPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.autoscalingPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToAutoscalingPolicyOp(opStrings, fieldDiffs, opts)
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
	var ops []autoscalingPolicyApiOperation
	if create {
		ops = append(ops, &createAutoscalingPolicyOperation{})
	} else if recreate {

		ops = append(ops, &deleteAutoscalingPolicyOperation{})

		ops = append(ops, &createAutoscalingPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAutoscalingPolicyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAutoscalingPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAutoscalingPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAutoscalingPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAutoscalingPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAutoscalingPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAutoscalingPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAutoscalingPolicy(c, newDesired, newState)
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
