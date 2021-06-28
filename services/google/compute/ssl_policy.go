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
package compute

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type SslPolicy struct {
	Id             *int64                      `json:"id"`
	SelfLink       *string                     `json:"selfLink"`
	Name           *string                     `json:"name"`
	Description    *string                     `json:"description"`
	Profile        *SslPolicyProfileEnum       `json:"profile"`
	MinTlsVersion  *SslPolicyMinTlsVersionEnum `json:"minTlsVersion"`
	EnabledFeature []string                    `json:"enabledFeature"`
	CustomFeature  []string                    `json:"customFeature"`
	Warning        []SslPolicyWarning          `json:"warning"`
	Project        *string                     `json:"project"`
}

func (r *SslPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum SslPolicyProfileEnum.
type SslPolicyProfileEnum string

// SslPolicyProfileEnumRef returns a *SslPolicyProfileEnum with the value of string s
// If the empty string is provided, nil is returned.
func SslPolicyProfileEnumRef(s string) *SslPolicyProfileEnum {
	if s == "" {
		return nil
	}

	v := SslPolicyProfileEnum(s)
	return &v
}

func (v SslPolicyProfileEnum) Validate() error {
	for _, s := range []string{"COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "SslPolicyProfileEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum SslPolicyMinTlsVersionEnum.
type SslPolicyMinTlsVersionEnum string

// SslPolicyMinTlsVersionEnumRef returns a *SslPolicyMinTlsVersionEnum with the value of string s
// If the empty string is provided, nil is returned.
func SslPolicyMinTlsVersionEnumRef(s string) *SslPolicyMinTlsVersionEnum {
	if s == "" {
		return nil
	}

	v := SslPolicyMinTlsVersionEnum(s)
	return &v
}

func (v SslPolicyMinTlsVersionEnum) Validate() error {
	for _, s := range []string{"TLS_1_0", "TLS_1_1", "TLS_1_2"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "SslPolicyMinTlsVersionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type SslPolicyWarning struct {
	empty   bool                   `json:"-"`
	Code    *string                `json:"code"`
	Message *string                `json:"message"`
	Data    []SslPolicyWarningData `json:"data"`
}

type jsonSslPolicyWarning SslPolicyWarning

func (r *SslPolicyWarning) UnmarshalJSON(data []byte) error {
	var res jsonSslPolicyWarning
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptySslPolicyWarning
	} else {

		r.Code = res.Code

		r.Message = res.Message

		r.Data = res.Data

	}
	return nil
}

// This object is used to assert a desired state where this SslPolicyWarning is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySslPolicyWarning *SslPolicyWarning = &SslPolicyWarning{empty: true}

func (r *SslPolicyWarning) Empty() bool {
	return r.empty
}

func (r *SslPolicyWarning) String() string {
	return dcl.SprintResource(r)
}

func (r *SslPolicyWarning) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type SslPolicyWarningData struct {
	empty bool    `json:"-"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type jsonSslPolicyWarningData SslPolicyWarningData

func (r *SslPolicyWarningData) UnmarshalJSON(data []byte) error {
	var res jsonSslPolicyWarningData
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptySslPolicyWarningData
	} else {

		r.Key = res.Key

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this SslPolicyWarningData is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySslPolicyWarningData *SslPolicyWarningData = &SslPolicyWarningData{empty: true}

func (r *SslPolicyWarningData) Empty() bool {
	return r.empty
}

func (r *SslPolicyWarningData) String() string {
	return dcl.SprintResource(r)
}

func (r *SslPolicyWarningData) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *SslPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "SslPolicy",
		Version: "compute",
	}
}

const SslPolicyMaxPage = -1

type SslPolicyList struct {
	Items []*SslPolicy

	nextToken string

	pageSize int32

	project string
}

func (l *SslPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *SslPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listSslPolicy(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListSslPolicy(ctx context.Context, project string) (*SslPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListSslPolicyWithMaxResults(ctx, project, SslPolicyMaxPage)

}

func (c *Client) ListSslPolicyWithMaxResults(ctx context.Context, project string, pageSize int32) (*SslPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listSslPolicy(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &SslPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetSslPolicy(ctx context.Context, r *SslPolicy) (*SslPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getSslPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalSslPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name
	if dcl.IsZeroValue(result.Profile) {
		result.Profile = SslPolicyProfileEnumRef("COMPATIBLE")
	}
	if dcl.IsZeroValue(result.MinTlsVersion) {
		result.MinTlsVersion = SslPolicyMinTlsVersionEnumRef("TLS_1_0")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeSslPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteSslPolicy(ctx context.Context, r *SslPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("SslPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting SslPolicy...")
	deleteOp := deleteSslPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllSslPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllSslPolicy(ctx context.Context, project string, filter func(*SslPolicy) bool) error {
	listObj, err := c.ListSslPolicy(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllSslPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllSslPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplySslPolicy(ctx context.Context, rawDesired *SslPolicy, opts ...dcl.ApplyOption) (*SslPolicy, error) {

	var resultNewState *SslPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applySslPolicyHelper(c, ctx, rawDesired, opts...)
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

func applySslPolicyHelper(c *Client, ctx context.Context, rawDesired *SslPolicy, opts ...dcl.ApplyOption) (*SslPolicy, error) {
	c.Config.Logger.Info("Beginning ApplySslPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.sslPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	for _, fd := range fieldDiffs {
		fmt.Printf("fd: %+v\n", fd)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToSslPolicyOp(opStrings, fieldDiffs, opts)
	fmt.Printf("diffs: %+v, opStrings: %v\n", diffs, opStrings)
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
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []sslPolicyApiOperation
	if create {
		ops = append(ops, &createSslPolicyOperation{})
	} else if recreate {
		ops = append(ops, &deleteSslPolicyOperation{})
		ops = append(ops, &createSslPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeSslPolicyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetSslPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createSslPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapSslPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeSslPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeSslPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeSslPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffSslPolicy(c, newDesired, newState)
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
