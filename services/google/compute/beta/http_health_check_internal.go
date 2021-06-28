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
	return dcl.URL("projects/{{project}}/global/httpHealthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func httpHealthCheckListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func httpHealthCheckCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func httpHealthCheckDeleteURL(userBasePath string, r *HttpHealthCheck) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/httpHealthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
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
	Diffs        []*dcl.FieldDiff
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
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
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
		res, err := unmarshalMapHttpHealthCheck(v, c)
		if err != nil {
			return nil, m.Token, err
		}
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
		_, err = c.GetHttpHealthCheck(ctx, r.urlNormalized())
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
type createHttpHealthCheckOperation struct {
	response map[string]interface{}
}

func (op *createHttpHealthCheckOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

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

	if _, err := c.GetHttpHealthCheck(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
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

func (c *Client) httpHealthCheckDiffsForRawDesired(ctx context.Context, rawDesired *HttpHealthCheck, opts ...dcl.ApplyOption) (initial, desired *HttpHealthCheck, diffs []*dcl.FieldDiff, err error) {
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
	fmt.Printf("newDiffs: %v\n", diffs)
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

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.IsZeroValue(rawDesired.CheckIntervalSec) {
		rawDesired.CheckIntervalSec = rawInitial.CheckIntervalSec
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.HealthyThreshold) {
		rawDesired.HealthyThreshold = rawInitial.HealthyThreshold
	}
	if dcl.StringCanonicalize(rawDesired.Host, rawInitial.Host) {
		rawDesired.Host = rawInitial.Host
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = rawInitial.Port
	}
	if dcl.StringCanonicalize(rawDesired.RequestPath, rawInitial.RequestPath) {
		rawDesired.RequestPath = rawInitial.RequestPath
	}
	if dcl.IsZeroValue(rawDesired.TimeoutSec) {
		rawDesired.TimeoutSec = rawInitial.TimeoutSec
	}
	if dcl.IsZeroValue(rawDesired.UnhealthyThreshold) {
		rawDesired.UnhealthyThreshold = rawInitial.UnhealthyThreshold
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
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
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthyThreshold) && dcl.IsEmptyValueIndirect(rawDesired.HealthyThreshold) {
		rawNew.HealthyThreshold = rawDesired.HealthyThreshold
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Host) && dcl.IsEmptyValueIndirect(rawDesired.Host) {
		rawNew.Host = rawDesired.Host
	} else {
		if dcl.StringCanonicalize(rawDesired.Host, rawNew.Host) {
			rawNew.Host = rawDesired.Host
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RequestPath) && dcl.IsEmptyValueIndirect(rawDesired.RequestPath) {
		rawNew.RequestPath = rawDesired.RequestPath
	} else {
		if dcl.StringCanonicalize(rawDesired.RequestPath, rawNew.RequestPath) {
			rawNew.RequestPath = rawDesired.RequestPath
		}
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
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
		}
	}

	rawNew.Project = rawDesired.Project

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
func diffHttpHealthCheck(c *Client, desired, actual *HttpHealthCheck, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.CheckIntervalSec, actual.CheckIntervalSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("CheckIntervalSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthyThreshold, actual.HealthyThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("HealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestPath, actual.RequestPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("RequestPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSec, actual.TimeoutSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("TimeoutSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnhealthyThreshold, actual.UnhealthyThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpHealthCheckUpdateOperation")}, fn.AddNest("UnhealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *HttpHealthCheck) urlNormalized() *HttpHealthCheck {
	normalized := dcl.Copy(*r).(HttpHealthCheck)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Host = dcl.SelfLinkToName(r.Host)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.RequestPath = dcl.SelfLinkToName(r.RequestPath)
	normalized.CreationTimestamp = dcl.SelfLinkToName(r.CreationTimestamp)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
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
		return dcl.URL("projects/{{project}}/global/httpHealthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

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
	return unmarshalMapHttpHealthCheck(m, c)
}

func unmarshalMapHttpHealthCheck(m map[string]interface{}, c *Client) (*HttpHealthCheck, error) {

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
	} else if v != nil {
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

	res := &HttpHealthCheck{}
	res.CheckIntervalSec = dcl.FlattenInteger(m["checkIntervalSec"])
	if _, ok := m["checkIntervalSec"]; !ok {
		c.Config.Logger.Info("Using default value for checkIntervalSec")
		res.CheckIntervalSec = dcl.Int64(5)
	}
	res.Description = dcl.FlattenString(m["description"])
	res.HealthyThreshold = dcl.FlattenInteger(m["healthyThreshold"])
	if _, ok := m["healthyThreshold"]; !ok {
		c.Config.Logger.Info("Using default value for healthyThreshold")
		res.HealthyThreshold = dcl.Int64(2)
	}
	res.Host = dcl.FlattenString(m["host"])
	res.Name = dcl.FlattenString(m["name"])
	res.Port = dcl.FlattenInteger(m["port"])
	if _, ok := m["port"]; !ok {
		c.Config.Logger.Info("Using default value for port")
		res.Port = dcl.Int64(80)
	}
	res.RequestPath = dcl.FlattenString(m["requestPath"])
	if _, ok := m["requestPath"]; !ok {
		c.Config.Logger.Info("Using default value for requestPath")
		res.RequestPath = dcl.String("/")
	}
	res.TimeoutSec = dcl.FlattenInteger(m["timeoutSec"])
	if _, ok := m["timeoutSec"]; !ok {
		c.Config.Logger.Info("Using default value for timeoutSec")
		res.TimeoutSec = dcl.Int64(5)
	}
	res.UnhealthyThreshold = dcl.FlattenInteger(m["unhealthyThreshold"])
	if _, ok := m["unhealthyThreshold"]; !ok {
		c.Config.Logger.Info("Using default value for unhealthyThreshold")
		res.UnhealthyThreshold = dcl.Int64(2)
	}
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Project = dcl.FlattenString(m["project"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])

	return res
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

type httpHealthCheckDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         httpHealthCheckApiOperation
}

func convertFieldDiffToHttpHealthCheckOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]httpHealthCheckDiff, error) {
	var diffs []httpHealthCheckDiff
	for _, op := range ops {
		diff := httpHealthCheckDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTohttpHealthCheckApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTohttpHealthCheckApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (httpHealthCheckApiOperation, error) {
	switch op {

	case "updateHttpHealthCheckUpdateOperation":
		return &updateHttpHealthCheckUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
