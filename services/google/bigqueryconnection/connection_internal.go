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
package bigqueryconnection

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

func (r *Connection) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.CloudSql) {
		if err := r.CloudSql.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConnectionCloudSql) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Credential) {
		if err := r.Credential.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConnectionCloudSqlCredential) validate() error {
	return nil
}

func connectionGetURL(userBasePath string, r *Connection) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connections/{{name}}", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, params), nil
}

func connectionListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connections", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, params), nil

}

func connectionCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connections?connectionId={{name}}", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, params), nil

}

func connectionDeleteURL(userBasePath string, r *Connection) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connections/{{name}}", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, params), nil
}

func (r *Connection) SetPolicyURL(userBasePath string) string {
	n := r.URLNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connections/{{name}}:setIamPolicy", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, fields)
}

func (r *Connection) SetPolicyVerb() string {
	return "POST"
}

func (r *Connection) getPolicyURL(userBasePath string) string {
	n := r.URLNormalized()
	fields := map[string]interface{}{
		"project":  *n.Project,
		"location": *n.Location,
		"name":     *n.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connections/{{name}}:getIamPolicy", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, fields)
}

func (r *Connection) IAMPolicyVersion() int {
	return 3
}

// connectionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type connectionApiOperation interface {
	do(context.Context, *Connection, *Client) error
}

// newUpdateConnectionUpdateConnectionRequest creates a request for an
// Connection resource's UpdateConnection update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateConnectionUpdateConnectionRequest(ctx context.Context, f *Connection, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/locations/%s/connections/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.FriendlyName; !dcl.IsEmptyValueIndirect(v) {
		req["friendlyName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandConnectionCloudSql(c, f.CloudSql); err != nil {
		return nil, fmt.Errorf("error expanding CloudSql into cloudSql: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["cloudSql"] = v
	}
	return req, nil
}

