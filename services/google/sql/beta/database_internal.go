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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Database) validate() error {

	if err := dcl.Required(r, "instance"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}

func databaseGetURL(userBasePath string, r *Database) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"instance": dcl.ValueOrEmptyString(r.Instance),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/databases/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

func databaseListURL(userBasePath, project, instance string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"instance": instance,
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/databases", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func databaseCreateURL(userBasePath, project, instance string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"instance": instance,
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/databases", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func databaseDeleteURL(userBasePath string, r *Database) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"instance": dcl.ValueOrEmptyString(r.Instance),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{instance}}/databases/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

// databaseApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type databaseApiOperation interface {
	do(context.Context, *Database, *Client) error
}

// newUpdateDatabaseUpdateRequest creates a request for an
// Database resource's Update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDatabaseUpdateRequest(ctx context.Context, f *Database, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Charset; !dcl.IsEmptyValueIndirect(v) {
		req["charset"] = v
	}
	if v := f.Collation; !dcl.IsEmptyValueIndirect(v) {
		req["collation"] = v
	}
	return req, nil
}

// marshalUpdateDatabaseUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateDatabaseUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDatabaseUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDatabaseUpdateOperation) do(ctx context.Context, r *Database, c *Client) error {
	_, err := c.GetDatabase(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Update")
	if err != nil {
		return err
	}

	req, err := newUpdateDatabaseUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateDatabaseUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listDatabaseRaw(ctx context.Context, project, instance, pageToken string, pageSize int32) ([]byte, error) {
	u, err := databaseListURL(c.Config.BasePath, project, instance)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DatabaseMaxPage {
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

type listDatabaseOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listDatabase(ctx context.Context, project, instance, pageToken string, pageSize int32) ([]*Database, string, error) {
	b, err := c.listDatabaseRaw(ctx, project, instance, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDatabaseOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Database
	for _, v := range m.Items {
		res, err := unmarshalMapDatabase(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Instance = &instance
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDatabase(ctx context.Context, f func(*Database) bool, resources []*Database) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDatabase(ctx, res)
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

type deleteDatabaseOperation struct{}

func (op *deleteDatabaseOperation) do(ctx context.Context, r *Database, c *Client) error {
	r, err := c.GetDatabase(ctx, r.URLNormalized())
	if err != nil {
		if IsNotFoundOrInstanceNotRunning(err) {
			c.Config.Logger.Infof("Database not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetDatabase checking for existence. error: %v", err)
		return err
	}

	u, err := databaseDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetDatabase(ctx, r.URLNormalized())
		if !IsNotFoundOrInstanceNotRunning(err) {
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
type createDatabaseOperation struct {
	response map[string]interface{}
}

func (op *createDatabaseOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDatabaseOperation) do(ctx context.Context, r *Database, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, instance := r.createFields()
	u, err := databaseCreateURL(c.Config.BasePath, project, instance)

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
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetDatabase(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDatabaseRaw(ctx context.Context, r *Database) ([]byte, error) {

	u, err := databaseGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) databaseDiffsForRawDesired(ctx context.Context, rawDesired *Database, opts ...dcl.ApplyOption) (initial, desired *Database, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Database
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Database); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Database, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDatabase(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !IsNotFoundOrInstanceNotRunning(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Database resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Database resource: %v", err)
		}
		c.Config.Logger.Info("Found that Database resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDatabaseDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Database: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Database: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDatabaseInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Database: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDatabaseDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Database: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDatabase(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDatabaseInitialState(rawInitial, rawDesired *Database) (*Database, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDatabaseDesiredState(rawDesired, rawInitial *Database, opts ...dcl.ApplyOption) (*Database, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.Charset, rawInitial.Charset) {
		rawDesired.Charset = rawInitial.Charset
	}
	if dcl.StringCanonicalize(rawDesired.Collation, rawInitial.Collation) {
		rawDesired.Collation = rawInitial.Collation
	}
	if dcl.StringCanonicalize(rawDesired.Instance, rawInitial.Instance) {
		rawDesired.Instance = rawInitial.Instance
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeDatabaseNewState(c *Client, rawNew, rawDesired *Database) (*Database, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Charset) && dcl.IsEmptyValueIndirect(rawDesired.Charset) {
		rawNew.Charset = rawDesired.Charset
	} else {
		if dcl.StringCanonicalize(rawDesired.Charset, rawNew.Charset) {
			rawNew.Charset = rawDesired.Charset
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Collation) && dcl.IsEmptyValueIndirect(rawDesired.Collation) {
		rawNew.Collation = rawDesired.Collation
	} else {
		if dcl.StringCanonicalize(rawDesired.Collation, rawNew.Collation) {
			rawNew.Collation = rawDesired.Collation
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Instance) && dcl.IsEmptyValueIndirect(rawDesired.Instance) {
		rawNew.Instance = rawDesired.Instance
	} else {
		if dcl.StringCanonicalize(rawDesired.Instance, rawNew.Instance) {
			rawNew.Instance = rawDesired.Instance
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
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
func diffDatabase(c *Client, desired, actual *Database, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Charset, actual.Charset, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatabaseUpdateOperation")}, fn.AddNest("Charset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Collation, actual.Collation, dcl.Info{OperationSelector: dcl.TriggersOperation("updateDatabaseUpdateOperation")}, fn.AddNest("Collation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Instance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

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

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *Database) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name)
}

func (r *Database) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance)
}

func (r *Database) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Instance), dcl.ValueOrEmptyString(n.Name)
}

func (r *Database) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "Update" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"instance": dcl.ValueOrEmptyString(n.Instance),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/instances/{{instance}}/databases/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Database resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Database) marshal(c *Client) ([]byte, error) {
	m, err := expandDatabase(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Database: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDatabase decodes JSON responses into the Database resource schema.
func unmarshalDatabase(b []byte, c *Client) (*Database, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDatabase(m, c)
}

func unmarshalMapDatabase(m map[string]interface{}, c *Client) (*Database, error) {

	return flattenDatabase(c, m), nil
}

// expandDatabase expands Database into a JSON request object.
func expandDatabase(c *Client, f *Database) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Charset; !dcl.IsEmptyValueIndirect(v) {
		m["charset"] = v
	}
	if v := f.Collation; !dcl.IsEmptyValueIndirect(v) {
		m["collation"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}

	return m, nil
}

// flattenDatabase flattens Database from a JSON request object into the
// Database type.
func flattenDatabase(c *Client, i interface{}) *Database {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Database{}
	res.Charset = dcl.FlattenString(m["charset"])
	res.Collation = dcl.FlattenString(m["collation"])
	res.Instance = dcl.FlattenString(m["instance"])
	res.Name = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["project"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Database) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDatabase(b, c)
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
		if nr.Instance == nil && ncr.Instance == nil {
			c.Config.Logger.Info("Both Instance fields null - considering equal.")
		} else if nr.Instance == nil || ncr.Instance == nil {
			c.Config.Logger.Info("Only one Instance field is null - considering unequal.")
			return false
		} else if *nr.Instance != *ncr.Instance {
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

type databaseDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         databaseApiOperation
}

func convertFieldDiffToDatabaseOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]databaseDiff, error) {
	var diffs []databaseDiff
	for _, op := range ops {
		diff := databaseDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTodatabaseApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTodatabaseApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (databaseApiOperation, error) {
	switch op {

	case "updateDatabaseUpdateOperation":
		return &updateDatabaseUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
