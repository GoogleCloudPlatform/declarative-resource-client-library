// Copyright 2025 Google LLC. All Rights Reserved.
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
package networkservices

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type TlsRoute struct {
	Name        *string         `json:"name"`
	SelfLink    *string         `json:"selfLink"`
	CreateTime  *string         `json:"createTime"`
	UpdateTime  *string         `json:"updateTime"`
	Description *string         `json:"description"`
	Rules       []TlsRouteRules `json:"rules"`
	Meshes      []string        `json:"meshes"`
	Gateways    []string        `json:"gateways"`
	Project     *string         `json:"project"`
	Location    *string         `json:"location"`
}

func (r *TlsRoute) String() string {
	return dcl.SprintResource(r)
}

type TlsRouteRules struct {
	empty   bool                   `json:"-"`
	Matches []TlsRouteRulesMatches `json:"matches"`
	Action  *TlsRouteRulesAction   `json:"action"`
}

type jsonTlsRouteRules TlsRouteRules

func (r *TlsRouteRules) UnmarshalJSON(data []byte) error {
	var res jsonTlsRouteRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTlsRouteRules
	} else {

		r.Matches = res.Matches

		r.Action = res.Action

	}
	return nil
}

// This object is used to assert a desired state where this TlsRouteRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTlsRouteRules *TlsRouteRules = &TlsRouteRules{empty: true}

func (r *TlsRouteRules) Empty() bool {
	return r.empty
}

func (r *TlsRouteRules) String() string {
	return dcl.SprintResource(r)
}

func (r *TlsRouteRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TlsRouteRulesMatches struct {
	empty   bool     `json:"-"`
	SniHost []string `json:"sniHost"`
	Alpn    []string `json:"alpn"`
}

type jsonTlsRouteRulesMatches TlsRouteRulesMatches

func (r *TlsRouteRulesMatches) UnmarshalJSON(data []byte) error {
	var res jsonTlsRouteRulesMatches
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTlsRouteRulesMatches
	} else {

		r.SniHost = res.SniHost

		r.Alpn = res.Alpn

	}
	return nil
}

// This object is used to assert a desired state where this TlsRouteRulesMatches is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTlsRouteRulesMatches *TlsRouteRulesMatches = &TlsRouteRulesMatches{empty: true}

func (r *TlsRouteRulesMatches) Empty() bool {
	return r.empty
}

func (r *TlsRouteRulesMatches) String() string {
	return dcl.SprintResource(r)
}

func (r *TlsRouteRulesMatches) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TlsRouteRulesAction struct {
	empty        bool                              `json:"-"`
	Destinations []TlsRouteRulesActionDestinations `json:"destinations"`
}

type jsonTlsRouteRulesAction TlsRouteRulesAction

func (r *TlsRouteRulesAction) UnmarshalJSON(data []byte) error {
	var res jsonTlsRouteRulesAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTlsRouteRulesAction
	} else {

		r.Destinations = res.Destinations

	}
	return nil
}

// This object is used to assert a desired state where this TlsRouteRulesAction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTlsRouteRulesAction *TlsRouteRulesAction = &TlsRouteRulesAction{empty: true}

func (r *TlsRouteRulesAction) Empty() bool {
	return r.empty
}

func (r *TlsRouteRulesAction) String() string {
	return dcl.SprintResource(r)
}

func (r *TlsRouteRulesAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TlsRouteRulesActionDestinations struct {
	empty       bool    `json:"-"`
	ServiceName *string `json:"serviceName"`
	Weight      *int64  `json:"weight"`
}

type jsonTlsRouteRulesActionDestinations TlsRouteRulesActionDestinations

func (r *TlsRouteRulesActionDestinations) UnmarshalJSON(data []byte) error {
	var res jsonTlsRouteRulesActionDestinations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTlsRouteRulesActionDestinations
	} else {

		r.ServiceName = res.ServiceName

		r.Weight = res.Weight

	}
	return nil
}

// This object is used to assert a desired state where this TlsRouteRulesActionDestinations is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTlsRouteRulesActionDestinations *TlsRouteRulesActionDestinations = &TlsRouteRulesActionDestinations{empty: true}

func (r *TlsRouteRulesActionDestinations) Empty() bool {
	return r.empty
}

func (r *TlsRouteRulesActionDestinations) String() string {
	return dcl.SprintResource(r)
}

func (r *TlsRouteRulesActionDestinations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *TlsRoute) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "TlsRoute",
		Version: "networkservices",
	}
}

func (r *TlsRoute) ID() (string, error) {
	if err := extractTlsRouteFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"self_link":   dcl.ValueOrEmptyString(nr.SelfLink),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time": dcl.ValueOrEmptyString(nr.UpdateTime),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"rules":       dcl.ValueOrEmptyString(nr.Rules),
		"meshes":      dcl.ValueOrEmptyString(nr.Meshes),
		"gateways":    dcl.ValueOrEmptyString(nr.Gateways),
		"project":     dcl.ValueOrEmptyString(nr.Project),
		"location":    dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}", params), nil
}

const TlsRouteMaxPage = -1

type TlsRouteList struct {
	Items []*TlsRoute

	nextToken string

	pageSize int32

	resource *TlsRoute
}

func (l *TlsRouteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TlsRouteList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTlsRoute(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTlsRoute(ctx context.Context, project, location string) (*TlsRouteList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTlsRouteWithMaxResults(ctx, project, location, TlsRouteMaxPage)

}

func (c *Client) ListTlsRouteWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*TlsRouteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &TlsRoute{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listTlsRoute(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TlsRouteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTlsRoute(ctx context.Context, r *TlsRoute) (*TlsRoute, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractTlsRouteFields(r)

	b, err := c.getTlsRouteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTlsRoute(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTlsRouteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractTlsRouteFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTlsRoute(ctx context.Context, r *TlsRoute) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("TlsRoute resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting TlsRoute...")
	deleteOp := deleteTlsRouteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTlsRoute deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTlsRoute(ctx context.Context, project, location string, filter func(*TlsRoute) bool) error {
	listObj, err := c.ListTlsRoute(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllTlsRoute(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTlsRoute(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTlsRoute(ctx context.Context, rawDesired *TlsRoute, opts ...dcl.ApplyOption) (*TlsRoute, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *TlsRoute
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTlsRouteHelper(c, ctx, rawDesired, opts...)
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

func applyTlsRouteHelper(c *Client, ctx context.Context, rawDesired *TlsRoute, opts ...dcl.ApplyOption) (*TlsRoute, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTlsRoute...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTlsRouteFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.tlsRouteDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTlsRouteDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
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
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []tlsRouteApiOperation
	if create {
		ops = append(ops, &createTlsRouteOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyTlsRouteDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyTlsRouteDiff(c *Client, ctx context.Context, desired *TlsRoute, rawDesired *TlsRoute, ops []tlsRouteApiOperation, opts ...dcl.ApplyOption) (*TlsRoute, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTlsRoute(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTlsRouteOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTlsRoute(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTlsRouteNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTlsRouteNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTlsRouteDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractTlsRouteFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractTlsRouteFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTlsRoute(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
