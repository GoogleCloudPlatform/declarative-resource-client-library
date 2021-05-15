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
package apigee

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Environment struct {
	Name           *string                `json:"name"`
	Description    *string                `json:"description"`
	CreatedAt      *int64                 `json:"createdAt"`
	LastModifiedAt *int64                 `json:"lastModifiedAt"`
	Properties     *EnvironmentProperties `json:"properties"`
	DisplayName    *string                `json:"displayName"`
	State          *EnvironmentStateEnum  `json:"state"`
	Organization   *string                `json:"organization"`
}

func (r *Environment) String() string {
	return dcl.SprintResource(r)
}

// The enum EnvironmentStateEnum.
type EnvironmentStateEnum string

// EnvironmentStateEnumRef returns a *EnvironmentStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func EnvironmentStateEnumRef(s string) *EnvironmentStateEnum {
	if s == "" {
		return nil
	}

	v := EnvironmentStateEnum(s)
	return &v
}

func (v EnvironmentStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "ACTIVE", "DELETING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EnvironmentStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type EnvironmentProperties struct {
	empty    bool                            `json:"-"`
	Property []EnvironmentPropertiesProperty `json:"property"`
}

type jsonEnvironmentProperties EnvironmentProperties

func (r *EnvironmentProperties) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentProperties
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentProperties
	} else {

		r.Property = res.Property

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentProperties *EnvironmentProperties = &EnvironmentProperties{empty: true}

func (r *EnvironmentProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EnvironmentPropertiesProperty struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type jsonEnvironmentPropertiesProperty EnvironmentPropertiesProperty

func (r *EnvironmentPropertiesProperty) UnmarshalJSON(data []byte) error {
	var res jsonEnvironmentPropertiesProperty
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEnvironmentPropertiesProperty
	} else {

		r.Name = res.Name

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this EnvironmentPropertiesProperty is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEnvironmentPropertiesProperty *EnvironmentPropertiesProperty = &EnvironmentPropertiesProperty{empty: true}

func (r *EnvironmentPropertiesProperty) String() string {
	return dcl.SprintResource(r)
}

func (r *EnvironmentPropertiesProperty) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Environment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "apigee",
		Type:    "Environment",
		Version: "apigee",
	}
}

const EnvironmentMaxPage = -1

type EnvironmentList struct {
	Items []*Environment

	nextToken string

	pageSize int32

	organization string
}

func (l *EnvironmentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EnvironmentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEnvironment(ctx, l.organization, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEnvironment(ctx context.Context, organization string) (*EnvironmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListEnvironmentWithMaxResults(ctx, organization, EnvironmentMaxPage)

}

func (c *Client) ListEnvironmentWithMaxResults(ctx context.Context, organization string, pageSize int32) (*EnvironmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listEnvironment(ctx, organization, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EnvironmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		organization: organization,
	}, nil
}

func (c *Client) GetEnvironment(ctx context.Context, r *Environment) (*Environment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getEnvironmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEnvironment(b, c)
	if err != nil {
		return nil, err
	}
	result.Organization = r.Organization
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEnvironmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEnvironment(ctx context.Context, r *Environment) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Environment resource is nil")
	}
	c.Config.Logger.Info("Deleting Environment...")
	deleteOp := deleteEnvironmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEnvironment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEnvironment(ctx context.Context, organization string, filter func(*Environment) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListEnvironment(ctx, organization)
	if err != nil {
		return err
	}

	err = c.deleteAllEnvironment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEnvironment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEnvironment(ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (*Environment, error) {
	c.Config.Logger.Info("Beginning ApplyEnvironment...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.environmentDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []environmentApiOperation
	if create {
		ops = append(ops, &createEnvironmentOperation{})
	} else if recreate {

		ops = append(ops, &deleteEnvironmentOperation{})

		ops = append(ops, &createEnvironmentOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeEnvironmentDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetEnvironment(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEnvironmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEnvironment(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEnvironmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEnvironmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEnvironmentDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEnvironment(c, newDesired, newState)
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
func (r *Environment) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "POST", body, nil
}
