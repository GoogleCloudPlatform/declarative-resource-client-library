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

func (r *WorkerPool) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.WorkerConfig) {
		if err := r.WorkerConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkConfig) {
		if err := r.NetworkConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkerPoolWorkerConfig) validate() error {
	return nil
}
func (r *WorkerPoolNetworkConfig) validate() error {
	if err := dcl.Required(r, "peeredNetwork"); err != nil {
		return err
	}
	return nil
}
func (r *WorkerPool) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudbuild.googleapis.com/v1beta1/", params)
}

func (r *WorkerPool) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workerPools/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *WorkerPool) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workerPools", nr.basePath(), userBasePath, params), nil

}

func (r *WorkerPool) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workerPools?workerPoolId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *WorkerPool) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workerPools/{{name}}", nr.basePath(), userBasePath, params), nil
}

// workerPoolApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type workerPoolApiOperation interface {
	do(context.Context, *WorkerPool, *Client) error
}

// newUpdateWorkerPoolUpdateWorkerPoolRequest creates a request for an
// WorkerPool resource's UpdateWorkerPool update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWorkerPoolUpdateWorkerPoolRequest(ctx context.Context, f *WorkerPool, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandWorkerPoolWorkerConfig(c, f.WorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkerConfig into workerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["workerConfig"] = v
	}
	return req, nil
}

