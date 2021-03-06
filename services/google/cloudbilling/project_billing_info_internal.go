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
package cloudbilling

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *ProjectBillingInfo) validate() error {

	return nil
}

func projectBillingInfoGetURL(userBasePath string, r *ProjectBillingInfo) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{name}}/billingInfo", "https://cloudbilling.googleapis.com/v1/", userBasePath, params), nil
}

func projectBillingInfoDeleteURL(userBasePath string, r *ProjectBillingInfo) (string, error) {
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{name}}/billingInfo", "https://cloudbilling.googleapis.com/v1/", userBasePath, params), nil
}

// projectBillingInfoApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type projectBillingInfoApiOperation interface {
	do(context.Context, *ProjectBillingInfo, *Client) error
}

// newUpdateProjectBillingInfoUpdateProjectBillingInfoRequest creates a request for an
// ProjectBillingInfo resource's UpdateProjectBillingInfo update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateProjectBillingInfoUpdateProjectBillingInfoRequest(ctx context.Context, f *ProjectBillingInfo, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("billingAccounts/%s", f.BillingAccountName, f.BillingAccountName); err != nil {
		return nil, fmt.Errorf("error expanding BillingAccountName into billingAccountName: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["billingAccountName"] = v
	}
	return req, nil
}

// marshalUpdateProjectBillingInfoUpdateProjectBillingInfoRequest converts the update into
// the final JSON request body.
func marshalUpdateProjectBillingInfoUpdateProjectBillingInfoRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateProjectBillingInfoUpdateProjectBillingInfoOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateProjectBillingInfoUpdateProjectBillingInfoOperation) do(ctx context.Context, r *ProjectBillingInfo, c *Client) error {
	_, err := c.GetProjectBillingInfo(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateProjectBillingInfo")
	if err != nil {
		return err
	}

	req, err := newUpdateProjectBillingInfoUpdateProjectBillingInfoRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateProjectBillingInfoUpdateProjectBillingInfoRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteAllProjectBillingInfo(ctx context.Context, f func(*ProjectBillingInfo) bool, resources []*ProjectBillingInfo) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteProjectBillingInfo(ctx, res)
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

type deleteProjectBillingInfoOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createProjectBillingInfoOperation struct {
	response map[string]interface{}
}

func (op *createProjectBillingInfoOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getProjectBillingInfoRaw(ctx context.Context, r *ProjectBillingInfo) ([]byte, error) {

	u, err := projectBillingInfoGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) projectBillingInfoDiffsForRawDesired(ctx context.Context, rawDesired *ProjectBillingInfo, opts ...dcl.ApplyOption) (initial, desired *ProjectBillingInfo, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ProjectBillingInfo
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ProjectBillingInfo); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ProjectBillingInfo, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetProjectBillingInfo(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ProjectBillingInfo resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ProjectBillingInfo resource: %v", err)
		}
		c.Config.Logger.Info("Found that ProjectBillingInfo resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeProjectBillingInfoDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for ProjectBillingInfo: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ProjectBillingInfo: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeProjectBillingInfoInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ProjectBillingInfo: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeProjectBillingInfoDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ProjectBillingInfo: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffProjectBillingInfo(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeProjectBillingInfoInitialState(rawInitial, rawDesired *ProjectBillingInfo) (*ProjectBillingInfo, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeProjectBillingInfoDesiredState(rawDesired, rawInitial *ProjectBillingInfo, opts ...dcl.ApplyOption) (*ProjectBillingInfo, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &ProjectBillingInfo{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.BillingAccountName, rawInitial.BillingAccountName) {
		canonicalDesired.BillingAccountName = rawInitial.BillingAccountName
	} else {
		canonicalDesired.BillingAccountName = rawDesired.BillingAccountName
	}

	return canonicalDesired, nil
}

func canonicalizeProjectBillingInfoNewState(c *Client, rawNew, rawDesired *ProjectBillingInfo) (*ProjectBillingInfo, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.BillingAccountName) && dcl.IsEmptyValueIndirect(rawDesired.BillingAccountName) {
		rawNew.BillingAccountName = rawDesired.BillingAccountName
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.BillingAccountName, rawNew.BillingAccountName) {
			rawNew.BillingAccountName = rawDesired.BillingAccountName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.BillingEnabled) && dcl.IsEmptyValueIndirect(rawDesired.BillingEnabled) {
		rawNew.BillingEnabled = rawDesired.BillingEnabled
	} else {
		if dcl.StringCanonicalize(rawDesired.BillingEnabled, rawNew.BillingEnabled) {
			rawNew.BillingEnabled = rawDesired.BillingEnabled
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
func diffProjectBillingInfo(c *Client, desired, actual *ProjectBillingInfo, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BillingAccountName, actual.BillingAccountName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateProjectBillingInfoUpdateProjectBillingInfoOperation")}, fn.AddNest("BillingAccountName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BillingEnabled, actual.BillingEnabled, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BillingEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *ProjectBillingInfo) getFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *ProjectBillingInfo) createFields() string {
	return ""
}

func (r *ProjectBillingInfo) deleteFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Name)
}

func (r *ProjectBillingInfo) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateProjectBillingInfo" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{name}}/billingInfo", "https://cloudbilling.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ProjectBillingInfo resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ProjectBillingInfo) marshal(c *Client) ([]byte, error) {
	m, err := expandProjectBillingInfo(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ProjectBillingInfo: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalProjectBillingInfo decodes JSON responses into the ProjectBillingInfo resource schema.
func unmarshalProjectBillingInfo(b []byte, c *Client) (*ProjectBillingInfo, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapProjectBillingInfo(m, c)
}

func unmarshalMapProjectBillingInfo(m map[string]interface{}, c *Client) (*ProjectBillingInfo, error) {

	return flattenProjectBillingInfo(c, m), nil
}

// expandProjectBillingInfo expands ProjectBillingInfo into a JSON request object.
func expandProjectBillingInfo(c *Client, f *ProjectBillingInfo) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into projectId: %w", err)
	} else if v != nil {
		m["projectId"] = v
	}
	if v, err := dcl.DeriveField("billingAccounts/%s", f.BillingAccountName, f.BillingAccountName); err != nil {
		return nil, fmt.Errorf("error expanding BillingAccountName into billingAccountName: %w", err)
	} else {
		m["billingAccountName"] = v
	}
	if v := f.BillingEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["billingEnabled"] = v
	}

	return m, nil
}

// flattenProjectBillingInfo flattens ProjectBillingInfo from a JSON request object into the
// ProjectBillingInfo type.
func flattenProjectBillingInfo(c *Client, i interface{}) *ProjectBillingInfo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &ProjectBillingInfo{}
	res.Name = dcl.FlattenString(m["projectId"])
	res.BillingAccountName = dcl.FlattenString(m["billingAccountName"])
	res.BillingEnabled = dcl.FlattenString(m["billingEnabled"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ProjectBillingInfo) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalProjectBillingInfo(b, c)
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

type projectBillingInfoDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         projectBillingInfoApiOperation
}

func convertFieldDiffsToProjectBillingInfoDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]projectBillingInfoDiff, error) {
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
	var diffs []projectBillingInfoDiff
	// For each operation name, create a projectBillingInfoDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := projectBillingInfoDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToProjectBillingInfoApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToProjectBillingInfoApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (projectBillingInfoApiOperation, error) {
	switch opName {

	case "updateProjectBillingInfoUpdateProjectBillingInfoOperation":
		return &updateProjectBillingInfoUpdateProjectBillingInfoOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
