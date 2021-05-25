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
package iam

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

type ServiceAccount struct {
	Name           *string                       `json:"name"`
	Project        *string                       `json:"project"`
	UniqueId       *string                       `json:"uniqueId"`
	Email          *string                       `json:"email"`
	DisplayName    *string                       `json:"displayName"`
	Description    *string                       `json:"description"`
	OAuth2ClientId *string                       `json:"oauth2ClientId"`
	ActasResources *ServiceAccountActasResources `json:"actasResources"`
	Disabled       *bool                         `json:"disabled"`
}

func (r *ServiceAccount) String() string {
	return dcl.SprintResource(r)
}

type ServiceAccountActasResources struct {
	empty     bool                                    `json:"-"`
	Resources []ServiceAccountActasResourcesResources `json:"resources"`
}

type jsonServiceAccountActasResources ServiceAccountActasResources

func (r *ServiceAccountActasResources) UnmarshalJSON(data []byte) error {
	var res jsonServiceAccountActasResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceAccountActasResources
	} else {

		r.Resources = res.Resources

	}
	return nil
}

// This object is used to assert a desired state where this ServiceAccountActasResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceAccountActasResources *ServiceAccountActasResources = &ServiceAccountActasResources{empty: true}

func (r *ServiceAccountActasResources) Empty() bool {
	return r.empty
}

func (r *ServiceAccountActasResources) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceAccountActasResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceAccountActasResourcesResources struct {
	empty            bool    `json:"-"`
	FullResourceName *string `json:"fullResourceName"`
}

type jsonServiceAccountActasResourcesResources ServiceAccountActasResourcesResources

func (r *ServiceAccountActasResourcesResources) UnmarshalJSON(data []byte) error {
	var res jsonServiceAccountActasResourcesResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceAccountActasResourcesResources
	} else {

		r.FullResourceName = res.FullResourceName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceAccountActasResourcesResources is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceAccountActasResourcesResources *ServiceAccountActasResourcesResources = &ServiceAccountActasResourcesResources{empty: true}

func (r *ServiceAccountActasResourcesResources) Empty() bool {
	return r.empty
}

func (r *ServiceAccountActasResourcesResources) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceAccountActasResourcesResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ServiceAccount) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "iam",
		Type:    "ServiceAccount",
		Version: "iam",
	}
}

const ServiceAccountMaxPage = -1

type ServiceAccountList struct {
	Items []*ServiceAccount

	nextToken string

	pageSize int32

	project string
}

func (l *ServiceAccountList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServiceAccountList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listServiceAccount(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListServiceAccount(ctx context.Context, project string) (*ServiceAccountList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListServiceAccountWithMaxResults(ctx, project, ServiceAccountMaxPage)

}

func (c *Client) ListServiceAccountWithMaxResults(ctx context.Context, project string, pageSize int32) (*ServiceAccountList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listServiceAccount(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServiceAccountList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetServiceAccount(ctx context.Context, r *ServiceAccount) (*ServiceAccount, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getServiceAccountRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalServiceAccount(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServiceAccountNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteServiceAccount(ctx context.Context, r *ServiceAccount) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ServiceAccount resource is nil")
	}
	c.Config.Logger.Info("Deleting ServiceAccount...")
	deleteOp := deleteServiceAccountOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllServiceAccount deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllServiceAccount(ctx context.Context, project string, filter func(*ServiceAccount) bool) error {
	listObj, err := c.ListServiceAccount(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllServiceAccount(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllServiceAccount(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyServiceAccount(ctx context.Context, rawDesired *ServiceAccount, opts ...dcl.ApplyOption) (*ServiceAccount, error) {

	var resultNewState *ServiceAccount
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyServiceAccountHelper(c, ctx, rawDesired, opts...)
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

func applyServiceAccountHelper(c *Client, ctx context.Context, rawDesired *ServiceAccount, opts ...dcl.ApplyOption) (*ServiceAccount, error) {
	c.Config.Logger.Info("Beginning ApplyServiceAccount...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.serviceAccountDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []serviceAccountApiOperation
	if create {
		ops = append(ops, &createServiceAccountOperation{})
	} else if recreate {

		ops = append(ops, &deleteServiceAccountOperation{})

		ops = append(ops, &createServiceAccountOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeServiceAccountDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetServiceAccount(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServiceAccountOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapServiceAccount(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServiceAccountNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServiceAccountNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServiceAccountDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffServiceAccount(c, newDesired, newState)
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
func (r *ServiceAccount) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	body.WriteString(fmt.Sprintf(`{"options":{"requestedPolicyVersion": %d}}`, r.IAMPolicyVersion()))
	return u, "POST", body, nil
}
