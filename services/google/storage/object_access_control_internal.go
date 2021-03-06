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

func (r *ObjectAccessControl) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "bucket"); err != nil {
		return err
	}
	if err := dcl.Required(r, "entity"); err != nil {
		return err
	}
	if err := dcl.Required(r, "role"); err != nil {
		return err
	}
	if err := dcl.Required(r, "object"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ProjectTeam) {
		if err := r.ProjectTeam.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ObjectAccessControlProjectTeam) validate() error {
	return nil
}

func objectAccessControlGetURL(userBasePath string, r *ObjectAccessControl) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"bucket":  dcl.ValueOrEmptyString(r.Bucket),
		"entity":  dcl.ValueOrEmptyString(r.Entity),
		"object":  dcl.ValueOrEmptyString(r.Object),
	}
	return dcl.URL("b/{{bucket}}/o/{{object}}/acl/{{entity}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

func objectAccessControlListURL(userBasePath, project, bucket, object string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"bucket":  bucket,
		"object":  object,
	}
	return dcl.URL("b/{{bucket}}/o/{{object}}/acl?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func objectAccessControlCreateURL(userBasePath, project, bucket, object string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"bucket":  bucket,
		"object":  object,
	}
	return dcl.URL("b/{{bucket}}/o/{{object}}/acl?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func objectAccessControlDeleteURL(userBasePath string, r *ObjectAccessControl) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"bucket":  dcl.ValueOrEmptyString(r.Bucket),
		"entity":  dcl.ValueOrEmptyString(r.Entity),
		"object":  dcl.ValueOrEmptyString(r.Object),
	}
	return dcl.URL("b/{{bucket}}/o/{{object}}/acl/{{entity}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

// objectAccessControlApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type objectAccessControlApiOperation interface {
	do(context.Context, *ObjectAccessControl, *Client) error
}

// newUpdateObjectAccessControlUpdateRequest creates a request for an
// ObjectAccessControl resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateObjectAccessControlUpdateRequest(ctx context.Context, f *ObjectAccessControl, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Bucket; !dcl.IsEmptyValueIndirect(v) {
		req["bucket"] = v
	}
	if v := f.Entity; !dcl.IsEmptyValueIndirect(v) {
		req["entity"] = v
	}
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		req["role"] = v
	}
	if v := f.Object; !dcl.IsEmptyValueIndirect(v) {
		req["object"] = v
	}
	return req, nil
}

// marshalUpdateObjectAccessControlUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateObjectAccessControlUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateObjectAccessControlUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateObjectAccessControlUpdateOperation) do(ctx context.Context, r *ObjectAccessControl, c *Client) error {
	_, err := c.GetObjectAccessControl(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateObjectAccessControlUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateObjectAccessControlUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listObjectAccessControlRaw(ctx context.Context, project, bucket, object, pageToken string, pageSize int32) ([]byte, error) {
	u, err := objectAccessControlListURL(c.Config.BasePath, project, bucket, object)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ObjectAccessControlMaxPage {
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

type listObjectAccessControlOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listObjectAccessControl(ctx context.Context, project, bucket, object, pageToken string, pageSize int32) ([]*ObjectAccessControl, string, error) {
	b, err := c.listObjectAccessControlRaw(ctx, project, bucket, object, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listObjectAccessControlOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ObjectAccessControl
	for _, v := range m.Items {
		res, err := unmarshalMapObjectAccessControl(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Bucket = &bucket
		res.Object = &object
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllObjectAccessControl(ctx context.Context, f func(*ObjectAccessControl) bool, resources []*ObjectAccessControl) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteObjectAccessControl(ctx, res)
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

type deleteObjectAccessControlOperation struct{}

func (op *deleteObjectAccessControlOperation) do(ctx context.Context, r *ObjectAccessControl, c *Client) error {
	r, err := c.GetObjectAccessControl(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ObjectAccessControl not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetObjectAccessControl checking for existence. error: %v", err)
		return err
	}

	u, err := objectAccessControlDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete ObjectAccessControl: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetObjectAccessControl(ctx, r.URLNormalized())
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
type createObjectAccessControlOperation struct {
	response map[string]interface{}
}

func (op *createObjectAccessControlOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createObjectAccessControlOperation) do(ctx context.Context, r *ObjectAccessControl, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, bucket, object := r.createFields()
	u, err := objectAccessControlCreateURL(c.Config.BasePath, project, bucket, object)

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

	if _, err := c.GetObjectAccessControl(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getObjectAccessControlRaw(ctx context.Context, r *ObjectAccessControl) ([]byte, error) {

	u, err := objectAccessControlGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) objectAccessControlDiffsForRawDesired(ctx context.Context, rawDesired *ObjectAccessControl, opts ...dcl.ApplyOption) (initial, desired *ObjectAccessControl, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ObjectAccessControl
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ObjectAccessControl); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ObjectAccessControl, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetObjectAccessControl(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ObjectAccessControl resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ObjectAccessControl resource: %v", err)
		}
		c.Config.Logger.Info("Found that ObjectAccessControl resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeObjectAccessControlDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for ObjectAccessControl: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ObjectAccessControl: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeObjectAccessControlInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ObjectAccessControl: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeObjectAccessControlDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ObjectAccessControl: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffObjectAccessControl(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeObjectAccessControlInitialState(rawInitial, rawDesired *ObjectAccessControl) (*ObjectAccessControl, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeObjectAccessControlDesiredState(rawDesired, rawInitial *ObjectAccessControl, opts ...dcl.ApplyOption) (*ObjectAccessControl, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ProjectTeam = canonicalizeObjectAccessControlProjectTeam(rawDesired.ProjectTeam, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &ObjectAccessControl{}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Bucket, rawInitial.Bucket) {
		canonicalDesired.Bucket = rawInitial.Bucket
	} else {
		canonicalDesired.Bucket = rawDesired.Bucket
	}
	if dcl.StringCanonicalize(rawDesired.Entity, rawInitial.Entity) {
		canonicalDesired.Entity = rawInitial.Entity
	} else {
		canonicalDesired.Entity = rawDesired.Entity
	}
	if dcl.IsZeroValue(rawDesired.Role) {
		canonicalDesired.Role = rawInitial.Role
	} else {
		canonicalDesired.Role = rawDesired.Role
	}
	if dcl.StringCanonicalize(rawDesired.Object, rawInitial.Object) {
		canonicalDesired.Object = rawInitial.Object
	} else {
		canonicalDesired.Object = rawDesired.Object
	}

	return canonicalDesired, nil
}

func canonicalizeObjectAccessControlNewState(c *Client, rawNew, rawDesired *ObjectAccessControl) (*ObjectAccessControl, error) {

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Bucket) && dcl.IsEmptyValueIndirect(rawDesired.Bucket) {
		rawNew.Bucket = rawDesired.Bucket
	} else {
		if dcl.NameToSelfLink(rawDesired.Bucket, rawNew.Bucket) {
			rawNew.Bucket = rawDesired.Bucket
		}
	}

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
		rawNew.ProjectTeam = canonicalizeNewObjectAccessControlProjectTeam(c, rawDesired.ProjectTeam, rawNew.ProjectTeam)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Role) && dcl.IsEmptyValueIndirect(rawDesired.Role) {
		rawNew.Role = rawDesired.Role
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Object) && dcl.IsEmptyValueIndirect(rawDesired.Object) {
		rawNew.Object = rawDesired.Object
	} else {
		if dcl.StringCanonicalize(rawDesired.Object, rawNew.Object) {
			rawNew.Object = rawDesired.Object
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Generation) && dcl.IsEmptyValueIndirect(rawDesired.Generation) {
		rawNew.Generation = rawDesired.Generation
	} else {
	}

	return rawNew, nil
}

func canonicalizeObjectAccessControlProjectTeam(des, initial *ObjectAccessControlProjectTeam, opts ...dcl.ApplyOption) *ObjectAccessControlProjectTeam {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ObjectAccessControlProjectTeam{}

	if dcl.NameToSelfLink(des.ProjectNumber, initial.ProjectNumber) || dcl.IsZeroValue(des.ProjectNumber) {
		cDes.ProjectNumber = initial.ProjectNumber
	} else {
		cDes.ProjectNumber = des.ProjectNumber
	}
	if dcl.IsZeroValue(des.Team) {
		des.Team = initial.Team
	} else {
		cDes.Team = des.Team
	}

	return cDes
}

func canonicalizeNewObjectAccessControlProjectTeam(c *Client, des, nw *ObjectAccessControlProjectTeam) *ObjectAccessControlProjectTeam {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ProjectNumber, nw.ProjectNumber) {
		nw.ProjectNumber = des.ProjectNumber
	}
	if dcl.IsZeroValue(nw.Team) {
		nw.Team = des.Team
	}

	return nw
}

func canonicalizeNewObjectAccessControlProjectTeamSet(c *Client, des, nw []ObjectAccessControlProjectTeam) []ObjectAccessControlProjectTeam {
	if des == nil {
		return nw
	}
	var reorderedNew []ObjectAccessControlProjectTeam
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareObjectAccessControlProjectTeamNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewObjectAccessControlProjectTeamSlice(c *Client, des, nw []ObjectAccessControlProjectTeam) []ObjectAccessControlProjectTeam {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ObjectAccessControlProjectTeam
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewObjectAccessControlProjectTeam(c, &d, &n))
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
func diffObjectAccessControl(c *Client, desired, actual *ObjectAccessControl, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Bucket, actual.Bucket, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateObjectAccessControlUpdateOperation")}, fn.AddNest("Bucket")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Entity, actual.Entity, dcl.Info{OperationSelector: dcl.TriggersOperation("updateObjectAccessControlUpdateOperation")}, fn.AddNest("Entity")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ProjectTeam, actual.ProjectTeam, dcl.Info{OutputOnly: true, ObjectFunction: compareObjectAccessControlProjectTeamNewStyle, EmptyObject: EmptyObjectAccessControlProjectTeam, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectTeam")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Role, actual.Role, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateObjectAccessControlUpdateOperation")}, fn.AddNest("Role")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Object, actual.Object, dcl.Info{OperationSelector: dcl.TriggersOperation("updateObjectAccessControlUpdateOperation")}, fn.AddNest("Object")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareObjectAccessControlProjectTeamNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ObjectAccessControlProjectTeam)
	if !ok {
		desiredNotPointer, ok := d.(ObjectAccessControlProjectTeam)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ObjectAccessControlProjectTeam or *ObjectAccessControlProjectTeam", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ObjectAccessControlProjectTeam)
	if !ok {
		actualNotPointer, ok := a.(ObjectAccessControlProjectTeam)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ObjectAccessControlProjectTeam", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectNumber, actual.ProjectNumber, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectNumber")); len(ds) != 0 || err != nil {
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

func (r *ObjectAccessControl) getFields() (string, string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity), dcl.ValueOrEmptyString(n.Object)
}

func (r *ObjectAccessControl) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Object)
}

func (r *ObjectAccessControl) deleteFields() (string, string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity), dcl.ValueOrEmptyString(n.Object)
}

func (r *ObjectAccessControl) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"bucket":  dcl.ValueOrEmptyString(n.Bucket),
			"entity":  dcl.ValueOrEmptyString(n.Entity),
			"object":  dcl.ValueOrEmptyString(n.Object),
		}
		return dcl.URL("b/{{bucket}}/o/{{object}}/acl/{{entity}}?userProject={{project}}", "https://www.googleapis.com/storage/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ObjectAccessControl resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ObjectAccessControl) marshal(c *Client) ([]byte, error) {
	m, err := expandObjectAccessControl(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ObjectAccessControl: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalObjectAccessControl decodes JSON responses into the ObjectAccessControl resource schema.
func unmarshalObjectAccessControl(b []byte, c *Client) (*ObjectAccessControl, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapObjectAccessControl(m, c)
}

func unmarshalMapObjectAccessControl(m map[string]interface{}, c *Client) (*ObjectAccessControl, error) {

	return flattenObjectAccessControl(c, m), nil
}

// expandObjectAccessControl expands ObjectAccessControl into a JSON request object.
func expandObjectAccessControl(c *Client, f *ObjectAccessControl) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.Bucket; !dcl.IsEmptyValueIndirect(v) {
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
	if v, err := expandObjectAccessControlProjectTeam(c, f.ProjectTeam); err != nil {
		return nil, fmt.Errorf("error expanding ProjectTeam into projectTeam: %w", err)
	} else if v != nil {
		m["projectTeam"] = v
	}
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		m["role"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Object; !dcl.IsEmptyValueIndirect(v) {
		m["object"] = v
	}
	if v := f.Generation; !dcl.IsEmptyValueIndirect(v) {
		m["generation"] = v
	}

	return m, nil
}

// flattenObjectAccessControl flattens ObjectAccessControl from a JSON request object into the
// ObjectAccessControl type.
func flattenObjectAccessControl(c *Client, i interface{}) *ObjectAccessControl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &ObjectAccessControl{}
	res.Project = dcl.FlattenString(m["project"])
	res.Bucket = dcl.FlattenString(m["bucket"])
	res.Domain = dcl.FlattenString(m["domain"])
	res.Email = dcl.FlattenString(m["email"])
	res.Entity = dcl.FlattenString(m["entity"])
	res.EntityId = dcl.FlattenString(m["entityId"])
	res.ProjectTeam = flattenObjectAccessControlProjectTeam(c, m["projectTeam"])
	res.Role = flattenObjectAccessControlRoleEnum(m["role"])
	res.Id = dcl.FlattenString(m["id"])
	res.Object = dcl.FlattenString(m["object"])
	res.Generation = dcl.FlattenInteger(m["generation"])

	return res
}

// expandObjectAccessControlProjectTeamMap expands the contents of ObjectAccessControlProjectTeam into a JSON
// request object.
func expandObjectAccessControlProjectTeamMap(c *Client, f map[string]ObjectAccessControlProjectTeam) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandObjectAccessControlProjectTeam(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandObjectAccessControlProjectTeamSlice expands the contents of ObjectAccessControlProjectTeam into a JSON
// request object.
func expandObjectAccessControlProjectTeamSlice(c *Client, f []ObjectAccessControlProjectTeam) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandObjectAccessControlProjectTeam(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenObjectAccessControlProjectTeamMap flattens the contents of ObjectAccessControlProjectTeam from a JSON
// response object.
func flattenObjectAccessControlProjectTeamMap(c *Client, i interface{}) map[string]ObjectAccessControlProjectTeam {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ObjectAccessControlProjectTeam{}
	}

	if len(a) == 0 {
		return map[string]ObjectAccessControlProjectTeam{}
	}

	items := make(map[string]ObjectAccessControlProjectTeam)
	for k, item := range a {
		items[k] = *flattenObjectAccessControlProjectTeam(c, item.(map[string]interface{}))
	}

	return items
}

// flattenObjectAccessControlProjectTeamSlice flattens the contents of ObjectAccessControlProjectTeam from a JSON
// response object.
func flattenObjectAccessControlProjectTeamSlice(c *Client, i interface{}) []ObjectAccessControlProjectTeam {
	a, ok := i.([]interface{})
	if !ok {
		return []ObjectAccessControlProjectTeam{}
	}

	if len(a) == 0 {
		return []ObjectAccessControlProjectTeam{}
	}

	items := make([]ObjectAccessControlProjectTeam, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenObjectAccessControlProjectTeam(c, item.(map[string]interface{})))
	}

	return items
}

// expandObjectAccessControlProjectTeam expands an instance of ObjectAccessControlProjectTeam into a JSON
// request object.
func expandObjectAccessControlProjectTeam(c *Client, f *ObjectAccessControlProjectTeam) (map[string]interface{}, error) {
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

// flattenObjectAccessControlProjectTeam flattens an instance of ObjectAccessControlProjectTeam from a JSON
// response object.
func flattenObjectAccessControlProjectTeam(c *Client, i interface{}) *ObjectAccessControlProjectTeam {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ObjectAccessControlProjectTeam{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyObjectAccessControlProjectTeam
	}
	r.ProjectNumber = dcl.FlattenString(m["projectNumber"])
	r.Team = flattenObjectAccessControlProjectTeamTeamEnum(m["team"])

	return r
}

// flattenObjectAccessControlProjectTeamTeamEnumMap flattens the contents of ObjectAccessControlProjectTeamTeamEnum from a JSON
// response object.
func flattenObjectAccessControlProjectTeamTeamEnumMap(c *Client, i interface{}) map[string]ObjectAccessControlProjectTeamTeamEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ObjectAccessControlProjectTeamTeamEnum{}
	}

	if len(a) == 0 {
		return map[string]ObjectAccessControlProjectTeamTeamEnum{}
	}

	items := make(map[string]ObjectAccessControlProjectTeamTeamEnum)
	for k, item := range a {
		items[k] = *flattenObjectAccessControlProjectTeamTeamEnum(item.(interface{}))
	}

	return items
}

// flattenObjectAccessControlProjectTeamTeamEnumSlice flattens the contents of ObjectAccessControlProjectTeamTeamEnum from a JSON
// response object.
func flattenObjectAccessControlProjectTeamTeamEnumSlice(c *Client, i interface{}) []ObjectAccessControlProjectTeamTeamEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ObjectAccessControlProjectTeamTeamEnum{}
	}

	if len(a) == 0 {
		return []ObjectAccessControlProjectTeamTeamEnum{}
	}

	items := make([]ObjectAccessControlProjectTeamTeamEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenObjectAccessControlProjectTeamTeamEnum(item.(interface{})))
	}

	return items
}

// flattenObjectAccessControlProjectTeamTeamEnum asserts that an interface is a string, and returns a
// pointer to a *ObjectAccessControlProjectTeamTeamEnum with the same value as that string.
func flattenObjectAccessControlProjectTeamTeamEnum(i interface{}) *ObjectAccessControlProjectTeamTeamEnum {
	s, ok := i.(string)
	if !ok {
		return ObjectAccessControlProjectTeamTeamEnumRef("")
	}

	return ObjectAccessControlProjectTeamTeamEnumRef(s)
}

// flattenObjectAccessControlRoleEnumMap flattens the contents of ObjectAccessControlRoleEnum from a JSON
// response object.
func flattenObjectAccessControlRoleEnumMap(c *Client, i interface{}) map[string]ObjectAccessControlRoleEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ObjectAccessControlRoleEnum{}
	}

	if len(a) == 0 {
		return map[string]ObjectAccessControlRoleEnum{}
	}

	items := make(map[string]ObjectAccessControlRoleEnum)
	for k, item := range a {
		items[k] = *flattenObjectAccessControlRoleEnum(item.(interface{}))
	}

	return items
}

