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
package storage

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type DefaultObjectAccessControl struct {
	Project     *string                                `json:"project"`
	Bucket      *string                                `json:"bucket"`
	Domain      *string                                `json:"domain"`
	Email       *string                                `json:"email"`
	Entity      *string                                `json:"entity"`
	EntityId    *string                                `json:"entityId"`
	ProjectTeam *DefaultObjectAccessControlProjectTeam `json:"projectTeam"`
	Role        *DefaultObjectAccessControlRoleEnum    `json:"role"`
}

func (r *DefaultObjectAccessControl) String() string {
	return dcl.SprintResource(r)
}

// The enum DefaultObjectAccessControlProjectTeamTeamEnum.
type DefaultObjectAccessControlProjectTeamTeamEnum string

// DefaultObjectAccessControlProjectTeamTeamEnumRef returns a *DefaultObjectAccessControlProjectTeamTeamEnum with the value of string s
// If the empty string is provided, nil is returned.
func DefaultObjectAccessControlProjectTeamTeamEnumRef(s string) *DefaultObjectAccessControlProjectTeamTeamEnum {
	if s == "" {
		return nil
	}

	v := DefaultObjectAccessControlProjectTeamTeamEnum(s)
	return &v
}

func (v DefaultObjectAccessControlProjectTeamTeamEnum) Validate() error {
	for _, s := range []string{"editors", "owners", "viewers"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DefaultObjectAccessControlProjectTeamTeamEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DefaultObjectAccessControlRoleEnum.
type DefaultObjectAccessControlRoleEnum string

// DefaultObjectAccessControlRoleEnumRef returns a *DefaultObjectAccessControlRoleEnum with the value of string s
// If the empty string is provided, nil is returned.
func DefaultObjectAccessControlRoleEnumRef(s string) *DefaultObjectAccessControlRoleEnum {
	if s == "" {
		return nil
	}

	v := DefaultObjectAccessControlRoleEnum(s)
	return &v
}

func (v DefaultObjectAccessControlRoleEnum) Validate() error {
	for _, s := range []string{"OWNER", "READER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DefaultObjectAccessControlRoleEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type DefaultObjectAccessControlProjectTeam struct {
	empty         bool                                           `json:"-"`
	ProjectNumber *string                                        `json:"projectNumber"`
	Team          *DefaultObjectAccessControlProjectTeamTeamEnum `json:"team"`
}

type jsonDefaultObjectAccessControlProjectTeam DefaultObjectAccessControlProjectTeam

func (r *DefaultObjectAccessControlProjectTeam) UnmarshalJSON(data []byte) error {
	var res jsonDefaultObjectAccessControlProjectTeam
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDefaultObjectAccessControlProjectTeam
	} else {

		r.ProjectNumber = res.ProjectNumber

		r.Team = res.Team

	}
	return nil
}

// This object is used to assert a desired state where this DefaultObjectAccessControlProjectTeam is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDefaultObjectAccessControlProjectTeam *DefaultObjectAccessControlProjectTeam = &DefaultObjectAccessControlProjectTeam{empty: true}

func (r *DefaultObjectAccessControlProjectTeam) Empty() bool {
	return r.empty
}

func (r *DefaultObjectAccessControlProjectTeam) String() string {
	return dcl.SprintResource(r)
}

func (r *DefaultObjectAccessControlProjectTeam) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *DefaultObjectAccessControl) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "storage",
		Type:    "DefaultObjectAccessControl",
		Version: "storage",
	}
}

const DefaultObjectAccessControlMaxPage = -1

type DefaultObjectAccessControlList struct {
	Items []*DefaultObjectAccessControl

	nextToken string

	pageSize int32

	project string

	bucket string
}

func (l *DefaultObjectAccessControlList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DefaultObjectAccessControlList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDefaultObjectAccessControl(ctx, l.project, l.bucket, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDefaultObjectAccessControl(ctx context.Context, project, bucket string) (*DefaultObjectAccessControlList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDefaultObjectAccessControlWithMaxResults(ctx, project, bucket, DefaultObjectAccessControlMaxPage)

}

func (c *Client) ListDefaultObjectAccessControlWithMaxResults(ctx context.Context, project, bucket string, pageSize int32) (*DefaultObjectAccessControlList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listDefaultObjectAccessControl(ctx, project, bucket, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DefaultObjectAccessControlList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		bucket: bucket,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *DefaultObjectAccessControl) URLNormalized() *DefaultObjectAccessControl {
	normalized := dcl.Copy(*r).(DefaultObjectAccessControl)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Bucket = dcl.SelfLinkToName(r.Bucket)
	normalized.Domain = dcl.SelfLinkToName(r.Domain)
	normalized.Email = dcl.SelfLinkToName(r.Email)
	normalized.Entity = dcl.SelfLinkToName(r.Entity)
	normalized.EntityId = dcl.SelfLinkToName(r.EntityId)
	return &normalized
}

func (c *Client) GetDefaultObjectAccessControl(ctx context.Context, r *DefaultObjectAccessControl) (*DefaultObjectAccessControl, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getDefaultObjectAccessControlRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDefaultObjectAccessControl(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Bucket = r.Bucket
	result.Entity = r.Entity

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDefaultObjectAccessControlNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDefaultObjectAccessControl(ctx context.Context, r *DefaultObjectAccessControl) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("DefaultObjectAccessControl resource is nil")
	}
	c.Config.Logger.Info("Deleting DefaultObjectAccessControl...")
	deleteOp := deleteDefaultObjectAccessControlOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDefaultObjectAccessControl deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDefaultObjectAccessControl(ctx context.Context, project, bucket string, filter func(*DefaultObjectAccessControl) bool) error {
	listObj, err := c.ListDefaultObjectAccessControl(ctx, project, bucket)
	if err != nil {
		return err
	}

	err = c.deleteAllDefaultObjectAccessControl(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDefaultObjectAccessControl(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDefaultObjectAccessControl(ctx context.Context, rawDesired *DefaultObjectAccessControl, opts ...dcl.ApplyOption) (*DefaultObjectAccessControl, error) {

	var resultNewState *DefaultObjectAccessControl
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDefaultObjectAccessControlHelper(c, ctx, rawDesired, opts...)
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

func applyDefaultObjectAccessControlHelper(c *Client, ctx context.Context, rawDesired *DefaultObjectAccessControl, opts ...dcl.ApplyOption) (*DefaultObjectAccessControl, error) {
	c.Config.Logger.Info("Beginning ApplyDefaultObjectAccessControl...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.defaultObjectAccessControlDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToDefaultObjectAccessControlOp(opStrings, fieldDiffs, opts)
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
	var ops []defaultObjectAccessControlApiOperation
	if create {
		ops = append(ops, &createDefaultObjectAccessControlOperation{})
	} else if recreate {
		ops = append(ops, &deleteDefaultObjectAccessControlOperation{})
		ops = append(ops, &createDefaultObjectAccessControlOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeDefaultObjectAccessControlDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetDefaultObjectAccessControl(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDefaultObjectAccessControlOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDefaultObjectAccessControl(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDefaultObjectAccessControlNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDefaultObjectAccessControlNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDefaultObjectAccessControlDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDefaultObjectAccessControl(c, newDesired, newState)
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
