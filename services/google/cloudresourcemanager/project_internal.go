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
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Project) validate() error {

	if !dcl.IsEmptyValueIndirect(r.Parent) {
		if err := r.Parent.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ProjectParent) validate() error {
	return nil
}

func projectGetURL(userBasePath string, r *Project) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v1/projects/{{name}}", "https://cloudresourcemanager.googleapis.com/", userBasePath, params), nil
}

func projectListURL(userBasePath string) (string, error) {
	params := map[string]interface{}{}
	return dcl.URL("v1/projects", "https://cloudresourcemanager.googleapis.com/", userBasePath, params), nil

}

func projectCreateURL(userBasePath string, _ ...interface{}) (string, error) {
	return dcl.URL("v1/projects", "https://cloudresourcemanager.googleapis.com/", userBasePath, map[string]interface{}{}), nil
}

func projectDeleteURL(userBasePath string, r *Project) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v1/projects/{{name}}", "https://cloudresourcemanager.googleapis.com/", userBasePath, params), nil
}

func (r *Project) SetPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"name": *n.Name,
	}
	return dcl.URL("v1/projects/{{name}}:setIamPolicy", "https://cloudresourcemanager.googleapis.com/", userBasePath, fields)
}

func (r *Project) getPolicyURL(userBasePath string) string {
	n := r.urlNormalized()
	fields := map[string]interface{}{
		"name": *n.Name,
	}
	return dcl.URL("v1/projects/{{name}}:getIamPolicy", "https://cloudresourcemanager.googleapis.com/", userBasePath, fields)
}

func (r *Project) IAMPolicyVersion() int {
	return 3
}

// projectApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type projectApiOperation interface {
	do(context.Context, *Project, *Client) error
}

