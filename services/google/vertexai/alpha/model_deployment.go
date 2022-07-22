// Copyright 2022 Google LLC. All Rights Reserved.
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

type ModelDeployment struct {
	Model              *string                            `json:"model"`
	DeployedModelId    *string                            `json:"deployedModelId"`
	DedicatedResources *ModelDeploymentDedicatedResources `json:"dedicatedResources"`
	Endpoint           *string                            `json:"endpoint"`
	Location           *string                            `json:"location"`
	Project            *string                            `json:"project"`
}

func (r *ModelDeployment) String() string {
	return dcl.SprintResource(r)
}

type ModelDeploymentDedicatedResources struct {
	empty           bool                                          `json:"-"`
	MachineSpec     *ModelDeploymentDedicatedResourcesMachineSpec `json:"machineSpec"`
	MinReplicaCount *int64                                        `json:"minReplicaCount"`
	MaxReplicaCount *int64                                        `json:"maxReplicaCount"`
}

type jsonModelDeploymentDedicatedResources ModelDeploymentDedicatedResources

func (r *ModelDeploymentDedicatedResources) UnmarshalJSON(data []byte) error {
	var res jsonModelDeploymentDedicatedResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelDeploymentDedicatedResources
	} else {

		r.MachineSpec = res.MachineSpec

		r.MinReplicaCount = res.MinReplicaCount

		r.MaxReplicaCount = res.MaxReplicaCount

	}
	return nil
}

// This object is used to assert a desired state where this ModelDeploymentDedicatedResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelDeploymentDedicatedResources *ModelDeploymentDedicatedResources = &ModelDeploymentDedicatedResources{empty: true}

func (r *ModelDeploymentDedicatedResources) Empty() bool {
	return r.empty
}

func (r *ModelDeploymentDedicatedResources) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelDeploymentDedicatedResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ModelDeploymentDedicatedResourcesMachineSpec struct {
	empty       bool    `json:"-"`
	MachineType *string `json:"machineType"`
}

type jsonModelDeploymentDedicatedResourcesMachineSpec ModelDeploymentDedicatedResourcesMachineSpec

func (r *ModelDeploymentDedicatedResourcesMachineSpec) UnmarshalJSON(data []byte) error {
	var res jsonModelDeploymentDedicatedResourcesMachineSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyModelDeploymentDedicatedResourcesMachineSpec
	} else {

		r.MachineType = res.MachineType

	}
	return nil
}

// This object is used to assert a desired state where this ModelDeploymentDedicatedResourcesMachineSpec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyModelDeploymentDedicatedResourcesMachineSpec *ModelDeploymentDedicatedResourcesMachineSpec = &ModelDeploymentDedicatedResourcesMachineSpec{empty: true}

func (r *ModelDeploymentDedicatedResourcesMachineSpec) Empty() bool {
	return r.empty
}

func (r *ModelDeploymentDedicatedResourcesMachineSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *ModelDeploymentDedicatedResourcesMachineSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ModelDeployment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vertex_ai",
		Type:    "ModelDeployment",
		Version: "alpha",
	}
}

func (r *ModelDeployment) ID() (string, error) {
	if err := extractModelDeploymentFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"model":               dcl.ValueOrEmptyString(nr.Model),
		"deployed_model_id":   dcl.ValueOrEmptyString(nr.DeployedModelId),
		"dedicated_resources": dcl.ValueOrEmptyString(nr.DedicatedResources),
		"endpoint":            dcl.ValueOrEmptyString(nr.Endpoint),
		"location":            dcl.ValueOrEmptyString(nr.Location),
		"project":             dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}/models/{{model}}", params), nil
}

const ModelDeploymentMaxPage = -1

type ModelDeploymentList struct {
	Items []*ModelDeployment

	nextToken string

	pageSize int32

	resource *ModelDeployment
}

func (l *ModelDeploymentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ModelDeploymentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listModelDeployment(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListModelDeployment(ctx context.Context, project, location, endpoint string) (*ModelDeploymentList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListModelDeploymentWithMaxResults(ctx, project, location, endpoint, ModelDeploymentMaxPage)

}

func (c *Client) ListModelDeploymentWithMaxResults(ctx context.Context, project, location, endpoint string, pageSize int32) (*ModelDeploymentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &ModelDeployment{
		Project:  &project,
		Location: &location,
		Endpoint: &endpoint,
	}
	items, token, err := c.listModelDeployment(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ModelDeploymentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetModelDeployment(ctx context.Context, r *ModelDeployment) (*ModelDeployment, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractModelDeploymentFields(r)

	b, err := c.getModelDeploymentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalModelDeployment(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Endpoint = r.Endpoint
	result.Model = r.Model

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeModelDeploymentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractModelDeploymentFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteModelDeployment(ctx context.Context, r *ModelDeployment) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ModelDeployment resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting ModelDeployment...")
	deleteOp := deleteModelDeploymentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllModelDeployment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllModelDeployment(ctx context.Context, project, location, endpoint string, filter func(*ModelDeployment) bool) error {
	listObj, err := c.ListModelDeployment(ctx, project, location, endpoint)
	if err != nil {
		return err
	}

	err = c.deleteAllModelDeployment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllModelDeployment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyModelDeployment(ctx context.Context, rawDesired *ModelDeployment, opts ...dcl.ApplyOption) (*ModelDeployment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *ModelDeployment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyModelDeploymentHelper(c, ctx, rawDesired, opts...)
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

func applyModelDeploymentHelper(c *Client, ctx context.Context, rawDesired *ModelDeployment, opts ...dcl.ApplyOption) (*ModelDeployment, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyModelDeployment...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractModelDeploymentFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.modelDeploymentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToModelDeploymentDiffs(c.Config, fieldDiffs, opts)
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
	var ops []modelDeploymentApiOperation
	if create {
		ops = append(ops, &createModelDeploymentOperation{})
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
	return applyModelDeploymentDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyModelDeploymentDiff(c *Client, ctx context.Context, desired *ModelDeployment, rawDesired *ModelDeployment, ops []modelDeploymentApiOperation, opts ...dcl.ApplyOption) (*ModelDeployment, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetModelDeployment(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createModelDeploymentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapModelDeployment(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeModelDeploymentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeModelDeploymentNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeModelDeploymentDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractModelDeploymentFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractModelDeploymentFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffModelDeployment(c, newDesired, newState)
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
