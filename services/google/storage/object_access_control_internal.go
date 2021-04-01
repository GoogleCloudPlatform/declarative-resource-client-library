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
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateObjectAccessControlUpdateOperation) do(ctx context.Context, r *ObjectAccessControl, c *Client) error {
	_, err := c.GetObjectAccessControl(ctx, r.urlNormalized())
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
		res := flattenObjectAccessControl(c, v)
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

	_, err := c.GetObjectAccessControl(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ObjectAccessControl not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetObjectAccessControl checking for existence. error: %v", err)
		return err
	}

	u, err := objectAccessControlDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetObjectAccessControl(ctx, r.urlNormalized())
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

	if _, err := c.GetObjectAccessControl(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getObjectAccessControlRaw(ctx context.Context, r *ObjectAccessControl) ([]byte, error) {

	u, err := objectAccessControlGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) objectAccessControlDiffsForRawDesired(ctx context.Context, rawDesired *ObjectAccessControl, opts ...dcl.ApplyOption) (initial, desired *ObjectAccessControl, diffs []objectAccessControlDiff, err error) {
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
	rawInitial, err := c.GetObjectAccessControl(ctx, fetchState.urlNormalized())
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
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Bucket, rawInitial.Bucket) {
		rawDesired.Bucket = rawInitial.Bucket
	}
	if dcl.StringCanonicalize(rawDesired.Domain, rawInitial.Domain) {
		rawDesired.Domain = rawInitial.Domain
	}
	if dcl.StringCanonicalize(rawDesired.Email, rawInitial.Email) {
		rawDesired.Email = rawInitial.Email
	}
	if dcl.StringCanonicalize(rawDesired.Entity, rawInitial.Entity) {
		rawDesired.Entity = rawInitial.Entity
	}
	if dcl.StringCanonicalize(rawDesired.EntityId, rawInitial.EntityId) {
		rawDesired.EntityId = rawInitial.EntityId
	}
	rawDesired.ProjectTeam = canonicalizeObjectAccessControlProjectTeam(rawDesired.ProjectTeam, rawInitial.ProjectTeam, opts...)
	if dcl.IsZeroValue(rawDesired.Role) {
		rawDesired.Role = rawInitial.Role
	}
	if dcl.StringCanonicalize(rawDesired.Id, rawInitial.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Object, rawInitial.Object) {
		rawDesired.Object = rawInitial.Object
	}
	if dcl.IsZeroValue(rawDesired.Generation) {
		rawDesired.Generation = rawInitial.Generation
	}

	return rawDesired, nil
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

	if dcl.StringCanonicalize(des.ProjectNumber, initial.ProjectNumber) || dcl.IsZeroValue(des.ProjectNumber) {
		des.ProjectNumber = initial.ProjectNumber
	}
	if dcl.IsZeroValue(des.Team) {
		des.Team = initial.Team
	}

	return des
}

func canonicalizeNewObjectAccessControlProjectTeam(c *Client, des, nw *ObjectAccessControlProjectTeam) *ObjectAccessControlProjectTeam {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ProjectNumber, nw.ProjectNumber) {
		nw.ProjectNumber = des.ProjectNumber
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
			if !compareObjectAccessControlProjectTeam(c, &d, &n) {
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
		return des
	}

	var items []ObjectAccessControlProjectTeam
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewObjectAccessControlProjectTeam(c, &d, &n))
	}

	return items
}

type objectAccessControlDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         objectAccessControlApiOperation
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
func diffObjectAccessControl(c *Client, desired, actual *ObjectAccessControl, opts ...dcl.ApplyOption) ([]objectAccessControlDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []objectAccessControlDiff
	if !dcl.IsZeroValue(desired.Bucket) && !dcl.NameToSelfLink(desired.Bucket, actual.Bucket) {
		c.Config.Logger.Infof("Detected diff in Bucket.\nDESIRED: %v\nACTUAL: %v", desired.Bucket, actual.Bucket)

		diffs = append(diffs, objectAccessControlDiff{
			UpdateOp:  &updateObjectAccessControlUpdateOperation{},
			FieldName: "Bucket",
		})

	}
	if !dcl.IsZeroValue(desired.Entity) && !dcl.StringCanonicalize(desired.Entity, actual.Entity) {
		c.Config.Logger.Infof("Detected diff in Entity.\nDESIRED: %v\nACTUAL: %v", desired.Entity, actual.Entity)

		diffs = append(diffs, objectAccessControlDiff{
			UpdateOp:  &updateObjectAccessControlUpdateOperation{},
			FieldName: "Entity",
		})

	}
	if !reflect.DeepEqual(desired.Role, actual.Role) {
		c.Config.Logger.Infof("Detected diff in Role.\nDESIRED: %v\nACTUAL: %v", desired.Role, actual.Role)

		diffs = append(diffs, objectAccessControlDiff{
			UpdateOp:  &updateObjectAccessControlUpdateOperation{},
			FieldName: "Role",
		})

	}
	if !dcl.IsZeroValue(desired.Object) && !dcl.StringCanonicalize(desired.Object, actual.Object) {
		c.Config.Logger.Infof("Detected diff in Object.\nDESIRED: %v\nACTUAL: %v", desired.Object, actual.Object)

		diffs = append(diffs, objectAccessControlDiff{
			UpdateOp:  &updateObjectAccessControlUpdateOperation{},
			FieldName: "Object",
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
	var deduped []objectAccessControlDiff
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
func compareObjectAccessControlProjectTeam(c *Client, desired, actual *ObjectAccessControlProjectTeam) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ProjectNumber == nil && desired.ProjectNumber != nil && !dcl.IsEmptyValueIndirect(desired.ProjectNumber) {
		c.Config.Logger.Infof("desired ProjectNumber %s - but actually nil", dcl.SprintResource(desired.ProjectNumber))
		return true
	}
	if !dcl.StringCanonicalize(desired.ProjectNumber, actual.ProjectNumber) && !dcl.IsZeroValue(desired.ProjectNumber) {
		c.Config.Logger.Infof("Diff in ProjectNumber. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ProjectNumber), dcl.SprintResource(actual.ProjectNumber))
		return true
	}
	if actual.Team == nil && desired.Team != nil && !dcl.IsEmptyValueIndirect(desired.Team) {
		c.Config.Logger.Infof("desired Team %s - but actually nil", dcl.SprintResource(desired.Team))
		return true
	}
	if !reflect.DeepEqual(desired.Team, actual.Team) && !dcl.IsZeroValue(desired.Team) {
		c.Config.Logger.Infof("Diff in Team. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Team), dcl.SprintResource(actual.Team))
		return true
	}
	return false
}

func compareObjectAccessControlProjectTeamSlice(c *Client, desired, actual []ObjectAccessControlProjectTeam) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ObjectAccessControlProjectTeam, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareObjectAccessControlProjectTeam(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ObjectAccessControlProjectTeam, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareObjectAccessControlProjectTeamMap(c *Client, desired, actual map[string]ObjectAccessControlProjectTeam) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ObjectAccessControlProjectTeam, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ObjectAccessControlProjectTeam, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareObjectAccessControlProjectTeam(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ObjectAccessControlProjectTeam, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareObjectAccessControlProjectTeamTeamEnumSlice(c *Client, desired, actual []ObjectAccessControlProjectTeamTeamEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ObjectAccessControlProjectTeamTeamEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareObjectAccessControlProjectTeamTeamEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ObjectAccessControlProjectTeamTeamEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareObjectAccessControlProjectTeamTeamEnum(c *Client, desired, actual *ObjectAccessControlProjectTeamTeamEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareObjectAccessControlRoleEnumSlice(c *Client, desired, actual []ObjectAccessControlRoleEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ObjectAccessControlRoleEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareObjectAccessControlRoleEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ObjectAccessControlRoleEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareObjectAccessControlRoleEnum(c *Client, desired, actual *ObjectAccessControlRoleEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ObjectAccessControl) urlNormalized() *ObjectAccessControl {
	normalized := deepcopy.Copy(*r).(ObjectAccessControl)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Bucket = dcl.SelfLinkToName(r.Bucket)
	normalized.Domain = dcl.SelfLinkToName(r.Domain)
	normalized.Email = dcl.SelfLinkToName(r.Email)
	normalized.Entity = dcl.SelfLinkToName(r.Entity)
	normalized.EntityId = dcl.SelfLinkToName(r.EntityId)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.Object = dcl.SelfLinkToName(r.Object)
	return &normalized
}

func (r *ObjectAccessControl) getFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity), dcl.ValueOrEmptyString(n.Object)
}

func (r *ObjectAccessControl) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Object)
}

func (r *ObjectAccessControl) deleteFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity), dcl.ValueOrEmptyString(n.Object)
}

func (r *ObjectAccessControl) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &ObjectAccessControl{}
	r.Project = dcl.FlattenString(m["project"])
	r.Bucket = dcl.FlattenString(m["bucket"])
	r.Domain = dcl.FlattenString(m["domain"])
	r.Email = dcl.FlattenString(m["email"])
	r.Entity = dcl.FlattenString(m["entity"])
	r.EntityId = dcl.FlattenString(m["entityId"])
	r.ProjectTeam = flattenObjectAccessControlProjectTeam(c, m["projectTeam"])
	r.Role = flattenObjectAccessControlRoleEnum(m["role"])
	r.Id = dcl.FlattenString(m["id"])
	r.Object = dcl.FlattenString(m["object"])
	r.Generation = dcl.FlattenInteger(m["generation"])

	return r
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
	r.ProjectNumber = dcl.FlattenString(m["projectNumber"])
	r.Team = flattenObjectAccessControlProjectTeamTeamEnum(m["team"])

	return r
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
