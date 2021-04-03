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
package accesscontextmanager

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

func (r *AccessLevel) validate() error {

	if err := dcl.Required(r, "title"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Basic) {
		if err := r.Basic.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AccessLevelBasic) validate() error {
	return nil
}
func (r *AccessLevelBasicConditions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DevicePolicy) {
		if err := r.DevicePolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AccessLevelBasicConditionsDevicePolicy) validate() error {
	return nil
}
func (r *AccessLevelBasicConditionsDevicePolicyOsConstraints) validate() error {
	return nil
}

func accessLevelGetURL(userBasePath string, r *AccessLevel) (string, error) {
	params := map[string]interface{}{
		"policy": dcl.ValueOrEmptyString(r.Policy),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("accessPolicies/{{policy}}/accessLevels/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil
}

func accessLevelListURL(userBasePath, policy string) (string, error) {
	params := map[string]interface{}{
		"policy": policy,
	}
	return dcl.URL("{{policy}}/accessLevels", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil

}

func accessLevelCreateURL(userBasePath, policy string) (string, error) {
	params := map[string]interface{}{
		"policy": policy,
	}
	return dcl.URL("accessPolicies/{{policy}}/accessLevels", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil

}

func accessLevelDeleteURL(userBasePath string, r *AccessLevel) (string, error) {
	params := map[string]interface{}{
		"policy": dcl.ValueOrEmptyString(r.Policy),
		"name":   dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("accessPolicies/{{policy}}/accessLevels/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, params), nil
}

// accessLevelApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type accessLevelApiOperation interface {
	do(context.Context, *AccessLevel, *Client) error
}

// newUpdateAccessLevelUpdateRequest creates a request for an
// AccessLevel resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAccessLevelUpdateRequest(ctx context.Context, f *AccessLevel, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		req["title"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandAccessLevelBasic(c, f.Basic); err != nil {
		return nil, fmt.Errorf("error expanding Basic into basic: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["basic"] = v
	}
	if v, err := dcl.DeriveField("%s/accessLevels/%s", f.Name, f.Policy, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	return req, nil
}

// marshalUpdateAccessLevelUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateAccessLevelUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAccessLevelUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAccessLevelUpdateOperation) do(ctx context.Context, r *AccessLevel, c *Client) error {
	_, err := c.GetAccessLevel(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"title", "description", "basic", "name"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateAccessLevelUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateAccessLevelUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAccessLevelRaw(ctx context.Context, policy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := accessLevelListURL(c.Config.BasePath, policy)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AccessLevelMaxPage {
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

type listAccessLevelOperation struct {
	AccessLevels []map[string]interface{} `json:"accessLevels"`
	Token        string                   `json:"nextPageToken"`
}

func (c *Client) listAccessLevel(ctx context.Context, policy, pageToken string, pageSize int32) ([]*AccessLevel, string, error) {
	b, err := c.listAccessLevelRaw(ctx, policy, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAccessLevelOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AccessLevel
	for _, v := range m.AccessLevels {
		res := flattenAccessLevel(c, v)
		res.Policy = &policy
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAccessLevel(ctx context.Context, f func(*AccessLevel) bool, resources []*AccessLevel) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAccessLevel(ctx, res)
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

type deleteAccessLevelOperation struct{}

func (op *deleteAccessLevelOperation) do(ctx context.Context, r *AccessLevel, c *Client) error {

	_, err := c.GetAccessLevel(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("AccessLevel not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetAccessLevel checking for existence. error: %v", err)
		return err
	}

	u, err := accessLevelDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetAccessLevel(ctx, r.urlNormalized())
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
type createAccessLevelOperation struct {
	response map[string]interface{}
}

func (op *createAccessLevelOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAccessLevelOperation) do(ctx context.Context, r *AccessLevel, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	policy := r.createFields()
	u, err := accessLevelCreateURL(c.Config.BasePath, policy)

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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://accesscontextmanager.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetAccessLevel(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAccessLevelRaw(ctx context.Context, r *AccessLevel) ([]byte, error) {

	u, err := accessLevelGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) accessLevelDiffsForRawDesired(ctx context.Context, rawDesired *AccessLevel, opts ...dcl.ApplyOption) (initial, desired *AccessLevel, diffs []accessLevelDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AccessLevel
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AccessLevel); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected AccessLevel, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAccessLevel(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a AccessLevel resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AccessLevel resource: %v", err)
		}
		c.Config.Logger.Info("Found that AccessLevel resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAccessLevelDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for AccessLevel: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for AccessLevel: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAccessLevelInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for AccessLevel: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAccessLevelDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for AccessLevel: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAccessLevel(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAccessLevelInitialState(rawInitial, rawDesired *AccessLevel) (*AccessLevel, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAccessLevelDesiredState(rawDesired, rawInitial *AccessLevel, opts ...dcl.ApplyOption) (*AccessLevel, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Basic = canonicalizeAccessLevelBasic(rawDesired.Basic, nil, opts...)

		return rawDesired, nil
	}
	if dcl.StringCanonicalize(rawDesired.Title, rawInitial.Title) {
		rawDesired.Title = rawInitial.Title
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.Basic = canonicalizeAccessLevelBasic(rawDesired.Basic, rawInitial.Basic, opts...)
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Policy, rawInitial.Policy) {
		rawDesired.Policy = rawInitial.Policy
	}

	return rawDesired, nil
}

func canonicalizeAccessLevelNewState(c *Client, rawNew, rawDesired *AccessLevel) (*AccessLevel, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Title) && dcl.IsEmptyValueIndirect(rawDesired.Title) {
		rawNew.Title = rawDesired.Title
	} else {
		if dcl.StringCanonicalize(rawDesired.Title, rawNew.Title) {
			rawNew.Title = rawDesired.Title
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Basic) && dcl.IsEmptyValueIndirect(rawDesired.Basic) {
		rawNew.Basic = rawDesired.Basic
	} else {
		rawNew.Basic = canonicalizeNewAccessLevelBasic(c, rawDesired.Basic, rawNew.Basic)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.Policy = rawDesired.Policy

	return rawNew, nil
}

func canonicalizeAccessLevelBasic(des, initial *AccessLevelBasic, opts ...dcl.ApplyOption) *AccessLevelBasic {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.CombiningFunction) {
		des.CombiningFunction = AccessLevelBasicCombiningFunctionEnumRef("AND")
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.CombiningFunction) {
		des.CombiningFunction = initial.CombiningFunction
	}
	if dcl.IsZeroValue(des.Conditions) {
		des.Conditions = initial.Conditions
	}

	return des
}

func canonicalizeNewAccessLevelBasic(c *Client, des, nw *AccessLevelBasic) *AccessLevelBasic {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.CombiningFunction) {
		nw.CombiningFunction = AccessLevelBasicCombiningFunctionEnumRef("AND")
	}

	nw.Conditions = canonicalizeNewAccessLevelBasicConditionsSlice(c, des.Conditions, nw.Conditions)

	return nw
}

func canonicalizeNewAccessLevelBasicSet(c *Client, des, nw []AccessLevelBasic) []AccessLevelBasic {
	if des == nil {
		return nw
	}
	var reorderedNew []AccessLevelBasic
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAccessLevelBasic(c, &d, &n) {
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

func canonicalizeNewAccessLevelBasicSlice(c *Client, des, nw []AccessLevelBasic) []AccessLevelBasic {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AccessLevelBasic
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAccessLevelBasic(c, &d, &n))
	}

	return items
}

func canonicalizeAccessLevelBasicConditions(des, initial *AccessLevelBasicConditions, opts ...dcl.ApplyOption) *AccessLevelBasicConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Regions) {
		des.Regions = initial.Regions
	}
	if dcl.IsZeroValue(des.IPSubnetworks) {
		des.IPSubnetworks = initial.IPSubnetworks
	}
	if dcl.IsZeroValue(des.RequiredAccessLevels) {
		des.RequiredAccessLevels = initial.RequiredAccessLevels
	}
	if dcl.IsZeroValue(des.Members) {
		des.Members = initial.Members
	}
	if dcl.BoolCanonicalize(des.Negate, initial.Negate) || dcl.IsZeroValue(des.Negate) {
		des.Negate = initial.Negate
	}
	des.DevicePolicy = canonicalizeAccessLevelBasicConditionsDevicePolicy(des.DevicePolicy, initial.DevicePolicy, opts...)

	return des
}

func canonicalizeNewAccessLevelBasicConditions(c *Client, des, nw *AccessLevelBasicConditions) *AccessLevelBasicConditions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.Negate, nw.Negate) {
		nw.Negate = des.Negate
	}
	nw.DevicePolicy = canonicalizeNewAccessLevelBasicConditionsDevicePolicy(c, des.DevicePolicy, nw.DevicePolicy)

	return nw
}

func canonicalizeNewAccessLevelBasicConditionsSet(c *Client, des, nw []AccessLevelBasicConditions) []AccessLevelBasicConditions {
	if des == nil {
		return nw
	}
	var reorderedNew []AccessLevelBasicConditions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAccessLevelBasicConditions(c, &d, &n) {
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

func canonicalizeNewAccessLevelBasicConditionsSlice(c *Client, des, nw []AccessLevelBasicConditions) []AccessLevelBasicConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AccessLevelBasicConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAccessLevelBasicConditions(c, &d, &n))
	}

	return items
}

func canonicalizeAccessLevelBasicConditionsDevicePolicy(des, initial *AccessLevelBasicConditionsDevicePolicy, opts ...dcl.ApplyOption) *AccessLevelBasicConditionsDevicePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.RequireScreenlock, initial.RequireScreenlock) || dcl.IsZeroValue(des.RequireScreenlock) {
		des.RequireScreenlock = initial.RequireScreenlock
	}
	if dcl.BoolCanonicalize(des.RequireAdminApproval, initial.RequireAdminApproval) || dcl.IsZeroValue(des.RequireAdminApproval) {
		des.RequireAdminApproval = initial.RequireAdminApproval
	}
	if dcl.BoolCanonicalize(des.RequireCorpOwned, initial.RequireCorpOwned) || dcl.IsZeroValue(des.RequireCorpOwned) {
		des.RequireCorpOwned = initial.RequireCorpOwned
	}
	if dcl.IsZeroValue(des.AllowedEncryptionStatuses) {
		des.AllowedEncryptionStatuses = initial.AllowedEncryptionStatuses
	}
	if dcl.IsZeroValue(des.AllowedDeviceManagementLevels) {
		des.AllowedDeviceManagementLevels = initial.AllowedDeviceManagementLevels
	}
	if dcl.IsZeroValue(des.OsConstraints) {
		des.OsConstraints = initial.OsConstraints
	}

	return des
}

func canonicalizeNewAccessLevelBasicConditionsDevicePolicy(c *Client, des, nw *AccessLevelBasicConditionsDevicePolicy) *AccessLevelBasicConditionsDevicePolicy {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.RequireScreenlock, nw.RequireScreenlock) {
		nw.RequireScreenlock = des.RequireScreenlock
	}
	if dcl.BoolCanonicalize(des.RequireAdminApproval, nw.RequireAdminApproval) {
		nw.RequireAdminApproval = des.RequireAdminApproval
	}
	if dcl.BoolCanonicalize(des.RequireCorpOwned, nw.RequireCorpOwned) {
		nw.RequireCorpOwned = des.RequireCorpOwned
	}
	nw.OsConstraints = canonicalizeNewAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c, des.OsConstraints, nw.OsConstraints)

	return nw
}

func canonicalizeNewAccessLevelBasicConditionsDevicePolicySet(c *Client, des, nw []AccessLevelBasicConditionsDevicePolicy) []AccessLevelBasicConditionsDevicePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []AccessLevelBasicConditionsDevicePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAccessLevelBasicConditionsDevicePolicy(c, &d, &n) {
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

func canonicalizeNewAccessLevelBasicConditionsDevicePolicySlice(c *Client, des, nw []AccessLevelBasicConditionsDevicePolicy) []AccessLevelBasicConditionsDevicePolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AccessLevelBasicConditionsDevicePolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAccessLevelBasicConditionsDevicePolicy(c, &d, &n))
	}

	return items
}

