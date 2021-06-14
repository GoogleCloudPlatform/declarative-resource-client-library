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
package monitoring

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

func (r *NotificationChannel) validate() error {

	return nil
}

func notificationChannelGetURL(userBasePath string, r *NotificationChannel) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/notificationChannels/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

func notificationChannelListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/notificationChannels", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func notificationChannelCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/notificationChannels", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func notificationChannelDeleteURL(userBasePath string, r *NotificationChannel) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/notificationChannels/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

// notificationChannelApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type notificationChannelApiOperation interface {
	do(context.Context, *NotificationChannel, *Client) error
}

// newUpdateNotificationChannelUpdateRequest creates a request for an
// NotificationChannel resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNotificationChannelUpdateRequest(ctx context.Context, f *NotificationChannel, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		req["enabled"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		req["userLabels"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/notificationChannels/%s", *f.Project, *f.Name)

	return req, nil
}

// marshalUpdateNotificationChannelUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateNotificationChannelUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNotificationChannelUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNotificationChannelUpdateOperation) do(ctx context.Context, r *NotificationChannel, c *Client) error {
	_, err := c.GetNotificationChannel(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.Diffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateNotificationChannelUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNotificationChannelUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listNotificationChannelRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := notificationChannelListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NotificationChannelMaxPage {
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

type listNotificationChannelOperation struct {
	NotificationChannels []map[string]interface{} `json:"notificationChannels"`
	Token                string                   `json:"nextPageToken"`
}

func (c *Client) listNotificationChannel(ctx context.Context, project, pageToken string, pageSize int32) ([]*NotificationChannel, string, error) {
	b, err := c.listNotificationChannelRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNotificationChannelOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*NotificationChannel
	for _, v := range m.NotificationChannels {
		res, err := unmarshalMapNotificationChannel(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNotificationChannel(ctx context.Context, f func(*NotificationChannel) bool, resources []*NotificationChannel) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNotificationChannel(ctx, res)
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

type deleteNotificationChannelOperation struct{}

func (op *deleteNotificationChannelOperation) do(ctx context.Context, r *NotificationChannel, c *Client) error {

	_, err := c.GetNotificationChannel(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("NotificationChannel not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetNotificationChannel checking for existence. error: %v", err)
		return err
	}

	u, err := notificationChannelDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete NotificationChannel: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetNotificationChannel(ctx, r.urlNormalized())
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
type createNotificationChannelOperation struct {
	response map[string]interface{}
}

func (op *createNotificationChannelOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNotificationChannelOperation) do(ctx context.Context, r *NotificationChannel, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := notificationChannelCreateURL(c.Config.BasePath, project)

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

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetNotificationChannel(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNotificationChannelRaw(ctx context.Context, r *NotificationChannel) ([]byte, error) {
	if dcl.IsZeroValue(r.Enabled) {
		r.Enabled = dcl.Bool(true)
	}

	u, err := notificationChannelGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) notificationChannelDiffsForRawDesired(ctx context.Context, rawDesired *NotificationChannel, opts ...dcl.ApplyOption) (initial, desired *NotificationChannel, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *NotificationChannel
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*NotificationChannel); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected NotificationChannel, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeNotificationChannelDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNotificationChannel(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a NotificationChannel resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve NotificationChannel resource: %v", err)
		}
		c.Config.Logger.Info("Found that NotificationChannel resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNotificationChannelDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for NotificationChannel: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for NotificationChannel: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNotificationChannelInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for NotificationChannel: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNotificationChannelDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for NotificationChannel: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNotificationChannel(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNotificationChannelInitialState(rawInitial, rawDesired *NotificationChannel) (*NotificationChannel, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNotificationChannelDesiredState(rawDesired, rawInitial *NotificationChannel, opts ...dcl.ApplyOption) (*NotificationChannel, error) {

	if dcl.IsZeroValue(rawDesired.Enabled) {
		rawDesired.Enabled = dcl.Bool(true)
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.BoolCanonicalize(rawDesired.Enabled, rawInitial.Enabled) {
		rawDesired.Enabled = rawInitial.Enabled
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.Type, rawInitial.Type) {
		rawDesired.Type = rawInitial.Type
	}
	if dcl.IsZeroValue(rawDesired.UserLabels) {
		rawDesired.UserLabels = rawInitial.UserLabels
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeNotificationChannelNewState(c *Client, rawNew, rawDesired *NotificationChannel) (*NotificationChannel, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Enabled) && dcl.IsEmptyValueIndirect(rawDesired.Enabled) {
		rawNew.Enabled = rawDesired.Enabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Enabled, rawNew.Enabled) {
			rawNew.Enabled = rawDesired.Enabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
		if dcl.StringCanonicalize(rawDesired.Type, rawNew.Type) {
			rawNew.Type = rawDesired.Type
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UserLabels) && dcl.IsEmptyValueIndirect(rawDesired.UserLabels) {
		rawNew.UserLabels = rawDesired.UserLabels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VerificationStatus) && dcl.IsEmptyValueIndirect(rawDesired.VerificationStatus) {
		rawNew.VerificationStatus = rawDesired.VerificationStatus
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffNotificationChannel(c *Client, desired, actual *NotificationChannel, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNotificationChannelUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNotificationChannelUpdateOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNotificationChannelUpdateOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNotificationChannelUpdateOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNotificationChannelUpdateOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserLabels, actual.UserLabels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNotificationChannelUpdateOperation")}, fn.AddNest("UserLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VerificationStatus, actual.VerificationStatus, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VerificationStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *NotificationChannel) urlNormalized() *NotificationChannel {
	normalized := dcl.Copy(*r).(NotificationChannel)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Type = dcl.SelfLinkToName(r.Type)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *NotificationChannel) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *NotificationChannel) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *NotificationChannel) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *NotificationChannel) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("v3/projects/{{project}}/notificationChannels/{{name}}", "https://monitoring.googleapis.com/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the NotificationChannel resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *NotificationChannel) marshal(c *Client) ([]byte, error) {
	m, err := expandNotificationChannel(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling NotificationChannel: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNotificationChannel decodes JSON responses into the NotificationChannel resource schema.
func unmarshalNotificationChannel(b []byte, c *Client) (*NotificationChannel, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNotificationChannel(m, c)
}

func unmarshalMapNotificationChannel(m map[string]interface{}, c *Client) (*NotificationChannel, error) {

	return flattenNotificationChannel(c, m), nil
}

// expandNotificationChannel expands NotificationChannel into a JSON request object.
func expandNotificationChannel(c *Client, f *NotificationChannel) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		m["userLabels"] = v
	}
	if v := f.VerificationStatus; !dcl.IsEmptyValueIndirect(v) {
		m["verificationStatus"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenNotificationChannel flattens NotificationChannel from a JSON request object into the
// NotificationChannel type.
func flattenNotificationChannel(c *Client, i interface{}) *NotificationChannel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &NotificationChannel{}
	res.Description = dcl.FlattenString(m["description"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Enabled = dcl.FlattenBool(m["enabled"])
	if _, ok := m["enabled"]; !ok {
		c.Config.Logger.Info("Using default value for enabled")
		res.Enabled = dcl.Bool(true)
	}
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.Type = dcl.FlattenString(m["type"])
	res.UserLabels = dcl.FlattenKeyValuePairs(m["userLabels"])
	res.VerificationStatus = flattenNotificationChannelVerificationStatusEnum(m["verificationStatus"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// flattenNotificationChannelVerificationStatusEnumSlice flattens the contents of NotificationChannelVerificationStatusEnum from a JSON
// response object.
func flattenNotificationChannelVerificationStatusEnumSlice(c *Client, i interface{}) []NotificationChannelVerificationStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotificationChannelVerificationStatusEnum{}
	}

	if len(a) == 0 {
		return []NotificationChannelVerificationStatusEnum{}
	}

	items := make([]NotificationChannelVerificationStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotificationChannelVerificationStatusEnum(item.(interface{})))
	}

	return items
}

// flattenNotificationChannelVerificationStatusEnum asserts that an interface is a string, and returns a
// pointer to a *NotificationChannelVerificationStatusEnum with the same value as that string.
func flattenNotificationChannelVerificationStatusEnum(i interface{}) *NotificationChannelVerificationStatusEnum {
	s, ok := i.(string)
	if !ok {
		return NotificationChannelVerificationStatusEnumRef("")
	}

	return NotificationChannelVerificationStatusEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *NotificationChannel) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNotificationChannel(b, c)
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

type notificationChannelDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         notificationChannelApiOperation
}

func convertFieldDiffToNotificationChannelOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]notificationChannelDiff, error) {
	var diffs []notificationChannelDiff
	for _, op := range ops {
		diff := notificationChannelDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTonotificationChannelApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTonotificationChannelApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (notificationChannelApiOperation, error) {
	switch op {

	case "updateNotificationChannelUpdateOperation":
		return &updateNotificationChannelUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
