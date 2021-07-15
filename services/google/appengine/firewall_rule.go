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
package appengine

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type FirewallRule struct {
	Action      *FirewallRuleActionEnum `json:"action"`
	Description *string                 `json:"description"`
	Priority    *int64                  `json:"priority"`
	SourceRange *string                 `json:"sourceRange"`
	App         *string                 `json:"app"`
}

func (r *FirewallRule) String() string {
	return dcl.SprintResource(r)
}

// The enum FirewallRuleActionEnum.
type FirewallRuleActionEnum string

// FirewallRuleActionEnumRef returns a *FirewallRuleActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func FirewallRuleActionEnumRef(s string) *FirewallRuleActionEnum {
	if s == "" {
		return nil
	}

	v := FirewallRuleActionEnum(s)
	return &v
}

func (v FirewallRuleActionEnum) Validate() error {
	for _, s := range []string{"UNSPECIFIED_ACTION", "ALLOW", "DENY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FirewallRuleActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *FirewallRule) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "app_engine",
		Type:    "FirewallRule",
		Version: "appengine",
	}
}

const FirewallRuleMaxPage = -1

type FirewallRuleList struct {
	Items []*FirewallRule

	nextToken string

	pageSize int32

	app string
}

func (l *FirewallRuleList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FirewallRuleList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFirewallRule(ctx, l.app, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFirewallRule(ctx context.Context, app string) (*FirewallRuleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListFirewallRuleWithMaxResults(ctx, app, FirewallRuleMaxPage)

}

func (c *Client) ListFirewallRuleWithMaxResults(ctx context.Context, app string, pageSize int32) (*FirewallRuleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listFirewallRule(ctx, app, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FirewallRuleList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		app: app,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *FirewallRule) URLNormalized() *FirewallRule {
	normalized := dcl.Copy(*r).(FirewallRule)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceRange = dcl.SelfLinkToName(r.SourceRange)
	normalized.App = dcl.SelfLinkToName(r.App)
	return &normalized
}

func (c *Client) GetFirewallRule(ctx context.Context, r *FirewallRule) (*FirewallRule, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getFirewallRuleRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFirewallRule(b, c)
	if err != nil {
		return nil, err
	}
	result.App = r.App
	result.Priority = r.Priority

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFirewallRuleNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFirewallRule(ctx context.Context, r *FirewallRule) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("FirewallRule resource is nil")
	}
	c.Config.Logger.Info("Deleting FirewallRule...")
	deleteOp := deleteFirewallRuleOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFirewallRule deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFirewallRule(ctx context.Context, app string, filter func(*FirewallRule) bool) error {
	listObj, err := c.ListFirewallRule(ctx, app)
	if err != nil {
		return err
	}

	err = c.deleteAllFirewallRule(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFirewallRule(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFirewallRule(ctx context.Context, rawDesired *FirewallRule, opts ...dcl.ApplyOption) (*FirewallRule, error) {
	var resultNewState *FirewallRule
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFirewallRuleHelper(c, ctx, rawDesired, opts...)
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

func applyFirewallRuleHelper(c *Client, ctx context.Context, rawDesired *FirewallRule, opts ...dcl.ApplyOption) (*FirewallRule, error) {
	c.Config.Logger.Info("Beginning ApplyFirewallRule...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.firewallRuleDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFirewallRuleDiffs(c.Config, fieldDiffs, opts)
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
	var ops []firewallRuleApiOperation
	if create {
		ops = append(ops, &createFirewallRuleOperation{})
	} else if recreate {
		ops = append(ops, &deleteFirewallRuleOperation{})
		ops = append(ops, &createFirewallRuleOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeFirewallRuleDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetFirewallRule(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFirewallRuleOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFirewallRule(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFirewallRuleNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFirewallRuleNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFirewallRuleDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFirewallRule(c, newDesired, newState)
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
