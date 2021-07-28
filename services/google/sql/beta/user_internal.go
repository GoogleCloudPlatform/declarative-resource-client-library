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
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateUserUpdateOperation) do(ctx context.Context, r *User, c *Client) error {
	_, err := c.GetUser(ctx, r.URLNormalized())
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
		res, err := unmarshalMapUser(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
	r, err := c.GetUser(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("User not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetUser checking for existence. error: %v", err)
		return err
	}

	u, err := userDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetUser(ctx, r.URLNormalized())
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

	if _, err := c.GetUser(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getUserRaw(ctx context.Context, r *User) ([]byte, error) {

	u, err := userGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) userDiffsForRawDesired(ctx context.Context, rawDesired *User, opts ...dcl.ApplyOption) (initial, desired *User, diffs []*dcl.FieldDiff, err error) {
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
	rawInitial, err := c.GetUser(ctx, fetchState.URLNormalized())
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
	canonicalDesired := &User{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Password, rawInitial.Password) {
		canonicalDesired.Password = rawInitial.Password
	} else {
		canonicalDesired.Password = rawDesired.Password
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Instance, rawInitial.Instance) {
		canonicalDesired.Instance = rawInitial.Instance
	} else {
		canonicalDesired.Instance = rawDesired.Instance
	}
	canonicalDesired.SqlserverUserDetails = canonicalizeUserSqlserverUserDetails(rawDesired.SqlserverUserDetails, rawInitial.SqlserverUserDetails, opts...)
	if dcl.IsZeroValue(rawDesired.Type) {
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		canonicalDesired.Etag = rawInitial.Etag
	} else {
		canonicalDesired.Etag = rawDesired.Etag
	}
	if dcl.StringCanonicalize(rawDesired.Host, rawInitial.Host) {
		canonicalDesired.Host = rawInitial.Host
	} else {
		canonicalDesired.Host = rawDesired.Host
	}

	return canonicalDesired, nil
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

	cDes := &UserSqlserverUserDetails{}

	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		cDes.Disabled = initial.Disabled
	} else {
		cDes.Disabled = des.Disabled
	}
	if dcl.IsZeroValue(des.ServerRoles) {
		des.ServerRoles = initial.ServerRoles
	} else {
		cDes.ServerRoles = des.ServerRoles
	}

	return cDes
}

func canonicalizeNewUserSqlserverUserDetails(c *Client, des, nw *UserSqlserverUserDetails) *UserSqlserverUserDetails {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}
	if dcl.IsZeroValue(nw.ServerRoles) {
		nw.ServerRoles = des.ServerRoles
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
			if diffs, _ := compareUserSqlserverUserDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []UserSqlserverUserDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewUserSqlserverUserDetails(c, &d, &n))
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
func diffUser(c *Client, desired, actual *User, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Password, actual.Password, dcl.Info{Ignore: true, OperationSelector: dcl.TriggersOperation("updateUserUpdateOperation")}, fn.AddNest("Password")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Instance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SqlserverUserDetails, actual.SqlserverUserDetails, dcl.Info{ObjectFunction: compareUserSqlserverUserDetailsNewStyle, EmptyObject: EmptyUserSqlserverUserDetails, OperationSelector: dcl.TriggersOperation("updateUserUpdateOperation")}, fn.AddNest("SqlserverUserDetails")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateUserUpdateOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareUserSqlserverUserDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*UserSqlserverUserDetails)
	if !ok {
		desiredNotPointer, ok := d.(UserSqlserverUserDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UserSqlserverUserDetails or *UserSqlserverUserDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*UserSqlserverUserDetails)
	if !ok {
		actualNotPointer, ok := a.(UserSqlserverUserDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a UserSqlserverUserDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerRoles, actual.ServerRoles, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerRoles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *User) getFields() (string, string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name), dcl.ValueOrEmptyString(n.Host)
}

func (r *User) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance)
}

func (r *User) deleteFields() (string, string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Host), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name)
}

func (r *User) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
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

	res := &User{}
	res.Name = dcl.FlattenString(m["name"])
	res.Password = dcl.FlattenSecretValue(m["password"])
	res.Project = dcl.FlattenString(m["project"])
	res.Instance = dcl.FlattenString(m["instance"])
	res.SqlserverUserDetails = flattenUserSqlserverUserDetails(c, m["sqlserverUserDetails"])
	res.Type = flattenUserTypeEnum(m["type"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Host = dcl.FlattenString(m["host"])

	return res
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
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v := f.ServerRoles; v != nil {
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyUserSqlserverUserDetails
	}
	r.Disabled = dcl.FlattenBool(m["disabled"])
	r.ServerRoles = dcl.FlattenStringSlice(m["serverRoles"])

	return r
}

// flattenUserTypeEnumMap flattens the contents of UserTypeEnum from a JSON
// response object.
func flattenUserTypeEnumMap(c *Client, i interface{}) map[string]UserTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UserTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]UserTypeEnum{}
	}

	items := make(map[string]UserTypeEnum)
	for k, item := range a {
		items[k] = *flattenUserTypeEnum(item.(interface{}))
	}

	return items
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

type userDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         userApiOperation
}

func convertFieldDiffsToUserDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]userDiff, error) {
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
	var diffs []userDiff
	// For each operation name, create a userDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := userDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToUserApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToUserApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (userApiOperation, error) {
	switch opName {

	case "updateUserUpdateOperation":
		return &updateUserUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
