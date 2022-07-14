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

type EndpointTrafficSplit struct {
	Endpoint     *string                            `json:"endpoint"`
	Project      *string                            `json:"project"`
	Location     *string                            `json:"location"`
	Etag         *string                            `json:"etag"`
	TrafficSplit []EndpointTrafficSplitTrafficSplit `json:"trafficSplit"`
}

func (r *EndpointTrafficSplit) String() string {
	return dcl.SprintResource(r)
}

type EndpointTrafficSplitTrafficSplit struct {
	empty             bool    `json:"-"`
	DeployedModelId   *string `json:"deployedModelId"`
	TrafficPercentage *int64  `json:"trafficPercentage"`
}

type jsonEndpointTrafficSplitTrafficSplit EndpointTrafficSplitTrafficSplit

func (r *EndpointTrafficSplitTrafficSplit) UnmarshalJSON(data []byte) error {
	var res jsonEndpointTrafficSplitTrafficSplit
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointTrafficSplitTrafficSplit
	} else {

		r.DeployedModelId = res.DeployedModelId

		r.TrafficPercentage = res.TrafficPercentage

	}
	return nil
}

// This object is used to assert a desired state where this EndpointTrafficSplitTrafficSplit is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointTrafficSplitTrafficSplit *EndpointTrafficSplitTrafficSplit = &EndpointTrafficSplitTrafficSplit{empty: true}

func (r *EndpointTrafficSplitTrafficSplit) Empty() bool {
	return r.empty
}

func (r *EndpointTrafficSplitTrafficSplit) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointTrafficSplitTrafficSplit) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *EndpointTrafficSplit) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vertex_ai",
		Type:    "EndpointTrafficSplit",
		Version: "beta",
	}
}

func (r *EndpointTrafficSplit) ID() (string, error) {
	if err := extractEndpointTrafficSplitFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"endpoint":      dcl.ValueOrEmptyString(nr.Endpoint),
		"project":       dcl.ValueOrEmptyString(nr.Project),
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"etag":          dcl.ValueOrEmptyString(nr.Etag),
		"traffic_split": dcl.ValueOrEmptyString(nr.TrafficSplit),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}", params), nil
}

const EndpointTrafficSplitMaxPage = -1

type EndpointTrafficSplitList struct {
	Items []*EndpointTrafficSplit

	nextToken string

	resource *EndpointTrafficSplit
}

func (c *Client) GetEndpointTrafficSplit(ctx context.Context, r *EndpointTrafficSplit) (*EndpointTrafficSplit, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractEndpointTrafficSplitFields(r)

	b, err := c.getEndpointTrafficSplitRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEndpointTrafficSplit(b, c, r)
	if err != nil {
		return nil, err
	}
	nr := r.urlNormalized()
	result.Project = nr.Project
	result.Location = nr.Location
	result.Endpoint = nr.Endpoint

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEndpointTrafficSplitNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractEndpointTrafficSplitFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyEndpointTrafficSplit(ctx context.Context, rawDesired *EndpointTrafficSplit, opts ...dcl.ApplyOption) (*EndpointTrafficSplit, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *EndpointTrafficSplit
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEndpointTrafficSplitHelper(c, ctx, rawDesired, opts...)
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

func applyEndpointTrafficSplitHelper(c *Client, ctx context.Context, rawDesired *EndpointTrafficSplit, opts ...dcl.ApplyOption) (*EndpointTrafficSplit, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyEndpointTrafficSplit...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEndpointTrafficSplitFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.endpointTrafficSplitDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEndpointTrafficSplitDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		return nil, dcl.ApplyInfeasibleError{Message: "No initial state found for singleton resource."}
	} else {
		for _, d := range diffs {
			if d.UpdateOp == nil {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) no update method found for field", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}
	var ops []endpointTrafficSplitApiOperation
	for _, d := range diffs {
		ops = append(ops, d.UpdateOp)
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
	return applyEndpointTrafficSplitDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyEndpointTrafficSplitDiff(c *Client, ctx context.Context, desired *EndpointTrafficSplit, rawDesired *EndpointTrafficSplit, ops []endpointTrafficSplitApiOperation, opts ...dcl.ApplyOption) (*EndpointTrafficSplit, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetEndpointTrafficSplit(ctx, desired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEndpointTrafficSplitNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEndpointTrafficSplitDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractEndpointTrafficSplitFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractEndpointTrafficSplitFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEndpointTrafficSplit(c, newDesired, newState)
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
