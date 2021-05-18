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

func (r *HealthCheck) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Http2HealthCheck) {
		if err := r.Http2HealthCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpHealthCheck) {
		if err := r.HttpHealthCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpsHealthCheck) {
		if err := r.HttpsHealthCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SslHealthCheck) {
		if err := r.SslHealthCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TcpHealthCheck) {
		if err := r.TcpHealthCheck.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *HealthCheckHttp2HealthCheck) validate() error {
	return nil
}
func (r *HealthCheckHttpHealthCheck) validate() error {
	return nil
}
func (r *HealthCheckHttpsHealthCheck) validate() error {
	return nil
}
func (r *HealthCheckSslHealthCheck) validate() error {
	return nil
}
func (r *HealthCheckTcpHealthCheck) validate() error {
	return nil
}

func healthCheckGetURL(userBasePath string, r *HealthCheck) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/healthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/healthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func healthCheckListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/healthChecks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/healthChecks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func healthCheckCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/healthChecks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/healthChecks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func healthCheckDeleteURL(userBasePath string, r *HealthCheck) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/healthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/healthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// healthCheckApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type healthCheckApiOperation interface {
	do(context.Context, *HealthCheck, *Client) error
}

// newUpdateHealthCheckUpdateRequest creates a request for an
// HealthCheck resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateHealthCheckUpdateRequest(ctx context.Context, f *HealthCheck, c *Client) (map[string]interface{}, error) {
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
	if v, err := expandHealthCheckHttp2HealthCheck(c, f.Http2HealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding Http2HealthCheck into http2HealthCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["http2HealthCheck"] = v
	}
	if v, err := expandHealthCheckHttpHealthCheck(c, f.HttpHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpHealthCheck into httpHealthCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpHealthCheck"] = v
	}
	if v, err := expandHealthCheckHttpsHealthCheck(c, f.HttpsHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpsHealthCheck into httpsHealthCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpsHealthCheck"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandHealthCheckSslHealthCheck(c, f.SslHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding SslHealthCheck into sslHealthCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["sslHealthCheck"] = v
	}
	if v, err := expandHealthCheckTcpHealthCheck(c, f.TcpHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding TcpHealthCheck into tcpHealthCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["tcpHealthCheck"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	if v := f.UnhealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		req["unhealthyThreshold"] = v
	}
	if v := f.TimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		req["timeoutSec"] = v
	}
	return req, nil
}

// marshalUpdateHealthCheckUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateHealthCheckUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateHealthCheckUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateHealthCheckUpdateOperation) do(ctx context.Context, r *HealthCheck, c *Client) error {
	_, err := c.GetHealthCheck(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateHealthCheckUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateHealthCheckUpdateRequest(c, req)
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

func (c *Client) listHealthCheckRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := healthCheckListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != HealthCheckMaxPage {
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

type listHealthCheckOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listHealthCheck(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*HealthCheck, string, error) {
	b, err := c.listHealthCheckRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listHealthCheckOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*HealthCheck
	for _, v := range m.Items {
		res := flattenHealthCheck(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllHealthCheck(ctx context.Context, f func(*HealthCheck) bool, resources []*HealthCheck) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteHealthCheck(ctx, res)
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

type deleteHealthCheckOperation struct{}

func (op *deleteHealthCheckOperation) do(ctx context.Context, r *HealthCheck, c *Client) error {

	_, err := c.GetHealthCheck(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("HealthCheck not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetHealthCheck checking for existence. error: %v", err)
		return err
	}

	u, err := healthCheckDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetHealthCheck(ctx, r.urlNormalized())
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
type createHealthCheckOperation struct {
	response map[string]interface{}
}

func (op *createHealthCheckOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createHealthCheckOperation) do(ctx context.Context, r *HealthCheck, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := healthCheckCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetHealthCheck(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getHealthCheckRaw(ctx context.Context, r *HealthCheck) ([]byte, error) {

	u, err := healthCheckGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) healthCheckDiffsForRawDesired(ctx context.Context, rawDesired *HealthCheck, opts ...dcl.ApplyOption) (initial, desired *HealthCheck, diffs []healthCheckDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *HealthCheck
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HealthCheck); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected HealthCheck, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetHealthCheck(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a HealthCheck resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve HealthCheck resource: %v", err)
		}
		c.Config.Logger.Info("Found that HealthCheck resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeHealthCheckDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for HealthCheck: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for HealthCheck: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeHealthCheckInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for HealthCheck: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeHealthCheckDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for HealthCheck: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffHealthCheck(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeHealthCheckInitialState(rawInitial, rawDesired *HealthCheck) (*HealthCheck, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeHealthCheckDesiredState(rawDesired, rawInitial *HealthCheck, opts ...dcl.ApplyOption) (*HealthCheck, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Http2HealthCheck = canonicalizeHealthCheckHttp2HealthCheck(rawDesired.Http2HealthCheck, nil, opts...)
		rawDesired.HttpHealthCheck = canonicalizeHealthCheckHttpHealthCheck(rawDesired.HttpHealthCheck, nil, opts...)
		rawDesired.HttpsHealthCheck = canonicalizeHealthCheckHttpsHealthCheck(rawDesired.HttpsHealthCheck, nil, opts...)
		rawDesired.SslHealthCheck = canonicalizeHealthCheckSslHealthCheck(rawDesired.SslHealthCheck, nil, opts...)
		rawDesired.TcpHealthCheck = canonicalizeHealthCheckTcpHealthCheck(rawDesired.TcpHealthCheck, nil, opts...)

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
	rawDesired.Http2HealthCheck = canonicalizeHealthCheckHttp2HealthCheck(rawDesired.Http2HealthCheck, rawInitial.Http2HealthCheck, opts...)
	rawDesired.HttpHealthCheck = canonicalizeHealthCheckHttpHealthCheck(rawDesired.HttpHealthCheck, rawInitial.HttpHealthCheck, opts...)
	rawDesired.HttpsHealthCheck = canonicalizeHealthCheckHttpsHealthCheck(rawDesired.HttpsHealthCheck, rawInitial.HttpsHealthCheck, opts...)
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.SslHealthCheck = canonicalizeHealthCheckSslHealthCheck(rawDesired.SslHealthCheck, rawInitial.SslHealthCheck, opts...)
	rawDesired.TcpHealthCheck = canonicalizeHealthCheckTcpHealthCheck(rawDesired.TcpHealthCheck, rawInitial.TcpHealthCheck, opts...)
	if dcl.IsZeroValue(rawDesired.Type) {
		rawDesired.Type = rawInitial.Type
	}
	if dcl.IsZeroValue(rawDesired.UnhealthyThreshold) {
		rawDesired.UnhealthyThreshold = rawInitial.UnhealthyThreshold
	}
	if dcl.IsZeroValue(rawDesired.TimeoutSec) {
		rawDesired.TimeoutSec = rawInitial.TimeoutSec
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeHealthCheckNewState(c *Client, rawNew, rawDesired *HealthCheck) (*HealthCheck, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Http2HealthCheck) && dcl.IsEmptyValueIndirect(rawDesired.Http2HealthCheck) {
		rawNew.Http2HealthCheck = rawDesired.Http2HealthCheck
	} else {
		rawNew.Http2HealthCheck = canonicalizeNewHealthCheckHttp2HealthCheck(c, rawDesired.Http2HealthCheck, rawNew.Http2HealthCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpHealthCheck) && dcl.IsEmptyValueIndirect(rawDesired.HttpHealthCheck) {
		rawNew.HttpHealthCheck = rawDesired.HttpHealthCheck
	} else {
		rawNew.HttpHealthCheck = canonicalizeNewHealthCheckHttpHealthCheck(c, rawDesired.HttpHealthCheck, rawNew.HttpHealthCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpsHealthCheck) && dcl.IsEmptyValueIndirect(rawDesired.HttpsHealthCheck) {
		rawNew.HttpsHealthCheck = rawDesired.HttpsHealthCheck
	} else {
		rawNew.HttpsHealthCheck = canonicalizeNewHealthCheckHttpsHealthCheck(c, rawDesired.HttpsHealthCheck, rawNew.HttpsHealthCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SslHealthCheck) && dcl.IsEmptyValueIndirect(rawDesired.SslHealthCheck) {
		rawNew.SslHealthCheck = rawDesired.SslHealthCheck
	} else {
		rawNew.SslHealthCheck = canonicalizeNewHealthCheckSslHealthCheck(c, rawDesired.SslHealthCheck, rawNew.SslHealthCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TcpHealthCheck) && dcl.IsEmptyValueIndirect(rawDesired.TcpHealthCheck) {
		rawNew.TcpHealthCheck = rawDesired.TcpHealthCheck
	} else {
		rawNew.TcpHealthCheck = canonicalizeNewHealthCheckTcpHealthCheck(c, rawDesired.TcpHealthCheck, rawNew.TcpHealthCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UnhealthyThreshold) && dcl.IsEmptyValueIndirect(rawDesired.UnhealthyThreshold) {
		rawNew.UnhealthyThreshold = rawDesired.UnhealthyThreshold
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeoutSec) && dcl.IsEmptyValueIndirect(rawDesired.TimeoutSec) {
		rawNew.TimeoutSec = rawDesired.TimeoutSec
	} else {
	}

	rawNew.Region = rawDesired.Region

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeHealthCheckHttp2HealthCheck(des, initial *HealthCheckHttp2HealthCheck, opts ...dcl.ApplyOption) *HealthCheckHttp2HealthCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.PortName, initial.PortName) || dcl.IsZeroValue(des.PortName) {
		des.PortName = initial.PortName
	}
	if dcl.IsZeroValue(des.PortSpecification) {
		des.PortSpecification = initial.PortSpecification
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.StringCanonicalize(des.RequestPath, initial.RequestPath) || dcl.IsZeroValue(des.RequestPath) {
		des.RequestPath = initial.RequestPath
	}
	if dcl.IsZeroValue(des.ProxyHeader) {
		des.ProxyHeader = initial.ProxyHeader
	}
	if dcl.StringCanonicalize(des.Response, initial.Response) || dcl.IsZeroValue(des.Response) {
		des.Response = initial.Response
	}

	return des
}

func canonicalizeNewHealthCheckHttp2HealthCheck(c *Client, des, nw *HealthCheckHttp2HealthCheck) *HealthCheckHttp2HealthCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}
	if dcl.StringCanonicalize(des.PortName, nw.PortName) {
		nw.PortName = des.PortName
	}
	if dcl.IsZeroValue(nw.PortSpecification) {
		nw.PortSpecification = des.PortSpecification
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.RequestPath, nw.RequestPath) {
		nw.RequestPath = des.RequestPath
	}
	if dcl.IsZeroValue(nw.ProxyHeader) {
		nw.ProxyHeader = des.ProxyHeader
	}
	if dcl.StringCanonicalize(des.Response, nw.Response) {
		nw.Response = des.Response
	}

	return nw
}

func canonicalizeNewHealthCheckHttp2HealthCheckSet(c *Client, des, nw []HealthCheckHttp2HealthCheck) []HealthCheckHttp2HealthCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []HealthCheckHttp2HealthCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareHealthCheckHttp2HealthCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewHealthCheckHttp2HealthCheckSlice(c *Client, des, nw []HealthCheckHttp2HealthCheck) []HealthCheckHttp2HealthCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HealthCheckHttp2HealthCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHealthCheckHttp2HealthCheck(c, &d, &n))
	}

	return items
}

func canonicalizeHealthCheckHttpHealthCheck(des, initial *HealthCheckHttpHealthCheck, opts ...dcl.ApplyOption) *HealthCheckHttpHealthCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.PortName, initial.PortName) || dcl.IsZeroValue(des.PortName) {
		des.PortName = initial.PortName
	}
	if dcl.IsZeroValue(des.PortSpecification) {
		des.PortSpecification = initial.PortSpecification
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.StringCanonicalize(des.RequestPath, initial.RequestPath) || dcl.IsZeroValue(des.RequestPath) {
		des.RequestPath = initial.RequestPath
	}
	if dcl.IsZeroValue(des.ProxyHeader) {
		des.ProxyHeader = initial.ProxyHeader
	}
	if dcl.StringCanonicalize(des.Response, initial.Response) || dcl.IsZeroValue(des.Response) {
		des.Response = initial.Response
	}

	return des
}

func canonicalizeNewHealthCheckHttpHealthCheck(c *Client, des, nw *HealthCheckHttpHealthCheck) *HealthCheckHttpHealthCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}
	if dcl.StringCanonicalize(des.PortName, nw.PortName) {
		nw.PortName = des.PortName
	}
	if dcl.IsZeroValue(nw.PortSpecification) {
		nw.PortSpecification = des.PortSpecification
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.RequestPath, nw.RequestPath) {
		nw.RequestPath = des.RequestPath
	}
	if dcl.IsZeroValue(nw.ProxyHeader) {
		nw.ProxyHeader = des.ProxyHeader
	}
	if dcl.StringCanonicalize(des.Response, nw.Response) {
		nw.Response = des.Response
	}

	return nw
}

func canonicalizeNewHealthCheckHttpHealthCheckSet(c *Client, des, nw []HealthCheckHttpHealthCheck) []HealthCheckHttpHealthCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []HealthCheckHttpHealthCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareHealthCheckHttpHealthCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewHealthCheckHttpHealthCheckSlice(c *Client, des, nw []HealthCheckHttpHealthCheck) []HealthCheckHttpHealthCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HealthCheckHttpHealthCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHealthCheckHttpHealthCheck(c, &d, &n))
	}

	return items
}

func canonicalizeHealthCheckHttpsHealthCheck(des, initial *HealthCheckHttpsHealthCheck, opts ...dcl.ApplyOption) *HealthCheckHttpsHealthCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.PortName, initial.PortName) || dcl.IsZeroValue(des.PortName) {
		des.PortName = initial.PortName
	}
	if dcl.IsZeroValue(des.PortSpecification) {
		des.PortSpecification = initial.PortSpecification
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		des.Host = initial.Host
	}
	if dcl.StringCanonicalize(des.RequestPath, initial.RequestPath) || dcl.IsZeroValue(des.RequestPath) {
		des.RequestPath = initial.RequestPath
	}
	if dcl.IsZeroValue(des.ProxyHeader) {
		des.ProxyHeader = initial.ProxyHeader
	}
	if dcl.StringCanonicalize(des.Response, initial.Response) || dcl.IsZeroValue(des.Response) {
		des.Response = initial.Response
	}

	return des
}

func canonicalizeNewHealthCheckHttpsHealthCheck(c *Client, des, nw *HealthCheckHttpsHealthCheck) *HealthCheckHttpsHealthCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}
	if dcl.StringCanonicalize(des.PortName, nw.PortName) {
		nw.PortName = des.PortName
	}
	if dcl.IsZeroValue(nw.PortSpecification) {
		nw.PortSpecification = des.PortSpecification
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.RequestPath, nw.RequestPath) {
		nw.RequestPath = des.RequestPath
	}
	if dcl.IsZeroValue(nw.ProxyHeader) {
		nw.ProxyHeader = des.ProxyHeader
	}
	if dcl.StringCanonicalize(des.Response, nw.Response) {
		nw.Response = des.Response
	}

	return nw
}

