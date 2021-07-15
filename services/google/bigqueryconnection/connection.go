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
package bigqueryconnection

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

type Connection struct {
	Name             *string             `json:"name"`
	FriendlyName     *string             `json:"friendlyName"`
	Description      *string             `json:"description"`
	CloudSql         *ConnectionCloudSql `json:"cloudSql"`
	CreationTime     *int64              `json:"creationTime"`
	LastModifiedTime *int64              `json:"lastModifiedTime"`
	HasCredential    *bool               `json:"hasCredential"`
	Project          *string             `json:"project"`
	Location         *string             `json:"location"`
}

func (r *Connection) String() string {
	return dcl.SprintResource(r)
}

// The enum ConnectionCloudSqlTypeEnum.
type ConnectionCloudSqlTypeEnum string

// ConnectionCloudSqlTypeEnumRef returns a *ConnectionCloudSqlTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConnectionCloudSqlTypeEnumRef(s string) *ConnectionCloudSqlTypeEnum {
	if s == "" {
		return nil
	}

	v := ConnectionCloudSqlTypeEnum(s)
	return &v
}

func (v ConnectionCloudSqlTypeEnum) Validate() error {
	for _, s := range []string{"DATABASE_TYPE_UNSPECIFIED", "POSTGRES", "MYSQL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConnectionCloudSqlTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ConnectionCloudSql struct {
	empty      bool                          `json:"-"`
	InstanceId *string                       `json:"instanceId"`
	Database   *string                       `json:"database"`
	Type       *ConnectionCloudSqlTypeEnum   `json:"type"`
	Credential *ConnectionCloudSqlCredential `json:"credential"`
}

type jsonConnectionCloudSql ConnectionCloudSql

func (r *ConnectionCloudSql) UnmarshalJSON(data []byte) error {
	var res jsonConnectionCloudSql
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConnectionCloudSql
	} else {

		r.InstanceId = res.InstanceId

		r.Database = res.Database

		r.Type = res.Type

		r.Credential = res.Credential

	}
	return nil
}

// This object is used to assert a desired state where this ConnectionCloudSql is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyConnectionCloudSql *ConnectionCloudSql = &ConnectionCloudSql{empty: true}

func (r *ConnectionCloudSql) Empty() bool {
	return r.empty
}

func (r *ConnectionCloudSql) String() string {
	return dcl.SprintResource(r)
}

func (r *ConnectionCloudSql) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConnectionCloudSqlCredential struct {
	empty    bool    `json:"-"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type jsonConnectionCloudSqlCredential ConnectionCloudSqlCredential

func (r *ConnectionCloudSqlCredential) UnmarshalJSON(data []byte) error {
	var res jsonConnectionCloudSqlCredential
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConnectionCloudSqlCredential
	} else {

		r.Username = res.Username

		r.Password = res.Password

	}
	return nil
}

// This object is used to assert a desired state where this ConnectionCloudSqlCredential is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyConnectionCloudSqlCredential *ConnectionCloudSqlCredential = &ConnectionCloudSqlCredential{empty: true}

func (r *ConnectionCloudSqlCredential) Empty() bool {
	return r.empty
}

func (r *ConnectionCloudSqlCredential) String() string {
	return dcl.SprintResource(r)
}

func (r *ConnectionCloudSqlCredential) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Connection) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery_connection",
		Type:    "Connection",
		Version: "bigqueryconnection",
	}
}

const ConnectionMaxPage = -1

type ConnectionList struct {
	Items []*Connection

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *ConnectionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ConnectionList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listConnection(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListConnection(ctx context.Context, project, location string) (*ConnectionList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListConnectionWithMaxResults(ctx, project, location, ConnectionMaxPage)

}

func (c *Client) ListConnectionWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ConnectionList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listConnection(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ConnectionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Connection) URLNormalized() *Connection {
	normalized := dcl.Copy(*r).(Connection)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.FriendlyName = dcl.SelfLinkToName(r.FriendlyName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetConnection(ctx context.Context, r *Connection) (*Connection, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getConnectionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalConnection(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeConnectionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteConnection(ctx context.Context, r *Connection) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Connection resource is nil")
	}
	c.Config.Logger.Info("Deleting Connection...")
	deleteOp := deleteConnectionOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllConnection deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllConnection(ctx context.Context, project, location string, filter func(*Connection) bool) error {
	listObj, err := c.ListConnection(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllConnection(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllConnection(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyConnection(ctx context.Context, rawDesired *Connection, opts ...dcl.ApplyOption) (*Connection, error) {

	var resultNewState *Connection
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyConnectionHelper(c, ctx, rawDesired, opts...)
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

func applyConnectionHelper(c *Client, ctx context.Context, rawDesired *Connection, opts ...dcl.ApplyOption) (*Connection, error) {
	c.Config.Logger.Info("Beginning ApplyConnection...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.connectionDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToConnectionDiffs(c.Config, fieldDiffs, opts)
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
	var ops []connectionApiOperation
	if create {
		ops = append(ops, &createConnectionOperation{})
	} else if recreate {
		ops = append(ops, &deleteConnectionOperation{})
		ops = append(ops, &createConnectionOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeConnectionDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetConnection(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createConnectionOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapConnection(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeConnectionNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeConnectionNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeConnectionDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffConnection(c, newDesired, newState)
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
func (r *Connection) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "POST", body, nil
}
