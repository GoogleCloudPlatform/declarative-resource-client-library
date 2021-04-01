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
package iam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Role) validate() error {

	if !dcl.IsEmptyValueIndirect(r.LocalizedValues) {
		if err := r.LocalizedValues.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RoleLocalizedValues) validate() error {
	return nil
}

func roleGetURL(userBasePath string, r *Role) (string, error) {
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(r.Parent),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("{{parent}}/roles/{{name}}", "https://iam.googleapis.com/v1/", userBasePath, params), nil
}

func roleListURL(userBasePath, parent string) (string, error) {
	params := map[string]interface{}{
		"parent": parent,
	}
	return dcl.URL("{{parent}}/roles", "https://iam.googleapis.com/v1/", userBasePath, params), nil

}

func roleCreateURL(userBasePath, parent string) (string, error) {
	params := map[string]interface{}{
		"parent": parent,
	}
	return dcl.URL("{{parent}}/roles", "https://iam.googleapis.com/v1/", userBasePath, params), nil

}

func roleDeleteURL(userBasePath string, r *Role) (string, error) {
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(r.Parent),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("{{parent}}/roles/{{name}}", "https://iam.googleapis.com/v1/", userBasePath, params), nil
}

// roleApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type roleApiOperation interface {
	do(context.Context, *Role, *Client) error
}

