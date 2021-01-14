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
package compute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mohae/deepcopy"
	"io/ioutil"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
	"reflect"
	"strings"
)

func (r *HttpHealthCheck) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func httpHealthCheckGetURL(userBasePath string, r *HttpHealthCheck) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func httpHealthCheckListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func httpHealthCheckCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func httpHealthCheckDeleteURL(userBasePath string, r *HttpHealthCheck) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// httpHealthCheckApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type httpHealthCheckApiOperation interface {
	do(context.Context, *HttpHealthCheck, *Client) error
}

// newUpdateHttpHealthCheckUpdateRequest creates a request for an
// HttpHealthCheck resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateHttpHealthCheckUpdateRequest(ctx context.Context, f *HttpHealthCheck, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.CheckIntervalSec; !dcl.IsEmptyValueIndirect(v) {
		req["checkIntervalSec"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.HealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		req["healthyThreshold"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		req["host"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		req["port"] = v
	}
	if v := f.RequestPath; !dcl.IsEmptyValueIndirect(v) {
		req["requestPath"] = v
	}
	if v := f.TimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		req["timeoutSec"] = v
	}
	if v := f.UnhealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		req["unhealthyThreshold"] = v
	}
	return req, nil
}

// marshalUpdateHttpHealthCheckUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateHttpHealthCheckUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateHttpHealthCheckUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateHttpHealthCheckUpdateOperation) do(ctx context.Context, r *HttpHealthCheck, c *Client) error {
	_, err := c.GetHttpHealthCheck(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateHttpHealthCheckUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateHttpHealthCheckUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listHttpHealthCheckRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := httpHealthCheckListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != HttpHealthCheckMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listHttpHealthCheckOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listHttpHealthCheck(ctx context.Context, project, pageToken string, pageSize int32) ([]*HttpHealthCheck, string, error) {
	b, err := c.listHttpHealthCheckRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listHttpHealthCheckOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*HttpHealthCheck
	for _, v := range m.Items {
		res := flattenHttpHealthCheck(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllHttpHealthCheck(ctx context.Context, f func(*HttpHealthCheck) bool, resources []*HttpHealthCheck) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteHttpHealthCheck(ctx, res)
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

type deleteHttpHealthCheckOperation struct{}

func (op *deleteHttpHealthCheckOperation) do(ctx context.Context, r *HttpHealthCheck, c *Client) error {

	_, err := c.GetHttpHealthCheck(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("HttpHealthCheck not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetHttpHealthCheck checking for existence. error: %v", err)
		return err
	}

	u, err := httpHealthCheckDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}
	_, err = c.GetHttpHealthCheck(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createHttpHealthCheckOperation struct{}

func (op *createHttpHealthCheckOperation) do(ctx context.Context, r *HttpHealthCheck, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := httpHealthCheckCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	if _, err := c.GetHttpHealthCheck(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getHttpHealthCheckRaw(ctx context.Context, r *HttpHealthCheck) ([]byte, error) {
	if dcl.IsZeroValue(r.CheckIntervalSec) {
		r.CheckIntervalSec = dcl.Int64(5)
	}
	if dcl.IsZeroValue(r.HealthyThreshold) {
		r.HealthyThreshold = dcl.Int64(2)
	}
	if dcl.IsZeroValue(r.Port) {
		r.Port = dcl.Int64(80)
	}
	if dcl.IsZeroValue(r.RequestPath) {
		r.RequestPath = dcl.String("/")
	}
	if dcl.IsZeroValue(r.TimeoutSec) {
		r.TimeoutSec = dcl.Int64(5)
	}
	if dcl.IsZeroValue(r.UnhealthyThreshold) {
		r.UnhealthyThreshold = dcl.Int64(2)
	}

	u, err := httpHealthCheckGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
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

func (c *Client) httpHealthCheckDiffsForRawDesired(ctx context.Context, rawDesired *HttpHealthCheck, opts ...dcl.ApplyOption) (initial, desired *HttpHealthCheck, diffs []httpHealthCheckDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *HttpHealthCheck
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HttpHealthCheck); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected HttpHealthCheck, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetHttpHealthCheck(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a HttpHealthCheck resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve HttpHealthCheck resource: %v", err)
		}

		c.Config.Logger.Info("Found that HttpHealthCheck resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeHttpHealthCheckDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for HttpHealthCheck: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for HttpHealthCheck: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeHttpHealthCheckInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for HttpHealthCheck: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeHttpHealthCheckDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for HttpHealthCheck: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffHttpHealthCheck(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeHttpHealthCheckInitialState(rawInitial, rawDesired *HttpHealthCheck) (*HttpHealthCheck, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeHttpHealthCheckDesiredState(rawDesired, rawInitial *HttpHealthCheck, opts ...dcl.ApplyOption) (*HttpHealthCheck, error) {

	if dcl.IsZeroValue(rawDesired.CheckIntervalSec) {
		rawDesired.CheckIntervalSec = dcl.Int64(5)
	}

	if dcl.IsZeroValue(rawDesired.HealthyThreshold) {
		rawDesired.HealthyThreshold = dcl.Int64(2)
	}

	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = dcl.Int64(80)
	}

	if dcl.IsZeroValue(rawDesired.RequestPath) {
		rawDesired.RequestPath = dcl.String("/")
	}

	if dcl.IsZeroValue(rawDesired.TimeoutSec) {
		rawDesired.TimeoutSec = dcl.Int64(5)
	}

	if dcl.IsZeroValue(rawDesired.UnhealthyThreshold) {
		rawDesired.UnhealthyThreshold = dcl.Int64(2)
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HttpHealthCheck); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected HttpHealthCheck, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.CheckIntervalSec) {
		rawDesired.CheckIntervalSec = rawInitial.CheckIntervalSec
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.HealthyThreshold) {
		rawDesired.HealthyThreshold = rawInitial.HealthyThreshold
	}
	if dcl.IsZeroValue(rawDesired.Host) {
		rawDesired.Host = rawInitial.Host
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = rawInitial.Port
	}
	if dcl.IsZeroValue(rawDesired.RequestPath) {
		rawDesired.RequestPath = rawInitial.RequestPath
	}
	if dcl.IsZeroValue(rawDesired.TimeoutSec) {
		rawDesired.TimeoutSec = rawInitial.TimeoutSec
	}
	if dcl.IsZeroValue(rawDesired.UnhealthyThreshold) {
		rawDesired.UnhealthyThreshold = rawInitial.UnhealthyThreshold
	}
	if dcl.IsZeroValue(rawDesired.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}

	return rawDesired, nil
}

func canonicalizeHttpHealthCheckNewState(c *Client, rawNew, rawDesired *HttpHealthCheck) (*HttpHealthCheck, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CheckIntervalSec) && dcl.IsEmptyValueIndirect(rawDesired.CheckIntervalSec) {
		rawNew.CheckIntervalSec = rawDesired.CheckIntervalSec
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthyThreshold) && dcl.IsEmptyValueIndirect(rawDesired.HealthyThreshold) {
		rawNew.HealthyThreshold = rawDesired.HealthyThreshold
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Host) && dcl.IsEmptyValueIndirect(rawDesired.Host) {
		rawNew.Host = rawDesired.Host
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RequestPath) && dcl.IsEmptyValueIndirect(rawDesired.RequestPath) {
		rawNew.RequestPath = rawDesired.RequestPath
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeoutSec) && dcl.IsEmptyValueIndirect(rawDesired.TimeoutSec) {
		rawNew.TimeoutSec = rawDesired.TimeoutSec
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UnhealthyThreshold) && dcl.IsEmptyValueIndirect(rawDesired.UnhealthyThreshold) {
		rawNew.UnhealthyThreshold = rawDesired.UnhealthyThreshold
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
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
	}

	return rawNew, nil
}

type httpHealthCheckDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         httpHealthCheckApiOperation
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
func diffHttpHealthCheck(c *Client, desired, actual *HttpHealthCheck, opts ...dcl.ApplyOption) ([]httpHealthCheckDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []httpHealthCheckDiff
	if !dcl.IsZeroValue(desired.CheckIntervalSec) && (dcl.IsZeroValue(actual.CheckIntervalSec) || !reflect.DeepEqual(*desired.CheckIntervalSec, *actual.CheckIntervalSec)) {
		c.Config.Logger.Infof("Detected diff in CheckIntervalSec.\nDESIRED: %#v\nACTUAL: %#v", desired.CheckIntervalSec, actual.CheckIntervalSec)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "CheckIntervalSec",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %#v\nACTUAL: %#v", desired.Description, actual.Description)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "Description",
		})

	}
	if !dcl.IsZeroValue(desired.HealthyThreshold) && (dcl.IsZeroValue(actual.HealthyThreshold) || !reflect.DeepEqual(*desired.HealthyThreshold, *actual.HealthyThreshold)) {
		c.Config.Logger.Infof("Detected diff in HealthyThreshold.\nDESIRED: %#v\nACTUAL: %#v", desired.HealthyThreshold, actual.HealthyThreshold)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "HealthyThreshold",
		})

	}
	if !dcl.IsZeroValue(desired.Host) && (dcl.IsZeroValue(actual.Host) || !reflect.DeepEqual(*desired.Host, *actual.Host)) {
		c.Config.Logger.Infof("Detected diff in Host.\nDESIRED: %#v\nACTUAL: %#v", desired.Host, actual.Host)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "Host",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %#v\nACTUAL: %#v", desired.Name, actual.Name)
		diffs = append(diffs, httpHealthCheckDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Port) && (dcl.IsZeroValue(actual.Port) || !reflect.DeepEqual(*desired.Port, *actual.Port)) {
		c.Config.Logger.Infof("Detected diff in Port.\nDESIRED: %#v\nACTUAL: %#v", desired.Port, actual.Port)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "Port",
		})

	}
	if !dcl.IsZeroValue(desired.RequestPath) && (dcl.IsZeroValue(actual.RequestPath) || !reflect.DeepEqual(*desired.RequestPath, *actual.RequestPath)) {
		c.Config.Logger.Infof("Detected diff in RequestPath.\nDESIRED: %#v\nACTUAL: %#v", desired.RequestPath, actual.RequestPath)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "RequestPath",
		})

	}
	if !dcl.IsZeroValue(desired.TimeoutSec) && (dcl.IsZeroValue(actual.TimeoutSec) || !reflect.DeepEqual(*desired.TimeoutSec, *actual.TimeoutSec)) {
		c.Config.Logger.Infof("Detected diff in TimeoutSec.\nDESIRED: %#v\nACTUAL: %#v", desired.TimeoutSec, actual.TimeoutSec)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "TimeoutSec",
		})

	}
	if !dcl.IsZeroValue(desired.UnhealthyThreshold) && (dcl.IsZeroValue(actual.UnhealthyThreshold) || !reflect.DeepEqual(*desired.UnhealthyThreshold, *actual.UnhealthyThreshold)) {
		c.Config.Logger.Infof("Detected diff in UnhealthyThreshold.\nDESIRED: %#v\nACTUAL: %#v", desired.UnhealthyThreshold, actual.UnhealthyThreshold)

		diffs = append(diffs, httpHealthCheckDiff{
			UpdateOp:  &updateHttpHealthCheckUpdateOperation{},
			FieldName: "UnhealthyThreshold",
		})

	}
	if !dcl.IsZeroValue(desired.Project) && !dcl.NameToSelfLink(desired.Project, actual.Project) {
		c.Config.Logger.Infof("Detected diff in Project.\nDESIRED: %#v\nACTUAL: %#v", desired.Project, actual.Project)
		diffs = append(diffs, httpHealthCheckDiff{
			RequiresRecreate: true,
			FieldName:        "Project",
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
	var deduped []httpHealthCheckDiff
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
func (r *HttpHealthCheck) urlNormalized() *HttpHealthCheck {
	normalized := deepcopy.Copy(*r).(HttpHealthCheck)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *HttpHealthCheck) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpHealthCheck) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *HttpHealthCheck) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpHealthCheck) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/httpHealthChecks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the HttpHealthCheck resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *HttpHealthCheck) marshal(c *Client) ([]byte, error) {
	m, err := expandHttpHealthCheck(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling HttpHealthCheck: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalHttpHealthCheck decodes JSON responses into the HttpHealthCheck resource schema.
func unmarshalHttpHealthCheck(b []byte, c *Client) (*HttpHealthCheck, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenHttpHealthCheck(c, m), nil
}

// expandHttpHealthCheck expands HttpHealthCheck into a JSON request object.
func expandHttpHealthCheck(c *Client, f *HttpHealthCheck) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CheckIntervalSec; !dcl.IsEmptyValueIndirect(v) {
		m["checkIntervalSec"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.HealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["healthyThreshold"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.RequestPath; !dcl.IsEmptyValueIndirect(v) {
		m["requestPath"] = v
	}
	if v := f.TimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["timeoutSec"] = v
	}
	if v := f.UnhealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["unhealthyThreshold"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}

	return m, nil
}

// flattenHttpHealthCheck flattens HttpHealthCheck from a JSON request object into the
// HttpHealthCheck type.
func flattenHttpHealthCheck(c *Client, i interface{}) *HttpHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &HttpHealthCheck{}
	r.CheckIntervalSec = dcl.FlattenInteger(m["checkIntervalSec"])
	if _, ok := m["checkIntervalSec"]; !ok {
		c.Config.Logger.Info("Using default value for checkIntervalSec")
		r.CheckIntervalSec = dcl.Int64(5)
	}
	r.Description = dcl.FlattenString(m["description"])
	r.HealthyThreshold = dcl.FlattenInteger(m["healthyThreshold"])
	if _, ok := m["healthyThreshold"]; !ok {
		c.Config.Logger.Info("Using default value for healthyThreshold")
		r.HealthyThreshold = dcl.Int64(2)
	}
	r.Host = dcl.FlattenString(m["host"])
	r.Name = dcl.FlattenString(m["name"])
	r.Port = dcl.FlattenInteger(m["port"])
	if _, ok := m["port"]; !ok {
		c.Config.Logger.Info("Using default value for port")
		r.Port = dcl.Int64(80)
	}
	r.RequestPath = dcl.FlattenString(m["requestPath"])
	if _, ok := m["requestPath"]; !ok {
		c.Config.Logger.Info("Using default value for requestPath")
		r.RequestPath = dcl.String("/")
	}
	r.TimeoutSec = dcl.FlattenInteger(m["timeoutSec"])
	if _, ok := m["timeoutSec"]; !ok {
		c.Config.Logger.Info("Using default value for timeoutSec")
		r.TimeoutSec = dcl.Int64(5)
	}
	r.UnhealthyThreshold = dcl.FlattenInteger(m["unhealthyThreshold"])
	if _, ok := m["unhealthyThreshold"]; !ok {
		c.Config.Logger.Info("Using default value for unhealthyThreshold")
		r.UnhealthyThreshold = dcl.Int64(2)
	}
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.Project = dcl.FlattenString(m["project"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *HttpHealthCheck) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalHttpHealthCheck(b, c)
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
