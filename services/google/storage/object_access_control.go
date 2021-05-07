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
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ObjectAccessControl struct {
	Project     *string                         `json:"project"`
	Bucket      *string                         `json:"bucket"`
	Domain      *string                         `json:"domain"`
	Email       *string                         `json:"email"`
	Entity      *string                         `json:"entity"`
	EntityId    *string                         `json:"entityId"`
	ProjectTeam *ObjectAccessControlProjectTeam `json:"projectTeam"`
	Role        *ObjectAccessControlRoleEnum    `json:"role"`
	Id          *string                         `json:"id"`
	Object      *string                         `json:"object"`
	Generation  *int64                          `json:"generation"`
}

func (r *ObjectAccessControl) String() string {
	return dcl.SprintResource(r)
}

// The enum ObjectAccessControlProjectTeamTeamEnum.
type ObjectAccessControlProjectTeamTeamEnum string

// ObjectAccessControlProjectTeamTeamEnumRef returns a *ObjectAccessControlProjectTeamTeamEnum with the value of string s
// If the empty string is provided, nil is returned.
func ObjectAccessControlProjectTeamTeamEnumRef(s string) *ObjectAccessControlProjectTeamTeamEnum {
	if s == "" {
		return nil
	}

	v := ObjectAccessControlProjectTeamTeamEnum(s)
	return &v
}

func (v ObjectAccessControlProjectTeamTeamEnum) Validate() error {
	for _, s := range []string{"editors", "owners", "viewers"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ObjectAccessControlProjectTeamTeamEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ObjectAccessControlRoleEnum.
type ObjectAccessControlRoleEnum string

// ObjectAccessControlRoleEnumRef returns a *ObjectAccessControlRoleEnum with the value of string s
// If the empty string is provided, nil is returned.
func ObjectAccessControlRoleEnumRef(s string) *ObjectAccessControlRoleEnum {
	if s == "" {
		return nil
	}

	v := ObjectAccessControlRoleEnum(s)
	return &v
}

func (v ObjectAccessControlRoleEnum) Validate() error {
	for _, s := range []string{"OWNER", "READER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ObjectAccessControlRoleEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ObjectAccessControlProjectTeam struct {
	empty         bool                                    `json:"-"`
	ProjectNumber *string                                 `json:"projectNumber"`
	Team          *ObjectAccessControlProjectTeamTeamEnum `json:"team"`
}

// This object is used to assert a desired state where this ObjectAccessControlProjectTeam is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyObjectAccessControlProjectTeam *ObjectAccessControlProjectTeam = &ObjectAccessControlProjectTeam{empty: true}

func (r *ObjectAccessControlProjectTeam) String() string {
	return dcl.SprintResource(r)
}

func (r *ObjectAccessControlProjectTeam) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ObjectAccessControl) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "storage",
		Type:    "ObjectAccessControl",
		Version: "storage",
	}
}

const ObjectAccessControlMaxPage = -1

type ObjectAccessControlList struct {
	Items []*ObjectAccessControl

	nextToken string

	pageSize int32

	project string

	bucket string

	object string
}

func (l *ObjectAccessControlList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ObjectAccessControlList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listObjectAccessControl(ctx, l.project, l.bucket, l.object, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListObjectAccessControl(ctx context.Context, project, bucket, object string) (*ObjectAccessControlList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListObjectAccessControlWithMaxResults(ctx, project, bucket, object, ObjectAccessControlMaxPage)

}

func (c *Client) ListObjectAccessControlWithMaxResults(ctx context.Context, project, bucket, object string, pageSize int32) (*ObjectAccessControlList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listObjectAccessControl(ctx, project, bucket, object, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ObjectAccessControlList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		bucket: bucket,

		object: object,
	}, nil
}

func (c *Client) GetObjectAccessControl(ctx context.Context, r *ObjectAccessControl) (*ObjectAccessControl, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getObjectAccessControlRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalObjectAccessControl(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Bucket = r.Bucket
	result.Entity = r.Entity
	result.Object = r.Object

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeObjectAccessControlNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteObjectAccessControl(ctx context.Context, r *ObjectAccessControl) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("ObjectAccessControl resource is nil")
	}
	c.Config.Logger.Info("Deleting ObjectAccessControl...")
	deleteOp := deleteObjectAccessControlOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllObjectAccessControl deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllObjectAccessControl(ctx context.Context, project, bucket, object string, filter func(*ObjectAccessControl) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListObjectAccessControl(ctx, project, bucket, object)
	if err != nil {
		return err
	}

	err = c.deleteAllObjectAccessControl(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllObjectAccessControl(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyObjectAccessControl(ctx context.Context, rawDesired *ObjectAccessControl, opts ...dcl.ApplyOption) (*ObjectAccessControl, error) {
	c.Config.Logger.Info("Beginning ApplyObjectAccessControl...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.objectAccessControlDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []objectAccessControlApiOperation
	if create {
		ops = append(ops, &createObjectAccessControlOperation{})
	} else if recreate {

		ops = append(ops, &deleteObjectAccessControlOperation{})

		ops = append(ops, &createObjectAccessControlOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeObjectAccessControlDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetObjectAccessControl(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createObjectAccessControlOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapObjectAccessControl(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeObjectAccessControlNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeObjectAccessControlNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeObjectAccessControlDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffObjectAccessControl(c, newDesired, newState)
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
