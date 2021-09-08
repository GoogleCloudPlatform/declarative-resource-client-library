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
package alpha

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

func (r *InstanceGroupManager) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DistributionPolicy) {
		if err := r.DistributionPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CurrentActions) {
		if err := r.CurrentActions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UpdatePolicy) {
		if err := r.UpdatePolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StatefulPolicy) {
		if err := r.StatefulPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerDistributionPolicy) validate() error {
	return nil
}
func (r *InstanceGroupManagerDistributionPolicyZones) validate() error {
	return nil
}
func (r *InstanceGroupManagerVersions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TargetSize) {
		if err := r.TargetSize.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerFixedOrPercent) validate() error {
	return nil
}
func (r *InstanceGroupManagerCurrentActions) validate() error {
	return nil
}
func (r *InstanceGroupManagerStatus) validate() error {
	if !dcl.IsEmptyValueIndirect(r.VersionTarget) {
		if err := r.VersionTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Stateful) {
		if err := r.Stateful.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerStatusVersionTarget) validate() error {
	return nil
}
func (r *InstanceGroupManagerStatusStateful) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PerInstanceConfigs) {
		if err := r.PerInstanceConfigs.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerStatusStatefulPerInstanceConfigs) validate() error {
	return nil
}
func (r *InstanceGroupManagerAutoHealingPolicies) validate() error {
	return nil
}
func (r *InstanceGroupManagerUpdatePolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MaxSurge) {
		if err := r.MaxSurge.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaxUnavailable) {
		if err := r.MaxUnavailable.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerNamedPorts) validate() error {
	return nil
}
func (r *InstanceGroupManagerStatefulPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PreservedState) {
		if err := r.PreservedState.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceGroupManagerStatefulPolicyPreservedState) validate() error {
	return nil
}
func (r *InstanceGroupManagerStatefulPolicyPreservedStateDisks) validate() error {
	return nil
}
func (r *InstanceGroupManager) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://www.googleapis.com/compute/beta/", params)
}

func (r *InstanceGroupManager) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	if dcl.IsZone(nr.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Get URL found")

}

func (r *InstanceGroupManager) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers", nr.basePath(), userBasePath, params), nil
	}

	if dcl.IsZone(nr.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers", nr.basePath(), userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid List URL found")

}

func (r *InstanceGroupManager) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers", nr.basePath(), userBasePath, params), nil
	}

	if dcl.IsZone(nr.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers", nr.basePath(), userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Create URL found")

}

func (r *InstanceGroupManager) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	if dcl.IsZone(nr.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Delete URL found")

}

// instanceGroupManagerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceGroupManagerApiOperation interface {
	do(context.Context, *InstanceGroupManager, *Client) error
}

