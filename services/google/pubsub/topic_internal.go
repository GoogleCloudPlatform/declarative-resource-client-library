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
package pubsub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Topic) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MessageStoragePolicy) {
		if err := r.MessageStoragePolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TopicMessageStoragePolicy) validate() error {
	return nil
}

func topicGetURL(userBasePath string, r *Topic) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/topics/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil
}

func topicListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/topics", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil

}

func topicCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/topics/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil

}

func topicDeleteURL(userBasePath string, r *Topic) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/topics/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil
}

// topicApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type topicApiOperation interface {
	do(context.Context, *Topic, *Client) error
}

// newUpdateTopicUpdateRequest creates a request for an
// Topic resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTopicUpdateRequest(ctx context.Context, f *Topic, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		req["kmsKeyName"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandTopicMessageStoragePolicy(c, f.MessageStoragePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MessageStoragePolicy into messageStoragePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["messageStoragePolicy"] = v
	}
	return req, nil
}

// marshalUpdateTopicUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateTopicUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(map[string]interface{}{"topic": m})

}

type updateTopicUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTopicUpdateOperation) do(ctx context.Context, r *Topic, c *Client) error {
	_, err := c.GetTopic(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateTopicUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTopicUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTopicRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := topicListURL(c.Config.BasePath, project)
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

func (c *Client) listTopic(ctx context.Context, project, pageToken string, pageSize int32) ([]*Topic, string, error) {
	b, err := c.listTopicRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTopicOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Topic
	for _, v := range m.Topics {
		res, err := unmarshalMapTopic(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
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
	r, err := c.GetTopic(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Topic not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTopic checking for existence. error: %v", err)
		return err
	}

	u, err := topicDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetTopic(ctx, r.URLNormalized())
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

	project, name := r.createFields()
	u, err := topicCreateURL(c.Config.BasePath, project, name)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(req), c.Config.RetryProvider)
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
	requiredSuccesses := 1
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := topicGetURL(c.Config.BasePath, r.URLNormalized())
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		getResp.Response.Body.Close()
		requiredSuccesses--
		if requiredSuccesses > 0 {
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return getResp, nil
	}, c.Config.RetryProvider)

	if _, err := c.GetTopic(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTopicRaw(ctx context.Context, r *Topic) ([]byte, error) {

	u, err := topicGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) topicDiffsForRawDesired(ctx context.Context, rawDesired *Topic, opts ...dcl.ApplyOption) (initial, desired *Topic, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetTopic(ctx, fetchState.URLNormalized())
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
		rawDesired.MessageStoragePolicy = canonicalizeTopicMessageStoragePolicy(rawDesired.MessageStoragePolicy, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Topic{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.KmsKeyName, rawInitial.KmsKeyName) {
		canonicalDesired.KmsKeyName = rawInitial.KmsKeyName
	} else {
		canonicalDesired.KmsKeyName = rawDesired.KmsKeyName
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.MessageStoragePolicy = canonicalizeTopicMessageStoragePolicy(rawDesired.MessageStoragePolicy, rawInitial.MessageStoragePolicy, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeTopicNewState(c *Client, rawNew, rawDesired *Topic) (*Topic, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.KmsKeyName) && dcl.IsEmptyValueIndirect(rawDesired.KmsKeyName) {
		rawNew.KmsKeyName = rawDesired.KmsKeyName
	} else {
		if dcl.StringCanonicalize(rawDesired.KmsKeyName, rawNew.KmsKeyName) {
			rawNew.KmsKeyName = rawDesired.KmsKeyName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MessageStoragePolicy) && dcl.IsEmptyValueIndirect(rawDesired.MessageStoragePolicy) {
		rawNew.MessageStoragePolicy = rawDesired.MessageStoragePolicy
	} else {
		rawNew.MessageStoragePolicy = canonicalizeNewTopicMessageStoragePolicy(c, rawDesired.MessageStoragePolicy, rawNew.MessageStoragePolicy)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeTopicMessageStoragePolicy(des, initial *TopicMessageStoragePolicy, opts ...dcl.ApplyOption) *TopicMessageStoragePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TopicMessageStoragePolicy{}

	if dcl.IsZeroValue(des.AllowedPersistenceRegions) {
		des.AllowedPersistenceRegions = initial.AllowedPersistenceRegions
	} else {
		cDes.AllowedPersistenceRegions = des.AllowedPersistenceRegions
	}

	return cDes
}

func canonicalizeNewTopicMessageStoragePolicy(c *Client, des, nw *TopicMessageStoragePolicy) *TopicMessageStoragePolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowedPersistenceRegions) {
		nw.AllowedPersistenceRegions = des.AllowedPersistenceRegions
	}

	return nw
}

func canonicalizeNewTopicMessageStoragePolicySet(c *Client, des, nw []TopicMessageStoragePolicy) []TopicMessageStoragePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []TopicMessageStoragePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTopicMessageStoragePolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewTopicMessageStoragePolicySlice(c *Client, des, nw []TopicMessageStoragePolicy) []TopicMessageStoragePolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TopicMessageStoragePolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTopicMessageStoragePolicy(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffTopic(c *Client, desired, actual *Topic, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTopicUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTopicUpdateOperation")}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTopicUpdateOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MessageStoragePolicy, actual.MessageStoragePolicy, dcl.Info{ObjectFunction: compareTopicMessageStoragePolicyNewStyle, EmptyObject: EmptyTopicMessageStoragePolicy, OperationSelector: dcl.TriggersOperation("updateTopicUpdateOperation")}, fn.AddNest("MessageStoragePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareTopicMessageStoragePolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TopicMessageStoragePolicy)
	if !ok {
		desiredNotPointer, ok := d.(TopicMessageStoragePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TopicMessageStoragePolicy or *TopicMessageStoragePolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TopicMessageStoragePolicy)
	if !ok {
		actualNotPointer, ok := a.(TopicMessageStoragePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TopicMessageStoragePolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedPersistenceRegions, actual.AllowedPersistenceRegions, dcl.Info{Type: "Set", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedPersistenceRegions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Topic) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Topic) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Topic) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Topic) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/topics/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, fields), nil

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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandTopicMessageStoragePolicy(c, f.MessageStoragePolicy); err != nil {
		return nil, fmt.Errorf("error expanding MessageStoragePolicy into messageStoragePolicy: %w", err)
	} else if v != nil {
		m["messageStoragePolicy"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
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

	res := &Topic{}
	res.Name = dcl.FlattenString(m["name"])
	res.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.MessageStoragePolicy = flattenTopicMessageStoragePolicy(c, m["messageStoragePolicy"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandTopicMessageStoragePolicyMap expands the contents of TopicMessageStoragePolicy into a JSON
// request object.
func expandTopicMessageStoragePolicyMap(c *Client, f map[string]TopicMessageStoragePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTopicMessageStoragePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTopicMessageStoragePolicySlice expands the contents of TopicMessageStoragePolicy into a JSON
// request object.
func expandTopicMessageStoragePolicySlice(c *Client, f []TopicMessageStoragePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTopicMessageStoragePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTopicMessageStoragePolicyMap flattens the contents of TopicMessageStoragePolicy from a JSON
// response object.
func flattenTopicMessageStoragePolicyMap(c *Client, i interface{}) map[string]TopicMessageStoragePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TopicMessageStoragePolicy{}
	}

	if len(a) == 0 {
		return map[string]TopicMessageStoragePolicy{}
	}

	items := make(map[string]TopicMessageStoragePolicy)
	for k, item := range a {
		items[k] = *flattenTopicMessageStoragePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTopicMessageStoragePolicySlice flattens the contents of TopicMessageStoragePolicy from a JSON
// response object.
func flattenTopicMessageStoragePolicySlice(c *Client, i interface{}) []TopicMessageStoragePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []TopicMessageStoragePolicy{}
	}

	if len(a) == 0 {
		return []TopicMessageStoragePolicy{}
	}

	items := make([]TopicMessageStoragePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTopicMessageStoragePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandTopicMessageStoragePolicy expands an instance of TopicMessageStoragePolicy into a JSON
// request object.
func expandTopicMessageStoragePolicy(c *Client, f *TopicMessageStoragePolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowedPersistenceRegions; !dcl.IsEmptyValueIndirect(v) {
		m["allowedPersistenceRegions"] = v
	}

	return m, nil
}

// flattenTopicMessageStoragePolicy flattens an instance of TopicMessageStoragePolicy from a JSON
// response object.
func flattenTopicMessageStoragePolicy(c *Client, i interface{}) *TopicMessageStoragePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TopicMessageStoragePolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTopicMessageStoragePolicy
	}
	r.AllowedPersistenceRegions = dcl.FlattenStringSlice(m["allowedPersistenceRegions"])

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
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
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

type topicDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         topicApiOperation
}

func convertFieldDiffsToTopicDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]topicDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []topicDiff
	// For each operation name, create a topicDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := topicDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTopicApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTopicApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (topicApiOperation, error) {
	switch opName {

	case "updateTopicUpdateOperation":
		return &updateTopicUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
