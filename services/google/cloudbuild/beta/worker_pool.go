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
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type WorkerPool struct {
	Name          *string                  `json:"name"`
	State         *WorkerPoolStateEnum     `json:"state"`
	CreateTime    *string                  `json:"createTime"`
	UpdateTime    *string                  `json:"updateTime"`
	DeleteTime    *string                  `json:"deleteTime"`
	WorkerConfig  *WorkerPoolWorkerConfig  `json:"workerConfig"`
	NetworkConfig *WorkerPoolNetworkConfig `json:"networkConfig"`
	Project       *string                  `json:"project"`
	Location      *string                  `json:"location"`
}

func (r *WorkerPool) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkerPoolStateEnum.
type WorkerPoolStateEnum string

// WorkerPoolStateEnumRef returns a *WorkerPoolStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkerPoolStateEnumRef(s string) *WorkerPoolStateEnum {
	if s == "" {
		return nil
	}

	v := WorkerPoolStateEnum(s)
	return &v
}

func (v WorkerPoolStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PENDING", "APPROVED", "REJECTED", "CANCELLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkerPoolStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type WorkerPoolWorkerConfig struct {
	empty        bool    `json:"-"`
	MachineType  *string `json:"machineType"`
	DiskSizeGb   *int64  `json:"diskSizeGb"`
	NoExternalIP *bool   `json:"noExternalIP"`
}

type jsonWorkerPoolWorkerConfig WorkerPoolWorkerConfig

func (r *WorkerPoolWorkerConfig) UnmarshalJSON(data []byte) error {
	var res jsonWorkerPoolWorkerConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkerPoolWorkerConfig
	} else {

		r.MachineType = res.MachineType

		r.DiskSizeGb = res.DiskSizeGb

		r.NoExternalIP = res.NoExternalIP

	}
	return nil
}

// This object is used to assert a desired state where this WorkerPoolWorkerConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkerPoolWorkerConfig *WorkerPoolWorkerConfig = &WorkerPoolWorkerConfig{empty: true}

func (r *WorkerPoolWorkerConfig) Empty() bool {
	return r.empty
}

func (r *WorkerPoolWorkerConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkerPoolWorkerConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkerPoolNetworkConfig struct {
	empty         bool    `json:"-"`
	PeeredNetwork *string `json:"peeredNetwork"`
}

type jsonWorkerPoolNetworkConfig WorkerPoolNetworkConfig

func (r *WorkerPoolNetworkConfig) UnmarshalJSON(data []byte) error {
	var res jsonWorkerPoolNetworkConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkerPoolNetworkConfig
	} else {

		r.PeeredNetwork = res.PeeredNetwork

	}
	return nil
}

// This object is used to assert a desired state where this WorkerPoolNetworkConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyWorkerPoolNetworkConfig *WorkerPoolNetworkConfig = &WorkerPoolNetworkConfig{empty: true}

func (r *WorkerPoolNetworkConfig) Empty() bool {
	return r.empty
}

func (r *WorkerPoolNetworkConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkerPoolNetworkConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *WorkerPool) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloud_build",
		Type:    "WorkerPool",
		Version: "beta",
	}
}

const WorkerPoolMaxPage = -1

type WorkerPoolList struct {
	Items []*WorkerPool

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *WorkerPoolList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkerPoolList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkerPool(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkerPool(ctx context.Context, project, location string) (*WorkerPoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListWorkerPoolWithMaxResults(ctx, project, location, WorkerPoolMaxPage)

}

func (c *Client) ListWorkerPoolWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*WorkerPoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listWorkerPool(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkerPoolList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *WorkerPool) URLNormalized() *WorkerPool {
	normalized := dcl.Copy(*r).(WorkerPool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetWorkerPool(ctx context.Context, r *WorkerPool) (*WorkerPool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getWorkerPoolRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkerPool(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkerPoolNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteWorkerPool(ctx context.Context, r *WorkerPool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("WorkerPool resource is nil")
	}
	c.Config.Logger.Info("Deleting WorkerPool...")
	deleteOp := deleteWorkerPoolOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkerPool deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkerPool(ctx context.Context, project, location string, filter func(*WorkerPool) bool) error {
	listObj, err := c.ListWorkerPool(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkerPool(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkerPool(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkerPool(ctx context.Context, rawDesired *WorkerPool, opts ...dcl.ApplyOption) (*WorkerPool, error) {
	var resultNewState *WorkerPool
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyWorkerPoolHelper(c, ctx, rawDesired, opts...)
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

func applyWorkerPoolHelper(c *Client, ctx context.Context, rawDesired *WorkerPool, opts ...dcl.ApplyOption) (*WorkerPool, error) {
	c.Config.Logger.Info("Beginning ApplyWorkerPool...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.workerPoolDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToWorkerPoolDiffs(c.Config, fieldDiffs, opts)
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
	var ops []workerPoolApiOperation
	if create {
		ops = append(ops, &createWorkerPoolOperation{})
	} else if recreate {
		ops = append(ops, &deleteWorkerPoolOperation{})
		ops = append(ops, &createWorkerPoolOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeWorkerPoolDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetWorkerPool(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createWorkerPoolOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapWorkerPool(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeWorkerPoolNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkerPoolNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkerPoolDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkerPool(c, newDesired, newState)
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
