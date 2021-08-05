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
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Backup struct {
	Name               *string                       `json:"name"`
	Description        *string                       `json:"description"`
	State              *BackupStateEnum              `json:"state"`
	CreateTime         *string                       `json:"createTime"`
	Labels             map[string]string             `json:"labels"`
	CapacityGb         *int64                        `json:"capacityGb"`
	StorageBytes       *int64                        `json:"storageBytes"`
	SourceInstance     *string                       `json:"sourceInstance"`
	SourceFileShare    *string                       `json:"sourceFileShare"`
	SourceInstanceTier *BackupSourceInstanceTierEnum `json:"sourceInstanceTier"`
	DownloadBytes      *int64                        `json:"downloadBytes"`
	Project            *string                       `json:"project"`
	Location           *string                       `json:"location"`
}

func (r *Backup) String() string {
	return dcl.SprintResource(r)
}

// The enum BackupStateEnum.
type BackupStateEnum string

// BackupStateEnumRef returns a *BackupStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackupStateEnumRef(s string) *BackupStateEnum {
	if s == "" {
		return nil
	}

	v := BackupStateEnum(s)
	return &v
}

func (v BackupStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "READY", "REPAIRING", "DELETING", "ERROR", "RESTORING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackupStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BackupSourceInstanceTierEnum.
type BackupSourceInstanceTierEnum string

// BackupSourceInstanceTierEnumRef returns a *BackupSourceInstanceTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func BackupSourceInstanceTierEnumRef(s string) *BackupSourceInstanceTierEnum {
	if s == "" {
		return nil
	}

	v := BackupSourceInstanceTierEnum(s)
	return &v
}

func (v BackupSourceInstanceTierEnum) Validate() error {
	for _, s := range []string{"TIER_UNSPECIFIED", "STANDARD", "PREMIUM", "BASIC_HDD", "BASIC_SSD", "HIGH_SCALE_SSD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BackupSourceInstanceTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Backup) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "file",
		Type:    "Backup",
		Version: "beta",
	}
}

const BackupMaxPage = -1

type BackupList struct {
	Items []*Backup

	nextToken string

	pageSize int32

	resource *Backup
}

func (l *BackupList) HasNext() bool {
	return l.nextToken != ""
}

func (l *BackupList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listBackup(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListBackup(ctx context.Context, r *Backup) (*BackupList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListBackupWithMaxResults(ctx, r, BackupMaxPage)

}

func (c *Client) ListBackupWithMaxResults(ctx context.Context, r *Backup, pageSize int32) (*BackupList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listBackup(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &BackupList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetBackup(ctx context.Context, r *Backup) (*Backup, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getBackupRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalBackup(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBackupNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBackup(ctx context.Context, r *Backup) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Backup resource is nil")
	}
	c.Config.Logger.Info("Deleting Backup...")
	deleteOp := deleteBackupOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBackup deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBackup(ctx context.Context, r *Backup, filter func(*Backup) bool) error {
	listObj, err := c.ListBackup(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllBackup(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllBackup(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyBackup(ctx context.Context, rawDesired *Backup, opts ...dcl.ApplyOption) (*Backup, error) {
	var resultNewState *Backup
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyBackupHelper(c, ctx, rawDesired, opts...)
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

func applyBackupHelper(c *Client, ctx context.Context, rawDesired *Backup, opts ...dcl.ApplyOption) (*Backup, error) {
	c.Config.Logger.Info("Beginning ApplyBackup...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.backupDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToBackupDiffs(c.Config, fieldDiffs, opts)
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
	var ops []backupApiOperation
	if create {
		ops = append(ops, &createBackupOperation{})
	} else if recreate {
		ops = append(ops, &deleteBackupOperation{})
		ops = append(ops, &createBackupOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeBackupDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetBackup(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createBackupOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapBackup(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeBackupNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBackupNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBackupDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBackup(c, newDesired, newState)
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
