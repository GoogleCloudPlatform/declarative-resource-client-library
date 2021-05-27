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
package apigee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Environment) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Organization, "Organization"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Properties) {
		if err := r.Properties.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EnvironmentProperties) validate() error {
	return nil
}
func (r *EnvironmentPropertiesProperty) validate() error {
	return nil
}

func environmentGetURL(userBasePath string, r *Environment) (string, error) {
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(r.Organization),
		"name":         dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("organizations/{{organization}}/environments/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

func environmentListURL(userBasePath, organization string) (string, error) {
	params := map[string]interface{}{
		"organization": organization,
	}
	return dcl.URL("organizations/{{organization}}/environments", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func environmentCreateURL(userBasePath, organization string) (string, error) {
	params := map[string]interface{}{
		"organization": organization,
	}
	return dcl.URL("organizations/{{organization}}/environments", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func environmentDeleteURL(userBasePath string, r *Environment) (string, error) {
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(r.Organization),
		"name":         dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("organizations/{{organization}}/environments/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

func (r *Environment) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"organization": *n.Organization,
		"name":         *n.Name,
	}
	return dcl.URL("organizations/{{organization}}/environments/{{name}}:setIamPolicy", "https://apigee.googleapis.com/v1/", userBasePath, fields)
}

func (r *Environment) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"organization": *n.Organization,
		"name":         *n.Name,
	}
	return dcl.URL("organizations/{{organization}}/environments/{{name}}:getIamPolicy", "https://apigee.googleapis.com/v1/", userBasePath, fields)
}

func (r *Environment) IAMPolicyVersion() int {
	return 3
}

// environmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type environmentApiOperation interface {
	do(context.Context, *Environment, *Client) error
}

// newUpdateEnvironmentUpdateEnvironmentRequest creates a request for an
// Environment resource's UpdateEnvironment update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEnvironmentUpdateEnvironmentRequest(ctx context.Context, f *Environment, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateEnvironmentUpdateEnvironmentRequest converts the update into
// the final JSON request body.
func marshalUpdateEnvironmentUpdateEnvironmentRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEnvironmentUpdateEnvironmentOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEnvironmentUpdateEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	_, err := c.GetEnvironment(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEnvironment")
	if err != nil {
		return err
	}

	req, err := newUpdateEnvironmentUpdateEnvironmentRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateEnvironmentUpdateEnvironmentRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listEnvironmentRaw(ctx context.Context, organization, pageToken string, pageSize int32) ([]byte, error) {
	u, err := environmentListURL(c.Config.BasePath, organization)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EnvironmentMaxPage {
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

type listEnvironmentOperation struct {
	Environments []map[string]interface{} `json:"environments"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listEnvironment(ctx context.Context, organization, pageToken string, pageSize int32) ([]*Environment, string, error) {
	b, err := c.listEnvironmentRaw(ctx, organization, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEnvironmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Environment
	for _, v := range m.Environments {
		res := flattenEnvironment(c, v)
		res.Organization = &organization
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEnvironment(ctx context.Context, f func(*Environment) bool, resources []*Environment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEnvironment(ctx, res)
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

type deleteEnvironmentOperation struct{}

func (op *deleteEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {

	_, err := c.GetEnvironment(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Environment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetEnvironment checking for existence. error: %v", err)
		return err
	}

	u, err := environmentDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetEnvironment(ctx, r.urlNormalized())
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
type createEnvironmentOperation struct {
	response map[string]interface{}
}

func (op *createEnvironmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	organization := r.createFields()
	u, err := environmentCreateURL(c.Config.BasePath, organization)

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
	if err := o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetEnvironment(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEnvironmentRaw(ctx context.Context, r *Environment) ([]byte, error) {

	u, err := environmentGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) environmentDiffsForRawDesired(ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (initial, desired *Environment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Environment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Environment); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Environment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEnvironment(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Environment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Environment resource: %v", err)
		}
		c.Config.Logger.Info("Found that Environment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEnvironmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Environment: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Environment: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEnvironmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Environment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEnvironmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Environment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEnvironment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEnvironmentInitialState(rawInitial, rawDesired *Environment) (*Environment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEnvironmentDesiredState(rawDesired, rawInitial *Environment, opts ...dcl.ApplyOption) (*Environment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Properties = canonicalizeEnvironmentProperties(rawDesired.Properties, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.Properties = canonicalizeEnvironmentProperties(rawDesired.Properties, rawInitial.Properties, opts...)
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.NameToSelfLink(rawDesired.Organization, rawInitial.Organization) {
		rawDesired.Organization = rawInitial.Organization
	}

	return rawDesired, nil
}

func canonicalizeEnvironmentNewState(c *Client, rawNew, rawDesired *Environment) (*Environment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
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

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedAt) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedAt) {
		rawNew.LastModifiedAt = rawDesired.LastModifiedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Properties) && dcl.IsEmptyValueIndirect(rawDesired.Properties) {
		rawNew.Properties = rawDesired.Properties
	} else {
		rawNew.Properties = canonicalizeNewEnvironmentProperties(c, rawDesired.Properties, rawNew.Properties)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Organization = rawDesired.Organization

	return rawNew, nil
}

func canonicalizeEnvironmentProperties(des, initial *EnvironmentProperties, opts ...dcl.ApplyOption) *EnvironmentProperties {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Property) {
		des.Property = initial.Property
	}

	return des
}

func canonicalizeNewEnvironmentProperties(c *Client, des, nw *EnvironmentProperties) *EnvironmentProperties {
	if des == nil || nw == nil {
		return nw
	}

	nw.Property = canonicalizeNewEnvironmentPropertiesPropertySet(c, des.Property, nw.Property)

	return nw
}

func canonicalizeNewEnvironmentPropertiesSet(c *Client, des, nw []EnvironmentProperties) []EnvironmentProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentPropertiesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentPropertiesSlice(c *Client, des, nw []EnvironmentProperties) []EnvironmentProperties {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentProperties
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentProperties(c, &d, &n))
	}

	return items
}

func canonicalizeEnvironmentPropertiesProperty(des, initial *EnvironmentPropertiesProperty, opts ...dcl.ApplyOption) *EnvironmentPropertiesProperty {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewEnvironmentPropertiesProperty(c *Client, des, nw *EnvironmentPropertiesProperty) *EnvironmentPropertiesProperty {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewEnvironmentPropertiesPropertySet(c *Client, des, nw []EnvironmentPropertiesProperty) []EnvironmentPropertiesProperty {
	if des == nil {
		return nw
	}
	var reorderedNew []EnvironmentPropertiesProperty
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEnvironmentPropertiesPropertyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewEnvironmentPropertiesPropertySlice(c *Client, des, nw []EnvironmentPropertiesProperty) []EnvironmentPropertiesProperty {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EnvironmentPropertiesProperty
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEnvironmentPropertiesProperty(c, &d, &n))
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
func diffEnvironment(c *Client, desired, actual *Environment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedAt, actual.CreatedAt, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreatedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedAt, actual.LastModifiedAt, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.Info{ObjectFunction: compareEnvironmentPropertiesNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Organization, actual.Organization, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Organization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareEnvironmentPropertiesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentProperties)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentProperties or *EnvironmentProperties", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentProperties)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentProperties", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Property, actual.Property, dcl.Info{Type: "Set", ObjectFunction: compareEnvironmentPropertiesPropertyNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Property")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEnvironmentPropertiesPropertyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EnvironmentPropertiesProperty)
	if !ok {
		desiredNotPointer, ok := d.(EnvironmentPropertiesProperty)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentPropertiesProperty or *EnvironmentPropertiesProperty", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EnvironmentPropertiesProperty)
	if !ok {
		actualNotPointer, ok := a.(EnvironmentPropertiesProperty)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EnvironmentPropertiesProperty", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Environment) urlNormalized() *Environment {
	normalized := dcl.Copy(*r).(Environment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Organization = dcl.SelfLinkToName(r.Organization)
	return &normalized
}

func (r *Environment) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Organization), dcl.ValueOrEmptyString(n.Name)
}

func (r *Environment) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Organization)
}

func (r *Environment) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Organization), dcl.ValueOrEmptyString(n.Name)
}

func (r *Environment) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateEnvironment" {
		fields := map[string]interface{}{
			"organization": dcl.ValueOrEmptyString(n.Organization),
			"name":         dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("organizations/{{organization}}/environments/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Environment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Environment) marshal(c *Client) ([]byte, error) {
	m, err := expandEnvironment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Environment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEnvironment decodes JSON responses into the Environment resource schema.
func unmarshalEnvironment(b []byte, c *Client) (*Environment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEnvironment(m, c)
}

func unmarshalMapEnvironment(m map[string]interface{}, c *Client) (*Environment, error) {

	return flattenEnvironment(c, m), nil
}

// expandEnvironment expands Environment into a JSON request object.
func expandEnvironment(c *Client, f *Environment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.CreatedAt; !dcl.IsEmptyValueIndirect(v) {
		m["createdAt"] = v
	}
	if v := f.LastModifiedAt; !dcl.IsEmptyValueIndirect(v) {
		m["lastModifiedAt"] = v
	}
	if v, err := expandEnvironmentProperties(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if v != nil {
		m["properties"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Organization into organization: %w", err)
	} else if v != nil {
		m["organization"] = v
	}

	return m, nil
}

// flattenEnvironment flattens Environment from a JSON request object into the
// Environment type.
func flattenEnvironment(c *Client, i interface{}) *Environment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Environment{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	res.LastModifiedAt = dcl.FlattenInteger(m["lastModifiedAt"])
	res.Properties = flattenEnvironmentProperties(c, m["properties"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.State = flattenEnvironmentStateEnum(m["state"])
	res.Organization = dcl.FlattenString(m["organization"])

	return res
}

// expandEnvironmentPropertiesMap expands the contents of EnvironmentProperties into a JSON
// request object.
func expandEnvironmentPropertiesMap(c *Client, f map[string]EnvironmentProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentPropertiesSlice expands the contents of EnvironmentProperties into a JSON
// request object.
func expandEnvironmentPropertiesSlice(c *Client, f []EnvironmentProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentPropertiesMap flattens the contents of EnvironmentProperties from a JSON
// response object.
func flattenEnvironmentPropertiesMap(c *Client, i interface{}) map[string]EnvironmentProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentProperties{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentProperties{}
	}

	items := make(map[string]EnvironmentProperties)
	for k, item := range a {
		items[k] = *flattenEnvironmentProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentPropertiesSlice flattens the contents of EnvironmentProperties from a JSON
// response object.
func flattenEnvironmentPropertiesSlice(c *Client, i interface{}) []EnvironmentProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentProperties{}
	}

	if len(a) == 0 {
		return []EnvironmentProperties{}
	}

	items := make([]EnvironmentProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentProperties expands an instance of EnvironmentProperties into a JSON
// request object.
func expandEnvironmentProperties(c *Client, f *EnvironmentProperties) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandEnvironmentPropertiesPropertySlice(c, f.Property); err != nil {
		return nil, fmt.Errorf("error expanding Property into property: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["property"] = v
	}

	return m, nil
}

// flattenEnvironmentProperties flattens an instance of EnvironmentProperties from a JSON
// response object.
func flattenEnvironmentProperties(c *Client, i interface{}) *EnvironmentProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentProperties{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentProperties
	}
	r.Property = flattenEnvironmentPropertiesPropertySlice(c, m["property"])

	return r
}

// expandEnvironmentPropertiesPropertyMap expands the contents of EnvironmentPropertiesProperty into a JSON
// request object.
func expandEnvironmentPropertiesPropertyMap(c *Client, f map[string]EnvironmentPropertiesProperty) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEnvironmentPropertiesProperty(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEnvironmentPropertiesPropertySlice expands the contents of EnvironmentPropertiesProperty into a JSON
// request object.
func expandEnvironmentPropertiesPropertySlice(c *Client, f []EnvironmentPropertiesProperty) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEnvironmentPropertiesProperty(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEnvironmentPropertiesPropertyMap flattens the contents of EnvironmentPropertiesProperty from a JSON
// response object.
func flattenEnvironmentPropertiesPropertyMap(c *Client, i interface{}) map[string]EnvironmentPropertiesProperty {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentPropertiesProperty{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentPropertiesProperty{}
	}

	items := make(map[string]EnvironmentPropertiesProperty)
	for k, item := range a {
		items[k] = *flattenEnvironmentPropertiesProperty(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEnvironmentPropertiesPropertySlice flattens the contents of EnvironmentPropertiesProperty from a JSON
// response object.
func flattenEnvironmentPropertiesPropertySlice(c *Client, i interface{}) []EnvironmentPropertiesProperty {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentPropertiesProperty{}
	}

	if len(a) == 0 {
		return []EnvironmentPropertiesProperty{}
	}

	items := make([]EnvironmentPropertiesProperty, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentPropertiesProperty(c, item.(map[string]interface{})))
	}

	return items
}

// expandEnvironmentPropertiesProperty expands an instance of EnvironmentPropertiesProperty into a JSON
// request object.
func expandEnvironmentPropertiesProperty(c *Client, f *EnvironmentPropertiesProperty) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenEnvironmentPropertiesProperty flattens an instance of EnvironmentPropertiesProperty from a JSON
// response object.
func flattenEnvironmentPropertiesProperty(c *Client, i interface{}) *EnvironmentPropertiesProperty {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EnvironmentPropertiesProperty{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEnvironmentPropertiesProperty
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenEnvironmentStateEnumSlice flattens the contents of EnvironmentStateEnum from a JSON
// response object.
func flattenEnvironmentStateEnumSlice(c *Client, i interface{}) []EnvironmentStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentStateEnum{}
	}

	if len(a) == 0 {
		return []EnvironmentStateEnum{}
	}

	items := make([]EnvironmentStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentStateEnum(item.(interface{})))
	}

	return items
}

// flattenEnvironmentStateEnum asserts that an interface is a string, and returns a
// pointer to a *EnvironmentStateEnum with the same value as that string.
func flattenEnvironmentStateEnum(i interface{}) *EnvironmentStateEnum {
	s, ok := i.(string)
	if !ok {
		return EnvironmentStateEnumRef("")
	}

	return EnvironmentStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Environment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEnvironment(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Organization == nil && ncr.Organization == nil {
			c.Config.Logger.Info("Both Organization fields null - considering equal.")
		} else if nr.Organization == nil || ncr.Organization == nil {
			c.Config.Logger.Info("Only one Organization field is null - considering unequal.")
			return false
		} else if *nr.Organization != *ncr.Organization {
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

type environmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         environmentApiOperation
}

func convertFieldDiffToEnvironmentOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]environmentDiff, error) {
	var diffs []environmentDiff
	for _, op := range ops {
		diff := environmentDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToenvironmentApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToenvironmentApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (environmentApiOperation, error) {
	switch op {

	case "updateEnvironmentUpdateEnvironmentOperation":
		return &updateEnvironmentUpdateEnvironmentOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
