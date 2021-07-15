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
package pubsublite

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Topic struct {
	Name            *string               `json:"name"`
	PartitionConfig *TopicPartitionConfig `json:"partitionConfig"`
	RetentionConfig *TopicRetentionConfig `json:"retentionConfig"`
	Project         *string               `json:"project"`
	Location        *string               `json:"location"`
}

func (r *Topic) String() string {
	return dcl.SprintResource(r)
}

type TopicPartitionConfig struct {
	empty    bool                          `json:"-"`
	Count    *int64                        `json:"count"`
	Capacity *TopicPartitionConfigCapacity `json:"capacity"`
}

type jsonTopicPartitionConfig TopicPartitionConfig

func (r *TopicPartitionConfig) UnmarshalJSON(data []byte) error {
	var res jsonTopicPartitionConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTopicPartitionConfig
	} else {

		r.Count = res.Count

		r.Capacity = res.Capacity

	}
	return nil
}

// This object is used to assert a desired state where this TopicPartitionConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTopicPartitionConfig *TopicPartitionConfig = &TopicPartitionConfig{empty: true}

func (r *TopicPartitionConfig) Empty() bool {
	return r.empty
}

func (r *TopicPartitionConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *TopicPartitionConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TopicPartitionConfigCapacity struct {
	empty              bool   `json:"-"`
	PublishMibPerSec   *int64 `json:"publishMibPerSec"`
	SubscribeMibPerSec *int64 `json:"subscribeMibPerSec"`
}

type jsonTopicPartitionConfigCapacity TopicPartitionConfigCapacity

func (r *TopicPartitionConfigCapacity) UnmarshalJSON(data []byte) error {
	var res jsonTopicPartitionConfigCapacity
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTopicPartitionConfigCapacity
	} else {

		r.PublishMibPerSec = res.PublishMibPerSec

		r.SubscribeMibPerSec = res.SubscribeMibPerSec

	}
	return nil
}

// This object is used to assert a desired state where this TopicPartitionConfigCapacity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTopicPartitionConfigCapacity *TopicPartitionConfigCapacity = &TopicPartitionConfigCapacity{empty: true}

func (r *TopicPartitionConfigCapacity) Empty() bool {
	return r.empty
}

func (r *TopicPartitionConfigCapacity) String() string {
	return dcl.SprintResource(r)
}

func (r *TopicPartitionConfigCapacity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TopicRetentionConfig struct {
	empty             bool    `json:"-"`
	PerPartitionBytes *int64  `json:"perPartitionBytes"`
	Period            *string `json:"period"`
}

type jsonTopicRetentionConfig TopicRetentionConfig

func (r *TopicRetentionConfig) UnmarshalJSON(data []byte) error {
	var res jsonTopicRetentionConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTopicRetentionConfig
	} else {

		r.PerPartitionBytes = res.PerPartitionBytes

		r.Period = res.Period

	}
	return nil
}

// This object is used to assert a desired state where this TopicRetentionConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTopicRetentionConfig *TopicRetentionConfig = &TopicRetentionConfig{empty: true}

func (r *TopicRetentionConfig) Empty() bool {
	return r.empty
}

func (r *TopicRetentionConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *TopicRetentionConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Topic) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "pubsub_lite",
		Type:    "Topic",
		Version: "pubsublite",
	}
}

const TopicMaxPage = -1

type TopicList struct {
	Items []*Topic

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *TopicList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TopicList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTopic(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTopic(ctx context.Context, project, location string) (*TopicList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTopicWithMaxResults(ctx, project, location, TopicMaxPage)

}

func (c *Client) ListTopicWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*TopicList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listTopic(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TopicList{
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
func (r *Topic) URLNormalized() *Topic {
	normalized := dcl.Copy(*r).(Topic)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (c *Client) GetTopic(ctx context.Context, r *Topic) (*Topic, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getTopicRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTopic(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTopicNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTopic(ctx context.Context, r *Topic) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Topic resource is nil")
	}
	c.Config.Logger.Info("Deleting Topic...")
	deleteOp := deleteTopicOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTopic deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTopic(ctx context.Context, project, location string, filter func(*Topic) bool) error {
	listObj, err := c.ListTopic(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllTopic(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTopic(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTopic(ctx context.Context, rawDesired *Topic, opts ...dcl.ApplyOption) (*Topic, error) {

	var resultNewState *Topic
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTopicHelper(c, ctx, rawDesired, opts...)
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

func applyTopicHelper(c *Client, ctx context.Context, rawDesired *Topic, opts ...dcl.ApplyOption) (*Topic, error) {
	c.Config.Logger.Info("Beginning ApplyTopic...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.topicDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTopicDiffs(c.Config, fieldDiffs, opts)
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
	var ops []topicApiOperation
	if create {
		ops = append(ops, &createTopicOperation{})
	} else if recreate {
		ops = append(ops, &deleteTopicOperation{})
		ops = append(ops, &createTopicOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTopicDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetTopic(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTopicOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTopic(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTopicNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTopicNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTopicDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTopic(c, newDesired, newState)
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