// marshalUpdateConnectionUpdateConnectionRequest converts the update into
// the final JSON request body.
func marshalUpdateConnectionUpdateConnectionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateConnectionUpdateConnectionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateConnectionUpdateConnectionOperation) do(ctx context.Context, r *Connection, c *Client) error {
	_, err := c.GetConnection(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateConnection")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateConnectionUpdateConnectionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateConnectionUpdateConnectionRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listConnectionRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := connectionListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ConnectionMaxPage {
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

type listConnectionOperation struct {
	Connections []map[string]interface{} `json:"connections"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listConnection(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Connection, string, error) {
	b, err := c.listConnectionRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listConnectionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Connection
	for _, v := range m.Connections {
		res, err := unmarshalMapConnection(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllConnection(ctx context.Context, f func(*Connection) bool, resources []*Connection) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteConnection(ctx, res)
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

type deleteConnectionOperation struct{}

func (op *deleteConnectionOperation) do(ctx context.Context, r *Connection, c *Client) error {
	r, err := c.GetConnection(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Connection not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetConnection checking for existence. error: %v", err)
		return err
	}

	u, err := connectionDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Connection: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetConnection(ctx, r.URLNormalized())
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
type createConnectionOperation struct {
	response map[string]interface{}
}

func (op *createConnectionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createConnectionOperation) do(ctx context.Context, r *Connection, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := connectionCreateURL(c.Config.BasePath, project, location, name)

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

	if _, err := c.GetConnection(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getConnectionRaw(ctx context.Context, r *Connection) ([]byte, error) {

	u, err := connectionGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) connectionDiffsForRawDesired(ctx context.Context, rawDesired *Connection, opts ...dcl.ApplyOption) (initial, desired *Connection, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Connection
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Connection); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Connection, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetConnection(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Connection resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Connection resource: %v", err)
		}
		c.Config.Logger.Info("Found that Connection resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeConnectionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Connection: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Connection: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeConnectionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Connection: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeConnectionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Connection: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffConnection(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeConnectionInitialState(rawInitial, rawDesired *Connection) (*Connection, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeConnectionDesiredState(rawDesired, rawInitial *Connection, opts ...dcl.ApplyOption) (*Connection, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.CloudSql = canonicalizeConnectionCloudSql(rawDesired.CloudSql, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Connection{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.FriendlyName, rawInitial.FriendlyName) {
		canonicalDesired.FriendlyName = rawInitial.FriendlyName
	} else {
		canonicalDesired.FriendlyName = rawDesired.FriendlyName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.CloudSql = canonicalizeConnectionCloudSql(rawDesired.CloudSql, rawInitial.CloudSql, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeConnectionNewState(c *Client, rawNew, rawDesired *Connection) (*Connection, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.FriendlyName) && dcl.IsEmptyValueIndirect(rawDesired.FriendlyName) {
		rawNew.FriendlyName = rawDesired.FriendlyName
	} else {
		if dcl.StringCanonicalize(rawDesired.FriendlyName, rawNew.FriendlyName) {
			rawNew.FriendlyName = rawDesired.FriendlyName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CloudSql) && dcl.IsEmptyValueIndirect(rawDesired.CloudSql) {
		rawNew.CloudSql = rawDesired.CloudSql
	} else {
		rawNew.CloudSql = canonicalizeNewConnectionCloudSql(c, rawDesired.CloudSql, rawNew.CloudSql)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTime) && dcl.IsEmptyValueIndirect(rawDesired.CreationTime) {
		rawNew.CreationTime = rawDesired.CreationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedTime) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedTime) {
		rawNew.LastModifiedTime = rawDesired.LastModifiedTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.HasCredential) && dcl.IsEmptyValueIndirect(rawDesired.HasCredential) {
		rawNew.HasCredential = rawDesired.HasCredential
	} else {
		if dcl.BoolCanonicalize(rawDesired.HasCredential, rawNew.HasCredential) {
			rawNew.HasCredential = rawDesired.HasCredential
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeConnectionCloudSql(des, initial *ConnectionCloudSql, opts ...dcl.ApplyOption) *ConnectionCloudSql {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConnectionCloudSql{}

	if dcl.NameToSelfLink(des.InstanceId, initial.InstanceId) || dcl.IsZeroValue(des.InstanceId) {
		cDes.InstanceId = initial.InstanceId
	} else {
		cDes.InstanceId = des.InstanceId
	}
	if dcl.NameToSelfLink(des.Database, initial.Database) || dcl.IsZeroValue(des.Database) {
		cDes.Database = initial.Database
	} else {
		cDes.Database = des.Database
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	cDes.Credential = canonicalizeConnectionCloudSqlCredential(des.Credential, initial.Credential, opts...)

	return cDes
}

func canonicalizeNewConnectionCloudSql(c *Client, des, nw *ConnectionCloudSql) *ConnectionCloudSql {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.InstanceId, nw.InstanceId) {
		nw.InstanceId = des.InstanceId
	}
	if dcl.NameToSelfLink(des.Database, nw.Database) {
		nw.Database = des.Database
	}
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}
	nw.Credential = canonicalizeNewConnectionCloudSqlCredential(c, des.Credential, nw.Credential)

	return nw
}

func canonicalizeNewConnectionCloudSqlSet(c *Client, des, nw []ConnectionCloudSql) []ConnectionCloudSql {
	if des == nil {
		return nw
	}
	var reorderedNew []ConnectionCloudSql
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareConnectionCloudSqlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewConnectionCloudSqlSlice(c *Client, des, nw []ConnectionCloudSql) []ConnectionCloudSql {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConnectionCloudSql
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConnectionCloudSql(c, &d, &n))
	}

	return items
}

func canonicalizeConnectionCloudSqlCredential(des, initial *ConnectionCloudSqlCredential, opts ...dcl.ApplyOption) *ConnectionCloudSqlCredential {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConnectionCloudSqlCredential{}

	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		cDes.Username = initial.Username
	} else {
		cDes.Username = des.Username
	}
	if dcl.StringCanonicalize(des.Password, initial.Password) || dcl.IsZeroValue(des.Password) {
		cDes.Password = initial.Password
	} else {
		cDes.Password = des.Password
	}

	return cDes
}

func canonicalizeNewConnectionCloudSqlCredential(c *Client, des, nw *ConnectionCloudSqlCredential) *ConnectionCloudSqlCredential {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}
	if dcl.StringCanonicalize(des.Password, nw.Password) {
		nw.Password = des.Password
	}

	return nw
}

func canonicalizeNewConnectionCloudSqlCredentialSet(c *Client, des, nw []ConnectionCloudSqlCredential) []ConnectionCloudSqlCredential {
	if des == nil {
		return nw
	}
	var reorderedNew []ConnectionCloudSqlCredential
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareConnectionCloudSqlCredentialNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewConnectionCloudSqlCredentialSlice(c *Client, des, nw []ConnectionCloudSqlCredential) []ConnectionCloudSqlCredential {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConnectionCloudSqlCredential
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConnectionCloudSqlCredential(c, &d, &n))
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
func diffConnection(c *Client, desired, actual *Connection, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateConnectionUpdateConnectionOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FriendlyName, actual.FriendlyName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateConnectionUpdateConnectionOperation")}, fn.AddNest("FriendlyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateConnectionUpdateConnectionOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudSql, actual.CloudSql, dcl.Info{ObjectFunction: compareConnectionCloudSqlNewStyle, EmptyObject: EmptyConnectionCloudSql, OperationSelector: dcl.TriggersOperation("updateConnectionUpdateConnectionOperation")}, fn.AddNest("CloudSql")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTime, actual.CreationTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedTime, actual.LastModifiedTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HasCredential, actual.HasCredential, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HasCredential")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareConnectionCloudSqlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConnectionCloudSql)
	if !ok {
		desiredNotPointer, ok := d.(ConnectionCloudSql)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectionCloudSql or *ConnectionCloudSql", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConnectionCloudSql)
	if !ok {
		actualNotPointer, ok := a.(ConnectionCloudSql)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectionCloudSql", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstanceId, actual.InstanceId, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Database, actual.Database, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Database")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Credential, actual.Credential, dcl.Info{Ignore: true, ObjectFunction: compareConnectionCloudSqlCredentialNewStyle, EmptyObject: EmptyConnectionCloudSqlCredential, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Credential")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConnectionCloudSqlCredentialNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConnectionCloudSqlCredential)
	if !ok {
		desiredNotPointer, ok := d.(ConnectionCloudSqlCredential)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectionCloudSqlCredential or *ConnectionCloudSqlCredential", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConnectionCloudSqlCredential)
	if !ok {
		actualNotPointer, ok := a.(ConnectionCloudSqlCredential)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectionCloudSqlCredential", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Username, actual.Username, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Username")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Password, actual.Password, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Password")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Connection) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connection) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connection) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Connection) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateConnection" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/connections/{{name}}", "https://bigqueryconnection.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Connection resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Connection) marshal(c *Client) ([]byte, error) {
	m, err := expandConnection(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Connection: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalConnection decodes JSON responses into the Connection resource schema.
func unmarshalConnection(b []byte, c *Client) (*Connection, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapConnection(m, c)
}

func unmarshalMapConnection(m map[string]interface{}, c *Client) (*Connection, error) {

	return flattenConnection(c, m), nil
}

// expandConnection expands Connection into a JSON request object.
func expandConnection(c *Client, f *Connection) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/connections/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.FriendlyName; !dcl.IsEmptyValueIndirect(v) {
		m["friendlyName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandConnectionCloudSql(c, f.CloudSql); err != nil {
		return nil, fmt.Errorf("error expanding CloudSql into cloudSql: %w", err)
	} else if v != nil {
		m["cloudSql"] = v
	}
	if v := f.CreationTime; !dcl.IsEmptyValueIndirect(v) {
		m["creationTime"] = v
	}
	if v := f.LastModifiedTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastModifiedTime"] = v
	}
	if v := f.HasCredential; !dcl.IsEmptyValueIndirect(v) {
		m["hasCredential"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenConnection flattens Connection from a JSON request object into the
// Connection type.
func flattenConnection(c *Client, i interface{}) *Connection {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Connection{}
	res.Name = dcl.FlattenString(m["name"])
	res.FriendlyName = dcl.FlattenString(m["friendlyName"])
	res.Description = dcl.FlattenString(m["description"])
	res.CloudSql = flattenConnectionCloudSql(c, m["cloudSql"])
	res.CreationTime = dcl.FlattenInteger(m["creationTime"])
	res.LastModifiedTime = dcl.FlattenInteger(m["lastModifiedTime"])
	res.HasCredential = dcl.FlattenBool(m["hasCredential"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandConnectionCloudSqlMap expands the contents of ConnectionCloudSql into a JSON
// request object.
func expandConnectionCloudSqlMap(c *Client, f map[string]ConnectionCloudSql) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConnectionCloudSql(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConnectionCloudSqlSlice expands the contents of ConnectionCloudSql into a JSON
// request object.
func expandConnectionCloudSqlSlice(c *Client, f []ConnectionCloudSql) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConnectionCloudSql(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConnectionCloudSqlMap flattens the contents of ConnectionCloudSql from a JSON
// response object.
func flattenConnectionCloudSqlMap(c *Client, i interface{}) map[string]ConnectionCloudSql {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConnectionCloudSql{}
	}

	if len(a) == 0 {
		return map[string]ConnectionCloudSql{}
	}

	items := make(map[string]ConnectionCloudSql)
	for k, item := range a {
		items[k] = *flattenConnectionCloudSql(c, item.(map[string]interface{}))
	}

	return items
}

// flattenConnectionCloudSqlSlice flattens the contents of ConnectionCloudSql from a JSON
// response object.
func flattenConnectionCloudSqlSlice(c *Client, i interface{}) []ConnectionCloudSql {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectionCloudSql{}
	}

	if len(a) == 0 {
		return []ConnectionCloudSql{}
	}

	items := make([]ConnectionCloudSql, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectionCloudSql(c, item.(map[string]interface{})))
	}

	return items
}

// expandConnectionCloudSql expands an instance of ConnectionCloudSql into a JSON
// request object.
func expandConnectionCloudSql(c *Client, f *ConnectionCloudSql) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.InstanceId; !dcl.IsEmptyValueIndirect(v) {
		m["instanceId"] = v
	}
	if v := f.Database; !dcl.IsEmptyValueIndirect(v) {
		m["database"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v, err := expandConnectionCloudSqlCredential(c, f.Credential); err != nil {
		return nil, fmt.Errorf("error expanding Credential into credential: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["credential"] = v
	}

	return m, nil
}

// flattenConnectionCloudSql flattens an instance of ConnectionCloudSql from a JSON
// response object.
func flattenConnectionCloudSql(c *Client, i interface{}) *ConnectionCloudSql {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConnectionCloudSql{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConnectionCloudSql
	}
	r.InstanceId = dcl.FlattenString(m["instanceId"])
	r.Database = dcl.FlattenString(m["database"])
	r.Type = flattenConnectionCloudSqlTypeEnum(m["type"])
	r.Credential = flattenConnectionCloudSqlCredential(c, m["credential"])

	return r
}

// expandConnectionCloudSqlCredentialMap expands the contents of ConnectionCloudSqlCredential into a JSON
// request object.
func expandConnectionCloudSqlCredentialMap(c *Client, f map[string]ConnectionCloudSqlCredential) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConnectionCloudSqlCredential(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConnectionCloudSqlCredentialSlice expands the contents of ConnectionCloudSqlCredential into a JSON
// request object.
func expandConnectionCloudSqlCredentialSlice(c *Client, f []ConnectionCloudSqlCredential) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConnectionCloudSqlCredential(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConnectionCloudSqlCredentialMap flattens the contents of ConnectionCloudSqlCredential from a JSON
// response object.
func flattenConnectionCloudSqlCredentialMap(c *Client, i interface{}) map[string]ConnectionCloudSqlCredential {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConnectionCloudSqlCredential{}
	}

	if len(a) == 0 {
		return map[string]ConnectionCloudSqlCredential{}
	}

	items := make(map[string]ConnectionCloudSqlCredential)
	for k, item := range a {
		items[k] = *flattenConnectionCloudSqlCredential(c, item.(map[string]interface{}))
	}

	return items
}

// flattenConnectionCloudSqlCredentialSlice flattens the contents of ConnectionCloudSqlCredential from a JSON
// response object.
func flattenConnectionCloudSqlCredentialSlice(c *Client, i interface{}) []ConnectionCloudSqlCredential {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectionCloudSqlCredential{}
	}

	if len(a) == 0 {
		return []ConnectionCloudSqlCredential{}
	}

	items := make([]ConnectionCloudSqlCredential, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectionCloudSqlCredential(c, item.(map[string]interface{})))
	}

	return items
}

// expandConnectionCloudSqlCredential expands an instance of ConnectionCloudSqlCredential into a JSON
// request object.
func expandConnectionCloudSqlCredential(c *Client, f *ConnectionCloudSqlCredential) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}

	return m, nil
}

// flattenConnectionCloudSqlCredential flattens an instance of ConnectionCloudSqlCredential from a JSON
// response object.
func flattenConnectionCloudSqlCredential(c *Client, i interface{}) *ConnectionCloudSqlCredential {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConnectionCloudSqlCredential{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConnectionCloudSqlCredential
	}
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])

	return r
}

// flattenConnectionCloudSqlTypeEnumMap flattens the contents of ConnectionCloudSqlTypeEnum from a JSON
// response object.
func flattenConnectionCloudSqlTypeEnumMap(c *Client, i interface{}) map[string]ConnectionCloudSqlTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConnectionCloudSqlTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ConnectionCloudSqlTypeEnum{}
	}

	items := make(map[string]ConnectionCloudSqlTypeEnum)
	for k, item := range a {
		items[k] = *flattenConnectionCloudSqlTypeEnum(item.(interface{}))
	}

	return items
}

// flattenConnectionCloudSqlTypeEnumSlice flattens the contents of ConnectionCloudSqlTypeEnum from a JSON
// response object.
func flattenConnectionCloudSqlTypeEnumSlice(c *Client, i interface{}) []ConnectionCloudSqlTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectionCloudSqlTypeEnum{}
	}

	if len(a) == 0 {
		return []ConnectionCloudSqlTypeEnum{}
	}

	items := make([]ConnectionCloudSqlTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectionCloudSqlTypeEnum(item.(interface{})))
	}

	return items
}

// flattenConnectionCloudSqlTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ConnectionCloudSqlTypeEnum with the same value as that string.
func flattenConnectionCloudSqlTypeEnum(i interface{}) *ConnectionCloudSqlTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ConnectionCloudSqlTypeEnumRef("")
	}

	return ConnectionCloudSqlTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Connection) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalConnection(b, c)
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

type connectionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         connectionApiOperation
}

func convertFieldDiffsToConnectionDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]connectionDiff, error) {
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
	var diffs []connectionDiff
	// For each operation name, create a connectionDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := connectionDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToConnectionApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToConnectionApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (connectionApiOperation, error) {
	switch opName {

	case "updateConnectionUpdateConnectionOperation":
		return &updateConnectionUpdateConnectionOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
