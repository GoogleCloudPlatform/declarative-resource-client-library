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
package spanner

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

func (r *Instance) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "config"); err != nil {
		return err
	}
	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://spanner.googleapis.com/v1/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances", "https://spanner.googleapis.com/v1/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances", "https://spanner.googleapis.com/v1/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://spanner.googleapis.com/v1/", userBasePath, params), nil
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceUpdateInstanceRequest creates a request for an
// Instance resource's UpdateInstance update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateInstanceRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("projects/%s/instanceConfigs/%s", f.Config, f.Project, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["config"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		req["nodeCount"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/instances/%s", *f.Project, *f.Name)

	return req, nil
}

// marshalUpdateInstanceUpdateInstanceRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateInstanceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	m = EncodeSpannerInstanceUpdateRequest(m)
	return json.Marshal(m)
}

type updateInstanceUpdateInstanceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateInstance")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateInstanceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateInstanceRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://spanner.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceMaxPage {
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

type listInstanceOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, project, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Instance
	for _, v := range m.Instances {
		res, err := unmarshalMapInstance(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstance(ctx context.Context, f func(*Instance) bool, resources []*Instance) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstance(ctx, res)
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

type deleteInstanceOperation struct{}

func (op *deleteInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	r, err := c.GetInstance(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Instance not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstance checking for existence. error: %v", err)
		return err
	}

	u, err := instanceDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Instance: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetInstance(ctx, r.URLNormalized())
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
type createInstanceOperation struct {
	response map[string]interface{}
}

func (op *createInstanceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := instanceCreateURL(c.Config.BasePath, project)

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
	if err := o.Wait(ctx, c.Config, "https://spanner.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetInstance(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

	u, err := instanceGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) instanceDiffsForRawDesired(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (initial, desired *Instance, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Instance
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Instance); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Instance, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstance(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Instance resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Instance resource: %v", err)
		}
		c.Config.Logger.Info("Found that Instance resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Instance: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Instance: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Instance: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Instance: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstance(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceInitialState(rawInitial, rawDesired *Instance) (*Instance, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceDesiredState(rawDesired, rawInitial *Instance, opts ...dcl.ApplyOption) (*Instance, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Instance{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Config, rawInitial.Config) {
		canonicalDesired.Config = rawInitial.Config
	} else {
		canonicalDesired.Config = rawDesired.Config
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.IsZeroValue(rawDesired.NodeCount) {
		canonicalDesired.NodeCount = rawInitial.NodeCount
	} else {
		canonicalDesired.NodeCount = rawDesired.NodeCount
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}

	return canonicalDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Config, rawNew.Config) {
			rawNew.Config = rawDesired.Config
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodeCount) && dcl.IsEmptyValueIndirect(rawDesired.NodeCount) {
		rawNew.NodeCount = rawDesired.NodeCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
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
func diffInstance(c *Client, desired, actual *Instance, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeCount, actual.NodeCount, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("NodeCount")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *Instance) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Instance) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "UpdateInstance" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/instances/{{name}}", "https://spanner.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Instance resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Instance) marshal(c *Client) ([]byte, error) {
	m, err := expandInstance(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Instance: %w", err)
	}
	m = EncodeSpannerInstanceCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalInstance decodes JSON responses into the Instance resource schema.
func unmarshalInstance(b []byte, c *Client) (*Instance, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstance(m, c)
}

func unmarshalMapInstance(m map[string]interface{}, c *Client) (*Instance, error) {

	return flattenInstance(c, m), nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/instanceConfigs/%s", f.Config, f.Project, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if v != nil {
		m["config"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["nodeCount"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenInstance flattens Instance from a JSON request object into the
// Instance type.
func flattenInstance(c *Client, i interface{}) *Instance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Instance{}
	res.Name = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["project"])
	res.Config = dcl.FlattenString(m["config"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.NodeCount = dcl.FlattenInteger(m["nodeCount"])
	res.State = flattenInstanceStateEnum(m["state"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])

	return res
}

// flattenInstanceStateEnumMap flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumMap(c *Client, i interface{}) map[string]InstanceStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceStateEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceStateEnum{}
	}

	items := make(map[string]InstanceStateEnum)
	for k, item := range a {
		items[k] = *flattenInstanceStateEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceStateEnumSlice flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumSlice(c *Client, i interface{}) []InstanceStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceStateEnum{}
	}

	if len(a) == 0 {
		return []InstanceStateEnum{}
	}

	items := make([]InstanceStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceStateEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceStateEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceStateEnum with the same value as that string.
func flattenInstanceStateEnum(i interface{}) *InstanceStateEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceStateEnumRef("")
	}

	return InstanceStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Instance) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstance(b, c)
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

type instanceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceApiOperation
}

func convertFieldDiffsToInstanceDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]instanceDiff, error) {
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
	var diffs []instanceDiff
	// For each operation name, create a instanceDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := instanceDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToInstanceApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToInstanceApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (instanceApiOperation, error) {
	switch opName {

	case "updateInstanceUpdateInstanceOperation":
		return &updateInstanceUpdateInstanceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
