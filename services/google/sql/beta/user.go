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
package beta

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type User struct {
	Name                 *string                   `json:"name"`
	Password             *string                   `json:"password"`
	Project              *string                   `json:"project"`
	Instance             *string                   `json:"instance"`
	SqlserverUserDetails *UserSqlserverUserDetails `json:"sqlserverUserDetails"`
	Type                 *UserTypeEnum             `json:"type"`
	Etag                 *string                   `json:"etag"`
	Host                 *string                   `json:"host"`
}

func (r *User) String() string {
	return dcl.SprintResource(r)
}

// The enum UserTypeEnum.
type UserTypeEnum string

// UserTypeEnumRef returns a *UserTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func UserTypeEnumRef(s string) *UserTypeEnum {
	if s == "" {
		return nil
	}

	v := UserTypeEnum(s)
	return &v
}

func (v UserTypeEnum) Validate() error {
	for _, s := range []string{"NATIVE", "CLOUD_IAM_USER", "CLOUD_IAM_SERVICE_ACCOUNT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "UserTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type UserSqlserverUserDetails struct {
	empty       bool     `json:"-"`
	Disabled    *bool    `json:"disabled"`
	ServerRoles []string `json:"serverRoles"`
}

// This object is used to assert a desired state where this UserSqlserverUserDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyUserSqlserverUserDetails *UserSqlserverUserDetails = &UserSqlserverUserDetails{empty: true}

func (r *UserSqlserverUserDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *UserSqlserverUserDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *User) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "sql",
		Type:    "User",
		Version: "beta",
	}
}

const UserMaxPage = -1

type UserList struct {
	Items []*User

	nextToken string

	pageSize int32

	project string

	instance string
}

func (l *UserList) HasNext() bool {
	return l.nextToken != ""
}

func (l *UserList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listUser(ctx, l.project, l.instance, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListUser(ctx context.Context, project, instance string) (*UserList, error) {

	return c.ListUserWithMaxResults(ctx, project, instance, UserMaxPage)

}

func (c *Client) ListUserWithMaxResults(ctx context.Context, project, instance string, pageSize int32) (*UserList, error) {
	items, token, err := c.listUser(ctx, project, instance, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &UserList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		instance: instance,
	}, nil
}

func (c *Client) GetUser(ctx context.Context, r *User) (*User, error) {
	b, err := c.getUserRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalUser(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Instance = r.Instance
	result.Name = r.Name
	result.Host = r.Host

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeUserNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteUser(ctx context.Context, r *User) error {
	if r == nil {
		return fmt.Errorf("User resource is nil")
	}
	c.Config.Logger.Info("Deleting User...")
	deleteOp := deleteUserOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllUser deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllUser(ctx context.Context, project, instance string, filter func(*User) bool) error {
	listObj, err := c.ListUser(ctx, project, instance)
	if err != nil {
		return err
	}

	err = c.deleteAllUser(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllUser(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyUser(ctx context.Context, rawDesired *User, opts ...dcl.ApplyOption) (*User, error) {
	c.Config.Logger.Info("Beginning ApplyUser...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.userDiffsForRawDesired(ctx, rawDesired, opts...)
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
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []userApiOperation
	if create {
		ops = append(ops, &createUserOperation{})
	} else if recreate {

		ops = append(ops, &deleteUserOperation{})

		ops = append(ops, &createUserOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeUserDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetUser(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeUserNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeUserDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffUser(c, newDesired, newState)
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
