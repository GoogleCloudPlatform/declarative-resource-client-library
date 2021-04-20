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
package beta

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

func (r *User) validate() error {

	if !dcl.IsEmptyValueIndirect(r.SqlserverUserDetails) {
		if err := r.SqlserverUserDetails.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UserSqlserverUserDetails) validate() error {
	return nil
}

func userGetURL(userBasePath string, r *User) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"instance": dcl.ValueOrEmptyString(r.Instance),
		"name":     dcl.ValueOrEmptyString(r.Name),
		"host":     dcl.ValueOrEmptyString(r.Host),
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/users?fetchId=%s%s", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

func userListURL(userBasePath, project, instance string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"instance": instance,
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/users", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func userCreateURL(userBasePath, project, instance string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"instance": instance,
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/users", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func userDeleteURL(userBasePath string, r *User) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"host":     dcl.ValueOrEmptyString(r.Host),
		"instance": dcl.ValueOrEmptyString(r.Instance),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/users?host={{host}}&name={{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

// userApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type userApiOperation interface {
	do(context.Context, *User, *Client) error
}

// newUpdateUserUpdateRequest creates a request for an
// User resource's Update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateUserUpdateRequest(ctx context.Context, f *User, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		req["password"] = v
	}
	if v, err := expandUserSqlserverUserDetails(c, f.SqlserverUserDetails); err != nil {
		return nil, fmt.Errorf("error expanding SqlserverUserDetails into sqlserverUserDetails: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["sqlserverUserDetails"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	return req, nil
}

// marshalUpdateUserUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateUserUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateUserUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateUserUpdateOperation) do(ctx context.Context, r *User, c *Client) error {
	_, err := c.GetUser(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Update")
	if err != nil {
		return err
	}

	req, err := newUpdateUserUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateUserUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listUserRaw(ctx context.Context, project, instance, pageToken string, pageSize int32) ([]byte, error) {
	u, err := userListURL(c.Config.BasePath, project, instance)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != UserMaxPage {
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

type listUserOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listUser(ctx context.Context, project, instance, pageToken string, pageSize int32) ([]*User, string, error) {
	b, err := c.listUserRaw(ctx, project, instance, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listUserOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*User
	for _, v := range m.Items {
		res := flattenUser(c, v)
		res.Project = &project
		res.Instance = &instance
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllUser(ctx context.Context, f func(*User) bool, resources []*User) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteUser(ctx, res)
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

type deleteUserOperation struct{}

func (op *deleteUserOperation) do(ctx context.Context, r *User, c *Client) error {

	_, err := c.GetUser(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("User not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetUser checking for existence. error: %v", err)
		return err
	}

	u, err := userDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetUser(ctx, r.urlNormalized())
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
type createUserOperation struct {
	response map[string]interface{}
}

func (op *createUserOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createUserOperation) do(ctx context.Context, r *User, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, instance := r.createFields()
	u, err := userCreateURL(c.Config.BasePath, project, instance)

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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetUser(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getUserRaw(ctx context.Context, r *User) ([]byte, error) {

	u, err := userGetURL(c.Config.BasePath, r.urlNormalized())
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

	b, err = dcl.ExtractElementFromList(b, "items", r.matcher(c))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) userDiffsForRawDesired(ctx context.Context, rawDesired *User, opts ...dcl.ApplyOption) (initial, desired *User, diffs []userDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *User
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*User); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected User, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetUser(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a User resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve User resource: %v", err)
		}
		c.Config.Logger.Info("Found that User resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeUserDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for User: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for User: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeUserInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for User: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeUserDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for User: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffUser(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeUserInitialState(rawInitial, rawDesired *User) (*User, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeUserDesiredState(rawDesired, rawInitial *User, opts ...dcl.ApplyOption) (*User, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SqlserverUserDetails = canonicalizeUserSqlserverUserDetails(rawDesired.SqlserverUserDetails, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Password, rawInitial.Password) {
		rawDesired.Password = rawInitial.Password
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Instance, rawInitial.Instance) {
		rawDesired.Instance = rawInitial.Instance
	}
	rawDesired.SqlserverUserDetails = canonicalizeUserSqlserverUserDetails(rawDesired.SqlserverUserDetails, rawInitial.SqlserverUserDetails, opts...)
	if dcl.IsZeroValue(rawDesired.Type) {
		rawDesired.Type = rawInitial.Type
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.StringCanonicalize(rawDesired.Host, rawInitial.Host) {
		rawDesired.Host = rawInitial.Host
	}

	return rawDesired, nil
}

func canonicalizeUserNewState(c *Client, rawNew, rawDesired *User) (*User, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.Password = rawDesired.Password

	rawNew.Project = rawDesired.Project

	rawNew.Instance = rawDesired.Instance

	if dcl.IsEmptyValueIndirect(rawNew.SqlserverUserDetails) && dcl.IsEmptyValueIndirect(rawDesired.SqlserverUserDetails) {
		rawNew.SqlserverUserDetails = rawDesired.SqlserverUserDetails
	} else {
		rawNew.SqlserverUserDetails = canonicalizeNewUserSqlserverUserDetails(c, rawDesired.SqlserverUserDetails, rawNew.SqlserverUserDetails)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Host) && dcl.IsEmptyValueIndirect(rawDesired.Host) {
		rawNew.Host = rawDesired.Host
	} else {
		if dcl.StringCanonicalize(rawDesired.Host, rawNew.Host) {
			rawNew.Host = rawDesired.Host
		}
	}

	return rawNew, nil
}

func canonicalizeUserSqlserverUserDetails(des, initial *UserSqlserverUserDetails, opts ...dcl.ApplyOption) *UserSqlserverUserDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		des.Disabled = initial.Disabled
	}
	if dcl.IsZeroValue(des.ServerRoles) {
		des.ServerRoles = initial.ServerRoles
	}

	return des
}

func canonicalizeNewUserSqlserverUserDetails(c *Client, des, nw *UserSqlserverUserDetails) *UserSqlserverUserDetails {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewUserSqlserverUserDetailsSet(c *Client, des, nw []UserSqlserverUserDetails) []UserSqlserverUserDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []UserSqlserverUserDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUserSqlserverUserDetails(c, &d, &n) {
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

func canonicalizeNewUserSqlserverUserDetailsSlice(c *Client, des, nw []UserSqlserverUserDetails) []UserSqlserverUserDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []UserSqlserverUserDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUserSqlserverUserDetails(c, &d, &n))
	}

	return items
}

type userDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         userApiOperation
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
func diffUser(c *Client, desired, actual *User, opts ...dcl.ApplyOption) ([]userDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []userDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Password, actual.Password, dcl.Info{Ignore: true, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "password"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{
			UpdateOp: &updateUserUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "project"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "instance"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "EnumType", FieldName: "type"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{
			UpdateOp: &updateUserUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "etag"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "host"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, userDiff{RequiresRecreate: true, Diffs: ds})
	}

	if compareUserSqlserverUserDetails(c, desired.SqlserverUserDetails, actual.SqlserverUserDetails) {
		c.Config.Logger.Infof("Detected diff in SqlserverUserDetails.\nDESIRED: %v\nACTUAL: %v", desired.SqlserverUserDetails, actual.SqlserverUserDetails)

		diffs = append(diffs, userDiff{
			UpdateOp:  &updateUserUpdateOperation{},
			FieldName: "SqlserverUserDetails",
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
	var deduped []userDiff
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
func compareUserSqlserverUserDetails(c *Client, desired, actual *UserSqlserverUserDetails) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.Disabled, actual.Disabled) && !dcl.IsZeroValue(desired.Disabled) {
		c.Config.Logger.Infof("Diff in Disabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Disabled), dcl.SprintResource(actual.Disabled))
		return true
	}
	if !dcl.StringSliceEquals(desired.ServerRoles, actual.ServerRoles) && !dcl.IsZeroValue(desired.ServerRoles) {
		c.Config.Logger.Infof("Diff in ServerRoles.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServerRoles), dcl.SprintResource(actual.ServerRoles))
		return true
	}
	return false
}

func compareUserSqlserverUserDetailsSlice(c *Client, desired, actual []UserSqlserverUserDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UserSqlserverUserDetails, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUserSqlserverUserDetails(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UserSqlserverUserDetails, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUserSqlserverUserDetailsMap(c *Client, desired, actual map[string]UserSqlserverUserDetails) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UserSqlserverUserDetails, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in UserSqlserverUserDetails, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareUserSqlserverUserDetails(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in UserSqlserverUserDetails, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareUserTypeEnumSlice(c *Client, desired, actual []UserTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UserTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUserTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UserTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUserTypeEnum(c *Client, desired, actual *UserTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *User) urlNormalized() *User {
	normalized := dcl.Copy(*r).(User)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Password = dcl.SelfLinkToName(r.Password)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Instance = dcl.SelfLinkToName(r.Instance)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Host = dcl.SelfLinkToName(r.Host)
	return &normalized
}

func (r *User) getFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name), dcl.ValueOrEmptyString(n.Host)
}

func (r *User) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance)
}

func (r *User) deleteFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Host), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name)
}

func (r *User) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "Update" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"host":     dcl.ValueOrEmptyString(n.Host),
			"instance": dcl.ValueOrEmptyString(n.Instance),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/instances/{{instance}}/users?host={{host}}&name={{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the User resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *User) marshal(c *Client) ([]byte, error) {
	m, err := expandUser(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling User: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalUser decodes JSON responses into the User resource schema.
func unmarshalUser(b []byte, c *Client) (*User, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapUser(m, c)
}

func unmarshalMapUser(m map[string]interface{}, c *Client) (*User, error) {

	return flattenUser(c, m), nil
}

// expandUser expands User into a JSON request object.
func expandUser(c *Client, f *User) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Instance into instance: %w", err)
	} else if v != nil {
		m["instance"] = v
	}
	if v, err := expandUserSqlserverUserDetails(c, f.SqlserverUserDetails); err != nil {
		return nil, fmt.Errorf("error expanding SqlserverUserDetails into sqlserverUserDetails: %w", err)
	} else if v != nil {
		m["sqlserverUserDetails"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}

	return m, nil
}

// flattenUser flattens User from a JSON request object into the
// User type.
func flattenUser(c *Client, i interface{}) *User {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &User{}
	r.Name = dcl.FlattenString(m["name"])
	r.Password = dcl.FlattenSecretValue(m["password"])
	r.Project = dcl.FlattenString(m["project"])
	r.Instance = dcl.FlattenString(m["instance"])
	r.SqlserverUserDetails = flattenUserSqlserverUserDetails(c, m["sqlserverUserDetails"])
	r.Type = flattenUserTypeEnum(m["type"])
	r.Etag = dcl.FlattenString(m["etag"])
	r.Host = dcl.FlattenString(m["host"])

	return r
}

// expandUserSqlserverUserDetailsMap expands the contents of UserSqlserverUserDetails into a JSON
// request object.
func expandUserSqlserverUserDetailsMap(c *Client, f map[string]UserSqlserverUserDetails) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUserSqlserverUserDetails(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUserSqlserverUserDetailsSlice expands the contents of UserSqlserverUserDetails into a JSON
// request object.
func expandUserSqlserverUserDetailsSlice(c *Client, f []UserSqlserverUserDetails) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUserSqlserverUserDetails(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUserSqlserverUserDetailsMap flattens the contents of UserSqlserverUserDetails from a JSON
// response object.
func flattenUserSqlserverUserDetailsMap(c *Client, i interface{}) map[string]UserSqlserverUserDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UserSqlserverUserDetails{}
	}

	if len(a) == 0 {
		return map[string]UserSqlserverUserDetails{}
	}

	items := make(map[string]UserSqlserverUserDetails)
	for k, item := range a {
		items[k] = *flattenUserSqlserverUserDetails(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUserSqlserverUserDetailsSlice flattens the contents of UserSqlserverUserDetails from a JSON
// response object.
func flattenUserSqlserverUserDetailsSlice(c *Client, i interface{}) []UserSqlserverUserDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []UserSqlserverUserDetails{}
	}

	if len(a) == 0 {
		return []UserSqlserverUserDetails{}
	}

	items := make([]UserSqlserverUserDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUserSqlserverUserDetails(c, item.(map[string]interface{})))
	}

	return items
}

// expandUserSqlserverUserDetails expands an instance of UserSqlserverUserDetails into a JSON
// request object.
func expandUserSqlserverUserDetails(c *Client, f *UserSqlserverUserDetails) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v := f.ServerRoles; !dcl.IsEmptyValueIndirect(v) {
		m["serverRoles"] = v
	}

	return m, nil
}

// flattenUserSqlserverUserDetails flattens an instance of UserSqlserverUserDetails from a JSON
// response object.
func flattenUserSqlserverUserDetails(c *Client, i interface{}) *UserSqlserverUserDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UserSqlserverUserDetails{}
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.ServerRoles = dcl.FlattenStringSlice(m["serverRoles"])

	return r
}

// flattenUserTypeEnumSlice flattens the contents of UserTypeEnum from a JSON
// response object.
func flattenUserTypeEnumSlice(c *Client, i interface{}) []UserTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UserTypeEnum{}
	}

	if len(a) == 0 {
		return []UserTypeEnum{}
	}

	items := make([]UserTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUserTypeEnum(item.(interface{})))
	}

	return items
}

// flattenUserTypeEnum asserts that an interface is a string, and returns a
// pointer to a *UserTypeEnum with the same value as that string.
func flattenUserTypeEnum(i interface{}) *UserTypeEnum {
	s, ok := i.(string)
	if !ok {
		return UserTypeEnumRef("")
	}

	return UserTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *User) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalUser(b, c)
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
		if nr.Instance == nil && ncr.Instance == nil {
			c.Config.Logger.Info("Both Instance fields null - considering equal.")
		} else if nr.Instance == nil || ncr.Instance == nil {
			c.Config.Logger.Info("Only one Instance field is null - considering unequal.")
			return false
		} else if *nr.Instance != *ncr.Instance {
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
		if nr.Host == nil && ncr.Host == nil {
			c.Config.Logger.Info("Both Host fields null - considering equal.")
		} else if nr.Host == nil || ncr.Host == nil {
			c.Config.Logger.Info("Only one Host field is null - considering unequal.")
			return false
		} else if *nr.Host != *ncr.Host {
			return false
		}
		return true
	}
}