// newUpdateRoleUpdateRoleRequest creates a request for an
// Role resource's UpdateRole update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateRoleUpdateRoleRequest(ctx context.Context, f *Role, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	b, err := c.getRoleRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateRoleUpdateRoleRequest converts the update into
// the final JSON request body.
func marshalUpdateRoleUpdateRoleRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateRoleUpdateRoleOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateRoleUpdateRoleOperation) do(ctx context.Context, r *Role, c *Client) error {
	_, err := c.GetRole(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateRole")
	if err != nil {
		return err
	}

	req, err := newUpdateRoleUpdateRoleRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateRoleUpdateRoleRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listRoleRaw(ctx context.Context, parent, pageToken string, pageSize int32) ([]byte, error) {
	u, err := roleListURL(c.Config.BasePath, parent)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RoleMaxPage {
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

type listRoleOperation struct {
	Roles []map[string]interface{} `json:"roles"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listRole(ctx context.Context, parent, pageToken string, pageSize int32) ([]*Role, string, error) {
	b, err := c.listRoleRaw(ctx, parent, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRoleOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Role
	for _, v := range m.Roles {
		res := flattenRole(c, v)
		res.Parent = &parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRole(ctx context.Context, f func(*Role) bool, resources []*Role) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRole(ctx, res)
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

type deleteRoleOperation struct{}

func (op *deleteRoleOperation) do(ctx context.Context, r *Role, c *Client) error {

	_, err := c.GetRole(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Role not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetRole checking for existence. error: %v", err)
		return err
	}

	u, err := roleDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Role: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createRoleOperation struct {
	response map[string]interface{}
}

func (op *createRoleOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRoleOperation) do(ctx context.Context, r *Role, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	parent := r.createFields()
	u, err := roleCreateURL(c.Config.BasePath, parent)

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

	if _, err := c.GetRole(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRoleRaw(ctx context.Context, r *Role) ([]byte, error) {

	u, err := roleGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) roleDiffsForRawDesired(ctx context.Context, rawDesired *Role, opts ...dcl.ApplyOption) (initial, desired *Role, diffs []roleDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Role
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Role); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Role, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRole(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Role resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Role resource: %v", err)
		}
		c.Config.Logger.Info("Found that Role resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRoleDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Role: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Role: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRoleInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Role: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRoleDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Role: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRole(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeRoleInitialState(rawInitial, rawDesired *Role) (*Role, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRoleDesiredState(rawDesired, rawInitial *Role, opts ...dcl.ApplyOption) (*Role, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.LocalizedValues = canonicalizeRoleLocalizedValues(rawDesired.LocalizedValues, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Title, rawInitial.Title) {
		rawDesired.Title = rawInitial.Title
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.LocalizedValues = canonicalizeRoleLocalizedValues(rawDesired.LocalizedValues, rawInitial.LocalizedValues, opts...)
	if dcl.StringCanonicalize(rawDesired.LifecyclePhase, rawInitial.LifecyclePhase) {
		rawDesired.LifecyclePhase = rawInitial.LifecyclePhase
	}
	if dcl.StringCanonicalize(rawDesired.GroupName, rawInitial.GroupName) {
		rawDesired.GroupName = rawInitial.GroupName
	}
	if dcl.StringCanonicalize(rawDesired.GroupTitle, rawInitial.GroupTitle) {
		rawDesired.GroupTitle = rawInitial.GroupTitle
	}
	if dcl.IsZeroValue(rawDesired.IncludedPermissions) {
		rawDesired.IncludedPermissions = rawInitial.IncludedPermissions
	}
	if dcl.IsZeroValue(rawDesired.Stage) {
		rawDesired.Stage = rawInitial.Stage
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.BoolCanonicalize(rawDesired.Deleted, rawInitial.Deleted) {
		rawDesired.Deleted = rawInitial.Deleted
	}
	if dcl.IsZeroValue(rawDesired.IncludedRoles) {
		rawDesired.IncludedRoles = rawInitial.IncludedRoles
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		rawDesired.Parent = rawInitial.Parent
	}

	return rawDesired, nil
}

func canonicalizeRoleNewState(c *Client, rawNew, rawDesired *Role) (*Role, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Title) && dcl.IsEmptyValueIndirect(rawDesired.Title) {
		rawNew.Title = rawDesired.Title
	} else {
		if dcl.StringCanonicalize(rawDesired.Title, rawNew.Title) {
			rawNew.Title = rawDesired.Title
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocalizedValues) && dcl.IsEmptyValueIndirect(rawDesired.LocalizedValues) {
		rawNew.LocalizedValues = rawDesired.LocalizedValues
	} else {
		rawNew.LocalizedValues = canonicalizeNewRoleLocalizedValues(c, rawDesired.LocalizedValues, rawNew.LocalizedValues)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LifecyclePhase) && dcl.IsEmptyValueIndirect(rawDesired.LifecyclePhase) {
		rawNew.LifecyclePhase = rawDesired.LifecyclePhase
	} else {
		if dcl.StringCanonicalize(rawDesired.LifecyclePhase, rawNew.LifecyclePhase) {
			rawNew.LifecyclePhase = rawDesired.LifecyclePhase
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GroupName) && dcl.IsEmptyValueIndirect(rawDesired.GroupName) {
		rawNew.GroupName = rawDesired.GroupName
	} else {
		if dcl.StringCanonicalize(rawDesired.GroupName, rawNew.GroupName) {
			rawNew.GroupName = rawDesired.GroupName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GroupTitle) && dcl.IsEmptyValueIndirect(rawDesired.GroupTitle) {
		rawNew.GroupTitle = rawDesired.GroupTitle
	} else {
		if dcl.StringCanonicalize(rawDesired.GroupTitle, rawNew.GroupTitle) {
			rawNew.GroupTitle = rawDesired.GroupTitle
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IncludedPermissions) && dcl.IsEmptyValueIndirect(rawDesired.IncludedPermissions) {
		rawNew.IncludedPermissions = rawDesired.IncludedPermissions
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Stage) && dcl.IsEmptyValueIndirect(rawDesired.Stage) {
		rawNew.Stage = rawDesired.Stage
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deleted) && dcl.IsEmptyValueIndirect(rawDesired.Deleted) {
		rawNew.Deleted = rawDesired.Deleted
	} else {
		if dcl.BoolCanonicalize(rawDesired.Deleted, rawNew.Deleted) {
			rawNew.Deleted = rawDesired.Deleted
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IncludedRoles) && dcl.IsEmptyValueIndirect(rawDesired.IncludedRoles) {
		rawNew.IncludedRoles = rawDesired.IncludedRoles
	} else {
	}

	rawNew.Parent = rawDesired.Parent

	return rawNew, nil
}

func canonicalizeRoleLocalizedValues(des, initial *RoleLocalizedValues, opts ...dcl.ApplyOption) *RoleLocalizedValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.LocalizedTitle, initial.LocalizedTitle) || dcl.IsZeroValue(des.LocalizedTitle) {
		des.LocalizedTitle = initial.LocalizedTitle
	}
	if dcl.StringCanonicalize(des.LocalizedDescription, initial.LocalizedDescription) || dcl.IsZeroValue(des.LocalizedDescription) {
		des.LocalizedDescription = initial.LocalizedDescription
	}

	return des
}

func canonicalizeNewRoleLocalizedValues(c *Client, des, nw *RoleLocalizedValues) *RoleLocalizedValues {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.LocalizedTitle, nw.LocalizedTitle) {
		nw.LocalizedTitle = des.LocalizedTitle
	}
	if dcl.StringCanonicalize(des.LocalizedDescription, nw.LocalizedDescription) {
		nw.LocalizedDescription = des.LocalizedDescription
	}

	return nw
}

func canonicalizeNewRoleLocalizedValuesSet(c *Client, des, nw []RoleLocalizedValues) []RoleLocalizedValues {
	if des == nil {
		return nw
	}
	var reorderedNew []RoleLocalizedValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRoleLocalizedValues(c, &d, &n) {
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

func canonicalizeNewRoleLocalizedValuesSlice(c *Client, des, nw []RoleLocalizedValues) []RoleLocalizedValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RoleLocalizedValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRoleLocalizedValues(c, &d, &n))
	}

	return items
}

type roleDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         roleApiOperation
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
func diffRole(c *Client, desired, actual *Role, opts ...dcl.ApplyOption) ([]roleDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []roleDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Title) && !dcl.StringCanonicalize(desired.Title, actual.Title) {
		c.Config.Logger.Infof("Detected diff in Title.\nDESIRED: %v\nACTUAL: %v", desired.Title, actual.Title)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "Title",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if compareRoleLocalizedValues(c, desired.LocalizedValues, actual.LocalizedValues) {
		c.Config.Logger.Infof("Detected diff in LocalizedValues.\nDESIRED: %v\nACTUAL: %v", desired.LocalizedValues, actual.LocalizedValues)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "LocalizedValues",
		})
	}
	if !dcl.IsZeroValue(desired.LifecyclePhase) && !dcl.StringCanonicalize(desired.LifecyclePhase, actual.LifecyclePhase) {
		c.Config.Logger.Infof("Detected diff in LifecyclePhase.\nDESIRED: %v\nACTUAL: %v", desired.LifecyclePhase, actual.LifecyclePhase)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "LifecyclePhase",
		})
	}
	if !dcl.IsZeroValue(desired.GroupName) && !dcl.StringCanonicalize(desired.GroupName, actual.GroupName) {
		c.Config.Logger.Infof("Detected diff in GroupName.\nDESIRED: %v\nACTUAL: %v", desired.GroupName, actual.GroupName)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "GroupName",
		})
	}
	if !dcl.IsZeroValue(desired.GroupTitle) && !dcl.StringCanonicalize(desired.GroupTitle, actual.GroupTitle) {
		c.Config.Logger.Infof("Detected diff in GroupTitle.\nDESIRED: %v\nACTUAL: %v", desired.GroupTitle, actual.GroupTitle)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "GroupTitle",
		})
	}
	if !dcl.StringSliceEquals(desired.IncludedPermissions, actual.IncludedPermissions) {
		c.Config.Logger.Infof("Detected diff in IncludedPermissions.\nDESIRED: %v\nACTUAL: %v", desired.IncludedPermissions, actual.IncludedPermissions)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "IncludedPermissions",
		})
	}
	if !reflect.DeepEqual(desired.Stage, actual.Stage) {
		c.Config.Logger.Infof("Detected diff in Stage.\nDESIRED: %v\nACTUAL: %v", desired.Stage, actual.Stage)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "Stage",
		})
	}
	if !dcl.IsZeroValue(desired.Etag) && !dcl.StringCanonicalize(desired.Etag, actual.Etag) {
		c.Config.Logger.Infof("Detected diff in Etag.\nDESIRED: %v\nACTUAL: %v", desired.Etag, actual.Etag)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "Etag",
		})
	}
	if !dcl.IsZeroValue(desired.Deleted) && !dcl.BoolCanonicalize(desired.Deleted, actual.Deleted) {
		c.Config.Logger.Infof("Detected diff in Deleted.\nDESIRED: %v\nACTUAL: %v", desired.Deleted, actual.Deleted)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "Deleted",
		})
	}
	if !dcl.StringSliceEquals(desired.IncludedRoles, actual.IncludedRoles) {
		c.Config.Logger.Infof("Detected diff in IncludedRoles.\nDESIRED: %v\nACTUAL: %v", desired.IncludedRoles, actual.IncludedRoles)
		diffs = append(diffs, roleDiff{
			RequiresRecreate: true,
			FieldName:        "IncludedRoles",
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
	var deduped []roleDiff
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
func compareRoleLocalizedValues(c *Client, desired, actual *RoleLocalizedValues) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.LocalizedTitle == nil && desired.LocalizedTitle != nil && !dcl.IsEmptyValueIndirect(desired.LocalizedTitle) {
		c.Config.Logger.Infof("desired LocalizedTitle %s - but actually nil", dcl.SprintResource(desired.LocalizedTitle))
		return true
	}
	if !dcl.StringCanonicalize(desired.LocalizedTitle, actual.LocalizedTitle) && !dcl.IsZeroValue(desired.LocalizedTitle) {
		c.Config.Logger.Infof("Diff in LocalizedTitle. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalizedTitle), dcl.SprintResource(actual.LocalizedTitle))
		return true
	}
	if actual.LocalizedDescription == nil && desired.LocalizedDescription != nil && !dcl.IsEmptyValueIndirect(desired.LocalizedDescription) {
		c.Config.Logger.Infof("desired LocalizedDescription %s - but actually nil", dcl.SprintResource(desired.LocalizedDescription))
		return true
	}
	if !dcl.StringCanonicalize(desired.LocalizedDescription, actual.LocalizedDescription) && !dcl.IsZeroValue(desired.LocalizedDescription) {
		c.Config.Logger.Infof("Diff in LocalizedDescription. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalizedDescription), dcl.SprintResource(actual.LocalizedDescription))
		return true
	}
	return false
}

func compareRoleLocalizedValuesSlice(c *Client, desired, actual []RoleLocalizedValues) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RoleLocalizedValues, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRoleLocalizedValues(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RoleLocalizedValues, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRoleLocalizedValuesMap(c *Client, desired, actual map[string]RoleLocalizedValues) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RoleLocalizedValues, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RoleLocalizedValues, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRoleLocalizedValues(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RoleLocalizedValues, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRoleStageEnumSlice(c *Client, desired, actual []RoleStageEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RoleStageEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRoleStageEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RoleStageEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRoleStageEnum(c *Client, desired, actual *RoleStageEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Role) urlNormalized() *Role {
	normalized := deepcopy.Copy(*r).(Role)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Title = dcl.SelfLinkToName(r.Title)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LifecyclePhase = dcl.SelfLinkToName(r.LifecyclePhase)
	normalized.GroupName = dcl.SelfLinkToName(r.GroupName)
	normalized.GroupTitle = dcl.SelfLinkToName(r.GroupTitle)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Parent = r.Parent
	return &normalized
}

func (r *Role) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Parent), dcl.ValueOrEmptyString(n.Name)
}

func (r *Role) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Parent)
}

func (r *Role) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Parent), dcl.ValueOrEmptyString(n.Name)
}

func (r *Role) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateRole" {
		fields := map[string]interface{}{
			"parent": dcl.ValueOrEmptyString(n.Parent),
			"name":   dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("{{parent}}/roles/{{name}}", "https://iam.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Role resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Role) marshal(c *Client) ([]byte, error) {
	m, err := expandRole(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Role: %w", err)
	}
	m = EncodeRoleCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalRole decodes JSON responses into the Role resource schema.
func unmarshalRole(b []byte, c *Client) (*Role, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRole(m, c)
}

func unmarshalMapRole(m map[string]interface{}, c *Client) (*Role, error) {

	return flattenRole(c, m), nil
}

// expandRole expands Role into a JSON request object.
func expandRole(c *Client, f *Role) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("%s/roles/%s", f.Name, f.Parent, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		m["title"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandRoleLocalizedValues(c, f.LocalizedValues); err != nil {
		return nil, fmt.Errorf("error expanding LocalizedValues into localizedValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localizedValues"] = v
	}
	if v := f.LifecyclePhase; !dcl.IsEmptyValueIndirect(v) {
		m["lifecyclePhase"] = v
	}
	if v := f.GroupName; !dcl.IsEmptyValueIndirect(v) {
		m["groupName"] = v
	}
	if v := f.GroupTitle; !dcl.IsEmptyValueIndirect(v) {
		m["groupTitle"] = v
	}
	if v := f.IncludedPermissions; !dcl.IsEmptyValueIndirect(v) {
		m["includedPermissions"] = v
	}
	if v := f.Stage; !dcl.IsEmptyValueIndirect(v) {
		m["stage"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v := f.Deleted; !dcl.IsEmptyValueIndirect(v) {
		m["deleted"] = v
	}
	if v := f.IncludedRoles; !dcl.IsEmptyValueIndirect(v) {
		m["includedRoles"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}

	return m, nil
}

// flattenRole flattens Role from a JSON request object into the
// Role type.
func flattenRole(c *Client, i interface{}) *Role {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Role{}
	r.Name = dcl.FlattenString(m["name"])
	r.Title = dcl.FlattenString(m["title"])
	r.Description = dcl.FlattenString(m["description"])
	r.LocalizedValues = flattenRoleLocalizedValues(c, m["localizedValues"])
	r.LifecyclePhase = dcl.FlattenString(m["lifecyclePhase"])
	r.GroupName = dcl.FlattenString(m["groupName"])
	r.GroupTitle = dcl.FlattenString(m["groupTitle"])
	r.IncludedPermissions = dcl.FlattenStringSlice(m["includedPermissions"])
	r.Stage = flattenRoleStageEnum(m["stage"])
	r.Etag = dcl.FlattenString(m["etag"])
	r.Deleted = dcl.FlattenBool(m["deleted"])
	r.IncludedRoles = dcl.FlattenStringSlice(m["includedRoles"])
	r.Parent = dcl.FlattenString(m["parent"])

	return r
}

// expandRoleLocalizedValuesMap expands the contents of RoleLocalizedValues into a JSON
// request object.
func expandRoleLocalizedValuesMap(c *Client, f map[string]RoleLocalizedValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRoleLocalizedValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRoleLocalizedValuesSlice expands the contents of RoleLocalizedValues into a JSON
// request object.
func expandRoleLocalizedValuesSlice(c *Client, f []RoleLocalizedValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRoleLocalizedValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRoleLocalizedValuesMap flattens the contents of RoleLocalizedValues from a JSON
// response object.
func flattenRoleLocalizedValuesMap(c *Client, i interface{}) map[string]RoleLocalizedValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoleLocalizedValues{}
	}

	if len(a) == 0 {
		return map[string]RoleLocalizedValues{}
	}

	items := make(map[string]RoleLocalizedValues)
	for k, item := range a {
		items[k] = *flattenRoleLocalizedValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRoleLocalizedValuesSlice flattens the contents of RoleLocalizedValues from a JSON
// response object.
func flattenRoleLocalizedValuesSlice(c *Client, i interface{}) []RoleLocalizedValues {
	a, ok := i.([]interface{})
	if !ok {
		return []RoleLocalizedValues{}
	}

	if len(a) == 0 {
		return []RoleLocalizedValues{}
	}

	items := make([]RoleLocalizedValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoleLocalizedValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandRoleLocalizedValues expands an instance of RoleLocalizedValues into a JSON
// request object.
func expandRoleLocalizedValues(c *Client, f *RoleLocalizedValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LocalizedTitle; !dcl.IsEmptyValueIndirect(v) {
		m["localizedTitle"] = v
	}
	if v := f.LocalizedDescription; !dcl.IsEmptyValueIndirect(v) {
		m["localizedDescription"] = v
	}

	return m, nil
}

// flattenRoleLocalizedValues flattens an instance of RoleLocalizedValues from a JSON
// response object.
func flattenRoleLocalizedValues(c *Client, i interface{}) *RoleLocalizedValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RoleLocalizedValues{}
	r.LocalizedTitle = dcl.FlattenString(m["localizedTitle"])
	r.LocalizedDescription = dcl.FlattenString(m["localizedDescription"])

	return r
}

// flattenRoleStageEnumSlice flattens the contents of RoleStageEnum from a JSON
// response object.
func flattenRoleStageEnumSlice(c *Client, i interface{}) []RoleStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoleStageEnum{}
	}

	if len(a) == 0 {
		return []RoleStageEnum{}
	}

	items := make([]RoleStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoleStageEnum(item.(interface{})))
	}

	return items
}

// flattenRoleStageEnum asserts that an interface is a string, and returns a
// pointer to a *RoleStageEnum with the same value as that string.
func flattenRoleStageEnum(i interface{}) *RoleStageEnum {
	s, ok := i.(string)
	if !ok {
		return RoleStageEnumRef("")
	}

	return RoleStageEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Role) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRole(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

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