// newUpdateInstanceGroupManagerPatchRequest creates a request for an
// InstanceGroupManager resource's Patch update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceGroupManagerPatchRequest(ctx context.Context, f *InstanceGroupManager, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandInstanceGroupManagerDistributionPolicy(c, f.DistributionPolicy); err != nil {
		return nil, fmt.Errorf("error expanding DistributionPolicy into distributionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["distributionPolicy"] = v
	}
	if v := f.InstanceTemplate; !dcl.IsEmptyValueIndirect(v) {
		req["instanceTemplate"] = v
	}
	if v, err := expandInstanceGroupManagerVersionsSlice(c, f.Versions); err != nil {
		return nil, fmt.Errorf("error expanding Versions into versions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["versions"] = v
	}
	if v := f.TargetPools; v != nil {
		req["targetPools"] = v
	}
	if v := f.BaseInstanceName; !dcl.IsEmptyValueIndirect(v) {
		req["baseInstanceName"] = v
	}
	if v, err := expandInstanceGroupManagerStatus(c, f.Status); err != nil {
		return nil, fmt.Errorf("error expanding Status into status: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["status"] = v
	}
	if v := f.TargetSize; !dcl.IsEmptyValueIndirect(v) {
		req["targetSize"] = v
	}
	if v, err := expandInstanceGroupManagerAutoHealingPoliciesSlice(c, f.AutoHealingPolicies); err != nil {
		return nil, fmt.Errorf("error expanding AutoHealingPolicies into autoHealingPolicies: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["autoHealingPolicies"] = v
	}
	if v, err := expandInstanceGroupManagerUpdatePolicy(c, f.UpdatePolicy); err != nil {
		return nil, fmt.Errorf("error expanding UpdatePolicy into updatePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["updatePolicy"] = v
	}
	if v, err := expandInstanceGroupManagerStatefulPolicy(c, f.StatefulPolicy); err != nil {
		return nil, fmt.Errorf("error expanding StatefulPolicy into statefulPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["statefulPolicy"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		req["serviceAccount"] = v
	}
	if v := f.FailoverAction; !dcl.IsEmptyValueIndirect(v) {
		req["failoverAction"] = v
	}
	b, err := c.getInstanceGroupManagerRaw(ctx, f)
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
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["fingerprint"] = rawFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateInstanceGroupManagerPatchRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceGroupManagerPatchRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceGroupManagerPatchOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceGroupManagerPatchOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	_, err := c.GetInstanceGroupManager(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "Patch")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceGroupManagerPatchRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateInstanceGroupManagerPatchRequest(c, req)
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
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceGroupManagerSetInstanceTemplateRequest creates a request for an
// InstanceGroupManager resource's setInstanceTemplate update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceGroupManagerSetInstanceTemplateRequest(ctx context.Context, f *InstanceGroupManager, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateInstanceGroupManagerSetInstanceTemplateRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceGroupManagerSetInstanceTemplateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceGroupManagerSetInstanceTemplateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceGroupManagerSetInstanceTemplateOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	_, err := c.GetInstanceGroupManager(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setInstanceTemplate")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceGroupManagerSetInstanceTemplateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateInstanceGroupManagerSetInstanceTemplateRequest(c, req)
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
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateInstanceGroupManagerSetTargetPoolsRequest creates a request for an
// InstanceGroupManager resource's setTargetPools update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceGroupManagerSetTargetPoolsRequest(ctx context.Context, f *InstanceGroupManager, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	return req, nil
}

// marshalUpdateInstanceGroupManagerSetTargetPoolsRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceGroupManagerSetTargetPoolsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceGroupManagerSetTargetPoolsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceGroupManagerSetTargetPoolsOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	_, err := c.GetInstanceGroupManager(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setTargetPools")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceGroupManagerSetTargetPoolsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateInstanceGroupManagerSetTargetPoolsRequest(c, req)
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
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceGroupManagerRaw(ctx context.Context, r *InstanceGroupManager, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceGroupManagerMaxPage {
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

type listInstanceGroupManagerOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listInstanceGroupManager(ctx context.Context, r *InstanceGroupManager, pageToken string, pageSize int32) ([]*InstanceGroupManager, string, error) {
	b, err := c.listInstanceGroupManagerRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceGroupManagerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*InstanceGroupManager
	for _, v := range m.Items {
		res, err := unmarshalMapInstanceGroupManager(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstanceGroupManager(ctx context.Context, f func(*InstanceGroupManager) bool, resources []*InstanceGroupManager) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstanceGroupManager(ctx, res)
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

type deleteInstanceGroupManagerOperation struct{}

func (op *deleteInstanceGroupManagerOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	r, err := c.GetInstanceGroupManager(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "InstanceGroupManager not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetInstanceGroupManager checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetInstanceGroupManager(ctx, r)
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
type createInstanceGroupManagerOperation struct {
	response map[string]interface{}
}

func (op *createInstanceGroupManagerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInstanceGroupManagerOperation) do(ctx context.Context, r *InstanceGroupManager, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetInstanceGroupManager(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceGroupManagerRaw(ctx context.Context, r *InstanceGroupManager) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) instanceGroupManagerDiffsForRawDesired(ctx context.Context, rawDesired *InstanceGroupManager, opts ...dcl.ApplyOption) (initial, desired *InstanceGroupManager, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *InstanceGroupManager
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*InstanceGroupManager); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected InstanceGroupManager, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstanceGroupManager(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a InstanceGroupManager resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve InstanceGroupManager resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that InstanceGroupManager resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceGroupManagerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for InstanceGroupManager: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for InstanceGroupManager: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceGroupManagerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for InstanceGroupManager: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceGroupManagerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for InstanceGroupManager: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstanceGroupManager(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceGroupManagerInitialState(rawInitial, rawDesired *InstanceGroupManager) (*InstanceGroupManager, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceGroupManagerDesiredState(rawDesired, rawInitial *InstanceGroupManager, opts ...dcl.ApplyOption) (*InstanceGroupManager, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DistributionPolicy = canonicalizeInstanceGroupManagerDistributionPolicy(rawDesired.DistributionPolicy, nil, opts...)
		rawDesired.CurrentActions = canonicalizeInstanceGroupManagerCurrentActions(rawDesired.CurrentActions, nil, opts...)
		rawDesired.Status = canonicalizeInstanceGroupManagerStatus(rawDesired.Status, nil, opts...)
		rawDesired.UpdatePolicy = canonicalizeInstanceGroupManagerUpdatePolicy(rawDesired.UpdatePolicy, nil, opts...)
		rawDesired.StatefulPolicy = canonicalizeInstanceGroupManagerStatefulPolicy(rawDesired.StatefulPolicy, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &InstanceGroupManager{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.DistributionPolicy = canonicalizeInstanceGroupManagerDistributionPolicy(rawDesired.DistributionPolicy, rawInitial.DistributionPolicy, opts...)
	if dcl.NameToSelfLink(rawDesired.InstanceTemplate, rawInitial.InstanceTemplate) {
		canonicalDesired.InstanceTemplate = rawInitial.InstanceTemplate
	} else {
		canonicalDesired.InstanceTemplate = rawDesired.InstanceTemplate
	}
	canonicalDesired.Versions = canonicalizeInstanceGroupManagerVersionsSlice(rawDesired.Versions, rawInitial.Versions, opts...)
	if dcl.IsZeroValue(rawDesired.TargetPools) {
		canonicalDesired.TargetPools = rawInitial.TargetPools
	} else {
		canonicalDesired.TargetPools = rawDesired.TargetPools
	}
	if dcl.StringCanonicalize(rawDesired.BaseInstanceName, rawInitial.BaseInstanceName) {
		canonicalDesired.BaseInstanceName = rawInitial.BaseInstanceName
	} else {
		canonicalDesired.BaseInstanceName = rawDesired.BaseInstanceName
	}
	if dcl.IsZeroValue(rawDesired.TargetSize) {
		canonicalDesired.TargetSize = rawInitial.TargetSize
	} else {
		canonicalDesired.TargetSize = rawDesired.TargetSize
	}
	canonicalDesired.AutoHealingPolicies = canonicalizeInstanceGroupManagerAutoHealingPoliciesSlice(rawDesired.AutoHealingPolicies, rawInitial.AutoHealingPolicies, opts...)
	canonicalDesired.UpdatePolicy = canonicalizeInstanceGroupManagerUpdatePolicy(rawDesired.UpdatePolicy, rawInitial.UpdatePolicy, opts...)
	canonicalDesired.NamedPorts = canonicalizeInstanceGroupManagerNamedPortsSlice(rawDesired.NamedPorts, rawInitial.NamedPorts, opts...)
	canonicalDesired.StatefulPolicy = canonicalizeInstanceGroupManagerStatefulPolicy(rawDesired.StatefulPolicy, rawInitial.StatefulPolicy, opts...)
	if dcl.NameToSelfLink(rawDesired.ServiceAccount, rawInitial.ServiceAccount) {
		canonicalDesired.ServiceAccount = rawInitial.ServiceAccount
	} else {
		canonicalDesired.ServiceAccount = rawDesired.ServiceAccount
	}
	if dcl.IsZeroValue(rawDesired.FailoverAction) {
		canonicalDesired.FailoverAction = rawInitial.FailoverAction
	} else {
		canonicalDesired.FailoverAction = rawDesired.FailoverAction
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeInstanceGroupManagerNewState(c *Client, rawNew, rawDesired *InstanceGroupManager) (*InstanceGroupManager, error) {

	if dcl.IsNotReturnedByServer(rawNew.Id) && dcl.IsNotReturnedByServer(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.CreationTimestamp) && dcl.IsNotReturnedByServer(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
		if dcl.StringCanonicalize(rawDesired.CreationTimestamp, rawNew.CreationTimestamp) {
			rawNew.CreationTimestamp = rawDesired.CreationTimestamp
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Zone) && dcl.IsNotReturnedByServer(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Region) && dcl.IsNotReturnedByServer(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.DistributionPolicy) && dcl.IsNotReturnedByServer(rawDesired.DistributionPolicy) {
		rawNew.DistributionPolicy = rawDesired.DistributionPolicy
	} else {
		rawNew.DistributionPolicy = canonicalizeNewInstanceGroupManagerDistributionPolicy(c, rawDesired.DistributionPolicy, rawNew.DistributionPolicy)
	}

	if dcl.IsNotReturnedByServer(rawNew.InstanceTemplate) && dcl.IsNotReturnedByServer(rawDesired.InstanceTemplate) {
		rawNew.InstanceTemplate = rawDesired.InstanceTemplate
	} else {
		if dcl.NameToSelfLink(rawDesired.InstanceTemplate, rawNew.InstanceTemplate) {
			rawNew.InstanceTemplate = rawDesired.InstanceTemplate
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Versions) && dcl.IsNotReturnedByServer(rawDesired.Versions) {
		rawNew.Versions = rawDesired.Versions
	} else {
		rawNew.Versions = canonicalizeNewInstanceGroupManagerVersionsSlice(c, rawDesired.Versions, rawNew.Versions)
	}

	if dcl.IsNotReturnedByServer(rawNew.InstanceGroup) && dcl.IsNotReturnedByServer(rawDesired.InstanceGroup) {
		rawNew.InstanceGroup = rawDesired.InstanceGroup
	} else {
		if dcl.NameToSelfLink(rawDesired.InstanceGroup, rawNew.InstanceGroup) {
			rawNew.InstanceGroup = rawDesired.InstanceGroup
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.TargetPools) && dcl.IsNotReturnedByServer(rawDesired.TargetPools) {
		rawNew.TargetPools = rawDesired.TargetPools
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.BaseInstanceName) && dcl.IsNotReturnedByServer(rawDesired.BaseInstanceName) {
		rawNew.BaseInstanceName = rawDesired.BaseInstanceName
	} else {
		if dcl.StringCanonicalize(rawDesired.BaseInstanceName, rawNew.BaseInstanceName) {
			rawNew.BaseInstanceName = rawDesired.BaseInstanceName
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Fingerprint) && dcl.IsNotReturnedByServer(rawDesired.Fingerprint) {
		rawNew.Fingerprint = rawDesired.Fingerprint
	} else {
		if dcl.StringCanonicalize(rawDesired.Fingerprint, rawNew.Fingerprint) {
			rawNew.Fingerprint = rawDesired.Fingerprint
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CurrentActions) && dcl.IsNotReturnedByServer(rawDesired.CurrentActions) {
		rawNew.CurrentActions = rawDesired.CurrentActions
	} else {
		rawNew.CurrentActions = canonicalizeNewInstanceGroupManagerCurrentActions(c, rawDesired.CurrentActions, rawNew.CurrentActions)
	}

	if dcl.IsNotReturnedByServer(rawNew.Status) && dcl.IsNotReturnedByServer(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewInstanceGroupManagerStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsNotReturnedByServer(rawNew.TargetSize) && dcl.IsNotReturnedByServer(rawDesired.TargetSize) {
		rawNew.TargetSize = rawDesired.TargetSize
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.SelfLink) && dcl.IsNotReturnedByServer(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.AutoHealingPolicies) && dcl.IsNotReturnedByServer(rawDesired.AutoHealingPolicies) {
		rawNew.AutoHealingPolicies = rawDesired.AutoHealingPolicies
	} else {
		rawNew.AutoHealingPolicies = canonicalizeNewInstanceGroupManagerAutoHealingPoliciesSlice(c, rawDesired.AutoHealingPolicies, rawNew.AutoHealingPolicies)
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdatePolicy) && dcl.IsNotReturnedByServer(rawDesired.UpdatePolicy) {
		rawNew.UpdatePolicy = rawDesired.UpdatePolicy
	} else {
		rawNew.UpdatePolicy = canonicalizeNewInstanceGroupManagerUpdatePolicy(c, rawDesired.UpdatePolicy, rawNew.UpdatePolicy)
	}

	if dcl.IsNotReturnedByServer(rawNew.NamedPorts) && dcl.IsNotReturnedByServer(rawDesired.NamedPorts) {
		rawNew.NamedPorts = rawDesired.NamedPorts
	} else {
		rawNew.NamedPorts = canonicalizeNewInstanceGroupManagerNamedPortsSlice(c, rawDesired.NamedPorts, rawNew.NamedPorts)
	}

	if dcl.IsNotReturnedByServer(rawNew.StatefulPolicy) && dcl.IsNotReturnedByServer(rawDesired.StatefulPolicy) {
		rawNew.StatefulPolicy = rawDesired.StatefulPolicy
	} else {
		rawNew.StatefulPolicy = canonicalizeNewInstanceGroupManagerStatefulPolicy(c, rawDesired.StatefulPolicy, rawNew.StatefulPolicy)
	}

	if dcl.IsNotReturnedByServer(rawNew.ServiceAccount) && dcl.IsNotReturnedByServer(rawDesired.ServiceAccount) {
		rawNew.ServiceAccount = rawDesired.ServiceAccount
	} else {
		if dcl.NameToSelfLink(rawDesired.ServiceAccount, rawNew.ServiceAccount) {
			rawNew.ServiceAccount = rawDesired.ServiceAccount
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.FailoverAction) && dcl.IsNotReturnedByServer(rawDesired.FailoverAction) {
		rawNew.FailoverAction = rawDesired.FailoverAction
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceGroupManagerDistributionPolicy(des, initial *InstanceGroupManagerDistributionPolicy, opts ...dcl.ApplyOption) *InstanceGroupManagerDistributionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerDistributionPolicy{}

	cDes.Zones = canonicalizeInstanceGroupManagerDistributionPolicyZonesSlice(des.Zones, initial.Zones, opts...)
	if dcl.IsZeroValue(des.TargetShape) {
		des.TargetShape = initial.TargetShape
	} else {
		cDes.TargetShape = des.TargetShape
	}

	return cDes
}

func canonicalizeInstanceGroupManagerDistributionPolicySlice(des, initial []InstanceGroupManagerDistributionPolicy, opts ...dcl.ApplyOption) []InstanceGroupManagerDistributionPolicy {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerDistributionPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerDistributionPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerDistributionPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerDistributionPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerDistributionPolicy(c *Client, des, nw *InstanceGroupManagerDistributionPolicy) *InstanceGroupManagerDistributionPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerDistributionPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Zones = canonicalizeNewInstanceGroupManagerDistributionPolicyZonesSlice(c, des.Zones, nw.Zones)

	return nw
}

func canonicalizeNewInstanceGroupManagerDistributionPolicySet(c *Client, des, nw []InstanceGroupManagerDistributionPolicy) []InstanceGroupManagerDistributionPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerDistributionPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerDistributionPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerDistributionPolicySlice(c *Client, des, nw []InstanceGroupManagerDistributionPolicy) []InstanceGroupManagerDistributionPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerDistributionPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerDistributionPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerDistributionPolicyZones(des, initial *InstanceGroupManagerDistributionPolicyZones, opts ...dcl.ApplyOption) *InstanceGroupManagerDistributionPolicyZones {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerDistributionPolicyZones{}

	if dcl.StringCanonicalize(des.Zone, initial.Zone) || dcl.IsZeroValue(des.Zone) {
		cDes.Zone = initial.Zone
	} else {
		cDes.Zone = des.Zone
	}

	return cDes
}

func canonicalizeInstanceGroupManagerDistributionPolicyZonesSlice(des, initial []InstanceGroupManagerDistributionPolicyZones, opts ...dcl.ApplyOption) []InstanceGroupManagerDistributionPolicyZones {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerDistributionPolicyZones, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerDistributionPolicyZones(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerDistributionPolicyZones, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerDistributionPolicyZones(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerDistributionPolicyZones(c *Client, des, nw *InstanceGroupManagerDistributionPolicyZones) *InstanceGroupManagerDistributionPolicyZones {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerDistributionPolicyZones while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Zone, nw.Zone) {
		nw.Zone = des.Zone
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerDistributionPolicyZonesSet(c *Client, des, nw []InstanceGroupManagerDistributionPolicyZones) []InstanceGroupManagerDistributionPolicyZones {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerDistributionPolicyZones
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerDistributionPolicyZonesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerDistributionPolicyZonesSlice(c *Client, des, nw []InstanceGroupManagerDistributionPolicyZones) []InstanceGroupManagerDistributionPolicyZones {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerDistributionPolicyZones
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerDistributionPolicyZones(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerVersions(des, initial *InstanceGroupManagerVersions, opts ...dcl.ApplyOption) *InstanceGroupManagerVersions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerVersions{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.NameToSelfLink(des.InstanceTemplate, initial.InstanceTemplate) || dcl.IsZeroValue(des.InstanceTemplate) {
		cDes.InstanceTemplate = initial.InstanceTemplate
	} else {
		cDes.InstanceTemplate = des.InstanceTemplate
	}
	cDes.TargetSize = canonicalizeInstanceGroupManagerFixedOrPercent(des.TargetSize, initial.TargetSize, opts...)

	return cDes
}

func canonicalizeInstanceGroupManagerVersionsSlice(des, initial []InstanceGroupManagerVersions, opts ...dcl.ApplyOption) []InstanceGroupManagerVersions {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerVersions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerVersions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerVersions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerVersions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerVersions(c *Client, des, nw *InstanceGroupManagerVersions) *InstanceGroupManagerVersions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerVersions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.NameToSelfLink(des.InstanceTemplate, nw.InstanceTemplate) {
		nw.InstanceTemplate = des.InstanceTemplate
	}
	nw.TargetSize = canonicalizeNewInstanceGroupManagerFixedOrPercent(c, des.TargetSize, nw.TargetSize)

	return nw
}

func canonicalizeNewInstanceGroupManagerVersionsSet(c *Client, des, nw []InstanceGroupManagerVersions) []InstanceGroupManagerVersions {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerVersions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerVersionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerVersionsSlice(c *Client, des, nw []InstanceGroupManagerVersions) []InstanceGroupManagerVersions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerVersions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerVersions(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerFixedOrPercent(des, initial *InstanceGroupManagerFixedOrPercent, opts ...dcl.ApplyOption) *InstanceGroupManagerFixedOrPercent {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerFixedOrPercent{}

	if dcl.IsZeroValue(des.Fixed) {
		des.Fixed = initial.Fixed
	} else {
		cDes.Fixed = des.Fixed
	}
	if dcl.IsZeroValue(des.Percent) {
		des.Percent = initial.Percent
	} else {
		cDes.Percent = des.Percent
	}

	return cDes
}

func canonicalizeInstanceGroupManagerFixedOrPercentSlice(des, initial []InstanceGroupManagerFixedOrPercent, opts ...dcl.ApplyOption) []InstanceGroupManagerFixedOrPercent {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerFixedOrPercent, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerFixedOrPercent(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerFixedOrPercent, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerFixedOrPercent(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerFixedOrPercent(c *Client, des, nw *InstanceGroupManagerFixedOrPercent) *InstanceGroupManagerFixedOrPercent {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerFixedOrPercent while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerFixedOrPercentSet(c *Client, des, nw []InstanceGroupManagerFixedOrPercent) []InstanceGroupManagerFixedOrPercent {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerFixedOrPercent
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerFixedOrPercentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerFixedOrPercentSlice(c *Client, des, nw []InstanceGroupManagerFixedOrPercent) []InstanceGroupManagerFixedOrPercent {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerFixedOrPercent
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerFixedOrPercent(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerCurrentActions(des, initial *InstanceGroupManagerCurrentActions, opts ...dcl.ApplyOption) *InstanceGroupManagerCurrentActions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerCurrentActions{}

	return cDes
}

func canonicalizeInstanceGroupManagerCurrentActionsSlice(des, initial []InstanceGroupManagerCurrentActions, opts ...dcl.ApplyOption) []InstanceGroupManagerCurrentActions {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerCurrentActions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerCurrentActions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerCurrentActions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerCurrentActions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerCurrentActions(c *Client, des, nw *InstanceGroupManagerCurrentActions) *InstanceGroupManagerCurrentActions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerCurrentActions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerCurrentActionsSet(c *Client, des, nw []InstanceGroupManagerCurrentActions) []InstanceGroupManagerCurrentActions {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerCurrentActions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerCurrentActionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerCurrentActionsSlice(c *Client, des, nw []InstanceGroupManagerCurrentActions) []InstanceGroupManagerCurrentActions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerCurrentActions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerCurrentActions(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatus(des, initial *InstanceGroupManagerStatus, opts ...dcl.ApplyOption) *InstanceGroupManagerStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatus{}

	return cDes
}

func canonicalizeInstanceGroupManagerStatusSlice(des, initial []InstanceGroupManagerStatus, opts ...dcl.ApplyOption) []InstanceGroupManagerStatus {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatus(c *Client, des, nw *InstanceGroupManagerStatus) *InstanceGroupManagerStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.IsStable, nw.IsStable) {
		nw.IsStable = des.IsStable
	}
	nw.VersionTarget = canonicalizeNewInstanceGroupManagerStatusVersionTarget(c, des.VersionTarget, nw.VersionTarget)
	nw.Stateful = canonicalizeNewInstanceGroupManagerStatusStateful(c, des.Stateful, nw.Stateful)
	if dcl.StringCanonicalize(des.Autoscaler, nw.Autoscaler) {
		nw.Autoscaler = des.Autoscaler
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatusSet(c *Client, des, nw []InstanceGroupManagerStatus) []InstanceGroupManagerStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatusSlice(c *Client, des, nw []InstanceGroupManagerStatus) []InstanceGroupManagerStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatus(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatusVersionTarget(des, initial *InstanceGroupManagerStatusVersionTarget, opts ...dcl.ApplyOption) *InstanceGroupManagerStatusVersionTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatusVersionTarget{}

	return cDes
}

func canonicalizeInstanceGroupManagerStatusVersionTargetSlice(des, initial []InstanceGroupManagerStatusVersionTarget, opts ...dcl.ApplyOption) []InstanceGroupManagerStatusVersionTarget {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatusVersionTarget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatusVersionTarget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatusVersionTarget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatusVersionTarget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatusVersionTarget(c *Client, des, nw *InstanceGroupManagerStatusVersionTarget) *InstanceGroupManagerStatusVersionTarget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatusVersionTarget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.IsReached, nw.IsReached) {
		nw.IsReached = des.IsReached
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatusVersionTargetSet(c *Client, des, nw []InstanceGroupManagerStatusVersionTarget) []InstanceGroupManagerStatusVersionTarget {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatusVersionTarget
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatusVersionTargetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatusVersionTargetSlice(c *Client, des, nw []InstanceGroupManagerStatusVersionTarget) []InstanceGroupManagerStatusVersionTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatusVersionTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatusVersionTarget(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatusStateful(des, initial *InstanceGroupManagerStatusStateful, opts ...dcl.ApplyOption) *InstanceGroupManagerStatusStateful {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatusStateful{}

	return cDes
}

func canonicalizeInstanceGroupManagerStatusStatefulSlice(des, initial []InstanceGroupManagerStatusStateful, opts ...dcl.ApplyOption) []InstanceGroupManagerStatusStateful {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatusStateful, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatusStateful(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatusStateful, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatusStateful(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatusStateful(c *Client, des, nw *InstanceGroupManagerStatusStateful) *InstanceGroupManagerStatusStateful {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatusStateful while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.HasStatefulConfig, nw.HasStatefulConfig) {
		nw.HasStatefulConfig = des.HasStatefulConfig
	}
	nw.PerInstanceConfigs = canonicalizeNewInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, des.PerInstanceConfigs, nw.PerInstanceConfigs)
	if dcl.BoolCanonicalize(des.IsStateful, nw.IsStateful) {
		nw.IsStateful = des.IsStateful
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatusStatefulSet(c *Client, des, nw []InstanceGroupManagerStatusStateful) []InstanceGroupManagerStatusStateful {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatusStateful
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatusStatefulNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatusStatefulSlice(c *Client, des, nw []InstanceGroupManagerStatusStateful) []InstanceGroupManagerStatusStateful {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatusStateful
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatusStateful(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatusStatefulPerInstanceConfigs(des, initial *InstanceGroupManagerStatusStatefulPerInstanceConfigs, opts ...dcl.ApplyOption) *InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatusStatefulPerInstanceConfigs{}

	if dcl.BoolCanonicalize(des.AllEffective, initial.AllEffective) || dcl.IsZeroValue(des.AllEffective) {
		cDes.AllEffective = initial.AllEffective
	} else {
		cDes.AllEffective = des.AllEffective
	}

	return cDes
}

func canonicalizeInstanceGroupManagerStatusStatefulPerInstanceConfigsSlice(des, initial []InstanceGroupManagerStatusStatefulPerInstanceConfigs, opts ...dcl.ApplyOption) []InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatusStatefulPerInstanceConfigs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatusStatefulPerInstanceConfigs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatusStatefulPerInstanceConfigs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatusStatefulPerInstanceConfigs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatusStatefulPerInstanceConfigs(c *Client, des, nw *InstanceGroupManagerStatusStatefulPerInstanceConfigs) *InstanceGroupManagerStatusStatefulPerInstanceConfigs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatusStatefulPerInstanceConfigs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.AllEffective, nw.AllEffective) {
		nw.AllEffective = des.AllEffective
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatusStatefulPerInstanceConfigsSet(c *Client, des, nw []InstanceGroupManagerStatusStatefulPerInstanceConfigs) []InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatusStatefulPerInstanceConfigs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatusStatefulPerInstanceConfigsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatusStatefulPerInstanceConfigsSlice(c *Client, des, nw []InstanceGroupManagerStatusStatefulPerInstanceConfigs) []InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatusStatefulPerInstanceConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerAutoHealingPolicies(des, initial *InstanceGroupManagerAutoHealingPolicies, opts ...dcl.ApplyOption) *InstanceGroupManagerAutoHealingPolicies {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerAutoHealingPolicies{}

	if dcl.NameToSelfLink(des.HealthCheck, initial.HealthCheck) || dcl.IsZeroValue(des.HealthCheck) {
		cDes.HealthCheck = initial.HealthCheck
	} else {
		cDes.HealthCheck = des.HealthCheck
	}
	if dcl.IsZeroValue(des.InitialDelaySec) {
		des.InitialDelaySec = initial.InitialDelaySec
	} else {
		cDes.InitialDelaySec = des.InitialDelaySec
	}

	return cDes
}

func canonicalizeInstanceGroupManagerAutoHealingPoliciesSlice(des, initial []InstanceGroupManagerAutoHealingPolicies, opts ...dcl.ApplyOption) []InstanceGroupManagerAutoHealingPolicies {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerAutoHealingPolicies, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerAutoHealingPolicies(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerAutoHealingPolicies, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerAutoHealingPolicies(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerAutoHealingPolicies(c *Client, des, nw *InstanceGroupManagerAutoHealingPolicies) *InstanceGroupManagerAutoHealingPolicies {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerAutoHealingPolicies while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.NameToSelfLink(des.HealthCheck, nw.HealthCheck) {
		nw.HealthCheck = des.HealthCheck
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerAutoHealingPoliciesSet(c *Client, des, nw []InstanceGroupManagerAutoHealingPolicies) []InstanceGroupManagerAutoHealingPolicies {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerAutoHealingPolicies
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerAutoHealingPoliciesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerAutoHealingPoliciesSlice(c *Client, des, nw []InstanceGroupManagerAutoHealingPolicies) []InstanceGroupManagerAutoHealingPolicies {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerAutoHealingPolicies
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerAutoHealingPolicies(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerUpdatePolicy(des, initial *InstanceGroupManagerUpdatePolicy, opts ...dcl.ApplyOption) *InstanceGroupManagerUpdatePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerUpdatePolicy{}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.IsZeroValue(des.InstanceRedistributionType) {
		des.InstanceRedistributionType = initial.InstanceRedistributionType
	} else {
		cDes.InstanceRedistributionType = des.InstanceRedistributionType
	}
	if dcl.IsZeroValue(des.MinimalAction) {
		des.MinimalAction = initial.MinimalAction
	} else {
		cDes.MinimalAction = des.MinimalAction
	}
	cDes.MaxSurge = canonicalizeInstanceGroupManagerFixedOrPercent(des.MaxSurge, initial.MaxSurge, opts...)
	cDes.MaxUnavailable = canonicalizeInstanceGroupManagerFixedOrPercent(des.MaxUnavailable, initial.MaxUnavailable, opts...)
	if dcl.IsZeroValue(des.ReplacementMethod) {
		des.ReplacementMethod = initial.ReplacementMethod
	} else {
		cDes.ReplacementMethod = des.ReplacementMethod
	}
	if dcl.IsZeroValue(des.MostDisruptiveAllowedAction) {
		des.MostDisruptiveAllowedAction = initial.MostDisruptiveAllowedAction
	} else {
		cDes.MostDisruptiveAllowedAction = des.MostDisruptiveAllowedAction
	}
	if dcl.IsZeroValue(des.MinReadySec) {
		des.MinReadySec = initial.MinReadySec
	} else {
		cDes.MinReadySec = des.MinReadySec
	}

	return cDes
}

func canonicalizeInstanceGroupManagerUpdatePolicySlice(des, initial []InstanceGroupManagerUpdatePolicy, opts ...dcl.ApplyOption) []InstanceGroupManagerUpdatePolicy {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerUpdatePolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerUpdatePolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerUpdatePolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerUpdatePolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerUpdatePolicy(c *Client, des, nw *InstanceGroupManagerUpdatePolicy) *InstanceGroupManagerUpdatePolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerUpdatePolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MaxSurge = canonicalizeNewInstanceGroupManagerFixedOrPercent(c, des.MaxSurge, nw.MaxSurge)
	nw.MaxUnavailable = canonicalizeNewInstanceGroupManagerFixedOrPercent(c, des.MaxUnavailable, nw.MaxUnavailable)

	return nw
}

func canonicalizeNewInstanceGroupManagerUpdatePolicySet(c *Client, des, nw []InstanceGroupManagerUpdatePolicy) []InstanceGroupManagerUpdatePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerUpdatePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerUpdatePolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerUpdatePolicySlice(c *Client, des, nw []InstanceGroupManagerUpdatePolicy) []InstanceGroupManagerUpdatePolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerUpdatePolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerUpdatePolicy(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerNamedPorts(des, initial *InstanceGroupManagerNamedPorts, opts ...dcl.ApplyOption) *InstanceGroupManagerNamedPorts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerNamedPorts{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	} else {
		cDes.Port = des.Port
	}

	return cDes
}

func canonicalizeInstanceGroupManagerNamedPortsSlice(des, initial []InstanceGroupManagerNamedPorts, opts ...dcl.ApplyOption) []InstanceGroupManagerNamedPorts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerNamedPorts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerNamedPorts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerNamedPorts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerNamedPorts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerNamedPorts(c *Client, des, nw *InstanceGroupManagerNamedPorts) *InstanceGroupManagerNamedPorts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerNamedPorts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerNamedPortsSet(c *Client, des, nw []InstanceGroupManagerNamedPorts) []InstanceGroupManagerNamedPorts {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerNamedPorts
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerNamedPortsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerNamedPortsSlice(c *Client, des, nw []InstanceGroupManagerNamedPorts) []InstanceGroupManagerNamedPorts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerNamedPorts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerNamedPorts(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatefulPolicy(des, initial *InstanceGroupManagerStatefulPolicy, opts ...dcl.ApplyOption) *InstanceGroupManagerStatefulPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatefulPolicy{}

	cDes.PreservedState = canonicalizeInstanceGroupManagerStatefulPolicyPreservedState(des.PreservedState, initial.PreservedState, opts...)

	return cDes
}

func canonicalizeInstanceGroupManagerStatefulPolicySlice(des, initial []InstanceGroupManagerStatefulPolicy, opts ...dcl.ApplyOption) []InstanceGroupManagerStatefulPolicy {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatefulPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatefulPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatefulPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatefulPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatefulPolicy(c *Client, des, nw *InstanceGroupManagerStatefulPolicy) *InstanceGroupManagerStatefulPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatefulPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.PreservedState = canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedState(c, des.PreservedState, nw.PreservedState)

	return nw
}

func canonicalizeNewInstanceGroupManagerStatefulPolicySet(c *Client, des, nw []InstanceGroupManagerStatefulPolicy) []InstanceGroupManagerStatefulPolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatefulPolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatefulPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatefulPolicySlice(c *Client, des, nw []InstanceGroupManagerStatefulPolicy) []InstanceGroupManagerStatefulPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatefulPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatefulPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatefulPolicyPreservedState(des, initial *InstanceGroupManagerStatefulPolicyPreservedState, opts ...dcl.ApplyOption) *InstanceGroupManagerStatefulPolicyPreservedState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatefulPolicyPreservedState{}

	if dcl.IsZeroValue(des.Disks) {
		des.Disks = initial.Disks
	} else {
		cDes.Disks = des.Disks
	}

	return cDes
}

func canonicalizeInstanceGroupManagerStatefulPolicyPreservedStateSlice(des, initial []InstanceGroupManagerStatefulPolicyPreservedState, opts ...dcl.ApplyOption) []InstanceGroupManagerStatefulPolicyPreservedState {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatefulPolicyPreservedState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatefulPolicyPreservedState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatefulPolicyPreservedState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatefulPolicyPreservedState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedState(c *Client, des, nw *InstanceGroupManagerStatefulPolicyPreservedState) *InstanceGroupManagerStatefulPolicyPreservedState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatefulPolicyPreservedState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedStateSet(c *Client, des, nw []InstanceGroupManagerStatefulPolicyPreservedState) []InstanceGroupManagerStatefulPolicyPreservedState {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatefulPolicyPreservedState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatefulPolicyPreservedStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedStateSlice(c *Client, des, nw []InstanceGroupManagerStatefulPolicyPreservedState) []InstanceGroupManagerStatefulPolicyPreservedState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatefulPolicyPreservedState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedState(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceGroupManagerStatefulPolicyPreservedStateDisks(des, initial *InstanceGroupManagerStatefulPolicyPreservedStateDisks, opts ...dcl.ApplyOption) *InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceGroupManagerStatefulPolicyPreservedStateDisks{}

	if dcl.IsZeroValue(des.AutoDelete) {
		des.AutoDelete = initial.AutoDelete
	} else {
		cDes.AutoDelete = des.AutoDelete
	}

	return cDes
}

func canonicalizeInstanceGroupManagerStatefulPolicyPreservedStateDisksSlice(des, initial []InstanceGroupManagerStatefulPolicyPreservedStateDisks, opts ...dcl.ApplyOption) []InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceGroupManagerStatefulPolicyPreservedStateDisks, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceGroupManagerStatefulPolicyPreservedStateDisks(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceGroupManagerStatefulPolicyPreservedStateDisks, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceGroupManagerStatefulPolicyPreservedStateDisks(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedStateDisks(c *Client, des, nw *InstanceGroupManagerStatefulPolicyPreservedStateDisks) *InstanceGroupManagerStatefulPolicyPreservedStateDisks {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceGroupManagerStatefulPolicyPreservedStateDisks while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedStateDisksSet(c *Client, des, nw []InstanceGroupManagerStatefulPolicyPreservedStateDisks) []InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceGroupManagerStatefulPolicyPreservedStateDisks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceGroupManagerStatefulPolicyPreservedStateDisksNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedStateDisksSlice(c *Client, des, nw []InstanceGroupManagerStatefulPolicyPreservedStateDisks) []InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceGroupManagerStatefulPolicyPreservedStateDisks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceGroupManagerStatefulPolicyPreservedStateDisks(c, &d, &n))
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
func diffInstanceGroupManager(c *Client, desired, actual *InstanceGroupManager, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DistributionPolicy, actual.DistributionPolicy, dcl.Info{ObjectFunction: compareInstanceGroupManagerDistributionPolicyNewStyle, EmptyObject: EmptyInstanceGroupManagerDistributionPolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DistributionPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceTemplate, actual.InstanceTemplate, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("InstanceTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Versions, actual.Versions, dcl.Info{ObjectFunction: compareInstanceGroupManagerVersionsNewStyle, EmptyObject: EmptyInstanceGroupManagerVersions, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("Versions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceGroup, actual.InstanceGroup, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstanceGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetPools, actual.TargetPools, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("TargetPools")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BaseInstanceName, actual.BaseInstanceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("BaseInstanceName")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CurrentActions, actual.CurrentActions, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceGroupManagerCurrentActionsNewStyle, EmptyObject: EmptyInstanceGroupManagerCurrentActions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CurrentActions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceGroupManagerStatusNewStyle, EmptyObject: EmptyInstanceGroupManagerStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetSize, actual.TargetSize, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("TargetSize")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.AutoHealingPolicies, actual.AutoHealingPolicies, dcl.Info{ObjectFunction: compareInstanceGroupManagerAutoHealingPoliciesNewStyle, EmptyObject: EmptyInstanceGroupManagerAutoHealingPolicies, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("AutoHealingPolicies")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdatePolicy, actual.UpdatePolicy, dcl.Info{ObjectFunction: compareInstanceGroupManagerUpdatePolicyNewStyle, EmptyObject: EmptyInstanceGroupManagerUpdatePolicy, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("UpdatePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NamedPorts, actual.NamedPorts, dcl.Info{ObjectFunction: compareInstanceGroupManagerNamedPortsNewStyle, EmptyObject: EmptyInstanceGroupManagerNamedPorts, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NamedPorts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatefulPolicy, actual.StatefulPolicy, dcl.Info{ObjectFunction: compareInstanceGroupManagerStatefulPolicyNewStyle, EmptyObject: EmptyInstanceGroupManagerStatefulPolicy, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("StatefulPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailoverAction, actual.FailoverAction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("FailoverAction")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareInstanceGroupManagerDistributionPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerDistributionPolicy)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerDistributionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerDistributionPolicy or *InstanceGroupManagerDistributionPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerDistributionPolicy)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerDistributionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerDistributionPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Zones, actual.Zones, dcl.Info{ObjectFunction: compareInstanceGroupManagerDistributionPolicyZonesNewStyle, EmptyObject: EmptyInstanceGroupManagerDistributionPolicyZones, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zones")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetShape, actual.TargetShape, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("TargetShape")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerDistributionPolicyZonesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerDistributionPolicyZones)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerDistributionPolicyZones)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerDistributionPolicyZones or *InstanceGroupManagerDistributionPolicyZones", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerDistributionPolicyZones)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerDistributionPolicyZones)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerDistributionPolicyZones", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerVersionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerVersions)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerVersions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerVersions or *InstanceGroupManagerVersions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerVersions)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerVersions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerVersions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceTemplate, actual.InstanceTemplate, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("InstanceTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetSize, actual.TargetSize, dcl.Info{ObjectFunction: compareInstanceGroupManagerFixedOrPercentNewStyle, EmptyObject: EmptyInstanceGroupManagerFixedOrPercent, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("TargetSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerFixedOrPercentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerFixedOrPercent)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerFixedOrPercent)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerFixedOrPercent or *InstanceGroupManagerFixedOrPercent", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerFixedOrPercent)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerFixedOrPercent)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerFixedOrPercent", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fixed, actual.Fixed, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("Fixed")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Calculated, actual.Calculated, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Calculated")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerCurrentActionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerCurrentActions)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerCurrentActions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerCurrentActions or *InstanceGroupManagerCurrentActions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerCurrentActions)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerCurrentActions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerCurrentActions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.None, actual.None, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("None")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Creating, actual.Creating, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Creating")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatingWithoutRetries, actual.CreatingWithoutRetries, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreatingWithoutRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Verifying, actual.Verifying, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Verifying")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Recreating, actual.Recreating, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Recreating")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deleting, actual.Deleting, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Deleting")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Abandoning, actual.Abandoning, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Abandoning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Restarting, actual.Restarting, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Restarting")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Refreshing, actual.Refreshing, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Refreshing")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatus)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatus or *InstanceGroupManagerStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatus)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IsStable, actual.IsStable, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsStable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionTarget, actual.VersionTarget, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceGroupManagerStatusVersionTargetNewStyle, EmptyObject: EmptyInstanceGroupManagerStatusVersionTarget, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Stateful, actual.Stateful, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceGroupManagerStatusStatefulNewStyle, EmptyObject: EmptyInstanceGroupManagerStatusStateful, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Stateful")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Autoscaler, actual.Autoscaler, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Autoscaler")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatusVersionTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatusVersionTarget)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatusVersionTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatusVersionTarget or *InstanceGroupManagerStatusVersionTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatusVersionTarget)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatusVersionTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatusVersionTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IsReached, actual.IsReached, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsReached")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatusStatefulNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatusStateful)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatusStateful)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatusStateful or *InstanceGroupManagerStatusStateful", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatusStateful)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatusStateful)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatusStateful", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HasStatefulConfig, actual.HasStatefulConfig, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HasStatefulConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerInstanceConfigs, actual.PerInstanceConfigs, dcl.Info{OutputOnly: true, ObjectFunction: compareInstanceGroupManagerStatusStatefulPerInstanceConfigsNewStyle, EmptyObject: EmptyInstanceGroupManagerStatusStatefulPerInstanceConfigs, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PerInstanceConfigs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IsStateful, actual.IsStateful, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsStateful")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatusStatefulPerInstanceConfigsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatusStatefulPerInstanceConfigs)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatusStatefulPerInstanceConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatusStatefulPerInstanceConfigs or *InstanceGroupManagerStatusStatefulPerInstanceConfigs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatusStatefulPerInstanceConfigs)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatusStatefulPerInstanceConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatusStatefulPerInstanceConfigs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllEffective, actual.AllEffective, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("AllEffective")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerAutoHealingPoliciesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerAutoHealingPolicies)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerAutoHealingPolicies)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerAutoHealingPolicies or *InstanceGroupManagerAutoHealingPolicies", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerAutoHealingPolicies)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerAutoHealingPolicies)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerAutoHealingPolicies", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HealthCheck, actual.HealthCheck, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("HealthCheck")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitialDelaySec, actual.InitialDelaySec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("InitialDelaySec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerUpdatePolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerUpdatePolicy)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerUpdatePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerUpdatePolicy or *InstanceGroupManagerUpdatePolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerUpdatePolicy)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerUpdatePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerUpdatePolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InstanceRedistributionType, actual.InstanceRedistributionType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("InstanceRedistributionType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinimalAction, actual.MinimalAction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("MinimalAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxSurge, actual.MaxSurge, dcl.Info{ObjectFunction: compareInstanceGroupManagerFixedOrPercentNewStyle, EmptyObject: EmptyInstanceGroupManagerFixedOrPercent, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("MaxSurge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxUnavailable, actual.MaxUnavailable, dcl.Info{ObjectFunction: compareInstanceGroupManagerFixedOrPercentNewStyle, EmptyObject: EmptyInstanceGroupManagerFixedOrPercent, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("MaxUnavailable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReplacementMethod, actual.ReplacementMethod, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("ReplacementMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MostDisruptiveAllowedAction, actual.MostDisruptiveAllowedAction, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("MostDisruptiveAllowedAction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinReadySec, actual.MinReadySec, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("MinReadySec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerNamedPortsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerNamedPorts)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerNamedPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerNamedPorts or *InstanceGroupManagerNamedPorts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerNamedPorts)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerNamedPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerNamedPorts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatefulPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatefulPolicy)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatefulPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatefulPolicy or *InstanceGroupManagerStatefulPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatefulPolicy)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatefulPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatefulPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PreservedState, actual.PreservedState, dcl.Info{ObjectFunction: compareInstanceGroupManagerStatefulPolicyPreservedStateNewStyle, EmptyObject: EmptyInstanceGroupManagerStatefulPolicyPreservedState, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("PreservedState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatefulPolicyPreservedStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatefulPolicyPreservedState)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatefulPolicyPreservedState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatefulPolicyPreservedState or *InstanceGroupManagerStatefulPolicyPreservedState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatefulPolicyPreservedState)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatefulPolicyPreservedState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatefulPolicyPreservedState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Disks, actual.Disks, dcl.Info{ObjectFunction: compareInstanceGroupManagerStatefulPolicyPreservedStateDisksNewStyle, EmptyObject: EmptyInstanceGroupManagerStatefulPolicyPreservedStateDisks, OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("Disks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceGroupManagerStatefulPolicyPreservedStateDisksNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceGroupManagerStatefulPolicyPreservedStateDisks)
	if !ok {
		desiredNotPointer, ok := d.(InstanceGroupManagerStatefulPolicyPreservedStateDisks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatefulPolicyPreservedStateDisks or *InstanceGroupManagerStatefulPolicyPreservedStateDisks", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceGroupManagerStatefulPolicyPreservedStateDisks)
	if !ok {
		actualNotPointer, ok := a.(InstanceGroupManagerStatefulPolicyPreservedStateDisks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceGroupManagerStatefulPolicyPreservedStateDisks", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutoDelete, actual.AutoDelete, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceGroupManagerPatchOperation")}, fn.AddNest("AutoDelete")); len(ds) != 0 || err != nil {
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
func (r *InstanceGroupManager) urlNormalized() *InstanceGroupManager {
	normalized := dcl.Copy(*r).(InstanceGroupManager)
	normalized.CreationTimestamp = dcl.SelfLinkToName(r.CreationTimestamp)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.InstanceTemplate = r.InstanceTemplate
	normalized.InstanceGroup = dcl.SelfLinkToName(r.InstanceGroup)
	normalized.BaseInstanceName = dcl.SelfLinkToName(r.BaseInstanceName)
	normalized.Fingerprint = dcl.SelfLinkToName(r.Fingerprint)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.ServiceAccount = dcl.SelfLinkToName(r.ServiceAccount)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *InstanceGroupManager) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "Patch" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		if dcl.IsRegion(nr.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}", nr.basePath(), userBasePath, fields), nil
		}

		if dcl.IsZone(nr.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}", nr.basePath(), userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	if updateName == "setInstanceTemplate" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		if dcl.IsRegion(nr.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}/setInstanceTemplate", nr.basePath(), userBasePath, fields), nil
		}

		if dcl.IsZone(nr.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}/setInstanceTemplate", nr.basePath(), userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}
	if updateName == "setTargetPools" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		if dcl.IsRegion(nr.Location) {
			return dcl.URL("projects/{{project}}/regions/{{location}}/instanceGroupManagers/{{name}}/setTargetPools", nr.basePath(), userBasePath, fields), nil
		}

		if dcl.IsZone(nr.Location) {
			return dcl.URL("projects/{{project}}/zones/{{location}}/instanceGroupManagers/{{name}}/setTargetPools", nr.basePath(), userBasePath, fields), nil
		}

		return "", fmt.Errorf("No valid update URL for %s", updateName)

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the InstanceGroupManager resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *InstanceGroupManager) marshal(c *Client) ([]byte, error) {
	m, err := expandInstanceGroupManager(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling InstanceGroupManager: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInstanceGroupManager decodes JSON responses into the InstanceGroupManager resource schema.
func unmarshalInstanceGroupManager(b []byte, c *Client) (*InstanceGroupManager, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstanceGroupManager(m, c)
}

func unmarshalMapInstanceGroupManager(m map[string]interface{}, c *Client) (*InstanceGroupManager, error) {

	return flattenInstanceGroupManager(c, m), nil
}

// expandInstanceGroupManager expands InstanceGroupManager into a JSON request object.
func expandInstanceGroupManager(c *Client, f *InstanceGroupManager) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandInstanceGroupManagerDistributionPolicy(c, f.DistributionPolicy); err != nil {
		return nil, fmt.Errorf("error expanding DistributionPolicy into distributionPolicy: %w", err)
	} else if v != nil {
		m["distributionPolicy"] = v
	}
	if v := f.InstanceTemplate; !dcl.IsEmptyValueIndirect(v) {
		m["instanceTemplate"] = v
	}
	if v, err := expandInstanceGroupManagerVersionsSlice(c, f.Versions); err != nil {
		return nil, fmt.Errorf("error expanding Versions into versions: %w", err)
	} else {
		m["versions"] = v
	}
	m["targetPools"] = f.TargetPools
	if v := f.BaseInstanceName; !dcl.IsEmptyValueIndirect(v) {
		m["baseInstanceName"] = v
	}
	if v := f.TargetSize; !dcl.IsEmptyValueIndirect(v) {
		m["targetSize"] = v
	}
	if v, err := expandInstanceGroupManagerAutoHealingPoliciesSlice(c, f.AutoHealingPolicies); err != nil {
		return nil, fmt.Errorf("error expanding AutoHealingPolicies into autoHealingPolicies: %w", err)
	} else {
		m["autoHealingPolicies"] = v
	}
	if v, err := expandInstanceGroupManagerUpdatePolicy(c, f.UpdatePolicy); err != nil {
		return nil, fmt.Errorf("error expanding UpdatePolicy into updatePolicy: %w", err)
	} else if v != nil {
		m["updatePolicy"] = v
	}
	if v, err := expandInstanceGroupManagerNamedPortsSlice(c, f.NamedPorts); err != nil {
		return nil, fmt.Errorf("error expanding NamedPorts into namedPorts: %w", err)
	} else {
		m["namedPorts"] = v
	}
	if v, err := expandInstanceGroupManagerStatefulPolicy(c, f.StatefulPolicy); err != nil {
		return nil, fmt.Errorf("error expanding StatefulPolicy into statefulPolicy: %w", err)
	} else if v != nil {
		m["statefulPolicy"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.FailoverAction; !dcl.IsEmptyValueIndirect(v) {
		m["failoverAction"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenInstanceGroupManager flattens InstanceGroupManager from a JSON request object into the
// InstanceGroupManager type.
func flattenInstanceGroupManager(c *Client, i interface{}) *InstanceGroupManager {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &InstanceGroupManager{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Zone = dcl.FlattenString(m["zone"])
	res.Region = dcl.FlattenString(m["region"])
	res.DistributionPolicy = flattenInstanceGroupManagerDistributionPolicy(c, m["distributionPolicy"])
	res.InstanceTemplate = dcl.FlattenString(m["instanceTemplate"])
	res.Versions = flattenInstanceGroupManagerVersionsSlice(c, m["versions"])
	res.InstanceGroup = dcl.FlattenString(m["instanceGroup"])
	res.TargetPools = dcl.FlattenStringSlice(m["targetPools"])
	res.BaseInstanceName = dcl.FlattenString(m["baseInstanceName"])
	res.Fingerprint = dcl.FlattenString(m["fingerprint"])
	res.CurrentActions = flattenInstanceGroupManagerCurrentActions(c, m["currentActions"])
	res.Status = flattenInstanceGroupManagerStatus(c, m["status"])
	res.TargetSize = dcl.FlattenInteger(m["targetSize"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.AutoHealingPolicies = flattenInstanceGroupManagerAutoHealingPoliciesSlice(c, m["autoHealingPolicies"])
	res.UpdatePolicy = flattenInstanceGroupManagerUpdatePolicy(c, m["updatePolicy"])
	res.NamedPorts = flattenInstanceGroupManagerNamedPortsSlice(c, m["namedPorts"])
	res.StatefulPolicy = flattenInstanceGroupManagerStatefulPolicy(c, m["statefulPolicy"])
	res.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	res.FailoverAction = flattenInstanceGroupManagerFailoverActionEnum(m["failoverAction"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandInstanceGroupManagerDistributionPolicyMap expands the contents of InstanceGroupManagerDistributionPolicy into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyMap(c *Client, f map[string]InstanceGroupManagerDistributionPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerDistributionPolicySlice expands the contents of InstanceGroupManagerDistributionPolicy into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicySlice(c *Client, f []InstanceGroupManagerDistributionPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerDistributionPolicyMap flattens the contents of InstanceGroupManagerDistributionPolicy from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyMap(c *Client, i interface{}) map[string]InstanceGroupManagerDistributionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerDistributionPolicy{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerDistributionPolicy{}
	}

	items := make(map[string]InstanceGroupManagerDistributionPolicy)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerDistributionPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerDistributionPolicySlice flattens the contents of InstanceGroupManagerDistributionPolicy from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicySlice(c *Client, i interface{}) []InstanceGroupManagerDistributionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerDistributionPolicy{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerDistributionPolicy{}
	}

	items := make([]InstanceGroupManagerDistributionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerDistributionPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerDistributionPolicy expands an instance of InstanceGroupManagerDistributionPolicy into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicy(c *Client, f *InstanceGroupManagerDistributionPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceGroupManagerDistributionPolicyZonesSlice(c, f.Zones); err != nil {
		return nil, fmt.Errorf("error expanding Zones into zones: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["zones"] = v
	}
	if v := f.TargetShape; !dcl.IsEmptyValueIndirect(v) {
		m["targetShape"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerDistributionPolicy flattens an instance of InstanceGroupManagerDistributionPolicy from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicy(c *Client, i interface{}) *InstanceGroupManagerDistributionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerDistributionPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerDistributionPolicy
	}
	r.Zones = flattenInstanceGroupManagerDistributionPolicyZonesSlice(c, m["zones"])
	r.TargetShape = flattenInstanceGroupManagerDistributionPolicyTargetShapeEnum(m["targetShape"])

	return r
}

// expandInstanceGroupManagerDistributionPolicyZonesMap expands the contents of InstanceGroupManagerDistributionPolicyZones into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyZonesMap(c *Client, f map[string]InstanceGroupManagerDistributionPolicyZones) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicyZones(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerDistributionPolicyZonesSlice expands the contents of InstanceGroupManagerDistributionPolicyZones into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyZonesSlice(c *Client, f []InstanceGroupManagerDistributionPolicyZones) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerDistributionPolicyZones(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerDistributionPolicyZonesMap flattens the contents of InstanceGroupManagerDistributionPolicyZones from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyZonesMap(c *Client, i interface{}) map[string]InstanceGroupManagerDistributionPolicyZones {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerDistributionPolicyZones{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerDistributionPolicyZones{}
	}

	items := make(map[string]InstanceGroupManagerDistributionPolicyZones)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerDistributionPolicyZones(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerDistributionPolicyZonesSlice flattens the contents of InstanceGroupManagerDistributionPolicyZones from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyZonesSlice(c *Client, i interface{}) []InstanceGroupManagerDistributionPolicyZones {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerDistributionPolicyZones{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerDistributionPolicyZones{}
	}

	items := make([]InstanceGroupManagerDistributionPolicyZones, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerDistributionPolicyZones(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerDistributionPolicyZones expands an instance of InstanceGroupManagerDistributionPolicyZones into a JSON
// request object.
func expandInstanceGroupManagerDistributionPolicyZones(c *Client, f *InstanceGroupManagerDistributionPolicyZones) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerDistributionPolicyZones flattens an instance of InstanceGroupManagerDistributionPolicyZones from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyZones(c *Client, i interface{}) *InstanceGroupManagerDistributionPolicyZones {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerDistributionPolicyZones{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerDistributionPolicyZones
	}
	r.Zone = dcl.FlattenString(m["zone"])

	return r
}

// expandInstanceGroupManagerVersionsMap expands the contents of InstanceGroupManagerVersions into a JSON
// request object.
func expandInstanceGroupManagerVersionsMap(c *Client, f map[string]InstanceGroupManagerVersions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerVersions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerVersionsSlice expands the contents of InstanceGroupManagerVersions into a JSON
// request object.
func expandInstanceGroupManagerVersionsSlice(c *Client, f []InstanceGroupManagerVersions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerVersions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerVersionsMap flattens the contents of InstanceGroupManagerVersions from a JSON
// response object.
func flattenInstanceGroupManagerVersionsMap(c *Client, i interface{}) map[string]InstanceGroupManagerVersions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerVersions{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerVersions{}
	}

	items := make(map[string]InstanceGroupManagerVersions)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerVersions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerVersionsSlice flattens the contents of InstanceGroupManagerVersions from a JSON
// response object.
func flattenInstanceGroupManagerVersionsSlice(c *Client, i interface{}) []InstanceGroupManagerVersions {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerVersions{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerVersions{}
	}

	items := make([]InstanceGroupManagerVersions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerVersions(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerVersions expands an instance of InstanceGroupManagerVersions into a JSON
// request object.
func expandInstanceGroupManagerVersions(c *Client, f *InstanceGroupManagerVersions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.InstanceTemplate; !dcl.IsEmptyValueIndirect(v) {
		m["instanceTemplate"] = v
	}
	if v, err := expandInstanceGroupManagerFixedOrPercent(c, f.TargetSize); err != nil {
		return nil, fmt.Errorf("error expanding TargetSize into targetSize: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["targetSize"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerVersions flattens an instance of InstanceGroupManagerVersions from a JSON
// response object.
func flattenInstanceGroupManagerVersions(c *Client, i interface{}) *InstanceGroupManagerVersions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerVersions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerVersions
	}
	r.Name = dcl.FlattenString(m["name"])
	r.InstanceTemplate = dcl.FlattenString(m["instanceTemplate"])
	r.TargetSize = flattenInstanceGroupManagerFixedOrPercent(c, m["targetSize"])

	return r
}

// expandInstanceGroupManagerFixedOrPercentMap expands the contents of InstanceGroupManagerFixedOrPercent into a JSON
// request object.
func expandInstanceGroupManagerFixedOrPercentMap(c *Client, f map[string]InstanceGroupManagerFixedOrPercent) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerFixedOrPercent(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerFixedOrPercentSlice expands the contents of InstanceGroupManagerFixedOrPercent into a JSON
// request object.
func expandInstanceGroupManagerFixedOrPercentSlice(c *Client, f []InstanceGroupManagerFixedOrPercent) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerFixedOrPercent(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerFixedOrPercentMap flattens the contents of InstanceGroupManagerFixedOrPercent from a JSON
// response object.
func flattenInstanceGroupManagerFixedOrPercentMap(c *Client, i interface{}) map[string]InstanceGroupManagerFixedOrPercent {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerFixedOrPercent{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerFixedOrPercent{}
	}

	items := make(map[string]InstanceGroupManagerFixedOrPercent)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerFixedOrPercent(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerFixedOrPercentSlice flattens the contents of InstanceGroupManagerFixedOrPercent from a JSON
// response object.
func flattenInstanceGroupManagerFixedOrPercentSlice(c *Client, i interface{}) []InstanceGroupManagerFixedOrPercent {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerFixedOrPercent{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerFixedOrPercent{}
	}

	items := make([]InstanceGroupManagerFixedOrPercent, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerFixedOrPercent(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerFixedOrPercent expands an instance of InstanceGroupManagerFixedOrPercent into a JSON
// request object.
func expandInstanceGroupManagerFixedOrPercent(c *Client, f *InstanceGroupManagerFixedOrPercent) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fixed; v != nil {
		m["fixed"] = v
	}
	if v := f.Percent; v != nil {
		m["percent"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerFixedOrPercent flattens an instance of InstanceGroupManagerFixedOrPercent from a JSON
// response object.
func flattenInstanceGroupManagerFixedOrPercent(c *Client, i interface{}) *InstanceGroupManagerFixedOrPercent {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerFixedOrPercent{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerFixedOrPercent
	}
	r.Fixed = dcl.FlattenInteger(m["fixed"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Calculated = dcl.FlattenInteger(m["calculated"])

	return r
}

// expandInstanceGroupManagerCurrentActionsMap expands the contents of InstanceGroupManagerCurrentActions into a JSON
// request object.
func expandInstanceGroupManagerCurrentActionsMap(c *Client, f map[string]InstanceGroupManagerCurrentActions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerCurrentActions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerCurrentActionsSlice expands the contents of InstanceGroupManagerCurrentActions into a JSON
// request object.
func expandInstanceGroupManagerCurrentActionsSlice(c *Client, f []InstanceGroupManagerCurrentActions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerCurrentActions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerCurrentActionsMap flattens the contents of InstanceGroupManagerCurrentActions from a JSON
// response object.
func flattenInstanceGroupManagerCurrentActionsMap(c *Client, i interface{}) map[string]InstanceGroupManagerCurrentActions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerCurrentActions{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerCurrentActions{}
	}

	items := make(map[string]InstanceGroupManagerCurrentActions)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerCurrentActions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerCurrentActionsSlice flattens the contents of InstanceGroupManagerCurrentActions from a JSON
// response object.
func flattenInstanceGroupManagerCurrentActionsSlice(c *Client, i interface{}) []InstanceGroupManagerCurrentActions {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerCurrentActions{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerCurrentActions{}
	}

	items := make([]InstanceGroupManagerCurrentActions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerCurrentActions(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerCurrentActions expands an instance of InstanceGroupManagerCurrentActions into a JSON
// request object.
func expandInstanceGroupManagerCurrentActions(c *Client, f *InstanceGroupManagerCurrentActions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenInstanceGroupManagerCurrentActions flattens an instance of InstanceGroupManagerCurrentActions from a JSON
// response object.
func flattenInstanceGroupManagerCurrentActions(c *Client, i interface{}) *InstanceGroupManagerCurrentActions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerCurrentActions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerCurrentActions
	}
	r.None = dcl.FlattenInteger(m["none"])
	r.Creating = dcl.FlattenInteger(m["creating"])
	r.CreatingWithoutRetries = dcl.FlattenInteger(m["creatingWithoutRetries"])
	r.Verifying = dcl.FlattenInteger(m["verifying"])
	r.Recreating = dcl.FlattenInteger(m["recreating"])
	r.Deleting = dcl.FlattenInteger(m["deleting"])
	r.Abandoning = dcl.FlattenInteger(m["abandoning"])
	r.Restarting = dcl.FlattenInteger(m["restarting"])
	r.Refreshing = dcl.FlattenInteger(m["refreshing"])

	return r
}

// expandInstanceGroupManagerStatusMap expands the contents of InstanceGroupManagerStatus into a JSON
// request object.
func expandInstanceGroupManagerStatusMap(c *Client, f map[string]InstanceGroupManagerStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatusSlice expands the contents of InstanceGroupManagerStatus into a JSON
// request object.
func expandInstanceGroupManagerStatusSlice(c *Client, f []InstanceGroupManagerStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatusMap flattens the contents of InstanceGroupManagerStatus from a JSON
// response object.
func flattenInstanceGroupManagerStatusMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatus{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatus{}
	}

	items := make(map[string]InstanceGroupManagerStatus)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatusSlice flattens the contents of InstanceGroupManagerStatus from a JSON
// response object.
func flattenInstanceGroupManagerStatusSlice(c *Client, i interface{}) []InstanceGroupManagerStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatus{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatus{}
	}

	items := make([]InstanceGroupManagerStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatus expands an instance of InstanceGroupManagerStatus into a JSON
// request object.
func expandInstanceGroupManagerStatus(c *Client, f *InstanceGroupManagerStatus) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenInstanceGroupManagerStatus flattens an instance of InstanceGroupManagerStatus from a JSON
// response object.
func flattenInstanceGroupManagerStatus(c *Client, i interface{}) *InstanceGroupManagerStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatus
	}
	r.IsStable = dcl.FlattenBool(m["isStable"])
	r.VersionTarget = flattenInstanceGroupManagerStatusVersionTarget(c, m["versionTarget"])
	r.Stateful = flattenInstanceGroupManagerStatusStateful(c, m["stateful"])
	r.Autoscaler = dcl.FlattenString(m["autoscaler"])

	return r
}

// expandInstanceGroupManagerStatusVersionTargetMap expands the contents of InstanceGroupManagerStatusVersionTarget into a JSON
// request object.
func expandInstanceGroupManagerStatusVersionTargetMap(c *Client, f map[string]InstanceGroupManagerStatusVersionTarget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatusVersionTarget(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatusVersionTargetSlice expands the contents of InstanceGroupManagerStatusVersionTarget into a JSON
// request object.
func expandInstanceGroupManagerStatusVersionTargetSlice(c *Client, f []InstanceGroupManagerStatusVersionTarget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatusVersionTarget(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatusVersionTargetMap flattens the contents of InstanceGroupManagerStatusVersionTarget from a JSON
// response object.
func flattenInstanceGroupManagerStatusVersionTargetMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatusVersionTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatusVersionTarget{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatusVersionTarget{}
	}

	items := make(map[string]InstanceGroupManagerStatusVersionTarget)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatusVersionTarget(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatusVersionTargetSlice flattens the contents of InstanceGroupManagerStatusVersionTarget from a JSON
// response object.
func flattenInstanceGroupManagerStatusVersionTargetSlice(c *Client, i interface{}) []InstanceGroupManagerStatusVersionTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatusVersionTarget{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatusVersionTarget{}
	}

	items := make([]InstanceGroupManagerStatusVersionTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatusVersionTarget(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatusVersionTarget expands an instance of InstanceGroupManagerStatusVersionTarget into a JSON
// request object.
func expandInstanceGroupManagerStatusVersionTarget(c *Client, f *InstanceGroupManagerStatusVersionTarget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenInstanceGroupManagerStatusVersionTarget flattens an instance of InstanceGroupManagerStatusVersionTarget from a JSON
// response object.
func flattenInstanceGroupManagerStatusVersionTarget(c *Client, i interface{}) *InstanceGroupManagerStatusVersionTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatusVersionTarget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatusVersionTarget
	}
	r.IsReached = dcl.FlattenBool(m["isReached"])

	return r
}

// expandInstanceGroupManagerStatusStatefulMap expands the contents of InstanceGroupManagerStatusStateful into a JSON
// request object.
func expandInstanceGroupManagerStatusStatefulMap(c *Client, f map[string]InstanceGroupManagerStatusStateful) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatusStateful(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatusStatefulSlice expands the contents of InstanceGroupManagerStatusStateful into a JSON
// request object.
func expandInstanceGroupManagerStatusStatefulSlice(c *Client, f []InstanceGroupManagerStatusStateful) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatusStateful(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatusStatefulMap flattens the contents of InstanceGroupManagerStatusStateful from a JSON
// response object.
func flattenInstanceGroupManagerStatusStatefulMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatusStateful {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatusStateful{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatusStateful{}
	}

	items := make(map[string]InstanceGroupManagerStatusStateful)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatusStateful(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatusStatefulSlice flattens the contents of InstanceGroupManagerStatusStateful from a JSON
// response object.
func flattenInstanceGroupManagerStatusStatefulSlice(c *Client, i interface{}) []InstanceGroupManagerStatusStateful {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatusStateful{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatusStateful{}
	}

	items := make([]InstanceGroupManagerStatusStateful, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatusStateful(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatusStateful expands an instance of InstanceGroupManagerStatusStateful into a JSON
// request object.
func expandInstanceGroupManagerStatusStateful(c *Client, f *InstanceGroupManagerStatusStateful) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenInstanceGroupManagerStatusStateful flattens an instance of InstanceGroupManagerStatusStateful from a JSON
// response object.
func flattenInstanceGroupManagerStatusStateful(c *Client, i interface{}) *InstanceGroupManagerStatusStateful {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatusStateful{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatusStateful
	}
	r.HasStatefulConfig = dcl.FlattenBool(m["hasStatefulConfig"])
	r.PerInstanceConfigs = flattenInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, m["perInstanceConfigs"])
	r.IsStateful = dcl.FlattenBool(m["isStateful"])

	return r
}

// expandInstanceGroupManagerStatusStatefulPerInstanceConfigsMap expands the contents of InstanceGroupManagerStatusStatefulPerInstanceConfigs into a JSON
// request object.
func expandInstanceGroupManagerStatusStatefulPerInstanceConfigsMap(c *Client, f map[string]InstanceGroupManagerStatusStatefulPerInstanceConfigs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatusStatefulPerInstanceConfigsSlice expands the contents of InstanceGroupManagerStatusStatefulPerInstanceConfigs into a JSON
// request object.
func expandInstanceGroupManagerStatusStatefulPerInstanceConfigsSlice(c *Client, f []InstanceGroupManagerStatusStatefulPerInstanceConfigs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatusStatefulPerInstanceConfigsMap flattens the contents of InstanceGroupManagerStatusStatefulPerInstanceConfigs from a JSON
// response object.
func flattenInstanceGroupManagerStatusStatefulPerInstanceConfigsMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	}

	items := make(map[string]InstanceGroupManagerStatusStatefulPerInstanceConfigs)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatusStatefulPerInstanceConfigsSlice flattens the contents of InstanceGroupManagerStatusStatefulPerInstanceConfigs from a JSON
// response object.
func flattenInstanceGroupManagerStatusStatefulPerInstanceConfigsSlice(c *Client, i interface{}) []InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	}

	items := make([]InstanceGroupManagerStatusStatefulPerInstanceConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatusStatefulPerInstanceConfigs(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatusStatefulPerInstanceConfigs expands an instance of InstanceGroupManagerStatusStatefulPerInstanceConfigs into a JSON
// request object.
func expandInstanceGroupManagerStatusStatefulPerInstanceConfigs(c *Client, f *InstanceGroupManagerStatusStatefulPerInstanceConfigs) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllEffective; !dcl.IsEmptyValueIndirect(v) {
		m["allEffective"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerStatusStatefulPerInstanceConfigs flattens an instance of InstanceGroupManagerStatusStatefulPerInstanceConfigs from a JSON
// response object.
func flattenInstanceGroupManagerStatusStatefulPerInstanceConfigs(c *Client, i interface{}) *InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatusStatefulPerInstanceConfigs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatusStatefulPerInstanceConfigs
	}
	r.AllEffective = dcl.FlattenBool(m["allEffective"])

	return r
}

// expandInstanceGroupManagerAutoHealingPoliciesMap expands the contents of InstanceGroupManagerAutoHealingPolicies into a JSON
// request object.
func expandInstanceGroupManagerAutoHealingPoliciesMap(c *Client, f map[string]InstanceGroupManagerAutoHealingPolicies) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerAutoHealingPolicies(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerAutoHealingPoliciesSlice expands the contents of InstanceGroupManagerAutoHealingPolicies into a JSON
// request object.
func expandInstanceGroupManagerAutoHealingPoliciesSlice(c *Client, f []InstanceGroupManagerAutoHealingPolicies) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerAutoHealingPolicies(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerAutoHealingPoliciesMap flattens the contents of InstanceGroupManagerAutoHealingPolicies from a JSON
// response object.
func flattenInstanceGroupManagerAutoHealingPoliciesMap(c *Client, i interface{}) map[string]InstanceGroupManagerAutoHealingPolicies {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerAutoHealingPolicies{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerAutoHealingPolicies{}
	}

	items := make(map[string]InstanceGroupManagerAutoHealingPolicies)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerAutoHealingPolicies(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerAutoHealingPoliciesSlice flattens the contents of InstanceGroupManagerAutoHealingPolicies from a JSON
// response object.
func flattenInstanceGroupManagerAutoHealingPoliciesSlice(c *Client, i interface{}) []InstanceGroupManagerAutoHealingPolicies {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerAutoHealingPolicies{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerAutoHealingPolicies{}
	}

	items := make([]InstanceGroupManagerAutoHealingPolicies, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerAutoHealingPolicies(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerAutoHealingPolicies expands an instance of InstanceGroupManagerAutoHealingPolicies into a JSON
// request object.
func expandInstanceGroupManagerAutoHealingPolicies(c *Client, f *InstanceGroupManagerAutoHealingPolicies) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HealthCheck; !dcl.IsEmptyValueIndirect(v) {
		m["healthCheck"] = v
	}
	if v := f.InitialDelaySec; !dcl.IsEmptyValueIndirect(v) {
		m["initialDelaySec"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerAutoHealingPolicies flattens an instance of InstanceGroupManagerAutoHealingPolicies from a JSON
// response object.
func flattenInstanceGroupManagerAutoHealingPolicies(c *Client, i interface{}) *InstanceGroupManagerAutoHealingPolicies {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerAutoHealingPolicies{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerAutoHealingPolicies
	}
	r.HealthCheck = dcl.FlattenString(m["healthCheck"])
	r.InitialDelaySec = dcl.FlattenInteger(m["initialDelaySec"])

	return r
}

// expandInstanceGroupManagerUpdatePolicyMap expands the contents of InstanceGroupManagerUpdatePolicy into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicyMap(c *Client, f map[string]InstanceGroupManagerUpdatePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerUpdatePolicySlice expands the contents of InstanceGroupManagerUpdatePolicy into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicySlice(c *Client, f []InstanceGroupManagerUpdatePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerUpdatePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerUpdatePolicyMap flattens the contents of InstanceGroupManagerUpdatePolicy from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicy{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicy{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicy)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicySlice flattens the contents of InstanceGroupManagerUpdatePolicy from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicySlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicy{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicy{}
	}

	items := make([]InstanceGroupManagerUpdatePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerUpdatePolicy expands an instance of InstanceGroupManagerUpdatePolicy into a JSON
// request object.
func expandInstanceGroupManagerUpdatePolicy(c *Client, f *InstanceGroupManagerUpdatePolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.InstanceRedistributionType; !dcl.IsEmptyValueIndirect(v) {
		m["instanceRedistributionType"] = v
	}
	if v := f.MinimalAction; !dcl.IsEmptyValueIndirect(v) {
		m["minimalAction"] = v
	}
	if v, err := expandInstanceGroupManagerFixedOrPercent(c, f.MaxSurge); err != nil {
		return nil, fmt.Errorf("error expanding MaxSurge into maxSurge: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxSurge"] = v
	}
	if v, err := expandInstanceGroupManagerFixedOrPercent(c, f.MaxUnavailable); err != nil {
		return nil, fmt.Errorf("error expanding MaxUnavailable into maxUnavailable: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxUnavailable"] = v
	}
	if v := f.ReplacementMethod; !dcl.IsEmptyValueIndirect(v) {
		m["replacementMethod"] = v
	}
	if v := f.MostDisruptiveAllowedAction; !dcl.IsEmptyValueIndirect(v) {
		m["mostDisruptiveAllowedAction"] = v
	}
	if v := f.MinReadySec; !dcl.IsEmptyValueIndirect(v) {
		m["minReadySec"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerUpdatePolicy flattens an instance of InstanceGroupManagerUpdatePolicy from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicy(c *Client, i interface{}) *InstanceGroupManagerUpdatePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerUpdatePolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerUpdatePolicy
	}
	r.Type = flattenInstanceGroupManagerUpdatePolicyTypeEnum(m["type"])
	r.InstanceRedistributionType = flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(m["instanceRedistributionType"])
	r.MinimalAction = flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(m["minimalAction"])
	r.MaxSurge = flattenInstanceGroupManagerFixedOrPercent(c, m["maxSurge"])
	r.MaxUnavailable = flattenInstanceGroupManagerFixedOrPercent(c, m["maxUnavailable"])
	r.ReplacementMethod = flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnum(m["replacementMethod"])
	r.MostDisruptiveAllowedAction = flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(m["mostDisruptiveAllowedAction"])
	r.MinReadySec = dcl.FlattenInteger(m["minReadySec"])

	return r
}

// expandInstanceGroupManagerNamedPortsMap expands the contents of InstanceGroupManagerNamedPorts into a JSON
// request object.
func expandInstanceGroupManagerNamedPortsMap(c *Client, f map[string]InstanceGroupManagerNamedPorts) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerNamedPorts(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerNamedPortsSlice expands the contents of InstanceGroupManagerNamedPorts into a JSON
// request object.
func expandInstanceGroupManagerNamedPortsSlice(c *Client, f []InstanceGroupManagerNamedPorts) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerNamedPorts(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerNamedPortsMap flattens the contents of InstanceGroupManagerNamedPorts from a JSON
// response object.
func flattenInstanceGroupManagerNamedPortsMap(c *Client, i interface{}) map[string]InstanceGroupManagerNamedPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerNamedPorts{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerNamedPorts{}
	}

	items := make(map[string]InstanceGroupManagerNamedPorts)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerNamedPorts(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerNamedPortsSlice flattens the contents of InstanceGroupManagerNamedPorts from a JSON
// response object.
func flattenInstanceGroupManagerNamedPortsSlice(c *Client, i interface{}) []InstanceGroupManagerNamedPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerNamedPorts{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerNamedPorts{}
	}

	items := make([]InstanceGroupManagerNamedPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerNamedPorts(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerNamedPorts expands an instance of InstanceGroupManagerNamedPorts into a JSON
// request object.
func expandInstanceGroupManagerNamedPorts(c *Client, f *InstanceGroupManagerNamedPorts) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerNamedPorts flattens an instance of InstanceGroupManagerNamedPorts from a JSON
// response object.
func flattenInstanceGroupManagerNamedPorts(c *Client, i interface{}) *InstanceGroupManagerNamedPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerNamedPorts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerNamedPorts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Port = dcl.FlattenInteger(m["port"])

	return r
}

// expandInstanceGroupManagerStatefulPolicyMap expands the contents of InstanceGroupManagerStatefulPolicy into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyMap(c *Client, f map[string]InstanceGroupManagerStatefulPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatefulPolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatefulPolicySlice expands the contents of InstanceGroupManagerStatefulPolicy into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicySlice(c *Client, f []InstanceGroupManagerStatefulPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatefulPolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatefulPolicyMap flattens the contents of InstanceGroupManagerStatefulPolicy from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatefulPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatefulPolicy{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatefulPolicy{}
	}

	items := make(map[string]InstanceGroupManagerStatefulPolicy)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatefulPolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatefulPolicySlice flattens the contents of InstanceGroupManagerStatefulPolicy from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicySlice(c *Client, i interface{}) []InstanceGroupManagerStatefulPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatefulPolicy{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatefulPolicy{}
	}

	items := make([]InstanceGroupManagerStatefulPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatefulPolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatefulPolicy expands an instance of InstanceGroupManagerStatefulPolicy into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicy(c *Client, f *InstanceGroupManagerStatefulPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceGroupManagerStatefulPolicyPreservedState(c, f.PreservedState); err != nil {
		return nil, fmt.Errorf("error expanding PreservedState into preservedState: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["preservedState"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerStatefulPolicy flattens an instance of InstanceGroupManagerStatefulPolicy from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicy(c *Client, i interface{}) *InstanceGroupManagerStatefulPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatefulPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatefulPolicy
	}
	r.PreservedState = flattenInstanceGroupManagerStatefulPolicyPreservedState(c, m["preservedState"])

	return r
}

// expandInstanceGroupManagerStatefulPolicyPreservedStateMap expands the contents of InstanceGroupManagerStatefulPolicyPreservedState into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyPreservedStateMap(c *Client, f map[string]InstanceGroupManagerStatefulPolicyPreservedState) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatefulPolicyPreservedState(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatefulPolicyPreservedStateSlice expands the contents of InstanceGroupManagerStatefulPolicyPreservedState into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyPreservedStateSlice(c *Client, f []InstanceGroupManagerStatefulPolicyPreservedState) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatefulPolicyPreservedState(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateMap flattens the contents of InstanceGroupManagerStatefulPolicyPreservedState from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatefulPolicyPreservedState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatefulPolicyPreservedState{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatefulPolicyPreservedState{}
	}

	items := make(map[string]InstanceGroupManagerStatefulPolicyPreservedState)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatefulPolicyPreservedState(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateSlice flattens the contents of InstanceGroupManagerStatefulPolicyPreservedState from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateSlice(c *Client, i interface{}) []InstanceGroupManagerStatefulPolicyPreservedState {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatefulPolicyPreservedState{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatefulPolicyPreservedState{}
	}

	items := make([]InstanceGroupManagerStatefulPolicyPreservedState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatefulPolicyPreservedState(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatefulPolicyPreservedState expands an instance of InstanceGroupManagerStatefulPolicyPreservedState into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyPreservedState(c *Client, f *InstanceGroupManagerStatefulPolicyPreservedState) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceGroupManagerStatefulPolicyPreservedStateDisksMap(c, f.Disks); err != nil {
		return nil, fmt.Errorf("error expanding Disks into disks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["disks"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerStatefulPolicyPreservedState flattens an instance of InstanceGroupManagerStatefulPolicyPreservedState from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedState(c *Client, i interface{}) *InstanceGroupManagerStatefulPolicyPreservedState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatefulPolicyPreservedState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatefulPolicyPreservedState
	}
	r.Disks = flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksMap(c, m["disks"])

	return r
}

// expandInstanceGroupManagerStatefulPolicyPreservedStateDisksMap expands the contents of InstanceGroupManagerStatefulPolicyPreservedStateDisks into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyPreservedStateDisksMap(c *Client, f map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceGroupManagerStatefulPolicyPreservedStateDisks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceGroupManagerStatefulPolicyPreservedStateDisksSlice expands the contents of InstanceGroupManagerStatefulPolicyPreservedStateDisks into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyPreservedStateDisksSlice(c *Client, f []InstanceGroupManagerStatefulPolicyPreservedStateDisks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceGroupManagerStatefulPolicyPreservedStateDisks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksMap flattens the contents of InstanceGroupManagerStatefulPolicyPreservedStateDisks from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	}

	items := make(map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisks)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatefulPolicyPreservedStateDisks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksSlice flattens the contents of InstanceGroupManagerStatefulPolicyPreservedStateDisks from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksSlice(c *Client, i interface{}) []InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	}

	items := make([]InstanceGroupManagerStatefulPolicyPreservedStateDisks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatefulPolicyPreservedStateDisks(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceGroupManagerStatefulPolicyPreservedStateDisks expands an instance of InstanceGroupManagerStatefulPolicyPreservedStateDisks into a JSON
// request object.
func expandInstanceGroupManagerStatefulPolicyPreservedStateDisks(c *Client, f *InstanceGroupManagerStatefulPolicyPreservedStateDisks) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutoDelete; !dcl.IsEmptyValueIndirect(v) {
		m["autoDelete"] = v
	}

	return m, nil
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateDisks flattens an instance of InstanceGroupManagerStatefulPolicyPreservedStateDisks from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateDisks(c *Client, i interface{}) *InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceGroupManagerStatefulPolicyPreservedStateDisks{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceGroupManagerStatefulPolicyPreservedStateDisks
	}
	r.AutoDelete = flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(m["autoDelete"])

	return r
}

// flattenInstanceGroupManagerDistributionPolicyTargetShapeEnumMap flattens the contents of InstanceGroupManagerDistributionPolicyTargetShapeEnum from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyTargetShapeEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerDistributionPolicyTargetShapeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerDistributionPolicyTargetShapeEnum{}
	}

	items := make(map[string]InstanceGroupManagerDistributionPolicyTargetShapeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerDistributionPolicyTargetShapeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerDistributionPolicyTargetShapeEnumSlice flattens the contents of InstanceGroupManagerDistributionPolicyTargetShapeEnum from a JSON
// response object.
func flattenInstanceGroupManagerDistributionPolicyTargetShapeEnumSlice(c *Client, i interface{}) []InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerDistributionPolicyTargetShapeEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerDistributionPolicyTargetShapeEnum{}
	}

	items := make([]InstanceGroupManagerDistributionPolicyTargetShapeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerDistributionPolicyTargetShapeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerDistributionPolicyTargetShapeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerDistributionPolicyTargetShapeEnum with the same value as that string.
func flattenInstanceGroupManagerDistributionPolicyTargetShapeEnum(i interface{}) *InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerDistributionPolicyTargetShapeEnumRef("")
	}

	return InstanceGroupManagerDistributionPolicyTargetShapeEnumRef(s)
}

// flattenInstanceGroupManagerUpdatePolicyTypeEnumMap flattens the contents of InstanceGroupManagerUpdatePolicyTypeEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyTypeEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyTypeEnum{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyTypeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyTypeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyTypeEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyTypeEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyTypeEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyTypeEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyTypeEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyTypeEnum(i interface{}) *InstanceGroupManagerUpdatePolicyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyTypeEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyTypeEnumRef(s)
}

// flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumMap flattens the contents of InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(i interface{}) *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef(s)
}

// flattenInstanceGroupManagerUpdatePolicyMinimalActionEnumMap flattens the contents of InstanceGroupManagerUpdatePolicyMinimalActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMinimalActionEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyMinimalActionEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyMinimalActionEnum{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyMinimalActionEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMinimalActionEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyMinimalActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMinimalActionEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyMinimalActionEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyMinimalActionEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyMinimalActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyMinimalActionEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyMinimalActionEnum(i interface{}) *InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyMinimalActionEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyMinimalActionEnumRef(s)
}

// flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnumMap flattens the contents of InstanceGroupManagerUpdatePolicyReplacementMethodEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyReplacementMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyReplacementMethodEnum{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyReplacementMethodEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyReplacementMethodEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyReplacementMethodEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyReplacementMethodEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyReplacementMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyReplacementMethodEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyReplacementMethodEnum(i interface{}) *InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyReplacementMethodEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyReplacementMethodEnumRef(s)
}

// flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumMap flattens the contents of InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum{}
	}

	items := make(map[string]InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumSlice flattens the contents of InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumSlice(c *Client, i interface{}) []InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum{}
	}

	items := make([]InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum with the same value as that string.
func flattenInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(i interface{}) *InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumRef("")
	}

	return InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumRef(s)
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumMap flattens the contents of InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum{}
	}

	items := make(map[string]InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumSlice flattens the contents of InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum from a JSON
// response object.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumSlice(c *Client, i interface{}) []InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum{}
	}

	items := make([]InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum with the same value as that string.
func flattenInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(i interface{}) *InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumRef("")
	}

	return InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumRef(s)
}

// flattenInstanceGroupManagerFailoverActionEnumMap flattens the contents of InstanceGroupManagerFailoverActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerFailoverActionEnumMap(c *Client, i interface{}) map[string]InstanceGroupManagerFailoverActionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceGroupManagerFailoverActionEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceGroupManagerFailoverActionEnum{}
	}

	items := make(map[string]InstanceGroupManagerFailoverActionEnum)
	for k, item := range a {
		items[k] = *flattenInstanceGroupManagerFailoverActionEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceGroupManagerFailoverActionEnumSlice flattens the contents of InstanceGroupManagerFailoverActionEnum from a JSON
// response object.
func flattenInstanceGroupManagerFailoverActionEnumSlice(c *Client, i interface{}) []InstanceGroupManagerFailoverActionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceGroupManagerFailoverActionEnum{}
	}

	if len(a) == 0 {
		return []InstanceGroupManagerFailoverActionEnum{}
	}

	items := make([]InstanceGroupManagerFailoverActionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceGroupManagerFailoverActionEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceGroupManagerFailoverActionEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceGroupManagerFailoverActionEnum with the same value as that string.
func flattenInstanceGroupManagerFailoverActionEnum(i interface{}) *InstanceGroupManagerFailoverActionEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceGroupManagerFailoverActionEnumRef("")
	}

	return InstanceGroupManagerFailoverActionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *InstanceGroupManager) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstanceGroupManager(b, c)
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

type instanceGroupManagerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceGroupManagerApiOperation
}

func convertFieldDiffsToInstanceGroupManagerDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]instanceGroupManagerDiff, error) {
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
	var diffs []instanceGroupManagerDiff
	// For each operation name, create a instanceGroupManagerDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := instanceGroupManagerDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToInstanceGroupManagerApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToInstanceGroupManagerApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (instanceGroupManagerApiOperation, error) {
	switch opName {

	case "updateInstanceGroupManagerPatchOperation":
		return &updateInstanceGroupManagerPatchOperation{FieldDiffs: fieldDiffs}, nil

	case "updateInstanceGroupManagerSetInstanceTemplateOperation":
		return &updateInstanceGroupManagerSetInstanceTemplateOperation{FieldDiffs: fieldDiffs}, nil

	case "updateInstanceGroupManagerSetTargetPoolsOperation":
		return &updateInstanceGroupManagerSetTargetPoolsOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractInstanceGroupManagerFields(r *InstanceGroupManager) error {
	return nil
}
