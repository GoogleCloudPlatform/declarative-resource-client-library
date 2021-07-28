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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Firewall) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "network"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LogConfig) {
		if err := r.LogConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FirewallLogConfig) validate() error {
	return nil
}
func (r *FirewallAllowed) validate() error {
	if err := dcl.Required(r, "ipProtocol"); err != nil {
		return err
	}
	return nil
}
func (r *FirewallDenied) validate() error {
	if err := dcl.Required(r, "ipProtocol"); err != nil {
		return err
	}
	return nil
}

func firewallGetURL(userBasePath string, r *Firewall) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/firewalls/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

func firewallListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/firewalls", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func firewallCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/global/firewalls", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil

}

func firewallDeleteURL(userBasePath string, r *Firewall) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/global/firewalls/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, params), nil
}

// firewallApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type firewallApiOperation interface {
	do(context.Context, *Firewall, *Client) error
}

// newUpdateFirewallUpdateRequest creates a request for an
// Firewall resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFirewallUpdateRequest(ctx context.Context, f *Firewall, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		req["id"] = v
	}
	if v, err := expandFirewallLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["logConfig"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		req["network"] = v
	}
	if v := f.Priority; !dcl.IsEmptyValueIndirect(v) {
		req["priority"] = v
	}
	if v, err := expandFirewallAllowedSlice(c, f.Allowed); err != nil {
		return nil, fmt.Errorf("error expanding Allowed into allowed: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["allowed"] = v
	}
	if v, err := expandFirewallDeniedSlice(c, f.Denied); err != nil {
		return nil, fmt.Errorf("error expanding Denied into denied: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["denied"] = v
	}
	if v := f.SourceRanges; !dcl.IsEmptyValueIndirect(v) {
		req["sourceRanges"] = v
	}
	if v := f.SourceServiceAccounts; !dcl.IsEmptyValueIndirect(v) {
		req["sourceServiceAccounts"] = v
	}
	if v := f.SourceTags; !dcl.IsEmptyValueIndirect(v) {
		req["sourceTags"] = v
	}
	if v := f.TargetServiceAccounts; !dcl.IsEmptyValueIndirect(v) {
		req["targetServiceAccounts"] = v
	}
	if v := f.TargetTags; !dcl.IsEmptyValueIndirect(v) {
		req["targetTags"] = v
	}
	return req, nil
}

// marshalUpdateFirewallUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateFirewallUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFirewallUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFirewallUpdateOperation) do(ctx context.Context, r *Firewall, c *Client) error {
	_, err := c.GetFirewall(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateFirewallUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFirewallUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listFirewallRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := firewallListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FirewallMaxPage {
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

type listFirewallOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listFirewall(ctx context.Context, project, pageToken string, pageSize int32) ([]*Firewall, string, error) {
	b, err := c.listFirewallRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFirewallOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Firewall
	for _, v := range m.Items {
		res, err := unmarshalMapFirewall(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFirewall(ctx context.Context, f func(*Firewall) bool, resources []*Firewall) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFirewall(ctx, res)
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

type deleteFirewallOperation struct{}

func (op *deleteFirewallOperation) do(ctx context.Context, r *Firewall, c *Client) error {
	r, err := c.GetFirewall(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Firewall not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetFirewall checking for existence. error: %v", err)
		return err
	}

	u, err := firewallDeleteURL(c.Config.BasePath, r.URLNormalized())
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
		_, err = c.GetFirewall(ctx, r.URLNormalized())
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
type createFirewallOperation struct {
	response map[string]interface{}
}

func (op *createFirewallOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFirewallOperation) do(ctx context.Context, r *Firewall, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := firewallCreateURL(c.Config.BasePath, project)

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

	if _, err := c.GetFirewall(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFirewallRaw(ctx context.Context, r *Firewall) ([]byte, error) {
	if dcl.IsZeroValue(r.Priority) {
		r.Priority = dcl.Int64(1000)
	}

	u, err := firewallGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) firewallDiffsForRawDesired(ctx context.Context, rawDesired *Firewall, opts ...dcl.ApplyOption) (initial, desired *Firewall, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Firewall
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Firewall); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Firewall, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFirewall(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Firewall resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Firewall resource: %v", err)
		}
		c.Config.Logger.Info("Found that Firewall resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFirewallDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Firewall: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Firewall: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFirewallInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Firewall: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFirewallDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Firewall: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFirewall(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFirewallInitialState(rawInitial, rawDesired *Firewall) (*Firewall, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFirewallDesiredState(rawDesired, rawInitial *Firewall, opts ...dcl.ApplyOption) (*Firewall, error) {

	if dcl.IsZeroValue(rawDesired.Priority) {
		rawDesired.Priority = dcl.Int64(1000)
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.LogConfig = canonicalizeFirewallLogConfig(rawDesired.LogConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Firewall{}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Direction) {
		canonicalDesired.Direction = rawInitial.Direction
	} else {
		canonicalDesired.Direction = rawDesired.Direction
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		canonicalDesired.Disabled = rawInitial.Disabled
	} else {
		canonicalDesired.Disabled = rawDesired.Disabled
	}
	if dcl.StringCanonicalize(rawDesired.Id, rawInitial.Id) {
		canonicalDesired.Id = rawInitial.Id
	} else {
		canonicalDesired.Id = rawDesired.Id
	}
	canonicalDesired.LogConfig = canonicalizeFirewallLogConfig(rawDesired.LogConfig, rawInitial.LogConfig, opts...)
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Network) {
		canonicalDesired.Network = rawInitial.Network
	} else {
		canonicalDesired.Network = rawDesired.Network
	}
	if dcl.IsZeroValue(rawDesired.Priority) {
		canonicalDesired.Priority = rawInitial.Priority
	} else {
		canonicalDesired.Priority = rawDesired.Priority
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.IsZeroValue(rawDesired.Allowed) {
		canonicalDesired.Allowed = rawInitial.Allowed
	} else {
		canonicalDesired.Allowed = rawDesired.Allowed
	}
	if dcl.IsZeroValue(rawDesired.Denied) {
		canonicalDesired.Denied = rawInitial.Denied
	} else {
		canonicalDesired.Denied = rawDesired.Denied
	}
	if dcl.IsZeroValue(rawDesired.DestinationRanges) {
		canonicalDesired.DestinationRanges = rawInitial.DestinationRanges
	} else {
		canonicalDesired.DestinationRanges = rawDesired.DestinationRanges
	}
	if dcl.IsZeroValue(rawDesired.SourceRanges) {
		canonicalDesired.SourceRanges = rawInitial.SourceRanges
	} else {
		canonicalDesired.SourceRanges = rawDesired.SourceRanges
	}
	if dcl.IsZeroValue(rawDesired.SourceServiceAccounts) {
		canonicalDesired.SourceServiceAccounts = rawInitial.SourceServiceAccounts
	} else {
		canonicalDesired.SourceServiceAccounts = rawDesired.SourceServiceAccounts
	}
	if dcl.IsZeroValue(rawDesired.SourceTags) {
		canonicalDesired.SourceTags = rawInitial.SourceTags
	} else {
		canonicalDesired.SourceTags = rawDesired.SourceTags
	}
	if dcl.IsZeroValue(rawDesired.TargetServiceAccounts) {
		canonicalDesired.TargetServiceAccounts = rawInitial.TargetServiceAccounts
	} else {
		canonicalDesired.TargetServiceAccounts = rawDesired.TargetServiceAccounts
	}
	if dcl.IsZeroValue(rawDesired.TargetTags) {
		canonicalDesired.TargetTags = rawInitial.TargetTags
	} else {
		canonicalDesired.TargetTags = rawDesired.TargetTags
	}

	return canonicalDesired, nil
}

func canonicalizeFirewallNewState(c *Client, rawNew, rawDesired *Firewall) (*Firewall, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Direction) && dcl.IsEmptyValueIndirect(rawDesired.Direction) {
		rawNew.Direction = rawDesired.Direction
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LogConfig) && dcl.IsEmptyValueIndirect(rawDesired.LogConfig) {
		rawNew.LogConfig = rawDesired.LogConfig
	} else {
		rawNew.LogConfig = canonicalizeNewFirewallLogConfig(c, rawDesired.LogConfig, rawNew.LogConfig)
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
	}

	if dcl.IsEmptyValueIndirect(rawNew.Priority) && dcl.IsEmptyValueIndirect(rawDesired.Priority) {
		rawNew.Priority = rawDesired.Priority
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Allowed) && dcl.IsEmptyValueIndirect(rawDesired.Allowed) {
		rawNew.Allowed = rawDesired.Allowed
	} else {
		rawNew.Allowed = canonicalizeNewFirewallAllowedSet(c, rawDesired.Allowed, rawNew.Allowed)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Denied) && dcl.IsEmptyValueIndirect(rawDesired.Denied) {
		rawNew.Denied = rawDesired.Denied
	} else {
		rawNew.Denied = canonicalizeNewFirewallDeniedSet(c, rawDesired.Denied, rawNew.Denied)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DestinationRanges) && dcl.IsEmptyValueIndirect(rawDesired.DestinationRanges) {
		rawNew.DestinationRanges = rawDesired.DestinationRanges
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceRanges) && dcl.IsEmptyValueIndirect(rawDesired.SourceRanges) {
		rawNew.SourceRanges = rawDesired.SourceRanges
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceServiceAccounts) && dcl.IsEmptyValueIndirect(rawDesired.SourceServiceAccounts) {
		rawNew.SourceServiceAccounts = rawDesired.SourceServiceAccounts
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceTags) && dcl.IsEmptyValueIndirect(rawDesired.SourceTags) {
		rawNew.SourceTags = rawDesired.SourceTags
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetServiceAccounts) && dcl.IsEmptyValueIndirect(rawDesired.TargetServiceAccounts) {
		rawNew.TargetServiceAccounts = rawDesired.TargetServiceAccounts
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TargetTags) && dcl.IsEmptyValueIndirect(rawDesired.TargetTags) {
		rawNew.TargetTags = rawDesired.TargetTags
	} else {
	}

	return rawNew, nil
}

func canonicalizeFirewallLogConfig(des, initial *FirewallLogConfig, opts ...dcl.ApplyOption) *FirewallLogConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FirewallLogConfig{}

	if dcl.BoolCanonicalize(des.Enable, initial.Enable) || dcl.IsZeroValue(des.Enable) {
		cDes.Enable = initial.Enable
	} else {
		cDes.Enable = des.Enable
	}

	return cDes
}

func canonicalizeNewFirewallLogConfig(c *Client, des, nw *FirewallLogConfig) *FirewallLogConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Enable, nw.Enable) {
		nw.Enable = des.Enable
	}

	return nw
}

func canonicalizeNewFirewallLogConfigSet(c *Client, des, nw []FirewallLogConfig) []FirewallLogConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []FirewallLogConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFirewallLogConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFirewallLogConfigSlice(c *Client, des, nw []FirewallLogConfig) []FirewallLogConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FirewallLogConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFirewallLogConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFirewallAllowed(des, initial *FirewallAllowed, opts ...dcl.ApplyOption) *FirewallAllowed {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FirewallAllowed{}

	if dcl.StringCanonicalize(des.IPProtocol, initial.IPProtocol) || dcl.IsZeroValue(des.IPProtocol) {
		cDes.IPProtocol = initial.IPProtocol
	} else {
		cDes.IPProtocol = des.IPProtocol
	}
	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	} else {
		cDes.Ports = des.Ports
	}
	if dcl.IsZeroValue(des.IPProtocolAlt) {
		des.IPProtocolAlt = initial.IPProtocolAlt
	} else {
		cDes.IPProtocolAlt = des.IPProtocolAlt
	}

	return cDes
}

func canonicalizeNewFirewallAllowed(c *Client, des, nw *FirewallAllowed) *FirewallAllowed {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPProtocol, nw.IPProtocol) {
		nw.IPProtocol = des.IPProtocol
	}
	if dcl.IsZeroValue(nw.Ports) {
		nw.Ports = des.Ports
	}
	if dcl.IsZeroValue(nw.IPProtocolAlt) {
		nw.IPProtocolAlt = des.IPProtocolAlt
	}

	return nw
}

func canonicalizeNewFirewallAllowedSet(c *Client, des, nw []FirewallAllowed) []FirewallAllowed {
	if des == nil {
		return nw
	}
	var reorderedNew []FirewallAllowed
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFirewallAllowedNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFirewallAllowedSlice(c *Client, des, nw []FirewallAllowed) []FirewallAllowed {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FirewallAllowed
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFirewallAllowed(c, &d, &n))
	}

	return items
}

func canonicalizeFirewallDenied(des, initial *FirewallDenied, opts ...dcl.ApplyOption) *FirewallDenied {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FirewallDenied{}

	if dcl.StringCanonicalize(des.IPProtocol, initial.IPProtocol) || dcl.IsZeroValue(des.IPProtocol) {
		cDes.IPProtocol = initial.IPProtocol
	} else {
		cDes.IPProtocol = des.IPProtocol
	}
	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	} else {
		cDes.Ports = des.Ports
	}
	if dcl.IsZeroValue(des.IPProtocolAlt) {
		des.IPProtocolAlt = initial.IPProtocolAlt
	} else {
		cDes.IPProtocolAlt = des.IPProtocolAlt
	}

	return cDes
}

func canonicalizeNewFirewallDenied(c *Client, des, nw *FirewallDenied) *FirewallDenied {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPProtocol, nw.IPProtocol) {
		nw.IPProtocol = des.IPProtocol
	}
	if dcl.IsZeroValue(nw.Ports) {
		nw.Ports = des.Ports
	}
	if dcl.IsZeroValue(nw.IPProtocolAlt) {
		nw.IPProtocolAlt = des.IPProtocolAlt
	}

	return nw
}

func canonicalizeNewFirewallDeniedSet(c *Client, des, nw []FirewallDenied) []FirewallDenied {
	if des == nil {
		return nw
	}
	var reorderedNew []FirewallDenied
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFirewallDeniedNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewFirewallDeniedSlice(c *Client, des, nw []FirewallDenied) []FirewallDenied {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FirewallDenied
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFirewallDenied(c, &d, &n))
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
func diffFirewall(c *Client, desired, actual *Firewall, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Direction, actual.Direction, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Direction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogConfig, actual.LogConfig, dcl.Info{ObjectFunction: compareFirewallLogConfigNewStyle, EmptyObject: EmptyFirewallLogConfig, OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("LogConfig")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Priority, actual.Priority, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("Priority")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Allowed, actual.Allowed, dcl.Info{Type: "Set", ObjectFunction: compareFirewallAllowedNewStyle, EmptyObject: EmptyFirewallAllowed, OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("Allowed")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Denied, actual.Denied, dcl.Info{Type: "Set", ObjectFunction: compareFirewallDeniedNewStyle, EmptyObject: EmptyFirewallDenied, OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("Denied")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DestinationRanges, actual.DestinationRanges, dcl.Info{Type: "Set", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DestinationRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceRanges, actual.SourceRanges, dcl.Info{Type: "Set", OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("SourceRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceServiceAccounts, actual.SourceServiceAccounts, dcl.Info{Type: "Set", OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("SourceServiceAccounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceTags, actual.SourceTags, dcl.Info{Type: "Set", OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("SourceTags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetServiceAccounts, actual.TargetServiceAccounts, dcl.Info{Type: "Set", OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("TargetServiceAccounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetTags, actual.TargetTags, dcl.Info{Type: "Set", OperationSelector: dcl.TriggersOperation("updateFirewallUpdateOperation")}, fn.AddNest("TargetTags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareFirewallLogConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FirewallLogConfig)
	if !ok {
		desiredNotPointer, ok := d.(FirewallLogConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirewallLogConfig or *FirewallLogConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FirewallLogConfig)
	if !ok {
		actualNotPointer, ok := a.(FirewallLogConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirewallLogConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enable, actual.Enable, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFirewallAllowedNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FirewallAllowed)
	if !ok {
		desiredNotPointer, ok := d.(FirewallAllowed)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirewallAllowed or *FirewallAllowed", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FirewallAllowed)
	if !ok {
		actualNotPointer, ok := a.(FirewallAllowed)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirewallAllowed", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IPProtocol, actual.IPProtocol, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPProtocol")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPProtocolAlt, actual.IPProtocolAlt, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpProtocolAlt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFirewallDeniedNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FirewallDenied)
	if !ok {
		desiredNotPointer, ok := d.(FirewallDenied)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirewallDenied or *FirewallDenied", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FirewallDenied)
	if !ok {
		actualNotPointer, ok := a.(FirewallDenied)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirewallDenied", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IPProtocol, actual.IPProtocol, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IPProtocol")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPProtocolAlt, actual.IPProtocolAlt, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpProtocolAlt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *Firewall) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Firewall) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Firewall) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Firewall) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/global/firewalls/{{name}}", "https://www.googleapis.com/compute/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Firewall resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Firewall) marshal(c *Client) ([]byte, error) {
	m, err := expandFirewall(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Firewall: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFirewall decodes JSON responses into the Firewall resource schema.
func unmarshalFirewall(b []byte, c *Client) (*Firewall, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFirewall(m, c)
}

func unmarshalMapFirewall(m map[string]interface{}, c *Client) (*Firewall, error) {

	return flattenFirewall(c, m), nil
}

// expandFirewall expands Firewall into a JSON request object.
func expandFirewall(c *Client, f *Firewall) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.CreationTimestamp; !dcl.IsEmptyValueIndirect(v) {
		m["creationTimestamp"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.Direction; !dcl.IsEmptyValueIndirect(v) {
		m["direction"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v, err := expandFirewallLogConfig(c, f.LogConfig); err != nil {
		return nil, fmt.Errorf("error expanding LogConfig into logConfig: %w", err)
	} else if v != nil {
		m["logConfig"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Priority; !dcl.IsEmptyValueIndirect(v) {
		m["priority"] = v
	}
	if v := f.SelfLink; !dcl.IsEmptyValueIndirect(v) {
		m["selfLink"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := expandFirewallAllowedSlice(c, f.Allowed); err != nil {
		return nil, fmt.Errorf("error expanding Allowed into allowed: %w", err)
	} else {
		m["allowed"] = v
	}
	if v, err := expandFirewallDeniedSlice(c, f.Denied); err != nil {
		return nil, fmt.Errorf("error expanding Denied into denied: %w", err)
	} else {
		m["denied"] = v
	}
	m["destinationRanges"] = f.DestinationRanges
	m["sourceRanges"] = f.SourceRanges
	m["sourceServiceAccounts"] = f.SourceServiceAccounts
	m["sourceTags"] = f.SourceTags
	m["targetServiceAccounts"] = f.TargetServiceAccounts
	m["targetTags"] = f.TargetTags

	return m, nil
}

// flattenFirewall flattens Firewall from a JSON request object into the
// Firewall type.
func flattenFirewall(c *Client, i interface{}) *Firewall {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Firewall{}
	res.CreationTimestamp = dcl.FlattenString(m["creationTimestamp"])
	res.Description = dcl.FlattenString(m["description"])
	res.Direction = flattenFirewallDirectionEnum(m["direction"])
	res.Disabled = dcl.FlattenBool(m["disabled"])
	res.Id = dcl.FlattenString(m["id"])
	res.LogConfig = flattenFirewallLogConfig(c, m["logConfig"])
	res.Name = dcl.FlattenString(m["name"])
	res.Network = dcl.FlattenString(m["network"])
	res.Priority = dcl.FlattenInteger(m["priority"])
	if _, ok := m["priority"]; !ok {
		c.Config.Logger.Info("Using default value for priority")
		res.Priority = dcl.Int64(1000)
	}
	res.SelfLink = dcl.FlattenString(m["selfLink"])
	res.Project = dcl.FlattenString(m["project"])
	res.Allowed = flattenFirewallAllowedSlice(c, m["allowed"])
	res.Denied = flattenFirewallDeniedSlice(c, m["denied"])
	res.DestinationRanges = dcl.FlattenStringSlice(m["destinationRanges"])
	res.SourceRanges = dcl.FlattenStringSlice(m["sourceRanges"])
	res.SourceServiceAccounts = dcl.FlattenStringSlice(m["sourceServiceAccounts"])
	res.SourceTags = dcl.FlattenStringSlice(m["sourceTags"])
	res.TargetServiceAccounts = dcl.FlattenStringSlice(m["targetServiceAccounts"])
	res.TargetTags = dcl.FlattenStringSlice(m["targetTags"])

	return res
}

// expandFirewallLogConfigMap expands the contents of FirewallLogConfig into a JSON
// request object.
func expandFirewallLogConfigMap(c *Client, f map[string]FirewallLogConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFirewallLogConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFirewallLogConfigSlice expands the contents of FirewallLogConfig into a JSON
// request object.
func expandFirewallLogConfigSlice(c *Client, f []FirewallLogConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFirewallLogConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFirewallLogConfigMap flattens the contents of FirewallLogConfig from a JSON
// response object.
func flattenFirewallLogConfigMap(c *Client, i interface{}) map[string]FirewallLogConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirewallLogConfig{}
	}

	if len(a) == 0 {
		return map[string]FirewallLogConfig{}
	}

	items := make(map[string]FirewallLogConfig)
	for k, item := range a {
		items[k] = *flattenFirewallLogConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFirewallLogConfigSlice flattens the contents of FirewallLogConfig from a JSON
// response object.
func flattenFirewallLogConfigSlice(c *Client, i interface{}) []FirewallLogConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallLogConfig{}
	}

	if len(a) == 0 {
		return []FirewallLogConfig{}
	}

	items := make([]FirewallLogConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallLogConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandFirewallLogConfig expands an instance of FirewallLogConfig into a JSON
// request object.
func expandFirewallLogConfig(c *Client, f *FirewallLogConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enable; !dcl.IsEmptyValueIndirect(v) {
		m["enable"] = v
	}

	return m, nil
}

// flattenFirewallLogConfig flattens an instance of FirewallLogConfig from a JSON
// response object.
func flattenFirewallLogConfig(c *Client, i interface{}) *FirewallLogConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FirewallLogConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFirewallLogConfig
	}
	r.Enable = dcl.FlattenBool(m["enable"])

	return r
}

// expandFirewallAllowedMap expands the contents of FirewallAllowed into a JSON
// request object.
func expandFirewallAllowedMap(c *Client, f map[string]FirewallAllowed) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFirewallAllowed(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFirewallAllowedSlice expands the contents of FirewallAllowed into a JSON
// request object.
func expandFirewallAllowedSlice(c *Client, f []FirewallAllowed) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFirewallAllowed(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFirewallAllowedMap flattens the contents of FirewallAllowed from a JSON
// response object.
func flattenFirewallAllowedMap(c *Client, i interface{}) map[string]FirewallAllowed {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirewallAllowed{}
	}

	if len(a) == 0 {
		return map[string]FirewallAllowed{}
	}

	items := make(map[string]FirewallAllowed)
	for k, item := range a {
		items[k] = *flattenFirewallAllowed(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFirewallAllowedSlice flattens the contents of FirewallAllowed from a JSON
// response object.
func flattenFirewallAllowedSlice(c *Client, i interface{}) []FirewallAllowed {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallAllowed{}
	}

	if len(a) == 0 {
		return []FirewallAllowed{}
	}

	items := make([]FirewallAllowed, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallAllowed(c, item.(map[string]interface{})))
	}

	return items
}

// expandFirewallAllowed expands an instance of FirewallAllowed into a JSON
// request object.
func expandFirewallAllowed(c *Client, f *FirewallAllowed) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPProtocol; !dcl.IsEmptyValueIndirect(v) {
		m["IPProtocol"] = v
	}
	if v := f.Ports; v != nil {
		m["ports"] = v
	}
	if v := f.IPProtocolAlt; v != nil {
		m["ipProtocolAlt"] = v
	}

	return m, nil
}

// flattenFirewallAllowed flattens an instance of FirewallAllowed from a JSON
// response object.
func flattenFirewallAllowed(c *Client, i interface{}) *FirewallAllowed {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FirewallAllowed{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFirewallAllowed
	}
	r.IPProtocol = dcl.FlattenString(m["IPProtocol"])
	r.Ports = dcl.FlattenStringSlice(m["ports"])
	r.IPProtocolAlt = dcl.FlattenStringSlice(m["ipProtocolAlt"])

	return r
}

// expandFirewallDeniedMap expands the contents of FirewallDenied into a JSON
// request object.
func expandFirewallDeniedMap(c *Client, f map[string]FirewallDenied) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFirewallDenied(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFirewallDeniedSlice expands the contents of FirewallDenied into a JSON
// request object.
func expandFirewallDeniedSlice(c *Client, f []FirewallDenied) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFirewallDenied(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFirewallDeniedMap flattens the contents of FirewallDenied from a JSON
// response object.
func flattenFirewallDeniedMap(c *Client, i interface{}) map[string]FirewallDenied {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirewallDenied{}
	}

	if len(a) == 0 {
		return map[string]FirewallDenied{}
	}

	items := make(map[string]FirewallDenied)
	for k, item := range a {
		items[k] = *flattenFirewallDenied(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFirewallDeniedSlice flattens the contents of FirewallDenied from a JSON
// response object.
func flattenFirewallDeniedSlice(c *Client, i interface{}) []FirewallDenied {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallDenied{}
	}

	if len(a) == 0 {
		return []FirewallDenied{}
	}

	items := make([]FirewallDenied, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallDenied(c, item.(map[string]interface{})))
	}

	return items
}

// expandFirewallDenied expands an instance of FirewallDenied into a JSON
// request object.
func expandFirewallDenied(c *Client, f *FirewallDenied) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPProtocol; !dcl.IsEmptyValueIndirect(v) {
		m["IPProtocol"] = v
	}
	if v := f.Ports; v != nil {
		m["ports"] = v
	}
	if v := f.IPProtocolAlt; v != nil {
		m["ipProtocolAlt"] = v
	}

	return m, nil
}

// flattenFirewallDenied flattens an instance of FirewallDenied from a JSON
// response object.
func flattenFirewallDenied(c *Client, i interface{}) *FirewallDenied {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FirewallDenied{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFirewallDenied
	}
	r.IPProtocol = dcl.FlattenString(m["IPProtocol"])
	r.Ports = dcl.FlattenStringSlice(m["ports"])
	r.IPProtocolAlt = dcl.FlattenStringSlice(m["ipProtocolAlt"])

	return r
}

// flattenFirewallDirectionEnumMap flattens the contents of FirewallDirectionEnum from a JSON
// response object.
func flattenFirewallDirectionEnumMap(c *Client, i interface{}) map[string]FirewallDirectionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirewallDirectionEnum{}
	}

	if len(a) == 0 {
		return map[string]FirewallDirectionEnum{}
	}

	items := make(map[string]FirewallDirectionEnum)
	for k, item := range a {
		items[k] = *flattenFirewallDirectionEnum(item.(interface{}))
	}

	return items
}

// flattenFirewallDirectionEnumSlice flattens the contents of FirewallDirectionEnum from a JSON
// response object.
func flattenFirewallDirectionEnumSlice(c *Client, i interface{}) []FirewallDirectionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FirewallDirectionEnum{}
	}

	if len(a) == 0 {
		return []FirewallDirectionEnum{}
	}

	items := make([]FirewallDirectionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirewallDirectionEnum(item.(interface{})))
	}

	return items
}

// flattenFirewallDirectionEnum asserts that an interface is a string, and returns a
// pointer to a *FirewallDirectionEnum with the same value as that string.
func flattenFirewallDirectionEnum(i interface{}) *FirewallDirectionEnum {
	s, ok := i.(string)
	if !ok {
		return FirewallDirectionEnumRef("")
	}

	return FirewallDirectionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Firewall) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFirewall(b, c)
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

type firewallDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         firewallApiOperation
}

func convertFieldDiffsToFirewallDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]firewallDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []firewallDiff
	// For each operation name, create a firewallDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := firewallDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFirewallApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFirewallApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (firewallApiOperation, error) {
	switch opName {

	case "updateFirewallUpdateOperation":
		return &updateFirewallUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
