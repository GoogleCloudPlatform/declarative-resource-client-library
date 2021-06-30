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
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Route struct {
	Id               *int64         `json:"id"`
	Name             *string        `json:"name"`
	Description      *string        `json:"description"`
	Network          *string        `json:"network"`
	Tag              []string       `json:"tag"`
	DestRange        *string        `json:"destRange"`
	Priority         *int64         `json:"priority"`
	NextHopInstance  *string        `json:"nextHopInstance"`
	NextHopIP        *string        `json:"nextHopIP"`
	NextHopNetwork   *string        `json:"nextHopNetwork"`
	NextHopGateway   *string        `json:"nextHopGateway"`
	NextHopPeering   *string        `json:"nextHopPeering"`
	NextHopIlb       *string        `json:"nextHopIlb"`
	Warning          []RouteWarning `json:"warning"`
	NextHopVpnTunnel *string        `json:"nextHopVpnTunnel"`
	SelfLink         *string        `json:"selfLink"`
	Project          *string        `json:"project"`
}

func (r *Route) String() string {
	return dcl.SprintResource(r)
}

// The enum RouteWarningCodeEnum.
type RouteWarningCodeEnum string

// RouteWarningCodeEnumRef returns a *RouteWarningCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func RouteWarningCodeEnumRef(s string) *RouteWarningCodeEnum {
	if s == "" {
		return nil
	}

	v := RouteWarningCodeEnum(s)
	return &v
}

func (v RouteWarningCodeEnum) Validate() error {
	for _, s := range []string{"BAD_REQUEST", "FORBIDDEN", "NOT_FOUND", "CONFLICT", "GONE", "PRECONDITION_FAILED", "INTERNAL_ERROR", "SERVICE_UNAVAILABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RouteWarningCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type RouteWarning struct {
	empty   bool                  `json:"-"`
	Code    *RouteWarningCodeEnum `json:"code"`
	Message *string               `json:"message"`
	Data    map[string]string     `json:"data"`
}

type jsonRouteWarning RouteWarning

func (r *RouteWarning) UnmarshalJSON(data []byte) error {
	var res jsonRouteWarning
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRouteWarning
	} else {

		r.Code = res.Code

		r.Message = res.Message

		r.Data = res.Data

	}
	return nil
}

// This object is used to assert a desired state where this RouteWarning is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyRouteWarning *RouteWarning = &RouteWarning{empty: true}

func (r *RouteWarning) Empty() bool {
	return r.empty
}

func (r *RouteWarning) String() string {
	return dcl.SprintResource(r)
}

func (r *RouteWarning) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Route) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Route",
		Version: "compute",
	}
}

const RouteMaxPage = -1

type RouteList struct {
	Items []*Route

	nextToken string

	pageSize int32

	project string
}

func (l *RouteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *RouteList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listRoute(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListRoute(ctx context.Context, project string) (*RouteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListRouteWithMaxResults(ctx, project, RouteMaxPage)

}

func (c *Client) ListRouteWithMaxResults(ctx context.Context, project string, pageSize int32) (*RouteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listRoute(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &RouteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Route) URLNormalized() *Route {
	normalized := dcl.Copy(*r).(Route)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.DestRange = dcl.SelfLinkToName(r.DestRange)
	normalized.NextHopInstance = dcl.SelfLinkToName(r.NextHopInstance)
	normalized.NextHopIP = dcl.SelfLinkToName(r.NextHopIP)
	normalized.NextHopNetwork = dcl.SelfLinkToName(r.NextHopNetwork)
	normalized.NextHopGateway = dcl.SelfLinkToName(r.NextHopGateway)
	normalized.NextHopPeering = dcl.SelfLinkToName(r.NextHopPeering)
	normalized.NextHopIlb = dcl.SelfLinkToName(r.NextHopIlb)
	normalized.NextHopVpnTunnel = dcl.SelfLinkToName(r.NextHopVpnTunnel)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}
func (c *Client) GetRoute(ctx context.Context, r *Route) (*Route, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getRouteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalRoute(b, c)
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
	result, err = canonicalizeRouteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteRoute(ctx context.Context, r *Route) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Route resource is nil")
	}
	c.Config.Logger.Info("Deleting Route...")
	deleteOp := deleteRouteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllRoute deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllRoute(ctx context.Context, project string, filter func(*Route) bool) error {
	listObj, err := c.ListRoute(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllRoute(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllRoute(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyRoute(ctx context.Context, rawDesired *Route, opts ...dcl.ApplyOption) (*Route, error) {

	var resultNewState *Route
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyRouteHelper(c, ctx, rawDesired, opts...)
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

func applyRouteHelper(c *Client, ctx context.Context, rawDesired *Route, opts ...dcl.ApplyOption) (*Route, error) {
	c.Config.Logger.Info("Beginning ApplyRoute...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.routeDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToRouteOp(opStrings, fieldDiffs, opts)
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
	var ops []routeApiOperation
	if create {
		ops = append(ops, &createRouteOperation{})
	} else if recreate {
		ops = append(ops, &deleteRouteOperation{})
		ops = append(ops, &createRouteOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeRouteDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetRoute(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createRouteOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapRoute(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeRouteNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeRouteNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeRouteDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffRoute(c, newDesired, newState)
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
