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

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type TargetVpnGateway struct {
	Id             *int64                      `json:"id"`
	Name           *string                     `json:"name"`
	Description    *string                     `json:"description"`
	Region         *string                     `json:"region"`
	Network        *string                     `json:"network"`
	Tunnel         []string                    `json:"tunnel"`
	Status         *TargetVpnGatewayStatusEnum `json:"status"`
	SelfLink       *string                     `json:"selfLink"`
	ForwardingRule []string                    `json:"forwardingRule"`
	Project        *string                     `json:"project"`
}

func (r *TargetVpnGateway) String() string {
	return dcl.SprintResource(r)
}

// The enum TargetVpnGatewayStatusEnum.
type TargetVpnGatewayStatusEnum string

// TargetVpnGatewayStatusEnumRef returns a *TargetVpnGatewayStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func TargetVpnGatewayStatusEnumRef(s string) *TargetVpnGatewayStatusEnum {
	if s == "" {
		return nil
	}

	v := TargetVpnGatewayStatusEnum(s)
	return &v
}

func (v TargetVpnGatewayStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "RUNNING", "DONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TargetVpnGatewayStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *TargetVpnGateway) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "TargetVpnGateway",
		Version: "beta",
	}
}

const TargetVpnGatewayMaxPage = -1

type TargetVpnGatewayList struct {
	Items []*TargetVpnGateway

	nextToken string

	pageSize int32

	project string

	region string
}

func (l *TargetVpnGatewayList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TargetVpnGatewayList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTargetVpnGateway(ctx, l.project, l.region, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTargetVpnGateway(ctx context.Context, project, region string) (*TargetVpnGatewayList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTargetVpnGatewayWithMaxResults(ctx, project, region, TargetVpnGatewayMaxPage)

}

func (c *Client) ListTargetVpnGatewayWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*TargetVpnGatewayList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listTargetVpnGateway(ctx, project, region, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TargetVpnGatewayList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		region: region,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *TargetVpnGateway) URLNormalized() *TargetVpnGateway {
	normalized := dcl.Copy(*r).(TargetVpnGateway)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetTargetVpnGateway(ctx context.Context, r *TargetVpnGateway) (*TargetVpnGateway, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getTargetVpnGatewayRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTargetVpnGateway(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTargetVpnGatewayNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTargetVpnGateway(ctx context.Context, r *TargetVpnGateway) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("TargetVpnGateway resource is nil")
	}
	c.Config.Logger.Info("Deleting TargetVpnGateway...")
	deleteOp := deleteTargetVpnGatewayOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTargetVpnGateway deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTargetVpnGateway(ctx context.Context, project, region string, filter func(*TargetVpnGateway) bool) error {
	listObj, err := c.ListTargetVpnGateway(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllTargetVpnGateway(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTargetVpnGateway(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTargetVpnGateway(ctx context.Context, rawDesired *TargetVpnGateway, opts ...dcl.ApplyOption) (*TargetVpnGateway, error) {
	var resultNewState *TargetVpnGateway
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTargetVpnGatewayHelper(c, ctx, rawDesired, opts...)
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

func applyTargetVpnGatewayHelper(c *Client, ctx context.Context, rawDesired *TargetVpnGateway, opts ...dcl.ApplyOption) (*TargetVpnGateway, error) {
	c.Config.Logger.Info("Beginning ApplyTargetVpnGateway...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.targetVpnGatewayDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTargetVpnGatewayDiffs(c.Config, fieldDiffs, opts)
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
	var ops []targetVpnGatewayApiOperation
	if create {
		ops = append(ops, &createTargetVpnGatewayOperation{})
	} else if recreate {
		ops = append(ops, &deleteTargetVpnGatewayOperation{})
		ops = append(ops, &createTargetVpnGatewayOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTargetVpnGatewayDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetTargetVpnGateway(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTargetVpnGatewayOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTargetVpnGateway(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTargetVpnGatewayNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTargetVpnGatewayNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTargetVpnGatewayDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTargetVpnGateway(c, newDesired, newState)
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
