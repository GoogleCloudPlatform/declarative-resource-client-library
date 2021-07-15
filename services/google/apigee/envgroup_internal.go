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
package apigee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Envgroup) validate() error {

	if err := dcl.Required(r, "hostnames"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Organization, "Organization"); err != nil {
		return err
	}
	return nil
}

func envgroupGetURL(userBasePath string, r *Envgroup) (string, error) {
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(r.Organization),
		"name":         dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("organizations/{{organization}}/envgroups/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

func envgroupListURL(userBasePath, organization string) (string, error) {
	params := map[string]interface{}{
		"organization": organization,
	}
	return dcl.URL("organizations/{{organization}}/envgroups", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func envgroupCreateURL(userBasePath, organization string) (string, error) {
	params := map[string]interface{}{
		"organization": organization,
	}
	return dcl.URL("organizations/{{organization}}/envgroups", "https://apigee.googleapis.com/v1/", userBasePath, params), nil

}

func envgroupDeleteURL(userBasePath string, r *Envgroup) (string, error) {
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(r.Organization),
		"name":         dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("organizations/{{organization}}/envgroups/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, params), nil
}

// envgroupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type envgroupApiOperation interface {
	do(context.Context, *Envgroup, *Client) error
}

// newUpdateEnvgroupPatchEnvironmentGroupRequest creates a request for an
// Envgroup resource's PatchEnvironmentGroup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEnvgroupPatchEnvironmentGroupRequest(ctx context.Context, f *Envgroup, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateEnvgroupPatchEnvironmentGroupRequest converts the update into
// the final JSON request body.
func marshalUpdateEnvgroupPatchEnvironmentGroupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEnvgroupPatchEnvironmentGroupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEnvgroupPatchEnvironmentGroupOperation) do(ctx context.Context, r *Envgroup, c *Client) error {
	_, err := c.GetEnvgroup(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchEnvironmentGroup")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateEnvgroupPatchEnvironmentGroupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateEnvgroupPatchEnvironmentGroupRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listEnvgroupRaw(ctx context.Context, organization, pageToken string, pageSize int32) ([]byte, error) {
	u, err := envgroupListURL(c.Config.BasePath, organization)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EnvgroupMaxPage {
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

type listEnvgroupOperation struct {
	EnvironmentGroups []map[string]interface{} `json:"environmentGroups"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listEnvgroup(ctx context.Context, organization, pageToken string, pageSize int32) ([]*Envgroup, string, error) {
	b, err := c.listEnvgroupRaw(ctx, organization, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEnvgroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Envgroup
	for _, v := range m.EnvironmentGroups {
		res, err := unmarshalMapEnvgroup(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Organization = &organization
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEnvgroup(ctx context.Context, f func(*Envgroup) bool, resources []*Envgroup) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEnvgroup(ctx, res)
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

type deleteEnvgroupOperation struct{}

func (op *deleteEnvgroupOperation) do(ctx context.Context, r *Envgroup, c *Client) error {
	r, err := c.GetEnvgroup(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Envgroup not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetEnvgroup checking for existence. error: %v", err)
		return err
	}

	u, err := envgroupDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetEnvgroup(ctx, r.URLNormalized())
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
type createEnvgroupOperation struct {
	response map[string]interface{}
}

func (op *createEnvgroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEnvgroupOperation) do(ctx context.Context, r *Envgroup, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	organization := r.createFields()
	u, err := envgroupCreateURL(c.Config.BasePath, organization)

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
	if err := o.Wait(ctx, c.Config, "https://apigee.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetEnvgroup(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEnvgroupRaw(ctx context.Context, r *Envgroup) ([]byte, error) {

	u, err := envgroupGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) envgroupDiffsForRawDesired(ctx context.Context, rawDesired *Envgroup, opts ...dcl.ApplyOption) (initial, desired *Envgroup, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Envgroup
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Envgroup); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Envgroup, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEnvgroup(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Envgroup resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Envgroup resource: %v", err)
		}
		c.Config.Logger.Info("Found that Envgroup resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEnvgroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Envgroup: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Envgroup: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEnvgroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Envgroup: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEnvgroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Envgroup: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEnvgroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEnvgroupInitialState(rawInitial, rawDesired *Envgroup) (*Envgroup, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEnvgroupDesiredState(rawDesired, rawInitial *Envgroup, opts ...dcl.ApplyOption) (*Envgroup, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Envgroup{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Hostnames) {
		canonicalDesired.Hostnames = rawInitial.Hostnames
	} else {
		canonicalDesired.Hostnames = rawDesired.Hostnames
	}
	if dcl.NameToSelfLink(rawDesired.Organization, rawInitial.Organization) {
		canonicalDesired.Organization = rawInitial.Organization
	} else {
		canonicalDesired.Organization = rawDesired.Organization
	}

	return canonicalDesired, nil
}

func canonicalizeEnvgroupNewState(c *Client, rawNew, rawDesired *Envgroup) (*Envgroup, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Hostnames) && dcl.IsEmptyValueIndirect(rawDesired.Hostnames) {
		rawNew.Hostnames = rawDesired.Hostnames
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedAt) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedAt) {
		rawNew.LastModifiedAt = rawDesired.LastModifiedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Organization = rawDesired.Organization

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffEnvgroup(c *Client, desired, actual *Envgroup, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Hostnames, actual.Hostnames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Hostnames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedAt, actual.CreatedAt, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreatedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedAt, actual.LastModifiedAt, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedAt")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Organization, actual.Organization, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Organization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *Envgroup) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Organization), dcl.ValueOrEmptyString(n.Name)
}

func (r *Envgroup) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Organization)
}

func (r *Envgroup) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Organization), dcl.ValueOrEmptyString(n.Name)
}

func (r *Envgroup) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "PatchEnvironmentGroup" {
		fields := map[string]interface{}{
			"organization": dcl.ValueOrEmptyString(n.Organization),
			"name":         dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("organizations/{{organization}}/envgroups/{{name}}", "https://apigee.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Envgroup resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Envgroup) marshal(c *Client) ([]byte, error) {
	m, err := expandEnvgroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Envgroup: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEnvgroup decodes JSON responses into the Envgroup resource schema.
func unmarshalEnvgroup(b []byte, c *Client) (*Envgroup, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEnvgroup(m, c)
}

func unmarshalMapEnvgroup(m map[string]interface{}, c *Client) (*Envgroup, error) {

	return flattenEnvgroup(c, m), nil
}

// expandEnvgroup expands Envgroup into a JSON request object.
func expandEnvgroup(c *Client, f *Envgroup) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Hostnames; !dcl.IsEmptyValueIndirect(v) {
		m["hostnames"] = v
	}
	if v := f.CreatedAt; !dcl.IsEmptyValueIndirect(v) {
		m["createdAt"] = v
	}
	if v := f.LastModifiedAt; !dcl.IsEmptyValueIndirect(v) {
		m["lastModifiedAt"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Organization into organization: %w", err)
	} else if v != nil {
		m["organization"] = v
	}

	return m, nil
}

// flattenEnvgroup flattens Envgroup from a JSON request object into the
// Envgroup type.
func flattenEnvgroup(c *Client, i interface{}) *Envgroup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Envgroup{}
	res.Name = dcl.FlattenString(m["name"])
	res.Hostnames = dcl.FlattenStringSlice(m["hostnames"])
	res.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	res.LastModifiedAt = dcl.FlattenInteger(m["lastModifiedAt"])
	res.State = flattenEnvgroupStateEnum(m["state"])
	res.Organization = dcl.FlattenString(m["organization"])

	return res
}

// flattenEnvgroupStateEnumMap flattens the contents of EnvgroupStateEnum from a JSON
// response object.
func flattenEnvgroupStateEnumMap(c *Client, i interface{}) map[string]EnvgroupStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvgroupStateEnum{}
	}

	if len(a) == 0 {
		return map[string]EnvgroupStateEnum{}
	}

	items := make(map[string]EnvgroupStateEnum)
	for k, item := range a {
		items[k] = *flattenEnvgroupStateEnum(item.(interface{}))
	}

	return items
}

// flattenEnvgroupStateEnumSlice flattens the contents of EnvgroupStateEnum from a JSON
// response object.
func flattenEnvgroupStateEnumSlice(c *Client, i interface{}) []EnvgroupStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvgroupStateEnum{}
	}

	if len(a) == 0 {
		return []EnvgroupStateEnum{}
	}

	items := make([]EnvgroupStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvgroupStateEnum(item.(interface{})))
	}

	return items
}

// flattenEnvgroupStateEnum asserts that an interface is a string, and returns a
// pointer to a *EnvgroupStateEnum with the same value as that string.
func flattenEnvgroupStateEnum(i interface{}) *EnvgroupStateEnum {
	s, ok := i.(string)
	if !ok {
		return EnvgroupStateEnumRef("")
	}

	return EnvgroupStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Envgroup) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEnvgroup(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Organization == nil && ncr.Organization == nil {
			c.Config.Logger.Info("Both Organization fields null - considering equal.")
		} else if nr.Organization == nil || ncr.Organization == nil {
			c.Config.Logger.Info("Only one Organization field is null - considering unequal.")
			return false
		} else if *nr.Organization != *ncr.Organization {
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

type envgroupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         envgroupApiOperation
}

func convertFieldDiffsToEnvgroupDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]envgroupDiff, error) {
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
	var diffs []envgroupDiff
	// For each operation name, create a envgroupDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := envgroupDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEnvgroupApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEnvgroupApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (envgroupApiOperation, error) {
	switch opName {

	case "updateEnvgroupPatchEnvironmentGroupOperation":
		return &updateEnvgroupPatchEnvironmentGroupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
