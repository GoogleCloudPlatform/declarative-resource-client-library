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

func (r *Subscription) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "topic"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ExpirationPolicy) {
		if err := r.ExpirationPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DeadLetterPolicy) {
		if err := r.DeadLetterPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PushConfig) {
		if err := r.PushConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *SubscriptionExpirationPolicy) validate() error {
	return nil
}
func (r *SubscriptionDeadLetterPolicy) validate() error {
	return nil
}
func (r *SubscriptionPushConfig) validate() error {
	if err := dcl.Required(r, "pushEndpoint"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.OidcToken) {
		if err := r.OidcToken.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *SubscriptionPushConfigOidcToken) validate() error {
	return nil
}

func subscriptionGetURL(userBasePath string, r *Subscription) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/subscriptions/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil
}

func subscriptionListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/subscriptions", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil

}

func subscriptionCreateURL(userBasePath, project, name string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"name":    name,
	}
	return dcl.URL("projects/{{project}}/subscriptions/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil

}

func subscriptionDeleteURL(userBasePath string, r *Subscription) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/subscriptions/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, params), nil
}

// subscriptionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type subscriptionApiOperation interface {
	do(context.Context, *Subscription, *Client) error
}

// newUpdateSubscriptionUpdateRequest creates a request for an
// Subscription resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSubscriptionUpdateRequest(ctx context.Context, f *Subscription, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.MessageRetentionDuration; !dcl.IsEmptyValueIndirect(v) {
		req["messageRetentionDuration"] = v
	}
	if v := f.RetainAckedMessages; !dcl.IsEmptyValueIndirect(v) {
		req["retainAckedMessages"] = v
	}
	if v, err := expandSubscriptionExpirationPolicy(c, f.ExpirationPolicy); err != nil {
		return nil, fmt.Errorf("error expanding ExpirationPolicy into expirationPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["expirationPolicy"] = v
	}
	return req, nil
}

// marshalUpdateSubscriptionUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateSubscriptionUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(map[string]interface{}{"subscription": m})

}

type updateSubscriptionUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSubscriptionUpdateOperation) do(ctx context.Context, r *Subscription, c *Client) error {
	_, err := c.GetSubscription(ctx, r.URLNormalized())
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

	req, err := newUpdateSubscriptionUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSubscriptionUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listSubscriptionRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := subscriptionListURL(c.Config.BasePath, project)
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

func (c *Client) listSubscription(ctx context.Context, project, pageToken string, pageSize int32) ([]*Subscription, string, error) {
	b, err := c.listSubscriptionRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSubscriptionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Subscription
	for _, v := range m.Subscriptions {
		res, err := unmarshalMapSubscription(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
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
	r, err := c.GetSubscription(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Subscription not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSubscription checking for existence. error: %v", err)
		return err
	}

	u, err := subscriptionDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetSubscription(ctx, r.URLNormalized())
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

	project, name := r.createFields()
	u, err := subscriptionCreateURL(c.Config.BasePath, project, name)

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

	// Poll for the Subscription resource to be created. Subscription resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	requiredSuccesses := 1
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := subscriptionGetURL(c.Config.BasePath, r.URLNormalized())
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

	if _, err := c.GetSubscription(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSubscriptionRaw(ctx context.Context, r *Subscription) ([]byte, error) {

	u, err := subscriptionGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) subscriptionDiffsForRawDesired(ctx context.Context, rawDesired *Subscription, opts ...dcl.ApplyOption) (initial, desired *Subscription, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetSubscription(ctx, fetchState.URLNormalized())
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
		rawDesired.ExpirationPolicy = canonicalizeSubscriptionExpirationPolicy(rawDesired.ExpirationPolicy, nil, opts...)
		rawDesired.DeadLetterPolicy = canonicalizeSubscriptionDeadLetterPolicy(rawDesired.DeadLetterPolicy, nil, opts...)
		rawDesired.PushConfig = canonicalizeSubscriptionPushConfig(rawDesired.PushConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Subscription{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Topic, rawInitial.Topic) {
		canonicalDesired.Topic = rawInitial.Topic
	} else {
		canonicalDesired.Topic = rawDesired.Topic
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.StringCanonicalize(rawDesired.MessageRetentionDuration, rawInitial.MessageRetentionDuration) {
		canonicalDesired.MessageRetentionDuration = rawInitial.MessageRetentionDuration
	} else {
		canonicalDesired.MessageRetentionDuration = rawDesired.MessageRetentionDuration
	}
	if dcl.BoolCanonicalize(rawDesired.RetainAckedMessages, rawInitial.RetainAckedMessages) {
		canonicalDesired.RetainAckedMessages = rawInitial.RetainAckedMessages
	} else {
		canonicalDesired.RetainAckedMessages = rawDesired.RetainAckedMessages
	}
	canonicalDesired.ExpirationPolicy = canonicalizeSubscriptionExpirationPolicy(rawDesired.ExpirationPolicy, rawInitial.ExpirationPolicy, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	canonicalDesired.DeadLetterPolicy = canonicalizeSubscriptionDeadLetterPolicy(rawDesired.DeadLetterPolicy, rawInitial.DeadLetterPolicy, opts...)
	canonicalDesired.PushConfig = canonicalizeSubscriptionPushConfig(rawDesired.PushConfig, rawInitial.PushConfig, opts...)
	if dcl.IsZeroValue(rawDesired.AckDeadlineSeconds) {
		canonicalDesired.AckDeadlineSeconds = rawInitial.AckDeadlineSeconds
	} else {
		canonicalDesired.AckDeadlineSeconds = rawDesired.AckDeadlineSeconds
	}

	return canonicalDesired, nil
}

func canonicalizeSubscriptionNewState(c *Client, rawNew, rawDesired *Subscription) (*Subscription, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Topic) && dcl.IsEmptyValueIndirect(rawDesired.Topic) {
		rawNew.Topic = rawDesired.Topic
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Topic, rawNew.Topic) {
			rawNew.Topic = rawDesired.Topic
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MessageRetentionDuration) && dcl.IsEmptyValueIndirect(rawDesired.MessageRetentionDuration) {
		rawNew.MessageRetentionDuration = rawDesired.MessageRetentionDuration
	} else {
		if dcl.StringCanonicalize(rawDesired.MessageRetentionDuration, rawNew.MessageRetentionDuration) {
			rawNew.MessageRetentionDuration = rawDesired.MessageRetentionDuration
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RetainAckedMessages) && dcl.IsEmptyValueIndirect(rawDesired.RetainAckedMessages) {
		rawNew.RetainAckedMessages = rawDesired.RetainAckedMessages
	} else {
		if dcl.BoolCanonicalize(rawDesired.RetainAckedMessages, rawNew.RetainAckedMessages) {
			rawNew.RetainAckedMessages = rawDesired.RetainAckedMessages
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpirationPolicy) && dcl.IsEmptyValueIndirect(rawDesired.ExpirationPolicy) {
		rawNew.ExpirationPolicy = rawDesired.ExpirationPolicy
	} else {
		rawNew.ExpirationPolicy = canonicalizeNewSubscriptionExpirationPolicy(c, rawDesired.ExpirationPolicy, rawNew.ExpirationPolicy)
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.DeadLetterPolicy) && dcl.IsEmptyValueIndirect(rawDesired.DeadLetterPolicy) {
		rawNew.DeadLetterPolicy = rawDesired.DeadLetterPolicy
	} else {
		rawNew.DeadLetterPolicy = canonicalizeNewSubscriptionDeadLetterPolicy(c, rawDesired.DeadLetterPolicy, rawNew.DeadLetterPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PushConfig) && dcl.IsEmptyValueIndirect(rawDesired.PushConfig) {
		rawNew.PushConfig = rawDesired.PushConfig
	} else {
		rawNew.PushConfig = canonicalizeNewSubscriptionPushConfig(c, rawDesired.PushConfig, rawNew.PushConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AckDeadlineSeconds) && dcl.IsEmptyValueIndirect(rawDesired.AckDeadlineSeconds) {
		rawNew.AckDeadlineSeconds = rawDesired.AckDeadlineSeconds
	} else {
	}

	return rawNew, nil
}

func canonicalizeSubscriptionExpirationPolicy(des, initial *SubscriptionExpirationPolicy, opts ...dcl.ApplyOption) *SubscriptionExpirationPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &SubscriptionExpirationPolicy{}

	if dcl.StringCanonicalize(des.Ttl, initial.Ttl) || dcl.IsZeroValue(des.Ttl) {
		cDes.Ttl = initial.Ttl
	} else {
		cDes.Ttl = des.Ttl
	}

	return cDes
}

func canonicalizeNewSubscriptionExpirationPolicy(c *Client, des, nw *SubscriptionExpirationPolicy) *SubscriptionExpirationPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Ttl, nw.Ttl) {
		nw.Ttl = des.Ttl
	}

	return nw
}

func canonicalizeNewSubscriptionExpirationPolicySet(c *Client, des, nw []SubscriptionExpirationPolicy) []SubscriptionExpirationPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []SubscriptionExpirationPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSubscriptionExpirationPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSubscriptionExpirationPolicySlice(c *Client, des, nw []SubscriptionExpirationPolicy) []SubscriptionExpirationPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SubscriptionExpirationPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubscriptionExpirationPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeSubscriptionDeadLetterPolicy(des, initial *SubscriptionDeadLetterPolicy, opts ...dcl.ApplyOption) *SubscriptionDeadLetterPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &SubscriptionDeadLetterPolicy{}

	if dcl.NameToSelfLink(des.DeadLetterTopic, initial.DeadLetterTopic) || dcl.IsZeroValue(des.DeadLetterTopic) {
		cDes.DeadLetterTopic = initial.DeadLetterTopic
	} else {
		cDes.DeadLetterTopic = des.DeadLetterTopic
	}
	if dcl.IsZeroValue(des.MaxDeliveryAttempts) {
		des.MaxDeliveryAttempts = initial.MaxDeliveryAttempts
	} else {
		cDes.MaxDeliveryAttempts = des.MaxDeliveryAttempts
	}

	return cDes
}

func canonicalizeNewSubscriptionDeadLetterPolicy(c *Client, des, nw *SubscriptionDeadLetterPolicy) *SubscriptionDeadLetterPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.DeadLetterTopic, nw.DeadLetterTopic) {
		nw.DeadLetterTopic = des.DeadLetterTopic
	}

	return nw
}

func canonicalizeNewSubscriptionDeadLetterPolicySet(c *Client, des, nw []SubscriptionDeadLetterPolicy) []SubscriptionDeadLetterPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []SubscriptionDeadLetterPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSubscriptionDeadLetterPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSubscriptionDeadLetterPolicySlice(c *Client, des, nw []SubscriptionDeadLetterPolicy) []SubscriptionDeadLetterPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SubscriptionDeadLetterPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubscriptionDeadLetterPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeSubscriptionPushConfig(des, initial *SubscriptionPushConfig, opts ...dcl.ApplyOption) *SubscriptionPushConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.Attributes) {
		des.Attributes = map[string]string{"x-goog-version": "v1"}
	}

	if initial == nil {
		return des
	}

	cDes := &SubscriptionPushConfig{}

	if dcl.StringCanonicalize(des.PushEndpoint, initial.PushEndpoint) || dcl.IsZeroValue(des.PushEndpoint) {
		cDes.PushEndpoint = initial.PushEndpoint
	} else {
		cDes.PushEndpoint = des.PushEndpoint
	}
	if dcl.IsZeroValue(des.Attributes) {
		des.Attributes = initial.Attributes
	} else {
		cDes.Attributes = des.Attributes
	}
	cDes.OidcToken = canonicalizeSubscriptionPushConfigOidcToken(des.OidcToken, initial.OidcToken, opts...)

	return cDes
}

func canonicalizeNewSubscriptionPushConfig(c *Client, des, nw *SubscriptionPushConfig) *SubscriptionPushConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Attributes) {
		nw.Attributes = map[string]string{"x-goog-version": "v1"}
	}

	if dcl.StringCanonicalize(des.PushEndpoint, nw.PushEndpoint) {
		nw.PushEndpoint = des.PushEndpoint
	}
	nw.OidcToken = canonicalizeNewSubscriptionPushConfigOidcToken(c, des.OidcToken, nw.OidcToken)

	return nw
}

func canonicalizeNewSubscriptionPushConfigSet(c *Client, des, nw []SubscriptionPushConfig) []SubscriptionPushConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []SubscriptionPushConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSubscriptionPushConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSubscriptionPushConfigSlice(c *Client, des, nw []SubscriptionPushConfig) []SubscriptionPushConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SubscriptionPushConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubscriptionPushConfig(c, &d, &n))
	}

	return items
}