// newUpdateProjectUpdateProjectRequest creates a request for an
// Project resource's UpdateProject update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateProjectUpdateProjectRequest(ctx context.Context, f *Project, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateProjectUpdateProjectRequest converts the update into
// the final JSON request body.
func marshalUpdateProjectUpdateProjectRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateProjectUpdateProjectOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateProjectUpdateProjectOperation) do(ctx context.Context, r *Project, c *Client) error {
	_, err := c.GetProject(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateProject")
	if err != nil {
		return err
	}

	req, err := newUpdateProjectUpdateProjectRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateProjectUpdateProjectRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listProjectRaw(ctx context.Context, pageToken string, pageSize int32) ([]byte, error) {
	u, err := projectListURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ProjectMaxPage {
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

type listProjectOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listProject(ctx context.Context, pageToken string, pageSize int32) ([]*Project, string, error) {
	b, err := c.listProjectRaw(ctx, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listProjectOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Project
	for _, v := range m.Items {
		res := flattenProject(c, v)
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllProject(ctx context.Context, f func(*Project) bool, resources []*Project) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteProject(ctx, res)
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

type deleteProjectOperation struct{}

func (op *deleteProjectOperation) do(ctx context.Context, r *Project, c *Client) error {

	r, err := c.GetProject(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.Infof("Project not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetProject checking for existence. error: %v", err)
		return err
	}

	if projectDeletePrecondition(r) {
		return nil
	}

	u, err := projectDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Project: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createProjectOperation struct {
	response map[string]interface{}
}

func (op *createProjectOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createProjectOperation) do(ctx context.Context, r *Project, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	u, err := projectCreateURL(c.Config.BasePath)

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

	if _, err := c.GetProject(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getProjectRaw(ctx context.Context, r *Project) ([]byte, error) {

	u, err := projectGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) projectDiffsForRawDesired(ctx context.Context, rawDesired *Project, opts ...dcl.ApplyOption) (initial, desired *Project, diffs []projectDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Project
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Project); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Project, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetProject(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Project resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Project resource: %v", err)
		}
		c.Config.Logger.Info("Found that Project resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeProjectDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Project: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Project: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeProjectInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Project: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeProjectDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Project: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffProject(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeProjectInitialState(rawInitial, rawDesired *Project) (*Project, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeProjectDesiredState(rawDesired, rawInitial *Project, opts ...dcl.ApplyOption) (*Project, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Parent = canonicalizeProjectParent(rawDesired.Parent, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.LifecycleState) {
		rawDesired.LifecycleState = rawInitial.LifecycleState
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	rawDesired.Parent = canonicalizeProjectParent(rawDesired.Parent, rawInitial.Parent, opts...)
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.ProjectNumber) {
		rawDesired.ProjectNumber = rawInitial.ProjectNumber
	}

	return rawDesired, nil
}

func canonicalizeProjectNewState(c *Client, rawNew, rawDesired *Project) (*Project, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LifecycleState) && dcl.IsEmptyValueIndirect(rawDesired.LifecycleState) {
		rawNew.LifecycleState = rawDesired.LifecycleState
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parent) && dcl.IsEmptyValueIndirect(rawDesired.Parent) {
		rawNew.Parent = rawDesired.Parent
	} else {
		rawNew.Parent = canonicalizeNewProjectParent(c, rawDesired.Parent, rawNew.Parent)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProjectNumber) && dcl.IsEmptyValueIndirect(rawDesired.ProjectNumber) {
		rawNew.ProjectNumber = rawDesired.ProjectNumber
	} else {
	}

	return rawNew, nil
}

func canonicalizeProjectParent(des, initial *ProjectParent, opts ...dcl.ApplyOption) *ProjectParent {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		des.Id = initial.Id
	}

	return des
}

func canonicalizeNewProjectParent(c *Client, des, nw *ProjectParent) *ProjectParent {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) || dcl.IsZeroValue(des.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Id, nw.Id) || dcl.IsZeroValue(des.Id) {
		nw.Id = des.Id
	}

	return nw
}

func canonicalizeNewProjectParentSet(c *Client, des, nw []ProjectParent) []ProjectParent {
	if des == nil {
		return nw
	}
	var reorderedNew []ProjectParent
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareProjectParent(c, &d, &n) {
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

func canonicalizeNewProjectParentSlice(c *Client, des, nw []ProjectParent) []ProjectParent {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ProjectParent
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewProjectParent(c, &d, &n))
	}

	return items
}

type projectDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         projectApiOperation
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
func diffProject(c *Client, desired, actual *Project, opts ...dcl.ApplyOption) ([]projectDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []projectDiff
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, projectDiff{
			UpdateOp:  &updateProjectUpdateProjectOperation{},
			FieldName: "Labels",
		})

	}
	if !dcl.IsZeroValue(desired.DisplayName) && !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %v\nACTUAL: %v", desired.DisplayName, actual.DisplayName)
		diffs = append(diffs, projectDiff{
			RequiresRecreate: true,
			FieldName:        "DisplayName",
		})
	}
	if compareProjectParent(c, desired.Parent, actual.Parent) {
		c.Config.Logger.Infof("Detected diff in Parent.\nDESIRED: %v\nACTUAL: %v", desired.Parent, actual.Parent)
		diffs = append(diffs, projectDiff{
			RequiresRecreate: true,
			FieldName:        "Parent",
		})
	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, projectDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
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
	var deduped []projectDiff
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
func compareProjectParent(c *Client, desired, actual *ProjectParent) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !dcl.StringCanonicalize(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if actual.Id == nil && desired.Id != nil && !dcl.IsEmptyValueIndirect(desired.Id) {
		c.Config.Logger.Infof("desired Id %s - but actually nil", dcl.SprintResource(desired.Id))
		return true
	}
	if !dcl.StringCanonicalize(desired.Id, actual.Id) && !dcl.IsZeroValue(desired.Id) {
		c.Config.Logger.Infof("Diff in Id. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Id), dcl.SprintResource(actual.Id))
		return true
	}
	return false
}

func compareProjectParentSlice(c *Client, desired, actual []ProjectParent) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ProjectParent, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareProjectParent(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ProjectParent, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareProjectParentMap(c *Client, desired, actual map[string]ProjectParent) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ProjectParent, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ProjectParent, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareProjectParent(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ProjectParent, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareProjectLifecycleStateEnumSlice(c *Client, desired, actual []ProjectLifecycleStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ProjectLifecycleStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareProjectLifecycleStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ProjectLifecycleStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareProjectLifecycleStateEnum(c *Client, desired, actual *ProjectLifecycleStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Project) urlNormalized() *Project {
	normalized := deepcopy.Copy(*r).(Project)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	return &normalized
}

func (r *Project) getFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *Project) createFields() string {
	return ""
}

func (r *Project) deleteFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *Project) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateProject" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("v1/projects/{{name}}", "https://cloudresourcemanager.googleapis.com/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Project resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Project) marshal(c *Client) ([]byte, error) {
	m, err := expandProject(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Project: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalProject decodes JSON responses into the Project resource schema.
func unmarshalProject(b []byte, c *Client) (*Project, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapProject(m, c)
}

func unmarshalMapProject(m map[string]interface{}, c *Client) (*Project, error) {

	return flattenProject(c, m), nil
}

// expandProject expands Project into a JSON request object.
func expandProject(c *Client, f *Project) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LifecycleState; !dcl.IsEmptyValueIndirect(v) {
		m["lifecycleState"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandProjectParent(c, f.Parent); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.ProjectNumber; !dcl.IsEmptyValueIndirect(v) {
		m["projectNumber"] = v
	}

	return m, nil
}

// flattenProject flattens Project from a JSON request object into the
// Project type.
func flattenProject(c *Client, i interface{}) *Project {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Project{}
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.LifecycleState = flattenProjectLifecycleStateEnum(m["lifecycleState"])
	r.DisplayName = dcl.FlattenString(m["name"])
	r.Parent = flattenProjectParent(c, m["parent"])
	r.Name = dcl.FlattenString(m["projectId"])
	r.ProjectNumber = dcl.FlattenInteger(m["projectNumber"])

	return r
}

// expandProjectParentMap expands the contents of ProjectParent into a JSON
// request object.
func expandProjectParentMap(c *Client, f map[string]ProjectParent) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandProjectParent(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandProjectParentSlice expands the contents of ProjectParent into a JSON
// request object.
func expandProjectParentSlice(c *Client, f []ProjectParent) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandProjectParent(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenProjectParentMap flattens the contents of ProjectParent from a JSON
// response object.
func flattenProjectParentMap(c *Client, i interface{}) map[string]ProjectParent {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ProjectParent{}
	}

	if len(a) == 0 {
		return map[string]ProjectParent{}
	}

	items := make(map[string]ProjectParent)
	for k, item := range a {
		items[k] = *flattenProjectParent(c, item.(map[string]interface{}))
	}

	return items
}

// flattenProjectParentSlice flattens the contents of ProjectParent from a JSON
// response object.
func flattenProjectParentSlice(c *Client, i interface{}) []ProjectParent {
	a, ok := i.([]interface{})
	if !ok {
		return []ProjectParent{}
	}

	if len(a) == 0 {
		return []ProjectParent{}
	}

	items := make([]ProjectParent, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenProjectParent(c, item.(map[string]interface{})))
	}

	return items
}

// expandProjectParent expands an instance of ProjectParent into a JSON
// request object.
func expandProjectParent(c *Client, f *ProjectParent) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}

	return m, nil
}

// flattenProjectParent flattens an instance of ProjectParent from a JSON
// response object.
func flattenProjectParent(c *Client, i interface{}) *ProjectParent {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ProjectParent{}
	r.Type = dcl.FlattenString(m["type"])
	r.Id = dcl.FlattenString(m["id"])

	return r
}

// flattenProjectLifecycleStateEnumSlice flattens the contents of ProjectLifecycleStateEnum from a JSON
// response object.
func flattenProjectLifecycleStateEnumSlice(c *Client, i interface{}) []ProjectLifecycleStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ProjectLifecycleStateEnum{}
	}

	if len(a) == 0 {
		return []ProjectLifecycleStateEnum{}
	}

	items := make([]ProjectLifecycleStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenProjectLifecycleStateEnum(item.(interface{})))
	}

	return items
}

// flattenProjectLifecycleStateEnum asserts that an interface is a string, and returns a
// pointer to a *ProjectLifecycleStateEnum with the same value as that string.
func flattenProjectLifecycleStateEnum(i interface{}) *ProjectLifecycleStateEnum {
	s, ok := i.(string)
	if !ok {
		return ProjectLifecycleStateEnumRef("")
	}

	return ProjectLifecycleStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Project) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalProject(b, c)
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
