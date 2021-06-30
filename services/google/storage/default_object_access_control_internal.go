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
package storage

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

func (r *DefaultObjectAccessControl) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Bucket, "Bucket"); err != nil {
		return err
	}
	if err := dcl.Required(r, "entity"); err != nil {
		return err
	}
	if err := dcl.Required(r, "role"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ProjectTeam) {
		if err := r.ProjectTeam.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DefaultObjectAccessControlProjectTeam) validate() error {
	return nil
}

func defaultObjectAccessControlGetURL(userBasePath string, r *DefaultObjectAccessControl) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"bucket":  dcl.ValueOrEmptyString(r.Bucket),
		"entity":  dcl.ValueOrEmptyString(r.Entity),
	}
	return dcl.URL("b/{{bucket}}/defaultObjectAcl/{{entity}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

func defaultObjectAccessControlListURL(userBasePath, project, bucket string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"bucket":  bucket,
	}
	return dcl.URL("b/{{bucket}}/defaultObjectAcl?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func defaultObjectAccessControlCreateURL(userBasePath, project, bucket string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"bucket":  bucket,
	}
	return dcl.URL("b/{{bucket}}/defaultObjectAcl?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func defaultObjectAccessControlDeleteURL(userBasePath string, r *DefaultObjectAccessControl) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"bucket":  dcl.ValueOrEmptyString(r.Bucket),
		"entity":  dcl.ValueOrEmptyString(r.Entity),
	}
	return dcl.URL("b/{{bucket}}/defaultObjectAcl/{{entity}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

// defaultObjectAccessControlApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type defaultObjectAccessControlApiOperation interface {
	do(context.Context, *DefaultObjectAccessControl, *Client) error
}

// newUpdateDefaultObjectAccessControlUpdateRequest creates a request for an
// DefaultObjectAccessControl resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDefaultObjectAccessControlUpdateRequest(ctx context.Context, f *DefaultObjectAccessControl, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Entity; !dcl.IsEmptyValueIndirect(v) {
		req["entity"] = v
	}
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		req["role"] = v
	}
	return req, nil
}

// marshalUpdateDefaultObjectAccessControlUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateDefaultObjectAccessControlUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDefaultObjectAccessControlUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDefaultObjectAccessControlUpdateOperation) do(ctx context.Context, r *DefaultObjectAccessControl, c *Client) error {
	_, err := c.GetDefaultObjectAccessControl(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateDefaultObjectAccessControlUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateDefaultObjectAccessControlUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listDefaultObjectAccessControlRaw(ctx context.Context, project, bucket, pageToken string, pageSize int32) ([]byte, error) {
	u, err := defaultObjectAccessControlListURL(c.Config.BasePath, project, bucket)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DefaultObjectAccessControlMaxPage {
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

type listDefaultObjectAccessControlOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listDefaultObjectAccessControl(ctx context.Context, project, bucket, pageToken string, pageSize int32) ([]*DefaultObjectAccessControl, string, error) {
	b, err := c.listDefaultObjectAccessControlRaw(ctx, project, bucket, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDefaultObjectAccessControlOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*DefaultObjectAccessControl
	for _, v := range m.Items {
		res, err := unmarshalMapDefaultObjectAccessControl(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Bucket = &bucket
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDefaultObjectAccessControl(ctx context.Context, f func(*DefaultObjectAccessControl) bool, resources []*DefaultObjectAccessControl) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDefaultObjectAccessControl(ctx, res)
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

type deleteDefaultObjectAccessControlOperation struct{}

func (op *deleteDefaultObjectAccessControlOperation) do(ctx context.Context, r *DefaultObjectAccessControl, c *Client) error {
	r, err := c.GetDefaultObjectAccessControl(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("DefaultObjectAccessControl not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetDefaultObjectAccessControl checking for existence. error: %v", err)
		return err
	}

	u, err := defaultObjectAccessControlDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete DefaultObjectAccessControl: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetDefaultObjectAccessControl(ctx, r.URLNormalized())
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
type createDefaultObjectAccessControlOperation struct {
	response map[string]interface{}
}

func (op *createDefaultObjectAccessControlOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDefaultObjectAccessControlOperation) do(ctx context.Context, r *DefaultObjectAccessControl, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, bucket := r.createFields()
	u, err := defaultObjectAccessControlCreateURL(c.Config.BasePath, project, bucket)

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

	if _, err := c.GetDefaultObjectAccessControl(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDefaultObjectAccessControlRaw(ctx context.Context, r *DefaultObjectAccessControl) ([]byte, error) {

	u, err := defaultObjectAccessControlGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) defaultObjectAccessControlDiffsForRawDesired(ctx context.Context, rawDesired *DefaultObjectAccessControl, opts ...dcl.ApplyOption) (initial, desired *DefaultObjectAccessControl, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *DefaultObjectAccessControl
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*DefaultObjectAccessControl); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected DefaultObjectAccessControl, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDefaultObjectAccessControl(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a DefaultObjectAccessControl resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve DefaultObjectAccessControl resource: %v", err)
		}
		c.Config.Logger.Info("Found that DefaultObjectAccessControl resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDefaultObjectAccessControlDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for DefaultObjectAccessControl: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for DefaultObjectAccessControl: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDefaultObjectAccessControlInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for DefaultObjectAccessControl: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDefaultObjectAccessControlDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for DefaultObjectAccessControl: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDefaultObjectAccessControl(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDefaultObjectAccessControlInitialState(rawInitial, rawDesired *DefaultObjectAccessControl) (*DefaultObjectAccessControl, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDefaultObjectAccessControlDesiredState(rawDesired, rawInitial *DefaultObjectAccessControl, opts ...dcl.ApplyOption) (*DefaultObjectAccessControl, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ProjectTeam = canonicalizeDefaultObjectAccessControlProjectTeam(rawDesired.ProjectTeam, nil, opts...)

		return rawDesired, nil
	}

	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Bucket, rawInitial.Bucket) {
		rawDesired.Bucket = rawInitial.Bucket
	}
	if dcl.StringCanonicalize(rawDesired.Entity, rawInitial.Entity) {
		rawDesired.Entity = rawInitial.Entity
	}
	if dcl.IsZeroValue(rawDesired.Role) {
		rawDesired.Role = rawInitial.Role
	}

	return rawDesired, nil
}

func canonicalizeDefaultObjectAccessControlNewState(c *Client, rawNew, rawDesired *DefaultObjectAccessControl) (*DefaultObjectAccessControl, error) {

	rawNew.Project = rawDesired.Project

	rawNew.Bucket = rawDesired.Bucket

	if dcl.IsEmptyValueIndirect(rawNew.Domain) && dcl.IsEmptyValueIndirect(rawDesired.Domain) {
		rawNew.Domain = rawDesired.Domain
	} else {
		if dcl.StringCanonicalize(rawDesired.Domain, rawNew.Domain) {
			rawNew.Domain = rawDesired.Domain
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Email) && dcl.IsEmptyValueIndirect(rawDesired.Email) {
		rawNew.Email = rawDesired.Email
	} else {
		if dcl.StringCanonicalize(rawDesired.Email, rawNew.Email) {
			rawNew.Email = rawDesired.Email
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Entity) && dcl.IsEmptyValueIndirect(rawDesired.Entity) {
		rawNew.Entity = rawDesired.Entity
	} else {
		if dcl.StringCanonicalize(rawDesired.Entity, rawNew.Entity) {
			rawNew.Entity = rawDesired.Entity
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EntityId) && dcl.IsEmptyValueIndirect(rawDesired.EntityId) {
		rawNew.EntityId = rawDesired.EntityId
	} else {
		if dcl.StringCanonicalize(rawDesired.EntityId, rawNew.EntityId) {
			rawNew.EntityId = rawDesired.EntityId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProjectTeam) && dcl.IsEmptyValueIndirect(rawDesired.ProjectTeam) {
		rawNew.ProjectTeam = rawDesired.ProjectTeam
	} else {
		rawNew.ProjectTeam = canonicalizeNewDefaultObjectAccessControlProjectTeam(c, rawDesired.ProjectTeam, rawNew.ProjectTeam)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Role) && dcl.IsEmptyValueIndirect(rawDesired.Role) {
		rawNew.Role = rawDesired.Role
	} else {
	}

	return rawNew, nil
}

func canonicalizeDefaultObjectAccessControlProjectTeam(des, initial *DefaultObjectAccessControlProjectTeam, opts ...dcl.ApplyOption) *DefaultObjectAccessControlProjectTeam {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ProjectNumber, initial.ProjectNumber) || dcl.IsZeroValue(des.ProjectNumber) {
		des.ProjectNumber = initial.ProjectNumber
	}
	if dcl.IsZeroValue(des.Team) {
		des.Team = initial.Team
	}

	return des
}

func canonicalizeNewDefaultObjectAccessControlProjectTeam(c *Client, des, nw *DefaultObjectAccessControlProjectTeam) *DefaultObjectAccessControlProjectTeam {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ProjectNumber, nw.ProjectNumber) {
		nw.ProjectNumber = des.ProjectNumber
	}
	if dcl.IsZeroValue(nw.Team) {
		nw.Team = des.Team
	}

	return nw
}

func canonicalizeNewDefaultObjectAccessControlProjectTeamSet(c *Client, des, nw []DefaultObjectAccessControlProjectTeam) []DefaultObjectAccessControlProjectTeam {
	if des == nil {
		return nw
	}
	var reorderedNew []DefaultObjectAccessControlProjectTeam
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareDefaultObjectAccessControlProjectTeamNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewDefaultObjectAccessControlProjectTeamSlice(c *Client, des, nw []DefaultObjectAccessControlProjectTeam) []DefaultObjectAccessControlProjectTeam {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DefaultObjectAccessControlProjectTeam
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDefaultObjectAccessControlProjectTeam(c, &d, &n))
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
func diffDefaultObjectAccessControl(c *Client, desired, actual *DefaultObjectAccessControl, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Domain, actual.Domain, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Domain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Email, actual.Email, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Email")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Entity, actual.Entity, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDefaultObjectAccessControlUpdateOperation")}, fn.AddNest("Entity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EntityId, actual.EntityId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EntityId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProjectTeam, actual.ProjectTeam, dcl.Info{OutputOnly: true, ObjectFunction: compareDefaultObjectAccessControlProjectTeamNewStyle, EmptyObject: EmptyDefaultObjectAccessControlProjectTeam, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectTeam")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Role, actual.Role, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateDefaultObjectAccessControlUpdateOperation")}, fn.AddNest("Role")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareDefaultObjectAccessControlProjectTeamNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DefaultObjectAccessControlProjectTeam)
	if !ok {
		desiredNotPointer, ok := d.(DefaultObjectAccessControlProjectTeam)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DefaultObjectAccessControlProjectTeam or *DefaultObjectAccessControlProjectTeam", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DefaultObjectAccessControlProjectTeam)
	if !ok {
		actualNotPointer, ok := a.(DefaultObjectAccessControlProjectTeam)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DefaultObjectAccessControlProjectTeam", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectNumber, actual.ProjectNumber, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Team, actual.Team, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Team")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *DefaultObjectAccessControl) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity)
}

func (r *DefaultObjectAccessControl) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket)
}

func (r *DefaultObjectAccessControl) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity)
}

func (r *DefaultObjectAccessControl) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"bucket":  dcl.ValueOrEmptyString(n.Bucket),
			"entity":  dcl.ValueOrEmptyString(n.Entity),
		}
		return dcl.URL("b/{{bucket}}/defaultObjectAcl/{{entity}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the DefaultObjectAccessControl resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *DefaultObjectAccessControl) marshal(c *Client) ([]byte, error) {
	m, err := expandDefaultObjectAccessControl(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling DefaultObjectAccessControl: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDefaultObjectAccessControl decodes JSON responses into the DefaultObjectAccessControl resource schema.
func unmarshalDefaultObjectAccessControl(b []byte, c *Client) (*DefaultObjectAccessControl, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDefaultObjectAccessControl(m, c)
}

func unmarshalMapDefaultObjectAccessControl(m map[string]interface{}, c *Client) (*DefaultObjectAccessControl, error) {

	return flattenDefaultObjectAccessControl(c, m), nil
}

// expandDefaultObjectAccessControl expands DefaultObjectAccessControl into a JSON request object.
func expandDefaultObjectAccessControl(c *Client, f *DefaultObjectAccessControl) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Bucket into bucket: %w", err)
	} else if v != nil {
		m["bucket"] = v
	}
	if v := f.Domain; !dcl.IsEmptyValueIndirect(v) {
		m["domain"] = v
	}
	if v := f.Email; !dcl.IsEmptyValueIndirect(v) {
		m["email"] = v
	}
	if v := f.Entity; !dcl.IsEmptyValueIndirect(v) {
		m["entity"] = v
	}
	if v := f.EntityId; !dcl.IsEmptyValueIndirect(v) {
		m["entityId"] = v
	}
	if v, err := expandDefaultObjectAccessControlProjectTeam(c, f.ProjectTeam); err != nil {
		return nil, fmt.Errorf("error expanding ProjectTeam into projectTeam: %w", err)
	} else if v != nil {
		m["projectTeam"] = v
	}
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		m["role"] = v
	}

	return m, nil
}

// flattenDefaultObjectAccessControl flattens DefaultObjectAccessControl from a JSON request object into the
// DefaultObjectAccessControl type.
func flattenDefaultObjectAccessControl(c *Client, i interface{}) *DefaultObjectAccessControl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &DefaultObjectAccessControl{}
	res.Project = dcl.FlattenString(m["project"])
	res.Bucket = dcl.FlattenString(m["bucket"])
	res.Domain = dcl.FlattenString(m["domain"])
	res.Email = dcl.FlattenString(m["email"])
	res.Entity = dcl.FlattenString(m["entity"])
	res.EntityId = dcl.FlattenString(m["entityId"])
	res.ProjectTeam = flattenDefaultObjectAccessControlProjectTeam(c, m["projectTeam"])
	res.Role = flattenDefaultObjectAccessControlRoleEnum(m["role"])

	return res
}

// expandDefaultObjectAccessControlProjectTeamMap expands the contents of DefaultObjectAccessControlProjectTeam into a JSON
// request object.
func expandDefaultObjectAccessControlProjectTeamMap(c *Client, f map[string]DefaultObjectAccessControlProjectTeam) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDefaultObjectAccessControlProjectTeam(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDefaultObjectAccessControlProjectTeamSlice expands the contents of DefaultObjectAccessControlProjectTeam into a JSON
// request object.
func expandDefaultObjectAccessControlProjectTeamSlice(c *Client, f []DefaultObjectAccessControlProjectTeam) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDefaultObjectAccessControlProjectTeam(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDefaultObjectAccessControlProjectTeamMap flattens the contents of DefaultObjectAccessControlProjectTeam from a JSON
// response object.
func flattenDefaultObjectAccessControlProjectTeamMap(c *Client, i interface{}) map[string]DefaultObjectAccessControlProjectTeam {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DefaultObjectAccessControlProjectTeam{}
	}

	if len(a) == 0 {
		return map[string]DefaultObjectAccessControlProjectTeam{}
	}

	items := make(map[string]DefaultObjectAccessControlProjectTeam)
	for k, item := range a {
		items[k] = *flattenDefaultObjectAccessControlProjectTeam(c, item.(map[string]interface{}))
	}

	return items
}

// flattenDefaultObjectAccessControlProjectTeamSlice flattens the contents of DefaultObjectAccessControlProjectTeam from a JSON
// response object.
func flattenDefaultObjectAccessControlProjectTeamSlice(c *Client, i interface{}) []DefaultObjectAccessControlProjectTeam {
	a, ok := i.([]interface{})
	if !ok {
		return []DefaultObjectAccessControlProjectTeam{}
	}

	if len(a) == 0 {
		return []DefaultObjectAccessControlProjectTeam{}
	}

	items := make([]DefaultObjectAccessControlProjectTeam, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDefaultObjectAccessControlProjectTeam(c, item.(map[string]interface{})))
	}

	return items
}

// expandDefaultObjectAccessControlProjectTeam expands an instance of DefaultObjectAccessControlProjectTeam into a JSON
// request object.
func expandDefaultObjectAccessControlProjectTeam(c *Client, f *DefaultObjectAccessControlProjectTeam) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectNumber; !dcl.IsEmptyValueIndirect(v) {
		m["projectNumber"] = v
	}
	if v := f.Team; !dcl.IsEmptyValueIndirect(v) {
		m["team"] = v
	}

	return m, nil
}

