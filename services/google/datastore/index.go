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
package datastore

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Index struct {
	Ancestor   *IndexAncestorEnum `json:"ancestor"`
	IndexId    *string            `json:"indexId"`
	Kind       *string            `json:"kind"`
	Project    *string            `json:"project"`
	Properties []IndexProperties  `json:"properties"`
	State      *IndexStateEnum    `json:"state"`
}

func (r *Index) String() string {
	return dcl.SprintResource(r)
}

// The enum IndexAncestorEnum.
type IndexAncestorEnum string

// IndexAncestorEnumRef returns a *IndexAncestorEnum with the value of string s
// If the empty string is provided, nil is returned.
func IndexAncestorEnumRef(s string) *IndexAncestorEnum {
	if s == "" {
		return nil
	}

	v := IndexAncestorEnum(s)
	return &v
}

func (v IndexAncestorEnum) Validate() error {
	for _, s := range []string{"NONE", "ALL_ANCESTORS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "IndexAncestorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum IndexPropertiesDirectionEnum.
type IndexPropertiesDirectionEnum string

// IndexPropertiesDirectionEnumRef returns a *IndexPropertiesDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func IndexPropertiesDirectionEnumRef(s string) *IndexPropertiesDirectionEnum {
	if s == "" {
		return nil
	}

	v := IndexPropertiesDirectionEnum(s)
	return &v
}

func (v IndexPropertiesDirectionEnum) Validate() error {
	for _, s := range []string{"ASCENDING", "DESCENDING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "IndexPropertiesDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum IndexStateEnum.
type IndexStateEnum string

// IndexStateEnumRef returns a *IndexStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func IndexStateEnumRef(s string) *IndexStateEnum {
	if s == "" {
		return nil
	}

	v := IndexStateEnum(s)
	return &v
}

func (v IndexStateEnum) Validate() error {
	for _, s := range []string{"CREATING", "READY", "DELETING", "ERROR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "IndexStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type IndexProperties struct {
	empty     bool                          `json:"-"`
	Name      *string                       `json:"name"`
	Direction *IndexPropertiesDirectionEnum `json:"direction"`
}

type jsonIndexProperties IndexProperties

func (r *IndexProperties) UnmarshalJSON(data []byte) error {
	var res jsonIndexProperties
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyIndexProperties
	} else {

		r.Name = res.Name

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this IndexProperties is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyIndexProperties *IndexProperties = &IndexProperties{empty: true}

func (r *IndexProperties) Empty() bool {
	return r.empty
}

func (r *IndexProperties) String() string {
	return dcl.SprintResource(r)
}

func (r *IndexProperties) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Index) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "datastore",
		Type:    "Index",
		Version: "datastore",
	}
}

const IndexMaxPage = -1

type IndexList struct {
	Items []*Index

	nextToken string

	pageSize int32

	project string
}

func (l *IndexList) HasNext() bool {
	return l.nextToken != ""
}

func (l *IndexList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listIndex(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListIndex(ctx context.Context, project string) (*IndexList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListIndexWithMaxResults(ctx, project, IndexMaxPage)

}

func (c *Client) ListIndexWithMaxResults(ctx context.Context, project string, pageSize int32) (*IndexList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listIndex(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &IndexList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Index) URLNormalized() *Index {
	normalized := dcl.Copy(*r).(Index)
	normalized.IndexId = dcl.SelfLinkToName(r.IndexId)
	normalized.Kind = dcl.SelfLinkToName(r.Kind)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetIndex(ctx context.Context, r *Index) (*Index, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getIndexRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalIndex(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.IndexId = r.IndexId
	if dcl.IsZeroValue(result.Ancestor) {
		result.Ancestor = IndexAncestorEnumRef("NONE")
	}

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeIndexNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteIndex(ctx context.Context, r *Index) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Index resource is nil")
	}
	c.Config.Logger.Info("Deleting Index...")
	deleteOp := deleteIndexOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllIndex deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllIndex(ctx context.Context, project string, filter func(*Index) bool) error {
	listObj, err := c.ListIndex(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllIndex(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllIndex(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyIndex(ctx context.Context, rawDesired *Index, opts ...dcl.ApplyOption) (*Index, error) {
	var resultNewState *Index
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyIndexHelper(c, ctx, rawDesired, opts...)
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

func applyIndexHelper(c *Client, ctx context.Context, rawDesired *Index, opts ...dcl.ApplyOption) (*Index, error) {
	c.Config.Logger.Info("Beginning ApplyIndex...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.indexDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToIndexDiffs(c.Config, fieldDiffs, opts)
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
	var ops []indexApiOperation
	if create {
		ops = append(ops, &createIndexOperation{})
	} else if recreate {
		ops = append(ops, &deleteIndexOperation{})
		ops = append(ops, &createIndexOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeIndexDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetIndex(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createIndexOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapIndex(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeIndexNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeIndexNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeIndexDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffIndex(c, newDesired, newState)
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
