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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Subscription) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "topic"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DeliveryConfig) {
		if err := r.DeliveryConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *SubscriptionDeliveryConfig) validate() error {
	return nil
}

func subscriptionGetURL(userBasePath string, r *Subscription) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/subscriptions/{{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil
}

func subscriptionListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/subscriptions", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil

}

func subscriptionCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/subscriptions?subscriptionId={{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil

}

func subscriptionDeleteURL(userBasePath string, r *Subscription) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("admin/projects/{{project}}/locations/{{location}}/subscriptions/{{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, params), nil
}

// subscriptionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type subscriptionApiOperation interface {
	do(context.Context, *Subscription, *Client) error
}

// newUpdateSubscriptionUpdateSubscriptionRequest creates a request for an
// Subscription resource's UpdateSubscription update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSubscriptionUpdateSubscriptionRequest(ctx context.Context, f *Subscription, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandSubscriptionDeliveryConfig(c, f.DeliveryConfig); err != nil {
		return nil, fmt.Errorf("error expanding DeliveryConfig into deliveryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["deliveryConfig"] = v
	}
	return req, nil
}

// marshalUpdateSubscriptionUpdateSubscriptionRequest converts the update into
// the final JSON request body.
func marshalUpdateSubscriptionUpdateSubscriptionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateSubscriptionUpdateSubscriptionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSubscriptionUpdateSubscriptionOperation) do(ctx context.Context, r *Subscription, c *Client) error {
	_, err := c.GetSubscription(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateSubscription")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"deliveryConfig"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateSubscriptionUpdateSubscriptionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSubscriptionUpdateSubscriptionRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listSubscriptionRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := subscriptionListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != SubscriptionMaxPage {
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

type listSubscriptionOperation struct {
	Subscriptions []map[string]interface{} `json:"subscriptions"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listSubscription(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Subscription, string, error) {
	b, err := c.listSubscriptionRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSubscriptionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Subscription
	for _, v := range m.Subscriptions {
		res := flattenSubscription(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllSubscription(ctx context.Context, f func(*Subscription) bool, resources []*Subscription) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteSubscription(ctx, res)
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

type deleteSubscriptionOperation struct{}

func (op *deleteSubscriptionOperation) do(ctx context.Context, r *Subscription, c *Client) error {

	_, err := c.GetSubscription(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Subscription not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSubscription checking for existence. error: %v", err)
		return err
	}

	u, err := subscriptionDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Subscription: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetSubscription(ctx, r.urlNormalized())
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
type createSubscriptionOperation struct {
	response map[string]interface{}
}

func (op *createSubscriptionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createSubscriptionOperation) do(ctx context.Context, r *Subscription, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := subscriptionCreateURL(c.Config.BasePath, project, location, name)

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

	// Poll for the Subscription resource to be created. Subscription resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := subscriptionGetURL(c.Config.BasePath, r.urlNormalized())
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
		return getResp, nil
	}, c.Config.RetryProvider)

	if _, err := c.GetSubscription(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSubscriptionRaw(ctx context.Context, r *Subscription) ([]byte, error) {

	u, err := subscriptionGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) subscriptionDiffsForRawDesired(ctx context.Context, rawDesired *Subscription, opts ...dcl.ApplyOption) (initial, desired *Subscription, diffs []subscriptionDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Subscription
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Subscription); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Subscription, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetSubscription(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Subscription resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Subscription resource: %v", err)
		}
		c.Config.Logger.Info("Found that Subscription resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeSubscriptionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Subscription: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Subscription: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeSubscriptionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Subscription: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeSubscriptionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Subscription: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffSubscription(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeSubscriptionInitialState(rawInitial, rawDesired *Subscription) (*Subscription, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeSubscriptionDesiredState(rawDesired, rawInitial *Subscription, opts ...dcl.ApplyOption) (*Subscription, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DeliveryConfig = canonicalizeSubscriptionDeliveryConfig(rawDesired.DeliveryConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Topic, rawInitial.Topic) {
		rawDesired.Topic = rawInitial.Topic
	}
	rawDesired.DeliveryConfig = canonicalizeSubscriptionDeliveryConfig(rawDesired.DeliveryConfig, rawInitial.DeliveryConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeSubscriptionNewState(c *Client, rawNew, rawDesired *Subscription) (*Subscription, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Topic) && dcl.IsEmptyValueIndirect(rawDesired.Topic) {
		rawNew.Topic = rawDesired.Topic
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Topic, rawNew.Topic) {
			rawNew.Topic = rawDesired.Topic
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeliveryConfig) && dcl.IsEmptyValueIndirect(rawDesired.DeliveryConfig) {
		rawNew.DeliveryConfig = rawDesired.DeliveryConfig
	} else {
		rawNew.DeliveryConfig = canonicalizeNewSubscriptionDeliveryConfig(c, rawDesired.DeliveryConfig, rawNew.DeliveryConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeSubscriptionDeliveryConfig(des, initial *SubscriptionDeliveryConfig, opts ...dcl.ApplyOption) *SubscriptionDeliveryConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DeliveryRequirement) {
		des.DeliveryRequirement = initial.DeliveryRequirement
	}

	return des
}

func canonicalizeNewSubscriptionDeliveryConfig(c *Client, des, nw *SubscriptionDeliveryConfig) *SubscriptionDeliveryConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewSubscriptionDeliveryConfigSet(c *Client, des, nw []SubscriptionDeliveryConfig) []SubscriptionDeliveryConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []SubscriptionDeliveryConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareSubscriptionDeliveryConfig(c, &d, &n) {
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

func canonicalizeNewSubscriptionDeliveryConfigSlice(c *Client, des, nw []SubscriptionDeliveryConfig) []SubscriptionDeliveryConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []SubscriptionDeliveryConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubscriptionDeliveryConfig(c, &d, &n))
	}

	return items
}

type subscriptionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         subscriptionApiOperation
	Diffs            []*dcl.FieldDiff
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
func diffSubscription(c *Client, desired, actual *Subscription, opts ...dcl.ApplyOption) ([]subscriptionDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []subscriptionDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, subscriptionDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.Topic, actual.Topic, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Topic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, subscriptionDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Topic",
		})
	}

	if ds, err := dcl.Diff(desired.DeliveryConfig, actual.DeliveryConfig, dcl.Info{ObjectFunction: compareSubscriptionDeliveryConfigNewStyle}, fn.AddNest("DeliveryConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, subscriptionDiff{
			UpdateOp: &updateSubscriptionUpdateSubscriptionOperation{}, Diffs: ds,
			FieldName: "DeliveryConfig",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, subscriptionDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
		})
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, subscriptionDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Location",
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
	var deduped []subscriptionDiff
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
func compareSubscriptionDeliveryConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubscriptionDeliveryConfig)
	if !ok {
		desiredNotPointer, ok := d.(SubscriptionDeliveryConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionDeliveryConfig or *SubscriptionDeliveryConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubscriptionDeliveryConfig)
	if !ok {
		actualNotPointer, ok := a.(SubscriptionDeliveryConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionDeliveryConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DeliveryRequirement, actual.DeliveryRequirement, dcl.Info{Type: "EnumType"}, fn.AddNest("DeliveryRequirement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSubscriptionDeliveryConfig(c *Client, desired, actual *SubscriptionDeliveryConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.DeliveryRequirement, actual.DeliveryRequirement) && !dcl.IsZeroValue(desired.DeliveryRequirement) {
		c.Config.Logger.Infof("Diff in DeliveryRequirement.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DeliveryRequirement), dcl.SprintResource(actual.DeliveryRequirement))
		return true
	}
	return false
}

func compareSubscriptionDeliveryConfigSlice(c *Client, desired, actual []SubscriptionDeliveryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SubscriptionDeliveryConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareSubscriptionDeliveryConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in SubscriptionDeliveryConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareSubscriptionDeliveryConfigMap(c *Client, desired, actual map[string]SubscriptionDeliveryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SubscriptionDeliveryConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in SubscriptionDeliveryConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareSubscriptionDeliveryConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in SubscriptionDeliveryConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareSubscriptionDeliveryConfigDeliveryRequirementEnumSlice(c *Client, desired, actual []SubscriptionDeliveryConfigDeliveryRequirementEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in SubscriptionDeliveryConfigDeliveryRequirementEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareSubscriptionDeliveryConfigDeliveryRequirementEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in SubscriptionDeliveryConfigDeliveryRequirementEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareSubscriptionDeliveryConfigDeliveryRequirementEnum(c *Client, desired, actual *SubscriptionDeliveryConfigDeliveryRequirementEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Subscription) urlNormalized() *Subscription {
	normalized := dcl.Copy(*r).(Subscription)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Topic = dcl.SelfLinkToName(r.Topic)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Subscription) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subscription) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subscription) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subscription) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateSubscription" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("admin/projects/{{project}}/locations/{{location}}/subscriptions/{{name}}", "https://us-central1-pubsublite.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Subscription resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Subscription) marshal(c *Client) ([]byte, error) {
	m, err := expandSubscription(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Subscription: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalSubscription decodes JSON responses into the Subscription resource schema.
func unmarshalSubscription(b []byte, c *Client) (*Subscription, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapSubscription(m, c)
}

func unmarshalMapSubscription(m map[string]interface{}, c *Client) (*Subscription, error) {

	return flattenSubscription(c, m), nil
}

// expandSubscription expands Subscription into a JSON request object.
func expandSubscription(c *Client, f *Subscription) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/subscriptions/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/locations/%s/topics/%s", f.Topic, f.Project, f.Location, f.Topic); err != nil {
		return nil, fmt.Errorf("error expanding Topic into topic: %w", err)
	} else if v != nil {
		m["topic"] = v
	}
	if v, err := expandSubscriptionDeliveryConfig(c, f.DeliveryConfig); err != nil {
		return nil, fmt.Errorf("error expanding DeliveryConfig into deliveryConfig: %w", err)
	} else if v != nil {
		m["deliveryConfig"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenSubscription flattens Subscription from a JSON request object into the
// Subscription type.
func flattenSubscription(c *Client, i interface{}) *Subscription {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Subscription{}
	r.Name = dcl.FlattenString(m["name"])
	r.Topic = dcl.FlattenString(m["topic"])
	r.DeliveryConfig = flattenSubscriptionDeliveryConfig(c, m["deliveryConfig"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandSubscriptionDeliveryConfigMap expands the contents of SubscriptionDeliveryConfig into a JSON
// request object.
func expandSubscriptionDeliveryConfigMap(c *Client, f map[string]SubscriptionDeliveryConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubscriptionDeliveryConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubscriptionDeliveryConfigSlice expands the contents of SubscriptionDeliveryConfig into a JSON
// request object.
func expandSubscriptionDeliveryConfigSlice(c *Client, f []SubscriptionDeliveryConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubscriptionDeliveryConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubscriptionDeliveryConfigMap flattens the contents of SubscriptionDeliveryConfig from a JSON
// response object.
func flattenSubscriptionDeliveryConfigMap(c *Client, i interface{}) map[string]SubscriptionDeliveryConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubscriptionDeliveryConfig{}
	}

	if len(a) == 0 {
		return map[string]SubscriptionDeliveryConfig{}
	}

	items := make(map[string]SubscriptionDeliveryConfig)
	for k, item := range a {
		items[k] = *flattenSubscriptionDeliveryConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubscriptionDeliveryConfigSlice flattens the contents of SubscriptionDeliveryConfig from a JSON
// response object.
func flattenSubscriptionDeliveryConfigSlice(c *Client, i interface{}) []SubscriptionDeliveryConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []SubscriptionDeliveryConfig{}
	}

	if len(a) == 0 {
		return []SubscriptionDeliveryConfig{}
	}

	items := make([]SubscriptionDeliveryConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubscriptionDeliveryConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubscriptionDeliveryConfig expands an instance of SubscriptionDeliveryConfig into a JSON
// request object.
func expandSubscriptionDeliveryConfig(c *Client, f *SubscriptionDeliveryConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.DeliveryRequirement; !dcl.IsEmptyValueIndirect(v) {
		m["deliveryRequirement"] = v
	}

	return m, nil
}

// flattenSubscriptionDeliveryConfig flattens an instance of SubscriptionDeliveryConfig from a JSON
// response object.
func flattenSubscriptionDeliveryConfig(c *Client, i interface{}) *SubscriptionDeliveryConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubscriptionDeliveryConfig{}
	r.DeliveryRequirement = flattenSubscriptionDeliveryConfigDeliveryRequirementEnum(m["deliveryRequirement"])

	return r
}

// flattenSubscriptionDeliveryConfigDeliveryRequirementEnumSlice flattens the contents of SubscriptionDeliveryConfigDeliveryRequirementEnum from a JSON
// response object.
func flattenSubscriptionDeliveryConfigDeliveryRequirementEnumSlice(c *Client, i interface{}) []SubscriptionDeliveryConfigDeliveryRequirementEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SubscriptionDeliveryConfigDeliveryRequirementEnum{}
	}

	if len(a) == 0 {
		return []SubscriptionDeliveryConfigDeliveryRequirementEnum{}
	}

	items := make([]SubscriptionDeliveryConfigDeliveryRequirementEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubscriptionDeliveryConfigDeliveryRequirementEnum(item.(interface{})))
	}

	return items
}

// flattenSubscriptionDeliveryConfigDeliveryRequirementEnum asserts that an interface is a string, and returns a
// pointer to a *SubscriptionDeliveryConfigDeliveryRequirementEnum with the same value as that string.
func flattenSubscriptionDeliveryConfigDeliveryRequirementEnum(i interface{}) *SubscriptionDeliveryConfigDeliveryRequirementEnum {
	s, ok := i.(string)
	if !ok {
		return SubscriptionDeliveryConfigDeliveryRequirementEnumRef("")
	}

	return SubscriptionDeliveryConfigDeliveryRequirementEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Subscription) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalSubscription(b, c)
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
