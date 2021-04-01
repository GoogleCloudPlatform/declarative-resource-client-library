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
	"reflect"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Router) validate() error {

	if !dcl.IsEmptyValueIndirect(r.Bgp) {
		if err := r.Bgp.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RouterNats) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LogConfig) {
		if err := r.LogConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RouterNatsLogConfig) validate() error {
	return nil
}
func (r *RouterNatsSubnetworks) validate() error {
	return nil
}
func (r *RouterInterfaces) validate() error {
	return nil
}
func (r *RouterBgpPeers) validate() error {
	return nil
}
func (r *RouterBgpPeersAdvertisedIPRanges) validate() error {
	return nil
}
func (r *RouterBgp) validate() error {
	return nil
}
func (r *RouterBgpAdvertisedIPRanges) validate() error {
	return nil
}

func routerGetURL(userBasePath string, r *Router) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func routerListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func routerCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func routerDeleteURL(userBasePath string, r *Router) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// routerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type routerApiOperation interface {
	do(context.Context, *Router, *Client) error
}

// newUpdateRouterUpdateRequest creates a request for an
// Router resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateRouterUpdateRequest(ctx context.Context, f *Router, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandRouterNatsSlice(c, f.Nats); err != nil {
		return nil, fmt.Errorf("error expanding Nats into nats: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["nats"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandRouterInterfacesSlice(c, f.Interfaces); err != nil {
		return nil, fmt.Errorf("error expanding Interfaces into interfaces: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["interfaces"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandRouterBgpPeersSlice(c, f.BgpPeers); err != nil {
		return nil, fmt.Errorf("error expanding BgpPeers into bgpPeers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["bgpPeers"] = v
	}
	if v, err := expandRouterBgp(c, f.Bgp); err != nil {
		return nil, fmt.Errorf("error expanding Bgp into bgp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["bgp"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		req["region"] = v
	}
	return req, nil
}

// marshalUpdateRouterUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateRouterUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateRouterUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateRouterUpdateOperation) do(ctx context.Context, r *Router, c *Client) error {
	_, err := c.GetRouter(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateRouterUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateRouterUpdateRequest(c, req)
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

func (c *Client) listRouterRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := routerListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RouterMaxPage {
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

type listRouterOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listRouter(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*Router, string, error) {
	b, err := c.listRouterRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRouterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Router
	for _, v := range m.Items {
		res := flattenRouter(c, v)
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRouter(ctx context.Context, f func(*Router) bool, resources []*Router) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRouter(ctx, res)
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

type deleteRouterOperation struct{}

func (op *deleteRouterOperation) do(ctx context.Context, r *Router, c *Client) error {

	_, err := c.GetRouter(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Router not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetRouter checking for existence. error: %v", err)
		return err
	}

	u, err := routerDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetRouter(ctx, r.urlNormalized())
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
type createRouterOperation struct {
	response map[string]interface{}
}

func (op *createRouterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRouterOperation) do(ctx context.Context, r *Router, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := routerCreateURL(c.Config.BasePath, project, region)

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

	if _, err := c.GetRouter(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRouterRaw(ctx context.Context, r *Router) ([]byte, error) {

	u, err := routerGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) routerDiffsForRawDesired(ctx context.Context, rawDesired *Router, opts ...dcl.ApplyOption) (initial, desired *Router, diffs []routerDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Router
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Router); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Router, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRouter(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Router resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Router resource: %v", err)
		}
		c.Config.Logger.Info("Found that Router resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRouterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Router: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Router: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRouterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Router: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRouterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Router: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRouter(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeRouterInitialState(rawInitial, rawDesired *Router) (*Router, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRouterDesiredState(rawDesired, rawInitial *Router, opts ...dcl.ApplyOption) (*Router, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Bgp = canonicalizeRouterBgp(rawDesired.Bgp, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.CreationTimestamp) {
		rawDesired.CreationTimestamp = rawInitial.CreationTimestamp
	}
	if dcl.IsZeroValue(rawDesired.Nats) {
		rawDesired.Nats = rawInitial.Nats
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.Interfaces) {
		rawDesired.Interfaces = rawInitial.Interfaces
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.BgpPeers) {
		rawDesired.BgpPeers = rawInitial.BgpPeers
	}
	rawDesired.Bgp = canonicalizeRouterBgp(rawDesired.Bgp, rawInitial.Bgp, opts...)
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}

	return rawDesired, nil
}

func canonicalizeRouterNewState(c *Client, rawNew, rawDesired *Router) (*Router, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Nats) && dcl.IsEmptyValueIndirect(rawDesired.Nats) {
		rawNew.Nats = rawDesired.Nats
	} else {
		rawNew.Nats = canonicalizeNewRouterNatsSlice(c, rawDesired.Nats, rawNew.Nats)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Interfaces) && dcl.IsEmptyValueIndirect(rawDesired.Interfaces) {
		rawNew.Interfaces = rawDesired.Interfaces
	} else {
		rawNew.Interfaces = canonicalizeNewRouterInterfacesSlice(c, rawDesired.Interfaces, rawNew.Interfaces)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.BgpPeers) && dcl.IsEmptyValueIndirect(rawDesired.BgpPeers) {
		rawNew.BgpPeers = rawDesired.BgpPeers
	} else {
		rawNew.BgpPeers = canonicalizeNewRouterBgpPeersSlice(c, rawDesired.BgpPeers, rawNew.BgpPeers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Bgp) && dcl.IsEmptyValueIndirect(rawDesired.Bgp) {
		rawNew.Bgp = rawDesired.Bgp
	} else {
		rawNew.Bgp = canonicalizeNewRouterBgp(c, rawDesired.Bgp, rawNew.Bgp)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.NameToSelfLink(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
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

func canonicalizeRouterNats(des, initial *RouterNats, opts ...dcl.ApplyOption) *RouterNats {
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
	des.LogConfig = canonicalizeRouterNatsLogConfig(des.LogConfig, initial.LogConfig, opts...)
	if dcl.IsZeroValue(des.SourceSubnetworkIPRangesToNat) {
		des.SourceSubnetworkIPRangesToNat = initial.SourceSubnetworkIPRangesToNat
	}
	if dcl.IsZeroValue(des.NatIps) {
		des.NatIps = initial.NatIps
	}
	if dcl.IsZeroValue(des.DrainNatIps) {
		des.DrainNatIps = initial.DrainNatIps
	}
	if dcl.IsZeroValue(des.NatIPAllocateOption) {
		des.NatIPAllocateOption = initial.NatIPAllocateOption
	}
	if dcl.IsZeroValue(des.MinPortsPerVm) {
		des.MinPortsPerVm = initial.MinPortsPerVm
	}
	if dcl.IsZeroValue(des.UdpIdleTimeoutSec) {
		des.UdpIdleTimeoutSec = initial.UdpIdleTimeoutSec
	}
	if dcl.IsZeroValue(des.IcmpIdleTimeoutSec) {
		des.IcmpIdleTimeoutSec = initial.IcmpIdleTimeoutSec
	}
	if dcl.IsZeroValue(des.TcpEstablishedIdleTimeoutSec) {
		des.TcpEstablishedIdleTimeoutSec = initial.TcpEstablishedIdleTimeoutSec
	}
	if dcl.IsZeroValue(des.TcpTransitoryIdleTimeoutSec) {
		des.TcpTransitoryIdleTimeoutSec = initial.TcpTransitoryIdleTimeoutSec
	}
	if dcl.IsZeroValue(des.Subnetworks) {
		des.Subnetworks = initial.Subnetworks
	}

	return des
}

func canonicalizeNewRouterNats(c *Client, des, nw *RouterNats) *RouterNats {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.LogConfig = canonicalizeNewRouterNatsLogConfig(c, des.LogConfig, nw.LogConfig)
	nw.Subnetworks = canonicalizeNewRouterNatsSubnetworksSlice(c, des.Subnetworks, nw.Subnetworks)

	return nw
}

func canonicalizeNewRouterNatsSet(c *Client, des, nw []RouterNats) []RouterNats {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterNats
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterNats(c, &d, &n) {
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

func canonicalizeNewRouterNatsSlice(c *Client, des, nw []RouterNats) []RouterNats {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterNats
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterNats(c, &d, &n))
	}

	return items
}

func canonicalizeRouterNatsLogConfig(des, initial *RouterNatsLogConfig, opts ...dcl.ApplyOption) *RouterNatsLogConfig {
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
	if dcl.IsZeroValue(des.Filter) {
		des.Filter = initial.Filter
	}

	return des
}

func canonicalizeNewRouterNatsLogConfig(c *Client, des, nw *RouterNatsLogConfig) *RouterNatsLogConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enable, nw.Enable) {
		nw.Enable = des.Enable
	}

	return nw
}

func canonicalizeNewRouterNatsLogConfigSet(c *Client, des, nw []RouterNatsLogConfig) []RouterNatsLogConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterNatsLogConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterNatsLogConfig(c, &d, &n) {
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

func canonicalizeNewRouterNatsLogConfigSlice(c *Client, des, nw []RouterNatsLogConfig) []RouterNatsLogConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterNatsLogConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterNatsLogConfig(c, &d, &n))
	}

	return items
}

func canonicalizeRouterNatsSubnetworks(des, initial *RouterNatsSubnetworks, opts ...dcl.ApplyOption) *RouterNatsSubnetworks {
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
	if dcl.StringCanonicalize(des.SourceIPRangesToNat, initial.SourceIPRangesToNat) || dcl.IsZeroValue(des.SourceIPRangesToNat) {
		des.SourceIPRangesToNat = initial.SourceIPRangesToNat
	}
	if dcl.StringCanonicalize(des.SecondaryIPRangeNames, initial.SecondaryIPRangeNames) || dcl.IsZeroValue(des.SecondaryIPRangeNames) {
		des.SecondaryIPRangeNames = initial.SecondaryIPRangeNames
	}

	return des
}

func canonicalizeNewRouterNatsSubnetworks(c *Client, des, nw *RouterNatsSubnetworks) *RouterNatsSubnetworks {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.SourceIPRangesToNat, nw.SourceIPRangesToNat) {
		nw.SourceIPRangesToNat = des.SourceIPRangesToNat
	}
	if dcl.StringCanonicalize(des.SecondaryIPRangeNames, nw.SecondaryIPRangeNames) {
		nw.SecondaryIPRangeNames = des.SecondaryIPRangeNames
	}

	return nw
}

func canonicalizeNewRouterNatsSubnetworksSet(c *Client, des, nw []RouterNatsSubnetworks) []RouterNatsSubnetworks {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterNatsSubnetworks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterNatsSubnetworks(c, &d, &n) {
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

func canonicalizeNewRouterNatsSubnetworksSlice(c *Client, des, nw []RouterNatsSubnetworks) []RouterNatsSubnetworks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterNatsSubnetworks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterNatsSubnetworks(c, &d, &n))
	}

	return items
}

func canonicalizeRouterInterfaces(des, initial *RouterInterfaces, opts ...dcl.ApplyOption) *RouterInterfaces {
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
	if dcl.NameToSelfLink(des.LinkedVpnTunnel, initial.LinkedVpnTunnel) || dcl.IsZeroValue(des.LinkedVpnTunnel) {
		des.LinkedVpnTunnel = initial.LinkedVpnTunnel
	}
	if dcl.StringCanonicalize(des.IPRange, initial.IPRange) || dcl.IsZeroValue(des.IPRange) {
		des.IPRange = initial.IPRange
	}
	if dcl.IsZeroValue(des.ManagementType) {
		des.ManagementType = initial.ManagementType
	}

	return des
}

func canonicalizeNewRouterInterfaces(c *Client, des, nw *RouterInterfaces) *RouterInterfaces {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.NameToSelfLink(des.LinkedVpnTunnel, nw.LinkedVpnTunnel) {
		nw.LinkedVpnTunnel = des.LinkedVpnTunnel
	}
	if dcl.StringCanonicalize(des.IPRange, nw.IPRange) {
		nw.IPRange = des.IPRange
	}

	return nw
}

func canonicalizeNewRouterInterfacesSet(c *Client, des, nw []RouterInterfaces) []RouterInterfaces {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterInterfaces
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterInterfaces(c, &d, &n) {
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

func canonicalizeNewRouterInterfacesSlice(c *Client, des, nw []RouterInterfaces) []RouterInterfaces {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterInterfaces
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterInterfaces(c, &d, &n))
	}

	return items
}

func canonicalizeRouterBgpPeers(des, initial *RouterBgpPeers, opts ...dcl.ApplyOption) *RouterBgpPeers {
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
	if dcl.StringCanonicalize(des.InterfaceName, initial.InterfaceName) || dcl.IsZeroValue(des.InterfaceName) {
		des.InterfaceName = initial.InterfaceName
	}
	if dcl.StringCanonicalize(des.IPAddress, initial.IPAddress) || dcl.IsZeroValue(des.IPAddress) {
		des.IPAddress = initial.IPAddress
	}
	if dcl.StringCanonicalize(des.PeerIPAddress, initial.PeerIPAddress) || dcl.IsZeroValue(des.PeerIPAddress) {
		des.PeerIPAddress = initial.PeerIPAddress
	}
	if dcl.IsZeroValue(des.PeerAsn) {
		des.PeerAsn = initial.PeerAsn
	}
	if dcl.IsZeroValue(des.AdvertisedRoutePriority) {
		des.AdvertisedRoutePriority = initial.AdvertisedRoutePriority
	}
	if dcl.StringCanonicalize(des.AdvertiseMode, initial.AdvertiseMode) || dcl.IsZeroValue(des.AdvertiseMode) {
		des.AdvertiseMode = initial.AdvertiseMode
	}
	if dcl.StringCanonicalize(des.ManagementType, initial.ManagementType) || dcl.IsZeroValue(des.ManagementType) {
		des.ManagementType = initial.ManagementType
	}
	if dcl.IsZeroValue(des.AdvertisedGroups) {
		des.AdvertisedGroups = initial.AdvertisedGroups
	}
	if dcl.IsZeroValue(des.AdvertisedIPRanges) {
		des.AdvertisedIPRanges = initial.AdvertisedIPRanges
	}

	return des
}

func canonicalizeNewRouterBgpPeers(c *Client, des, nw *RouterBgpPeers) *RouterBgpPeers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.InterfaceName, nw.InterfaceName) {
		nw.InterfaceName = des.InterfaceName
	}
	if dcl.StringCanonicalize(des.IPAddress, nw.IPAddress) {
		nw.IPAddress = des.IPAddress
	}
	if dcl.StringCanonicalize(des.PeerIPAddress, nw.PeerIPAddress) {
		nw.PeerIPAddress = des.PeerIPAddress
	}
	if dcl.StringCanonicalize(des.AdvertiseMode, nw.AdvertiseMode) {
		nw.AdvertiseMode = des.AdvertiseMode
	}
	if dcl.StringCanonicalize(des.ManagementType, nw.ManagementType) {
		nw.ManagementType = des.ManagementType
	}
	nw.AdvertisedIPRanges = canonicalizeNewRouterBgpPeersAdvertisedIPRangesSlice(c, des.AdvertisedIPRanges, nw.AdvertisedIPRanges)

	return nw
}

func canonicalizeNewRouterBgpPeersSet(c *Client, des, nw []RouterBgpPeers) []RouterBgpPeers {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterBgpPeers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterBgpPeers(c, &d, &n) {
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

func canonicalizeNewRouterBgpPeersSlice(c *Client, des, nw []RouterBgpPeers) []RouterBgpPeers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterBgpPeers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterBgpPeers(c, &d, &n))
	}

	return items
}

func canonicalizeRouterBgpPeersAdvertisedIPRanges(des, initial *RouterBgpPeersAdvertisedIPRanges, opts ...dcl.ApplyOption) *RouterBgpPeersAdvertisedIPRanges {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Range, initial.Range) || dcl.IsZeroValue(des.Range) {
		des.Range = initial.Range
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewRouterBgpPeersAdvertisedIPRanges(c *Client, des, nw *RouterBgpPeersAdvertisedIPRanges) *RouterBgpPeersAdvertisedIPRanges {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Range, nw.Range) {
		nw.Range = des.Range
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewRouterBgpPeersAdvertisedIPRangesSet(c *Client, des, nw []RouterBgpPeersAdvertisedIPRanges) []RouterBgpPeersAdvertisedIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterBgpPeersAdvertisedIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterBgpPeersAdvertisedIPRanges(c, &d, &n) {
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

func canonicalizeNewRouterBgpPeersAdvertisedIPRangesSlice(c *Client, des, nw []RouterBgpPeersAdvertisedIPRanges) []RouterBgpPeersAdvertisedIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterBgpPeersAdvertisedIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterBgpPeersAdvertisedIPRanges(c, &d, &n))
	}

	return items
}

func canonicalizeRouterBgp(des, initial *RouterBgp, opts ...dcl.ApplyOption) *RouterBgp {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Asn) {
		des.Asn = initial.Asn
	}
	if dcl.IsZeroValue(des.AdvertiseMode) {
		des.AdvertiseMode = initial.AdvertiseMode
	}
	if dcl.IsZeroValue(des.AdvertisedGroups) {
		des.AdvertisedGroups = initial.AdvertisedGroups
	}
	if dcl.IsZeroValue(des.AdvertisedIPRanges) {
		des.AdvertisedIPRanges = initial.AdvertisedIPRanges
	}

	return des
}

