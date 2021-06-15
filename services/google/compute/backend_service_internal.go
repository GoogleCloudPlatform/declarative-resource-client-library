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

func (r *BackendService) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Iap", "CdnPolicy"}, r.Iap, r.CdnPolicy); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.FailoverPolicy) {
		if err := r.FailoverPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConnectionDraining) {
		if err := r.ConnectionDraining.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Iap) {
		if err := r.Iap.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CdnPolicy) {
		if err := r.CdnPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LogConfig) {
		if err := r.LogConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SecuritySettings) {
		if err := r.SecuritySettings.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConsistentHash) {
		if err := r.ConsistentHash.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CircuitBreakers) {
		if err := r.CircuitBreakers.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.OutlierDetection) {
		if err := r.OutlierDetection.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaxStreamDuration) {
		if err := r.MaxStreamDuration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BackendServiceBackends) validate() error {
	return nil
}
func (r *BackendServiceFailoverPolicy) validate() error {
	return nil
}
func (r *BackendServiceConnectionDraining) validate() error {
	return nil
}
func (r *BackendServiceIap) validate() error {
	return nil
}
func (r *BackendServiceCdnPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CacheKeyPolicy) {
		if err := r.CacheKeyPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BackendServiceCdnPolicyCacheKeyPolicy) validate() error {
	return nil
}
func (r *BackendServiceCdnPolicyNegativeCachingPolicy) validate() error {
	return nil
}
func (r *BackendServiceCdnPolicyBypassCacheOnRequestHeaders) validate() error {
	return nil
}
func (r *BackendServiceLogConfig) validate() error {
	return nil
}
func (r *BackendServiceSecuritySettings) validate() error {
	return nil
}
func (r *BackendServiceConsistentHash) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HttpCookie) {
		if err := r.HttpCookie.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BackendServiceConsistentHashHttpCookie) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Ttl) {
		if err := r.Ttl.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BackendServiceConsistentHashHttpCookieTtl) validate() error {
	return nil
}
func (r *BackendServiceCircuitBreakers) validate() error {
	return nil
}
func (r *BackendServiceOutlierDetection) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Interval) {
		if err := r.Interval.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BaseEjectionTime) {
		if err := r.BaseEjectionTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BackendServiceOutlierDetectionInterval) validate() error {
	return nil
}
func (r *BackendServiceOutlierDetectionBaseEjectionTime) validate() error {
	return nil
}
func (r *BackendServiceMaxStreamDuration) validate() error {
	return nil
}

func backendServiceGetURL(userBasePath string, r *BackendService) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/backendServices/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/backendServices/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func backendServiceListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/backendServices", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/backendServices", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func backendServiceCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/backendServices", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/backendServices", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func backendServiceDeleteURL(userBasePath string, r *BackendService) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/backendServices/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
	}

	return dcl.URL("projects/{{project}}/global/backendServices/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// backendServiceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type backendServiceApiOperation interface {
	do(context.Context, *BackendService, *Client) error
}

