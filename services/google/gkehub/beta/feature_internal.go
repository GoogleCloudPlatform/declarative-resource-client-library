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

func (r *Feature) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ResourceState) {
		if err := r.ResourceState.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Spec) {
		if err := r.Spec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureResourceState) validate() error {
	return nil
}
func (r *FeatureSpec) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Multiclusteringress) {
		if err := r.Multiclusteringress.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Helloworld) {
		if err := r.Helloworld.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecMulticlusteringress) validate() error {
	return nil
}
func (r *FeatureSpecHelloworld) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FeatureTest) {
		if err := r.FeatureTest.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecHelloworldFeatureTest) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Fifth) {
		if err := r.Fifth.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecHelloworldFeatureTestFifth) validate() error {
	return nil
}
func (r *FeatureSpecHelloworldFeatureTestEighth) validate() error {
	return nil
}
func (r *FeatureState) validate() error {
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Helloworld) {
		if err := r.Helloworld.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureStateState) validate() error {
	return nil
}
func (r *FeatureStateHelloworld) validate() error {
	return nil
}

// featureApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type featureApiOperation interface {
	do(context.Context, *Feature, *Client) error
}

// newUpdateFeatureUpdateFeatureRequest creates a request for an
// Feature resource's UpdateFeature update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFeatureUpdateFeatureRequest(ctx context.Context, f *Feature, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["spec"] = v
	}
	if v := f.MembershipSpecs; !dcl.IsEmptyValueIndirect(v) {
		req["membershipSpecs"] = v
	}
	return req, nil
}

