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
package container

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *NodePool) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"NodeCount", "Autoscaling"}, r.NodeCount, r.Autoscaling); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Cluster, "Cluster"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Config) {
		if err := r.Config.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Autoscaling) {
		if err := r.Autoscaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Management) {
		if err := r.Management.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaxPodsConstraint) {
		if err := r.MaxPodsConstraint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UpgradeSettings) {
		if err := r.UpgradeSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NodePoolConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SandboxConfig) {
		if err := r.SandboxConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReservationAffinity) {
		if err := r.ReservationAffinity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ShieldedInstanceConfig) {
		if err := r.ShieldedInstanceConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NodePoolConfigAccelerators) validate() error {
	return nil
}
func (r *NodePoolConfigTaints) validate() error {
	return nil
}
func (r *NodePoolConfigSandboxConfig) validate() error {
	return nil
}
func (r *NodePoolConfigReservationAffinity) validate() error {
	return nil
}
func (r *NodePoolConfigShieldedInstanceConfig) validate() error {
	return nil
}
func (r *NodePoolAutoscaling) validate() error {
	return nil
}
func (r *NodePoolManagement) validate() error {
	if !dcl.IsEmptyValueIndirect(r.UpgradeOptions) {
		if err := r.UpgradeOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NodePoolManagementUpgradeOptions) validate() error {
	return nil
}
func (r *NodePoolMaxPodsConstraint) validate() error {
	return nil
}
func (r *NodePoolConditions) validate() error {
	return nil
}
func (r *NodePoolUpgradeSettings) validate() error {
	return nil
}

func nodePoolGetURL(userBasePath string, r *NodePool) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"cluster":  dcl.ValueOrEmptyString(r.Cluster),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}", "https://container.googleapis.com/v1/", userBasePath, params), nil
}

func nodePoolListURL(userBasePath, project, location, cluster string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"cluster":  cluster,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools", "https://container.googleapis.com/v1/", userBasePath, params), nil

}

func nodePoolCreateURL(userBasePath, project, location, cluster string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"cluster":  cluster,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools", "https://container.googleapis.com/v1/", userBasePath, params), nil

}

func nodePoolDeleteURL(userBasePath string, r *NodePool) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"cluster":  dcl.ValueOrEmptyString(r.Cluster),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}", "https://container.googleapis.com/v1/", userBasePath, params), nil
}

// nodePoolApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type nodePoolApiOperation interface {
	do(context.Context, *NodePool, *Client) error
}

// newUpdateNodePoolSetAutoscalingRequest creates a request for an
// NodePool resource's setAutoscaling update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNodePoolSetAutoscalingRequest(ctx context.Context, f *NodePool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandNodePoolAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["autoscaling"] = v
	}
	return req, nil
}

// marshalUpdateNodePoolSetAutoscalingRequest converts the update into
// the final JSON request body.
func marshalUpdateNodePoolSetAutoscalingRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNodePoolSetAutoscalingOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNodePoolSetAutoscalingOperation) do(ctx context.Context, r *NodePool, c *Client) error {
	_, err := c.GetNodePool(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setAutoscaling")
	if err != nil {
		return err
	}

	req, err := newUpdateNodePoolSetAutoscalingRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNodePoolSetAutoscalingRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateNodePoolSetManagementRequest creates a request for an
// NodePool resource's setManagement update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNodePoolSetManagementRequest(ctx context.Context, f *NodePool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandNodePoolManagement(c, f.Management); err != nil {
		return nil, fmt.Errorf("error expanding Management into management: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["management"] = v
	}
	return req, nil
}

// marshalUpdateNodePoolSetManagementRequest converts the update into
// the final JSON request body.
func marshalUpdateNodePoolSetManagementRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNodePoolSetManagementOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNodePoolSetManagementOperation) do(ctx context.Context, r *NodePool, c *Client) error {
	_, err := c.GetNodePool(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setManagement")
	if err != nil {
		return err
	}

	req, err := newUpdateNodePoolSetManagementRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNodePoolSetManagementRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateNodePoolSetSizeRequest creates a request for an
// NodePool resource's setSize update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNodePoolSetSizeRequest(ctx context.Context, f *NodePool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		req["nodeCount"] = v
	}
	return req, nil
}

// marshalUpdateNodePoolSetSizeRequest converts the update into
// the final JSON request body.
func marshalUpdateNodePoolSetSizeRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNodePoolSetSizeOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNodePoolSetSizeOperation) do(ctx context.Context, r *NodePool, c *Client) error {
	_, err := c.GetNodePool(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setSize")
	if err != nil {
		return err
	}

	req, err := newUpdateNodePoolSetSizeRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNodePoolSetSizeRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.ComputeOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

// newUpdateNodePoolUpdateRequest creates a request for an
// NodePool resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNodePoolUpdateRequest(ctx context.Context, f *NodePool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Locations; !dcl.IsEmptyValueIndirect(v) {
		req["locations"] = v
	}
	return req, nil
}

// marshalUpdateNodePoolUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateNodePoolUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNodePoolUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNodePoolUpdateOperation) do(ctx context.Context, r *NodePool, c *Client) error {
	_, err := c.GetNodePool(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateNodePoolUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNodePoolUpdateRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	return nil
}

