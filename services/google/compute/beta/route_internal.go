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
		res := flattenRoute(c, v)
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

	_, err := c.GetRoute(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Route not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetRoute checking for existence. error: %v", err)
		return err
	}

	u, err := routeDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetRoute(ctx, r.urlNormalized())
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

	if _, err := c.GetRoute(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRouteRaw(ctx context.Context, r *Route) ([]byte, error) {
	if dcl.IsZeroValue(r.Priority) {
		r.Priority = dcl.Int64(1000)
	}

	u, err := routeGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) routeDiffsForRawDesired(ctx context.Context, rawDesired *Route, opts ...dcl.ApplyOption) (initial, desired *Route, diffs []routeDiff, err error) {
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
	rawInitial, err := c.GetRoute(ctx, fetchState.urlNormalized())
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
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
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
	if dcl.StringCanonicalize(rawDesired.NextHopPeering, rawInitial.NextHopPeering) {
		rawDesired.NextHopPeering = rawInitial.NextHopPeering
	}
	if dcl.StringCanonicalize(rawDesired.NextHopIlb, rawInitial.NextHopIlb) {
		rawDesired.NextHopIlb = rawInitial.NextHopIlb
	}
	if dcl.IsZeroValue(rawDesired.Warning) {
		rawDesired.Warning = rawInitial.Warning
	}
	if dcl.StringCanonicalize(rawDesired.NextHopVpnTunnel, rawInitial.NextHopVpnTunnel) {
		rawDesired.NextHopVpnTunnel = rawInitial.NextHopVpnTunnel
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
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

	if dcl.IsZeroValue(des.Code) {
		des.Code = initial.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		des.Message = initial.Message
	}
	if dcl.IsZeroValue(des.Data) {
		des.Data = initial.Data
	}

	return des
}

func canonicalizeNewRouteWarning(c *Client, des, nw *RouteWarning) *RouteWarning {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
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
			if !compareRouteWarning(c, &d, &n) {
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
		return des
	}

	var items []RouteWarning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRouteWarning(c, &d, &n))
	}

	return items
}

type routeDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         routeApiOperation
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
func diffRoute(c *Client, desired, actual *Route, opts ...dcl.ApplyOption) ([]routeDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []routeDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.PartialSelfLinkToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !dcl.StringSliceEquals(desired.Tag, actual.Tag) {
		c.Config.Logger.Infof("Detected diff in Tag.\nDESIRED: %v\nACTUAL: %v", desired.Tag, actual.Tag)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "Tag",
		})
	}
	if !dcl.IsZeroValue(desired.DestRange) && !dcl.StringCanonicalize(desired.DestRange, actual.DestRange) {
		c.Config.Logger.Infof("Detected diff in DestRange.\nDESIRED: %v\nACTUAL: %v", desired.DestRange, actual.DestRange)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "DestRange",
		})
	}
	if !reflect.DeepEqual(desired.Priority, actual.Priority) {
		c.Config.Logger.Infof("Detected diff in Priority.\nDESIRED: %v\nACTUAL: %v", desired.Priority, actual.Priority)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "Priority",
		})
	}
	if !dcl.IsZeroValue(desired.NextHopInstance) && !dcl.StringCanonicalize(desired.NextHopInstance, actual.NextHopInstance) {
		c.Config.Logger.Infof("Detected diff in NextHopInstance.\nDESIRED: %v\nACTUAL: %v", desired.NextHopInstance, actual.NextHopInstance)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "NextHopInstance",
		})
	}
	if !dcl.IsZeroValue(desired.NextHopIP) && !dcl.StringCanonicalize(desired.NextHopIP, actual.NextHopIP) {
		c.Config.Logger.Infof("Detected diff in NextHopIP.\nDESIRED: %v\nACTUAL: %v", desired.NextHopIP, actual.NextHopIP)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "NextHopIP",
		})
	}
	if !dcl.IsZeroValue(desired.NextHopNetwork) && !dcl.StringCanonicalize(desired.NextHopNetwork, actual.NextHopNetwork) {
		c.Config.Logger.Infof("Detected diff in NextHopNetwork.\nDESIRED: %v\nACTUAL: %v", desired.NextHopNetwork, actual.NextHopNetwork)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "NextHopNetwork",
		})
	}
	if !dcl.IsZeroValue(desired.NextHopGateway) && !dcl.PartialSelfLinkToSelfLink(desired.NextHopGateway, actual.NextHopGateway) {
		c.Config.Logger.Infof("Detected diff in NextHopGateway.\nDESIRED: %v\nACTUAL: %v", desired.NextHopGateway, actual.NextHopGateway)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "NextHopGateway",
		})
	}
	if !dcl.IsZeroValue(desired.NextHopIlb) && !dcl.StringCanonicalize(desired.NextHopIlb, actual.NextHopIlb) {
		c.Config.Logger.Infof("Detected diff in NextHopIlb.\nDESIRED: %v\nACTUAL: %v", desired.NextHopIlb, actual.NextHopIlb)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "NextHopIlb",
		})
	}
	if !dcl.IsZeroValue(desired.NextHopVpnTunnel) && !dcl.StringCanonicalize(desired.NextHopVpnTunnel, actual.NextHopVpnTunnel) {
		c.Config.Logger.Infof("Detected diff in NextHopVpnTunnel.\nDESIRED: %v\nACTUAL: %v", desired.NextHopVpnTunnel, actual.NextHopVpnTunnel)
		diffs = append(diffs, routeDiff{
			RequiresRecreate: true,
			FieldName:        "NextHopVpnTunnel",
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
	var deduped []routeDiff
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
func compareRouteWarning(c *Client, desired, actual *RouteWarning) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareRouteWarningSlice(c *Client, desired, actual []RouteWarning) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouteWarning, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouteWarning(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouteWarning, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouteWarningMap(c *Client, desired, actual map[string]RouteWarning) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouteWarning, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in RouteWarning, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareRouteWarning(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in RouteWarning, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareRouteWarningCodeEnumSlice(c *Client, desired, actual []RouteWarningCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in RouteWarningCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareRouteWarningCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in RouteWarningCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareRouteWarningCodeEnum(c *Client, desired, actual *RouteWarningCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Route) urlNormalized() *Route {
	normalized := deepcopy.Copy(*r).(Route)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.DestRange = dcl.SelfLinkToName(r.DestRange)
	normalized.NextHopInstance = dcl.SelfLinkToName(r.NextHopInstance)
	normalized.NextHopIP = dcl.SelfLinkToName(r.NextHopIP)
	normalized.NextHopNetwork = dcl.SelfLinkToName(r.NextHopNetwork)
	normalized.NextHopGateway = dcl.SelfLinkToName(r.NextHopGateway)
	normalized.NextHopPeering = dcl.SelfLinkToName(r.NextHopPeering)
	normalized.NextHopIlb = dcl.SelfLinkToName(r.NextHopIlb)
	normalized.NextHopVpnTunnel = dcl.SelfLinkToName(r.NextHopVpnTunnel)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Route) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Route) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Route) deleteFields() (string, string) {
	n := r.urlNormalized()
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
	if v, err := dcl.DeriveField("projects/%s/global/networks/%s", f.Network, f.Project, f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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
	} else if !dcl.IsEmptyValueIndirect(v) {
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

	r := &Route{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Network = dcl.FlattenString(m["network"])
	r.Tag = dcl.FlattenStringSlice(m["tags"])
	r.DestRange = dcl.FlattenString(m["destRange"])
	r.Priority = dcl.FlattenInteger(m["priority"])
	if _, ok := m["priority"]; !ok {
		c.Config.Logger.Info("Using default value for priority")
		r.Priority = dcl.Int64(1000)
	}
	r.NextHopInstance = dcl.FlattenString(m["nextHopInstance"])
	r.NextHopIP = dcl.FlattenString(m["nextHopIp"])
	r.NextHopNetwork = dcl.FlattenString(m["nextHopNetwork"])
	r.NextHopGateway = dcl.FlattenString(m["nextHopGateway"])
	r.NextHopPeering = dcl.FlattenString(m["nextHopPeering"])
	r.NextHopIlb = dcl.FlattenString(m["nextHopIlb"])
	r.Warning = flattenRouteWarningSlice(c, m["warnings"])
	r.NextHopVpnTunnel = dcl.FlattenString(m["nextHopVpnTunnel"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.Project = dcl.FlattenString(m["project"])

	return r
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
