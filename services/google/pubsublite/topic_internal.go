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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Topic) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.PartitionConfig) {
		if err := r.PartitionConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetentionConfig) {
		if err := r.RetentionConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TopicPartitionConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Capacity) {
		if err := r.Capacity.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TopicPartitionConfigCapacity) validate() error {
	return nil
}
func (r *TopicRetentionConfig) validate() error {
	return nil
}

func topicGetURL(userBasePath string, r *Topic) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/topics/{{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil
}

func topicListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/topics", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil

}

func topicCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/topics?topicId={{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil

}

func topicDeleteURL(userBasePath string, r *Topic) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/topics/{{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil
}

// topicApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type topicApiOperation interface {
	do(context.Context, *Topic, *Client) error
}

// newUpdateTopicUpdateTopicRequest creates a request for an
// Topic resource's UpdateTopic update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTopicUpdateTopicRequest(ctx context.Context, f *Topic, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandTopicPartitionConfig(c, f.PartitionConfig); err != nil {
		return nil, fmt.Errorf("error expanding PartitionConfig into partitionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["partitionConfig"] = v
	}
	if v, err := expandTopicRetentionConfig(c, f.RetentionConfig); err != nil {
		return nil, fmt.Errorf("error expanding RetentionConfig into retentionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["retentionConfig"] = v
	}
	return req, nil
}

// marshalUpdateTopicUpdateTopicRequest converts the update into
// the final JSON request body.
func marshalUpdateTopicUpdateTopicRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(map[string]interface{}{"topic": m})

}

type updateTopicUpdateTopicOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTopicUpdateTopicOperation) do(ctx context.Context, r *Topic, c *Client) error {
	_, err := c.GetTopic(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTopic")
	if err != nil {
		return err
	}

	req, err := newUpdateTopicUpdateTopicRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTopicUpdateTopicRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTopicRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := topicListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TopicMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listTopicOperation struct {
	Topics []map[string]interface{} `json:"topics"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listTopic(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Topic, string, error) {
	b, err := c.listTopicRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTopicOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Topic
	for _, v := range m.Topics {
		res := flattenTopic(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTopic(ctx context.Context, f func(*Topic) bool, resources []*Topic) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTopic(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteTopicOperation struct{}

func (op *deleteTopicOperation) do(ctx context.Context, r *Topic, c *Client) error {

	_, err := c.GetTopic(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Topic not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTopic checking for existence. error: %v", err)
		return err
	}

	u, err := topicDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Topic: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetTopic(ctx, r.urlNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createTopicOperation struct {
	response map[string]interface{}
}

func (op *createTopicOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTopicOperation) do(ctx context.Context, r *Topic, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := topicCreateURL(c.Config.BasePath, project, location, name)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o
	// Poll for the Topic resource to be created. Topic resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := topicGetURL(c.Config.BasePath, r)
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true) {
				return nil, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		getResp.Response.Body.Close()
		return getResp, nil
	}, c.Config.RetryProvider)

	if _, err := c.GetTopic(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTopicRaw(ctx context.Context, r *Topic) ([]byte, error) {

	u, err := topicGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) topicDiffsForRawDesired(ctx context.Context, rawDesired *Topic, opts ...dcl.ApplyOption) (initial, desired *Topic, diffs []topicDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Topic
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Topic); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Topic, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTopic(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Topic resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Topic resource: %v", err)
		}
		c.Config.Logger.Info("Found that Topic resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTopicDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Topic: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Topic: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTopicInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Topic: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTopicDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Topic: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTopic(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTopicInitialState(rawInitial, rawDesired *Topic) (*Topic, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTopicDesiredState(rawDesired, rawInitial *Topic, opts ...dcl.ApplyOption) (*Topic, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.PartitionConfig = canonicalizeTopicPartitionConfig(rawDesired.PartitionConfig, nil, opts...)
		rawDesired.RetentionConfig = canonicalizeTopicRetentionConfig(rawDesired.RetentionConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.PartitionConfig = canonicalizeTopicPartitionConfig(rawDesired.PartitionConfig, rawInitial.PartitionConfig, opts...)
	rawDesired.RetentionConfig = canonicalizeTopicRetentionConfig(rawDesired.RetentionConfig, rawInitial.RetentionConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeTopicNewState(c *Client, rawNew, rawDesired *Topic) (*Topic, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PartitionConfig) && dcl.IsEmptyValueIndirect(rawDesired.PartitionConfig) {
		rawNew.PartitionConfig = rawDesired.PartitionConfig
	} else {
		rawNew.PartitionConfig = canonicalizeNewTopicPartitionConfig(c, rawDesired.PartitionConfig, rawNew.PartitionConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.RetentionConfig) && dcl.IsEmptyValueIndirect(rawDesired.RetentionConfig) {
		rawNew.RetentionConfig = rawDesired.RetentionConfig
	} else {
		rawNew.RetentionConfig = canonicalizeNewTopicRetentionConfig(c, rawDesired.RetentionConfig, rawNew.RetentionConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeTopicPartitionConfig(des, initial *TopicPartitionConfig, opts ...dcl.ApplyOption) *TopicPartitionConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Count) {
		des.Count = initial.Count
	}
	des.Capacity = canonicalizeTopicPartitionConfigCapacity(des.Capacity, initial.Capacity, opts...)

	return des
}

func canonicalizeNewTopicPartitionConfig(c *Client, des, nw *TopicPartitionConfig) *TopicPartitionConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.Capacity = canonicalizeNewTopicPartitionConfigCapacity(c, des.Capacity, nw.Capacity)

	return nw
}

func canonicalizeNewTopicPartitionConfigSet(c *Client, des, nw []TopicPartitionConfig) []TopicPartitionConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []TopicPartitionConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareTopicPartitionConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTopicPartitionConfigSlice(c *Client, des, nw []TopicPartitionConfig) []TopicPartitionConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []TopicPartitionConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTopicPartitionConfig(c, &d, &n))
	}

	return items
}

func canonicalizeTopicPartitionConfigCapacity(des, initial *TopicPartitionConfigCapacity, opts ...dcl.ApplyOption) *TopicPartitionConfigCapacity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PublishMibPerSec) {
		des.PublishMibPerSec = initial.PublishMibPerSec
	}
	if dcl.IsZeroValue(des.SubscribeMibPerSec) {
		des.SubscribeMibPerSec = initial.SubscribeMibPerSec
	}

	return des
}

func canonicalizeNewTopicPartitionConfigCapacity(c *Client, des, nw *TopicPartitionConfigCapacity) *TopicPartitionConfigCapacity {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewTopicPartitionConfigCapacitySet(c *Client, des, nw []TopicPartitionConfigCapacity) []TopicPartitionConfigCapacity {
	if des == nil {
		return nw
	}
	var reorderedNew []TopicPartitionConfigCapacity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareTopicPartitionConfigCapacity(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTopicPartitionConfigCapacitySlice(c *Client, des, nw []TopicPartitionConfigCapacity) []TopicPartitionConfigCapacity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []TopicPartitionConfigCapacity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTopicPartitionConfigCapacity(c, &d, &n))
	}

	return items
}

func canonicalizeTopicRetentionConfig(des, initial *TopicRetentionConfig, opts ...dcl.ApplyOption) *TopicRetentionConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PerPartitionBytes) {
		des.PerPartitionBytes = initial.PerPartitionBytes
	}
	if dcl.StringCanonicalize(des.Period, initial.Period) || dcl.IsZeroValue(des.Period) {
		des.Period = initial.Period
	}

	return des
}

func canonicalizeNewTopicRetentionConfig(c *Client, des, nw *TopicRetentionConfig) *TopicRetentionConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Period, nw.Period) || dcl.IsZeroValue(des.Period) {
		nw.Period = des.Period
	}

	return nw
}

func canonicalizeNewTopicRetentionConfigSet(c *Client, des, nw []TopicRetentionConfig) []TopicRetentionConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []TopicRetentionConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareTopicRetentionConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTopicRetentionConfigSlice(c *Client, des, nw []TopicRetentionConfig) []TopicRetentionConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []TopicRetentionConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTopicRetentionConfig(c, &d, &n))
	}

	return items
}

type topicDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         topicApiOperation
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffTopic(c *Client, desired, actual *Topic, opts ...dcl.ApplyOption) ([]topicDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []topicDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, topicDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareTopicPartitionConfig(c, desired.PartitionConfig, actual.PartitionConfig) {
		c.Config.Logger.Infof("Detected diff in PartitionConfig.\nDESIRED: %v\nACTUAL: %v", desired.PartitionConfig, actual.PartitionConfig)

		diffs = append(diffs, topicDiff{
			UpdateOp:  &updateTopicUpdateTopicOperation{},
			FieldName: "PartitionConfig",
		})

	}
	if compareTopicRetentionConfig(c, desired.RetentionConfig, actual.RetentionConfig) {
		c.Config.Logger.Infof("Detected diff in RetentionConfig.\nDESIRED: %v\nACTUAL: %v", desired.RetentionConfig, actual.RetentionConfig)

		diffs = append(diffs, topicDiff{
			UpdateOp:  &updateTopicUpdateTopicOperation{},
			FieldName: "RetentionConfig",
		})

	}
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []topicDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareTopicPartitionConfig(c *Client, desired, actual *TopicPartitionConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Count == nil && desired.Count != nil && !dcl.IsEmptyValueIndirect(desired.Count) {
		c.Config.Logger.Infof("desired Count %s - but actually nil", dcl.SprintResource(desired.Count))
		return true
	}
	if !reflect.DeepEqual(desired.Count, actual.Count) && !dcl.IsZeroValue(desired.Count) {
		c.Config.Logger.Infof("Diff in Count. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Count), dcl.SprintResource(actual.Count))
		return true
	}
	if actual.Capacity == nil && desired.Capacity != nil && !dcl.IsEmptyValueIndirect(desired.Capacity) {
		c.Config.Logger.Infof("desired Capacity %s - but actually nil", dcl.SprintResource(desired.Capacity))
		return true
	}
	if compareTopicPartitionConfigCapacity(c, desired.Capacity, actual.Capacity) && !dcl.IsZeroValue(desired.Capacity) {
		c.Config.Logger.Infof("Diff in Capacity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Capacity), dcl.SprintResource(actual.Capacity))
		return true
	}
	return false
}

func compareTopicPartitionConfigSlice(c *Client, desired, actual []TopicPartitionConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TopicPartitionConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTopicPartitionConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TopicPartitionConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTopicPartitionConfigMap(c *Client, desired, actual map[string]TopicPartitionConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TopicPartitionConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TopicPartitionConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTopicPartitionConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TopicPartitionConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTopicPartitionConfigCapacity(c *Client, desired, actual *TopicPartitionConfigCapacity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.PublishMibPerSec == nil && desired.PublishMibPerSec != nil && !dcl.IsEmptyValueIndirect(desired.PublishMibPerSec) {
		c.Config.Logger.Infof("desired PublishMibPerSec %s - but actually nil", dcl.SprintResource(desired.PublishMibPerSec))
		return true
	}
	if !reflect.DeepEqual(desired.PublishMibPerSec, actual.PublishMibPerSec) && !dcl.IsZeroValue(desired.PublishMibPerSec) {
		c.Config.Logger.Infof("Diff in PublishMibPerSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PublishMibPerSec), dcl.SprintResource(actual.PublishMibPerSec))
		return true
	}
	if actual.SubscribeMibPerSec == nil && desired.SubscribeMibPerSec != nil && !dcl.IsEmptyValueIndirect(desired.SubscribeMibPerSec) {
		c.Config.Logger.Infof("desired SubscribeMibPerSec %s - but actually nil", dcl.SprintResource(desired.SubscribeMibPerSec))
		return true
	}
	if !reflect.DeepEqual(desired.SubscribeMibPerSec, actual.SubscribeMibPerSec) && !dcl.IsZeroValue(desired.SubscribeMibPerSec) {
		c.Config.Logger.Infof("Diff in SubscribeMibPerSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SubscribeMibPerSec), dcl.SprintResource(actual.SubscribeMibPerSec))
		return true
	}
	return false
}

func compareTopicPartitionConfigCapacitySlice(c *Client, desired, actual []TopicPartitionConfigCapacity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TopicPartitionConfigCapacity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTopicPartitionConfigCapacity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TopicPartitionConfigCapacity, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTopicPartitionConfigCapacityMap(c *Client, desired, actual map[string]TopicPartitionConfigCapacity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TopicPartitionConfigCapacity, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TopicPartitionConfigCapacity, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTopicPartitionConfigCapacity(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TopicPartitionConfigCapacity, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTopicRetentionConfig(c *Client, desired, actual *TopicRetentionConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.PerPartitionBytes == nil && desired.PerPartitionBytes != nil && !dcl.IsEmptyValueIndirect(desired.PerPartitionBytes) {
		c.Config.Logger.Infof("desired PerPartitionBytes %s - but actually nil", dcl.SprintResource(desired.PerPartitionBytes))
		return true
	}
	if !reflect.DeepEqual(desired.PerPartitionBytes, actual.PerPartitionBytes) && !dcl.IsZeroValue(desired.PerPartitionBytes) {
		c.Config.Logger.Infof("Diff in PerPartitionBytes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PerPartitionBytes), dcl.SprintResource(actual.PerPartitionBytes))
		return true
	}
	if actual.Period == nil && desired.Period != nil && !dcl.IsEmptyValueIndirect(desired.Period) {
		c.Config.Logger.Infof("desired Period %s - but actually nil", dcl.SprintResource(desired.Period))
		return true
	}
	if !dcl.StringCanonicalize(desired.Period, actual.Period) && !dcl.IsZeroValue(desired.Period) {
		c.Config.Logger.Infof("Diff in Period. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Period), dcl.SprintResource(actual.Period))
		return true
	}
	return false
}

func compareTopicRetentionConfigSlice(c *Client, desired, actual []TopicRetentionConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TopicRetentionConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTopicRetentionConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TopicRetentionConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTopicRetentionConfigMap(c *Client, desired, actual map[string]TopicRetentionConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TopicRetentionConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TopicRetentionConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTopicRetentionConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TopicRetentionConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Topic) urlNormalized() *Topic {
	normalized := deepcopy.Copy(*r).(Topic)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Topic) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Topic) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Topic) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Topic) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateTopic" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("admin/projects/{{project}}/locations/{{location}}/topics/{{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Topic resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Topic) marshal(c *Client) ([]byte, error) {
	m, err := expandTopic(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Topic: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTopic decodes JSON responses into the Topic resource schema.
func unmarshalTopic(b []byte, c *Client) (*Topic, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTopic(m, c)
}

func unmarshalMapTopic(m map[string]interface{}, c *Client) (*Topic, error) {

	return flattenTopic(c, m), nil
}

// expandTopic expands Topic into a JSON request object.
func expandTopic(c *Client, f *Topic) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/topics/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandTopicPartitionConfig(c, f.PartitionConfig); err != nil {
		return nil, fmt.Errorf("error expanding PartitionConfig into partitionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["partitionConfig"] = v
	}
	if v, err := expandTopicRetentionConfig(c, f.RetentionConfig); err != nil {
		return nil, fmt.Errorf("error expanding RetentionConfig into retentionConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retentionConfig"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenTopic flattens Topic from a JSON request object into the
// Topic type.
func flattenTopic(c *Client, i interface{}) *Topic {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Topic{}
	r.Name = dcl.FlattenString(m["name"])
	r.PartitionConfig = flattenTopicPartitionConfig(c, m["partitionConfig"])
	r.RetentionConfig = flattenTopicRetentionConfig(c, m["retentionConfig"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandTopicPartitionConfigMap expands the contents of TopicPartitionConfig into a JSON
// request object.
func expandTopicPartitionConfigMap(c *Client, f map[string]TopicPartitionConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTopicPartitionConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTopicPartitionConfigSlice expands the contents of TopicPartitionConfig into a JSON
// request object.
func expandTopicPartitionConfigSlice(c *Client, f []TopicPartitionConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTopicPartitionConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTopicPartitionConfigMap flattens the contents of TopicPartitionConfig from a JSON
// response object.
func flattenTopicPartitionConfigMap(c *Client, i interface{}) map[string]TopicPartitionConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TopicPartitionConfig{}
	}

	if len(a) == 0 {
		return map[string]TopicPartitionConfig{}
	}

	items := make(map[string]TopicPartitionConfig)
	for k, item := range a {
		items[k] = *flattenTopicPartitionConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTopicPartitionConfigSlice flattens the contents of TopicPartitionConfig from a JSON
// response object.
func flattenTopicPartitionConfigSlice(c *Client, i interface{}) []TopicPartitionConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []TopicPartitionConfig{}
	}

	if len(a) == 0 {
		return []TopicPartitionConfig{}
	}

	items := make([]TopicPartitionConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTopicPartitionConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandTopicPartitionConfig expands an instance of TopicPartitionConfig into a JSON
// request object.
func expandTopicPartitionConfig(c *Client, f *TopicPartitionConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Count; !dcl.IsEmptyValueIndirect(v) {
		m["count"] = v
	}
	if v, err := expandTopicPartitionConfigCapacity(c, f.Capacity); err != nil {
		return nil, fmt.Errorf("error expanding Capacity into capacity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["capacity"] = v
	}

	return m, nil
}

// flattenTopicPartitionConfig flattens an instance of TopicPartitionConfig from a JSON
// response object.
func flattenTopicPartitionConfig(c *Client, i interface{}) *TopicPartitionConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TopicPartitionConfig{}
	r.Count = dcl.FlattenInteger(m["count"])
	r.Capacity = flattenTopicPartitionConfigCapacity(c, m["capacity"])

	return r
}

// expandTopicPartitionConfigCapacityMap expands the contents of TopicPartitionConfigCapacity into a JSON
// request object.
func expandTopicPartitionConfigCapacityMap(c *Client, f map[string]TopicPartitionConfigCapacity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTopicPartitionConfigCapacity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTopicPartitionConfigCapacitySlice expands the contents of TopicPartitionConfigCapacity into a JSON
// request object.
func expandTopicPartitionConfigCapacitySlice(c *Client, f []TopicPartitionConfigCapacity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTopicPartitionConfigCapacity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTopicPartitionConfigCapacityMap flattens the contents of TopicPartitionConfigCapacity from a JSON
// response object.
func flattenTopicPartitionConfigCapacityMap(c *Client, i interface{}) map[string]TopicPartitionConfigCapacity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TopicPartitionConfigCapacity{}
	}

	if len(a) == 0 {
		return map[string]TopicPartitionConfigCapacity{}
	}

	items := make(map[string]TopicPartitionConfigCapacity)
	for k, item := range a {
		items[k] = *flattenTopicPartitionConfigCapacity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTopicPartitionConfigCapacitySlice flattens the contents of TopicPartitionConfigCapacity from a JSON
// response object.
func flattenTopicPartitionConfigCapacitySlice(c *Client, i interface{}) []TopicPartitionConfigCapacity {
	a, ok := i.([]interface{})
	if !ok {
		return []TopicPartitionConfigCapacity{}
	}

	if len(a) == 0 {
		return []TopicPartitionConfigCapacity{}
	}

	items := make([]TopicPartitionConfigCapacity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTopicPartitionConfigCapacity(c, item.(map[string]interface{})))
	}

	return items
}

// expandTopicPartitionConfigCapacity expands an instance of TopicPartitionConfigCapacity into a JSON
// request object.
func expandTopicPartitionConfigCapacity(c *Client, f *TopicPartitionConfigCapacity) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PublishMibPerSec; !dcl.IsEmptyValueIndirect(v) {
		m["publishMibPerSec"] = v
	}
	if v := f.SubscribeMibPerSec; !dcl.IsEmptyValueIndirect(v) {
		m["subscribeMibPerSec"] = v
	}

	return m, nil
}

// flattenTopicPartitionConfigCapacity flattens an instance of TopicPartitionConfigCapacity from a JSON
// response object.
func flattenTopicPartitionConfigCapacity(c *Client, i interface{}) *TopicPartitionConfigCapacity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TopicPartitionConfigCapacity{}
	r.PublishMibPerSec = dcl.FlattenInteger(m["publishMibPerSec"])
	r.SubscribeMibPerSec = dcl.FlattenInteger(m["subscribeMibPerSec"])

	return r
}

// expandTopicRetentionConfigMap expands the contents of TopicRetentionConfig into a JSON
// request object.
func expandTopicRetentionConfigMap(c *Client, f map[string]TopicRetentionConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTopicRetentionConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTopicRetentionConfigSlice expands the contents of TopicRetentionConfig into a JSON
// request object.
func expandTopicRetentionConfigSlice(c *Client, f []TopicRetentionConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTopicRetentionConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTopicRetentionConfigMap flattens the contents of TopicRetentionConfig from a JSON
// response object.
func flattenTopicRetentionConfigMap(c *Client, i interface{}) map[string]TopicRetentionConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TopicRetentionConfig{}
	}

	if len(a) == 0 {
		return map[string]TopicRetentionConfig{}
	}

	items := make(map[string]TopicRetentionConfig)
	for k, item := range a {
		items[k] = *flattenTopicRetentionConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTopicRetentionConfigSlice flattens the contents of TopicRetentionConfig from a JSON
// response object.
func flattenTopicRetentionConfigSlice(c *Client, i interface{}) []TopicRetentionConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []TopicRetentionConfig{}
	}

	if len(a) == 0 {
		return []TopicRetentionConfig{}
	}

	items := make([]TopicRetentionConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTopicRetentionConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandTopicRetentionConfig expands an instance of TopicRetentionConfig into a JSON
// request object.
func expandTopicRetentionConfig(c *Client, f *TopicRetentionConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PerPartitionBytes; !dcl.IsEmptyValueIndirect(v) {
		m["perPartitionBytes"] = v
	}
	if v := f.Period; !dcl.IsEmptyValueIndirect(v) {
		m["period"] = v
	}

	return m, nil
}

// flattenTopicRetentionConfig flattens an instance of TopicRetentionConfig from a JSON
// response object.
func flattenTopicRetentionConfig(c *Client, i interface{}) *TopicRetentionConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TopicRetentionConfig{}
	r.PerPartitionBytes = dcl.FlattenInteger(m["perPartitionBytes"])
	r.Period = dcl.FlattenString(m["period"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Topic) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTopic(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}
