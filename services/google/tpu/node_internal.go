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
package tpu

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

func (r *Node) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "acceleratorType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "tensorflowVersion"); err != nil {
		return err
	}
	if err := dcl.Required(r, "cidrBlock"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.CreateTime) {
		if err := r.CreateTime.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SchedulingConfig) {
		if err := r.SchedulingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NodeCreateTime) validate() error {
	return nil
}
func (r *NodeSchedulingConfig) validate() error {
	if err := dcl.Required(r, "preemptible"); err != nil {
		return err
	}
	return nil
}
func (r *NodeNetworkEndpoints) validate() error {
	return nil
}
func (r *NodeSymptoms) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CreateTime) {
		if err := r.CreateTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NodeSymptomsCreateTime) validate() error {
	return nil
}

func nodeGetURL(userBasePath string, r *Node) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/nodes/{{name}}", "https://tpu.googleapis.com/v1/", userBasePath, params), nil
}

func nodeListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/nodes", "https://tpu.googleapis.com/v1/", userBasePath, params), nil

}

func nodeCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/nodes?nodeId={{name}}", "https://tpu.googleapis.com/v1/", userBasePath, params), nil

}

func nodeDeleteURL(userBasePath string, r *Node) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/nodes/{{name}}", "https://tpu.googleapis.com/v1/", userBasePath, params), nil
}

// nodeApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type nodeApiOperation interface {
	do(context.Context, *Node, *Client) error
}

// newUpdateNodeReimageNodeRequest creates a request for an
// Node resource's ReimageNode update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNodeReimageNodeRequest(ctx context.Context, f *Node, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.TensorflowVersion; !dcl.IsEmptyValueIndirect(v) {
		req["tensorflowVersion"] = v
	}
	return req, nil
}

