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
package compute

import (
	"context"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type HttpsHealthCheck struct {
	CheckIntervalSec   *int64  `json:"checkIntervalSec"`
	Description        *string `json:"description"`
	HealthyThreshold   *int64  `json:"healthyThreshold"`
	Host               *string `json:"host"`
	Name               *string `json:"name"`
	Port               *int64  `json:"port"`
	RequestPath        *string `json:"requestPath"`
	TimeoutSec         *int64  `json:"timeoutSec"`
	UnhealthyThreshold *int64  `json:"unhealthyThreshold"`
	Project            *string `json:"project"`
	SelfLink           *string `json:"selfLink"`
	CreationTimestamp  *string `json:"creationTimestamp"`
}

func (r *HttpsHealthCheck) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *HttpsHealthCheck) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "HttpsHealthCheck",
		Version: "compute",
	}
}

const HttpsHealthCheckMaxPage = -1

type HttpsHealthCheckList struct {
	Items []*HttpsHealthCheck

	nextToken string

	pageSize int32

	project string
}

func (l *HttpsHealthCheckList) HasNext() bool {
	return l.nextToken != ""
}

func (l *HttpsHealthCheckList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listHttpsHealthCheck(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListHttpsHealthCheck(ctx context.Context, project string) (*HttpsHealthCheckList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListHttpsHealthCheckWithMaxResults(ctx, project, HttpsHealthCheckMaxPage)

}

func (c *Client) ListHttpsHealthCheckWithMaxResults(ctx context.Context, project string, pageSize int32) (*HttpsHealthCheckList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listHttpsHealthCheck(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &HttpsHealthCheckList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetHttpsHealthCheck(ctx context.Context, r *HttpsHealthCheck) (*HttpsHealthCheck, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getHttpsHealthCheckRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalHttpsHealthCheck(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.CheckIntervalSec) {
		result.CheckIntervalSec = dcl.Int64(5)
	}
	if dcl.IsZeroValue(result.HealthyThreshold) {
		result.HealthyThreshold = dcl.Int64(2)
	}
	if dcl.IsZeroValue(result.Port) {
		result.Port = dcl.Int64(443)
	}
	if dcl.IsZeroValue(result.RequestPath) {
		result.RequestPath = dcl.String("/")
	}
	if dcl.IsZeroValue(result.TimeoutSec) {
		result.TimeoutSec = dcl.Int64(5)
	}
	if dcl.IsZeroValue(result.UnhealthyThreshold) {
		result.UnhealthyThreshold = dcl.Int64(2)
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeHttpsHealthCheckNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteHttpsHealthCheck(ctx context.Context, r *HttpsHealthCheck) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("HttpsHealthCheck resource is nil")
	}
	c.Config.Logger.Info("Deleting HttpsHealthCheck...")
	deleteOp := deleteHttpsHealthCheckOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllHttpsHealthCheck deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllHttpsHealthCheck(ctx context.Context, project string, filter func(*HttpsHealthCheck) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListHttpsHealthCheck(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllHttpsHealthCheck(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllHttpsHealthCheck(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyHttpsHealthCheck(ctx context.Context, rawDesired *HttpsHealthCheck, opts ...dcl.ApplyOption) (*HttpsHealthCheck, error) {
	c.Config.Logger.Info("Beginning ApplyHttpsHealthCheck...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.httpsHealthCheckDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []httpsHealthCheckApiOperation
	if create {
		ops = append(ops, &createHttpsHealthCheckOperation{})
	} else if recreate {

		ops = append(ops, &deleteHttpsHealthCheckOperation{})

		ops = append(ops, &createHttpsHealthCheckOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeHttpsHealthCheckDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetHttpsHealthCheck(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createHttpsHealthCheckOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapHttpsHealthCheck(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeHttpsHealthCheckNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeHttpsHealthCheckNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeHttpsHealthCheckDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffHttpsHealthCheck(c, newDesired, newState)
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