// marshalUpdateFeatureUpdateFeatureRequest converts the update into
// the final JSON request body.
func marshalUpdateFeatureUpdateFeatureRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFeatureUpdateFeatureOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listFeatureRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := featureListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FeatureMaxPage {
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

type listFeatureOperation struct {
	Resources []map[string]interface{} `json:"resources"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listFeature(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*Feature, string, error) {
	b, err := c.listFeatureRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFeatureOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Feature
	for _, v := range m.Resources {
		res := flattenFeature(c, v)
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFeature(ctx context.Context, f func(*Feature) bool, resources []*Feature) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFeature(ctx, res)
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

type deleteFeatureOperation struct{}

func (op *deleteFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {

	_, err := c.GetFeature(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Feature not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetFeature checking for existence. error: %v", err)
		return err
	}

	u, err := featureDeleteURL(c.Config.BasePath, r.urlNormalized())
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
	if err := o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetFeature(ctx, r.urlNormalized())
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
type createFeatureOperation struct {
	response map[string]interface{}
}

func (op *createFeatureOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, location, name := r.createFields()
	u, err := featureCreateURL(c.Config.BasePath, project, location, name)

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
	if err := o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetFeature(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFeatureRaw(ctx context.Context, r *Feature) ([]byte, error) {

	u, err := featureGetURL(c.Config.BasePath, r.urlNormalized())
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

func (c *Client) featureDiffsForRawDesired(ctx context.Context, rawDesired *Feature, opts ...dcl.ApplyOption) (initial, desired *Feature, diffs []featureDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Feature
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Feature); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Feature, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFeature(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Feature resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Feature resource: %v", err)
		}
		c.Config.Logger.Info("Found that Feature resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFeatureDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Feature: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Feature: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFeatureInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Feature: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFeatureDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Feature: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFeature(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFeatureInitialState(rawInitial, rawDesired *Feature) (*Feature, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFeatureDesiredState(rawDesired, rawInitial *Feature, opts ...dcl.ApplyOption) (*Feature, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ResourceState = canonicalizeFeatureResourceState(rawDesired.ResourceState, nil, opts...)
		rawDesired.Spec = canonicalizeFeatureSpec(rawDesired.Spec, nil, opts...)
		rawDesired.State = canonicalizeFeatureState(rawDesired.State, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	rawDesired.ResourceState = canonicalizeFeatureResourceState(rawDesired.ResourceState, rawInitial.ResourceState, opts...)
	rawDesired.Spec = canonicalizeFeatureSpec(rawDesired.Spec, rawInitial.Spec, opts...)
	if dcl.IsZeroValue(rawDesired.MembershipSpecs) {
		rawDesired.MembershipSpecs = rawInitial.MembershipSpecs
	}
	rawDesired.State = canonicalizeFeatureState(rawDesired.State, rawInitial.State, opts...)
	if dcl.IsZeroValue(rawDesired.MembershipStates) {
		rawDesired.MembershipStates = rawInitial.MembershipStates
	}
	if dcl.IsZeroValue(rawDesired.CreateTime) {
		rawDesired.CreateTime = rawInitial.CreateTime
	}
	if dcl.IsZeroValue(rawDesired.UpdateTime) {
		rawDesired.UpdateTime = rawInitial.UpdateTime
	}
	if dcl.IsZeroValue(rawDesired.DeleteTime) {
		rawDesired.DeleteTime = rawInitial.DeleteTime
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeFeatureNewState(c *Client, rawNew, rawDesired *Feature) (*Feature, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceState) && dcl.IsEmptyValueIndirect(rawDesired.ResourceState) {
		rawNew.ResourceState = rawDesired.ResourceState
	} else {
		rawNew.ResourceState = canonicalizeNewFeatureResourceState(c, rawDesired.ResourceState, rawNew.ResourceState)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Spec) && dcl.IsEmptyValueIndirect(rawDesired.Spec) {
		rawNew.Spec = rawDesired.Spec
	} else {
		rawNew.Spec = canonicalizeNewFeatureSpec(c, rawDesired.Spec, rawNew.Spec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MembershipSpecs) && dcl.IsEmptyValueIndirect(rawDesired.MembershipSpecs) {
		rawNew.MembershipSpecs = rawDesired.MembershipSpecs
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewFeatureState(c, rawDesired.State, rawNew.State)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MembershipStates) && dcl.IsEmptyValueIndirect(rawDesired.MembershipStates) {
		rawNew.MembershipStates = rawDesired.MembershipStates
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeFeatureResourceState(des, initial *FeatureResourceState, opts ...dcl.ApplyOption) *FeatureResourceState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.State) {
		des.State = initial.State
	}

	return des
}

func canonicalizeNewFeatureResourceState(c *Client, des, nw *FeatureResourceState) *FeatureResourceState {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewFeatureResourceStateSet(c *Client, des, nw []FeatureResourceState) []FeatureResourceState {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureResourceState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureResourceState(c, &d, &n) {
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

func canonicalizeNewFeatureResourceStateSlice(c *Client, des, nw []FeatureResourceState) []FeatureResourceState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureResourceState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureResourceState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpec(des, initial *FeatureSpec, opts ...dcl.ApplyOption) *FeatureSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Multiclusteringress = canonicalizeFeatureSpecMulticlusteringress(des.Multiclusteringress, initial.Multiclusteringress, opts...)
	des.Helloworld = canonicalizeFeatureSpecHelloworld(des.Helloworld, initial.Helloworld, opts...)

	return des
}

func canonicalizeNewFeatureSpec(c *Client, des, nw *FeatureSpec) *FeatureSpec {
	if des == nil || nw == nil {
		return nw
	}

	nw.Multiclusteringress = canonicalizeNewFeatureSpecMulticlusteringress(c, des.Multiclusteringress, nw.Multiclusteringress)
	nw.Helloworld = canonicalizeNewFeatureSpecHelloworld(c, des.Helloworld, nw.Helloworld)

	return nw
}

func canonicalizeNewFeatureSpecSet(c *Client, des, nw []FeatureSpec) []FeatureSpec {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureSpec(c, &d, &n) {
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

func canonicalizeNewFeatureSpecSlice(c *Client, des, nw []FeatureSpec) []FeatureSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpec(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecMulticlusteringress(des, initial *FeatureSpecMulticlusteringress, opts ...dcl.ApplyOption) *FeatureSpecMulticlusteringress {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ConfigMembership, initial.ConfigMembership) || dcl.IsZeroValue(des.ConfigMembership) {
		des.ConfigMembership = initial.ConfigMembership
	}
	if dcl.IsZeroValue(des.Billing) {
		des.Billing = initial.Billing
	}

	return des
}

func canonicalizeNewFeatureSpecMulticlusteringress(c *Client, des, nw *FeatureSpecMulticlusteringress) *FeatureSpecMulticlusteringress {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ConfigMembership, nw.ConfigMembership) {
		nw.ConfigMembership = des.ConfigMembership
	}

	return nw
}

func canonicalizeNewFeatureSpecMulticlusteringressSet(c *Client, des, nw []FeatureSpecMulticlusteringress) []FeatureSpecMulticlusteringress {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecMulticlusteringress
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureSpecMulticlusteringress(c, &d, &n) {
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

func canonicalizeNewFeatureSpecMulticlusteringressSlice(c *Client, des, nw []FeatureSpecMulticlusteringress) []FeatureSpecMulticlusteringress {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureSpecMulticlusteringress
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecMulticlusteringress(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecHelloworld(des, initial *FeatureSpecHelloworld, opts ...dcl.ApplyOption) *FeatureSpecHelloworld {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.FeatureTest = canonicalizeFeatureSpecHelloworldFeatureTest(des.FeatureTest, initial.FeatureTest, opts...)
	if dcl.StringCanonicalize(des.CustomConfig, initial.CustomConfig) || dcl.IsZeroValue(des.CustomConfig) {
		des.CustomConfig = initial.CustomConfig
	}

	return des
}

func canonicalizeNewFeatureSpecHelloworld(c *Client, des, nw *FeatureSpecHelloworld) *FeatureSpecHelloworld {
	if des == nil || nw == nil {
		return nw
	}

	nw.FeatureTest = canonicalizeNewFeatureSpecHelloworldFeatureTest(c, des.FeatureTest, nw.FeatureTest)
	if dcl.StringCanonicalize(des.CustomConfig, nw.CustomConfig) {
		nw.CustomConfig = des.CustomConfig
	}

	return nw
}

func canonicalizeNewFeatureSpecHelloworldSet(c *Client, des, nw []FeatureSpecHelloworld) []FeatureSpecHelloworld {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecHelloworld
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureSpecHelloworld(c, &d, &n) {
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

func canonicalizeNewFeatureSpecHelloworldSlice(c *Client, des, nw []FeatureSpecHelloworld) []FeatureSpecHelloworld {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureSpecHelloworld
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecHelloworld(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecHelloworldFeatureTest(des, initial *FeatureSpecHelloworldFeatureTest, opts ...dcl.ApplyOption) *FeatureSpecHelloworldFeatureTest {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.First, initial.First) || dcl.IsZeroValue(des.First) {
		des.First = initial.First
	}
	if dcl.IsZeroValue(des.Second) {
		des.Second = initial.Second
	}
	if dcl.IsZeroValue(des.Third) {
		des.Third = initial.Third
	}
	if dcl.StringCanonicalize(des.Fourth, initial.Fourth) || dcl.IsZeroValue(des.Fourth) {
		des.Fourth = initial.Fourth
	}
	des.Fifth = canonicalizeFeatureSpecHelloworldFeatureTestFifth(des.Fifth, initial.Fifth, opts...)
	if dcl.IsZeroValue(des.Sixth) {
		des.Sixth = initial.Sixth
	}
	if dcl.StringCanonicalize(des.Seventh, initial.Seventh) || dcl.IsZeroValue(des.Seventh) {
		des.Seventh = initial.Seventh
	}
	if dcl.IsZeroValue(des.Eighth) {
		des.Eighth = initial.Eighth
	}
	if dcl.IsZeroValue(des.Ninth) {
		des.Ninth = initial.Ninth
	}

	return des
}

func canonicalizeNewFeatureSpecHelloworldFeatureTest(c *Client, des, nw *FeatureSpecHelloworldFeatureTest) *FeatureSpecHelloworldFeatureTest {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.First, nw.First) {
		nw.First = des.First
	}
	if dcl.StringCanonicalize(des.Fourth, nw.Fourth) {
		nw.Fourth = des.Fourth
	}
	nw.Fifth = canonicalizeNewFeatureSpecHelloworldFeatureTestFifth(c, des.Fifth, nw.Fifth)
	if dcl.StringCanonicalize(des.Seventh, nw.Seventh) {
		nw.Seventh = des.Seventh
	}
	nw.Eighth = canonicalizeNewFeatureSpecHelloworldFeatureTestEighthSlice(c, des.Eighth, nw.Eighth)

	return nw
}

func canonicalizeNewFeatureSpecHelloworldFeatureTestSet(c *Client, des, nw []FeatureSpecHelloworldFeatureTest) []FeatureSpecHelloworldFeatureTest {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecHelloworldFeatureTest
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureSpecHelloworldFeatureTest(c, &d, &n) {
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

func canonicalizeNewFeatureSpecHelloworldFeatureTestSlice(c *Client, des, nw []FeatureSpecHelloworldFeatureTest) []FeatureSpecHelloworldFeatureTest {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureSpecHelloworldFeatureTest
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecHelloworldFeatureTest(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecHelloworldFeatureTestFifth(des, initial *FeatureSpecHelloworldFeatureTestFifth, opts ...dcl.ApplyOption) *FeatureSpecHelloworldFeatureTestFifth {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.TypeUrl, initial.TypeUrl) || dcl.IsZeroValue(des.TypeUrl) {
		des.TypeUrl = initial.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewFeatureSpecHelloworldFeatureTestFifth(c *Client, des, nw *FeatureSpecHelloworldFeatureTestFifth) *FeatureSpecHelloworldFeatureTestFifth {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.TypeUrl, nw.TypeUrl) {
		nw.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewFeatureSpecHelloworldFeatureTestFifthSet(c *Client, des, nw []FeatureSpecHelloworldFeatureTestFifth) []FeatureSpecHelloworldFeatureTestFifth {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecHelloworldFeatureTestFifth
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureSpecHelloworldFeatureTestFifth(c, &d, &n) {
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

func canonicalizeNewFeatureSpecHelloworldFeatureTestFifthSlice(c *Client, des, nw []FeatureSpecHelloworldFeatureTestFifth) []FeatureSpecHelloworldFeatureTestFifth {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureSpecHelloworldFeatureTestFifth
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecHelloworldFeatureTestFifth(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecHelloworldFeatureTestEighth(des, initial *FeatureSpecHelloworldFeatureTestEighth, opts ...dcl.ApplyOption) *FeatureSpecHelloworldFeatureTestEighth {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.First, initial.First) || dcl.IsZeroValue(des.First) {
		des.First = initial.First
	}
	if dcl.IsZeroValue(des.Second) {
		des.Second = initial.Second
	}

	return des
}

func canonicalizeNewFeatureSpecHelloworldFeatureTestEighth(c *Client, des, nw *FeatureSpecHelloworldFeatureTestEighth) *FeatureSpecHelloworldFeatureTestEighth {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.First, nw.First) {
		nw.First = des.First
	}

	return nw
}

func canonicalizeNewFeatureSpecHelloworldFeatureTestEighthSet(c *Client, des, nw []FeatureSpecHelloworldFeatureTestEighth) []FeatureSpecHelloworldFeatureTestEighth {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureSpecHelloworldFeatureTestEighth
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureSpecHelloworldFeatureTestEighth(c, &d, &n) {
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

func canonicalizeNewFeatureSpecHelloworldFeatureTestEighthSlice(c *Client, des, nw []FeatureSpecHelloworldFeatureTestEighth) []FeatureSpecHelloworldFeatureTestEighth {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureSpecHelloworldFeatureTestEighth
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecHelloworldFeatureTestEighth(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureState(des, initial *FeatureState, opts ...dcl.ApplyOption) *FeatureState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.State = canonicalizeFeatureStateState(des.State, initial.State, opts...)
	des.Helloworld = canonicalizeFeatureStateHelloworld(des.Helloworld, initial.Helloworld, opts...)

	return des
}

func canonicalizeNewFeatureState(c *Client, des, nw *FeatureState) *FeatureState {
	if des == nil || nw == nil {
		return nw
	}

	nw.State = canonicalizeNewFeatureStateState(c, des.State, nw.State)
	nw.Helloworld = canonicalizeNewFeatureStateHelloworld(c, des.Helloworld, nw.Helloworld)

	return nw
}

func canonicalizeNewFeatureStateSet(c *Client, des, nw []FeatureState) []FeatureState {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureState(c, &d, &n) {
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

func canonicalizeNewFeatureStateSlice(c *Client, des, nw []FeatureState) []FeatureState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateState(des, initial *FeatureStateState, opts ...dcl.ApplyOption) *FeatureStateState {
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
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		des.Description = initial.Description
	}
	if dcl.IsZeroValue(des.UpdateTime) {
		des.UpdateTime = initial.UpdateTime
	}

	return des
}

func canonicalizeNewFeatureStateState(c *Client, des, nw *FeatureStateState) *FeatureStateState {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewFeatureStateStateSet(c *Client, des, nw []FeatureStateState) []FeatureStateState {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateState
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureStateState(c, &d, &n) {
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

func canonicalizeNewFeatureStateStateSlice(c *Client, des, nw []FeatureStateState) []FeatureStateState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureStateState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateHelloworld(des, initial *FeatureStateHelloworld, opts ...dcl.ApplyOption) *FeatureStateHelloworld {
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

func canonicalizeNewFeatureStateHelloworld(c *Client, des, nw *FeatureStateHelloworld) *FeatureStateHelloworld {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewFeatureStateHelloworldSet(c *Client, des, nw []FeatureStateHelloworld) []FeatureStateHelloworld {
	if des == nil {
		return nw
	}
	var reorderedNew []FeatureStateHelloworld
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareFeatureStateHelloworld(c, &d, &n) {
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

func canonicalizeNewFeatureStateHelloworldSlice(c *Client, des, nw []FeatureStateHelloworld) []FeatureStateHelloworld {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []FeatureStateHelloworld
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateHelloworld(c, &d, &n))
	}

	return items
}

type featureDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         featureApiOperation
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
func diffFeature(c *Client, desired, actual *Feature, opts ...dcl.ApplyOption) ([]featureDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []featureDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.PartialSelfLinkToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, featureDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.MapEquals(desired.Labels, actual.Labels, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in Labels.\nDESIRED: %v\nACTUAL: %v", desired.Labels, actual.Labels)

		diffs = append(diffs, featureDiff{
			UpdateOp:  &updateFeatureUpdateFeatureOperation{},
			FieldName: "Labels",
		})

	}
	if compareFeatureSpec(c, desired.Spec, actual.Spec) {
		c.Config.Logger.Infof("Detected diff in Spec.\nDESIRED: %v\nACTUAL: %v", desired.Spec, actual.Spec)

		diffs = append(diffs, featureDiff{
			UpdateOp:  &updateFeatureUpdateFeatureOperation{},
			FieldName: "Spec",
		})

	}
	if !dcl.MapEquals(desired.MembershipSpecs, actual.MembershipSpecs, []string(nil)) {
		c.Config.Logger.Infof("Detected diff in MembershipSpecs.\nDESIRED: %v\nACTUAL: %v", desired.MembershipSpecs, actual.MembershipSpecs)

		diffs = append(diffs, featureDiff{
			UpdateOp:  &updateFeatureUpdateFeatureOperation{},
			FieldName: "MembershipSpecs",
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
	var deduped []featureDiff
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
func compareFeatureResourceState(c *Client, desired, actual *FeatureResourceState) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.State == nil && desired.State != nil && !dcl.IsEmptyValueIndirect(desired.State) {
		c.Config.Logger.Infof("desired State %s - but actually nil", dcl.SprintResource(desired.State))
		return true
	}
	if !reflect.DeepEqual(desired.State, actual.State) && !dcl.IsZeroValue(desired.State) {
		c.Config.Logger.Infof("Diff in State. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.State), dcl.SprintResource(actual.State))
		return true
	}
	return false
}

func compareFeatureResourceStateSlice(c *Client, desired, actual []FeatureResourceState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureResourceState, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureResourceState(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureResourceState, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureResourceStateMap(c *Client, desired, actual map[string]FeatureResourceState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureResourceState, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureResourceState, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureResourceState(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureResourceState, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureSpec(c *Client, desired, actual *FeatureSpec) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Multiclusteringress == nil && desired.Multiclusteringress != nil && !dcl.IsEmptyValueIndirect(desired.Multiclusteringress) {
		c.Config.Logger.Infof("desired Multiclusteringress %s - but actually nil", dcl.SprintResource(desired.Multiclusteringress))
		return true
	}
	if compareFeatureSpecMulticlusteringress(c, desired.Multiclusteringress, actual.Multiclusteringress) && !dcl.IsZeroValue(desired.Multiclusteringress) {
		c.Config.Logger.Infof("Diff in Multiclusteringress. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Multiclusteringress), dcl.SprintResource(actual.Multiclusteringress))
		return true
	}
	if actual.Helloworld == nil && desired.Helloworld != nil && !dcl.IsEmptyValueIndirect(desired.Helloworld) {
		c.Config.Logger.Infof("desired Helloworld %s - but actually nil", dcl.SprintResource(desired.Helloworld))
		return true
	}
	if compareFeatureSpecHelloworld(c, desired.Helloworld, actual.Helloworld) && !dcl.IsZeroValue(desired.Helloworld) {
		c.Config.Logger.Infof("Diff in Helloworld. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Helloworld), dcl.SprintResource(actual.Helloworld))
		return true
	}
	return false
}

func compareFeatureSpecSlice(c *Client, desired, actual []FeatureSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpec, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpec(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpec, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecMap(c *Client, desired, actual map[string]FeatureSpec) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpec, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureSpec, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureSpec(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureSpec, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureSpecMulticlusteringress(c *Client, desired, actual *FeatureSpecMulticlusteringress) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ConfigMembership == nil && desired.ConfigMembership != nil && !dcl.IsEmptyValueIndirect(desired.ConfigMembership) {
		c.Config.Logger.Infof("desired ConfigMembership %s - but actually nil", dcl.SprintResource(desired.ConfigMembership))
		return true
	}
	if !dcl.StringCanonicalize(desired.ConfigMembership, actual.ConfigMembership) && !dcl.IsZeroValue(desired.ConfigMembership) {
		c.Config.Logger.Infof("Diff in ConfigMembership. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConfigMembership), dcl.SprintResource(actual.ConfigMembership))
		return true
	}
	if actual.Billing == nil && desired.Billing != nil && !dcl.IsEmptyValueIndirect(desired.Billing) {
		c.Config.Logger.Infof("desired Billing %s - but actually nil", dcl.SprintResource(desired.Billing))
		return true
	}
	if !reflect.DeepEqual(desired.Billing, actual.Billing) && !dcl.IsZeroValue(desired.Billing) {
		c.Config.Logger.Infof("Diff in Billing. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Billing), dcl.SprintResource(actual.Billing))
		return true
	}
	return false
}

func compareFeatureSpecMulticlusteringressSlice(c *Client, desired, actual []FeatureSpecMulticlusteringress) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecMulticlusteringress, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecMulticlusteringress(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecMulticlusteringress, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecMulticlusteringressMap(c *Client, desired, actual map[string]FeatureSpecMulticlusteringress) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecMulticlusteringress, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureSpecMulticlusteringress, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureSpecMulticlusteringress(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureSpecMulticlusteringress, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworld(c *Client, desired, actual *FeatureSpecHelloworld) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.FeatureTest == nil && desired.FeatureTest != nil && !dcl.IsEmptyValueIndirect(desired.FeatureTest) {
		c.Config.Logger.Infof("desired FeatureTest %s - but actually nil", dcl.SprintResource(desired.FeatureTest))
		return true
	}
	if compareFeatureSpecHelloworldFeatureTest(c, desired.FeatureTest, actual.FeatureTest) && !dcl.IsZeroValue(desired.FeatureTest) {
		c.Config.Logger.Infof("Diff in FeatureTest. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FeatureTest), dcl.SprintResource(actual.FeatureTest))
		return true
	}
	if actual.CustomConfig == nil && desired.CustomConfig != nil && !dcl.IsEmptyValueIndirect(desired.CustomConfig) {
		c.Config.Logger.Infof("desired CustomConfig %s - but actually nil", dcl.SprintResource(desired.CustomConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.CustomConfig, actual.CustomConfig) && !dcl.IsZeroValue(desired.CustomConfig) {
		c.Config.Logger.Infof("Diff in CustomConfig. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CustomConfig), dcl.SprintResource(actual.CustomConfig))
		return true
	}
	return false
}

func compareFeatureSpecHelloworldSlice(c *Client, desired, actual []FeatureSpecHelloworld) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworld, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecHelloworld(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworld, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldMap(c *Client, desired, actual map[string]FeatureSpecHelloworld) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworld, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworld, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureSpecHelloworld(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworld, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTest(c *Client, desired, actual *FeatureSpecHelloworldFeatureTest) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.First == nil && desired.First != nil && !dcl.IsEmptyValueIndirect(desired.First) {
		c.Config.Logger.Infof("desired First %s - but actually nil", dcl.SprintResource(desired.First))
		return true
	}
	if !dcl.StringCanonicalize(desired.First, actual.First) && !dcl.IsZeroValue(desired.First) {
		c.Config.Logger.Infof("Diff in First. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.First), dcl.SprintResource(actual.First))
		return true
	}
	if actual.Second == nil && desired.Second != nil && !dcl.IsEmptyValueIndirect(desired.Second) {
		c.Config.Logger.Infof("desired Second %s - but actually nil", dcl.SprintResource(desired.Second))
		return true
	}
	if !reflect.DeepEqual(desired.Second, actual.Second) && !dcl.IsZeroValue(desired.Second) {
		c.Config.Logger.Infof("Diff in Second. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Second), dcl.SprintResource(actual.Second))
		return true
	}
	if actual.Third == nil && desired.Third != nil && !dcl.IsEmptyValueIndirect(desired.Third) {
		c.Config.Logger.Infof("desired Third %s - but actually nil", dcl.SprintResource(desired.Third))
		return true
	}
	if !reflect.DeepEqual(desired.Third, actual.Third) && !dcl.IsZeroValue(desired.Third) {
		c.Config.Logger.Infof("Diff in Third. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Third), dcl.SprintResource(actual.Third))
		return true
	}
	if actual.Fourth == nil && desired.Fourth != nil && !dcl.IsEmptyValueIndirect(desired.Fourth) {
		c.Config.Logger.Infof("desired Fourth %s - but actually nil", dcl.SprintResource(desired.Fourth))
		return true
	}
	if !dcl.StringCanonicalize(desired.Fourth, actual.Fourth) && !dcl.IsZeroValue(desired.Fourth) {
		c.Config.Logger.Infof("Diff in Fourth. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fourth), dcl.SprintResource(actual.Fourth))
		return true
	}
	if actual.Fifth == nil && desired.Fifth != nil && !dcl.IsEmptyValueIndirect(desired.Fifth) {
		c.Config.Logger.Infof("desired Fifth %s - but actually nil", dcl.SprintResource(desired.Fifth))
		return true
	}
	if compareFeatureSpecHelloworldFeatureTestFifth(c, desired.Fifth, actual.Fifth) && !dcl.IsZeroValue(desired.Fifth) {
		c.Config.Logger.Infof("Diff in Fifth. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Fifth), dcl.SprintResource(actual.Fifth))
		return true
	}
	if actual.Sixth == nil && desired.Sixth != nil && !dcl.IsEmptyValueIndirect(desired.Sixth) {
		c.Config.Logger.Infof("desired Sixth %s - but actually nil", dcl.SprintResource(desired.Sixth))
		return true
	}
	if !reflect.DeepEqual(desired.Sixth, actual.Sixth) && !dcl.IsZeroValue(desired.Sixth) {
		c.Config.Logger.Infof("Diff in Sixth. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sixth), dcl.SprintResource(actual.Sixth))
		return true
	}
	if actual.Seventh == nil && desired.Seventh != nil && !dcl.IsEmptyValueIndirect(desired.Seventh) {
		c.Config.Logger.Infof("desired Seventh %s - but actually nil", dcl.SprintResource(desired.Seventh))
		return true
	}
	if !dcl.StringCanonicalize(desired.Seventh, actual.Seventh) && !dcl.IsZeroValue(desired.Seventh) {
		c.Config.Logger.Infof("Diff in Seventh. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seventh), dcl.SprintResource(actual.Seventh))
		return true
	}
	if actual.Eighth == nil && desired.Eighth != nil && !dcl.IsEmptyValueIndirect(desired.Eighth) {
		c.Config.Logger.Infof("desired Eighth %s - but actually nil", dcl.SprintResource(desired.Eighth))
		return true
	}
	if compareFeatureSpecHelloworldFeatureTestEighthSlice(c, desired.Eighth, actual.Eighth) && !dcl.IsZeroValue(desired.Eighth) {
		c.Config.Logger.Infof("Diff in Eighth. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Eighth), dcl.SprintResource(actual.Eighth))
		return true
	}
	if actual.Ninth == nil && desired.Ninth != nil && !dcl.IsEmptyValueIndirect(desired.Ninth) {
		c.Config.Logger.Infof("desired Ninth %s - but actually nil", dcl.SprintResource(desired.Ninth))
		return true
	}
	if !dcl.MapEquals(desired.Ninth, actual.Ninth, []string(nil)) && !dcl.IsZeroValue(desired.Ninth) {
		c.Config.Logger.Infof("Diff in Ninth. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Ninth), dcl.SprintResource(actual.Ninth))
		return true
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestSlice(c *Client, desired, actual []FeatureSpecHelloworldFeatureTest) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTest, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecHelloworldFeatureTest(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTest, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestMap(c *Client, desired, actual map[string]FeatureSpecHelloworldFeatureTest) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTest, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTest, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureSpecHelloworldFeatureTest(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTest, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestFifth(c *Client, desired, actual *FeatureSpecHelloworldFeatureTestFifth) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.TypeUrl == nil && desired.TypeUrl != nil && !dcl.IsEmptyValueIndirect(desired.TypeUrl) {
		c.Config.Logger.Infof("desired TypeUrl %s - but actually nil", dcl.SprintResource(desired.TypeUrl))
		return true
	}
	if !dcl.StringCanonicalize(desired.TypeUrl, actual.TypeUrl) && !dcl.IsZeroValue(desired.TypeUrl) {
		c.Config.Logger.Infof("Diff in TypeUrl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TypeUrl), dcl.SprintResource(actual.TypeUrl))
		return true
	}
	if actual.Value == nil && desired.Value != nil && !dcl.IsEmptyValueIndirect(desired.Value) {
		c.Config.Logger.Infof("desired Value %s - but actually nil", dcl.SprintResource(desired.Value))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestFifthSlice(c *Client, desired, actual []FeatureSpecHelloworldFeatureTestFifth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTestFifth, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecHelloworldFeatureTestFifth(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestFifth, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestFifthMap(c *Client, desired, actual map[string]FeatureSpecHelloworldFeatureTestFifth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTestFifth, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestFifth, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureSpecHelloworldFeatureTestFifth(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestFifth, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestEighth(c *Client, desired, actual *FeatureSpecHelloworldFeatureTestEighth) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.First == nil && desired.First != nil && !dcl.IsEmptyValueIndirect(desired.First) {
		c.Config.Logger.Infof("desired First %s - but actually nil", dcl.SprintResource(desired.First))
		return true
	}
	if !dcl.StringCanonicalize(desired.First, actual.First) && !dcl.IsZeroValue(desired.First) {
		c.Config.Logger.Infof("Diff in First. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.First), dcl.SprintResource(actual.First))
		return true
	}
	if actual.Second == nil && desired.Second != nil && !dcl.IsEmptyValueIndirect(desired.Second) {
		c.Config.Logger.Infof("desired Second %s - but actually nil", dcl.SprintResource(desired.Second))
		return true
	}
	if !reflect.DeepEqual(desired.Second, actual.Second) && !dcl.IsZeroValue(desired.Second) {
		c.Config.Logger.Infof("Diff in Second. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Second), dcl.SprintResource(actual.Second))
		return true
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestEighthSlice(c *Client, desired, actual []FeatureSpecHelloworldFeatureTestEighth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTestEighth, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecHelloworldFeatureTestEighth(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestEighth, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestEighthMap(c *Client, desired, actual map[string]FeatureSpecHelloworldFeatureTestEighth) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTestEighth, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestEighth, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureSpecHelloworldFeatureTestEighth(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestEighth, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureState(c *Client, desired, actual *FeatureState) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Helloworld == nil && desired.Helloworld != nil && !dcl.IsEmptyValueIndirect(desired.Helloworld) {
		c.Config.Logger.Infof("desired Helloworld %s - but actually nil", dcl.SprintResource(desired.Helloworld))
		return true
	}
	if compareFeatureStateHelloworld(c, desired.Helloworld, actual.Helloworld) && !dcl.IsZeroValue(desired.Helloworld) {
		c.Config.Logger.Infof("Diff in Helloworld. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Helloworld), dcl.SprintResource(actual.Helloworld))
		return true
	}
	return false
}

func compareFeatureStateSlice(c *Client, desired, actual []FeatureState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureState, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureState(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureState, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureStateMap(c *Client, desired, actual map[string]FeatureState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureState, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureState, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureState(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureState, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureStateState(c *Client, desired, actual *FeatureStateState) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Code == nil && desired.Code != nil && !dcl.IsEmptyValueIndirect(desired.Code) {
		c.Config.Logger.Infof("desired Code %s - but actually nil", dcl.SprintResource(desired.Code))
		return true
	}
	if !reflect.DeepEqual(desired.Code, actual.Code) && !dcl.IsZeroValue(desired.Code) {
		c.Config.Logger.Infof("Diff in Code. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Code), dcl.SprintResource(actual.Code))
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
	if actual.UpdateTime == nil && desired.UpdateTime != nil && !dcl.IsEmptyValueIndirect(desired.UpdateTime) {
		c.Config.Logger.Infof("desired UpdateTime %s - but actually nil", dcl.SprintResource(desired.UpdateTime))
		return true
	}
	if !reflect.DeepEqual(desired.UpdateTime, actual.UpdateTime) && !dcl.IsZeroValue(desired.UpdateTime) {
		c.Config.Logger.Infof("Diff in UpdateTime. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpdateTime), dcl.SprintResource(actual.UpdateTime))
		return true
	}
	return false
}

func compareFeatureStateStateSlice(c *Client, desired, actual []FeatureStateState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureStateState, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureStateState(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureStateState, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureStateStateMap(c *Client, desired, actual map[string]FeatureStateState) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureStateState, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureStateState, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureStateState(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureStateState, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureStateHelloworld(c *Client, desired, actual *FeatureStateHelloworld) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	return false
}

func compareFeatureStateHelloworldSlice(c *Client, desired, actual []FeatureStateHelloworld) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureStateHelloworld, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureStateHelloworld(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureStateHelloworld, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureStateHelloworldMap(c *Client, desired, actual map[string]FeatureStateHelloworld) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureStateHelloworld, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in FeatureStateHelloworld, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareFeatureStateHelloworld(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in FeatureStateHelloworld, key %s. \nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareFeatureResourceStateStateEnumSlice(c *Client, desired, actual []FeatureResourceStateStateEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureResourceStateStateEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureResourceStateStateEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureResourceStateStateEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureResourceStateStateEnum(c *Client, desired, actual *FeatureResourceStateStateEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareFeatureSpecMulticlusteringressBillingEnumSlice(c *Client, desired, actual []FeatureSpecMulticlusteringressBillingEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecMulticlusteringressBillingEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecMulticlusteringressBillingEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecMulticlusteringressBillingEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecMulticlusteringressBillingEnum(c *Client, desired, actual *FeatureSpecMulticlusteringressBillingEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareFeatureSpecHelloworldFeatureTestThirdEnumSlice(c *Client, desired, actual []FeatureSpecHelloworldFeatureTestThirdEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureSpecHelloworldFeatureTestThirdEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureSpecHelloworldFeatureTestThirdEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureSpecHelloworldFeatureTestThirdEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureSpecHelloworldFeatureTestThirdEnum(c *Client, desired, actual *FeatureSpecHelloworldFeatureTestThirdEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareFeatureStateStateCodeEnumSlice(c *Client, desired, actual []FeatureStateStateCodeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in FeatureStateStateCodeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareFeatureStateStateCodeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in FeatureStateStateCodeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareFeatureStateStateCodeEnum(c *Client, desired, actual *FeatureStateStateCodeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Feature) urlNormalized() *Feature {
	normalized := deepcopy.Copy(*r).(Feature)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Feature) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Feature) createFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Feature) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *Feature) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "UpdateFeature" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(n.Project),
			"location": dcl.ValueOrEmptyString(n.Location),
			"name":     dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", "https://gkehub.googleapis.com/v1beta1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Feature resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Feature) marshal(c *Client) ([]byte, error) {
	m, err := expandFeature(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Feature: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFeature decodes JSON responses into the Feature resource schema.
func unmarshalFeature(b []byte, c *Client) (*Feature, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFeature(m, c)
}

func unmarshalMapFeature(m map[string]interface{}, c *Client) (*Feature, error) {

	return flattenFeature(c, m), nil
}

// expandFeature expands Feature into a JSON request object.
func expandFeature(c *Client, f *Feature) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/features/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandFeatureResourceState(c, f.ResourceState); err != nil {
		return nil, fmt.Errorf("error expanding ResourceState into resourceState: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceState"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["spec"] = v
	}
	if v := f.MembershipSpecs; !dcl.IsEmptyValueIndirect(v) {
		m["membershipSpecs"] = v
	}
	if v, err := expandFeatureState(c, f.State); err != nil {
		return nil, fmt.Errorf("error expanding State into state: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.MembershipStates; !dcl.IsEmptyValueIndirect(v) {
		m["membershipStates"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.DeleteTime; !dcl.IsEmptyValueIndirect(v) {
		m["deleteTime"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenFeature flattens Feature from a JSON request object into the
// Feature type.
func flattenFeature(c *Client, i interface{}) *Feature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Feature{}
	r.Name = dcl.FlattenString(m["name"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.ResourceState = flattenFeatureResourceState(c, m["resourceState"])
	r.Spec = flattenFeatureSpec(c, m["spec"])
	r.MembershipSpecs = dcl.FlattenKeyValuePairs(m["membershipSpecs"])
	r.State = flattenFeatureState(c, m["state"])
	r.MembershipStates = dcl.FlattenKeyValuePairs(m["membershipStates"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])
	r.DeleteTime = dcl.FlattenString(m["deleteTime"])
	r.Project = dcl.FlattenString(m["project"])
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandFeatureResourceStateMap expands the contents of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceStateMap(c *Client, f map[string]FeatureResourceState) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureResourceState(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureResourceStateSlice expands the contents of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceStateSlice(c *Client, f []FeatureResourceState) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureResourceState(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureResourceStateMap flattens the contents of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceStateMap(c *Client, i interface{}) map[string]FeatureResourceState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureResourceState{}
	}

	if len(a) == 0 {
		return map[string]FeatureResourceState{}
	}

	items := make(map[string]FeatureResourceState)
	for k, item := range a {
		items[k] = *flattenFeatureResourceState(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureResourceStateSlice flattens the contents of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceStateSlice(c *Client, i interface{}) []FeatureResourceState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureResourceState{}
	}

	if len(a) == 0 {
		return []FeatureResourceState{}
	}

	items := make([]FeatureResourceState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureResourceState(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureResourceState expands an instance of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceState(c *Client, f *FeatureResourceState) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}

	return m, nil
}

// flattenFeatureResourceState flattens an instance of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceState(c *Client, i interface{}) *FeatureResourceState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureResourceState{}
	r.State = flattenFeatureResourceStateStateEnum(m["state"])

	return r
}

// expandFeatureSpecMap expands the contents of FeatureSpec into a JSON
// request object.
func expandFeatureSpecMap(c *Client, f map[string]FeatureSpec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecSlice expands the contents of FeatureSpec into a JSON
// request object.
func expandFeatureSpecSlice(c *Client, f []FeatureSpec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecMap flattens the contents of FeatureSpec from a JSON
// response object.
func flattenFeatureSpecMap(c *Client, i interface{}) map[string]FeatureSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpec{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpec{}
	}

	items := make(map[string]FeatureSpec)
	for k, item := range a {
		items[k] = *flattenFeatureSpec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureSpecSlice flattens the contents of FeatureSpec from a JSON
// response object.
func flattenFeatureSpecSlice(c *Client, i interface{}) []FeatureSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpec{}
	}

	if len(a) == 0 {
		return []FeatureSpec{}
	}

	items := make([]FeatureSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpec(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureSpec expands an instance of FeatureSpec into a JSON
// request object.
func expandFeatureSpec(c *Client, f *FeatureSpec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureSpecMulticlusteringress(c, f.Multiclusteringress); err != nil {
		return nil, fmt.Errorf("error expanding Multiclusteringress into multiclusteringress: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["multiclusteringress"] = v
	}
	if v, err := expandFeatureSpecHelloworld(c, f.Helloworld); err != nil {
		return nil, fmt.Errorf("error expanding Helloworld into helloworld: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["helloworld"] = v
	}

	return m, nil
}

// flattenFeatureSpec flattens an instance of FeatureSpec from a JSON
// response object.
func flattenFeatureSpec(c *Client, i interface{}) *FeatureSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpec{}
	r.Multiclusteringress = flattenFeatureSpecMulticlusteringress(c, m["multiclusteringress"])
	r.Helloworld = flattenFeatureSpecHelloworld(c, m["helloworld"])

	return r
}

// expandFeatureSpecMulticlusteringressMap expands the contents of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringressMap(c *Client, f map[string]FeatureSpecMulticlusteringress) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecMulticlusteringress(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecMulticlusteringressSlice expands the contents of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringressSlice(c *Client, f []FeatureSpecMulticlusteringress) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecMulticlusteringress(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecMulticlusteringressMap flattens the contents of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressMap(c *Client, i interface{}) map[string]FeatureSpecMulticlusteringress {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecMulticlusteringress{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecMulticlusteringress{}
	}

	items := make(map[string]FeatureSpecMulticlusteringress)
	for k, item := range a {
		items[k] = *flattenFeatureSpecMulticlusteringress(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureSpecMulticlusteringressSlice flattens the contents of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressSlice(c *Client, i interface{}) []FeatureSpecMulticlusteringress {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecMulticlusteringress{}
	}

	if len(a) == 0 {
		return []FeatureSpecMulticlusteringress{}
	}

	items := make([]FeatureSpecMulticlusteringress, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecMulticlusteringress(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureSpecMulticlusteringress expands an instance of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringress(c *Client, f *FeatureSpecMulticlusteringress) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConfigMembership; !dcl.IsEmptyValueIndirect(v) {
		m["configMembership"] = v
	}
	if v := f.Billing; !dcl.IsEmptyValueIndirect(v) {
		m["billing"] = v
	}

	return m, nil
}

// flattenFeatureSpecMulticlusteringress flattens an instance of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringress(c *Client, i interface{}) *FeatureSpecMulticlusteringress {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecMulticlusteringress{}
	r.ConfigMembership = dcl.FlattenString(m["configMembership"])
	r.Billing = flattenFeatureSpecMulticlusteringressBillingEnum(m["billing"])

	return r
}

// expandFeatureSpecHelloworldMap expands the contents of FeatureSpecHelloworld into a JSON
// request object.
func expandFeatureSpecHelloworldMap(c *Client, f map[string]FeatureSpecHelloworld) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecHelloworld(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecHelloworldSlice expands the contents of FeatureSpecHelloworld into a JSON
// request object.
func expandFeatureSpecHelloworldSlice(c *Client, f []FeatureSpecHelloworld) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecHelloworld(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecHelloworldMap flattens the contents of FeatureSpecHelloworld from a JSON
// response object.
func flattenFeatureSpecHelloworldMap(c *Client, i interface{}) map[string]FeatureSpecHelloworld {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecHelloworld{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecHelloworld{}
	}

	items := make(map[string]FeatureSpecHelloworld)
	for k, item := range a {
		items[k] = *flattenFeatureSpecHelloworld(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureSpecHelloworldSlice flattens the contents of FeatureSpecHelloworld from a JSON
// response object.
func flattenFeatureSpecHelloworldSlice(c *Client, i interface{}) []FeatureSpecHelloworld {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecHelloworld{}
	}

	if len(a) == 0 {
		return []FeatureSpecHelloworld{}
	}

	items := make([]FeatureSpecHelloworld, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecHelloworld(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureSpecHelloworld expands an instance of FeatureSpecHelloworld into a JSON
// request object.
func expandFeatureSpecHelloworld(c *Client, f *FeatureSpecHelloworld) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureSpecHelloworldFeatureTest(c, f.FeatureTest); err != nil {
		return nil, fmt.Errorf("error expanding FeatureTest into featureTest: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["featureTest"] = v
	}
	if v := f.CustomConfig; !dcl.IsEmptyValueIndirect(v) {
		m["customConfig"] = v
	}

	return m, nil
}

// flattenFeatureSpecHelloworld flattens an instance of FeatureSpecHelloworld from a JSON
// response object.
func flattenFeatureSpecHelloworld(c *Client, i interface{}) *FeatureSpecHelloworld {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecHelloworld{}
	r.FeatureTest = flattenFeatureSpecHelloworldFeatureTest(c, m["featureTest"])
	r.CustomConfig = dcl.FlattenString(m["customConfig"])

	return r
}

// expandFeatureSpecHelloworldFeatureTestMap expands the contents of FeatureSpecHelloworldFeatureTest into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestMap(c *Client, f map[string]FeatureSpecHelloworldFeatureTest) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecHelloworldFeatureTest(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecHelloworldFeatureTestSlice expands the contents of FeatureSpecHelloworldFeatureTest into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestSlice(c *Client, f []FeatureSpecHelloworldFeatureTest) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecHelloworldFeatureTest(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecHelloworldFeatureTestMap flattens the contents of FeatureSpecHelloworldFeatureTest from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestMap(c *Client, i interface{}) map[string]FeatureSpecHelloworldFeatureTest {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecHelloworldFeatureTest{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecHelloworldFeatureTest{}
	}

	items := make(map[string]FeatureSpecHelloworldFeatureTest)
	for k, item := range a {
		items[k] = *flattenFeatureSpecHelloworldFeatureTest(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureSpecHelloworldFeatureTestSlice flattens the contents of FeatureSpecHelloworldFeatureTest from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestSlice(c *Client, i interface{}) []FeatureSpecHelloworldFeatureTest {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecHelloworldFeatureTest{}
	}

	if len(a) == 0 {
		return []FeatureSpecHelloworldFeatureTest{}
	}

	items := make([]FeatureSpecHelloworldFeatureTest, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecHelloworldFeatureTest(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureSpecHelloworldFeatureTest expands an instance of FeatureSpecHelloworldFeatureTest into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTest(c *Client, f *FeatureSpecHelloworldFeatureTest) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.First; !dcl.IsEmptyValueIndirect(v) {
		m["first"] = v
	}
	if v := f.Second; !dcl.IsEmptyValueIndirect(v) {
		m["second"] = v
	}
	if v := f.Third; !dcl.IsEmptyValueIndirect(v) {
		m["third"] = v
	}
	if v := f.Fourth; !dcl.IsEmptyValueIndirect(v) {
		m["fourth"] = v
	}
	if v, err := expandFeatureSpecHelloworldFeatureTestFifth(c, f.Fifth); err != nil {
		return nil, fmt.Errorf("error expanding Fifth into fifth: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fifth"] = v
	}
	if v := f.Sixth; !dcl.IsEmptyValueIndirect(v) {
		m["sixth"] = v
	}
	if v := f.Seventh; !dcl.IsEmptyValueIndirect(v) {
		m["seventh"] = v
	}
	if v, err := expandFeatureSpecHelloworldFeatureTestEighthSlice(c, f.Eighth); err != nil {
		return nil, fmt.Errorf("error expanding Eighth into eighth: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["eighth"] = v
	}
	if v := f.Ninth; !dcl.IsEmptyValueIndirect(v) {
		m["ninth"] = v
	}

	return m, nil
}

// flattenFeatureSpecHelloworldFeatureTest flattens an instance of FeatureSpecHelloworldFeatureTest from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTest(c *Client, i interface{}) *FeatureSpecHelloworldFeatureTest {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecHelloworldFeatureTest{}
	r.First = dcl.FlattenString(m["first"])
	r.Second = dcl.FlattenInteger(m["second"])
	r.Third = flattenFeatureSpecHelloworldFeatureTestThirdEnum(m["third"])
	r.Fourth = dcl.FlattenString(m["fourth"])
	r.Fifth = flattenFeatureSpecHelloworldFeatureTestFifth(c, m["fifth"])
	r.Sixth = dcl.FlattenInteger(m["sixth"])
	r.Seventh = dcl.FlattenString(m["seventh"])
	r.Eighth = flattenFeatureSpecHelloworldFeatureTestEighthSlice(c, m["eighth"])
	r.Ninth = dcl.FlattenKeyValuePairs(m["ninth"])

	return r
}

// expandFeatureSpecHelloworldFeatureTestFifthMap expands the contents of FeatureSpecHelloworldFeatureTestFifth into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestFifthMap(c *Client, f map[string]FeatureSpecHelloworldFeatureTestFifth) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecHelloworldFeatureTestFifth(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecHelloworldFeatureTestFifthSlice expands the contents of FeatureSpecHelloworldFeatureTestFifth into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestFifthSlice(c *Client, f []FeatureSpecHelloworldFeatureTestFifth) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecHelloworldFeatureTestFifth(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecHelloworldFeatureTestFifthMap flattens the contents of FeatureSpecHelloworldFeatureTestFifth from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestFifthMap(c *Client, i interface{}) map[string]FeatureSpecHelloworldFeatureTestFifth {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecHelloworldFeatureTestFifth{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecHelloworldFeatureTestFifth{}
	}

	items := make(map[string]FeatureSpecHelloworldFeatureTestFifth)
	for k, item := range a {
		items[k] = *flattenFeatureSpecHelloworldFeatureTestFifth(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureSpecHelloworldFeatureTestFifthSlice flattens the contents of FeatureSpecHelloworldFeatureTestFifth from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestFifthSlice(c *Client, i interface{}) []FeatureSpecHelloworldFeatureTestFifth {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecHelloworldFeatureTestFifth{}
	}

	if len(a) == 0 {
		return []FeatureSpecHelloworldFeatureTestFifth{}
	}

	items := make([]FeatureSpecHelloworldFeatureTestFifth, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecHelloworldFeatureTestFifth(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureSpecHelloworldFeatureTestFifth expands an instance of FeatureSpecHelloworldFeatureTestFifth into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestFifth(c *Client, f *FeatureSpecHelloworldFeatureTestFifth) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TypeUrl; !dcl.IsEmptyValueIndirect(v) {
		m["typeUrl"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenFeatureSpecHelloworldFeatureTestFifth flattens an instance of FeatureSpecHelloworldFeatureTestFifth from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestFifth(c *Client, i interface{}) *FeatureSpecHelloworldFeatureTestFifth {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecHelloworldFeatureTestFifth{}
	r.TypeUrl = dcl.FlattenString(m["typeUrl"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandFeatureSpecHelloworldFeatureTestEighthMap expands the contents of FeatureSpecHelloworldFeatureTestEighth into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestEighthMap(c *Client, f map[string]FeatureSpecHelloworldFeatureTestEighth) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecHelloworldFeatureTestEighth(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecHelloworldFeatureTestEighthSlice expands the contents of FeatureSpecHelloworldFeatureTestEighth into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestEighthSlice(c *Client, f []FeatureSpecHelloworldFeatureTestEighth) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecHelloworldFeatureTestEighth(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecHelloworldFeatureTestEighthMap flattens the contents of FeatureSpecHelloworldFeatureTestEighth from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestEighthMap(c *Client, i interface{}) map[string]FeatureSpecHelloworldFeatureTestEighth {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecHelloworldFeatureTestEighth{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecHelloworldFeatureTestEighth{}
	}

	items := make(map[string]FeatureSpecHelloworldFeatureTestEighth)
	for k, item := range a {
		items[k] = *flattenFeatureSpecHelloworldFeatureTestEighth(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureSpecHelloworldFeatureTestEighthSlice flattens the contents of FeatureSpecHelloworldFeatureTestEighth from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestEighthSlice(c *Client, i interface{}) []FeatureSpecHelloworldFeatureTestEighth {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecHelloworldFeatureTestEighth{}
	}

	if len(a) == 0 {
		return []FeatureSpecHelloworldFeatureTestEighth{}
	}

	items := make([]FeatureSpecHelloworldFeatureTestEighth, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecHelloworldFeatureTestEighth(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureSpecHelloworldFeatureTestEighth expands an instance of FeatureSpecHelloworldFeatureTestEighth into a JSON
// request object.
func expandFeatureSpecHelloworldFeatureTestEighth(c *Client, f *FeatureSpecHelloworldFeatureTestEighth) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.First; !dcl.IsEmptyValueIndirect(v) {
		m["first"] = v
	}
	if v := f.Second; !dcl.IsEmptyValueIndirect(v) {
		m["second"] = v
	}

	return m, nil
}

// flattenFeatureSpecHelloworldFeatureTestEighth flattens an instance of FeatureSpecHelloworldFeatureTestEighth from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestEighth(c *Client, i interface{}) *FeatureSpecHelloworldFeatureTestEighth {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecHelloworldFeatureTestEighth{}
	r.First = dcl.FlattenString(m["first"])
	r.Second = dcl.FlattenInteger(m["second"])

	return r
}

// expandFeatureStateMap expands the contents of FeatureState into a JSON
// request object.
func expandFeatureStateMap(c *Client, f map[string]FeatureState) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureState(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateSlice expands the contents of FeatureState into a JSON
// request object.
func expandFeatureStateSlice(c *Client, f []FeatureState) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureState(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateMap flattens the contents of FeatureState from a JSON
// response object.
func flattenFeatureStateMap(c *Client, i interface{}) map[string]FeatureState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureState{}
	}

	if len(a) == 0 {
		return map[string]FeatureState{}
	}

	items := make(map[string]FeatureState)
	for k, item := range a {
		items[k] = *flattenFeatureState(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureStateSlice flattens the contents of FeatureState from a JSON
// response object.
func flattenFeatureStateSlice(c *Client, i interface{}) []FeatureState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureState{}
	}

	if len(a) == 0 {
		return []FeatureState{}
	}

	items := make([]FeatureState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureState(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureState expands an instance of FeatureState into a JSON
// request object.
func expandFeatureState(c *Client, f *FeatureState) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureStateState(c, f.State); err != nil {
		return nil, fmt.Errorf("error expanding State into state: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandFeatureStateHelloworld(c, f.Helloworld); err != nil {
		return nil, fmt.Errorf("error expanding Helloworld into helloworld: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["helloworld"] = v
	}

	return m, nil
}

// flattenFeatureState flattens an instance of FeatureState from a JSON
// response object.
func flattenFeatureState(c *Client, i interface{}) *FeatureState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureState{}
	r.State = flattenFeatureStateState(c, m["state"])
	r.Helloworld = flattenFeatureStateHelloworld(c, m["helloworld"])

	return r
}

// expandFeatureStateStateMap expands the contents of FeatureStateState into a JSON
// request object.
func expandFeatureStateStateMap(c *Client, f map[string]FeatureStateState) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateState(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateStateSlice expands the contents of FeatureStateState into a JSON
// request object.
func expandFeatureStateStateSlice(c *Client, f []FeatureStateState) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateState(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateStateMap flattens the contents of FeatureStateState from a JSON
// response object.
func flattenFeatureStateStateMap(c *Client, i interface{}) map[string]FeatureStateState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateState{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateState{}
	}

	items := make(map[string]FeatureStateState)
	for k, item := range a {
		items[k] = *flattenFeatureStateState(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureStateStateSlice flattens the contents of FeatureStateState from a JSON
// response object.
func flattenFeatureStateStateSlice(c *Client, i interface{}) []FeatureStateState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateState{}
	}

	if len(a) == 0 {
		return []FeatureStateState{}
	}

	items := make([]FeatureStateState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateState(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureStateState expands an instance of FeatureStateState into a JSON
// request object.
func expandFeatureStateState(c *Client, f *FeatureStateState) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}

	return m, nil
}

// flattenFeatureStateState flattens an instance of FeatureStateState from a JSON
// response object.
func flattenFeatureStateState(c *Client, i interface{}) *FeatureStateState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateState{}
	r.Code = flattenFeatureStateStateCodeEnum(m["code"])
	r.Description = dcl.FlattenString(m["description"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// expandFeatureStateHelloworldMap expands the contents of FeatureStateHelloworld into a JSON
// request object.
func expandFeatureStateHelloworldMap(c *Client, f map[string]FeatureStateHelloworld) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateHelloworld(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateHelloworldSlice expands the contents of FeatureStateHelloworld into a JSON
// request object.
func expandFeatureStateHelloworldSlice(c *Client, f []FeatureStateHelloworld) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateHelloworld(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateHelloworldMap flattens the contents of FeatureStateHelloworld from a JSON
// response object.
func flattenFeatureStateHelloworldMap(c *Client, i interface{}) map[string]FeatureStateHelloworld {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateHelloworld{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateHelloworld{}
	}

	items := make(map[string]FeatureStateHelloworld)
	for k, item := range a {
		items[k] = *flattenFeatureStateHelloworld(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFeatureStateHelloworldSlice flattens the contents of FeatureStateHelloworld from a JSON
// response object.
func flattenFeatureStateHelloworldSlice(c *Client, i interface{}) []FeatureStateHelloworld {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateHelloworld{}
	}

	if len(a) == 0 {
		return []FeatureStateHelloworld{}
	}

	items := make([]FeatureStateHelloworld, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateHelloworld(c, item.(map[string]interface{})))
	}

	return items
}

// expandFeatureStateHelloworld expands an instance of FeatureStateHelloworld into a JSON
// request object.
func expandFeatureStateHelloworld(c *Client, f *FeatureStateHelloworld) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateHelloworld flattens an instance of FeatureStateHelloworld from a JSON
// response object.
func flattenFeatureStateHelloworld(c *Client, i interface{}) *FeatureStateHelloworld {

	r := &FeatureStateHelloworld{}

	return r
}

// flattenFeatureResourceStateStateEnumSlice flattens the contents of FeatureResourceStateStateEnum from a JSON
// response object.
func flattenFeatureResourceStateStateEnumSlice(c *Client, i interface{}) []FeatureResourceStateStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureResourceStateStateEnum{}
	}

	if len(a) == 0 {
		return []FeatureResourceStateStateEnum{}
	}

	items := make([]FeatureResourceStateStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureResourceStateStateEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureResourceStateStateEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureResourceStateStateEnum with the same value as that string.
func flattenFeatureResourceStateStateEnum(i interface{}) *FeatureResourceStateStateEnum {
	s, ok := i.(string)
	if !ok {
		return FeatureResourceStateStateEnumRef("")
	}

	return FeatureResourceStateStateEnumRef(s)
}

// flattenFeatureSpecMulticlusteringressBillingEnumSlice flattens the contents of FeatureSpecMulticlusteringressBillingEnum from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressBillingEnumSlice(c *Client, i interface{}) []FeatureSpecMulticlusteringressBillingEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecMulticlusteringressBillingEnum{}
	}

	if len(a) == 0 {
		return []FeatureSpecMulticlusteringressBillingEnum{}
	}

	items := make([]FeatureSpecMulticlusteringressBillingEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecMulticlusteringressBillingEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureSpecMulticlusteringressBillingEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureSpecMulticlusteringressBillingEnum with the same value as that string.
func flattenFeatureSpecMulticlusteringressBillingEnum(i interface{}) *FeatureSpecMulticlusteringressBillingEnum {
	s, ok := i.(string)
	if !ok {
		return FeatureSpecMulticlusteringressBillingEnumRef("")
	}

	return FeatureSpecMulticlusteringressBillingEnumRef(s)
}

// flattenFeatureSpecHelloworldFeatureTestThirdEnumSlice flattens the contents of FeatureSpecHelloworldFeatureTestThirdEnum from a JSON
// response object.
func flattenFeatureSpecHelloworldFeatureTestThirdEnumSlice(c *Client, i interface{}) []FeatureSpecHelloworldFeatureTestThirdEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecHelloworldFeatureTestThirdEnum{}
	}

	if len(a) == 0 {
		return []FeatureSpecHelloworldFeatureTestThirdEnum{}
	}

	items := make([]FeatureSpecHelloworldFeatureTestThirdEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecHelloworldFeatureTestThirdEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureSpecHelloworldFeatureTestThirdEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureSpecHelloworldFeatureTestThirdEnum with the same value as that string.
func flattenFeatureSpecHelloworldFeatureTestThirdEnum(i interface{}) *FeatureSpecHelloworldFeatureTestThirdEnum {
	s, ok := i.(string)
	if !ok {
		return FeatureSpecHelloworldFeatureTestThirdEnumRef("")
	}

	return FeatureSpecHelloworldFeatureTestThirdEnumRef(s)
}

// flattenFeatureStateStateCodeEnumSlice flattens the contents of FeatureStateStateCodeEnum from a JSON
// response object.
func flattenFeatureStateStateCodeEnumSlice(c *Client, i interface{}) []FeatureStateStateCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateStateCodeEnum{}
	}

	if len(a) == 0 {
		return []FeatureStateStateCodeEnum{}
	}

	items := make([]FeatureStateStateCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateStateCodeEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureStateStateCodeEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureStateStateCodeEnum with the same value as that string.
func flattenFeatureStateStateCodeEnum(i interface{}) *FeatureStateStateCodeEnum {
	s, ok := i.(string)
	if !ok {
		return FeatureStateStateCodeEnumRef("")
	}

	return FeatureStateStateCodeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Feature) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFeature(b, c)
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
