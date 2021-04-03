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
package monitoring

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
)

func (r *Group) validate() error {

	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func groupGetURL(userBasePath string, r *Group) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/groups/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

func groupListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/groups", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func groupCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/groups", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func groupDeleteURL(userBasePath string, r *Group) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/groups/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

// groupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type groupApiOperation interface {
	do(context.Context, *Group, *Client) error
}

// newUpdateGroupUpdateRequest creates a request for an
// Group resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateGroupUpdateRequest(ctx context.Context, f *Group, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		req["filter"] = v
	}
	if v := f.IsCluster; !dcl.IsEmptyValueIndirect(v) {
		req["isCluster"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/groups/%s", f.ParentName, f.Project, f.ParentName); err != nil {
		return nil, fmt.Errorf("error expanding ParentName into parentName: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["parentName"] = v
	}
	return req, nil
}

// marshalUpdateGroupUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateGroupUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateGroupUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateGroupUpdateOperation) do(ctx context.Context, r *Group, c *Client) error {
	_, err := c.GetGroup(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateGroupUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateGroupUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listGroupRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := groupListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != GroupMaxPage {
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

type listGroupOperation struct {
	Group []map[string]interface{} `json:"group"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listGroup(ctx context.Context, project, pageToken string, pageSize int32) ([]*Group, string, error) {
	b, err := c.listGroupRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listGroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Group
	for _, v := range m.Group {
		res := flattenGroup(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllGroup(ctx context.Context, f func(*Group) bool, resources []*Group) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteGroup(ctx, res)
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

type deleteGroupOperation struct{}

func (op *deleteGroupOperation) do(ctx context.Context, r *Group, c *Client) error {

	_, err := c.GetGroup(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Group not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetGroup checking for existence. error: %v", err)
		return err
	}

	u, err := groupDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Group: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetGroup(ctx, r.urlNormalized())
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
type createGroupOperation struct {
	response map[string]interface{}
}

func (op *createGroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := groupCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetGroup(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getGroupRaw(ctx context.Context, r *Group) ([]byte, error) {

	u, err := groupGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) groupDiffsForRawDesired(ctx context.Context, rawDesired *Group, opts ...dcl.ApplyOption) (initial, desired *Group, diffs []groupDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Group
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Group); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Group, got %T", sh)
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
		desired, err := canonicalizeGroupDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetGroup(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Group resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Group resource: %v", err)
		}
		c.Config.Logger.Info("Found that Group resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeGroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Group: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Group: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeGroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Group: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeGroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Group: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffGroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeGroupInitialState(rawInitial, rawDesired *Group) (*Group, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeGroupDesiredState(rawDesired, rawInitial *Group, opts ...dcl.ApplyOption) (*Group, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Filter, rawInitial.Filter) {
		rawDesired.Filter = rawInitial.Filter
	}
	if dcl.BoolCanonicalize(rawDesired.IsCluster, rawInitial.IsCluster) {
		rawDesired.IsCluster = rawInitial.IsCluster
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.ParentName, rawInitial.ParentName) {
		rawDesired.ParentName = rawInitial.ParentName
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeGroupNewState(c *Client, rawNew, rawDesired *Group) (*Group, error) {

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Filter) && dcl.IsEmptyValueIndirect(rawDesired.Filter) {
		rawNew.Filter = rawDesired.Filter
	} else {
		if dcl.StringCanonicalize(rawDesired.Filter, rawNew.Filter) {
			rawNew.Filter = rawDesired.Filter
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IsCluster) && dcl.IsEmptyValueIndirect(rawDesired.IsCluster) {
		rawNew.IsCluster = rawDesired.IsCluster
	} else {
		if dcl.BoolCanonicalize(rawDesired.IsCluster, rawNew.IsCluster) {
			rawNew.IsCluster = rawDesired.IsCluster
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ParentName) && dcl.IsEmptyValueIndirect(rawDesired.ParentName) {
		rawNew.ParentName = rawDesired.ParentName
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.ParentName, rawNew.ParentName) {
			rawNew.ParentName = rawDesired.ParentName
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

type groupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         groupApiOperation
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
func diffGroup(c *Client, desired, actual *Group, opts ...dcl.ApplyOption) ([]groupDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []groupDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.DisplayName, actual.DisplayName, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, groupDiff{
			UpdateOp: &updateGroupUpdateOperation{}, FieldName: "DisplayName",
		})
	}

	if d, err := dcl.Diff(desired.Filter, actual.Filter, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, groupDiff{
			UpdateOp: &updateGroupUpdateOperation{}, FieldName: "Filter",
		})
	}

	if d, err := dcl.Diff(desired.IsCluster, actual.IsCluster, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, groupDiff{
			UpdateOp: &updateGroupUpdateOperation{}, FieldName: "IsCluster",
		})
	}

	if d, err := dcl.Diff(desired.ParentName, actual.ParentName, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, groupDiff{
			UpdateOp: &updateGroupUpdateOperation{}, FieldName: "ParentName",
		})
	}

	if !dcl.IsZeroValue(desired.DisplayName) && !dcl.StringCanonicalize(desired.DisplayName, actual.DisplayName) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %v\nACTUAL: %v", desired.DisplayName, actual.DisplayName)

		diffs = append(diffs, groupDiff{
			UpdateOp:  &updateGroupUpdateOperation{},
			FieldName: "DisplayName",
		})

	}
	if !dcl.IsZeroValue(desired.Filter) && !dcl.StringCanonicalize(desired.Filter, actual.Filter) {
		c.Config.Logger.Infof("Detected diff in Filter.\nDESIRED: %v\nACTUAL: %v", desired.Filter, actual.Filter)

		diffs = append(diffs, groupDiff{
			UpdateOp:  &updateGroupUpdateOperation{},
			FieldName: "Filter",
		})

	}
	if !dcl.IsZeroValue(desired.IsCluster) && !dcl.BoolCanonicalize(desired.IsCluster, actual.IsCluster) {
		c.Config.Logger.Infof("Detected diff in IsCluster.\nDESIRED: %v\nACTUAL: %v", desired.IsCluster, actual.IsCluster)

		diffs = append(diffs, groupDiff{
			UpdateOp:  &updateGroupUpdateOperation{},
			FieldName: "IsCluster",
		})

	}
	if !dcl.StringEqualsWithSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, groupDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.ParentName) && !dcl.PartialSelfLinkToSelfLink(desired.ParentName, actual.ParentName) {
		c.Config.Logger.Infof("Detected diff in ParentName.\nDESIRED: %v\nACTUAL: %v", desired.ParentName, actual.ParentName)

		diffs = append(diffs, groupDiff{
			UpdateOp:  &updateGroupUpdateOperation{},
			FieldName: "ParentName",
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
	var deduped []groupDiff
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
func (r *Group) urlNormalized() *Group {
	normalized := deepcopy.Copy(*r).(Group)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Filter = dcl.SelfLinkToName(r.Filter)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ParentName = dcl.SelfLinkToName(r.ParentName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Group) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Group) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Group) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Group) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("v3/projects/{{project}}/groups/{{name}}", "https://monitoring.googleapis.com/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Group resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Group) marshal(c *Client) ([]byte, error) {
	m, err := expandGroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Group: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalGroup decodes JSON responses into the Group resource schema.
func unmarshalGroup(b []byte, c *Client) (*Group, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapGroup(m, c)
}

func unmarshalMapGroup(m map[string]interface{}, c *Client) (*Group, error) {

	return flattenGroup(c, m), nil
}

// expandGroup expands Group into a JSON request object.
func expandGroup(c *Client, f *Group) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}
	if v := f.IsCluster; !dcl.IsEmptyValueIndirect(v) {
		m["isCluster"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/groups/%s", f.ParentName, f.Project, f.ParentName); err != nil {
		return nil, fmt.Errorf("error expanding ParentName into parentName: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parentName"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenGroup flattens Group from a JSON request object into the
// Group type.
func flattenGroup(c *Client, i interface{}) *Group {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Group{}
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.Filter = dcl.FlattenString(m["filter"])
	r.IsCluster = dcl.FlattenBool(m["isCluster"])
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	r.ParentName = dcl.FlattenString(m["parentName"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Group) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalGroup(b, c)
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
