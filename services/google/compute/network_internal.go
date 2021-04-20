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
	"reflect"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Network) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.RoutingConfig) {
		if err := r.RoutingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NetworkRoutingConfig) validate() error {
	return nil
}

func networkGetURL(userBasePath string, r *Network) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/networks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func networkListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/networks", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func networkCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/networks", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func networkDeleteURL(userBasePath string, r *Network) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/networks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// networkApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type networkApiOperation interface {
	do(context.Context, *Network, *Client) error
}

// newUpdateNetworkUpdateRequest creates a request for an
// Network resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNetworkUpdateRequest(ctx context.Context, f *Network, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandNetworkRoutingConfig(c, f.RoutingConfig); err != nil {
		return nil, fmt.Errorf("error expanding RoutingConfig into routingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["routingConfig"] = v
	}
	return req, nil
}

// marshalUpdateNetworkUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateNetworkUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNetworkUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNetworkUpdateOperation) do(ctx context.Context, r *Network, c *Client) error {
	_, err := c.GetNetwork(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateNetworkUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateNetworkUpdateRequest(c, req)
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

func (c *Client) listNetworkRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := networkListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NetworkMaxPage {
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

type listNetworkOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNetwork(ctx context.Context, project, pageToken string, pageSize int32) ([]*Network, string, error) {
	b, err := c.listNetworkRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNetworkOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Network
	for _, v := range m.Items {
		res := flattenNetwork(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNetwork(ctx context.Context, f func(*Network) bool, resources []*Network) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNetwork(ctx, res)
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

type deleteNetworkOperation struct{}

func (op *deleteNetworkOperation) do(ctx context.Context, r *Network, c *Client) error {

	_, err := c.GetNetwork(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Network not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetNetwork checking for existence. error: %v", err)
		return err
	}

	u, err := networkDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetNetwork(ctx, r.urlNormalized())
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
type createNetworkOperation struct {
	response map[string]interface{}
}

func (op *createNetworkOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNetworkOperation) do(ctx context.Context, r *Network, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := networkCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetNetwork(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNetworkRaw(ctx context.Context, r *Network) ([]byte, error) {
	if dcl.IsZeroValue(r.AutoCreateSubnetworks) {
		r.AutoCreateSubnetworks = dcl.Bool(true)
	}

	u, err := networkGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) networkDiffsForRawDesired(ctx context.Context, rawDesired *Network, opts ...dcl.ApplyOption) (initial, desired *Network, diffs []networkDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Network
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Network); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Network, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNetwork(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Network resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Network resource: %v", err)
		}
		c.Config.Logger.Info("Found that Network resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNetworkDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Network: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Network: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNetworkInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Network: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNetworkDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Network: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNetwork(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNetworkInitialState(rawInitial, rawDesired *Network) (*Network, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNetworkDesiredState(rawDesired, rawInitial *Network, opts ...dcl.ApplyOption) (*Network, error) {

	if dcl.IsZeroValue(rawDesired.AutoCreateSubnetworks) {
		rawDesired.AutoCreateSubnetworks = dcl.Bool(true)
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.RoutingConfig = canonicalizeNetworkRoutingConfig(rawDesired.RoutingConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.GatewayIPv4, rawInitial.GatewayIPv4) {
		rawDesired.GatewayIPv4 = rawInitial.GatewayIPv4
	}
	if dcl.StringCanonicalize(rawDesired.IPv4Range, rawInitial.IPv4Range) {
		rawDesired.IPv4Range = rawInitial.IPv4Range
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.BoolCanonicalize(rawDesired.AutoCreateSubnetworks, rawInitial.AutoCreateSubnetworks) {
		rawDesired.AutoCreateSubnetworks = rawInitial.AutoCreateSubnetworks
	}
	rawDesired.RoutingConfig = canonicalizeNetworkRoutingConfig(rawDesired.RoutingConfig, rawInitial.RoutingConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}

	return rawDesired, nil
}

func canonicalizeNetworkNewState(c *Client, rawNew, rawDesired *Network) (*Network, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GatewayIPv4) && dcl.IsEmptyValueIndirect(rawDesired.GatewayIPv4) {
		rawNew.GatewayIPv4 = rawDesired.GatewayIPv4
	} else {
		if dcl.StringCanonicalize(rawDesired.GatewayIPv4, rawNew.GatewayIPv4) {
			rawNew.GatewayIPv4 = rawDesired.GatewayIPv4
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPv4Range) && dcl.IsEmptyValueIndirect(rawDesired.IPv4Range) {
		rawNew.IPv4Range = rawDesired.IPv4Range
	} else {
		if dcl.StringCanonicalize(rawDesired.IPv4Range, rawNew.IPv4Range) {
			rawNew.IPv4Range = rawDesired.IPv4Range
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AutoCreateSubnetworks) && dcl.IsEmptyValueIndirect(rawDesired.AutoCreateSubnetworks) {
		rawNew.AutoCreateSubnetworks = rawDesired.AutoCreateSubnetworks
	} else {
		if dcl.BoolCanonicalize(rawDesired.AutoCreateSubnetworks, rawNew.AutoCreateSubnetworks) {
			rawNew.AutoCreateSubnetworks = rawDesired.AutoCreateSubnetworks
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RoutingConfig) && dcl.IsEmptyValueIndirect(rawDesired.RoutingConfig) {
		rawNew.RoutingConfig = rawDesired.RoutingConfig
	} else {
		rawNew.RoutingConfig = canonicalizeNewNetworkRoutingConfig(c, rawDesired.RoutingConfig, rawNew.RoutingConfig)
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	return rawNew, nil
}

func canonicalizeNetworkRoutingConfig(des, initial *NetworkRoutingConfig, opts ...dcl.ApplyOption) *NetworkRoutingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RoutingMode) {
		des.RoutingMode = initial.RoutingMode
	}

	return des
}

func canonicalizeNewNetworkRoutingConfig(c *Client, des, nw *NetworkRoutingConfig) *NetworkRoutingConfig {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewNetworkRoutingConfigSet(c *Client, des, nw []NetworkRoutingConfig) []NetworkRoutingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []NetworkRoutingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNetworkRoutingConfig(c, &d, &n) {
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

func canonicalizeNewNetworkRoutingConfigSlice(c *Client, des, nw []NetworkRoutingConfig) []NetworkRoutingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NetworkRoutingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNetworkRoutingConfig(c, &d, &n))
	}

	return items
}

type networkDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         networkApiOperation
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
func diffNetwork(c *Client, desired, actual *Network, opts ...dcl.ApplyOption) ([]networkDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []networkDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "description"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.GatewayIPv4, actual.GatewayIPv4, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "", FieldName: "gateway_ipv4"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.IPv4Range, actual.IPv4Range, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "", FieldName: "ipv4_range"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.AutoCreateSubnetworks, actual.AutoCreateSubnetworks, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "auto_create_subnetworks"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "project"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "", FieldName: "self_link"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkDiff{RequiresRecreate: true, Diffs: ds})
	}

	if compareNetworkRoutingConfig(c, desired.RoutingConfig, actual.RoutingConfig) {
		c.Config.Logger.Infof("Detected diff in RoutingConfig.\nDESIRED: %v\nACTUAL: %v", desired.RoutingConfig, actual.RoutingConfig)

		diffs = append(diffs, networkDiff{
			UpdateOp:  &updateNetworkUpdateOperation{},
			FieldName: "RoutingConfig",
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
	var deduped []networkDiff
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
func compareNetworkRoutingConfig(c *Client, desired, actual *NetworkRoutingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.RoutingMode, actual.RoutingMode) && !dcl.IsZeroValue(desired.RoutingMode) {
		c.Config.Logger.Infof("Diff in RoutingMode.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RoutingMode), dcl.SprintResource(actual.RoutingMode))
		return true
	}
	return false
}

func compareNetworkRoutingConfigSlice(c *Client, desired, actual []NetworkRoutingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkRoutingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNetworkRoutingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NetworkRoutingConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNetworkRoutingConfigMap(c *Client, desired, actual map[string]NetworkRoutingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkRoutingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NetworkRoutingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNetworkRoutingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NetworkRoutingConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNetworkRoutingConfigRoutingModeEnumSlice(c *Client, desired, actual []NetworkRoutingConfigRoutingModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkRoutingConfigRoutingModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNetworkRoutingConfigRoutingModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NetworkRoutingConfigRoutingModeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNetworkRoutingConfigRoutingModeEnum(c *Client, desired, actual *NetworkRoutingConfigRoutingModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Network) urlNormalized() *Network {
	normalized := dcl.Copy(*r).(Network)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.GatewayIPv4 = dcl.SelfLinkToName(r.GatewayIPv4)
	normalized.IPv4Range = dcl.SelfLinkToName(r.IPv4Range)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *Network) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Network) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Network) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Network) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/networks/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Network resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Network) marshal(c *Client) ([]byte, error) {
	m, err := expandNetwork(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Network: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNetwork decodes JSON responses into the Network resource schema.
func unmarshalNetwork(b []byte, c *Client) (*Network, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNetwork(m, c)
}

func unmarshalMapNetwork(m map[string]interface{}, c *Client) (*Network, error) {

	return flattenNetwork(c, m), nil
}

// expandNetwork expands Network into a JSON request object.
func expandNetwork(c *Client, f *Network) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.GatewayIPv4; !dcl.IsEmptyValueIndirect(v) {
		m["gatewayIPv4"] = v
	}
	if v := f.IPv4Range; !dcl.IsEmptyValueIndirect(v) {
		m["ipv4Range"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	m["autoCreateSubnetworks"] = f.AutoCreateSubnetworks
	if v, err := expandNetworkRoutingConfig(c, f.RoutingConfig); err != nil {
		return nil, fmt.Errorf("error expanding RoutingConfig into routingConfig: %w", err)
	} else if v != nil {
		m["routingConfig"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}

	return m, nil
}

// flattenNetwork flattens Network from a JSON request object into the
// Network type.
func flattenNetwork(c *Client, i interface{}) *Network {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Network{}
	r.Description = dcl.FlattenString(m["description"])
	r.GatewayIPv4 = dcl.FlattenString(m["gatewayIPv4"])
	r.IPv4Range = dcl.FlattenString(m["ipv4Range"])
	r.Name = dcl.FlattenString(m["name"])
	r.AutoCreateSubnetworks = dcl.FlattenBool(m["autoCreateSubnetworks"])
	if _, ok := m["autoCreateSubnetworks"]; !ok {
		c.Config.Logger.Info("Using default value for autoCreateSubnetworks")
		r.AutoCreateSubnetworks = dcl.Bool(true)
	}
	r.RoutingConfig = flattenNetworkRoutingConfig(c, m["routingConfig"])
	r.Project = dcl.FlattenString(m["project"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])

	return r
}

// expandNetworkRoutingConfigMap expands the contents of NetworkRoutingConfig into a JSON
// request object.
func expandNetworkRoutingConfigMap(c *Client, f map[string]NetworkRoutingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNetworkRoutingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNetworkRoutingConfigSlice expands the contents of NetworkRoutingConfig into a JSON
// request object.
func expandNetworkRoutingConfigSlice(c *Client, f []NetworkRoutingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNetworkRoutingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNetworkRoutingConfigMap flattens the contents of NetworkRoutingConfig from a JSON
// response object.
func flattenNetworkRoutingConfigMap(c *Client, i interface{}) map[string]NetworkRoutingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NetworkRoutingConfig{}
	}

	if len(a) == 0 {
		return map[string]NetworkRoutingConfig{}
	}

	items := make(map[string]NetworkRoutingConfig)
	for k, item := range a {
		items[k] = *flattenNetworkRoutingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNetworkRoutingConfigSlice flattens the contents of NetworkRoutingConfig from a JSON
// response object.
func flattenNetworkRoutingConfigSlice(c *Client, i interface{}) []NetworkRoutingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []NetworkRoutingConfig{}
	}

	if len(a) == 0 {
		return []NetworkRoutingConfig{}
	}

	items := make([]NetworkRoutingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNetworkRoutingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandNetworkRoutingConfig expands an instance of NetworkRoutingConfig into a JSON
// request object.
func expandNetworkRoutingConfig(c *Client, f *NetworkRoutingConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.RoutingMode; !dcl.IsEmptyValueIndirect(v) {
		m["routingMode"] = v
	}

	return m, nil
}

// flattenNetworkRoutingConfig flattens an instance of NetworkRoutingConfig from a JSON
// response object.
func flattenNetworkRoutingConfig(c *Client, i interface{}) *NetworkRoutingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NetworkRoutingConfig{}
	r.RoutingMode = flattenNetworkRoutingConfigRoutingModeEnum(m["routingMode"])

	return r
}

// flattenNetworkRoutingConfigRoutingModeEnumSlice flattens the contents of NetworkRoutingConfigRoutingModeEnum from a JSON
// response object.
func flattenNetworkRoutingConfigRoutingModeEnumSlice(c *Client, i interface{}) []NetworkRoutingConfigRoutingModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NetworkRoutingConfigRoutingModeEnum{}
	}

	if len(a) == 0 {
		return []NetworkRoutingConfigRoutingModeEnum{}
	}

	items := make([]NetworkRoutingConfigRoutingModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNetworkRoutingConfigRoutingModeEnum(item.(interface{})))
	}

	return items
}

// flattenNetworkRoutingConfigRoutingModeEnum asserts that an interface is a string, and returns a
// pointer to a *NetworkRoutingConfigRoutingModeEnum with the same value as that string.
func flattenNetworkRoutingConfigRoutingModeEnum(i interface{}) *NetworkRoutingConfigRoutingModeEnum {
	s, ok := i.(string)
	if !ok {
		return NetworkRoutingConfigRoutingModeEnumRef("")
	}

	return NetworkRoutingConfigRoutingModeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Network) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNetwork(b, c)
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