func canonicalizeAccessLevelBasicConditionsDevicePolicyOsConstraints(des, initial *AccessLevelBasicConditionsDevicePolicyOsConstraints, opts ...dcl.ApplyOption) *AccessLevelBasicConditionsDevicePolicyOsConstraints {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.MinimumVersion, initial.MinimumVersion) || dcl.IsZeroValue(des.MinimumVersion) {
		des.MinimumVersion = initial.MinimumVersion
	}
	if dcl.IsZeroValue(des.OsType) {
		des.OsType = initial.OsType
	}
	if dcl.BoolCanonicalize(des.RequireVerifiedChromeOs, initial.RequireVerifiedChromeOs) || dcl.IsZeroValue(des.RequireVerifiedChromeOs) {
		des.RequireVerifiedChromeOs = initial.RequireVerifiedChromeOs
	}

	return des
}

func canonicalizeNewAccessLevelBasicConditionsDevicePolicyOsConstraints(c *Client, des, nw *AccessLevelBasicConditionsDevicePolicyOsConstraints) *AccessLevelBasicConditionsDevicePolicyOsConstraints {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.MinimumVersion, nw.MinimumVersion) {
		nw.MinimumVersion = des.MinimumVersion
	}
	if dcl.BoolCanonicalize(des.RequireVerifiedChromeOs, nw.RequireVerifiedChromeOs) {
		nw.RequireVerifiedChromeOs = des.RequireVerifiedChromeOs
	}

	return nw
}