func canonicalizeNewRouterBgp(c *Client, des, nw *RouterBgp) *RouterBgp {
	if des == nil || nw == nil {
		return nw
	}

	nw.AdvertisedIPRanges = canonicalizeNewRouterBgpAdvertisedIPRangesSlice(c, des.AdvertisedIPRanges, nw.AdvertisedIPRanges)

	return nw
}

func canonicalizeNewRouterBgpSet(c *Client, des, nw []RouterBgp) []RouterBgp {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterBgp
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterBgp(c, &d, &n) {
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

func canonicalizeNewRouterBgpSlice(c *Client, des, nw []RouterBgp) []RouterBgp {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterBgp
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterBgp(c, &d, &n))
	}

	return items
}

func canonicalizeRouterBgpAdvertisedIPRanges(des, initial *RouterBgpAdvertisedIPRanges, opts ...dcl.ApplyOption) *RouterBgpAdvertisedIPRanges {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Range, initial.Range) || dcl.IsZeroValue(des.Range) {
		des.Range = initial.Range
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}

	return des
}

func canonicalizeNewRouterBgpAdvertisedIPRanges(c *Client, des, nw *RouterBgpAdvertisedIPRanges) *RouterBgpAdvertisedIPRanges {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Range, nw.Range) {
		nw.Range = des.Range
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewRouterBgpAdvertisedIPRangesSet(c *Client, des, nw []RouterBgpAdvertisedIPRanges) []RouterBgpAdvertisedIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterBgpAdvertisedIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareRouterBgpAdvertisedIPRanges(c, &d, &n) {
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

func canonicalizeNewRouterBgpAdvertisedIPRangesSlice(c *Client, des, nw []RouterBgpAdvertisedIPRanges) []RouterBgpAdvertisedIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []RouterBgpAdvertisedIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterBgpAdvertisedIPRanges(c, &d, &n))
	}

	return items
}

type routerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         routerApiOperation
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
func diffRouter(c *Client, desired, actual *Router, opts ...dcl.ApplyOption) ([]routerDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []routerDiff
	if compareRouterNatsSlice(c, desired.Nats, actual.Nats) {
		c.Config.Logger.Infof("Detected diff in Nats.\nDESIRED: %v\nACTUAL: %v", desired.Nats, actual.Nats)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "Nats",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "Name",
		})

	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.NameToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, routerDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if compareRouterInterfacesSlice(c, desired.Interfaces, actual.Interfaces) {
		c.Config.Logger.Infof("Detected diff in Interfaces.\nDESIRED: %v\nACTUAL: %v", desired.Interfaces, actual.Interfaces)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "Interfaces",
		})

	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "Description",
		})

	}
	if compareRouterBgpPeersSlice(c, desired.BgpPeers, actual.BgpPeers) {
		c.Config.Logger.Infof("Detected diff in BgpPeers.\nDESIRED: %v\nACTUAL: %v", desired.BgpPeers, actual.BgpPeers)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "BgpPeers",
		})

	}
	if compareRouterBgp(c, desired.Bgp, actual.Bgp) {
		c.Config.Logger.Infof("Detected diff in Bgp.\nDESIRED: %v\nACTUAL: %v", desired.Bgp, actual.Bgp)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "Bgp",
		})

	}
	if !dcl.IsZeroValue(desired.Region) && !dcl.NameToSelfLink(desired.Region, actual.Region) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)

		diffs = append(diffs, routerDiff{
			UpdateOp:  &updateRouterUpdateOperation{},
			FieldName: "Region",
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
	var deduped []routerDiff
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
func compareRouterNats(c *Client, desired, actual *RouterNats) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.LogConfig == nil && desired.LogConfig != nil && !dcl.IsEmptyValueIndirect(desired.LogConfig) {
		c.Config.Logger.Infof("desired LogConfig %s - but actually nil", dcl.SprintResource(desired.LogConfig))
		return true
	}
	if compareRouterNatsLogConfig(c, desired.LogConfig, actual.LogConfig) && !dcl.IsZeroValue(desired.LogConfig) {
		c.Config.Logger.Infof("Diff in LogConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LogConfig), dcl.SprintResource(actual.LogConfig))
		return true
	}
	if actual.SourceSubnetworkIPRangesToNat == nil && desired.SourceSubnetworkIPRangesToNat != nil && !dcl.IsEmptyValueIndirect(desired.SourceSubnetworkIPRangesToNat) {
		c.Config.Logger.Infof("desired SourceSubnetworkIPRangesToNat %s - but actually nil", dcl.SprintResource(desired.SourceSubnetworkIPRangesToNat))
		return true
	}
	if !reflect.DeepEqual(desired.SourceSubnetworkIPRangesToNat, actual.SourceSubnetworkIPRangesToNat) && !dcl.IsZeroValue(desired.SourceSubnetworkIPRangesToNat) {
		c.Config.Logger.Infof("Diff in SourceSubnetworkIPRangesToNat. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceSubnetworkIPRangesToNat), dcl.SprintResource(actual.SourceSubnetworkIPRangesToNat))
		return true
	}
	if actual.NatIps == nil && desired.NatIps != nil && !dcl.IsEmptyValueIndirect(desired.NatIps) {
		c.Config.Logger.Infof("desired NatIps %s - but actually nil", dcl.SprintResource(desired.NatIps))
		return true
	}
	if !dcl.StringSliceEquals(desired.NatIps, actual.NatIps) && !dcl.IsZeroValue(desired.NatIps) {
		c.Config.Logger.Infof("Diff in NatIps. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NatIps), dcl.SprintResource(actual.NatIps))
		return true
	}
	if actual.DrainNatIps == nil && desired.DrainNatIps != nil && !dcl.IsEmptyValueIndirect(desired.DrainNatIps) {
		c.Config.Logger.Infof("desired DrainNatIps %s - but actually nil", dcl.SprintResource(desired.DrainNatIps))
		return true
	}
	if !dcl.StringSliceEquals(desired.DrainNatIps, actual.DrainNatIps) && !dcl.IsZeroValue(desired.DrainNatIps) {
		c.Config.Logger.Infof("Diff in DrainNatIps. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DrainNatIps), dcl.SprintResource(actual.DrainNatIps))
		return true
	}
	if actual.NatIPAllocateOption == nil && desired.NatIPAllocateOption != nil && !dcl.IsEmptyValueIndirect(desired.NatIPAllocateOption) {
		c.Config.Logger.Infof("desired NatIPAllocateOption %s - but actually nil", dcl.SprintResource(desired.NatIPAllocateOption))
		return true
	}
	if compareRouterNatsNatIPAllocateOptionEnumSlice(c, desired.NatIPAllocateOption, actual.NatIPAllocateOption) && !dcl.IsZeroValue(desired.NatIPAllocateOption) {
		c.Config.Logger.Infof("Diff in NatIPAllocateOption. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NatIPAllocateOption), dcl.SprintResource(actual.NatIPAllocateOption))
		return true
	}
	if actual.MinPortsPerVm == nil && desired.MinPortsPerVm != nil && !dcl.IsEmptyValueIndirect(desired.MinPortsPerVm) {
		c.Config.Logger.Infof("desired MinPortsPerVm %s - but actually nil", dcl.SprintResource(desired.MinPortsPerVm))
		return true
	}
	if !reflect.DeepEqual(desired.MinPortsPerVm, actual.MinPortsPerVm) && !dcl.IsZeroValue(desired.MinPortsPerVm) {
		c.Config.Logger.Infof("Diff in MinPortsPerVm. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinPortsPerVm), dcl.SprintResource(actual.MinPortsPerVm))
		return true
	}
	if actual.UdpIdleTimeoutSec == nil && desired.UdpIdleTimeoutSec != nil && !dcl.IsEmptyValueIndirect(desired.UdpIdleTimeoutSec) {
		c.Config.Logger.Infof("desired UdpIdleTimeoutSec %s - but actually nil", dcl.SprintResource(desired.UdpIdleTimeoutSec))
		return true
	}
	if !reflect.DeepEqual(desired.UdpIdleTimeoutSec, actual.UdpIdleTimeoutSec) && !dcl.IsZeroValue(desired.UdpIdleTimeoutSec) {
		c.Config.Logger.Infof("Diff in UdpIdleTimeoutSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UdpIdleTimeoutSec), dcl.SprintResource(actual.UdpIdleTimeoutSec))
		return true
	}
	if actual.IcmpIdleTimeoutSec == nil && desired.IcmpIdleTimeoutSec != nil && !dcl.IsEmptyValueIndirect(desired.IcmpIdleTimeoutSec) {
		c.Config.Logger.Infof("desired IcmpIdleTimeoutSec %s - but actually nil", dcl.SprintResource(desired.IcmpIdleTimeoutSec))
		return true
	}
	if !reflect.DeepEqual(desired.IcmpIdleTimeoutSec, actual.IcmpIdleTimeoutSec) && !dcl.IsZeroValue(desired.IcmpIdleTimeoutSec) {
		c.Config.Logger.Infof("Diff in IcmpIdleTimeoutSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IcmpIdleTimeoutSec), dcl.SprintResource(actual.IcmpIdleTimeoutSec))
		return true
	}
	if actual.TcpEstablishedIdleTimeoutSec == nil && desired.TcpEstablishedIdleTimeoutSec != nil && !dcl.IsEmptyValueIndirect(desired.TcpEstablishedIdleTimeoutSec) {
		c.Config.Logger.Infof("desired TcpEstablishedIdleTimeoutSec %s - but actually nil", dcl.SprintResource(desired.TcpEstablishedIdleTimeoutSec))
		return true
	}
	if !reflect.DeepEqual(desired.TcpEstablishedIdleTimeoutSec, actual.TcpEstablishedIdleTimeoutSec) && !dcl.IsZeroValue(desired.TcpEstablishedIdleTimeoutSec) {
		c.Config.Logger.Infof("Diff in TcpEstablishedIdleTimeoutSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TcpEstablishedIdleTimeoutSec), dcl.SprintResource(actual.TcpEstablishedIdleTimeoutSec))
		return true
	}
	if actual.TcpTransitoryIdleTimeoutSec == nil && desired.TcpTransitoryIdleTimeoutSec != nil && !dcl.IsEmptyValueIndirect(desired.TcpTransitoryIdleTimeoutSec) {
		c.Config.Logger.Infof("desired TcpTransitoryIdleTimeoutSec %s - but actually nil", dcl.SprintResource(desired.TcpTransitoryIdleTimeoutSec))
		return true
	}
	if !reflect.DeepEqual(desired.TcpTransitoryIdleTimeoutSec, actual.TcpTransitoryIdleTimeoutSec) && !dcl.IsZeroValue(desired.TcpTransitoryIdleTimeoutSec) {
		c.Config.Logger.Infof("Diff in TcpTransitoryIdleTimeoutSec. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TcpTransitoryIdleTimeoutSec), dcl.SprintResource(actual.TcpTransitoryIdleTimeoutSec))
		return true
	}
	if actual.Subnetworks == nil && desired.Subnetworks != nil && !dcl.IsEmptyValueIndirect(desired.Subnetworks) {
		c.Config.Logger.Infof("desired Subnetworks %s - but actually nil", dcl.SprintResource(desired.Subnetworks))
		return true
	}
	if compareRouterNatsSubnetworksSlice(c, desired.Subnetworks, actual.Subnetworks) && !dcl.IsZeroValue(desired.Subnetworks) {
		c.Config.Logger.Infof("Diff in Subnetworks. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Subnetworks), dcl.SprintResource(actual.Subnetworks))
		return true
	}
	return false
}

func compareRouterNatsSlice(c *Client, desired, actual []RouterNats) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNats, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterNats(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterNats, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterNatsMap(c *Client, desired, actual map[string]RouterNats) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNats, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterNats, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterNats(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterNats, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterNatsLogConfig(c *Client, desired, actual *RouterNatsLogConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Enable == nil && desired.Enable != nil && !dcl.IsEmptyValueIndirect(desired.Enable) {
		c.Config.Logger.Infof("desired Enable %s - but actually nil", dcl.SprintResource(desired.Enable))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enable, actual.Enable) && !dcl.IsZeroValue(desired.Enable) {
		c.Config.Logger.Infof("Diff in Enable. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enable), dcl.SprintResource(actual.Enable))
		return true
	}
	if actual.Filter == nil && desired.Filter != nil && !dcl.IsEmptyValueIndirect(desired.Filter) {
		c.Config.Logger.Infof("desired Filter %s - but actually nil", dcl.SprintResource(desired.Filter))
		return true
	}
	if !reflect.DeepEqual(desired.Filter, actual.Filter) && !dcl.IsZeroValue(desired.Filter) {
		c.Config.Logger.Infof("Diff in Filter. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Filter), dcl.SprintResource(actual.Filter))
		return true
	}
	return false
}

func compareRouterNatsLogConfigSlice(c *Client, desired, actual []RouterNatsLogConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsLogConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterNatsLogConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterNatsLogConfig, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterNatsLogConfigMap(c *Client, desired, actual map[string]RouterNatsLogConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsLogConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterNatsLogConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterNatsLogConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterNatsLogConfig, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterNatsSubnetworks(c *Client, desired, actual *RouterNatsSubnetworks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.SourceIPRangesToNat == nil && desired.SourceIPRangesToNat != nil && !dcl.IsEmptyValueIndirect(desired.SourceIPRangesToNat) {
		c.Config.Logger.Infof("desired SourceIPRangesToNat %s - but actually nil", dcl.SprintResource(desired.SourceIPRangesToNat))
		return true
	}
	if !dcl.StringCanonicalize(desired.SourceIPRangesToNat, actual.SourceIPRangesToNat) && !dcl.IsZeroValue(desired.SourceIPRangesToNat) {
		c.Config.Logger.Infof("Diff in SourceIPRangesToNat. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SourceIPRangesToNat), dcl.SprintResource(actual.SourceIPRangesToNat))
		return true
	}
	if actual.SecondaryIPRangeNames == nil && desired.SecondaryIPRangeNames != nil && !dcl.IsEmptyValueIndirect(desired.SecondaryIPRangeNames) {
		c.Config.Logger.Infof("desired SecondaryIPRangeNames %s - but actually nil", dcl.SprintResource(desired.SecondaryIPRangeNames))
		return true
	}
	if !dcl.StringCanonicalize(desired.SecondaryIPRangeNames, actual.SecondaryIPRangeNames) && !dcl.IsZeroValue(desired.SecondaryIPRangeNames) {
		c.Config.Logger.Infof("Diff in SecondaryIPRangeNames. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SecondaryIPRangeNames), dcl.SprintResource(actual.SecondaryIPRangeNames))
		return true
	}
	return false
}

func compareRouterNatsSubnetworksSlice(c *Client, desired, actual []RouterNatsSubnetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsSubnetworks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterNatsSubnetworks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterNatsSubnetworks, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterNatsSubnetworksMap(c *Client, desired, actual map[string]RouterNatsSubnetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsSubnetworks, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterNatsSubnetworks, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterNatsSubnetworks(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterNatsSubnetworks, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterInterfaces(c *Client, desired, actual *RouterInterfaces) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.LinkedVpnTunnel == nil && desired.LinkedVpnTunnel != nil && !dcl.IsEmptyValueIndirect(desired.LinkedVpnTunnel) {
		c.Config.Logger.Infof("desired LinkedVpnTunnel %s - but actually nil", dcl.SprintResource(desired.LinkedVpnTunnel))
		return true
	}
	if !dcl.NameToSelfLink(desired.LinkedVpnTunnel, actual.LinkedVpnTunnel) && !dcl.IsZeroValue(desired.LinkedVpnTunnel) {
		c.Config.Logger.Infof("Diff in LinkedVpnTunnel. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LinkedVpnTunnel), dcl.SprintResource(actual.LinkedVpnTunnel))
		return true
	}
	if actual.IPRange == nil && desired.IPRange != nil && !dcl.IsEmptyValueIndirect(desired.IPRange) {
		c.Config.Logger.Infof("desired IPRange %s - but actually nil", dcl.SprintResource(desired.IPRange))
		return true
	}
	if !dcl.StringCanonicalize(desired.IPRange, actual.IPRange) && !dcl.IsZeroValue(desired.IPRange) {
		c.Config.Logger.Infof("Diff in IPRange. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPRange), dcl.SprintResource(actual.IPRange))
		return true
	}
	return false
}

func compareRouterInterfacesSlice(c *Client, desired, actual []RouterInterfaces) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterInterfaces, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterInterfaces(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterInterfaces, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterInterfacesMap(c *Client, desired, actual map[string]RouterInterfaces) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterInterfaces, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterInterfaces, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterInterfaces(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterInterfaces, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterBgpPeers(c *Client, desired, actual *RouterBgpPeers) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Name == nil && desired.Name != nil && !dcl.IsEmptyValueIndirect(desired.Name) {
		c.Config.Logger.Infof("desired Name %s - but actually nil", dcl.SprintResource(desired.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if actual.InterfaceName == nil && desired.InterfaceName != nil && !dcl.IsEmptyValueIndirect(desired.InterfaceName) {
		c.Config.Logger.Infof("desired InterfaceName %s - but actually nil", dcl.SprintResource(desired.InterfaceName))
		return true
	}
	if !dcl.StringCanonicalize(desired.InterfaceName, actual.InterfaceName) && !dcl.IsZeroValue(desired.InterfaceName) {
		c.Config.Logger.Infof("Diff in InterfaceName. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InterfaceName), dcl.SprintResource(actual.InterfaceName))
		return true
	}
	if actual.IPAddress == nil && desired.IPAddress != nil && !dcl.IsEmptyValueIndirect(desired.IPAddress) {
		c.Config.Logger.Infof("desired IPAddress %s - but actually nil", dcl.SprintResource(desired.IPAddress))
		return true
	}
	if !dcl.StringCanonicalize(desired.IPAddress, actual.IPAddress) && !dcl.IsZeroValue(desired.IPAddress) {
		c.Config.Logger.Infof("Diff in IPAddress. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPAddress), dcl.SprintResource(actual.IPAddress))
		return true
	}
	if actual.PeerAsn == nil && desired.PeerAsn != nil && !dcl.IsEmptyValueIndirect(desired.PeerAsn) {
		c.Config.Logger.Infof("desired PeerAsn %s - but actually nil", dcl.SprintResource(desired.PeerAsn))
		return true
	}
	if !reflect.DeepEqual(desired.PeerAsn, actual.PeerAsn) && !dcl.IsZeroValue(desired.PeerAsn) {
		c.Config.Logger.Infof("Diff in PeerAsn. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PeerAsn), dcl.SprintResource(actual.PeerAsn))
		return true
	}
	if actual.AdvertisedRoutePriority == nil && desired.AdvertisedRoutePriority != nil && !dcl.IsEmptyValueIndirect(desired.AdvertisedRoutePriority) {
		c.Config.Logger.Infof("desired AdvertisedRoutePriority %s - but actually nil", dcl.SprintResource(desired.AdvertisedRoutePriority))
		return true
	}
	if !reflect.DeepEqual(desired.AdvertisedRoutePriority, actual.AdvertisedRoutePriority) && !dcl.IsZeroValue(desired.AdvertisedRoutePriority) {
		c.Config.Logger.Infof("Diff in AdvertisedRoutePriority. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertisedRoutePriority), dcl.SprintResource(actual.AdvertisedRoutePriority))
		return true
	}
	if actual.AdvertiseMode == nil && desired.AdvertiseMode != nil && !dcl.IsEmptyValueIndirect(desired.AdvertiseMode) {
		c.Config.Logger.Infof("desired AdvertiseMode %s - but actually nil", dcl.SprintResource(desired.AdvertiseMode))
		return true
	}
	if !dcl.StringCanonicalize(desired.AdvertiseMode, actual.AdvertiseMode) && !dcl.IsZeroValue(desired.AdvertiseMode) {
		c.Config.Logger.Infof("Diff in AdvertiseMode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertiseMode), dcl.SprintResource(actual.AdvertiseMode))
		return true
	}
	if actual.AdvertisedGroups == nil && desired.AdvertisedGroups != nil && !dcl.IsEmptyValueIndirect(desired.AdvertisedGroups) {
		c.Config.Logger.Infof("desired AdvertisedGroups %s - but actually nil", dcl.SprintResource(desired.AdvertisedGroups))
		return true
	}
	if compareRouterBgpPeersAdvertisedGroupsEnumSlice(c, desired.AdvertisedGroups, actual.AdvertisedGroups) && !dcl.IsZeroValue(desired.AdvertisedGroups) {
		c.Config.Logger.Infof("Diff in AdvertisedGroups. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertisedGroups), dcl.SprintResource(actual.AdvertisedGroups))
		return true
	}
	if actual.AdvertisedIPRanges == nil && desired.AdvertisedIPRanges != nil && !dcl.IsEmptyValueIndirect(desired.AdvertisedIPRanges) {
		c.Config.Logger.Infof("desired AdvertisedIPRanges %s - but actually nil", dcl.SprintResource(desired.AdvertisedIPRanges))
		return true
	}
	if compareRouterBgpPeersAdvertisedIPRangesSlice(c, desired.AdvertisedIPRanges, actual.AdvertisedIPRanges) && !dcl.IsZeroValue(desired.AdvertisedIPRanges) {
		c.Config.Logger.Infof("Diff in AdvertisedIPRanges. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertisedIPRanges), dcl.SprintResource(actual.AdvertisedIPRanges))
		return true
	}
	return false
}

func compareRouterBgpPeersSlice(c *Client, desired, actual []RouterBgpPeers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpPeers, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterBgpPeers(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterBgpPeers, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterBgpPeersMap(c *Client, desired, actual map[string]RouterBgpPeers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpPeers, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterBgpPeers, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterBgpPeers(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterBgpPeers, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterBgpPeersAdvertisedIPRanges(c *Client, desired, actual *RouterBgpPeersAdvertisedIPRanges) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Range == nil && desired.Range != nil && !dcl.IsEmptyValueIndirect(desired.Range) {
		c.Config.Logger.Infof("desired Range %s - but actually nil", dcl.SprintResource(desired.Range))
		return true
	}
	if !dcl.StringCanonicalize(desired.Range, actual.Range) && !dcl.IsZeroValue(desired.Range) {
		c.Config.Logger.Infof("Diff in Range. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Range), dcl.SprintResource(actual.Range))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	return false
}

func compareRouterBgpPeersAdvertisedIPRangesSlice(c *Client, desired, actual []RouterBgpPeersAdvertisedIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpPeersAdvertisedIPRanges, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterBgpPeersAdvertisedIPRanges(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterBgpPeersAdvertisedIPRanges, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterBgpPeersAdvertisedIPRangesMap(c *Client, desired, actual map[string]RouterBgpPeersAdvertisedIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpPeersAdvertisedIPRanges, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterBgpPeersAdvertisedIPRanges, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterBgpPeersAdvertisedIPRanges(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterBgpPeersAdvertisedIPRanges, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterBgp(c *Client, desired, actual *RouterBgp) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Asn == nil && desired.Asn != nil && !dcl.IsEmptyValueIndirect(desired.Asn) {
		c.Config.Logger.Infof("desired Asn %s - but actually nil", dcl.SprintResource(desired.Asn))
		return true
	}
	if !reflect.DeepEqual(desired.Asn, actual.Asn) && !dcl.IsZeroValue(desired.Asn) {
		c.Config.Logger.Infof("Diff in Asn. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Asn), dcl.SprintResource(actual.Asn))
		return true
	}
	if actual.AdvertiseMode == nil && desired.AdvertiseMode != nil && !dcl.IsEmptyValueIndirect(desired.AdvertiseMode) {
		c.Config.Logger.Infof("desired AdvertiseMode %s - but actually nil", dcl.SprintResource(desired.AdvertiseMode))
		return true
	}
	if !reflect.DeepEqual(desired.AdvertiseMode, actual.AdvertiseMode) && !dcl.IsZeroValue(desired.AdvertiseMode) {
		c.Config.Logger.Infof("Diff in AdvertiseMode. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertiseMode), dcl.SprintResource(actual.AdvertiseMode))
		return true
	}
	if actual.AdvertisedGroups == nil && desired.AdvertisedGroups != nil && !dcl.IsEmptyValueIndirect(desired.AdvertisedGroups) {
		c.Config.Logger.Infof("desired AdvertisedGroups %s - but actually nil", dcl.SprintResource(desired.AdvertisedGroups))
		return true
	}
	if !dcl.StringSliceEquals(desired.AdvertisedGroups, actual.AdvertisedGroups) && !dcl.IsZeroValue(desired.AdvertisedGroups) {
		c.Config.Logger.Infof("Diff in AdvertisedGroups. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertisedGroups), dcl.SprintResource(actual.AdvertisedGroups))
		return true
	}
	if actual.AdvertisedIPRanges == nil && desired.AdvertisedIPRanges != nil && !dcl.IsEmptyValueIndirect(desired.AdvertisedIPRanges) {
		c.Config.Logger.Infof("desired AdvertisedIPRanges %s - but actually nil", dcl.SprintResource(desired.AdvertisedIPRanges))
		return true
	}
	if compareRouterBgpAdvertisedIPRangesSlice(c, desired.AdvertisedIPRanges, actual.AdvertisedIPRanges) && !dcl.IsZeroValue(desired.AdvertisedIPRanges) {
		c.Config.Logger.Infof("Diff in AdvertisedIPRanges. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AdvertisedIPRanges), dcl.SprintResource(actual.AdvertisedIPRanges))
		return true
	}
	return false
}

func compareRouterBgpSlice(c *Client, desired, actual []RouterBgp) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgp, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterBgp(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterBgp, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterBgpMap(c *Client, desired, actual map[string]RouterBgp) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgp, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterBgp, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterBgp(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterBgp, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterBgpAdvertisedIPRanges(c *Client, desired, actual *RouterBgpAdvertisedIPRanges) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Range == nil && desired.Range != nil && !dcl.IsEmptyValueIndirect(desired.Range) {
		c.Config.Logger.Infof("desired Range %s - but actually nil", dcl.SprintResource(desired.Range))
		return true
	}
	if !dcl.StringCanonicalize(desired.Range, actual.Range) && !dcl.IsZeroValue(desired.Range) {
		c.Config.Logger.Infof("Diff in Range. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Range), dcl.SprintResource(actual.Range))
		return true
	}
	if actual.Description == nil && desired.Description != nil && !dcl.IsEmptyValueIndirect(desired.Description) {
		c.Config.Logger.Infof("desired Description %s - but actually nil", dcl.SprintResource(desired.Description))
		return true
	}
	if !dcl.StringCanonicalize(desired.Description, actual.Description) && !dcl.IsZeroValue(desired.Description) {
		c.Config.Logger.Infof("Diff in Description. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Description), dcl.SprintResource(actual.Description))
		return true
	}
	return false
}

func compareRouterBgpAdvertisedIPRangesSlice(c *Client, desired, actual []RouterBgpAdvertisedIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpAdvertisedIPRanges, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterBgpAdvertisedIPRanges(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterBgpAdvertisedIPRanges, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterBgpAdvertisedIPRangesMap(c *Client, desired, actual map[string]RouterBgpAdvertisedIPRanges) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpAdvertisedIPRanges, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouterBgpAdvertisedIPRanges, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouterBgpAdvertisedIPRanges(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouterBgpAdvertisedIPRanges, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouterNatsLogConfigFilterEnumSlice(c *Client, desired, actual []RouterNatsLogConfigFilterEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsLogConfigFilterEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterNatsLogConfigFilterEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterNatsLogConfigFilterEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterNatsLogConfigFilterEnum(c *Client, desired, actual *RouterNatsLogConfigFilterEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareRouterNatsSourceSubnetworkIPRangesToNatEnumSlice(c *Client, desired, actual []RouterNatsSourceSubnetworkIPRangesToNatEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsSourceSubnetworkIPRangesToNatEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterNatsSourceSubnetworkIPRangesToNatEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterNatsSourceSubnetworkIPRangesToNatEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterNatsSourceSubnetworkIPRangesToNatEnum(c *Client, desired, actual *RouterNatsSourceSubnetworkIPRangesToNatEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareRouterNatsNatIPAllocateOptionEnumSlice(c *Client, desired, actual []RouterNatsNatIPAllocateOptionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterNatsNatIPAllocateOptionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterNatsNatIPAllocateOptionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterNatsNatIPAllocateOptionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterNatsNatIPAllocateOptionEnum(c *Client, desired, actual *RouterNatsNatIPAllocateOptionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareRouterInterfacesManagementTypeEnumSlice(c *Client, desired, actual []RouterInterfacesManagementTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterInterfacesManagementTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterInterfacesManagementTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterInterfacesManagementTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterInterfacesManagementTypeEnum(c *Client, desired, actual *RouterInterfacesManagementTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareRouterBgpPeersAdvertisedGroupsEnumSlice(c *Client, desired, actual []RouterBgpPeersAdvertisedGroupsEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpPeersAdvertisedGroupsEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterBgpPeersAdvertisedGroupsEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterBgpPeersAdvertisedGroupsEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterBgpPeersAdvertisedGroupsEnum(c *Client, desired, actual *RouterBgpPeersAdvertisedGroupsEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareRouterBgpAdvertiseModeEnumSlice(c *Client, desired, actual []RouterBgpAdvertiseModeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouterBgpAdvertiseModeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouterBgpAdvertiseModeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouterBgpAdvertiseModeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouterBgpAdvertiseModeEnum(c *Client, desired, actual *RouterBgpAdvertiseModeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Router) urlNormalized() *Router {
	normalized := deepcopy.Copy(*r).(Router)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *Router) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Router) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *Router) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Router) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Router resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Router) marshal(c *Client) ([]byte, error) {
	m, err := expandRouter(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Router: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalRouter decodes JSON responses into the Router resource schema.
func unmarshalRouter(b []byte, c *Client) (*Router, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRouter(m, c)
}

func unmarshalMapRouter(m map[string]interface{}, c *Client) (*Router, error) {

	return flattenRouter(c, m), nil
}

// expandRouter expands Router into a JSON request object.
func expandRouter(c *Client, f *Router) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v, err := expandRouterNatsSlice(c, f.Nats); err != nil {
		return nil, fmt.Errorf("error expanding Nats into nats: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["nats"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v, err := expandRouterInterfacesSlice(c, f.Interfaces); err != nil {
		return nil, fmt.Errorf("error expanding Interfaces into interfaces: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["interfaces"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandRouterBgpPeersSlice(c, f.BgpPeers); err != nil {
		return nil, fmt.Errorf("error expanding BgpPeers into bgpPeers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bgpPeers"] = v
	}
	if v, err := expandRouterBgp(c, f.Bgp); err != nil {
		return nil, fmt.Errorf("error expanding Bgp into bgp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bgp"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}

	return m, nil
}

// flattenRouter flattens Router from a JSON request object into the
// Router type.
func flattenRouter(c *Client, i interface{}) *Router {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Router{}
	r.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	r.Nats = flattenRouterNatsSlice(c, m["nats"])
	r.Name = dcl.FlattenString(m["name"])
	r.Network = dcl.FlattenString(m["network"])
	r.Interfaces = flattenRouterInterfacesSlice(c, m["interfaces"])
	r.Description = dcl.FlattenString(m["description"])
	r.BgpPeers = flattenRouterBgpPeersSlice(c, m["bgpPeers"])
	r.Bgp = flattenRouterBgp(c, m["bgp"])
	r.Region = dcl.FlattenString(m["region"])
	r.Project = dcl.FlattenString(m["project"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])

	return r
}

// expandRouterNatsMap expands the contents of RouterNats into a JSON
// request object.
func expandRouterNatsMap(c *Client, f map[string]RouterNats) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterNats(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterNatsSlice expands the contents of RouterNats into a JSON
// request object.
func expandRouterNatsSlice(c *Client, f []RouterNats) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterNats(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterNatsMap flattens the contents of RouterNats from a JSON
// response object.
func flattenRouterNatsMap(c *Client, i interface{}) map[string]RouterNats {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterNats{}
	}

	if len(a) == 0 {
		return map[string]RouterNats{}
	}

	items := make(map[string]RouterNats)
	for k, item := range a {
		items[k] = *flattenRouterNats(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterNatsSlice flattens the contents of RouterNats from a JSON
// response object.
func flattenRouterNatsSlice(c *Client, i interface{}) []RouterNats {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterNats{}
	}

	if len(a) == 0 {
		return []RouterNats{}
	}

	items := make([]RouterNats, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterNats(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterNats expands an instance of RouterNats into a JSON
// request object.
func expandRouterNats(c *Client, f *RouterNats) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandRouterNatsLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["logConfig"] = v
	}
	if v := f.SourceSubnetworkIPRangesToNat; !dcl.IsEmptyValueIndirect(v) {
		m["sourceSubnetworkIPRangesToNat"] = v
	}
	if v := f.NatIps; !dcl.IsEmptyValueIndirect(v) {
		m["natIps"] = v
	}
	if v := f.DrainNatIps; !dcl.IsEmptyValueIndirect(v) {
		m["drainNatIps"] = v
	}
	if v := f.NatIPAllocateOption; !dcl.IsEmptyValueIndirect(v) {
		m["natIPAllocateOption"] = v
	}
	if v := f.MinPortsPerVm; !dcl.IsEmptyValueIndirect(v) {
		m["minPortsPerVm"] = v
	}
	if v := f.UdpIdleTimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["udpIdleTimeoutSec"] = v
	}
	if v := f.IcmpIdleTimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["icmpIdleTimeoutSec"] = v
	}
	if v := f.TcpEstablishedIdleTimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["tcpEstablishedIdleTimeoutSec"] = v
	}
	if v := f.TcpTransitoryIdleTimeoutSec; !dcl.IsEmptyValueIndirect(v) {
		m["tcpTransitoryIdleTimeoutSec"] = v
	}
	if v, err := expandRouterNatsSubnetworksSlice(c, f.Subnetworks); err != nil {
		return nil, fmt.Errorf("error expanding Subnetworks into subnetworks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subnetworks"] = v
	}

	return m, nil
}

// flattenRouterNats flattens an instance of RouterNats from a JSON
// response object.
func flattenRouterNats(c *Client, i interface{}) *RouterNats {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterNats{}
	r.Name = dcl.FlattenString(m["name"])
	r.LogConfig = flattenRouterNatsLogConfig(c, m["logConfig"])
	r.SourceSubnetworkIPRangesToNat = flattenRouterNatsSourceSubnetworkIPRangesToNatEnum(m["sourceSubnetworkIPRangesToNat"])
	r.NatIps = dcl.FlattenStringSlice(m["natIps"])
	r.DrainNatIps = dcl.FlattenStringSlice(m["drainNatIps"])
	r.NatIPAllocateOption = flattenRouterNatsNatIPAllocateOptionEnumSlice(c, m["natIPAllocateOption"])
	r.MinPortsPerVm = dcl.FlattenInteger(m["minPortsPerVm"])
	r.UdpIdleTimeoutSec = dcl.FlattenInteger(m["udpIdleTimeoutSec"])
	r.IcmpIdleTimeoutSec = dcl.FlattenInteger(m["icmpIdleTimeoutSec"])
	r.TcpEstablishedIdleTimeoutSec = dcl.FlattenInteger(m["tcpEstablishedIdleTimeoutSec"])
	r.TcpTransitoryIdleTimeoutSec = dcl.FlattenInteger(m["tcpTransitoryIdleTimeoutSec"])
	r.Subnetworks = flattenRouterNatsSubnetworksSlice(c, m["subnetworks"])

	return r
}

// expandRouterNatsLogConfigMap expands the contents of RouterNatsLogConfig into a JSON
// request object.
func expandRouterNatsLogConfigMap(c *Client, f map[string]RouterNatsLogConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterNatsLogConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterNatsLogConfigSlice expands the contents of RouterNatsLogConfig into a JSON
// request object.
func expandRouterNatsLogConfigSlice(c *Client, f []RouterNatsLogConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterNatsLogConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterNatsLogConfigMap flattens the contents of RouterNatsLogConfig from a JSON
// response object.
func flattenRouterNatsLogConfigMap(c *Client, i interface{}) map[string]RouterNatsLogConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterNatsLogConfig{}
	}

	if len(a) == 0 {
		return map[string]RouterNatsLogConfig{}
	}

	items := make(map[string]RouterNatsLogConfig)
	for k, item := range a {
		items[k] = *flattenRouterNatsLogConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterNatsLogConfigSlice flattens the contents of RouterNatsLogConfig from a JSON
// response object.
func flattenRouterNatsLogConfigSlice(c *Client, i interface{}) []RouterNatsLogConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterNatsLogConfig{}
	}

	if len(a) == 0 {
		return []RouterNatsLogConfig{}
	}

	items := make([]RouterNatsLogConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterNatsLogConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterNatsLogConfig expands an instance of RouterNatsLogConfig into a JSON
// request object.
func expandRouterNatsLogConfig(c *Client, f *RouterNatsLogConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enable; !dcl.IsEmptyValueIndirect(v) {
		m["enable"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		m["filter"] = v
	}

	return m, nil
}

// flattenRouterNatsLogConfig flattens an instance of RouterNatsLogConfig from a JSON
// response object.
func flattenRouterNatsLogConfig(c *Client, i interface{}) *RouterNatsLogConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterNatsLogConfig{}
	r.Enable = dcl.FlattenBool(m["enable"])
	r.Filter = flattenRouterNatsLogConfigFilterEnum(m["filter"])

	return r
}

// expandRouterNatsSubnetworksMap expands the contents of RouterNatsSubnetworks into a JSON
// request object.
func expandRouterNatsSubnetworksMap(c *Client, f map[string]RouterNatsSubnetworks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterNatsSubnetworks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterNatsSubnetworksSlice expands the contents of RouterNatsSubnetworks into a JSON
// request object.
func expandRouterNatsSubnetworksSlice(c *Client, f []RouterNatsSubnetworks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterNatsSubnetworks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterNatsSubnetworksMap flattens the contents of RouterNatsSubnetworks from a JSON
// response object.
func flattenRouterNatsSubnetworksMap(c *Client, i interface{}) map[string]RouterNatsSubnetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterNatsSubnetworks{}
	}

	if len(a) == 0 {
		return map[string]RouterNatsSubnetworks{}
	}

	items := make(map[string]RouterNatsSubnetworks)
	for k, item := range a {
		items[k] = *flattenRouterNatsSubnetworks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterNatsSubnetworksSlice flattens the contents of RouterNatsSubnetworks from a JSON
// response object.
func flattenRouterNatsSubnetworksSlice(c *Client, i interface{}) []RouterNatsSubnetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterNatsSubnetworks{}
	}

	if len(a) == 0 {
		return []RouterNatsSubnetworks{}
	}

	items := make([]RouterNatsSubnetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterNatsSubnetworks(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterNatsSubnetworks expands an instance of RouterNatsSubnetworks into a JSON
// request object.
func expandRouterNatsSubnetworks(c *Client, f *RouterNatsSubnetworks) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.SourceIPRangesToNat; !dcl.IsEmptyValueIndirect(v) {
		m["sourceIPRangesToNat"] = v
	}
	if v := f.SecondaryIPRangeNames; !dcl.IsEmptyValueIndirect(v) {
		m["secondaryIPRangeNames"] = v
	}

	return m, nil
}

// flattenRouterNatsSubnetworks flattens an instance of RouterNatsSubnetworks from a JSON
// response object.
func flattenRouterNatsSubnetworks(c *Client, i interface{}) *RouterNatsSubnetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterNatsSubnetworks{}
	r.Name = dcl.FlattenString(m["name"])
	r.SourceIPRangesToNat = dcl.FlattenString(m["sourceIPRangesToNat"])
	r.SecondaryIPRangeNames = dcl.FlattenString(m["secondaryIPRangeNames"])

	return r
}

// expandRouterInterfacesMap expands the contents of RouterInterfaces into a JSON
// request object.
func expandRouterInterfacesMap(c *Client, f map[string]RouterInterfaces) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterInterfaces(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterInterfacesSlice expands the contents of RouterInterfaces into a JSON
// request object.
func expandRouterInterfacesSlice(c *Client, f []RouterInterfaces) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterInterfaces(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterInterfacesMap flattens the contents of RouterInterfaces from a JSON
// response object.
func flattenRouterInterfacesMap(c *Client, i interface{}) map[string]RouterInterfaces {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterInterfaces{}
	}

	if len(a) == 0 {
		return map[string]RouterInterfaces{}
	}

	items := make(map[string]RouterInterfaces)
	for k, item := range a {
		items[k] = *flattenRouterInterfaces(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterInterfacesSlice flattens the contents of RouterInterfaces from a JSON
// response object.
func flattenRouterInterfacesSlice(c *Client, i interface{}) []RouterInterfaces {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterInterfaces{}
	}

	if len(a) == 0 {
		return []RouterInterfaces{}
	}

	items := make([]RouterInterfaces, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterInterfaces(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterInterfaces expands an instance of RouterInterfaces into a JSON
// request object.
func expandRouterInterfaces(c *Client, f *RouterInterfaces) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.LinkedVpnTunnel; !dcl.IsEmptyValueIndirect(v) {
		m["linkedVpnTunnel"] = v
	}
	if v := f.IPRange; !dcl.IsEmptyValueIndirect(v) {
		m["ipRange"] = v
	}
	if v := f.ManagementType; !dcl.IsEmptyValueIndirect(v) {
		m["managementType"] = v
	}

	return m, nil
}

// flattenRouterInterfaces flattens an instance of RouterInterfaces from a JSON
// response object.
func flattenRouterInterfaces(c *Client, i interface{}) *RouterInterfaces {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterInterfaces{}
	r.Name = dcl.FlattenString(m["name"])
	r.LinkedVpnTunnel = dcl.FlattenString(m["linkedVpnTunnel"])
	r.IPRange = dcl.FlattenString(m["ipRange"])
	r.ManagementType = flattenRouterInterfacesManagementTypeEnum(m["managementType"])

	return r
}

// expandRouterBgpPeersMap expands the contents of RouterBgpPeers into a JSON
// request object.
func expandRouterBgpPeersMap(c *Client, f map[string]RouterBgpPeers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterBgpPeers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterBgpPeersSlice expands the contents of RouterBgpPeers into a JSON
// request object.
func expandRouterBgpPeersSlice(c *Client, f []RouterBgpPeers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterBgpPeers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterBgpPeersMap flattens the contents of RouterBgpPeers from a JSON
// response object.
func flattenRouterBgpPeersMap(c *Client, i interface{}) map[string]RouterBgpPeers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterBgpPeers{}
	}

	if len(a) == 0 {
		return map[string]RouterBgpPeers{}
	}

	items := make(map[string]RouterBgpPeers)
	for k, item := range a {
		items[k] = *flattenRouterBgpPeers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterBgpPeersSlice flattens the contents of RouterBgpPeers from a JSON
// response object.
func flattenRouterBgpPeersSlice(c *Client, i interface{}) []RouterBgpPeers {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterBgpPeers{}
	}

	if len(a) == 0 {
		return []RouterBgpPeers{}
	}

	items := make([]RouterBgpPeers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterBgpPeers(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterBgpPeers expands an instance of RouterBgpPeers into a JSON
// request object.
func expandRouterBgpPeers(c *Client, f *RouterBgpPeers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.InterfaceName; !dcl.IsEmptyValueIndirect(v) {
		m["interfaceName"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}
	if v := f.PeerIPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["peerIPAddress"] = v
	}
	if v := f.PeerAsn; !dcl.IsEmptyValueIndirect(v) {
		m["peerAsn"] = v
	}
	if v := f.AdvertisedRoutePriority; !dcl.IsEmptyValueIndirect(v) {
		m["advertisedRoutePriority"] = v
	}
	if v := f.AdvertiseMode; !dcl.IsEmptyValueIndirect(v) {
		m["advertiseMode"] = v
	}
	if v := f.ManagementType; !dcl.IsEmptyValueIndirect(v) {
		m["managementType"] = v
	}
	if v := f.AdvertisedGroups; !dcl.IsEmptyValueIndirect(v) {
		m["advertisedGroups"] = v
	}
	if v, err := expandRouterBgpPeersAdvertisedIPRangesSlice(c, f.AdvertisedIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding AdvertisedIPRanges into advertisedIpRanges: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["advertisedIpRanges"] = v
	}

	return m, nil
}

// flattenRouterBgpPeers flattens an instance of RouterBgpPeers from a JSON
// response object.
func flattenRouterBgpPeers(c *Client, i interface{}) *RouterBgpPeers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterBgpPeers{}
	r.Name = dcl.FlattenString(m["name"])
	r.InterfaceName = dcl.FlattenString(m["interfaceName"])
	r.IPAddress = dcl.FlattenString(m["ipAddress"])
	r.PeerIPAddress = dcl.FlattenSecretValue(m["peerIPAddress"])
	r.PeerAsn = dcl.FlattenInteger(m["peerAsn"])
	r.AdvertisedRoutePriority = dcl.FlattenInteger(m["advertisedRoutePriority"])
	r.AdvertiseMode = dcl.FlattenString(m["advertiseMode"])
	r.ManagementType = dcl.FlattenString(m["managementType"])
	r.AdvertisedGroups = flattenRouterBgpPeersAdvertisedGroupsEnumSlice(c, m["advertisedGroups"])
	r.AdvertisedIPRanges = flattenRouterBgpPeersAdvertisedIPRangesSlice(c, m["advertisedIpRanges"])

	return r
}

// expandRouterBgpPeersAdvertisedIPRangesMap expands the contents of RouterBgpPeersAdvertisedIPRanges into a JSON
// request object.
func expandRouterBgpPeersAdvertisedIPRangesMap(c *Client, f map[string]RouterBgpPeersAdvertisedIPRanges) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterBgpPeersAdvertisedIPRanges(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterBgpPeersAdvertisedIPRangesSlice expands the contents of RouterBgpPeersAdvertisedIPRanges into a JSON
// request object.
func expandRouterBgpPeersAdvertisedIPRangesSlice(c *Client, f []RouterBgpPeersAdvertisedIPRanges) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterBgpPeersAdvertisedIPRanges(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterBgpPeersAdvertisedIPRangesMap flattens the contents of RouterBgpPeersAdvertisedIPRanges from a JSON
// response object.
func flattenRouterBgpPeersAdvertisedIPRangesMap(c *Client, i interface{}) map[string]RouterBgpPeersAdvertisedIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterBgpPeersAdvertisedIPRanges{}
	}

	if len(a) == 0 {
		return map[string]RouterBgpPeersAdvertisedIPRanges{}
	}

	items := make(map[string]RouterBgpPeersAdvertisedIPRanges)
	for k, item := range a {
		items[k] = *flattenRouterBgpPeersAdvertisedIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterBgpPeersAdvertisedIPRangesSlice flattens the contents of RouterBgpPeersAdvertisedIPRanges from a JSON
// response object.
func flattenRouterBgpPeersAdvertisedIPRangesSlice(c *Client, i interface{}) []RouterBgpPeersAdvertisedIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterBgpPeersAdvertisedIPRanges{}
	}

	if len(a) == 0 {
		return []RouterBgpPeersAdvertisedIPRanges{}
	}

	items := make([]RouterBgpPeersAdvertisedIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterBgpPeersAdvertisedIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterBgpPeersAdvertisedIPRanges expands an instance of RouterBgpPeersAdvertisedIPRanges into a JSON
// request object.
func expandRouterBgpPeersAdvertisedIPRanges(c *Client, f *RouterBgpPeersAdvertisedIPRanges) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Range; !dcl.IsEmptyValueIndirect(v) {
		m["range"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenRouterBgpPeersAdvertisedIPRanges flattens an instance of RouterBgpPeersAdvertisedIPRanges from a JSON
// response object.
func flattenRouterBgpPeersAdvertisedIPRanges(c *Client, i interface{}) *RouterBgpPeersAdvertisedIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterBgpPeersAdvertisedIPRanges{}
	r.Range = dcl.FlattenString(m["range"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandRouterBgpMap expands the contents of RouterBgp into a JSON
// request object.
func expandRouterBgpMap(c *Client, f map[string]RouterBgp) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterBgp(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterBgpSlice expands the contents of RouterBgp into a JSON
// request object.
func expandRouterBgpSlice(c *Client, f []RouterBgp) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterBgp(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterBgpMap flattens the contents of RouterBgp from a JSON
// response object.
func flattenRouterBgpMap(c *Client, i interface{}) map[string]RouterBgp {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterBgp{}
	}

	if len(a) == 0 {
		return map[string]RouterBgp{}
	}

	items := make(map[string]RouterBgp)
	for k, item := range a {
		items[k] = *flattenRouterBgp(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterBgpSlice flattens the contents of RouterBgp from a JSON
// response object.
func flattenRouterBgpSlice(c *Client, i interface{}) []RouterBgp {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterBgp{}
	}

	if len(a) == 0 {
		return []RouterBgp{}
	}

	items := make([]RouterBgp, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterBgp(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterBgp expands an instance of RouterBgp into a JSON
// request object.
func expandRouterBgp(c *Client, f *RouterBgp) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Asn; !dcl.IsEmptyValueIndirect(v) {
		m["asn"] = v
	}
	if v := f.AdvertiseMode; !dcl.IsEmptyValueIndirect(v) {
		m["advertiseMode"] = v
	}
	if v := f.AdvertisedGroups; !dcl.IsEmptyValueIndirect(v) {
		m["advertisedGroups"] = v
	}
	if v, err := expandRouterBgpAdvertisedIPRangesSlice(c, f.AdvertisedIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding AdvertisedIPRanges into advertisedIpRanges: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["advertisedIpRanges"] = v
	}

	return m, nil
}

// flattenRouterBgp flattens an instance of RouterBgp from a JSON
// response object.
func flattenRouterBgp(c *Client, i interface{}) *RouterBgp {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterBgp{}
	r.Asn = dcl.FlattenInteger(m["asn"])
	r.AdvertiseMode = flattenRouterBgpAdvertiseModeEnum(m["advertiseMode"])
	r.AdvertisedGroups = dcl.FlattenStringSlice(m["advertisedGroups"])
	r.AdvertisedIPRanges = flattenRouterBgpAdvertisedIPRangesSlice(c, m["advertisedIpRanges"])

	return r
}

// expandRouterBgpAdvertisedIPRangesMap expands the contents of RouterBgpAdvertisedIPRanges into a JSON
// request object.
func expandRouterBgpAdvertisedIPRangesMap(c *Client, f map[string]RouterBgpAdvertisedIPRanges) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterBgpAdvertisedIPRanges(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterBgpAdvertisedIPRangesSlice expands the contents of RouterBgpAdvertisedIPRanges into a JSON
// request object.
func expandRouterBgpAdvertisedIPRangesSlice(c *Client, f []RouterBgpAdvertisedIPRanges) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterBgpAdvertisedIPRanges(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterBgpAdvertisedIPRangesMap flattens the contents of RouterBgpAdvertisedIPRanges from a JSON
// response object.
func flattenRouterBgpAdvertisedIPRangesMap(c *Client, i interface{}) map[string]RouterBgpAdvertisedIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterBgpAdvertisedIPRanges{}
	}

	if len(a) == 0 {
		return map[string]RouterBgpAdvertisedIPRanges{}
	}

	items := make(map[string]RouterBgpAdvertisedIPRanges)
	for k, item := range a {
		items[k] = *flattenRouterBgpAdvertisedIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterBgpAdvertisedIPRangesSlice flattens the contents of RouterBgpAdvertisedIPRanges from a JSON
// response object.
func flattenRouterBgpAdvertisedIPRangesSlice(c *Client, i interface{}) []RouterBgpAdvertisedIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterBgpAdvertisedIPRanges{}
	}

	if len(a) == 0 {
		return []RouterBgpAdvertisedIPRanges{}
	}

	items := make([]RouterBgpAdvertisedIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterBgpAdvertisedIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterBgpAdvertisedIPRanges expands an instance of RouterBgpAdvertisedIPRanges into a JSON
// request object.
func expandRouterBgpAdvertisedIPRanges(c *Client, f *RouterBgpAdvertisedIPRanges) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Range; !dcl.IsEmptyValueIndirect(v) {
		m["range"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenRouterBgpAdvertisedIPRanges flattens an instance of RouterBgpAdvertisedIPRanges from a JSON
// response object.
func flattenRouterBgpAdvertisedIPRanges(c *Client, i interface{}) *RouterBgpAdvertisedIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterBgpAdvertisedIPRanges{}
	r.Range = dcl.FlattenString(m["range"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// flattenRouterNatsLogConfigFilterEnumSlice flattens the contents of RouterNatsLogConfigFilterEnum from a JSON
// response object.
func flattenRouterNatsLogConfigFilterEnumSlice(c *Client, i interface{}) []RouterNatsLogConfigFilterEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterNatsLogConfigFilterEnum{}
	}

	if len(a) == 0 {
		return []RouterNatsLogConfigFilterEnum{}
	}

	items := make([]RouterNatsLogConfigFilterEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterNatsLogConfigFilterEnum(item.(interface{})))
	}

	return items
}

// flattenRouterNatsLogConfigFilterEnum asserts that an interface is a string, and returns a
// pointer to a *RouterNatsLogConfigFilterEnum with the same value as that string.
func flattenRouterNatsLogConfigFilterEnum(i interface{}) *RouterNatsLogConfigFilterEnum {
	s, ok := i.(string)
	if !ok {
		return RouterNatsLogConfigFilterEnumRef("")
	}

	return RouterNatsLogConfigFilterEnumRef(s)
}

// flattenRouterNatsSourceSubnetworkIPRangesToNatEnumSlice flattens the contents of RouterNatsSourceSubnetworkIPRangesToNatEnum from a JSON
// response object.
func flattenRouterNatsSourceSubnetworkIPRangesToNatEnumSlice(c *Client, i interface{}) []RouterNatsSourceSubnetworkIPRangesToNatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterNatsSourceSubnetworkIPRangesToNatEnum{}
	}

	if len(a) == 0 {
		return []RouterNatsSourceSubnetworkIPRangesToNatEnum{}
	}

	items := make([]RouterNatsSourceSubnetworkIPRangesToNatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterNatsSourceSubnetworkIPRangesToNatEnum(item.(interface{})))
	}

	return items
}

// flattenRouterNatsSourceSubnetworkIPRangesToNatEnum asserts that an interface is a string, and returns a
// pointer to a *RouterNatsSourceSubnetworkIPRangesToNatEnum with the same value as that string.
func flattenRouterNatsSourceSubnetworkIPRangesToNatEnum(i interface{}) *RouterNatsSourceSubnetworkIPRangesToNatEnum {
	s, ok := i.(string)
	if !ok {
		return RouterNatsSourceSubnetworkIPRangesToNatEnumRef("")
	}

	return RouterNatsSourceSubnetworkIPRangesToNatEnumRef(s)
}

// flattenRouterNatsNatIPAllocateOptionEnumSlice flattens the contents of RouterNatsNatIPAllocateOptionEnum from a JSON
// response object.
func flattenRouterNatsNatIPAllocateOptionEnumSlice(c *Client, i interface{}) []RouterNatsNatIPAllocateOptionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterNatsNatIPAllocateOptionEnum{}
	}

	if len(a) == 0 {
		return []RouterNatsNatIPAllocateOptionEnum{}
	}

	items := make([]RouterNatsNatIPAllocateOptionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterNatsNatIPAllocateOptionEnum(item.(interface{})))
	}

	return items
}

// flattenRouterNatsNatIPAllocateOptionEnum asserts that an interface is a string, and returns a
// pointer to a *RouterNatsNatIPAllocateOptionEnum with the same value as that string.
func flattenRouterNatsNatIPAllocateOptionEnum(i interface{}) *RouterNatsNatIPAllocateOptionEnum {
	s, ok := i.(string)
	if !ok {
		return RouterNatsNatIPAllocateOptionEnumRef("")
	}

	return RouterNatsNatIPAllocateOptionEnumRef(s)
}

// flattenRouterInterfacesManagementTypeEnumSlice flattens the contents of RouterInterfacesManagementTypeEnum from a JSON
// response object.
func flattenRouterInterfacesManagementTypeEnumSlice(c *Client, i interface{}) []RouterInterfacesManagementTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterInterfacesManagementTypeEnum{}
	}

	if len(a) == 0 {
		return []RouterInterfacesManagementTypeEnum{}
	}

	items := make([]RouterInterfacesManagementTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterInterfacesManagementTypeEnum(item.(interface{})))
	}

	return items
}

// flattenRouterInterfacesManagementTypeEnum asserts that an interface is a string, and returns a
// pointer to a *RouterInterfacesManagementTypeEnum with the same value as that string.
func flattenRouterInterfacesManagementTypeEnum(i interface{}) *RouterInterfacesManagementTypeEnum {
	s, ok := i.(string)
	if !ok {
		return RouterInterfacesManagementTypeEnumRef("")
	}

	return RouterInterfacesManagementTypeEnumRef(s)
}

// flattenRouterBgpPeersAdvertisedGroupsEnumSlice flattens the contents of RouterBgpPeersAdvertisedGroupsEnum from a JSON
// response object.
func flattenRouterBgpPeersAdvertisedGroupsEnumSlice(c *Client, i interface{}) []RouterBgpPeersAdvertisedGroupsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterBgpPeersAdvertisedGroupsEnum{}
	}

	if len(a) == 0 {
		return []RouterBgpPeersAdvertisedGroupsEnum{}
	}

	items := make([]RouterBgpPeersAdvertisedGroupsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterBgpPeersAdvertisedGroupsEnum(item.(interface{})))
	}

	return items
}

// flattenRouterBgpPeersAdvertisedGroupsEnum asserts that an interface is a string, and returns a
// pointer to a *RouterBgpPeersAdvertisedGroupsEnum with the same value as that string.
func flattenRouterBgpPeersAdvertisedGroupsEnum(i interface{}) *RouterBgpPeersAdvertisedGroupsEnum {
	s, ok := i.(string)
	if !ok {
		return RouterBgpPeersAdvertisedGroupsEnumRef("")
	}

	return RouterBgpPeersAdvertisedGroupsEnumRef(s)
}

// flattenRouterBgpAdvertiseModeEnumSlice flattens the contents of RouterBgpAdvertiseModeEnum from a JSON
// response object.
func flattenRouterBgpAdvertiseModeEnumSlice(c *Client, i interface{}) []RouterBgpAdvertiseModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterBgpAdvertiseModeEnum{}
	}

	if len(a) == 0 {
		return []RouterBgpAdvertiseModeEnum{}
	}

	items := make([]RouterBgpAdvertiseModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterBgpAdvertiseModeEnum(item.(interface{})))
	}

	return items
}

// flattenRouterBgpAdvertiseModeEnum asserts that an interface is a string, and returns a
// pointer to a *RouterBgpAdvertiseModeEnum with the same value as that string.
func flattenRouterBgpAdvertiseModeEnum(i interface{}) *RouterBgpAdvertiseModeEnum {
	s, ok := i.(string)
	if !ok {
		return RouterBgpAdvertiseModeEnumRef("")
	}

	return RouterBgpAdvertiseModeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Router) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRouter(b, c)
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
