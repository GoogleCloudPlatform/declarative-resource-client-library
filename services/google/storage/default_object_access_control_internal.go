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
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDefaultObjectAccessControlUpdateOperation) do(ctx context.Context, r *DefaultObjectAccessControl, c *Client) error {
	_, err := c.GetDefaultObjectAccessControl(ctx, r.urlNormalized())
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
		res := flattenDefaultObjectAccessControl(c, v)
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

	_, err := c.GetDefaultObjectAccessControl(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("DefaultObjectAccessControl not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetDefaultObjectAccessControl checking for existence. error: %v", err)
		return err
	}

	u, err := defaultObjectAccessControlDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetDefaultObjectAccessControl(ctx, r.urlNormalized())
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

	if _, err := c.GetDefaultObjectAccessControl(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDefaultObjectAccessControlRaw(ctx context.Context, r *DefaultObjectAccessControl) ([]byte, error) {

	u, err := defaultObjectAccessControlGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) defaultObjectAccessControlDiffsForRawDesired(ctx context.Context, rawDesired *DefaultObjectAccessControl, opts ...dcl.ApplyOption) (initial, desired *DefaultObjectAccessControl, diffs []defaultObjectAccessControlDiff, err error) {
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
	rawInitial, err := c.GetDefaultObjectAccessControl(ctx, fetchState.urlNormalized())
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
	rawDesired.ProjectTeam = canonicalizeDefaultObjectAccessControlProjectTeam(rawDesired.ProjectTeam, rawInitial.ProjectTeam, opts...)
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
			if !compareDefaultObjectAccessControlProjectTeam(c, &d, &n) {
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
		return des
	}

	var items []DefaultObjectAccessControlProjectTeam
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDefaultObjectAccessControlProjectTeam(c, &d, &n))
	}

	return items
}

type defaultObjectAccessControlDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         defaultObjectAccessControlApiOperation
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
func diffDefaultObjectAccessControl(c *Client, desired, actual *DefaultObjectAccessControl, opts ...dcl.ApplyOption) ([]defaultObjectAccessControlDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []defaultObjectAccessControlDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Domain, actual.Domain, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, defaultObjectAccessControlDiff{RequiresRecreate: true, FieldName: "Domain"})
	}

	if d, err := dcl.Diff(desired.Email, actual.Email, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, defaultObjectAccessControlDiff{RequiresRecreate: true, FieldName: "Email"})
	}

	if d, err := dcl.Diff(desired.Entity, actual.Entity, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, defaultObjectAccessControlDiff{
			UpdateOp: &updateDefaultObjectAccessControlUpdateOperation{}, FieldName: "Entity",
		})
	}

	if d, err := dcl.Diff(desired.EntityId, actual.EntityId, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, defaultObjectAccessControlDiff{RequiresRecreate: true, FieldName: "EntityId"})
	}

	if !dcl.IsZeroValue(desired.Entity) && !dcl.StringCanonicalize(desired.Entity, actual.Entity) {
		c.Config.Logger.Infof("Detected diff in Entity.\nDESIRED: %v\nACTUAL: %v", desired.Entity, actual.Entity)

		diffs = append(diffs, defaultObjectAccessControlDiff{
			UpdateOp:  &updateDefaultObjectAccessControlUpdateOperation{},
			FieldName: "Entity",
		})

	}
	if !reflect.DeepEqual(desired.Role, actual.Role) {
		c.Config.Logger.Infof("Detected diff in Role.\nDESIRED: %v\nACTUAL: %v", desired.Role, actual.Role)

		diffs = append(diffs, defaultObjectAccessControlDiff{
			UpdateOp:  &updateDefaultObjectAccessControlUpdateOperation{},
			FieldName: "Role",
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
	var deduped []defaultObjectAccessControlDiff
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
func compareDefaultObjectAccessControlProjectTeam(c *Client, desired, actual *DefaultObjectAccessControlProjectTeam) bool {
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

func compareDefaultObjectAccessControlProjectTeamSlice(c *Client, desired, actual []DefaultObjectAccessControlProjectTeam) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in DefaultObjectAccessControlProjectTeam, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareDefaultObjectAccessControlProjectTeam(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in DefaultObjectAccessControlProjectTeam, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareDefaultObjectAccessControlProjectTeamMap(c *Client, desired, actual map[string]DefaultObjectAccessControlProjectTeam) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in DefaultObjectAccessControlProjectTeam, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in DefaultObjectAccessControlProjectTeam, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareDefaultObjectAccessControlProjectTeam(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in DefaultObjectAccessControlProjectTeam, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareDefaultObjectAccessControlProjectTeamTeamEnumSlice(c *Client, desired, actual []DefaultObjectAccessControlProjectTeamTeamEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in DefaultObjectAccessControlProjectTeamTeamEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareDefaultObjectAccessControlProjectTeamTeamEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in DefaultObjectAccessControlProjectTeamTeamEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareDefaultObjectAccessControlProjectTeamTeamEnum(c *Client, desired, actual *DefaultObjectAccessControlProjectTeamTeamEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareDefaultObjectAccessControlRoleEnumSlice(c *Client, desired, actual []DefaultObjectAccessControlRoleEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in DefaultObjectAccessControlRoleEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareDefaultObjectAccessControlRoleEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in DefaultObjectAccessControlRoleEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareDefaultObjectAccessControlRoleEnum(c *Client, desired, actual *DefaultObjectAccessControlRoleEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *DefaultObjectAccessControl) urlNormalized() *DefaultObjectAccessControl {
	normalized := deepcopy.Copy(*r).(DefaultObjectAccessControl)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Bucket = dcl.SelfLinkToName(r.Bucket)
	normalized.Domain = dcl.SelfLinkToName(r.Domain)
	normalized.Email = dcl.SelfLinkToName(r.Email)
	normalized.Entity = dcl.SelfLinkToName(r.Entity)
	normalized.EntityId = dcl.SelfLinkToName(r.EntityId)
	return &normalized
}

func (r *DefaultObjectAccessControl) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity)
}

func (r *DefaultObjectAccessControl) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket)
}

func (r *DefaultObjectAccessControl) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Bucket), dcl.ValueOrEmptyString(n.Entity)
}

func (r *DefaultObjectAccessControl) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Bucket into bucket: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &DefaultObjectAccessControl{}
	r.Project = dcl.FlattenString(m["project"])
	r.Bucket = dcl.FlattenString(m["bucket"])
	r.Domain = dcl.FlattenString(m["domain"])
	r.Email = dcl.FlattenString(m["email"])
	r.Entity = dcl.FlattenString(m["entity"])
	r.EntityId = dcl.FlattenString(m["entityId"])
	r.ProjectTeam = flattenDefaultObjectAccessControlProjectTeam(c, m["projectTeam"])
	r.Role = flattenDefaultObjectAccessControlRoleEnum(m["role"])

	return r
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
		return true
	}
}
