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
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type FirewallPolicyAssociation struct {
	Name             *string `json:"name"`
	AttachmentTarget *string `json:"attachmentTarget"`
	FirewallPolicy   *string `json:"firewallPolicy"`
	ShortName        *string `json:"shortName"`
}

func (r *FirewallPolicyAssociation) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *FirewallPolicyAssociation) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "FirewallPolicyAssociation",
		Version: "compute",
	}
}

const FirewallPolicyAssociationMaxPage = -1

type FirewallPolicyAssociationList struct {
	Items []*FirewallPolicyAssociation

	nextToken string

	pageSize int32

	resource *FirewallPolicyAssociation
}

func (l *FirewallPolicyAssociationList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FirewallPolicyAssociationList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFirewallPolicyAssociation(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFirewallPolicyAssociation(ctx context.Context, r *FirewallPolicyAssociation) (*FirewallPolicyAssociationList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListFirewallPolicyAssociationWithMaxResults(ctx, r, FirewallPolicyAssociationMaxPage)

}

func (c *Client) ListFirewallPolicyAssociationWithMaxResults(ctx context.Context, r *FirewallPolicyAssociation, pageSize int32) (*FirewallPolicyAssociationList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listFirewallPolicyAssociation(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FirewallPolicyAssociationList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetFirewallPolicyAssociation(ctx context.Context, r *FirewallPolicyAssociation) (*FirewallPolicyAssociation, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getFirewallPolicyAssociationRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 400) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFirewallPolicyAssociation(b, c)
	if err != nil {
		return nil, err
	}
	result.FirewallPolicy = r.FirewallPolicy
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFirewallPolicyAssociationNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFirewallPolicyAssociation(ctx context.Context, r *FirewallPolicyAssociation) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("FirewallPolicyAssociation resource is nil")
	}
	c.Config.Logger.Info("Deleting FirewallPolicyAssociation...")
	deleteOp := deleteFirewallPolicyAssociationOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFirewallPolicyAssociation deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFirewallPolicyAssociation(ctx context.Context, r *FirewallPolicyAssociation, filter func(*FirewallPolicyAssociation) bool) error {
	listObj, err := c.ListFirewallPolicyAssociation(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllFirewallPolicyAssociation(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFirewallPolicyAssociation(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFirewallPolicyAssociation(ctx context.Context, rawDesired *FirewallPolicyAssociation, opts ...dcl.ApplyOption) (*FirewallPolicyAssociation, error) {
	var resultNewState *FirewallPolicyAssociation
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFirewallPolicyAssociationHelper(c, ctx, rawDesired, opts...)
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

func applyFirewallPolicyAssociationHelper(c *Client, ctx context.Context, rawDesired *FirewallPolicyAssociation, opts ...dcl.ApplyOption) (*FirewallPolicyAssociation, error) {
	c.Config.Logger.Info("Beginning ApplyFirewallPolicyAssociation...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.firewallPolicyAssociationDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFirewallPolicyAssociationDiffs(c.Config, fieldDiffs, opts)
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
	var ops []firewallPolicyAssociationApiOperation
	if create {
		ops = append(ops, &createFirewallPolicyAssociationOperation{})
	} else if recreate {
		ops = append(ops, &deleteFirewallPolicyAssociationOperation{})
		ops = append(ops, &createFirewallPolicyAssociationOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeFirewallPolicyAssociationDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetFirewallPolicyAssociation(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFirewallPolicyAssociationOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFirewallPolicyAssociation(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFirewallPolicyAssociationNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFirewallPolicyAssociationNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFirewallPolicyAssociationDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFirewallPolicyAssociation(c, newDesired, newState)
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
