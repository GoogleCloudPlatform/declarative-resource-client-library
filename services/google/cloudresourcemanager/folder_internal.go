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
package cloudresourcemanager

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Folder) validate() error {

	if err := dcl.Required(r, "parent"); err != nil {
		return err
	}
	return nil
}

func (r *Folder) SetPolicyURL(userBasePath string) string {
	n := r.URLNormalized()
	fields := map[string]interface{}{
		"name": *n.Name,
	}
	return dcl.URL("v2/folders/{{name}}:setIamPolicy", "https://cloudresourcemanager.googleapis.com/", userBasePath, fields)
}

func (r *Folder) SetPolicyVerb() string {
	return "POST"
}

func (r *Folder) getPolicyURL(userBasePath string) string {
	n := r.URLNormalized()
	fields := map[string]interface{}{
		"name": *n.Name,
	}
	return dcl.URL("v2/folders/{{name}}:getIamPolicy", "https://cloudresourcemanager.googleapis.com/", userBasePath, fields)
}

func (r *Folder) IAMPolicyVersion() int {
	return 3
}

// folderApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type folderApiOperation interface {
	do(context.Context, *Folder, *Client) error
}

// newUpdateFolderMoveFolderRequest creates a request for an
// Folder resource's MoveFolder update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFolderMoveFolderRequest(ctx context.Context, f *Folder, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Parent; !dcl.IsEmptyValueIndirect(v) {
		req["parent"] = v
	}
	return req, nil
}

// marshalUpdateFolderMoveFolderRequest converts the update into
// the final JSON request body.
func marshalUpdateFolderMoveFolderRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFolderMoveFolderOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.
// newUpdateFolderUpdateFolderRequest creates a request for an
// Folder resource's UpdateFolder update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFolderUpdateFolderRequest(ctx context.Context, f *Folder, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	return req, nil
}

// marshalUpdateFolderUpdateFolderRequest converts the update into
// the final JSON request body.
func marshalUpdateFolderUpdateFolderRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFolderUpdateFolderOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFolderUpdateFolderOperation) do(ctx context.Context, r *Folder, c *Client) error {
	_, err := c.GetFolder(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateFolder")
	if err != nil {
		return err
	}

	req, err := newUpdateFolderUpdateFolderRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFolderUpdateFolderRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.CRMOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://cloudresourcemanager.googleapis.com/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listFolderRaw(ctx context.Context, parent, pageToken string, pageSize int32) ([]byte, error) {
	u, err := folderListURL(c.Config.BasePath, parent)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FolderMaxPage {
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

type listFolderOperation struct {
	Folders []map[string]interface{} `json:"folders"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listFolder(ctx context.Context, parent, pageToken string, pageSize int32) ([]*Folder, string, error) {
	b, err := c.listFolderRaw(ctx, parent, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFolderOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Folder
	for _, v := range m.Folders {
		res, err := unmarshalMapFolder(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Parent = &parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFolder(ctx context.Context, f func(*Folder) bool, resources []*Folder) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFolder(ctx, res)
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

type deleteFolderOperation struct{}

func (op *deleteFolderOperation) do(ctx context.Context, r *Folder, c *Client) error {
	r, err := c.GetFolder(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Folder not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetFolder checking for existence. error: %v", err)
		return err
	}

	u, err := folderDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Folder: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFolderOperation struct {
	response map[string]interface{}
}

func (op *createFolderOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFolderOperation) do(ctx context.Context, r *Folder, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	parent := r.createFields()
	u, err := folderCreateURL(c.Config.BasePath, parent)

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
	var o operations.CRMOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://cloudresourcemanager.googleapis.com/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetFolder(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFolderRaw(ctx context.Context, r *Folder) ([]byte, error) {

	u, err := folderGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) folderDiffsForRawDesired(ctx context.Context, rawDesired *Folder, opts ...dcl.ApplyOption) (initial, desired *Folder, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Folder
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Folder); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Folder, got %T", sh)
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
		desired, err := canonicalizeFolderDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFolder(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Folder resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Folder resource: %v", err)
		}
		c.Config.Logger.Info("Found that Folder resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFolderDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Folder: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Folder: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFolderInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Folder: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFolderDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Folder: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFolder(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFolderInitialState(rawInitial, rawDesired *Folder) (*Folder, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFolderDesiredState(rawDesired, rawInitial *Folder, opts ...dcl.ApplyOption) (*Folder, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.Parent, rawInitial.Parent) {
		rawDesired.Parent = rawInitial.Parent
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}

	return rawDesired, nil
}

func canonicalizeFolderNewState(c *Client, rawNew, rawDesired *Folder) (*Folder, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parent) && dcl.IsEmptyValueIndirect(rawDesired.Parent) {
		rawNew.Parent = rawDesired.Parent
	} else {
		if dcl.StringCanonicalize(rawDesired.Parent, rawNew.Parent) {
			rawNew.Parent = rawDesired.Parent
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffFolder(c *Client, desired, actual *Folder, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFolderMoveFolderOperation")}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFolderUpdateFolderOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *Folder) getFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *Folder) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Parent)
}

func (r *Folder) deleteFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

// marshal encodes the Folder resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Folder) marshal(c *Client) ([]byte, error) {
	m, err := expandFolder(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Folder: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFolder decodes JSON responses into the Folder resource schema.
func unmarshalFolder(b []byte, c *Client) (*Folder, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFolder(m, c)
}

func unmarshalMapFolder(m map[string]interface{}, c *Client) (*Folder, error) {

	return flattenFolder(c, m), nil
}

// expandFolder expands Folder into a JSON request object.
func expandFolder(c *Client, f *Folder) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Parent; !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.DeleteTime; !dcl.IsEmptyValueIndirect(v) {
		m["deleteTime"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}

	return m, nil
}

// flattenFolder flattens Folder from a JSON request object into the
// Folder type.
func flattenFolder(c *Client, i interface{}) *Folder {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Folder{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.Parent = dcl.FlattenString(m["parent"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.State = flattenFolderStateEnum(m["state"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.Etag = dcl.FlattenString(m["etag"])

	return res
}

// flattenFolderStateEnumSlice flattens the contents of FolderStateEnum from a JSON
// response object.
func flattenFolderStateEnumSlice(c *Client, i interface{}) []FolderStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FolderStateEnum{}
	}

	if len(a) == 0 {
		return []FolderStateEnum{}
	}

	items := make([]FolderStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFolderStateEnum(item.(interface{})))
	}

	return items
}

// flattenFolderStateEnum asserts that an interface is a string, and returns a
// pointer to a *FolderStateEnum with the same value as that string.
func flattenFolderStateEnum(i interface{}) *FolderStateEnum {
	s, ok := i.(string)
	if !ok {
		return FolderStateEnumRef("")
	}

	return FolderStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Folder) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFolder(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
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

type folderDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         folderApiOperation
}

func convertFieldDiffToFolderOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]folderDiff, error) {
	var diffs []folderDiff
	for _, op := range ops {
		diff := folderDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTofolderApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTofolderApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (folderApiOperation, error) {
	switch op {

	case "updateFolderMoveFolderOperation":
		return &updateFolderMoveFolderOperation{Diffs: diffs}, nil

	case "updateFolderUpdateFolderOperation":
		return &updateFolderUpdateFolderOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
