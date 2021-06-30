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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Variable) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Text", "Value"}, r.Text, r.Value); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.RuntimeConfig, "RuntimeConfig"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func variableGetURL(userBasePath string, r *Variable) (string, error) {
	params := map[string]interface{}{
		"project":       dcl.ValueOrEmptyString(r.Project),
		"runtimeConfig": dcl.ValueOrEmptyString(r.RuntimeConfig),
		"name":          dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/configs/{{runtimeConfig}}/variables/{{name}}", "https://runtimeconfig.googleapis.com/v1beta1/", userBasePath, params), nil
}

func variableListURL(userBasePath, project, runtimeConfig string) (string, error) {
	params := map[string]interface{}{
		"project":       project,
		"runtimeConfig": runtimeConfig,
	}
	return dcl.URL("projects/{{project}}/configs/{{runtimeConfig}}/variables", "https://runtimeconfig.googleapis.com/v1beta1/", userBasePath, params), nil

}

func variableCreateURL(userBasePath, project, runtimeConfig string) (string, error) {
	params := map[string]interface{}{
		"project":       project,
		"runtimeConfig": runtimeConfig,
	}
	return dcl.URL("projects/{{project}}/configs/{{runtimeConfig}}/variables", "https://runtimeconfig.googleapis.com/v1beta1/", userBasePath, params), nil

}

func variableDeleteURL(userBasePath string, r *Variable) (string, error) {
	params := map[string]interface{}{
		"project":       dcl.ValueOrEmptyString(r.Project),
		"runtimeConfig": dcl.ValueOrEmptyString(r.RuntimeConfig),
		"name":          dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/configs/{{runtimeConfig}}/variables/{{name}}", "https://runtimeconfig.googleapis.com/v1beta1/", userBasePath, params), nil
}

// variableApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type variableApiOperation interface {
	do(context.Context, *Variable, *Client) error
}

