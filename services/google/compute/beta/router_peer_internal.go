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

func (r *RouterPeer) validate() error {

	return nil
}
func (r *RouterPeerAdvertisedIPRanges) validate() error {
	return nil
}

func routerPeerGetURL(userBasePath string, r *RouterPeer) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"router":  dcl.ValueOrEmptyString(r.Router),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{router}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func routerPeerListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func routerPeerCreateURL(userBasePath, project, region, router string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
		"router":  router,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{router}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func routerPeerDeleteURL(userBasePath string, r *RouterPeer) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"router":  dcl.ValueOrEmptyString(r.Router),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{router}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// routerPeerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type routerPeerApiOperation interface {
	do(context.Context, *RouterPeer, *Client) error
}

// newUpdateRouterPeerUpdateRequest creates a request for an
// RouterPeer resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateRouterPeerUpdateRequest(ctx context.Context, f *RouterPeer, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.InterfaceName; !dcl.IsEmptyValueIndirect(v) {
		req["interfaceName"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		req["ipAddress"] = v
	}
	if v := f.PeerIPAddress; !dcl.IsEmptyValueIndirect(v) {
		req["peerIpAddress"] = v
	}
	if v := f.PeerAsn; !dcl.IsEmptyValueIndirect(v) {
		req["peerAsn"] = v
	}
	if v := f.AdvertisedRoutePriority; !dcl.IsEmptyValueIndirect(v) {
		req["advertisedRoutePriority"] = v
	}
	if v := f.AdvertiseMode; !dcl.IsEmptyValueIndirect(v) {
		req["advertiseMode"] = v
	}
	if v := f.AdvertisedGroups; !dcl.IsEmptyValueIndirect(v) {
		req["advertisedGroups"] = v
	}
	if v, err := expandRouterPeerAdvertisedIPRangesSlice(c, f.AdvertisedIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding AdvertisedIPRanges into advertisedIpRanges: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["advertisedIpRanges"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		req["region"] = v
	}
	return req, nil
}

// marshalUpdateRouterPeerUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateRouterPeerUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateRouterPeerUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateRouterPeerUpdateOperation) do(ctx context.Context, r *RouterPeer, c *Client) error {
	_, err := c.GetRouterPeer(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateRouterPeerUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateRouterPeerUpdateRequest(c, req)
	if err != nil {
		return err
	}
	body, err = routerPeerNestedRequest(ctx, r, c, false)
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

func (c *Client) listRouterPeerRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := routerPeerListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RouterPeerMaxPage {
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

type listRouterPeerOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listRouterPeer(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*RouterPeer, string, error) {
	b, err := c.listRouterPeerRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRouterPeerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*RouterPeer
	for _, v := range m.Items {
		res, err := unmarshalMapRouterPeer(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRouterPeer(ctx context.Context, f func(*RouterPeer) bool, resources []*RouterPeer) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRouterPeer(ctx, res)
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

type deleteRouterPeerOperation struct{}

func (op *deleteRouterPeerOperation) do(ctx context.Context, r *RouterPeer, c *Client) error {
	r, err := c.GetRouterPeer(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("RouterPeer not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetRouterPeer checking for existence. error: %v", err)
		return err
	}

	u, err := routerPeerDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	b, err := routerPeerNestedRequest(ctx, r, c, true)
	if err != nil {
		return err
	}
	body = bytes.NewBuffer(b)

	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, body, c.Config.RetryProvider)
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
		_, err = c.GetRouterPeer(ctx, r.URLNormalized())
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
type createRouterPeerOperation struct {
	response map[string]interface{}
}

func (op *createRouterPeerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRouterPeerOperation) do(ctx context.Context, r *RouterPeer, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region, router := r.createFields()
	u, err := routerPeerCreateURL(c.Config.BasePath, project, region, router)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	req, err = routerPeerNestedRequest(ctx, r, c, false)
	if err != nil {
		return err
	}

	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(req), c.Config.RetryProvider)
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

	if _, err := c.GetRouterPeer(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRouterPeerRaw(ctx context.Context, r *RouterPeer) ([]byte, error) {

	u, err := routerPeerGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) routerPeerDiffsForRawDesired(ctx context.Context, rawDesired *RouterPeer, opts ...dcl.ApplyOption) (initial, desired *RouterPeer, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *RouterPeer
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*RouterPeer); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected RouterPeer, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRouterPeer(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a RouterPeer resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve RouterPeer resource: %v", err)
		}
		c.Config.Logger.Info("Found that RouterPeer resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRouterPeerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for RouterPeer: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for RouterPeer: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRouterPeerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for RouterPeer: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRouterPeerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for RouterPeer: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRouterPeer(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func fetchRouterPeerIndexFromParentArray(r *RouterPeer, arr []interface{}) (int, error) {
	for i, in := range arr {
		item := in.(map[string]interface{})

		if dcl.ValueOrEmptyString(r.Name) != item["name"] {
			continue
		}

		return i, nil
	}
	return -1, dcl.NotFoundError{Cause: fmt.Errorf("No such item found in %v with matching keys [name]", arr)}
}

func routerPeerFromParent(r *RouterPeer, parent map[string]interface{}) (map[string]interface{}, error) {
	itemAtQuery, err := dcl.GetMapEntry(parent, []string{"bgpPeers"})
	if err != nil {
		return nil, dcl.NotFoundError{Cause: err}
	}
	arr := itemAtQuery.([]interface{})
	index, err := fetchRouterPeerIndexFromParentArray(r, arr)
	if err != nil {
		return nil, err
	}
	return arr[index].(map[string]interface{}), nil
}

// routerPeerNestedRequest takes in a {$.Name}} + client and returns a
// JSON request ready to be sent.
// This request is assembled by fetching the parent resource and finding/replacing/deleting
// the RouterPeer resource r as dictated by 'delete' and the presence r in the parent resource.
func routerPeerNestedRequest(ctx context.Context, r *RouterPeer, c *Client, delete bool) ([]byte, error) {
	// Fetch parent resource.
	pb, err := c.getRouterPeerRaw(ctx, r.URLNormalized())
	if err != nil {
		return nil, err
	}
	var parent map[string]interface{}
	err = json.Unmarshal(pb, &parent)
	if err != nil {
		return nil, err
	}

	arrayAtQuery, err := dcl.GetMapEntry(parent, []string{"bgpPeers"})
	// Error returned means that no field exists currently exist.
	if err != nil {
		arrayAtQuery = make([]interface{}, 0)
	}
	arr := arrayAtQuery.([]interface{})

	rAsMap, err := expandRouterPeer(c, r)
	if err != nil {
		return nil, err
	}

	index, err := fetchRouterPeerIndexFromParentArray(r, arr)

	// Not found.
	if err != nil {
		arr = append(arr, rAsMap)
	} else {
		if delete {
			arr[index] = nil
		} else {
			arr[index] = rAsMap
		}
	}

	err = dcl.PutMapEntry(parent, []string{"bgpPeers"}, arr)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(parent)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func canonicalizeRouterPeerInitialState(rawInitial, rawDesired *RouterPeer) (*RouterPeer, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRouterPeerDesiredState(rawDesired, rawInitial *RouterPeer, opts ...dcl.ApplyOption) (*RouterPeer, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	if dcl.NameToSelfLink(rawDesired.Router, rawInitial.Router) {
		rawDesired.Router = rawInitial.Router
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.InterfaceName, rawInitial.InterfaceName) {
		rawDesired.InterfaceName = rawInitial.InterfaceName
	}
	if dcl.StringCanonicalize(rawDesired.IPAddress, rawInitial.IPAddress) {
		rawDesired.IPAddress = rawInitial.IPAddress
	}
	if dcl.StringCanonicalize(rawDesired.PeerIPAddress, rawInitial.PeerIPAddress) {
		rawDesired.PeerIPAddress = rawInitial.PeerIPAddress
	}
	if dcl.IsZeroValue(rawDesired.PeerAsn) {
		rawDesired.PeerAsn = rawInitial.PeerAsn
	}
	if dcl.IsZeroValue(rawDesired.AdvertisedRoutePriority) {
		rawDesired.AdvertisedRoutePriority = rawInitial.AdvertisedRoutePriority
	}
	if dcl.StringCanonicalize(rawDesired.AdvertiseMode, rawInitial.AdvertiseMode) {
		rawDesired.AdvertiseMode = rawInitial.AdvertiseMode
	}
	if dcl.IsZeroValue(rawDesired.AdvertisedGroups) {
		rawDesired.AdvertisedGroups = rawInitial.AdvertisedGroups
	}
	if dcl.IsZeroValue(rawDesired.AdvertisedIPRanges) {
		rawDesired.AdvertisedIPRanges = rawInitial.AdvertisedIPRanges
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeRouterPeerNewState(c *Client, rawNew, rawDesired *RouterPeer) (*RouterPeer, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
	}

	rawNew.Router = rawDesired.Router

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InterfaceName) && dcl.IsEmptyValueIndirect(rawDesired.InterfaceName) {
		rawNew.InterfaceName = rawDesired.InterfaceName
	} else {
		if dcl.StringCanonicalize(rawDesired.InterfaceName, rawNew.InterfaceName) {
			rawNew.InterfaceName = rawDesired.InterfaceName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddress) && dcl.IsEmptyValueIndirect(rawDesired.IPAddress) {
		rawNew.IPAddress = rawDesired.IPAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.IPAddress, rawNew.IPAddress) {
			rawNew.IPAddress = rawDesired.IPAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerIPAddress) && dcl.IsEmptyValueIndirect(rawDesired.PeerIPAddress) {
		rawNew.PeerIPAddress = rawDesired.PeerIPAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.PeerIPAddress, rawNew.PeerIPAddress) {
			rawNew.PeerIPAddress = rawDesired.PeerIPAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerAsn) && dcl.IsEmptyValueIndirect(rawDesired.PeerAsn) {
		rawNew.PeerAsn = rawDesired.PeerAsn
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AdvertisedRoutePriority) && dcl.IsEmptyValueIndirect(rawDesired.AdvertisedRoutePriority) {
		rawNew.AdvertisedRoutePriority = rawDesired.AdvertisedRoutePriority
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AdvertiseMode) && dcl.IsEmptyValueIndirect(rawDesired.AdvertiseMode) {
		rawNew.AdvertiseMode = rawDesired.AdvertiseMode
	} else {
		if dcl.StringCanonicalize(rawDesired.AdvertiseMode, rawNew.AdvertiseMode) {
			rawNew.AdvertiseMode = rawDesired.AdvertiseMode
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ManagementType) && dcl.IsEmptyValueIndirect(rawDesired.ManagementType) {
		rawNew.ManagementType = rawDesired.ManagementType
	} else {
		if dcl.StringCanonicalize(rawDesired.ManagementType, rawNew.ManagementType) {
			rawNew.ManagementType = rawDesired.ManagementType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AdvertisedGroups) && dcl.IsEmptyValueIndirect(rawDesired.AdvertisedGroups) {
		rawNew.AdvertisedGroups = rawDesired.AdvertisedGroups
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AdvertisedIPRanges) && dcl.IsEmptyValueIndirect(rawDesired.AdvertisedIPRanges) {
		rawNew.AdvertisedIPRanges = rawDesired.AdvertisedIPRanges
	} else {
		rawNew.AdvertisedIPRanges = canonicalizeNewRouterPeerAdvertisedIPRangesSlice(c, rawDesired.AdvertisedIPRanges, rawNew.AdvertisedIPRanges)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.NameToSelfLink(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeRouterPeerAdvertisedIPRanges(des, initial *RouterPeerAdvertisedIPRanges, opts ...dcl.ApplyOption) *RouterPeerAdvertisedIPRanges {
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

func canonicalizeNewRouterPeerAdvertisedIPRanges(c *Client, des, nw *RouterPeerAdvertisedIPRanges) *RouterPeerAdvertisedIPRanges {
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

func canonicalizeNewRouterPeerAdvertisedIPRangesSet(c *Client, des, nw []RouterPeerAdvertisedIPRanges) []RouterPeerAdvertisedIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []RouterPeerAdvertisedIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareRouterPeerAdvertisedIPRangesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewRouterPeerAdvertisedIPRangesSlice(c *Client, des, nw []RouterPeerAdvertisedIPRanges) []RouterPeerAdvertisedIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RouterPeerAdvertisedIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouterPeerAdvertisedIPRanges(c, &d, &n))
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
func diffRouterPeer(c *Client, desired, actual *RouterPeer, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.CreationTimestamp, actual.CreationTimestamp, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTimestamp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Router, actual.Router, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Router")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InterfaceName, actual.InterfaceName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("InterfaceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAddress, actual.IPAddress, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("IpAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeerIPAddress, actual.PeerIPAddress, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("PeerIpAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PeerAsn, actual.PeerAsn, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("PeerAsn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdvertisedRoutePriority, actual.AdvertisedRoutePriority, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("AdvertisedRoutePriority")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdvertiseMode, actual.AdvertiseMode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("AdvertiseMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManagementType, actual.ManagementType, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagementType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdvertisedGroups, actual.AdvertisedGroups, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("AdvertisedGroups")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdvertisedIPRanges, actual.AdvertisedIPRanges, dcl.Info{ObjectFunction: compareRouterPeerAdvertisedIPRangesNewStyle, EmptyObject: EmptyRouterPeerAdvertisedIPRanges, OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("AdvertisedIpRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareRouterPeerAdvertisedIPRangesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RouterPeerAdvertisedIPRanges)
	if !ok {
		desiredNotPointer, ok := d.(RouterPeerAdvertisedIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RouterPeerAdvertisedIPRanges or *RouterPeerAdvertisedIPRanges", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RouterPeerAdvertisedIPRanges)
	if !ok {
		actualNotPointer, ok := a.(RouterPeerAdvertisedIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RouterPeerAdvertisedIPRanges", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Range, actual.Range, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("Range")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateRouterPeerUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *RouterPeer) getFields() (string, string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Router), dcl.ValueOrEmptyString(n.Name)
}

func (r *RouterPeer) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Router)
}

func (r *RouterPeer) deleteFields() (string, string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Router), dcl.ValueOrEmptyString(n.Name)
}

func (r *RouterPeer) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"router":  dcl.ValueOrEmptyString(n.Router),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/routers/{{router}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the RouterPeer resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *RouterPeer) marshal(c *Client) ([]byte, error) {
	m, err := expandRouterPeer(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling RouterPeer: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalRouterPeer decodes JSON responses into the RouterPeer resource schema.
func unmarshalRouterPeer(b []byte, c *Client) (*RouterPeer, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRouterPeer(m, c)
}

func unmarshalMapRouterPeer(m map[string]interface{}, c *Client) (*RouterPeer, error) {

	return flattenRouterPeer(c, m), nil
}

// expandRouterPeer expands RouterPeer into a JSON request object.
func expandRouterPeer(c *Client, f *RouterPeer) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Router into router: %w", err)
	} else if v != nil {
		m["router"] = v
	}
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
		m["peerIpAddress"] = v
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
	if v, err := expandRouterPeerAdvertisedIPRangesSlice(c, f.AdvertisedIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding AdvertisedIPRanges into advertisedIpRanges: %w", err)
	} else if v != nil {
		m["advertisedIpRanges"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenRouterPeer flattens RouterPeer from a JSON request object into the
// RouterPeer type.
func flattenRouterPeer(c *Client, i interface{}) *RouterPeer {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &RouterPeer{}
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Router = dcl.FlattenString(m["router"])
	res.Name = dcl.FlattenString(m["name"])
	res.InterfaceName = dcl.FlattenString(m["interfaceName"])
	res.IPAddress = dcl.FlattenString(m["ipAddress"])
	res.PeerIPAddress = dcl.FlattenString(m["peerIpAddress"])
	res.PeerAsn = dcl.FlattenInteger(m["peerAsn"])
	res.AdvertisedRoutePriority = dcl.FlattenInteger(m["advertisedRoutePriority"])
	res.AdvertiseMode = dcl.FlattenString(m["advertiseMode"])
	res.ManagementType = dcl.FlattenString(m["managementType"])
	res.AdvertisedGroups = dcl.FlattenStringSlice(m["advertisedGroups"])
	res.AdvertisedIPRanges = flattenRouterPeerAdvertisedIPRangesSlice(c, m["advertisedIpRanges"])
	res.Region = dcl.FlattenString(m["region"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandRouterPeerAdvertisedIPRangesMap expands the contents of RouterPeerAdvertisedIPRanges into a JSON
// request object.
func expandRouterPeerAdvertisedIPRangesMap(c *Client, f map[string]RouterPeerAdvertisedIPRanges) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouterPeerAdvertisedIPRanges(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouterPeerAdvertisedIPRangesSlice expands the contents of RouterPeerAdvertisedIPRanges into a JSON
// request object.
func expandRouterPeerAdvertisedIPRangesSlice(c *Client, f []RouterPeerAdvertisedIPRanges) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouterPeerAdvertisedIPRanges(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouterPeerAdvertisedIPRangesMap flattens the contents of RouterPeerAdvertisedIPRanges from a JSON
// response object.
func flattenRouterPeerAdvertisedIPRangesMap(c *Client, i interface{}) map[string]RouterPeerAdvertisedIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouterPeerAdvertisedIPRanges{}
	}

	if len(a) == 0 {
		return map[string]RouterPeerAdvertisedIPRanges{}
	}

	items := make(map[string]RouterPeerAdvertisedIPRanges)
	for k, item := range a {
		items[k] = *flattenRouterPeerAdvertisedIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouterPeerAdvertisedIPRangesSlice flattens the contents of RouterPeerAdvertisedIPRanges from a JSON
// response object.
func flattenRouterPeerAdvertisedIPRangesSlice(c *Client, i interface{}) []RouterPeerAdvertisedIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []RouterPeerAdvertisedIPRanges{}
	}

	if len(a) == 0 {
		return []RouterPeerAdvertisedIPRanges{}
	}

	items := make([]RouterPeerAdvertisedIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouterPeerAdvertisedIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouterPeerAdvertisedIPRanges expands an instance of RouterPeerAdvertisedIPRanges into a JSON
// request object.
func expandRouterPeerAdvertisedIPRanges(c *Client, f *RouterPeerAdvertisedIPRanges) (map[string]interface{}, error) {
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

// flattenRouterPeerAdvertisedIPRanges flattens an instance of RouterPeerAdvertisedIPRanges from a JSON
// response object.
func flattenRouterPeerAdvertisedIPRanges(c *Client, i interface{}) *RouterPeerAdvertisedIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouterPeerAdvertisedIPRanges{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRouterPeerAdvertisedIPRanges
	}
	r.Range = dcl.FlattenString(m["range"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *RouterPeer) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRouterPeer(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
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
		if nr.Router == nil && ncr.Router == nil {
			c.Config.Logger.Info("Both Router fields null - considering equal.")
		} else if nr.Router == nil || ncr.Router == nil {
			c.Config.Logger.Info("Only one Router field is null - considering unequal.")
			return false
		} else if *nr.Router != *ncr.Router {
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

type routerPeerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         routerPeerApiOperation
}

func convertFieldDiffToRouterPeerOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]routerPeerDiff, error) {
	var diffs []routerPeerDiff
	for _, op := range ops {
		diff := routerPeerDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTorouterPeerApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTorouterPeerApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (routerPeerApiOperation, error) {
	switch op {

	case "updateRouterPeerUpdateOperation":
		return &updateRouterPeerUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
