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

func (r *NetworkEndpointGroup) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"CloudRun", "AppEngine", "CloudFunction"}, r.CloudRun, r.AppEngine, r.CloudFunction); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.CloudRun) {
		if err := r.CloudRun.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AppEngine) {
		if err := r.AppEngine.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudFunction) {
		if err := r.CloudFunction.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NetworkEndpointGroupCloudRun) validate() error {
	return nil
}
func (r *NetworkEndpointGroupAppEngine) validate() error {
	return nil
}
func (r *NetworkEndpointGroupCloudFunction) validate() error {
	return nil
}

func networkEndpointGroupGetURL(userBasePath string, r *NetworkEndpointGroup) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Get URL found")

}

func networkEndpointGroupListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid List URL found")

}

func networkEndpointGroupCreateURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	if dcl.IsRegion(&location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(&location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Create URL found")

}

func networkEndpointGroupDeleteURL(userBasePath string, r *NetworkEndpointGroup) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	if dcl.IsRegion(r.Location) {
		return dcl.URL("projects/{{project}}/regions/{{location}}/networkEndpointGroups/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	if dcl.IsZone(r.Location) {
		return dcl.URL("projects/{{project}}/zones/{{location}}/networkEndpointGroups/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
	}

	return "", fmt.Errorf("No valid Delete URL found")

}

// networkEndpointGroupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type networkEndpointGroupApiOperation interface {
	do(context.Context, *NetworkEndpointGroup, *Client) error
}

func (c *Client) listNetworkEndpointGroupRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := networkEndpointGroupListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NetworkEndpointGroupMaxPage {
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

type listNetworkEndpointGroupOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNetworkEndpointGroup(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*NetworkEndpointGroup, string, error) {
	b, err := c.listNetworkEndpointGroupRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNetworkEndpointGroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*NetworkEndpointGroup
	for _, v := range m.Items {
		res := flattenNetworkEndpointGroup(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNetworkEndpointGroup(ctx context.Context, f func(*NetworkEndpointGroup) bool, resources []*NetworkEndpointGroup) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNetworkEndpointGroup(ctx, res)
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

type deleteNetworkEndpointGroupOperation struct{}

func (op *deleteNetworkEndpointGroupOperation) do(ctx context.Context, r *NetworkEndpointGroup, c *Client) error {

	_, err := c.GetNetworkEndpointGroup(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("NetworkEndpointGroup not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetNetworkEndpointGroup checking for existence. error: %v", err)
		return err
	}

	u, err := networkEndpointGroupDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetNetworkEndpointGroup(ctx, r.urlNormalized())
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
type createNetworkEndpointGroupOperation struct {
	response map[string]interface{}
}

func (op *createNetworkEndpointGroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNetworkEndpointGroupOperation) do(ctx context.Context, r *NetworkEndpointGroup, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location := r.createFields()
	u, err := networkEndpointGroupCreateURL(c.Config.BasePath, project, location)

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

	if _, err := c.GetNetworkEndpointGroup(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNetworkEndpointGroupRaw(ctx context.Context, r *NetworkEndpointGroup) ([]byte, error) {

	u, err := networkEndpointGroupGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) networkEndpointGroupDiffsForRawDesired(ctx context.Context, rawDesired *NetworkEndpointGroup, opts ...dcl.ApplyOption) (initial, desired *NetworkEndpointGroup, diffs []networkEndpointGroupDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *NetworkEndpointGroup
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*NetworkEndpointGroup); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected NetworkEndpointGroup, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNetworkEndpointGroup(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a NetworkEndpointGroup resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve NetworkEndpointGroup resource: %v", err)
		}
		c.Config.Logger.Info("Found that NetworkEndpointGroup resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNetworkEndpointGroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for NetworkEndpointGroup: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for NetworkEndpointGroup: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNetworkEndpointGroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for NetworkEndpointGroup: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNetworkEndpointGroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for NetworkEndpointGroup: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNetworkEndpointGroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNetworkEndpointGroupInitialState(rawInitial, rawDesired *NetworkEndpointGroup) (*NetworkEndpointGroup, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.CloudRun) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.AppEngine, rawInitial.CloudFunction) {
			rawInitial.CloudRun = EmptyNetworkEndpointGroupCloudRun
		}
	}

	if dcl.IsZeroValue(rawInitial.AppEngine) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.CloudRun, rawInitial.CloudFunction) {
			rawInitial.AppEngine = EmptyNetworkEndpointGroupAppEngine
		}
	}

	if dcl.IsZeroValue(rawInitial.CloudFunction) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.CloudRun, rawInitial.AppEngine) {
			rawInitial.CloudFunction = EmptyNetworkEndpointGroupCloudFunction
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNetworkEndpointGroupDesiredState(rawDesired, rawInitial *NetworkEndpointGroup, opts ...dcl.ApplyOption) (*NetworkEndpointGroup, error) {

	if dcl.IsZeroValue(rawDesired.CloudRun) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.AppEngine, rawDesired.CloudFunction) {
			rawDesired.CloudRun = EmptyNetworkEndpointGroupCloudRun
		}
	}

	if dcl.IsZeroValue(rawDesired.AppEngine) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.CloudRun, rawDesired.CloudFunction) {
			rawDesired.AppEngine = EmptyNetworkEndpointGroupAppEngine
		}
	}

	if dcl.IsZeroValue(rawDesired.CloudFunction) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.CloudRun, rawDesired.AppEngine) {
			rawDesired.CloudFunction = EmptyNetworkEndpointGroupCloudFunction
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.CloudRun = canonicalizeNetworkEndpointGroupCloudRun(rawDesired.CloudRun, nil, opts...)
		rawDesired.AppEngine = canonicalizeNetworkEndpointGroupAppEngine(rawDesired.AppEngine, nil, opts...)
		rawDesired.CloudFunction = canonicalizeNetworkEndpointGroupCloudFunction(rawDesired.CloudFunction, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Id) {
		rawDesired.Id = rawInitial.Id
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.StringCanonicalize(rawDesired.SelfLinkWithId, rawInitial.SelfLinkWithId) {
		rawDesired.SelfLinkWithId = rawInitial.SelfLinkWithId
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.IsZeroValue(rawDesired.NetworkEndpointType) {
		rawDesired.NetworkEndpointType = rawInitial.NetworkEndpointType
	}
	if dcl.IsZeroValue(rawDesired.Size) {
		rawDesired.Size = rawInitial.Size
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.NameToSelfLink(rawDesired.Subnetwork, rawInitial.Subnetwork) {
		rawDesired.Subnetwork = rawInitial.Subnetwork
	}
	if dcl.IsZeroValue(rawDesired.DefaultPort) {
		rawDesired.DefaultPort = rawInitial.DefaultPort
	}
	if dcl.IsZeroValue(rawDesired.Annotations) {
		rawDesired.Annotations = rawInitial.Annotations
	}
	rawDesired.CloudRun = canonicalizeNetworkEndpointGroupCloudRun(rawDesired.CloudRun, rawInitial.CloudRun, opts...)
	rawDesired.AppEngine = canonicalizeNetworkEndpointGroupAppEngine(rawDesired.AppEngine, rawInitial.AppEngine, opts...)
	rawDesired.CloudFunction = canonicalizeNetworkEndpointGroupCloudFunction(rawDesired.CloudFunction, rawInitial.CloudFunction, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeNetworkEndpointGroupNewState(c *Client, rawNew, rawDesired *NetworkEndpointGroup) (*NetworkEndpointGroup, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLinkWithId) && dcl.IsEmptyValueIndirect(rawDesired.SelfLinkWithId) {
		rawNew.SelfLinkWithId = rawDesired.SelfLinkWithId
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLinkWithId, rawNew.SelfLinkWithId) {
			rawNew.SelfLinkWithId = rawDesired.SelfLinkWithId
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.NetworkEndpointType) && dcl.IsEmptyValueIndirect(rawDesired.NetworkEndpointType) {
		rawNew.NetworkEndpointType = rawDesired.NetworkEndpointType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Size) && dcl.IsEmptyValueIndirect(rawDesired.Size) {
		rawNew.Size = rawDesired.Size
	} else {
	}

	rawNew.Location = rawDesired.Location

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnetwork) && dcl.IsEmptyValueIndirect(rawDesired.Subnetwork) {
		rawNew.Subnetwork = rawDesired.Subnetwork
	} else {
		if dcl.NameToSelfLink(rawDesired.Subnetwork, rawNew.Subnetwork) {
			rawNew.Subnetwork = rawDesired.Subnetwork
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultPort) && dcl.IsEmptyValueIndirect(rawDesired.DefaultPort) {
		rawNew.DefaultPort = rawDesired.DefaultPort
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CloudRun) && dcl.IsEmptyValueIndirect(rawDesired.CloudRun) {
		rawNew.CloudRun = rawDesired.CloudRun
	} else {
		rawNew.CloudRun = canonicalizeNewNetworkEndpointGroupCloudRun(c, rawDesired.CloudRun, rawNew.CloudRun)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AppEngine) && dcl.IsEmptyValueIndirect(rawDesired.AppEngine) {
		rawNew.AppEngine = rawDesired.AppEngine
	} else {
		rawNew.AppEngine = canonicalizeNewNetworkEndpointGroupAppEngine(c, rawDesired.AppEngine, rawNew.AppEngine)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CloudFunction) && dcl.IsEmptyValueIndirect(rawDesired.CloudFunction) {
		rawNew.CloudFunction = rawDesired.CloudFunction
	} else {
		rawNew.CloudFunction = canonicalizeNewNetworkEndpointGroupCloudFunction(c, rawDesired.CloudFunction, rawNew.CloudFunction)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeNetworkEndpointGroupCloudRun(des, initial *NetworkEndpointGroupCloudRun, opts ...dcl.ApplyOption) *NetworkEndpointGroupCloudRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		des.Tag = initial.Tag
	}
	if dcl.StringCanonicalize(des.UrlMask, initial.UrlMask) || dcl.IsZeroValue(des.UrlMask) {
		des.UrlMask = initial.UrlMask
	}

	return des
}

func canonicalizeNewNetworkEndpointGroupCloudRun(c *Client, des, nw *NetworkEndpointGroupCloudRun) *NetworkEndpointGroupCloudRun {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.StringCanonicalize(des.UrlMask, nw.UrlMask) {
		nw.UrlMask = des.UrlMask
	}

	return nw
}

func canonicalizeNewNetworkEndpointGroupCloudRunSet(c *Client, des, nw []NetworkEndpointGroupCloudRun) []NetworkEndpointGroupCloudRun {
	if des == nil {
		return nw
	}
	var reorderedNew []NetworkEndpointGroupCloudRun
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNetworkEndpointGroupCloudRun(c, &d, &n) {
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

func canonicalizeNewNetworkEndpointGroupCloudRunSlice(c *Client, des, nw []NetworkEndpointGroupCloudRun) []NetworkEndpointGroupCloudRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NetworkEndpointGroupCloudRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNetworkEndpointGroupCloudRun(c, &d, &n))
	}

	return items
}

func canonicalizeNetworkEndpointGroupAppEngine(des, initial *NetworkEndpointGroupAppEngine, opts ...dcl.ApplyOption) *NetworkEndpointGroupAppEngine {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.NameToSelfLink(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		des.Version = initial.Version
	}
	if dcl.StringCanonicalize(des.UrlMask, initial.UrlMask) || dcl.IsZeroValue(des.UrlMask) {
		des.UrlMask = initial.UrlMask
	}

	return des
}

func canonicalizeNewNetworkEndpointGroupAppEngine(c *Client, des, nw *NetworkEndpointGroupAppEngine) *NetworkEndpointGroupAppEngine {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.NameToSelfLink(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.UrlMask, nw.UrlMask) {
		nw.UrlMask = des.UrlMask
	}

	return nw
}

func canonicalizeNewNetworkEndpointGroupAppEngineSet(c *Client, des, nw []NetworkEndpointGroupAppEngine) []NetworkEndpointGroupAppEngine {
	if des == nil {
		return nw
	}
	var reorderedNew []NetworkEndpointGroupAppEngine
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNetworkEndpointGroupAppEngine(c, &d, &n) {
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

func canonicalizeNewNetworkEndpointGroupAppEngineSlice(c *Client, des, nw []NetworkEndpointGroupAppEngine) []NetworkEndpointGroupAppEngine {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NetworkEndpointGroupAppEngine
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNetworkEndpointGroupAppEngine(c, &d, &n))
	}

	return items
}

func canonicalizeNetworkEndpointGroupCloudFunction(des, initial *NetworkEndpointGroupCloudFunction, opts ...dcl.ApplyOption) *NetworkEndpointGroupCloudFunction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.Function, initial.Function) || dcl.IsZeroValue(des.Function) {
		des.Function = initial.Function
	}
	if dcl.StringCanonicalize(des.UrlMask, initial.UrlMask) || dcl.IsZeroValue(des.UrlMask) {
		des.UrlMask = initial.UrlMask
	}

	return des
}

func canonicalizeNewNetworkEndpointGroupCloudFunction(c *Client, des, nw *NetworkEndpointGroupCloudFunction) *NetworkEndpointGroupCloudFunction {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.Function, nw.Function) {
		nw.Function = des.Function
	}
	if dcl.StringCanonicalize(des.UrlMask, nw.UrlMask) {
		nw.UrlMask = des.UrlMask
	}

	return nw
}

func canonicalizeNewNetworkEndpointGroupCloudFunctionSet(c *Client, des, nw []NetworkEndpointGroupCloudFunction) []NetworkEndpointGroupCloudFunction {
	if des == nil {
		return nw
	}
	var reorderedNew []NetworkEndpointGroupCloudFunction
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareNetworkEndpointGroupCloudFunction(c, &d, &n) {
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

func canonicalizeNewNetworkEndpointGroupCloudFunctionSlice(c *Client, des, nw []NetworkEndpointGroupCloudFunction) []NetworkEndpointGroupCloudFunction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []NetworkEndpointGroupCloudFunction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNetworkEndpointGroupCloudFunction(c, &d, &n))
	}

	return items
}

type networkEndpointGroupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         networkEndpointGroupApiOperation
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
func diffNetworkEndpointGroup(c *Client, desired, actual *NetworkEndpointGroup, opts ...dcl.ApplyOption) ([]networkEndpointGroupDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []networkEndpointGroupDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Id, actual.Id, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Id"})
	}

	if d, err := dcl.Diff(desired.SelfLink, actual.SelfLink, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "SelfLink"})
	}

	if d, err := dcl.Diff(desired.SelfLinkWithId, actual.SelfLinkWithId, &dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "SelfLinkWithId"})
	}

	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Name"})
	}

	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Description"})
	}

	if d, err := dcl.Diff(desired.Size, actual.Size, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Size"})
	}

	if d, err := dcl.Diff(desired.Network, actual.Network, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Network"})
	}

	if d, err := dcl.Diff(desired.Subnetwork, actual.Subnetwork, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "ReferenceType"}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Subnetwork"})
	}

	if d, err := dcl.Diff(desired.DefaultPort, actual.DefaultPort, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "DefaultPort"})
	}

	if d, err := dcl.Diff(desired.Annotations, actual.Annotations, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, networkEndpointGroupDiff{RequiresRecreate: true, FieldName: "Annotations"})
	}

	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !reflect.DeepEqual(desired.NetworkEndpointType, actual.NetworkEndpointType) {
		c.Config.Logger.Infof("Detected diff in NetworkEndpointType.\nDESIRED: %v\nACTUAL: %v", desired.NetworkEndpointType, actual.NetworkEndpointType)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "NetworkEndpointType",
		})
	}
	if !reflect.DeepEqual(desired.Size, actual.Size) {
		c.Config.Logger.Infof("Detected diff in Size.\nDESIRED: %v\nACTUAL: %v", desired.Size, actual.Size)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "Size",
		})
	}
	if !dcl.IsZeroValue(desired.Network) && !dcl.NameToSelfLink(desired.Network, actual.Network) {
		c.Config.Logger.Infof("Detected diff in Network.\nDESIRED: %v\nACTUAL: %v", desired.Network, actual.Network)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "Network",
		})
	}
	if !dcl.IsZeroValue(desired.Subnetwork) && !dcl.NameToSelfLink(desired.Subnetwork, actual.Subnetwork) {
		c.Config.Logger.Infof("Detected diff in Subnetwork.\nDESIRED: %v\nACTUAL: %v", desired.Subnetwork, actual.Subnetwork)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "Subnetwork",
		})
	}
	if !reflect.DeepEqual(desired.DefaultPort, actual.DefaultPort) {
		c.Config.Logger.Infof("Detected diff in DefaultPort.\nDESIRED: %v\nACTUAL: %v", desired.DefaultPort, actual.DefaultPort)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "DefaultPort",
		})
	}
	if !dcl.MapEquals(desired.Annotations, actual.Annotations, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Annotations.\nDESIRED: %v\nACTUAL: %v", desired.Annotations, actual.Annotations)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "Annotations",
		})
	}
	if compareNetworkEndpointGroupCloudRun(c, desired.CloudRun, actual.CloudRun) {
		c.Config.Logger.Infof("Detected diff in CloudRun.\nDESIRED: %v\nACTUAL: %v", desired.CloudRun, actual.CloudRun)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "CloudRun",
		})
	}
	if compareNetworkEndpointGroupAppEngine(c, desired.AppEngine, actual.AppEngine) {
		c.Config.Logger.Infof("Detected diff in AppEngine.\nDESIRED: %v\nACTUAL: %v", desired.AppEngine, actual.AppEngine)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "AppEngine",
		})
	}
	if compareNetworkEndpointGroupCloudFunction(c, desired.CloudFunction, actual.CloudFunction) {
		c.Config.Logger.Infof("Detected diff in CloudFunction.\nDESIRED: %v\nACTUAL: %v", desired.CloudFunction, actual.CloudFunction)
		diffs = append(diffs, networkEndpointGroupDiff{
			RequiresRecreate: true,
			FieldName:        "CloudFunction",
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
	var deduped []networkEndpointGroupDiff
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
func compareNetworkEndpointGroupCloudRun(c *Client, desired, actual *NetworkEndpointGroupCloudRun) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Service == nil && desired.Service != nil && !dcl.IsEmptyValueIndirect(desired.Service) {
		c.Config.Logger.Infof("desired Service %s - but actually nil", dcl.SprintResource(desired.Service))
		return true
	}
	if !dcl.StringCanonicalize(desired.Service, actual.Service) && !dcl.IsZeroValue(desired.Service) {
		c.Config.Logger.Infof("Diff in Service. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Service), dcl.SprintResource(actual.Service))
		return true
	}
	if actual.Tag == nil && desired.Tag != nil && !dcl.IsEmptyValueIndirect(desired.Tag) {
		c.Config.Logger.Infof("desired Tag %s - but actually nil", dcl.SprintResource(desired.Tag))
		return true
	}
	if !dcl.StringCanonicalize(desired.Tag, actual.Tag) && !dcl.IsZeroValue(desired.Tag) {
		c.Config.Logger.Infof("Diff in Tag. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tag), dcl.SprintResource(actual.Tag))
		return true
	}
	if actual.UrlMask == nil && desired.UrlMask != nil && !dcl.IsEmptyValueIndirect(desired.UrlMask) {
		c.Config.Logger.Infof("desired UrlMask %s - but actually nil", dcl.SprintResource(desired.UrlMask))
		return true
	}
	if !dcl.StringCanonicalize(desired.UrlMask, actual.UrlMask) && !dcl.IsZeroValue(desired.UrlMask) {
		c.Config.Logger.Infof("Diff in UrlMask. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlMask), dcl.SprintResource(actual.UrlMask))
		return true
	}
	return false
}

func compareNetworkEndpointGroupCloudRunSlice(c *Client, desired, actual []NetworkEndpointGroupCloudRun) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupCloudRun, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNetworkEndpointGroupCloudRun(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupCloudRun, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupCloudRunMap(c *Client, desired, actual map[string]NetworkEndpointGroupCloudRun) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupCloudRun, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupCloudRun, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNetworkEndpointGroupCloudRun(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupCloudRun, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupAppEngine(c *Client, desired, actual *NetworkEndpointGroupAppEngine) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Service == nil && desired.Service != nil && !dcl.IsEmptyValueIndirect(desired.Service) {
		c.Config.Logger.Infof("desired Service %s - but actually nil", dcl.SprintResource(desired.Service))
		return true
	}
	if !dcl.StringCanonicalize(desired.Service, actual.Service) && !dcl.IsZeroValue(desired.Service) {
		c.Config.Logger.Infof("Diff in Service. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Service), dcl.SprintResource(actual.Service))
		return true
	}
	if actual.Version == nil && desired.Version != nil && !dcl.IsEmptyValueIndirect(desired.Version) {
		c.Config.Logger.Infof("desired Version %s - but actually nil", dcl.SprintResource(desired.Version))
		return true
	}
	if !dcl.NameToSelfLink(desired.Version, actual.Version) && !dcl.IsZeroValue(desired.Version) {
		c.Config.Logger.Infof("Diff in Version. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Version), dcl.SprintResource(actual.Version))
		return true
	}
	if actual.UrlMask == nil && desired.UrlMask != nil && !dcl.IsEmptyValueIndirect(desired.UrlMask) {
		c.Config.Logger.Infof("desired UrlMask %s - but actually nil", dcl.SprintResource(desired.UrlMask))
		return true
	}
	if !dcl.StringCanonicalize(desired.UrlMask, actual.UrlMask) && !dcl.IsZeroValue(desired.UrlMask) {
		c.Config.Logger.Infof("Diff in UrlMask. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlMask), dcl.SprintResource(actual.UrlMask))
		return true
	}
	return false
}

func compareNetworkEndpointGroupAppEngineSlice(c *Client, desired, actual []NetworkEndpointGroupAppEngine) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupAppEngine, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNetworkEndpointGroupAppEngine(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupAppEngine, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupAppEngineMap(c *Client, desired, actual map[string]NetworkEndpointGroupAppEngine) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupAppEngine, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupAppEngine, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNetworkEndpointGroupAppEngine(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupAppEngine, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupCloudFunction(c *Client, desired, actual *NetworkEndpointGroupCloudFunction) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Function == nil && desired.Function != nil && !dcl.IsEmptyValueIndirect(desired.Function) {
		c.Config.Logger.Infof("desired Function %s - but actually nil", dcl.SprintResource(desired.Function))
		return true
	}
	if !dcl.NameToSelfLink(desired.Function, actual.Function) && !dcl.IsZeroValue(desired.Function) {
		c.Config.Logger.Infof("Diff in Function. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Function), dcl.SprintResource(actual.Function))
		return true
	}
	if actual.UrlMask == nil && desired.UrlMask != nil && !dcl.IsEmptyValueIndirect(desired.UrlMask) {
		c.Config.Logger.Infof("desired UrlMask %s - but actually nil", dcl.SprintResource(desired.UrlMask))
		return true
	}
	if !dcl.StringCanonicalize(desired.UrlMask, actual.UrlMask) && !dcl.IsZeroValue(desired.UrlMask) {
		c.Config.Logger.Infof("Diff in UrlMask. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UrlMask), dcl.SprintResource(actual.UrlMask))
		return true
	}
	return false
}

func compareNetworkEndpointGroupCloudFunctionSlice(c *Client, desired, actual []NetworkEndpointGroupCloudFunction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupCloudFunction, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNetworkEndpointGroupCloudFunction(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupCloudFunction, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupCloudFunctionMap(c *Client, desired, actual map[string]NetworkEndpointGroupCloudFunction) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupCloudFunction, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupCloudFunction, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareNetworkEndpointGroupCloudFunction(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupCloudFunction, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupNetworkEndpointTypeEnumSlice(c *Client, desired, actual []NetworkEndpointGroupNetworkEndpointTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in NetworkEndpointGroupNetworkEndpointTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareNetworkEndpointGroupNetworkEndpointTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in NetworkEndpointGroupNetworkEndpointTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareNetworkEndpointGroupNetworkEndpointTypeEnum(c *Client, desired, actual *NetworkEndpointGroupNetworkEndpointTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *NetworkEndpointGroup) urlNormalized() *NetworkEndpointGroup {
	normalized := deepcopy.Copy(*r).(NetworkEndpointGroup)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.SelfLinkWithId = dcl.SelfLinkToName(r.SelfLinkWithId)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.Subnetwork = dcl.SelfLinkToName(r.Subnetwork)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *NetworkEndpointGroup) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *NetworkEndpointGroup) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location)
}

func (r *NetworkEndpointGroup) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *NetworkEndpointGroup) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the NetworkEndpointGroup resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *NetworkEndpointGroup) marshal(c *Client) ([]byte, error) {
	m, err := expandNetworkEndpointGroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling NetworkEndpointGroup: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNetworkEndpointGroup decodes JSON responses into the NetworkEndpointGroup resource schema.
func unmarshalNetworkEndpointGroup(b []byte, c *Client) (*NetworkEndpointGroup, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNetworkEndpointGroup(m, c)
}

func unmarshalMapNetworkEndpointGroup(m map[string]interface{}, c *Client) (*NetworkEndpointGroup, error) {

	return flattenNetworkEndpointGroup(c, m), nil
}

// expandNetworkEndpointGroup expands NetworkEndpointGroup into a JSON request object.
func expandNetworkEndpointGroup(c *Client, f *NetworkEndpointGroup) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.SelfLinkWithId; !dcl.IsEmptyValueIndirect(v) {
		m["selfLinkWithId"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.NetworkEndpointType; !dcl.IsEmptyValueIndirect(v) {
		m["networkEndpointType"] = v
	}
	if v := f.Size; !dcl.IsEmptyValueIndirect(v) {
		m["size"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Subnetwork; !dcl.IsEmptyValueIndirect(v) {
		m["subnetwork"] = v
	}
	if v := f.DefaultPort; !dcl.IsEmptyValueIndirect(v) {
		m["defaultPort"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandNetworkEndpointGroupCloudRun(c, f.CloudRun); err != nil {
		return nil, fmt.Errorf("error expanding CloudRun into cloudRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudRun"] = v
	}
	if v, err := expandNetworkEndpointGroupAppEngine(c, f.AppEngine); err != nil {
		return nil, fmt.Errorf("error expanding AppEngine into appEngine: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["appEngine"] = v
	}
	if v, err := expandNetworkEndpointGroupCloudFunction(c, f.CloudFunction); err != nil {
		return nil, fmt.Errorf("error expanding CloudFunction into cloudFunction: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudFunction"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenNetworkEndpointGroup flattens NetworkEndpointGroup from a JSON request object into the
// NetworkEndpointGroup type.
func flattenNetworkEndpointGroup(c *Client, i interface{}) *NetworkEndpointGroup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &NetworkEndpointGroup{}
	r.Id = dcl.FlattenInteger(m["id"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.SelfLinkWithId = dcl.FlattenString(m["selfLinkWithId"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.NetworkEndpointType = flattenNetworkEndpointGroupNetworkEndpointTypeEnum(m["networkEndpointType"])
	r.Size = dcl.FlattenInteger(m["size"])
	r.Location = dcl.FlattenString(m["location"])
	r.Network = dcl.FlattenString(m["network"])
	r.Subnetwork = dcl.FlattenString(m["subnetwork"])
	r.DefaultPort = dcl.FlattenInteger(m["defaultPort"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.CloudRun = flattenNetworkEndpointGroupCloudRun(c, m["cloudRun"])
	r.AppEngine = flattenNetworkEndpointGroupAppEngine(c, m["appEngine"])
	r.CloudFunction = flattenNetworkEndpointGroupCloudFunction(c, m["cloudFunction"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandNetworkEndpointGroupCloudRunMap expands the contents of NetworkEndpointGroupCloudRun into a JSON
// request object.
func expandNetworkEndpointGroupCloudRunMap(c *Client, f map[string]NetworkEndpointGroupCloudRun) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNetworkEndpointGroupCloudRun(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNetworkEndpointGroupCloudRunSlice expands the contents of NetworkEndpointGroupCloudRun into a JSON
// request object.
func expandNetworkEndpointGroupCloudRunSlice(c *Client, f []NetworkEndpointGroupCloudRun) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNetworkEndpointGroupCloudRun(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNetworkEndpointGroupCloudRunMap flattens the contents of NetworkEndpointGroupCloudRun from a JSON
// response object.
func flattenNetworkEndpointGroupCloudRunMap(c *Client, i interface{}) map[string]NetworkEndpointGroupCloudRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NetworkEndpointGroupCloudRun{}
	}

	if len(a) == 0 {
		return map[string]NetworkEndpointGroupCloudRun{}
	}

	items := make(map[string]NetworkEndpointGroupCloudRun)
	for k, item := range a {
		items[k] = *flattenNetworkEndpointGroupCloudRun(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNetworkEndpointGroupCloudRunSlice flattens the contents of NetworkEndpointGroupCloudRun from a JSON
// response object.
func flattenNetworkEndpointGroupCloudRunSlice(c *Client, i interface{}) []NetworkEndpointGroupCloudRun {
	a, ok := i.([]interface{})
	if !ok {
		return []NetworkEndpointGroupCloudRun{}
	}

	if len(a) == 0 {
		return []NetworkEndpointGroupCloudRun{}
	}

	items := make([]NetworkEndpointGroupCloudRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNetworkEndpointGroupCloudRun(c, item.(map[string]interface{})))
	}

	return items
}

// expandNetworkEndpointGroupCloudRun expands an instance of NetworkEndpointGroupCloudRun into a JSON
// request object.
func expandNetworkEndpointGroupCloudRun(c *Client, f *NetworkEndpointGroupCloudRun) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}
	if v := f.UrlMask; !dcl.IsEmptyValueIndirect(v) {
		m["urlMask"] = v
	}

	return m, nil
}

// flattenNetworkEndpointGroupCloudRun flattens an instance of NetworkEndpointGroupCloudRun from a JSON
// response object.
func flattenNetworkEndpointGroupCloudRun(c *Client, i interface{}) *NetworkEndpointGroupCloudRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NetworkEndpointGroupCloudRun{}
	r.Service = dcl.FlattenString(m["service"])
	r.Tag = dcl.FlattenString(m["tag"])
	r.UrlMask = dcl.FlattenString(m["urlMask"])

	return r
}

// expandNetworkEndpointGroupAppEngineMap expands the contents of NetworkEndpointGroupAppEngine into a JSON
// request object.
func expandNetworkEndpointGroupAppEngineMap(c *Client, f map[string]NetworkEndpointGroupAppEngine) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNetworkEndpointGroupAppEngine(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNetworkEndpointGroupAppEngineSlice expands the contents of NetworkEndpointGroupAppEngine into a JSON
// request object.
func expandNetworkEndpointGroupAppEngineSlice(c *Client, f []NetworkEndpointGroupAppEngine) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNetworkEndpointGroupAppEngine(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNetworkEndpointGroupAppEngineMap flattens the contents of NetworkEndpointGroupAppEngine from a JSON
// response object.
func flattenNetworkEndpointGroupAppEngineMap(c *Client, i interface{}) map[string]NetworkEndpointGroupAppEngine {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NetworkEndpointGroupAppEngine{}
	}

	if len(a) == 0 {
		return map[string]NetworkEndpointGroupAppEngine{}
	}

	items := make(map[string]NetworkEndpointGroupAppEngine)
	for k, item := range a {
		items[k] = *flattenNetworkEndpointGroupAppEngine(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNetworkEndpointGroupAppEngineSlice flattens the contents of NetworkEndpointGroupAppEngine from a JSON
// response object.
func flattenNetworkEndpointGroupAppEngineSlice(c *Client, i interface{}) []NetworkEndpointGroupAppEngine {
	a, ok := i.([]interface{})
	if !ok {
		return []NetworkEndpointGroupAppEngine{}
	}

	if len(a) == 0 {
		return []NetworkEndpointGroupAppEngine{}
	}

	items := make([]NetworkEndpointGroupAppEngine, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNetworkEndpointGroupAppEngine(c, item.(map[string]interface{})))
	}

	return items
}

// expandNetworkEndpointGroupAppEngine expands an instance of NetworkEndpointGroupAppEngine into a JSON
// request object.
func expandNetworkEndpointGroupAppEngine(c *Client, f *NetworkEndpointGroupAppEngine) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.UrlMask; !dcl.IsEmptyValueIndirect(v) {
		m["urlMask"] = v
	}

	return m, nil
}

// flattenNetworkEndpointGroupAppEngine flattens an instance of NetworkEndpointGroupAppEngine from a JSON
// response object.
func flattenNetworkEndpointGroupAppEngine(c *Client, i interface{}) *NetworkEndpointGroupAppEngine {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NetworkEndpointGroupAppEngine{}
	r.Service = dcl.FlattenString(m["service"])
	r.Version = dcl.FlattenString(m["version"])
	r.UrlMask = dcl.FlattenString(m["urlMask"])

	return r
}

// expandNetworkEndpointGroupCloudFunctionMap expands the contents of NetworkEndpointGroupCloudFunction into a JSON
// request object.
func expandNetworkEndpointGroupCloudFunctionMap(c *Client, f map[string]NetworkEndpointGroupCloudFunction) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNetworkEndpointGroupCloudFunction(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNetworkEndpointGroupCloudFunctionSlice expands the contents of NetworkEndpointGroupCloudFunction into a JSON
// request object.
func expandNetworkEndpointGroupCloudFunctionSlice(c *Client, f []NetworkEndpointGroupCloudFunction) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNetworkEndpointGroupCloudFunction(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNetworkEndpointGroupCloudFunctionMap flattens the contents of NetworkEndpointGroupCloudFunction from a JSON
// response object.
func flattenNetworkEndpointGroupCloudFunctionMap(c *Client, i interface{}) map[string]NetworkEndpointGroupCloudFunction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NetworkEndpointGroupCloudFunction{}
	}

	if len(a) == 0 {
		return map[string]NetworkEndpointGroupCloudFunction{}
	}

	items := make(map[string]NetworkEndpointGroupCloudFunction)
	for k, item := range a {
		items[k] = *flattenNetworkEndpointGroupCloudFunction(c, item.(map[string]interface{}))
	}

	return items
}

// flattenNetworkEndpointGroupCloudFunctionSlice flattens the contents of NetworkEndpointGroupCloudFunction from a JSON
// response object.
func flattenNetworkEndpointGroupCloudFunctionSlice(c *Client, i interface{}) []NetworkEndpointGroupCloudFunction {
	a, ok := i.([]interface{})
	if !ok {
		return []NetworkEndpointGroupCloudFunction{}
	}

	if len(a) == 0 {
		return []NetworkEndpointGroupCloudFunction{}
	}

	items := make([]NetworkEndpointGroupCloudFunction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNetworkEndpointGroupCloudFunction(c, item.(map[string]interface{})))
	}

	return items
}

// expandNetworkEndpointGroupCloudFunction expands an instance of NetworkEndpointGroupCloudFunction into a JSON
// request object.
func expandNetworkEndpointGroupCloudFunction(c *Client, f *NetworkEndpointGroupCloudFunction) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Function; !dcl.IsEmptyValueIndirect(v) {
		m["function"] = v
	}
	if v := f.UrlMask; !dcl.IsEmptyValueIndirect(v) {
		m["urlMask"] = v
	}

	return m, nil
}

// flattenNetworkEndpointGroupCloudFunction flattens an instance of NetworkEndpointGroupCloudFunction from a JSON
// response object.
func flattenNetworkEndpointGroupCloudFunction(c *Client, i interface{}) *NetworkEndpointGroupCloudFunction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NetworkEndpointGroupCloudFunction{}
	r.Function = dcl.FlattenString(m["function"])
	r.UrlMask = dcl.FlattenString(m["urlMask"])

	return r
}

// flattenNetworkEndpointGroupNetworkEndpointTypeEnumSlice flattens the contents of NetworkEndpointGroupNetworkEndpointTypeEnum from a JSON
// response object.
func flattenNetworkEndpointGroupNetworkEndpointTypeEnumSlice(c *Client, i interface{}) []NetworkEndpointGroupNetworkEndpointTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NetworkEndpointGroupNetworkEndpointTypeEnum{}
	}

	if len(a) == 0 {
		return []NetworkEndpointGroupNetworkEndpointTypeEnum{}
	}

	items := make([]NetworkEndpointGroupNetworkEndpointTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNetworkEndpointGroupNetworkEndpointTypeEnum(item.(interface{})))
	}

	return items
}

// flattenNetworkEndpointGroupNetworkEndpointTypeEnum asserts that an interface is a string, and returns a
// pointer to a *NetworkEndpointGroupNetworkEndpointTypeEnum with the same value as that string.
func flattenNetworkEndpointGroupNetworkEndpointTypeEnum(i interface{}) *NetworkEndpointGroupNetworkEndpointTypeEnum {
	s, ok := i.(string)
	if !ok {
		return NetworkEndpointGroupNetworkEndpointTypeEnumRef("")
	}

	return NetworkEndpointGroupNetworkEndpointTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *NetworkEndpointGroup) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNetworkEndpointGroup(b, c)
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