// flattenDefaultObjectAccessControlProjectTeam flattens an instance of DefaultObjectAccessControlProjectTeam from a JSON
// response object.
func flattenDefaultObjectAccessControlProjectTeam(c *Client, i interface{}) *DefaultObjectAccessControlProjectTeam {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DefaultObjectAccessControlProjectTeam{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDefaultObjectAccessControlProjectTeam
	}
	r.ProjectNumber = dcl.FlattenString(m["projectNumber"])
	r.Team = flattenDefaultObjectAccessControlProjectTeamTeamEnum(m["team"])

	return r
}

// flattenDefaultObjectAccessControlProjectTeamTeamEnumSlice flattens the contents of DefaultObjectAccessControlProjectTeamTeamEnum from a JSON
// response object.
func flattenDefaultObjectAccessControlProjectTeamTeamEnumSlice(c *Client, i interface{}) []DefaultObjectAccessControlProjectTeamTeamEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DefaultObjectAccessControlProjectTeamTeamEnum{}
	}

	if len(a) == 0 {
		return []DefaultObjectAccessControlProjectTeamTeamEnum{}
	}

	items := make([]DefaultObjectAccessControlProjectTeamTeamEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDefaultObjectAccessControlProjectTeamTeamEnum(item.(interface{})))
	}

	return items
}

