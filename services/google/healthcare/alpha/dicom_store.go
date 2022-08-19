// Copyright 2022 Google LLC. All Rights Reserved.
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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type DicomStore struct {
	Name               *string                       `json:"name"`
	NotificationConfig *DicomStoreNotificationConfig `json:"notificationConfig"`
	Labels             map[string]string             `json:"labels"`
	Project            *string                       `json:"project"`
	Location           *string                       `json:"location"`
	Dataset            *string                       `json:"dataset"`
}

func (r *DicomStore) String() string {
	return dcl.SprintResource(r)
}

type DicomStoreNotificationConfig struct {
	empty       bool    `json:"-"`
	PubsubTopic *string `json:"pubsubTopic"`
}

type jsonDicomStoreNotificationConfig DicomStoreNotificationConfig

func (r *DicomStoreNotificationConfig) UnmarshalJSON(data []byte) error {
	var res jsonDicomStoreNotificationConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDicomStoreNotificationConfig
	} else {

		r.PubsubTopic = res.PubsubTopic

	}
	return nil
}

// This object is used to assert a desired state where this DicomStoreNotificationConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDicomStoreNotificationConfig *DicomStoreNotificationConfig = &DicomStoreNotificationConfig{empty: true}

func (r *DicomStoreNotificationConfig) Empty() bool {
	return r.empty
}

func (r *DicomStoreNotificationConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *DicomStoreNotificationConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *DicomStore) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "healthcare",
		Type:    "DicomStore",
		Version: "alpha",
	}
}

func (r *DicomStore) ID() (string, error) {
	if err := extractDicomStoreFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                dcl.ValueOrEmptyString(nr.Name),
		"notification_config": dcl.ValueOrEmptyString(nr.NotificationConfig),
		"labels":              dcl.ValueOrEmptyString(nr.Labels),
		"project":             dcl.ValueOrEmptyString(nr.Project),
		"location":            dcl.ValueOrEmptyString(nr.Location),
		"dataset":             dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/datasets/{{dataset}}/dicomStores/{{name}}", params), nil
}

const DicomStoreMaxPage = -1

type DicomStoreList struct {
	Items []*DicomStore

	nextToken string

	pageSize int32

	resource *DicomStore
}

func (l *DicomStoreList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DicomStoreList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDicomStore(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDicomStore(ctx context.Context, project, location, dataset string) (*DicomStoreList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDicomStoreWithMaxResults(ctx, project, location, dataset, DicomStoreMaxPage)

}

func (c *Client) ListDicomStoreWithMaxResults(ctx context.Context, project, location, dataset string, pageSize int32) (*DicomStoreList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &DicomStore{
		Project:  &project,
		Location: &location,
		Dataset:  &dataset,
	}
	items, token, err := c.listDicomStore(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DicomStoreList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetDicomStore(ctx context.Context, r *DicomStore) (*DicomStore, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractDicomStoreFields(r)

	b, err := c.getDicomStoreRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDicomStore(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Dataset = r.Dataset
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDicomStoreNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractDicomStoreFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDicomStore(ctx context.Context, r *DicomStore) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("DicomStore resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting DicomStore...")
	deleteOp := deleteDicomStoreOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDicomStore deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDicomStore(ctx context.Context, project, location, dataset string, filter func(*DicomStore) bool) error {
	listObj, err := c.ListDicomStore(ctx, project, location, dataset)
	if err != nil {
		return err
	}

	err = c.deleteAllDicomStore(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDicomStore(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDicomStore(ctx context.Context, rawDesired *DicomStore, opts ...dcl.ApplyOption) (*DicomStore, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *DicomStore
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDicomStoreHelper(c, ctx, rawDesired, opts...)
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

func applyDicomStoreHelper(c *Client, ctx context.Context, rawDesired *DicomStore, opts ...dcl.ApplyOption) (*DicomStore, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyDicomStore...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractDicomStoreFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.dicomStoreDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToDicomStoreDiffs(c.Config, fieldDiffs, opts)
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
	var ops []dicomStoreApiOperation
	if create {
		ops = append(ops, &createDicomStoreOperation{})
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
	return applyDicomStoreDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyDicomStoreDiff(c *Client, ctx context.Context, desired *DicomStore, rawDesired *DicomStore, ops []dicomStoreApiOperation, opts ...dcl.ApplyOption) (*DicomStore, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetDicomStore(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDicomStoreOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDicomStore(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDicomStoreNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDicomStoreNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDicomStoreDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractDicomStoreFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractDicomStoreFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDicomStore(c, newDesired, newState)
	if err != nil {
		return newState, err
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
