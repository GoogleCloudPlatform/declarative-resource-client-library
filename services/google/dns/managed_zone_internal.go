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
package dns

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
)

func (r *ManagedZone) validate() error {

	if err := dcl.Required(r, "description"); err != nil {
		return err
	}
	if err := dcl.Required(r, "dnsName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DnssecConfig) {
		if err := r.DnssecConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PrivateVisibilityConfig) {
		if err := r.PrivateVisibilityConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ForwardingConfig) {
		if err := r.ForwardingConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PeeringConfig) {
		if err := r.PeeringConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ManagedZoneDnssecConfig) validate() error {
	return nil
}
func (r *ManagedZoneDnssecConfigDefaultKeySpecs) validate() error {
	return nil
}
func (r *ManagedZonePrivateVisibilityConfig) validate() error {
	return nil
}
func (r *ManagedZonePrivateVisibilityConfigNetworks) validate() error {
	return nil
}
func (r *ManagedZoneForwardingConfig) validate() error {
	return nil
}
func (r *ManagedZoneForwardingConfigTargetNameServers) validate() error {
	return nil
}
func (r *ManagedZonePeeringConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TargetNetwork) {
		if err := r.TargetNetwork.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ManagedZonePeeringConfigTargetNetwork) validate() error {
	return nil
}

func managedZoneGetURL(userBasePath string, r *ManagedZone) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/managedZones/{{name}}", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil
}

func managedZoneListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/managedZones", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil

}

func managedZoneCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/managedZones", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil

}

func managedZoneDeleteURL(userBasePath string, r *ManagedZone) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/managedZones/{{name}}", "https://www.googleapis.com/dns/v1/", userBasePath, params), nil
}

// managedZoneApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type managedZoneApiOperation interface {
	do(context.Context, *ManagedZone, *Client) error
}

