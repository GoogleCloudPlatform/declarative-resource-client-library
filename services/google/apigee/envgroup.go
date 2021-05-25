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
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Envgroup struct {
	Name           *string            `json:"name"`
	Hostnames      []string           `json:"hostnames"`
	CreatedAt      *int64             `json:"createdAt"`
	LastModifiedAt *int64             `json:"lastModifiedAt"`
	State          *EnvgroupStateEnum `json:"state"`
	Organization   *string            `json:"organization"`
}

func (r *Envgroup) String() string {
	return dcl.SprintResource(r)
}

// The enum EnvgroupStateEnum.
type EnvgroupStateEnum string

// EnvgroupStateEnumRef returns a *EnvgroupStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func EnvgroupStateEnumRef(s string) *EnvgroupStateEnum {
	if s == "" {
		return nil
	}

	v := EnvgroupStateEnum(s)
	return &v
}

func (v EnvgroupStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "ACTIVE", "DELETING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EnvgroupStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Envgroup) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "apigee",
		Type:    "Envgroup",
		Version: "apigee",
	}
}

const EnvgroupMaxPage = -1

type EnvgroupList struct {
	Items []*Envgroup

	nextToken string

	pageSize int32

	organization string
}

func (l *EnvgroupList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EnvgroupList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEnvgroup(ctx, l.organization, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEnvgroup(ctx context.Context, organization string) (*EnvgroupList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEnvgroupWithMaxResults(ctx, organization, EnvgroupMaxPage)

}

func (c *Client) ListEnvgroupWithMaxResults(ctx context.Context, organization string, pageSize int32) (*EnvgroupList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listEnvgroup(ctx, organization, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EnvgroupList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		organization: organization,
	}, nil
}

func (c *Client) GetEnvgroup(ctx context.Context, r *Envgroup) (*Envgroup, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getEnvgroupRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEnvgroup(b, c)
	if err != nil {
		return nil, err
	}
	result.Organization = r.Organization
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEnvgroupNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEnvgroup(ctx context.Context, r *Envgroup) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Envgroup resource is nil")
	}
	c.Config.Logger.Info("Deleting Envgroup...")
	deleteOp := deleteEnvgroupOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEnvgroup deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEnvgroup(ctx context.Context, organization string, filter func(*Envgroup) bool) error {
	listObj, err := c.ListEnvgroup(ctx, organization)
	if err != nil {
		return err
	}

	err = c.deleteAllEnvgroup(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEnvgroup(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEnvgroup(ctx context.Context, rawDesired *Envgroup, opts ...dcl.ApplyOption) (*Envgroup, error) {

	var resultNewState *Envgroup
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEnvgroupHelper(c, ctx, rawDesired, opts...)
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

func applyEnvgroupHelper(c *Client, ctx context.Context, rawDesired *Envgroup, opts ...dcl.ApplyOption) (*Envgroup, error) {
	c.Config.Logger.Info("Beginning ApplyEnvgroup...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.envgroupDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []envgroupApiOperation
	if create {
		ops = append(ops, &createEnvgroupOperation{})
	} else if recreate {

		ops = append(ops, &deleteEnvgroupOperation{})

		ops = append(ops, &createEnvgroupOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeEnvgroupDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetEnvgroup(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEnvgroupOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEnvgroup(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEnvgroupNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEnvgroupNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEnvgroupDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEnvgroup(c, newDesired, newState)
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
