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
package identitytoolkit

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

func (r *Tenant) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MfaConfig) {
		if err := r.MfaConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TenantMfaConfig) validate() error {
	return nil
}

func tenantGetURL(userBasePath string, r *Tenant) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil
}

func tenantListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/tenants", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil

}

func tenantCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/tenants", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil

}

func tenantDeleteURL(userBasePath string, r *Tenant) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, params), nil
}

// tenantApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type tenantApiOperation interface {
	do(context.Context, *Tenant, *Client) error
}

// newUpdateTenantUpdateTenantRequest creates a request for an
// Tenant resource's UpdateTenant update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTenantUpdateTenantRequest(ctx context.Context, f *Tenant, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.AllowPasswordSignup; !dcl.IsEmptyValueIndirect(v) {
		req["allowPasswordSignup"] = v
	}
	if v := f.EnableEmailLinkSignin; !dcl.IsEmptyValueIndirect(v) {
		req["enableEmailLinkSignin"] = v
	}
	if v := f.DisableAuth; !dcl.IsEmptyValueIndirect(v) {
		req["disableAuth"] = v
	}
	if v := f.EnableAnonymousUser; !dcl.IsEmptyValueIndirect(v) {
		req["enableAnonymousUser"] = v
	}
	if v, err := expandTenantMfaConfig(c, f.MfaConfig); err != nil {
		return nil, fmt.Errorf("error expanding MfaConfig into mfaConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["mfaConfig"] = v
	}
	if v := f.TestPhoneNumbers; !dcl.IsEmptyValueIndirect(v) {
		req["testPhoneNumbers"] = v
	}
	return req, nil
}

