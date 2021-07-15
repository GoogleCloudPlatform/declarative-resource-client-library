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

type Firewall struct {
	CreationTimestamp     *string                `json:"creationTimestamp"`
	Description           *string                `json:"description"`
	Direction             *FirewallDirectionEnum `json:"direction"`
	Disabled              *bool                  `json:"disabled"`
	Id                    *string                `json:"id"`
	LogConfig             *FirewallLogConfig     `json:"logConfig"`
	Name                  *string                `json:"name"`
	Network               *string                `json:"network"`
	Priority              *int64                 `json:"priority"`
	SelfLink              *string                `json:"selfLink"`
	Project               *string                `json:"project"`
	Allowed               []FirewallAllowed      `json:"allowed"`
	Denied                []FirewallDenied       `json:"denied"`
	DestinationRanges     []string               `json:"destinationRanges"`
	SourceRanges          []string               `json:"sourceRanges"`
	SourceServiceAccounts []string               `json:"sourceServiceAccounts"`
	SourceTags            []string               `json:"sourceTags"`
	TargetServiceAccounts []string               `json:"targetServiceAccounts"`
	TargetTags            []string               `json:"targetTags"`
}

func (r *Firewall) String() string {
	return dcl.SprintResource(r)
}

// The enum FirewallDirectionEnum.
type FirewallDirectionEnum string

// FirewallDirectionEnumRef returns a *FirewallDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func FirewallDirectionEnumRef(s string) *FirewallDirectionEnum {
	if s == "" {
		return nil
	}

	v := FirewallDirectionEnum(s)
	return &v
}

func (v FirewallDirectionEnum) Validate() error {
	for _, s := range []string{"INGRESS", "EGRESS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FirewallDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FirewallLogConfig struct {
	empty  bool  `json:"-"`
	Enable *bool `json:"enable"`
}

type jsonFirewallLogConfig FirewallLogConfig

func (r *FirewallLogConfig) UnmarshalJSON(data []byte) error {
	var res jsonFirewallLogConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFirewallLogConfig
	} else {

		r.Enable = res.Enable

	}
	return nil
}

// This object is used to assert a desired state where this FirewallLogConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFirewallLogConfig *FirewallLogConfig = &FirewallLogConfig{empty: true}

func (r *FirewallLogConfig) Empty() bool {
	return r.empty
}

func (r *FirewallLogConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *FirewallLogConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FirewallAllowed struct {
	empty         bool     `json:"-"`
	IPProtocol    *string  `json:"ipProtocol"`
	Ports         []string `json:"ports"`
	IPProtocolAlt []string `json:"ipProtocolAlt"`
}

type jsonFirewallAllowed FirewallAllowed

func (r *FirewallAllowed) UnmarshalJSON(data []byte) error {
	var res jsonFirewallAllowed
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFirewallAllowed
	} else {

		r.IPProtocol = res.IPProtocol

		r.Ports = res.Ports

		r.IPProtocolAlt = res.IPProtocolAlt

	}
	return nil
}

// This object is used to assert a desired state where this FirewallAllowed is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFirewallAllowed *FirewallAllowed = &FirewallAllowed{empty: true}

func (r *FirewallAllowed) Empty() bool {
	return r.empty
}

func (r *FirewallAllowed) String() string {
	return dcl.SprintResource(r)
}

func (r *FirewallAllowed) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FirewallDenied struct {
	empty         bool     `json:"-"`
	IPProtocol    *string  `json:"ipProtocol"`
	Ports         []string `json:"ports"`
	IPProtocolAlt []string `json:"ipProtocolAlt"`
}

type jsonFirewallDenied FirewallDenied

func (r *FirewallDenied) UnmarshalJSON(data []byte) error {
	var res jsonFirewallDenied
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFirewallDenied
	} else {

		r.IPProtocol = res.IPProtocol

		r.Ports = res.Ports

		r.IPProtocolAlt = res.IPProtocolAlt

	}
	return nil
}

// This object is used to assert a desired state where this FirewallDenied is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyFirewallDenied *FirewallDenied = &FirewallDenied{empty: true}

func (r *FirewallDenied) Empty() bool {
	return r.empty
}

func (r *FirewallDenied) String() string {
	return dcl.SprintResource(r)
}

func (r *FirewallDenied) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Firewall) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Firewall",
		Version: "beta",
	}
}

const FirewallMaxPage = -1

type FirewallList struct {
	Items []*Firewall

	nextToken string

	pageSize int32

	project string
}

func (l *FirewallList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FirewallList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFirewall(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFirewall(ctx context.Context, project string) (*FirewallList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListFirewallWithMaxResults(ctx, project, FirewallMaxPage)

}

func (c *Client) ListFirewallWithMaxResults(ctx context.Context, project string, pageSize int32) (*FirewallList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listFirewall(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FirewallList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Firewall) URLNormalized() *Firewall {
	normalized := dcl.Copy(*r).(Firewall)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetFirewall(ctx context.Context, r *Firewall) (*Firewall, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getFirewallRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFirewall(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.Priority) {
		result.Priority = dcl.Int64(1000)
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFirewallNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFirewall(ctx context.Context, r *Firewall) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Firewall resource is nil")
	}
	c.Config.Logger.Info("Deleting Firewall...")
	deleteOp := deleteFirewallOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFirewall deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFirewall(ctx context.Context, project string, filter func(*Firewall) bool) error {
	listObj, err := c.ListFirewall(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllFirewall(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFirewall(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFirewall(ctx context.Context, rawDesired *Firewall, opts ...dcl.ApplyOption) (*Firewall, error) {

	var resultNewState *Firewall
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFirewallHelper(c, ctx, rawDesired, opts...)
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

func applyFirewallHelper(c *Client, ctx context.Context, rawDesired *Firewall, opts ...dcl.ApplyOption) (*Firewall, error) {
	c.Config.Logger.Info("Beginning ApplyFirewall...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.firewallDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFirewallDiffs(c.Config, fieldDiffs, opts)
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
	var ops []firewallApiOperation
	if create {
		ops = append(ops, &createFirewallOperation{})
	} else if recreate {
		ops = append(ops, &deleteFirewallOperation{})
		ops = append(ops, &createFirewallOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeFirewallDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetFirewall(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFirewallOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFirewall(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFirewallNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFirewallNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFirewallDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFirewall(c, newDesired, newState)
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
