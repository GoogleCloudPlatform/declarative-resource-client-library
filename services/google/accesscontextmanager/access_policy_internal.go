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
package accesscontextmanager

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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *AccessPolicy) validate() error {

	if err := dcl.Required(r, "title"); err != nil {
		return err
	}
	return nil
}

func accessPolicyGetURL(userBasePath string, r *AccessPolicy) (string, error) {
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(r.Parent),
	}
	return dcl.URL("accessPolicies?parent=organizations/{{parent}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil
}

func accessPolicyListURL(userBasePath, parent string) (string, error) {
	params := map[string]interface{}{
		"parent": parent,
	}
	return dcl.URL("accessPolicies?parent=organizations/{{parent}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil

}

func accessPolicyCreateURL(userBasePath string, _ ...interface{}) (string, error) {
	return dcl.URL("accessPolicies", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, map[string]interface{}{}), nil
}

func accessPolicyDeleteURL(userBasePath string, r *AccessPolicy) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("accessPolicies/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil
}

// accessPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type accessPolicyApiOperation interface {
	do(context.Context, *AccessPolicy, *Client) error
}

// newUpdateAccessPolicyUpdateRequest creates a request for an
// AccessPolicy resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAccessPolicyUpdateRequest(ctx context.Context, f *AccessPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		req["title"] = v
	}
	return req, nil
}

// marshalUpdateAccessPolicyUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateAccessPolicyUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAccessPolicyUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAccessPolicyUpdateOperation) do(ctx context.Context, r *AccessPolicy, c *Client) error {
	_, err := c.GetAccessPolicy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"title"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateAccessPolicyUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAccessPolicyUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAccessPolicyRaw(ctx context.Context, parent, pageToken string, pageSize int32) ([]byte, error) {
	u, err := accessPolicyListURL(c.Config.BasePath, parent)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AccessPolicyMaxPage {
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

type listAccessPolicyOperation struct {
	AccessPolicies []map[string]interface{} `json:"accessPolicies"`
	Token          string                   `json:"nextPageToken"`
}

func (c *Client) listAccessPolicy(ctx context.Context, parent, pageToken string, pageSize int32) ([]*AccessPolicy, string, error) {
	b, err := c.listAccessPolicyRaw(ctx, parent, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAccessPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AccessPolicy
	for _, v := range m.AccessPolicies {
		res := flattenAccessPolicy(c, v)
		res.Parent = &parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAccessPolicy(ctx context.Context, f func(*AccessPolicy) bool, resources []*AccessPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAccessPolicy(ctx, res)
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

type deleteAccessPolicyOperation struct{}

func (op *deleteAccessPolicyOperation) do(ctx context.Context, r *AccessPolicy, c *Client) error {

	_, err := c.GetAccessPolicy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AccessPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAccessPolicy checking for existence. error: %v", err)
		return err
	}

	u, err := accessPolicyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAccessPolicy(ctx, r.urlNormalized())
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
type createAccessPolicyOperation struct {
	response map[string]interface{}
}

func (op *createAccessPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAccessPolicyOperation) do(ctx context.Context, r *AccessPolicy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	u, err := accessPolicyCreateURL(c.Config.BasePath)

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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string")
	}
	r.Name = &name

	if _, err := c.GetAccessPolicy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAccessPolicyRaw(ctx context.Context, r *AccessPolicy) ([]byte, error) {

	u, err := accessPolicyGetURL(c.Config.BasePath, r.urlNormalized())
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

	b, err = dcl.ExtractElementFromList(b, "accessPolicies", r.matcher(c))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) accessPolicyDiffsForRawDesired(ctx context.Context, rawDesired *AccessPolicy, opts ...dcl.ApplyOption) (initial, desired *AccessPolicy, diffs []accessPolicyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AccessPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AccessPolicy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AccessPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		fetchState.Name = GetAccessPolicyNameResource(ctx, c, rawDesired)
	}
	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeAccessPolicyDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAccessPolicy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AccessPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AccessPolicy resource: %v", err)
		}
		c.Config.Logger.Info("Found that AccessPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAccessPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for AccessPolicy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AccessPolicy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAccessPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AccessPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAccessPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AccessPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAccessPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAccessPolicyInitialState(rawInitial, rawDesired *AccessPolicy) (*AccessPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAccessPolicyDesiredState(rawDesired, rawInitial *AccessPolicy, opts ...dcl.ApplyOption) (*AccessPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		rawDesired.Parent = rawInitial.Parent
	}
	if dcl.StringCanonicalize(rawDesired.Title, rawInitial.Title) {
		rawDesired.Title = rawInitial.Title
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}

	return rawDesired, nil
}

func canonicalizeAccessPolicyNewState(c *Client, rawNew, rawDesired *AccessPolicy) (*AccessPolicy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parent) && dcl.IsEmptyValueIndirect(rawDesired.Parent) {
		rawNew.Parent = rawDesired.Parent
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Parent, rawNew.Parent) {
			rawNew.Parent = rawDesired.Parent
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Title) && dcl.IsEmptyValueIndirect(rawDesired.Title) {
		rawNew.Title = rawDesired.Title
	} else {
		if dcl.StringCanonicalize(rawDesired.Title, rawNew.Title) {
			rawNew.Title = rawDesired.Title
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	return rawNew, nil
}

type accessPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         accessPolicyApiOperation
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
func diffAccessPolicy(c *Client, desired, actual *AccessPolicy, opts ...dcl.ApplyOption) ([]accessPolicyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []accessPolicyDiff
	if !dcl.IsZeroValue(desired.Parent) && !dcl.PartialSelfLinkToSelfLink(desired.Parent, actual.Parent) {
		c.Config.Logger.Infof("Detected diff in Parent.\nDESIRED: %v\nACTUAL: %v", desired.Parent, actual.Parent)
		diffs = append(diffs, accessPolicyDiff{
			RequiresRecreate: true,
			FieldName:        "Parent",
		})
	}
	if !dcl.IsZeroValue(desired.Title) && !dcl.StringCanonicalize(desired.Title, actual.Title) {
		c.Config.Logger.Infof("Detected diff in Title.\nDESIRED: %v\nACTUAL: %v", desired.Title, actual.Title)

		diffs = append(diffs, accessPolicyDiff{
			UpdateOp:  &updateAccessPolicyUpdateOperation{},
			FieldName: "Title",
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
	var deduped []accessPolicyDiff
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
func (r *AccessPolicy) urlNormalized() *AccessPolicy {
	normalized := deepcopy.Copy(*r).(AccessPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Parent = dcl.SelfLinkToName(r.Parent)
	normalized.Title = dcl.SelfLinkToName(r.Title)
	return &normalized
}

func (r *AccessPolicy) getFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Parent)
}

func (r *AccessPolicy) createFields() string {
	return ""
}

func (r *AccessPolicy) deleteFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *AccessPolicy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("accessPolicies/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AccessPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AccessPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandAccessPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AccessPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAccessPolicy decodes JSON responses into the AccessPolicy resource schema.
func unmarshalAccessPolicy(b []byte, c *Client) (*AccessPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAccessPolicy(m, c)
}

func unmarshalMapAccessPolicy(m map[string]interface{}, c *Client) (*AccessPolicy, error) {

	return flattenAccessPolicy(c, m), nil
}

// expandAccessPolicy expands AccessPolicy into a JSON request object.
func expandAccessPolicy(c *Client, f *AccessPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.DeriveField("organizations/%s", f.Parent, f.Parent); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		m["title"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}

	return m, nil
}

// flattenAccessPolicy flattens AccessPolicy from a JSON request object into the
// AccessPolicy type.
func flattenAccessPolicy(c *Client, i interface{}) *AccessPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &AccessPolicy{}
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	r.Parent = dcl.FlattenString(m["parent"])
	r.Title = dcl.FlattenString(m["title"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AccessPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAccessPolicy(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Parent == nil && ncr.Parent == nil {
			c.Config.Logger.Info("Both Parent fields null - considering equal.")
		} else if nr.Parent == nil || ncr.Parent == nil {
			c.Config.Logger.Info("Only one Parent field is null - considering unequal.")
			return false
		} else if *nr.Parent != *ncr.Parent {
			return false
		}
		return true
	}
}
