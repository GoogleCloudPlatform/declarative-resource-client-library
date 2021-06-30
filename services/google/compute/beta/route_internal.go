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

func (r *Route) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "network"); err != nil {
		return err
	}
	if err := dcl.Required(r, "destRange"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}
func (r *RouteWarning) validate() error {
	return nil
}

func routeGetURL(userBasePath string, r *Route) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/routes/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func routeListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/routes", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func routeCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/routes", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func routeDeleteURL(userBasePath string, r *Route) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/routes/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// routeApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type routeApiOperation interface {
	do(context.Context, *Route, *Client) error
}

func (c *Client) listRouteRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := routeListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RouteMaxPage {
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

type listRouteOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listRoute(ctx context.Context, project, pageToken string, pageSize int32) ([]*Route, string, error) {
	b, err := c.listRouteRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRouteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Route
	for _, v := range m.Items {
		res, err := unmarshalMapRoute(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRoute(ctx context.Context, f func(*Route) bool, resources []*Route) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRoute(ctx, res)
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

type deleteRouteOperation struct{}

func (op *deleteRouteOperation) do(ctx context.Context, r *Route, c *Client) error {
	r, err := c.GetRoute(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Route not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetRoute checking for existence. error: %v", err)
		return err
	}

	u, err := routeDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetRoute(ctx, r.URLNormalized())
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
type createRouteOperation struct {
	response map[string]interface{}
}

func (op *createRouteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRouteOperation) do(ctx context.Context, r *Route, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := routeCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetRoute(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRouteRaw(ctx context.Context, r *Route) ([]byte, error) {
	if dcl.IsZeroValue(r.Priority) {
		r.Priority = dcl.Int64(1000)
	}

	u, err := routeGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) routeDiffsForRawDesired(ctx context.Context, rawDesired *Route, opts ...dcl.ApplyOption) (initial, desired *Route, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Route
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Route); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Route, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRoute(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Route resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Route resource: %v", err)
		}
		c.Config.Logger.Info("Found that Route resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRouteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Route: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Route: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRouteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Route: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRouteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Route: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRoute(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeRouteInitialState(rawInitial, rawDesired *Route) (*Route, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRouteDesiredState(rawDesired, rawInitial *Route, opts ...dcl.ApplyOption) (*Route, error) {

	if dcl.IsZeroValue(rawDesired.Priority) {
		rawDesired.Priority = dcl.Int64(1000)
	}

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
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.Tag) {
		rawDesired.Tag = rawInitial.Tag
	}
	if dcl.StringCanonicalize(rawDesired.DestRange, rawInitial.DestRange) {
		rawDesired.DestRange = rawInitial.DestRange
	}
	if dcl.IsZeroValue(rawDesired.Priority) {
		rawDesired.Priority = rawInitial.Priority
	}
	if dcl.StringCanonicalize(rawDesired.NextHopInstance, rawInitial.NextHopInstance) {
		rawDesired.NextHopInstance = rawInitial.NextHopInstance
	}
	if dcl.StringCanonicalize(rawDesired.NextHopIP, rawInitial.NextHopIP) {
		rawDesired.NextHopIP = rawInitial.NextHopIP
	}
	if dcl.StringCanonicalize(rawDesired.NextHopNetwork, rawInitial.NextHopNetwork) {
		rawDesired.NextHopNetwork = rawInitial.NextHopNetwork
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.NextHopGateway, rawInitial.NextHopGateway) {
		rawDesired.NextHopGateway = rawInitial.NextHopGateway
	}
	if dcl.StringCanonicalize(rawDesired.NextHopIlb, rawInitial.NextHopIlb) {
		rawDesired.NextHopIlb = rawInitial.NextHopIlb
	}
	if dcl.StringCanonicalize(rawDesired.NextHopVpnTunnel, rawInitial.NextHopVpnTunnel) {
		rawDesired.NextHopVpnTunnel = rawInitial.NextHopVpnTunnel
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeRouteNewState(c *Client, rawNew, rawDesired *Route) (*Route, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Tag) && dcl.IsEmptyValueIndirect(rawDesired.Tag) {
		rawNew.Tag = rawDesired.Tag
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DestRange) && dcl.IsEmptyValueIndirect(rawDesired.DestRange) {
		rawNew.DestRange = rawDesired.DestRange
	} else {
		if dcl.StringCanonicalize(rawDesired.DestRange, rawNew.DestRange) {
			rawNew.DestRange = rawDesired.DestRange
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Priority) && dcl.IsEmptyValueIndirect(rawDesired.Priority) {
		rawNew.Priority = rawDesired.Priority
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopInstance) && dcl.IsEmptyValueIndirect(rawDesired.NextHopInstance) {
		rawNew.NextHopInstance = rawDesired.NextHopInstance
	} else {
		if dcl.StringCanonicalize(rawDesired.NextHopInstance, rawNew.NextHopInstance) {
			rawNew.NextHopInstance = rawDesired.NextHopInstance
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopIP) && dcl.IsEmptyValueIndirect(rawDesired.NextHopIP) {
		rawNew.NextHopIP = rawDesired.NextHopIP
	} else {
		if dcl.StringCanonicalize(rawDesired.NextHopIP, rawNew.NextHopIP) {
			rawNew.NextHopIP = rawDesired.NextHopIP
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopNetwork) && dcl.IsEmptyValueIndirect(rawDesired.NextHopNetwork) {
		rawNew.NextHopNetwork = rawDesired.NextHopNetwork
	} else {
		if dcl.StringCanonicalize(rawDesired.NextHopNetwork, rawNew.NextHopNetwork) {
			rawNew.NextHopNetwork = rawDesired.NextHopNetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopGateway) && dcl.IsEmptyValueIndirect(rawDesired.NextHopGateway) {
		rawNew.NextHopGateway = rawDesired.NextHopGateway
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.NextHopGateway, rawNew.NextHopGateway) {
			rawNew.NextHopGateway = rawDesired.NextHopGateway
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopPeering) && dcl.IsEmptyValueIndirect(rawDesired.NextHopPeering) {
		rawNew.NextHopPeering = rawDesired.NextHopPeering
	} else {
		if dcl.StringCanonicalize(rawDesired.NextHopPeering, rawNew.NextHopPeering) {
			rawNew.NextHopPeering = rawDesired.NextHopPeering
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopIlb) && dcl.IsEmptyValueIndirect(rawDesired.NextHopIlb) {
		rawNew.NextHopIlb = rawDesired.NextHopIlb
	} else {
		if dcl.StringCanonicalize(rawDesired.NextHopIlb, rawNew.NextHopIlb) {
			rawNew.NextHopIlb = rawDesired.NextHopIlb
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Warning) && dcl.IsEmptyValueIndirect(rawDesired.Warning) {
		rawNew.Warning = rawDesired.Warning
	} else {
		rawNew.Warning = canonicalizeNewRouteWarningSlice(c, rawDesired.Warning, rawNew.Warning)
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextHopVpnTunnel) && dcl.IsEmptyValueIndirect(rawDesired.NextHopVpnTunnel) {
		rawNew.NextHopVpnTunnel = rawDesired.NextHopVpnTunnel
	} else {
		if dcl.StringCanonicalize(rawDesired.NextHopVpnTunnel, rawNew.NextHopVpnTunnel) {
			rawNew.NextHopVpnTunnel = rawDesired.NextHopVpnTunnel
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

	return rawNew, nil
}

func canonicalizeRouteWarning(des, initial *RouteWarning, opts ...dcl.ApplyOption) *RouteWarning {
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

func canonicalizeNewRouteWarning(c *Client, des, nw *RouteWarning) *RouteWarning {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	if dcl.IsZeroValue(nw.Data) {
		nw.Data = des.Data
	}

	return nw
}

func canonicalizeNewRouteWarningSet(c *Client, des, nw []RouteWarning) []RouteWarning {
	if des == nil {
		return nw
	}
	var reorderedNew []RouteWarning
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareRouteWarningNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewRouteWarningSlice(c *Client, des, nw []RouteWarning) []RouteWarning {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RouteWarning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouteWarning(c, &d, &n))
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
func diffRoute(c *Client, desired, actual *Route, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DestRange, actual.DestRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DestRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Priority, actual.Priority, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Priority")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopInstance, actual.NextHopInstance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopIP, actual.NextHopIP, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopNetwork, actual.NextHopNetwork, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopNetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopGateway, actual.NextHopGateway, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopGateway")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopPeering, actual.NextHopPeering, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopPeering")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopIlb, actual.NextHopIlb, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopIlb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Warning, actual.Warning, dcl.Info{OutputOnly: true, ObjectFunction: compareRouteWarningNewStyle, EmptyObject: EmptyRouteWarning, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Warnings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NextHopVpnTunnel, actual.NextHopVpnTunnel, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextHopVpnTunnel")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareRouteWarningNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RouteWarning)
	if !ok {
		desiredNotPointer, ok := d.(RouteWarning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RouteWarning or *RouteWarning", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RouteWarning)
	if !ok {
		actualNotPointer, ok := a.(RouteWarning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RouteWarning", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Data, actual.Data, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Data")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Route) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Route) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Route) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Route) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Route resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Route) marshal(c *Client) ([]byte, error) {
	m, err := expandRoute(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Route: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalRoute decodes JSON responses into the Route resource schema.
func unmarshalRoute(b []byte, c *Client) (*Route, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRoute(m, c)
}

func unmarshalMapRoute(m map[string]interface{}, c *Client) (*Route, error) {
	if v, err := dcl.MapFromListOfKeyValues(m, []string{"warning", "data", "items"}); err != nil {
		return nil, err
	} else {
		dcl.PutMapEntry(
			m,
			[]string{"warning", "data"},
			v,
		)
	}

	return flattenRoute(c, m), nil
}

// expandRoute expands Route into a JSON request object.
func expandRoute(c *Client, f *Route) (map[string]interface{}, error) {
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
	if v, err := dcl.DeriveField("global/networks/%s", f.Network, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if v != nil {
		m["network"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}
	if v := f.DestRange; !dcl.IsEmptyValueIndirect(v) {
		m["destRange"] = v
	}
	if v := f.Priority; !dcl.IsEmptyValueIndirect(v) {
		m["priority"] = v
	}
	if v := f.NextHopInstance; !dcl.IsEmptyValueIndirect(v) {
		m["nextHopInstance"] = v
	}
	if v := f.NextHopIP; !dcl.IsEmptyValueIndirect(v) {
		m["nextHopIp"] = v
	}
	if v := f.NextHopNetwork; !dcl.IsEmptyValueIndirect(v) {
		m["nextHopNetwork"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/global/gateways/%s", f.NextHopGateway, f.Project, f.NextHopGateway); err != nil {
		return nil, fmt.Errorf("error expanding NextHopGateway into nextHopGateway: %w", err)
	} else if v != nil {
		m["nextHopGateway"] = v
	}
	if v := f.NextHopPeering; !dcl.IsEmptyValueIndirect(v) {
		m["nextHopPeering"] = v
	}
	if v := f.NextHopIlb; !dcl.IsEmptyValueIndirect(v) {
		m["nextHopIlb"] = v
	}
	if v, err := expandRouteWarningSlice(c, f.Warning); err != nil {
		return nil, fmt.Errorf("error expanding Warning into warnings: %w", err)
	} else if v != nil {
		m["warnings"] = v
	}
	if v := f.NextHopVpnTunnel; !dcl.IsEmptyValueIndirect(v) {
		m["nextHopVpnTunnel"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenRoute flattens Route from a JSON request object into the
// Route type.
func flattenRoute(c *Client, i interface{}) *Route {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Route{}
	res.Id = dcl.FlattenInteger(m["id"])
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Network = dcl.FlattenString(m["network"])
	res.Tag = dcl.FlattenStringSlice(m["tags"])
	res.DestRange = dcl.FlattenString(m["destRange"])
	res.Priority = dcl.FlattenInteger(m["priority"])
	if _, ok := m["priority"]; !ok {
		c.Config.Logger.Info("Using default value for priority")
		res.Priority = dcl.Int64(1000)
	}
	res.NextHopInstance = dcl.FlattenString(m["nextHopInstance"])
	res.NextHopIP = dcl.FlattenString(m["nextHopIp"])
	res.NextHopNetwork = dcl.FlattenString(m["nextHopNetwork"])
	res.NextHopGateway = dcl.FlattenString(m["nextHopGateway"])
	res.NextHopPeering = dcl.FlattenString(m["nextHopPeering"])
	res.NextHopIlb = dcl.FlattenString(m["nextHopIlb"])
	res.Warning = flattenRouteWarningSlice(c, m["warnings"])
	res.NextHopVpnTunnel = dcl.FlattenString(m["nextHopVpnTunnel"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandRouteWarningMap expands the contents of RouteWarning into a JSON
// request object.
func expandRouteWarningMap(c *Client, f map[string]RouteWarning) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRouteWarning(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRouteWarningSlice expands the contents of RouteWarning into a JSON
// request object.
func expandRouteWarningSlice(c *Client, f []RouteWarning) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRouteWarning(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRouteWarningMap flattens the contents of RouteWarning from a JSON
// response object.
func flattenRouteWarningMap(c *Client, i interface{}) map[string]RouteWarning {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RouteWarning{}
	}

	if len(a) == 0 {
		return map[string]RouteWarning{}
	}

	items := make(map[string]RouteWarning)
	for k, item := range a {
		items[k] = *flattenRouteWarning(c, item.(map[string]interface{}))
	}

	return items
}

// flattenRouteWarningSlice flattens the contents of RouteWarning from a JSON
// response object.
func flattenRouteWarningSlice(c *Client, i interface{}) []RouteWarning {
	a, ok := i.([]interface{})
	if !ok {
		return []RouteWarning{}
	}

	if len(a) == 0 {
		return []RouteWarning{}
	}

	items := make([]RouteWarning, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouteWarning(c, item.(map[string]interface{})))
	}

	return items
}

// expandRouteWarning expands an instance of RouteWarning into a JSON
// request object.
func expandRouteWarning(c *Client, f *RouteWarning) (map[string]interface{}, error) {
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
	if v, err := dcl.ListOfKeyValuesFromMapInItemsStruct(f.Data); err != nil {
		return nil, fmt.Errorf("error expanding Data into data: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["data"] = v
	}

	return m, nil
}

// flattenRouteWarning flattens an instance of RouteWarning from a JSON
// response object.
func flattenRouteWarning(c *Client, i interface{}) *RouteWarning {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RouteWarning{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRouteWarning
	}
	r.Code = flattenRouteWarningCodeEnum(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Data = dcl.FlattenKeyValuePairs(m["data"])

	return r
}

// flattenRouteWarningCodeEnumSlice flattens the contents of RouteWarningCodeEnum from a JSON
// response object.
func flattenRouteWarningCodeEnumSlice(c *Client, i interface{}) []RouteWarningCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RouteWarningCodeEnum{}
	}

	if len(a) == 0 {
		return []RouteWarningCodeEnum{}
	}

	items := make([]RouteWarningCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRouteWarningCodeEnum(item.(interface{})))
	}

	return items
}

// flattenRouteWarningCodeEnum asserts that an interface is a string, and returns a
// pointer to a *RouteWarningCodeEnum with the same value as that string.
func flattenRouteWarningCodeEnum(i interface{}) *RouteWarningCodeEnum {
	s, ok := i.(string)
	if !ok {
		return RouteWarningCodeEnumRef("")
	}

	return RouteWarningCodeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Route) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRoute(b, c)
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

type routeDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         routeApiOperation
}

func convertFieldDiffToRouteOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]routeDiff, error) {
	var diffs []routeDiff
	for _, op := range ops {
		diff := routeDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTorouteApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTorouteApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (routeApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
