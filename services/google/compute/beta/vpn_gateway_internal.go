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

func (r *VpnGateway) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "network"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}
func (r *VpnGatewayVpnInterface) validate() error {
	return nil
}

func vpnGatewayGetURL(userBasePath string, r *VpnGateway) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func vpnGatewayListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnGateways", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func vpnGatewayCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnGateways", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func vpnGatewayDeleteURL(userBasePath string, r *VpnGateway) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// vpnGatewayApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type vpnGatewayApiOperation interface {
	do(context.Context, *VpnGateway, *Client) error
}

// newUpdateVpnGatewaySetLabelsRequest creates a request for an
// VpnGateway resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateVpnGatewaySetLabelsRequest(ctx context.Context, f *VpnGateway, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	b, err := c.getVpnGatewayRaw(ctx, f.urlNormalized())
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawLabelFingerprint, err := dcl.GetMapEntry(
		m,
		[]string{"labelFingerprint"},
	)
	if err != nil {
		c.Config.Logger.Warningf("Failed to fetch from JSON Path: %v", err)
	} else {
		req["labelFingerprint"] = rawLabelFingerprint.(string)
	}
	return req, nil
}

// marshalUpdateVpnGatewaySetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateVpnGatewaySetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateVpnGatewaySetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateVpnGatewaySetLabelsOperation) do(ctx context.Context, r *VpnGateway, c *Client) error {
	_, err := c.GetVpnGateway(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateVpnGatewaySetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateVpnGatewaySetLabelsRequest(c, req)
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
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/compute/beta/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listVpnGatewayRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := vpnGatewayListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != VpnGatewayMaxPage {
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

type listVpnGatewayOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listVpnGateway(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*VpnGateway, string, error) {
	b, err := c.listVpnGatewayRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listVpnGatewayOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*VpnGateway
	for _, v := range m.Items {
		res := flattenVpnGateway(c, v)
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllVpnGateway(ctx context.Context, f func(*VpnGateway) bool, resources []*VpnGateway) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteVpnGateway(ctx, res)
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

type deleteVpnGatewayOperation struct{}

func (op *deleteVpnGatewayOperation) do(ctx context.Context, r *VpnGateway, c *Client) error {

	_, err := c.GetVpnGateway(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("VpnGateway not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetVpnGateway checking for existence. error: %v", err)
		return err
	}

	u, err := vpnGatewayDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetVpnGateway(ctx, r.urlNormalized())
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
type createVpnGatewayOperation struct {
	response map[string]interface{}
}

func (op *createVpnGatewayOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createVpnGatewayOperation) do(ctx context.Context, r *VpnGateway, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := vpnGatewayCreateURL(c.Config.BasePath, project, region)

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

	if _, err := c.GetVpnGateway(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getVpnGatewayRaw(ctx context.Context, r *VpnGateway) ([]byte, error) {

	u, err := vpnGatewayGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) vpnGatewayDiffsForRawDesired(ctx context.Context, rawDesired *VpnGateway, opts ...dcl.ApplyOption) (initial, desired *VpnGateway, diffs []vpnGatewayDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *VpnGateway
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*VpnGateway); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected VpnGateway, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetVpnGateway(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a VpnGateway resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve VpnGateway resource: %v", err)
		}
		c.Config.Logger.Info("Found that VpnGateway resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeVpnGatewayDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for VpnGateway: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for VpnGateway: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeVpnGatewayInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for VpnGateway: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeVpnGatewayDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for VpnGateway: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffVpnGateway(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeVpnGatewayInitialState(rawInitial, rawDesired *VpnGateway) (*VpnGateway, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeVpnGatewayDesiredState(rawDesired, rawInitial *VpnGateway, opts ...dcl.ApplyOption) (*VpnGateway, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}

	return rawDesired, nil
}

func canonicalizeVpnGatewayNewState(c *Client, rawNew, rawDesired *VpnGateway) (*VpnGateway, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VpnInterface) && dcl.IsEmptyValueIndirect(rawDesired.VpnInterface) {
		rawNew.VpnInterface = rawDesired.VpnInterface
	} else {
		rawNew.VpnInterface = canonicalizeNewVpnGatewayVpnInterfaceSlice(c, rawDesired.VpnInterface, rawNew.VpnInterface)
	}

	return rawNew, nil
}

func canonicalizeVpnGatewayVpnInterface(des, initial *VpnGatewayVpnInterface, opts ...dcl.ApplyOption) *VpnGatewayVpnInterface {
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

func canonicalizeNewVpnGatewayVpnInterface(c *Client, des, nw *VpnGatewayVpnInterface) *VpnGatewayVpnInterface {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.IPAddress, nw.IPAddress) {
		nw.IPAddress = des.IPAddress
	}

	return nw
}

func canonicalizeNewVpnGatewayVpnInterfaceSet(c *Client, des, nw []VpnGatewayVpnInterface) []VpnGatewayVpnInterface {
	if des == nil {
		return nw
	}
	var reorderedNew []VpnGatewayVpnInterface
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareVpnGatewayVpnInterfaceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewVpnGatewayVpnInterfaceSlice(c *Client, des, nw []VpnGatewayVpnInterface) []VpnGatewayVpnInterface {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VpnGatewayVpnInterface
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVpnGatewayVpnInterface(c, &d, &n))
	}

	return items
}

type vpnGatewayDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         vpnGatewayApiOperation
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
func diffVpnGateway(c *Client, desired, actual *VpnGateway, opts ...dcl.ApplyOption) ([]vpnGatewayDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []vpnGatewayDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.VpnInterface, actual.VpnInterface, dcl.Info{OutputOnly: true, ObjectFunction: compareVpnGatewayVpnInterfaceNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VpnInterface")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToVpnGatewayDiff(ds, opts...)
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
	var deduped []vpnGatewayDiff
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
func compareVpnGatewayVpnInterfaceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VpnGatewayVpnInterface)
	if !ok {
		desiredNotPointer, ok := d.(VpnGatewayVpnInterface)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VpnGatewayVpnInterface or *VpnGatewayVpnInterface", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VpnGatewayVpnInterface)
	if !ok {
		actualNotPointer, ok := a.(VpnGatewayVpnInterface)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VpnGatewayVpnInterface", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAddress, actual.IPAddress, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPAddress")); len(ds) != 0 || err != nil {
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
func (r *VpnGateway) urlNormalized() *VpnGateway {
	normalized := dcl.Copy(*r).(VpnGateway)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *VpnGateway) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *VpnGateway) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *VpnGateway) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *VpnGateway) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the VpnGateway resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *VpnGateway) marshal(c *Client) ([]byte, error) {
	m, err := expandVpnGateway(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling VpnGateway: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalVpnGateway decodes JSON responses into the VpnGateway resource schema.
func unmarshalVpnGateway(b []byte, c *Client) (*VpnGateway, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapVpnGateway(m, c)
}

func unmarshalMapVpnGateway(m map[string]interface{}, c *Client) (*VpnGateway, error) {

	return flattenVpnGateway(c, m), nil
}

// expandVpnGateway expands VpnGateway into a JSON request object.
func expandVpnGateway(c *Client, f *VpnGateway) (map[string]interface{}, error) {
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
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/global/networks/%s", f.Network, f.Project, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if v != nil {
		m["network"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandVpnGatewayVpnInterfaceSlice(c, f.VpnInterface); err != nil {
		return nil, fmt.Errorf("error expanding VpnInterface into vpnInterfaces: %w", err)
	} else if v != nil {
		m["vpnInterfaces"] = v
	}

	return m, nil
}

// flattenVpnGateway flattens VpnGateway from a JSON request object into the
// VpnGateway type.
func flattenVpnGateway(c *Client, i interface{}) *VpnGateway {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &VpnGateway{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Region = dcl.FlattenString(m["region"])
	res.Network = dcl.FlattenString(m["network"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Project = dcl.FlattenString(m["project"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.VpnInterface = flattenVpnGatewayVpnInterfaceSlice(c, m["vpnInterfaces"])

	return res
}

// expandVpnGatewayVpnInterfaceMap expands the contents of VpnGatewayVpnInterface into a JSON
// request object.
func expandVpnGatewayVpnInterfaceMap(c *Client, f map[string]VpnGatewayVpnInterface) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVpnGatewayVpnInterface(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVpnGatewayVpnInterfaceSlice expands the contents of VpnGatewayVpnInterface into a JSON
// request object.
func expandVpnGatewayVpnInterfaceSlice(c *Client, f []VpnGatewayVpnInterface) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVpnGatewayVpnInterface(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVpnGatewayVpnInterfaceMap flattens the contents of VpnGatewayVpnInterface from a JSON
// response object.
func flattenVpnGatewayVpnInterfaceMap(c *Client, i interface{}) map[string]VpnGatewayVpnInterface {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VpnGatewayVpnInterface{}
	}

	if len(a) == 0 {
		return map[string]VpnGatewayVpnInterface{}
	}

	items := make(map[string]VpnGatewayVpnInterface)
	for k, item := range a {
		items[k] = *flattenVpnGatewayVpnInterface(c, item.(map[string]interface{}))
	}

	return items
}

// flattenVpnGatewayVpnInterfaceSlice flattens the contents of VpnGatewayVpnInterface from a JSON
// response object.
func flattenVpnGatewayVpnInterfaceSlice(c *Client, i interface{}) []VpnGatewayVpnInterface {
	a, ok := i.([]interface{})
	if !ok {
		return []VpnGatewayVpnInterface{}
	}

	if len(a) == 0 {
		return []VpnGatewayVpnInterface{}
	}

	items := make([]VpnGatewayVpnInterface, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVpnGatewayVpnInterface(c, item.(map[string]interface{})))
	}

	return items
}

// expandVpnGatewayVpnInterface expands an instance of VpnGatewayVpnInterface into a JSON
// request object.
func expandVpnGatewayVpnInterface(c *Client, f *VpnGatewayVpnInterface) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}

	return m, nil
}

// flattenVpnGatewayVpnInterface flattens an instance of VpnGatewayVpnInterface from a JSON
// response object.
func flattenVpnGatewayVpnInterface(c *Client, i interface{}) *VpnGatewayVpnInterface {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VpnGatewayVpnInterface{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.IPAddress = dcl.FlattenString(m["ipAddress"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *VpnGateway) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalVpnGateway(b, c)
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
		if nr.Region == nil && ncr.Region == nil {
			c.Config.Logger.Info("Both Region fields null - considering equal.")
		} else if nr.Region == nil || ncr.Region == nil {
			c.Config.Logger.Info("Only one Region field is null - considering unequal.")
			return false
		} else if *nr.Region != *ncr.Region {
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

func convertFieldDiffToVpnGatewayDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]vpnGatewayDiff, error) {
	var diffs []vpnGatewayDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := vpnGatewayDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTovpnGatewayApiOperation(op, opts...)
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

func convertOpNameTovpnGatewayApiOperation(op string, opts ...dcl.ApplyOption) (vpnGatewayApiOperation, error) {
	switch op {

	case "updateVpnGatewaySetLabelsOperation":
		return &updateVpnGatewaySetLabelsOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