// flattenDefaultObjectAccessControlProjectTeamTeamEnum asserts that an interface is a string, and returns a
// pointer to a *DefaultObjectAccessControlProjectTeamTeamEnum with the same value as that string.
func flattenDefaultObjectAccessControlProjectTeamTeamEnum(i interface{}) *DefaultObjectAccessControlProjectTeamTeamEnum {
	s, ok := i.(string)
	if !ok {
		return DefaultObjectAccessControlProjectTeamTeamEnumRef("")
	}

	return DefaultObjectAccessControlProjectTeamTeamEnumRef(s)
}

// flattenDefaultObjectAccessControlRoleEnumSlice flattens the contents of DefaultObjectAccessControlRoleEnum from a JSON
// response object.
func flattenDefaultObjectAccessControlRoleEnumSlice(c *Client, i interface{}) []DefaultObjectAccessControlRoleEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []DefaultObjectAccessControlRoleEnum{}
	}

	if len(a) == 0 {
		return []DefaultObjectAccessControlRoleEnum{}
	}

	items := make([]DefaultObjectAccessControlRoleEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDefaultObjectAccessControlRoleEnum(item.(interface{})))
	}

	return items
}

// flattenDefaultObjectAccessControlRoleEnum asserts that an interface is a string, and returns a
// pointer to a *DefaultObjectAccessControlRoleEnum with the same value as that string.
func flattenDefaultObjectAccessControlRoleEnum(i interface{}) *DefaultObjectAccessControlRoleEnum {
	s, ok := i.(string)
	if !ok {
		return DefaultObjectAccessControlRoleEnumRef("")
	}

	return DefaultObjectAccessControlRoleEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *DefaultObjectAccessControl) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDefaultObjectAccessControl(b, c)
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
		if nr.Bucket == nil && ncr.Bucket == nil {
			c.Config.Logger.Info("Both Bucket fields null - considering equal.")
		} else if nr.Bucket == nil || ncr.Bucket == nil {
			c.Config.Logger.Info("Only one Bucket field is null - considering unequal.")
			return false
		} else if *nr.Bucket != *ncr.Bucket {
			return false
		}
		if nr.Entity == nil && ncr.Entity == nil {
			c.Config.Logger.Info("Both Entity fields null - considering equal.")
		} else if nr.Entity == nil || ncr.Entity == nil {
			c.Config.Logger.Info("Only one Entity field is null - considering unequal.")
			return false
		} else if *nr.Entity != *ncr.Entity {
			return false
		}
		return true
	}
}

type defaultObjectAccessControlDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         defaultObjectAccessControlApiOperation
}

func convertFieldDiffToDefaultObjectAccessControlOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]defaultObjectAccessControlDiff, error) {
	var diffs []defaultObjectAccessControlDiff
	for _, op := range ops {
		diff := defaultObjectAccessControlDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTodefaultObjectAccessControlApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTodefaultObjectAccessControlApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (defaultObjectAccessControlApiOperation, error) {
	switch op {

	case "updateDefaultObjectAccessControlUpdateOperation":
		return &updateDefaultObjectAccessControlUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
