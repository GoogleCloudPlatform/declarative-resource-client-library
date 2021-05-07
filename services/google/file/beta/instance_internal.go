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

func (r *Instance) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *InstanceFileShares) validate() error {
	return nil
}
func (r *InstanceFileSharesNfsExportOptions) validate() error {
	return nil
}
func (r *InstanceNetworks) validate() error {
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances?instanceId={{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, params), nil
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceUpdateInstanceRequest creates a request for an
// Instance resource's UpdateInstance update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateInstanceRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandInstanceFileSharesSlice(c, f.FileShares); err != nil {
		return nil, fmt.Errorf("error expanding FileShares into fileShares: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["fileShares"] = v
	}
	return req, nil
}

// marshalUpdateInstanceUpdateInstanceRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateInstanceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceUpdateInstanceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateInstance")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"description", "labels", "fileShares"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateInstanceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateInstanceRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://file.googleapis.com/v1beta1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceMaxPage {
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

type listInstanceOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Instance
	for _, v := range m.Instances {
		res := flattenInstance(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstance(ctx context.Context, f func(*Instance) bool, resources []*Instance) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstance(ctx, res)
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

type deleteInstanceOperation struct{}

func (op *deleteInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {

	_, err := c.GetInstance(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Instance not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstance checking for existence. error: %v", err)
		return err
	}

	u, err := instanceDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://file.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetInstance(ctx, r.urlNormalized())
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
type createInstanceOperation struct {
	response map[string]interface{}
}

func (op *createInstanceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := instanceCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://file.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetInstance(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

	u, err := instanceGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) instanceDiffsForRawDesired(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (initial, desired *Instance, diffs []instanceDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Instance
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Instance); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Instance, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstance(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Instance resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Instance resource: %v", err)
		}
		c.Config.Logger.Info("Found that Instance resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Instance: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Instance: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Instance: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Instance: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstance(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceInitialState(rawInitial, rawDesired *Instance) (*Instance, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceDesiredState(rawDesired, rawInitial *Instance, opts ...dcl.ApplyOption) (*Instance, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.Tier) {
		rawDesired.Tier = rawInitial.Tier
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.FileShares) {
		rawDesired.FileShares = rawInitial.FileShares
	}
	if dcl.IsZeroValue(rawDesired.Networks) {
		rawDesired.Networks = rawInitial.Networks
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StatusMessage) && dcl.IsEmptyValueIndirect(rawDesired.StatusMessage) {
		rawNew.StatusMessage = rawDesired.StatusMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StatusMessage, rawNew.StatusMessage) {
			rawNew.StatusMessage = rawDesired.StatusMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Tier) && dcl.IsEmptyValueIndirect(rawDesired.Tier) {
		rawNew.Tier = rawDesired.Tier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FileShares) && dcl.IsEmptyValueIndirect(rawDesired.FileShares) {
		rawNew.FileShares = rawDesired.FileShares
	} else {
		rawNew.FileShares = canonicalizeNewInstanceFileSharesSlice(c, rawDesired.FileShares, rawNew.FileShares)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Networks) && dcl.IsEmptyValueIndirect(rawDesired.Networks) {
		rawNew.Networks = rawDesired.Networks
	} else {
		rawNew.Networks = canonicalizeNewInstanceNetworksSlice(c, rawDesired.Networks, rawNew.Networks)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceFileShares(des, initial *InstanceFileShares, opts ...dcl.ApplyOption) *InstanceFileShares {
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
	if dcl.IsZeroValue(des.CapacityGb) {
		des.CapacityGb = initial.CapacityGb
	}
	if dcl.NameToSelfLink(des.SourceBackup, initial.SourceBackup) || dcl.IsZeroValue(des.SourceBackup) {
		des.SourceBackup = initial.SourceBackup
	}
	if dcl.IsZeroValue(des.NfsExportOptions) {
		des.NfsExportOptions = initial.NfsExportOptions
	}

	return des
}

func canonicalizeNewInstanceFileShares(c *Client, des, nw *InstanceFileShares) *InstanceFileShares {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.IsZeroValue(nw.CapacityGb) {
		nw.CapacityGb = des.CapacityGb
	}
	if dcl.NameToSelfLink(des.SourceBackup, nw.SourceBackup) {
		nw.SourceBackup = des.SourceBackup
	}
	nw.NfsExportOptions = canonicalizeNewInstanceFileSharesNfsExportOptionsSlice(c, des.NfsExportOptions, nw.NfsExportOptions)

	return nw
}

func canonicalizeNewInstanceFileSharesSet(c *Client, des, nw []InstanceFileShares) []InstanceFileShares {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceFileShares
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceFileSharesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceFileSharesSlice(c *Client, des, nw []InstanceFileShares) []InstanceFileShares {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceFileShares
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceFileShares(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceFileSharesNfsExportOptions(des, initial *InstanceFileSharesNfsExportOptions, opts ...dcl.ApplyOption) *InstanceFileSharesNfsExportOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.IPRanges) {
		des.IPRanges = initial.IPRanges
	}
	if dcl.IsZeroValue(des.AccessMode) {
		des.AccessMode = initial.AccessMode
	}
	if dcl.IsZeroValue(des.SquashMode) {
		des.SquashMode = initial.SquashMode
	}
	if dcl.IsZeroValue(des.AnonUid) {
		des.AnonUid = initial.AnonUid
	}
	if dcl.IsZeroValue(des.AnonGid) {
		des.AnonGid = initial.AnonGid
	}

	return des
}

func canonicalizeNewInstanceFileSharesNfsExportOptions(c *Client, des, nw *InstanceFileSharesNfsExportOptions) *InstanceFileSharesNfsExportOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.IPRanges) {
		nw.IPRanges = des.IPRanges
	}
	if dcl.IsZeroValue(nw.AccessMode) {
		nw.AccessMode = des.AccessMode
	}
	if dcl.IsZeroValue(nw.SquashMode) {
		nw.SquashMode = des.SquashMode
	}
	if dcl.IsZeroValue(nw.AnonUid) {
		nw.AnonUid = des.AnonUid
	}
	if dcl.IsZeroValue(nw.AnonGid) {
		nw.AnonGid = des.AnonGid
	}

	return nw
}

func canonicalizeNewInstanceFileSharesNfsExportOptionsSet(c *Client, des, nw []InstanceFileSharesNfsExportOptions) []InstanceFileSharesNfsExportOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceFileSharesNfsExportOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceFileSharesNfsExportOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceFileSharesNfsExportOptionsSlice(c *Client, des, nw []InstanceFileSharesNfsExportOptions) []InstanceFileSharesNfsExportOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceFileSharesNfsExportOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceFileSharesNfsExportOptions(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceNetworks(des, initial *InstanceNetworks, opts ...dcl.ApplyOption) *InstanceNetworks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Network, initial.Network) || dcl.IsZeroValue(des.Network) {
		des.Network = initial.Network
	}
	if dcl.IsZeroValue(des.Modes) {
		des.Modes = initial.Modes
	}
	if dcl.StringCanonicalize(des.ReservedIPRange, initial.ReservedIPRange) || dcl.IsZeroValue(des.ReservedIPRange) {
		des.ReservedIPRange = initial.ReservedIPRange
	}

	return des
}

func canonicalizeNewInstanceNetworks(c *Client, des, nw *InstanceNetworks) *InstanceNetworks {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Network, nw.Network) {
		nw.Network = des.Network
	}
	if dcl.IsZeroValue(nw.Modes) {
		nw.Modes = des.Modes
	}
	if dcl.StringCanonicalize(des.ReservedIPRange, nw.ReservedIPRange) {
		nw.ReservedIPRange = des.ReservedIPRange
	}
	if dcl.IsZeroValue(nw.IPAddresses) {
		nw.IPAddresses = des.IPAddresses
	}

	return nw
}

func canonicalizeNewInstanceNetworksSet(c *Client, des, nw []InstanceNetworks) []InstanceNetworks {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceNetworks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareInstanceNetworksNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewInstanceNetworksSlice(c *Client, des, nw []InstanceNetworks) []InstanceNetworks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceNetworks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworks(c, &d, &n))
	}

	return items
}

type instanceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceApiOperation
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
func diffInstance(c *Client, desired, actual *Instance, opts ...dcl.ApplyOption) ([]instanceDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []instanceDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StatusMessage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Tier, actual.Tier, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.FileShares, actual.FileShares, dcl.Info{ObjectFunction: compareInstanceFileSharesNewStyle, OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("FileShares")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Networks, actual.Networks, dcl.Info{ObjectFunction: compareInstanceNetworksNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Networks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToInstanceDiff(ds, opts...)
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
	var deduped []instanceDiff
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
func compareInstanceFileSharesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceFileShares)
	if !ok {
		desiredNotPointer, ok := d.(InstanceFileShares)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileShares or *InstanceFileShares", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceFileShares)
	if !ok {
		actualNotPointer, ok := a.(InstanceFileShares)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileShares", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CapacityGb, actual.CapacityGb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CapacityGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceBackup, actual.SourceBackup, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceBackup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NfsExportOptions, actual.NfsExportOptions, dcl.Info{ObjectFunction: compareInstanceFileSharesNfsExportOptionsNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NfsExportOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceFileSharesNfsExportOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceFileSharesNfsExportOptions)
	if !ok {
		desiredNotPointer, ok := d.(InstanceFileSharesNfsExportOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileSharesNfsExportOptions or *InstanceFileSharesNfsExportOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceFileSharesNfsExportOptions)
	if !ok {
		actualNotPointer, ok := a.(InstanceFileSharesNfsExportOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileSharesNfsExportOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IPRanges, actual.IPRanges, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AccessMode, actual.AccessMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AccessMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SquashMode, actual.SquashMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SquashMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnonUid, actual.AnonUid, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AnonUid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnonGid, actual.AnonGid, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AnonGid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceNetworksNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceNetworks)
	if !ok {
		desiredNotPointer, ok := d.(InstanceNetworks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworks or *InstanceNetworks", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceNetworks)
	if !ok {
		actualNotPointer, ok := a.(InstanceNetworks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworks", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Modes, actual.Modes, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Modes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservedIPRange, actual.ReservedIPRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReservedIPRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAddresses, actual.IPAddresses, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPAddresses")); len(ds) != 0 || err != nil {
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
func (r *Instance) urlNormalized() *Instance {
	normalized := dcl.Copy(*r).(Instance)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Instance) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateInstance" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", "https://file.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Instance resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Instance) marshal(c *Client) ([]byte, error) {
	m, err := expandInstance(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Instance: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInstance decodes JSON responses into the Instance resource schema.
func unmarshalInstance(b []byte, c *Client) (*Instance, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstance(m, c)
}

func unmarshalMapInstance(m map[string]interface{}, c *Client) (*Instance, error) {

	return flattenInstance(c, m), nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.StatusMessage; !dcl.IsEmptyValueIndirect(v) {
		m["statusMessage"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.Tier; !dcl.IsEmptyValueIndirect(v) {
		m["tier"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandInstanceFileSharesSlice(c, f.FileShares); err != nil {
		return nil, fmt.Errorf("error expanding FileShares into fileShares: %w", err)
	} else if v != nil {
		m["fileShares"] = v
	}
	if v, err := expandInstanceNetworksSlice(c, f.Networks); err != nil {
		return nil, fmt.Errorf("error expanding Networks into networks: %w", err)
	} else if v != nil {
		m["networks"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
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

// flattenInstance flattens Instance from a JSON request object into the
// Instance type.
func flattenInstance(c *Client, i interface{}) *Instance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Instance{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.State = flattenInstanceStateEnum(m["state"])
	res.StatusMessage = dcl.FlattenString(m["statusMessage"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.Tier = flattenInstanceTierEnum(m["tier"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.FileShares = flattenInstanceFileSharesSlice(c, m["fileShares"])
	res.Networks = flattenInstanceNetworksSlice(c, m["networks"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandInstanceFileSharesMap expands the contents of InstanceFileShares into a JSON
// request object.
func expandInstanceFileSharesMap(c *Client, f map[string]InstanceFileShares) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFileShares(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFileSharesSlice expands the contents of InstanceFileShares into a JSON
// request object.
func expandInstanceFileSharesSlice(c *Client, f []InstanceFileShares) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFileShares(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFileSharesMap flattens the contents of InstanceFileShares from a JSON
// response object.
func flattenInstanceFileSharesMap(c *Client, i interface{}) map[string]InstanceFileShares {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFileShares{}
	}

	if len(a) == 0 {
		return map[string]InstanceFileShares{}
	}

	items := make(map[string]InstanceFileShares)
	for k, item := range a {
		items[k] = *flattenInstanceFileShares(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceFileSharesSlice flattens the contents of InstanceFileShares from a JSON
// response object.
func flattenInstanceFileSharesSlice(c *Client, i interface{}) []InstanceFileShares {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileShares{}
	}

	if len(a) == 0 {
		return []InstanceFileShares{}
	}

	items := make([]InstanceFileShares, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileShares(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceFileShares expands an instance of InstanceFileShares into a JSON
// request object.
func expandInstanceFileShares(c *Client, f *InstanceFileShares) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.CapacityGb; !dcl.IsEmptyValueIndirect(v) {
		m["capacityGb"] = v
	}
	if v := f.SourceBackup; !dcl.IsEmptyValueIndirect(v) {
		m["sourceBackup"] = v
	}
	if v, err := expandInstanceFileSharesNfsExportOptionsSlice(c, f.NfsExportOptions); err != nil {
		return nil, fmt.Errorf("error expanding NfsExportOptions into nfsExportOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nfsExportOptions"] = v
	}

	return m, nil
}

// flattenInstanceFileShares flattens an instance of InstanceFileShares from a JSON
// response object.
func flattenInstanceFileShares(c *Client, i interface{}) *InstanceFileShares {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFileShares{}
	r.Name = dcl.FlattenString(m["name"])
	r.CapacityGb = dcl.FlattenInteger(m["capacityGb"])
	r.SourceBackup = dcl.FlattenString(m["sourceBackup"])
	r.NfsExportOptions = flattenInstanceFileSharesNfsExportOptionsSlice(c, m["nfsExportOptions"])

	return r
}

// expandInstanceFileSharesNfsExportOptionsMap expands the contents of InstanceFileSharesNfsExportOptions into a JSON
// request object.
func expandInstanceFileSharesNfsExportOptionsMap(c *Client, f map[string]InstanceFileSharesNfsExportOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFileSharesNfsExportOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFileSharesNfsExportOptionsSlice expands the contents of InstanceFileSharesNfsExportOptions into a JSON
// request object.
func expandInstanceFileSharesNfsExportOptionsSlice(c *Client, f []InstanceFileSharesNfsExportOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFileSharesNfsExportOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFileSharesNfsExportOptionsMap flattens the contents of InstanceFileSharesNfsExportOptions from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsMap(c *Client, i interface{}) map[string]InstanceFileSharesNfsExportOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFileSharesNfsExportOptions{}
	}

	if len(a) == 0 {
		return map[string]InstanceFileSharesNfsExportOptions{}
	}

	items := make(map[string]InstanceFileSharesNfsExportOptions)
	for k, item := range a {
		items[k] = *flattenInstanceFileSharesNfsExportOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsSlice flattens the contents of InstanceFileSharesNfsExportOptions from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsSlice(c *Client, i interface{}) []InstanceFileSharesNfsExportOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileSharesNfsExportOptions{}
	}

	if len(a) == 0 {
		return []InstanceFileSharesNfsExportOptions{}
	}

	items := make([]InstanceFileSharesNfsExportOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileSharesNfsExportOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceFileSharesNfsExportOptions expands an instance of InstanceFileSharesNfsExportOptions into a JSON
// request object.
func expandInstanceFileSharesNfsExportOptions(c *Client, f *InstanceFileSharesNfsExportOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPRanges; !dcl.IsEmptyValueIndirect(v) {
		m["ipRanges"] = v
	}
	if v := f.AccessMode; !dcl.IsEmptyValueIndirect(v) {
		m["accessMode"] = v
	}
	if v := f.SquashMode; !dcl.IsEmptyValueIndirect(v) {
		m["squashMode"] = v
	}
	if v := f.AnonUid; !dcl.IsEmptyValueIndirect(v) {
		m["anonUid"] = v
	}
	if v := f.AnonGid; !dcl.IsEmptyValueIndirect(v) {
		m["anonGid"] = v
	}

	return m, nil
}

// flattenInstanceFileSharesNfsExportOptions flattens an instance of InstanceFileSharesNfsExportOptions from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptions(c *Client, i interface{}) *InstanceFileSharesNfsExportOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFileSharesNfsExportOptions{}
	r.IPRanges = dcl.FlattenStringSlice(m["ipRanges"])
	r.AccessMode = flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(m["accessMode"])
	r.SquashMode = flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(m["squashMode"])
	r.AnonUid = dcl.FlattenInteger(m["anonUid"])
	r.AnonGid = dcl.FlattenInteger(m["anonGid"])

	return r
}

// expandInstanceNetworksMap expands the contents of InstanceNetworks into a JSON
// request object.
func expandInstanceNetworksMap(c *Client, f map[string]InstanceNetworks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworksSlice expands the contents of InstanceNetworks into a JSON
// request object.
func expandInstanceNetworksSlice(c *Client, f []InstanceNetworks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworksMap flattens the contents of InstanceNetworks from a JSON
// response object.
func flattenInstanceNetworksMap(c *Client, i interface{}) map[string]InstanceNetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworks{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworks{}
	}

	items := make(map[string]InstanceNetworks)
	for k, item := range a {
		items[k] = *flattenInstanceNetworks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceNetworksSlice flattens the contents of InstanceNetworks from a JSON
// response object.
func flattenInstanceNetworksSlice(c *Client, i interface{}) []InstanceNetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworks{}
	}

	if len(a) == 0 {
		return []InstanceNetworks{}
	}

	items := make([]InstanceNetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworks(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceNetworks expands an instance of InstanceNetworks into a JSON
// request object.
func expandInstanceNetworks(c *Client, f *InstanceNetworks) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Modes; !dcl.IsEmptyValueIndirect(v) {
		m["modes"] = v
	}
	if v := f.ReservedIPRange; !dcl.IsEmptyValueIndirect(v) {
		m["reservedIpRange"] = v
	}
	if v := f.IPAddresses; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddresses"] = v
	}

	return m, nil
}

// flattenInstanceNetworks flattens an instance of InstanceNetworks from a JSON
// response object.
func flattenInstanceNetworks(c *Client, i interface{}) *InstanceNetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworks{}
	r.Network = dcl.FlattenString(m["network"])
	r.Modes = flattenInstanceNetworksModesEnumSlice(c, m["modes"])
	r.ReservedIPRange = dcl.FlattenString(m["reservedIpRange"])
	r.IPAddresses = dcl.FlattenStringSlice(m["ipAddresses"])

	return r
}

// flattenInstanceStateEnumSlice flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumSlice(c *Client, i interface{}) []InstanceStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceStateEnum{}
	}

	if len(a) == 0 {
		return []InstanceStateEnum{}
	}

	items := make([]InstanceStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceStateEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceStateEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceStateEnum with the same value as that string.
func flattenInstanceStateEnum(i interface{}) *InstanceStateEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceStateEnumRef("")
	}

	return InstanceStateEnumRef(s)
}

// flattenInstanceTierEnumSlice flattens the contents of InstanceTierEnum from a JSON
// response object.
func flattenInstanceTierEnumSlice(c *Client, i interface{}) []InstanceTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTierEnum{}
	}

	if len(a) == 0 {
		return []InstanceTierEnum{}
	}

	items := make([]InstanceTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTierEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceTierEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTierEnum with the same value as that string.
func flattenInstanceTierEnum(i interface{}) *InstanceTierEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceTierEnumRef("")
	}

	return InstanceTierEnumRef(s)
}

// flattenInstanceFileSharesNfsExportOptionsAccessModeEnumSlice flattens the contents of InstanceFileSharesNfsExportOptionsAccessModeEnum from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsAccessModeEnumSlice(c *Client, i interface{}) []InstanceFileSharesNfsExportOptionsAccessModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileSharesNfsExportOptionsAccessModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceFileSharesNfsExportOptionsAccessModeEnum{}
	}

	items := make([]InstanceFileSharesNfsExportOptionsAccessModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsAccessModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceFileSharesNfsExportOptionsAccessModeEnum with the same value as that string.
func flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(i interface{}) *InstanceFileSharesNfsExportOptionsAccessModeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceFileSharesNfsExportOptionsAccessModeEnumRef("")
	}

	return InstanceFileSharesNfsExportOptionsAccessModeEnumRef(s)
}

// flattenInstanceFileSharesNfsExportOptionsSquashModeEnumSlice flattens the contents of InstanceFileSharesNfsExportOptionsSquashModeEnum from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsSquashModeEnumSlice(c *Client, i interface{}) []InstanceFileSharesNfsExportOptionsSquashModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileSharesNfsExportOptionsSquashModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceFileSharesNfsExportOptionsSquashModeEnum{}
	}

	items := make([]InstanceFileSharesNfsExportOptionsSquashModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsSquashModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceFileSharesNfsExportOptionsSquashModeEnum with the same value as that string.
func flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(i interface{}) *InstanceFileSharesNfsExportOptionsSquashModeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceFileSharesNfsExportOptionsSquashModeEnumRef("")
	}

	return InstanceFileSharesNfsExportOptionsSquashModeEnumRef(s)
}

// flattenInstanceNetworksModesEnumSlice flattens the contents of InstanceNetworksModesEnum from a JSON
// response object.
func flattenInstanceNetworksModesEnumSlice(c *Client, i interface{}) []InstanceNetworksModesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworksModesEnum{}
	}

	if len(a) == 0 {
		return []InstanceNetworksModesEnum{}
	}

	items := make([]InstanceNetworksModesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworksModesEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceNetworksModesEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceNetworksModesEnum with the same value as that string.
func flattenInstanceNetworksModesEnum(i interface{}) *InstanceNetworksModesEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceNetworksModesEnumRef("")
	}

	return InstanceNetworksModesEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Instance) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstance(b, c)
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

func convertFieldDiffToInstanceDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]instanceDiff, error) {
	var diffs []instanceDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := instanceDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameToinstanceApiOperation(op, opts...)
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

func convertOpNameToinstanceApiOperation(op string, opts ...dcl.ApplyOption) (instanceApiOperation, error) {
	switch op {

	case "updateInstanceUpdateInstanceOperation":
		return &updateInstanceUpdateInstanceOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