// marshalUpdateTenantUpdateTenantRequest converts the update into
// the final JSON request body.
func marshalUpdateTenantUpdateTenantRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTenantUpdateTenantOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTenantUpdateTenantOperation) do(ctx context.Context, r *Tenant, c *Client) error {
	_, err := c.GetTenant(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTenant")
	if err != nil {
		return err
	}

	req, err := newUpdateTenantUpdateTenantRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTenantUpdateTenantRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTenantRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := tenantListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TenantMaxPage {
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

type listTenantOperation struct {
	Tenants []map[string]interface{} `json:"tenants"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listTenant(ctx context.Context, project, pageToken string, pageSize int32) ([]*Tenant, string, error) {
	b, err := c.listTenantRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTenantOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Tenant
	for _, v := range m.Tenants {
		res := flattenTenant(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTenant(ctx context.Context, f func(*Tenant) bool, resources []*Tenant) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTenant(ctx, res)
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

type deleteTenantOperation struct{}

func (op *deleteTenantOperation) do(ctx context.Context, r *Tenant, c *Client) error {

	_, err := c.GetTenant(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Tenant not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTenant checking for existence. error: %v", err)
		return err
	}

	u, err := tenantDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Tenant: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetTenant(ctx, r.urlNormalized())
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
type createTenantOperation struct {
	response map[string]interface{}
}

func (op *createTenantOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTenantOperation) do(ctx context.Context, r *Tenant, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := tenantCreateURL(c.Config.BasePath, project)

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
		return fmt.Errorf("expected name to be a string")
	}
	r.Name = &name

	if _, err := c.GetTenant(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTenantRaw(ctx context.Context, r *Tenant) ([]byte, error) {

	u, err := tenantGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) tenantDiffsForRawDesired(ctx context.Context, rawDesired *Tenant, opts ...dcl.ApplyOption) (initial, desired *Tenant, diffs []tenantDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Tenant
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Tenant); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Tenant, got %T", sh)
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
		desired, err := canonicalizeTenantDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTenant(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Tenant resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Tenant resource: %v", err)
		}
		c.Config.Logger.Info("Found that Tenant resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTenantDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Tenant: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Tenant: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTenantInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Tenant: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTenantDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Tenant: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTenant(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTenantInitialState(rawInitial, rawDesired *Tenant) (*Tenant, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTenantDesiredState(rawDesired, rawInitial *Tenant, opts ...dcl.ApplyOption) (*Tenant, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MfaConfig = canonicalizeTenantMfaConfig(rawDesired.MfaConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.BoolCanonicalize(rawDesired.AllowPasswordSignup, rawInitial.AllowPasswordSignup) {
		rawDesired.AllowPasswordSignup = rawInitial.AllowPasswordSignup
	}
	if dcl.BoolCanonicalize(rawDesired.EnableEmailLinkSignin, rawInitial.EnableEmailLinkSignin) {
		rawDesired.EnableEmailLinkSignin = rawInitial.EnableEmailLinkSignin
	}
	if dcl.BoolCanonicalize(rawDesired.DisableAuth, rawInitial.DisableAuth) {
		rawDesired.DisableAuth = rawInitial.DisableAuth
	}
	if dcl.BoolCanonicalize(rawDesired.EnableAnonymousUser, rawInitial.EnableAnonymousUser) {
		rawDesired.EnableAnonymousUser = rawInitial.EnableAnonymousUser
	}
	rawDesired.MfaConfig = canonicalizeTenantMfaConfig(rawDesired.MfaConfig, rawInitial.MfaConfig, opts...)
	if dcl.IsZeroValue(rawDesired.TestPhoneNumbers) {
		rawDesired.TestPhoneNumbers = rawInitial.TestPhoneNumbers
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeTenantNewState(c *Client, rawNew, rawDesired *Tenant) (*Tenant, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AllowPasswordSignup) && dcl.IsEmptyValueIndirect(rawDesired.AllowPasswordSignup) {
		rawNew.AllowPasswordSignup = rawDesired.AllowPasswordSignup
	} else {
		if dcl.BoolCanonicalize(rawDesired.AllowPasswordSignup, rawNew.AllowPasswordSignup) {
			rawNew.AllowPasswordSignup = rawDesired.AllowPasswordSignup
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableEmailLinkSignin) && dcl.IsEmptyValueIndirect(rawDesired.EnableEmailLinkSignin) {
		rawNew.EnableEmailLinkSignin = rawDesired.EnableEmailLinkSignin
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableEmailLinkSignin, rawNew.EnableEmailLinkSignin) {
			rawNew.EnableEmailLinkSignin = rawDesired.EnableEmailLinkSignin
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisableAuth) && dcl.IsEmptyValueIndirect(rawDesired.DisableAuth) {
		rawNew.DisableAuth = rawDesired.DisableAuth
	} else {
		if dcl.BoolCanonicalize(rawDesired.DisableAuth, rawNew.DisableAuth) {
			rawNew.DisableAuth = rawDesired.DisableAuth
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableAnonymousUser) && dcl.IsEmptyValueIndirect(rawDesired.EnableAnonymousUser) {
		rawNew.EnableAnonymousUser = rawDesired.EnableAnonymousUser
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableAnonymousUser, rawNew.EnableAnonymousUser) {
			rawNew.EnableAnonymousUser = rawDesired.EnableAnonymousUser
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MfaConfig) && dcl.IsEmptyValueIndirect(rawDesired.MfaConfig) {
		rawNew.MfaConfig = rawDesired.MfaConfig
	} else {
		rawNew.MfaConfig = canonicalizeNewTenantMfaConfig(c, rawDesired.MfaConfig, rawNew.MfaConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TestPhoneNumbers) && dcl.IsEmptyValueIndirect(rawDesired.TestPhoneNumbers) {
		rawNew.TestPhoneNumbers = rawDesired.TestPhoneNumbers
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeTenantMfaConfig(des, initial *TenantMfaConfig, opts ...dcl.ApplyOption) *TenantMfaConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.EnabledProviders) {
		des.EnabledProviders = initial.EnabledProviders
	}

	return des
}

func canonicalizeNewTenantMfaConfig(c *Client, des, nw *TenantMfaConfig) *TenantMfaConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewTenantMfaConfigSet(c *Client, des, nw []TenantMfaConfig) []TenantMfaConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []TenantMfaConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareTenantMfaConfig(c, &d, &n) {
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

func canonicalizeNewTenantMfaConfigSlice(c *Client, des, nw []TenantMfaConfig) []TenantMfaConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []TenantMfaConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTenantMfaConfig(c, &d, &n))
	}

	return items
}

type tenantDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         tenantApiOperation
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
func diffTenant(c *Client, desired, actual *Tenant, opts ...dcl.ApplyOption) ([]tenantDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []tenantDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "DisplayName",
		})
	}

	if ds, err := dcl.Diff(desired.AllowPasswordSignup, actual.AllowPasswordSignup, dcl.Info{}, fn.AddNest("AllowPasswordSignup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "AllowPasswordSignup",
		})
	}

	if ds, err := dcl.Diff(desired.EnableEmailLinkSignin, actual.EnableEmailLinkSignin, dcl.Info{}, fn.AddNest("EnableEmailLinkSignin")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "EnableEmailLinkSignin",
		})
	}

	if ds, err := dcl.Diff(desired.DisableAuth, actual.DisableAuth, dcl.Info{}, fn.AddNest("DisableAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "DisableAuth",
		})
	}

	if ds, err := dcl.Diff(desired.EnableAnonymousUser, actual.EnableAnonymousUser, dcl.Info{}, fn.AddNest("EnableAnonymousUser")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "EnableAnonymousUser",
		})
	}

	if ds, err := dcl.Diff(desired.MfaConfig, actual.MfaConfig, dcl.Info{ObjectFunction: compareTenantMfaConfigNewStyle}, fn.AddNest("MfaConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "MfaConfig",
		})
	}

	if ds, err := dcl.Diff(desired.TestPhoneNumbers, actual.TestPhoneNumbers, dcl.Info{}, fn.AddNest("TestPhoneNumbers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{
			UpdateOp: &updateTenantUpdateTenantOperation{}, Diffs: ds,
			FieldName: "TestPhoneNumbers",
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, tenantDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Project",
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
	var deduped []tenantDiff
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
func compareTenantMfaConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TenantMfaConfig)
	if !ok {
		desiredNotPointer, ok := d.(TenantMfaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TenantMfaConfig or *TenantMfaConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TenantMfaConfig)
	if !ok {
		actualNotPointer, ok := a.(TenantMfaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TenantMfaConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType"}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnabledProviders, actual.EnabledProviders, dcl.Info{Type: "EnumType"}, fn.AddNest("EnabledProviders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTenantMfaConfig(c *Client, desired, actual *TenantMfaConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.State, actual.State) && !dcl.IsZeroValue(desired.State) {
		c.Config.Logger.Infof("Diff in State.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.State), dcl.SprintResource(actual.State))
		return true
	}
	if compareTenantMfaConfigEnabledProvidersEnumSlice(c, desired.EnabledProviders, actual.EnabledProviders) && !dcl.IsZeroValue(desired.EnabledProviders) {
		c.Config.Logger.Infof("Diff in EnabledProviders.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnabledProviders), dcl.SprintResource(actual.EnabledProviders))
		return true
	}
	return false
}

func compareTenantMfaConfigSlice(c *Client, desired, actual []TenantMfaConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TenantMfaConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTenantMfaConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TenantMfaConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTenantMfaConfigMap(c *Client, desired, actual map[string]TenantMfaConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TenantMfaConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in TenantMfaConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareTenantMfaConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in TenantMfaConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareTenantMfaConfigStateEnumSlice(c *Client, desired, actual []TenantMfaConfigStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TenantMfaConfigStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTenantMfaConfigStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TenantMfaConfigStateEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTenantMfaConfigStateEnum(c *Client, desired, actual *TenantMfaConfigStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareTenantMfaConfigEnabledProvidersEnumSlice(c *Client, desired, actual []TenantMfaConfigEnabledProvidersEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in TenantMfaConfigEnabledProvidersEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareTenantMfaConfigEnabledProvidersEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in TenantMfaConfigEnabledProvidersEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareTenantMfaConfigEnabledProvidersEnum(c *Client, desired, actual *TenantMfaConfigEnabledProvidersEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Tenant) urlNormalized() *Tenant {
	normalized := dcl.Copy(*r).(Tenant)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Tenant) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Tenant) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Tenant) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Tenant) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateTenant" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/tenants/{{name}}", "https://identitytoolkit.googleapis.com/v2/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Tenant resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Tenant) marshal(c *Client) ([]byte, error) {
	m, err := expandTenant(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Tenant: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTenant decodes JSON responses into the Tenant resource schema.
func unmarshalTenant(b []byte, c *Client) (*Tenant, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTenant(m, c)
}

func unmarshalMapTenant(m map[string]interface{}, c *Client) (*Tenant, error) {

	return flattenTenant(c, m), nil
}

// expandTenant expands Tenant into a JSON request object.
func expandTenant(c *Client, f *Tenant) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/tenants/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.AllowPasswordSignup; !dcl.IsEmptyValueIndirect(v) {
		m["allowPasswordSignup"] = v
	}
	if v := f.EnableEmailLinkSignin; !dcl.IsEmptyValueIndirect(v) {
		m["enableEmailLinkSignin"] = v
	}
	if v := f.DisableAuth; !dcl.IsEmptyValueIndirect(v) {
		m["disableAuth"] = v
	}
	if v := f.EnableAnonymousUser; !dcl.IsEmptyValueIndirect(v) {
		m["enableAnonymousUser"] = v
	}
	if v, err := expandTenantMfaConfig(c, f.MfaConfig); err != nil {
		return nil, fmt.Errorf("error expanding MfaConfig into mfaConfig: %w", err)
	} else if v != nil {
		m["mfaConfig"] = v
	}
	if v := f.TestPhoneNumbers; !dcl.IsEmptyValueIndirect(v) {
		m["testPhoneNumbers"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenTenant flattens Tenant from a JSON request object into the
// Tenant type.
func flattenTenant(c *Client, i interface{}) *Tenant {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Tenant{}
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.AllowPasswordSignup = dcl.FlattenBool(m["allowPasswordSignup"])
	r.EnableEmailLinkSignin = dcl.FlattenBool(m["enableEmailLinkSignin"])
	r.DisableAuth = dcl.FlattenBool(m["disableAuth"])
	r.EnableAnonymousUser = dcl.FlattenBool(m["enableAnonymousUser"])
	r.MfaConfig = flattenTenantMfaConfig(c, m["mfaConfig"])
	r.TestPhoneNumbers = dcl.FlattenKeyValuePairs(m["testPhoneNumbers"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandTenantMfaConfigMap expands the contents of TenantMfaConfig into a JSON
// request object.
func expandTenantMfaConfigMap(c *Client, f map[string]TenantMfaConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTenantMfaConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTenantMfaConfigSlice expands the contents of TenantMfaConfig into a JSON
// request object.
func expandTenantMfaConfigSlice(c *Client, f []TenantMfaConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTenantMfaConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTenantMfaConfigMap flattens the contents of TenantMfaConfig from a JSON
// response object.
func flattenTenantMfaConfigMap(c *Client, i interface{}) map[string]TenantMfaConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TenantMfaConfig{}
	}

	if len(a) == 0 {
		return map[string]TenantMfaConfig{}
	}

	items := make(map[string]TenantMfaConfig)
	for k, item := range a {
		items[k] = *flattenTenantMfaConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenTenantMfaConfigSlice flattens the contents of TenantMfaConfig from a JSON
// response object.
func flattenTenantMfaConfigSlice(c *Client, i interface{}) []TenantMfaConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantMfaConfig{}
	}

	if len(a) == 0 {
		return []TenantMfaConfig{}
	}

	items := make([]TenantMfaConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantMfaConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandTenantMfaConfig expands an instance of TenantMfaConfig into a JSON
// request object.
func expandTenantMfaConfig(c *Client, f *TenantMfaConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.EnabledProviders; !dcl.IsEmptyValueIndirect(v) {
		m["enabledProviders"] = v
	}

	return m, nil
}

// flattenTenantMfaConfig flattens an instance of TenantMfaConfig from a JSON
// response object.
func flattenTenantMfaConfig(c *Client, i interface{}) *TenantMfaConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TenantMfaConfig{}
	r.State = flattenTenantMfaConfigStateEnum(m["state"])
	r.EnabledProviders = flattenTenantMfaConfigEnabledProvidersEnumSlice(c, m["enabledProviders"])

	return r
}

// flattenTenantMfaConfigStateEnumSlice flattens the contents of TenantMfaConfigStateEnum from a JSON
// response object.
func flattenTenantMfaConfigStateEnumSlice(c *Client, i interface{}) []TenantMfaConfigStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantMfaConfigStateEnum{}
	}

	if len(a) == 0 {
		return []TenantMfaConfigStateEnum{}
	}

	items := make([]TenantMfaConfigStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantMfaConfigStateEnum(item.(interface{})))
	}

	return items
}

// flattenTenantMfaConfigStateEnum asserts that an interface is a string, and returns a
// pointer to a *TenantMfaConfigStateEnum with the same value as that string.
func flattenTenantMfaConfigStateEnum(i interface{}) *TenantMfaConfigStateEnum {
	s, ok := i.(string)
	if !ok {
		return TenantMfaConfigStateEnumRef("")
	}

	return TenantMfaConfigStateEnumRef(s)
}

// flattenTenantMfaConfigEnabledProvidersEnumSlice flattens the contents of TenantMfaConfigEnabledProvidersEnum from a JSON
// response object.
func flattenTenantMfaConfigEnabledProvidersEnumSlice(c *Client, i interface{}) []TenantMfaConfigEnabledProvidersEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantMfaConfigEnabledProvidersEnum{}
	}

	if len(a) == 0 {
		return []TenantMfaConfigEnabledProvidersEnum{}
	}

	items := make([]TenantMfaConfigEnabledProvidersEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantMfaConfigEnabledProvidersEnum(item.(interface{})))
	}

	return items
}

// flattenTenantMfaConfigEnabledProvidersEnum asserts that an interface is a string, and returns a
// pointer to a *TenantMfaConfigEnabledProvidersEnum with the same value as that string.
func flattenTenantMfaConfigEnabledProvidersEnum(i interface{}) *TenantMfaConfigEnabledProvidersEnum {
	s, ok := i.(string)
	if !ok {
		return TenantMfaConfigEnabledProvidersEnumRef("")
	}

	return TenantMfaConfigEnabledProvidersEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Tenant) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTenant(b, c)
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
