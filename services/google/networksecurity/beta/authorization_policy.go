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
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type AuthorizationPolicy struct {
	Name        *string                        `json:"name"`
	Description *string                        `json:"description"`
	CreateTime  *string                        `json:"createTime"`
	UpdateTime  *string                        `json:"updateTime"`
	Labels      map[string]string              `json:"labels"`
	Action      *AuthorizationPolicyActionEnum `json:"action"`
	Rules       []AuthorizationPolicyRules     `json:"rules"`
	Project     *string                        `json:"project"`
	Location    *string                        `json:"location"`
}

func (r *AuthorizationPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum AuthorizationPolicyActionEnum.
type AuthorizationPolicyActionEnum string

// AuthorizationPolicyActionEnumRef returns a *AuthorizationPolicyActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func AuthorizationPolicyActionEnumRef(s string) *AuthorizationPolicyActionEnum {
	if s == "" {
		return nil
	}

	v := AuthorizationPolicyActionEnum(s)
	return &v
}

func (v AuthorizationPolicyActionEnum) Validate() error {
	for _, s := range []string{"ACTION_UNSPECIFIED", "ALLOW", "DENY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AuthorizationPolicyActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AuthorizationPolicyRules struct {
	empty        bool                                   `json:"-"`
	Sources      []AuthorizationPolicyRulesSources      `json:"sources"`
	Destinations []AuthorizationPolicyRulesDestinations `json:"destinations"`
}

type jsonAuthorizationPolicyRules AuthorizationPolicyRules

func (r *AuthorizationPolicyRules) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRules
	} else {

		r.Sources = res.Sources

		r.Destinations = res.Destinations

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRules is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRules *AuthorizationPolicyRules = &AuthorizationPolicyRules{empty: true}

func (r *AuthorizationPolicyRules) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRules) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AuthorizationPolicyRulesSources struct {
	empty      bool     `json:"-"`
	Principals []string `json:"principals"`
	IPBlocks   []string `json:"ipBlocks"`
}

type jsonAuthorizationPolicyRulesSources AuthorizationPolicyRulesSources

func (r *AuthorizationPolicyRulesSources) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRulesSources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRulesSources
	} else {

		r.Principals = res.Principals

		r.IPBlocks = res.IPBlocks

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRulesSources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRulesSources *AuthorizationPolicyRulesSources = &AuthorizationPolicyRulesSources{empty: true}

func (r *AuthorizationPolicyRulesSources) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRulesSources) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRulesSources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AuthorizationPolicyRulesDestinations struct {
	empty           bool                                                 `json:"-"`
	Hosts           []string                                             `json:"hosts"`
	Ports           []int64                                              `json:"ports"`
	Methods         []string                                             `json:"methods"`
	HttpHeaderMatch *AuthorizationPolicyRulesDestinationsHttpHeaderMatch `json:"httpHeaderMatch"`
}

type jsonAuthorizationPolicyRulesDestinations AuthorizationPolicyRulesDestinations

func (r *AuthorizationPolicyRulesDestinations) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRulesDestinations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRulesDestinations
	} else {

		r.Hosts = res.Hosts

		r.Ports = res.Ports

		r.Methods = res.Methods

		r.HttpHeaderMatch = res.HttpHeaderMatch

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRulesDestinations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRulesDestinations *AuthorizationPolicyRulesDestinations = &AuthorizationPolicyRulesDestinations{empty: true}

func (r *AuthorizationPolicyRulesDestinations) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRulesDestinations) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRulesDestinations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AuthorizationPolicyRulesDestinationsHttpHeaderMatch struct {
	empty      bool    `json:"-"`
	HeaderName *string `json:"headerName"`
	RegexMatch *string `json:"regexMatch"`
}

type jsonAuthorizationPolicyRulesDestinationsHttpHeaderMatch AuthorizationPolicyRulesDestinationsHttpHeaderMatch

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRulesDestinationsHttpHeaderMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch
	} else {

		r.HeaderName = res.HeaderName

		r.RegexMatch = res.RegexMatch

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRulesDestinationsHttpHeaderMatch is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch *AuthorizationPolicyRulesDestinationsHttpHeaderMatch = &AuthorizationPolicyRulesDestinationsHttpHeaderMatch{empty: true}

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AuthorizationPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_security",
		Type:    "AuthorizationPolicy",
		Version: "beta",
	}
}

const AuthorizationPolicyMaxPage = -1

type AuthorizationPolicyList struct {
	Items []*AuthorizationPolicy

	nextToken string

	pageSize int32

	resource *AuthorizationPolicy
}

func (l *AuthorizationPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AuthorizationPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAuthorizationPolicy(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAuthorizationPolicy(ctx context.Context, r *AuthorizationPolicy) (*AuthorizationPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAuthorizationPolicyWithMaxResults(ctx, r, AuthorizationPolicyMaxPage)

}

func (c *Client) ListAuthorizationPolicyWithMaxResults(ctx context.Context, r *AuthorizationPolicy, pageSize int32) (*AuthorizationPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAuthorizationPolicy(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AuthorizationPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAuthorizationPolicy(ctx context.Context, r *AuthorizationPolicy) (*AuthorizationPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAuthorizationPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAuthorizationPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAuthorizationPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAuthorizationPolicy(ctx context.Context, r *AuthorizationPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AuthorizationPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting AuthorizationPolicy...")
	deleteOp := deleteAuthorizationPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAuthorizationPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAuthorizationPolicy(ctx context.Context, project, location string, filter func(*AuthorizationPolicy) bool) error {
	r := &AuthorizationPolicy{

		Project: &project,

		Location: &location,
	}
	listObj, err := c.ListAuthorizationPolicy(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllAuthorizationPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAuthorizationPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAuthorizationPolicy(ctx context.Context, rawDesired *AuthorizationPolicy, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {
	var resultNewState *AuthorizationPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAuthorizationPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyAuthorizationPolicyHelper(c *Client, ctx context.Context, rawDesired *AuthorizationPolicy, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {
	c.Config.Logger.Info("Beginning ApplyAuthorizationPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.authorizationPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAuthorizationPolicyDiffs(c.Config, fieldDiffs, opts)
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
	var ops []authorizationPolicyApiOperation
	if create {
		ops = append(ops, &createAuthorizationPolicyOperation{})
	} else if recreate {
		ops = append(ops, &deleteAuthorizationPolicyOperation{})
		ops = append(ops, &createAuthorizationPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAuthorizationPolicyDesiredState(rawDesired, nil)
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
		c.Config.Logger.Infof("Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetAuthorizationPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAuthorizationPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAuthorizationPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAuthorizationPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAuthorizationPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAuthorizationPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAuthorizationPolicy(c, newDesired, newState)
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
func (r *AuthorizationPolicy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "POST", body, nil
}
