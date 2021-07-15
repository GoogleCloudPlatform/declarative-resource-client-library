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

type Snapshot struct {
	Name                    *string                          `json:"name"`
	Description             *string                          `json:"description"`
	SourceDisk              *string                          `json:"sourceDisk"`
	DiskSizeGb              *int64                           `json:"diskSizeGb"`
	StorageBytes            *int64                           `json:"storageBytes"`
	License                 []string                         `json:"license"`
	SnapshotEncryptionKey   *SnapshotSnapshotEncryptionKey   `json:"snapshotEncryptionKey"`
	SourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey `json:"sourceDiskEncryptionKey"`
	SelfLink                *string                          `json:"selfLink"`
	Labels                  map[string]string                `json:"labels"`
	Project                 *string                          `json:"project"`
	Zone                    *string                          `json:"zone"`
	Id                      *int64                           `json:"id"`
}

func (r *Snapshot) String() string {
	return dcl.SprintResource(r)
}

type SnapshotSnapshotEncryptionKey struct {
	empty  bool    `json:"-"`
	RawKey *string `json:"rawKey"`
	Sha256 *string `json:"sha256"`
}

type jsonSnapshotSnapshotEncryptionKey SnapshotSnapshotEncryptionKey

func (r *SnapshotSnapshotEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonSnapshotSnapshotEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptySnapshotSnapshotEncryptionKey
	} else {

		r.RawKey = res.RawKey

		r.Sha256 = res.Sha256

	}
	return nil
}

// This object is used to assert a desired state where this SnapshotSnapshotEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySnapshotSnapshotEncryptionKey *SnapshotSnapshotEncryptionKey = &SnapshotSnapshotEncryptionKey{empty: true}

func (r *SnapshotSnapshotEncryptionKey) Empty() bool {
	return r.empty
}

func (r *SnapshotSnapshotEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *SnapshotSnapshotEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type SnapshotSourceDiskEncryptionKey struct {
	empty  bool    `json:"-"`
	RawKey *string `json:"rawKey"`
}

type jsonSnapshotSourceDiskEncryptionKey SnapshotSourceDiskEncryptionKey

func (r *SnapshotSourceDiskEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonSnapshotSourceDiskEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptySnapshotSourceDiskEncryptionKey
	} else {

		r.RawKey = res.RawKey

	}
	return nil
}

// This object is used to assert a desired state where this SnapshotSourceDiskEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptySnapshotSourceDiskEncryptionKey *SnapshotSourceDiskEncryptionKey = &SnapshotSourceDiskEncryptionKey{empty: true}

func (r *SnapshotSourceDiskEncryptionKey) Empty() bool {
	return r.empty
}

func (r *SnapshotSourceDiskEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *SnapshotSourceDiskEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Snapshot) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Snapshot",
		Version: "compute",
	}
}

const SnapshotMaxPage = -1

type SnapshotList struct {
	Items []*Snapshot

	nextToken string

	pageSize int32

	project string
}

func (l *SnapshotList) HasNext() bool {
	return l.nextToken != ""
}

func (l *SnapshotList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listSnapshot(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListSnapshot(ctx context.Context, project string) (*SnapshotList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListSnapshotWithMaxResults(ctx, project, SnapshotMaxPage)

}

func (c *Client) ListSnapshotWithMaxResults(ctx context.Context, project string, pageSize int32) (*SnapshotList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listSnapshot(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &SnapshotList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Snapshot) URLNormalized() *Snapshot {
	normalized := dcl.Copy(*r).(Snapshot)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceDisk = dcl.SelfLinkToName(r.SourceDisk)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	return &normalized
}

func (c *Client) GetSnapshot(ctx context.Context, r *Snapshot) (*Snapshot, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getSnapshotRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalSnapshot(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeSnapshotNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteSnapshot(ctx context.Context, r *Snapshot) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Snapshot resource is nil")
	}
	c.Config.Logger.Info("Deleting Snapshot...")
	deleteOp := deleteSnapshotOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllSnapshot deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllSnapshot(ctx context.Context, project string, filter func(*Snapshot) bool) error {
	listObj, err := c.ListSnapshot(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllSnapshot(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllSnapshot(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplySnapshot(ctx context.Context, rawDesired *Snapshot, opts ...dcl.ApplyOption) (*Snapshot, error) {
	var resultNewState *Snapshot
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applySnapshotHelper(c, ctx, rawDesired, opts...)
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

func applySnapshotHelper(c *Client, ctx context.Context, rawDesired *Snapshot, opts ...dcl.ApplyOption) (*Snapshot, error) {
	c.Config.Logger.Info("Beginning ApplySnapshot...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.snapshotDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToSnapshotDiffs(c.Config, fieldDiffs, opts)
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
	var ops []snapshotApiOperation
	if create {
		ops = append(ops, &createSnapshotOperation{})
	} else if recreate {
		ops = append(ops, &deleteSnapshotOperation{})
		ops = append(ops, &createSnapshotOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeSnapshotDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetSnapshot(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createSnapshotOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapSnapshot(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeSnapshotNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeSnapshotNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeSnapshotDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffSnapshot(c, newDesired, newState)
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
