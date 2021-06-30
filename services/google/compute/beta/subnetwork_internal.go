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

func (r *Subnetwork) validate() error {

	if err := dcl.Required(r, "ipCidrRange"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "network"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LogConfig) {
		if err := r.LogConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *SubnetworkSecondaryIPRanges) validate() error {
	if err := dcl.Required(r, "rangeName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "ipCidrRange"); err != nil {
		return err
	}
	return nil
}
func (r *SubnetworkLogConfig) validate() error {
	return nil
}

func subnetworkGetURL(userBasePath string, r *Subnetwork) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

func subnetworkListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func subnetworkCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil

}

func subnetworkDeleteURL(userBasePath string, r *Subnetwork) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, params), nil
}

// subnetworkApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type subnetworkApiOperation interface {
	do(context.Context, *Subnetwork, *Client) error
}

// newUpdateSubnetworkExpandIpCidrRangeRequest creates a request for an
// Subnetwork resource's expandIpCidrRange update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSubnetworkExpandIpCidrRangeRequest(ctx context.Context, f *Subnetwork, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.IPCidrRange; !dcl.IsEmptyValueIndirect(v) {
		req["ipCidrRange"] = v
	}
	return req, nil
}

// marshalUpdateSubnetworkExpandIpCidrRangeRequest converts the update into
// the final JSON request body.
func marshalUpdateSubnetworkExpandIpCidrRangeRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateSubnetworkExpandIpCidrRangeOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSubnetworkExpandIpCidrRangeOperation) do(ctx context.Context, r *Subnetwork, c *Client) error {
	_, err := c.GetSubnetwork(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "expandIpCidrRange")
	if err != nil {
		return err
	}

	req, err := newUpdateSubnetworkExpandIpCidrRangeRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSubnetworkExpandIpCidrRangeRequest(c, req)
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

// newUpdateSubnetworkSetPrivateIpGoogleAccessRequest creates a request for an
// Subnetwork resource's setPrivateIpGoogleAccess update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSubnetworkSetPrivateIpGoogleAccessRequest(ctx context.Context, f *Subnetwork, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.PrivateIPGoogleAccess; !dcl.IsEmptyValueIndirect(v) {
		req["privateIpGoogleAccess"] = v
	}
	return req, nil
}

// marshalUpdateSubnetworkSetPrivateIpGoogleAccessRequest converts the update into
// the final JSON request body.
func marshalUpdateSubnetworkSetPrivateIpGoogleAccessRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateSubnetworkSetPrivateIpGoogleAccessOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSubnetworkSetPrivateIpGoogleAccessOperation) do(ctx context.Context, r *Subnetwork, c *Client) error {
	_, err := c.GetSubnetwork(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "setPrivateIpGoogleAccess")
	if err != nil {
		return err
	}

	req, err := newUpdateSubnetworkSetPrivateIpGoogleAccessRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSubnetworkSetPrivateIpGoogleAccessRequest(c, req)
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

// newUpdateSubnetworkUpdateRequest creates a request for an
// Subnetwork resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateSubnetworkUpdateRequest(ctx context.Context, f *Subnetwork, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		req["fingerprint"] = v
	}
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		req["role"] = v
	}
	if v, err := expandSubnetworkSecondaryIPRangesSlice(c, f.SecondaryIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryIPRanges into secondaryIpRanges: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["secondaryIpRanges"] = v
	}
	if v, err := expandSubnetworkLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["logConfig"] = v
	}
	return req, nil
}

// marshalUpdateSubnetworkUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateSubnetworkUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateSubnetworkUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateSubnetworkUpdateOperation) do(ctx context.Context, r *Subnetwork, c *Client) error {
	_, err := c.GetSubnetwork(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateSubnetworkUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateSubnetworkUpdateRequest(c, req)
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

func (c *Client) listSubnetworkRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := subnetworkListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != SubnetworkMaxPage {
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

type listSubnetworkOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listSubnetwork(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*Subnetwork, string, error) {
	b, err := c.listSubnetworkRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listSubnetworkOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Subnetwork
	for _, v := range m.Items {
		res, err := unmarshalMapSubnetwork(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllSubnetwork(ctx context.Context, f func(*Subnetwork) bool, resources []*Subnetwork) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteSubnetwork(ctx, res)
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

type deleteSubnetworkOperation struct{}

func (op *deleteSubnetworkOperation) do(ctx context.Context, r *Subnetwork, c *Client) error {
	r, err := c.GetSubnetwork(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Subnetwork not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetSubnetwork checking for existence. error: %v", err)
		return err
	}

	u, err := subnetworkDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetSubnetwork(ctx, r.URLNormalized())
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
type createSubnetworkOperation struct {
	response map[string]interface{}
}

func (op *createSubnetworkOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createSubnetworkOperation) do(ctx context.Context, r *Subnetwork, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := subnetworkCreateURL(c.Config.BasePath, project, region)

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

	if _, err := c.GetSubnetwork(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getSubnetworkRaw(ctx context.Context, r *Subnetwork) ([]byte, error) {

	u, err := subnetworkGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) subnetworkDiffsForRawDesired(ctx context.Context, rawDesired *Subnetwork, opts ...dcl.ApplyOption) (initial, desired *Subnetwork, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Subnetwork
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Subnetwork); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Subnetwork, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetSubnetwork(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Subnetwork resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Subnetwork resource: %v", err)
		}
		c.Config.Logger.Info("Found that Subnetwork resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeSubnetworkDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Subnetwork: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Subnetwork: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeSubnetworkInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Subnetwork: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeSubnetworkDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Subnetwork: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffSubnetwork(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeSubnetworkInitialState(rawInitial, rawDesired *Subnetwork) (*Subnetwork, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeSubnetworkDesiredState(rawDesired, rawInitial *Subnetwork, opts ...dcl.ApplyOption) (*Subnetwork, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.LogConfig = canonicalizeSubnetworkLogConfig(rawDesired.LogConfig, nil, opts...)

		return rawDesired, nil
	}

	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.IPCidrRange, rawInitial.IPCidrRange) {
		rawDesired.IPCidrRange = rawInitial.IPCidrRange
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.Purpose) {
		rawDesired.Purpose = rawInitial.Purpose
	}
	if dcl.IsZeroValue(rawDesired.Role) {
		rawDesired.Role = rawInitial.Role
	}
	if dcl.IsZeroValue(rawDesired.SecondaryIPRanges) {
		rawDesired.SecondaryIPRanges = rawInitial.SecondaryIPRanges
	}
	if dcl.BoolCanonicalize(rawDesired.PrivateIPGoogleAccess, rawInitial.PrivateIPGoogleAccess) {
		rawDesired.PrivateIPGoogleAccess = rawInitial.PrivateIPGoogleAccess
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	rawDesired.LogConfig = canonicalizeSubnetworkLogConfig(rawDesired.LogConfig, rawInitial.LogConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.BoolCanonicalize(rawDesired.EnableFlowLogs, rawInitial.EnableFlowLogs) {
		rawDesired.EnableFlowLogs = rawInitial.EnableFlowLogs
	}

	return rawDesired, nil
}

func canonicalizeSubnetworkNewState(c *Client, rawNew, rawDesired *Subnetwork) (*Subnetwork, error) {

	if dcl.IsEmptyValueIndirect(rawNew.CreationTimestamp) && dcl.IsEmptyValueIndirect(rawDesired.CreationTimestamp) {
		rawNew.CreationTimestamp = rawDesired.CreationTimestamp
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GatewayAddress) && dcl.IsEmptyValueIndirect(rawDesired.GatewayAddress) {
		rawNew.GatewayAddress = rawDesired.GatewayAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.GatewayAddress, rawNew.GatewayAddress) {
			rawNew.GatewayAddress = rawDesired.GatewayAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPCidrRange) && dcl.IsEmptyValueIndirect(rawDesired.IPCidrRange) {
		rawNew.IPCidrRange = rawDesired.IPCidrRange
	} else {
		if dcl.StringCanonicalize(rawDesired.IPCidrRange, rawNew.IPCidrRange) {
			rawNew.IPCidrRange = rawDesired.IPCidrRange
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.Fingerprint) && dcl.IsEmptyValueIndirect(rawDesired.Fingerprint) {
		rawNew.Fingerprint = rawDesired.Fingerprint
	} else {
		if dcl.StringCanonicalize(rawDesired.Fingerprint, rawNew.Fingerprint) {
			rawNew.Fingerprint = rawDesired.Fingerprint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Purpose) && dcl.IsEmptyValueIndirect(rawDesired.Purpose) {
		rawNew.Purpose = rawDesired.Purpose
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Role) && dcl.IsEmptyValueIndirect(rawDesired.Role) {
		rawNew.Role = rawDesired.Role
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SecondaryIPRanges) && dcl.IsEmptyValueIndirect(rawDesired.SecondaryIPRanges) {
		rawNew.SecondaryIPRanges = rawDesired.SecondaryIPRanges
	} else {
		rawNew.SecondaryIPRanges = canonicalizeNewSubnetworkSecondaryIPRangesSlice(c, rawDesired.SecondaryIPRanges, rawNew.SecondaryIPRanges)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateIPGoogleAccess) && dcl.IsEmptyValueIndirect(rawDesired.PrivateIPGoogleAccess) {
		rawNew.PrivateIPGoogleAccess = rawDesired.PrivateIPGoogleAccess
	} else {
		if dcl.BoolCanonicalize(rawDesired.PrivateIPGoogleAccess, rawNew.PrivateIPGoogleAccess) {
			rawNew.PrivateIPGoogleAccess = rawDesired.PrivateIPGoogleAccess
		}
	}

	rawNew.Region = rawDesired.Region

	if dcl.IsEmptyValueIndirect(rawNew.LogConfig) && dcl.IsEmptyValueIndirect(rawDesired.LogConfig) {
		rawNew.LogConfig = rawDesired.LogConfig
	} else {
		rawNew.LogConfig = canonicalizeNewSubnetworkLogConfig(c, rawDesired.LogConfig, rawNew.LogConfig)
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableFlowLogs) && dcl.IsEmptyValueIndirect(rawDesired.EnableFlowLogs) {
		rawNew.EnableFlowLogs = rawDesired.EnableFlowLogs
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableFlowLogs, rawNew.EnableFlowLogs) {
			rawNew.EnableFlowLogs = rawDesired.EnableFlowLogs
		}
	}

	return rawNew, nil
}

func canonicalizeSubnetworkSecondaryIPRanges(des, initial *SubnetworkSecondaryIPRanges, opts ...dcl.ApplyOption) *SubnetworkSecondaryIPRanges {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.RangeName, initial.RangeName) || dcl.IsZeroValue(des.RangeName) {
		des.RangeName = initial.RangeName
	}
	if dcl.StringCanonicalize(des.IPCidrRange, initial.IPCidrRange) || dcl.IsZeroValue(des.IPCidrRange) {
		des.IPCidrRange = initial.IPCidrRange
	}

	return des
}

func canonicalizeNewSubnetworkSecondaryIPRanges(c *Client, des, nw *SubnetworkSecondaryIPRanges) *SubnetworkSecondaryIPRanges {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.RangeName, nw.RangeName) {
		nw.RangeName = des.RangeName
	}
	if dcl.StringCanonicalize(des.IPCidrRange, nw.IPCidrRange) {
		nw.IPCidrRange = des.IPCidrRange
	}

	return nw
}

func canonicalizeNewSubnetworkSecondaryIPRangesSet(c *Client, des, nw []SubnetworkSecondaryIPRanges) []SubnetworkSecondaryIPRanges {
	if des == nil {
		return nw
	}
	var reorderedNew []SubnetworkSecondaryIPRanges
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSubnetworkSecondaryIPRangesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSubnetworkSecondaryIPRangesSlice(c *Client, des, nw []SubnetworkSecondaryIPRanges) []SubnetworkSecondaryIPRanges {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SubnetworkSecondaryIPRanges
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubnetworkSecondaryIPRanges(c, &d, &n))
	}

	return items
}

func canonicalizeSubnetworkLogConfig(des, initial *SubnetworkLogConfig, opts ...dcl.ApplyOption) *SubnetworkLogConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.AggregationInterval) {
		des.AggregationInterval = SubnetworkLogConfigAggregationIntervalEnumRef("INTERVAL_5_SEC")
	}

	if dcl.IsZeroValue(des.FlowSampling) {
		des.FlowSampling = dcl.Float64(0.5)
	}

	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = SubnetworkLogConfigMetadataEnumRef("INCLUDE_ALL_METADATA")
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AggregationInterval) {
		des.AggregationInterval = initial.AggregationInterval
	}
	if dcl.IsZeroValue(des.FlowSampling) {
		des.FlowSampling = initial.FlowSampling
	}
	if dcl.IsZeroValue(des.Metadata) {
		des.Metadata = initial.Metadata
	}

	return des
}

func canonicalizeNewSubnetworkLogConfig(c *Client, des, nw *SubnetworkLogConfig) *SubnetworkLogConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AggregationInterval) {
		nw.AggregationInterval = SubnetworkLogConfigAggregationIntervalEnumRef("INTERVAL_5_SEC")
	}

	if dcl.IsZeroValue(nw.FlowSampling) {
		nw.FlowSampling = dcl.Float64(0.5)
	}

	if dcl.IsZeroValue(nw.Metadata) {
		nw.Metadata = SubnetworkLogConfigMetadataEnumRef("INCLUDE_ALL_METADATA")
	}

	if dcl.IsZeroValue(nw.AggregationInterval) {
		nw.AggregationInterval = des.AggregationInterval
	}
	if dcl.IsZeroValue(nw.FlowSampling) {
		nw.FlowSampling = des.FlowSampling
	}
	if dcl.IsZeroValue(nw.Metadata) {
		nw.Metadata = des.Metadata
	}

	return nw
}

func canonicalizeNewSubnetworkLogConfigSet(c *Client, des, nw []SubnetworkLogConfig) []SubnetworkLogConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []SubnetworkLogConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareSubnetworkLogConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewSubnetworkLogConfigSlice(c *Client, des, nw []SubnetworkLogConfig) []SubnetworkLogConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []SubnetworkLogConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewSubnetworkLogConfig(c, &d, &n))
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
func diffSubnetwork(c *Client, desired, actual *Subnetwork, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GatewayAddress, actual.GatewayAddress, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GatewayAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPCidrRange, actual.IPCidrRange, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubnetworkExpandIpCidrRangeOperation")}, fn.AddNest("IpCidrRange")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fingerprint, actual.Fingerprint, dcl.Info{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Purpose, actual.Purpose, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Purpose")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Role, actual.Role, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("Role")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecondaryIPRanges, actual.SecondaryIPRanges, dcl.Info{ObjectFunction: compareSubnetworkSecondaryIPRangesNewStyle, EmptyObject: EmptySubnetworkSecondaryIPRanges, OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("SecondaryIpRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateIPGoogleAccess, actual.PrivateIPGoogleAccess, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubnetworkSetPrivateIpGoogleAccessOperation")}, fn.AddNest("PrivateIpGoogleAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogConfig, actual.LogConfig, dcl.Info{ObjectFunction: compareSubnetworkLogConfigNewStyle, EmptyObject: EmptySubnetworkLogConfig, OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("LogConfig")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableFlowLogs, actual.EnableFlowLogs, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableFlowLogs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareSubnetworkSecondaryIPRangesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubnetworkSecondaryIPRanges)
	if !ok {
		desiredNotPointer, ok := d.(SubnetworkSecondaryIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubnetworkSecondaryIPRanges or *SubnetworkSecondaryIPRanges", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubnetworkSecondaryIPRanges)
	if !ok {
		actualNotPointer, ok := a.(SubnetworkSecondaryIPRanges)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubnetworkSecondaryIPRanges", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RangeName, actual.RangeName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RangeName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPCidrRange, actual.IPCidrRange, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpCidrRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareSubnetworkLogConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*SubnetworkLogConfig)
	if !ok {
		desiredNotPointer, ok := d.(SubnetworkLogConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubnetworkLogConfig or *SubnetworkLogConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*SubnetworkLogConfig)
	if !ok {
		actualNotPointer, ok := a.(SubnetworkLogConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a SubnetworkLogConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AggregationInterval, actual.AggregationInterval, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("AggregationInterval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FlowSampling, actual.FlowSampling, dcl.Info{OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("FlowSampling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Metadata, actual.Metadata, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateSubnetworkUpdateOperation")}, fn.AddNest("Metadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Subnetwork) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subnetwork) createFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *Subnetwork) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Subnetwork) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "expandIpCidrRange" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks/{{name}}/expandIpCidrRange", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	if updateName == "setPrivateIpGoogleAccess" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks/{{name}}/setPrivateIpGoogleAccess", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/regions/{{region}}/subnetworks/{{name}}", "https://www.googleapis.com/compute/beta/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Subnetwork resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Subnetwork) marshal(c *Client) ([]byte, error) {
	m, err := expandSubnetwork(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Subnetwork: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalSubnetwork decodes JSON responses into the Subnetwork resource schema.
func unmarshalSubnetwork(b []byte, c *Client) (*Subnetwork, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapSubnetwork(m, c)
}

func unmarshalMapSubnetwork(m map[string]interface{}, c *Client) (*Subnetwork, error) {

	return flattenSubnetwork(c, m), nil
}

// expandSubnetwork expands Subnetwork into a JSON request object.
func expandSubnetwork(c *Client, f *Subnetwork) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.GatewayAddress; !dcl.IsEmptyValueIndirect(v) {
		m["gatewayAddress"] = v
	}
	if v := f.IPCidrRange; !dcl.IsEmptyValueIndirect(v) {
		m["ipCidrRange"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["fingerprint"] = v
	}
	if v := f.Purpose; !dcl.IsEmptyValueIndirect(v) {
		m["purpose"] = v
	}
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		m["role"] = v
	}
	if v, err := expandSubnetworkSecondaryIPRangesSlice(c, f.SecondaryIPRanges); err != nil {
		return nil, fmt.Errorf("error expanding SecondaryIPRanges into secondaryIpRanges: %w", err)
	} else if v != nil {
		m["secondaryIpRanges"] = v
	}
	if v := f.PrivateIPGoogleAccess; !dcl.IsEmptyValueIndirect(v) {
		m["privateIpGoogleAccess"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if v != nil {
		m["region"] = v
	}
	if v, err := expandSubnetworkLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if v != nil {
		m["logConfig"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v := f.EnableFlowLogs; !dcl.IsEmptyValueIndirect(v) {
		m["enableFlowLogs"] = v
	}

	return m, nil
}

// flattenSubnetwork flattens Subnetwork from a JSON request object into the
// Subnetwork type.
func flattenSubnetwork(c *Client, i interface{}) *Subnetwork {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Subnetwork{}
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Description = dcl.FlattenString(m["description"])
	res.GatewayAddress = dcl.FlattenString(m["gatewayAddress"])
	res.IPCidrRange = dcl.FlattenString(m["ipCidrRange"])
	res.Name = dcl.FlattenString(m["name"])
	res.Network = dcl.FlattenString(m["network"])
	res.Fingerprint = dcl.FlattenString(m["fingerprint"])
	res.Purpose = flattenSubnetworkPurposeEnum(m["purpose"])
	res.Role = flattenSubnetworkRoleEnum(m["role"])
	res.SecondaryIPRanges = flattenSubnetworkSecondaryIPRangesSlice(c, m["secondaryIpRanges"])
	res.PrivateIPGoogleAccess = dcl.FlattenBool(m["privateIpGoogleAccess"])
	res.Region = dcl.FlattenString(m["region"])
	res.LogConfig = flattenSubnetworkLogConfig(c, m["logConfig"])
	res.Project = dcl.FlattenString(m["project"])
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.EnableFlowLogs = dcl.FlattenBool(m["enableFlowLogs"])

	return res
}

// expandSubnetworkSecondaryIPRangesMap expands the contents of SubnetworkSecondaryIPRanges into a JSON
// request object.
func expandSubnetworkSecondaryIPRangesMap(c *Client, f map[string]SubnetworkSecondaryIPRanges) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubnetworkSecondaryIPRanges(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubnetworkSecondaryIPRangesSlice expands the contents of SubnetworkSecondaryIPRanges into a JSON
// request object.
func expandSubnetworkSecondaryIPRangesSlice(c *Client, f []SubnetworkSecondaryIPRanges) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubnetworkSecondaryIPRanges(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubnetworkSecondaryIPRangesMap flattens the contents of SubnetworkSecondaryIPRanges from a JSON
// response object.
func flattenSubnetworkSecondaryIPRangesMap(c *Client, i interface{}) map[string]SubnetworkSecondaryIPRanges {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubnetworkSecondaryIPRanges{}
	}

	if len(a) == 0 {
		return map[string]SubnetworkSecondaryIPRanges{}
	}

	items := make(map[string]SubnetworkSecondaryIPRanges)
	for k, item := range a {
		items[k] = *flattenSubnetworkSecondaryIPRanges(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubnetworkSecondaryIPRangesSlice flattens the contents of SubnetworkSecondaryIPRanges from a JSON
// response object.
func flattenSubnetworkSecondaryIPRangesSlice(c *Client, i interface{}) []SubnetworkSecondaryIPRanges {
	a, ok := i.([]interface{})
	if !ok {
		return []SubnetworkSecondaryIPRanges{}
	}

	if len(a) == 0 {
		return []SubnetworkSecondaryIPRanges{}
	}

	items := make([]SubnetworkSecondaryIPRanges, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubnetworkSecondaryIPRanges(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubnetworkSecondaryIPRanges expands an instance of SubnetworkSecondaryIPRanges into a JSON
// request object.
func expandSubnetworkSecondaryIPRanges(c *Client, f *SubnetworkSecondaryIPRanges) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RangeName; !dcl.IsEmptyValueIndirect(v) {
		m["rangeName"] = v
	}
	if v := f.IPCidrRange; !dcl.IsEmptyValueIndirect(v) {
		m["ipCidrRange"] = v
	}

	return m, nil
}

// flattenSubnetworkSecondaryIPRanges flattens an instance of SubnetworkSecondaryIPRanges from a JSON
// response object.
func flattenSubnetworkSecondaryIPRanges(c *Client, i interface{}) *SubnetworkSecondaryIPRanges {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubnetworkSecondaryIPRanges{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySubnetworkSecondaryIPRanges
	}
	r.RangeName = dcl.FlattenString(m["rangeName"])
	r.IPCidrRange = dcl.FlattenString(m["ipCidrRange"])

	return r
}

// expandSubnetworkLogConfigMap expands the contents of SubnetworkLogConfig into a JSON
// request object.
func expandSubnetworkLogConfigMap(c *Client, f map[string]SubnetworkLogConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandSubnetworkLogConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandSubnetworkLogConfigSlice expands the contents of SubnetworkLogConfig into a JSON
// request object.
func expandSubnetworkLogConfigSlice(c *Client, f []SubnetworkLogConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandSubnetworkLogConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenSubnetworkLogConfigMap flattens the contents of SubnetworkLogConfig from a JSON
// response object.
func flattenSubnetworkLogConfigMap(c *Client, i interface{}) map[string]SubnetworkLogConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]SubnetworkLogConfig{}
	}

	if len(a) == 0 {
		return map[string]SubnetworkLogConfig{}
	}

	items := make(map[string]SubnetworkLogConfig)
	for k, item := range a {
		items[k] = *flattenSubnetworkLogConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenSubnetworkLogConfigSlice flattens the contents of SubnetworkLogConfig from a JSON
// response object.
func flattenSubnetworkLogConfigSlice(c *Client, i interface{}) []SubnetworkLogConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []SubnetworkLogConfig{}
	}

	if len(a) == 0 {
		return []SubnetworkLogConfig{}
	}

	items := make([]SubnetworkLogConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubnetworkLogConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandSubnetworkLogConfig expands an instance of SubnetworkLogConfig into a JSON
// request object.
func expandSubnetworkLogConfig(c *Client, f *SubnetworkLogConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AggregationInterval; !dcl.IsEmptyValueIndirect(v) {
		m["aggregationInterval"] = v
	}
	if v := f.FlowSampling; !dcl.IsEmptyValueIndirect(v) {
		m["flowSampling"] = v
	}
	if v := f.Metadata; !dcl.IsEmptyValueIndirect(v) {
		m["metadata"] = v
	}

	return m, nil
}

// flattenSubnetworkLogConfig flattens an instance of SubnetworkLogConfig from a JSON
// response object.
func flattenSubnetworkLogConfig(c *Client, i interface{}) *SubnetworkLogConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &SubnetworkLogConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptySubnetworkLogConfig
	}
	r.AggregationInterval = flattenSubnetworkLogConfigAggregationIntervalEnum(m["aggregationInterval"])
	if dcl.IsEmptyValueIndirect(m["aggregationInterval"]) {
		c.Config.Logger.Info("Using default value for aggregationInterval.")
		r.AggregationInterval = SubnetworkLogConfigAggregationIntervalEnumRef("INTERVAL_5_SEC")
	}
	r.FlowSampling = dcl.FlattenDouble(m["flowSampling"])
	if dcl.IsEmptyValueIndirect(m["flowSampling"]) {
		c.Config.Logger.Info("Using default value for flowSampling.")
		r.FlowSampling = dcl.Float64(0.5)
	}
	r.Metadata = flattenSubnetworkLogConfigMetadataEnum(m["metadata"])
	if dcl.IsEmptyValueIndirect(m["metadata"]) {
		c.Config.Logger.Info("Using default value for metadata.")
		r.Metadata = SubnetworkLogConfigMetadataEnumRef("INCLUDE_ALL_METADATA")
	}

	return r
}

// flattenSubnetworkPurposeEnumSlice flattens the contents of SubnetworkPurposeEnum from a JSON
// response object.
func flattenSubnetworkPurposeEnumSlice(c *Client, i interface{}) []SubnetworkPurposeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SubnetworkPurposeEnum{}
	}

	if len(a) == 0 {
		return []SubnetworkPurposeEnum{}
	}

	items := make([]SubnetworkPurposeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubnetworkPurposeEnum(item.(interface{})))
	}

	return items
}

// flattenSubnetworkPurposeEnum asserts that an interface is a string, and returns a
// pointer to a *SubnetworkPurposeEnum with the same value as that string.
func flattenSubnetworkPurposeEnum(i interface{}) *SubnetworkPurposeEnum {
	s, ok := i.(string)
	if !ok {
		return SubnetworkPurposeEnumRef("")
	}

	return SubnetworkPurposeEnumRef(s)
}

// flattenSubnetworkRoleEnumSlice flattens the contents of SubnetworkRoleEnum from a JSON
// response object.
func flattenSubnetworkRoleEnumSlice(c *Client, i interface{}) []SubnetworkRoleEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SubnetworkRoleEnum{}
	}

	if len(a) == 0 {
		return []SubnetworkRoleEnum{}
	}

	items := make([]SubnetworkRoleEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubnetworkRoleEnum(item.(interface{})))
	}

	return items
}

// flattenSubnetworkRoleEnum asserts that an interface is a string, and returns a
// pointer to a *SubnetworkRoleEnum with the same value as that string.
func flattenSubnetworkRoleEnum(i interface{}) *SubnetworkRoleEnum {
	s, ok := i.(string)
	if !ok {
		return SubnetworkRoleEnumRef("")
	}

	return SubnetworkRoleEnumRef(s)
}

// flattenSubnetworkLogConfigAggregationIntervalEnumSlice flattens the contents of SubnetworkLogConfigAggregationIntervalEnum from a JSON
// response object.
func flattenSubnetworkLogConfigAggregationIntervalEnumSlice(c *Client, i interface{}) []SubnetworkLogConfigAggregationIntervalEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SubnetworkLogConfigAggregationIntervalEnum{}
	}

	if len(a) == 0 {
		return []SubnetworkLogConfigAggregationIntervalEnum{}
	}

	items := make([]SubnetworkLogConfigAggregationIntervalEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubnetworkLogConfigAggregationIntervalEnum(item.(interface{})))
	}

	return items
}

// flattenSubnetworkLogConfigAggregationIntervalEnum asserts that an interface is a string, and returns a
// pointer to a *SubnetworkLogConfigAggregationIntervalEnum with the same value as that string.
func flattenSubnetworkLogConfigAggregationIntervalEnum(i interface{}) *SubnetworkLogConfigAggregationIntervalEnum {
	s, ok := i.(string)
	if !ok {
		return SubnetworkLogConfigAggregationIntervalEnumRef("")
	}

	return SubnetworkLogConfigAggregationIntervalEnumRef(s)
}

// flattenSubnetworkLogConfigMetadataEnumSlice flattens the contents of SubnetworkLogConfigMetadataEnum from a JSON
// response object.
func flattenSubnetworkLogConfigMetadataEnumSlice(c *Client, i interface{}) []SubnetworkLogConfigMetadataEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []SubnetworkLogConfigMetadataEnum{}
	}

	if len(a) == 0 {
		return []SubnetworkLogConfigMetadataEnum{}
	}

	items := make([]SubnetworkLogConfigMetadataEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenSubnetworkLogConfigMetadataEnum(item.(interface{})))
	}

	return items
}

// flattenSubnetworkLogConfigMetadataEnum asserts that an interface is a string, and returns a
// pointer to a *SubnetworkLogConfigMetadataEnum with the same value as that string.
func flattenSubnetworkLogConfigMetadataEnum(i interface{}) *SubnetworkLogConfigMetadataEnum {
	s, ok := i.(string)
	if !ok {
		return SubnetworkLogConfigMetadataEnumRef("")
	}

	return SubnetworkLogConfigMetadataEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Subnetwork) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalSubnetwork(b, c)
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

type subnetworkDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         subnetworkApiOperation
}

func convertFieldDiffToSubnetworkOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]subnetworkDiff, error) {
	var diffs []subnetworkDiff
	for _, op := range ops {
		diff := subnetworkDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTosubnetworkApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTosubnetworkApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (subnetworkApiOperation, error) {
	switch op {

	case "updateSubnetworkExpandIpCidrRangeOperation":
		return &updateSubnetworkExpandIpCidrRangeOperation{Diffs: diffs}, nil

	case "updateSubnetworkSetPrivateIpGoogleAccessOperation":
		return &updateSubnetworkSetPrivateIpGoogleAccessOperation{Diffs: diffs}, nil

	case "updateSubnetworkUpdateOperation":
		return &updateSubnetworkUpdateOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
