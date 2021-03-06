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
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *HttpsHealthCheck) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func httpsHealthCheckGetURL(userBasePath string, r *HttpsHealthCheck) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/httpsHealthChecks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func httpsHealthCheckListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/httpsHealthChecks", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func httpsHealthCheckCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/httpsHealthChecks", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func httpsHealthCheckDeleteURL(userBasePath string, r *HttpsHealthCheck) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/httpsHealthChecks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// httpsHealthCheckApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type httpsHealthCheckApiOperation interface {
	do(context.Context, *HttpsHealthCheck, *Client) error
}

// newUpdateHttpsHealthCheckUpdateRequest creates a request for an
// HttpsHealthCheck resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateHttpsHealthCheckUpdateRequest(ctx context.Context, f *HttpsHealthCheck, c *Client) (map[string]interface{}, error) {
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

// marshalUpdateHttpsHealthCheckUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateHttpsHealthCheckUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateHttpsHealthCheckUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateHttpsHealthCheckUpdateOperation) do(ctx context.Context, r *HttpsHealthCheck, c *Client) error {
	_, err := c.GetHttpsHealthCheck(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateHttpsHealthCheckUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateHttpsHealthCheckUpdateRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listHttpsHealthCheckRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := httpsHealthCheckListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != HttpsHealthCheckMaxPage {
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

type listHttpsHealthCheckOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listHttpsHealthCheck(ctx context.Context, project, pageToken string, pageSize int32) ([]*HttpsHealthCheck, string, error) {
	b, err := c.listHttpsHealthCheckRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listHttpsHealthCheckOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*HttpsHealthCheck
	for _, v := range m.Items {
		res, err := unmarshalMapHttpsHealthCheck(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllHttpsHealthCheck(ctx context.Context, f func(*HttpsHealthCheck) bool, resources []*HttpsHealthCheck) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteHttpsHealthCheck(ctx, res)
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

type deleteHttpsHealthCheckOperation struct{}

func (op *deleteHttpsHealthCheckOperation) do(ctx context.Context, r *HttpsHealthCheck, c *Client) error {
	r, err := c.GetHttpsHealthCheck(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("HttpsHealthCheck not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetHttpsHealthCheck checking for existence. error: %v", err)
		return err
	}

	u, err := httpsHealthCheckDeleteURL(c.Config.BasePath, r.URLNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetHttpsHealthCheck(ctx, r.URLNormalized())
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
type createHttpsHealthCheckOperation struct {
	response map[string]interface{}
}

func (op *createHttpsHealthCheckOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createHttpsHealthCheckOperation) do(ctx context.Context, r *HttpsHealthCheck, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := httpsHealthCheckCreateURL(c.Config.BasePath, project)

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
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetHttpsHealthCheck(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getHttpsHealthCheckRaw(ctx context.Context, r *HttpsHealthCheck) ([]byte, error) {
	if dcl.IsZeroValue(r.CheckIntervalSec) {
		r.CheckIntervalSec = dcl.Int64(5)
	}
	if dcl.IsZeroValue(r.HealthyThreshold) {
		r.HealthyThreshold = dcl.Int64(2)
	}
	if dcl.IsZeroValue(r.Port) {
		r.Port = dcl.Int64(443)
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

	u, err := httpsHealthCheckGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) httpsHealthCheckDiffsForRawDesired(ctx context.Context, rawDesired *HttpsHealthCheck, opts ...dcl.ApplyOption) (initial, desired *HttpsHealthCheck, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *HttpsHealthCheck
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HttpsHealthCheck); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected HttpsHealthCheck, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetHttpsHealthCheck(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a HttpsHealthCheck resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve HttpsHealthCheck resource: %v", err)
		}
		c.Config.Logger.Info("Found that HttpsHealthCheck resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeHttpsHealthCheckDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for HttpsHealthCheck: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for HttpsHealthCheck: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeHttpsHealthCheckInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for HttpsHealthCheck: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeHttpsHealthCheckDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for HttpsHealthCheck: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffHttpsHealthCheck(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeHttpsHealthCheckInitialState(rawInitial, rawDesired *HttpsHealthCheck) (*HttpsHealthCheck, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeHttpsHealthCheckDesiredState(rawDesired, rawInitial *HttpsHealthCheck, opts ...dcl.ApplyOption) (*HttpsHealthCheck, error) {

	if dcl.IsZeroValue(rawDesired.CheckIntervalSec) {
		rawDesired.CheckIntervalSec = dcl.Int64(5)
	}

	if dcl.IsZeroValue(rawDesired.HealthyThreshold) {
		rawDesired.HealthyThreshold = dcl.Int64(2)
	}

	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = dcl.Int64(443)
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
	canonicalDesired := &HttpsHealthCheck{}
	if dcl.IsZeroValue(rawDesired.CheckIntervalSec) {
		canonicalDesired.CheckIntervalSec = rawInitial.CheckIntervalSec
	} else {
		canonicalDesired.CheckIntervalSec = rawDesired.CheckIntervalSec
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.HealthyThreshold) {
		canonicalDesired.HealthyThreshold = rawInitial.HealthyThreshold
	} else {
		canonicalDesired.HealthyThreshold = rawDesired.HealthyThreshold
	}
	if dcl.StringCanonicalize(rawDesired.Host, rawInitial.Host) {
		canonicalDesired.Host = rawInitial.Host
	} else {
		canonicalDesired.Host = rawDesired.Host
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		canonicalDesired.Port = rawInitial.Port
	} else {
		canonicalDesired.Port = rawDesired.Port
	}
	if dcl.StringCanonicalize(rawDesired.RequestPath, rawInitial.RequestPath) {
		canonicalDesired.RequestPath = rawInitial.RequestPath
	} else {
		canonicalDesired.RequestPath = rawDesired.RequestPath
	}
	if dcl.IsZeroValue(rawDesired.TimeoutSec) {
		canonicalDesired.TimeoutSec = rawInitial.TimeoutSec
	} else {
		canonicalDesired.TimeoutSec = rawDesired.TimeoutSec
	}
	if dcl.IsZeroValue(rawDesired.UnhealthyThreshold) {
		canonicalDesired.UnhealthyThreshold = rawInitial.UnhealthyThreshold
	} else {
		canonicalDesired.UnhealthyThreshold = rawDesired.UnhealthyThreshold
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeHttpsHealthCheckNewState(c *Client, rawNew, rawDesired *HttpsHealthCheck) (*HttpsHealthCheck, error) {

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

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
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
func diffHttpsHealthCheck(c *Client, desired, actual *HttpsHealthCheck, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.CheckIntervalSec, actual.CheckIntervalSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("CheckIntervalSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthyThreshold, actual.HealthyThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("HealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestPath, actual.RequestPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("RequestPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSec, actual.TimeoutSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("TimeoutSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnhealthyThreshold, actual.UnhealthyThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHttpsHealthCheckUpdateOperation")}, fn.AddNest("UnhealthyThreshold")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *HttpsHealthCheck) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpsHealthCheck) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *HttpsHealthCheck) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *HttpsHealthCheck) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/httpsHealthChecks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the HttpsHealthCheck resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *HttpsHealthCheck) marshal(c *Client) ([]byte, error) {
	m, err := expandHttpsHealthCheck(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling HttpsHealthCheck: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalHttpsHealthCheck decodes JSON responses into the HttpsHealthCheck resource schema.
func unmarshalHttpsHealthCheck(b []byte, c *Client) (*HttpsHealthCheck, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapHttpsHealthCheck(m, c)
}

func unmarshalMapHttpsHealthCheck(m map[string]interface{}, c *Client) (*HttpsHealthCheck, error) {

	return flattenHttpsHealthCheck(c, m), nil
}

// expandHttpsHealthCheck expands HttpsHealthCheck into a JSON request object.
func expandHttpsHealthCheck(c *Client, f *HttpsHealthCheck) (map[string]interface{}, error) {
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}

	return m, nil
}

// flattenHttpsHealthCheck flattens HttpsHealthCheck from a JSON request object into the
// HttpsHealthCheck type.
func flattenHttpsHealthCheck(c *Client, i interface{}) *HttpsHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &HttpsHealthCheck{}
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
		res.Port = dcl.Int64(443)
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
	res.Project = dcl.FlattenString(m["project"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *HttpsHealthCheck) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalHttpsHealthCheck(b, c)
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

type httpsHealthCheckDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         httpsHealthCheckApiOperation
}

func convertFieldDiffsToHttpsHealthCheckDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]httpsHealthCheckDiff, error) {
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
	var diffs []httpsHealthCheckDiff
	// For each operation name, create a httpsHealthCheckDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := httpsHealthCheckDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToHttpsHealthCheckApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToHttpsHealthCheckApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (httpsHealthCheckApiOperation, error) {
	switch opName {

	case "updateHttpsHealthCheckUpdateOperation":
		return &updateHttpsHealthCheckUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
