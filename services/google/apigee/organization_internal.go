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
	"reflect"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Organization) validate() error {

	if err := dcl.Required(r, "analyticsRegion"); err != nil {
		return err
	}
	if err := dcl.Required(r, "runtimeType"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Properties) {
		if err := r.Properties.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OrganizationProperties) validate() error {
	return nil
}
func (r *OrganizationPropertiesProperty) validate() error {
	return nil
}

func organizationGetURL(userBasePath string, r *Organization) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("organizations/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

func organizationListURL(userBasePath string) (string, error) {
	params := map[string]interface{}{}
	return dcl.URL("organizations", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func organizationCreateURL(userBasePath, parent string) (string, error) {
	params := map[string]interface{}{
		"parent": parent,
	}
	return dcl.URL("organizations?parent={{parent}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func organizationDeleteURL(userBasePath string, r *Organization) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("organizations/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

// organizationApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type organizationApiOperation interface {
	do(context.Context, *Organization, *Client) error
}

// newUpdateOrganizationUpdateOrganizationRequest creates a request for an
// Organization resource's UpdateOrganization update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateOrganizationUpdateOrganizationRequest(ctx context.Context, f *Organization, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandOrganizationProperties(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["properties"] = v
	}
	return req, nil
}

// marshalUpdateOrganizationUpdateOrganizationRequest converts the update into
// the final JSON request body.
func marshalUpdateOrganizationUpdateOrganizationRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateOrganizationUpdateOrganizationOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listOrganizationRaw(ctx context.Context, pageToken string, pageSize int32) ([]byte, error) {
	u, err := organizationListURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != OrganizationMaxPage {
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

type listOrganizationOperation struct {
	Organizations []map[string]interface{} `json:"organizations"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listOrganization(ctx context.Context, pageToken string, pageSize int32) ([]*Organization, string, error) {
	b, err := c.listOrganizationRaw(ctx, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listOrganizationOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Organization
	for _, v := range m.Organizations {
		res := flattenOrganization(c, v)
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllOrganization(ctx context.Context, f func(*Organization) bool, resources []*Organization) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteOrganization(ctx, res)
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

type deleteOrganizationOperation struct{}

func (op *deleteOrganizationOperation) do(ctx context.Context, r *Organization, c *Client) error {

	_, err := c.GetOrganization(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.Infof("Organization not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetOrganization checking for existence. error: %v", err)
		return err
	}

	u, err := organizationDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetOrganization(ctx, r.urlNormalized())
		if !dcl.IsNotFoundOrCode(err, 403) {
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
type createOrganizationOperation struct {
	response map[string]interface{}
}

func (op *createOrganizationOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createOrganizationOperation) do(ctx context.Context, r *Organization, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	parent := r.createFields()
	u, err := organizationCreateURL(c.Config.BasePath, parent)

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

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string")
	}
	r.Name = &name

	if _, err := c.GetOrganization(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getOrganizationRaw(ctx context.Context, r *Organization) ([]byte, error) {

	u, err := organizationGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) organizationDiffsForRawDesired(ctx context.Context, rawDesired *Organization, opts ...dcl.ApplyOption) (initial, desired *Organization, diffs []organizationDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Organization
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Organization); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Organization, got %T", sh)
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
		desired, err := canonicalizeOrganizationDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetOrganization(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Organization resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Organization resource: %v", err)
		}
		c.Config.Logger.Info("Found that Organization resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeOrganizationDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Organization: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Organization: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeOrganizationInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Organization: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeOrganizationDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Organization: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffOrganization(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeOrganizationInitialState(rawInitial, rawDesired *Organization) (*Organization, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeOrganizationDesiredState(rawDesired, rawInitial *Organization, opts ...dcl.ApplyOption) (*Organization, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Properties = canonicalizeOrganizationProperties(rawDesired.Properties, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.CreatedAt) {
		rawDesired.CreatedAt = rawInitial.CreatedAt
	}
	if dcl.IsZeroValue(rawDesired.LastModifiedAt) {
		rawDesired.LastModifiedAt = rawInitial.LastModifiedAt
	}
	if dcl.IsZeroValue(rawDesired.ExpiresAt) {
		rawDesired.ExpiresAt = rawInitial.ExpiresAt
	}
	if dcl.IsZeroValue(rawDesired.Environments) {
		rawDesired.Environments = rawInitial.Environments
	}
	rawDesired.Properties = canonicalizeOrganizationProperties(rawDesired.Properties, rawInitial.Properties, opts...)
	if dcl.StringCanonicalize(rawDesired.AnalyticsRegion, rawInitial.AnalyticsRegion) {
		rawDesired.AnalyticsRegion = rawInitial.AnalyticsRegion
	}
	if dcl.StringCanonicalize(rawDesired.AuthorizedNetwork, rawInitial.AuthorizedNetwork) {
		rawDesired.AuthorizedNetwork = rawInitial.AuthorizedNetwork
	}
	if dcl.IsZeroValue(rawDesired.RuntimeType) {
		rawDesired.RuntimeType = rawInitial.RuntimeType
	}
	if dcl.IsZeroValue(rawDesired.SubscriptionType) {
		rawDesired.SubscriptionType = rawInitial.SubscriptionType
	}
	if dcl.IsZeroValue(rawDesired.BillingType) {
		rawDesired.BillingType = rawInitial.BillingType
	}
	if dcl.StringCanonicalize(rawDesired.CaCertificate, rawInitial.CaCertificate) {
		rawDesired.CaCertificate = rawInitial.CaCertificate
	}
	if dcl.StringCanonicalize(rawDesired.RuntimeDatabaseEncryptionKeyName, rawInitial.RuntimeDatabaseEncryptionKeyName) {
		rawDesired.RuntimeDatabaseEncryptionKeyName = rawInitial.RuntimeDatabaseEncryptionKeyName
	}
	if dcl.StringCanonicalize(rawDesired.ProjectId, rawInitial.ProjectId) {
		rawDesired.ProjectId = rawInitial.ProjectId
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		rawDesired.Parent = rawInitial.Parent
	}

	return rawDesired, nil
}

func canonicalizeOrganizationNewState(c *Client, rawNew, rawDesired *Organization) (*Organization, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.ExpiresAt) && dcl.IsEmptyValueIndirect(rawDesired.ExpiresAt) {
		rawNew.ExpiresAt = rawDesired.ExpiresAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Environments) && dcl.IsEmptyValueIndirect(rawDesired.Environments) {
		rawNew.Environments = rawDesired.Environments
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Properties) && dcl.IsEmptyValueIndirect(rawDesired.Properties) {
		rawNew.Properties = rawDesired.Properties
	} else {
		rawNew.Properties = canonicalizeNewOrganizationProperties(c, rawDesired.Properties, rawNew.Properties)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AnalyticsRegion) && dcl.IsEmptyValueIndirect(rawDesired.AnalyticsRegion) {
		rawNew.AnalyticsRegion = rawDesired.AnalyticsRegion
	} else {
		if dcl.StringCanonicalize(rawDesired.AnalyticsRegion, rawNew.AnalyticsRegion) {
			rawNew.AnalyticsRegion = rawDesired.AnalyticsRegion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthorizedNetwork) && dcl.IsEmptyValueIndirect(rawDesired.AuthorizedNetwork) {
		rawNew.AuthorizedNetwork = rawDesired.AuthorizedNetwork
	} else {
		if dcl.StringCanonicalize(rawDesired.AuthorizedNetwork, rawNew.AuthorizedNetwork) {
			rawNew.AuthorizedNetwork = rawDesired.AuthorizedNetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeType) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeType) {
		rawNew.RuntimeType = rawDesired.RuntimeType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubscriptionType) && dcl.IsEmptyValueIndirect(rawDesired.SubscriptionType) {
		rawNew.SubscriptionType = rawDesired.SubscriptionType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BillingType) && dcl.IsEmptyValueIndirect(rawDesired.BillingType) {
		rawNew.BillingType = rawDesired.BillingType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CaCertificate) && dcl.IsEmptyValueIndirect(rawDesired.CaCertificate) {
		rawNew.CaCertificate = rawDesired.CaCertificate
	} else {
		if dcl.StringCanonicalize(rawDesired.CaCertificate, rawNew.CaCertificate) {
			rawNew.CaCertificate = rawDesired.CaCertificate
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeDatabaseEncryptionKeyName) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeDatabaseEncryptionKeyName) {
		rawNew.RuntimeDatabaseEncryptionKeyName = rawDesired.RuntimeDatabaseEncryptionKeyName
	} else {
		if dcl.StringCanonicalize(rawDesired.RuntimeDatabaseEncryptionKeyName, rawNew.RuntimeDatabaseEncryptionKeyName) {
			rawNew.RuntimeDatabaseEncryptionKeyName = rawDesired.RuntimeDatabaseEncryptionKeyName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProjectId) && dcl.IsEmptyValueIndirect(rawDesired.ProjectId) {
		rawNew.ProjectId = rawDesired.ProjectId
	} else {
		if dcl.StringCanonicalize(rawDesired.ProjectId, rawNew.ProjectId) {
			rawNew.ProjectId = rawDesired.ProjectId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Parent = rawDesired.Parent

	return rawNew, nil
}

func canonicalizeOrganizationProperties(des, initial *OrganizationProperties, opts ...dcl.ApplyOption) *OrganizationProperties {
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

func canonicalizeNewOrganizationProperties(c *Client, des, nw *OrganizationProperties) *OrganizationProperties {
	if des == nil || nw == nil {
		return nw
	}

	nw.Property = canonicalizeNewOrganizationPropertiesPropertySet(c, des.Property, nw.Property)

	return nw
}

func canonicalizeNewOrganizationPropertiesSet(c *Client, des, nw []OrganizationProperties) []OrganizationProperties {
	if des == nil {
		return nw
	}
	var reorderedNew []OrganizationProperties
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareOrganizationProperties(c, &d, &n) {
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

func canonicalizeNewOrganizationPropertiesSlice(c *Client, des, nw []OrganizationProperties) []OrganizationProperties {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []OrganizationProperties
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOrganizationProperties(c, &d, &n))
	}

	return items
}

func canonicalizeOrganizationPropertiesProperty(des, initial *OrganizationPropertiesProperty, opts ...dcl.ApplyOption) *OrganizationPropertiesProperty {
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

func canonicalizeNewOrganizationPropertiesProperty(c *Client, des, nw *OrganizationPropertiesProperty) *OrganizationPropertiesProperty {
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

func canonicalizeNewOrganizationPropertiesPropertySet(c *Client, des, nw []OrganizationPropertiesProperty) []OrganizationPropertiesProperty {
	if des == nil {
		return nw
	}
	var reorderedNew []OrganizationPropertiesProperty
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareOrganizationPropertiesProperty(c, &d, &n) {
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

func canonicalizeNewOrganizationPropertiesPropertySlice(c *Client, des, nw []OrganizationPropertiesProperty) []OrganizationPropertiesProperty {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []OrganizationPropertiesProperty
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOrganizationPropertiesProperty(c, &d, &n))
	}

	return items
}

type organizationDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         organizationApiOperation
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
func diffOrganization(c *Client, desired, actual *Organization, opts ...dcl.ApplyOption) ([]organizationDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []organizationDiff

	var fn dcl.FieldName

	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType"}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Name",
		})
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{
			UpdateOp: &updateOrganizationUpdateOrganizationOperation{}, Diffs: ds,
			FieldName: "DisplayName",
		})
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{
			UpdateOp: &updateOrganizationUpdateOrganizationOperation{}, Diffs: ds,
			FieldName: "Description",
		})
	}

	if ds, err := dcl.Diff(desired.CreatedAt, actual.CreatedAt, dcl.Info{OutputOnly: true}, fn.AddNest("CreatedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CreatedAt",
		})
	}

	if ds, err := dcl.Diff(desired.LastModifiedAt, actual.LastModifiedAt, dcl.Info{OutputOnly: true}, fn.AddNest("LastModifiedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "LastModifiedAt",
		})
	}

	if ds, err := dcl.Diff(desired.ExpiresAt, actual.ExpiresAt, dcl.Info{OutputOnly: true}, fn.AddNest("ExpiresAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ExpiresAt",
		})
	}

	if ds, err := dcl.Diff(desired.Environments, actual.Environments, dcl.Info{OutputOnly: true}, fn.AddNest("Environments")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Environments",
		})
	}

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.Info{ObjectFunction: compareOrganizationPropertiesNewStyle}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{
			UpdateOp: &updateOrganizationUpdateOrganizationOperation{}, Diffs: ds,
			FieldName: "Properties",
		})
	}

	if ds, err := dcl.Diff(desired.AnalyticsRegion, actual.AnalyticsRegion, dcl.Info{}, fn.AddNest("AnalyticsRegion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "AnalyticsRegion",
		})
	}

	if ds, err := dcl.Diff(desired.AuthorizedNetwork, actual.AuthorizedNetwork, dcl.Info{}, fn.AddNest("AuthorizedNetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "AuthorizedNetwork",
		})
	}

	if ds, err := dcl.Diff(desired.RuntimeType, actual.RuntimeType, dcl.Info{Type: "EnumType"}, fn.AddNest("RuntimeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "RuntimeType",
		})
	}

	if ds, err := dcl.Diff(desired.SubscriptionType, actual.SubscriptionType, dcl.Info{OutputOnly: true, Type: "EnumType"}, fn.AddNest("SubscriptionType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "SubscriptionType",
		})
	}

	if ds, err := dcl.Diff(desired.BillingType, actual.BillingType, dcl.Info{OutputOnly: true, Type: "EnumType"}, fn.AddNest("BillingType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "BillingType",
		})
	}

	if ds, err := dcl.Diff(desired.CaCertificate, actual.CaCertificate, dcl.Info{OutputOnly: true}, fn.AddNest("CaCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "CaCertificate",
		})
	}

	if ds, err := dcl.Diff(desired.RuntimeDatabaseEncryptionKeyName, actual.RuntimeDatabaseEncryptionKeyName, dcl.Info{}, fn.AddNest("RuntimeDatabaseEncryptionKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "RuntimeDatabaseEncryptionKeyName",
		})
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.Info{OutputOnly: true}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "ProjectId",
		})
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType"}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "State",
		})
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.Info{}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, organizationDiff{RequiresRecreate: true, Diffs: ds,
			FieldName: "Parent",
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
	var deduped []organizationDiff
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
func compareOrganizationPropertiesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OrganizationProperties)
	if !ok {
		desiredNotPointer, ok := d.(OrganizationProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationProperties or *OrganizationProperties", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OrganizationProperties)
	if !ok {
		actualNotPointer, ok := a.(OrganizationProperties)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationProperties", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Property, actual.Property, dcl.Info{Type: "Set", ObjectFunction: compareOrganizationPropertiesPropertyNewStyle}, fn.AddNest("Property")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOrganizationProperties(c *Client, desired, actual *OrganizationProperties) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if toAdd, toRemove := compareOrganizationPropertiesPropertySets(c, desired.Property, actual.Property); len(toAdd)+len(toRemove) > 0 {
		c.Config.Logger.Infof("Diff in Property.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Property), dcl.SprintResource(actual.Property))
		return true
	}
	return false
}

func compareOrganizationPropertiesPropertySets(c *Client, desired, actual []OrganizationPropertiesProperty) (toAdd, toRemove []OrganizationPropertiesProperty) {
	if actual == nil {
		return desired, nil
	}
	toAdd = make([]OrganizationPropertiesProperty, 0)
	toRemove = make([]OrganizationPropertiesProperty, 0)

	for _, act := range actual {
		found := false
		for _, des := range desired {
			if !compareOrganizationPropertiesProperty(c, &des, &act) {
				found = true
				break
			}
		}
		if !found {
			toRemove = append(toRemove, act)
		}
	}

	for _, des := range desired {
		found := false
		for _, act := range actual {
			if !compareOrganizationPropertiesProperty(c, &des, &act) {
				found = true
				break
			}
		}
		if !found {
			toAdd = append(toAdd, des)
		}
	}

	return toAdd, toRemove
}

func compareOrganizationPropertiesSlice(c *Client, desired, actual []OrganizationProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationProperties, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOrganizationProperties(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OrganizationProperties, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOrganizationPropertiesMap(c *Client, desired, actual map[string]OrganizationProperties) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationProperties, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in OrganizationProperties, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareOrganizationProperties(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in OrganizationProperties, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareOrganizationPropertiesPropertyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OrganizationPropertiesProperty)
	if !ok {
		desiredNotPointer, ok := d.(OrganizationPropertiesProperty)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationPropertiesProperty or *OrganizationPropertiesProperty", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OrganizationPropertiesProperty)
	if !ok {
		actualNotPointer, ok := a.(OrganizationPropertiesProperty)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationPropertiesProperty", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOrganizationPropertiesProperty(c *Client, desired, actual *OrganizationPropertiesProperty) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareOrganizationPropertiesPropertySlice(c *Client, desired, actual []OrganizationPropertiesProperty) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationPropertiesProperty, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOrganizationPropertiesProperty(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OrganizationPropertiesProperty, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOrganizationPropertiesPropertyMap(c *Client, desired, actual map[string]OrganizationPropertiesProperty) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationPropertiesProperty, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in OrganizationPropertiesProperty, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareOrganizationPropertiesProperty(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in OrganizationPropertiesProperty, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareOrganizationRuntimeTypeEnumSlice(c *Client, desired, actual []OrganizationRuntimeTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationRuntimeTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOrganizationRuntimeTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OrganizationRuntimeTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOrganizationRuntimeTypeEnum(c *Client, desired, actual *OrganizationRuntimeTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareOrganizationSubscriptionTypeEnumSlice(c *Client, desired, actual []OrganizationSubscriptionTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationSubscriptionTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOrganizationSubscriptionTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OrganizationSubscriptionTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOrganizationSubscriptionTypeEnum(c *Client, desired, actual *OrganizationSubscriptionTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareOrganizationBillingTypeEnumSlice(c *Client, desired, actual []OrganizationBillingTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationBillingTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOrganizationBillingTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OrganizationBillingTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOrganizationBillingTypeEnum(c *Client, desired, actual *OrganizationBillingTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareOrganizationStateEnumSlice(c *Client, desired, actual []OrganizationStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in OrganizationStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareOrganizationStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in OrganizationStateEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareOrganizationStateEnum(c *Client, desired, actual *OrganizationStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Organization) urlNormalized() *Organization {
	normalized := dcl.Copy(*r).(Organization)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AnalyticsRegion = dcl.SelfLinkToName(r.AnalyticsRegion)
	normalized.AuthorizedNetwork = dcl.SelfLinkToName(r.AuthorizedNetwork)
	normalized.CaCertificate = dcl.SelfLinkToName(r.CaCertificate)
	normalized.RuntimeDatabaseEncryptionKeyName = dcl.SelfLinkToName(r.RuntimeDatabaseEncryptionKeyName)
	normalized.ProjectId = dcl.SelfLinkToName(r.ProjectId)
	normalized.Parent = r.Parent
	return &normalized
}

func (r *Organization) getFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *Organization) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Parent)
}

func (r *Organization) deleteFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *Organization) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateOrganization" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("organizations/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Organization resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Organization) marshal(c *Client) ([]byte, error) {
	m, err := expandOrganization(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Organization: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalOrganization decodes JSON responses into the Organization resource schema.
func unmarshalOrganization(b []byte, c *Client) (*Organization, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapOrganization(m, c)
}

func unmarshalMapOrganization(m map[string]interface{}, c *Client) (*Organization, error) {

	return flattenOrganization(c, m), nil
}

// expandOrganization expands Organization into a JSON request object.
func expandOrganization(c *Client, f *Organization) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("organizations/%s", f.Name, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
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
	if v := f.ExpiresAt; !dcl.IsEmptyValueIndirect(v) {
		m["expiresAt"] = v
	}
	if v := f.Environments; !dcl.IsEmptyValueIndirect(v) {
		m["environments"] = v
	}
	if v, err := expandOrganizationProperties(c, f.Properties); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if v != nil {
		m["properties"] = v
	}
	if v := f.AnalyticsRegion; !dcl.IsEmptyValueIndirect(v) {
		m["analyticsRegion"] = v
	}
	if v := f.AuthorizedNetwork; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedNetwork"] = v
	}
	if v := f.RuntimeType; !dcl.IsEmptyValueIndirect(v) {
		m["runtimeType"] = v
	}
	if v := f.SubscriptionType; !dcl.IsEmptyValueIndirect(v) {
		m["subscriptionType"] = v
	}
	if v := f.BillingType; !dcl.IsEmptyValueIndirect(v) {
		m["billingType"] = v
	}
	if v := f.CaCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["caCertificate"] = v
	}
	if v := f.RuntimeDatabaseEncryptionKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["runtimeDatabaseEncryptionKeyName"] = v
	}
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if v != nil {
		m["parent"] = v
	}

	return m, nil
}

// flattenOrganization flattens Organization from a JSON request object into the
// Organization type.
func flattenOrganization(c *Client, i interface{}) *Organization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Organization{}
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Description = dcl.FlattenString(m["description"])
	r.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	r.LastModifiedAt = dcl.FlattenInteger(m["lastModifiedAt"])
	r.ExpiresAt = dcl.FlattenInteger(m["expiresAt"])
	r.Environments = dcl.FlattenStringSlice(m["environments"])
	r.Properties = flattenOrganizationProperties(c, m["properties"])
	r.AnalyticsRegion = dcl.FlattenString(m["analyticsRegion"])
	r.AuthorizedNetwork = dcl.FlattenString(m["authorizedNetwork"])
	r.RuntimeType = flattenOrganizationRuntimeTypeEnum(m["runtimeType"])
	r.SubscriptionType = flattenOrganizationSubscriptionTypeEnum(m["subscriptionType"])
	r.BillingType = flattenOrganizationBillingTypeEnum(m["billingType"])
	r.CaCertificate = dcl.FlattenString(m["caCertificate"])
	r.RuntimeDatabaseEncryptionKeyName = dcl.FlattenString(m["runtimeDatabaseEncryptionKeyName"])
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.State = flattenOrganizationStateEnum(m["state"])
	r.Parent = dcl.FlattenString(m["parent"])

	return r
}

// expandOrganizationPropertiesMap expands the contents of OrganizationProperties into a JSON
// request object.
func expandOrganizationPropertiesMap(c *Client, f map[string]OrganizationProperties) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOrganizationProperties(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOrganizationPropertiesSlice expands the contents of OrganizationProperties into a JSON
// request object.
func expandOrganizationPropertiesSlice(c *Client, f []OrganizationProperties) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOrganizationProperties(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOrganizationPropertiesMap flattens the contents of OrganizationProperties from a JSON
// response object.
func flattenOrganizationPropertiesMap(c *Client, i interface{}) map[string]OrganizationProperties {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationProperties{}
	}

	if len(a) == 0 {
		return map[string]OrganizationProperties{}
	}

	items := make(map[string]OrganizationProperties)
	for k, item := range a {
		items[k] = *flattenOrganizationProperties(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOrganizationPropertiesSlice flattens the contents of OrganizationProperties from a JSON
// response object.
func flattenOrganizationPropertiesSlice(c *Client, i interface{}) []OrganizationProperties {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationProperties{}
	}

	if len(a) == 0 {
		return []OrganizationProperties{}
	}

	items := make([]OrganizationProperties, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationProperties(c, item.(map[string]interface{})))
	}

	return items
}

// expandOrganizationProperties expands an instance of OrganizationProperties into a JSON
// request object.
func expandOrganizationProperties(c *Client, f *OrganizationProperties) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandOrganizationPropertiesPropertySlice(c, f.Property); err != nil {
		return nil, fmt.Errorf("error expanding Property into property: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["property"] = v
	}

	return m, nil
}

// flattenOrganizationProperties flattens an instance of OrganizationProperties from a JSON
// response object.
func flattenOrganizationProperties(c *Client, i interface{}) *OrganizationProperties {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OrganizationProperties{}
	r.Property = flattenOrganizationPropertiesPropertySlice(c, m["property"])

	return r
}

// expandOrganizationPropertiesPropertyMap expands the contents of OrganizationPropertiesProperty into a JSON
// request object.
func expandOrganizationPropertiesPropertyMap(c *Client, f map[string]OrganizationPropertiesProperty) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOrganizationPropertiesProperty(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOrganizationPropertiesPropertySlice expands the contents of OrganizationPropertiesProperty into a JSON
// request object.
func expandOrganizationPropertiesPropertySlice(c *Client, f []OrganizationPropertiesProperty) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOrganizationPropertiesProperty(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOrganizationPropertiesPropertyMap flattens the contents of OrganizationPropertiesProperty from a JSON
// response object.
func flattenOrganizationPropertiesPropertyMap(c *Client, i interface{}) map[string]OrganizationPropertiesProperty {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationPropertiesProperty{}
	}

	if len(a) == 0 {
		return map[string]OrganizationPropertiesProperty{}
	}

	items := make(map[string]OrganizationPropertiesProperty)
	for k, item := range a {
		items[k] = *flattenOrganizationPropertiesProperty(c, item.(map[string]interface{}))
	}

	return items
}

// flattenOrganizationPropertiesPropertySlice flattens the contents of OrganizationPropertiesProperty from a JSON
// response object.
func flattenOrganizationPropertiesPropertySlice(c *Client, i interface{}) []OrganizationPropertiesProperty {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationPropertiesProperty{}
	}

	if len(a) == 0 {
		return []OrganizationPropertiesProperty{}
	}

	items := make([]OrganizationPropertiesProperty, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationPropertiesProperty(c, item.(map[string]interface{})))
	}

	return items
}

// expandOrganizationPropertiesProperty expands an instance of OrganizationPropertiesProperty into a JSON
// request object.
func expandOrganizationPropertiesProperty(c *Client, f *OrganizationPropertiesProperty) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenOrganizationPropertiesProperty flattens an instance of OrganizationPropertiesProperty from a JSON
// response object.
func flattenOrganizationPropertiesProperty(c *Client, i interface{}) *OrganizationPropertiesProperty {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OrganizationPropertiesProperty{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenOrganizationRuntimeTypeEnumSlice flattens the contents of OrganizationRuntimeTypeEnum from a JSON
// response object.
func flattenOrganizationRuntimeTypeEnumSlice(c *Client, i interface{}) []OrganizationRuntimeTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationRuntimeTypeEnum{}
	}

	if len(a) == 0 {
		return []OrganizationRuntimeTypeEnum{}
	}

	items := make([]OrganizationRuntimeTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationRuntimeTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationRuntimeTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationRuntimeTypeEnum with the same value as that string.
func flattenOrganizationRuntimeTypeEnum(i interface{}) *OrganizationRuntimeTypeEnum {
	s, ok := i.(string)
	if !ok {
		return OrganizationRuntimeTypeEnumRef("")
	}

	return OrganizationRuntimeTypeEnumRef(s)
}

// flattenOrganizationSubscriptionTypeEnumSlice flattens the contents of OrganizationSubscriptionTypeEnum from a JSON
// response object.
func flattenOrganizationSubscriptionTypeEnumSlice(c *Client, i interface{}) []OrganizationSubscriptionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationSubscriptionTypeEnum{}
	}

	if len(a) == 0 {
		return []OrganizationSubscriptionTypeEnum{}
	}

	items := make([]OrganizationSubscriptionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationSubscriptionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationSubscriptionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationSubscriptionTypeEnum with the same value as that string.
func flattenOrganizationSubscriptionTypeEnum(i interface{}) *OrganizationSubscriptionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return OrganizationSubscriptionTypeEnumRef("")
	}

	return OrganizationSubscriptionTypeEnumRef(s)
}

// flattenOrganizationBillingTypeEnumSlice flattens the contents of OrganizationBillingTypeEnum from a JSON
// response object.
func flattenOrganizationBillingTypeEnumSlice(c *Client, i interface{}) []OrganizationBillingTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationBillingTypeEnum{}
	}

	if len(a) == 0 {
		return []OrganizationBillingTypeEnum{}
	}

	items := make([]OrganizationBillingTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationBillingTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationBillingTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationBillingTypeEnum with the same value as that string.
func flattenOrganizationBillingTypeEnum(i interface{}) *OrganizationBillingTypeEnum {
	s, ok := i.(string)
	if !ok {
		return OrganizationBillingTypeEnumRef("")
	}

	return OrganizationBillingTypeEnumRef(s)
}

// flattenOrganizationStateEnumSlice flattens the contents of OrganizationStateEnum from a JSON
// response object.
func flattenOrganizationStateEnumSlice(c *Client, i interface{}) []OrganizationStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationStateEnum{}
	}

	if len(a) == 0 {
		return []OrganizationStateEnum{}
	}

	items := make([]OrganizationStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationStateEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationStateEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationStateEnum with the same value as that string.
func flattenOrganizationStateEnum(i interface{}) *OrganizationStateEnum {
	s, ok := i.(string)
	if !ok {
		return OrganizationStateEnumRef("")
	}

	return OrganizationStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Organization) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalOrganization(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

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
