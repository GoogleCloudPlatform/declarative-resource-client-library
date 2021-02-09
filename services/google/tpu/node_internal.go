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
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.Retry)
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
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listNodeOperation struct {
	Items []map[string]interface{} `json:"items"`
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
	for _, v := range m.Items {
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
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
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
	_, err = c.GetNode(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createNodeOperation struct{}

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
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
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

	if _, err := c.GetNode(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getNodeRaw(ctx context.Context, r *Node) ([]byte, error) {

	u, err := nodeGetURL(c.Config.BasePath, r.urlNormalized())
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Node); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected Node, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.CreateTime = canonicalizeNodeCreateTime(rawDesired.CreateTime, nil, opts...)
		rawDesired.SchedulingConfig = canonicalizeNodeSchedulingConfig(rawDesired.SchedulingConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.AcceleratorType) {
		rawDesired.AcceleratorType = rawInitial.AcceleratorType
	}
	if dcl.IsZeroValue(rawDesired.IPAddress) {
		rawDesired.IPAddress = rawInitial.IPAddress
	}
	if dcl.IsZeroValue(rawDesired.Port) {
		rawDesired.Port = rawInitial.Port
	}
	if dcl.IsZeroValue(rawDesired.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.IsZeroValue(rawDesired.HealthDescription) {
		rawDesired.HealthDescription = rawInitial.HealthDescription
	}
	if dcl.IsZeroValue(rawDesired.TensorflowVersion) {
		rawDesired.TensorflowVersion = rawInitial.TensorflowVersion
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.CidrBlock) {
		rawDesired.CidrBlock = rawInitial.CidrBlock
	}
	if dcl.IsZeroValue(rawDesired.ServiceAccount) {
		rawDesired.ServiceAccount = rawInitial.ServiceAccount
	}
	rawDesired.CreateTime = canonicalizeNodeCreateTime(rawDesired.CreateTime, rawInitial.CreateTime, opts...)
	rawDesired.SchedulingConfig = canonicalizeNodeSchedulingConfig(rawDesired.SchedulingConfig, rawInitial.SchedulingConfig, opts...)
	if dcl.IsZeroValue(rawDesired.NetworkEndpoints) {
		rawDesired.NetworkEndpoints = rawInitial.NetworkEndpoints
	}
	if dcl.IsZeroValue(rawDesired.Health) {
		rawDesired.Health = rawInitial.Health
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.UseServiceNetworking) {
		rawDesired.UseServiceNetworking = rawInitial.UseServiceNetworking
	}
	if dcl.IsZeroValue(rawDesired.Symptoms) {
		rawDesired.Symptoms = rawInitial.Symptoms
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AcceleratorType) && dcl.IsEmptyValueIndirect(rawDesired.AcceleratorType) {
		rawNew.AcceleratorType = rawDesired.AcceleratorType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddress) && dcl.IsEmptyValueIndirect(rawDesired.IPAddress) {
		rawNew.IPAddress = rawDesired.IPAddress
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Port) && dcl.IsEmptyValueIndirect(rawDesired.Port) {
		rawNew.Port = rawDesired.Port
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.HealthDescription) && dcl.IsEmptyValueIndirect(rawDesired.HealthDescription) {
		rawNew.HealthDescription = rawDesired.HealthDescription
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TensorflowVersion) && dcl.IsEmptyValueIndirect(rawDesired.TensorflowVersion) {
		rawNew.TensorflowVersion = rawDesired.TensorflowVersion
	} else {
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccount) {
		rawNew.ServiceAccount = rawDesired.ServiceAccount
	} else {
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Symptoms) && dcl.IsEmptyValueIndirect(rawDesired.Symptoms) {
		rawNew.Symptoms = rawDesired.Symptoms
	} else {
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

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Node)
		_ = r
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
			if !compareNodeCreateTime(c, &d, &n) {
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

func canonicalizeNodeSchedulingConfig(des, initial *NodeSchedulingConfig, opts ...dcl.ApplyOption) *NodeSchedulingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Node)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Preemptible) {
		des.Preemptible = initial.Preemptible
	}
	if dcl.IsZeroValue(des.Reserved) {
		des.Reserved = initial.Reserved
	}

	return des
}

func canonicalizeNewNodeSchedulingConfig(c *Client, des, nw *NodeSchedulingConfig) *NodeSchedulingConfig {
	if des == nil || nw == nil {
		return nw
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
			if !compareNodeSchedulingConfig(c, &d, &n) {
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

func canonicalizeNodeNetworkEndpoints(des, initial *NodeNetworkEndpoints, opts ...dcl.ApplyOption) *NodeNetworkEndpoints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Node)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IPAddress) {
		des.IPAddress = initial.IPAddress
	}
	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}

	return des
}

func canonicalizeNewNodeNetworkEndpoints(c *Client, des, nw *NodeNetworkEndpoints) *NodeNetworkEndpoints {
	if des == nil || nw == nil {
		return nw
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
			if !compareNodeNetworkEndpoints(c, &d, &n) {
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

func canonicalizeNodeSymptoms(des, initial *NodeSymptoms, opts ...dcl.ApplyOption) *NodeSymptoms {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Node)
		_ = r
	}

	if initial == nil {
		return des
	}

	des.CreateTime = canonicalizeNodeSymptomsCreateTime(des.CreateTime, initial.CreateTime, opts...)
	if dcl.IsZeroValue(des.SymptomType) {
		des.SymptomType = initial.SymptomType
	}
	if dcl.IsZeroValue(des.Details) {
		des.Details = initial.Details
	}
	if dcl.IsZeroValue(des.WorkerId) {
		des.WorkerId = initial.WorkerId
	}

	return des
}

func canonicalizeNewNodeSymptoms(c *Client, des, nw *NodeSymptoms) *NodeSymptoms {
	if des == nil || nw == nil {
		return nw
	}

	nw.CreateTime = canonicalizeNewNodeSymptomsCreateTime(c, des.CreateTime, nw.CreateTime)

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
			if !compareNodeSymptoms(c, &d, &n) {
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

func canonicalizeNodeSymptomsCreateTime(des, initial *NodeSymptomsCreateTime, opts ...dcl.ApplyOption) *NodeSymptomsCreateTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*Node)
		_ = r
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
			if !compareNodeSymptomsCreateTime(c, &d, &n) {
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

type nodeDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         nodeApiOperation
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
	if !dcl.IsZeroValue(desired.Name) && (dcl.IsZeroValue(actual.Name) || !reflect.DeepEqual(*desired.Name, *actual.Name)) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && (dcl.IsZeroValue(actual.Description) || !reflect.DeepEqual(*desired.Description, *actual.Description)) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.AcceleratorType) && (dcl.IsZeroValue(actual.AcceleratorType) || !reflect.DeepEqual(*desired.AcceleratorType, *actual.AcceleratorType)) {
		c.Config.Logger.Infof("Detected diff in AcceleratorType.\nDESIRED: %v\nACTUAL: %v", desired.AcceleratorType, actual.AcceleratorType)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "AcceleratorType",
		})
	}
	if !dcl.IsZeroValue(desired.TensorflowVersion) && (dcl.IsZeroValue(actual.TensorflowVersion) || !reflect.DeepEqual(*desired.TensorflowVersion, *actual.TensorflowVersion)) {
		c.Config.Logger.Infof("Detected diff in TensorflowVersion.\nDESIRED: %v\nACTUAL: %v", desired.TensorflowVersion, actual.TensorflowVersion)

		diffs = append(diffs, nodeDiff{
			UpdateOp:  &updateNodeReimageNodeOperation{},
			FieldName: "TensorflowVersion",
		})

	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.NameToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !dcl.IsZeroValue(desired.CidrBlock) && (dcl.IsZeroValue(actual.CidrBlock) || !reflect.DeepEqual(*desired.CidrBlock, *actual.CidrBlock)) {
		c.Config.Logger.Infof("Detected diff in CidrBlock.\nDESIRED: %v\nACTUAL: %v", desired.CidrBlock, actual.CidrBlock)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "CidrBlock",
		})
	}
	if compareNodeSchedulingConfig(c, desired.SchedulingConfig, actual.SchedulingConfig) {
		c.Config.Logger.Infof("Detected diff in SchedulingConfig.\nDESIRED: %v\nACTUAL: %v", desired.SchedulingConfig, actual.SchedulingConfig)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "SchedulingConfig",
		})
	}
	if !reflect.DeepEqual(desired.Labels, actual.Labels) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "Labels",
		})
	}
	if !dcl.IsZeroValue(desired.UseServiceNetworking) && (dcl.IsZeroValue(actual.UseServiceNetworking) || !reflect.DeepEqual(*desired.UseServiceNetworking, *actual.UseServiceNetworking)) {
		c.Config.Logger.Infof("Detected diff in UseServiceNetworking.\nDESIRED: %v\nACTUAL: %v", desired.UseServiceNetworking, actual.UseServiceNetworking)
		diffs = append(diffs, nodeDiff{
			RequiresRecreate: true,
			FieldName:        "UseServiceNetworking",
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
func compareNodeCreateTimeSlice(c *Client, desired, actual []NodeCreateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeCreateTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeCreateTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeCreateTime, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeCreateTime(c *Client, desired, actual *NodeCreateTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareNodeSchedulingConfigSlice(c *Client, desired, actual []NodeSchedulingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeSchedulingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeSchedulingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeSchedulingConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeSchedulingConfig(c *Client, desired, actual *NodeSchedulingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Preemptible == nil && desired.Preemptible != nil && !dcl.IsEmptyValueIndirect(desired.Preemptible) {
		c.Config.Logger.Infof("desired Preemptible %s - but actually nil", dcl.SprintResource(desired.Preemptible))
		return true
	}
	if !reflect.DeepEqual(desired.Preemptible, actual.Preemptible) && !dcl.IsZeroValue(desired.Preemptible) && !(dcl.IsEmptyValueIndirect(desired.Preemptible) && dcl.IsZeroValue(actual.Preemptible)) {
		c.Config.Logger.Infof("Diff in Preemptible. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Preemptible), dcl.SprintResource(actual.Preemptible))
		return true
	}
	if actual.Reserved == nil && desired.Reserved != nil && !dcl.IsEmptyValueIndirect(desired.Reserved) {
		c.Config.Logger.Infof("desired Reserved %s - but actually nil", dcl.SprintResource(desired.Reserved))
		return true
	}
	if !reflect.DeepEqual(desired.Reserved, actual.Reserved) && !dcl.IsZeroValue(desired.Reserved) && !(dcl.IsEmptyValueIndirect(desired.Reserved) && dcl.IsZeroValue(actual.Reserved)) {
		c.Config.Logger.Infof("Diff in Reserved. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Reserved), dcl.SprintResource(actual.Reserved))
		return true
	}
	return false
}
func compareNodeNetworkEndpointsSlice(c *Client, desired, actual []NodeNetworkEndpoints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeNetworkEndpoints, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeNetworkEndpoints(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeNetworkEndpoints, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeNetworkEndpoints(c *Client, desired, actual *NodeNetworkEndpoints) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}
func compareNodeSymptomsSlice(c *Client, desired, actual []NodeSymptoms) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeSymptoms, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeSymptoms(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeSymptoms, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeSymptoms(c *Client, desired, actual *NodeSymptoms) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CreateTime == nil && desired.CreateTime != nil && !dcl.IsEmptyValueIndirect(desired.CreateTime) {
		c.Config.Logger.Infof("desired CreateTime %s - but actually nil", dcl.SprintResource(desired.CreateTime))
		return true
	}
	if compareNodeSymptomsCreateTime(c, desired.CreateTime, actual.CreateTime) && !dcl.IsZeroValue(desired.CreateTime) {
		c.Config.Logger.Infof("Diff in CreateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateTime), dcl.SprintResource(actual.CreateTime))
		return true
	}
	if actual.SymptomType == nil && desired.SymptomType != nil && !dcl.IsEmptyValueIndirect(desired.SymptomType) {
		c.Config.Logger.Infof("desired SymptomType %s - but actually nil", dcl.SprintResource(desired.SymptomType))
		return true
	}
	if !reflect.DeepEqual(desired.SymptomType, actual.SymptomType) && !dcl.IsZeroValue(desired.SymptomType) && !(dcl.IsEmptyValueIndirect(desired.SymptomType) && dcl.IsZeroValue(actual.SymptomType)) {
		c.Config.Logger.Infof("Diff in SymptomType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SymptomType), dcl.SprintResource(actual.SymptomType))
		return true
	}
	if actual.Details == nil && desired.Details != nil && !dcl.IsEmptyValueIndirect(desired.Details) {
		c.Config.Logger.Infof("desired Details %s - but actually nil", dcl.SprintResource(desired.Details))
		return true
	}
	if !reflect.DeepEqual(desired.Details, actual.Details) && !dcl.IsZeroValue(desired.Details) && !(dcl.IsEmptyValueIndirect(desired.Details) && dcl.IsZeroValue(actual.Details)) {
		c.Config.Logger.Infof("Diff in Details. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Details), dcl.SprintResource(actual.Details))
		return true
	}
	if actual.WorkerId == nil && desired.WorkerId != nil && !dcl.IsEmptyValueIndirect(desired.WorkerId) {
		c.Config.Logger.Infof("desired WorkerId %s - but actually nil", dcl.SprintResource(desired.WorkerId))
		return true
	}
	if !reflect.DeepEqual(desired.WorkerId, actual.WorkerId) && !dcl.IsZeroValue(desired.WorkerId) && !(dcl.IsEmptyValueIndirect(desired.WorkerId) && dcl.IsZeroValue(actual.WorkerId)) {
		c.Config.Logger.Infof("Diff in WorkerId. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.WorkerId), dcl.SprintResource(actual.WorkerId))
		return true
	}
	return false
}
func compareNodeSymptomsCreateTimeSlice(c *Client, desired, actual []NodeSymptomsCreateTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeSymptomsCreateTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeSymptomsCreateTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeSymptomsCreateTime, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeSymptomsCreateTime(c *Client, desired, actual *NodeSymptomsCreateTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Seconds == nil && desired.Seconds != nil && !dcl.IsEmptyValueIndirect(desired.Seconds) {
		c.Config.Logger.Infof("desired Seconds %s - but actually nil", dcl.SprintResource(desired.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) && !(dcl.IsEmptyValueIndirect(desired.Seconds) && dcl.IsZeroValue(actual.Seconds)) {
		c.Config.Logger.Infof("Diff in Seconds. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if actual.Nanos == nil && desired.Nanos != nil && !dcl.IsEmptyValueIndirect(desired.Nanos) {
		c.Config.Logger.Infof("desired Nanos %s - but actually nil", dcl.SprintResource(desired.Nanos))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) && !(dcl.IsEmptyValueIndirect(desired.Nanos) && dcl.IsZeroValue(actual.Nanos)) {
		c.Config.Logger.Infof("Diff in Nanos. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}
func compareNodeStateEnumSlice(c *Client, desired, actual []NodeStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeStateEnum(c *Client, desired, actual *NodeStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNodeHealthEnumSlice(c *Client, desired, actual []NodeHealthEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeHealthEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeHealthEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeHealthEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeHealthEnum(c *Client, desired, actual *NodeHealthEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareNodeSymptomsSymptomTypeEnumSlice(c *Client, desired, actual []NodeSymptomsSymptomTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NodeSymptomsSymptomTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNodeSymptomsSymptomTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NodeSymptomsSymptomTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNodeSymptomsSymptomTypeEnum(c *Client, desired, actual *NodeSymptomsSymptomTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Node) urlNormalized() *Node {
	normalized := deepcopy.Copy(*r).(Node)
	normalized.Network = dcl.SelfLinkToName(r.Network)
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v, err := expandNodeSchedulingConfig(c, f.SchedulingConfig); err != nil {
		return nil, fmt.Errorf("error expanding SchedulingConfig into schedulingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["schedulingConfig"] = v
	}
	if v, err := expandNodeNetworkEndpointsSlice(c, f.NetworkEndpoints); err != nil {
		return nil, fmt.Errorf("error expanding NetworkEndpoints into networkEndpoints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["symptoms"] = v
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

	r := &Node{}
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.AcceleratorType = dcl.FlattenString(m["acceleratorType"])
	r.IPAddress = dcl.FlattenString(m["ipAddress"])
	r.Port = dcl.FlattenString(m["port"])
	r.State = flattenNodeStateEnum(m["state"])
	r.HealthDescription = dcl.FlattenString(m["healthDescription"])
	r.TensorflowVersion = dcl.FlattenString(m["tensorflowVersion"])
	r.Network = dcl.FlattenString(m["network"])
	r.CidrBlock = dcl.FlattenString(m["cidrBlock"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.CreateTime = flattenNodeCreateTime(c, m["createTime"])
	r.SchedulingConfig = flattenNodeSchedulingConfig(c, m["schedulingConfig"])
	r.NetworkEndpoints = flattenNodeNetworkEndpointsSlice(c, m["networkEndpoints"])
	r.Health = flattenNodeHealthEnum(m["health"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.UseServiceNetworking = dcl.FlattenBool(m["useServiceNetworking"])
	r.Symptoms = flattenNodeSymptomsSlice(c, m["symptoms"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
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
		items = append(items, *flattenNodeStateEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenNodeHealthEnum(item.(map[string]interface{})))
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
		items = append(items, *flattenNodeSymptomsSymptomTypeEnum(item.(map[string]interface{})))
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