// newUpdateManagedZoneUpdateRequest creates a request for an
// ManagedZone resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateManagedZoneUpdateRequest(ctx context.Context, f *ManagedZone, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandManagedZonePrivateVisibilityConfig(c, f.PrivateVisibilityConfig); err != nil {
		return nil, fmt.Errorf("error expanding PrivateVisibilityConfig into privateVisibilityConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["privateVisibilityConfig"] = v
	}
	if v, err := expandManagedZoneForwardingConfig(c, f.ForwardingConfig); err != nil {
		return nil, fmt.Errorf("error expanding ForwardingConfig into forwardingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["forwardingConfig"] = v
	}
	if v, err := expandManagedZonePeeringConfig(c, f.PeeringConfig); err != nil {
		return nil, fmt.Errorf("error expanding PeeringConfig into peeringConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["peeringConfig"] = v
	}
	return req, nil
}

// marshalUpdateManagedZoneUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateManagedZoneUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateManagedZoneUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateManagedZoneUpdateOperation) do(ctx context.Context, r *ManagedZone, c *Client) error {
	_, err := c.GetManagedZone(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateManagedZoneUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateManagedZoneUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listManagedZoneRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := managedZoneListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ManagedZoneMaxPage {
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

type listManagedZoneOperation struct {
	ManagedZones []map[string]interface{} `json:"managedZones"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listManagedZone(ctx context.Context, project, pageToken string, pageSize int32) ([]*ManagedZone, string, error) {
	b, err := c.listManagedZoneRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listManagedZoneOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ManagedZone
	for _, v := range m.ManagedZones {
		res := flattenManagedZone(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllManagedZone(ctx context.Context, f func(*ManagedZone) bool, resources []*ManagedZone) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteManagedZone(ctx, res)
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

type deleteManagedZoneOperation struct{}

func (op *deleteManagedZoneOperation) do(ctx context.Context, r *ManagedZone, c *Client) error {

	_, err := c.GetManagedZone(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ManagedZone not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetManagedZone checking for existence. error: %v", err)
		return err
	}

	u, err := managedZoneDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete ManagedZone: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetManagedZone(ctx, r.urlNormalized())
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
type createManagedZoneOperation struct {
	response map[string]interface{}
}

func (op *createManagedZoneOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createManagedZoneOperation) do(ctx context.Context, r *ManagedZone, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := managedZoneCreateURL(c.Config.BasePath, project)

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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetManagedZone(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getManagedZoneRaw(ctx context.Context, r *ManagedZone) ([]byte, error) {

	u, err := managedZoneGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) managedZoneDiffsForRawDesired(ctx context.Context, rawDesired *ManagedZone, opts ...dcl.ApplyOption) (initial, desired *ManagedZone, diffs []managedZoneDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ManagedZone
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ManagedZone); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ManagedZone, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetManagedZone(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ManagedZone resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ManagedZone resource: %v", err)
		}
		c.Config.Logger.Info("Found that ManagedZone resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeManagedZoneDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for ManagedZone: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ManagedZone: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeManagedZoneInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ManagedZone: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeManagedZoneDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ManagedZone: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffManagedZone(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeManagedZoneInitialState(rawInitial, rawDesired *ManagedZone) (*ManagedZone, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeManagedZoneDesiredState(rawDesired, rawInitial *ManagedZone, opts ...dcl.ApplyOption) (*ManagedZone, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DnssecConfig = canonicalizeManagedZoneDnssecConfig(rawDesired.DnssecConfig, nil, opts...)
		rawDesired.PrivateVisibilityConfig = canonicalizeManagedZonePrivateVisibilityConfig(rawDesired.PrivateVisibilityConfig, nil, opts...)
		rawDesired.ForwardingConfig = canonicalizeManagedZoneForwardingConfig(rawDesired.ForwardingConfig, nil, opts...)
		rawDesired.PeeringConfig = canonicalizeManagedZonePeeringConfig(rawDesired.PeeringConfig, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.DnsName, rawInitial.DnsName) {
		rawDesired.DnsName = rawInitial.DnsName
	}
	rawDesired.DnssecConfig = canonicalizeManagedZoneDnssecConfig(rawDesired.DnssecConfig, rawInitial.DnssecConfig, opts...)
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.NameServers) {
		rawDesired.NameServers = rawInitial.NameServers
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.Visibility) {
		rawDesired.Visibility = rawInitial.Visibility
	}
	rawDesired.PrivateVisibilityConfig = canonicalizeManagedZonePrivateVisibilityConfig(rawDesired.PrivateVisibilityConfig, rawInitial.PrivateVisibilityConfig, opts...)
	rawDesired.ForwardingConfig = canonicalizeManagedZoneForwardingConfig(rawDesired.ForwardingConfig, rawInitial.ForwardingConfig, opts...)
	if dcl.BoolCanonicalize(rawDesired.ReverseLookup, rawInitial.ReverseLookup) {
		rawDesired.ReverseLookup = rawInitial.ReverseLookup
	}
	rawDesired.PeeringConfig = canonicalizeManagedZonePeeringConfig(rawDesired.PeeringConfig, rawInitial.PeeringConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeManagedZoneNewState(c *Client, rawNew, rawDesired *ManagedZone) (*ManagedZone, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DnsName) && dcl.IsEmptyValueIndirect(rawDesired.DnsName) {
		rawNew.DnsName = rawDesired.DnsName
	} else {
		if dcl.StringCanonicalize(rawDesired.DnsName, rawNew.DnsName) {
			rawNew.DnsName = rawDesired.DnsName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DnssecConfig) && dcl.IsEmptyValueIndirect(rawDesired.DnssecConfig) {
		rawNew.DnssecConfig = rawDesired.DnssecConfig
	} else {
		rawNew.DnssecConfig = canonicalizeNewManagedZoneDnssecConfig(c, rawDesired.DnssecConfig, rawNew.DnssecConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NameServers) && dcl.IsEmptyValueIndirect(rawDesired.NameServers) {
		rawNew.NameServers = rawDesired.NameServers
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Visibility) && dcl.IsEmptyValueIndirect(rawDesired.Visibility) {
		rawNew.Visibility = rawDesired.Visibility
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateVisibilityConfig) && dcl.IsEmptyValueIndirect(rawDesired.PrivateVisibilityConfig) {
		rawNew.PrivateVisibilityConfig = rawDesired.PrivateVisibilityConfig
	} else {
		rawNew.PrivateVisibilityConfig = canonicalizeNewManagedZonePrivateVisibilityConfig(c, rawDesired.PrivateVisibilityConfig, rawNew.PrivateVisibilityConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ForwardingConfig) && dcl.IsEmptyValueIndirect(rawDesired.ForwardingConfig) {
		rawNew.ForwardingConfig = rawDesired.ForwardingConfig
	} else {
		rawNew.ForwardingConfig = canonicalizeNewManagedZoneForwardingConfig(c, rawDesired.ForwardingConfig, rawNew.ForwardingConfig)
	}

	rawNew.ReverseLookup = rawDesired.ReverseLookup

	if dcl.IsEmptyValueIndirect(rawNew.PeeringConfig) && dcl.IsEmptyValueIndirect(rawDesired.PeeringConfig) {
		rawNew.PeeringConfig = rawDesired.PeeringConfig
	} else {
		rawNew.PeeringConfig = canonicalizeNewManagedZonePeeringConfig(c, rawDesired.PeeringConfig, rawNew.PeeringConfig)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeManagedZoneDnssecConfig(des, initial *ManagedZoneDnssecConfig, opts ...dcl.ApplyOption) *ManagedZoneDnssecConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.NonExistence) {
		des.NonExistence = initial.NonExistence
	}
	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}
	if dcl.IsZeroValue(des.DefaultKeySpecs) {
		des.DefaultKeySpecs = initial.DefaultKeySpecs
	}

	return des
}

func canonicalizeNewManagedZoneDnssecConfig(c *Client, des, nw *ManagedZoneDnssecConfig) *ManagedZoneDnssecConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	nw.DefaultKeySpecs = canonicalizeNewManagedZoneDnssecConfigDefaultKeySpecsSlice(c, des.DefaultKeySpecs, nw.DefaultKeySpecs)

	return nw
}

func canonicalizeNewManagedZoneDnssecConfigSet(c *Client, des, nw []ManagedZoneDnssecConfig) []ManagedZoneDnssecConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZoneDnssecConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZoneDnssecConfig(c, &d, &n) {
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

func canonicalizeNewManagedZoneDnssecConfigSlice(c *Client, des, nw []ManagedZoneDnssecConfig) []ManagedZoneDnssecConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZoneDnssecConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZoneDnssecConfig(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZoneDnssecConfigDefaultKeySpecs(des, initial *ManagedZoneDnssecConfigDefaultKeySpecs, opts ...dcl.ApplyOption) *ManagedZoneDnssecConfigDefaultKeySpecs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Algorithm) {
		des.Algorithm = initial.Algorithm
	}
	if dcl.IsZeroValue(des.KeyLength) {
		des.KeyLength = initial.KeyLength
	}
	if dcl.IsZeroValue(des.KeyType) {
		des.KeyType = initial.KeyType
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewManagedZoneDnssecConfigDefaultKeySpecs(c *Client, des, nw *ManagedZoneDnssecConfigDefaultKeySpecs) *ManagedZoneDnssecConfigDefaultKeySpecs {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewManagedZoneDnssecConfigDefaultKeySpecsSet(c *Client, des, nw []ManagedZoneDnssecConfigDefaultKeySpecs) []ManagedZoneDnssecConfigDefaultKeySpecs {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZoneDnssecConfigDefaultKeySpecs
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZoneDnssecConfigDefaultKeySpecs(c, &d, &n) {
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

func canonicalizeNewManagedZoneDnssecConfigDefaultKeySpecsSlice(c *Client, des, nw []ManagedZoneDnssecConfigDefaultKeySpecs) []ManagedZoneDnssecConfigDefaultKeySpecs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZoneDnssecConfigDefaultKeySpecs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZoneDnssecConfigDefaultKeySpecs(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZonePrivateVisibilityConfig(des, initial *ManagedZonePrivateVisibilityConfig, opts ...dcl.ApplyOption) *ManagedZonePrivateVisibilityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Networks) {
		des.Networks = initial.Networks
	}

	return des
}

func canonicalizeNewManagedZonePrivateVisibilityConfig(c *Client, des, nw *ManagedZonePrivateVisibilityConfig) *ManagedZonePrivateVisibilityConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.Networks = canonicalizeNewManagedZonePrivateVisibilityConfigNetworksSlice(c, des.Networks, nw.Networks)

	return nw
}

func canonicalizeNewManagedZonePrivateVisibilityConfigSet(c *Client, des, nw []ManagedZonePrivateVisibilityConfig) []ManagedZonePrivateVisibilityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZonePrivateVisibilityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZonePrivateVisibilityConfig(c, &d, &n) {
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

func canonicalizeNewManagedZonePrivateVisibilityConfigSlice(c *Client, des, nw []ManagedZonePrivateVisibilityConfig) []ManagedZonePrivateVisibilityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZonePrivateVisibilityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZonePrivateVisibilityConfig(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZonePrivateVisibilityConfigNetworks(des, initial *ManagedZonePrivateVisibilityConfigNetworks, opts ...dcl.ApplyOption) *ManagedZonePrivateVisibilityConfigNetworks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.NetworkUrl, initial.NetworkUrl) || dcl.IsZeroValue(des.NetworkUrl) {
		des.NetworkUrl = initial.NetworkUrl
	}

	return des
}

func canonicalizeNewManagedZonePrivateVisibilityConfigNetworks(c *Client, des, nw *ManagedZonePrivateVisibilityConfigNetworks) *ManagedZonePrivateVisibilityConfigNetworks {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.NetworkUrl, nw.NetworkUrl) {
		nw.NetworkUrl = des.NetworkUrl
	}

	return nw
}

func canonicalizeNewManagedZonePrivateVisibilityConfigNetworksSet(c *Client, des, nw []ManagedZonePrivateVisibilityConfigNetworks) []ManagedZonePrivateVisibilityConfigNetworks {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZonePrivateVisibilityConfigNetworks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZonePrivateVisibilityConfigNetworks(c, &d, &n) {
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

func canonicalizeNewManagedZonePrivateVisibilityConfigNetworksSlice(c *Client, des, nw []ManagedZonePrivateVisibilityConfigNetworks) []ManagedZonePrivateVisibilityConfigNetworks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZonePrivateVisibilityConfigNetworks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZonePrivateVisibilityConfigNetworks(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZoneForwardingConfig(des, initial *ManagedZoneForwardingConfig, opts ...dcl.ApplyOption) *ManagedZoneForwardingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.TargetNameServers) {
		des.TargetNameServers = initial.TargetNameServers
	}

	return des
}

func canonicalizeNewManagedZoneForwardingConfig(c *Client, des, nw *ManagedZoneForwardingConfig) *ManagedZoneForwardingConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.TargetNameServers = canonicalizeNewManagedZoneForwardingConfigTargetNameServersSlice(c, des.TargetNameServers, nw.TargetNameServers)

	return nw
}

func canonicalizeNewManagedZoneForwardingConfigSet(c *Client, des, nw []ManagedZoneForwardingConfig) []ManagedZoneForwardingConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZoneForwardingConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZoneForwardingConfig(c, &d, &n) {
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

func canonicalizeNewManagedZoneForwardingConfigSlice(c *Client, des, nw []ManagedZoneForwardingConfig) []ManagedZoneForwardingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZoneForwardingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZoneForwardingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZoneForwardingConfigTargetNameServers(des, initial *ManagedZoneForwardingConfigTargetNameServers, opts ...dcl.ApplyOption) *ManagedZoneForwardingConfigTargetNameServers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.IPv4Address, initial.IPv4Address) || dcl.IsZeroValue(des.IPv4Address) {
		des.IPv4Address = initial.IPv4Address
	}
	if dcl.IsZeroValue(des.ForwardingPath) {
		des.ForwardingPath = initial.ForwardingPath
	}

	return des
}

func canonicalizeNewManagedZoneForwardingConfigTargetNameServers(c *Client, des, nw *ManagedZoneForwardingConfigTargetNameServers) *ManagedZoneForwardingConfigTargetNameServers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPv4Address, nw.IPv4Address) {
		nw.IPv4Address = des.IPv4Address
	}

	return nw
}

func canonicalizeNewManagedZoneForwardingConfigTargetNameServersSet(c *Client, des, nw []ManagedZoneForwardingConfigTargetNameServers) []ManagedZoneForwardingConfigTargetNameServers {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZoneForwardingConfigTargetNameServers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZoneForwardingConfigTargetNameServers(c, &d, &n) {
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

func canonicalizeNewManagedZoneForwardingConfigTargetNameServersSlice(c *Client, des, nw []ManagedZoneForwardingConfigTargetNameServers) []ManagedZoneForwardingConfigTargetNameServers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZoneForwardingConfigTargetNameServers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZoneForwardingConfigTargetNameServers(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZonePeeringConfig(des, initial *ManagedZonePeeringConfig, opts ...dcl.ApplyOption) *ManagedZonePeeringConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.TargetNetwork = canonicalizeManagedZonePeeringConfigTargetNetwork(des.TargetNetwork, initial.TargetNetwork, opts...)

	return des
}

func canonicalizeNewManagedZonePeeringConfig(c *Client, des, nw *ManagedZonePeeringConfig) *ManagedZonePeeringConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.TargetNetwork = canonicalizeNewManagedZonePeeringConfigTargetNetwork(c, des.TargetNetwork, nw.TargetNetwork)

	return nw
}

func canonicalizeNewManagedZonePeeringConfigSet(c *Client, des, nw []ManagedZonePeeringConfig) []ManagedZonePeeringConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZonePeeringConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZonePeeringConfig(c, &d, &n) {
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

func canonicalizeNewManagedZonePeeringConfigSlice(c *Client, des, nw []ManagedZonePeeringConfig) []ManagedZonePeeringConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZonePeeringConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZonePeeringConfig(c, &d, &n))
	}

	return items
}

func canonicalizeManagedZonePeeringConfigTargetNetwork(des, initial *ManagedZonePeeringConfigTargetNetwork, opts ...dcl.ApplyOption) *ManagedZonePeeringConfigTargetNetwork {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.NetworkUrl, initial.NetworkUrl) || dcl.IsZeroValue(des.NetworkUrl) {
		des.NetworkUrl = initial.NetworkUrl
	}

	return des
}

func canonicalizeNewManagedZonePeeringConfigTargetNetwork(c *Client, des, nw *ManagedZonePeeringConfigTargetNetwork) *ManagedZonePeeringConfigTargetNetwork {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.NetworkUrl, nw.NetworkUrl) {
		nw.NetworkUrl = des.NetworkUrl
	}

	return nw
}

func canonicalizeNewManagedZonePeeringConfigTargetNetworkSet(c *Client, des, nw []ManagedZonePeeringConfigTargetNetwork) []ManagedZonePeeringConfigTargetNetwork {
	if des == nil {
		return nw
	}
	var reorderedNew []ManagedZonePeeringConfigTargetNetwork
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareManagedZonePeeringConfigTargetNetwork(c, &d, &n) {
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

func canonicalizeNewManagedZonePeeringConfigTargetNetworkSlice(c *Client, des, nw []ManagedZonePeeringConfigTargetNetwork) []ManagedZonePeeringConfigTargetNetwork {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []ManagedZonePeeringConfigTargetNetwork
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewManagedZonePeeringConfigTargetNetwork(c, &d, &n))
	}

	return items
}

type managedZoneDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         managedZoneApiOperation
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
func diffManagedZone(c *Client, desired, actual *ManagedZone, opts ...dcl.ApplyOption) ([]managedZoneDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []managedZoneDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "description"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{
			UpdateOp: &updateManagedZoneUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.DnsName, actual.DnsName, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "dns_name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.NameServers, actual.NameServers, dcl.Info{Ignore: false, OutputOnly: true, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name_servers"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "labels"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{
			UpdateOp: &updateManagedZoneUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.Visibility, actual.Visibility, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "EnumType", FieldName: "visibility"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.ReverseLookup, actual.ReverseLookup, dcl.Info{Ignore: true, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "reverse_lookup"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "project"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, managedZoneDiff{RequiresRecreate: true, Diffs: ds})
	}

	if compareManagedZoneDnssecConfig(c, desired.DnssecConfig, actual.DnssecConfig) {
		c.Config.Logger.Infof("Detected diff in DnssecConfig.\nDESIRED: %v\nACTUAL: %v", desired.DnssecConfig, actual.DnssecConfig)
		diffs = append(diffs, managedZoneDiff{
			RequiresRecreate: true,
			FieldName:        "DnssecConfig",
		})
	}
	if compareManagedZonePrivateVisibilityConfig(c, desired.PrivateVisibilityConfig, actual.PrivateVisibilityConfig) {
		c.Config.Logger.Infof("Detected diff in PrivateVisibilityConfig.\nDESIRED: %v\nACTUAL: %v", desired.PrivateVisibilityConfig, actual.PrivateVisibilityConfig)

		diffs = append(diffs, managedZoneDiff{
			UpdateOp:  &updateManagedZoneUpdateOperation{},
			FieldName: "PrivateVisibilityConfig",
		})

	}
	if compareManagedZoneForwardingConfig(c, desired.ForwardingConfig, actual.ForwardingConfig) {
		c.Config.Logger.Infof("Detected diff in ForwardingConfig.\nDESIRED: %v\nACTUAL: %v", desired.ForwardingConfig, actual.ForwardingConfig)

		diffs = append(diffs, managedZoneDiff{
			UpdateOp:  &updateManagedZoneUpdateOperation{},
			FieldName: "ForwardingConfig",
		})

	}
	if compareManagedZonePeeringConfig(c, desired.PeeringConfig, actual.PeeringConfig) {
		c.Config.Logger.Infof("Detected diff in PeeringConfig.\nDESIRED: %v\nACTUAL: %v", desired.PeeringConfig, actual.PeeringConfig)

		diffs = append(diffs, managedZoneDiff{
			UpdateOp:  &updateManagedZoneUpdateOperation{},
			FieldName: "PeeringConfig",
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
	var deduped []managedZoneDiff
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
func compareManagedZoneDnssecConfig(c *Client, desired, actual *ManagedZoneDnssecConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.NonExistence, actual.NonExistence) && !dcl.IsZeroValue(desired.NonExistence) {
		c.Config.Logger.Infof("Diff in NonExistence.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NonExistence), dcl.SprintResource(actual.NonExistence))
		return true
	}
	if !reflect.DeepEqual(desired.State, actual.State) && !dcl.IsZeroValue(desired.State) {
		c.Config.Logger.Infof("Diff in State.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.State), dcl.SprintResource(actual.State))
		return true
	}
	if compareManagedZoneDnssecConfigDefaultKeySpecsSlice(c, desired.DefaultKeySpecs, actual.DefaultKeySpecs) && !dcl.IsZeroValue(desired.DefaultKeySpecs) {
		c.Config.Logger.Infof("Diff in DefaultKeySpecs.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DefaultKeySpecs), dcl.SprintResource(actual.DefaultKeySpecs))
		return true
	}
	return false
}

func compareManagedZoneDnssecConfigSlice(c *Client, desired, actual []ManagedZoneDnssecConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneDnssecConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigMap(c *Client, desired, actual map[string]ManagedZoneDnssecConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZoneDnssecConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigDefaultKeySpecs(c *Client, desired, actual *ManagedZoneDnssecConfigDefaultKeySpecs) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Algorithm, actual.Algorithm) && !dcl.IsZeroValue(desired.Algorithm) {
		c.Config.Logger.Infof("Diff in Algorithm.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Algorithm), dcl.SprintResource(actual.Algorithm))
		return true
	}
	if !reflect.DeepEqual(desired.KeyLength, actual.KeyLength) && !dcl.IsZeroValue(desired.KeyLength) {
		c.Config.Logger.Infof("Diff in KeyLength.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KeyLength), dcl.SprintResource(actual.KeyLength))
		return true
	}
	if !reflect.DeepEqual(desired.KeyType, actual.KeyType) && !dcl.IsZeroValue(desired.KeyType) {
		c.Config.Logger.Infof("Diff in KeyType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KeyType), dcl.SprintResource(actual.KeyType))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareManagedZoneDnssecConfigDefaultKeySpecsSlice(c *Client, desired, actual []ManagedZoneDnssecConfigDefaultKeySpecs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfigDefaultKeySpecs, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneDnssecConfigDefaultKeySpecs(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigDefaultKeySpecs, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigDefaultKeySpecsMap(c *Client, desired, actual map[string]ManagedZoneDnssecConfigDefaultKeySpecs) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfigDefaultKeySpecs, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigDefaultKeySpecs, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZoneDnssecConfigDefaultKeySpecs(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigDefaultKeySpecs, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZonePrivateVisibilityConfig(c *Client, desired, actual *ManagedZonePrivateVisibilityConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareManagedZonePrivateVisibilityConfigNetworksSlice(c, desired.Networks, actual.Networks) && !dcl.IsZeroValue(desired.Networks) {
		c.Config.Logger.Infof("Diff in Networks.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Networks), dcl.SprintResource(actual.Networks))
		return true
	}
	return false
}

func compareManagedZonePrivateVisibilityConfigSlice(c *Client, desired, actual []ManagedZonePrivateVisibilityConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePrivateVisibilityConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZonePrivateVisibilityConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZonePrivateVisibilityConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZonePrivateVisibilityConfigMap(c *Client, desired, actual map[string]ManagedZonePrivateVisibilityConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePrivateVisibilityConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZonePrivateVisibilityConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZonePrivateVisibilityConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZonePrivateVisibilityConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZonePrivateVisibilityConfigNetworks(c *Client, desired, actual *ManagedZonePrivateVisibilityConfigNetworks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.NameToSelfLink(desired.NetworkUrl, actual.NetworkUrl) && !dcl.IsZeroValue(desired.NetworkUrl) {
		c.Config.Logger.Infof("Diff in NetworkUrl.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkUrl), dcl.SprintResource(actual.NetworkUrl))
		return true
	}
	return false
}

func compareManagedZonePrivateVisibilityConfigNetworksSlice(c *Client, desired, actual []ManagedZonePrivateVisibilityConfigNetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePrivateVisibilityConfigNetworks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZonePrivateVisibilityConfigNetworks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZonePrivateVisibilityConfigNetworks, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZonePrivateVisibilityConfigNetworksMap(c *Client, desired, actual map[string]ManagedZonePrivateVisibilityConfigNetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePrivateVisibilityConfigNetworks, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZonePrivateVisibilityConfigNetworks, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZonePrivateVisibilityConfigNetworks(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZonePrivateVisibilityConfigNetworks, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZoneForwardingConfig(c *Client, desired, actual *ManagedZoneForwardingConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareManagedZoneForwardingConfigTargetNameServersSlice(c, desired.TargetNameServers, actual.TargetNameServers) && !dcl.IsZeroValue(desired.TargetNameServers) {
		c.Config.Logger.Infof("Diff in TargetNameServers.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetNameServers), dcl.SprintResource(actual.TargetNameServers))
		return true
	}
	return false
}

func compareManagedZoneForwardingConfigSlice(c *Client, desired, actual []ManagedZoneForwardingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneForwardingConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneForwardingConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneForwardingConfigMap(c *Client, desired, actual map[string]ManagedZoneForwardingConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneForwardingConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZoneForwardingConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZoneForwardingConfigTargetNameServers(c *Client, desired, actual *ManagedZoneForwardingConfigTargetNameServers) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.IPv4Address, actual.IPv4Address) && !dcl.IsZeroValue(desired.IPv4Address) {
		c.Config.Logger.Infof("Diff in IPv4Address.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPv4Address), dcl.SprintResource(actual.IPv4Address))
		return true
	}
	if !reflect.DeepEqual(desired.ForwardingPath, actual.ForwardingPath) && !dcl.IsZeroValue(desired.ForwardingPath) {
		c.Config.Logger.Infof("Diff in ForwardingPath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ForwardingPath), dcl.SprintResource(actual.ForwardingPath))
		return true
	}
	return false
}

func compareManagedZoneForwardingConfigTargetNameServersSlice(c *Client, desired, actual []ManagedZoneForwardingConfigTargetNameServers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneForwardingConfigTargetNameServers, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneForwardingConfigTargetNameServers(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfigTargetNameServers, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneForwardingConfigTargetNameServersMap(c *Client, desired, actual map[string]ManagedZoneForwardingConfigTargetNameServers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneForwardingConfigTargetNameServers, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfigTargetNameServers, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZoneForwardingConfigTargetNameServers(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfigTargetNameServers, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZonePeeringConfig(c *Client, desired, actual *ManagedZonePeeringConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareManagedZonePeeringConfigTargetNetwork(c, desired.TargetNetwork, actual.TargetNetwork) && !dcl.IsZeroValue(desired.TargetNetwork) {
		c.Config.Logger.Infof("Diff in TargetNetwork.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetNetwork), dcl.SprintResource(actual.TargetNetwork))
		return true
	}
	return false
}

func compareManagedZonePeeringConfigSlice(c *Client, desired, actual []ManagedZonePeeringConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePeeringConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZonePeeringConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZonePeeringConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZonePeeringConfigMap(c *Client, desired, actual map[string]ManagedZonePeeringConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePeeringConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZonePeeringConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZonePeeringConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZonePeeringConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZonePeeringConfigTargetNetwork(c *Client, desired, actual *ManagedZonePeeringConfigTargetNetwork) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.NameToSelfLink(desired.NetworkUrl, actual.NetworkUrl) && !dcl.IsZeroValue(desired.NetworkUrl) {
		c.Config.Logger.Infof("Diff in NetworkUrl.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.NetworkUrl), dcl.SprintResource(actual.NetworkUrl))
		return true
	}
	return false
}

func compareManagedZonePeeringConfigTargetNetworkSlice(c *Client, desired, actual []ManagedZonePeeringConfigTargetNetwork) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePeeringConfigTargetNetwork, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZonePeeringConfigTargetNetwork(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZonePeeringConfigTargetNetwork, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZonePeeringConfigTargetNetworkMap(c *Client, desired, actual map[string]ManagedZonePeeringConfigTargetNetwork) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZonePeeringConfigTargetNetwork, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in ManagedZonePeeringConfigTargetNetwork, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareManagedZonePeeringConfigTargetNetwork(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in ManagedZonePeeringConfigTargetNetwork, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigNonExistenceEnumSlice(c *Client, desired, actual []ManagedZoneDnssecConfigNonExistenceEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfigNonExistenceEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneDnssecConfigNonExistenceEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigNonExistenceEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigNonExistenceEnum(c *Client, desired, actual *ManagedZoneDnssecConfigNonExistenceEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareManagedZoneDnssecConfigStateEnumSlice(c *Client, desired, actual []ManagedZoneDnssecConfigStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfigStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneDnssecConfigStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigStateEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigStateEnum(c *Client, desired, actual *ManagedZoneDnssecConfigStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumSlice(c *Client, desired, actual []ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(c *Client, desired, actual *ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumSlice(c *Client, desired, actual []ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(c *Client, desired, actual *ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareManagedZoneVisibilityEnumSlice(c *Client, desired, actual []ManagedZoneVisibilityEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneVisibilityEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneVisibilityEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneVisibilityEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneVisibilityEnum(c *Client, desired, actual *ManagedZoneVisibilityEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareManagedZoneForwardingConfigTargetNameServersForwardingPathEnumSlice(c *Client, desired, actual []ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(c *Client, desired, actual *ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *ManagedZone) urlNormalized() *ManagedZone {
	normalized := dcl.Copy(*r).(ManagedZone)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DnsName = dcl.SelfLinkToName(r.DnsName)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *ManagedZone) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *ManagedZone) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *ManagedZone) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *ManagedZone) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/managedZones/{{name}}", "https://www.googleapis.com/dns/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ManagedZone resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ManagedZone) marshal(c *Client) ([]byte, error) {
	m, err := expandManagedZone(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ManagedZone: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalManagedZone decodes JSON responses into the ManagedZone resource schema.
func unmarshalManagedZone(b []byte, c *Client) (*ManagedZone, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapManagedZone(m, c)
}

func unmarshalMapManagedZone(m map[string]interface{}, c *Client) (*ManagedZone, error) {

	return flattenManagedZone(c, m), nil
}

// expandManagedZone expands ManagedZone into a JSON request object.
func expandManagedZone(c *Client, f *ManagedZone) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.DnsName; !dcl.IsEmptyValueIndirect(v) {
		m["dnsName"] = v
	}
	if v, err := expandManagedZoneDnssecConfig(c, f.DnssecConfig); err != nil {
		return nil, fmt.Errorf("error expanding DnssecConfig into dnssecConfig: %w", err)
	} else if v != nil {
		m["dnssecConfig"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.NameServers; !dcl.IsEmptyValueIndirect(v) {
		m["nameServers"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Visibility; !dcl.IsEmptyValueIndirect(v) {
		m["visibility"] = v
	}
	if v, err := expandManagedZonePrivateVisibilityConfig(c, f.PrivateVisibilityConfig); err != nil {
		return nil, fmt.Errorf("error expanding PrivateVisibilityConfig into privateVisibilityConfig: %w", err)
	} else if v != nil {
		m["privateVisibilityConfig"] = v
	}
	if v, err := expandManagedZoneForwardingConfig(c, f.ForwardingConfig); err != nil {
		return nil, fmt.Errorf("error expanding ForwardingConfig into forwardingConfig: %w", err)
	} else if v != nil {
		m["forwardingConfig"] = v
	}
	if v, err := expandManagedZoneReverseLookupConfig(f, f.ReverseLookup); err != nil {
		return nil, fmt.Errorf("error expanding ReverseLookup into reverseLookupConfig: %w", err)
	} else if v != nil {
		m["reverseLookupConfig"] = v
	}
	if v, err := expandManagedZonePeeringConfig(c, f.PeeringConfig); err != nil {
		return nil, fmt.Errorf("error expanding PeeringConfig into peeringConfig: %w", err)
	} else if v != nil {
		m["peeringConfig"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenManagedZone flattens ManagedZone from a JSON request object into the
// ManagedZone type.
func flattenManagedZone(c *Client, i interface{}) *ManagedZone {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &ManagedZone{}
	r.Description = dcl.FlattenString(m["description"])
	r.DnsName = dcl.FlattenString(m["dnsName"])
	r.DnssecConfig = flattenManagedZoneDnssecConfig(c, m["dnssecConfig"])
	r.Name = dcl.FlattenString(m["name"])
	r.NameServers = dcl.FlattenStringSlice(m["nameServers"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Visibility = flattenManagedZoneVisibilityEnum(m["visibility"])
	r.PrivateVisibilityConfig = flattenManagedZonePrivateVisibilityConfig(c, m["privateVisibilityConfig"])
	r.ForwardingConfig = flattenManagedZoneForwardingConfig(c, m["forwardingConfig"])
	r.ReverseLookup = flattenManagedZoneReverseLookupConfig(m["reverseLookupConfig"])
	r.PeeringConfig = flattenManagedZonePeeringConfig(c, m["peeringConfig"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandManagedZoneDnssecConfigMap expands the contents of ManagedZoneDnssecConfig into a JSON
// request object.
func expandManagedZoneDnssecConfigMap(c *Client, f map[string]ManagedZoneDnssecConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZoneDnssecConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZoneDnssecConfigSlice expands the contents of ManagedZoneDnssecConfig into a JSON
// request object.
func expandManagedZoneDnssecConfigSlice(c *Client, f []ManagedZoneDnssecConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZoneDnssecConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZoneDnssecConfigMap flattens the contents of ManagedZoneDnssecConfig from a JSON
// response object.
func flattenManagedZoneDnssecConfigMap(c *Client, i interface{}) map[string]ManagedZoneDnssecConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZoneDnssecConfig{}
	}

	if len(a) == 0 {
		return map[string]ManagedZoneDnssecConfig{}
	}

	items := make(map[string]ManagedZoneDnssecConfig)
	for k, item := range a {
		items[k] = *flattenManagedZoneDnssecConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZoneDnssecConfigSlice flattens the contents of ManagedZoneDnssecConfig from a JSON
// response object.
func flattenManagedZoneDnssecConfigSlice(c *Client, i interface{}) []ManagedZoneDnssecConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneDnssecConfig{}
	}

	if len(a) == 0 {
		return []ManagedZoneDnssecConfig{}
	}

	items := make([]ManagedZoneDnssecConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneDnssecConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZoneDnssecConfig expands an instance of ManagedZoneDnssecConfig into a JSON
// request object.
func expandManagedZoneDnssecConfig(c *Client, f *ManagedZoneDnssecConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.NonExistence; !dcl.IsEmptyValueIndirect(v) {
		m["nonExistence"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandManagedZoneDnssecConfigDefaultKeySpecsSlice(c, f.DefaultKeySpecs); err != nil {
		return nil, fmt.Errorf("error expanding DefaultKeySpecs into defaultKeySpecs: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultKeySpecs"] = v
	}

	return m, nil
}

// flattenManagedZoneDnssecConfig flattens an instance of ManagedZoneDnssecConfig from a JSON
// response object.
func flattenManagedZoneDnssecConfig(c *Client, i interface{}) *ManagedZoneDnssecConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZoneDnssecConfig{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.NonExistence = flattenManagedZoneDnssecConfigNonExistenceEnum(m["nonExistence"])
	r.State = flattenManagedZoneDnssecConfigStateEnum(m["state"])
	r.DefaultKeySpecs = flattenManagedZoneDnssecConfigDefaultKeySpecsSlice(c, m["defaultKeySpecs"])

	return r
}

// expandManagedZoneDnssecConfigDefaultKeySpecsMap expands the contents of ManagedZoneDnssecConfigDefaultKeySpecs into a JSON
// request object.
func expandManagedZoneDnssecConfigDefaultKeySpecsMap(c *Client, f map[string]ManagedZoneDnssecConfigDefaultKeySpecs) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZoneDnssecConfigDefaultKeySpecs(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZoneDnssecConfigDefaultKeySpecsSlice expands the contents of ManagedZoneDnssecConfigDefaultKeySpecs into a JSON
// request object.
func expandManagedZoneDnssecConfigDefaultKeySpecsSlice(c *Client, f []ManagedZoneDnssecConfigDefaultKeySpecs) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZoneDnssecConfigDefaultKeySpecs(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZoneDnssecConfigDefaultKeySpecsMap flattens the contents of ManagedZoneDnssecConfigDefaultKeySpecs from a JSON
// response object.
func flattenManagedZoneDnssecConfigDefaultKeySpecsMap(c *Client, i interface{}) map[string]ManagedZoneDnssecConfigDefaultKeySpecs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZoneDnssecConfigDefaultKeySpecs{}
	}

	if len(a) == 0 {
		return map[string]ManagedZoneDnssecConfigDefaultKeySpecs{}
	}

	items := make(map[string]ManagedZoneDnssecConfigDefaultKeySpecs)
	for k, item := range a {
		items[k] = *flattenManagedZoneDnssecConfigDefaultKeySpecs(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZoneDnssecConfigDefaultKeySpecsSlice flattens the contents of ManagedZoneDnssecConfigDefaultKeySpecs from a JSON
// response object.
func flattenManagedZoneDnssecConfigDefaultKeySpecsSlice(c *Client, i interface{}) []ManagedZoneDnssecConfigDefaultKeySpecs {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneDnssecConfigDefaultKeySpecs{}
	}

	if len(a) == 0 {
		return []ManagedZoneDnssecConfigDefaultKeySpecs{}
	}

	items := make([]ManagedZoneDnssecConfigDefaultKeySpecs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneDnssecConfigDefaultKeySpecs(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZoneDnssecConfigDefaultKeySpecs expands an instance of ManagedZoneDnssecConfigDefaultKeySpecs into a JSON
// request object.
func expandManagedZoneDnssecConfigDefaultKeySpecs(c *Client, f *ManagedZoneDnssecConfigDefaultKeySpecs) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Algorithm; !dcl.IsEmptyValueIndirect(v) {
		m["algorithm"] = v
	}
	if v := f.KeyLength; !dcl.IsEmptyValueIndirect(v) {
		m["keyLength"] = v
	}
	if v := f.KeyType; !dcl.IsEmptyValueIndirect(v) {
		m["keyType"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenManagedZoneDnssecConfigDefaultKeySpecs flattens an instance of ManagedZoneDnssecConfigDefaultKeySpecs from a JSON
// response object.
func flattenManagedZoneDnssecConfigDefaultKeySpecs(c *Client, i interface{}) *ManagedZoneDnssecConfigDefaultKeySpecs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZoneDnssecConfigDefaultKeySpecs{}
	r.Algorithm = flattenManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(m["algorithm"])
	r.KeyLength = dcl.FlattenInteger(m["keyLength"])
	r.KeyType = flattenManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(m["keyType"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandManagedZonePrivateVisibilityConfigMap expands the contents of ManagedZonePrivateVisibilityConfig into a JSON
// request object.
func expandManagedZonePrivateVisibilityConfigMap(c *Client, f map[string]ManagedZonePrivateVisibilityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZonePrivateVisibilityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZonePrivateVisibilityConfigSlice expands the contents of ManagedZonePrivateVisibilityConfig into a JSON
// request object.
func expandManagedZonePrivateVisibilityConfigSlice(c *Client, f []ManagedZonePrivateVisibilityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZonePrivateVisibilityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZonePrivateVisibilityConfigMap flattens the contents of ManagedZonePrivateVisibilityConfig from a JSON
// response object.
func flattenManagedZonePrivateVisibilityConfigMap(c *Client, i interface{}) map[string]ManagedZonePrivateVisibilityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZonePrivateVisibilityConfig{}
	}

	if len(a) == 0 {
		return map[string]ManagedZonePrivateVisibilityConfig{}
	}

	items := make(map[string]ManagedZonePrivateVisibilityConfig)
	for k, item := range a {
		items[k] = *flattenManagedZonePrivateVisibilityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZonePrivateVisibilityConfigSlice flattens the contents of ManagedZonePrivateVisibilityConfig from a JSON
// response object.
func flattenManagedZonePrivateVisibilityConfigSlice(c *Client, i interface{}) []ManagedZonePrivateVisibilityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZonePrivateVisibilityConfig{}
	}

	if len(a) == 0 {
		return []ManagedZonePrivateVisibilityConfig{}
	}

	items := make([]ManagedZonePrivateVisibilityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZonePrivateVisibilityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZonePrivateVisibilityConfig expands an instance of ManagedZonePrivateVisibilityConfig into a JSON
// request object.
func expandManagedZonePrivateVisibilityConfig(c *Client, f *ManagedZonePrivateVisibilityConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandManagedZonePrivateVisibilityConfigNetworksSlice(c, f.Networks); err != nil {
		return nil, fmt.Errorf("error expanding Networks into networks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networks"] = v
	}

	return m, nil
}

// flattenManagedZonePrivateVisibilityConfig flattens an instance of ManagedZonePrivateVisibilityConfig from a JSON
// response object.
func flattenManagedZonePrivateVisibilityConfig(c *Client, i interface{}) *ManagedZonePrivateVisibilityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZonePrivateVisibilityConfig{}
	r.Networks = flattenManagedZonePrivateVisibilityConfigNetworksSlice(c, m["networks"])

	return r
}

// expandManagedZonePrivateVisibilityConfigNetworksMap expands the contents of ManagedZonePrivateVisibilityConfigNetworks into a JSON
// request object.
func expandManagedZonePrivateVisibilityConfigNetworksMap(c *Client, f map[string]ManagedZonePrivateVisibilityConfigNetworks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZonePrivateVisibilityConfigNetworks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZonePrivateVisibilityConfigNetworksSlice expands the contents of ManagedZonePrivateVisibilityConfigNetworks into a JSON
// request object.
func expandManagedZonePrivateVisibilityConfigNetworksSlice(c *Client, f []ManagedZonePrivateVisibilityConfigNetworks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZonePrivateVisibilityConfigNetworks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZonePrivateVisibilityConfigNetworksMap flattens the contents of ManagedZonePrivateVisibilityConfigNetworks from a JSON
// response object.
func flattenManagedZonePrivateVisibilityConfigNetworksMap(c *Client, i interface{}) map[string]ManagedZonePrivateVisibilityConfigNetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZonePrivateVisibilityConfigNetworks{}
	}

	if len(a) == 0 {
		return map[string]ManagedZonePrivateVisibilityConfigNetworks{}
	}

	items := make(map[string]ManagedZonePrivateVisibilityConfigNetworks)
	for k, item := range a {
		items[k] = *flattenManagedZonePrivateVisibilityConfigNetworks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZonePrivateVisibilityConfigNetworksSlice flattens the contents of ManagedZonePrivateVisibilityConfigNetworks from a JSON
// response object.
func flattenManagedZonePrivateVisibilityConfigNetworksSlice(c *Client, i interface{}) []ManagedZonePrivateVisibilityConfigNetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZonePrivateVisibilityConfigNetworks{}
	}

	if len(a) == 0 {
		return []ManagedZonePrivateVisibilityConfigNetworks{}
	}

	items := make([]ManagedZonePrivateVisibilityConfigNetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZonePrivateVisibilityConfigNetworks(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZonePrivateVisibilityConfigNetworks expands an instance of ManagedZonePrivateVisibilityConfigNetworks into a JSON
// request object.
func expandManagedZonePrivateVisibilityConfigNetworks(c *Client, f *ManagedZonePrivateVisibilityConfigNetworks) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.NetworkUrl; !dcl.IsEmptyValueIndirect(v) {
		m["networkUrl"] = v
	}

	return m, nil
}

// flattenManagedZonePrivateVisibilityConfigNetworks flattens an instance of ManagedZonePrivateVisibilityConfigNetworks from a JSON
// response object.
func flattenManagedZonePrivateVisibilityConfigNetworks(c *Client, i interface{}) *ManagedZonePrivateVisibilityConfigNetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZonePrivateVisibilityConfigNetworks{}
	r.NetworkUrl = dcl.FlattenString(m["networkUrl"])

	return r
}

// expandManagedZoneForwardingConfigMap expands the contents of ManagedZoneForwardingConfig into a JSON
// request object.
func expandManagedZoneForwardingConfigMap(c *Client, f map[string]ManagedZoneForwardingConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZoneForwardingConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZoneForwardingConfigSlice expands the contents of ManagedZoneForwardingConfig into a JSON
// request object.
func expandManagedZoneForwardingConfigSlice(c *Client, f []ManagedZoneForwardingConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZoneForwardingConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZoneForwardingConfigMap flattens the contents of ManagedZoneForwardingConfig from a JSON
// response object.
func flattenManagedZoneForwardingConfigMap(c *Client, i interface{}) map[string]ManagedZoneForwardingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZoneForwardingConfig{}
	}

	if len(a) == 0 {
		return map[string]ManagedZoneForwardingConfig{}
	}

	items := make(map[string]ManagedZoneForwardingConfig)
	for k, item := range a {
		items[k] = *flattenManagedZoneForwardingConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZoneForwardingConfigSlice flattens the contents of ManagedZoneForwardingConfig from a JSON
// response object.
func flattenManagedZoneForwardingConfigSlice(c *Client, i interface{}) []ManagedZoneForwardingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneForwardingConfig{}
	}

	if len(a) == 0 {
		return []ManagedZoneForwardingConfig{}
	}

	items := make([]ManagedZoneForwardingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneForwardingConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZoneForwardingConfig expands an instance of ManagedZoneForwardingConfig into a JSON
// request object.
func expandManagedZoneForwardingConfig(c *Client, f *ManagedZoneForwardingConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandManagedZoneForwardingConfigTargetNameServersSlice(c, f.TargetNameServers); err != nil {
		return nil, fmt.Errorf("error expanding TargetNameServers into targetNameServers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["targetNameServers"] = v
	}

	return m, nil
}

// flattenManagedZoneForwardingConfig flattens an instance of ManagedZoneForwardingConfig from a JSON
// response object.
func flattenManagedZoneForwardingConfig(c *Client, i interface{}) *ManagedZoneForwardingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZoneForwardingConfig{}
	r.TargetNameServers = flattenManagedZoneForwardingConfigTargetNameServersSlice(c, m["targetNameServers"])

	return r
}

// expandManagedZoneForwardingConfigTargetNameServersMap expands the contents of ManagedZoneForwardingConfigTargetNameServers into a JSON
// request object.
func expandManagedZoneForwardingConfigTargetNameServersMap(c *Client, f map[string]ManagedZoneForwardingConfigTargetNameServers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZoneForwardingConfigTargetNameServers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZoneForwardingConfigTargetNameServersSlice expands the contents of ManagedZoneForwardingConfigTargetNameServers into a JSON
// request object.
func expandManagedZoneForwardingConfigTargetNameServersSlice(c *Client, f []ManagedZoneForwardingConfigTargetNameServers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZoneForwardingConfigTargetNameServers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZoneForwardingConfigTargetNameServersMap flattens the contents of ManagedZoneForwardingConfigTargetNameServers from a JSON
// response object.
func flattenManagedZoneForwardingConfigTargetNameServersMap(c *Client, i interface{}) map[string]ManagedZoneForwardingConfigTargetNameServers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZoneForwardingConfigTargetNameServers{}
	}

	if len(a) == 0 {
		return map[string]ManagedZoneForwardingConfigTargetNameServers{}
	}

	items := make(map[string]ManagedZoneForwardingConfigTargetNameServers)
	for k, item := range a {
		items[k] = *flattenManagedZoneForwardingConfigTargetNameServers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZoneForwardingConfigTargetNameServersSlice flattens the contents of ManagedZoneForwardingConfigTargetNameServers from a JSON
// response object.
func flattenManagedZoneForwardingConfigTargetNameServersSlice(c *Client, i interface{}) []ManagedZoneForwardingConfigTargetNameServers {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneForwardingConfigTargetNameServers{}
	}

	if len(a) == 0 {
		return []ManagedZoneForwardingConfigTargetNameServers{}
	}

	items := make([]ManagedZoneForwardingConfigTargetNameServers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneForwardingConfigTargetNameServers(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZoneForwardingConfigTargetNameServers expands an instance of ManagedZoneForwardingConfigTargetNameServers into a JSON
// request object.
func expandManagedZoneForwardingConfigTargetNameServers(c *Client, f *ManagedZoneForwardingConfigTargetNameServers) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.IPv4Address; !dcl.IsEmptyValueIndirect(v) {
		m["ipv4Address"] = v
	}
	if v := f.ForwardingPath; !dcl.IsEmptyValueIndirect(v) {
		m["forwardingPath"] = v
	}

	return m, nil
}

// flattenManagedZoneForwardingConfigTargetNameServers flattens an instance of ManagedZoneForwardingConfigTargetNameServers from a JSON
// response object.
func flattenManagedZoneForwardingConfigTargetNameServers(c *Client, i interface{}) *ManagedZoneForwardingConfigTargetNameServers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZoneForwardingConfigTargetNameServers{}
	r.IPv4Address = dcl.FlattenString(m["ipv4Address"])
	r.ForwardingPath = flattenManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(m["forwardingPath"])

	return r
}

// expandManagedZonePeeringConfigMap expands the contents of ManagedZonePeeringConfig into a JSON
// request object.
func expandManagedZonePeeringConfigMap(c *Client, f map[string]ManagedZonePeeringConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZonePeeringConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZonePeeringConfigSlice expands the contents of ManagedZonePeeringConfig into a JSON
// request object.
func expandManagedZonePeeringConfigSlice(c *Client, f []ManagedZonePeeringConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZonePeeringConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZonePeeringConfigMap flattens the contents of ManagedZonePeeringConfig from a JSON
// response object.
func flattenManagedZonePeeringConfigMap(c *Client, i interface{}) map[string]ManagedZonePeeringConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZonePeeringConfig{}
	}

	if len(a) == 0 {
		return map[string]ManagedZonePeeringConfig{}
	}

	items := make(map[string]ManagedZonePeeringConfig)
	for k, item := range a {
		items[k] = *flattenManagedZonePeeringConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZonePeeringConfigSlice flattens the contents of ManagedZonePeeringConfig from a JSON
// response object.
func flattenManagedZonePeeringConfigSlice(c *Client, i interface{}) []ManagedZonePeeringConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZonePeeringConfig{}
	}

	if len(a) == 0 {
		return []ManagedZonePeeringConfig{}
	}

	items := make([]ManagedZonePeeringConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZonePeeringConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZonePeeringConfig expands an instance of ManagedZonePeeringConfig into a JSON
// request object.
func expandManagedZonePeeringConfig(c *Client, f *ManagedZonePeeringConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandManagedZonePeeringConfigTargetNetwork(c, f.TargetNetwork); err != nil {
		return nil, fmt.Errorf("error expanding TargetNetwork into targetNetwork: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["targetNetwork"] = v
	}

	return m, nil
}

// flattenManagedZonePeeringConfig flattens an instance of ManagedZonePeeringConfig from a JSON
// response object.
func flattenManagedZonePeeringConfig(c *Client, i interface{}) *ManagedZonePeeringConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZonePeeringConfig{}
	r.TargetNetwork = flattenManagedZonePeeringConfigTargetNetwork(c, m["targetNetwork"])

	return r
}

// expandManagedZonePeeringConfigTargetNetworkMap expands the contents of ManagedZonePeeringConfigTargetNetwork into a JSON
// request object.
func expandManagedZonePeeringConfigTargetNetworkMap(c *Client, f map[string]ManagedZonePeeringConfigTargetNetwork) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandManagedZonePeeringConfigTargetNetwork(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandManagedZonePeeringConfigTargetNetworkSlice expands the contents of ManagedZonePeeringConfigTargetNetwork into a JSON
// request object.
func expandManagedZonePeeringConfigTargetNetworkSlice(c *Client, f []ManagedZonePeeringConfigTargetNetwork) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandManagedZonePeeringConfigTargetNetwork(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenManagedZonePeeringConfigTargetNetworkMap flattens the contents of ManagedZonePeeringConfigTargetNetwork from a JSON
// response object.
func flattenManagedZonePeeringConfigTargetNetworkMap(c *Client, i interface{}) map[string]ManagedZonePeeringConfigTargetNetwork {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ManagedZonePeeringConfigTargetNetwork{}
	}

	if len(a) == 0 {
		return map[string]ManagedZonePeeringConfigTargetNetwork{}
	}

	items := make(map[string]ManagedZonePeeringConfigTargetNetwork)
	for k, item := range a {
		items[k] = *flattenManagedZonePeeringConfigTargetNetwork(c, item.(map[string]interface{}))
	}

	return items
}

// flattenManagedZonePeeringConfigTargetNetworkSlice flattens the contents of ManagedZonePeeringConfigTargetNetwork from a JSON
// response object.
func flattenManagedZonePeeringConfigTargetNetworkSlice(c *Client, i interface{}) []ManagedZonePeeringConfigTargetNetwork {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZonePeeringConfigTargetNetwork{}
	}

	if len(a) == 0 {
		return []ManagedZonePeeringConfigTargetNetwork{}
	}

	items := make([]ManagedZonePeeringConfigTargetNetwork, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZonePeeringConfigTargetNetwork(c, item.(map[string]interface{})))
	}

	return items
}

// expandManagedZonePeeringConfigTargetNetwork expands an instance of ManagedZonePeeringConfigTargetNetwork into a JSON
// request object.
func expandManagedZonePeeringConfigTargetNetwork(c *Client, f *ManagedZonePeeringConfigTargetNetwork) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.NetworkUrl; !dcl.IsEmptyValueIndirect(v) {
		m["networkUrl"] = v
	}

	return m, nil
}

// flattenManagedZonePeeringConfigTargetNetwork flattens an instance of ManagedZonePeeringConfigTargetNetwork from a JSON
// response object.
func flattenManagedZonePeeringConfigTargetNetwork(c *Client, i interface{}) *ManagedZonePeeringConfigTargetNetwork {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ManagedZonePeeringConfigTargetNetwork{}
	r.NetworkUrl = dcl.FlattenString(m["networkUrl"])

	return r
}

// flattenManagedZoneDnssecConfigNonExistenceEnumSlice flattens the contents of ManagedZoneDnssecConfigNonExistenceEnum from a JSON
// response object.
func flattenManagedZoneDnssecConfigNonExistenceEnumSlice(c *Client, i interface{}) []ManagedZoneDnssecConfigNonExistenceEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneDnssecConfigNonExistenceEnum{}
	}

	if len(a) == 0 {
		return []ManagedZoneDnssecConfigNonExistenceEnum{}
	}

	items := make([]ManagedZoneDnssecConfigNonExistenceEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneDnssecConfigNonExistenceEnum(item.(interface{})))
	}

	return items
}

// flattenManagedZoneDnssecConfigNonExistenceEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedZoneDnssecConfigNonExistenceEnum with the same value as that string.
func flattenManagedZoneDnssecConfigNonExistenceEnum(i interface{}) *ManagedZoneDnssecConfigNonExistenceEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedZoneDnssecConfigNonExistenceEnumRef("")
	}

	return ManagedZoneDnssecConfigNonExistenceEnumRef(s)
}

// flattenManagedZoneDnssecConfigStateEnumSlice flattens the contents of ManagedZoneDnssecConfigStateEnum from a JSON
// response object.
func flattenManagedZoneDnssecConfigStateEnumSlice(c *Client, i interface{}) []ManagedZoneDnssecConfigStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneDnssecConfigStateEnum{}
	}

	if len(a) == 0 {
		return []ManagedZoneDnssecConfigStateEnum{}
	}

	items := make([]ManagedZoneDnssecConfigStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneDnssecConfigStateEnum(item.(interface{})))
	}

	return items
}

// flattenManagedZoneDnssecConfigStateEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedZoneDnssecConfigStateEnum with the same value as that string.
func flattenManagedZoneDnssecConfigStateEnum(i interface{}) *ManagedZoneDnssecConfigStateEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedZoneDnssecConfigStateEnumRef("")
	}

	return ManagedZoneDnssecConfigStateEnumRef(s)
}

// flattenManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumSlice flattens the contents of ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum from a JSON
// response object.
func flattenManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumSlice(c *Client, i interface{}) []ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum{}
	}

	items := make([]ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum with the same value as that string.
func flattenManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(i interface{}) *ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumRef("")
	}

	return ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumRef(s)
}

// flattenManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumSlice flattens the contents of ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum from a JSON
// response object.
func flattenManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumSlice(c *Client, i interface{}) []ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum{}
	}

	if len(a) == 0 {
		return []ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum{}
	}

	items := make([]ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum with the same value as that string.
func flattenManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(i interface{}) *ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumRef("")
	}

	return ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumRef(s)
}

// flattenManagedZoneVisibilityEnumSlice flattens the contents of ManagedZoneVisibilityEnum from a JSON
// response object.
func flattenManagedZoneVisibilityEnumSlice(c *Client, i interface{}) []ManagedZoneVisibilityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneVisibilityEnum{}
	}

	if len(a) == 0 {
		return []ManagedZoneVisibilityEnum{}
	}

	items := make([]ManagedZoneVisibilityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneVisibilityEnum(item.(interface{})))
	}

	return items
}

// flattenManagedZoneVisibilityEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedZoneVisibilityEnum with the same value as that string.
func flattenManagedZoneVisibilityEnum(i interface{}) *ManagedZoneVisibilityEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedZoneVisibilityEnumRef("")
	}

	return ManagedZoneVisibilityEnumRef(s)
}

// flattenManagedZoneForwardingConfigTargetNameServersForwardingPathEnumSlice flattens the contents of ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum from a JSON
// response object.
func flattenManagedZoneForwardingConfigTargetNameServersForwardingPathEnumSlice(c *Client, i interface{}) []ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum{}
	}

	if len(a) == 0 {
		return []ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum{}
	}

	items := make([]ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(item.(interface{})))
	}

	return items
}

// flattenManagedZoneForwardingConfigTargetNameServersForwardingPathEnum asserts that an interface is a string, and returns a
// pointer to a *ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum with the same value as that string.
func flattenManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(i interface{}) *ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum {
	s, ok := i.(string)
	if !ok {
		return ManagedZoneForwardingConfigTargetNameServersForwardingPathEnumRef("")
	}

	return ManagedZoneForwardingConfigTargetNameServersForwardingPathEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ManagedZone) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalManagedZone(b, c)
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
