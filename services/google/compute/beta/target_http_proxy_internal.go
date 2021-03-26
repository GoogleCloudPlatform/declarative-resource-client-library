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

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *TargetHttpProxy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "urlMap"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func targetHttpProxyGetURL(userBasePath string, r *TargetHttpProxy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/targetHttpProxies/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func targetHttpProxyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/targetHttpProxies", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func targetHttpProxyCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/targetHttpProxies", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func targetHttpProxyDeleteURL(userBasePath string, r *TargetHttpProxy) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/targetHttpProxies/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// targetHttpProxyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type targetHttpProxyApiOperation interface {
	do(context.Context, *TargetHttpProxy, *Client) error
}

// newUpdateTargetHttpProxySetURLMapRequest creates a request for an
// TargetHttpProxy resource's SetURLMap update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetHttpProxySetURLMapRequest(ctx context.Context, f *TargetHttpProxy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/global/urlMaps/%s", f.UrlMap, f.Project, f.UrlMap); err != nil {
		return nil, fmt.Errorf("error expanding UrlMap into urlMap: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["urlMap"] = v
	}
	return req, nil
}

// marshalUpdateTargetHttpProxySetURLMapRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetHttpProxySetURLMapRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTargetHttpProxySetURLMapOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetHttpProxySetURLMapOperation) do(ctx context.Context, r *TargetHttpProxy, c *Client) error {
	_, err := c.GetTargetHttpProxy(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "SetURLMap")
	if err != nil {
		return err
	}

	req, err := newUpdateTargetHttpProxySetURLMapRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateTargetHttpProxySetURLMapRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTargetHttpProxyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := targetHttpProxyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TargetHttpProxyMaxPage {
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

type listTargetHttpProxyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listTargetHttpProxy(ctx context.Context, project, pageToken string, pageSize int32) ([]*TargetHttpProxy, string, error) {
	b, err := c.listTargetHttpProxyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTargetHttpProxyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TargetHttpProxy
	for _, v := range m.Items {
		res := flattenTargetHttpProxy(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTargetHttpProxy(ctx context.Context, f func(*TargetHttpProxy) bool, resources []*TargetHttpProxy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTargetHttpProxy(ctx, res)
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

type deleteTargetHttpProxyOperation struct{}

func (op *deleteTargetHttpProxyOperation) do(ctx context.Context, r *TargetHttpProxy, c *Client) error {

	_, err := c.GetTargetHttpProxy(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("TargetHttpProxy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetTargetHttpProxy checking for existence. error: %v", err)
		return err
	}

	u, err := targetHttpProxyDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetTargetHttpProxy(ctx, r.urlNormalized())
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
type createTargetHttpProxyOperation struct {
	response map[string]interface{}
}

func (op *createTargetHttpProxyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTargetHttpProxyOperation) do(ctx context.Context, r *TargetHttpProxy, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := targetHttpProxyCreateURL(c.Config.BasePath, project)

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
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetTargetHttpProxy(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTargetHttpProxyRaw(ctx context.Context, r *TargetHttpProxy) ([]byte, error) {

	u, err := targetHttpProxyGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) targetHttpProxyDiffsForRawDesired(ctx context.Context, rawDesired *TargetHttpProxy, opts ...dcl.ApplyOption) (initial, desired *TargetHttpProxy, diffs []targetHttpProxyDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TargetHttpProxy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TargetHttpProxy); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected TargetHttpProxy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTargetHttpProxy(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a TargetHttpProxy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TargetHttpProxy resource: %v", err)
		}
		c.Config.Logger.Info("Found that TargetHttpProxy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTargetHttpProxyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for TargetHttpProxy: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for TargetHttpProxy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTargetHttpProxyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for TargetHttpProxy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTargetHttpProxyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for TargetHttpProxy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTargetHttpProxy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTargetHttpProxyInitialState(rawInitial, rawDesired *TargetHttpProxy) (*TargetHttpProxy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTargetHttpProxyDesiredState(rawDesired, rawInitial *TargetHttpProxy, opts ...dcl.ApplyOption) (*TargetHttpProxy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.UrlMap, rawInitial.UrlMap) {
		rawDesired.UrlMap = rawInitial.UrlMap
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeTargetHttpProxyNewState(c *Client, rawNew, rawDesired *TargetHttpProxy) (*TargetHttpProxy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UrlMap) && dcl.IsEmptyValueIndirect(rawDesired.UrlMap) {
		rawNew.UrlMap = rawDesired.UrlMap
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.UrlMap, rawNew.UrlMap) {
			rawNew.UrlMap = rawDesired.UrlMap
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

type targetHttpProxyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         targetHttpProxyApiOperation
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
func diffTargetHttpProxy(c *Client, desired, actual *TargetHttpProxy, opts ...dcl.ApplyOption) ([]targetHttpProxyDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []targetHttpProxyDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, targetHttpProxyDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, targetHttpProxyDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.UrlMap) && !dcl.PartialSelfLinkToSelfLink(desired.UrlMap, actual.UrlMap) {
		c.Config.Logger.Infof("Detected diff in UrlMap.\nDESIRED: %v\nACTUAL: %v", desired.UrlMap, actual.UrlMap)

		diffs = append(diffs, targetHttpProxyDiff{
			UpdateOp:  &updateTargetHttpProxySetURLMapOperation{},
			FieldName: "UrlMap",
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
	var deduped []targetHttpProxyDiff
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
func (r *TargetHttpProxy) urlNormalized() *TargetHttpProxy {
	normalized := deepcopy.Copy(*r).(TargetHttpProxy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.UrlMap = dcl.SelfLinkToName(r.UrlMap)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *TargetHttpProxy) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetHttpProxy) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *TargetHttpProxy) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *TargetHttpProxy) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "SetURLMap" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/targetHttpProxies/{{name}}/setUrlMap", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TargetHttpProxy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TargetHttpProxy) marshal(c *Client) ([]byte, error) {
	m, err := expandTargetHttpProxy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TargetHttpProxy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTargetHttpProxy decodes JSON responses into the TargetHttpProxy resource schema.
func unmarshalTargetHttpProxy(b []byte, c *Client) (*TargetHttpProxy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTargetHttpProxy(m, c)
}

func unmarshalMapTargetHttpProxy(m map[string]interface{}, c *Client) (*TargetHttpProxy, error) {

	return flattenTargetHttpProxy(c, m), nil
}

// expandTargetHttpProxy expands TargetHttpProxy into a JSON request object.
func expandTargetHttpProxy(c *Client, f *TargetHttpProxy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/global/urlMaps/%s", f.UrlMap, f.Project, f.UrlMap); err != nil {
		return nil, fmt.Errorf("error expanding UrlMap into urlMap: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlMap"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenTargetHttpProxy flattens TargetHttpProxy from a JSON request object into the
// TargetHttpProxy type.
func flattenTargetHttpProxy(c *Client, i interface{}) *TargetHttpProxy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &TargetHttpProxy{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.UrlMap = dcl.FlattenString(m["urlMap"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TargetHttpProxy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTargetHttpProxy(b, c)
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