// marshalUpdateWorkerPoolUpdateWorkerPoolRequest converts the update into
// the final JSON request body.
func marshalUpdateWorkerPoolUpdateWorkerPoolRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWorkerPoolUpdateWorkerPoolOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWorkerPoolUpdateWorkerPoolOperation) do(ctx context.Context, r *WorkerPool, c *Client) error {
	_, err := c.GetWorkerPool(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWorkerPool")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateWorkerPoolUpdateWorkerPoolRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateWorkerPoolUpdateWorkerPoolRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listWorkerPoolRaw(ctx context.Context, r *WorkerPool, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WorkerPoolMaxPage {
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

type listWorkerPoolOperation struct {
	WorkerPools []map[string]interface{} `json:"workerPools"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listWorkerPool(ctx context.Context, r *WorkerPool, pageToken string, pageSize int32) ([]*WorkerPool, string, error) {
	b, err := c.listWorkerPoolRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWorkerPoolOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*WorkerPool
	for _, v := range m.WorkerPools {
		res, err := unmarshalMapWorkerPool(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWorkerPool(ctx context.Context, f func(*WorkerPool) bool, resources []*WorkerPool) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWorkerPool(ctx, res)
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

type deleteWorkerPoolOperation struct{}

func (op *deleteWorkerPoolOperation) do(ctx context.Context, r *WorkerPool, c *Client) error {
	r, err := c.GetWorkerPool(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "WorkerPool not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetWorkerPool checking for existence. error: %v", err)
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
	var o operations.StandardGCPOperation
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
		_, err = c.GetWorkerPool(ctx, r)
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
type createWorkerPoolOperation struct {
	response map[string]interface{}
}

func (op *createWorkerPoolOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWorkerPoolOperation) do(ctx context.Context, r *WorkerPool, c *Client) error {
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetWorkerPool(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWorkerPoolRaw(ctx context.Context, r *WorkerPool) ([]byte, error) {

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

func (c *Client) workerPoolDiffsForRawDesired(ctx context.Context, rawDesired *WorkerPool, opts ...dcl.ApplyOption) (initial, desired *WorkerPool, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *WorkerPool
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WorkerPool); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected WorkerPool, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWorkerPool(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a WorkerPool resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve WorkerPool resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that WorkerPool resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWorkerPoolDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for WorkerPool: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for WorkerPool: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWorkerPoolInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for WorkerPool: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWorkerPoolDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for WorkerPool: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWorkerPool(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWorkerPoolInitialState(rawInitial, rawDesired *WorkerPool) (*WorkerPool, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeWorkerPoolDesiredState(rawDesired, rawInitial *WorkerPool, opts ...dcl.ApplyOption) (*WorkerPool, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.WorkerConfig = canonicalizeWorkerPoolWorkerConfig(rawDesired.WorkerConfig, nil, opts...)
		rawDesired.NetworkConfig = canonicalizeWorkerPoolNetworkConfig(rawDesired.NetworkConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &WorkerPool{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.WorkerConfig = canonicalizeWorkerPoolWorkerConfig(rawDesired.WorkerConfig, rawInitial.WorkerConfig, opts...)
	canonicalDesired.NetworkConfig = canonicalizeWorkerPoolNetworkConfig(rawDesired.NetworkConfig, rawInitial.NetworkConfig, opts...)
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

func canonicalizeWorkerPoolNewState(c *Client, rawNew, rawDesired *WorkerPool) (*WorkerPool, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.WorkerConfig) && dcl.IsEmptyValueIndirect(rawDesired.WorkerConfig) {
		rawNew.WorkerConfig = rawDesired.WorkerConfig
	} else {
		rawNew.WorkerConfig = canonicalizeNewWorkerPoolWorkerConfig(c, rawDesired.WorkerConfig, rawNew.WorkerConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkConfig) && dcl.IsEmptyValueIndirect(rawDesired.NetworkConfig) {
		rawNew.NetworkConfig = rawDesired.NetworkConfig
	} else {
		rawNew.NetworkConfig = canonicalizeNewWorkerPoolNetworkConfig(c, rawDesired.NetworkConfig, rawNew.NetworkConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeWorkerPoolWorkerConfig(des, initial *WorkerPoolWorkerConfig, opts ...dcl.ApplyOption) *WorkerPoolWorkerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkerPoolWorkerConfig{}

	if dcl.StringCanonicalize(des.MachineType, initial.MachineType) || dcl.IsZeroValue(des.MachineType) {
		cDes.MachineType = initial.MachineType
	} else {
		cDes.MachineType = des.MachineType
	}
	if dcl.IsZeroValue(des.DiskSizeGb) {
		des.DiskSizeGb = initial.DiskSizeGb
	} else {
		cDes.DiskSizeGb = des.DiskSizeGb
	}
	if dcl.BoolCanonicalize(des.NoExternalIP, initial.NoExternalIP) || dcl.IsZeroValue(des.NoExternalIP) {
		cDes.NoExternalIP = initial.NoExternalIP
	} else {
		cDes.NoExternalIP = des.NoExternalIP
	}

	return cDes
}

func canonicalizeNewWorkerPoolWorkerConfig(c *Client, des, nw *WorkerPoolWorkerConfig) *WorkerPoolWorkerConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MachineType, nw.MachineType) {
		nw.MachineType = des.MachineType
	}
	if dcl.BoolCanonicalize(des.NoExternalIP, nw.NoExternalIP) {
		nw.NoExternalIP = des.NoExternalIP
	}

	return nw
}

func canonicalizeNewWorkerPoolWorkerConfigSet(c *Client, des, nw []WorkerPoolWorkerConfig) []WorkerPoolWorkerConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkerPoolWorkerConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareWorkerPoolWorkerConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewWorkerPoolWorkerConfigSlice(c *Client, des, nw []WorkerPoolWorkerConfig) []WorkerPoolWorkerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkerPoolWorkerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkerPoolWorkerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeWorkerPoolNetworkConfig(des, initial *WorkerPoolNetworkConfig, opts ...dcl.ApplyOption) *WorkerPoolNetworkConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkerPoolNetworkConfig{}

	if dcl.NameToSelfLink(des.PeeredNetwork, initial.PeeredNetwork) || dcl.IsZeroValue(des.PeeredNetwork) {
		cDes.PeeredNetwork = initial.PeeredNetwork
	} else {
		cDes.PeeredNetwork = des.PeeredNetwork
	}

	return cDes
}

func canonicalizeNewWorkerPoolNetworkConfig(c *Client, des, nw *WorkerPoolNetworkConfig) *WorkerPoolNetworkConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.PeeredNetwork, nw.PeeredNetwork) {
		nw.PeeredNetwork = des.PeeredNetwork
	}

	return nw
}

func canonicalizeNewWorkerPoolNetworkConfigSet(c *Client, des, nw []WorkerPoolNetworkConfig) []WorkerPoolNetworkConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkerPoolNetworkConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareWorkerPoolNetworkConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewWorkerPoolNetworkConfigSlice(c *Client, des, nw []WorkerPoolNetworkConfig) []WorkerPoolNetworkConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkerPoolNetworkConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkerPoolNetworkConfig(c, &d, &n))
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
func diffWorkerPool(c *Client, desired, actual *WorkerPool, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkerConfig, actual.WorkerConfig, dcl.Info{ObjectFunction: compareWorkerPoolWorkerConfigNewStyle, EmptyObject: EmptyWorkerPoolWorkerConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkerConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkConfig, actual.NetworkConfig, dcl.Info{ObjectFunction: compareWorkerPoolNetworkConfigNewStyle, EmptyObject: EmptyWorkerPoolNetworkConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkConfig")); len(ds) != 0 || err != nil {
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
func compareWorkerPoolWorkerConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkerPoolWorkerConfig)
	if !ok {
		desiredNotPointer, ok := d.(WorkerPoolWorkerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkerPoolWorkerConfig or *WorkerPoolWorkerConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkerPoolWorkerConfig)
	if !ok {
		actualNotPointer, ok := a.(WorkerPoolWorkerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkerPoolWorkerConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.Info{OperationSelector: dcl.TriggersOperation("updateWorkerPoolUpdateWorkerPoolOperation")}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DiskSizeGb, actual.DiskSizeGb, dcl.Info{OperationSelector: dcl.TriggersOperation("updateWorkerPoolUpdateWorkerPoolOperation")}, fn.AddNest("DiskSizeGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NoExternalIP, actual.NoExternalIP, dcl.Info{OperationSelector: dcl.TriggersOperation("updateWorkerPoolUpdateWorkerPoolOperation")}, fn.AddNest("NoExternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkerPoolNetworkConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkerPoolNetworkConfig)
	if !ok {
		desiredNotPointer, ok := d.(WorkerPoolNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkerPoolNetworkConfig or *WorkerPoolNetworkConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkerPoolNetworkConfig)
	if !ok {
		actualNotPointer, ok := a.(WorkerPoolNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkerPoolNetworkConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PeeredNetwork, actual.PeeredNetwork, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PeeredNetwork")); len(ds) != 0 || err != nil {
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
func (r *WorkerPool) urlNormalized() *WorkerPool {
	normalized := dcl.Copy(*r).(WorkerPool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *WorkerPool) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateWorkerPool" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/workerPools/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the WorkerPool resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *WorkerPool) marshal(c *Client) ([]byte, error) {
	m, err := expandWorkerPool(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling WorkerPool: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWorkerPool decodes JSON responses into the WorkerPool resource schema.
func unmarshalWorkerPool(b []byte, c *Client) (*WorkerPool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWorkerPool(m, c)
}

func unmarshalMapWorkerPool(m map[string]interface{}, c *Client) (*WorkerPool, error) {

	return flattenWorkerPool(c, m), nil
}

// expandWorkerPool expands WorkerPool into a JSON request object.
func expandWorkerPool(c *Client, f *WorkerPool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/workerPools/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.DeleteTime; !dcl.IsEmptyValueIndirect(v) {
		m["deleteTime"] = v
	}
	if v, err := expandWorkerPoolWorkerConfig(c, f.WorkerConfig); err != nil {
		return nil, fmt.Errorf("error expanding WorkerConfig into workerConfig: %w", err)
	} else if v != nil {
		m["workerConfig"] = v
	}
	if v, err := expandWorkerPoolNetworkConfig(c, f.NetworkConfig); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if v != nil {
		m["networkConfig"] = v
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

// flattenWorkerPool flattens WorkerPool from a JSON request object into the
// WorkerPool type.
func flattenWorkerPool(c *Client, i interface{}) *WorkerPool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &WorkerPool{}
	res.Name = dcl.FlattenString(m["name"])
	res.State = flattenWorkerPoolStateEnum(m["state"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.WorkerConfig = flattenWorkerPoolWorkerConfig(c, m["workerConfig"])
	res.NetworkConfig = flattenWorkerPoolNetworkConfig(c, m["networkConfig"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandWorkerPoolWorkerConfigMap expands the contents of WorkerPoolWorkerConfig into a JSON
// request object.
func expandWorkerPoolWorkerConfigMap(c *Client, f map[string]WorkerPoolWorkerConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkerPoolWorkerConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkerPoolWorkerConfigSlice expands the contents of WorkerPoolWorkerConfig into a JSON
// request object.
func expandWorkerPoolWorkerConfigSlice(c *Client, f []WorkerPoolWorkerConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkerPoolWorkerConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkerPoolWorkerConfigMap flattens the contents of WorkerPoolWorkerConfig from a JSON
// response object.
func flattenWorkerPoolWorkerConfigMap(c *Client, i interface{}) map[string]WorkerPoolWorkerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkerPoolWorkerConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkerPoolWorkerConfig{}
	}

	items := make(map[string]WorkerPoolWorkerConfig)
	for k, item := range a {
		items[k] = *flattenWorkerPoolWorkerConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkerPoolWorkerConfigSlice flattens the contents of WorkerPoolWorkerConfig from a JSON
// response object.
func flattenWorkerPoolWorkerConfigSlice(c *Client, i interface{}) []WorkerPoolWorkerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkerPoolWorkerConfig{}
	}

	if len(a) == 0 {
		return []WorkerPoolWorkerConfig{}
	}

	items := make([]WorkerPoolWorkerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkerPoolWorkerConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkerPoolWorkerConfig expands an instance of WorkerPoolWorkerConfig into a JSON
// request object.
func expandWorkerPoolWorkerConfig(c *Client, f *WorkerPoolWorkerConfig) (map[string]interface{}, error) {
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
	if v := f.NoExternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["noExternalIp"] = v
	}

	return m, nil
}

// flattenWorkerPoolWorkerConfig flattens an instance of WorkerPoolWorkerConfig from a JSON
// response object.
func flattenWorkerPoolWorkerConfig(c *Client, i interface{}) *WorkerPoolWorkerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkerPoolWorkerConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkerPoolWorkerConfig
	}
	r.MachineType = dcl.FlattenString(m["machineType"])
	r.DiskSizeGb = dcl.FlattenInteger(m["diskSizeGb"])
	r.NoExternalIP = dcl.FlattenBool(m["noExternalIp"])

	return r
}

// expandWorkerPoolNetworkConfigMap expands the contents of WorkerPoolNetworkConfig into a JSON
// request object.
func expandWorkerPoolNetworkConfigMap(c *Client, f map[string]WorkerPoolNetworkConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkerPoolNetworkConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkerPoolNetworkConfigSlice expands the contents of WorkerPoolNetworkConfig into a JSON
// request object.
func expandWorkerPoolNetworkConfigSlice(c *Client, f []WorkerPoolNetworkConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkerPoolNetworkConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkerPoolNetworkConfigMap flattens the contents of WorkerPoolNetworkConfig from a JSON
// response object.
func flattenWorkerPoolNetworkConfigMap(c *Client, i interface{}) map[string]WorkerPoolNetworkConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkerPoolNetworkConfig{}
	}

	if len(a) == 0 {
		return map[string]WorkerPoolNetworkConfig{}
	}

	items := make(map[string]WorkerPoolNetworkConfig)
	for k, item := range a {
		items[k] = *flattenWorkerPoolNetworkConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenWorkerPoolNetworkConfigSlice flattens the contents of WorkerPoolNetworkConfig from a JSON
// response object.
func flattenWorkerPoolNetworkConfigSlice(c *Client, i interface{}) []WorkerPoolNetworkConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkerPoolNetworkConfig{}
	}

	if len(a) == 0 {
		return []WorkerPoolNetworkConfig{}
	}

	items := make([]WorkerPoolNetworkConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkerPoolNetworkConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandWorkerPoolNetworkConfig expands an instance of WorkerPoolNetworkConfig into a JSON
// request object.
func expandWorkerPoolNetworkConfig(c *Client, f *WorkerPoolNetworkConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PeeredNetwork; !dcl.IsEmptyValueIndirect(v) {
		m["peeredNetwork"] = v
	}

	return m, nil
}

// flattenWorkerPoolNetworkConfig flattens an instance of WorkerPoolNetworkConfig from a JSON
// response object.
func flattenWorkerPoolNetworkConfig(c *Client, i interface{}) *WorkerPoolNetworkConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkerPoolNetworkConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkerPoolNetworkConfig
	}
	r.PeeredNetwork = dcl.FlattenString(m["peeredNetwork"])

	return r
}

// flattenWorkerPoolStateEnumMap flattens the contents of WorkerPoolStateEnum from a JSON
// response object.
func flattenWorkerPoolStateEnumMap(c *Client, i interface{}) map[string]WorkerPoolStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkerPoolStateEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkerPoolStateEnum{}
	}

	items := make(map[string]WorkerPoolStateEnum)
	for k, item := range a {
		items[k] = *flattenWorkerPoolStateEnum(item.(interface{}))
	}

	return items
}

// flattenWorkerPoolStateEnumSlice flattens the contents of WorkerPoolStateEnum from a JSON
// response object.
func flattenWorkerPoolStateEnumSlice(c *Client, i interface{}) []WorkerPoolStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkerPoolStateEnum{}
	}

	if len(a) == 0 {
		return []WorkerPoolStateEnum{}
	}

	items := make([]WorkerPoolStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkerPoolStateEnum(item.(interface{})))
	}

	return items
}

// flattenWorkerPoolStateEnum asserts that an interface is a string, and returns a
// pointer to a *WorkerPoolStateEnum with the same value as that string.
func flattenWorkerPoolStateEnum(i interface{}) *WorkerPoolStateEnum {
	s, ok := i.(string)
	if !ok {
		return WorkerPoolStateEnumRef("")
	}

	return WorkerPoolStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *WorkerPool) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWorkerPool(b, c)
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

type workerPoolDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         workerPoolApiOperation
}

func convertFieldDiffsToWorkerPoolDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]workerPoolDiff, error) {
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
	var diffs []workerPoolDiff
	// For each operation name, create a workerPoolDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := workerPoolDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToWorkerPoolApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToWorkerPoolApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (workerPoolApiOperation, error) {
	switch opName {

	case "updateWorkerPoolUpdateWorkerPoolOperation":
		return &updateWorkerPoolUpdateWorkerPoolOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractWorkerPoolFields(r *WorkerPool) error {
	return nil
}