func canonicalizeSubscriptionPushConfigOidcToken(des, initial *SubscriptionPushConfigOidcToken, opts ...dcl.ApplyOption) *SubscriptionPushConfigOidcToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &SubscriptionPushConfigOidcToken{}

	if dcl.StringCanonicalize(des.ServiceAccountEmail, initial.ServiceAccountEmail) || dcl.IsZeroValue(des.ServiceAccountEmail) {
		cDes.ServiceAccountEmail = initial.ServiceAccountEmail
	} else {
		cDes.ServiceAccountEmail = des.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Audience, initial.Audience) || dcl.IsZeroValue(des.Audience) {
		cDes.Audience = initial.Audience
	} else {
		cDes.Audience = des.Audience
	}

	return cDes
}

func canonicalizeNewSubscriptionPushConfigOidcToken(c *Client, des, nw *SubscriptionPushConfigOidcToken) *SubscriptionPushConfigOidcToken {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ServiceAccountEmail, nw.ServiceAccountEmail) {
		nw.ServiceAccountEmail = des.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Audience, nw.Audience) {
		nw.Audience = des.Audience
	}

	return nw
}

func canonicalizeNewSubscriptionPushConfigOidcTokenSet(c *Client, des, nw []SubscriptionPushConfigOidcToken) []SubscriptionPushConfigOidcToken {
	if des == nil {
		return nw
	}
	var reorderedNew []SubscriptionPushConfigOidcToken
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSubscriptionPushConfigOidcTokenNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSubscriptionPushConfigOidcTokenSlice(c *Client, des, nw []SubscriptionPushConfigOidcToken) []SubscriptionPushConfigOidcToken {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SubscriptionPushConfigOidcToken
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubscriptionPushConfigOidcToken(c, &d, &n))
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
func diffSubscription(c *Client, desired, actual *Subscription, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubscriptionUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Topic, actual.Topic, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Topic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubscriptionUpdateOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MessageRetentionDuration, actual.MessageRetentionDuration, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubscriptionUpdateOperation")}, fn.AddNest("MessageRetentionDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetainAckedMessages, actual.RetainAckedMessages, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubscriptionUpdateOperation")}, fn.AddNest("RetainAckedMessages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpirationPolicy, actual.ExpirationPolicy, dcl.Info{ObjectFunction: compareSubscriptionExpirationPolicyNewStyle, EmptyObject: EmptySubscriptionExpirationPolicy, OperationSelector: dcl.TriggersOperation("updateSubscriptionUpdateOperation")}, fn.AddNest("ExpirationPolicy")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DeadLetterPolicy, actual.DeadLetterPolicy, dcl.Info{ObjectFunction: compareSubscriptionDeadLetterPolicyNewStyle, EmptyObject: EmptySubscriptionDeadLetterPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeadLetterPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PushConfig, actual.PushConfig, dcl.Info{ObjectFunction: compareSubscriptionPushConfigNewStyle, EmptyObject: EmptySubscriptionPushConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PushConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AckDeadlineSeconds, actual.AckDeadlineSeconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AckDeadlineSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareSubscriptionExpirationPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubscriptionExpirationPolicy)
	if !ok {
		desiredNotPointer, ok := d.(SubscriptionExpirationPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionExpirationPolicy or *SubscriptionExpirationPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubscriptionExpirationPolicy)
	if !ok {
		actualNotPointer, ok := a.(SubscriptionExpirationPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionExpirationPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Ttl, actual.Ttl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubscriptionUpdateOperation")}, fn.AddNest("Ttl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSubscriptionDeadLetterPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubscriptionDeadLetterPolicy)
	if !ok {
		desiredNotPointer, ok := d.(SubscriptionDeadLetterPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionDeadLetterPolicy or *SubscriptionDeadLetterPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubscriptionDeadLetterPolicy)
	if !ok {
		actualNotPointer, ok := a.(SubscriptionDeadLetterPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionDeadLetterPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DeadLetterTopic, actual.DeadLetterTopic, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeadLetterTopic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxDeliveryAttempts, actual.MaxDeliveryAttempts, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxDeliveryAttempts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSubscriptionPushConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubscriptionPushConfig)
	if !ok {
		desiredNotPointer, ok := d.(SubscriptionPushConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionPushConfig or *SubscriptionPushConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubscriptionPushConfig)
	if !ok {
		actualNotPointer, ok := a.(SubscriptionPushConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionPushConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PushEndpoint, actual.PushEndpoint, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PushEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Attributes, actual.Attributes, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Attributes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OidcToken, actual.OidcToken, dcl.Info{ObjectFunction: compareSubscriptionPushConfigOidcTokenNewStyle, EmptyObject: EmptySubscriptionPushConfigOidcToken, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OidcToken")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSubscriptionPushConfigOidcTokenNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubscriptionPushConfigOidcToken)
	if !ok {
		desiredNotPointer, ok := d.(SubscriptionPushConfigOidcToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionPushConfigOidcToken or *SubscriptionPushConfigOidcToken", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubscriptionPushConfigOidcToken)
	if !ok {
		actualNotPointer, ok := a.(SubscriptionPushConfigOidcToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubscriptionPushConfigOidcToken", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Audience, actual.Audience, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Audience")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Subscription) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subscription) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subscription) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subscription) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/subscriptions/{{name}}", "https://pubsub.googleapis.com/v1/", userBasePath, fields), nil

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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/topics/%s", f.Topic, f.Project, f.Topic); err != nil {
		return nil, fmt.Errorf("error expanding Topic into topic: %w", err)
	} else if v != nil {
		m["topic"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.MessageRetentionDuration; !dcl.IsEmptyValueIndirect(v) {
		m["messageRetentionDuration"] = v
	}
	if v := f.RetainAckedMessages; !dcl.IsEmptyValueIndirect(v) {
		m["retainAckedMessages"] = v
	}
	if v, err := expandSubscriptionExpirationPolicy(c, f.ExpirationPolicy); err != nil {
		return nil, fmt.Errorf("error expanding ExpirationPolicy into expirationPolicy: %w", err)
	} else if v != nil {
		m["expirationPolicy"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := expandSubscriptionDeadLetterPolicy(c, f.DeadLetterPolicy); err != nil {
		return nil, fmt.Errorf("error expanding DeadLetterPolicy into deadLetterPolicy: %w", err)
	} else if v != nil {
		m["deadLetterPolicy"] = v
	}
	if v, err := expandSubscriptionPushConfig(c, f.PushConfig); err != nil {
		return nil, fmt.Errorf("error expanding PushConfig into pushConfig: %w", err)
	} else if v != nil {
		m["pushConfig"] = v
	}
	if v := f.AckDeadlineSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["ackDeadlineSeconds"] = v
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

	res := &Subscription{}
	res.Name = dcl.FlattenString(m["name"])
	res.Topic = dcl.FlattenString(m["topic"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.MessageRetentionDuration = dcl.FlattenString(m["messageRetentionDuration"])
	res.RetainAckedMessages = dcl.FlattenBool(m["retainAckedMessages"])
	res.ExpirationPolicy = flattenSubscriptionExpirationPolicy(c, m["expirationPolicy"])
	res.Project = dcl.FlattenString(m["project"])
	res.DeadLetterPolicy = flattenSubscriptionDeadLetterPolicy(c, m["deadLetterPolicy"])
	res.PushConfig = flattenSubscriptionPushConfig(c, m["pushConfig"])
	res.AckDeadlineSeconds = dcl.FlattenInteger(m["ackDeadlineSeconds"])

	return res
}

// expandSubscriptionExpirationPolicyMap expands the contents of SubscriptionExpirationPolicy into a JSON
// request object.
func expandSubscriptionExpirationPolicyMap(c *Client, f map[string]SubscriptionExpirationPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubscriptionExpirationPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubscriptionExpirationPolicySlice expands the contents of SubscriptionExpirationPolicy into a JSON
// request object.
func expandSubscriptionExpirationPolicySlice(c *Client, f []SubscriptionExpirationPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubscriptionExpirationPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubscriptionExpirationPolicyMap flattens the contents of SubscriptionExpirationPolicy from a JSON
// response object.
func flattenSubscriptionExpirationPolicyMap(c *Client, i interface{}) map[string]SubscriptionExpirationPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubscriptionExpirationPolicy{}
	}

	if len(a) == 0 {
		return map[string]SubscriptionExpirationPolicy{}
	}

	items := make(map[string]SubscriptionExpirationPolicy)
	for k, item := range a {
		items[k] = *flattenSubscriptionExpirationPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubscriptionExpirationPolicySlice flattens the contents of SubscriptionExpirationPolicy from a JSON
// response object.
func flattenSubscriptionExpirationPolicySlice(c *Client, i interface{}) []SubscriptionExpirationPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []SubscriptionExpirationPolicy{}
	}

	if len(a) == 0 {
		return []SubscriptionExpirationPolicy{}
	}

	items := make([]SubscriptionExpirationPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubscriptionExpirationPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubscriptionExpirationPolicy expands an instance of SubscriptionExpirationPolicy into a JSON
// request object.
func expandSubscriptionExpirationPolicy(c *Client, f *SubscriptionExpirationPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Ttl; !dcl.IsEmptyValueIndirect(v) {
		m["ttl"] = v
	}

	return m, nil
}

// flattenSubscriptionExpirationPolicy flattens an instance of SubscriptionExpirationPolicy from a JSON
// response object.
func flattenSubscriptionExpirationPolicy(c *Client, i interface{}) *SubscriptionExpirationPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubscriptionExpirationPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySubscriptionExpirationPolicy
	}
	r.Ttl = dcl.FlattenString(m["ttl"])

	return r
}

// expandSubscriptionDeadLetterPolicyMap expands the contents of SubscriptionDeadLetterPolicy into a JSON
// request object.
func expandSubscriptionDeadLetterPolicyMap(c *Client, f map[string]SubscriptionDeadLetterPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubscriptionDeadLetterPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubscriptionDeadLetterPolicySlice expands the contents of SubscriptionDeadLetterPolicy into a JSON
// request object.
func expandSubscriptionDeadLetterPolicySlice(c *Client, f []SubscriptionDeadLetterPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubscriptionDeadLetterPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubscriptionDeadLetterPolicyMap flattens the contents of SubscriptionDeadLetterPolicy from a JSON
// response object.
func flattenSubscriptionDeadLetterPolicyMap(c *Client, i interface{}) map[string]SubscriptionDeadLetterPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubscriptionDeadLetterPolicy{}
	}

	if len(a) == 0 {
		return map[string]SubscriptionDeadLetterPolicy{}
	}

	items := make(map[string]SubscriptionDeadLetterPolicy)
	for k, item := range a {
		items[k] = *flattenSubscriptionDeadLetterPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubscriptionDeadLetterPolicySlice flattens the contents of SubscriptionDeadLetterPolicy from a JSON
// response object.
func flattenSubscriptionDeadLetterPolicySlice(c *Client, i interface{}) []SubscriptionDeadLetterPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []SubscriptionDeadLetterPolicy{}
	}

	if len(a) == 0 {
		return []SubscriptionDeadLetterPolicy{}
	}

	items := make([]SubscriptionDeadLetterPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubscriptionDeadLetterPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubscriptionDeadLetterPolicy expands an instance of SubscriptionDeadLetterPolicy into a JSON
// request object.
func expandSubscriptionDeadLetterPolicy(c *Client, f *SubscriptionDeadLetterPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DeadLetterTopic; !dcl.IsEmptyValueIndirect(v) {
		m["deadLetterTopic"] = v
	}
	if v := f.MaxDeliveryAttempts; !dcl.IsEmptyValueIndirect(v) {
		m["maxDeliveryAttempts"] = v
	}

	return m, nil
}

// flattenSubscriptionDeadLetterPolicy flattens an instance of SubscriptionDeadLetterPolicy from a JSON
// response object.
func flattenSubscriptionDeadLetterPolicy(c *Client, i interface{}) *SubscriptionDeadLetterPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubscriptionDeadLetterPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySubscriptionDeadLetterPolicy
	}
	r.DeadLetterTopic = dcl.FlattenString(m["deadLetterTopic"])
	r.MaxDeliveryAttempts = dcl.FlattenInteger(m["maxDeliveryAttempts"])

	return r
}

// expandSubscriptionPushConfigMap expands the contents of SubscriptionPushConfig into a JSON
// request object.
func expandSubscriptionPushConfigMap(c *Client, f map[string]SubscriptionPushConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubscriptionPushConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubscriptionPushConfigSlice expands the contents of SubscriptionPushConfig into a JSON
// request object.
func expandSubscriptionPushConfigSlice(c *Client, f []SubscriptionPushConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubscriptionPushConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubscriptionPushConfigMap flattens the contents of SubscriptionPushConfig from a JSON
// response object.
func flattenSubscriptionPushConfigMap(c *Client, i interface{}) map[string]SubscriptionPushConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubscriptionPushConfig{}
	}

	if len(a) == 0 {
		return map[string]SubscriptionPushConfig{}
	}

	items := make(map[string]SubscriptionPushConfig)
	for k, item := range a {
		items[k] = *flattenSubscriptionPushConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubscriptionPushConfigSlice flattens the contents of SubscriptionPushConfig from a JSON
// response object.
func flattenSubscriptionPushConfigSlice(c *Client, i interface{}) []SubscriptionPushConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []SubscriptionPushConfig{}
	}

	if len(a) == 0 {
		return []SubscriptionPushConfig{}
	}

	items := make([]SubscriptionPushConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubscriptionPushConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubscriptionPushConfig expands an instance of SubscriptionPushConfig into a JSON
// request object.
func expandSubscriptionPushConfig(c *Client, f *SubscriptionPushConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PushEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["pushEndpoint"] = v
	}
	if v := f.Attributes; !dcl.IsEmptyValueIndirect(v) {
		m["attributes"] = v
	}
	if v, err := expandSubscriptionPushConfigOidcToken(c, f.OidcToken); err != nil {
		return nil, fmt.Errorf("error expanding OidcToken into oidcToken: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oidcToken"] = v
	}

	return m, nil
}

// flattenSubscriptionPushConfig flattens an instance of SubscriptionPushConfig from a JSON
// response object.
func flattenSubscriptionPushConfig(c *Client, i interface{}) *SubscriptionPushConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubscriptionPushConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySubscriptionPushConfig
	}
	r.PushEndpoint = dcl.FlattenString(m["pushEndpoint"])
	r.Attributes = dcl.FlattenKeyValuePairs(m["attributes"])
	if dcl.IsEmptyValueIndirect(m["attributes"]) {
		c.Config.Logger.Info("Using default value for attributes.")
		r.Attributes = map[string]string{"x-goog-version": "v1"}
	}
	r.OidcToken = flattenSubscriptionPushConfigOidcToken(c, m["oidcToken"])

	return r
}

// expandSubscriptionPushConfigOidcTokenMap expands the contents of SubscriptionPushConfigOidcToken into a JSON
// request object.
func expandSubscriptionPushConfigOidcTokenMap(c *Client, f map[string]SubscriptionPushConfigOidcToken) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubscriptionPushConfigOidcToken(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubscriptionPushConfigOidcTokenSlice expands the contents of SubscriptionPushConfigOidcToken into a JSON
// request object.
func expandSubscriptionPushConfigOidcTokenSlice(c *Client, f []SubscriptionPushConfigOidcToken) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubscriptionPushConfigOidcToken(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubscriptionPushConfigOidcTokenMap flattens the contents of SubscriptionPushConfigOidcToken from a JSON
// response object.
func flattenSubscriptionPushConfigOidcTokenMap(c *Client, i interface{}) map[string]SubscriptionPushConfigOidcToken {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubscriptionPushConfigOidcToken{}
	}

	if len(a) == 0 {
		return map[string]SubscriptionPushConfigOidcToken{}
	}

	items := make(map[string]SubscriptionPushConfigOidcToken)
	for k, item := range a {
		items[k] = *flattenSubscriptionPushConfigOidcToken(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubscriptionPushConfigOidcTokenSlice flattens the contents of SubscriptionPushConfigOidcToken from a JSON
// response object.
func flattenSubscriptionPushConfigOidcTokenSlice(c *Client, i interface{}) []SubscriptionPushConfigOidcToken {
	a, ok := i.([]interface{})
	if !ok {
		return []SubscriptionPushConfigOidcToken{}
	}

	if len(a) == 0 {
		return []SubscriptionPushConfigOidcToken{}
	}

	items := make([]SubscriptionPushConfigOidcToken, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubscriptionPushConfigOidcToken(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubscriptionPushConfigOidcToken expands an instance of SubscriptionPushConfigOidcToken into a JSON
// request object.
func expandSubscriptionPushConfigOidcToken(c *Client, f *SubscriptionPushConfigOidcToken) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.Audience; !dcl.IsEmptyValueIndirect(v) {
		m["audience"] = v
	}

	return m, nil
}

// flattenSubscriptionPushConfigOidcToken flattens an instance of SubscriptionPushConfigOidcToken from a JSON
// response object.
func flattenSubscriptionPushConfigOidcToken(c *Client, i interface{}) *SubscriptionPushConfigOidcToken {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubscriptionPushConfigOidcToken{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySubscriptionPushConfigOidcToken
	}
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	r.Audience = dcl.FlattenString(m["audience"])

	return r
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

type subscriptionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         subscriptionApiOperation
}

func convertFieldDiffsToSubscriptionDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]subscriptionDiff, error) {
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
	var diffs []subscriptionDiff
	// For each operation name, create a subscriptionDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := subscriptionDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToSubscriptionApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToSubscriptionApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (subscriptionApiOperation, error) {
	switch opName {

	case "updateSubscriptionUpdateOperation":
		return &updateSubscriptionUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