// newUpdateVariableUpdateRequest creates a request for an
// Variable resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateVariableUpdateRequest(ctx context.Context, f *Variable, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Text; !dcl.IsEmptyValueIndirect(v) {
		req["text"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		req["value"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/configs/%s/variables/%s", *f.Project, *f.RuntimeConfig, *f.Name)

	return req, nil
}

// marshalUpdateVariableUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateVariableUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateVariableUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateVariableUpdateOperation) do(ctx context.Context, r *Variable, c *Client) error {
	_, err := c.GetVariable(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateVariableUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateVariableUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listVariableRaw(ctx context.Context, project, runtimeConfig, pageToken string, pageSize int32) ([]byte, error) {
	u, err := variableListURL(c.Config.BasePath, project, runtimeConfig)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != VariableMaxPage {
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

type listVariableOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listVariable(ctx context.Context, project, runtimeConfig, pageToken string, pageSize int32) ([]*Variable, string, error) {
	b, err := c.listVariableRaw(ctx, project, runtimeConfig, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listVariableOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Variable
	for _, v := range m.Instances {
		res, err := unmarshalMapVariable(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.RuntimeConfig = &runtimeConfig
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllVariable(ctx context.Context, f func(*Variable) bool, resources []*Variable) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteVariable(ctx, res)
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

type deleteVariableOperation struct{}

func (op *deleteVariableOperation) do(ctx context.Context, r *Variable, c *Client) error {
	r, err := c.GetVariable(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Variable not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetVariable checking for existence. error: %v", err)
		return err
	}

	u, err := variableDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Variable: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetVariable(ctx, r.URLNormalized())
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
type createVariableOperation struct {
	response map[string]interface{}
}

func (op *createVariableOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createVariableOperation) do(ctx context.Context, r *Variable, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, runtimeConfig := r.createFields()
	u, err := variableCreateURL(c.Config.BasePath, project, runtimeConfig)

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

	if _, err := c.GetVariable(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getVariableRaw(ctx context.Context, r *Variable) ([]byte, error) {

	u, err := variableGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) variableDiffsForRawDesired(ctx context.Context, rawDesired *Variable, opts ...dcl.ApplyOption) (initial, desired *Variable, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Variable
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Variable); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Variable, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetVariable(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Variable resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Variable resource: %v", err)
		}
		c.Config.Logger.Info("Found that Variable resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeVariableDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Variable: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Variable: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeVariableInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Variable: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeVariableDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Variable: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffVariable(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeVariableInitialState(rawInitial, rawDesired *Variable) (*Variable, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.Text) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Value) {
			rawInitial.Text = dcl.String("")
		}
	}

	if dcl.IsZeroValue(rawInitial.Value) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Text) {
			rawInitial.Value = dcl.String("")
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeVariableDesiredState(rawDesired, rawInitial *Variable, opts ...dcl.ApplyOption) (*Variable, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if rawDesired.Text != nil || rawInitial.Text != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Value) {
			rawDesired.Text = nil
			rawInitial.Text = nil
		}
	}

	if rawDesired.Value != nil || rawInitial.Value != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Text) {
			rawDesired.Value = nil
			rawInitial.Value = nil
		}
	}

	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.RuntimeConfig, rawInitial.RuntimeConfig) {
		rawDesired.RuntimeConfig = rawInitial.RuntimeConfig
	}
	if dcl.StringCanonicalize(rawDesired.Text, rawInitial.Text) {
		rawDesired.Text = rawInitial.Text
	}
	if dcl.StringCanonicalize(rawDesired.Value, rawInitial.Value) {
		rawDesired.Value = rawInitial.Value
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeVariableNewState(c *Client, rawNew, rawDesired *Variable) (*Variable, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.RuntimeConfig = rawDesired.RuntimeConfig

	if dcl.IsEmptyValueIndirect(rawNew.Text) && dcl.IsEmptyValueIndirect(rawDesired.Text) {
		rawNew.Text = rawDesired.Text
	} else {
		if dcl.StringCanonicalize(rawDesired.Text, rawNew.Text) {
			rawNew.Text = rawDesired.Text
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Value) && dcl.IsEmptyValueIndirect(rawDesired.Value) {
		rawNew.Value = rawDesired.Value
	} else {
		if dcl.StringCanonicalize(rawDesired.Value, rawNew.Value) {
			rawNew.Value = rawDesired.Value
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
		if dcl.StringCanonicalize(rawDesired.UpdateTime, rawNew.UpdateTime) {
			rawNew.UpdateTime = rawDesired.UpdateTime
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffVariable(c *Client, desired, actual *Variable, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.RuntimeConfig, actual.RuntimeConfig, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RuntimeConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Text, actual.Text, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVariableUpdateOperation")}, fn.AddNest("Text")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.TriggersOperation("updateVariableUpdateOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *Variable) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.RuntimeConfig), dcl.ValueOrEmptyString(n.Name)
}

func (r *Variable) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.RuntimeConfig)
}

func (r *Variable) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.RuntimeConfig), dcl.ValueOrEmptyString(n.Name)
}

func (r *Variable) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project":       dcl.ValueOrEmptyString(n.Project),
			"runtimeConfig": dcl.ValueOrEmptyString(n.RuntimeConfig),
			"name":          dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/configs/{{runtimeConfig}}/variables/{{name}}", "https://runtimeconfig.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Variable resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Variable) marshal(c *Client) ([]byte, error) {
	m, err := expandVariable(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Variable: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalVariable decodes JSON responses into the Variable resource schema.
func unmarshalVariable(b []byte, c *Client) (*Variable, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapVariable(m, c)
}

func unmarshalMapVariable(m map[string]interface{}, c *Client) (*Variable, error) {

	return flattenVariable(c, m), nil
}

// expandVariable expands Variable into a JSON request object.
func expandVariable(c *Client, f *Variable) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveFromPattern("projects/%s/configs/%s/variables/%s", f.Name, f.Project, f.RuntimeConfig, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding RuntimeConfig into runtimeConfig: %w", err)
	} else if v != nil {
		m["runtimeConfig"] = v
	}
	if v := f.Text; !dcl.IsEmptyValueIndirect(v) {
		m["text"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenVariable flattens Variable from a JSON request object into the
// Variable type.
func flattenVariable(c *Client, i interface{}) *Variable {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Variable{}
	res.Name = dcl.FlattenString(m["name"])
	res.RuntimeConfig = dcl.FlattenString(m["runtimeConfig"])
	res.Text = dcl.FlattenString(m["text"])
	res.Value = dcl.FlattenString(m["value"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Variable) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalVariable(b, c)
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
		if nr.RuntimeConfig == nil && ncr.RuntimeConfig == nil {
			c.Config.Logger.Info("Both RuntimeConfig fields null - considering equal.")
		} else if nr.RuntimeConfig == nil || ncr.RuntimeConfig == nil {
			c.Config.Logger.Info("Only one RuntimeConfig field is null - considering unequal.")
			return false
		} else if *nr.RuntimeConfig != *ncr.RuntimeConfig {
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

type variableDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         variableApiOperation
}

func convertFieldDiffToVariableOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]variableDiff, error) {
	var diffs []variableDiff
	for _, op := range ops {
		diff := variableDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTovariableApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTovariableApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (variableApiOperation, error) {
	switch op {

	case "updateVariableUpdateOperation":
		return &updateVariableUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
