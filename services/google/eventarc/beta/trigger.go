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

type Trigger struct {
	Name             *string                   `json:"name"`
	CreateTime       *string                   `json:"createTime"`
	UpdateTime       *string                   `json:"updateTime"`
	ServiceAccount   *string                   `json:"serviceAccount"`
	Destination      *TriggerDestination       `json:"destination"`
	Transport        *TriggerTransport         `json:"transport"`
	Labels           map[string]string         `json:"labels"`
	Etag             *string                   `json:"etag"`
	MatchingCriteria []TriggerMatchingCriteria `json:"matchingCriteria"`
	Project          *string                   `json:"project"`
	Location         *string                   `json:"location"`
}

func (r *Trigger) String() string {
	return dcl.SprintResource(r)
}

type TriggerDestination struct {
	empty           bool                               `json:"-"`
	CloudRunService *TriggerDestinationCloudRunService `json:"cloudRunService"`
}

// This object is used to assert a desired state where this TriggerDestination is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerDestination *TriggerDestination = &TriggerDestination{empty: true}

func (r *TriggerDestination) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerDestination) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerDestinationCloudRunService struct {
	empty   bool    `json:"-"`
	Service *string `json:"service"`
	Path    *string `json:"path"`
	Region  *string `json:"region"`
}

// This object is used to assert a desired state where this TriggerDestinationCloudRunService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerDestinationCloudRunService *TriggerDestinationCloudRunService = &TriggerDestinationCloudRunService{empty: true}

func (r *TriggerDestinationCloudRunService) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerDestinationCloudRunService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerTransport struct {
	empty  bool                    `json:"-"`
	Pubsub *TriggerTransportPubsub `json:"pubsub"`
}

// This object is used to assert a desired state where this TriggerTransport is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerTransport *TriggerTransport = &TriggerTransport{empty: true}

func (r *TriggerTransport) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerTransport) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerTransportPubsub struct {
	empty        bool    `json:"-"`
	Topic        *string `json:"topic"`
	Subscription *string `json:"subscription"`
}

// This object is used to assert a desired state where this TriggerTransportPubsub is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerTransportPubsub *TriggerTransportPubsub = &TriggerTransportPubsub{empty: true}

func (r *TriggerTransportPubsub) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerTransportPubsub) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerMatchingCriteria struct {
	empty     bool    `json:"-"`
	Attribute *string `json:"attribute"`
	Value     *string `json:"value"`
}

// This object is used to assert a desired state where this TriggerMatchingCriteria is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerMatchingCriteria *TriggerMatchingCriteria = &TriggerMatchingCriteria{empty: true}

func (r *TriggerMatchingCriteria) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerMatchingCriteria) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Trigger) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "eventarc",
		Type:    "Trigger",
		Version: "beta",
	}
}

const TriggerMaxPage = -1

type TriggerList struct {
	Items []*Trigger

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *TriggerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TriggerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTrigger(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTrigger(ctx context.Context, project, location string) (*TriggerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListTriggerWithMaxResults(ctx, project, location, TriggerMaxPage)

}

func (c *Client) ListTriggerWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*TriggerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listTrigger(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TriggerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetTrigger(ctx context.Context, r *Trigger) (*Trigger, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getTriggerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTrigger(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTriggerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTrigger(ctx context.Context, r *Trigger) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Trigger resource is nil")
	}
	c.Config.Logger.Info("Deleting Trigger...")
	deleteOp := deleteTriggerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTrigger deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTrigger(ctx context.Context, project, location string, filter func(*Trigger) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListTrigger(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllTrigger(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTrigger(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTrigger(ctx context.Context, rawDesired *Trigger, opts ...dcl.ApplyOption) (*Trigger, error) {
	c.Config.Logger.Info("Beginning ApplyTrigger...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.triggerDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []triggerApiOperation
	if create {
		ops = append(ops, &createTriggerOperation{})
	} else if recreate {

		ops = append(ops, &deleteTriggerOperation{})

		ops = append(ops, &createTriggerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTriggerDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetTrigger(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTriggerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTrigger(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTriggerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTriggerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTriggerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTrigger(c, newDesired, newState)
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
