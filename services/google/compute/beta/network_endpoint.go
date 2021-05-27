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
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type NetworkEndpoint struct {
	Port        *int64            `json:"port"`
	IPAddress   *string           `json:"ipAddress"`
	Fqdn        *string           `json:"fqdn"`
	Instance    *string           `json:"instance"`
	Annotations map[string]string `json:"annotations"`
	Project     *string           `json:"project"`
	Location    *string           `json:"location"`
	Group       *string           `json:"group"`
}

func (r *NetworkEndpoint) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *NetworkEndpoint) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "NetworkEndpoint",
		Version: "beta",
	}
}

const NetworkEndpointMaxPage = -1

type NetworkEndpointList struct {
	Items []*NetworkEndpoint

	nextToken string

	pageSize int32

	project string

	location string

	group string
}

func (l *NetworkEndpointList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NetworkEndpointList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNetworkEndpoint(ctx, l.project, l.location, l.group, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNetworkEndpoint(ctx context.Context, project, location, group string) (*NetworkEndpointList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListNetworkEndpointWithMaxResults(ctx, project, location, group, NetworkEndpointMaxPage)

}

func (c *Client) ListNetworkEndpointWithMaxResults(ctx context.Context, project, location, group string, pageSize int32) (*NetworkEndpointList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listNetworkEndpoint(ctx, project, location, group, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &NetworkEndpointList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,

		group: group,
	}, nil
}

func (c *Client) DeleteNetworkEndpoint(ctx context.Context, r *NetworkEndpoint) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("NetworkEndpoint resource is nil")
	}
	c.Config.Logger.Info("Deleting NetworkEndpoint...")
	deleteOp := deleteNetworkEndpointOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNetworkEndpoint deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNetworkEndpoint(ctx context.Context, project, location, group string, filter func(*NetworkEndpoint) bool) error {
	listObj, err := c.ListNetworkEndpoint(ctx, project, location, group)
	if err != nil {
		return err
	}

	err = c.deleteAllNetworkEndpoint(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNetworkEndpoint(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNetworkEndpoint(ctx context.Context, rawDesired *NetworkEndpoint, opts ...dcl.ApplyOption) (*NetworkEndpoint, error) {

	var resultNewState *NetworkEndpoint
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyNetworkEndpointHelper(c, ctx, rawDesired, opts...)
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

func applyNetworkEndpointHelper(c *Client, ctx context.Context, rawDesired *NetworkEndpoint, opts ...dcl.ApplyOption) (*NetworkEndpoint, error) {
	c.Config.Logger.Info("Beginning ApplyNetworkEndpoint...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.networkEndpointDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToNetworkEndpointOp(opStrings, fieldDiffs, opts)
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
	var ops []networkEndpointApiOperation
	if create {
		ops = append(ops, &createNetworkEndpointOperation{})
	} else if recreate {

		ops = append(ops, &deleteNetworkEndpointOperation{})

		ops = append(ops, &createNetworkEndpointOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeNetworkEndpointDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetNetworkEndpoint(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createNetworkEndpointOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapNetworkEndpoint(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeNetworkEndpointNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNetworkEndpointNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNetworkEndpointDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNetworkEndpoint(c, newDesired, newState)
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
