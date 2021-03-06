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
package bigqueryreservation

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Assignment struct {
	Name        *string                `json:"name"`
	Assignee    *string                `json:"assignee"`
	JobType     *AssignmentJobTypeEnum `json:"jobType"`
	State       *AssignmentStateEnum   `json:"state"`
	Project     *string                `json:"project"`
	Location    *string                `json:"location"`
	Reservation *string                `json:"reservation"`
}

func (r *Assignment) String() string {
	return dcl.SprintResource(r)
}

// The enum AssignmentJobTypeEnum.
type AssignmentJobTypeEnum string

// AssignmentJobTypeEnumRef returns a *AssignmentJobTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AssignmentJobTypeEnumRef(s string) *AssignmentJobTypeEnum {
	if s == "" {
		return nil
	}

	v := AssignmentJobTypeEnum(s)
	return &v
}

func (v AssignmentJobTypeEnum) Validate() error {
	for _, s := range []string{"JOB_TYPE_UNSPECIFIED", "PIPELINE", "QUERY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AssignmentJobTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AssignmentStateEnum.
type AssignmentStateEnum string

// AssignmentStateEnumRef returns a *AssignmentStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func AssignmentStateEnumRef(s string) *AssignmentStateEnum {
	if s == "" {
		return nil
	}

	v := AssignmentStateEnum(s)
	return &v
}

func (v AssignmentStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PENDING", "ACTIVE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AssignmentStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Assignment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery_reservation",
		Type:    "Assignment",
		Version: "bigqueryreservation",
	}
}

const AssignmentMaxPage = -1

type AssignmentList struct {
	Items []*Assignment

	nextToken string

	pageSize int32

	project string

	location string

	reservation string
}

func (l *AssignmentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AssignmentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAssignment(ctx, l.project, l.location, l.reservation, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAssignment(ctx context.Context, project, location, reservation string) (*AssignmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAssignmentWithMaxResults(ctx, project, location, reservation, AssignmentMaxPage)

}

func (c *Client) ListAssignmentWithMaxResults(ctx context.Context, project, location, reservation string, pageSize int32) (*AssignmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAssignment(ctx, project, location, reservation, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AssignmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,

		reservation: reservation,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Assignment) URLNormalized() *Assignment {
	normalized := dcl.Copy(*r).(Assignment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Assignee = dcl.SelfLinkToName(r.Assignee)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Reservation = dcl.SelfLinkToName(r.Reservation)
	return &normalized
}

func (c *Client) GetAssignment(ctx context.Context, r *Assignment) (*Assignment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAssignmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAssignment(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Reservation = r.Reservation
	result.Assignee = r.Assignee
	result.JobType = r.JobType

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAssignmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAssignment(ctx context.Context, r *Assignment) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Assignment resource is nil")
	}
	c.Config.Logger.Info("Deleting Assignment...")
	deleteOp := deleteAssignmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAssignment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAssignment(ctx context.Context, project, location, reservation string, filter func(*Assignment) bool) error {
	listObj, err := c.ListAssignment(ctx, project, location, reservation)
	if err != nil {
		return err
	}

	err = c.deleteAllAssignment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAssignment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAssignment(ctx context.Context, rawDesired *Assignment, opts ...dcl.ApplyOption) (*Assignment, error) {
	var resultNewState *Assignment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAssignmentHelper(c, ctx, rawDesired, opts...)
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

func applyAssignmentHelper(c *Client, ctx context.Context, rawDesired *Assignment, opts ...dcl.ApplyOption) (*Assignment, error) {
	c.Config.Logger.Info("Beginning ApplyAssignment...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}
	vProject, err := dcl.ValueFromRegexOnField(rawDesired.Project, rawDesired.Reservation, "projects/([a-z0-9A-Z-]*)/locations/.*")
	if err != nil {
		return nil, err
	}
	rawDesired.Project = vProject
	vLocation, err := dcl.ValueFromRegexOnField(rawDesired.Location, rawDesired.Reservation, "projects/.*/locations/([a-z0-9A-Z-]*)/reservations/.*")
	if err != nil {
		return nil, err
	}
	rawDesired.Location = vLocation

	initial, desired, fieldDiffs, err := c.assignmentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAssignmentDiffs(c.Config, fieldDiffs, opts)
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
	var ops []assignmentApiOperation
	if create {
		ops = append(ops, &createAssignmentOperation{})
	} else if recreate {
		ops = append(ops, &deleteAssignmentOperation{})
		ops = append(ops, &createAssignmentOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAssignmentDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAssignment(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAssignmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAssignment(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAssignmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAssignmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAssignmentDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAssignment(c, newDesired, newState)
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
