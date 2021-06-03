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
	return nil
}
func (r *FeatureSpecMulticlusteringress) validate() error {
	return nil
}
func (r *FeatureState) validate() error {
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureStateState) validate() error {
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
	Diffs        []*dcl.FieldDiff
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

func (c *Client) featureDiffsForRawDesired(ctx context.Context, rawDesired *Feature, opts ...dcl.ApplyOption) (initial, desired *Feature, diffs []*dcl.FieldDiff, err error) {
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

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewFeatureState(c, rawDesired.State, rawNew.State)
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
	if dcl.BoolCanonicalize(des.HasResources, initial.HasResources) || dcl.IsZeroValue(des.HasResources) {
		des.HasResources = initial.HasResources
	}

	return des
}

func canonicalizeNewFeatureResourceState(c *Client, des, nw *FeatureResourceState) *FeatureResourceState {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.State) {
		nw.State = des.State
	}
	if dcl.BoolCanonicalize(des.HasResources, nw.HasResources) {
		nw.HasResources = des.HasResources
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
			if diffs, _ := compareFeatureResourceStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	return des
}

func canonicalizeNewFeatureSpec(c *Client, des, nw *FeatureSpec) *FeatureSpec {
	if des == nil || nw == nil {
		return nw
	}

	nw.Multiclusteringress = canonicalizeNewFeatureSpecMulticlusteringress(c, des.Multiclusteringress, nw.Multiclusteringress)

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
			if diffs, _ := compareFeatureSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	if dcl.NameToSelfLink(des.ConfigMembership, initial.ConfigMembership) || dcl.IsZeroValue(des.ConfigMembership) {
		des.ConfigMembership = initial.ConfigMembership
	}

	return des
}

func canonicalizeNewFeatureSpecMulticlusteringress(c *Client, des, nw *FeatureSpecMulticlusteringress) *FeatureSpecMulticlusteringress {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.ConfigMembership, nw.ConfigMembership) {
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
			if diffs, _ := compareFeatureSpecMulticlusteringressNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []FeatureSpecMulticlusteringress
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecMulticlusteringress(c, &d, &n))
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

	return des
}

func canonicalizeNewFeatureState(c *Client, des, nw *FeatureState) *FeatureState {
	if des == nil || nw == nil {
		return nw
	}

	nw.State = canonicalizeNewFeatureStateState(c, des.State, nw.State)

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
			if diffs, _ := compareFeatureStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
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

	return des
}

func canonicalizeNewFeatureStateState(c *Client, des, nw *FeatureStateState) *FeatureStateState {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Code) {
		nw.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.UpdateTime, nw.UpdateTime) {
		nw.UpdateTime = des.UpdateTime
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
			if diffs, _ := compareFeatureStateStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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
		return nw
	}

	var items []FeatureStateState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateState(c, &d, &n))
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
func diffFeature(c *Client, desired, actual *Feature, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceState, actual.ResourceState, dcl.Info{ObjectFunction: compareFeatureResourceStateNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.Info{ObjectFunction: compareFeatureSpecNewStyle, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, ObjectFunction: compareFeatureStateNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareFeatureResourceStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureResourceState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureResourceState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureResourceState or *FeatureResourceState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureResourceState)
	if !ok {
		actualNotPointer, ok := a.(FeatureResourceState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureResourceState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HasResources, actual.HasResources, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HasResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpec)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpec or *FeatureSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpec)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Multiclusteringress, actual.Multiclusteringress, dcl.Info{ObjectFunction: compareFeatureSpecMulticlusteringressNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Multiclusteringress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecMulticlusteringressNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecMulticlusteringress)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecMulticlusteringress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecMulticlusteringress or *FeatureSpecMulticlusteringress", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecMulticlusteringress)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecMulticlusteringress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecMulticlusteringress", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigMembership, actual.ConfigMembership, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigMembership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureState or *FeatureState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureState)
	if !ok {
		actualNotPointer, ok := a.(FeatureState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, ObjectFunction: compareFeatureStateStateNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateState or *FeatureStateState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateState)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Feature) urlNormalized() *Feature {
	normalized := dcl.Copy(*r).(Feature)
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
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandFeatureResourceState(c, f.ResourceState); err != nil {
		return nil, fmt.Errorf("error expanding ResourceState into resourceState: %w", err)
	} else if v != nil {
		m["resourceState"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if v != nil {
		m["spec"] = v
	}
	if v, err := expandFeatureState(c, f.State); err != nil {
		return nil, fmt.Errorf("error expanding State into state: %w", err)
	} else if v != nil {
		m["state"] = v
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
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
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

	res := &Feature{}
	res.Name = dcl.FlattenString(m["name"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.ResourceState = flattenFeatureResourceState(c, m["resourceState"])
	res.Spec = flattenFeatureSpec(c, m["spec"])
	res.State = flattenFeatureState(c, m["state"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
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
	if v := f.HasResources; !dcl.IsEmptyValueIndirect(v) {
		m["hasResources"] = v
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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureResourceState
	}
	r.State = flattenFeatureResourceStateStateEnum(m["state"])
	r.HasResources = dcl.FlattenBool(m["hasResources"])

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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpec
	}
	r.Multiclusteringress = flattenFeatureSpecMulticlusteringress(c, m["multiclusteringress"])

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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecMulticlusteringress
	}
	r.ConfigMembership = dcl.FlattenString(m["configMembership"])

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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureState
	}
	r.State = flattenFeatureStateState(c, m["state"])

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

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateState
	}
	r.Code = flattenFeatureStateStateCodeEnum(m["code"])
	r.Description = dcl.FlattenString(m["description"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

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

type featureDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         featureApiOperation
}

func convertFieldDiffToFeatureOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]featureDiff, error) {
	var diffs []featureDiff
	for _, op := range ops {
		diff := featureDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTofeatureApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTofeatureApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (featureApiOperation, error) {
	switch op {

	case "updateFeatureUpdateFeatureOperation":
		return &updateFeatureUpdateFeatureOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