func (c *Client) listNodePoolRaw(ctx context.Context, project, location, cluster, pageToken string) ([]byte, error) {
	u, err := nodePoolListURL(c.Config.BasePath, project, location, cluster)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
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

type listNodePoolOperation struct {
	NodePools []map[string]interface{} `json:"nodePools"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listNodePool(ctx context.Context, project, location, cluster, pageToken string) ([]*NodePool, string, error) {
	b, err := c.listNodePoolRaw(ctx, project, location, cluster, pageToken)
	if err != nil {
		return nil, "", err
	}

	var m listNodePoolOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*NodePool
	for _, v := range m.NodePools {
		res := flattenNodePool(c, v)
		res.Project = &project
		res.Location = &location
		res.Cluster = &cluster
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNodePool(ctx context.Context, f func(*NodePool) bool, resources []*NodePool) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNodePool(ctx, res)
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

type deleteNodePoolOperation struct{}

func (op *deleteNodePoolOperation) do(ctx context.Context, r *NodePool, c *Client) error {

	_, err := c.GetNodePool(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("NodePool not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetNodePool checking for existence. error: %v", err)
		return err
	}

	if err := nodePoolWaitForRestingState(ctx, r, c); err != nil {
		return err
	}

	u, err := nodePoolDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetNodePool(ctx, r.urlNormalized())
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
type createNodePoolOperation struct {
	response map[string]interface{}
}

func (op *createNodePoolOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNodePoolOperation) do(ctx context.Context, r *NodePool, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, cluster := r.createFields()
	u, err := nodePoolCreateURL(c.Config.BasePath, project, location, cluster)

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
	if err := o.Wait(ctx, c.Config, "https://container.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetNodePool(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNodePoolRaw(ctx context.Context, r *NodePool) ([]byte, error) {

	u, err := nodePoolGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) nodePoolDiffsForRawDesired(ctx context.Context, rawDesired *NodePool, opts ...dcl.ApplyOption) (initial, desired *NodePool, diffs []nodePoolDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *NodePool
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*NodePool); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected NodePool, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNodePool(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a NodePool resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve NodePool resource: %v", err)
		}
		c.Config.Logger.Info("Found that NodePool resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNodePoolDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for NodePool: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for NodePool: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNodePoolInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for NodePool: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNodePoolDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for NodePool: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNodePool(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func nodePoolCheckForErrorStateAndDelete(ctx context.Context, r *NodePool, c *Client) error {
	if dcl.FindStringInArray(*r.Status, []string{"STOPPING", "ERROR", "DEGRADED", "RUNNING_WITH_ERROR"}) {
		err := c.DeleteNodePool(ctx, r.urlNormalized())
		if err != nil {
			return fmt.Errorf("NodePool was is in error state %s, delete attempted with error: %v", *r.StatusMessage, err)
		} else {
			return fmt.Errorf("NodePool was is in error state %s, delete succeeded.", *r.StatusMessage)
		}
	}
	return nil
}

type nodePoolWaitOperation struct {
	Client   *Client
	Resource *NodePool
}

func (op *nodePoolWaitOperation) operate(ctx context.Context) (*dcl.RetryDetails, error) {
	current, err := op.Client.GetNodePool(ctx, op.Resource.urlNormalized())
	if err != nil {
		return nil, err
	}
	if dcl.FindStringInArray(*current.Status, []string{"PROVISIONING", "RECONCILING"}) {
		return nil, dcl.OperationNotDone{}
	}
	return &dcl.RetryDetails{}, nil
}

func nodePoolWaitForRestingState(ctx context.Context, r *NodePool, c *Client) error {
	op := nodePoolWaitOperation{
		Client:   c,
		Resource: r,
	}
	return dcl.Do(ctx, op.operate, c.Config.RetryProvider)
}

func canonicalizeNodePoolInitialState(rawInitial, rawDesired *NodePool) (*NodePool, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.NodeCount) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.Autoscaling) {
			rawInitial.NodeCount = dcl.Int64(0)
		}
	}

	if dcl.IsZeroValue(rawInitial.Autoscaling) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.NodeCount) {
			rawInitial.Autoscaling = EmptyNodePoolAutoscaling
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

func canonicalizeNodePoolDesiredState(rawDesired, rawInitial *NodePool, opts ...dcl.ApplyOption) (*NodePool, error) {

	if dcl.IsZeroValue(rawDesired.NodeCount) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.Autoscaling) {
			rawDesired.NodeCount = dcl.Int64(0)
		}
	}

	if dcl.IsZeroValue(rawDesired.Autoscaling) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.NodeCount) {
			rawDesired.Autoscaling = EmptyNodePoolAutoscaling
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeNodePoolConfig(rawDesired.Config, nil, opts...)
		rawDesired.Autoscaling = canonicalizeNodePoolAutoscaling(rawDesired.Autoscaling, nil, opts...)
		rawDesired.Management = canonicalizeNodePoolManagement(rawDesired.Management, nil, opts...)
		rawDesired.MaxPodsConstraint = canonicalizeNodePoolMaxPodsConstraint(rawDesired.MaxPodsConstraint, nil, opts...)
		rawDesired.UpgradeSettings = canonicalizeNodePoolUpgradeSettings(rawDesired.UpgradeSettings, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	rawDesired.Config = canonicalizeNodePoolConfig(rawDesired.Config, rawInitial.Config, opts...)
	if dcl.IsZeroValue(rawDesired.NodeCount) {
		rawDesired.NodeCount = rawInitial.NodeCount
	}
	if dcl.StringCanonicalize(rawDesired.Version, rawInitial.Version) {
		rawDesired.Version = rawInitial.Version
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.StringCanonicalize(rawDesired.StatusMessage, rawInitial.StatusMessage) {
		rawDesired.StatusMessage = rawInitial.StatusMessage
	}
	if dcl.IsZeroValue(rawDesired.Locations) {
		rawDesired.Locations = rawInitial.Locations
	}
	rawDesired.Autoscaling = canonicalizeNodePoolAutoscaling(rawDesired.Autoscaling, rawInitial.Autoscaling, opts...)
	rawDesired.Management = canonicalizeNodePoolManagement(rawDesired.Management, rawInitial.Management, opts...)
	rawDesired.MaxPodsConstraint = canonicalizeNodePoolMaxPodsConstraint(rawDesired.MaxPodsConstraint, rawInitial.MaxPodsConstraint, opts...)
	if dcl.IsZeroValue(rawDesired.Conditions) {
		rawDesired.Conditions = rawInitial.Conditions
	}
	if dcl.IsZeroValue(rawDesired.PodIPv4CidrSize) {
		rawDesired.PodIPv4CidrSize = rawInitial.PodIPv4CidrSize
	}
	rawDesired.UpgradeSettings = canonicalizeNodePoolUpgradeSettings(rawDesired.UpgradeSettings, rawInitial.UpgradeSettings, opts...)
	if dcl.NameToSelfLink(rawDesired.Cluster, rawInitial.Cluster) {
		rawDesired.Cluster = rawInitial.Cluster
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeNodePoolNewState(c *Client, rawNew, rawDesired *NodePool) (*NodePool, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		rawNew.Config = canonicalizeNewNodePoolConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NodeCount) && dcl.IsEmptyValueIndirect(rawDesired.NodeCount) {
		rawNew.NodeCount = rawDesired.NodeCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Version) && dcl.IsEmptyValueIndirect(rawDesired.Version) {
		rawNew.Version = rawDesired.Version
	} else {
		if dcl.StringCanonicalize(rawDesired.Version, rawNew.Version) {
			rawNew.Version = rawDesired.Version
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusMessage) && dcl.IsEmptyValueIndirect(rawDesired.StatusMessage) {
		rawNew.StatusMessage = rawDesired.StatusMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StatusMessage, rawNew.StatusMessage) {
			rawNew.StatusMessage = rawDesired.StatusMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Locations) && dcl.IsEmptyValueIndirect(rawDesired.Locations) {
		rawNew.Locations = rawDesired.Locations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Autoscaling) && dcl.IsEmptyValueIndirect(rawDesired.Autoscaling) {
		rawNew.Autoscaling = rawDesired.Autoscaling
	} else {
		rawNew.Autoscaling = canonicalizeNewNodePoolAutoscaling(c, rawDesired.Autoscaling, rawNew.Autoscaling)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Management) && dcl.IsEmptyValueIndirect(rawDesired.Management) {
		rawNew.Management = rawDesired.Management
	} else {
		rawNew.Management = canonicalizeNewNodePoolManagement(c, rawDesired.Management, rawNew.Management)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxPodsConstraint) && dcl.IsEmptyValueIndirect(rawDesired.MaxPodsConstraint) {
		rawNew.MaxPodsConstraint = rawDesired.MaxPodsConstraint
	} else {
		rawNew.MaxPodsConstraint = canonicalizeNewNodePoolMaxPodsConstraint(c, rawDesired.MaxPodsConstraint, rawNew.MaxPodsConstraint)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Conditions) && dcl.IsEmptyValueIndirect(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
		rawNew.Conditions = canonicalizeNewNodePoolConditionsSlice(c, rawDesired.Conditions, rawNew.Conditions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PodIPv4CidrSize) && dcl.IsEmptyValueIndirect(rawDesired.PodIPv4CidrSize) {
		rawNew.PodIPv4CidrSize = rawDesired.PodIPv4CidrSize
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpgradeSettings) && dcl.IsEmptyValueIndirect(rawDesired.UpgradeSettings) {
		rawNew.UpgradeSettings = rawDesired.UpgradeSettings
	} else {
		rawNew.UpgradeSettings = canonicalizeNewNodePoolUpgradeSettings(c, rawDesired.UpgradeSettings, rawNew.UpgradeSettings)
	}

	rawNew.Cluster = rawDesired.Cluster

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeNodePoolConfig(des, initial *NodePoolConfig, opts ...dcl.ApplyOption) *NodePoolConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		des.MachineType = initial.MachineType
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	}
	if dcl.IsZeroValue(des.OAuthScopes) {
		des.OAuthScopes = initial.OAuthScopes
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		des.ServiceAccount = initial.ServiceAccount
	}
	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = initial.Metadata
	}
	if dcl.StringCanonicalize(des.ImageType, initial.ImageType) || dcl.IsZeroValue(des.ImageType) {
		des.ImageType = initial.ImageType
	}
	if dcl.IsZeroValue(des.Labels) {
		des.Labels = initial.Labels
	}
	if dcl.IsZeroValue(des.LocalSsdCount) {
		des.LocalSsdCount = initial.LocalSsdCount
	}
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	}
	if dcl.BoolCanonicalize(des.Preemptible, initial.Preemptible) || dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}
	if dcl.IsZeroValue(des.Accelerators) {
		des.Accelerators = initial.Accelerators
	}
	if dcl.StringCanonicalize(des.DiskType, initial.DiskType) || dcl.IsZeroValue(des.DiskType) {
		des.DiskType = initial.DiskType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, initial.MinCpuPlatform) || dcl.IsZeroValue(des.MinCpuPlatform) {
		des.MinCpuPlatform = initial.MinCpuPlatform
	}
	if dcl.IsZeroValue(des.Taints) {
		des.Taints = initial.Taints
	}
	des.SandboxConfig = canonicalizeNodePoolConfigSandboxConfig(des.SandboxConfig, initial.SandboxConfig, opts...)
	des.ReservationAffinity = canonicalizeNodePoolConfigReservationAffinity(des.ReservationAffinity, initial.ReservationAffinity, opts...)
	des.ShieldedInstanceConfig = canonicalizeNodePoolConfigShieldedInstanceConfig(des.ShieldedInstanceConfig, initial.ShieldedInstanceConfig, opts...)

	return des
}

func canonicalizeNewNodePoolConfig(c *Client, des, nw *NodePoolConfig) *NodePoolConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringCanonicalize(des.ImageType, nw.ImageType) {
		nw.ImageType = des.ImageType
	}
	if dcl.BoolCanonicalize(des.Preemptible, nw.Preemptible) {
		nw.Preemptible = des.Preemptible
	}
	nw.Accelerators = canonicalizeNewNodePoolConfigAcceleratorsSlice(c, des.Accelerators, nw.Accelerators)
	if dcl.StringCanonicalize(des.DiskType, nw.DiskType) {
		nw.DiskType = des.DiskType
	}
	if dcl.StringCanonicalize(des.MinCpuPlatform, nw.MinCpuPlatform) {
		nw.MinCpuPlatform = des.MinCpuPlatform
	}
	nw.Taints = canonicalizeNewNodePoolConfigTaintsSlice(c, des.Taints, nw.Taints)
	nw.SandboxConfig = canonicalizeNewNodePoolConfigSandboxConfig(c, des.SandboxConfig, nw.SandboxConfig)
	nw.ReservationAffinity = canonicalizeNewNodePoolConfigReservationAffinity(c, des.ReservationAffinity, nw.ReservationAffinity)
	nw.ShieldedInstanceConfig = canonicalizeNewNodePoolConfigShieldedInstanceConfig(c, des.ShieldedInstanceConfig, nw.ShieldedInstanceConfig)

	return nw
}

func canonicalizeNewNodePoolConfigSet(c *Client, des, nw []NodePoolConfig) []NodePoolConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConfig(c, &d, &n) {
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

func canonicalizeNewNodePoolConfigSlice(c *Client, des, nw []NodePoolConfig) []NodePoolConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConfig(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolConfigAccelerators(des, initial *NodePoolConfigAccelerators, opts ...dcl.ApplyOption) *NodePoolConfigAccelerators {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AcceleratorCount) {
		des.AcceleratorCount = initial.AcceleratorCount
	}
	if dcl.StringCanonicalize(des.AcceleratorType, initial.AcceleratorType) || dcl.IsZeroValue(des.AcceleratorType) {
		des.AcceleratorType = initial.AcceleratorType
	}

	return des
}

func canonicalizeNewNodePoolConfigAccelerators(c *Client, des, nw *NodePoolConfigAccelerators) *NodePoolConfigAccelerators {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.AcceleratorType, nw.AcceleratorType) {
		nw.AcceleratorType = des.AcceleratorType
	}

	return nw
}

func canonicalizeNewNodePoolConfigAcceleratorsSet(c *Client, des, nw []NodePoolConfigAccelerators) []NodePoolConfigAccelerators {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConfigAccelerators
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConfigAccelerators(c, &d, &n) {
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

func canonicalizeNewNodePoolConfigAcceleratorsSlice(c *Client, des, nw []NodePoolConfigAccelerators) []NodePoolConfigAccelerators {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConfigAccelerators
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConfigAccelerators(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolConfigTaints(des, initial *NodePoolConfigTaints, opts ...dcl.ApplyOption) *NodePoolConfigTaints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}
	if dcl.StringCanonicalize(des.Effect, initial.Effect) || dcl.IsZeroValue(des.Effect) {
		des.Effect = initial.Effect
	}

	return des
}

func canonicalizeNewNodePoolConfigTaints(c *Client, des, nw *NodePoolConfigTaints) *NodePoolConfigTaints {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	if dcl.StringCanonicalize(des.Effect, nw.Effect) {
		nw.Effect = des.Effect
	}

	return nw
}

func canonicalizeNewNodePoolConfigTaintsSet(c *Client, des, nw []NodePoolConfigTaints) []NodePoolConfigTaints {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConfigTaints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConfigTaints(c, &d, &n) {
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

func canonicalizeNewNodePoolConfigTaintsSlice(c *Client, des, nw []NodePoolConfigTaints) []NodePoolConfigTaints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConfigTaints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConfigTaints(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolConfigSandboxConfig(des, initial *NodePoolConfigSandboxConfig, opts ...dcl.ApplyOption) *NodePoolConfigSandboxConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewNodePoolConfigSandboxConfig(c *Client, des, nw *NodePoolConfigSandboxConfig) *NodePoolConfigSandboxConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNodePoolConfigSandboxConfigSet(c *Client, des, nw []NodePoolConfigSandboxConfig) []NodePoolConfigSandboxConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConfigSandboxConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConfigSandboxConfig(c, &d, &n) {
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

func canonicalizeNewNodePoolConfigSandboxConfigSlice(c *Client, des, nw []NodePoolConfigSandboxConfig) []NodePoolConfigSandboxConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConfigSandboxConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConfigSandboxConfig(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolConfigReservationAffinity(des, initial *NodePoolConfigReservationAffinity, opts ...dcl.ApplyOption) *NodePoolConfigReservationAffinity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ConsumeReservationType) {
		des.ConsumeReservationType = initial.ConsumeReservationType
	}
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Values) {
		des.Values = initial.Values
	}

	return des
}

func canonicalizeNewNodePoolConfigReservationAffinity(c *Client, des, nw *NodePoolConfigReservationAffinity) *NodePoolConfigReservationAffinity {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}

	return nw
}

func canonicalizeNewNodePoolConfigReservationAffinitySet(c *Client, des, nw []NodePoolConfigReservationAffinity) []NodePoolConfigReservationAffinity {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConfigReservationAffinity
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConfigReservationAffinity(c, &d, &n) {
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

func canonicalizeNewNodePoolConfigReservationAffinitySlice(c *Client, des, nw []NodePoolConfigReservationAffinity) []NodePoolConfigReservationAffinity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConfigReservationAffinity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConfigReservationAffinity(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolConfigShieldedInstanceConfig(des, initial *NodePoolConfigShieldedInstanceConfig, opts ...dcl.ApplyOption) *NodePoolConfigShieldedInstanceConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, initial.EnableSecureBoot) || dcl.IsZeroValue(des.EnableSecureBoot) {
		des.EnableSecureBoot = initial.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, initial.EnableIntegrityMonitoring) || dcl.IsZeroValue(des.EnableIntegrityMonitoring) {
		des.EnableIntegrityMonitoring = initial.EnableIntegrityMonitoring
	}

	return des
}

func canonicalizeNewNodePoolConfigShieldedInstanceConfig(c *Client, des, nw *NodePoolConfigShieldedInstanceConfig) *NodePoolConfigShieldedInstanceConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.EnableSecureBoot, nw.EnableSecureBoot) {
		nw.EnableSecureBoot = des.EnableSecureBoot
	}
	if dcl.BoolCanonicalize(des.EnableIntegrityMonitoring, nw.EnableIntegrityMonitoring) {
		nw.EnableIntegrityMonitoring = des.EnableIntegrityMonitoring
	}

	return nw
}

func canonicalizeNewNodePoolConfigShieldedInstanceConfigSet(c *Client, des, nw []NodePoolConfigShieldedInstanceConfig) []NodePoolConfigShieldedInstanceConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConfigShieldedInstanceConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConfigShieldedInstanceConfig(c, &d, &n) {
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

func canonicalizeNewNodePoolConfigShieldedInstanceConfigSlice(c *Client, des, nw []NodePoolConfigShieldedInstanceConfig) []NodePoolConfigShieldedInstanceConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConfigShieldedInstanceConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConfigShieldedInstanceConfig(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolAutoscaling(des, initial *NodePoolAutoscaling, opts ...dcl.ApplyOption) *NodePoolAutoscaling {
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
	if dcl.IsZeroValue(des.MinNodeCount) {
		des.MinNodeCount = initial.MinNodeCount
	}
	if dcl.IsZeroValue(des.MaxNodeCount) {
		des.MaxNodeCount = initial.MaxNodeCount
	}
	if dcl.BoolCanonicalize(des.Autoprovisioned, initial.Autoprovisioned) || dcl.IsZeroValue(des.Autoprovisioned) {
		des.Autoprovisioned = initial.Autoprovisioned
	}

	return des
}

func canonicalizeNewNodePoolAutoscaling(c *Client, des, nw *NodePoolAutoscaling) *NodePoolAutoscaling {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.Autoprovisioned, nw.Autoprovisioned) {
		nw.Autoprovisioned = des.Autoprovisioned
	}

	return nw
}

func canonicalizeNewNodePoolAutoscalingSet(c *Client, des, nw []NodePoolAutoscaling) []NodePoolAutoscaling {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolAutoscaling
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolAutoscaling(c, &d, &n) {
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

func canonicalizeNewNodePoolAutoscalingSlice(c *Client, des, nw []NodePoolAutoscaling) []NodePoolAutoscaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolAutoscaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolAutoscaling(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolManagement(des, initial *NodePoolManagement, opts ...dcl.ApplyOption) *NodePoolManagement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.AutoUpgrade, initial.AutoUpgrade) || dcl.IsZeroValue(des.AutoUpgrade) {
		des.AutoUpgrade = initial.AutoUpgrade
	}
	if dcl.BoolCanonicalize(des.AutoRepair, initial.AutoRepair) || dcl.IsZeroValue(des.AutoRepair) {
		des.AutoRepair = initial.AutoRepair
	}
	des.UpgradeOptions = canonicalizeNodePoolManagementUpgradeOptions(des.UpgradeOptions, initial.UpgradeOptions, opts...)

	return des
}

func canonicalizeNewNodePoolManagement(c *Client, des, nw *NodePoolManagement) *NodePoolManagement {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.AutoUpgrade, nw.AutoUpgrade) {
		nw.AutoUpgrade = des.AutoUpgrade
	}
	if dcl.BoolCanonicalize(des.AutoRepair, nw.AutoRepair) {
		nw.AutoRepair = des.AutoRepair
	}
	nw.UpgradeOptions = canonicalizeNewNodePoolManagementUpgradeOptions(c, des.UpgradeOptions, nw.UpgradeOptions)

	return nw
}

func canonicalizeNewNodePoolManagementSet(c *Client, des, nw []NodePoolManagement) []NodePoolManagement {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolManagement
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolManagement(c, &d, &n) {
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

func canonicalizeNewNodePoolManagementSlice(c *Client, des, nw []NodePoolManagement) []NodePoolManagement {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolManagement
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolManagement(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolManagementUpgradeOptions(des, initial *NodePoolManagementUpgradeOptions, opts ...dcl.ApplyOption) *NodePoolManagementUpgradeOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AutoUpgradeStartTime) {
		des.AutoUpgradeStartTime = initial.AutoUpgradeStartTime
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewNodePoolManagementUpgradeOptions(c *Client, des, nw *NodePoolManagementUpgradeOptions) *NodePoolManagementUpgradeOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewNodePoolManagementUpgradeOptionsSet(c *Client, des, nw []NodePoolManagementUpgradeOptions) []NodePoolManagementUpgradeOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolManagementUpgradeOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolManagementUpgradeOptions(c, &d, &n) {
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

func canonicalizeNewNodePoolManagementUpgradeOptionsSlice(c *Client, des, nw []NodePoolManagementUpgradeOptions) []NodePoolManagementUpgradeOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolManagementUpgradeOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolManagementUpgradeOptions(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolMaxPodsConstraint(des, initial *NodePoolMaxPodsConstraint, opts ...dcl.ApplyOption) *NodePoolMaxPodsConstraint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxPodsPerNode) {
		des.MaxPodsPerNode = initial.MaxPodsPerNode
	}

	return des
}

func canonicalizeNewNodePoolMaxPodsConstraint(c *Client, des, nw *NodePoolMaxPodsConstraint) *NodePoolMaxPodsConstraint {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNodePoolMaxPodsConstraintSet(c *Client, des, nw []NodePoolMaxPodsConstraint) []NodePoolMaxPodsConstraint {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolMaxPodsConstraint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolMaxPodsConstraint(c, &d, &n) {
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

func canonicalizeNewNodePoolMaxPodsConstraintSlice(c *Client, des, nw []NodePoolMaxPodsConstraint) []NodePoolMaxPodsConstraint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolMaxPodsConstraint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolMaxPodsConstraint(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolConditions(des, initial *NodePoolConditions, opts ...dcl.ApplyOption) *NodePoolConditions {
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
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}

	return des
}

func canonicalizeNewNodePoolConditions(c *Client, des, nw *NodePoolConditions) *NodePoolConditions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewNodePoolConditionsSet(c *Client, des, nw []NodePoolConditions) []NodePoolConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolConditions(c, &d, &n) {
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

func canonicalizeNewNodePoolConditionsSlice(c *Client, des, nw []NodePoolConditions) []NodePoolConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolConditions(c, &d, &n))
	}

	return items
}

func canonicalizeNodePoolUpgradeSettings(des, initial *NodePoolUpgradeSettings, opts ...dcl.ApplyOption) *NodePoolUpgradeSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.MaxSurge) {
		des.MaxSurge = initial.MaxSurge
	}
	if dcl.IsZeroValue(des.MaxUnavailable) {
		des.MaxUnavailable = initial.MaxUnavailable
	}

	return des
}

func canonicalizeNewNodePoolUpgradeSettings(c *Client, des, nw *NodePoolUpgradeSettings) *NodePoolUpgradeSettings {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNodePoolUpgradeSettingsSet(c *Client, des, nw []NodePoolUpgradeSettings) []NodePoolUpgradeSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []NodePoolUpgradeSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNodePoolUpgradeSettings(c, &d, &n) {
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

func canonicalizeNewNodePoolUpgradeSettingsSlice(c *Client, des, nw []NodePoolUpgradeSettings) []NodePoolUpgradeSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NodePoolUpgradeSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodePoolUpgradeSettings(c, &d, &n))
	}

	return items
}

type nodePoolDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         nodePoolApiOperation
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
func diffNodePool(c *Client, desired, actual *NodePool, opts ...dcl.ApplyOption) ([]nodePoolDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []nodePoolDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, nodePoolDiff{RequiresRecreate: true, FieldName: "Name"})
	}

	if d, err := dcl.Diff(desired.NodeCount, actual.NodeCount, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, nodePoolDiff{
			UpdateOp: &updateNodePoolSetSizeOperation{}, FieldName: "NodeCount",
		})
	}

	if d, err := dcl.Diff(desired.Version, actual.Version, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, nodePoolDiff{RequiresRecreate: true, FieldName: "Version"})
	}

	if d, err := dcl.Diff(desired.Status, actual.Status, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, nodePoolDiff{RequiresRecreate: true, FieldName: "Status"})
	}

	if d, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, nodePoolDiff{RequiresRecreate: true, FieldName: "StatusMessage"})
	}

	if d, err := dcl.Diff(desired.PodIPv4CidrSize, actual.PodIPv4CidrSize, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, nodePoolDiff{RequiresRecreate: true, FieldName: "PodIPv4CidrSize"})
	}

	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, nodePoolDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if compareNodePoolConfig(c, desired.Config, actual.Config) {
		c.Config.Logger.Infof("Detected diff in Config.\nDESIRED: %v\nACTUAL: %v", desired.Config, actual.Config)
		diffs = append(diffs, nodePoolDiff{
			RequiresRecreate: true,
			FieldName:        "Config",
		})
	}
	if !reflect.DeepEqual(desired.NodeCount, actual.NodeCount) {
		c.Config.Logger.Infof("Detected diff in NodeCount.\nDESIRED: %v\nACTUAL: %v", desired.NodeCount, actual.NodeCount)

		diffs = append(diffs, nodePoolDiff{
			UpdateOp:  &updateNodePoolSetSizeOperation{},
			FieldName: "NodeCount",
		})

	}
	if !dcl.IsZeroValue(desired.Version) && !dcl.StringCanonicalize(desired.Version, actual.Version) {
		c.Config.Logger.Infof("Detected diff in Version.\nDESIRED: %v\nACTUAL: %v", desired.Version, actual.Version)
		diffs = append(diffs, nodePoolDiff{
			RequiresRecreate: true,
			FieldName:        "Version",
		})
	}
	if !dcl.StringSliceEquals(desired.Locations, actual.Locations) {
		c.Config.Logger.Infof("Detected diff in Locations.\nDESIRED: %v\nACTUAL: %v", desired.Locations, actual.Locations)

		diffs = append(diffs, nodePoolDiff{
			UpdateOp:  &updateNodePoolUpdateOperation{},
			FieldName: "Locations",
		})

	}
	if compareNodePoolAutoscaling(c, desired.Autoscaling, actual.Autoscaling) {
		c.Config.Logger.Infof("Detected diff in Autoscaling.\nDESIRED: %v\nACTUAL: %v", desired.Autoscaling, actual.Autoscaling)

		diffs = append(diffs, nodePoolDiff{
			UpdateOp:  &updateNodePoolSetAutoscalingOperation{},
			FieldName: "Autoscaling",
		})

	}
	if compareNodePoolManagement(c, desired.Management, actual.Management) {
		c.Config.Logger.Infof("Detected diff in Management.\nDESIRED: %v\nACTUAL: %v", desired.Management, actual.Management)

		diffs = append(diffs, nodePoolDiff{
			UpdateOp:  &updateNodePoolSetManagementOperation{},
			FieldName: "Management",
		})

	}
	if compareNodePoolMaxPodsConstraint(c, desired.MaxPodsConstraint, actual.MaxPodsConstraint) {
		c.Config.Logger.Infof("Detected diff in MaxPodsConstraint.\nDESIRED: %v\nACTUAL: %v", desired.MaxPodsConstraint, actual.MaxPodsConstraint)
		diffs = append(diffs, nodePoolDiff{
			RequiresRecreate: true,
			FieldName:        "MaxPodsConstraint",
		})
	}
	if compareNodePoolUpgradeSettings(c, desired.UpgradeSettings, actual.UpgradeSettings) {
		c.Config.Logger.Infof("Detected diff in UpgradeSettings.\nDESIRED: %v\nACTUAL: %v", desired.UpgradeSettings, actual.UpgradeSettings)
		diffs = append(diffs, nodePoolDiff{
			RequiresRecreate: true,
			FieldName:        "UpgradeSettings",
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
	var deduped []nodePoolDiff
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
func compareNodePoolConfig(c *Client, desired, actual *NodePoolConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MachineType == nil && desired.MachineType != nil && !dcl.IsEmptyValueIndirect(desired.MachineType) {
		c.Config.Logger.Infof("desired MachineType %s - but actually nil", dcl.SprintResource(desired.MachineType))
		return true
	}
	if !dcl.StringCanonicalize(desired.MachineType, actual.MachineType) && !dcl.IsZeroValue(desired.MachineType) {
		c.Config.Logger.Infof("Diff in MachineType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MachineType), dcl.SprintResource(actual.MachineType))
		return true
	}
	if actual.DiskSizeGb == nil && desired.DiskSizeGb != nil && !dcl.IsEmptyValueIndirect(desired.DiskSizeGb) {
		c.Config.Logger.Infof("desired DiskSizeGb %s - but actually nil", dcl.SprintResource(desired.DiskSizeGb))
		return true
	}
	if !reflect.DeepEqual(desired.DiskSizeGb, actual.DiskSizeGb) && !dcl.IsZeroValue(desired.DiskSizeGb) {
		c.Config.Logger.Infof("Diff in DiskSizeGb. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskSizeGb), dcl.SprintResource(actual.DiskSizeGb))
		return true
	}
	if actual.OAuthScopes == nil && desired.OAuthScopes != nil && !dcl.IsEmptyValueIndirect(desired.OAuthScopes) {
		c.Config.Logger.Infof("desired OAuthScopes %s - but actually nil", dcl.SprintResource(desired.OAuthScopes))
		return true
	}
	if !dcl.StringSliceEquals(desired.OAuthScopes, actual.OAuthScopes) && !dcl.IsZeroValue(desired.OAuthScopes) {
		c.Config.Logger.Infof("Diff in OAuthScopes. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OAuthScopes), dcl.SprintResource(actual.OAuthScopes))
		return true
	}
	if actual.ServiceAccount == nil && desired.ServiceAccount != nil && !dcl.IsEmptyValueIndirect(desired.ServiceAccount) {
		c.Config.Logger.Infof("desired ServiceAccount %s - but actually nil", dcl.SprintResource(desired.ServiceAccount))
		return true
	}
	if !dcl.StringCanonicalize(desired.ServiceAccount, actual.ServiceAccount) && !dcl.IsZeroValue(desired.ServiceAccount) {
		c.Config.Logger.Infof("Diff in ServiceAccount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ServiceAccount), dcl.SprintResource(actual.ServiceAccount))
		return true
	}
	if actual.Metadata == nil && desired.Metadata != nil && !dcl.IsEmptyValueIndirect(desired.Metadata) {
		c.Config.Logger.Infof("desired Metadata %s - but actually nil", dcl.SprintResource(desired.Metadata))
		return true
	}
	if !dcl.MapEquals(desired.Metadata, actual.Metadata, []string(nil)) && !dcl.IsZeroValue(desired.Metadata) {
		c.Config.Logger.Infof("Diff in Metadata. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Metadata), dcl.SprintResource(actual.Metadata))
		return true
	}
	if actual.ImageType == nil && desired.ImageType != nil && !dcl.IsEmptyValueIndirect(desired.ImageType) {
		c.Config.Logger.Infof("desired ImageType %s - but actually nil", dcl.SprintResource(desired.ImageType))
		return true
	}
	if !dcl.StringCanonicalize(desired.ImageType, actual.ImageType) && !dcl.IsZeroValue(desired.ImageType) {
		c.Config.Logger.Infof("Diff in ImageType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ImageType), dcl.SprintResource(actual.ImageType))
		return true
	}
	if actual.Labels == nil && desired.Labels != nil && !dcl.IsEmptyValueIndirect(desired.Labels) {
		c.Config.Logger.Infof("desired Labels %s - but actually nil", dcl.SprintResource(desired.Labels))
		return true
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) && !dcl.IsZeroValue(desired.Labels) {
		c.Config.Logger.Infof("Diff in Labels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Labels), dcl.SprintResource(actual.Labels))
		return true
	}
	if actual.LocalSsdCount == nil && desired.LocalSsdCount != nil && !dcl.IsEmptyValueIndirect(desired.LocalSsdCount) {
		c.Config.Logger.Infof("desired LocalSsdCount %s - but actually nil", dcl.SprintResource(desired.LocalSsdCount))
		return true
	}
	if !reflect.DeepEqual(desired.LocalSsdCount, actual.LocalSsdCount) && !dcl.IsZeroValue(desired.LocalSsdCount) {
		c.Config.Logger.Infof("Diff in LocalSsdCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocalSsdCount), dcl.SprintResource(actual.LocalSsdCount))
		return true
	}
	if actual.Tags == nil && desired.Tags != nil && !dcl.IsEmptyValueIndirect(desired.Tags) {
		c.Config.Logger.Infof("desired Tags %s - but actually nil", dcl.SprintResource(desired.Tags))
		return true
	}
	if !dcl.StringSliceEquals(desired.Tags, actual.Tags) && !dcl.IsZeroValue(desired.Tags) {
		c.Config.Logger.Infof("Diff in Tags. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tags), dcl.SprintResource(actual.Tags))
		return true
	}
	if actual.Preemptible == nil && desired.Preemptible != nil && !dcl.IsEmptyValueIndirect(desired.Preemptible) {
		c.Config.Logger.Infof("desired Preemptible %s - but actually nil", dcl.SprintResource(desired.Preemptible))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Preemptible, actual.Preemptible) && !dcl.IsZeroValue(desired.Preemptible) {
		c.Config.Logger.Infof("Diff in Preemptible. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptible), dcl.SprintResource(actual.Preemptible))
		return true
	}
	if actual.Accelerators == nil && desired.Accelerators != nil && !dcl.IsEmptyValueIndirect(desired.Accelerators) {
		c.Config.Logger.Infof("desired Accelerators %s - but actually nil", dcl.SprintResource(desired.Accelerators))
		return true
	}
	if compareNodePoolConfigAcceleratorsSlice(c, desired.Accelerators, actual.Accelerators) && !dcl.IsZeroValue(desired.Accelerators) {
		c.Config.Logger.Infof("Diff in Accelerators. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Accelerators), dcl.SprintResource(actual.Accelerators))
		return true
	}
	if actual.DiskType == nil && desired.DiskType != nil && !dcl.IsEmptyValueIndirect(desired.DiskType) {
		c.Config.Logger.Infof("desired DiskType %s - but actually nil", dcl.SprintResource(desired.DiskType))
		return true
	}
	if !dcl.StringCanonicalize(desired.DiskType, actual.DiskType) && !dcl.IsZeroValue(desired.DiskType) {
		c.Config.Logger.Infof("Diff in DiskType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DiskType), dcl.SprintResource(actual.DiskType))
		return true
	}
	if actual.MinCpuPlatform == nil && desired.MinCpuPlatform != nil && !dcl.IsEmptyValueIndirect(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("desired MinCpuPlatform %s - but actually nil", dcl.SprintResource(desired.MinCpuPlatform))
		return true
	}
	if !dcl.StringCanonicalize(desired.MinCpuPlatform, actual.MinCpuPlatform) && !dcl.IsZeroValue(desired.MinCpuPlatform) {
		c.Config.Logger.Infof("Diff in MinCpuPlatform. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinCpuPlatform), dcl.SprintResource(actual.MinCpuPlatform))
		return true
	}
	if actual.Taints == nil && desired.Taints != nil && !dcl.IsEmptyValueIndirect(desired.Taints) {
		c.Config.Logger.Infof("desired Taints %s - but actually nil", dcl.SprintResource(desired.Taints))
		return true
	}
	if compareNodePoolConfigTaintsSlice(c, desired.Taints, actual.Taints) && !dcl.IsZeroValue(desired.Taints) {
		c.Config.Logger.Infof("Diff in Taints. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Taints), dcl.SprintResource(actual.Taints))
		return true
	}
	if actual.SandboxConfig == nil && desired.SandboxConfig != nil && !dcl.IsEmptyValueIndirect(desired.SandboxConfig) {
		c.Config.Logger.Infof("desired SandboxConfig %s - but actually nil", dcl.SprintResource(desired.SandboxConfig))
		return true
	}
	if compareNodePoolConfigSandboxConfig(c, desired.SandboxConfig, actual.SandboxConfig) && !dcl.IsZeroValue(desired.SandboxConfig) {
		c.Config.Logger.Infof("Diff in SandboxConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SandboxConfig), dcl.SprintResource(actual.SandboxConfig))
		return true
	}
	if actual.ReservationAffinity == nil && desired.ReservationAffinity != nil && !dcl.IsEmptyValueIndirect(desired.ReservationAffinity) {
		c.Config.Logger.Infof("desired ReservationAffinity %s - but actually nil", dcl.SprintResource(desired.ReservationAffinity))
		return true
	}
	if compareNodePoolConfigReservationAffinity(c, desired.ReservationAffinity, actual.ReservationAffinity) && !dcl.IsZeroValue(desired.ReservationAffinity) {
		c.Config.Logger.Infof("Diff in ReservationAffinity. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReservationAffinity), dcl.SprintResource(actual.ReservationAffinity))
		return true
	}
	if actual.ShieldedInstanceConfig == nil && desired.ShieldedInstanceConfig != nil && !dcl.IsEmptyValueIndirect(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("desired ShieldedInstanceConfig %s - but actually nil", dcl.SprintResource(desired.ShieldedInstanceConfig))
		return true
	}
	if compareNodePoolConfigShieldedInstanceConfig(c, desired.ShieldedInstanceConfig, actual.ShieldedInstanceConfig) && !dcl.IsZeroValue(desired.ShieldedInstanceConfig) {
		c.Config.Logger.Infof("Diff in ShieldedInstanceConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ShieldedInstanceConfig), dcl.SprintResource(actual.ShieldedInstanceConfig))
		return true
	}
	return false
}

func compareNodePoolConfigSlice(c *Client, desired, actual []NodePoolConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigMap(c *Client, desired, actual map[string]NodePoolConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConfigAccelerators(c *Client, desired, actual *NodePoolConfigAccelerators) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AcceleratorCount == nil && desired.AcceleratorCount != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorCount) {
		c.Config.Logger.Infof("desired AcceleratorCount %s - but actually nil", dcl.SprintResource(desired.AcceleratorCount))
		return true
	}
	if !reflect.DeepEqual(desired.AcceleratorCount, actual.AcceleratorCount) && !dcl.IsZeroValue(desired.AcceleratorCount) {
		c.Config.Logger.Infof("Diff in AcceleratorCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorCount), dcl.SprintResource(actual.AcceleratorCount))
		return true
	}
	if actual.AcceleratorType == nil && desired.AcceleratorType != nil && !dcl.IsEmptyValueIndirect(desired.AcceleratorType) {
		c.Config.Logger.Infof("desired AcceleratorType %s - but actually nil", dcl.SprintResource(desired.AcceleratorType))
		return true
	}
	if !dcl.StringCanonicalize(desired.AcceleratorType, actual.AcceleratorType) && !dcl.IsZeroValue(desired.AcceleratorType) {
		c.Config.Logger.Infof("Diff in AcceleratorType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AcceleratorType), dcl.SprintResource(actual.AcceleratorType))
		return true
	}
	return false
}

func compareNodePoolConfigAcceleratorsSlice(c *Client, desired, actual []NodePoolConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigAccelerators, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigAccelerators(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigAccelerators, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigAcceleratorsMap(c *Client, desired, actual map[string]NodePoolConfigAccelerators) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigAccelerators, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConfigAccelerators, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConfigAccelerators(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConfigAccelerators, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConfigTaints(c *Client, desired, actual *NodePoolConfigTaints) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	if actual.Effect == nil && desired.Effect != nil && !dcl.IsEmptyValueIndirect(desired.Effect) {
		c.Config.Logger.Infof("desired Effect %s - but actually nil", dcl.SprintResource(desired.Effect))
		return true
	}
	if !dcl.StringCanonicalize(desired.Effect, actual.Effect) && !dcl.IsZeroValue(desired.Effect) {
		c.Config.Logger.Infof("Diff in Effect. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Effect), dcl.SprintResource(actual.Effect))
		return true
	}
	return false
}

func compareNodePoolConfigTaintsSlice(c *Client, desired, actual []NodePoolConfigTaints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigTaints, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigTaints(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigTaints, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigTaintsMap(c *Client, desired, actual map[string]NodePoolConfigTaints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigTaints, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConfigTaints, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConfigTaints(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConfigTaints, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConfigSandboxConfig(c *Client, desired, actual *NodePoolConfigSandboxConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	return false
}

func compareNodePoolConfigSandboxConfigSlice(c *Client, desired, actual []NodePoolConfigSandboxConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigSandboxConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigSandboxConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigSandboxConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigSandboxConfigMap(c *Client, desired, actual map[string]NodePoolConfigSandboxConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigSandboxConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConfigSandboxConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConfigSandboxConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConfigSandboxConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConfigReservationAffinity(c *Client, desired, actual *NodePoolConfigReservationAffinity) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ConsumeReservationType == nil && desired.ConsumeReservationType != nil && !dcl.IsEmptyValueIndirect(desired.ConsumeReservationType) {
		c.Config.Logger.Infof("desired ConsumeReservationType %s - but actually nil", dcl.SprintResource(desired.ConsumeReservationType))
		return true
	}
	if !reflect.DeepEqual(desired.ConsumeReservationType, actual.ConsumeReservationType) && !dcl.IsZeroValue(desired.ConsumeReservationType) {
		c.Config.Logger.Infof("Diff in ConsumeReservationType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConsumeReservationType), dcl.SprintResource(actual.ConsumeReservationType))
		return true
	}
	if actual.Key == nil && desired.Key != nil && !dcl.IsEmptyValueIndirect(desired.Key) {
		c.Config.Logger.Infof("desired Key %s - but actually nil", dcl.SprintResource(desired.Key))
		return true
	}
	if !dcl.StringCanonicalize(desired.Key, actual.Key) && !dcl.IsZeroValue(desired.Key) {
		c.Config.Logger.Infof("Diff in Key. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Key), dcl.SprintResource(actual.Key))
		return true
	}
	if actual.Values == nil && desired.Values != nil && !dcl.IsEmptyValueIndirect(desired.Values) {
		c.Config.Logger.Infof("desired Values %s - but actually nil", dcl.SprintResource(desired.Values))
		return true
	}
	if !dcl.StringSliceEquals(desired.Values, actual.Values) && !dcl.IsZeroValue(desired.Values) {
		c.Config.Logger.Infof("Diff in Values. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Values), dcl.SprintResource(actual.Values))
		return true
	}
	return false
}

func compareNodePoolConfigReservationAffinitySlice(c *Client, desired, actual []NodePoolConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigReservationAffinity, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigReservationAffinity(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigReservationAffinity, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigReservationAffinityMap(c *Client, desired, actual map[string]NodePoolConfigReservationAffinity) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigReservationAffinity, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConfigReservationAffinity, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConfigReservationAffinity(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConfigReservationAffinity, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConfigShieldedInstanceConfig(c *Client, desired, actual *NodePoolConfigShieldedInstanceConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.EnableSecureBoot == nil && desired.EnableSecureBoot != nil && !dcl.IsEmptyValueIndirect(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("desired EnableSecureBoot %s - but actually nil", dcl.SprintResource(desired.EnableSecureBoot))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableSecureBoot, actual.EnableSecureBoot) && !dcl.IsZeroValue(desired.EnableSecureBoot) {
		c.Config.Logger.Infof("Diff in EnableSecureBoot. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableSecureBoot), dcl.SprintResource(actual.EnableSecureBoot))
		return true
	}
	if actual.EnableIntegrityMonitoring == nil && desired.EnableIntegrityMonitoring != nil && !dcl.IsEmptyValueIndirect(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("desired EnableIntegrityMonitoring %s - but actually nil", dcl.SprintResource(desired.EnableIntegrityMonitoring))
		return true
	}
	if !dcl.BoolCanonicalize(desired.EnableIntegrityMonitoring, actual.EnableIntegrityMonitoring) && !dcl.IsZeroValue(desired.EnableIntegrityMonitoring) {
		c.Config.Logger.Infof("Diff in EnableIntegrityMonitoring. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EnableIntegrityMonitoring), dcl.SprintResource(actual.EnableIntegrityMonitoring))
		return true
	}
	return false
}

func compareNodePoolConfigShieldedInstanceConfigSlice(c *Client, desired, actual []NodePoolConfigShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigShieldedInstanceConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigShieldedInstanceConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigShieldedInstanceConfigMap(c *Client, desired, actual map[string]NodePoolConfigShieldedInstanceConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigShieldedInstanceConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConfigShieldedInstanceConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConfigShieldedInstanceConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConfigShieldedInstanceConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolAutoscaling(c *Client, desired, actual *NodePoolAutoscaling) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enabled == nil && desired.Enabled != nil && !dcl.IsEmptyValueIndirect(desired.Enabled) {
		c.Config.Logger.Infof("desired Enabled %s - but actually nil", dcl.SprintResource(desired.Enabled))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if actual.MinNodeCount == nil && desired.MinNodeCount != nil && !dcl.IsEmptyValueIndirect(desired.MinNodeCount) {
		c.Config.Logger.Infof("desired MinNodeCount %s - but actually nil", dcl.SprintResource(desired.MinNodeCount))
		return true
	}
	if !reflect.DeepEqual(desired.MinNodeCount, actual.MinNodeCount) && !dcl.IsZeroValue(desired.MinNodeCount) {
		c.Config.Logger.Infof("Diff in MinNodeCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinNodeCount), dcl.SprintResource(actual.MinNodeCount))
		return true
	}
	if actual.MaxNodeCount == nil && desired.MaxNodeCount != nil && !dcl.IsEmptyValueIndirect(desired.MaxNodeCount) {
		c.Config.Logger.Infof("desired MaxNodeCount %s - but actually nil", dcl.SprintResource(desired.MaxNodeCount))
		return true
	}
	if !reflect.DeepEqual(desired.MaxNodeCount, actual.MaxNodeCount) && !dcl.IsZeroValue(desired.MaxNodeCount) {
		c.Config.Logger.Infof("Diff in MaxNodeCount. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxNodeCount), dcl.SprintResource(actual.MaxNodeCount))
		return true
	}
	if actual.Autoprovisioned == nil && desired.Autoprovisioned != nil && !dcl.IsEmptyValueIndirect(desired.Autoprovisioned) {
		c.Config.Logger.Infof("desired Autoprovisioned %s - but actually nil", dcl.SprintResource(desired.Autoprovisioned))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Autoprovisioned, actual.Autoprovisioned) && !dcl.IsZeroValue(desired.Autoprovisioned) {
		c.Config.Logger.Infof("Diff in Autoprovisioned. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Autoprovisioned), dcl.SprintResource(actual.Autoprovisioned))
		return true
	}
	return false
}

func compareNodePoolAutoscalingSlice(c *Client, desired, actual []NodePoolAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolAutoscaling, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolAutoscaling(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolAutoscaling, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolAutoscalingMap(c *Client, desired, actual map[string]NodePoolAutoscaling) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolAutoscaling, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolAutoscaling, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolAutoscaling(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolAutoscaling, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolManagement(c *Client, desired, actual *NodePoolManagement) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.AutoUpgrade == nil && desired.AutoUpgrade != nil && !dcl.IsEmptyValueIndirect(desired.AutoUpgrade) {
		c.Config.Logger.Infof("desired AutoUpgrade %s - but actually nil", dcl.SprintResource(desired.AutoUpgrade))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoUpgrade, actual.AutoUpgrade) && !dcl.IsZeroValue(desired.AutoUpgrade) {
		c.Config.Logger.Infof("Diff in AutoUpgrade. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoUpgrade), dcl.SprintResource(actual.AutoUpgrade))
		return true
	}
	if actual.AutoRepair == nil && desired.AutoRepair != nil && !dcl.IsEmptyValueIndirect(desired.AutoRepair) {
		c.Config.Logger.Infof("desired AutoRepair %s - but actually nil", dcl.SprintResource(desired.AutoRepair))
		return true
	}
	if !dcl.BoolCanonicalize(desired.AutoRepair, actual.AutoRepair) && !dcl.IsZeroValue(desired.AutoRepair) {
		c.Config.Logger.Infof("Diff in AutoRepair. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoRepair), dcl.SprintResource(actual.AutoRepair))
		return true
	}
	if actual.UpgradeOptions == nil && desired.UpgradeOptions != nil && !dcl.IsEmptyValueIndirect(desired.UpgradeOptions) {
		c.Config.Logger.Infof("desired UpgradeOptions %s - but actually nil", dcl.SprintResource(desired.UpgradeOptions))
		return true
	}
	if compareNodePoolManagementUpgradeOptions(c, desired.UpgradeOptions, actual.UpgradeOptions) && !dcl.IsZeroValue(desired.UpgradeOptions) {
		c.Config.Logger.Infof("Diff in UpgradeOptions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpgradeOptions), dcl.SprintResource(actual.UpgradeOptions))
		return true
	}
	return false
}

func compareNodePoolManagementSlice(c *Client, desired, actual []NodePoolManagement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolManagement, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolManagement(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolManagement, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolManagementMap(c *Client, desired, actual map[string]NodePoolManagement) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolManagement, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolManagement, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolManagement(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolManagement, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolManagementUpgradeOptions(c *Client, desired, actual *NodePoolManagementUpgradeOptions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareNodePoolManagementUpgradeOptionsSlice(c *Client, desired, actual []NodePoolManagementUpgradeOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolManagementUpgradeOptions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolManagementUpgradeOptions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolManagementUpgradeOptions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolManagementUpgradeOptionsMap(c *Client, desired, actual map[string]NodePoolManagementUpgradeOptions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolManagementUpgradeOptions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolManagementUpgradeOptions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolManagementUpgradeOptions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolManagementUpgradeOptions, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolMaxPodsConstraint(c *Client, desired, actual *NodePoolMaxPodsConstraint) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MaxPodsPerNode == nil && desired.MaxPodsPerNode != nil && !dcl.IsEmptyValueIndirect(desired.MaxPodsPerNode) {
		c.Config.Logger.Infof("desired MaxPodsPerNode %s - but actually nil", dcl.SprintResource(desired.MaxPodsPerNode))
		return true
	}
	if !reflect.DeepEqual(desired.MaxPodsPerNode, actual.MaxPodsPerNode) && !dcl.IsZeroValue(desired.MaxPodsPerNode) {
		c.Config.Logger.Infof("Diff in MaxPodsPerNode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxPodsPerNode), dcl.SprintResource(actual.MaxPodsPerNode))
		return true
	}
	return false
}

func compareNodePoolMaxPodsConstraintSlice(c *Client, desired, actual []NodePoolMaxPodsConstraint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolMaxPodsConstraint, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolMaxPodsConstraint(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolMaxPodsConstraint, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolMaxPodsConstraintMap(c *Client, desired, actual map[string]NodePoolMaxPodsConstraint) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolMaxPodsConstraint, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolMaxPodsConstraint, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolMaxPodsConstraint(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolMaxPodsConstraint, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConditions(c *Client, desired, actual *NodePoolConditions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareNodePoolConditionsSlice(c *Client, desired, actual []NodePoolConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConditions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConditions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConditions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConditionsMap(c *Client, desired, actual map[string]NodePoolConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConditions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolConditions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolConditions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolConditions, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolUpgradeSettings(c *Client, desired, actual *NodePoolUpgradeSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MaxSurge == nil && desired.MaxSurge != nil && !dcl.IsEmptyValueIndirect(desired.MaxSurge) {
		c.Config.Logger.Infof("desired MaxSurge %s - but actually nil", dcl.SprintResource(desired.MaxSurge))
		return true
	}
	if !reflect.DeepEqual(desired.MaxSurge, actual.MaxSurge) && !dcl.IsZeroValue(desired.MaxSurge) {
		c.Config.Logger.Infof("Diff in MaxSurge. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxSurge), dcl.SprintResource(actual.MaxSurge))
		return true
	}
	if actual.MaxUnavailable == nil && desired.MaxUnavailable != nil && !dcl.IsEmptyValueIndirect(desired.MaxUnavailable) {
		c.Config.Logger.Infof("desired MaxUnavailable %s - but actually nil", dcl.SprintResource(desired.MaxUnavailable))
		return true
	}
	if !reflect.DeepEqual(desired.MaxUnavailable, actual.MaxUnavailable) && !dcl.IsZeroValue(desired.MaxUnavailable) {
		c.Config.Logger.Infof("Diff in MaxUnavailable. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxUnavailable), dcl.SprintResource(actual.MaxUnavailable))
		return true
	}
	return false
}

func compareNodePoolUpgradeSettingsSlice(c *Client, desired, actual []NodePoolUpgradeSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolUpgradeSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolUpgradeSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolUpgradeSettings, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolUpgradeSettingsMap(c *Client, desired, actual map[string]NodePoolUpgradeSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolUpgradeSettings, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NodePoolUpgradeSettings, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNodePoolUpgradeSettings(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NodePoolUpgradeSettings, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNodePoolConfigSandboxConfigTypeEnumSlice(c *Client, desired, actual []NodePoolConfigSandboxConfigTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigSandboxConfigTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigSandboxConfigTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigSandboxConfigTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigSandboxConfigTypeEnum(c *Client, desired, actual *NodePoolConfigSandboxConfigTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNodePoolConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, desired, actual []NodePoolConfigReservationAffinityConsumeReservationTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConfigReservationAffinityConsumeReservationTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConfigReservationAffinityConsumeReservationTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConfigReservationAffinityConsumeReservationTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConfigReservationAffinityConsumeReservationTypeEnum(c *Client, desired, actual *NodePoolConfigReservationAffinityConsumeReservationTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNodePoolConditionsCodeEnumSlice(c *Client, desired, actual []NodePoolConditionsCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodePoolConditionsCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodePoolConditionsCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodePoolConditionsCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodePoolConditionsCodeEnum(c *Client, desired, actual *NodePoolConditionsCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *NodePool) urlNormalized() *NodePool {
	normalized := deepcopy.Copy(*r).(NodePool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Version = dcl.SelfLinkToName(r.Version)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.Cluster = dcl.SelfLinkToName(r.Cluster)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *NodePool) getFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Cluster), dcl.ValueOrEmptyString(n.Name)
}

func (r *NodePool) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Cluster)
}

func (r *NodePool) deleteFields() (string, string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Cluster), dcl.ValueOrEmptyString(n.Name)
}

func (r *NodePool) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setAutoscaling" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"cluster":  dcl.ValueOrEmptyString(n.Cluster),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}:setAutoscaling", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "setManagement" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"cluster":  dcl.ValueOrEmptyString(n.Cluster),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}:setManagement", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "setSize" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"cluster":  dcl.ValueOrEmptyString(n.Cluster),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}:setSize", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	if updateName == "update" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"cluster":  dcl.ValueOrEmptyString(n.Cluster),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}", "https://container.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the NodePool resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *NodePool) marshal(c *Client) ([]byte, error) {
	m, err := expandNodePool(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling NodePool: %w", err)
	}
	m = EncodeNodePoolCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalNodePool decodes JSON responses into the NodePool resource schema.
func unmarshalNodePool(b []byte, c *Client) (*NodePool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNodePool(m, c)
}

func unmarshalMapNodePool(m map[string]interface{}, c *Client) (*NodePool, error) {

	return flattenNodePool(c, m), nil
}

// expandNodePool expands NodePool into a JSON request object.
func expandNodePool(c *Client, f *NodePool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandNodePoolConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["config"] = v
	}
	if v := f.NodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["initialNodeCount"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.StatusMessage; !dcl.IsEmptyValueIndirect(v) {
		m["statusMessage"] = v
	}
	if v := f.Locations; !dcl.IsEmptyValueIndirect(v) {
		m["locations"] = v
	}
	if v, err := expandNodePoolAutoscaling(c, f.Autoscaling); err != nil {
		return nil, fmt.Errorf("error expanding Autoscaling into autoscaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscaling"] = v
	}
	if v, err := expandNodePoolManagement(c, f.Management); err != nil {
		return nil, fmt.Errorf("error expanding Management into management: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["management"] = v
	}
	if v, err := expandNodePoolMaxPodsConstraint(c, f.MaxPodsConstraint); err != nil {
		return nil, fmt.Errorf("error expanding MaxPodsConstraint into maxPodsConstraint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsConstraint"] = v
	}
	if v, err := expandNodePoolConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditions"] = v
	}
	if v := f.PodIPv4CidrSize; !dcl.IsEmptyValueIndirect(v) {
		m["podIPv4CidrSize"] = v
	}
	if v, err := expandNodePoolUpgradeSettings(c, f.UpgradeSettings); err != nil {
		return nil, fmt.Errorf("error expanding UpgradeSettings into upgradeSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["upgradeSettings"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Cluster into cluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cluster"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenNodePool flattens NodePool from a JSON request object into the
// NodePool type.
func flattenNodePool(c *Client, i interface{}) *NodePool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &NodePool{}
	r.Name = dcl.FlattenString(m["name"])
	r.Config = flattenNodePoolConfig(c, m["config"])
	r.NodeCount = dcl.FlattenInteger(m["initialNodeCount"])
	r.Version = dcl.FlattenString(m["version"])
	r.Status = dcl.FlattenString(m["status"])
	r.StatusMessage = dcl.FlattenString(m["statusMessage"])
	r.Locations = dcl.FlattenStringSlice(m["locations"])
	r.Autoscaling = flattenNodePoolAutoscaling(c, m["autoscaling"])
	r.Management = flattenNodePoolManagement(c, m["management"])
	r.MaxPodsConstraint = flattenNodePoolMaxPodsConstraint(c, m["maxPodsConstraint"])
	r.Conditions = flattenNodePoolConditionsSlice(c, m["conditions"])
	r.PodIPv4CidrSize = dcl.FlattenInteger(m["podIPv4CidrSize"])
	r.UpgradeSettings = flattenNodePoolUpgradeSettings(c, m["upgradeSettings"])
	r.Cluster = dcl.FlattenString(m["cluster"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandNodePoolConfigMap expands the contents of NodePoolConfig into a JSON
// request object.
func expandNodePoolConfigMap(c *Client, f map[string]NodePoolConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConfigSlice expands the contents of NodePoolConfig into a JSON
// request object.
func expandNodePoolConfigSlice(c *Client, f []NodePoolConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConfigMap flattens the contents of NodePoolConfig from a JSON
// response object.
func flattenNodePoolConfigMap(c *Client, i interface{}) map[string]NodePoolConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConfig{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConfig{}
	}

	items := make(map[string]NodePoolConfig)
	for k, item := range a {
		items[k] = *flattenNodePoolConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConfigSlice flattens the contents of NodePoolConfig from a JSON
// response object.
func flattenNodePoolConfigSlice(c *Client, i interface{}) []NodePoolConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfig{}
	}

	if len(a) == 0 {
		return []NodePoolConfig{}
	}

	items := make([]NodePoolConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConfig expands an instance of NodePoolConfig into a JSON
// request object.
func expandNodePoolConfig(c *Client, f *NodePoolConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MachineType; !dcl.IsEmptyValueIndirect(v) {
		m["machineType"] = v
	}
	if v := f.DiskSizeGb; !dcl.IsEmptyValueIndirect(v) {
		m["diskSizeGb"] = v
	}
	if v := f.OAuthScopes; !dcl.IsEmptyValueIndirect(v) {
		m["oauthScopes"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}
	if v := f.ImageType; !dcl.IsEmptyValueIndirect(v) {
		m["imageType"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.LocalSsdCount; !dcl.IsEmptyValueIndirect(v) {
		m["localSsdCount"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.Preemptible; !dcl.IsEmptyValueIndirect(v) {
		m["preemptible"] = v
	}
	if v, err := expandNodePoolConfigAcceleratorsSlice(c, f.Accelerators); err != nil {
		return nil, fmt.Errorf("error expanding Accelerators into accelerators: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["accelerators"] = v
	}
	if v := f.DiskType; !dcl.IsEmptyValueIndirect(v) {
		m["diskType"] = v
	}
	if v := f.MinCpuPlatform; !dcl.IsEmptyValueIndirect(v) {
		m["minCpuPlatform"] = v
	}
	if v, err := expandNodePoolConfigTaintsSlice(c, f.Taints); err != nil {
		return nil, fmt.Errorf("error expanding Taints into taints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["taints"] = v
	}
	if v, err := expandNodePoolConfigSandboxConfig(c, f.SandboxConfig); err != nil {
		return nil, fmt.Errorf("error expanding SandboxConfig into sandboxConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sandboxConfig"] = v
	}
	if v, err := expandNodePoolConfigReservationAffinity(c, f.ReservationAffinity); err != nil {
		return nil, fmt.Errorf("error expanding ReservationAffinity into reservationAffinity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reservationAffinity"] = v
	}
	if v, err := expandNodePoolConfigShieldedInstanceConfig(c, f.ShieldedInstanceConfig); err != nil {
		return nil, fmt.Errorf("error expanding ShieldedInstanceConfig into shieldedInstanceConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["shieldedInstanceConfig"] = v
	}

	return m, nil
}

// flattenNodePoolConfig flattens an instance of NodePoolConfig from a JSON
// response object.
func flattenNodePoolConfig(c *Client, i interface{}) *NodePoolConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConfig{}
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.OAuthScopes = dcl.FlattenStringSlice(m["oauthScopes"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Metadata = dcl.FlattenKeyValuePairs(m["metadata"])
	r.ImageType = dcl.FlattenString(m["imageType"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.LocalSsdCount = dcl.FlattenInteger(m["localSsdCount"])
	r.Tags = dcl.FlattenStringSlice(m["tags"])
	r.Preemptible = dcl.FlattenBool(m["preemptible"])
	r.Accelerators = flattenNodePoolConfigAcceleratorsSlice(c, m["accelerators"])
	r.DiskType = dcl.FlattenString(m["diskType"])
	r.MinCpuPlatform = dcl.FlattenString(m["minCpuPlatform"])
	r.Taints = flattenNodePoolConfigTaintsSlice(c, m["taints"])
	r.SandboxConfig = flattenNodePoolConfigSandboxConfig(c, m["sandboxConfig"])
	r.ReservationAffinity = flattenNodePoolConfigReservationAffinity(c, m["reservationAffinity"])
	r.ShieldedInstanceConfig = flattenNodePoolConfigShieldedInstanceConfig(c, m["shieldedInstanceConfig"])

	return r
}

// expandNodePoolConfigAcceleratorsMap expands the contents of NodePoolConfigAccelerators into a JSON
// request object.
func expandNodePoolConfigAcceleratorsMap(c *Client, f map[string]NodePoolConfigAccelerators) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConfigAcceleratorsSlice expands the contents of NodePoolConfigAccelerators into a JSON
// request object.
func expandNodePoolConfigAcceleratorsSlice(c *Client, f []NodePoolConfigAccelerators) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConfigAccelerators(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConfigAcceleratorsMap flattens the contents of NodePoolConfigAccelerators from a JSON
// response object.
func flattenNodePoolConfigAcceleratorsMap(c *Client, i interface{}) map[string]NodePoolConfigAccelerators {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConfigAccelerators{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConfigAccelerators{}
	}

	items := make(map[string]NodePoolConfigAccelerators)
	for k, item := range a {
		items[k] = *flattenNodePoolConfigAccelerators(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConfigAcceleratorsSlice flattens the contents of NodePoolConfigAccelerators from a JSON
// response object.
func flattenNodePoolConfigAcceleratorsSlice(c *Client, i interface{}) []NodePoolConfigAccelerators {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigAccelerators{}
	}

	if len(a) == 0 {
		return []NodePoolConfigAccelerators{}
	}

	items := make([]NodePoolConfigAccelerators, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigAccelerators(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConfigAccelerators expands an instance of NodePoolConfigAccelerators into a JSON
// request object.
func expandNodePoolConfigAccelerators(c *Client, f *NodePoolConfigAccelerators) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AcceleratorCount; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorCount"] = v
	}
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}

	return m, nil
}

// flattenNodePoolConfigAccelerators flattens an instance of NodePoolConfigAccelerators from a JSON
// response object.
func flattenNodePoolConfigAccelerators(c *Client, i interface{}) *NodePoolConfigAccelerators {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConfigAccelerators{}
	r.AcceleratorCount = dcl.FlattenInteger(m["acceleratorCount"])
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])

	return r
}

// expandNodePoolConfigTaintsMap expands the contents of NodePoolConfigTaints into a JSON
// request object.
func expandNodePoolConfigTaintsMap(c *Client, f map[string]NodePoolConfigTaints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConfigTaintsSlice expands the contents of NodePoolConfigTaints into a JSON
// request object.
func expandNodePoolConfigTaintsSlice(c *Client, f []NodePoolConfigTaints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConfigTaints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConfigTaintsMap flattens the contents of NodePoolConfigTaints from a JSON
// response object.
func flattenNodePoolConfigTaintsMap(c *Client, i interface{}) map[string]NodePoolConfigTaints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConfigTaints{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConfigTaints{}
	}

	items := make(map[string]NodePoolConfigTaints)
	for k, item := range a {
		items[k] = *flattenNodePoolConfigTaints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConfigTaintsSlice flattens the contents of NodePoolConfigTaints from a JSON
// response object.
func flattenNodePoolConfigTaintsSlice(c *Client, i interface{}) []NodePoolConfigTaints {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigTaints{}
	}

	if len(a) == 0 {
		return []NodePoolConfigTaints{}
	}

	items := make([]NodePoolConfigTaints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigTaints(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConfigTaints expands an instance of NodePoolConfigTaints into a JSON
// request object.
func expandNodePoolConfigTaints(c *Client, f *NodePoolConfigTaints) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.Effect; !dcl.IsEmptyValueIndirect(v) {
		m["effect"] = v
	}

	return m, nil
}

// flattenNodePoolConfigTaints flattens an instance of NodePoolConfigTaints from a JSON
// response object.
func flattenNodePoolConfigTaints(c *Client, i interface{}) *NodePoolConfigTaints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConfigTaints{}
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])
	r.Effect = dcl.FlattenString(m["effect"])

	return r
}

// expandNodePoolConfigSandboxConfigMap expands the contents of NodePoolConfigSandboxConfig into a JSON
// request object.
func expandNodePoolConfigSandboxConfigMap(c *Client, f map[string]NodePoolConfigSandboxConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConfigSandboxConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConfigSandboxConfigSlice expands the contents of NodePoolConfigSandboxConfig into a JSON
// request object.
func expandNodePoolConfigSandboxConfigSlice(c *Client, f []NodePoolConfigSandboxConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConfigSandboxConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConfigSandboxConfigMap flattens the contents of NodePoolConfigSandboxConfig from a JSON
// response object.
func flattenNodePoolConfigSandboxConfigMap(c *Client, i interface{}) map[string]NodePoolConfigSandboxConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConfigSandboxConfig{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConfigSandboxConfig{}
	}

	items := make(map[string]NodePoolConfigSandboxConfig)
	for k, item := range a {
		items[k] = *flattenNodePoolConfigSandboxConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConfigSandboxConfigSlice flattens the contents of NodePoolConfigSandboxConfig from a JSON
// response object.
func flattenNodePoolConfigSandboxConfigSlice(c *Client, i interface{}) []NodePoolConfigSandboxConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigSandboxConfig{}
	}

	if len(a) == 0 {
		return []NodePoolConfigSandboxConfig{}
	}

	items := make([]NodePoolConfigSandboxConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigSandboxConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConfigSandboxConfig expands an instance of NodePoolConfigSandboxConfig into a JSON
// request object.
func expandNodePoolConfigSandboxConfig(c *Client, f *NodePoolConfigSandboxConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenNodePoolConfigSandboxConfig flattens an instance of NodePoolConfigSandboxConfig from a JSON
// response object.
func flattenNodePoolConfigSandboxConfig(c *Client, i interface{}) *NodePoolConfigSandboxConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConfigSandboxConfig{}
	r.Type = flattenNodePoolConfigSandboxConfigTypeEnum(m["type"])

	return r
}

// expandNodePoolConfigReservationAffinityMap expands the contents of NodePoolConfigReservationAffinity into a JSON
// request object.
func expandNodePoolConfigReservationAffinityMap(c *Client, f map[string]NodePoolConfigReservationAffinity) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConfigReservationAffinitySlice expands the contents of NodePoolConfigReservationAffinity into a JSON
// request object.
func expandNodePoolConfigReservationAffinitySlice(c *Client, f []NodePoolConfigReservationAffinity) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConfigReservationAffinity(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConfigReservationAffinityMap flattens the contents of NodePoolConfigReservationAffinity from a JSON
// response object.
func flattenNodePoolConfigReservationAffinityMap(c *Client, i interface{}) map[string]NodePoolConfigReservationAffinity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConfigReservationAffinity{}
	}

	items := make(map[string]NodePoolConfigReservationAffinity)
	for k, item := range a {
		items[k] = *flattenNodePoolConfigReservationAffinity(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConfigReservationAffinitySlice flattens the contents of NodePoolConfigReservationAffinity from a JSON
// response object.
func flattenNodePoolConfigReservationAffinitySlice(c *Client, i interface{}) []NodePoolConfigReservationAffinity {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigReservationAffinity{}
	}

	if len(a) == 0 {
		return []NodePoolConfigReservationAffinity{}
	}

	items := make([]NodePoolConfigReservationAffinity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigReservationAffinity(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConfigReservationAffinity expands an instance of NodePoolConfigReservationAffinity into a JSON
// request object.
func expandNodePoolConfigReservationAffinity(c *Client, f *NodePoolConfigReservationAffinity) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConsumeReservationType; !dcl.IsEmptyValueIndirect(v) {
		m["consumeReservationType"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}

	return m, nil
}

// flattenNodePoolConfigReservationAffinity flattens an instance of NodePoolConfigReservationAffinity from a JSON
// response object.
func flattenNodePoolConfigReservationAffinity(c *Client, i interface{}) *NodePoolConfigReservationAffinity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConfigReservationAffinity{}
	r.ConsumeReservationType = flattenNodePoolConfigReservationAffinityConsumeReservationTypeEnum(m["consumeReservationType"])
	r.Key = dcl.FlattenString(m["key"])
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandNodePoolConfigShieldedInstanceConfigMap expands the contents of NodePoolConfigShieldedInstanceConfig into a JSON
// request object.
func expandNodePoolConfigShieldedInstanceConfigMap(c *Client, f map[string]NodePoolConfigShieldedInstanceConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConfigShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConfigShieldedInstanceConfigSlice expands the contents of NodePoolConfigShieldedInstanceConfig into a JSON
// request object.
func expandNodePoolConfigShieldedInstanceConfigSlice(c *Client, f []NodePoolConfigShieldedInstanceConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConfigShieldedInstanceConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConfigShieldedInstanceConfigMap flattens the contents of NodePoolConfigShieldedInstanceConfig from a JSON
// response object.
func flattenNodePoolConfigShieldedInstanceConfigMap(c *Client, i interface{}) map[string]NodePoolConfigShieldedInstanceConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConfigShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConfigShieldedInstanceConfig{}
	}

	items := make(map[string]NodePoolConfigShieldedInstanceConfig)
	for k, item := range a {
		items[k] = *flattenNodePoolConfigShieldedInstanceConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConfigShieldedInstanceConfigSlice flattens the contents of NodePoolConfigShieldedInstanceConfig from a JSON
// response object.
func flattenNodePoolConfigShieldedInstanceConfigSlice(c *Client, i interface{}) []NodePoolConfigShieldedInstanceConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigShieldedInstanceConfig{}
	}

	if len(a) == 0 {
		return []NodePoolConfigShieldedInstanceConfig{}
	}

	items := make([]NodePoolConfigShieldedInstanceConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigShieldedInstanceConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConfigShieldedInstanceConfig expands an instance of NodePoolConfigShieldedInstanceConfig into a JSON
// request object.
func expandNodePoolConfigShieldedInstanceConfig(c *Client, f *NodePoolConfigShieldedInstanceConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnableSecureBoot; !dcl.IsEmptyValueIndirect(v) {
		m["enableSecureBoot"] = v
	}
	if v := f.EnableIntegrityMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["enableIntegrityMonitoring"] = v
	}

	return m, nil
}

// flattenNodePoolConfigShieldedInstanceConfig flattens an instance of NodePoolConfigShieldedInstanceConfig from a JSON
// response object.
func flattenNodePoolConfigShieldedInstanceConfig(c *Client, i interface{}) *NodePoolConfigShieldedInstanceConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConfigShieldedInstanceConfig{}
	r.EnableSecureBoot = dcl.FlattenBool(m["enableSecureBoot"])
	r.EnableIntegrityMonitoring = dcl.FlattenBool(m["enableIntegrityMonitoring"])

	return r
}

// expandNodePoolAutoscalingMap expands the contents of NodePoolAutoscaling into a JSON
// request object.
func expandNodePoolAutoscalingMap(c *Client, f map[string]NodePoolAutoscaling) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolAutoscalingSlice expands the contents of NodePoolAutoscaling into a JSON
// request object.
func expandNodePoolAutoscalingSlice(c *Client, f []NodePoolAutoscaling) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolAutoscaling(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolAutoscalingMap flattens the contents of NodePoolAutoscaling from a JSON
// response object.
func flattenNodePoolAutoscalingMap(c *Client, i interface{}) map[string]NodePoolAutoscaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolAutoscaling{}
	}

	if len(a) == 0 {
		return map[string]NodePoolAutoscaling{}
	}

	items := make(map[string]NodePoolAutoscaling)
	for k, item := range a {
		items[k] = *flattenNodePoolAutoscaling(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolAutoscalingSlice flattens the contents of NodePoolAutoscaling from a JSON
// response object.
func flattenNodePoolAutoscalingSlice(c *Client, i interface{}) []NodePoolAutoscaling {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolAutoscaling{}
	}

	if len(a) == 0 {
		return []NodePoolAutoscaling{}
	}

	items := make([]NodePoolAutoscaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolAutoscaling(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolAutoscaling expands an instance of NodePoolAutoscaling into a JSON
// request object.
func expandNodePoolAutoscaling(c *Client, f *NodePoolAutoscaling) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.MinNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["minNodeCount"] = v
	}
	if v := f.MaxNodeCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxNodeCount"] = v
	}
	if v := f.Autoprovisioned; !dcl.IsEmptyValueIndirect(v) {
		m["autoprovisioned"] = v
	}

	return m, nil
}

// flattenNodePoolAutoscaling flattens an instance of NodePoolAutoscaling from a JSON
// response object.
func flattenNodePoolAutoscaling(c *Client, i interface{}) *NodePoolAutoscaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolAutoscaling{}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.MinNodeCount = dcl.FlattenInteger(m["minNodeCount"])
	r.MaxNodeCount = dcl.FlattenInteger(m["maxNodeCount"])
	r.Autoprovisioned = dcl.FlattenBool(m["autoprovisioned"])

	return r
}

// expandNodePoolManagementMap expands the contents of NodePoolManagement into a JSON
// request object.
func expandNodePoolManagementMap(c *Client, f map[string]NodePoolManagement) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolManagement(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolManagementSlice expands the contents of NodePoolManagement into a JSON
// request object.
func expandNodePoolManagementSlice(c *Client, f []NodePoolManagement) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolManagement(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolManagementMap flattens the contents of NodePoolManagement from a JSON
// response object.
func flattenNodePoolManagementMap(c *Client, i interface{}) map[string]NodePoolManagement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolManagement{}
	}

	if len(a) == 0 {
		return map[string]NodePoolManagement{}
	}

	items := make(map[string]NodePoolManagement)
	for k, item := range a {
		items[k] = *flattenNodePoolManagement(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolManagementSlice flattens the contents of NodePoolManagement from a JSON
// response object.
func flattenNodePoolManagementSlice(c *Client, i interface{}) []NodePoolManagement {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolManagement{}
	}

	if len(a) == 0 {
		return []NodePoolManagement{}
	}

	items := make([]NodePoolManagement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolManagement(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolManagement expands an instance of NodePoolManagement into a JSON
// request object.
func expandNodePoolManagement(c *Client, f *NodePoolManagement) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutoUpgrade; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgrade"] = v
	}
	if v := f.AutoRepair; !dcl.IsEmptyValueIndirect(v) {
		m["autoRepair"] = v
	}
	if v, err := expandNodePoolManagementUpgradeOptions(c, f.UpgradeOptions); err != nil {
		return nil, fmt.Errorf("error expanding UpgradeOptions into upgradeOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["upgradeOptions"] = v
	}

	return m, nil
}

// flattenNodePoolManagement flattens an instance of NodePoolManagement from a JSON
// response object.
func flattenNodePoolManagement(c *Client, i interface{}) *NodePoolManagement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolManagement{}
	r.AutoUpgrade = dcl.FlattenBool(m["autoUpgrade"])
	r.AutoRepair = dcl.FlattenBool(m["autoRepair"])
	r.UpgradeOptions = flattenNodePoolManagementUpgradeOptions(c, m["upgradeOptions"])

	return r
}

// expandNodePoolManagementUpgradeOptionsMap expands the contents of NodePoolManagementUpgradeOptions into a JSON
// request object.
func expandNodePoolManagementUpgradeOptionsMap(c *Client, f map[string]NodePoolManagementUpgradeOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolManagementUpgradeOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolManagementUpgradeOptionsSlice expands the contents of NodePoolManagementUpgradeOptions into a JSON
// request object.
func expandNodePoolManagementUpgradeOptionsSlice(c *Client, f []NodePoolManagementUpgradeOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolManagementUpgradeOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolManagementUpgradeOptionsMap flattens the contents of NodePoolManagementUpgradeOptions from a JSON
// response object.
func flattenNodePoolManagementUpgradeOptionsMap(c *Client, i interface{}) map[string]NodePoolManagementUpgradeOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolManagementUpgradeOptions{}
	}

	if len(a) == 0 {
		return map[string]NodePoolManagementUpgradeOptions{}
	}

	items := make(map[string]NodePoolManagementUpgradeOptions)
	for k, item := range a {
		items[k] = *flattenNodePoolManagementUpgradeOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolManagementUpgradeOptionsSlice flattens the contents of NodePoolManagementUpgradeOptions from a JSON
// response object.
func flattenNodePoolManagementUpgradeOptionsSlice(c *Client, i interface{}) []NodePoolManagementUpgradeOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolManagementUpgradeOptions{}
	}

	if len(a) == 0 {
		return []NodePoolManagementUpgradeOptions{}
	}

	items := make([]NodePoolManagementUpgradeOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolManagementUpgradeOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolManagementUpgradeOptions expands an instance of NodePoolManagementUpgradeOptions into a JSON
// request object.
func expandNodePoolManagementUpgradeOptions(c *Client, f *NodePoolManagementUpgradeOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutoUpgradeStartTime; !dcl.IsEmptyValueIndirect(v) {
		m["autoUpgradeStartTime"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenNodePoolManagementUpgradeOptions flattens an instance of NodePoolManagementUpgradeOptions from a JSON
// response object.
func flattenNodePoolManagementUpgradeOptions(c *Client, i interface{}) *NodePoolManagementUpgradeOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolManagementUpgradeOptions{}
	r.AutoUpgradeStartTime = dcl.FlattenString(m["autoUpgradeStartTime"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandNodePoolMaxPodsConstraintMap expands the contents of NodePoolMaxPodsConstraint into a JSON
// request object.
func expandNodePoolMaxPodsConstraintMap(c *Client, f map[string]NodePoolMaxPodsConstraint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolMaxPodsConstraintSlice expands the contents of NodePoolMaxPodsConstraint into a JSON
// request object.
func expandNodePoolMaxPodsConstraintSlice(c *Client, f []NodePoolMaxPodsConstraint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolMaxPodsConstraint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolMaxPodsConstraintMap flattens the contents of NodePoolMaxPodsConstraint from a JSON
// response object.
func flattenNodePoolMaxPodsConstraintMap(c *Client, i interface{}) map[string]NodePoolMaxPodsConstraint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return map[string]NodePoolMaxPodsConstraint{}
	}

	items := make(map[string]NodePoolMaxPodsConstraint)
	for k, item := range a {
		items[k] = *flattenNodePoolMaxPodsConstraint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolMaxPodsConstraintSlice flattens the contents of NodePoolMaxPodsConstraint from a JSON
// response object.
func flattenNodePoolMaxPodsConstraintSlice(c *Client, i interface{}) []NodePoolMaxPodsConstraint {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolMaxPodsConstraint{}
	}

	if len(a) == 0 {
		return []NodePoolMaxPodsConstraint{}
	}

	items := make([]NodePoolMaxPodsConstraint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolMaxPodsConstraint(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolMaxPodsConstraint expands an instance of NodePoolMaxPodsConstraint into a JSON
// request object.
func expandNodePoolMaxPodsConstraint(c *Client, f *NodePoolMaxPodsConstraint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxPodsPerNode; !dcl.IsEmptyValueIndirect(v) {
		m["maxPodsPerNode"] = v
	}

	return m, nil
}

// flattenNodePoolMaxPodsConstraint flattens an instance of NodePoolMaxPodsConstraint from a JSON
// response object.
func flattenNodePoolMaxPodsConstraint(c *Client, i interface{}) *NodePoolMaxPodsConstraint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolMaxPodsConstraint{}
	r.MaxPodsPerNode = dcl.FlattenInteger(m["maxPodsPerNode"])

	return r
}

// expandNodePoolConditionsMap expands the contents of NodePoolConditions into a JSON
// request object.
func expandNodePoolConditionsMap(c *Client, f map[string]NodePoolConditions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolConditions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolConditionsSlice expands the contents of NodePoolConditions into a JSON
// request object.
func expandNodePoolConditionsSlice(c *Client, f []NodePoolConditions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolConditions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolConditionsMap flattens the contents of NodePoolConditions from a JSON
// response object.
func flattenNodePoolConditionsMap(c *Client, i interface{}) map[string]NodePoolConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolConditions{}
	}

	if len(a) == 0 {
		return map[string]NodePoolConditions{}
	}

	items := make(map[string]NodePoolConditions)
	for k, item := range a {
		items[k] = *flattenNodePoolConditions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolConditionsSlice flattens the contents of NodePoolConditions from a JSON
// response object.
func flattenNodePoolConditionsSlice(c *Client, i interface{}) []NodePoolConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConditions{}
	}

	if len(a) == 0 {
		return []NodePoolConditions{}
	}

	items := make([]NodePoolConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConditions(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolConditions expands an instance of NodePoolConditions into a JSON
// request object.
func expandNodePoolConditions(c *Client, f *NodePoolConditions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}

	return m, nil
}

// flattenNodePoolConditions flattens an instance of NodePoolConditions from a JSON
// response object.
func flattenNodePoolConditions(c *Client, i interface{}) *NodePoolConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolConditions{}
	r.Code = flattenNodePoolConditionsCodeEnum(m["code"])
	r.Message = dcl.FlattenString(m["message"])

	return r
}

// expandNodePoolUpgradeSettingsMap expands the contents of NodePoolUpgradeSettings into a JSON
// request object.
func expandNodePoolUpgradeSettingsMap(c *Client, f map[string]NodePoolUpgradeSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodePoolUpgradeSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodePoolUpgradeSettingsSlice expands the contents of NodePoolUpgradeSettings into a JSON
// request object.
func expandNodePoolUpgradeSettingsSlice(c *Client, f []NodePoolUpgradeSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodePoolUpgradeSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodePoolUpgradeSettingsMap flattens the contents of NodePoolUpgradeSettings from a JSON
// response object.
func flattenNodePoolUpgradeSettingsMap(c *Client, i interface{}) map[string]NodePoolUpgradeSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodePoolUpgradeSettings{}
	}

	if len(a) == 0 {
		return map[string]NodePoolUpgradeSettings{}
	}

	items := make(map[string]NodePoolUpgradeSettings)
	for k, item := range a {
		items[k] = *flattenNodePoolUpgradeSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodePoolUpgradeSettingsSlice flattens the contents of NodePoolUpgradeSettings from a JSON
// response object.
func flattenNodePoolUpgradeSettingsSlice(c *Client, i interface{}) []NodePoolUpgradeSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolUpgradeSettings{}
	}

	if len(a) == 0 {
		return []NodePoolUpgradeSettings{}
	}

	items := make([]NodePoolUpgradeSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolUpgradeSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodePoolUpgradeSettings expands an instance of NodePoolUpgradeSettings into a JSON
// request object.
func expandNodePoolUpgradeSettings(c *Client, f *NodePoolUpgradeSettings) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxSurge; !dcl.IsEmptyValueIndirect(v) {
		m["maxSurge"] = v
	}
	if v := f.MaxUnavailable; !dcl.IsEmptyValueIndirect(v) {
		m["maxUnavailable"] = v
	}

	return m, nil
}

// flattenNodePoolUpgradeSettings flattens an instance of NodePoolUpgradeSettings from a JSON
// response object.
func flattenNodePoolUpgradeSettings(c *Client, i interface{}) *NodePoolUpgradeSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodePoolUpgradeSettings{}
	r.MaxSurge = dcl.FlattenInteger(m["maxSurge"])
	r.MaxUnavailable = dcl.FlattenInteger(m["maxUnavailable"])

	return r
}

// flattenNodePoolConfigSandboxConfigTypeEnumSlice flattens the contents of NodePoolConfigSandboxConfigTypeEnum from a JSON
// response object.
func flattenNodePoolConfigSandboxConfigTypeEnumSlice(c *Client, i interface{}) []NodePoolConfigSandboxConfigTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigSandboxConfigTypeEnum{}
	}

	if len(a) == 0 {
		return []NodePoolConfigSandboxConfigTypeEnum{}
	}

	items := make([]NodePoolConfigSandboxConfigTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigSandboxConfigTypeEnum(item.(interface{})))
	}

	return items
}

// flattenNodePoolConfigSandboxConfigTypeEnum asserts that an interface is a string, and returns a
// pointer to a *NodePoolConfigSandboxConfigTypeEnum with the same value as that string.
func flattenNodePoolConfigSandboxConfigTypeEnum(i interface{}) *NodePoolConfigSandboxConfigTypeEnum {
	s, ok := i.(string)
	if !ok {
		return NodePoolConfigSandboxConfigTypeEnumRef("")
	}

	return NodePoolConfigSandboxConfigTypeEnumRef(s)
}

// flattenNodePoolConfigReservationAffinityConsumeReservationTypeEnumSlice flattens the contents of NodePoolConfigReservationAffinityConsumeReservationTypeEnum from a JSON
// response object.
func flattenNodePoolConfigReservationAffinityConsumeReservationTypeEnumSlice(c *Client, i interface{}) []NodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	if len(a) == 0 {
		return []NodePoolConfigReservationAffinityConsumeReservationTypeEnum{}
	}

	items := make([]NodePoolConfigReservationAffinityConsumeReservationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConfigReservationAffinityConsumeReservationTypeEnum(item.(interface{})))
	}

	return items
}

// flattenNodePoolConfigReservationAffinityConsumeReservationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *NodePoolConfigReservationAffinityConsumeReservationTypeEnum with the same value as that string.
func flattenNodePoolConfigReservationAffinityConsumeReservationTypeEnum(i interface{}) *NodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return NodePoolConfigReservationAffinityConsumeReservationTypeEnumRef("")
	}

	return NodePoolConfigReservationAffinityConsumeReservationTypeEnumRef(s)
}

// flattenNodePoolConditionsCodeEnumSlice flattens the contents of NodePoolConditionsCodeEnum from a JSON
// response object.
func flattenNodePoolConditionsCodeEnumSlice(c *Client, i interface{}) []NodePoolConditionsCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NodePoolConditionsCodeEnum{}
	}

	if len(a) == 0 {
		return []NodePoolConditionsCodeEnum{}
	}

	items := make([]NodePoolConditionsCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodePoolConditionsCodeEnum(item.(interface{})))
	}

	return items
}

// flattenNodePoolConditionsCodeEnum asserts that an interface is a string, and returns a
// pointer to a *NodePoolConditionsCodeEnum with the same value as that string.
func flattenNodePoolConditionsCodeEnum(i interface{}) *NodePoolConditionsCodeEnum {
	s, ok := i.(string)
	if !ok {
		return NodePoolConditionsCodeEnumRef("")
	}

	return NodePoolConditionsCodeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *NodePool) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNodePool(b, c)
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
		if nr.Cluster == nil && ncr.Cluster == nil {
			c.Config.Logger.Info("Both Cluster fields null - considering equal.")
		} else if nr.Cluster == nil || ncr.Cluster == nil {
			c.Config.Logger.Info("Only one Cluster field is null - considering unequal.")
			return false
		} else if *nr.Cluster != *ncr.Cluster {
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
