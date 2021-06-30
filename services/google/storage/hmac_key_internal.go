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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *HmacKey) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serviceAccountEmail"); err != nil {
		return err
	}
	return nil
}

func hmacKeyGetURL(userBasePath string, r *HmacKey) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/hmacKeys/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

func hmacKeyListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/hmacKeys", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func hmacKeyCreateURL(userBasePath, project, serviceAccountEmail string) (string, error) {
	params := map[string]interface{}{
		"project":             project,
		"serviceAccountEmail": serviceAccountEmail,
	}
	return dcl.URL("projects/{{project}}/hmacKeys?serviceAccountEmail={{serviceAccountEmail}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil

}

func hmacKeyDeleteURL(userBasePath string, r *HmacKey) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/hmacKeys/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, params), nil
}

// hmacKeyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type hmacKeyApiOperation interface {
	do(context.Context, *HmacKey, *Client) error
}

// newUpdateHmacKeyUpdateRequest creates a request for an
// HmacKey resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateHmacKeyUpdateRequest(ctx context.Context, f *HmacKey, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		req["state"] = v
	}
	return req, nil
}

// marshalUpdateHmacKeyUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateHmacKeyUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateHmacKeyUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateHmacKeyUpdateOperation) do(ctx context.Context, r *HmacKey, c *Client) error {
	_, err := c.GetHmacKey(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateHmacKeyUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateHmacKeyUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listHmacKeyRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := hmacKeyListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != HmacKeyMaxPage {
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

type listHmacKeyOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listHmacKey(ctx context.Context, project, pageToken string, pageSize int32) ([]*HmacKey, string, error) {
	b, err := c.listHmacKeyRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listHmacKeyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*HmacKey
	for _, v := range m.Items {
		res, err := unmarshalMapHmacKey(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllHmacKey(ctx context.Context, f func(*HmacKey) bool, resources []*HmacKey) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteHmacKey(ctx, res)
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

type deleteHmacKeyOperation struct{}

func (op *deleteHmacKeyOperation) do(ctx context.Context, r *HmacKey, c *Client) error {
	r, err := c.GetHmacKey(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("HmacKey not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetHmacKey checking for existence. error: %v", err)
		return err
	}

	if hmacKeyDeletePrecondition(r) {
		return nil
	}

	err = r.HmacKeyEnsureStateInactive(ctx, c)
	if err != nil {
		return err
	}
	u, err := hmacKeyDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete HmacKey: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createHmacKeyOperation struct {
	response map[string]interface{}
}

func (op *createHmacKeyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getHmacKeyRaw(ctx context.Context, r *HmacKey) ([]byte, error) {
	if dcl.IsZeroValue(r.State) {
		r.State = HmacKeyStateEnumRef("ACTIVE")
	}

	u, err := hmacKeyGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) hmacKeyDiffsForRawDesired(ctx context.Context, rawDesired *HmacKey, opts ...dcl.ApplyOption) (initial, desired *HmacKey, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *HmacKey
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HmacKey); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected HmacKey, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeHmacKeyDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetHmacKey(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a HmacKey resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve HmacKey resource: %v", err)
		}
		c.Config.Logger.Info("Found that HmacKey resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeHmacKeyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for HmacKey: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for HmacKey: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeHmacKeyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for HmacKey: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeHmacKeyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for HmacKey: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffHmacKey(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeHmacKeyInitialState(rawInitial, rawDesired *HmacKey) (*HmacKey, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeHmacKeyDesiredState(rawDesired, rawInitial *HmacKey, opts ...dcl.ApplyOption) (*HmacKey, error) {

	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = HmacKeyStateEnumRef("ACTIVE")
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.ServiceAccountEmail, rawInitial.ServiceAccountEmail) {
		rawDesired.ServiceAccountEmail = rawInitial.ServiceAccountEmail
	}

	return rawDesired, nil
}

func canonicalizeHmacKeyNewState(c *Client, rawNew, rawDesired *HmacKey) (*HmacKey, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeCreated) && dcl.IsEmptyValueIndirect(rawDesired.TimeCreated) {
		rawNew.TimeCreated = rawDesired.TimeCreated
	} else {
		if dcl.StringCanonicalize(rawDesired.TimeCreated, rawNew.TimeCreated) {
			rawNew.TimeCreated = rawDesired.TimeCreated
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Updated) && dcl.IsEmptyValueIndirect(rawDesired.Updated) {
		rawNew.Updated = rawDesired.Updated
	} else {
		if dcl.StringCanonicalize(rawDesired.Updated, rawNew.Updated) {
			rawNew.Updated = rawDesired.Updated
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Secret) && dcl.IsEmptyValueIndirect(rawDesired.Secret) {
		rawNew.Secret = rawDesired.Secret
	} else {
		if dcl.StringCanonicalize(rawDesired.Secret, rawNew.Secret) {
			rawNew.Secret = rawDesired.Secret
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmail) {
		rawNew.ServiceAccountEmail = rawDesired.ServiceAccountEmail
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceAccountEmail, rawNew.ServiceAccountEmail) {
			rawNew.ServiceAccountEmail = rawDesired.ServiceAccountEmail
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
func diffHmacKey(c *Client, desired, actual *HmacKey, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AccessId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeCreated, actual.TimeCreated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeCreated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Updated, actual.Updated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Updated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHmacKeyUpdateOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *HmacKey) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *HmacKey) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.ServiceAccountEmail)
}

func (r *HmacKey) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *HmacKey) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/hmacKeys/{{name}}", "https://www.googleapis.com/storage/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the HmacKey resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *HmacKey) marshal(c *Client) ([]byte, error) {
	m, err := expandHmacKey(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling HmacKey: %w", err)
	}
	m = EncodeStorageHmacKeyCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalHmacKey decodes JSON responses into the HmacKey resource schema.
func unmarshalHmacKey(b []byte, c *Client) (*HmacKey, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapHmacKey(m, c)
}

func unmarshalMapHmacKey(m map[string]interface{}, c *Client) (*HmacKey, error) {

	return flattenHmacKey(c, m), nil
}

// expandHmacKey expands HmacKey into a JSON request object.
func expandHmacKey(c *Client, f *HmacKey) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["accessId"] = v
	}
	if v := f.TimeCreated; !dcl.IsEmptyValueIndirect(v) {
		m["timeCreated"] = v
	}
	if v := f.Updated; !dcl.IsEmptyValueIndirect(v) {
		m["updated"] = v
	}
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}

	return m, nil
}

// flattenHmacKey flattens HmacKey from a JSON request object into the
// HmacKey type.
func flattenHmacKey(c *Client, i interface{}) *HmacKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &HmacKey{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["accessId"]))
	res.TimeCreated = dcl.FlattenString(m["timeCreated"])
	res.Updated = dcl.FlattenString(m["updated"])
	res.Secret = dcl.FlattenString(m["secret"])
	res.State = flattenHmacKeyStateEnum(m["state"])
	if _, ok := m["state"]; !ok {
		c.Config.Logger.Info("Using default value for state")
		res.State = HmacKeyStateEnumRef("ACTIVE")
	}
	res.Project = dcl.FlattenString(m["project"])
	res.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])

	return res
}