// marshalUpdateNodeReimageNodeRequest converts the update into
// the final JSON request body.
func marshalUpdateNodeReimageNodeRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNodeReimageNodeOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNodeReimageNodeOperation) do(ctx context.Context, r *Node, c *Client) error {
	_, err := c.GetNode(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "ReimageNode")
	if err != nil {
		return err
	}

	req, err := newUpdateNodeReimageNodeRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNodeReimageNodeRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://tpu.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listNodeRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := nodeListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NodeMaxPage {
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

type listNodeOperation struct {
	Nodes []map[string]interface{} `json:"nodes"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNode(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Node, string, error) {
	b, err := c.listNodeRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNodeOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Node
	for _, v := range m.Nodes {
		res := flattenNode(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNode(ctx context.Context, f func(*Node) bool, resources []*Node) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNode(ctx, res)
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

type deleteNodeOperation struct{}

func (op *deleteNodeOperation) do(ctx context.Context, r *Node, c *Client) error {

	_, err := c.GetNode(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Node not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetNode checking for existence. error: %v", err)
		return err
	}

	u, err := nodeDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://tpu.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetNode(ctx, r.urlNormalized())
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
type createNodeOperation struct {
	response map[string]interface{}
}

func (op *createNodeOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNodeOperation) do(ctx context.Context, r *Node, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := nodeCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://tpu.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetNode(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNodeRaw(ctx context.Context, r *Node) ([]byte, error) {

	u, err := nodeGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) nodeDiffsForRawDesired(ctx context.Context, rawDesired *Node, opts ...dcl.ApplyOption) (initial, desired *Node, diffs []nodeDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Node
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Node); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Node, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNode(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Node resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Node resource: %v", err)
		}
		c.Config.Logger.Info("Found that Node resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNodeDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Node: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Node: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNodeInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Node: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNodeDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Node: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNode(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNodeInitialState(rawInitial, rawDesired *Node) (*Node, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNodeDesiredState(rawDesired, rawInitial *Node, opts ...dcl.ApplyOption) (*Node, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.CreateTime = canonicalizeNodeCreateTime(rawDesired.CreateTime, nil, opts...)
		rawDesired.SchedulingConfig = canonicalizeNodeSchedulingConfig(rawDesired.SchedulingConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.AcceleratorType, rawInitial.AcceleratorType) {
		rawDesired.AcceleratorType = rawInitial.AcceleratorType
	}
	if dcl.StringCanonicalize(rawDesired.TensorflowVersion, rawInitial.TensorflowVersion) {
		rawDesired.TensorflowVersion = rawInitial.TensorflowVersion
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.StringCanonicalize(rawDesired.CidrBlock, rawInitial.CidrBlock) {
		rawDesired.CidrBlock = rawInitial.CidrBlock
	}
	rawDesired.SchedulingConfig = canonicalizeNodeSchedulingConfig(rawDesired.SchedulingConfig, rawInitial.SchedulingConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.BoolCanonicalize(rawDesired.UseServiceNetworking, rawInitial.UseServiceNetworking) {
		rawDesired.UseServiceNetworking = rawInitial.UseServiceNetworking
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeNodeNewState(c *Client, rawNew, rawDesired *Node) (*Node, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.AcceleratorType) && dcl.IsEmptyValueIndirect(rawDesired.AcceleratorType) {
		rawNew.AcceleratorType = rawDesired.AcceleratorType
	} else {
		if dcl.StringCanonicalize(rawDesired.AcceleratorType, rawNew.AcceleratorType) {
			rawNew.AcceleratorType = rawDesired.AcceleratorType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddress) && dcl.IsEmptyValueIndirect(rawDesired.IPAddress) {
		rawNew.IPAddress = rawDesired.IPAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.IPAddress, rawNew.IPAddress) {
			rawNew.IPAddress = rawDesired.IPAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
		if dcl.StringCanonicalize(rawDesired.Port, rawNew.Port) {
			rawNew.Port = rawDesired.Port
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthDescription) && dcl.IsEmptyValueIndirect(rawDesired.HealthDescription) {
		rawNew.HealthDescription = rawDesired.HealthDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.HealthDescription, rawNew.HealthDescription) {
			rawNew.HealthDescription = rawDesired.HealthDescription
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TensorflowVersion) && dcl.IsEmptyValueIndirect(rawDesired.TensorflowVersion) {
		rawNew.TensorflowVersion = rawDesired.TensorflowVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.TensorflowVersion, rawNew.TensorflowVersion) {
			rawNew.TensorflowVersion = rawDesired.TensorflowVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CidrBlock) && dcl.IsEmptyValueIndirect(rawDesired.CidrBlock) {
		rawNew.CidrBlock = rawDesired.CidrBlock
	} else {
		if dcl.StringCanonicalize(rawDesired.CidrBlock, rawNew.CidrBlock) {
			rawNew.CidrBlock = rawDesired.CidrBlock
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccount) {
		rawNew.ServiceAccount = rawDesired.ServiceAccount
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceAccount, rawNew.ServiceAccount) {
			rawNew.ServiceAccount = rawDesired.ServiceAccount
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
		rawNew.CreateTime = canonicalizeNewNodeCreateTime(c, rawDesired.CreateTime, rawNew.CreateTime)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SchedulingConfig) && dcl.IsEmptyValueIndirect(rawDesired.SchedulingConfig) {
		rawNew.SchedulingConfig = rawDesired.SchedulingConfig
	} else {
		rawNew.SchedulingConfig = canonicalizeNewNodeSchedulingConfig(c, rawDesired.SchedulingConfig, rawNew.SchedulingConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkEndpoints) && dcl.IsEmptyValueIndirect(rawDesired.NetworkEndpoints) {
		rawNew.NetworkEndpoints = rawDesired.NetworkEndpoints
	} else {
		rawNew.NetworkEndpoints = canonicalizeNewNodeNetworkEndpointsSlice(c, rawDesired.NetworkEndpoints, rawNew.NetworkEndpoints)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Health) && dcl.IsEmptyValueIndirect(rawDesired.Health) {
		rawNew.Health = rawDesired.Health
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UseServiceNetworking) && dcl.IsEmptyValueIndirect(rawDesired.UseServiceNetworking) {
		rawNew.UseServiceNetworking = rawDesired.UseServiceNetworking
	} else {
		if dcl.BoolCanonicalize(rawDesired.UseServiceNetworking, rawNew.UseServiceNetworking) {
			rawNew.UseServiceNetworking = rawDesired.UseServiceNetworking
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Symptoms) && dcl.IsEmptyValueIndirect(rawDesired.Symptoms) {
		rawNew.Symptoms = rawDesired.Symptoms
	} else {
		rawNew.Symptoms = canonicalizeNewNodeSymptomsSlice(c, rawDesired.Symptoms, rawNew.Symptoms)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeNodeCreateTime(des, initial *NodeCreateTime, opts ...dcl.ApplyOption) *NodeCreateTime {
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

func canonicalizeNewNodeCreateTime(c *Client, des, nw *NodeCreateTime) *NodeCreateTime {
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

func canonicalizeNewNodeCreateTimeSet(c *Client, des, nw []NodeCreateTime) []NodeCreateTime {
	if des == nil {
		return nw
	}
	var reorderedNew []NodeCreateTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNodeCreateTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNodeCreateTimeSlice(c *Client, des, nw []NodeCreateTime) []NodeCreateTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NodeCreateTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodeCreateTime(c, &d, &n))
	}

	return items
}

func canonicalizeNodeSchedulingConfig(des, initial *NodeSchedulingConfig, opts ...dcl.ApplyOption) *NodeSchedulingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.Preemptible, initial.Preemptible) || dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}
	if dcl.BoolCanonicalize(des.Reserved, initial.Reserved) || dcl.IsZeroValue(des.Reserved) {
		des.Reserved = initial.Reserved
	}

	return des
}

func canonicalizeNewNodeSchedulingConfig(c *Client, des, nw *NodeSchedulingConfig) *NodeSchedulingConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Preemptible, nw.Preemptible) {
		nw.Preemptible = des.Preemptible
	}
	if dcl.BoolCanonicalize(des.Reserved, nw.Reserved) {
		nw.Reserved = des.Reserved
	}

	return nw
}

func canonicalizeNewNodeSchedulingConfigSet(c *Client, des, nw []NodeSchedulingConfig) []NodeSchedulingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []NodeSchedulingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNodeSchedulingConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNodeSchedulingConfigSlice(c *Client, des, nw []NodeSchedulingConfig) []NodeSchedulingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NodeSchedulingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodeSchedulingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeNodeNetworkEndpoints(des, initial *NodeNetworkEndpoints, opts ...dcl.ApplyOption) *NodeNetworkEndpoints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	return des
}

func canonicalizeNewNodeNetworkEndpoints(c *Client, des, nw *NodeNetworkEndpoints) *NodeNetworkEndpoints {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPAddress, nw.IPAddress) {
		nw.IPAddress = des.IPAddress
	}
	if dcl.IsZeroValue(nw.Port) {
		nw.Port = des.Port
	}

	return nw
}

func canonicalizeNewNodeNetworkEndpointsSet(c *Client, des, nw []NodeNetworkEndpoints) []NodeNetworkEndpoints {
	if des == nil {
		return nw
	}
	var reorderedNew []NodeNetworkEndpoints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNodeNetworkEndpointsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNodeNetworkEndpointsSlice(c *Client, des, nw []NodeNetworkEndpoints) []NodeNetworkEndpoints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NodeNetworkEndpoints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodeNetworkEndpoints(c, &d, &n))
	}

	return items
}

func canonicalizeNodeSymptoms(des, initial *NodeSymptoms, opts ...dcl.ApplyOption) *NodeSymptoms {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.CreateTime = canonicalizeNodeSymptomsCreateTime(des.CreateTime, initial.CreateTime, opts...)
	if dcl.IsZeroValue(des.SymptomType) {
		des.SymptomType = initial.SymptomType
	}
	if dcl.StringCanonicalize(des.Details, initial.Details) || dcl.IsZeroValue(des.Details) {
		des.Details = initial.Details
	}
	if dcl.StringCanonicalize(des.WorkerId, initial.WorkerId) || dcl.IsZeroValue(des.WorkerId) {
		des.WorkerId = initial.WorkerId
	}

	return des
}

func canonicalizeNewNodeSymptoms(c *Client, des, nw *NodeSymptoms) *NodeSymptoms {
	if des == nil || nw == nil {
		return nw
	}

	nw.CreateTime = canonicalizeNewNodeSymptomsCreateTime(c, des.CreateTime, nw.CreateTime)
	if dcl.IsZeroValue(nw.SymptomType) {
		nw.SymptomType = des.SymptomType
	}
	if dcl.StringCanonicalize(des.Details, nw.Details) {
		nw.Details = des.Details
	}
	if dcl.StringCanonicalize(des.WorkerId, nw.WorkerId) {
		nw.WorkerId = des.WorkerId
	}

	return nw
}

func canonicalizeNewNodeSymptomsSet(c *Client, des, nw []NodeSymptoms) []NodeSymptoms {
	if des == nil {
		return nw
	}
	var reorderedNew []NodeSymptoms
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNodeSymptomsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNodeSymptomsSlice(c *Client, des, nw []NodeSymptoms) []NodeSymptoms {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NodeSymptoms
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodeSymptoms(c, &d, &n))
	}

	return items
}

func canonicalizeNodeSymptomsCreateTime(des, initial *NodeSymptomsCreateTime, opts ...dcl.ApplyOption) *NodeSymptomsCreateTime {
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

func canonicalizeNewNodeSymptomsCreateTime(c *Client, des, nw *NodeSymptomsCreateTime) *NodeSymptomsCreateTime {
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

func canonicalizeNewNodeSymptomsCreateTimeSet(c *Client, des, nw []NodeSymptomsCreateTime) []NodeSymptomsCreateTime {
	if des == nil {
		return nw
	}
	var reorderedNew []NodeSymptomsCreateTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareNodeSymptomsCreateTimeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewNodeSymptomsCreateTimeSlice(c *Client, des, nw []NodeSymptomsCreateTime) []NodeSymptomsCreateTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NodeSymptomsCreateTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNodeSymptomsCreateTime(c, &d, &n))
	}

	return items
}

type nodeDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         nodeApiOperation
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
func diffNode(c *Client, desired, actual *Node, opts ...dcl.ApplyOption) ([]nodeDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []nodeDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.AcceleratorType, actual.AcceleratorType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcceleratorType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.IPAddress, actual.IPAddress, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.HealthDescription, actual.HealthDescription, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HealthDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.TensorflowVersion, actual.TensorflowVersion, dcl.Info{OperationSelector: dcl.TriggersOperation("updateNodeReimageNodeOperation")}, fn.AddNest("TensorflowVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.CidrBlock, actual.CidrBlock, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, ObjectFunction: compareNodeCreateTimeNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SchedulingConfig, actual.SchedulingConfig, dcl.Info{ObjectFunction: compareNodeSchedulingConfigNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SchedulingConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.NetworkEndpoints, actual.NetworkEndpoints, dcl.Info{OutputOnly: true, ObjectFunction: compareNodeNetworkEndpointsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkEndpoints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Health, actual.Health, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Health")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.UseServiceNetworking, actual.UseServiceNetworking, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UseServiceNetworking")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Symptoms, actual.Symptoms, dcl.Info{OutputOnly: true, ObjectFunction: compareNodeSymptomsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Symptoms")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToNodeDiff(ds, opts...)
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
	var deduped []nodeDiff
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
func compareNodeCreateTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NodeCreateTime)
	if !ok {
		desiredNotPointer, ok := d.(NodeCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeCreateTime or *NodeCreateTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NodeCreateTime)
	if !ok {
		actualNotPointer, ok := a.(NodeCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeCreateTime", a)
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

func compareNodeSchedulingConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NodeSchedulingConfig)
	if !ok {
		desiredNotPointer, ok := d.(NodeSchedulingConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeSchedulingConfig or *NodeSchedulingConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NodeSchedulingConfig)
	if !ok {
		actualNotPointer, ok := a.(NodeSchedulingConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeSchedulingConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Preemptible, actual.Preemptible, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Preemptible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reserved, actual.Reserved, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Reserved")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNodeNetworkEndpointsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NodeNetworkEndpoints)
	if !ok {
		desiredNotPointer, ok := d.(NodeNetworkEndpoints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeNetworkEndpoints or *NodeNetworkEndpoints", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NodeNetworkEndpoints)
	if !ok {
		actualNotPointer, ok := a.(NodeNetworkEndpoints)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeNetworkEndpoints", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IPAddress, actual.IPAddress, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNodeSymptomsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NodeSymptoms)
	if !ok {
		desiredNotPointer, ok := d.(NodeSymptoms)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeSymptoms or *NodeSymptoms", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NodeSymptoms)
	if !ok {
		actualNotPointer, ok := a.(NodeSymptoms)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeSymptoms", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{ObjectFunction: compareNodeSymptomsCreateTimeNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SymptomType, actual.SymptomType, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SymptomType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkerId, actual.WorkerId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkerId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNodeSymptomsCreateTimeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NodeSymptomsCreateTime)
	if !ok {
		desiredNotPointer, ok := d.(NodeSymptomsCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeSymptomsCreateTime or *NodeSymptomsCreateTime", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NodeSymptomsCreateTime)
	if !ok {
		actualNotPointer, ok := a.(NodeSymptomsCreateTime)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NodeSymptomsCreateTime", a)
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
func (r *Node) urlNormalized() *Node {
	normalized := dcl.Copy(*r).(Node)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AcceleratorType = dcl.SelfLinkToName(r.AcceleratorType)
	normalized.IPAddress = dcl.SelfLinkToName(r.IPAddress)
	normalized.Port = dcl.SelfLinkToName(r.Port)
	normalized.HealthDescription = dcl.SelfLinkToName(r.HealthDescription)
	normalized.TensorflowVersion = dcl.SelfLinkToName(r.TensorflowVersion)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.CidrBlock = dcl.SelfLinkToName(r.CidrBlock)
	normalized.ServiceAccount = dcl.SelfLinkToName(r.ServiceAccount)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Node) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Node) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Node) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Node) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "ReimageNode" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/nodes/{{name}}:reimage", "https://tpu.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Node resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Node) marshal(c *Client) ([]byte, error) {
	m, err := expandNode(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Node: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNode decodes JSON responses into the Node resource schema.
func unmarshalNode(b []byte, c *Client) (*Node, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNode(m, c)
}

func unmarshalMapNode(m map[string]interface{}, c *Client) (*Node, error) {

	return flattenNode(c, m), nil
}

// expandNode expands Node into a JSON request object.
func expandNode(c *Client, f *Node) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.AcceleratorType; !dcl.IsEmptyValueIndirect(v) {
		m["acceleratorType"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.HealthDescription; !dcl.IsEmptyValueIndirect(v) {
		m["healthDescription"] = v
	}
	if v := f.TensorflowVersion; !dcl.IsEmptyValueIndirect(v) {
		m["tensorflowVersion"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["cidrBlock"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandNodeCreateTime(c, f.CreateTime); err != nil {
		return nil, fmt.Errorf("error expanding CreateTime into createTime: %w", err)
	} else if v != nil {
		m["createTime"] = v
	}
	if v, err := expandNodeSchedulingConfig(c, f.SchedulingConfig); err != nil {
		return nil, fmt.Errorf("error expanding SchedulingConfig into schedulingConfig: %w", err)
	} else if v != nil {
		m["schedulingConfig"] = v
	}
	if v, err := expandNodeNetworkEndpointsSlice(c, f.NetworkEndpoints); err != nil {
		return nil, fmt.Errorf("error expanding NetworkEndpoints into networkEndpoints: %w", err)
	} else if v != nil {
		m["networkEndpoints"] = v
	}
	if v := f.Health; !dcl.IsEmptyValueIndirect(v) {
		m["health"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.UseServiceNetworking; !dcl.IsEmptyValueIndirect(v) {
		m["useServiceNetworking"] = v
	}
	if v, err := expandNodeSymptomsSlice(c, f.Symptoms); err != nil {
		return nil, fmt.Errorf("error expanding Symptoms into symptoms: %w", err)
	} else if v != nil {
		m["symptoms"] = v
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

// flattenNode flattens Node from a JSON request object into the
// Node type.
func flattenNode(c *Client, i interface{}) *Node {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Node{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.AcceleratorType = dcl.FlattenString(m["acceleratorType"])
	res.IPAddress = dcl.FlattenString(m["ipAddress"])
	res.Port = dcl.FlattenString(m["port"])
	res.State = flattenNodeStateEnum(m["state"])
	res.HealthDescription = dcl.FlattenString(m["healthDescription"])
	res.TensorflowVersion = dcl.FlattenString(m["tensorflowVersion"])
	res.Network = dcl.FlattenString(m["network"])
	res.CidrBlock = dcl.FlattenString(m["cidrBlock"])
	res.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	res.CreateTime = flattenNodeCreateTime(c, m["createTime"])
	res.SchedulingConfig = flattenNodeSchedulingConfig(c, m["schedulingConfig"])
	res.NetworkEndpoints = flattenNodeNetworkEndpointsSlice(c, m["networkEndpoints"])
	res.Health = flattenNodeHealthEnum(m["health"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.UseServiceNetworking = dcl.FlattenBool(m["useServiceNetworking"])
	res.Symptoms = flattenNodeSymptomsSlice(c, m["symptoms"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandNodeCreateTimeMap expands the contents of NodeCreateTime into a JSON
// request object.
func expandNodeCreateTimeMap(c *Client, f map[string]NodeCreateTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodeCreateTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodeCreateTimeSlice expands the contents of NodeCreateTime into a JSON
// request object.
func expandNodeCreateTimeSlice(c *Client, f []NodeCreateTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodeCreateTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodeCreateTimeMap flattens the contents of NodeCreateTime from a JSON
// response object.
func flattenNodeCreateTimeMap(c *Client, i interface{}) map[string]NodeCreateTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodeCreateTime{}
	}

	if len(a) == 0 {
		return map[string]NodeCreateTime{}
	}

	items := make(map[string]NodeCreateTime)
	for k, item := range a {
		items[k] = *flattenNodeCreateTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodeCreateTimeSlice flattens the contents of NodeCreateTime from a JSON
// response object.
func flattenNodeCreateTimeSlice(c *Client, i interface{}) []NodeCreateTime {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeCreateTime{}
	}

	if len(a) == 0 {
		return []NodeCreateTime{}
	}

	items := make([]NodeCreateTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeCreateTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodeCreateTime expands an instance of NodeCreateTime into a JSON
// request object.
func expandNodeCreateTime(c *Client, f *NodeCreateTime) (map[string]interface{}, error) {
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

// flattenNodeCreateTime flattens an instance of NodeCreateTime from a JSON
// response object.
func flattenNodeCreateTime(c *Client, i interface{}) *NodeCreateTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodeCreateTime{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNodeCreateTime
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandNodeSchedulingConfigMap expands the contents of NodeSchedulingConfig into a JSON
// request object.
func expandNodeSchedulingConfigMap(c *Client, f map[string]NodeSchedulingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodeSchedulingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodeSchedulingConfigSlice expands the contents of NodeSchedulingConfig into a JSON
// request object.
func expandNodeSchedulingConfigSlice(c *Client, f []NodeSchedulingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodeSchedulingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodeSchedulingConfigMap flattens the contents of NodeSchedulingConfig from a JSON
// response object.
func flattenNodeSchedulingConfigMap(c *Client, i interface{}) map[string]NodeSchedulingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodeSchedulingConfig{}
	}

	if len(a) == 0 {
		return map[string]NodeSchedulingConfig{}
	}

	items := make(map[string]NodeSchedulingConfig)
	for k, item := range a {
		items[k] = *flattenNodeSchedulingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodeSchedulingConfigSlice flattens the contents of NodeSchedulingConfig from a JSON
// response object.
func flattenNodeSchedulingConfigSlice(c *Client, i interface{}) []NodeSchedulingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeSchedulingConfig{}
	}

	if len(a) == 0 {
		return []NodeSchedulingConfig{}
	}

	items := make([]NodeSchedulingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeSchedulingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodeSchedulingConfig expands an instance of NodeSchedulingConfig into a JSON
// request object.
func expandNodeSchedulingConfig(c *Client, f *NodeSchedulingConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Preemptible; !dcl.IsEmptyValueIndirect(v) {
		m["preemptible"] = v
	}
	if v := f.Reserved; !dcl.IsEmptyValueIndirect(v) {
		m["reserved"] = v
	}

	return m, nil
}

// flattenNodeSchedulingConfig flattens an instance of NodeSchedulingConfig from a JSON
// response object.
func flattenNodeSchedulingConfig(c *Client, i interface{}) *NodeSchedulingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodeSchedulingConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNodeSchedulingConfig
	}
	r.Preemptible = dcl.FlattenBool(m["preemptible"])
	r.Reserved = dcl.FlattenBool(m["reserved"])

	return r
}

// expandNodeNetworkEndpointsMap expands the contents of NodeNetworkEndpoints into a JSON
// request object.
func expandNodeNetworkEndpointsMap(c *Client, f map[string]NodeNetworkEndpoints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodeNetworkEndpoints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodeNetworkEndpointsSlice expands the contents of NodeNetworkEndpoints into a JSON
// request object.
func expandNodeNetworkEndpointsSlice(c *Client, f []NodeNetworkEndpoints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodeNetworkEndpoints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodeNetworkEndpointsMap flattens the contents of NodeNetworkEndpoints from a JSON
// response object.
func flattenNodeNetworkEndpointsMap(c *Client, i interface{}) map[string]NodeNetworkEndpoints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodeNetworkEndpoints{}
	}

	if len(a) == 0 {
		return map[string]NodeNetworkEndpoints{}
	}

	items := make(map[string]NodeNetworkEndpoints)
	for k, item := range a {
		items[k] = *flattenNodeNetworkEndpoints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodeNetworkEndpointsSlice flattens the contents of NodeNetworkEndpoints from a JSON
// response object.
func flattenNodeNetworkEndpointsSlice(c *Client, i interface{}) []NodeNetworkEndpoints {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeNetworkEndpoints{}
	}

	if len(a) == 0 {
		return []NodeNetworkEndpoints{}
	}

	items := make([]NodeNetworkEndpoints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeNetworkEndpoints(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodeNetworkEndpoints expands an instance of NodeNetworkEndpoints into a JSON
// request object.
func expandNodeNetworkEndpoints(c *Client, f *NodeNetworkEndpoints) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}

	return m, nil
}

// flattenNodeNetworkEndpoints flattens an instance of NodeNetworkEndpoints from a JSON
// response object.
func flattenNodeNetworkEndpoints(c *Client, i interface{}) *NodeNetworkEndpoints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodeNetworkEndpoints{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNodeNetworkEndpoints
	}
	r.IPAddress = dcl.FlattenString(m["ipAddress"])
	r.Port = dcl.FlattenInteger(m["port"])

	return r
}

// expandNodeSymptomsMap expands the contents of NodeSymptoms into a JSON
// request object.
func expandNodeSymptomsMap(c *Client, f map[string]NodeSymptoms) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodeSymptoms(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodeSymptomsSlice expands the contents of NodeSymptoms into a JSON
// request object.
func expandNodeSymptomsSlice(c *Client, f []NodeSymptoms) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodeSymptoms(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodeSymptomsMap flattens the contents of NodeSymptoms from a JSON
// response object.
func flattenNodeSymptomsMap(c *Client, i interface{}) map[string]NodeSymptoms {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodeSymptoms{}
	}

	if len(a) == 0 {
		return map[string]NodeSymptoms{}
	}

	items := make(map[string]NodeSymptoms)
	for k, item := range a {
		items[k] = *flattenNodeSymptoms(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodeSymptomsSlice flattens the contents of NodeSymptoms from a JSON
// response object.
func flattenNodeSymptomsSlice(c *Client, i interface{}) []NodeSymptoms {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeSymptoms{}
	}

	if len(a) == 0 {
		return []NodeSymptoms{}
	}

	items := make([]NodeSymptoms, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeSymptoms(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodeSymptoms expands an instance of NodeSymptoms into a JSON
// request object.
func expandNodeSymptoms(c *Client, f *NodeSymptoms) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandNodeSymptomsCreateTime(c, f.CreateTime); err != nil {
		return nil, fmt.Errorf("error expanding CreateTime into createTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.SymptomType; !dcl.IsEmptyValueIndirect(v) {
		m["symptomType"] = v
	}
	if v := f.Details; !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}
	if v := f.WorkerId; !dcl.IsEmptyValueIndirect(v) {
		m["workerId"] = v
	}

	return m, nil
}

// flattenNodeSymptoms flattens an instance of NodeSymptoms from a JSON
// response object.
func flattenNodeSymptoms(c *Client, i interface{}) *NodeSymptoms {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodeSymptoms{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNodeSymptoms
	}
	r.CreateTime = flattenNodeSymptomsCreateTime(c, m["createTime"])
	r.SymptomType = flattenNodeSymptomsSymptomTypeEnum(m["symptomType"])
	r.Details = dcl.FlattenString(m["details"])
	r.WorkerId = dcl.FlattenString(m["workerId"])

	return r
}

// expandNodeSymptomsCreateTimeMap expands the contents of NodeSymptomsCreateTime into a JSON
// request object.
func expandNodeSymptomsCreateTimeMap(c *Client, f map[string]NodeSymptomsCreateTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNodeSymptomsCreateTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNodeSymptomsCreateTimeSlice expands the contents of NodeSymptomsCreateTime into a JSON
// request object.
func expandNodeSymptomsCreateTimeSlice(c *Client, f []NodeSymptomsCreateTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNodeSymptomsCreateTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNodeSymptomsCreateTimeMap flattens the contents of NodeSymptomsCreateTime from a JSON
// response object.
func flattenNodeSymptomsCreateTimeMap(c *Client, i interface{}) map[string]NodeSymptomsCreateTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NodeSymptomsCreateTime{}
	}

	if len(a) == 0 {
		return map[string]NodeSymptomsCreateTime{}
	}

	items := make(map[string]NodeSymptomsCreateTime)
	for k, item := range a {
		items[k] = *flattenNodeSymptomsCreateTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNodeSymptomsCreateTimeSlice flattens the contents of NodeSymptomsCreateTime from a JSON
// response object.
func flattenNodeSymptomsCreateTimeSlice(c *Client, i interface{}) []NodeSymptomsCreateTime {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeSymptomsCreateTime{}
	}

	if len(a) == 0 {
		return []NodeSymptomsCreateTime{}
	}

	items := make([]NodeSymptomsCreateTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeSymptomsCreateTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandNodeSymptomsCreateTime expands an instance of NodeSymptomsCreateTime into a JSON
// request object.
func expandNodeSymptomsCreateTime(c *Client, f *NodeSymptomsCreateTime) (map[string]interface{}, error) {
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

// flattenNodeSymptomsCreateTime flattens an instance of NodeSymptomsCreateTime from a JSON
// response object.
func flattenNodeSymptomsCreateTime(c *Client, i interface{}) *NodeSymptomsCreateTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NodeSymptomsCreateTime{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNodeSymptomsCreateTime
	}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// flattenNodeStateEnumSlice flattens the contents of NodeStateEnum from a JSON
// response object.
func flattenNodeStateEnumSlice(c *Client, i interface{}) []NodeStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeStateEnum{}
	}

	if len(a) == 0 {
		return []NodeStateEnum{}
	}

	items := make([]NodeStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeStateEnum(item.(interface{})))
	}

	return items
}

// flattenNodeStateEnum asserts that an interface is a string, and returns a
// pointer to a *NodeStateEnum with the same value as that string.
func flattenNodeStateEnum(i interface{}) *NodeStateEnum {
	s, ok := i.(string)
	if !ok {
		return NodeStateEnumRef("")
	}

	return NodeStateEnumRef(s)
}

// flattenNodeHealthEnumSlice flattens the contents of NodeHealthEnum from a JSON
// response object.
func flattenNodeHealthEnumSlice(c *Client, i interface{}) []NodeHealthEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeHealthEnum{}
	}

	if len(a) == 0 {
		return []NodeHealthEnum{}
	}

	items := make([]NodeHealthEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeHealthEnum(item.(interface{})))
	}

	return items
}

// flattenNodeHealthEnum asserts that an interface is a string, and returns a
// pointer to a *NodeHealthEnum with the same value as that string.
func flattenNodeHealthEnum(i interface{}) *NodeHealthEnum {
	s, ok := i.(string)
	if !ok {
		return NodeHealthEnumRef("")
	}

	return NodeHealthEnumRef(s)
}

// flattenNodeSymptomsSymptomTypeEnumSlice flattens the contents of NodeSymptomsSymptomTypeEnum from a JSON
// response object.
func flattenNodeSymptomsSymptomTypeEnumSlice(c *Client, i interface{}) []NodeSymptomsSymptomTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NodeSymptomsSymptomTypeEnum{}
	}

	if len(a) == 0 {
		return []NodeSymptomsSymptomTypeEnum{}
	}

	items := make([]NodeSymptomsSymptomTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNodeSymptomsSymptomTypeEnum(item.(interface{})))
	}

	return items
}

// flattenNodeSymptomsSymptomTypeEnum asserts that an interface is a string, and returns a
// pointer to a *NodeSymptomsSymptomTypeEnum with the same value as that string.
func flattenNodeSymptomsSymptomTypeEnum(i interface{}) *NodeSymptomsSymptomTypeEnum {
	s, ok := i.(string)
	if !ok {
		return NodeSymptomsSymptomTypeEnumRef("")
	}

	return NodeSymptomsSymptomTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Node) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNode(b, c)
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

func convertFieldDiffToNodeDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]nodeDiff, error) {
	var diffs []nodeDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := nodeDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTonodeApiOperation(op, opts...)
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

func convertOpNameTonodeApiOperation(op string, opts ...dcl.ApplyOption) (nodeApiOperation, error) {
	switch op {

	case "updateNodeReimageNodeOperation":
		return &updateNodeReimageNodeOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
