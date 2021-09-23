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
	v := BackupStateEnum(s)
	return &v
}

func (v BackupStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := BackupSourceInstanceTierEnum(s)
	return &v
}

func (v BackupSourceInstanceTierEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
		Service: "filestore",
		Type:    "Backup",
		Version: "beta",
	}
}

func (r *Backup) ID() (string, error) {
	if err := extractBackupFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":               dcl.ValueOrEmptyString(nr.Name),
		"description":        dcl.ValueOrEmptyString(nr.Description),
		"state":              dcl.ValueOrEmptyString(nr.State),
		"createTime":         dcl.ValueOrEmptyString(nr.CreateTime),
		"labels":             dcl.ValueOrEmptyString(nr.Labels),
		"capacityGb":         dcl.ValueOrEmptyString(nr.CapacityGb),
		"storageBytes":       dcl.ValueOrEmptyString(nr.StorageBytes),
		"sourceInstance":     dcl.ValueOrEmptyString(nr.SourceInstance),
		"sourceFileShare":    dcl.ValueOrEmptyString(nr.SourceFileShare),
		"sourceInstanceTier": dcl.ValueOrEmptyString(nr.SourceInstanceTier),
		"downloadBytes":      dcl.ValueOrEmptyString(nr.DownloadBytes),
		"project":            dcl.ValueOrEmptyString(nr.Project),
		"location":           dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/backups/{{name}}", params), nil
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

func (c *Client) ListBackup(ctx context.Context, project, location string) (*BackupList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListBackupWithMaxResults(ctx, project, location, BackupMaxPage)

}

func (c *Client) ListBackupWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*BackupList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Backup{
		Project:  &project,
		Location: &location,
	}
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
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractBackupFields(r)

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

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBackupNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBackup(ctx context.Context, r *Backup) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Backup resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Backup...")
	deleteOp := deleteBackupOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBackup deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBackup(ctx context.Context, project, location string, filter func(*Backup) bool) error {
	listObj, err := c.ListBackup(ctx, project, location)
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
	ctx = dcl.ContextWithRequestID(ctx)
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
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyBackup...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractBackupFields(rawDesired); err != nil {
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
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
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
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
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

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

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

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBackupNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBackupDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBackup(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}