func canonicalizeNewAccessLevelBasicConditionsDevicePolicyOsConstraintsSet(c *Client, des, nw []AccessLevelBasicConditionsDevicePolicyOsConstraints) []AccessLevelBasicConditionsDevicePolicyOsConstraints {
	if des == nil {
		return nw
	}
	var reorderedNew []AccessLevelBasicConditionsDevicePolicyOsConstraints
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareAccessLevelBasicConditionsDevicePolicyOsConstraints(c, &d, &n) {
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

func canonicalizeNewAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c *Client, des, nw []AccessLevelBasicConditionsDevicePolicyOsConstraints) []AccessLevelBasicConditionsDevicePolicyOsConstraints {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []AccessLevelBasicConditionsDevicePolicyOsConstraints
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAccessLevelBasicConditionsDevicePolicyOsConstraints(c, &d, &n))
	}

	return items
}

type accessLevelDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         accessLevelApiOperation
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
func diffAccessLevel(c *Client, desired, actual *AccessLevel, opts ...dcl.ApplyOption) ([]accessLevelDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []accessLevelDiff
	// New style diffs.
	if d, err := dcl.Diff(desired.Title, actual.Title, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, accessLevelDiff{
			UpdateOp: &updateAccessLevelUpdateOperation{}, FieldName: "Title",
		})
	}

	if d, err := dcl.Diff(desired.CreateTime, actual.CreateTime, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, accessLevelDiff{RequiresRecreate: true, FieldName: "CreateTime"})
	}

	if d, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, accessLevelDiff{RequiresRecreate: true, FieldName: "UpdateTime"})
	}

	if d, err := dcl.Diff(desired.Description, actual.Description, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, accessLevelDiff{
			UpdateOp: &updateAccessLevelUpdateOperation{}, FieldName: "Description",
		})
	}

	if d, err := dcl.Diff(desired.Name, actual.Name, &dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: ""}); d || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, accessLevelDiff{
			UpdateOp: &updateAccessLevelUpdateOperation{}, FieldName: "Name",
		})
	}

	if !dcl.IsZeroValue(desired.Title) && !dcl.StringCanonicalize(desired.Title, actual.Title) {
		c.Config.Logger.Infof("Detected diff in Title.\nDESIRED: %v\nACTUAL: %v", desired.Title, actual.Title)

		diffs = append(diffs, accessLevelDiff{
			UpdateOp:  &updateAccessLevelUpdateOperation{},
			FieldName: "Title",
		})

	}
	if !reflect.DeepEqual(desired.CreateTime, actual.CreateTime) {
		c.Config.Logger.Infof("Detected diff in CreateTime.\nDESIRED: %v\nACTUAL: %v", desired.CreateTime, actual.CreateTime)
		diffs = append(diffs, accessLevelDiff{
			RequiresRecreate: true,
			FieldName:        "CreateTime",
		})
	}
	if !reflect.DeepEqual(desired.UpdateTime, actual.UpdateTime) {
		c.Config.Logger.Infof("Detected diff in UpdateTime.\nDESIRED: %v\nACTUAL: %v", desired.UpdateTime, actual.UpdateTime)
		diffs = append(diffs, accessLevelDiff{
			RequiresRecreate: true,
			FieldName:        "UpdateTime",
		})
	}
	if !dcl.IsZeroValue(desired.Description) && !dcl.StringCanonicalize(desired.Description, actual.Description) {
		c.Config.Logger.Infof("Detected diff in Description.\nDESIRED: %v\nACTUAL: %v", desired.Description, actual.Description)

		diffs = append(diffs, accessLevelDiff{
			UpdateOp:  &updateAccessLevelUpdateOperation{},
			FieldName: "Description",
		})

	}
	if compareAccessLevelBasic(c, desired.Basic, actual.Basic) {
		c.Config.Logger.Infof("Detected diff in Basic.\nDESIRED: %v\nACTUAL: %v", desired.Basic, actual.Basic)

		diffs = append(diffs, accessLevelDiff{
			UpdateOp:  &updateAccessLevelUpdateOperation{},
			FieldName: "Basic",
		})

	}
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)

		diffs = append(diffs, accessLevelDiff{
			UpdateOp:  &updateAccessLevelUpdateOperation{},
			FieldName: "Name",
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
	var deduped []accessLevelDiff
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
func compareAccessLevelBasic(c *Client, desired, actual *AccessLevelBasic) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.CombiningFunction == nil && desired.CombiningFunction != nil && !dcl.IsEmptyValueIndirect(desired.CombiningFunction) {
		c.Config.Logger.Infof("desired CombiningFunction %s - but actually nil", dcl.SprintResource(desired.CombiningFunction))
		return true
	}
	if !reflect.DeepEqual(desired.CombiningFunction, actual.CombiningFunction) && !dcl.IsZeroValue(desired.CombiningFunction) {
		c.Config.Logger.Infof("Diff in CombiningFunction. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CombiningFunction), dcl.SprintResource(actual.CombiningFunction))
		return true
	}
	if actual.Conditions == nil && desired.Conditions != nil && !dcl.IsEmptyValueIndirect(desired.Conditions) {
		c.Config.Logger.Infof("desired Conditions %s - but actually nil", dcl.SprintResource(desired.Conditions))
		return true
	}
	if compareAccessLevelBasicConditionsSlice(c, desired.Conditions, actual.Conditions) && !dcl.IsZeroValue(desired.Conditions) {
		c.Config.Logger.Infof("Diff in Conditions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Conditions), dcl.SprintResource(actual.Conditions))
		return true
	}
	return false
}

func compareAccessLevelBasicSlice(c *Client, desired, actual []AccessLevelBasic) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasic, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasic(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasic, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicMap(c *Client, desired, actual map[string]AccessLevelBasic) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasic, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AccessLevelBasic, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAccessLevelBasic(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AccessLevelBasic, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditions(c *Client, desired, actual *AccessLevelBasicConditions) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Regions == nil && desired.Regions != nil && !dcl.IsEmptyValueIndirect(desired.Regions) {
		c.Config.Logger.Infof("desired Regions %s - but actually nil", dcl.SprintResource(desired.Regions))
		return true
	}
	if !dcl.StringSliceEquals(desired.Regions, actual.Regions) && !dcl.IsZeroValue(desired.Regions) {
		c.Config.Logger.Infof("Diff in Regions. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Regions), dcl.SprintResource(actual.Regions))
		return true
	}
	if actual.IPSubnetworks == nil && desired.IPSubnetworks != nil && !dcl.IsEmptyValueIndirect(desired.IPSubnetworks) {
		c.Config.Logger.Infof("desired IPSubnetworks %s - but actually nil", dcl.SprintResource(desired.IPSubnetworks))
		return true
	}
	if !dcl.StringSliceEquals(desired.IPSubnetworks, actual.IPSubnetworks) && !dcl.IsZeroValue(desired.IPSubnetworks) {
		c.Config.Logger.Infof("Diff in IPSubnetworks. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPSubnetworks), dcl.SprintResource(actual.IPSubnetworks))
		return true
	}
	if actual.RequiredAccessLevels == nil && desired.RequiredAccessLevels != nil && !dcl.IsEmptyValueIndirect(desired.RequiredAccessLevels) {
		c.Config.Logger.Infof("desired RequiredAccessLevels %s - but actually nil", dcl.SprintResource(desired.RequiredAccessLevels))
		return true
	}
	if !dcl.StringSliceEquals(desired.RequiredAccessLevels, actual.RequiredAccessLevels) && !dcl.IsZeroValue(desired.RequiredAccessLevels) {
		c.Config.Logger.Infof("Diff in RequiredAccessLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequiredAccessLevels), dcl.SprintResource(actual.RequiredAccessLevels))
		return true
	}
	if actual.Members == nil && desired.Members != nil && !dcl.IsEmptyValueIndirect(desired.Members) {
		c.Config.Logger.Infof("desired Members %s - but actually nil", dcl.SprintResource(desired.Members))
		return true
	}
	if !dcl.StringSliceEquals(desired.Members, actual.Members) && !dcl.IsZeroValue(desired.Members) {
		c.Config.Logger.Infof("Diff in Members. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Members), dcl.SprintResource(actual.Members))
		return true
	}
	if actual.Negate == nil && desired.Negate != nil && !dcl.IsEmptyValueIndirect(desired.Negate) {
		c.Config.Logger.Infof("desired Negate %s - but actually nil", dcl.SprintResource(desired.Negate))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Negate, actual.Negate) && !dcl.IsZeroValue(desired.Negate) {
		c.Config.Logger.Infof("Diff in Negate. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Negate), dcl.SprintResource(actual.Negate))
		return true
	}
	if actual.DevicePolicy == nil && desired.DevicePolicy != nil && !dcl.IsEmptyValueIndirect(desired.DevicePolicy) {
		c.Config.Logger.Infof("desired DevicePolicy %s - but actually nil", dcl.SprintResource(desired.DevicePolicy))
		return true
	}
	if compareAccessLevelBasicConditionsDevicePolicy(c, desired.DevicePolicy, actual.DevicePolicy) && !dcl.IsZeroValue(desired.DevicePolicy) {
		c.Config.Logger.Infof("Diff in DevicePolicy. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DevicePolicy), dcl.SprintResource(actual.DevicePolicy))
		return true
	}
	return false
}

func compareAccessLevelBasicConditionsSlice(c *Client, desired, actual []AccessLevelBasicConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditions, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicConditions(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditions, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsMap(c *Client, desired, actual map[string]AccessLevelBasicConditions) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditions, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditions, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAccessLevelBasicConditions(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditions, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicy(c *Client, desired, actual *AccessLevelBasicConditionsDevicePolicy) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RequireScreenlock == nil && desired.RequireScreenlock != nil && !dcl.IsEmptyValueIndirect(desired.RequireScreenlock) {
		c.Config.Logger.Infof("desired RequireScreenlock %s - but actually nil", dcl.SprintResource(desired.RequireScreenlock))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RequireScreenlock, actual.RequireScreenlock) && !dcl.IsZeroValue(desired.RequireScreenlock) {
		c.Config.Logger.Infof("Diff in RequireScreenlock. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequireScreenlock), dcl.SprintResource(actual.RequireScreenlock))
		return true
	}
	if actual.RequireAdminApproval == nil && desired.RequireAdminApproval != nil && !dcl.IsEmptyValueIndirect(desired.RequireAdminApproval) {
		c.Config.Logger.Infof("desired RequireAdminApproval %s - but actually nil", dcl.SprintResource(desired.RequireAdminApproval))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RequireAdminApproval, actual.RequireAdminApproval) && !dcl.IsZeroValue(desired.RequireAdminApproval) {
		c.Config.Logger.Infof("Diff in RequireAdminApproval. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequireAdminApproval), dcl.SprintResource(actual.RequireAdminApproval))
		return true
	}
	if actual.RequireCorpOwned == nil && desired.RequireCorpOwned != nil && !dcl.IsEmptyValueIndirect(desired.RequireCorpOwned) {
		c.Config.Logger.Infof("desired RequireCorpOwned %s - but actually nil", dcl.SprintResource(desired.RequireCorpOwned))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RequireCorpOwned, actual.RequireCorpOwned) && !dcl.IsZeroValue(desired.RequireCorpOwned) {
		c.Config.Logger.Infof("Diff in RequireCorpOwned. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequireCorpOwned), dcl.SprintResource(actual.RequireCorpOwned))
		return true
	}
	if actual.AllowedEncryptionStatuses == nil && desired.AllowedEncryptionStatuses != nil && !dcl.IsEmptyValueIndirect(desired.AllowedEncryptionStatuses) {
		c.Config.Logger.Infof("desired AllowedEncryptionStatuses %s - but actually nil", dcl.SprintResource(desired.AllowedEncryptionStatuses))
		return true
	}
	if compareAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumSlice(c, desired.AllowedEncryptionStatuses, actual.AllowedEncryptionStatuses) && !dcl.IsZeroValue(desired.AllowedEncryptionStatuses) {
		c.Config.Logger.Infof("Diff in AllowedEncryptionStatuses. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedEncryptionStatuses), dcl.SprintResource(actual.AllowedEncryptionStatuses))
		return true
	}
	if actual.AllowedDeviceManagementLevels == nil && desired.AllowedDeviceManagementLevels != nil && !dcl.IsEmptyValueIndirect(desired.AllowedDeviceManagementLevels) {
		c.Config.Logger.Infof("desired AllowedDeviceManagementLevels %s - but actually nil", dcl.SprintResource(desired.AllowedDeviceManagementLevels))
		return true
	}
	if compareAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumSlice(c, desired.AllowedDeviceManagementLevels, actual.AllowedDeviceManagementLevels) && !dcl.IsZeroValue(desired.AllowedDeviceManagementLevels) {
		c.Config.Logger.Infof("Diff in AllowedDeviceManagementLevels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AllowedDeviceManagementLevels), dcl.SprintResource(actual.AllowedDeviceManagementLevels))
		return true
	}
	if actual.OsConstraints == nil && desired.OsConstraints != nil && !dcl.IsEmptyValueIndirect(desired.OsConstraints) {
		c.Config.Logger.Infof("desired OsConstraints %s - but actually nil", dcl.SprintResource(desired.OsConstraints))
		return true
	}
	if compareAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c, desired.OsConstraints, actual.OsConstraints) && !dcl.IsZeroValue(desired.OsConstraints) {
		c.Config.Logger.Infof("Diff in OsConstraints. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OsConstraints), dcl.SprintResource(actual.OsConstraints))
		return true
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicySlice(c *Client, desired, actual []AccessLevelBasicConditionsDevicePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicy, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicConditionsDevicePolicy(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicy, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyMap(c *Client, desired, actual map[string]AccessLevelBasicConditionsDevicePolicy) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicy, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicy, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAccessLevelBasicConditionsDevicePolicy(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicy, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyOsConstraints(c *Client, desired, actual *AccessLevelBasicConditionsDevicePolicyOsConstraints) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.MinimumVersion == nil && desired.MinimumVersion != nil && !dcl.IsEmptyValueIndirect(desired.MinimumVersion) {
		c.Config.Logger.Infof("desired MinimumVersion %s - but actually nil", dcl.SprintResource(desired.MinimumVersion))
		return true
	}
	if !dcl.StringCanonicalize(desired.MinimumVersion, actual.MinimumVersion) && !dcl.IsZeroValue(desired.MinimumVersion) {
		c.Config.Logger.Infof("Diff in MinimumVersion. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinimumVersion), dcl.SprintResource(actual.MinimumVersion))
		return true
	}
	if actual.OsType == nil && desired.OsType != nil && !dcl.IsEmptyValueIndirect(desired.OsType) {
		c.Config.Logger.Infof("desired OsType %s - but actually nil", dcl.SprintResource(desired.OsType))
		return true
	}
	if !reflect.DeepEqual(desired.OsType, actual.OsType) && !dcl.IsZeroValue(desired.OsType) {
		c.Config.Logger.Infof("Diff in OsType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.OsType), dcl.SprintResource(actual.OsType))
		return true
	}
	if actual.RequireVerifiedChromeOs == nil && desired.RequireVerifiedChromeOs != nil && !dcl.IsEmptyValueIndirect(desired.RequireVerifiedChromeOs) {
		c.Config.Logger.Infof("desired RequireVerifiedChromeOs %s - but actually nil", dcl.SprintResource(desired.RequireVerifiedChromeOs))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RequireVerifiedChromeOs, actual.RequireVerifiedChromeOs) && !dcl.IsZeroValue(desired.RequireVerifiedChromeOs) {
		c.Config.Logger.Infof("Diff in RequireVerifiedChromeOs. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequireVerifiedChromeOs), dcl.SprintResource(actual.RequireVerifiedChromeOs))
		return true
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c *Client, desired, actual []AccessLevelBasicConditionsDevicePolicyOsConstraints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraints, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicConditionsDevicePolicyOsConstraints(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraints, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyOsConstraintsMap(c *Client, desired, actual map[string]AccessLevelBasicConditionsDevicePolicyOsConstraints) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraints, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraints, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareAccessLevelBasicConditionsDevicePolicyOsConstraints(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraints, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicCombiningFunctionEnumSlice(c *Client, desired, actual []AccessLevelBasicCombiningFunctionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicCombiningFunctionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicCombiningFunctionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicCombiningFunctionEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicCombiningFunctionEnum(c *Client, desired, actual *AccessLevelBasicCombiningFunctionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumSlice(c *Client, desired, actual []AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(c *Client, desired, actual *AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumSlice(c *Client, desired, actual []AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(c *Client, desired, actual *AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumSlice(c *Client, desired, actual []AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(c *Client, desired, actual *AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AccessLevel) urlNormalized() *AccessLevel {
	normalized := deepcopy.Copy(*r).(AccessLevel)
	normalized.Title = dcl.SelfLinkToName(r.Title)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Policy = dcl.SelfLinkToName(r.Policy)
	return &normalized
}

func (r *AccessLevel) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Policy), dcl.ValueOrEmptyString(n.Name)
}

func (r *AccessLevel) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Policy)
}

func (r *AccessLevel) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Policy), dcl.ValueOrEmptyString(n.Name)
}

func (r *AccessLevel) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"policy": dcl.ValueOrEmptyString(n.Policy),
			"name":   dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("accessPolicies/{{policy}}/accessLevels/{{name}}", "https://accesscontextmanager.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AccessLevel resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AccessLevel) marshal(c *Client) ([]byte, error) {
	m, err := expandAccessLevel(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AccessLevel: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAccessLevel decodes JSON responses into the AccessLevel resource schema.
func unmarshalAccessLevel(b []byte, c *Client) (*AccessLevel, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAccessLevel(m, c)
}

func unmarshalMapAccessLevel(m map[string]interface{}, c *Client) (*AccessLevel, error) {

	return flattenAccessLevel(c, m), nil
}

// expandAccessLevel expands AccessLevel into a JSON request object.
func expandAccessLevel(c *Client, f *AccessLevel) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Title; !dcl.IsEmptyValueIndirect(v) {
		m["title"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandAccessLevelBasic(c, f.Basic); err != nil {
		return nil, fmt.Errorf("error expanding Basic into basic: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["basic"] = v
	}
	if v, err := dcl.DeriveField("%s/accessLevels/%s", f.Name, f.Policy, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Policy into policy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policy"] = v
	}

	return m, nil
}

// flattenAccessLevel flattens AccessLevel from a JSON request object into the
// AccessLevel type.
func flattenAccessLevel(c *Client, i interface{}) *AccessLevel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &AccessLevel{}
	r.Title = dcl.FlattenString(m["title"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.Description = dcl.FlattenString(m["description"])
	r.Basic = flattenAccessLevelBasic(c, m["basic"])
	r.Name = dcl.FlattenString(m["name"])
	r.Policy = dcl.FlattenString(m["policy"])

	return r
}

// expandAccessLevelBasicMap expands the contents of AccessLevelBasic into a JSON
// request object.
func expandAccessLevelBasicMap(c *Client, f map[string]AccessLevelBasic) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAccessLevelBasic(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAccessLevelBasicSlice expands the contents of AccessLevelBasic into a JSON
// request object.
func expandAccessLevelBasicSlice(c *Client, f []AccessLevelBasic) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAccessLevelBasic(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAccessLevelBasicMap flattens the contents of AccessLevelBasic from a JSON
// response object.
func flattenAccessLevelBasicMap(c *Client, i interface{}) map[string]AccessLevelBasic {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AccessLevelBasic{}
	}

	if len(a) == 0 {
		return map[string]AccessLevelBasic{}
	}

	items := make(map[string]AccessLevelBasic)
	for k, item := range a {
		items[k] = *flattenAccessLevelBasic(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAccessLevelBasicSlice flattens the contents of AccessLevelBasic from a JSON
// response object.
func flattenAccessLevelBasicSlice(c *Client, i interface{}) []AccessLevelBasic {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasic{}
	}

	if len(a) == 0 {
		return []AccessLevelBasic{}
	}

	items := make([]AccessLevelBasic, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasic(c, item.(map[string]interface{})))
	}

	return items
}

// expandAccessLevelBasic expands an instance of AccessLevelBasic into a JSON
// request object.
func expandAccessLevelBasic(c *Client, f *AccessLevelBasic) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CombiningFunction; !dcl.IsEmptyValueIndirect(v) {
		m["combiningFunction"] = v
	}
	if v, err := expandAccessLevelBasicConditionsSlice(c, f.Conditions); err != nil {
		return nil, fmt.Errorf("error expanding Conditions into conditions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["conditions"] = v
	}

	return m, nil
}

// flattenAccessLevelBasic flattens an instance of AccessLevelBasic from a JSON
// response object.
func flattenAccessLevelBasic(c *Client, i interface{}) *AccessLevelBasic {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AccessLevelBasic{}
	r.CombiningFunction = flattenAccessLevelBasicCombiningFunctionEnum(m["combiningFunction"])
	if dcl.IsEmptyValueIndirect(m["combiningFunction"]) {
		c.Config.Logger.Info("Using default value for combiningFunction.")
		r.CombiningFunction = AccessLevelBasicCombiningFunctionEnumRef("AND")
	}
	r.Conditions = flattenAccessLevelBasicConditionsSlice(c, m["conditions"])

	return r
}

// expandAccessLevelBasicConditionsMap expands the contents of AccessLevelBasicConditions into a JSON
// request object.
func expandAccessLevelBasicConditionsMap(c *Client, f map[string]AccessLevelBasicConditions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAccessLevelBasicConditions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAccessLevelBasicConditionsSlice expands the contents of AccessLevelBasicConditions into a JSON
// request object.
func expandAccessLevelBasicConditionsSlice(c *Client, f []AccessLevelBasicConditions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAccessLevelBasicConditions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAccessLevelBasicConditionsMap flattens the contents of AccessLevelBasicConditions from a JSON
// response object.
func flattenAccessLevelBasicConditionsMap(c *Client, i interface{}) map[string]AccessLevelBasicConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AccessLevelBasicConditions{}
	}

	if len(a) == 0 {
		return map[string]AccessLevelBasicConditions{}
	}

	items := make(map[string]AccessLevelBasicConditions)
	for k, item := range a {
		items[k] = *flattenAccessLevelBasicConditions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAccessLevelBasicConditionsSlice flattens the contents of AccessLevelBasicConditions from a JSON
// response object.
func flattenAccessLevelBasicConditionsSlice(c *Client, i interface{}) []AccessLevelBasicConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicConditions{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicConditions{}
	}

	items := make([]AccessLevelBasicConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicConditions(c, item.(map[string]interface{})))
	}

	return items
}

// expandAccessLevelBasicConditions expands an instance of AccessLevelBasicConditions into a JSON
// request object.
func expandAccessLevelBasicConditions(c *Client, f *AccessLevelBasicConditions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Regions; !dcl.IsEmptyValueIndirect(v) {
		m["regions"] = v
	}
	if v := f.IPSubnetworks; !dcl.IsEmptyValueIndirect(v) {
		m["ipSubnetworks"] = v
	}
	if v := f.RequiredAccessLevels; !dcl.IsEmptyValueIndirect(v) {
		m["requiredAccessLevels"] = v
	}
	if v := f.Members; !dcl.IsEmptyValueIndirect(v) {
		m["members"] = v
	}
	if v := f.Negate; !dcl.IsEmptyValueIndirect(v) {
		m["negate"] = v
	}
	if v, err := expandAccessLevelBasicConditionsDevicePolicy(c, f.DevicePolicy); err != nil {
		return nil, fmt.Errorf("error expanding DevicePolicy into devicePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["devicePolicy"] = v
	}

	return m, nil
}

// flattenAccessLevelBasicConditions flattens an instance of AccessLevelBasicConditions from a JSON
// response object.
func flattenAccessLevelBasicConditions(c *Client, i interface{}) *AccessLevelBasicConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AccessLevelBasicConditions{}
	r.Regions = dcl.FlattenStringSlice(m["regions"])
	r.IPSubnetworks = dcl.FlattenStringSlice(m["ipSubnetworks"])
	r.RequiredAccessLevels = dcl.FlattenStringSlice(m["requiredAccessLevels"])
	r.Members = dcl.FlattenStringSlice(m["members"])
	r.Negate = dcl.FlattenBool(m["negate"])
	r.DevicePolicy = flattenAccessLevelBasicConditionsDevicePolicy(c, m["devicePolicy"])

	return r
}

// expandAccessLevelBasicConditionsDevicePolicyMap expands the contents of AccessLevelBasicConditionsDevicePolicy into a JSON
// request object.
func expandAccessLevelBasicConditionsDevicePolicyMap(c *Client, f map[string]AccessLevelBasicConditionsDevicePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAccessLevelBasicConditionsDevicePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAccessLevelBasicConditionsDevicePolicySlice expands the contents of AccessLevelBasicConditionsDevicePolicy into a JSON
// request object.
func expandAccessLevelBasicConditionsDevicePolicySlice(c *Client, f []AccessLevelBasicConditionsDevicePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAccessLevelBasicConditionsDevicePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAccessLevelBasicConditionsDevicePolicyMap flattens the contents of AccessLevelBasicConditionsDevicePolicy from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyMap(c *Client, i interface{}) map[string]AccessLevelBasicConditionsDevicePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AccessLevelBasicConditionsDevicePolicy{}
	}

	if len(a) == 0 {
		return map[string]AccessLevelBasicConditionsDevicePolicy{}
	}

	items := make(map[string]AccessLevelBasicConditionsDevicePolicy)
	for k, item := range a {
		items[k] = *flattenAccessLevelBasicConditionsDevicePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAccessLevelBasicConditionsDevicePolicySlice flattens the contents of AccessLevelBasicConditionsDevicePolicy from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicySlice(c *Client, i interface{}) []AccessLevelBasicConditionsDevicePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicConditionsDevicePolicy{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicConditionsDevicePolicy{}
	}

	items := make([]AccessLevelBasicConditionsDevicePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicConditionsDevicePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandAccessLevelBasicConditionsDevicePolicy expands an instance of AccessLevelBasicConditionsDevicePolicy into a JSON
// request object.
func expandAccessLevelBasicConditionsDevicePolicy(c *Client, f *AccessLevelBasicConditionsDevicePolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RequireScreenlock; !dcl.IsEmptyValueIndirect(v) {
		m["requireScreenlock"] = v
	}
	if v := f.RequireAdminApproval; !dcl.IsEmptyValueIndirect(v) {
		m["requireAdminApproval"] = v
	}
	if v := f.RequireCorpOwned; !dcl.IsEmptyValueIndirect(v) {
		m["requireCorpOwned"] = v
	}
	if v := f.AllowedEncryptionStatuses; !dcl.IsEmptyValueIndirect(v) {
		m["allowedEncryptionStatuses"] = v
	}
	if v := f.AllowedDeviceManagementLevels; !dcl.IsEmptyValueIndirect(v) {
		m["allowedDeviceManagementLevels"] = v
	}
	if v, err := expandAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c, f.OsConstraints); err != nil {
		return nil, fmt.Errorf("error expanding OsConstraints into osConstraints: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["osConstraints"] = v
	}

	return m, nil
}

// flattenAccessLevelBasicConditionsDevicePolicy flattens an instance of AccessLevelBasicConditionsDevicePolicy from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicy(c *Client, i interface{}) *AccessLevelBasicConditionsDevicePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AccessLevelBasicConditionsDevicePolicy{}
	r.RequireScreenlock = dcl.FlattenBool(m["requireScreenlock"])
	r.RequireAdminApproval = dcl.FlattenBool(m["requireAdminApproval"])
	r.RequireCorpOwned = dcl.FlattenBool(m["requireCorpOwned"])
	r.AllowedEncryptionStatuses = flattenAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumSlice(c, m["allowedEncryptionStatuses"])
	r.AllowedDeviceManagementLevels = flattenAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumSlice(c, m["allowedDeviceManagementLevels"])
	r.OsConstraints = flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c, m["osConstraints"])

	return r
}

// expandAccessLevelBasicConditionsDevicePolicyOsConstraintsMap expands the contents of AccessLevelBasicConditionsDevicePolicyOsConstraints into a JSON
// request object.
func expandAccessLevelBasicConditionsDevicePolicyOsConstraintsMap(c *Client, f map[string]AccessLevelBasicConditionsDevicePolicyOsConstraints) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAccessLevelBasicConditionsDevicePolicyOsConstraints(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice expands the contents of AccessLevelBasicConditionsDevicePolicyOsConstraints into a JSON
// request object.
func expandAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c *Client, f []AccessLevelBasicConditionsDevicePolicyOsConstraints) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAccessLevelBasicConditionsDevicePolicyOsConstraints(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsMap flattens the contents of AccessLevelBasicConditionsDevicePolicyOsConstraints from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsMap(c *Client, i interface{}) map[string]AccessLevelBasicConditionsDevicePolicyOsConstraints {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AccessLevelBasicConditionsDevicePolicyOsConstraints{}
	}

	if len(a) == 0 {
		return map[string]AccessLevelBasicConditionsDevicePolicyOsConstraints{}
	}

	items := make(map[string]AccessLevelBasicConditionsDevicePolicyOsConstraints)
	for k, item := range a {
		items[k] = *flattenAccessLevelBasicConditionsDevicePolicyOsConstraints(c, item.(map[string]interface{}))
	}

	return items
}

// flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice flattens the contents of AccessLevelBasicConditionsDevicePolicyOsConstraints from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsSlice(c *Client, i interface{}) []AccessLevelBasicConditionsDevicePolicyOsConstraints {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicConditionsDevicePolicyOsConstraints{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicConditionsDevicePolicyOsConstraints{}
	}

	items := make([]AccessLevelBasicConditionsDevicePolicyOsConstraints, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicConditionsDevicePolicyOsConstraints(c, item.(map[string]interface{})))
	}

	return items
}

// expandAccessLevelBasicConditionsDevicePolicyOsConstraints expands an instance of AccessLevelBasicConditionsDevicePolicyOsConstraints into a JSON
// request object.
func expandAccessLevelBasicConditionsDevicePolicyOsConstraints(c *Client, f *AccessLevelBasicConditionsDevicePolicyOsConstraints) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinimumVersion; !dcl.IsEmptyValueIndirect(v) {
		m["minimumVersion"] = v
	}
	if v := f.OsType; !dcl.IsEmptyValueIndirect(v) {
		m["osType"] = v
	}
	if v := f.RequireVerifiedChromeOs; !dcl.IsEmptyValueIndirect(v) {
		m["requireVerifiedChromeOs"] = v
	}

	return m, nil
}

// flattenAccessLevelBasicConditionsDevicePolicyOsConstraints flattens an instance of AccessLevelBasicConditionsDevicePolicyOsConstraints from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyOsConstraints(c *Client, i interface{}) *AccessLevelBasicConditionsDevicePolicyOsConstraints {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AccessLevelBasicConditionsDevicePolicyOsConstraints{}
	r.MinimumVersion = dcl.FlattenString(m["minimumVersion"])
	r.OsType = flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(m["osType"])
	r.RequireVerifiedChromeOs = dcl.FlattenBool(m["requireVerifiedChromeOs"])

	return r
}

// flattenAccessLevelBasicCombiningFunctionEnumSlice flattens the contents of AccessLevelBasicCombiningFunctionEnum from a JSON
// response object.
func flattenAccessLevelBasicCombiningFunctionEnumSlice(c *Client, i interface{}) []AccessLevelBasicCombiningFunctionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicCombiningFunctionEnum{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicCombiningFunctionEnum{}
	}

	items := make([]AccessLevelBasicCombiningFunctionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicCombiningFunctionEnum(item.(interface{})))
	}

	return items
}

// flattenAccessLevelBasicCombiningFunctionEnum asserts that an interface is a string, and returns a
// pointer to a *AccessLevelBasicCombiningFunctionEnum with the same value as that string.
func flattenAccessLevelBasicCombiningFunctionEnum(i interface{}) *AccessLevelBasicCombiningFunctionEnum {
	s, ok := i.(string)
	if !ok {
		return AccessLevelBasicCombiningFunctionEnumRef("")
	}

	return AccessLevelBasicCombiningFunctionEnumRef(s)
}

// flattenAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumSlice flattens the contents of AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumSlice(c *Client, i interface{}) []AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum{}
	}

	items := make([]AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(item.(interface{})))
	}

	return items
}

// flattenAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum asserts that an interface is a string, and returns a
// pointer to a *AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum with the same value as that string.
func flattenAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(i interface{}) *AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum {
	s, ok := i.(string)
	if !ok {
		return AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumRef("")
	}

	return AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumRef(s)
}

// flattenAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumSlice flattens the contents of AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumSlice(c *Client, i interface{}) []AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum{}
	}

	items := make([]AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(item.(interface{})))
	}

	return items
}

// flattenAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum asserts that an interface is a string, and returns a
// pointer to a *AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum with the same value as that string.
func flattenAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(i interface{}) *AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum {
	s, ok := i.(string)
	if !ok {
		return AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumRef("")
	}

	return AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumRef(s)
}

// flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumSlice flattens the contents of AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum from a JSON
// response object.
func flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumSlice(c *Client, i interface{}) []AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum{}
	}

	if len(a) == 0 {
		return []AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum{}
	}

	items := make([]AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum with the same value as that string.
func flattenAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(i interface{}) *AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum {
	s, ok := i.(string)
	if !ok {
		return AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumRef("")
	}

	return AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AccessLevel) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAccessLevel(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Policy == nil && ncr.Policy == nil {
			c.Config.Logger.Info("Both Policy fields null - considering equal.")
		} else if nr.Policy == nil || ncr.Policy == nil {
			c.Config.Logger.Info("Only one Policy field is null - considering unequal.")
			return false
		} else if *nr.Policy != *ncr.Policy {
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