// flattenHmacKeyStateEnumSlice flattens the contents of HmacKeyStateEnum from a JSON
// response object.
func flattenHmacKeyStateEnumSlice(c *Client, i interface{}) []HmacKeyStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HmacKeyStateEnum{}
	}

	if len(a) == 0 {
		return []HmacKeyStateEnum{}
	}

	items := make([]HmacKeyStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHmacKeyStateEnum(item.(interface{})))
	}

	return items
}

// flattenHmacKeyStateEnum asserts that an interface is a string, and returns a
// pointer to a *HmacKeyStateEnum with the same value as that string.
func flattenHmacKeyStateEnum(i interface{}) *HmacKeyStateEnum {
	s, ok := i.(string)
	if !ok {
		return HmacKeyStateEnumRef("")
	}

	return HmacKeyStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *HmacKey) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalHmacKey(b, c)
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

type hmacKeyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         hmacKeyApiOperation
}

func convertFieldDiffToHmacKeyOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]hmacKeyDiff, error) {
	var diffs []hmacKeyDiff
	for _, op := range ops {
		diff := hmacKeyDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTohmacKeyApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTohmacKeyApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (hmacKeyApiOperation, error) {
	switch op {

	case "updateHmacKeyUpdateOperation":
		return &updateHmacKeyUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
