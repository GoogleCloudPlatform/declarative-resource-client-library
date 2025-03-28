// Copyright 2025 Google LLC. All Rights Reserved.
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
package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *LogView) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Bucket, "Bucket"); err != nil {
		return err
	}
	return nil
}
func (r *LogView) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://logging.googleapis.com/v2/", params)
}

func (r *LogView) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"bucket":   dcl.ValueOrEmptyString(nr.Bucket),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *LogView) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"bucket":   dcl.ValueOrEmptyString(nr.Bucket),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/buckets/{{bucket}}/views", nr.basePath(), userBasePath, params), nil

}

func (r *LogView) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"bucket":   dcl.ValueOrEmptyString(nr.Bucket),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/buckets/{{bucket}}/views?viewId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *LogView) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"bucket":   dcl.ValueOrEmptyString(nr.Bucket),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}", nr.basePath(), userBasePath, params), nil
}

// logViewApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type logViewApiOperation interface {
	do(context.Context, *LogView, *Client) error
}

// newUpdateLogViewUpdateLogViewRequest creates a request for an
// LogView resource's UpdateLogView update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateLogViewUpdateLogViewRequest(ctx context.Context, f *LogView, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		req["filter"] = v
	}
	return req, nil
}

// marshalUpdateLogViewUpdateLogViewRequest converts the update into
// the final JSON request body.
func marshalUpdateLogViewUpdateLogViewRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateLogViewUpdateLogViewOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateLogViewUpdateLogViewOperation) do(ctx context.Context, r *LogView, c *Client) error {
	_, err := c.GetLogView(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateLogView")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateLogViewUpdateLogViewRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateLogViewUpdateLogViewRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listLogViewRaw(ctx context.Context, r *LogView, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != LogViewMaxPage {
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

type listLogViewOperation struct {
	Views []map[string]interface{} `json:"views"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listLogView(ctx context.Context, r *LogView, pageToken string, pageSize int32) ([]*LogView, string, error) {
	b, err := c.listLogViewRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listLogViewOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*LogView
	for _, v := range m.Views {
		res, err := unmarshalMapLogView(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.Bucket = r.Bucket
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllLogView(ctx context.Context, f func(*LogView) bool, resources []*LogView) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteLogView(ctx, res)
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

type deleteLogViewOperation struct{}

func (op *deleteLogViewOperation) do(ctx context.Context, r *LogView, c *Client) error {
	r, err := c.GetLogView(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "LogView not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetLogView checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete LogView: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetLogView(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createLogViewOperation struct {
	response map[string]interface{}
}

func (op *createLogViewOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createLogViewOperation) do(ctx context.Context, r *LogView, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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

	if _, err := c.GetLogView(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getLogViewRaw(ctx context.Context, r *LogView) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) logViewDiffsForRawDesired(ctx context.Context, rawDesired *LogView, opts ...dcl.ApplyOption) (initial, desired *LogView, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *LogView
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*LogView); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected LogView, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetLogView(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a LogView resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve LogView resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that LogView resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeLogViewDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for LogView: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for LogView: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractLogViewFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeLogViewInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for LogView: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeLogViewDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for LogView: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffLogView(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeLogViewInitialState(rawInitial, rawDesired *LogView) (*LogView, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeLogViewDesiredState(rawDesired, rawInitial *LogView, opts ...dcl.ApplyOption) (*LogView, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &LogView{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.Filter, rawInitial.Filter) {
		canonicalDesired.Filter = rawInitial.Filter
	} else {
		canonicalDesired.Filter = rawDesired.Filter
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	if dcl.NameToSelfLink(rawDesired.Bucket, rawInitial.Bucket) {
		canonicalDesired.Bucket = rawInitial.Bucket
	} else {
		canonicalDesired.Bucket = rawDesired.Bucket
	}
	return canonicalDesired, nil
}

func canonicalizeLogViewNewState(c *Client, rawNew, rawDesired *LogView) (*LogView, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Filter) && dcl.IsEmptyValueIndirect(rawDesired.Filter) {
		rawNew.Filter = rawDesired.Filter
	} else {
		if dcl.StringCanonicalize(rawDesired.Filter, rawNew.Filter) {
			rawNew.Filter = rawDesired.Filter
		}
	}

	rawNew.Parent = rawDesired.Parent

	rawNew.Location = rawDesired.Location

	rawNew.Bucket = rawDesired.Bucket

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffLogView(c *Client, desired, actual *LogView, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateLogViewUpdateLogViewOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateLogViewUpdateLogViewOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *LogView) urlNormalized() *LogView {
	normalized := dcl.Copy(*r).(LogView)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Filter = dcl.SelfLinkToName(r.Filter)
	normalized.Parent = r.Parent
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Bucket = dcl.SelfLinkToName(r.Bucket)
	return &normalized
}

func (r *LogView) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateLogView" {
		fields := map[string]interface{}{
			"location": dcl.ValueOrEmptyString(nr.Location),
			"bucket":   dcl.ValueOrEmptyString(nr.Bucket),
			"parent":   dcl.ValueOrEmptyString(nr.Parent),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the LogView resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *LogView) marshal(c *Client) ([]byte, error) {
	m, err := expandLogView(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling LogView: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalLogView decodes JSON responses into the LogView resource schema.
func unmarshalLogView(b []byte, c *Client, res *LogView) (*LogView, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapLogView(m, c, res)
}

func unmarshalMapLogView(m map[string]interface{}, c *Client, res *LogView) (*LogView, error) {

	flattened := flattenLogView(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandLogView expands LogView into a JSON request object.
func expandLogView(c *Client, f *LogView) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("%s/locations/%s/buckets/%s/views/%s", f.Name, f.Parent, dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Bucket), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Filter; dcl.ValueShouldBeSent(v) {
		m["filter"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Bucket into bucket: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bucket"] = v
	}

	return m, nil
}

// flattenLogView flattens LogView from a JSON request object into the
// LogView type.
func flattenLogView(c *Client, i interface{}, res *LogView) *LogView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &LogView{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Filter = dcl.FlattenString(m["filter"])
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Bucket = dcl.FlattenString(m["bucket"])

	return resultRes
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *LogView) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalLogView(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Bucket == nil && ncr.Bucket == nil {
			c.Config.Logger.Info("Both Bucket fields null - considering equal.")
		} else if nr.Bucket == nil || ncr.Bucket == nil {
			c.Config.Logger.Info("Only one Bucket field is null - considering unequal.")
			return false
		} else if *nr.Bucket != *ncr.Bucket {
			return false
		}
		if nr.Parent == nil && ncr.Parent == nil {
			c.Config.Logger.Info("Both Parent fields null - considering equal.")
		} else if nr.Parent == nil || ncr.Parent == nil {
			c.Config.Logger.Info("Only one Parent field is null - considering unequal.")
			return false
		} else if *nr.Parent != *ncr.Parent {
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

type logViewDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         logViewApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToLogViewDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]logViewDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []logViewDiff
	// For each operation name, create a logViewDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := logViewDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToLogViewApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToLogViewApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (logViewApiOperation, error) {
	switch opName {

	case "updateLogViewUpdateLogViewOperation":
		return &updateLogViewUpdateLogViewOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractLogViewFields(r *LogView) error {
	vParent, err := dcl.ValueFromRegexOnField("Parent", r.Parent, r.Bucket, "((projects|folders|organizations|billingAccounts)/[a-z0-9A-Z-]*)/locations/.*")
	if err != nil {
		return err
	}
	r.Parent = vParent
	vLocation, err := dcl.ValueFromRegexOnField("Location", r.Location, r.Bucket, "[a-zA-Z]*/[a-z0-9A-Z-]*/locations/([a-z0-9-]*)/buckets/.*")
	if err != nil {
		return err
	}
	r.Location = vLocation
	return nil
}

func postReadExtractLogViewFields(r *LogView) error {
	return nil
}
