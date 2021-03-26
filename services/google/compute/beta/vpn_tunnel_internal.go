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

func (r *VpnTunnel) validate() error {

	if err := dcl.Required(r, "sharedSecret"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}

func vpnTunnelGetURL(userBasePath string, r *VpnTunnel) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func vpnTunnelListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnTunnels", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func vpnTunnelCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnTunnels", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func vpnTunnelDeleteURL(userBasePath string, r *VpnTunnel) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// vpnTunnelApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type vpnTunnelApiOperation interface {
	do(context.Context, *VpnTunnel, *Client) error
}

// newUpdateVpnTunnelSetLabelsRequest creates a request for an
// VpnTunnel resource's setLabels update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateVpnTunnelSetLabelsRequest(ctx context.Context, f *VpnTunnel, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateVpnTunnelSetLabelsRequest converts the update into
// the final JSON request body.
func marshalUpdateVpnTunnelSetLabelsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateVpnTunnelSetLabelsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateVpnTunnelSetLabelsOperation) do(ctx context.Context, r *VpnTunnel, c *Client) error {
	_, err := c.GetVpnTunnel(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setLabels")
	if err != nil {
		return err
	}

	req, err := newUpdateVpnTunnelSetLabelsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateVpnTunnelSetLabelsRequest(c, req)
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

func (c *Client) listVpnTunnelRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := vpnTunnelListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != VpnTunnelMaxPage {
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

type listVpnTunnelOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listVpnTunnel(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*VpnTunnel, string, error) {
	b, err := c.listVpnTunnelRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listVpnTunnelOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*VpnTunnel
	for _, v := range m.Items {
		res := flattenVpnTunnel(c, v)
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllVpnTunnel(ctx context.Context, f func(*VpnTunnel) bool, resources []*VpnTunnel) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteVpnTunnel(ctx, res)
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

type deleteVpnTunnelOperation struct{}

func (op *deleteVpnTunnelOperation) do(ctx context.Context, r *VpnTunnel, c *Client) error {

	_, err := c.GetVpnTunnel(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("VpnTunnel not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetVpnTunnel checking for existence. error: %v", err)
		return err
	}

	u, err := vpnTunnelDeleteURL(c.Config.BasePath, r.urlNormalized())
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
		_, err = c.GetVpnTunnel(ctx, r.urlNormalized())
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
type createVpnTunnelOperation struct {
	response map[string]interface{}
}

func (op *createVpnTunnelOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createVpnTunnelOperation) do(ctx context.Context, r *VpnTunnel, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := vpnTunnelCreateURL(c.Config.BasePath, project, region)

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

	if _, err := c.GetVpnTunnel(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getVpnTunnelRaw(ctx context.Context, r *VpnTunnel) ([]byte, error) {
	if dcl.IsZeroValue(r.IkeVersion) {
		r.IkeVersion = dcl.Int64(2)
	}

	u, err := vpnTunnelGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) vpnTunnelDiffsForRawDesired(ctx context.Context, rawDesired *VpnTunnel, opts ...dcl.ApplyOption) (initial, desired *VpnTunnel, diffs []vpnTunnelDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *VpnTunnel
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*VpnTunnel); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected VpnTunnel, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetVpnTunnel(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a VpnTunnel resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve VpnTunnel resource: %v", err)
		}
		c.Config.Logger.Info("Found that VpnTunnel resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeVpnTunnelDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for VpnTunnel: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for VpnTunnel: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeVpnTunnelInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for VpnTunnel: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeVpnTunnelDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for VpnTunnel: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffVpnTunnel(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeVpnTunnelInitialState(rawInitial, rawDesired *VpnTunnel) (*VpnTunnel, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeVpnTunnelDesiredState(rawDesired, rawInitial *VpnTunnel, opts ...dcl.ApplyOption) (*VpnTunnel, error) {

	if dcl.IsZeroValue(rawDesired.IkeVersion) {
		rawDesired.IkeVersion = dcl.Int64(2)
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
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
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.TargetVpnGateway, rawInitial.TargetVpnGateway) {
		rawDesired.TargetVpnGateway = rawInitial.TargetVpnGateway
	}
	if dcl.NameToSelfLink(rawDesired.VpnGateway, rawInitial.VpnGateway) {
		rawDesired.VpnGateway = rawInitial.VpnGateway
	}
	if dcl.IsZeroValue(rawDesired.VpnGatewayInterface) {
		rawDesired.VpnGatewayInterface = rawInitial.VpnGatewayInterface
	}
	if dcl.StringCanonicalize(rawDesired.PeerExternalGateway, rawInitial.PeerExternalGateway) {
		rawDesired.PeerExternalGateway = rawInitial.PeerExternalGateway
	}
	if dcl.IsZeroValue(rawDesired.PeerExternalGatewayInterface) {
		rawDesired.PeerExternalGatewayInterface = rawInitial.PeerExternalGatewayInterface
	}
	if dcl.StringCanonicalize(rawDesired.PeerGcpGateway, rawInitial.PeerGcpGateway) {
		rawDesired.PeerGcpGateway = rawInitial.PeerGcpGateway
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Router, rawInitial.Router) {
		rawDesired.Router = rawInitial.Router
	}
	if dcl.StringCanonicalize(rawDesired.PeerIP, rawInitial.PeerIP) {
		rawDesired.PeerIP = rawInitial.PeerIP
	}
	if dcl.StringCanonicalize(rawDesired.SharedSecret, rawInitial.SharedSecret) {
		rawDesired.SharedSecret = rawInitial.SharedSecret
	}
	if dcl.StringCanonicalize(rawDesired.SharedSecretHash, rawInitial.SharedSecretHash) {
		rawDesired.SharedSecretHash = rawInitial.SharedSecretHash
	}
	if dcl.IsZeroValue(rawDesired.Status) {
		rawDesired.Status = rawInitial.Status
	}
	if dcl.StringCanonicalize(rawDesired.SelfLink, rawInitial.SelfLink) {
		rawDesired.SelfLink = rawInitial.SelfLink
	}
	if dcl.IsZeroValue(rawDesired.IkeVersion) {
		rawDesired.IkeVersion = rawInitial.IkeVersion
	}
	if dcl.StringCanonicalize(rawDesired.DetailedStatus, rawInitial.DetailedStatus) {
		rawDesired.DetailedStatus = rawInitial.DetailedStatus
	}
	if dcl.IsZeroValue(rawDesired.LocalTrafficSelector) {
		rawDesired.LocalTrafficSelector = rawInitial.LocalTrafficSelector
	}
	if dcl.IsZeroValue(rawDesired.RemoteTrafficSelector) {
		rawDesired.RemoteTrafficSelector = rawInitial.RemoteTrafficSelector
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeVpnTunnelNewState(c *Client, rawNew, rawDesired *VpnTunnel) (*VpnTunnel, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

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

	if dcl.IsEmptyValueIndirect(rawNew.TargetVpnGateway) && dcl.IsEmptyValueIndirect(rawDesired.TargetVpnGateway) {
		rawNew.TargetVpnGateway = rawDesired.TargetVpnGateway
	} else {
		if dcl.NameToSelfLink(rawDesired.TargetVpnGateway, rawNew.TargetVpnGateway) {
			rawNew.TargetVpnGateway = rawDesired.TargetVpnGateway
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VpnGateway) && dcl.IsEmptyValueIndirect(rawDesired.VpnGateway) {
		rawNew.VpnGateway = rawDesired.VpnGateway
	} else {
		if dcl.NameToSelfLink(rawDesired.VpnGateway, rawNew.VpnGateway) {
			rawNew.VpnGateway = rawDesired.VpnGateway
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VpnGatewayInterface) && dcl.IsEmptyValueIndirect(rawDesired.VpnGatewayInterface) {
		rawNew.VpnGatewayInterface = rawDesired.VpnGatewayInterface
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerExternalGateway) && dcl.IsEmptyValueIndirect(rawDesired.PeerExternalGateway) {
		rawNew.PeerExternalGateway = rawDesired.PeerExternalGateway
	} else {
		if dcl.StringCanonicalize(rawDesired.PeerExternalGateway, rawNew.PeerExternalGateway) {
			rawNew.PeerExternalGateway = rawDesired.PeerExternalGateway
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerExternalGatewayInterface) && dcl.IsEmptyValueIndirect(rawDesired.PeerExternalGatewayInterface) {
		rawNew.PeerExternalGatewayInterface = rawDesired.PeerExternalGatewayInterface
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerGcpGateway) && dcl.IsEmptyValueIndirect(rawDesired.PeerGcpGateway) {
		rawNew.PeerGcpGateway = rawDesired.PeerGcpGateway
	} else {
		if dcl.StringCanonicalize(rawDesired.PeerGcpGateway, rawNew.PeerGcpGateway) {
			rawNew.PeerGcpGateway = rawDesired.PeerGcpGateway
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Router) && dcl.IsEmptyValueIndirect(rawDesired.Router) {
		rawNew.Router = rawDesired.Router
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Router, rawNew.Router) {
			rawNew.Router = rawDesired.Router
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PeerIP) && dcl.IsEmptyValueIndirect(rawDesired.PeerIP) {
		rawNew.PeerIP = rawDesired.PeerIP
	} else {
		if dcl.StringCanonicalize(rawDesired.PeerIP, rawNew.PeerIP) {
			rawNew.PeerIP = rawDesired.PeerIP
		}
	}

	rawNew.SharedSecret = rawDesired.SharedSecret

	if dcl.IsEmptyValueIndirect(rawNew.SharedSecretHash) && dcl.IsEmptyValueIndirect(rawDesired.SharedSecretHash) {
		rawNew.SharedSecretHash = rawDesired.SharedSecretHash
	} else {
		if dcl.StringCanonicalize(rawDesired.SharedSecretHash, rawNew.SharedSecretHash) {
			rawNew.SharedSecretHash = rawDesired.SharedSecretHash
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IkeVersion) && dcl.IsEmptyValueIndirect(rawDesired.IkeVersion) {
		rawNew.IkeVersion = rawDesired.IkeVersion
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DetailedStatus) && dcl.IsEmptyValueIndirect(rawDesired.DetailedStatus) {
		rawNew.DetailedStatus = rawDesired.DetailedStatus
	} else {
		if dcl.StringCanonicalize(rawDesired.DetailedStatus, rawNew.DetailedStatus) {
			rawNew.DetailedStatus = rawDesired.DetailedStatus
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocalTrafficSelector) && dcl.IsEmptyValueIndirect(rawDesired.LocalTrafficSelector) {
		rawNew.LocalTrafficSelector = rawDesired.LocalTrafficSelector
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RemoteTrafficSelector) && dcl.IsEmptyValueIndirect(rawDesired.RemoteTrafficSelector) {
		rawNew.RemoteTrafficSelector = rawDesired.RemoteTrafficSelector
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

type vpnTunnelDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         vpnTunnelApiOperation
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
func diffVpnTunnel(c *Client, desired, actual *VpnTunnel, opts ...dcl.ApplyOption) ([]vpnTunnelDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []vpnTunnelDiff
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, vpnTunnelDiff{
			UpdateOp:  &updateVpnTunnelSetLabelsOperation{},
			FieldName: "Labels",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.StringCanonicalize(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "Description",
		})
	}
	if !dcl.IsZeroValue(desired.Region) && !dcl.StringCanonicalize(desired.Region, actual.Region) {
		c.Config.Logger.Infof("Detected diff in Region.\nDESIRED: %v\nACTUAL: %v", desired.Region, actual.Region)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "Region",
		})
	}
	if !dcl.IsZeroValue(desired.TargetVpnGateway) && !dcl.NameToSelfLink(desired.TargetVpnGateway, actual.TargetVpnGateway) {
		c.Config.Logger.Infof("Detected diff in TargetVpnGateway.\nDESIRED: %v\nACTUAL: %v", desired.TargetVpnGateway, actual.TargetVpnGateway)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "TargetVpnGateway",
		})
	}
	if !dcl.IsZeroValue(desired.VpnGateway) && !dcl.NameToSelfLink(desired.VpnGateway, actual.VpnGateway) {
		c.Config.Logger.Infof("Detected diff in VpnGateway.\nDESIRED: %v\nACTUAL: %v", desired.VpnGateway, actual.VpnGateway)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "VpnGateway",
		})
	}
	if !reflect.DeepEqual(desired.VpnGatewayInterface, actual.VpnGatewayInterface) {
		c.Config.Logger.Infof("Detected diff in VpnGatewayInterface.\nDESIRED: %v\nACTUAL: %v", desired.VpnGatewayInterface, actual.VpnGatewayInterface)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "VpnGatewayInterface",
		})
	}
	if !dcl.IsZeroValue(desired.PeerExternalGateway) && !dcl.StringCanonicalize(desired.PeerExternalGateway, actual.PeerExternalGateway) {
		c.Config.Logger.Infof("Detected diff in PeerExternalGateway.\nDESIRED: %v\nACTUAL: %v", desired.PeerExternalGateway, actual.PeerExternalGateway)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "PeerExternalGateway",
		})
	}
	if !reflect.DeepEqual(desired.PeerExternalGatewayInterface, actual.PeerExternalGatewayInterface) {
		c.Config.Logger.Infof("Detected diff in PeerExternalGatewayInterface.\nDESIRED: %v\nACTUAL: %v", desired.PeerExternalGatewayInterface, actual.PeerExternalGatewayInterface)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "PeerExternalGatewayInterface",
		})
	}
	if !dcl.IsZeroValue(desired.PeerGcpGateway) && !dcl.StringCanonicalize(desired.PeerGcpGateway, actual.PeerGcpGateway) {
		c.Config.Logger.Infof("Detected diff in PeerGcpGateway.\nDESIRED: %v\nACTUAL: %v", desired.PeerGcpGateway, actual.PeerGcpGateway)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "PeerGcpGateway",
		})
	}
	if !dcl.IsZeroValue(desired.Router) && !dcl.PartialSelfLinkToSelfLink(desired.Router, actual.Router) {
		c.Config.Logger.Infof("Detected diff in Router.\nDESIRED: %v\nACTUAL: %v", desired.Router, actual.Router)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "Router",
		})
	}
	if !dcl.IsZeroValue(desired.PeerIP) && !dcl.StringCanonicalize(desired.PeerIP, actual.PeerIP) {
		c.Config.Logger.Infof("Detected diff in PeerIP.\nDESIRED: %v\nACTUAL: %v", desired.PeerIP, actual.PeerIP)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "PeerIP",
		})
	}
	if !reflect.DeepEqual(desired.IkeVersion, actual.IkeVersion) {
		c.Config.Logger.Infof("Detected diff in IkeVersion.\nDESIRED: %v\nACTUAL: %v", desired.IkeVersion, actual.IkeVersion)
		diffs = append(diffs, vpnTunnelDiff{
			RequiresRecreate: true,
			FieldName:        "IkeVersion",
		})
	}
	if !dcl.StringSliceEquals(desired.LocalTrafficSelector, actual.LocalTrafficSelector) {
		c.Config.Logger.Infof("Detected diff in LocalTrafficSelector.\nDESIRED: %v\nACTUAL: %v", desired.LocalTrafficSelector, actual.LocalTrafficSelector)
		toAdd, toRemove := dcl.CompareStringSets(desired.LocalTrafficSelector, actual.LocalTrafficSelector)
		if len(toAdd) > 0 {
			diffs = append(diffs, vpnTunnelDiff{
				RequiresRecreate: true,
				FieldName:        "LocalTrafficSelector",
			})
		}
		if len(toRemove) > 0 {
			diffs = append(diffs, vpnTunnelDiff{
				RequiresRecreate: true,
				FieldName:        "LocalTrafficSelector",
			})
		}
	}
	if !dcl.StringSliceEquals(desired.RemoteTrafficSelector, actual.RemoteTrafficSelector) {
		c.Config.Logger.Infof("Detected diff in RemoteTrafficSelector.\nDESIRED: %v\nACTUAL: %v", desired.RemoteTrafficSelector, actual.RemoteTrafficSelector)
		toAdd, toRemove := dcl.CompareStringSets(desired.RemoteTrafficSelector, actual.RemoteTrafficSelector)
		if len(toAdd) > 0 {
			diffs = append(diffs, vpnTunnelDiff{
				RequiresRecreate: true,
				FieldName:        "RemoteTrafficSelector",
			})
		}
		if len(toRemove) > 0 {
			diffs = append(diffs, vpnTunnelDiff{
				RequiresRecreate: true,
				FieldName:        "RemoteTrafficSelector",
			})
		}
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
	var deduped []vpnTunnelDiff
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
func compareVpnTunnelStatusEnumSlice(c *Client, desired, actual []VpnTunnelStatusEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in VpnTunnelStatusEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareVpnTunnelStatusEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in VpnTunnelStatusEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareVpnTunnelStatusEnum(c *Client, desired, actual *VpnTunnelStatusEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *VpnTunnel) urlNormalized() *VpnTunnel {
	normalized := deepcopy.Copy(*r).(VpnTunnel)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.TargetVpnGateway = dcl.SelfLinkToName(r.TargetVpnGateway)
	normalized.VpnGateway = dcl.SelfLinkToName(r.VpnGateway)
	normalized.PeerExternalGateway = dcl.SelfLinkToName(r.PeerExternalGateway)
	normalized.PeerGcpGateway = dcl.SelfLinkToName(r.PeerGcpGateway)
	normalized.Router = dcl.SelfLinkToName(r.Router)
	normalized.PeerIP = dcl.SelfLinkToName(r.PeerIP)
	normalized.SharedSecret = dcl.SelfLinkToName(r.SharedSecret)
	normalized.SharedSecretHash = dcl.SelfLinkToName(r.SharedSecretHash)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.DetailedStatus = dcl.SelfLinkToName(r.DetailedStatus)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *VpnTunnel) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *VpnTunnel) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *VpnTunnel) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *VpnTunnel) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "setLabels" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/vpnTunnels/setLabels", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the VpnTunnel resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *VpnTunnel) marshal(c *Client) ([]byte, error) {
	m, err := expandVpnTunnel(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling VpnTunnel: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalVpnTunnel decodes JSON responses into the VpnTunnel resource schema.
func unmarshalVpnTunnel(b []byte, c *Client) (*VpnTunnel, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapVpnTunnel(m, c)
}

func unmarshalMapVpnTunnel(m map[string]interface{}, c *Client) (*VpnTunnel, error) {

	return flattenVpnTunnel(c, m), nil
}

// expandVpnTunnel expands VpnTunnel into a JSON request object.
func expandVpnTunnel(c *Client, f *VpnTunnel) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
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
	if v := f.TargetVpnGateway; !dcl.IsEmptyValueIndirect(v) {
		m["targetVpnGateway"] = v
	}
	if v := f.VpnGateway; !dcl.IsEmptyValueIndirect(v) {
		m["vpnGateway"] = v
	}
	if v := f.VpnGatewayInterface; !dcl.IsEmptyValueIndirect(v) {
		m["vpnGatewayInterface"] = v
	}
	if v := f.PeerExternalGateway; !dcl.IsEmptyValueIndirect(v) {
		m["peerExternalGateway"] = v
	}
	if v := f.PeerExternalGatewayInterface; !dcl.IsEmptyValueIndirect(v) {
		m["peerExternalGatewayInterface"] = v
	}
	if v := f.PeerGcpGateway; !dcl.IsEmptyValueIndirect(v) {
		m["peerGcpGateway"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/regions/%s/routers/%s", f.Router, f.Project, f.Region, f.Router); err != nil {
		return nil, fmt.Errorf("error expanding Router into router: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["router"] = v
	}
	if v := f.PeerIP; !dcl.IsEmptyValueIndirect(v) {
		m["peerIp"] = v
	}
	if v := f.SharedSecret; !dcl.IsEmptyValueIndirect(v) {
		m["sharedSecret"] = v
	}
	if v := f.SharedSecretHash; !dcl.IsEmptyValueIndirect(v) {
		m["sharedSecretHash"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.IkeVersion; !dcl.IsEmptyValueIndirect(v) {
		m["ikeVersion"] = v
	}
	if v := f.DetailedStatus; !dcl.IsEmptyValueIndirect(v) {
		m["detailedStatus"] = v
	}
	if v := f.LocalTrafficSelector; !dcl.IsEmptyValueIndirect(v) {
		m["localTrafficSelector"] = v
	}
	if v := f.RemoteTrafficSelector; !dcl.IsEmptyValueIndirect(v) {
		m["remoteTrafficSelector"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenVpnTunnel flattens VpnTunnel from a JSON request object into the
// VpnTunnel type.
func flattenVpnTunnel(c *Client, i interface{}) *VpnTunnel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &VpnTunnel{}
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Id = dcl.FlattenInteger(m["id"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.Region = dcl.FlattenString(m["region"])
	r.TargetVpnGateway = dcl.FlattenString(m["targetVpnGateway"])
	r.VpnGateway = dcl.FlattenString(m["vpnGateway"])
	r.VpnGatewayInterface = dcl.FlattenInteger(m["vpnGatewayInterface"])
	r.PeerExternalGateway = dcl.FlattenString(m["peerExternalGateway"])
	r.PeerExternalGatewayInterface = dcl.FlattenInteger(m["peerExternalGatewayInterface"])
	r.PeerGcpGateway = dcl.FlattenString(m["peerGcpGateway"])
	r.Router = dcl.FlattenString(m["router"])
	r.PeerIP = dcl.FlattenString(m["peerIp"])
	r.SharedSecret = dcl.FlattenSecretValue(m["sharedSecret"])
	r.SharedSecretHash = dcl.FlattenString(m["sharedSecretHash"])
	r.Status = flattenVpnTunnelStatusEnum(m["status"])
	r.SelfLink = dcl.FlattenString(m["selfLink"])
	r.IkeVersion = dcl.FlattenInteger(m["ikeVersion"])
	if _, ok := m["ikeVersion"]; !ok {
		c.Config.Logger.Info("Using default value for ikeVersion")
		r.IkeVersion = dcl.Int64(2)
	}
	r.DetailedStatus = dcl.FlattenString(m["detailedStatus"])
	r.LocalTrafficSelector = dcl.FlattenStringSlice(m["localTrafficSelector"])
	r.RemoteTrafficSelector = dcl.FlattenStringSlice(m["remoteTrafficSelector"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// flattenVpnTunnelStatusEnumSlice flattens the contents of VpnTunnelStatusEnum from a JSON
// response object.
func flattenVpnTunnelStatusEnumSlice(c *Client, i interface{}) []VpnTunnelStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VpnTunnelStatusEnum{}
	}

	if len(a) == 0 {
		return []VpnTunnelStatusEnum{}
	}

	items := make([]VpnTunnelStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVpnTunnelStatusEnum(item.(interface{})))
	}

	return items
}

// flattenVpnTunnelStatusEnum asserts that an interface is a string, and returns a
// pointer to a *VpnTunnelStatusEnum with the same value as that string.
func flattenVpnTunnelStatusEnum(i interface{}) *VpnTunnelStatusEnum {
	s, ok := i.(string)
	if !ok {
		return VpnTunnelStatusEnumRef("")
	}

	return VpnTunnelStatusEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *VpnTunnel) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalVpnTunnel(b, c)
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
