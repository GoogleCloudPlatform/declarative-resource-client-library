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
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type FirewallPolicyRule struct {
	Description           *string                          `json:"description"`
	Priority              *int64                           `json:"priority"`
	Match                 *FirewallPolicyRuleMatch         `json:"match"`
	Action                *string                          `json:"action"`
	Direction             *FirewallPolicyRuleDirectionEnum `json:"direction"`
	TargetResources       []string                         `json:"targetResources"`
	EnableLogging         *bool                            `json:"enableLogging"`
	RuleTupleCount        *int64                           `json:"ruleTupleCount"`
	TargetServiceAccounts []string                         `json:"targetServiceAccounts"`
	TargetSecureLabels    []string                         `json:"targetSecureLabels"`
	Disabled              *bool                            `json:"disabled"`
	Kind                  *string                          `json:"kind"`
	FirewallPolicy        *string                          `json:"firewallPolicy"`
}

func (r *FirewallPolicyRule) String() string {
	return dcl.SprintResource(r)
}

// The enum FirewallPolicyRuleDirectionEnum.
type FirewallPolicyRuleDirectionEnum string

// FirewallPolicyRuleDirectionEnumRef returns a *FirewallPolicyRuleDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func FirewallPolicyRuleDirectionEnumRef(s string) *FirewallPolicyRuleDirectionEnum {
	if s == "" {
		return nil
	}

	v := FirewallPolicyRuleDirectionEnum(s)
	return &v
}

func (v FirewallPolicyRuleDirectionEnum) Validate() error {
	for _, s := range []string{"INGRESS", "EGRESS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FirewallPolicyRuleDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FirewallPolicyRuleMatch struct {
	empty           bool                                   `json:"-"`
	SrcIPRanges     []string                               `json:"srcIPRanges"`
	DestIPRanges    []string                               `json:"destIPRanges"`
	Layer4Configs   []FirewallPolicyRuleMatchLayer4Configs `json:"layer4Configs"`
	SrcSecureLabels []string                               `json:"srcSecureLabels"`
}

type jsonFirewallPolicyRuleMatch FirewallPolicyRuleMatch

func (r *FirewallPolicyRuleMatch) UnmarshalJSON(data []byte) error {
	var res jsonFirewallPolicyRuleMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFirewallPolicyRuleMatch
	} else {

		r.SrcIPRanges = res.SrcIPRanges

		r.DestIPRanges = res.DestIPRanges

		r.Layer4Configs = res.Layer4Configs

		r.SrcSecureLabels = res.SrcSecureLabels

	}
	return nil
}

// This object is used to assert a desired state where this FirewallPolicyRuleMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFirewallPolicyRuleMatch *FirewallPolicyRuleMatch = &FirewallPolicyRuleMatch{empty: true}

func (r *FirewallPolicyRuleMatch) Empty() bool {
	return r.empty
}

func (r *FirewallPolicyRuleMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *FirewallPolicyRuleMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FirewallPolicyRuleMatchLayer4Configs struct {
	empty      bool     `json:"-"`
	IPProtocol *string  `json:"ipProtocol"`
	Ports      []string `json:"ports"`
}

type jsonFirewallPolicyRuleMatchLayer4Configs FirewallPolicyRuleMatchLayer4Configs

func (r *FirewallPolicyRuleMatchLayer4Configs) UnmarshalJSON(data []byte) error {
	var res jsonFirewallPolicyRuleMatchLayer4Configs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFirewallPolicyRuleMatchLayer4Configs
	} else {

		r.IPProtocol = res.IPProtocol

		r.Ports = res.Ports

	}
	return nil
}

// This object is used to assert a desired state where this FirewallPolicyRuleMatchLayer4Configs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFirewallPolicyRuleMatchLayer4Configs *FirewallPolicyRuleMatchLayer4Configs = &FirewallPolicyRuleMatchLayer4Configs{empty: true}

func (r *FirewallPolicyRuleMatchLayer4Configs) Empty() bool {
	return r.empty
}

func (r *FirewallPolicyRuleMatchLayer4Configs) String() string {
	return dcl.SprintResource(r)
}

func (r *FirewallPolicyRuleMatchLayer4Configs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *FirewallPolicyRule) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "FirewallPolicyRule",
		Version: "compute",
	}
}

const FirewallPolicyRuleMaxPage = -1

type FirewallPolicyRuleList struct {
	Items []*FirewallPolicyRule

	nextToken string

	pageSize int32

	firewallPolicy string
}

func (l *FirewallPolicyRuleList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FirewallPolicyRuleList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFirewallPolicyRule(ctx, l.firewallPolicy, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFirewallPolicyRule(ctx context.Context, firewallPolicy string) (*FirewallPolicyRuleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListFirewallPolicyRuleWithMaxResults(ctx, firewallPolicy, FirewallPolicyRuleMaxPage)

}

func (c *Client) ListFirewallPolicyRuleWithMaxResults(ctx context.Context, firewallPolicy string, pageSize int32) (*FirewallPolicyRuleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listFirewallPolicyRule(ctx, firewallPolicy, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FirewallPolicyRuleList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		firewallPolicy: firewallPolicy,
	}, nil
}

func (c *Client) GetFirewallPolicyRule(ctx context.Context, r *FirewallPolicyRule) (*FirewallPolicyRule, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getFirewallPolicyRuleRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 400) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFirewallPolicyRule(b, c)
	if err != nil {
		return nil, err
	}
	result.FirewallPolicy = r.FirewallPolicy
	result.Priority = r.Priority

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFirewallPolicyRuleNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFirewallPolicyRule(ctx context.Context, r *FirewallPolicyRule) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("FirewallPolicyRule resource is nil")
	}
	c.Config.Logger.Info("Deleting FirewallPolicyRule...")
	deleteOp := deleteFirewallPolicyRuleOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFirewallPolicyRule deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFirewallPolicyRule(ctx context.Context, firewallPolicy string, filter func(*FirewallPolicyRule) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListFirewallPolicyRule(ctx, firewallPolicy)
	if err != nil {
		return err
	}

	err = c.deleteAllFirewallPolicyRule(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFirewallPolicyRule(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFirewallPolicyRule(ctx context.Context, rawDesired *FirewallPolicyRule, opts ...dcl.ApplyOption) (*FirewallPolicyRule, error) {
	c.Config.Logger.Info("Beginning ApplyFirewallPolicyRule...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.firewallPolicyRuleDiffsForRawDesired(ctx, rawDesired, opts...)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []firewallPolicyRuleApiOperation
	if create {
		ops = append(ops, &createFirewallPolicyRuleOperation{})
	} else if recreate {

		ops = append(ops, &deleteFirewallPolicyRuleOperation{})

		ops = append(ops, &createFirewallPolicyRuleOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeFirewallPolicyRuleDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetFirewallPolicyRule(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFirewallPolicyRuleOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFirewallPolicyRule(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFirewallPolicyRuleNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFirewallPolicyRuleNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFirewallPolicyRuleDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFirewallPolicyRule(c, newDesired, newState)
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
