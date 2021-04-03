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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *HttpFilter) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "filterName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "configTypeUrl"); err != nil {
		return err
	}
	if err := dcl.Required(r, "config"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}

func httpFilterGetURL(userBasePath string, r *HttpFilter) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpFilters/{{name}}", "https://networkservices.googleapis.com/v1alpha1/", userBasePath, params), nil
}

func httpFilterListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpFilters", "https://networkservices.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func httpFilterCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpFilters?httpFilterId={{name}}", "https://networkservices.googleapis.com/v1alpha1/", userBasePath, params), nil

}

func httpFilterDeleteURL(userBasePath string, r *HttpFilter) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpFilters/{{name}}", "https://networkservices.googleapis.com/v1alpha1/", userBasePath, params), nil
}

// httpFilterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type httpFilterApiOperation interface {
	do(context.Context, *HttpFilter, *Client) error
}

// newUpdateHttpFilterUpdateHttpFilterRequest creates a request for an
// HttpFilter resource's UpdateHttpFilter update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateHttpFilterUpdateHttpFilterRequest(ctx context.Context, f *HttpFilter, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/locations/global/httpFilters/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.FilterName; !dcl.IsEmptyValueIndirect(v) {
		req["filterName"] = v
	}
	if v := f.ConfigTypeUrl; !dcl.IsEmptyValueIndirect(v) {
		req["configTypeUrl"] = v
	}
	if v := f.Config; !dcl.IsEmptyValueIndirect(v) {
		req["config"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	return req, nil
}

// marshalUpdateHttpFilterUpdateHttpFilterRequest converts the update into
// the final JSON request body.
func marshalUpdateHttpFilterUpdateHttpFilterRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateHttpFilterUpdateHttpFilterOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateHttpFilterUpdateHttpFilterOperation) do(ctx context.Context, r *HttpFilter, c *Client) error {
	_, err := c.GetHttpFilter(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateHttpFilter")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"name", "labels", "filterName", "configTypeUrl", "config", "description"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateHttpFilterUpdateHttpFilterRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateHttpFilterUpdateHttpFilterRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://networkservices.googleapis.com/v1alpha1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listHttpFilterRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := httpFilterListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != HttpFilterMaxPage {
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

type listHttpFilterOperation struct {
	HttpFilters []map[string]interface{} `json:"httpFilters"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listHttpFilter(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*HttpFilter, string, error) {
	b, err := c.listHttpFilterRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listHttpFilterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*HttpFilter
	for _, v := range m.HttpFilters {
		res := flattenHttpFilter(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllHttpFilter(ctx context.Context, f func(*HttpFilter) bool, resources []*HttpFilter) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteHttpFilter(ctx, res)
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

type deleteHttpFilterOperation struct{}

func (op *deleteHttpFilterOperation) do(ctx context.Context, r *HttpFilter, c *Client) error {

	_, err := c.GetHttpFilter(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("HttpFilter not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetHttpFilter checking for existence. error: %v", err)
		return err
	}

	u, err := httpFilterDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://networkservices.googleapis.com/v1alpha1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetHttpFilter(ctx, r.urlNormalized())
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
type createHttpFilterOperation struct {
	response map[string]interface{}
}

func (op *createHttpFilterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createHttpFilterOperation) do(ctx context.Context, r *HttpFilter, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := httpFilterCreateURL(c.Config.BasePath, project, location, name)

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
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://networkservices.googleapis.com/v1alpha1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetHttpFilter(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getHttpFilterRaw(ctx context.Context, r *HttpFilter) ([]byte, error) {

	u, err := httpFilterGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) httpFilterDiffsForRawDesired(ctx context.Context, rawDesired *HttpFilter, opts ...dcl.ApplyOption) (initial, desired *HttpFilter, diffs []httpFilterDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *HttpFilter
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HttpFilter); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected HttpFilter, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetHttpFilter(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a HttpFilter resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve HttpFilter resource: %v", err)
		}
		c.Config.Logger.Info("Found that HttpFilter resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeHttpFilterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for HttpFilter: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for HttpFilter: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeHttpFilterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for HttpFilter: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeHttpFilterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for HttpFilter: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffHttpFilter(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeHttpFilterInitialState(rawInitial, rawDesired *HttpFilter) (*HttpFilter, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeHttpFilterDesiredState(rawDesired, rawInitial *HttpFilter, opts ...dcl.ApplyOption) (*HttpFilter, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.StringCanonicalize(rawDesired.FilterName, rawInitial.FilterName) {
		rawDesired.FilterName = rawInitial.FilterName
	}
	if dcl.StringCanonicalize(rawDesired.ConfigTypeUrl, rawInitial.ConfigTypeUrl) {
		rawDesired.ConfigTypeUrl = rawInitial.ConfigTypeUrl
	}
	if dcl.StringCanonicalize(rawDesired.Config, rawInitial.Config) {
		rawDesired.Config = rawInitial.Config
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeHttpFilterNewState(c *Client, rawNew, rawDesired *HttpFilter) (*HttpFilter, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
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

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FilterName) && dcl.IsEmptyValueIndirect(rawDesired.FilterName) {
		rawNew.FilterName = rawDesired.FilterName
	} else {
		if dcl.StringCanonicalize(rawDesired.FilterName, rawNew.FilterName) {
			rawNew.FilterName = rawDesired.FilterName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConfigTypeUrl) && dcl.IsEmptyValueIndirect(rawDesired.ConfigTypeUrl) {
		rawNew.ConfigTypeUrl = rawDesired.ConfigTypeUrl
	} else {
		if dcl.StringCanonicalize(rawDesired.ConfigTypeUrl, rawNew.ConfigTypeUrl) {
			rawNew.ConfigTypeUrl = rawDesired.ConfigTypeUrl
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		if dcl.StringCanonicalize(rawDesired.Config, rawNew.Config) {
			rawNew.Config = rawDesired.Config
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

type httpFilterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         httpFilterApiOperation
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
func diffHttpFilter(c *Client, desired, actual *HttpFilter, opts ...dcl.ApplyOption) ([]httpFilterDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []httpFilterDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{
			UpdateOp: &updateHttpFilterUpdateHttpFilterOperation{}, FieldName: "Name",
		})
	}

	if d, err := dcl.Diff(desired.CreateTime, actual.CreateTime, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{RequiresRecreate: true, FieldName: "CreateTime"})
	}

	if d, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{RequiresRecreate: true, FieldName: "UpdateTime"})
	}

	if d, err := dcl.Diff(desired.Labels, actual.Labels, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{
			UpdateOp: &updateHttpFilterUpdateHttpFilterOperation{}, FieldName: "Labels",
		})
	}

	if d, err := dcl.Diff(desired.FilterName, actual.FilterName, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{
			UpdateOp: &updateHttpFilterUpdateHttpFilterOperation{}, FieldName: "FilterName",
		})
	}

	if d, err := dcl.Diff(desired.ConfigTypeUrl, actual.ConfigTypeUrl, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{
			UpdateOp: &updateHttpFilterUpdateHttpFilterOperation{}, FieldName: "ConfigTypeUrl",
		})
	}

	if d, err := dcl.Diff(desired.Config, actual.Config, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{
			UpdateOp: &updateHttpFilterUpdateHttpFilterOperation{}, FieldName: "Config",
		})
	}

	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, httpFilterDiff{
			UpdateOp: &updateHttpFilterUpdateHttpFilterOperation{}, FieldName: "Description",
		})
	}

	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, httpFilterDiff{
			UpdateOp:  &updateHttpFilterUpdateHttpFilterOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, httpFilterDiff{
			UpdateOp:  &updateHttpFilterUpdateHttpFilterOperation{},
			FieldName: "Labels",
		})

	}
	if !dcl.IsZeroValue(desired.FilterName) && !dcl.StringCanonicalize(desired.FilterName, actual.FilterName) {
		c.Config.Logger.Infof("Detected diff in FilterName.\nDESIRED: %v\nACTUAL: %v", desired.FilterName, actual.FilterName)

		diffs = append(diffs, httpFilterDiff{
			UpdateOp:  &updateHttpFilterUpdateHttpFilterOperation{},
			FieldName: "FilterName",
		})

	}
	if !dcl.IsZeroValue(desired.ConfigTypeUrl) && !dcl.StringCanonicalize(desired.ConfigTypeUrl, actual.ConfigTypeUrl) {
		c.Config.Logger.Infof("Detected diff in ConfigTypeUrl.\nDESIRED: %v\nACTUAL: %v", desired.ConfigTypeUrl, actual.ConfigTypeUrl)

		diffs = append(diffs, httpFilterDiff{
			UpdateOp:  &updateHttpFilterUpdateHttpFilterOperation{},
			FieldName: "ConfigTypeUrl",
		})

	}
	if !dcl.IsZeroValue(desired.Config) && !dcl.StringCanonicalize(desired.Config, actual.Config) {
		c.Config.Logger.Infof("Detected diff in Config.\nDESIRED: %v\nACTUAL: %v", desired.Config, actual.Config)

		diffs = append(diffs, httpFilterDiff{
			UpdateOp:  &updateHttpFilterUpdateHttpFilterOperation{},
			FieldName: "Config",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, httpFilterDiff{
			UpdateOp:  &updateHttpFilterUpdateHttpFilterOperation{},
			FieldName: "Description",
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
	var deduped []httpFilterDiff
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *HttpFilter) urlNormalized() *HttpFilter {
	normalized := deepcopy.Copy(*r).(HttpFilter)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.FilterName = dcl.SelfLinkToName(r.FilterName)
	normalized.ConfigTypeUrl = dcl.SelfLinkToName(r.ConfigTypeUrl)
	normalized.Config = dcl.SelfLinkToName(r.Config)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *HttpFilter) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpFilter) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpFilter) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpFilter) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateHttpFilter" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/httpFilters/{{name}}", "https://networkservices.googleapis.com/v1alpha1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the HttpFilter resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *HttpFilter) marshal(c *Client) ([]byte, error) {
	m, err := expandHttpFilter(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling HttpFilter: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalHttpFilter decodes JSON responses into the HttpFilter resource schema.
func unmarshalHttpFilter(b []byte, c *Client) (*HttpFilter, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapHttpFilter(m, c)
}

func unmarshalMapHttpFilter(m map[string]interface{}, c *Client) (*HttpFilter, error) {

	return flattenHttpFilter(c, m), nil
}

// expandHttpFilter expands HttpFilter into a JSON request object.
func expandHttpFilter(c *Client, f *HttpFilter) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/global/httpFilters/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.FilterName; !dcl.IsEmptyValueIndirect(v) {
		m["filterName"] = v
	}
	if v := f.ConfigTypeUrl; !dcl.IsEmptyValueIndirect(v) {
		m["configTypeUrl"] = v
	}
	if v := f.Config; !dcl.IsEmptyValueIndirect(v) {
		m["config"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
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

// flattenHttpFilter flattens HttpFilter from a JSON request object into the
// HttpFilter type.
func flattenHttpFilter(c *Client, i interface{}) *HttpFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &HttpFilter{}
	r.Name = dcl.FlattenString(m["name"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.FilterName = dcl.FlattenString(m["filterName"])
	r.ConfigTypeUrl = dcl.FlattenString(m["configTypeUrl"])
	r.Config = dcl.FlattenString(m["config"])
	r.Description = dcl.FlattenString(m["description"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *HttpFilter) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalHttpFilter(b, c)
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