func canonicalizeNewHealthCheckHttpsHealthCheckSet(c *Client, des, nw []HealthCheckHttpsHealthCheck) []HealthCheckHttpsHealthCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []HealthCheckHttpsHealthCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareHealthCheckHttpsHealthCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewHealthCheckHttpsHealthCheckSlice(c *Client, des, nw []HealthCheckHttpsHealthCheck) []HealthCheckHttpsHealthCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HealthCheckHttpsHealthCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHealthCheckHttpsHealthCheck(c, &d, &n))
	}

	return items
}

func canonicalizeHealthCheckSslHealthCheck(des, initial *HealthCheckSslHealthCheck, opts ...dcl.ApplyOption) *HealthCheckSslHealthCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.PortName, initial.PortName) || dcl.IsZeroValue(des.PortName) {
		des.PortName = initial.PortName
	}
	if dcl.IsZeroValue(des.PortSpecification) {
		des.PortSpecification = initial.PortSpecification
	}
	if dcl.StringCanonicalize(des.Request, initial.Request) || dcl.IsZeroValue(des.Request) {
		des.Request = initial.Request
	}
	if dcl.StringCanonicalize(des.Response, initial.Response) || dcl.IsZeroValue(des.Response) {
		des.Response = initial.Response
	}
	if dcl.IsZeroValue(des.ProxyHeader) {
		des.ProxyHeader = initial.ProxyHeader
	}

	return des
}

func canonicalizeNewHealthCheckSslHealthCheck(c *Client, des, nw *HealthCheckSslHealthCheck) *HealthCheckSslHealthCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}
	if dcl.StringCanonicalize(des.PortName, nw.PortName) {
		nw.PortName = des.PortName
	}
	if dcl.IsZeroValue(nw.PortSpecification) {
		nw.PortSpecification = des.PortSpecification
	}
	if dcl.StringCanonicalize(des.Request, nw.Request) {
		nw.Request = des.Request
	}
	if dcl.StringCanonicalize(des.Response, nw.Response) {
		nw.Response = des.Response
	}
	if dcl.IsZeroValue(nw.ProxyHeader) {
		nw.ProxyHeader = des.ProxyHeader
	}

	return nw
}

func canonicalizeNewHealthCheckSslHealthCheckSet(c *Client, des, nw []HealthCheckSslHealthCheck) []HealthCheckSslHealthCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []HealthCheckSslHealthCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareHealthCheckSslHealthCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewHealthCheckSslHealthCheckSlice(c *Client, des, nw []HealthCheckSslHealthCheck) []HealthCheckSslHealthCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HealthCheckSslHealthCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHealthCheckSslHealthCheck(c, &d, &n))
	}

	return items
}

func canonicalizeHealthCheckTcpHealthCheck(des, initial *HealthCheckTcpHealthCheck, opts ...dcl.ApplyOption) *HealthCheckTcpHealthCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	if dcl.StringCanonicalize(des.PortName, initial.PortName) || dcl.IsZeroValue(des.PortName) {
		des.PortName = initial.PortName
	}
	if dcl.IsZeroValue(des.PortSpecification) {
		des.PortSpecification = initial.PortSpecification
	}
	if dcl.StringCanonicalize(des.Request, initial.Request) || dcl.IsZeroValue(des.Request) {
		des.Request = initial.Request
	}
	if dcl.StringCanonicalize(des.Response, initial.Response) || dcl.IsZeroValue(des.Response) {
		des.Response = initial.Response
	}
	if dcl.IsZeroValue(des.ProxyHeader) {
		des.ProxyHeader = initial.ProxyHeader
	}

	return des
}

func canonicalizeNewHealthCheckTcpHealthCheck(c *Client, des, nw *HealthCheckTcpHealthCheck) *HealthCheckTcpHealthCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}
	if dcl.StringCanonicalize(des.PortName, nw.PortName) {
		nw.PortName = des.PortName
	}
	if dcl.IsZeroValue(nw.PortSpecification) {
		nw.PortSpecification = des.PortSpecification
	}
	if dcl.StringCanonicalize(des.Request, nw.Request) {
		nw.Request = des.Request
	}
	if dcl.StringCanonicalize(des.Response, nw.Response) {
		nw.Response = des.Response
	}
	if dcl.IsZeroValue(nw.ProxyHeader) {
		nw.ProxyHeader = des.ProxyHeader
	}

	return nw
}

func canonicalizeNewHealthCheckTcpHealthCheckSet(c *Client, des, nw []HealthCheckTcpHealthCheck) []HealthCheckTcpHealthCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []HealthCheckTcpHealthCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareHealthCheckTcpHealthCheckNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewHealthCheckTcpHealthCheckSlice(c *Client, des, nw []HealthCheckTcpHealthCheck) []HealthCheckTcpHealthCheck {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HealthCheckTcpHealthCheck
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHealthCheckTcpHealthCheck(c, &d, &n))
	}

	return items
}

type healthCheckDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         healthCheckApiOperation
	Diffs            []*dcl.FieldDiff
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
func diffHealthCheck(c *Client, desired, actual *HealthCheck, opts ...dcl.ApplyOption) ([]healthCheckDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []healthCheckDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.CheckIntervalSec, actual.CheckIntervalSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("CheckIntervalSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.HealthyThreshold, actual.HealthyThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("HealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Http2HealthCheck, actual.Http2HealthCheck, dcl.Info{ObjectFunction: compareHealthCheckHttp2HealthCheckNewStyle, OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Http2HealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.HttpHealthCheck, actual.HttpHealthCheck, dcl.Info{ObjectFunction: compareHealthCheckHttpHealthCheckNewStyle, OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("HttpHealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.HttpsHealthCheck, actual.HttpsHealthCheck, dcl.Info{ObjectFunction: compareHealthCheckHttpsHealthCheckNewStyle, OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("HttpsHealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SslHealthCheck, actual.SslHealthCheck, dcl.Info{ObjectFunction: compareHealthCheckSslHealthCheckNewStyle, OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("SslHealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.TcpHealthCheck, actual.TcpHealthCheck, dcl.Info{ObjectFunction: compareHealthCheckTcpHealthCheckNewStyle, OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("TcpHealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.UnhealthyThreshold, actual.UnhealthyThreshold, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("UnhealthyThreshold")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSec, actual.TimeoutSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("TimeoutSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToHealthCheckDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
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
	var deduped []healthCheckDiff
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
func compareHealthCheckHttp2HealthCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HealthCheckHttp2HealthCheck)
	if !ok {
		desiredNotPointer, ok := d.(HealthCheckHttp2HealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckHttp2HealthCheck or *HealthCheckHttp2HealthCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HealthCheckHttp2HealthCheck)
	if !ok {
		actualNotPointer, ok := a.(HealthCheckHttp2HealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckHttp2HealthCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortName, actual.PortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortSpecification, actual.PortSpecification, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortSpecification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestPath, actual.RequestPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("RequestPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyHeader, actual.ProxyHeader, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("ProxyHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Response, actual.Response, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Response")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHealthCheckHttpHealthCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HealthCheckHttpHealthCheck)
	if !ok {
		desiredNotPointer, ok := d.(HealthCheckHttpHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckHttpHealthCheck or *HealthCheckHttpHealthCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HealthCheckHttpHealthCheck)
	if !ok {
		actualNotPointer, ok := a.(HealthCheckHttpHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckHttpHealthCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortName, actual.PortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortSpecification, actual.PortSpecification, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortSpecification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestPath, actual.RequestPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("RequestPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyHeader, actual.ProxyHeader, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("ProxyHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Response, actual.Response, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Response")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHealthCheckHttpsHealthCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HealthCheckHttpsHealthCheck)
	if !ok {
		desiredNotPointer, ok := d.(HealthCheckHttpsHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckHttpsHealthCheck or *HealthCheckHttpsHealthCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HealthCheckHttpsHealthCheck)
	if !ok {
		actualNotPointer, ok := a.(HealthCheckHttpsHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckHttpsHealthCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortName, actual.PortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortSpecification, actual.PortSpecification, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortSpecification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestPath, actual.RequestPath, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("RequestPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyHeader, actual.ProxyHeader, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("ProxyHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Response, actual.Response, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Response")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHealthCheckSslHealthCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HealthCheckSslHealthCheck)
	if !ok {
		desiredNotPointer, ok := d.(HealthCheckSslHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckSslHealthCheck or *HealthCheckSslHealthCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HealthCheckSslHealthCheck)
	if !ok {
		actualNotPointer, ok := a.(HealthCheckSslHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckSslHealthCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortName, actual.PortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortSpecification, actual.PortSpecification, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortSpecification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Request, actual.Request, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Request")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Response, actual.Response, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Response")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyHeader, actual.ProxyHeader, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("ProxyHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHealthCheckTcpHealthCheckNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HealthCheckTcpHealthCheck)
	if !ok {
		desiredNotPointer, ok := d.(HealthCheckTcpHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckTcpHealthCheck or *HealthCheckTcpHealthCheck", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HealthCheckTcpHealthCheck)
	if !ok {
		actualNotPointer, ok := a.(HealthCheckTcpHealthCheck)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HealthCheckTcpHealthCheck", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortName, actual.PortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortSpecification, actual.PortSpecification, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("PortSpecification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Request, actual.Request, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Request")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Response, actual.Response, dcl.Info{OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("Response")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyHeader, actual.ProxyHeader, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHealthCheckUpdateOperation")}, fn.AddNest("ProxyHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *HealthCheck) urlNormalized() *HealthCheck {
	normalized := dcl.Copy(*r).(HealthCheck)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *HealthCheck) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *HealthCheck) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *HealthCheck) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *HealthCheck) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/healthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil
		}

		return dcl.URL("projects/{{project}}/global/healthChecks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the HealthCheck resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *HealthCheck) marshal(c *Client) ([]byte, error) {
	m, err := expandHealthCheck(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling HealthCheck: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalHealthCheck decodes JSON responses into the HealthCheck resource schema.
func unmarshalHealthCheck(b []byte, c *Client) (*HealthCheck, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapHealthCheck(m, c)
}

func unmarshalMapHealthCheck(m map[string]interface{}, c *Client) (*HealthCheck, error) {

	return flattenHealthCheck(c, m), nil
}

// expandHealthCheck expands HealthCheck into a JSON request object.
func expandHealthCheck(c *Client, f *HealthCheck) (map[string]interface{}, error) {
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
	if v, err := expandHealthCheckHttp2HealthCheck(c, f.Http2HealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding Http2HealthCheck into http2HealthCheck: %w", err)
	} else if v != nil {
		m["http2HealthCheck"] = v
	}
	if v, err := expandHealthCheckHttpHealthCheck(c, f.HttpHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpHealthCheck into httpHealthCheck: %w", err)
	} else if v != nil {
		m["httpHealthCheck"] = v
	}
	if v, err := expandHealthCheckHttpsHealthCheck(c, f.HttpsHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpsHealthCheck into httpsHealthCheck: %w", err)
	} else if v != nil {
		m["httpsHealthCheck"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandHealthCheckSslHealthCheck(c, f.SslHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding SslHealthCheck into sslHealthCheck: %w", err)
	} else if v != nil {
		m["sslHealthCheck"] = v
	}
	if v, err := expandHealthCheckTcpHealthCheck(c, f.TcpHealthCheck); err != nil {
		return nil, fmt.Errorf("error expanding TcpHealthCheck into tcpHealthCheck: %w", err)
	} else if v != nil {
		m["tcpHealthCheck"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.UnhealthyThreshold; !dcl.IsEmptyValueIndirect(v) {
		m["unhealthyThreshold"] = v
	}
	if v := f.TimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["timeoutSec"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if v != nil {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenHealthCheck flattens HealthCheck from a JSON request object into the
// HealthCheck type.
func flattenHealthCheck(c *Client, i interface{}) *HealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &HealthCheck{}
	res.CheckIntervalSec = dcl.FlattenInteger(m["checkIntervalSec"])
	res.Description = dcl.FlattenString(m["description"])
	res.HealthyThreshold = dcl.FlattenInteger(m["healthyThreshold"])
	res.Http2HealthCheck = flattenHealthCheckHttp2HealthCheck(c, m["http2HealthCheck"])
	res.HttpHealthCheck = flattenHealthCheckHttpHealthCheck(c, m["httpHealthCheck"])
	res.HttpsHealthCheck = flattenHealthCheckHttpsHealthCheck(c, m["httpsHealthCheck"])
	res.Name = dcl.FlattenString(m["name"])
	res.SslHealthCheck = flattenHealthCheckSslHealthCheck(c, m["sslHealthCheck"])
	res.TcpHealthCheck = flattenHealthCheckTcpHealthCheck(c, m["tcpHealthCheck"])
	res.Type = flattenHealthCheckTypeEnum(m["type"])
	res.UnhealthyThreshold = dcl.FlattenInteger(m["unhealthyThreshold"])
	res.TimeoutSec = dcl.FlattenInteger(m["timeoutSec"])
	res.Region = dcl.FlattenString(m["region"])
	res.Project = dcl.FlattenString(m["project"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandHealthCheckHttp2HealthCheckMap expands the contents of HealthCheckHttp2HealthCheck into a JSON
// request object.
func expandHealthCheckHttp2HealthCheckMap(c *Client, f map[string]HealthCheckHttp2HealthCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHealthCheckHttp2HealthCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHealthCheckHttp2HealthCheckSlice expands the contents of HealthCheckHttp2HealthCheck into a JSON
// request object.
func expandHealthCheckHttp2HealthCheckSlice(c *Client, f []HealthCheckHttp2HealthCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHealthCheckHttp2HealthCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHealthCheckHttp2HealthCheckMap flattens the contents of HealthCheckHttp2HealthCheck from a JSON
// response object.
func flattenHealthCheckHttp2HealthCheckMap(c *Client, i interface{}) map[string]HealthCheckHttp2HealthCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HealthCheckHttp2HealthCheck{}
	}

	if len(a) == 0 {
		return map[string]HealthCheckHttp2HealthCheck{}
	}

	items := make(map[string]HealthCheckHttp2HealthCheck)
	for k, item := range a {
		items[k] = *flattenHealthCheckHttp2HealthCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenHealthCheckHttp2HealthCheckSlice flattens the contents of HealthCheckHttp2HealthCheck from a JSON
// response object.
func flattenHealthCheckHttp2HealthCheckSlice(c *Client, i interface{}) []HealthCheckHttp2HealthCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttp2HealthCheck{}
	}

	if len(a) == 0 {
		return []HealthCheckHttp2HealthCheck{}
	}

	items := make([]HealthCheckHttp2HealthCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttp2HealthCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandHealthCheckHttp2HealthCheck expands an instance of HealthCheckHttp2HealthCheck into a JSON
// request object.
func expandHealthCheckHttp2HealthCheck(c *Client, f *HealthCheckHttp2HealthCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		m["portName"] = v
	}
	if v := f.PortSpecification; !dcl.IsEmptyValueIndirect(v) {
		m["portSpecification"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.RequestPath; !dcl.IsEmptyValueIndirect(v) {
		m["requestPath"] = v
	}
	if v := f.ProxyHeader; !dcl.IsEmptyValueIndirect(v) {
		m["proxyHeader"] = v
	}
	if v := f.Response; !dcl.IsEmptyValueIndirect(v) {
		m["response"] = v
	}

	return m, nil
}

// flattenHealthCheckHttp2HealthCheck flattens an instance of HealthCheckHttp2HealthCheck from a JSON
// response object.
func flattenHealthCheckHttp2HealthCheck(c *Client, i interface{}) *HealthCheckHttp2HealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HealthCheckHttp2HealthCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHealthCheckHttp2HealthCheck
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.PortName = dcl.FlattenString(m["portName"])
	r.PortSpecification = flattenHealthCheckHttp2HealthCheckPortSpecificationEnum(m["portSpecification"])
	r.Host = dcl.FlattenString(m["host"])
	r.RequestPath = dcl.FlattenString(m["requestPath"])
	r.ProxyHeader = flattenHealthCheckHttp2HealthCheckProxyHeaderEnum(m["proxyHeader"])
	r.Response = dcl.FlattenString(m["response"])

	return r
}

// expandHealthCheckHttpHealthCheckMap expands the contents of HealthCheckHttpHealthCheck into a JSON
// request object.
func expandHealthCheckHttpHealthCheckMap(c *Client, f map[string]HealthCheckHttpHealthCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHealthCheckHttpHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHealthCheckHttpHealthCheckSlice expands the contents of HealthCheckHttpHealthCheck into a JSON
// request object.
func expandHealthCheckHttpHealthCheckSlice(c *Client, f []HealthCheckHttpHealthCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHealthCheckHttpHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHealthCheckHttpHealthCheckMap flattens the contents of HealthCheckHttpHealthCheck from a JSON
// response object.
func flattenHealthCheckHttpHealthCheckMap(c *Client, i interface{}) map[string]HealthCheckHttpHealthCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HealthCheckHttpHealthCheck{}
	}

	if len(a) == 0 {
		return map[string]HealthCheckHttpHealthCheck{}
	}

	items := make(map[string]HealthCheckHttpHealthCheck)
	for k, item := range a {
		items[k] = *flattenHealthCheckHttpHealthCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenHealthCheckHttpHealthCheckSlice flattens the contents of HealthCheckHttpHealthCheck from a JSON
// response object.
func flattenHealthCheckHttpHealthCheckSlice(c *Client, i interface{}) []HealthCheckHttpHealthCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttpHealthCheck{}
	}

	if len(a) == 0 {
		return []HealthCheckHttpHealthCheck{}
	}

	items := make([]HealthCheckHttpHealthCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttpHealthCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandHealthCheckHttpHealthCheck expands an instance of HealthCheckHttpHealthCheck into a JSON
// request object.
func expandHealthCheckHttpHealthCheck(c *Client, f *HealthCheckHttpHealthCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		m["portName"] = v
	}
	if v := f.PortSpecification; !dcl.IsEmptyValueIndirect(v) {
		m["portSpecification"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.RequestPath; !dcl.IsEmptyValueIndirect(v) {
		m["requestPath"] = v
	}
	if v := f.ProxyHeader; !dcl.IsEmptyValueIndirect(v) {
		m["proxyHeader"] = v
	}
	if v := f.Response; !dcl.IsEmptyValueIndirect(v) {
		m["response"] = v
	}

	return m, nil
}

// flattenHealthCheckHttpHealthCheck flattens an instance of HealthCheckHttpHealthCheck from a JSON
// response object.
func flattenHealthCheckHttpHealthCheck(c *Client, i interface{}) *HealthCheckHttpHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HealthCheckHttpHealthCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHealthCheckHttpHealthCheck
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.PortName = dcl.FlattenString(m["portName"])
	r.PortSpecification = flattenHealthCheckHttpHealthCheckPortSpecificationEnum(m["portSpecification"])
	r.Host = dcl.FlattenString(m["host"])
	r.RequestPath = dcl.FlattenString(m["requestPath"])
	r.ProxyHeader = flattenHealthCheckHttpHealthCheckProxyHeaderEnum(m["proxyHeader"])
	r.Response = dcl.FlattenString(m["response"])

	return r
}

// expandHealthCheckHttpsHealthCheckMap expands the contents of HealthCheckHttpsHealthCheck into a JSON
// request object.
func expandHealthCheckHttpsHealthCheckMap(c *Client, f map[string]HealthCheckHttpsHealthCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHealthCheckHttpsHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHealthCheckHttpsHealthCheckSlice expands the contents of HealthCheckHttpsHealthCheck into a JSON
// request object.
func expandHealthCheckHttpsHealthCheckSlice(c *Client, f []HealthCheckHttpsHealthCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHealthCheckHttpsHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHealthCheckHttpsHealthCheckMap flattens the contents of HealthCheckHttpsHealthCheck from a JSON
// response object.
func flattenHealthCheckHttpsHealthCheckMap(c *Client, i interface{}) map[string]HealthCheckHttpsHealthCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HealthCheckHttpsHealthCheck{}
	}

	if len(a) == 0 {
		return map[string]HealthCheckHttpsHealthCheck{}
	}

	items := make(map[string]HealthCheckHttpsHealthCheck)
	for k, item := range a {
		items[k] = *flattenHealthCheckHttpsHealthCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenHealthCheckHttpsHealthCheckSlice flattens the contents of HealthCheckHttpsHealthCheck from a JSON
// response object.
func flattenHealthCheckHttpsHealthCheckSlice(c *Client, i interface{}) []HealthCheckHttpsHealthCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttpsHealthCheck{}
	}

	if len(a) == 0 {
		return []HealthCheckHttpsHealthCheck{}
	}

	items := make([]HealthCheckHttpsHealthCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttpsHealthCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandHealthCheckHttpsHealthCheck expands an instance of HealthCheckHttpsHealthCheck into a JSON
// request object.
func expandHealthCheckHttpsHealthCheck(c *Client, f *HealthCheckHttpsHealthCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		m["portName"] = v
	}
	if v := f.PortSpecification; !dcl.IsEmptyValueIndirect(v) {
		m["portSpecification"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.RequestPath; !dcl.IsEmptyValueIndirect(v) {
		m["requestPath"] = v
	}
	if v := f.ProxyHeader; !dcl.IsEmptyValueIndirect(v) {
		m["proxyHeader"] = v
	}
	if v := f.Response; !dcl.IsEmptyValueIndirect(v) {
		m["response"] = v
	}

	return m, nil
}

// flattenHealthCheckHttpsHealthCheck flattens an instance of HealthCheckHttpsHealthCheck from a JSON
// response object.
func flattenHealthCheckHttpsHealthCheck(c *Client, i interface{}) *HealthCheckHttpsHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HealthCheckHttpsHealthCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHealthCheckHttpsHealthCheck
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.PortName = dcl.FlattenString(m["portName"])
	r.PortSpecification = flattenHealthCheckHttpsHealthCheckPortSpecificationEnum(m["portSpecification"])
	r.Host = dcl.FlattenString(m["host"])
	r.RequestPath = dcl.FlattenString(m["requestPath"])
	r.ProxyHeader = flattenHealthCheckHttpsHealthCheckProxyHeaderEnum(m["proxyHeader"])
	r.Response = dcl.FlattenString(m["response"])

	return r
}

// expandHealthCheckSslHealthCheckMap expands the contents of HealthCheckSslHealthCheck into a JSON
// request object.
func expandHealthCheckSslHealthCheckMap(c *Client, f map[string]HealthCheckSslHealthCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHealthCheckSslHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHealthCheckSslHealthCheckSlice expands the contents of HealthCheckSslHealthCheck into a JSON
// request object.
func expandHealthCheckSslHealthCheckSlice(c *Client, f []HealthCheckSslHealthCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHealthCheckSslHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHealthCheckSslHealthCheckMap flattens the contents of HealthCheckSslHealthCheck from a JSON
// response object.
func flattenHealthCheckSslHealthCheckMap(c *Client, i interface{}) map[string]HealthCheckSslHealthCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HealthCheckSslHealthCheck{}
	}

	if len(a) == 0 {
		return map[string]HealthCheckSslHealthCheck{}
	}

	items := make(map[string]HealthCheckSslHealthCheck)
	for k, item := range a {
		items[k] = *flattenHealthCheckSslHealthCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenHealthCheckSslHealthCheckSlice flattens the contents of HealthCheckSslHealthCheck from a JSON
// response object.
func flattenHealthCheckSslHealthCheckSlice(c *Client, i interface{}) []HealthCheckSslHealthCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckSslHealthCheck{}
	}

	if len(a) == 0 {
		return []HealthCheckSslHealthCheck{}
	}

	items := make([]HealthCheckSslHealthCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckSslHealthCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandHealthCheckSslHealthCheck expands an instance of HealthCheckSslHealthCheck into a JSON
// request object.
func expandHealthCheckSslHealthCheck(c *Client, f *HealthCheckSslHealthCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		m["portName"] = v
	}
	if v := f.PortSpecification; !dcl.IsEmptyValueIndirect(v) {
		m["portSpecification"] = v
	}
	if v := f.Request; !dcl.IsEmptyValueIndirect(v) {
		m["request"] = v
	}
	if v := f.Response; !dcl.IsEmptyValueIndirect(v) {
		m["response"] = v
	}
	if v := f.ProxyHeader; !dcl.IsEmptyValueIndirect(v) {
		m["proxyHeader"] = v
	}

	return m, nil
}

// flattenHealthCheckSslHealthCheck flattens an instance of HealthCheckSslHealthCheck from a JSON
// response object.
func flattenHealthCheckSslHealthCheck(c *Client, i interface{}) *HealthCheckSslHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HealthCheckSslHealthCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHealthCheckSslHealthCheck
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.PortName = dcl.FlattenString(m["portName"])
	r.PortSpecification = flattenHealthCheckSslHealthCheckPortSpecificationEnum(m["portSpecification"])
	r.Request = dcl.FlattenString(m["request"])
	r.Response = dcl.FlattenString(m["response"])
	r.ProxyHeader = flattenHealthCheckSslHealthCheckProxyHeaderEnum(m["proxyHeader"])

	return r
}

// expandHealthCheckTcpHealthCheckMap expands the contents of HealthCheckTcpHealthCheck into a JSON
// request object.
func expandHealthCheckTcpHealthCheckMap(c *Client, f map[string]HealthCheckTcpHealthCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHealthCheckTcpHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHealthCheckTcpHealthCheckSlice expands the contents of HealthCheckTcpHealthCheck into a JSON
// request object.
func expandHealthCheckTcpHealthCheckSlice(c *Client, f []HealthCheckTcpHealthCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHealthCheckTcpHealthCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHealthCheckTcpHealthCheckMap flattens the contents of HealthCheckTcpHealthCheck from a JSON
// response object.
func flattenHealthCheckTcpHealthCheckMap(c *Client, i interface{}) map[string]HealthCheckTcpHealthCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HealthCheckTcpHealthCheck{}
	}

	if len(a) == 0 {
		return map[string]HealthCheckTcpHealthCheck{}
	}

	items := make(map[string]HealthCheckTcpHealthCheck)
	for k, item := range a {
		items[k] = *flattenHealthCheckTcpHealthCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenHealthCheckTcpHealthCheckSlice flattens the contents of HealthCheckTcpHealthCheck from a JSON
// response object.
func flattenHealthCheckTcpHealthCheckSlice(c *Client, i interface{}) []HealthCheckTcpHealthCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckTcpHealthCheck{}
	}

	if len(a) == 0 {
		return []HealthCheckTcpHealthCheck{}
	}

	items := make([]HealthCheckTcpHealthCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckTcpHealthCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandHealthCheckTcpHealthCheck expands an instance of HealthCheckTcpHealthCheck into a JSON
// request object.
func expandHealthCheckTcpHealthCheck(c *Client, f *HealthCheckTcpHealthCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		m["portName"] = v
	}
	if v := f.PortSpecification; !dcl.IsEmptyValueIndirect(v) {
		m["portSpecification"] = v
	}
	if v := f.Request; !dcl.IsEmptyValueIndirect(v) {
		m["request"] = v
	}
	if v := f.Response; !dcl.IsEmptyValueIndirect(v) {
		m["response"] = v
	}
	if v := f.ProxyHeader; !dcl.IsEmptyValueIndirect(v) {
		m["proxyHeader"] = v
	}

	return m, nil
}

// flattenHealthCheckTcpHealthCheck flattens an instance of HealthCheckTcpHealthCheck from a JSON
// response object.
func flattenHealthCheckTcpHealthCheck(c *Client, i interface{}) *HealthCheckTcpHealthCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HealthCheckTcpHealthCheck{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHealthCheckTcpHealthCheck
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.PortName = dcl.FlattenString(m["portName"])
	r.PortSpecification = flattenHealthCheckTcpHealthCheckPortSpecificationEnum(m["portSpecification"])
	r.Request = dcl.FlattenString(m["request"])
	r.Response = dcl.FlattenString(m["response"])
	r.ProxyHeader = flattenHealthCheckTcpHealthCheckProxyHeaderEnum(m["proxyHeader"])

	return r
}

// flattenHealthCheckHttp2HealthCheckPortSpecificationEnumSlice flattens the contents of HealthCheckHttp2HealthCheckPortSpecificationEnum from a JSON
// response object.
func flattenHealthCheckHttp2HealthCheckPortSpecificationEnumSlice(c *Client, i interface{}) []HealthCheckHttp2HealthCheckPortSpecificationEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttp2HealthCheckPortSpecificationEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckHttp2HealthCheckPortSpecificationEnum{}
	}

	items := make([]HealthCheckHttp2HealthCheckPortSpecificationEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttp2HealthCheckPortSpecificationEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckHttp2HealthCheckPortSpecificationEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckHttp2HealthCheckPortSpecificationEnum with the same value as that string.
func flattenHealthCheckHttp2HealthCheckPortSpecificationEnum(i interface{}) *HealthCheckHttp2HealthCheckPortSpecificationEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckHttp2HealthCheckPortSpecificationEnumRef("")
	}

	return HealthCheckHttp2HealthCheckPortSpecificationEnumRef(s)
}

// flattenHealthCheckHttp2HealthCheckProxyHeaderEnumSlice flattens the contents of HealthCheckHttp2HealthCheckProxyHeaderEnum from a JSON
// response object.
func flattenHealthCheckHttp2HealthCheckProxyHeaderEnumSlice(c *Client, i interface{}) []HealthCheckHttp2HealthCheckProxyHeaderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttp2HealthCheckProxyHeaderEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckHttp2HealthCheckProxyHeaderEnum{}
	}

	items := make([]HealthCheckHttp2HealthCheckProxyHeaderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttp2HealthCheckProxyHeaderEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckHttp2HealthCheckProxyHeaderEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckHttp2HealthCheckProxyHeaderEnum with the same value as that string.
func flattenHealthCheckHttp2HealthCheckProxyHeaderEnum(i interface{}) *HealthCheckHttp2HealthCheckProxyHeaderEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckHttp2HealthCheckProxyHeaderEnumRef("")
	}

	return HealthCheckHttp2HealthCheckProxyHeaderEnumRef(s)
}

// flattenHealthCheckHttpHealthCheckPortSpecificationEnumSlice flattens the contents of HealthCheckHttpHealthCheckPortSpecificationEnum from a JSON
// response object.
func flattenHealthCheckHttpHealthCheckPortSpecificationEnumSlice(c *Client, i interface{}) []HealthCheckHttpHealthCheckPortSpecificationEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttpHealthCheckPortSpecificationEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckHttpHealthCheckPortSpecificationEnum{}
	}

	items := make([]HealthCheckHttpHealthCheckPortSpecificationEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttpHealthCheckPortSpecificationEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckHttpHealthCheckPortSpecificationEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckHttpHealthCheckPortSpecificationEnum with the same value as that string.
func flattenHealthCheckHttpHealthCheckPortSpecificationEnum(i interface{}) *HealthCheckHttpHealthCheckPortSpecificationEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckHttpHealthCheckPortSpecificationEnumRef("")
	}

	return HealthCheckHttpHealthCheckPortSpecificationEnumRef(s)
}

// flattenHealthCheckHttpHealthCheckProxyHeaderEnumSlice flattens the contents of HealthCheckHttpHealthCheckProxyHeaderEnum from a JSON
// response object.
func flattenHealthCheckHttpHealthCheckProxyHeaderEnumSlice(c *Client, i interface{}) []HealthCheckHttpHealthCheckProxyHeaderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttpHealthCheckProxyHeaderEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckHttpHealthCheckProxyHeaderEnum{}
	}

	items := make([]HealthCheckHttpHealthCheckProxyHeaderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttpHealthCheckProxyHeaderEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckHttpHealthCheckProxyHeaderEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckHttpHealthCheckProxyHeaderEnum with the same value as that string.
func flattenHealthCheckHttpHealthCheckProxyHeaderEnum(i interface{}) *HealthCheckHttpHealthCheckProxyHeaderEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckHttpHealthCheckProxyHeaderEnumRef("")
	}

	return HealthCheckHttpHealthCheckProxyHeaderEnumRef(s)
}

// flattenHealthCheckHttpsHealthCheckPortSpecificationEnumSlice flattens the contents of HealthCheckHttpsHealthCheckPortSpecificationEnum from a JSON
// response object.
func flattenHealthCheckHttpsHealthCheckPortSpecificationEnumSlice(c *Client, i interface{}) []HealthCheckHttpsHealthCheckPortSpecificationEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttpsHealthCheckPortSpecificationEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckHttpsHealthCheckPortSpecificationEnum{}
	}

	items := make([]HealthCheckHttpsHealthCheckPortSpecificationEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttpsHealthCheckPortSpecificationEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckHttpsHealthCheckPortSpecificationEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckHttpsHealthCheckPortSpecificationEnum with the same value as that string.
func flattenHealthCheckHttpsHealthCheckPortSpecificationEnum(i interface{}) *HealthCheckHttpsHealthCheckPortSpecificationEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckHttpsHealthCheckPortSpecificationEnumRef("")
	}

	return HealthCheckHttpsHealthCheckPortSpecificationEnumRef(s)
}

// flattenHealthCheckHttpsHealthCheckProxyHeaderEnumSlice flattens the contents of HealthCheckHttpsHealthCheckProxyHeaderEnum from a JSON
// response object.
func flattenHealthCheckHttpsHealthCheckProxyHeaderEnumSlice(c *Client, i interface{}) []HealthCheckHttpsHealthCheckProxyHeaderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckHttpsHealthCheckProxyHeaderEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckHttpsHealthCheckProxyHeaderEnum{}
	}

	items := make([]HealthCheckHttpsHealthCheckProxyHeaderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckHttpsHealthCheckProxyHeaderEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckHttpsHealthCheckProxyHeaderEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckHttpsHealthCheckProxyHeaderEnum with the same value as that string.
func flattenHealthCheckHttpsHealthCheckProxyHeaderEnum(i interface{}) *HealthCheckHttpsHealthCheckProxyHeaderEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckHttpsHealthCheckProxyHeaderEnumRef("")
	}

	return HealthCheckHttpsHealthCheckProxyHeaderEnumRef(s)
}

// flattenHealthCheckSslHealthCheckPortSpecificationEnumSlice flattens the contents of HealthCheckSslHealthCheckPortSpecificationEnum from a JSON
// response object.
func flattenHealthCheckSslHealthCheckPortSpecificationEnumSlice(c *Client, i interface{}) []HealthCheckSslHealthCheckPortSpecificationEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckSslHealthCheckPortSpecificationEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckSslHealthCheckPortSpecificationEnum{}
	}

	items := make([]HealthCheckSslHealthCheckPortSpecificationEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckSslHealthCheckPortSpecificationEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckSslHealthCheckPortSpecificationEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckSslHealthCheckPortSpecificationEnum with the same value as that string.
func flattenHealthCheckSslHealthCheckPortSpecificationEnum(i interface{}) *HealthCheckSslHealthCheckPortSpecificationEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckSslHealthCheckPortSpecificationEnumRef("")
	}

	return HealthCheckSslHealthCheckPortSpecificationEnumRef(s)
}

// flattenHealthCheckSslHealthCheckProxyHeaderEnumSlice flattens the contents of HealthCheckSslHealthCheckProxyHeaderEnum from a JSON
// response object.
func flattenHealthCheckSslHealthCheckProxyHeaderEnumSlice(c *Client, i interface{}) []HealthCheckSslHealthCheckProxyHeaderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckSslHealthCheckProxyHeaderEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckSslHealthCheckProxyHeaderEnum{}
	}

	items := make([]HealthCheckSslHealthCheckProxyHeaderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckSslHealthCheckProxyHeaderEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckSslHealthCheckProxyHeaderEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckSslHealthCheckProxyHeaderEnum with the same value as that string.
func flattenHealthCheckSslHealthCheckProxyHeaderEnum(i interface{}) *HealthCheckSslHealthCheckProxyHeaderEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckSslHealthCheckProxyHeaderEnumRef("")
	}

	return HealthCheckSslHealthCheckProxyHeaderEnumRef(s)
}

// flattenHealthCheckTcpHealthCheckPortSpecificationEnumSlice flattens the contents of HealthCheckTcpHealthCheckPortSpecificationEnum from a JSON
// response object.
func flattenHealthCheckTcpHealthCheckPortSpecificationEnumSlice(c *Client, i interface{}) []HealthCheckTcpHealthCheckPortSpecificationEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckTcpHealthCheckPortSpecificationEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckTcpHealthCheckPortSpecificationEnum{}
	}

	items := make([]HealthCheckTcpHealthCheckPortSpecificationEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckTcpHealthCheckPortSpecificationEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckTcpHealthCheckPortSpecificationEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckTcpHealthCheckPortSpecificationEnum with the same value as that string.
func flattenHealthCheckTcpHealthCheckPortSpecificationEnum(i interface{}) *HealthCheckTcpHealthCheckPortSpecificationEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckTcpHealthCheckPortSpecificationEnumRef("")
	}

	return HealthCheckTcpHealthCheckPortSpecificationEnumRef(s)
}

// flattenHealthCheckTcpHealthCheckProxyHeaderEnumSlice flattens the contents of HealthCheckTcpHealthCheckProxyHeaderEnum from a JSON
// response object.
func flattenHealthCheckTcpHealthCheckProxyHeaderEnumSlice(c *Client, i interface{}) []HealthCheckTcpHealthCheckProxyHeaderEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckTcpHealthCheckProxyHeaderEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckTcpHealthCheckProxyHeaderEnum{}
	}

	items := make([]HealthCheckTcpHealthCheckProxyHeaderEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckTcpHealthCheckProxyHeaderEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckTcpHealthCheckProxyHeaderEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckTcpHealthCheckProxyHeaderEnum with the same value as that string.
func flattenHealthCheckTcpHealthCheckProxyHeaderEnum(i interface{}) *HealthCheckTcpHealthCheckProxyHeaderEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckTcpHealthCheckProxyHeaderEnumRef("")
	}

	return HealthCheckTcpHealthCheckProxyHeaderEnumRef(s)
}

// flattenHealthCheckTypeEnumSlice flattens the contents of HealthCheckTypeEnum from a JSON
// response object.
func flattenHealthCheckTypeEnumSlice(c *Client, i interface{}) []HealthCheckTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HealthCheckTypeEnum{}
	}

	if len(a) == 0 {
		return []HealthCheckTypeEnum{}
	}

	items := make([]HealthCheckTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHealthCheckTypeEnum(item.(interface{})))
	}

	return items
}

// flattenHealthCheckTypeEnum asserts that an interface is a string, and returns a
// pointer to a *HealthCheckTypeEnum with the same value as that string.
func flattenHealthCheckTypeEnum(i interface{}) *HealthCheckTypeEnum {
	s, ok := i.(string)
	if !ok {
		return HealthCheckTypeEnumRef("")
	}

	return HealthCheckTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *HealthCheck) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalHealthCheck(b, c)
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
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

func convertFieldDiffToHealthCheckDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]healthCheckDiff, error) {
	var diffs []healthCheckDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := healthCheckDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTohealthCheckApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameTohealthCheckApiOperation(op string, opts ...dcl.ApplyOption) (healthCheckApiOperation, error) {
	switch op {

	case "updateHealthCheckUpdateOperation":
		return &updateHealthCheckUpdateOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
