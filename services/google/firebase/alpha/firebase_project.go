// Copyright 2024 Google LLC. All Rights Reserved.
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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type FirebaseProject struct {
	ProjectId     *string                   `json:"projectId"`
	ProjectNumber *int64                    `json:"projectNumber"`
	DisplayName   *string                   `json:"displayName"`
	Resources     *FirebaseProjectResources `json:"resources"`
	State         *FirebaseProjectStateEnum `json:"state"`
	Annotations   map[string]string         `json:"annotations"`
	Project       *string                   `json:"project"`
}

func (r *FirebaseProject) String() string {
	return dcl.SprintResource(r)
}

// The enum FirebaseProjectStateEnum.
type FirebaseProjectStateEnum string

// FirebaseProjectStateEnumRef returns a *FirebaseProjectStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func FirebaseProjectStateEnumRef(s string) *FirebaseProjectStateEnum {
	v := FirebaseProjectStateEnum(s)
	return &v
}

func (v FirebaseProjectStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FirebaseProjectStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FirebaseProjectResources struct {
	empty                    bool    `json:"-"`
	HostingSite              *string `json:"hostingSite"`
	RealtimeDatabaseInstance *string `json:"realtimeDatabaseInstance"`
	StorageBucket            *string `json:"storageBucket"`
	LocationId               *string `json:"locationId"`
}

type jsonFirebaseProjectResources FirebaseProjectResources

func (r *FirebaseProjectResources) UnmarshalJSON(data []byte) error {
	var res jsonFirebaseProjectResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFirebaseProjectResources
	} else {

		r.HostingSite = res.HostingSite

		r.RealtimeDatabaseInstance = res.RealtimeDatabaseInstance

		r.StorageBucket = res.StorageBucket

		r.LocationId = res.LocationId

	}
	return nil
}

// This object is used to assert a desired state where this FirebaseProjectResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFirebaseProjectResources *FirebaseProjectResources = &FirebaseProjectResources{empty: true}

func (r *FirebaseProjectResources) Empty() bool {
	return r.empty
}

func (r *FirebaseProjectResources) String() string {
	return dcl.SprintResource(r)
}

func (r *FirebaseProjectResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *FirebaseProject) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "firebase",
		Type:    "FirebaseProject",
		Version: "alpha",
	}
}

func (r *FirebaseProject) ID() (string, error) {
	if err := extractFirebaseProjectFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project_id":     dcl.ValueOrEmptyString(nr.ProjectId),
		"project_number": dcl.ValueOrEmptyString(nr.ProjectNumber),
		"display_name":   dcl.ValueOrEmptyString(nr.DisplayName),
		"resources":      dcl.ValueOrEmptyString(nr.Resources),
		"state":          dcl.ValueOrEmptyString(nr.State),
		"annotations":    dcl.ValueOrEmptyString(nr.Annotations),
		"project":        dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}", params), nil
}

const FirebaseProjectMaxPage = -1

type FirebaseProjectList struct {
	Items []*FirebaseProject

	nextToken string

	pageSize int32

	resource *FirebaseProject
}

func (l *FirebaseProjectList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FirebaseProjectList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFirebaseProject(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFirebaseProject(ctx context.Context) (*FirebaseProjectList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListFirebaseProjectWithMaxResults(ctx, FirebaseProjectMaxPage)

}

func (c *Client) ListFirebaseProjectWithMaxResults(ctx context.Context, pageSize int32) (*FirebaseProjectList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &FirebaseProject{}
	items, token, err := c.listFirebaseProject(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FirebaseProjectList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetFirebaseProject(ctx context.Context, r *FirebaseProject) (*FirebaseProject, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractFirebaseProjectFields(r)

	b, err := c.getFirebaseProjectRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFirebaseProject(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFirebaseProjectNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractFirebaseProjectFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyFirebaseProject(ctx context.Context, rawDesired *FirebaseProject, opts ...dcl.ApplyOption) (*FirebaseProject, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *FirebaseProject
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFirebaseProjectHelper(c, ctx, rawDesired, opts...)
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

func applyFirebaseProjectHelper(c *Client, ctx context.Context, rawDesired *FirebaseProject, opts ...dcl.ApplyOption) (*FirebaseProject, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyFirebaseProject...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractFirebaseProjectFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.firebaseProjectDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFirebaseProjectDiffs(c.Config, fieldDiffs, opts)
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
	var ops []firebaseProjectApiOperation
	if create {
		ops = append(ops, &createFirebaseProjectOperation{})
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
	return applyFirebaseProjectDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyFirebaseProjectDiff(c *Client, ctx context.Context, desired *FirebaseProject, rawDesired *FirebaseProject, ops []firebaseProjectApiOperation, opts ...dcl.ApplyOption) (*FirebaseProject, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetFirebaseProject(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFirebaseProjectOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFirebaseProject(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFirebaseProjectNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFirebaseProjectNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFirebaseProjectDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractFirebaseProjectFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractFirebaseProjectFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFirebaseProject(c, newDesired, newState)
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