// flattenObjectAccessControlRoleEnumSlice flattens the contents of ObjectAccessControlRoleEnum from a JSON
// response object.
func flattenObjectAccessControlRoleEnumSlice(c *Client, i interface{}) []ObjectAccessControlRoleEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ObjectAccessControlRoleEnum{}
	}

	if len(a) == 0 {
		return []ObjectAccessControlRoleEnum{}
	}

	items := make([]ObjectAccessControlRoleEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenObjectAccessControlRoleEnum(item.(interface{})))
	}

	return items
}

// flattenObjectAccessControlRoleEnum asserts that an interface is a string, and returns a
// pointer to a *ObjectAccessControlRoleEnum with the same value as that string.
func flattenObjectAccessControlRoleEnum(i interface{}) *ObjectAccessControlRoleEnum {
	s, ok := i.(string)
	if !ok {
		return ObjectAccessControlRoleEnumRef("")
	}

	return ObjectAccessControlRoleEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ObjectAccessControl) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalObjectAccessControl(b, c)
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
		if nr.Object == nil && ncr.Object == nil {
			c.Config.Logger.Info("Both Object fields null - considering equal.")
		} else if nr.Object == nil || ncr.Object == nil {
			c.Config.Logger.Info("Only one Object field is null - considering unequal.")
			return false
		} else if *nr.Object != *ncr.Object {
			return false
		}
		return true
	}
}

type objectAccessControlDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         objectAccessControlApiOperation
}

func convertFieldDiffsToObjectAccessControlDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]objectAccessControlDiff, error) {
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
	var diffs []objectAccessControlDiff
	// For each operation name, create a objectAccessControlDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := objectAccessControlDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToObjectAccessControlApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToObjectAccessControlApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (objectAccessControlApiOperation, error) {
	switch opName {

	case "updateObjectAccessControlUpdateOperation":
		return &updateObjectAccessControlUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