// newUpdateBackendServiceUpdateRequest creates a request for an
// BackendService resource's Update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBackendServiceUpdateRequest(ctx context.Context, f *BackendService, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandBackendServiceBackendsSlice(c, f.Backends); err != nil {
		return nil, fmt.Errorf("error expanding Backends into backends: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["backends"] = v
	}
	if v := f.HealthChecks; !dcl.IsEmptyValueIndirect(v) {
		req["healthChecks"] = v
	}
	if v := f.TimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		req["timeoutSec"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		req["port"] = v
	}
	if v := f.Protocol; !dcl.IsEmptyValueIndirect(v) {
		req["protocol"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		req["portName"] = v
	}
	if v := f.EnableCdn; !dcl.IsEmptyValueIndirect(v) {
		req["enableCDN"] = v
	}
	if v := f.SessionAffinity; !dcl.IsEmptyValueIndirect(v) {
		req["sessionAffinity"] = v
	}
	if v := f.AffinityCookieTtlSec; !dcl.IsEmptyValueIndirect(v) {
		req["affinityCookieTtlSec"] = v
	}
	if v, err := expandBackendServiceFailoverPolicy(c, f.FailoverPolicy); err != nil {
		return nil, fmt.Errorf("error expanding FailoverPolicy into failoverPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["failoverPolicy"] = v
	}
	if v := f.LoadBalancingScheme; !dcl.IsEmptyValueIndirect(v) {
		req["loadBalancingScheme"] = v
	}
	if v, err := expandBackendServiceConnectionDraining(c, f.ConnectionDraining); err != nil {
		return nil, fmt.Errorf("error expanding ConnectionDraining into connectionDraining: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["connectionDraining"] = v
	}
	if v, err := expandBackendServiceIap(c, f.Iap); err != nil {
		return nil, fmt.Errorf("error expanding Iap into iap: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["iap"] = v
	}
	if v, err := expandBackendServiceCdnPolicy(c, f.CdnPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CdnPolicy into cdnPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["cdnPolicy"] = v
	}
	if v := f.CustomRequestHeaders; !dcl.IsEmptyValueIndirect(v) {
		req["customRequestHeaders"] = v
	}
	if v := f.CustomResponseHeaders; !dcl.IsEmptyValueIndirect(v) {
		req["customResponseHeaders"] = v
	}
	if v, err := expandBackendServiceLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["logConfig"] = v
	}
	if v := f.LocalityLbPolicy; !dcl.IsEmptyValueIndirect(v) {
		req["localityLbPolicy"] = v
	}
	if v, err := expandBackendServiceConsistentHash(c, f.ConsistentHash); err != nil {
		return nil, fmt.Errorf("error expanding ConsistentHash into consistentHash: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["consistentHash"] = v
	}
	if v, err := expandBackendServiceCircuitBreakers(c, f.CircuitBreakers); err != nil {
		return nil, fmt.Errorf("error expanding CircuitBreakers into circuitBreakers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["circuitBreakers"] = v
	}
	if v, err := expandBackendServiceOutlierDetection(c, f.OutlierDetection); err != nil {
		return nil, fmt.Errorf("error expanding OutlierDetection into outlierDetection: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["outlierDetection"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		req["network"] = v
	}
	if v, err := expandBackendServiceMaxStreamDuration(c, f.MaxStreamDuration); err != nil {
		return nil, fmt.Errorf("error expanding MaxStreamDuration into maxStreamDuration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["maxStreamDuration"] = v
	}
	b, err := c.getBackendServiceRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"fingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["fingerprint"] = rawFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateBackendServiceUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateBackendServiceUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBackendServiceUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBackendServiceUpdateOperation) do(ctx context.Context, r *BackendService, c *Client) error {
	_, err := c.GetBackendService(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Update")
	if err != nil {
		return err
	}

	req, err := newUpdateBackendServiceUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateBackendServiceUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listBackendServiceRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := backendServiceListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BackendServiceMaxPage {
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

type listBackendServiceOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listBackendService(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*BackendService, string, error) {
	b, err := c.listBackendServiceRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBackendServiceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*BackendService
	for _, v := range m.Items {
		res, err := unmarshalMapBackendService(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBackendService(ctx context.Context, f func(*BackendService) bool, resources []*BackendService) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBackendService(ctx, res)
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

type deleteBackendServiceOperation struct{}

func (op *deleteBackendServiceOperation) do(ctx context.Context, r *BackendService, c *Client) error {

	_, err := c.GetBackendService(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("BackendService not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetBackendService checking for existence. error: %v", err)
		return err
	}

	u, err := backendServiceDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetBackendService(ctx, r.urlNormalized())
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
type createBackendServiceOperation struct {
	response map[string]interface{}
}

func (op *createBackendServiceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBackendServiceOperation) do(ctx context.Context, r *BackendService, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := backendServiceCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetBackendService(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBackendServiceRaw(ctx context.Context, r *BackendService) ([]byte, error) {

	u, err := backendServiceGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) backendServiceDiffsForRawDesired(ctx context.Context, rawDesired *BackendService, opts ...dcl.ApplyOption) (initial, desired *BackendService, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *BackendService
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*BackendService); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected BackendService, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBackendService(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a BackendService resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve BackendService resource: %v", err)
		}
		c.Config.Logger.Info("Found that BackendService resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBackendServiceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for BackendService: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for BackendService: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBackendServiceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for BackendService: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBackendServiceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for BackendService: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBackendService(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBackendServiceInitialState(rawInitial, rawDesired *BackendService) (*BackendService, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.Iap) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.CdnPolicy) {
			rawInitial.Iap = EmptyBackendServiceIap
		}
	}

	if dcl.IsZeroValue(rawInitial.CdnPolicy) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Iap) {
			rawInitial.CdnPolicy = EmptyBackendServiceCdnPolicy
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

func canonicalizeBackendServiceDesiredState(rawDesired, rawInitial *BackendService, opts ...dcl.ApplyOption) (*BackendService, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.FailoverPolicy = canonicalizeBackendServiceFailoverPolicy(rawDesired.FailoverPolicy, nil, opts...)
		rawDesired.ConnectionDraining = canonicalizeBackendServiceConnectionDraining(rawDesired.ConnectionDraining, nil, opts...)
		rawDesired.Iap = canonicalizeBackendServiceIap(rawDesired.Iap, nil, opts...)
		rawDesired.CdnPolicy = canonicalizeBackendServiceCdnPolicy(rawDesired.CdnPolicy, nil, opts...)
		rawDesired.LogConfig = canonicalizeBackendServiceLogConfig(rawDesired.LogConfig, nil, opts...)
		rawDesired.SecuritySettings = canonicalizeBackendServiceSecuritySettings(rawDesired.SecuritySettings, nil, opts...)
		rawDesired.ConsistentHash = canonicalizeBackendServiceConsistentHash(rawDesired.ConsistentHash, nil, opts...)
		rawDesired.CircuitBreakers = canonicalizeBackendServiceCircuitBreakers(rawDesired.CircuitBreakers, nil, opts...)
		rawDesired.OutlierDetection = canonicalizeBackendServiceOutlierDetection(rawDesired.OutlierDetection, nil, opts...)
		rawDesired.MaxStreamDuration = canonicalizeBackendServiceMaxStreamDuration(rawDesired.MaxStreamDuration, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.Iap != nil || rawInitial.Iap != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.CdnPolicy) {
			rawDesired.Iap = nil
			rawInitial.Iap = nil
		}
	}

	if rawDesired.CdnPolicy != nil || rawInitial.CdnPolicy != nil {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Iap) {
			rawDesired.CdnPolicy = nil
			rawInitial.CdnPolicy = nil
		}
	}

	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Backends) {
		rawDesired.Backends = rawInitial.Backends
	}
	if dcl.IsZeroValue(rawDesired.HealthChecks) {
		rawDesired.HealthChecks = rawInitial.HealthChecks
	}
	if dcl.IsZeroValue(rawDesired.TimeoutSec) {
		rawDesired.TimeoutSec = rawInitial.TimeoutSec
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = rawInitial.Port
	}
	if dcl.IsZeroValue(rawDesired.Protocol) {
		rawDesired.Protocol = rawInitial.Protocol
	}
	if dcl.StringCanonicalize(rawDesired.PortName, rawInitial.PortName) {
		rawDesired.PortName = rawInitial.PortName
	}
	if dcl.BoolCanonicalize(rawDesired.EnableCdn, rawInitial.EnableCdn) {
		rawDesired.EnableCdn = rawInitial.EnableCdn
	}
	if dcl.IsZeroValue(rawDesired.SessionAffinity) {
		rawDesired.SessionAffinity = rawInitial.SessionAffinity
	}
	if dcl.IsZeroValue(rawDesired.AffinityCookieTtlSec) {
		rawDesired.AffinityCookieTtlSec = rawInitial.AffinityCookieTtlSec
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	rawDesired.FailoverPolicy = canonicalizeBackendServiceFailoverPolicy(rawDesired.FailoverPolicy, rawInitial.FailoverPolicy, opts...)
	if dcl.IsZeroValue(rawDesired.LoadBalancingScheme) {
		rawDesired.LoadBalancingScheme = rawInitial.LoadBalancingScheme
	}
	rawDesired.ConnectionDraining = canonicalizeBackendServiceConnectionDraining(rawDesired.ConnectionDraining, rawInitial.ConnectionDraining, opts...)
	rawDesired.Iap = canonicalizeBackendServiceIap(rawDesired.Iap, rawInitial.Iap, opts...)
	rawDesired.CdnPolicy = canonicalizeBackendServiceCdnPolicy(rawDesired.CdnPolicy, rawInitial.CdnPolicy, opts...)
	if dcl.IsZeroValue(rawDesired.CustomRequestHeaders) {
		rawDesired.CustomRequestHeaders = rawInitial.CustomRequestHeaders
	}
	if dcl.IsZeroValue(rawDesired.CustomResponseHeaders) {
		rawDesired.CustomResponseHeaders = rawInitial.CustomResponseHeaders
	}
	if dcl.NameToSelfLink(rawDesired.SecurityPolicy, rawInitial.SecurityPolicy) {
		rawDesired.SecurityPolicy = rawInitial.SecurityPolicy
	}
	rawDesired.LogConfig = canonicalizeBackendServiceLogConfig(rawDesired.LogConfig, rawInitial.LogConfig, opts...)
	rawDesired.SecuritySettings = canonicalizeBackendServiceSecuritySettings(rawDesired.SecuritySettings, rawInitial.SecuritySettings, opts...)
	if dcl.IsZeroValue(rawDesired.LocalityLbPolicy) {
		rawDesired.LocalityLbPolicy = rawInitial.LocalityLbPolicy
	}
	rawDesired.ConsistentHash = canonicalizeBackendServiceConsistentHash(rawDesired.ConsistentHash, rawInitial.ConsistentHash, opts...)
	rawDesired.CircuitBreakers = canonicalizeBackendServiceCircuitBreakers(rawDesired.CircuitBreakers, rawInitial.CircuitBreakers, opts...)
	rawDesired.OutlierDetection = canonicalizeBackendServiceOutlierDetection(rawDesired.OutlierDetection, rawInitial.OutlierDetection, opts...)
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	rawDesired.MaxStreamDuration = canonicalizeBackendServiceMaxStreamDuration(rawDesired.MaxStreamDuration, rawInitial.MaxStreamDuration, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeBackendServiceNewState(c *Client, rawNew, rawDesired *BackendService) (*BackendService, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLinkWithId) && dcl.IsEmptyValueIndirect(rawDesired.SelfLinkWithId) {
		rawNew.SelfLinkWithId = rawDesired.SelfLinkWithId
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLinkWithId, rawNew.SelfLinkWithId) {
			rawNew.SelfLinkWithId = rawDesired.SelfLinkWithId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Backends) && dcl.IsEmptyValueIndirect(rawDesired.Backends) {
		rawNew.Backends = rawDesired.Backends
	} else {
		rawNew.Backends = canonicalizeNewBackendServiceBackendsSlice(c, rawDesired.Backends, rawNew.Backends)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthChecks) && dcl.IsEmptyValueIndirect(rawDesired.HealthChecks) {
		rawNew.HealthChecks = rawDesired.HealthChecks
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeoutSec) && dcl.IsEmptyValueIndirect(rawDesired.TimeoutSec) {
		rawNew.TimeoutSec = rawDesired.TimeoutSec
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Protocol) && dcl.IsEmptyValueIndirect(rawDesired.Protocol) {
		rawNew.Protocol = rawDesired.Protocol
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Fingerprint) && dcl.IsEmptyValueIndirect(rawDesired.Fingerprint) {
		rawNew.Fingerprint = rawDesired.Fingerprint
	} else {
		if dcl.StringCanonicalize(rawDesired.Fingerprint, rawNew.Fingerprint) {
			rawNew.Fingerprint = rawDesired.Fingerprint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PortName) && dcl.IsEmptyValueIndirect(rawDesired.PortName) {
		rawNew.PortName = rawDesired.PortName
	} else {
		if dcl.StringCanonicalize(rawDesired.PortName, rawNew.PortName) {
			rawNew.PortName = rawDesired.PortName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableCdn) && dcl.IsEmptyValueIndirect(rawDesired.EnableCdn) {
		rawNew.EnableCdn = rawDesired.EnableCdn
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableCdn, rawNew.EnableCdn) {
			rawNew.EnableCdn = rawDesired.EnableCdn
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SessionAffinity) && dcl.IsEmptyValueIndirect(rawDesired.SessionAffinity) {
		rawNew.SessionAffinity = rawDesired.SessionAffinity
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AffinityCookieTtlSec) && dcl.IsEmptyValueIndirect(rawDesired.AffinityCookieTtlSec) {
		rawNew.AffinityCookieTtlSec = rawDesired.AffinityCookieTtlSec
	} else {
	}

	rawNew.Location = rawDesired.Location

	if dcl.IsEmptyValueIndirect(rawNew.FailoverPolicy) && dcl.IsEmptyValueIndirect(rawDesired.FailoverPolicy) {
		rawNew.FailoverPolicy = rawDesired.FailoverPolicy
	} else {
		rawNew.FailoverPolicy = canonicalizeNewBackendServiceFailoverPolicy(c, rawDesired.FailoverPolicy, rawNew.FailoverPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LoadBalancingScheme) && dcl.IsEmptyValueIndirect(rawDesired.LoadBalancingScheme) {
		rawNew.LoadBalancingScheme = rawDesired.LoadBalancingScheme
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConnectionDraining) && dcl.IsEmptyValueIndirect(rawDesired.ConnectionDraining) {
		rawNew.ConnectionDraining = rawDesired.ConnectionDraining
	} else {
		rawNew.ConnectionDraining = canonicalizeNewBackendServiceConnectionDraining(c, rawDesired.ConnectionDraining, rawNew.ConnectionDraining)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Iap) && dcl.IsEmptyValueIndirect(rawDesired.Iap) {
		rawNew.Iap = rawDesired.Iap
	} else {
		rawNew.Iap = canonicalizeNewBackendServiceIap(c, rawDesired.Iap, rawNew.Iap)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CdnPolicy) && dcl.IsEmptyValueIndirect(rawDesired.CdnPolicy) {
		rawNew.CdnPolicy = rawDesired.CdnPolicy
	} else {
		rawNew.CdnPolicy = canonicalizeNewBackendServiceCdnPolicy(c, rawDesired.CdnPolicy, rawNew.CdnPolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomRequestHeaders) && dcl.IsEmptyValueIndirect(rawDesired.CustomRequestHeaders) {
		rawNew.CustomRequestHeaders = rawDesired.CustomRequestHeaders
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomResponseHeaders) && dcl.IsEmptyValueIndirect(rawDesired.CustomResponseHeaders) {
		rawNew.CustomResponseHeaders = rawDesired.CustomResponseHeaders
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SecurityPolicy) && dcl.IsEmptyValueIndirect(rawDesired.SecurityPolicy) {
		rawNew.SecurityPolicy = rawDesired.SecurityPolicy
	} else {
		if dcl.NameToSelfLink(rawDesired.SecurityPolicy, rawNew.SecurityPolicy) {
			rawNew.SecurityPolicy = rawDesired.SecurityPolicy
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LogConfig) && dcl.IsEmptyValueIndirect(rawDesired.LogConfig) {
		rawNew.LogConfig = rawDesired.LogConfig
	} else {
		rawNew.LogConfig = canonicalizeNewBackendServiceLogConfig(c, rawDesired.LogConfig, rawNew.LogConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SecuritySettings) && dcl.IsEmptyValueIndirect(rawDesired.SecuritySettings) {
		rawNew.SecuritySettings = rawDesired.SecuritySettings
	} else {
		rawNew.SecuritySettings = canonicalizeNewBackendServiceSecuritySettings(c, rawDesired.SecuritySettings, rawNew.SecuritySettings)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocalityLbPolicy) && dcl.IsEmptyValueIndirect(rawDesired.LocalityLbPolicy) {
		rawNew.LocalityLbPolicy = rawDesired.LocalityLbPolicy
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConsistentHash) && dcl.IsEmptyValueIndirect(rawDesired.ConsistentHash) {
		rawNew.ConsistentHash = rawDesired.ConsistentHash
	} else {
		rawNew.ConsistentHash = canonicalizeNewBackendServiceConsistentHash(c, rawDesired.ConsistentHash, rawNew.ConsistentHash)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CircuitBreakers) && dcl.IsEmptyValueIndirect(rawDesired.CircuitBreakers) {
		rawNew.CircuitBreakers = rawDesired.CircuitBreakers
	} else {
		rawNew.CircuitBreakers = canonicalizeNewBackendServiceCircuitBreakers(c, rawDesired.CircuitBreakers, rawNew.CircuitBreakers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.OutlierDetection) && dcl.IsEmptyValueIndirect(rawDesired.OutlierDetection) {
		rawNew.OutlierDetection = rawDesired.OutlierDetection
	} else {
		rawNew.OutlierDetection = canonicalizeNewBackendServiceOutlierDetection(c, rawDesired.OutlierDetection, rawNew.OutlierDetection)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxStreamDuration) && dcl.IsEmptyValueIndirect(rawDesired.MaxStreamDuration) {
		rawNew.MaxStreamDuration = rawDesired.MaxStreamDuration
	} else {
		rawNew.MaxStreamDuration = canonicalizeNewBackendServiceMaxStreamDuration(c, rawDesired.MaxStreamDuration, rawNew.MaxStreamDuration)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeBackendServiceBackends(des, initial *BackendServiceBackends, opts ...dcl.ApplyOption) *BackendServiceBackends {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.NameToSelfLink(des.Group, initial.Group) || dcl.IsZeroValue(des.Group) {
		des.Group = initial.Group
	}
	if dcl.IsZeroValue(des.BalancingMode) {
		des.BalancingMode = initial.BalancingMode
	}
	if dcl.IsZeroValue(des.MaxUtilization) {
		des.MaxUtilization = initial.MaxUtilization
	}
	if dcl.IsZeroValue(des.MaxRate) {
		des.MaxRate = initial.MaxRate
	}
	if dcl.IsZeroValue(des.MaxRatePerInstance) {
		des.MaxRatePerInstance = initial.MaxRatePerInstance
	}
	if dcl.IsZeroValue(des.MaxRatePerEndpoint) {
		des.MaxRatePerEndpoint = initial.MaxRatePerEndpoint
	}
	if dcl.IsZeroValue(des.MaxConnections) {
		des.MaxConnections = initial.MaxConnections
	}
	if dcl.IsZeroValue(des.MaxConnectionsPerInstance) {
		des.MaxConnectionsPerInstance = initial.MaxConnectionsPerInstance
	}
	if dcl.IsZeroValue(des.MaxConnectionsPerEndpoint) {
		des.MaxConnectionsPerEndpoint = initial.MaxConnectionsPerEndpoint
	}
	if dcl.IsZeroValue(des.CapacityScaler) {
		des.CapacityScaler = initial.CapacityScaler
	}
	if dcl.BoolCanonicalize(des.Failover, initial.Failover) || dcl.IsZeroValue(des.Failover) {
		des.Failover = initial.Failover
	}

	return des
}

func canonicalizeNewBackendServiceBackends(c *Client, des, nw *BackendServiceBackends) *BackendServiceBackends {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.NameToSelfLink(des.Group, nw.Group) {
		nw.Group = des.Group
	}
	if dcl.IsZeroValue(nw.BalancingMode) {
		nw.BalancingMode = des.BalancingMode
	}
	if dcl.IsZeroValue(nw.MaxUtilization) {
		nw.MaxUtilization = des.MaxUtilization
	}
	if dcl.IsZeroValue(nw.MaxRate) {
		nw.MaxRate = des.MaxRate
	}
	if dcl.IsZeroValue(nw.MaxRatePerInstance) {
		nw.MaxRatePerInstance = des.MaxRatePerInstance
	}
	if dcl.IsZeroValue(nw.MaxRatePerEndpoint) {
		nw.MaxRatePerEndpoint = des.MaxRatePerEndpoint
	}
	if dcl.IsZeroValue(nw.MaxConnections) {
		nw.MaxConnections = des.MaxConnections
	}
	if dcl.IsZeroValue(nw.MaxConnectionsPerInstance) {
		nw.MaxConnectionsPerInstance = des.MaxConnectionsPerInstance
	}
	if dcl.IsZeroValue(nw.MaxConnectionsPerEndpoint) {
		nw.MaxConnectionsPerEndpoint = des.MaxConnectionsPerEndpoint
	}
	if dcl.IsZeroValue(nw.CapacityScaler) {
		nw.CapacityScaler = des.CapacityScaler
	}
	if dcl.BoolCanonicalize(des.Failover, nw.Failover) {
		nw.Failover = des.Failover
	}

	return nw
}

func canonicalizeNewBackendServiceBackendsSet(c *Client, des, nw []BackendServiceBackends) []BackendServiceBackends {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceBackends
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceBackendsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceBackendsSlice(c *Client, des, nw []BackendServiceBackends) []BackendServiceBackends {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceBackends
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceBackends(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceFailoverPolicy(des, initial *BackendServiceFailoverPolicy, opts ...dcl.ApplyOption) *BackendServiceFailoverPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.DisableConnectionDrainOnFailover, initial.DisableConnectionDrainOnFailover) || dcl.IsZeroValue(des.DisableConnectionDrainOnFailover) {
		des.DisableConnectionDrainOnFailover = initial.DisableConnectionDrainOnFailover
	}
	if dcl.BoolCanonicalize(des.DropTrafficIfUnhealthy, initial.DropTrafficIfUnhealthy) || dcl.IsZeroValue(des.DropTrafficIfUnhealthy) {
		des.DropTrafficIfUnhealthy = initial.DropTrafficIfUnhealthy
	}
	if dcl.IsZeroValue(des.FailoverRatio) {
		des.FailoverRatio = initial.FailoverRatio
	}

	return des
}

func canonicalizeNewBackendServiceFailoverPolicy(c *Client, des, nw *BackendServiceFailoverPolicy) *BackendServiceFailoverPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.DisableConnectionDrainOnFailover, nw.DisableConnectionDrainOnFailover) {
		nw.DisableConnectionDrainOnFailover = des.DisableConnectionDrainOnFailover
	}
	if dcl.BoolCanonicalize(des.DropTrafficIfUnhealthy, nw.DropTrafficIfUnhealthy) {
		nw.DropTrafficIfUnhealthy = des.DropTrafficIfUnhealthy
	}
	if dcl.IsZeroValue(nw.FailoverRatio) {
		nw.FailoverRatio = des.FailoverRatio
	}

	return nw
}

func canonicalizeNewBackendServiceFailoverPolicySet(c *Client, des, nw []BackendServiceFailoverPolicy) []BackendServiceFailoverPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceFailoverPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceFailoverPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceFailoverPolicySlice(c *Client, des, nw []BackendServiceFailoverPolicy) []BackendServiceFailoverPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceFailoverPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceFailoverPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceConnectionDraining(des, initial *BackendServiceConnectionDraining, opts ...dcl.ApplyOption) *BackendServiceConnectionDraining {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DrainingTimeoutSec) {
		des.DrainingTimeoutSec = initial.DrainingTimeoutSec
	}

	return des
}

func canonicalizeNewBackendServiceConnectionDraining(c *Client, des, nw *BackendServiceConnectionDraining) *BackendServiceConnectionDraining {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.DrainingTimeoutSec) {
		nw.DrainingTimeoutSec = des.DrainingTimeoutSec
	}

	return nw
}

func canonicalizeNewBackendServiceConnectionDrainingSet(c *Client, des, nw []BackendServiceConnectionDraining) []BackendServiceConnectionDraining {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceConnectionDraining
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceConnectionDrainingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceConnectionDrainingSlice(c *Client, des, nw []BackendServiceConnectionDraining) []BackendServiceConnectionDraining {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceConnectionDraining
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceConnectionDraining(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceIap(des, initial *BackendServiceIap, opts ...dcl.ApplyOption) *BackendServiceIap {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.StringCanonicalize(des.OAuth2ClientId, initial.OAuth2ClientId) || dcl.IsZeroValue(des.OAuth2ClientId) {
		des.OAuth2ClientId = initial.OAuth2ClientId
	}
	if dcl.StringCanonicalize(des.OAuth2ClientSecret, initial.OAuth2ClientSecret) || dcl.IsZeroValue(des.OAuth2ClientSecret) {
		des.OAuth2ClientSecret = initial.OAuth2ClientSecret
	}

	return des
}

func canonicalizeNewBackendServiceIap(c *Client, des, nw *BackendServiceIap) *BackendServiceIap {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.StringCanonicalize(des.OAuth2ClientId, nw.OAuth2ClientId) {
		nw.OAuth2ClientId = des.OAuth2ClientId
	}
	nw.OAuth2ClientSecret = des.OAuth2ClientSecret
	if dcl.StringCanonicalize(des.OAuth2ClientSecretSha256, nw.OAuth2ClientSecretSha256) {
		nw.OAuth2ClientSecretSha256 = des.OAuth2ClientSecretSha256
	}

	return nw
}

func canonicalizeNewBackendServiceIapSet(c *Client, des, nw []BackendServiceIap) []BackendServiceIap {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceIap
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceIapNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceIapSlice(c *Client, des, nw []BackendServiceIap) []BackendServiceIap {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceIap
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceIap(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceCdnPolicy(des, initial *BackendServiceCdnPolicy, opts ...dcl.ApplyOption) *BackendServiceCdnPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.CacheKeyPolicy = canonicalizeBackendServiceCdnPolicyCacheKeyPolicy(des.CacheKeyPolicy, initial.CacheKeyPolicy, opts...)
	if dcl.IsZeroValue(des.SignedUrlCacheMaxAgeSec) {
		des.SignedUrlCacheMaxAgeSec = initial.SignedUrlCacheMaxAgeSec
	}
	if dcl.BoolCanonicalize(des.RequestCoalescing, initial.RequestCoalescing) || dcl.IsZeroValue(des.RequestCoalescing) {
		des.RequestCoalescing = initial.RequestCoalescing
	}
	if dcl.IsZeroValue(des.CacheMode) {
		des.CacheMode = initial.CacheMode
	}
	if dcl.IsZeroValue(des.DefaultTtl) {
		des.DefaultTtl = initial.DefaultTtl
	}
	if dcl.IsZeroValue(des.MaxTtl) {
		des.MaxTtl = initial.MaxTtl
	}
	if dcl.IsZeroValue(des.ClientTtl) {
		des.ClientTtl = initial.ClientTtl
	}
	if dcl.BoolCanonicalize(des.NegativeCaching, initial.NegativeCaching) || dcl.IsZeroValue(des.NegativeCaching) {
		des.NegativeCaching = initial.NegativeCaching
	}
	if dcl.IsZeroValue(des.NegativeCachingPolicy) {
		des.NegativeCachingPolicy = initial.NegativeCachingPolicy
	}
	if dcl.IsZeroValue(des.BypassCacheOnRequestHeaders) {
		des.BypassCacheOnRequestHeaders = initial.BypassCacheOnRequestHeaders
	}
	if dcl.IsZeroValue(des.ServeWhileStale) {
		des.ServeWhileStale = initial.ServeWhileStale
	}

	return des
}

func canonicalizeNewBackendServiceCdnPolicy(c *Client, des, nw *BackendServiceCdnPolicy) *BackendServiceCdnPolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.CacheKeyPolicy = canonicalizeNewBackendServiceCdnPolicyCacheKeyPolicy(c, des.CacheKeyPolicy, nw.CacheKeyPolicy)
	if dcl.IsZeroValue(nw.SignedUrlKeyNames) {
		nw.SignedUrlKeyNames = des.SignedUrlKeyNames
	}
	if dcl.IsZeroValue(nw.SignedUrlCacheMaxAgeSec) {
		nw.SignedUrlCacheMaxAgeSec = des.SignedUrlCacheMaxAgeSec
	}
	if dcl.BoolCanonicalize(des.RequestCoalescing, nw.RequestCoalescing) {
		nw.RequestCoalescing = des.RequestCoalescing
	}
	if dcl.IsZeroValue(nw.CacheMode) {
		nw.CacheMode = des.CacheMode
	}
	if dcl.IsZeroValue(nw.DefaultTtl) {
		nw.DefaultTtl = des.DefaultTtl
	}
	if dcl.IsZeroValue(nw.MaxTtl) {
		nw.MaxTtl = des.MaxTtl
	}
	if dcl.IsZeroValue(nw.ClientTtl) {
		nw.ClientTtl = des.ClientTtl
	}
	if dcl.BoolCanonicalize(des.NegativeCaching, nw.NegativeCaching) {
		nw.NegativeCaching = des.NegativeCaching
	}
	nw.NegativeCachingPolicy = canonicalizeNewBackendServiceCdnPolicyNegativeCachingPolicySlice(c, des.NegativeCachingPolicy, nw.NegativeCachingPolicy)
	nw.BypassCacheOnRequestHeaders = canonicalizeNewBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice(c, des.BypassCacheOnRequestHeaders, nw.BypassCacheOnRequestHeaders)
	if dcl.IsZeroValue(nw.ServeWhileStale) {
		nw.ServeWhileStale = des.ServeWhileStale
	}

	return nw
}

func canonicalizeNewBackendServiceCdnPolicySet(c *Client, des, nw []BackendServiceCdnPolicy) []BackendServiceCdnPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceCdnPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceCdnPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceCdnPolicySlice(c *Client, des, nw []BackendServiceCdnPolicy) []BackendServiceCdnPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceCdnPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceCdnPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceCdnPolicyCacheKeyPolicy(des, initial *BackendServiceCdnPolicyCacheKeyPolicy, opts ...dcl.ApplyOption) *BackendServiceCdnPolicyCacheKeyPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IncludeProtocol, initial.IncludeProtocol) || dcl.IsZeroValue(des.IncludeProtocol) {
		des.IncludeProtocol = initial.IncludeProtocol
	}
	if dcl.BoolCanonicalize(des.IncludeHost, initial.IncludeHost) || dcl.IsZeroValue(des.IncludeHost) {
		des.IncludeHost = initial.IncludeHost
	}
	if dcl.BoolCanonicalize(des.IncludeQueryString, initial.IncludeQueryString) || dcl.IsZeroValue(des.IncludeQueryString) {
		des.IncludeQueryString = initial.IncludeQueryString
	}
	if dcl.IsZeroValue(des.QueryStringWhitelist) {
		des.QueryStringWhitelist = initial.QueryStringWhitelist
	}
	if dcl.IsZeroValue(des.QueryStringBlacklist) {
		des.QueryStringBlacklist = initial.QueryStringBlacklist
	}
	if dcl.IsZeroValue(des.IncludeHttpHeaders) {
		des.IncludeHttpHeaders = initial.IncludeHttpHeaders
	}
	if dcl.IsZeroValue(des.IncludeNamedCookies) {
		des.IncludeNamedCookies = initial.IncludeNamedCookies
	}

	return des
}

func canonicalizeNewBackendServiceCdnPolicyCacheKeyPolicy(c *Client, des, nw *BackendServiceCdnPolicyCacheKeyPolicy) *BackendServiceCdnPolicyCacheKeyPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IncludeProtocol, nw.IncludeProtocol) {
		nw.IncludeProtocol = des.IncludeProtocol
	}
	if dcl.BoolCanonicalize(des.IncludeHost, nw.IncludeHost) {
		nw.IncludeHost = des.IncludeHost
	}
	if dcl.BoolCanonicalize(des.IncludeQueryString, nw.IncludeQueryString) {
		nw.IncludeQueryString = des.IncludeQueryString
	}
	if dcl.IsZeroValue(nw.QueryStringWhitelist) {
		nw.QueryStringWhitelist = des.QueryStringWhitelist
	}
	if dcl.IsZeroValue(nw.QueryStringBlacklist) {
		nw.QueryStringBlacklist = des.QueryStringBlacklist
	}
	if dcl.IsZeroValue(nw.IncludeHttpHeaders) {
		nw.IncludeHttpHeaders = des.IncludeHttpHeaders
	}
	if dcl.IsZeroValue(nw.IncludeNamedCookies) {
		nw.IncludeNamedCookies = des.IncludeNamedCookies
	}

	return nw
}

func canonicalizeNewBackendServiceCdnPolicyCacheKeyPolicySet(c *Client, des, nw []BackendServiceCdnPolicyCacheKeyPolicy) []BackendServiceCdnPolicyCacheKeyPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceCdnPolicyCacheKeyPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceCdnPolicyCacheKeyPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceCdnPolicyCacheKeyPolicySlice(c *Client, des, nw []BackendServiceCdnPolicyCacheKeyPolicy) []BackendServiceCdnPolicyCacheKeyPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceCdnPolicyCacheKeyPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceCdnPolicyCacheKeyPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceCdnPolicyNegativeCachingPolicy(des, initial *BackendServiceCdnPolicyNegativeCachingPolicy, opts ...dcl.ApplyOption) *BackendServiceCdnPolicyNegativeCachingPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.IsZeroValue(des.Ttl) {
		des.Ttl = initial.Ttl
	}

	return des
}

func canonicalizeNewBackendServiceCdnPolicyNegativeCachingPolicy(c *Client, des, nw *BackendServiceCdnPolicyNegativeCachingPolicy) *BackendServiceCdnPolicyNegativeCachingPolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Code) {
		nw.Code = des.Code
	}
	if dcl.IsZeroValue(nw.Ttl) {
		nw.Ttl = des.Ttl
	}

	return nw
}

func canonicalizeNewBackendServiceCdnPolicyNegativeCachingPolicySet(c *Client, des, nw []BackendServiceCdnPolicyNegativeCachingPolicy) []BackendServiceCdnPolicyNegativeCachingPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceCdnPolicyNegativeCachingPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceCdnPolicyNegativeCachingPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceCdnPolicyNegativeCachingPolicySlice(c *Client, des, nw []BackendServiceCdnPolicyNegativeCachingPolicy) []BackendServiceCdnPolicyNegativeCachingPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceCdnPolicyNegativeCachingPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceCdnPolicyNegativeCachingPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceCdnPolicyBypassCacheOnRequestHeaders(des, initial *BackendServiceCdnPolicyBypassCacheOnRequestHeaders, opts ...dcl.ApplyOption) *BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.HeaderName, initial.HeaderName) || dcl.IsZeroValue(des.HeaderName) {
		des.HeaderName = initial.HeaderName
	}

	return des
}

func canonicalizeNewBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c *Client, des, nw *BackendServiceCdnPolicyBypassCacheOnRequestHeaders) *BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HeaderName, nw.HeaderName) {
		nw.HeaderName = des.HeaderName
	}

	return nw
}

func canonicalizeNewBackendServiceCdnPolicyBypassCacheOnRequestHeadersSet(c *Client, des, nw []BackendServiceCdnPolicyBypassCacheOnRequestHeaders) []BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceCdnPolicyBypassCacheOnRequestHeaders
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceCdnPolicyBypassCacheOnRequestHeadersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice(c *Client, des, nw []BackendServiceCdnPolicyBypassCacheOnRequestHeaders) []BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceCdnPolicyBypassCacheOnRequestHeaders
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceLogConfig(des, initial *BackendServiceLogConfig, opts ...dcl.ApplyOption) *BackendServiceLogConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Enable, initial.Enable) || dcl.IsZeroValue(des.Enable) {
		des.Enable = initial.Enable
	}
	if dcl.IsZeroValue(des.SampleRate) {
		des.SampleRate = initial.SampleRate
	}

	return des
}

func canonicalizeNewBackendServiceLogConfig(c *Client, des, nw *BackendServiceLogConfig) *BackendServiceLogConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enable, nw.Enable) {
		nw.Enable = des.Enable
	}
	if dcl.IsZeroValue(nw.SampleRate) {
		nw.SampleRate = des.SampleRate
	}

	return nw
}

func canonicalizeNewBackendServiceLogConfigSet(c *Client, des, nw []BackendServiceLogConfig) []BackendServiceLogConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceLogConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceLogConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceLogConfigSlice(c *Client, des, nw []BackendServiceLogConfig) []BackendServiceLogConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceLogConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceLogConfig(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceSecuritySettings(des, initial *BackendServiceSecuritySettings, opts ...dcl.ApplyOption) *BackendServiceSecuritySettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.ClientTlsPolicy, initial.ClientTlsPolicy) || dcl.IsZeroValue(des.ClientTlsPolicy) {
		des.ClientTlsPolicy = initial.ClientTlsPolicy
	}
	if dcl.IsZeroValue(des.SubjectAltNames) {
		des.SubjectAltNames = initial.SubjectAltNames
	}

	return des
}

func canonicalizeNewBackendServiceSecuritySettings(c *Client, des, nw *BackendServiceSecuritySettings) *BackendServiceSecuritySettings {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ClientTlsPolicy, nw.ClientTlsPolicy) {
		nw.ClientTlsPolicy = des.ClientTlsPolicy
	}
	if dcl.IsZeroValue(nw.SubjectAltNames) {
		nw.SubjectAltNames = des.SubjectAltNames
	}

	return nw
}

func canonicalizeNewBackendServiceSecuritySettingsSet(c *Client, des, nw []BackendServiceSecuritySettings) []BackendServiceSecuritySettings {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceSecuritySettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceSecuritySettingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceSecuritySettingsSlice(c *Client, des, nw []BackendServiceSecuritySettings) []BackendServiceSecuritySettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceSecuritySettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceSecuritySettings(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceConsistentHash(des, initial *BackendServiceConsistentHash, opts ...dcl.ApplyOption) *BackendServiceConsistentHash {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.HttpCookie = canonicalizeBackendServiceConsistentHashHttpCookie(des.HttpCookie, initial.HttpCookie, opts...)
	if dcl.StringCanonicalize(des.HttpHeaderName, initial.HttpHeaderName) || dcl.IsZeroValue(des.HttpHeaderName) {
		des.HttpHeaderName = initial.HttpHeaderName
	}
	if dcl.IsZeroValue(des.MinimumRingSize) {
		des.MinimumRingSize = initial.MinimumRingSize
	}

	return des
}

func canonicalizeNewBackendServiceConsistentHash(c *Client, des, nw *BackendServiceConsistentHash) *BackendServiceConsistentHash {
	if des == nil || nw == nil {
		return nw
	}

	nw.HttpCookie = canonicalizeNewBackendServiceConsistentHashHttpCookie(c, des.HttpCookie, nw.HttpCookie)
	if dcl.StringCanonicalize(des.HttpHeaderName, nw.HttpHeaderName) {
		nw.HttpHeaderName = des.HttpHeaderName
	}
	if dcl.IsZeroValue(nw.MinimumRingSize) {
		nw.MinimumRingSize = des.MinimumRingSize
	}

	return nw
}

func canonicalizeNewBackendServiceConsistentHashSet(c *Client, des, nw []BackendServiceConsistentHash) []BackendServiceConsistentHash {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceConsistentHash
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceConsistentHashNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceConsistentHashSlice(c *Client, des, nw []BackendServiceConsistentHash) []BackendServiceConsistentHash {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceConsistentHash
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceConsistentHash(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceConsistentHashHttpCookie(des, initial *BackendServiceConsistentHashHttpCookie, opts ...dcl.ApplyOption) *BackendServiceConsistentHashHttpCookie {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	des.Ttl = canonicalizeBackendServiceConsistentHashHttpCookieTtl(des.Ttl, initial.Ttl, opts...)

	return des
}

func canonicalizeNewBackendServiceConsistentHashHttpCookie(c *Client, des, nw *BackendServiceConsistentHashHttpCookie) *BackendServiceConsistentHashHttpCookie {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	nw.Ttl = canonicalizeNewBackendServiceConsistentHashHttpCookieTtl(c, des.Ttl, nw.Ttl)

	return nw
}

func canonicalizeNewBackendServiceConsistentHashHttpCookieSet(c *Client, des, nw []BackendServiceConsistentHashHttpCookie) []BackendServiceConsistentHashHttpCookie {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceConsistentHashHttpCookie
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceConsistentHashHttpCookieNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceConsistentHashHttpCookieSlice(c *Client, des, nw []BackendServiceConsistentHashHttpCookie) []BackendServiceConsistentHashHttpCookie {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceConsistentHashHttpCookie
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceConsistentHashHttpCookie(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceConsistentHashHttpCookieTtl(des, initial *BackendServiceConsistentHashHttpCookieTtl, opts ...dcl.ApplyOption) *BackendServiceConsistentHashHttpCookieTtl {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewBackendServiceConsistentHashHttpCookieTtl(c *Client, des, nw *BackendServiceConsistentHashHttpCookieTtl) *BackendServiceConsistentHashHttpCookieTtl {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewBackendServiceConsistentHashHttpCookieTtlSet(c *Client, des, nw []BackendServiceConsistentHashHttpCookieTtl) []BackendServiceConsistentHashHttpCookieTtl {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceConsistentHashHttpCookieTtl
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceConsistentHashHttpCookieTtlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceConsistentHashHttpCookieTtlSlice(c *Client, des, nw []BackendServiceConsistentHashHttpCookieTtl) []BackendServiceConsistentHashHttpCookieTtl {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceConsistentHashHttpCookieTtl
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceConsistentHashHttpCookieTtl(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceCircuitBreakers(des, initial *BackendServiceCircuitBreakers, opts ...dcl.ApplyOption) *BackendServiceCircuitBreakers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxRequestsPerConnection) {
		des.MaxRequestsPerConnection = initial.MaxRequestsPerConnection
	}
	if dcl.IsZeroValue(des.MaxConnections) {
		des.MaxConnections = initial.MaxConnections
	}
	if dcl.IsZeroValue(des.MaxPendingRequests) {
		des.MaxPendingRequests = initial.MaxPendingRequests
	}
	if dcl.IsZeroValue(des.MaxRequests) {
		des.MaxRequests = initial.MaxRequests
	}
	if dcl.IsZeroValue(des.MaxRetries) {
		des.MaxRetries = initial.MaxRetries
	}

	return des
}

func canonicalizeNewBackendServiceCircuitBreakers(c *Client, des, nw *BackendServiceCircuitBreakers) *BackendServiceCircuitBreakers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.MaxRequestsPerConnection) {
		nw.MaxRequestsPerConnection = des.MaxRequestsPerConnection
	}
	if dcl.IsZeroValue(nw.MaxConnections) {
		nw.MaxConnections = des.MaxConnections
	}
	if dcl.IsZeroValue(nw.MaxPendingRequests) {
		nw.MaxPendingRequests = des.MaxPendingRequests
	}
	if dcl.IsZeroValue(nw.MaxRequests) {
		nw.MaxRequests = des.MaxRequests
	}
	if dcl.IsZeroValue(nw.MaxRetries) {
		nw.MaxRetries = des.MaxRetries
	}

	return nw
}

func canonicalizeNewBackendServiceCircuitBreakersSet(c *Client, des, nw []BackendServiceCircuitBreakers) []BackendServiceCircuitBreakers {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceCircuitBreakers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceCircuitBreakersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceCircuitBreakersSlice(c *Client, des, nw []BackendServiceCircuitBreakers) []BackendServiceCircuitBreakers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceCircuitBreakers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceCircuitBreakers(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceOutlierDetection(des, initial *BackendServiceOutlierDetection, opts ...dcl.ApplyOption) *BackendServiceOutlierDetection {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ConsecutiveErrors) {
		des.ConsecutiveErrors = initial.ConsecutiveErrors
	}
	des.Interval = canonicalizeBackendServiceOutlierDetectionInterval(des.Interval, initial.Interval, opts...)
	des.BaseEjectionTime = canonicalizeBackendServiceOutlierDetectionBaseEjectionTime(des.BaseEjectionTime, initial.BaseEjectionTime, opts...)
	if dcl.IsZeroValue(des.MaxEjectionPercent) {
		des.MaxEjectionPercent = initial.MaxEjectionPercent
	}
	if dcl.IsZeroValue(des.EnforcingConsecutiveErrors) {
		des.EnforcingConsecutiveErrors = initial.EnforcingConsecutiveErrors
	}
	if dcl.IsZeroValue(des.EnforcingSuccessRate) {
		des.EnforcingSuccessRate = initial.EnforcingSuccessRate
	}
	if dcl.IsZeroValue(des.SuccessRateMinimumHosts) {
		des.SuccessRateMinimumHosts = initial.SuccessRateMinimumHosts
	}
	if dcl.IsZeroValue(des.SuccessRateRequestVolume) {
		des.SuccessRateRequestVolume = initial.SuccessRateRequestVolume
	}
	if dcl.IsZeroValue(des.SuccessRateStdevFactor) {
		des.SuccessRateStdevFactor = initial.SuccessRateStdevFactor
	}
	if dcl.IsZeroValue(des.ConsecutiveGatewayFailure) {
		des.ConsecutiveGatewayFailure = initial.ConsecutiveGatewayFailure
	}
	if dcl.IsZeroValue(des.EnforcingConsecutiveGatewayFailure) {
		des.EnforcingConsecutiveGatewayFailure = initial.EnforcingConsecutiveGatewayFailure
	}

	return des
}

func canonicalizeNewBackendServiceOutlierDetection(c *Client, des, nw *BackendServiceOutlierDetection) *BackendServiceOutlierDetection {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ConsecutiveErrors) {
		nw.ConsecutiveErrors = des.ConsecutiveErrors
	}
	nw.Interval = canonicalizeNewBackendServiceOutlierDetectionInterval(c, des.Interval, nw.Interval)
	nw.BaseEjectionTime = canonicalizeNewBackendServiceOutlierDetectionBaseEjectionTime(c, des.BaseEjectionTime, nw.BaseEjectionTime)
	if dcl.IsZeroValue(nw.MaxEjectionPercent) {
		nw.MaxEjectionPercent = des.MaxEjectionPercent
	}
	if dcl.IsZeroValue(nw.EnforcingConsecutiveErrors) {
		nw.EnforcingConsecutiveErrors = des.EnforcingConsecutiveErrors
	}
	if dcl.IsZeroValue(nw.EnforcingSuccessRate) {
		nw.EnforcingSuccessRate = des.EnforcingSuccessRate
	}
	if dcl.IsZeroValue(nw.SuccessRateMinimumHosts) {
		nw.SuccessRateMinimumHosts = des.SuccessRateMinimumHosts
	}
	if dcl.IsZeroValue(nw.SuccessRateRequestVolume) {
		nw.SuccessRateRequestVolume = des.SuccessRateRequestVolume
	}
	if dcl.IsZeroValue(nw.SuccessRateStdevFactor) {
		nw.SuccessRateStdevFactor = des.SuccessRateStdevFactor
	}
	if dcl.IsZeroValue(nw.ConsecutiveGatewayFailure) {
		nw.ConsecutiveGatewayFailure = des.ConsecutiveGatewayFailure
	}
	if dcl.IsZeroValue(nw.EnforcingConsecutiveGatewayFailure) {
		nw.EnforcingConsecutiveGatewayFailure = des.EnforcingConsecutiveGatewayFailure
	}

	return nw
}

func canonicalizeNewBackendServiceOutlierDetectionSet(c *Client, des, nw []BackendServiceOutlierDetection) []BackendServiceOutlierDetection {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceOutlierDetection
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceOutlierDetectionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceOutlierDetectionSlice(c *Client, des, nw []BackendServiceOutlierDetection) []BackendServiceOutlierDetection {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceOutlierDetection
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceOutlierDetection(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceOutlierDetectionInterval(des, initial *BackendServiceOutlierDetectionInterval, opts ...dcl.ApplyOption) *BackendServiceOutlierDetectionInterval {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewBackendServiceOutlierDetectionInterval(c *Client, des, nw *BackendServiceOutlierDetectionInterval) *BackendServiceOutlierDetectionInterval {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewBackendServiceOutlierDetectionIntervalSet(c *Client, des, nw []BackendServiceOutlierDetectionInterval) []BackendServiceOutlierDetectionInterval {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceOutlierDetectionInterval
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceOutlierDetectionIntervalNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceOutlierDetectionIntervalSlice(c *Client, des, nw []BackendServiceOutlierDetectionInterval) []BackendServiceOutlierDetectionInterval {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceOutlierDetectionInterval
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceOutlierDetectionInterval(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceOutlierDetectionBaseEjectionTime(des, initial *BackendServiceOutlierDetectionBaseEjectionTime, opts ...dcl.ApplyOption) *BackendServiceOutlierDetectionBaseEjectionTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewBackendServiceOutlierDetectionBaseEjectionTime(c *Client, des, nw *BackendServiceOutlierDetectionBaseEjectionTime) *BackendServiceOutlierDetectionBaseEjectionTime {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewBackendServiceOutlierDetectionBaseEjectionTimeSet(c *Client, des, nw []BackendServiceOutlierDetectionBaseEjectionTime) []BackendServiceOutlierDetectionBaseEjectionTime {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceOutlierDetectionBaseEjectionTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceOutlierDetectionBaseEjectionTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceOutlierDetectionBaseEjectionTimeSlice(c *Client, des, nw []BackendServiceOutlierDetectionBaseEjectionTime) []BackendServiceOutlierDetectionBaseEjectionTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceOutlierDetectionBaseEjectionTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceOutlierDetectionBaseEjectionTime(c, &d, &n))
	}

	return items
}

func canonicalizeBackendServiceMaxStreamDuration(des, initial *BackendServiceMaxStreamDuration, opts ...dcl.ApplyOption) *BackendServiceMaxStreamDuration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewBackendServiceMaxStreamDuration(c *Client, des, nw *BackendServiceMaxStreamDuration) *BackendServiceMaxStreamDuration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Seconds) {
		nw.Seconds = des.Seconds
	}
	if dcl.IsZeroValue(nw.Nanos) {
		nw.Nanos = des.Nanos
	}

	return nw
}

func canonicalizeNewBackendServiceMaxStreamDurationSet(c *Client, des, nw []BackendServiceMaxStreamDuration) []BackendServiceMaxStreamDuration {
	if des == nil {
		return nw
	}
	var reorderedNew []BackendServiceMaxStreamDuration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareBackendServiceMaxStreamDurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewBackendServiceMaxStreamDurationSlice(c *Client, des, nw []BackendServiceMaxStreamDuration) []BackendServiceMaxStreamDuration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BackendServiceMaxStreamDuration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBackendServiceMaxStreamDuration(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffBackendService(c *Client, desired, actual *BackendService, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SelfLinkWithId, actual.SelfLinkWithId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLinkWithId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Backends, actual.Backends, dcl.Info{ObjectFunction: compareBackendServiceBackendsNewStyle, EmptyObject: EmptyBackendServiceBackends, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Backends")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HealthChecks, actual.HealthChecks, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("HealthChecks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeoutSec, actual.TimeoutSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("TimeoutSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Protocol, actual.Protocol, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Protocol")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fingerprint, actual.Fingerprint, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortName, actual.PortName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("PortName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableCdn, actual.EnableCdn, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("EnableCDN")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SessionAffinity, actual.SessionAffinity, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SessionAffinity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AffinityCookieTtlSec, actual.AffinityCookieTtlSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("AffinityCookieTtlSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailoverPolicy, actual.FailoverPolicy, dcl.Info{ObjectFunction: compareBackendServiceFailoverPolicyNewStyle, EmptyObject: EmptyBackendServiceFailoverPolicy, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("FailoverPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LoadBalancingScheme, actual.LoadBalancingScheme, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("LoadBalancingScheme")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConnectionDraining, actual.ConnectionDraining, dcl.Info{ObjectFunction: compareBackendServiceConnectionDrainingNewStyle, EmptyObject: EmptyBackendServiceConnectionDraining, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ConnectionDraining")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Iap, actual.Iap, dcl.Info{ObjectFunction: compareBackendServiceIapNewStyle, EmptyObject: EmptyBackendServiceIap, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Iap")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CdnPolicy, actual.CdnPolicy, dcl.Info{ObjectFunction: compareBackendServiceCdnPolicyNewStyle, EmptyObject: EmptyBackendServiceCdnPolicy, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CdnPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomRequestHeaders, actual.CustomRequestHeaders, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CustomRequestHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomResponseHeaders, actual.CustomResponseHeaders, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CustomResponseHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityPolicy, actual.SecurityPolicy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecurityPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogConfig, actual.LogConfig, dcl.Info{ObjectFunction: compareBackendServiceLogConfigNewStyle, EmptyObject: EmptyBackendServiceLogConfig, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("LogConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecuritySettings, actual.SecuritySettings, dcl.Info{ObjectFunction: compareBackendServiceSecuritySettingsNewStyle, EmptyObject: EmptyBackendServiceSecuritySettings, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecuritySettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalityLbPolicy, actual.LocalityLbPolicy, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("LocalityLbPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConsistentHash, actual.ConsistentHash, dcl.Info{ObjectFunction: compareBackendServiceConsistentHashNewStyle, EmptyObject: EmptyBackendServiceConsistentHash, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ConsistentHash")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CircuitBreakers, actual.CircuitBreakers, dcl.Info{ObjectFunction: compareBackendServiceCircuitBreakersNewStyle, EmptyObject: EmptyBackendServiceCircuitBreakers, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CircuitBreakers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OutlierDetection, actual.OutlierDetection, dcl.Info{ObjectFunction: compareBackendServiceOutlierDetectionNewStyle, EmptyObject: EmptyBackendServiceOutlierDetection, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("OutlierDetection")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxStreamDuration, actual.MaxStreamDuration, dcl.Info{ObjectFunction: compareBackendServiceMaxStreamDurationNewStyle, EmptyObject: EmptyBackendServiceMaxStreamDuration, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxStreamDuration")); len(ds) != 0 || err != nil {
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
func compareBackendServiceBackendsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceBackends)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceBackends)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceBackends or *BackendServiceBackends", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceBackends)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceBackends)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceBackends", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Group, actual.Group, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Group")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BalancingMode, actual.BalancingMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("BalancingMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxUtilization, actual.MaxUtilization, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxUtilization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRate, actual.MaxRate, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRatePerInstance, actual.MaxRatePerInstance, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxRatePerInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRatePerEndpoint, actual.MaxRatePerEndpoint, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxRatePerEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxConnections, actual.MaxConnections, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxConnections")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxConnectionsPerInstance, actual.MaxConnectionsPerInstance, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxConnectionsPerInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxConnectionsPerEndpoint, actual.MaxConnectionsPerEndpoint, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxConnectionsPerEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CapacityScaler, actual.CapacityScaler, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CapacityScaler")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Failover, actual.Failover, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Failover")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceFailoverPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceFailoverPolicy)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceFailoverPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceFailoverPolicy or *BackendServiceFailoverPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceFailoverPolicy)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceFailoverPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceFailoverPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisableConnectionDrainOnFailover, actual.DisableConnectionDrainOnFailover, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("DisableConnectionDrainOnFailover")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DropTrafficIfUnhealthy, actual.DropTrafficIfUnhealthy, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("DropTrafficIfUnhealthy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailoverRatio, actual.FailoverRatio, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("FailoverRatio")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceConnectionDrainingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceConnectionDraining)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceConnectionDraining)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConnectionDraining or *BackendServiceConnectionDraining", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceConnectionDraining)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceConnectionDraining)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConnectionDraining", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DrainingTimeoutSec, actual.DrainingTimeoutSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("DrainingTimeoutSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceIapNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceIap)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceIap)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceIap or *BackendServiceIap", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceIap)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceIap)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceIap", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientId, actual.OAuth2ClientId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Oauth2ClientId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientSecret, actual.OAuth2ClientSecret, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Oauth2ClientSecret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientSecretSha256, actual.OAuth2ClientSecretSha256, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oauth2ClientSecretSha256")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceCdnPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceCdnPolicy)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceCdnPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicy or *BackendServiceCdnPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceCdnPolicy)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceCdnPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CacheKeyPolicy, actual.CacheKeyPolicy, dcl.Info{ObjectFunction: compareBackendServiceCdnPolicyCacheKeyPolicyNewStyle, EmptyObject: EmptyBackendServiceCdnPolicyCacheKeyPolicy, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CacheKeyPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignedUrlKeyNames, actual.SignedUrlKeyNames, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SignedUrlKeyNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignedUrlCacheMaxAgeSec, actual.SignedUrlCacheMaxAgeSec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SignedUrlCacheMaxAgeSec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestCoalescing, actual.RequestCoalescing, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("RequestCoalescing")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CacheMode, actual.CacheMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("CacheMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultTtl, actual.DefaultTtl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("DefaultTtl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxTtl, actual.MaxTtl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxTtl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientTtl, actual.ClientTtl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ClientTtl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NegativeCaching, actual.NegativeCaching, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("NegativeCaching")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NegativeCachingPolicy, actual.NegativeCachingPolicy, dcl.Info{ObjectFunction: compareBackendServiceCdnPolicyNegativeCachingPolicyNewStyle, EmptyObject: EmptyBackendServiceCdnPolicyNegativeCachingPolicy, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("NegativeCachingPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BypassCacheOnRequestHeaders, actual.BypassCacheOnRequestHeaders, dcl.Info{ObjectFunction: compareBackendServiceCdnPolicyBypassCacheOnRequestHeadersNewStyle, EmptyObject: EmptyBackendServiceCdnPolicyBypassCacheOnRequestHeaders, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("BypassCacheOnRequestHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServeWhileStale, actual.ServeWhileStale, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ServeWhileStale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceCdnPolicyCacheKeyPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceCdnPolicyCacheKeyPolicy)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceCdnPolicyCacheKeyPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicyCacheKeyPolicy or *BackendServiceCdnPolicyCacheKeyPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceCdnPolicyCacheKeyPolicy)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceCdnPolicyCacheKeyPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicyCacheKeyPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IncludeProtocol, actual.IncludeProtocol, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("IncludeProtocol")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeHost, actual.IncludeHost, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("IncludeHost")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeQueryString, actual.IncludeQueryString, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("IncludeQueryString")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QueryStringWhitelist, actual.QueryStringWhitelist, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("QueryStringWhitelist")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QueryStringBlacklist, actual.QueryStringBlacklist, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("QueryStringBlacklist")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeHttpHeaders, actual.IncludeHttpHeaders, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("IncludeHttpHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeNamedCookies, actual.IncludeNamedCookies, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("IncludeNamedCookies")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceCdnPolicyNegativeCachingPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceCdnPolicyNegativeCachingPolicy)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceCdnPolicyNegativeCachingPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicyNegativeCachingPolicy or *BackendServiceCdnPolicyNegativeCachingPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceCdnPolicyNegativeCachingPolicy)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceCdnPolicyNegativeCachingPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicyNegativeCachingPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ttl, actual.Ttl, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Ttl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceCdnPolicyBypassCacheOnRequestHeadersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceCdnPolicyBypassCacheOnRequestHeaders)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceCdnPolicyBypassCacheOnRequestHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicyBypassCacheOnRequestHeaders or *BackendServiceCdnPolicyBypassCacheOnRequestHeaders", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceCdnPolicyBypassCacheOnRequestHeaders)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceCdnPolicyBypassCacheOnRequestHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCdnPolicyBypassCacheOnRequestHeaders", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HeaderName, actual.HeaderName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("HeaderName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceLogConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceLogConfig)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceLogConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceLogConfig or *BackendServiceLogConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceLogConfig)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceLogConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceLogConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enable, actual.Enable, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Enable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SampleRate, actual.SampleRate, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SampleRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceSecuritySettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceSecuritySettings)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceSecuritySettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceSecuritySettings or *BackendServiceSecuritySettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceSecuritySettings)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceSecuritySettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceSecuritySettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ClientTlsPolicy, actual.ClientTlsPolicy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ClientTlsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubjectAltNames, actual.SubjectAltNames, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SubjectAltNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceConsistentHashNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceConsistentHash)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceConsistentHash)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConsistentHash or *BackendServiceConsistentHash", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceConsistentHash)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceConsistentHash)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConsistentHash", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpCookie, actual.HttpCookie, dcl.Info{ObjectFunction: compareBackendServiceConsistentHashHttpCookieNewStyle, EmptyObject: EmptyBackendServiceConsistentHashHttpCookie, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("HttpCookie")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpHeaderName, actual.HttpHeaderName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("HttpHeaderName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinimumRingSize, actual.MinimumRingSize, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MinimumRingSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceConsistentHashHttpCookieNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceConsistentHashHttpCookie)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceConsistentHashHttpCookie)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConsistentHashHttpCookie or *BackendServiceConsistentHashHttpCookie", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceConsistentHashHttpCookie)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceConsistentHashHttpCookie)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConsistentHashHttpCookie", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ttl, actual.Ttl, dcl.Info{ObjectFunction: compareBackendServiceConsistentHashHttpCookieTtlNewStyle, EmptyObject: EmptyBackendServiceConsistentHashHttpCookieTtl, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Ttl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceConsistentHashHttpCookieTtlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceConsistentHashHttpCookieTtl)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceConsistentHashHttpCookieTtl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConsistentHashHttpCookieTtl or *BackendServiceConsistentHashHttpCookieTtl", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceConsistentHashHttpCookieTtl)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceConsistentHashHttpCookieTtl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceConsistentHashHttpCookieTtl", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceCircuitBreakersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceCircuitBreakers)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceCircuitBreakers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCircuitBreakers or *BackendServiceCircuitBreakers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceCircuitBreakers)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceCircuitBreakers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceCircuitBreakers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxRequestsPerConnection, actual.MaxRequestsPerConnection, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxRequestsPerConnection")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxConnections, actual.MaxConnections, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxConnections")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxPendingRequests, actual.MaxPendingRequests, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxPendingRequests")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRequests, actual.MaxRequests, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxRequests")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRetries, actual.MaxRetries, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceOutlierDetectionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceOutlierDetection)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceOutlierDetection)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceOutlierDetection or *BackendServiceOutlierDetection", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceOutlierDetection)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceOutlierDetection)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceOutlierDetection", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConsecutiveErrors, actual.ConsecutiveErrors, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ConsecutiveErrors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interval, actual.Interval, dcl.Info{ObjectFunction: compareBackendServiceOutlierDetectionIntervalNewStyle, EmptyObject: EmptyBackendServiceOutlierDetectionInterval, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("Interval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BaseEjectionTime, actual.BaseEjectionTime, dcl.Info{ObjectFunction: compareBackendServiceOutlierDetectionBaseEjectionTimeNewStyle, EmptyObject: EmptyBackendServiceOutlierDetectionBaseEjectionTime, OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("BaseEjectionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxEjectionPercent, actual.MaxEjectionPercent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("MaxEjectionPercent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcingConsecutiveErrors, actual.EnforcingConsecutiveErrors, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("EnforcingConsecutiveErrors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcingSuccessRate, actual.EnforcingSuccessRate, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("EnforcingSuccessRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessRateMinimumHosts, actual.SuccessRateMinimumHosts, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SuccessRateMinimumHosts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessRateRequestVolume, actual.SuccessRateRequestVolume, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SuccessRateRequestVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuccessRateStdevFactor, actual.SuccessRateStdevFactor, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("SuccessRateStdevFactor")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConsecutiveGatewayFailure, actual.ConsecutiveGatewayFailure, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("ConsecutiveGatewayFailure")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcingConsecutiveGatewayFailure, actual.EnforcingConsecutiveGatewayFailure, dcl.Info{OperationSelector: dcl.TriggersOperation("updateBackendServiceUpdateOperation")}, fn.AddNest("EnforcingConsecutiveGatewayFailure")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceOutlierDetectionIntervalNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceOutlierDetectionInterval)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceOutlierDetectionInterval)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceOutlierDetectionInterval or *BackendServiceOutlierDetectionInterval", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceOutlierDetectionInterval)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceOutlierDetectionInterval)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceOutlierDetectionInterval", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceOutlierDetectionBaseEjectionTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceOutlierDetectionBaseEjectionTime)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceOutlierDetectionBaseEjectionTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceOutlierDetectionBaseEjectionTime or *BackendServiceOutlierDetectionBaseEjectionTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceOutlierDetectionBaseEjectionTime)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceOutlierDetectionBaseEjectionTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceOutlierDetectionBaseEjectionTime", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBackendServiceMaxStreamDurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BackendServiceMaxStreamDuration)
	if !ok {
		desiredNotPointer, ok := d.(BackendServiceMaxStreamDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceMaxStreamDuration or *BackendServiceMaxStreamDuration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BackendServiceMaxStreamDuration)
	if !ok {
		actualNotPointer, ok := a.(BackendServiceMaxStreamDuration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BackendServiceMaxStreamDuration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Seconds, actual.Seconds, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Seconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
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
func (r *BackendService) urlNormalized() *BackendService {
	normalized := dcl.Copy(*r).(BackendService)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.SelfLinkWithId = dcl.SelfLinkToName(r.SelfLinkWithId)
	normalized.Fingerprint = dcl.SelfLinkToName(r.Fingerprint)
	normalized.PortName = dcl.SelfLinkToName(r.PortName)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.SecurityPolicy = dcl.SelfLinkToName(r.SecurityPolicy)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *BackendService) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *BackendService) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *BackendService) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *BackendService) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "Update" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		if dcl.IsRegion(r.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/backendServices/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil
		}

		return dcl.URL("projects/{{project}}/global/backendServices/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the BackendService resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *BackendService) marshal(c *Client) ([]byte, error) {
	m, err := expandBackendService(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling BackendService: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBackendService decodes JSON responses into the BackendService resource schema.
func unmarshalBackendService(b []byte, c *Client) (*BackendService, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBackendService(m, c)
}

func unmarshalMapBackendService(m map[string]interface{}, c *Client) (*BackendService, error) {

	return flattenBackendService(c, m), nil
}

// expandBackendService expands BackendService into a JSON request object.
func expandBackendService(c *Client, f *BackendService) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.SelfLinkWithId; !dcl.IsEmptyValueIndirect(v) {
		m["selfLinkWithId"] = v
	}
	if v, err := expandBackendServiceBackendsSlice(c, f.Backends); err != nil {
		return nil, fmt.Errorf("error expanding Backends into backends: %w", err)
	} else if v != nil {
		m["backends"] = v
	}
	if v := f.HealthChecks; !dcl.IsEmptyValueIndirect(v) {
		m["healthChecks"] = v
	}
	if v := f.TimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["timeoutSec"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.Protocol; !dcl.IsEmptyValueIndirect(v) {
		m["protocol"] = v
	}
	if v := f.Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["fingerprint"] = v
	}
	if v := f.PortName; !dcl.IsEmptyValueIndirect(v) {
		m["portName"] = v
	}
	if v := f.EnableCdn; !dcl.IsEmptyValueIndirect(v) {
		m["enableCDN"] = v
	}
	if v := f.SessionAffinity; !dcl.IsEmptyValueIndirect(v) {
		m["sessionAffinity"] = v
	}
	if v := f.AffinityCookieTtlSec; !dcl.IsEmptyValueIndirect(v) {
		m["affinityCookieTtlSec"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into region: %w", err)
	} else if v != nil {
		m["region"] = v
	}
	if v, err := expandBackendServiceFailoverPolicy(c, f.FailoverPolicy); err != nil {
		return nil, fmt.Errorf("error expanding FailoverPolicy into failoverPolicy: %w", err)
	} else if v != nil {
		m["failoverPolicy"] = v
	}
	if v := f.LoadBalancingScheme; !dcl.IsEmptyValueIndirect(v) {
		m["loadBalancingScheme"] = v
	}
	if v, err := expandBackendServiceConnectionDraining(c, f.ConnectionDraining); err != nil {
		return nil, fmt.Errorf("error expanding ConnectionDraining into connectionDraining: %w", err)
	} else if v != nil {
		m["connectionDraining"] = v
	}
	if v, err := expandBackendServiceIap(c, f.Iap); err != nil {
		return nil, fmt.Errorf("error expanding Iap into iap: %w", err)
	} else if v != nil {
		m["iap"] = v
	}
	if v, err := expandBackendServiceCdnPolicy(c, f.CdnPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CdnPolicy into cdnPolicy: %w", err)
	} else if v != nil {
		m["cdnPolicy"] = v
	}
	if v := f.CustomRequestHeaders; !dcl.IsEmptyValueIndirect(v) {
		m["customRequestHeaders"] = v
	}
	if v := f.CustomResponseHeaders; !dcl.IsEmptyValueIndirect(v) {
		m["customResponseHeaders"] = v
	}
	if v := f.SecurityPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["securityPolicy"] = v
	}
	if v, err := expandBackendServiceLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if v != nil {
		m["logConfig"] = v
	}
	if v, err := expandBackendServiceSecuritySettings(c, f.SecuritySettings); err != nil {
		return nil, fmt.Errorf("error expanding SecuritySettings into securitySettings: %w", err)
	} else if v != nil {
		m["securitySettings"] = v
	}
	if v := f.LocalityLbPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["localityLbPolicy"] = v
	}
	if v, err := expandBackendServiceConsistentHash(c, f.ConsistentHash); err != nil {
		return nil, fmt.Errorf("error expanding ConsistentHash into consistentHash: %w", err)
	} else if v != nil {
		m["consistentHash"] = v
	}
	if v, err := expandBackendServiceCircuitBreakers(c, f.CircuitBreakers); err != nil {
		return nil, fmt.Errorf("error expanding CircuitBreakers into circuitBreakers: %w", err)
	} else if v != nil {
		m["circuitBreakers"] = v
	}
	if v, err := expandBackendServiceOutlierDetection(c, f.OutlierDetection); err != nil {
		return nil, fmt.Errorf("error expanding OutlierDetection into outlierDetection: %w", err)
	} else if v != nil {
		m["outlierDetection"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v, err := expandBackendServiceMaxStreamDuration(c, f.MaxStreamDuration); err != nil {
		return nil, fmt.Errorf("error expanding MaxStreamDuration into maxStreamDuration: %w", err)
	} else if v != nil {
		m["maxStreamDuration"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenBackendService flattens BackendService from a JSON request object into the
// BackendService type.
func flattenBackendService(c *Client, i interface{}) *BackendService {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &BackendService{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.SelfLinkWithId = dcl.FlattenString(m["selfLinkWithId"])
	res.Backends = flattenBackendServiceBackendsSlice(c, m["backends"])
	res.HealthChecks = dcl.FlattenStringSlice(m["healthChecks"])
	res.TimeoutSec = dcl.FlattenInteger(m["timeoutSec"])
	res.Port = dcl.FlattenInteger(m["port"])
	res.Protocol = flattenBackendServiceProtocolEnum(m["protocol"])
	res.Fingerprint = dcl.FlattenString(m["fingerprint"])
	res.PortName = dcl.FlattenString(m["portName"])
	res.EnableCdn = dcl.FlattenBool(m["enableCDN"])
	res.SessionAffinity = flattenBackendServiceSessionAffinityEnum(m["sessionAffinity"])
	res.AffinityCookieTtlSec = dcl.FlattenInteger(m["affinityCookieTtlSec"])
	res.Location = dcl.FlattenString(m["region"])
	res.FailoverPolicy = flattenBackendServiceFailoverPolicy(c, m["failoverPolicy"])
	res.LoadBalancingScheme = flattenBackendServiceLoadBalancingSchemeEnum(m["loadBalancingScheme"])
	res.ConnectionDraining = flattenBackendServiceConnectionDraining(c, m["connectionDraining"])
	res.Iap = flattenBackendServiceIap(c, m["iap"])
	res.CdnPolicy = flattenBackendServiceCdnPolicy(c, m["cdnPolicy"])
	res.CustomRequestHeaders = dcl.FlattenStringSlice(m["customRequestHeaders"])
	res.CustomResponseHeaders = dcl.FlattenStringSlice(m["customResponseHeaders"])
	res.SecurityPolicy = dcl.FlattenString(m["securityPolicy"])
	res.LogConfig = flattenBackendServiceLogConfig(c, m["logConfig"])
	res.SecuritySettings = flattenBackendServiceSecuritySettings(c, m["securitySettings"])
	res.LocalityLbPolicy = flattenBackendServiceLocalityLbPolicyEnum(m["localityLbPolicy"])
	res.ConsistentHash = flattenBackendServiceConsistentHash(c, m["consistentHash"])
	res.CircuitBreakers = flattenBackendServiceCircuitBreakers(c, m["circuitBreakers"])
	res.OutlierDetection = flattenBackendServiceOutlierDetection(c, m["outlierDetection"])
	res.Network = dcl.FlattenString(m["network"])
	res.MaxStreamDuration = flattenBackendServiceMaxStreamDuration(c, m["maxStreamDuration"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandBackendServiceBackendsMap expands the contents of BackendServiceBackends into a JSON
// request object.
func expandBackendServiceBackendsMap(c *Client, f map[string]BackendServiceBackends) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceBackends(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceBackendsSlice expands the contents of BackendServiceBackends into a JSON
// request object.
func expandBackendServiceBackendsSlice(c *Client, f []BackendServiceBackends) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceBackends(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceBackendsMap flattens the contents of BackendServiceBackends from a JSON
// response object.
func flattenBackendServiceBackendsMap(c *Client, i interface{}) map[string]BackendServiceBackends {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceBackends{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceBackends{}
	}

	items := make(map[string]BackendServiceBackends)
	for k, item := range a {
		items[k] = *flattenBackendServiceBackends(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceBackendsSlice flattens the contents of BackendServiceBackends from a JSON
// response object.
func flattenBackendServiceBackendsSlice(c *Client, i interface{}) []BackendServiceBackends {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceBackends{}
	}

	if len(a) == 0 {
		return []BackendServiceBackends{}
	}

	items := make([]BackendServiceBackends, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceBackends(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceBackends expands an instance of BackendServiceBackends into a JSON
// request object.
func expandBackendServiceBackends(c *Client, f *BackendServiceBackends) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Group; !dcl.IsEmptyValueIndirect(v) {
		m["group"] = v
	}
	if v := f.BalancingMode; !dcl.IsEmptyValueIndirect(v) {
		m["balancingMode"] = v
	}
	if v := f.MaxUtilization; !dcl.IsEmptyValueIndirect(v) {
		m["maxUtilization"] = v
	}
	if v := f.MaxRate; !dcl.IsEmptyValueIndirect(v) {
		m["maxRate"] = v
	}
	if v := f.MaxRatePerInstance; !dcl.IsEmptyValueIndirect(v) {
		m["maxRatePerInstance"] = v
	}
	if v := f.MaxRatePerEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["maxRatePerEndpoint"] = v
	}
	if v := f.MaxConnections; !dcl.IsEmptyValueIndirect(v) {
		m["maxConnections"] = v
	}
	if v := f.MaxConnectionsPerInstance; !dcl.IsEmptyValueIndirect(v) {
		m["maxConnectionsPerInstance"] = v
	}
	if v := f.MaxConnectionsPerEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["maxConnectionsPerEndpoint"] = v
	}
	if v := f.CapacityScaler; !dcl.IsEmptyValueIndirect(v) {
		m["capacityScaler"] = v
	}
	if v := f.Failover; !dcl.IsEmptyValueIndirect(v) {
		m["failover"] = v
	}

	return m, nil
}

// flattenBackendServiceBackends flattens an instance of BackendServiceBackends from a JSON
// response object.
func flattenBackendServiceBackends(c *Client, i interface{}) *BackendServiceBackends {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceBackends{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceBackends
	}
	r.Description = dcl.FlattenString(m["description"])
	r.Group = dcl.FlattenString(m["group"])
	r.BalancingMode = flattenBackendServiceBackendsBalancingModeEnum(m["balancingMode"])
	r.MaxUtilization = dcl.FlattenDouble(m["maxUtilization"])
	r.MaxRate = dcl.FlattenInteger(m["maxRate"])
	r.MaxRatePerInstance = dcl.FlattenDouble(m["maxRatePerInstance"])
	r.MaxRatePerEndpoint = dcl.FlattenDouble(m["maxRatePerEndpoint"])
	r.MaxConnections = dcl.FlattenInteger(m["maxConnections"])
	r.MaxConnectionsPerInstance = dcl.FlattenInteger(m["maxConnectionsPerInstance"])
	r.MaxConnectionsPerEndpoint = dcl.FlattenInteger(m["maxConnectionsPerEndpoint"])
	r.CapacityScaler = dcl.FlattenDouble(m["capacityScaler"])
	r.Failover = dcl.FlattenBool(m["failover"])

	return r
}

// expandBackendServiceFailoverPolicyMap expands the contents of BackendServiceFailoverPolicy into a JSON
// request object.
func expandBackendServiceFailoverPolicyMap(c *Client, f map[string]BackendServiceFailoverPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceFailoverPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceFailoverPolicySlice expands the contents of BackendServiceFailoverPolicy into a JSON
// request object.
func expandBackendServiceFailoverPolicySlice(c *Client, f []BackendServiceFailoverPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceFailoverPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceFailoverPolicyMap flattens the contents of BackendServiceFailoverPolicy from a JSON
// response object.
func flattenBackendServiceFailoverPolicyMap(c *Client, i interface{}) map[string]BackendServiceFailoverPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceFailoverPolicy{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceFailoverPolicy{}
	}

	items := make(map[string]BackendServiceFailoverPolicy)
	for k, item := range a {
		items[k] = *flattenBackendServiceFailoverPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceFailoverPolicySlice flattens the contents of BackendServiceFailoverPolicy from a JSON
// response object.
func flattenBackendServiceFailoverPolicySlice(c *Client, i interface{}) []BackendServiceFailoverPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceFailoverPolicy{}
	}

	if len(a) == 0 {
		return []BackendServiceFailoverPolicy{}
	}

	items := make([]BackendServiceFailoverPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceFailoverPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceFailoverPolicy expands an instance of BackendServiceFailoverPolicy into a JSON
// request object.
func expandBackendServiceFailoverPolicy(c *Client, f *BackendServiceFailoverPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DisableConnectionDrainOnFailover; !dcl.IsEmptyValueIndirect(v) {
		m["disableConnectionDrainOnFailover"] = v
	}
	if v := f.DropTrafficIfUnhealthy; !dcl.IsEmptyValueIndirect(v) {
		m["dropTrafficIfUnhealthy"] = v
	}
	if v := f.FailoverRatio; !dcl.IsEmptyValueIndirect(v) {
		m["failoverRatio"] = v
	}

	return m, nil
}

// flattenBackendServiceFailoverPolicy flattens an instance of BackendServiceFailoverPolicy from a JSON
// response object.
func flattenBackendServiceFailoverPolicy(c *Client, i interface{}) *BackendServiceFailoverPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceFailoverPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceFailoverPolicy
	}
	r.DisableConnectionDrainOnFailover = dcl.FlattenBool(m["disableConnectionDrainOnFailover"])
	r.DropTrafficIfUnhealthy = dcl.FlattenBool(m["dropTrafficIfUnhealthy"])
	r.FailoverRatio = dcl.FlattenDouble(m["failoverRatio"])

	return r
}

// expandBackendServiceConnectionDrainingMap expands the contents of BackendServiceConnectionDraining into a JSON
// request object.
func expandBackendServiceConnectionDrainingMap(c *Client, f map[string]BackendServiceConnectionDraining) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceConnectionDraining(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceConnectionDrainingSlice expands the contents of BackendServiceConnectionDraining into a JSON
// request object.
func expandBackendServiceConnectionDrainingSlice(c *Client, f []BackendServiceConnectionDraining) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceConnectionDraining(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceConnectionDrainingMap flattens the contents of BackendServiceConnectionDraining from a JSON
// response object.
func flattenBackendServiceConnectionDrainingMap(c *Client, i interface{}) map[string]BackendServiceConnectionDraining {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceConnectionDraining{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceConnectionDraining{}
	}

	items := make(map[string]BackendServiceConnectionDraining)
	for k, item := range a {
		items[k] = *flattenBackendServiceConnectionDraining(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceConnectionDrainingSlice flattens the contents of BackendServiceConnectionDraining from a JSON
// response object.
func flattenBackendServiceConnectionDrainingSlice(c *Client, i interface{}) []BackendServiceConnectionDraining {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceConnectionDraining{}
	}

	if len(a) == 0 {
		return []BackendServiceConnectionDraining{}
	}

	items := make([]BackendServiceConnectionDraining, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceConnectionDraining(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceConnectionDraining expands an instance of BackendServiceConnectionDraining into a JSON
// request object.
func expandBackendServiceConnectionDraining(c *Client, f *BackendServiceConnectionDraining) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DrainingTimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["drainingTimeoutSec"] = v
	}

	return m, nil
}

// flattenBackendServiceConnectionDraining flattens an instance of BackendServiceConnectionDraining from a JSON
// response object.
func flattenBackendServiceConnectionDraining(c *Client, i interface{}) *BackendServiceConnectionDraining {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceConnectionDraining{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceConnectionDraining
	}
	r.DrainingTimeoutSec = dcl.FlattenInteger(m["drainingTimeoutSec"])

	return r
}

// expandBackendServiceIapMap expands the contents of BackendServiceIap into a JSON
// request object.
func expandBackendServiceIapMap(c *Client, f map[string]BackendServiceIap) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceIap(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceIapSlice expands the contents of BackendServiceIap into a JSON
// request object.
func expandBackendServiceIapSlice(c *Client, f []BackendServiceIap) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceIap(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceIapMap flattens the contents of BackendServiceIap from a JSON
// response object.
func flattenBackendServiceIapMap(c *Client, i interface{}) map[string]BackendServiceIap {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceIap{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceIap{}
	}

	items := make(map[string]BackendServiceIap)
	for k, item := range a {
		items[k] = *flattenBackendServiceIap(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceIapSlice flattens the contents of BackendServiceIap from a JSON
// response object.
func flattenBackendServiceIapSlice(c *Client, i interface{}) []BackendServiceIap {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceIap{}
	}

	if len(a) == 0 {
		return []BackendServiceIap{}
	}

	items := make([]BackendServiceIap, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceIap(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceIap expands an instance of BackendServiceIap into a JSON
// request object.
func expandBackendServiceIap(c *Client, f *BackendServiceIap) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.OAuth2ClientId; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientId"] = v
	}
	if v := f.OAuth2ClientSecret; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientSecret"] = v
	}
	if v := f.OAuth2ClientSecretSha256; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientSecretSha256"] = v
	}

	return m, nil
}

// flattenBackendServiceIap flattens an instance of BackendServiceIap from a JSON
// response object.
func flattenBackendServiceIap(c *Client, i interface{}) *BackendServiceIap {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceIap{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceIap
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.OAuth2ClientId = dcl.FlattenString(m["oauth2ClientId"])
	r.OAuth2ClientSecret = dcl.FlattenString(m["oauth2ClientSecret"])
	r.OAuth2ClientSecretSha256 = dcl.FlattenString(m["oauth2ClientSecretSha256"])

	return r
}

// expandBackendServiceCdnPolicyMap expands the contents of BackendServiceCdnPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyMap(c *Client, f map[string]BackendServiceCdnPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceCdnPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceCdnPolicySlice expands the contents of BackendServiceCdnPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicySlice(c *Client, f []BackendServiceCdnPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceCdnPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceCdnPolicyMap flattens the contents of BackendServiceCdnPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyMap(c *Client, i interface{}) map[string]BackendServiceCdnPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceCdnPolicy{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceCdnPolicy{}
	}

	items := make(map[string]BackendServiceCdnPolicy)
	for k, item := range a {
		items[k] = *flattenBackendServiceCdnPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceCdnPolicySlice flattens the contents of BackendServiceCdnPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicySlice(c *Client, i interface{}) []BackendServiceCdnPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceCdnPolicy{}
	}

	if len(a) == 0 {
		return []BackendServiceCdnPolicy{}
	}

	items := make([]BackendServiceCdnPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceCdnPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceCdnPolicy expands an instance of BackendServiceCdnPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicy(c *Client, f *BackendServiceCdnPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBackendServiceCdnPolicyCacheKeyPolicy(c, f.CacheKeyPolicy); err != nil {
		return nil, fmt.Errorf("error expanding CacheKeyPolicy into cacheKeyPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cacheKeyPolicy"] = v
	}
	if v := f.SignedUrlKeyNames; !dcl.IsEmptyValueIndirect(v) {
		m["signedUrlKeyNames"] = v
	}
	if v := f.SignedUrlCacheMaxAgeSec; !dcl.IsEmptyValueIndirect(v) {
		m["signedUrlCacheMaxAgeSec"] = v
	}
	if v := f.RequestCoalescing; !dcl.IsEmptyValueIndirect(v) {
		m["requestCoalescing"] = v
	}
	if v := f.CacheMode; !dcl.IsEmptyValueIndirect(v) {
		m["cacheMode"] = v
	}
	if v := f.DefaultTtl; !dcl.IsEmptyValueIndirect(v) {
		m["defaultTtl"] = v
	}
	if v := f.MaxTtl; !dcl.IsEmptyValueIndirect(v) {
		m["maxTtl"] = v
	}
	if v := f.ClientTtl; !dcl.IsEmptyValueIndirect(v) {
		m["clientTtl"] = v
	}
	if v := f.NegativeCaching; !dcl.IsEmptyValueIndirect(v) {
		m["negativeCaching"] = v
	}
	if v, err := expandBackendServiceCdnPolicyNegativeCachingPolicySlice(c, f.NegativeCachingPolicy); err != nil {
		return nil, fmt.Errorf("error expanding NegativeCachingPolicy into negativeCachingPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["negativeCachingPolicy"] = v
	}
	if v, err := expandBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice(c, f.BypassCacheOnRequestHeaders); err != nil {
		return nil, fmt.Errorf("error expanding BypassCacheOnRequestHeaders into bypassCacheOnRequestHeaders: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bypassCacheOnRequestHeaders"] = v
	}
	if v := f.ServeWhileStale; !dcl.IsEmptyValueIndirect(v) {
		m["serveWhileStale"] = v
	}

	return m, nil
}

// flattenBackendServiceCdnPolicy flattens an instance of BackendServiceCdnPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicy(c *Client, i interface{}) *BackendServiceCdnPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceCdnPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceCdnPolicy
	}
	r.CacheKeyPolicy = flattenBackendServiceCdnPolicyCacheKeyPolicy(c, m["cacheKeyPolicy"])
	r.SignedUrlKeyNames = dcl.FlattenStringSlice(m["signedUrlKeyNames"])
	r.SignedUrlCacheMaxAgeSec = dcl.FlattenInteger(m["signedUrlCacheMaxAgeSec"])
	r.RequestCoalescing = dcl.FlattenBool(m["requestCoalescing"])
	r.CacheMode = flattenBackendServiceCdnPolicyCacheModeEnum(m["cacheMode"])
	r.DefaultTtl = dcl.FlattenInteger(m["defaultTtl"])
	r.MaxTtl = dcl.FlattenInteger(m["maxTtl"])
	r.ClientTtl = dcl.FlattenInteger(m["clientTtl"])
	r.NegativeCaching = dcl.FlattenBool(m["negativeCaching"])
	r.NegativeCachingPolicy = flattenBackendServiceCdnPolicyNegativeCachingPolicySlice(c, m["negativeCachingPolicy"])
	r.BypassCacheOnRequestHeaders = flattenBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice(c, m["bypassCacheOnRequestHeaders"])
	r.ServeWhileStale = dcl.FlattenInteger(m["serveWhileStale"])

	return r
}

// expandBackendServiceCdnPolicyCacheKeyPolicyMap expands the contents of BackendServiceCdnPolicyCacheKeyPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyCacheKeyPolicyMap(c *Client, f map[string]BackendServiceCdnPolicyCacheKeyPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceCdnPolicyCacheKeyPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceCdnPolicyCacheKeyPolicySlice expands the contents of BackendServiceCdnPolicyCacheKeyPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyCacheKeyPolicySlice(c *Client, f []BackendServiceCdnPolicyCacheKeyPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceCdnPolicyCacheKeyPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceCdnPolicyCacheKeyPolicyMap flattens the contents of BackendServiceCdnPolicyCacheKeyPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyCacheKeyPolicyMap(c *Client, i interface{}) map[string]BackendServiceCdnPolicyCacheKeyPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceCdnPolicyCacheKeyPolicy{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceCdnPolicyCacheKeyPolicy{}
	}

	items := make(map[string]BackendServiceCdnPolicyCacheKeyPolicy)
	for k, item := range a {
		items[k] = *flattenBackendServiceCdnPolicyCacheKeyPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceCdnPolicyCacheKeyPolicySlice flattens the contents of BackendServiceCdnPolicyCacheKeyPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyCacheKeyPolicySlice(c *Client, i interface{}) []BackendServiceCdnPolicyCacheKeyPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceCdnPolicyCacheKeyPolicy{}
	}

	if len(a) == 0 {
		return []BackendServiceCdnPolicyCacheKeyPolicy{}
	}

	items := make([]BackendServiceCdnPolicyCacheKeyPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceCdnPolicyCacheKeyPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceCdnPolicyCacheKeyPolicy expands an instance of BackendServiceCdnPolicyCacheKeyPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyCacheKeyPolicy(c *Client, f *BackendServiceCdnPolicyCacheKeyPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IncludeProtocol; !dcl.IsEmptyValueIndirect(v) {
		m["includeProtocol"] = v
	}
	if v := f.IncludeHost; !dcl.IsEmptyValueIndirect(v) {
		m["includeHost"] = v
	}
	if v := f.IncludeQueryString; !dcl.IsEmptyValueIndirect(v) {
		m["includeQueryString"] = v
	}
	if v := f.QueryStringWhitelist; !dcl.IsEmptyValueIndirect(v) {
		m["queryStringWhitelist"] = v
	}
	if v := f.QueryStringBlacklist; !dcl.IsEmptyValueIndirect(v) {
		m["queryStringBlacklist"] = v
	}
	if v := f.IncludeHttpHeaders; !dcl.IsEmptyValueIndirect(v) {
		m["includeHttpHeaders"] = v
	}
	if v := f.IncludeNamedCookies; !dcl.IsEmptyValueIndirect(v) {
		m["includeNamedCookies"] = v
	}

	return m, nil
}

// flattenBackendServiceCdnPolicyCacheKeyPolicy flattens an instance of BackendServiceCdnPolicyCacheKeyPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyCacheKeyPolicy(c *Client, i interface{}) *BackendServiceCdnPolicyCacheKeyPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceCdnPolicyCacheKeyPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceCdnPolicyCacheKeyPolicy
	}
	r.IncludeProtocol = dcl.FlattenBool(m["includeProtocol"])
	r.IncludeHost = dcl.FlattenBool(m["includeHost"])
	r.IncludeQueryString = dcl.FlattenBool(m["includeQueryString"])
	r.QueryStringWhitelist = dcl.FlattenStringSlice(m["queryStringWhitelist"])
	r.QueryStringBlacklist = dcl.FlattenStringSlice(m["queryStringBlacklist"])
	r.IncludeHttpHeaders = dcl.FlattenStringSlice(m["includeHttpHeaders"])
	r.IncludeNamedCookies = dcl.FlattenStringSlice(m["includeNamedCookies"])

	return r
}

// expandBackendServiceCdnPolicyNegativeCachingPolicyMap expands the contents of BackendServiceCdnPolicyNegativeCachingPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyNegativeCachingPolicyMap(c *Client, f map[string]BackendServiceCdnPolicyNegativeCachingPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceCdnPolicyNegativeCachingPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceCdnPolicyNegativeCachingPolicySlice expands the contents of BackendServiceCdnPolicyNegativeCachingPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyNegativeCachingPolicySlice(c *Client, f []BackendServiceCdnPolicyNegativeCachingPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceCdnPolicyNegativeCachingPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceCdnPolicyNegativeCachingPolicyMap flattens the contents of BackendServiceCdnPolicyNegativeCachingPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyNegativeCachingPolicyMap(c *Client, i interface{}) map[string]BackendServiceCdnPolicyNegativeCachingPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceCdnPolicyNegativeCachingPolicy{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceCdnPolicyNegativeCachingPolicy{}
	}

	items := make(map[string]BackendServiceCdnPolicyNegativeCachingPolicy)
	for k, item := range a {
		items[k] = *flattenBackendServiceCdnPolicyNegativeCachingPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceCdnPolicyNegativeCachingPolicySlice flattens the contents of BackendServiceCdnPolicyNegativeCachingPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyNegativeCachingPolicySlice(c *Client, i interface{}) []BackendServiceCdnPolicyNegativeCachingPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceCdnPolicyNegativeCachingPolicy{}
	}

	if len(a) == 0 {
		return []BackendServiceCdnPolicyNegativeCachingPolicy{}
	}

	items := make([]BackendServiceCdnPolicyNegativeCachingPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceCdnPolicyNegativeCachingPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceCdnPolicyNegativeCachingPolicy expands an instance of BackendServiceCdnPolicyNegativeCachingPolicy into a JSON
// request object.
func expandBackendServiceCdnPolicyNegativeCachingPolicy(c *Client, f *BackendServiceCdnPolicyNegativeCachingPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Ttl; !dcl.IsEmptyValueIndirect(v) {
		m["ttl"] = v
	}

	return m, nil
}

// flattenBackendServiceCdnPolicyNegativeCachingPolicy flattens an instance of BackendServiceCdnPolicyNegativeCachingPolicy from a JSON
// response object.
func flattenBackendServiceCdnPolicyNegativeCachingPolicy(c *Client, i interface{}) *BackendServiceCdnPolicyNegativeCachingPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceCdnPolicyNegativeCachingPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceCdnPolicyNegativeCachingPolicy
	}
	r.Code = dcl.FlattenInteger(m["code"])
	r.Ttl = dcl.FlattenInteger(m["ttl"])

	return r
}

// expandBackendServiceCdnPolicyBypassCacheOnRequestHeadersMap expands the contents of BackendServiceCdnPolicyBypassCacheOnRequestHeaders into a JSON
// request object.
func expandBackendServiceCdnPolicyBypassCacheOnRequestHeadersMap(c *Client, f map[string]BackendServiceCdnPolicyBypassCacheOnRequestHeaders) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice expands the contents of BackendServiceCdnPolicyBypassCacheOnRequestHeaders into a JSON
// request object.
func expandBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice(c *Client, f []BackendServiceCdnPolicyBypassCacheOnRequestHeaders) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceCdnPolicyBypassCacheOnRequestHeadersMap flattens the contents of BackendServiceCdnPolicyBypassCacheOnRequestHeaders from a JSON
// response object.
func flattenBackendServiceCdnPolicyBypassCacheOnRequestHeadersMap(c *Client, i interface{}) map[string]BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceCdnPolicyBypassCacheOnRequestHeaders{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceCdnPolicyBypassCacheOnRequestHeaders{}
	}

	items := make(map[string]BackendServiceCdnPolicyBypassCacheOnRequestHeaders)
	for k, item := range a {
		items[k] = *flattenBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice flattens the contents of BackendServiceCdnPolicyBypassCacheOnRequestHeaders from a JSON
// response object.
func flattenBackendServiceCdnPolicyBypassCacheOnRequestHeadersSlice(c *Client, i interface{}) []BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceCdnPolicyBypassCacheOnRequestHeaders{}
	}

	if len(a) == 0 {
		return []BackendServiceCdnPolicyBypassCacheOnRequestHeaders{}
	}

	items := make([]BackendServiceCdnPolicyBypassCacheOnRequestHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceCdnPolicyBypassCacheOnRequestHeaders expands an instance of BackendServiceCdnPolicyBypassCacheOnRequestHeaders into a JSON
// request object.
func expandBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c *Client, f *BackendServiceCdnPolicyBypassCacheOnRequestHeaders) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["headerName"] = v
	}

	return m, nil
}

// flattenBackendServiceCdnPolicyBypassCacheOnRequestHeaders flattens an instance of BackendServiceCdnPolicyBypassCacheOnRequestHeaders from a JSON
// response object.
func flattenBackendServiceCdnPolicyBypassCacheOnRequestHeaders(c *Client, i interface{}) *BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceCdnPolicyBypassCacheOnRequestHeaders{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceCdnPolicyBypassCacheOnRequestHeaders
	}
	r.HeaderName = dcl.FlattenString(m["headerName"])

	return r
}

// expandBackendServiceLogConfigMap expands the contents of BackendServiceLogConfig into a JSON
// request object.
func expandBackendServiceLogConfigMap(c *Client, f map[string]BackendServiceLogConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceLogConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceLogConfigSlice expands the contents of BackendServiceLogConfig into a JSON
// request object.
func expandBackendServiceLogConfigSlice(c *Client, f []BackendServiceLogConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceLogConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceLogConfigMap flattens the contents of BackendServiceLogConfig from a JSON
// response object.
func flattenBackendServiceLogConfigMap(c *Client, i interface{}) map[string]BackendServiceLogConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceLogConfig{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceLogConfig{}
	}

	items := make(map[string]BackendServiceLogConfig)
	for k, item := range a {
		items[k] = *flattenBackendServiceLogConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceLogConfigSlice flattens the contents of BackendServiceLogConfig from a JSON
// response object.
func flattenBackendServiceLogConfigSlice(c *Client, i interface{}) []BackendServiceLogConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceLogConfig{}
	}

	if len(a) == 0 {
		return []BackendServiceLogConfig{}
	}

	items := make([]BackendServiceLogConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceLogConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceLogConfig expands an instance of BackendServiceLogConfig into a JSON
// request object.
func expandBackendServiceLogConfig(c *Client, f *BackendServiceLogConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enable; !dcl.IsEmptyValueIndirect(v) {
		m["enable"] = v
	}
	if v := f.SampleRate; !dcl.IsEmptyValueIndirect(v) {
		m["sampleRate"] = v
	}

	return m, nil
}

// flattenBackendServiceLogConfig flattens an instance of BackendServiceLogConfig from a JSON
// response object.
func flattenBackendServiceLogConfig(c *Client, i interface{}) *BackendServiceLogConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceLogConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceLogConfig
	}
	r.Enable = dcl.FlattenBool(m["enable"])
	r.SampleRate = dcl.FlattenDouble(m["sampleRate"])

	return r
}

// expandBackendServiceSecuritySettingsMap expands the contents of BackendServiceSecuritySettings into a JSON
// request object.
func expandBackendServiceSecuritySettingsMap(c *Client, f map[string]BackendServiceSecuritySettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceSecuritySettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceSecuritySettingsSlice expands the contents of BackendServiceSecuritySettings into a JSON
// request object.
func expandBackendServiceSecuritySettingsSlice(c *Client, f []BackendServiceSecuritySettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceSecuritySettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceSecuritySettingsMap flattens the contents of BackendServiceSecuritySettings from a JSON
// response object.
func flattenBackendServiceSecuritySettingsMap(c *Client, i interface{}) map[string]BackendServiceSecuritySettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceSecuritySettings{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceSecuritySettings{}
	}

	items := make(map[string]BackendServiceSecuritySettings)
	for k, item := range a {
		items[k] = *flattenBackendServiceSecuritySettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceSecuritySettingsSlice flattens the contents of BackendServiceSecuritySettings from a JSON
// response object.
func flattenBackendServiceSecuritySettingsSlice(c *Client, i interface{}) []BackendServiceSecuritySettings {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceSecuritySettings{}
	}

	if len(a) == 0 {
		return []BackendServiceSecuritySettings{}
	}

	items := make([]BackendServiceSecuritySettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceSecuritySettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceSecuritySettings expands an instance of BackendServiceSecuritySettings into a JSON
// request object.
func expandBackendServiceSecuritySettings(c *Client, f *BackendServiceSecuritySettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ClientTlsPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["clientTlsPolicy"] = v
	}
	if v := f.SubjectAltNames; !dcl.IsEmptyValueIndirect(v) {
		m["subjectAltNames"] = v
	}

	return m, nil
}

// flattenBackendServiceSecuritySettings flattens an instance of BackendServiceSecuritySettings from a JSON
// response object.
func flattenBackendServiceSecuritySettings(c *Client, i interface{}) *BackendServiceSecuritySettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceSecuritySettings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceSecuritySettings
	}
	r.ClientTlsPolicy = dcl.FlattenString(m["clientTlsPolicy"])
	r.SubjectAltNames = dcl.FlattenStringSlice(m["subjectAltNames"])

	return r
}

// expandBackendServiceConsistentHashMap expands the contents of BackendServiceConsistentHash into a JSON
// request object.
func expandBackendServiceConsistentHashMap(c *Client, f map[string]BackendServiceConsistentHash) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceConsistentHash(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceConsistentHashSlice expands the contents of BackendServiceConsistentHash into a JSON
// request object.
func expandBackendServiceConsistentHashSlice(c *Client, f []BackendServiceConsistentHash) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceConsistentHash(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceConsistentHashMap flattens the contents of BackendServiceConsistentHash from a JSON
// response object.
func flattenBackendServiceConsistentHashMap(c *Client, i interface{}) map[string]BackendServiceConsistentHash {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceConsistentHash{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceConsistentHash{}
	}

	items := make(map[string]BackendServiceConsistentHash)
	for k, item := range a {
		items[k] = *flattenBackendServiceConsistentHash(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceConsistentHashSlice flattens the contents of BackendServiceConsistentHash from a JSON
// response object.
func flattenBackendServiceConsistentHashSlice(c *Client, i interface{}) []BackendServiceConsistentHash {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceConsistentHash{}
	}

	if len(a) == 0 {
		return []BackendServiceConsistentHash{}
	}

	items := make([]BackendServiceConsistentHash, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceConsistentHash(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceConsistentHash expands an instance of BackendServiceConsistentHash into a JSON
// request object.
func expandBackendServiceConsistentHash(c *Client, f *BackendServiceConsistentHash) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBackendServiceConsistentHashHttpCookie(c, f.HttpCookie); err != nil {
		return nil, fmt.Errorf("error expanding HttpCookie into httpCookie: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpCookie"] = v
	}
	if v := f.HttpHeaderName; !dcl.IsEmptyValueIndirect(v) {
		m["httpHeaderName"] = v
	}
	if v := f.MinimumRingSize; !dcl.IsEmptyValueIndirect(v) {
		m["minimumRingSize"] = v
	}

	return m, nil
}

// flattenBackendServiceConsistentHash flattens an instance of BackendServiceConsistentHash from a JSON
// response object.
func flattenBackendServiceConsistentHash(c *Client, i interface{}) *BackendServiceConsistentHash {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceConsistentHash{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceConsistentHash
	}
	r.HttpCookie = flattenBackendServiceConsistentHashHttpCookie(c, m["httpCookie"])
	r.HttpHeaderName = dcl.FlattenString(m["httpHeaderName"])
	r.MinimumRingSize = dcl.FlattenInteger(m["minimumRingSize"])

	return r
}

// expandBackendServiceConsistentHashHttpCookieMap expands the contents of BackendServiceConsistentHashHttpCookie into a JSON
// request object.
func expandBackendServiceConsistentHashHttpCookieMap(c *Client, f map[string]BackendServiceConsistentHashHttpCookie) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceConsistentHashHttpCookie(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceConsistentHashHttpCookieSlice expands the contents of BackendServiceConsistentHashHttpCookie into a JSON
// request object.
func expandBackendServiceConsistentHashHttpCookieSlice(c *Client, f []BackendServiceConsistentHashHttpCookie) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceConsistentHashHttpCookie(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceConsistentHashHttpCookieMap flattens the contents of BackendServiceConsistentHashHttpCookie from a JSON
// response object.
func flattenBackendServiceConsistentHashHttpCookieMap(c *Client, i interface{}) map[string]BackendServiceConsistentHashHttpCookie {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceConsistentHashHttpCookie{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceConsistentHashHttpCookie{}
	}

	items := make(map[string]BackendServiceConsistentHashHttpCookie)
	for k, item := range a {
		items[k] = *flattenBackendServiceConsistentHashHttpCookie(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceConsistentHashHttpCookieSlice flattens the contents of BackendServiceConsistentHashHttpCookie from a JSON
// response object.
func flattenBackendServiceConsistentHashHttpCookieSlice(c *Client, i interface{}) []BackendServiceConsistentHashHttpCookie {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceConsistentHashHttpCookie{}
	}

	if len(a) == 0 {
		return []BackendServiceConsistentHashHttpCookie{}
	}

	items := make([]BackendServiceConsistentHashHttpCookie, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceConsistentHashHttpCookie(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceConsistentHashHttpCookie expands an instance of BackendServiceConsistentHashHttpCookie into a JSON
// request object.
func expandBackendServiceConsistentHashHttpCookie(c *Client, f *BackendServiceConsistentHashHttpCookie) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v, err := expandBackendServiceConsistentHashHttpCookieTtl(c, f.Ttl); err != nil {
		return nil, fmt.Errorf("error expanding Ttl into ttl: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ttl"] = v
	}

	return m, nil
}

// flattenBackendServiceConsistentHashHttpCookie flattens an instance of BackendServiceConsistentHashHttpCookie from a JSON
// response object.
func flattenBackendServiceConsistentHashHttpCookie(c *Client, i interface{}) *BackendServiceConsistentHashHttpCookie {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceConsistentHashHttpCookie{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceConsistentHashHttpCookie
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Path = dcl.FlattenString(m["path"])
	r.Ttl = flattenBackendServiceConsistentHashHttpCookieTtl(c, m["ttl"])

	return r
}

// expandBackendServiceConsistentHashHttpCookieTtlMap expands the contents of BackendServiceConsistentHashHttpCookieTtl into a JSON
// request object.
func expandBackendServiceConsistentHashHttpCookieTtlMap(c *Client, f map[string]BackendServiceConsistentHashHttpCookieTtl) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceConsistentHashHttpCookieTtl(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceConsistentHashHttpCookieTtlSlice expands the contents of BackendServiceConsistentHashHttpCookieTtl into a JSON
// request object.
func expandBackendServiceConsistentHashHttpCookieTtlSlice(c *Client, f []BackendServiceConsistentHashHttpCookieTtl) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceConsistentHashHttpCookieTtl(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceConsistentHashHttpCookieTtlMap flattens the contents of BackendServiceConsistentHashHttpCookieTtl from a JSON
// response object.
func flattenBackendServiceConsistentHashHttpCookieTtlMap(c *Client, i interface{}) map[string]BackendServiceConsistentHashHttpCookieTtl {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceConsistentHashHttpCookieTtl{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceConsistentHashHttpCookieTtl{}
	}

	items := make(map[string]BackendServiceConsistentHashHttpCookieTtl)
	for k, item := range a {
		items[k] = *flattenBackendServiceConsistentHashHttpCookieTtl(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceConsistentHashHttpCookieTtlSlice flattens the contents of BackendServiceConsistentHashHttpCookieTtl from a JSON
// response object.
func flattenBackendServiceConsistentHashHttpCookieTtlSlice(c *Client, i interface{}) []BackendServiceConsistentHashHttpCookieTtl {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceConsistentHashHttpCookieTtl{}
	}

	if len(a) == 0 {
		return []BackendServiceConsistentHashHttpCookieTtl{}
	}

	items := make([]BackendServiceConsistentHashHttpCookieTtl, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceConsistentHashHttpCookieTtl(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceConsistentHashHttpCookieTtl expands an instance of BackendServiceConsistentHashHttpCookieTtl into a JSON
// request object.
func expandBackendServiceConsistentHashHttpCookieTtl(c *Client, f *BackendServiceConsistentHashHttpCookieTtl) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenBackendServiceConsistentHashHttpCookieTtl flattens an instance of BackendServiceConsistentHashHttpCookieTtl from a JSON
// response object.
func flattenBackendServiceConsistentHashHttpCookieTtl(c *Client, i interface{}) *BackendServiceConsistentHashHttpCookieTtl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceConsistentHashHttpCookieTtl{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceConsistentHashHttpCookieTtl
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandBackendServiceCircuitBreakersMap expands the contents of BackendServiceCircuitBreakers into a JSON
// request object.
func expandBackendServiceCircuitBreakersMap(c *Client, f map[string]BackendServiceCircuitBreakers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceCircuitBreakers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceCircuitBreakersSlice expands the contents of BackendServiceCircuitBreakers into a JSON
// request object.
func expandBackendServiceCircuitBreakersSlice(c *Client, f []BackendServiceCircuitBreakers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceCircuitBreakers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceCircuitBreakersMap flattens the contents of BackendServiceCircuitBreakers from a JSON
// response object.
func flattenBackendServiceCircuitBreakersMap(c *Client, i interface{}) map[string]BackendServiceCircuitBreakers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceCircuitBreakers{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceCircuitBreakers{}
	}

	items := make(map[string]BackendServiceCircuitBreakers)
	for k, item := range a {
		items[k] = *flattenBackendServiceCircuitBreakers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceCircuitBreakersSlice flattens the contents of BackendServiceCircuitBreakers from a JSON
// response object.
func flattenBackendServiceCircuitBreakersSlice(c *Client, i interface{}) []BackendServiceCircuitBreakers {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceCircuitBreakers{}
	}

	if len(a) == 0 {
		return []BackendServiceCircuitBreakers{}
	}

	items := make([]BackendServiceCircuitBreakers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceCircuitBreakers(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceCircuitBreakers expands an instance of BackendServiceCircuitBreakers into a JSON
// request object.
func expandBackendServiceCircuitBreakers(c *Client, f *BackendServiceCircuitBreakers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxRequestsPerConnection; !dcl.IsEmptyValueIndirect(v) {
		m["maxRequestsPerConnection"] = v
	}
	if v := f.MaxConnections; !dcl.IsEmptyValueIndirect(v) {
		m["maxConnections"] = v
	}
	if v := f.MaxPendingRequests; !dcl.IsEmptyValueIndirect(v) {
		m["maxPendingRequests"] = v
	}
	if v := f.MaxRequests; !dcl.IsEmptyValueIndirect(v) {
		m["maxRequests"] = v
	}
	if v := f.MaxRetries; !dcl.IsEmptyValueIndirect(v) {
		m["maxRetries"] = v
	}

	return m, nil
}

// flattenBackendServiceCircuitBreakers flattens an instance of BackendServiceCircuitBreakers from a JSON
// response object.
func flattenBackendServiceCircuitBreakers(c *Client, i interface{}) *BackendServiceCircuitBreakers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceCircuitBreakers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceCircuitBreakers
	}
	r.MaxRequestsPerConnection = dcl.FlattenInteger(m["maxRequestsPerConnection"])
	r.MaxConnections = dcl.FlattenInteger(m["maxConnections"])
	r.MaxPendingRequests = dcl.FlattenInteger(m["maxPendingRequests"])
	r.MaxRequests = dcl.FlattenInteger(m["maxRequests"])
	r.MaxRetries = dcl.FlattenInteger(m["maxRetries"])

	return r
}

// expandBackendServiceOutlierDetectionMap expands the contents of BackendServiceOutlierDetection into a JSON
// request object.
func expandBackendServiceOutlierDetectionMap(c *Client, f map[string]BackendServiceOutlierDetection) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceOutlierDetection(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceOutlierDetectionSlice expands the contents of BackendServiceOutlierDetection into a JSON
// request object.
func expandBackendServiceOutlierDetectionSlice(c *Client, f []BackendServiceOutlierDetection) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceOutlierDetection(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceOutlierDetectionMap flattens the contents of BackendServiceOutlierDetection from a JSON
// response object.
func flattenBackendServiceOutlierDetectionMap(c *Client, i interface{}) map[string]BackendServiceOutlierDetection {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceOutlierDetection{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceOutlierDetection{}
	}

	items := make(map[string]BackendServiceOutlierDetection)
	for k, item := range a {
		items[k] = *flattenBackendServiceOutlierDetection(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceOutlierDetectionSlice flattens the contents of BackendServiceOutlierDetection from a JSON
// response object.
func flattenBackendServiceOutlierDetectionSlice(c *Client, i interface{}) []BackendServiceOutlierDetection {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceOutlierDetection{}
	}

	if len(a) == 0 {
		return []BackendServiceOutlierDetection{}
	}

	items := make([]BackendServiceOutlierDetection, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceOutlierDetection(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceOutlierDetection expands an instance of BackendServiceOutlierDetection into a JSON
// request object.
func expandBackendServiceOutlierDetection(c *Client, f *BackendServiceOutlierDetection) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConsecutiveErrors; !dcl.IsEmptyValueIndirect(v) {
		m["consecutiveErrors"] = v
	}
	if v, err := expandBackendServiceOutlierDetectionInterval(c, f.Interval); err != nil {
		return nil, fmt.Errorf("error expanding Interval into interval: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["interval"] = v
	}
	if v, err := expandBackendServiceOutlierDetectionBaseEjectionTime(c, f.BaseEjectionTime); err != nil {
		return nil, fmt.Errorf("error expanding BaseEjectionTime into baseEjectionTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["baseEjectionTime"] = v
	}
	if v := f.MaxEjectionPercent; !dcl.IsEmptyValueIndirect(v) {
		m["maxEjectionPercent"] = v
	}
	if v := f.EnforcingConsecutiveErrors; !dcl.IsEmptyValueIndirect(v) {
		m["enforcingConsecutiveErrors"] = v
	}
	if v := f.EnforcingSuccessRate; !dcl.IsEmptyValueIndirect(v) {
		m["enforcingSuccessRate"] = v
	}
	if v := f.SuccessRateMinimumHosts; !dcl.IsEmptyValueIndirect(v) {
		m["successRateMinimumHosts"] = v
	}
	if v := f.SuccessRateRequestVolume; !dcl.IsEmptyValueIndirect(v) {
		m["successRateRequestVolume"] = v
	}
	if v := f.SuccessRateStdevFactor; !dcl.IsEmptyValueIndirect(v) {
		m["successRateStdevFactor"] = v
	}
	if v := f.ConsecutiveGatewayFailure; !dcl.IsEmptyValueIndirect(v) {
		m["consecutiveGatewayFailure"] = v
	}
	if v := f.EnforcingConsecutiveGatewayFailure; !dcl.IsEmptyValueIndirect(v) {
		m["enforcingConsecutiveGatewayFailure"] = v
	}

	return m, nil
}

// flattenBackendServiceOutlierDetection flattens an instance of BackendServiceOutlierDetection from a JSON
// response object.
func flattenBackendServiceOutlierDetection(c *Client, i interface{}) *BackendServiceOutlierDetection {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceOutlierDetection{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceOutlierDetection
	}
	r.ConsecutiveErrors = dcl.FlattenInteger(m["consecutiveErrors"])
	r.Interval = flattenBackendServiceOutlierDetectionInterval(c, m["interval"])
	r.BaseEjectionTime = flattenBackendServiceOutlierDetectionBaseEjectionTime(c, m["baseEjectionTime"])
	r.MaxEjectionPercent = dcl.FlattenInteger(m["maxEjectionPercent"])
	r.EnforcingConsecutiveErrors = dcl.FlattenInteger(m["enforcingConsecutiveErrors"])
	r.EnforcingSuccessRate = dcl.FlattenInteger(m["enforcingSuccessRate"])
	r.SuccessRateMinimumHosts = dcl.FlattenInteger(m["successRateMinimumHosts"])
	r.SuccessRateRequestVolume = dcl.FlattenInteger(m["successRateRequestVolume"])
	r.SuccessRateStdevFactor = dcl.FlattenInteger(m["successRateStdevFactor"])
	r.ConsecutiveGatewayFailure = dcl.FlattenInteger(m["consecutiveGatewayFailure"])
	r.EnforcingConsecutiveGatewayFailure = dcl.FlattenInteger(m["enforcingConsecutiveGatewayFailure"])

	return r
}

// expandBackendServiceOutlierDetectionIntervalMap expands the contents of BackendServiceOutlierDetectionInterval into a JSON
// request object.
func expandBackendServiceOutlierDetectionIntervalMap(c *Client, f map[string]BackendServiceOutlierDetectionInterval) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceOutlierDetectionInterval(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceOutlierDetectionIntervalSlice expands the contents of BackendServiceOutlierDetectionInterval into a JSON
// request object.
func expandBackendServiceOutlierDetectionIntervalSlice(c *Client, f []BackendServiceOutlierDetectionInterval) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceOutlierDetectionInterval(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceOutlierDetectionIntervalMap flattens the contents of BackendServiceOutlierDetectionInterval from a JSON
// response object.
func flattenBackendServiceOutlierDetectionIntervalMap(c *Client, i interface{}) map[string]BackendServiceOutlierDetectionInterval {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceOutlierDetectionInterval{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceOutlierDetectionInterval{}
	}

	items := make(map[string]BackendServiceOutlierDetectionInterval)
	for k, item := range a {
		items[k] = *flattenBackendServiceOutlierDetectionInterval(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceOutlierDetectionIntervalSlice flattens the contents of BackendServiceOutlierDetectionInterval from a JSON
// response object.
func flattenBackendServiceOutlierDetectionIntervalSlice(c *Client, i interface{}) []BackendServiceOutlierDetectionInterval {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceOutlierDetectionInterval{}
	}

	if len(a) == 0 {
		return []BackendServiceOutlierDetectionInterval{}
	}

	items := make([]BackendServiceOutlierDetectionInterval, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceOutlierDetectionInterval(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceOutlierDetectionInterval expands an instance of BackendServiceOutlierDetectionInterval into a JSON
// request object.
func expandBackendServiceOutlierDetectionInterval(c *Client, f *BackendServiceOutlierDetectionInterval) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenBackendServiceOutlierDetectionInterval flattens an instance of BackendServiceOutlierDetectionInterval from a JSON
// response object.
func flattenBackendServiceOutlierDetectionInterval(c *Client, i interface{}) *BackendServiceOutlierDetectionInterval {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceOutlierDetectionInterval{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceOutlierDetectionInterval
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandBackendServiceOutlierDetectionBaseEjectionTimeMap expands the contents of BackendServiceOutlierDetectionBaseEjectionTime into a JSON
// request object.
func expandBackendServiceOutlierDetectionBaseEjectionTimeMap(c *Client, f map[string]BackendServiceOutlierDetectionBaseEjectionTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceOutlierDetectionBaseEjectionTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceOutlierDetectionBaseEjectionTimeSlice expands the contents of BackendServiceOutlierDetectionBaseEjectionTime into a JSON
// request object.
func expandBackendServiceOutlierDetectionBaseEjectionTimeSlice(c *Client, f []BackendServiceOutlierDetectionBaseEjectionTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceOutlierDetectionBaseEjectionTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceOutlierDetectionBaseEjectionTimeMap flattens the contents of BackendServiceOutlierDetectionBaseEjectionTime from a JSON
// response object.
func flattenBackendServiceOutlierDetectionBaseEjectionTimeMap(c *Client, i interface{}) map[string]BackendServiceOutlierDetectionBaseEjectionTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceOutlierDetectionBaseEjectionTime{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceOutlierDetectionBaseEjectionTime{}
	}

	items := make(map[string]BackendServiceOutlierDetectionBaseEjectionTime)
	for k, item := range a {
		items[k] = *flattenBackendServiceOutlierDetectionBaseEjectionTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceOutlierDetectionBaseEjectionTimeSlice flattens the contents of BackendServiceOutlierDetectionBaseEjectionTime from a JSON
// response object.
func flattenBackendServiceOutlierDetectionBaseEjectionTimeSlice(c *Client, i interface{}) []BackendServiceOutlierDetectionBaseEjectionTime {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceOutlierDetectionBaseEjectionTime{}
	}

	if len(a) == 0 {
		return []BackendServiceOutlierDetectionBaseEjectionTime{}
	}

	items := make([]BackendServiceOutlierDetectionBaseEjectionTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceOutlierDetectionBaseEjectionTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceOutlierDetectionBaseEjectionTime expands an instance of BackendServiceOutlierDetectionBaseEjectionTime into a JSON
// request object.
func expandBackendServiceOutlierDetectionBaseEjectionTime(c *Client, f *BackendServiceOutlierDetectionBaseEjectionTime) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenBackendServiceOutlierDetectionBaseEjectionTime flattens an instance of BackendServiceOutlierDetectionBaseEjectionTime from a JSON
// response object.
func flattenBackendServiceOutlierDetectionBaseEjectionTime(c *Client, i interface{}) *BackendServiceOutlierDetectionBaseEjectionTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceOutlierDetectionBaseEjectionTime{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceOutlierDetectionBaseEjectionTime
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandBackendServiceMaxStreamDurationMap expands the contents of BackendServiceMaxStreamDuration into a JSON
// request object.
func expandBackendServiceMaxStreamDurationMap(c *Client, f map[string]BackendServiceMaxStreamDuration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBackendServiceMaxStreamDuration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBackendServiceMaxStreamDurationSlice expands the contents of BackendServiceMaxStreamDuration into a JSON
// request object.
func expandBackendServiceMaxStreamDurationSlice(c *Client, f []BackendServiceMaxStreamDuration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBackendServiceMaxStreamDuration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBackendServiceMaxStreamDurationMap flattens the contents of BackendServiceMaxStreamDuration from a JSON
// response object.
func flattenBackendServiceMaxStreamDurationMap(c *Client, i interface{}) map[string]BackendServiceMaxStreamDuration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackendServiceMaxStreamDuration{}
	}

	if len(a) == 0 {
		return map[string]BackendServiceMaxStreamDuration{}
	}

	items := make(map[string]BackendServiceMaxStreamDuration)
	for k, item := range a {
		items[k] = *flattenBackendServiceMaxStreamDuration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenBackendServiceMaxStreamDurationSlice flattens the contents of BackendServiceMaxStreamDuration from a JSON
// response object.
func flattenBackendServiceMaxStreamDurationSlice(c *Client, i interface{}) []BackendServiceMaxStreamDuration {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceMaxStreamDuration{}
	}

	if len(a) == 0 {
		return []BackendServiceMaxStreamDuration{}
	}

	items := make([]BackendServiceMaxStreamDuration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceMaxStreamDuration(c, item.(map[string]interface{})))
	}

	return items
}

// expandBackendServiceMaxStreamDuration expands an instance of BackendServiceMaxStreamDuration into a JSON
// request object.
func expandBackendServiceMaxStreamDuration(c *Client, f *BackendServiceMaxStreamDuration) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenBackendServiceMaxStreamDuration flattens an instance of BackendServiceMaxStreamDuration from a JSON
// response object.
func flattenBackendServiceMaxStreamDuration(c *Client, i interface{}) *BackendServiceMaxStreamDuration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BackendServiceMaxStreamDuration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBackendServiceMaxStreamDuration
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// flattenBackendServiceBackendsBalancingModeEnumSlice flattens the contents of BackendServiceBackendsBalancingModeEnum from a JSON
// response object.
func flattenBackendServiceBackendsBalancingModeEnumSlice(c *Client, i interface{}) []BackendServiceBackendsBalancingModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceBackendsBalancingModeEnum{}
	}

	if len(a) == 0 {
		return []BackendServiceBackendsBalancingModeEnum{}
	}

	items := make([]BackendServiceBackendsBalancingModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceBackendsBalancingModeEnum(item.(interface{})))
	}

	return items
}

// flattenBackendServiceBackendsBalancingModeEnum asserts that an interface is a string, and returns a
// pointer to a *BackendServiceBackendsBalancingModeEnum with the same value as that string.
func flattenBackendServiceBackendsBalancingModeEnum(i interface{}) *BackendServiceBackendsBalancingModeEnum {
	s, ok := i.(string)
	if !ok {
		return BackendServiceBackendsBalancingModeEnumRef("")
	}

	return BackendServiceBackendsBalancingModeEnumRef(s)
}

// flattenBackendServiceProtocolEnumSlice flattens the contents of BackendServiceProtocolEnum from a JSON
// response object.
func flattenBackendServiceProtocolEnumSlice(c *Client, i interface{}) []BackendServiceProtocolEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceProtocolEnum{}
	}

	if len(a) == 0 {
		return []BackendServiceProtocolEnum{}
	}

	items := make([]BackendServiceProtocolEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceProtocolEnum(item.(interface{})))
	}

	return items
}

// flattenBackendServiceProtocolEnum asserts that an interface is a string, and returns a
// pointer to a *BackendServiceProtocolEnum with the same value as that string.
func flattenBackendServiceProtocolEnum(i interface{}) *BackendServiceProtocolEnum {
	s, ok := i.(string)
	if !ok {
		return BackendServiceProtocolEnumRef("")
	}

	return BackendServiceProtocolEnumRef(s)
}

// flattenBackendServiceSessionAffinityEnumSlice flattens the contents of BackendServiceSessionAffinityEnum from a JSON
// response object.
func flattenBackendServiceSessionAffinityEnumSlice(c *Client, i interface{}) []BackendServiceSessionAffinityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceSessionAffinityEnum{}
	}

	if len(a) == 0 {
		return []BackendServiceSessionAffinityEnum{}
	}

	items := make([]BackendServiceSessionAffinityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceSessionAffinityEnum(item.(interface{})))
	}

	return items
}

// flattenBackendServiceSessionAffinityEnum asserts that an interface is a string, and returns a
// pointer to a *BackendServiceSessionAffinityEnum with the same value as that string.
func flattenBackendServiceSessionAffinityEnum(i interface{}) *BackendServiceSessionAffinityEnum {
	s, ok := i.(string)
	if !ok {
		return BackendServiceSessionAffinityEnumRef("")
	}

	return BackendServiceSessionAffinityEnumRef(s)
}

// flattenBackendServiceLoadBalancingSchemeEnumSlice flattens the contents of BackendServiceLoadBalancingSchemeEnum from a JSON
// response object.
func flattenBackendServiceLoadBalancingSchemeEnumSlice(c *Client, i interface{}) []BackendServiceLoadBalancingSchemeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceLoadBalancingSchemeEnum{}
	}

	if len(a) == 0 {
		return []BackendServiceLoadBalancingSchemeEnum{}
	}

	items := make([]BackendServiceLoadBalancingSchemeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceLoadBalancingSchemeEnum(item.(interface{})))
	}

	return items
}

// flattenBackendServiceLoadBalancingSchemeEnum asserts that an interface is a string, and returns a
// pointer to a *BackendServiceLoadBalancingSchemeEnum with the same value as that string.
func flattenBackendServiceLoadBalancingSchemeEnum(i interface{}) *BackendServiceLoadBalancingSchemeEnum {
	s, ok := i.(string)
	if !ok {
		return BackendServiceLoadBalancingSchemeEnumRef("")
	}

	return BackendServiceLoadBalancingSchemeEnumRef(s)
}

// flattenBackendServiceCdnPolicyCacheModeEnumSlice flattens the contents of BackendServiceCdnPolicyCacheModeEnum from a JSON
// response object.
func flattenBackendServiceCdnPolicyCacheModeEnumSlice(c *Client, i interface{}) []BackendServiceCdnPolicyCacheModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceCdnPolicyCacheModeEnum{}
	}

	if len(a) == 0 {
		return []BackendServiceCdnPolicyCacheModeEnum{}
	}

	items := make([]BackendServiceCdnPolicyCacheModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceCdnPolicyCacheModeEnum(item.(interface{})))
	}

	return items
}

// flattenBackendServiceCdnPolicyCacheModeEnum asserts that an interface is a string, and returns a
// pointer to a *BackendServiceCdnPolicyCacheModeEnum with the same value as that string.
func flattenBackendServiceCdnPolicyCacheModeEnum(i interface{}) *BackendServiceCdnPolicyCacheModeEnum {
	s, ok := i.(string)
	if !ok {
		return BackendServiceCdnPolicyCacheModeEnumRef("")
	}

	return BackendServiceCdnPolicyCacheModeEnumRef(s)
}

// flattenBackendServiceLocalityLbPolicyEnumSlice flattens the contents of BackendServiceLocalityLbPolicyEnum from a JSON
// response object.
func flattenBackendServiceLocalityLbPolicyEnumSlice(c *Client, i interface{}) []BackendServiceLocalityLbPolicyEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackendServiceLocalityLbPolicyEnum{}
	}

	if len(a) == 0 {
		return []BackendServiceLocalityLbPolicyEnum{}
	}

	items := make([]BackendServiceLocalityLbPolicyEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackendServiceLocalityLbPolicyEnum(item.(interface{})))
	}

	return items
}

// flattenBackendServiceLocalityLbPolicyEnum asserts that an interface is a string, and returns a
// pointer to a *BackendServiceLocalityLbPolicyEnum with the same value as that string.
func flattenBackendServiceLocalityLbPolicyEnum(i interface{}) *BackendServiceLocalityLbPolicyEnum {
	s, ok := i.(string)
	if !ok {
		return BackendServiceLocalityLbPolicyEnumRef("")
	}

	return BackendServiceLocalityLbPolicyEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *BackendService) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBackendService(b, c)
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

type backendServiceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         backendServiceApiOperation
}

func convertFieldDiffToBackendServiceOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]backendServiceDiff, error) {
	var diffs []backendServiceDiff
	for _, op := range ops {
		diff := backendServiceDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTobackendServiceApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTobackendServiceApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (backendServiceApiOperation, error) {
	switch op {

	case "updateBackendServiceUpdateOperation":
		return &updateBackendServiceUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
