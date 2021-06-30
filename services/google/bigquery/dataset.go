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
package bigquery

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Dataset struct {
	Etag                           *string                                `json:"etag"`
	Id                             *string                                `json:"id"`
	SelfLink                       *string                                `json:"selfLink"`
	Name                           *string                                `json:"name"`
	Project                        *string                                `json:"project"`
	FriendlyName                   *string                                `json:"friendlyName"`
	Description                    *string                                `json:"description"`
	DefaultTableExpirationMs       *string                                `json:"defaultTableExpirationMs"`
	DefaultPartitionExpirationMs   *string                                `json:"defaultPartitionExpirationMs"`
	Labels                         map[string]string                      `json:"labels"`
	Access                         []DatasetAccess                        `json:"access"`
	CreationTime                   *int64                                 `json:"creationTime"`
	LastModifiedTime               *int64                                 `json:"lastModifiedTime"`
	Location                       *string                                `json:"location"`
	Published                      *bool                                  `json:"published"`
	DefaultEncryptionConfiguration *DatasetDefaultEncryptionConfiguration `json:"defaultEncryptionConfiguration"`
}

func (r *Dataset) String() string {
	return dcl.SprintResource(r)
}

type DatasetAccess struct {
	empty        bool                  `json:"-"`
	Role         *string               `json:"role"`
	UserByEmail  *string               `json:"userByEmail"`
	GroupByEmail *string               `json:"groupByEmail"`
	Domain       *string               `json:"domain"`
	SpecialGroup *string               `json:"specialGroup"`
	IamMember    *string               `json:"iamMember"`
	View         *DatasetAccessView    `json:"view"`
	Routine      *DatasetAccessRoutine `json:"routine"`
}

type jsonDatasetAccess DatasetAccess

func (r *DatasetAccess) UnmarshalJSON(data []byte) error {
	var res jsonDatasetAccess
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetAccess
	} else {

		r.Role = res.Role

		r.UserByEmail = res.UserByEmail

		r.GroupByEmail = res.GroupByEmail

		r.Domain = res.Domain

		r.SpecialGroup = res.SpecialGroup

		r.IamMember = res.IamMember

		r.View = res.View

		r.Routine = res.Routine

	}
	return nil
}

// This object is used to assert a desired state where this DatasetAccess is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDatasetAccess *DatasetAccess = &DatasetAccess{empty: true}

func (r *DatasetAccess) Empty() bool {
	return r.empty
}

func (r *DatasetAccess) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetAccess) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DatasetAccessView struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	TableId   *string `json:"tableId"`
}

type jsonDatasetAccessView DatasetAccessView

func (r *DatasetAccessView) UnmarshalJSON(data []byte) error {
	var res jsonDatasetAccessView
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetAccessView
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.TableId = res.TableId

	}
	return nil
}

// This object is used to assert a desired state where this DatasetAccessView is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDatasetAccessView *DatasetAccessView = &DatasetAccessView{empty: true}

func (r *DatasetAccessView) Empty() bool {
	return r.empty
}

func (r *DatasetAccessView) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetAccessView) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DatasetAccessRoutine struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	RoutineId *string `json:"routineId"`
}

type jsonDatasetAccessRoutine DatasetAccessRoutine

func (r *DatasetAccessRoutine) UnmarshalJSON(data []byte) error {
	var res jsonDatasetAccessRoutine
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetAccessRoutine
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.RoutineId = res.RoutineId

	}
	return nil
}

// This object is used to assert a desired state where this DatasetAccessRoutine is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDatasetAccessRoutine *DatasetAccessRoutine = &DatasetAccessRoutine{empty: true}

func (r *DatasetAccessRoutine) Empty() bool {
	return r.empty
}

func (r *DatasetAccessRoutine) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetAccessRoutine) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DatasetDefaultEncryptionConfiguration struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonDatasetDefaultEncryptionConfiguration DatasetDefaultEncryptionConfiguration

func (r *DatasetDefaultEncryptionConfiguration) UnmarshalJSON(data []byte) error {
	var res jsonDatasetDefaultEncryptionConfiguration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetDefaultEncryptionConfiguration
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this DatasetDefaultEncryptionConfiguration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDatasetDefaultEncryptionConfiguration *DatasetDefaultEncryptionConfiguration = &DatasetDefaultEncryptionConfiguration{empty: true}

func (r *DatasetDefaultEncryptionConfiguration) Empty() bool {
	return r.empty
}

func (r *DatasetDefaultEncryptionConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetDefaultEncryptionConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Dataset) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery",
		Type:    "Dataset",
		Version: "bigquery",
	}
}

const DatasetMaxPage = -1

type DatasetList struct {
	Items []*Dataset

	nextToken string

	pageSize int32

	project string
}

func (l *DatasetList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DatasetList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDataset(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDataset(ctx context.Context, project string) (*DatasetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDatasetWithMaxResults(ctx, project, DatasetMaxPage)

}

func (c *Client) ListDatasetWithMaxResults(ctx context.Context, project string, pageSize int32) (*DatasetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listDataset(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DatasetList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Dataset) URLNormalized() *Dataset {
	normalized := dcl.Copy(*r).(Dataset)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.FriendlyName = dcl.SelfLinkToName(r.FriendlyName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DefaultTableExpirationMs = dcl.SelfLinkToName(r.DefaultTableExpirationMs)
	normalized.DefaultPartitionExpirationMs = dcl.SelfLinkToName(r.DefaultPartitionExpirationMs)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}
func (c *Client) GetDataset(ctx context.Context, r *Dataset) (*Dataset, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getDatasetRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDataset(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDatasetNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDataset(ctx context.Context, r *Dataset) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Dataset resource is nil")
	}
	c.Config.Logger.Info("Deleting Dataset...")
	deleteOp := deleteDatasetOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDataset deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDataset(ctx context.Context, project string, filter func(*Dataset) bool) error {
	listObj, err := c.ListDataset(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllDataset(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDataset(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDataset(ctx context.Context, rawDesired *Dataset, opts ...dcl.ApplyOption) (*Dataset, error) {

	var resultNewState *Dataset
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDatasetHelper(c, ctx, rawDesired, opts...)
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

func applyDatasetHelper(c *Client, ctx context.Context, rawDesired *Dataset, opts ...dcl.ApplyOption) (*Dataset, error) {
	c.Config.Logger.Info("Beginning ApplyDataset...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.datasetDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToDatasetOp(opStrings, fieldDiffs, opts)
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
	var ops []datasetApiOperation
	if create {
		ops = append(ops, &createDatasetOperation{})
	} else if recreate {
		ops = append(ops, &deleteDatasetOperation{})
		ops = append(ops, &createDatasetOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeDatasetDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetDataset(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDatasetOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDataset(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDatasetNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDatasetNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDatasetDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDataset(c, newDesired, newState)
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
